package decoder

import (
	"encoding/binary"
	"errors"
	fmt "fmt"
	"io"

	"github.com/vx-labs/mqtt-protocol/packet"
)

func readMessageBuffer(p *packet.Header, sizeBuf []byte, r io.Reader) (byte, []byte, int, error) {
	if len(sizeBuf) != 4 {
		return 0, nil, 0, errors.New("invalid header buffer size")
	}
	read := 0
	n, err := io.ReadFull(r, sizeBuf[0:1])
	read += n
	if err != nil {
		return 0, nil, read, err
	}
	fixedHeader := sizeBuf[0]
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
		n, err := io.ReadFull(r, sizeBuf[read:read+1])
		if n == 0 {
			return 0, nil, read, fmt.Errorf("null read")
		}
		read++
		if err != nil {
			return 0, nil, read, err
		}
		if sizeBuf[cur] < 0x80 {
			break
		}
	}
	remlen, _ := binary.Uvarint(sizeBuf[:read])
	buffer := make([]byte, remlen)
	n, err = io.ReadFull(r, buffer)
	return packetType, buffer, read, nil
}
