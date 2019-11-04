package packet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPublish_Decode_QoS0(t *testing.T) {
	b := []byte{0x0, 0x1, 'a', 0x1, 0x2, 0x3}
	p := &Publish{
		Header: &Header{Qos: 0},
	}
	n, err := UnmarshalPublish(p, b)
	assert.Nil(t, err)
	assert.Equal(t, 6, n)
	assert.Equal(t, "a", string(p.Topic))
	assert.Equal(t, []byte{0x1, 0x2, 0x3}, p.Payload)
}
func TestPublish_Decode_QoS1(t *testing.T) {
	b := []byte{0x0, 0x1, 'a', 0x0, 0x1, 0x1, 0x2, 0x3}
	p := &Publish{
		Header: &Header{Qos: 1},
	}
	n, err := UnmarshalPublish(p, b)
	assert.Nil(t, err)
	assert.Equal(t, 8, n)
	assert.Equal(t, int32(1), p.MessageId)
	assert.Equal(t, "a", string(p.Topic))
	assert.Equal(t, []byte{0x1, 0x2, 0x3}, p.Payload)
}
func BenchmarkPublish_Decode_QoS0(b *testing.B) {
	buff := []byte{0x0, 0x1, 'a', 0x1, 0x2, 0x3}
	p := &Publish{
		Header: &Header{Qos: 0},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		UnmarshalPublish(p, buff)
	}
}
func BenchmarkPublish_Decode_QoS1(b *testing.B) {
	buff := []byte{0x0, 0x1, 'a', 0x0, 0x1, 0x1, 0x2, 0x3}
	p := &Publish{
		Header: &Header{Qos: 1},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		UnmarshalPublish(p, buff)
	}
}
func TestPublish_Encode_QoS0(t *testing.T) {
	buff := make([]byte, 6)
	p := &Publish{
		Header:  &Header{Qos: 0},
		Topic:   []byte("a"),
		Payload: []byte{1, 2, 3},
	}
	n, err := EncodePublish(p, buff)
	assert.Nil(t, err)
	assert.Equal(t, 6, n)
	assert.Equal(t, []byte{0x0, 0x1, 'a', 0x1, 0x2, 0x3}, buff)
}
func TestPublish_Encode_QoS1(t *testing.T) {
	buff := make([]byte, 8)
	p := &Publish{
		Header:    &Header{Qos: 1},
		Topic:     []byte("a"),
		Payload:   []byte{1, 2, 3},
		MessageId: 9,
	}
	n, err := EncodePublish(p, buff)
	assert.Nil(t, err)
	assert.Equal(t, 8, n)
	assert.Equal(t, []byte{0x0, 0x1, 'a', 0x0, 0x9, 0x1, 0x2, 0x3}, buff)
}
func BenchmarkPublish_Encode_QoS0(b *testing.B) {
	buff := make([]byte, 8)
	p := &Publish{
		Header:    &Header{Qos: 1},
		Topic:     []byte("a"),
		Payload:   []byte{1, 2, 3},
		MessageId: 9,
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		EncodePublish(p, buff)
	}
}

func BenchmarkPublish_Length(b *testing.B) {
	p := &Publish{
		Header:    &Header{Qos: 1},
		Topic:     []byte("a"),
		Payload:   []byte{1, 2, 3},
		MessageId: 9,
	}
	for i := 0; i < b.N; i++ {
		PublishLength(p)
	}
}
