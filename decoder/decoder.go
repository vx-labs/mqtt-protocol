package decoder

import (
	"io"

	"github.com/vx-labs/mqtt-protocol/packet"
)

func Decode(r io.Reader, headerBuf []byte) (packet.Packet, error) {
	p, _, err := decodeEncodedPacket(headerBuf, r)
	return p, err
}
