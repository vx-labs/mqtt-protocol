package packet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingreq_Decode(t *testing.T) {
	var b []byte
	n, err := decodePingReq(&PingReq{}, b)
	assert.Nil(t, err)
	assert.Equal(t, 0, n)
}
