package packet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnAck_Encode(t *testing.T) {
	b := make([]byte, 2)
	p := &MqttConnAck{ReturnCode: CONNACK_CONNECTION_ACCEPTED}
	n, err := EncodeConnAck(p, b)
	assert.Nil(t, err)
	assert.Equal(t, len(b), n)
}

func BenchmarkConnAck_Encode(b *testing.B) {
	buff := make([]byte, 2)
	p := &MqttConnAck{ReturnCode: CONNACK_CONNECTION_ACCEPTED}
	for i := 0; i < b.N; i++ {
		EncodeConnAck(p, buff)
	}
}
