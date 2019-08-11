package packet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubAck_Decode(t *testing.T) {
	buff := []byte{
		0, 22, 2,
	}
	p := &SubAck{}
	n, err := decodeSubAck(p, buff)
	assert.Equal(t, 3, n)
	assert.Equal(t, int32(22), p.MessageId)
	assert.Equal(t, 1, len(p.Qos))
	assert.Equal(t, int32(2), p.Qos[0])
	assert.Nil(t, err)
}
func TestSuback_Encode(t *testing.T) {
	buff := make([]byte, 10)
	packet := &SubAck{
		Header:    &Header{},
		MessageId: 9,
		Qos:       []int32{1, 0, 2, 1},
	}
	n, err := EncodeSubAck(packet, buff)
	assert.Nil(t, err)
	assert.Equal(t, 6, n)
}
func BenchmarkSuback_Encode(b *testing.B) {
	buff := make([]byte, 10)
	packet := &SubAck{
		Header:    &Header{},
		MessageId: 9,
		Qos:       []int32{1, 0, 2, 1},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		EncodeSubAck(packet, buff)
	}
}
