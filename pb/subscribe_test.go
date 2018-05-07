package pb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubscribe_Decode(t *testing.T) {
	buff := []byte{
		0x0, 0x1,
		0x0, 0x2, 'a', 'b',
		0x1,
	}
	p := &MqttSubscribe{}
	n, err := decodeSubscribe(p, buff)
	assert.Equal(t, 7, n)
	assert.Nil(t, err)
}

func TestSubscribe_CountTopics(t *testing.T) {
	buff := []byte{
		0x0, 0x2, 'a', 'b', 0x1,
		0x0, 0x2, 'a', 'b', 0x1,
		0x0, 0x2, 'a', 'b', 0x1,
		0x0, 0x2, 'a', 'b', 0x1,
	}
	n, err := countTopics(buff)
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
		countTopics(buff)
	}
}

func BenchmarkSubscribe_Decode(b *testing.B) {
	buff := []byte{
		0x0, 0x1,
		0x0, 0x2, 'a', 'b',
		0x1,
	}
	p := &MqttSubscribe{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		decodeSubscribe(p, buff)
	}
}