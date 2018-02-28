package types

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
	"bytes"
)

func TestMqttInt_Length(t *testing.T) {
	i := MqttInt{
		lsb: 4,
		msb: 0,
	}
	assert.Equal(t, 2, i.Length())
}

func TestMqttInt_Bytes(t *testing.T) {
	i := MqttInt{
		lsb: 4,
		msb: 0,
	}
	assert.Equal(t, []byte{0, 4}, i.Bytes())
}
func TestMqttInt_Int(t *testing.T) {
	i := MqttInt{
		lsb: 4,
		msb: 0,
	}
	assert.Equal(t, 4, i.Int())
	j := MqttInt{
		lsb: 4,
		msb: 2,
	}
	assert.Equal(t, 516, j.Int())
}

func TestMqttInt_Encode(t *testing.T) {
	i := MqttInt{
		lsb: 4,
		msb: 2,
	}
	b, err := i.Encode()
	assert.Nil(t, err)
	assert.Equal(t, []byte{2, 4}, b)

}

func TestMqttInt_Decode(t *testing.T) {
	i := MqttInt{}
	b := []byte{2, 4}
	err := i.Decode(bytes.NewReader(b))
	assert.Nil(t, err)
	assert.Equal(t, byte(4), i.lsb)
	assert.Equal(t, byte(2), i.msb)
}

func TestMqttInt_FromInt(t *testing.T) {
	i := MqttInt{}
	assert.Equal(t, fmt.Errorf("value too big"), i.FromInt(8000000))
	assert.Equal(t, nil, i.FromInt(516))
	assert.Equal(t, byte(4), i.lsb)
	assert.Equal(t, byte(2), i.msb)
}
