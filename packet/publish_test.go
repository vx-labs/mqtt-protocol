package packet

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"bytes"
)

func TestPublish_Decode_NoHeader(t *testing.T) {
	var b []byte
	r := bytes.NewReader(b)
	p := Publish{}
	assert.NotNil(t, p.Decode(r))
}

func TestPublish_Decode_EmptyPayload(t *testing.T) {
	b := []byte{0x0, 0x1, 'a', 0x0, 0x1}
	r := bytes.NewReader(b)
	p := Publish{header: Header{
		remLength: 5,
		qos:       0x1,
	}}
	assert.Nil(t, p.Decode(r))
	assert.Equal(t, 1, p.messageId.Int())
	assert.Equal(t, "a", p.topicName.String())
	assert.Equal(t, []byte{}, p.Message())
}
func TestPublish_Decode(t *testing.T) {
	b := []byte{0x0, 0x1, 'a', 0x0, 0x1, 0x1, 0x2, 0x3}
	r := bytes.NewReader(b)
	p := Publish{header: Header{
		remLength: 8,
		qos:       0x1,
	}}
	assert.Nil(t, p.Decode(r))
	assert.Equal(t, p.messageId.Int(), 1)
	assert.Equal(t, p.topicName.String(), "a")
	assert.Equal(t, []byte{0x1, 0x2, 0x3}, p.Message())
}

func TestPublish_Payload(t *testing.T) {
	p := Publish{header: Header{
		remLength: 9,
		qos:       0x1,
	}}
	p.topicName.FromString("a")
	p.messageId.FromInt(1)
	assert.Equal(t, 1, len(p.Payload()))
	assert.Equal(t, 0, p.Payload()[0].Length())
	assert.Equal(t, 2, len(p.VariableHeaders()))
}

func TestPublish_WithTopicName(t *testing.T) {
	p := Publish{}
	c := p.WithTopicName("test")

	assert.Equal(t, "test", c.TopicName())
	assert.Equal(t, 6, c.topicName.Length())
}

func TestPublish_Length(t *testing.T) {
	p := Publish{}
	c := p.WithTopicName("test")
	c = c.WithMessage([]byte("mock"))

	if assert.Equal(t, 1, len(c.VariableHeaders())) {
		assert.Equal(t, 6, c.VariableHeaders()[0].Length())
	}
	assert.Equal(t, 4, len(c.Message()))
	assert.Equal(t, 1, len(c.Payload()))
	assert.Equal(t, 4, c.Payload()[0].Length())
}
