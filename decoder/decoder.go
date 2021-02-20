package decoder

import (
	"io"
	"sync"

	"github.com/vx-labs/mqtt-protocol/packet"
)

func Decode(r io.Reader, buf []byte) (packet.Packet, error) {
	p, _, err := decodeEncodedPacket(buf, r)
	return p, err
}

type Sync struct {
	headerBuf *sync.Pool
}

func New(maxPacketSize int) *Sync {
	pool := &sync.Pool{
		New: func() interface{} {
			return make([]byte, maxPacketSize)
		},
	}
	return &Sync{
		headerBuf: pool,
	}
}

func (s *Sync) Decode(r io.Reader) (packet.Packet, error) {
	buf := s.headerBuf.Get().([]byte)
	defer s.headerBuf.Put(buf)
	return Decode(r, buf)
}
