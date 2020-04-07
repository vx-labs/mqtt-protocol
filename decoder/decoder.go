package decoder

import (
	"io"

	"github.com/vx-labs/mqtt-protocol/packet"
)

type Decoder struct {
	headerBuf []byte
}

func New() *Decoder {
	return &Decoder{
		headerBuf: make([]byte, 4),
	}
}
func (a *Decoder) Decode(r io.Reader) (packet.Packet, error) {
	p, _, err := decodeEncodedPacket(a.headerBuf, r)
	return p, err
}
