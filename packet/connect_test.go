package packet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnect_Decode(t *testing.T) {
	t.Run("empty keepalive", func(t *testing.T) {
		b := []byte{0x0, 0x6, 'M', 'Q', 'I', 's', 'd', 'p', 0x3, 0xce, 0, 0xa, 0x0, 0x1, 'a', 0x0, 0x1, 'b', 0x0, 0x1, 'c', 0x0, 0x2, 'd', 'e', 0x0, 0x1, 'f'}
		p := &Connect{}
		n, err := unmarshalConnect(p, b)
		assert.Nil(t, err)
		assert.Equal(t, len(b), n)
	})
	t.Run("empty keepalive", func(t *testing.T) {
		b := []byte{0x0, 0x6, 'M', 'Q', 'I', 's', 'd', 'p', 0x3, 0xce, 0, 0, 0x0, 0x1, 'a', 0x0, 0x1, 'b', 0x0, 0x1, 'c', 0x0, 0x2, 'd', 'e', 0x0, 0x1, 'f'}
		p := &Connect{}
		_, err := unmarshalConnect(p, b)
		assert.Nil(t, err)
		assert.Equal(t, int32(30), p.KeepaliveTimer)
	})
}

func TestConnect_Encode(t *testing.T) {
	p := &Connect{
		Clean: true,
	}
	buff := make([]byte, ConnectLength(p))
	n, err := EncodeConnect(p, buff)
	assert.Nil(t, err)
	assert.Equal(t, ConnectLength(p), n)
	assert.Equal(t, []byte{0x0, 0x4, 0x4d, 0x51, 0x54, 0x54, 0x4, 0x0, 0x0, 0x0, 0x0, 0x0}, buff)
}

func BenchmarkConnect_Decode(b *testing.B) {
	buff := []byte{0x0, 0x6, 'M', 'Q', 'I', 's', 'd', 'p', 0x3, 0xce, 0, 0xa, 0x0, 0x1, 'a', 0x0, 0x1, 'b', 0x0, 0x1, 'c', 0x0, 0x2, 'd', 'e', 0x0, 0x1, 'f'}
	p := &Connect{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		unmarshalConnect(p, buff)
	}
}
