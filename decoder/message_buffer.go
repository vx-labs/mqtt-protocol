package decoder

import (
	"encoding/binary"
	"errors"
	fmt "fmt"
	"io"

	"github.com/vx-labs/mqtt-protocol/packet"
)

func readMessageBuffer(p *packet.Header, buf []byte, r io.Reader) (byte, int, error) {
	if len(buf) <= 4 {
		return 0, 0, errors.New("buffer too short for header")
	}
	read := 0
	n, err := io.ReadFull(r, buf[0:1])
	read += n
	if err != nil {
		return 0, read, err
	}
	fixedHeader := buf[0]
	p.Retain = fixedHeader&0x1 == 0x1
	p.Qos = int32(fixedHeader & (0x3 << 1) >> 1)
	p.Dup = fixedHeader&(0x1<<3)>>3 == 1
	packetType := fixedHeader & (0xf << 4) >> 4
	read = 0
	for {
		if read > 4 {
			return 0, read, fmt.Errorf("malformed remlength")
		}
		cur := read
		n, err := io.ReadFull(r, buf[read:read+1])
		if n == 0 {
			return 0, read, fmt.Errorf("null read")
		}
		read++
		if err != nil {
			return 0, read, err
		}
		if buf[cur] < 0x80 {
			break
		}
	}
	remlen, _ := binary.Uvarint(buf[:read])
	bufferLen := int(remlen)
	if len(buf) < bufferLen {
		return 0, read, errors.New("buffer too short for payload")
	}
	n, err = io.ReadFull(r, buf[:bufferLen])
	return packetType, bufferLen, nil
}
