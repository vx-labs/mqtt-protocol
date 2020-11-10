package packet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingresp_Decode(t *testing.T) {
	var b []byte
	n, err := UnmarshalPingResp(&PingResp{}, b)
	assert.Nil(t, err)
	assert.Equal(t, 0, n)
}
