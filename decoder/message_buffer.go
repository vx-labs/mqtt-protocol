package decoder

import (
	"encoding/binary"
	fmt "fmt"
	"io"

	"github.com/vx-labs/mqtt-protocol/packet"
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
func readMessageBuffer(p *packet.Header, r io.Reader) (byte, []byte, int, error) {
	sizeBuff := make([]byte, 4)
	read := 0
	n, err := readBits(r, sizeBuff[0:1])
	read += n
	if err != nil {
		return 0, nil, read, err
	}
	fixedHeader := sizeBuff[0]
	p.Retain = fixedHeader&0x1 == 0x1
	p.Qos = int32(fixedHeader & (0x3 << 1) >> 1)
	p.Dup = fixedHeader&(0x1<<3)>>3 == 1
	packetType := fixedHeader & (0xf << 4) >> 4

	read = 0
	for {
		if read > 4 {
			return 0, nil, read, fmt.Errorf("malformed remlength")
		}
		cur := read
		n, err := readBits(r, sizeBuff[read:read+1])
		if n == 0 {
			return 0, nil, read, fmt.Errorf("null read")
		}
		read++
		if err != nil {
			return 0, nil, read, err
		}
		if sizeBuff[cur] < 0x80 {
			break
		}
	}
	remlen, _ := binary.Uvarint(sizeBuff)
	buffer := make([]byte, remlen)
	n, err = readBits(r, buffer)
	return packetType, buffer, read, nil
}
