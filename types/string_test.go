package types

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"bytes"
)

func TestMqttString_Bytes(t *testing.T) {
	s := MqttString{
		len: MqttInt{lsb: 4, msb: 0},
		b:   []byte{'m', 'o', 'c', 'k'},
	}
	assert.Equal(t, []byte{'m', 'o', 'c', 'k'}, s.Bytes())
}
func TestMqttString_Length(t *testing.T) {
	s := MqttString{
		b: []byte{'m', 'o', 'c', 'k'},
	}
	assert.Equal(t, 6, s.Length())
}
func TestMqttString_String(t *testing.T) {
	s := MqttString{
		b: []byte{'m', 'o', 'c', 'k'},
	}
	assert.Equal(t, "mock", s.String())
}

func TestMqttString_Encode(t *testing.T) {
	s := MqttString{
		len: MqttInt{lsb: 4, msb: 0},
		b:   []byte{'m', 'o', 'c', 'k'},
	}
	b, err := s.Encode()
	assert.Nil(t, err)
	assert.Equal(t, []byte{0, 4, 'm', 'o', 'c', 'k'}, b)
}

func TestMqttString_Decode(t *testing.T) {
	b := []byte{0, 4, 'm', 'o', 'c', 'k'}
	s := MqttString{}
	assert.Nil(t, s.Decode(bytes.NewReader(b)))
	assert.Equal(t, 4, s.len.Int())
	assert.Equal(t, []byte{'m', 'o', 'c', 'k'}, s.b)
}
