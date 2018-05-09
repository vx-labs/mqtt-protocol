package decoder

import (
	"encoding/binary"
	fmt "fmt"
	"io"

	"github.com/vx-labs/mqtt-protocol/pb"
)

func readBits(r io.Reader, buff []byte) (int, error) {
	read := 0
	for read < len(buff) {
		n, err := r.Read(buff[read:])
		read += n
		if err != nil {
			return read, err
		}
	}
	return read, nil
}
func (d *Decoder) readMessageBuffer(p *pb.MqttHeader, r io.Reader, buffer []byte) (byte, []byte, error) {
	read := 0
	n, err := readBits(r, buffer[read:1])
	read += n
	if err != nil {
		return 0, nil, err
	}
	fixedHeader := buffer[0]
	p.Retain = fixedHeader&0x1 == 0x1
	p.Qos = int32(fixedHeader & (0x3 << 1) >> 1)
	p.Dup = fixedHeader&(0x1<<3)>>3 == 1
	packetType := fixedHeader & (0xf << 4) >> 4
	for {
		if read > 4 {
			return 0, nil, fmt.Errorf("malformed remlength")
		}
		cur := read
		n, err := readBits(r, buffer[read:read+1])
		if n == 0 {
			return 0, nil, fmt.Errorf("null read")
		}
		read++
		if err != nil {
			return 0, nil, err
		}
		if buffer[cur] < 0x80 {
			break
		}
	}
	remlen, _ := binary.Uvarint(buffer[1:read])
	if int(remlen) > len(buffer) {
		return 0, nil, fmt.Errorf("packet size (%d) greater than max packet size (%d)", remlen, len(buffer))
	}
	read = 0
	n, err = readBits(r, buffer[0:remlen])
	read += n
	return packetType, buffer[:read], nil
}
