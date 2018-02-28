package packet

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSuback_WithGrantedQoS(t *testing.T) {
	a := Suback{}
	c := a.WithGrantedQoS([]byte{0, 2})
	if assert.Equal(t, 2, c.grantedQos.Length()) {
		assert.Equal(t, []byte{0, 2}, c.grantedQos.Bytes())
	}
}
