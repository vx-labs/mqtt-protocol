package packet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeLP(t *testing.T) {
	buff := []byte{0, 2, 'a', 'b'}
	str, n, err := decodeLP(buff)
	assert.Nil(t, err)
	assert.Equal(t, 4, n)
	assert.Equal(t, "ab", string(str))
}

func BenchmarkDecodeLP(b *testing.B) {
	buff := []byte{0, 2, 'a', 'b'}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		decodeLP(buff)
	}
}

func TestDecodeString(t *testing.T) {
	buff := []byte{0, 2, 'a', 'b'}
	str, n, err := decodeString(buff)
	assert.Nil(t, err)
	assert.Equal(t, 4, n)
	assert.Equal(t, "ab", str)
}

func BenchmarkDecodeString(b *testing.B) {
	buff := []byte{0, 2, 'a', 'b'}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		decodeString(buff)
	}
}
func TestEncodeLP(t *testing.T) {
	buff := make([]byte, 4)
	value := []byte("ab")
	n, err := encodeLP(value, buff)
	assert.Nil(t, err)
	assert.Equal(t, 4, n)
	assert.Equal(t, buff, []byte{0, 2, 'a', 'b'})
}

func BenchmarkEncodeLP(b *testing.B) {
	buff := make([]byte, 4)
	value := []byte("ab")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		encodeLP(value, buff)
	}
}
