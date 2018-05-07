package pb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPublish_Decode_QoS0(t *testing.T) {
	b := []byte{0x0, 0x1, 'a', 0x1, 0x2, 0x3}
	p := &MqttPublish{
		Header: &MqttHeader{Qos: 0},
	}
	n, err := decodePublish(p, b)
	assert.Nil(t, err)
	assert.Equal(t, 6, n)
	assert.Equal(t, "a", string(p.Topic))
	assert.Equal(t, []byte{0x1, 0x2, 0x3}, p.Payload)
}
func TestPublish_Decode_QoS1(t *testing.T) {
	b := []byte{0x0, 0x1, 'a', 0x0, 0x1, 0x1, 0x2, 0x3}
	p := &MqttPublish{
		Header: &MqttHeader{Qos: 1},
	}
	n, err := decodePublish(p, b)
	assert.Nil(t, err)
	assert.Equal(t, 8, n)
	assert.Equal(t, int32(1), p.MessageId)
	assert.Equal(t, "a", string(p.Topic))
	assert.Equal(t, []byte{0x1, 0x2, 0x3}, p.Payload)
}
func BenchmarkPublish_Decode_QoS0(b *testing.B) {
	buff := []byte{0x0, 0x1, 'a', 0x1, 0x2, 0x3}
	p := &MqttPublish{
		Header: &MqttHeader{Qos: 0},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		decodePublish(p, buff)
	}
}
func BenchmarkPublish_Decode_QoS1(b *testing.B) {
	buff := []byte{0x0, 0x1, 'a', 0x0, 0x1, 0x1, 0x2, 0x3}
	p := &MqttPublish{
		Header: &MqttHeader{Qos: 1},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		decodePublish(p, buff)
	}
}
