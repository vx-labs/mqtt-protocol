package pb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnsubscribe_Decode(t *testing.T) {
	buff := []byte{
		0x0, 0x1,
		0x0, 0x2, 'a', 'b',
	}
	p := &MqttUnsubscribe{}
	n, err := decodeUnsubscribe(p, buff)
	assert.Equal(t, 6, n)
	assert.Nil(t, err)
}

func BenchmarkUnsubscribe_Decode(b *testing.B) {
	buff := []byte{
		0x0, 0x1,
		0x0, 0x2, 'a', 'b',
	}
	p := &MqttUnsubscribe{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		decodeUnsubscribe(p, buff)
	}
}
