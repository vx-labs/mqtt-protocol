package packet

import (
	"testing"
	"bytes"
	"github.com/stretchr/testify/assert"
)

func TestUnsubscribe_Decode(t *testing.T) {
	b := []byte{0x0, 0x1, 0x0, 0x1, 0x61}
	u := Unsubscribe{header:Header{remLength: 5}}
	err := u.Decode(bytes.NewReader(b))
	assert.Nil(t, err)
	assert.Equal(t, 1, u.MessageId())
	assert.Equal(t, "a", u.topics[0].String())
}
