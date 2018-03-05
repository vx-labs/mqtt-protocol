package packet

import (
	"github.com/vx-labs/mqtt-protocol/types"
	"fmt"
	"io"
)

type Publish struct {
	header    Header
	topicName types.MqttString
	messageId types.MqttInt
	message   types.MqttByteArray
}

func NewPublish() *Publish {
	return &Publish{}
}

func (p *Publish) WithFixedHeader(h Header) *Publish {
	c := p.clone()
	c.header = h
	return &c
}

func (h *Publish) clone() Publish {
	return *h
}
func (h *Publish) Retain() bool {
	return h.header.Retain()
}
func (h *Publish) Dup() bool {
	return h.header.Dup()
}
func (h *Publish) TopicName() string {
	return h.topicName.String()
}
func (h *Publish) WithTopicName(t string) *Publish {
	c := h.clone()
	c.topicName.FromString(t)
	return &c
}

func (h *Publish) WithMessage(m []byte) *Publish {
	c := h.clone()
	c.message.FromBytes(m)
	return &c
}
func (h *Publish) QoS() byte {
	return h.header.qos
}
func (h *Publish) WithMessageId(i int) *Publish {
	c := h.clone()
	c.messageId.FromInt(i)
	return &c
}
func (h *Publish) MessageId() int {
	return h.messageId.Int()
}
func (h *Publish) MessageString() string {
	return string(h.Message())
}
func (h *Publish) Message() []byte {
	return h.message.Bytes()
}

func (h *Publish) Type() byte {
	return PUBLISH
}
func (h *Publish) String() string {
	return "publish"
}
func (h *Publish) FixedHeader(l int) types.Field {
	return h.header.WithRemainingLength(l)
}

func (c *Publish) VariableHeaders() ([]types.Field) {
	b := []types.Field{
		&c.topicName,
	}
	if c.header.qos > types.QoS0 {
		b = append(b, &c.messageId)
	}
	return b
}
func (c *Publish) Payload() ([]types.Field) {
	return []types.Field{&c.message}
}

func (h *Publish) Decode(r io.Reader) error {
	if h.header.remLength == 0 {
		return fmt.Errorf("invalid header")
	}
	counter := h.header.remLength
	for _, f := range h.VariableHeaders() {
		err := f.Decode(r)
		if err != nil {
			return err
		}
		counter -= f.Length()
	}
	if counter == 0 {
		return nil
	}
	h.message.SetLength(counter)
	return h.message.Decode(r)
}

func (c *Publish) Packet() types.Packet {
	return c
}
