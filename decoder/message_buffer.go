package decoder

import (
	"encoding/binary"
	fmt "fmt"
	"io"

	"github.com/vx-labs/mqtt-protocol/pb"
)

func (d *Decoder) readMessageBuffer(p *pb.MqttHeader, r io.Reader) (byte, []byte, error) {
	read := 0
	for read < 1 {
		n, err := r.Read(d.buffer[read:1])
		read += n
		if err != nil {
			return 0, nil, err
		}
	}
	fixedHeader := d.buffer[0]
	p.Retain = fixedHeader&0x1 == 0x1
	p.Qos = int32(fixedHeader & (0x3 << 1) >> 1)
	p.Dup = fixedHeader&(0x1<<3)>>3 == 1
	packetType := fixedHeader & (0xf << 4) >> 4
	for {
		if read > 4 {
			return 0, nil, fmt.Errorf("malformed remlength")
		}
		cur := read
		n, err := r.Read(d.buffer[read : read+1])
		if n == 0 {
			return 0, nil, fmt.Errorf("null read")
		}
		read++
		if err != nil {
			return 0, nil, err
		}
		if d.buffer[cur] < 0x80 {
			break
		}
	}
	remlen, _ := binary.Uvarint(d.buffer[1:read])
	if int(remlen) > len(d.buffer) {
		return 0, nil, fmt.Errorf("packet size (%d) greater than max packet size (%d)", remlen, len(d.buffer))
	}
	read = 0
	for read < int(remlen) {
		n, err := r.Read(d.buffer[0:])
		if err != nil {
			return 0, nil, err
		}
		read += n
	}
	return packetType, d.buffer[:read], nil
}
