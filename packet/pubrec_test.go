package packet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPubRec_Decode(t *testing.T) {
	buff := []byte{
		0x0, 0x1,
	}
	p := &PubRec{}
	n, err := UnmarshalPubRec(p, buff)
	assert.Equal(t, 2, n)
	assert.Equal(t, int32(1), p.MessageId)
	assert.Nil(t, err)
}
func BenchmarkPubRec_Decode(b *testing.B) {
	buff := []byte{
		0x0, 0x1,
	}
	p := &PubRec{}
	for i := 0; i < b.N; i++ {
		UnmarshalPubRec(p, buff)
	}
}

func TestPubRec_Encode(t *testing.T) {
	buff := make([]byte, 2)
	p := &PubRec{
		Header:    &Header{Qos: 1},
		MessageId: 9,
	}
	n, err := EncodePubRec(p, buff)
	assert.Nil(t, err)
	assert.Equal(t, 2, n)
	assert.Equal(t, []byte{0x0, 0x9}, buff)
}
func BenchmarkPubRec_Encode(b *testing.B) {
	buff := make([]byte, 2)
	p := &PubRec{
		Header:    &Header{Qos: 1},
		MessageId: 9,
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		EncodePubRec(p, buff)
	}
}
