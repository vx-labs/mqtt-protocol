package packet

import (
	"github.com/vx-labs/mqtt-protocol/types"
	"io"
)

type Unsuback struct {
	header    Header
	messageId types.MqttInt
}

func (p *Unsuback) WithFixedHeader(h Header) *Unsuback {
	c := p.clone()
	c.header = h
	return &c
}

func (p *Unsuback) clone() Unsuback {
	return *p
}
func (p *Unsuback) Type() byte {
	return UNSUBACK
}
func (p *Unsuback) String() string {
	return "unbsuback"
}
func (p *Unsuback) FixedHeader(l int) types.Field {
	return p.header.WithRemainingLength(l)
}

func (p *Unsuback) VariableHeaders() ([]types.Field) {
	return []types.Field{
		&p.messageId,
	}
}

func (p *Unsuback) WithMessageId(i int) *Unsuback {
	c := p.clone()
	c.messageId.FromInt(i)
	return &c
}
func (p *Unsuback) MessageId() int {
	return p.messageId.Int()
}

func (p *Unsuback) Payload() ([]types.Field) {
	return []types.Field{}
}

func (p *Unsuback) Decode(r io.Reader) error {
	for _, f := range p.VariableHeaders() {
		err := f.Decode(r)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *Unsuback) Packet() types.Packet {
	return p
}
