package packet

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestEncoder_Encode(t *testing.T) {
	a := New().PingResp().Packet()
	b, err := NewEncoder().Encode(a)
	assert.Nil(t, err)
	assert.Equal(t, []byte{208, 0}, b)
}
