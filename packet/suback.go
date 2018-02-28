package packet

import (
	"github.com/vx-labs/mqtt-protocol/types"
	"io"
)

type Suback struct {
	header     Header
	messageId  types.MqttInt
	grantedQos types.MqttByteArray
}

func (p *Suback) clone() Suback {
	return *p
}
func (p *Suback) GrantedQoS() []byte {
	return p.grantedQos.Bytes()
}
func (p *Suback) WithGrantedQoS(q []byte) *Suback {
	c := p.clone()
	c.grantedQos.FromBytes(q)
	return &c
}
func (p *Suback) WithFixedHeader(h Header) *Suback {
	c := *p
	c.header = h
	return &c
}

func (p *Suback) Type() byte {
	return SUBACK
}
func (p *Suback) String() string {
	return "suback"
}
func (p *Suback) Length() int {
	return 2 + p.grantedQos.Length()
}

func (p *Suback) FixedHeader(l int) types.Field {
	return p.header.WithRemainingLength(l)
}

func (p *Suback) MessageId() int {
	return p.messageId.Int()
}

func (p *Suback) WithMessageId(i int) *Suback {
	c := p.clone()
	c.messageId.FromInt(i)
	return &c
}

func (p *Suback) VariableHeaders() []types.Field {
	return []types.Field{
		&p.messageId,
	}
}
func (p *Suback) Payload() []types.Field {
	b := make([]types.Field, 1)
	b[0] = &p.grantedQos
	return b
}

func (p *Suback) Decode(r io.Reader) error {
	for _, f := range p.VariableHeaders() {
		err := f.Decode(r)
		if err != nil {
			return err
		}
	}
	for _, f := range p.Payload() {
		err := f.Decode(r)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Suback) Packet() types.Packet {
	return c
}
