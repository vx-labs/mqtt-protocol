package packet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPubRel_Decode(t *testing.T) {
	buff := []byte{
		0x0, 0x1,
	}
	p := &PubRel{}
	n, err := UnmarshalPubRel(p, buff)
	assert.Equal(t, 2, n)
	assert.Equal(t, int32(1), p.MessageId)
	assert.Nil(t, err)
}
func BenchmarkPubRel_Decode(b *testing.B) {
	buff := []byte{
		0x0, 0x1,
	}
	p := &PubRel{}
	for i := 0; i < b.N; i++ {
		UnmarshalPubRel(p, buff)
	}
}

func TestPubRel_Encode(t *testing.T) {
	buff := make([]byte, 2)
	p := &PubRel{
		Header:    &Header{Qos: 1},
		MessageId: 9,
	}
	n, err := EncodePubRel(p, buff)
	assert.Nil(t, err)
	assert.Equal(t, 2, n)
	assert.Equal(t, []byte{0x0, 0x9}, buff)
}
func BenchmarkPubRel_Encode(b *testing.B) {
	buff := make([]byte, 2)
	p := &PubRel{
		Header:    &Header{Qos: 1},
		MessageId: 9,
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		EncodePubRel(p, buff)
	}
}
