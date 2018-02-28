package packet

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/vx-labs/mqtt-protocol/types"
	"bytes"
)

func TestHeader_Decode(t *testing.T) {
	b := []byte{0x10, 0}
	h := Header{}
	err := h.Decode(bytes.NewReader(b))
	assert.Nil(t, err)
	assert.Equal(t, CONNECT, h.packetType)
	assert.Equal(t, 0, h.remLength)
}

func TestHeader_Encode(t *testing.T) {
	b := []byte{0x46}
	h := Header{
		retain:     0,
		qos:        types.QoS0,
		dup:        0,
		packetType: CONNECT,
	}
	b, err := h.Encode()
	assert.Nil(t, err)
	assert.Equal(t, byte(0x10), b[0])

	h = Header{
		retain:     1,
		qos:        types.QoS2,
		dup:        0,
		packetType: CONNECT,
	}
	b, err = h.Encode()
	assert.Nil(t, err)
	assert.Equal(t, byte(0x15), b[0])
	h = Header{
		retain:     1,
		qos:        types.QoS2,
		dup:        0,
		packetType: PUBLISH,
	}
	b, err = h.Encode()
	assert.Nil(t, err)
	assert.Equal(t, byte(0x35), b[0])
}

func Test_recurseRemainingLength(t *testing.T) {
	b := []byte{2}

	value, err := recurseRemainingLength([]byte{193}, 0, 1, bytes.NewReader(b))
	assert.Nil(t, err)
	assert.Equal(t, 321, value)
}

func Test_encodingRemainingLength(t *testing.T) {
	i := encodingRemainingLength(321)
	assert.Equal(t, []byte{193, 2}, i)
}

func TestHeader_EncodeRemainingLength_ZeroLength(t *testing.T) {
	b := encodingRemainingLength(0)
	assert.Equal(t, 1, len(b))
}
