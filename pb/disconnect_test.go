package pb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDisconnect_Decode(t *testing.T) {
	var b []byte
	n, err := decodeDisconnect(&MqttDisconnect{}, b)
	assert.Nil(t, err)
	assert.Equal(t, 0, n)
}
