package types

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"bytes"
)

func TestMqttByte_Encode(t *testing.T) {
	q := MqttByte{b: 0x1}
	b, err := q.Encode()
	assert.Nil(t, err)
	assert.Equal(t, []byte{0x1}, b)
}
func TestMqttByte_Decode(t *testing.T) {
	a := []byte{0x1}
	q := MqttByte{}
	assert.Nil(t, q.Decode(bytes.NewReader([]byte{0x10})))
	assert.NotNil(t, q.Decode(bytes.NewReader([]byte{})))
	assert.Nil(t, q.Decode(bytes.NewReader(a)))
	assert.Equal(t, byte(0x1), q.b)
}
