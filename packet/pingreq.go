package packet

import (
	"github.com/vx-labs/mqtt-protocol/types"
	"io"
)

type Pingreq struct {
	header     Header
}

func (p *Pingreq) WithFixedHeader(h Header) *Pingreq{
	c := p.clone()
	c.header = h
	return &c
}

func NewPingreq() *Pingreq {
	return &Pingreq{}
}


func (h *Pingreq) clone() Pingreq  {
	return *h
}
func (h *Pingreq) Type() byte {
	return PINGREQ
}
func (h *Pingreq) String() string {
	return "pingreq"
}
func (h *Pingreq) FixedHeader(l int) types.Field {
	return h.header.WithRemainingLength(l)
}

func (c *Pingreq) VariableHeaders() ([]types.Field) {
	return []types.Field{}
}
func (c *Pingreq) Payload() ([]types.Field) {
	return []types.Field{}
}

func (h *Pingreq) Decode(r io.Reader) error {
	return nil
}

func (c *Pingreq) Packet() types.Packet {
	return c
}
