package packet

import (
	"fmt"
	"github.com/vx-labs/mqtt-protocol/types"
	"io"
	"bytes"
)

func NewDecoder() *Decoder {
	d := &Decoder{
	}
	return d
}

type Decoder struct {
}

func (d *Decoder) Decode(r io.Reader) (types.Packet, error) {
	h := Header{}
	err := h.Decode(r)
	if err != nil {
		return nil, err
	}
	b := make([]byte, h.remLength)
	read := 0
	for read < len(b) {
		n, err := r.Read(b[read:])
		if err != nil {
			return nil, err
		}
		read += n
	}
	reader := bytes.NewReader(b)
	var c types.Packet
	switch h.PacketType() {
	case CONNECT:
		c = NewConnect().WithFixedHeader(h)
	case SUBSCRIBE:
		c = NewSubscribe().WithFixedHeader(h)
	case PUBLISH:
		c = NewPublish().WithFixedHeader(h)
	case DISCONNECT:
		c = NewDisconnect().WithFixedHeader(h)
	case PINGREQ:
		c = NewPingreq().WithFixedHeader(h)
	case PINGRESP:
		c = NewPingresp().WithFixedHeader(h)
	case PUBACK:
		c = NewPuback().WithFixedHeader(h)
	case UNSUBSCRIBE:
		c = NewUnsubscribe().WithFixedHeader(h)
	default:
		return nil, fmt.Errorf("unsupported packet type by decoder: %d", h.PacketType())
	}
	err = c.Decode(reader)
	if err != nil {
		return nil, err
	}
	return c, nil
}
