package pb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnect_Decode(t *testing.T) {
	b := []byte{0x0, 0x6, 'M', 'Q', 'I', 's', 'd', 'p', 0x3, 0xce, 0, 0xa, 0x0, 0x1, 'a', 0x0, 0x1, 'b', 0x0, 0x1, 'c', 0x0, 0x2, 'd', 'e', 0x0, 0x1, 'f'}
	p := &MqttConnect{}
	n, err := decodeConnect(p, b)
	assert.Nil(t, err)
	assert.Equal(t, len(b), n)
}

func BenchmarkConnect_Decode(b *testing.B) {
	buff := []byte{0x0, 0x6, 'M', 'Q', 'I', 's', 'd', 'p', 0x3, 0xce, 0, 0xa, 0x0, 0x1, 'a', 0x0, 0x1, 'b', 0x0, 0x1, 'c', 0x0, 0x2, 'd', 'e', 0x0, 0x1, 'f'}
	p := &MqttConnect{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		decodeConnect(p, buff)
	}
}