package packet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnsubAck_Decode(t *testing.T) {
	buff := []byte{
		0x0, 0x1,
	}
	p := &UnsubAck{}
	n, err := decodeUnsubAck(p, buff)
	assert.Equal(t, 2, n)
	assert.Equal(t, int32(1), p.MessageId)
	assert.Nil(t, err)
}
func BenchmarkUnsubAck_Decode(b *testing.B) {
	buff := []byte{
		0x0, 0x1,
	}
	p := &UnsubAck{}
	for i := 0; i < b.N; i++ {
		decodeUnsubAck(p, buff)
	}
}

func TestUnsubAck_Encode(t *testing.T) {
	buff := make([]byte, 2)
	p := &UnsubAck{
		Header:    &Header{Qos: 1},
		MessageId: 9,
	}
	n, err := EncodeUnsubAck(p, buff)
	assert.Nil(t, err)
	assert.Equal(t, 2, n)
	assert.Equal(t, []byte{0x0, 0x9}, buff)
}
func BenchmarkUnsubAck_Encode(b *testing.B) {
	buff := make([]byte, 2)
	p := &UnsubAck{
		Header:    &Header{Qos: 1},
		MessageId: 9,
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		EncodeUnsubAck(p, buff)
	}
}
