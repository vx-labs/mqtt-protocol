package decoder

import (
	"io"

	"github.com/vx-labs/mqtt-protocol/packet"
)

func Decode(r io.Reader, headerBuf []byte, block bool) (packet.Packet, error) {
	p, _, err := decodeEncodedPacket(headerBuf, r, block)
	return p, err
}

type Sync struct {
	block     bool
	headerBuf []byte
}

func New(block bool) *Sync {
	return &Sync{
		headerBuf: make([]byte, 4),
		block:     block,
	}
}

func (s *Sync) Decode(r io.Reader) (packet.Packet, error) { return Decode(r, s.headerBuf, s.block) }
