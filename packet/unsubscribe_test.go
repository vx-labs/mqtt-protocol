package packet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnsubscribe_Decode(t *testing.T) {
	buff := []byte{
		0x0, 0x1,
		0x0, 0x2, 'a', 'b',
	}
	p := &Unsubscribe{}
	n, err := UnmarshalUnsubscribe(p, buff)
	assert.Equal(t, 6, n)
	assert.Nil(t, err)
}

func BenchmarkUnsubscribe_Decode(b *testing.B) {
	buff := []byte{
		0x0, 0x1,
		0x0, 0x2, 'a', 'b',
	}
	p := &Unsubscribe{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		UnmarshalUnsubscribe(p, buff)
	}
}
