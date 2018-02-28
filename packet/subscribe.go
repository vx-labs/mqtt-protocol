package packet

import (
	"github.com/vx-labs/mqtt-protocol/types"
	"io"
)

func NewSubscribePacket(headers Header) *Subscribe {
	return &Subscribe{
		header: headers,
	}
}

type Subscribe struct {
	header        Header
	messageId     types.MqttInt
	subscriptions []types.TopicSubscription
}

func NewSubscribe() *Subscribe {
	return &Subscribe{}
}

func (p *Subscribe) WithFixedHeader(h Header) *Subscribe{
	c := *p
	c.header = h
	return &c
}


func (p *Subscribe) Type() byte{
	return SUBSCRIBE
}

func (p *Subscribe) String() string {
	return "subscribe"
}

func (p *Subscribe) Subscriptions() []types.TopicSubscription {
	return p.subscriptions
}
func (p *Subscribe) MessageId() int {
	return p.messageId.Int()
}

func (h *Subscribe) FixedHeader(l int) types.Field {
	return h.header.WithRemainingLength(l)
}
func (h *Subscribe) VariableHeaders() []types.Field {
	return []types.Field{
		&h.messageId,
	}
}
func (h *Subscribe) Payload() []types.Field {
	p := make([]types.Field, len(h.subscriptions))
	for idx, t := range h.subscriptions {
		p[idx] = &t
	}
	return p
}

func (p *Subscribe) Decode(r io.Reader) error {
	err := p.messageId.Decode(r)
	if err != nil {
		return err
	}
	progress := 0
	idx := 0
	for progress < p.header.remLength - p.messageId.Length() {
		p.subscriptions = append(p.subscriptions, types.TopicSubscription{})
		p.subscriptions[idx].Topic.Decode(r)
		progress += p.subscriptions[idx].Topic.Length()
		p.subscriptions[idx].Qos.Decode(r)
		progress += 1
		idx += 1
	}
	return nil
}

func (c *Subscribe) Packet() types.Packet {
	return c
}