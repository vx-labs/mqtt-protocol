package packet

import (
	"testing"
	"bytes"
	"github.com/stretchr/testify/assert"
)

func TestPingreq_Decode(t *testing.T) {
	var b []byte
	r := bytes.NewReader(b)
	p := Pingreq{header: Header{
		remLength: 0,
	}}
	assert.Nil(t, p.Decode(r))
}