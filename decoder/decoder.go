package decoder

import (
	"io"

	"github.com/vx-labs/mqtt-protocol/packet"
)

func Decode(r io.Reader, headerBuf []byte) (packet.Packet, error) {
	p, _, err := decodeEncodedPacket(headerBuf, r)
	return p, err
}

type Sync struct {
	headerBuf []byte
}

func New() *Sync {
	return &Sync{
		headerBuf: make([]byte, 4),
	}
}

func (s *Sync) Decode(r io.Reader) (packet.Packet, error) { return Decode(r, s.headerBuf) }
