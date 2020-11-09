package packet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubscribe_Encode(t *testing.T) {
	p := &Subscribe{Header: &Header{}, MessageId: 1, Qos: []int32{1}, Topic: [][]byte{{'a', 'b'}}}
	buff := make([]byte, 7)
	p.Encode(buff)

	assert.Equal(t, []byte{
		0x0, 0x1,
		0x0, 0x2, 'a', 'b',
		0x1,
	}, buff)

}
func TestSubscribe_Decode(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		buff := []byte{
			0x0, 0x1,
			0x0, 0x2, 'a', 'b',
			0x1,
		}
		p := &Subscribe{}
		n, err := UnmarshalSubscribe(p, buff)
		assert.Equal(t, 7, n)
		assert.Nil(t, err)
	})
	t.Run("invalid", func(t *testing.T) {
		buff := []byte{
			0x0, 0x1,
			0x0, 0x2, 'a', 'b',
		}
		p := &Subscribe{}
		_, err := UnmarshalSubscribe(p, buff)
		assert.NotNil(t, err)
	})
}

func TestSubscribe_CountTopics(t *testing.T) {
	buff := []byte{
		0x0, 0x2, 'a', 'b', 0x1,
		0x0, 0x2, 'a', 'b', 0x1,
		0x0, 0x2, 'a', 'b', 0x1,
		0x0, 0x2, 'a', 'b', 0x1,
	}
	n, err := countSubscribeTopics(buff)
	assert.Nil(t, err)
	assert.Equal(t, 4, n)
}
func BenchmarkSubscribe_CountTopics(b *testing.B) {
	buff := []byte{}
	for i := 0; i < 100; i++ {
		buff = append(buff, 0x0, 0x2, 'a', 'b', 0x1)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		countSubscribeTopics(buff)
	}
}

func BenchmarkSubscribe_Decode(b *testing.B) {
	buff := []byte{
		0x0, 0x1,
		0x0, 0x2, 'a', 'b',
		0x1,
	}
	p := &Subscribe{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		UnmarshalSubscribe(p, buff)
	}
}
