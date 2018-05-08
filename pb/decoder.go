package pb

import (
	"encoding/binary"
	fmt "fmt"
	"io"

	"github.com/vx-labs/mqtt-protocol/types"
)

const (
	Max_Packet_Size = 8092
)

func NewDecoder() *Decoder {
	d := &Decoder{
		buffer: make([]byte, Max_Packet_Size),
	}
	return d
}

type Decoder struct {
	buffer []byte
}

func (d *Decoder) readMessageBuffer(p *MqttHeader, r io.Reader) (byte, []byte, error) {
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
	if remlen > Max_Packet_Size {
		return 0, nil, fmt.Errorf("packet size (%d) greater than max packet size (%d)", remlen, Max_Packet_Size)
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

func (d *Decoder) Decode(r io.Reader) (types.Packet, error) {
	p := &MqttHeader{}

	read := 0
	for read != 2 {
		n, err := r.Read(d.buffer[read:2])
		read += n
		if err != nil {
			return nil, err
		}
	}
	d.readMessageBuffer(p, r)
	var c types.Packet

	return c, nil
}
