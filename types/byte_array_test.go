package types

import (
	"testing"
	"bytes"
	"github.com/stretchr/testify/assert"
)

func TestMqttByteArray_Decode(t *testing.T) {
	a := []byte{0x0, 0x1, 0x2, 0x3}
	b := MqttByteArray{}
	b.SetLength(len(a))
	assert.Nil(t, b.Decode(bytes.NewReader(a)))
	assert.Equal(t, a, b.b)
}
func TestMqttByteArray_Encode(t *testing.T) {
	a := MqttByteArray{
		b: []byte{0x0, 0x1, 0x2},
		len: 3,
	}
	p, err := a.Encode()
	assert.Nil(t, err)
	assert.Equal(t, []byte{0x0, 0x1, 0x2}, p)
}