package packet

import (
	"github.com/vx-labs/mqtt-protocol/types"
	"io"
)

const (
	CONNACK_CONNECTION_ACCEPTED              byte = iota
	CONNACK_REFUSED_BAD_PROTOCOL_VERSION
	CONNACK_REFUSED_IDENTIFIER_REJECTED
	CONNACK_REFUSED_SERVER_UNAVAILABLE
	CONNACK_REFUSED_BAD_USERNAME_OR_PASSWORD
	CONNACK_REFUSED_NOT_AUTHORIZED
)

type Connack struct {
	header     Header
	reserved   types.MqttByte
	returnCode types.MqttByte
}

func (p *Connack) WithFixedHeader(h Header) *Connack{
	c := p.clone()
	c.header = h
	return &c
}

func (h *Connack) clone() Connack  {
	return *h
}
func (h *Connack) ReturnCode() byte  {
	return h.returnCode.Byte()
}
func (h *Connack) WithReturnCode(b byte) *Connack  {
	c := h.clone()
	c.returnCode.FromByte(b)
	return &c
}
func (h *Connack) Type() byte {
	return CONNACK
}
func (h *Connack) String() string {
	return "connack"
}
func (h *Connack) FixedHeader(l int) types.Field {
	return h.header.WithRemainingLength(l)
}

func (c *Connack) VariableHeaders() ([]types.Field) {
	return []types.Field{}
}
func (c *Connack) Payload() ([]types.Field) {
	return []types.Field{
		&c.reserved,
		&c.returnCode,
	}
}

func (h *Connack) Decode(r io.Reader) error {
	for _, f := range h.VariableHeaders() {
		err := f.Decode(r)
		if err != nil {
			return err
		}
	}
	for _, f := range h.Payload() {
		err := f.Decode(r)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Connack) Packet() types.Packet {
	return c
}
