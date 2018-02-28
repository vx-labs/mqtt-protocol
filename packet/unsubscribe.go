package packet

import (
	"github.com/vx-labs/mqtt-protocol/types"
	"io"
)

type Unsubscribe struct {
	header    Header
	messageId types.MqttInt
	topics    []types.MqttString
}

func NewUnsubscribe() *Unsubscribe {
	return &Unsubscribe{}
}

func (p *Unsubscribe) WithFixedHeader(h Header) *Unsubscribe {
	c := *p
	c.header = h
	return &c
}

func (p *Unsubscribe) Type() byte {
	return UNSUBSCRIBE
}

func (p *Unsubscribe) String() string{
	return "unsubscribe"
}

func (p *Unsubscribe) Topics() []string {
	topics := make([]string, len(p.topics))
	for idx, t := range p.topics {
		topics[idx] = t.String()
	}
	return topics
}
func (p *Unsubscribe) MessageId() int {
	return p.messageId.Int()
}

func (h *Unsubscribe) FixedHeader(l int) types.Field {
	return h.header.WithRemainingLength(l)
}
func (h *Unsubscribe) VariableHeaders() []types.Field {
	return []types.Field{
		&h.messageId,
	}
}
func (h *Unsubscribe) Payload() []types.Field {
	p := make([]types.Field, len(h.topics))
	for idx, t := range h.topics {
		p[idx] = &t
	}
	return p
}

func (p *Unsubscribe) Decode(r io.Reader) error {
	err := p.messageId.Decode(r)
	if err != nil {
		return err
	}
	progress := 0
	idx := 0
	for progress < p.header.remLength-p.messageId.Length() {
		p.topics = append(p.topics, types.MqttString{})
		p.topics[idx].Decode(r)
		progress += p.topics[idx].Length()
		idx += 1
	}
	return nil
}

func (c *Unsubscribe) Packet() types.Packet {
	return c
}
