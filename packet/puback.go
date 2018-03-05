package packet

import (
	"github.com/vx-labs/mqtt-protocol/types"
	"io"
)

type Puback struct {
	header    Header
	messageId types.MqttInt
}

func (p *Puback) WithFixedHeader(h Header) *Puback {
	c := p.clone()
	c.header = h
	return &c
}
func NewPuback() *Puback {
	return &Puback{}
}

func (p *Puback) clone() Puback {
	return *p
}
func (p *Puback) Type() byte {
	return PUBACK
}
func (p *Puback) String() string {
	return "puback"
}
func (p *Puback) FixedHeader(l int) types.Field {
	return p.header.WithRemainingLength(l)
}

func (p *Puback) VariableHeaders() ([]types.Field) {
	return []types.Field{
		&p.messageId,
	}
}

func (p *Puback) WithMessageId(i int) *Puback {
	c := p.clone()
	c.messageId.FromInt(i)
	return &c
}
func (p *Puback) MessageId() int {
	return p.messageId.Int()
}

func (p *Puback) Payload() ([]types.Field) {
	return []types.Field{}
}

func (p *Puback) Decode(r io.Reader) error {
	for _, f := range p.VariableHeaders() {
		err := f.Decode(r)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *Puback) Packet() types.Packet {
	return p
}
