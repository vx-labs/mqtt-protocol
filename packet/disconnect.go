package packet

import (
	"github.com/vx-labs/mqtt-protocol/types"
	"io"
)

type Disconnect struct {
	header     Header
}

func (p *Disconnect) WithFixedHeader(h Header) *Disconnect{
	c := p.clone()
	c.header = h
	return &c
}

func NewDisconnect() *Disconnect {
	return &Disconnect{}
}


func (h *Disconnect) clone() Disconnect  {
	return *h
}
func (h *Disconnect) Type() byte {
	return DISCONNECT
}
func (h *Disconnect) String() string {
	return "disconnect"
}
func (h *Disconnect) FixedHeader(l int) types.Field {
	return h.header.WithRemainingLength(l)
}

func (c *Disconnect) VariableHeaders() ([]types.Field) {
	return []types.Field{}
}
func (c *Disconnect) Payload() ([]types.Field) {
	return []types.Field{}
}

func (h *Disconnect) Decode(r io.Reader) error {
	return nil
}

func (c *Disconnect) Packet() types.Packet {
	return c
}
