package packet

import (
	"github.com/vx-labs/mqtt-protocol/types"
	"io"
)

type Pingresp struct {
	header Header
}

func (p *Pingresp) WithFixedHeader(h Header) *Pingresp {
	c := p.clone()
	c.header = h
	return &c
}

func NewPingresp() *Pingresp {
	return &Pingresp{}
}

func (h *Pingresp) clone() Pingresp {
	return *h
}
func (h *Pingresp) Type() byte {
	return PINGRESP
}
func (h *Pingresp) String() string {
	return "pingresp"
}
func (h *Pingresp) FixedHeader(l int) types.Field {
	return h.header.WithRemainingLength(l)
}

func (c *Pingresp) VariableHeaders() ([]types.Field) {
	return []types.Field{}
}
func (c *Pingresp) Payload() ([]types.Field) {
	return []types.Field{}
}

func (h *Pingresp) Decode(r io.Reader) error {
	return nil
}

func (c *Pingresp) Packet() types.Packet {
	return c
}
