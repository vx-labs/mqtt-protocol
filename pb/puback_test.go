package pb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPubAck_Decode(t *testing.T) {
	buff := []byte{
		0x0, 0x1,
	}
	p := &MqttPubAck{}
	n, err := decodePubAck(p, buff)
	assert.Equal(t, 2, n)
	assert.Equal(t, int32(1), p.MessageId)
	assert.Nil(t, err)
}
func BenchmarkPubAck_Decode(b *testing.B) {
	buff := []byte{
		0x0, 0x1,
	}
	p := &MqttPubAck{}
	for i := 0; i < b.N; i++ {
		decodePubAck(p, buff)
	}
}

func TestPubAck_Encode(t *testing.T) {
	buff := make([]byte, 2)
	p := &MqttPubAck{
		Header:    &MqttHeader{Qos: 1},
		MessageId: 9,
	}
	n, err := EncodePubAck(p, buff)
	assert.Nil(t, err)
	assert.Equal(t, 2, n)
	assert.Equal(t, []byte{0x0, 0x9}, buff)
}
func BenchmarkPubAck_Encode(b *testing.B) {
	buff := make([]byte, 2)
	p := &MqttPubAck{
		Header:    &MqttHeader{Qos: 1},
		MessageId: 9,
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		EncodePubAck(p, buff)
	}
}
