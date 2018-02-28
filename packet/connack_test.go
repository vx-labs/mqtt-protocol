package packet

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestConnack_WithReturnCode(t *testing.T) {
	w := CONNACK_CONNECTION_ACCEPTED
	p := New().Connack()
	c := p.WithReturnCode(w)
	assert.Equal(t, w, c.returnCode.Byte())
	assert.Equal(t, w, c.ReturnCode())
	assert.Equal(t, byte(0), p.returnCode.Byte())
}
func TestConnack_WithReturnCode_Oneline(t *testing.T) {
	w := CONNACK_CONNECTION_ACCEPTED
	c := New().Connack().WithReturnCode(w)
	assert.Equal(t, w, c.returnCode.Byte())
	assert.Equal(t, w, c.ReturnCode())
}
