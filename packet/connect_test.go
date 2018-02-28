package packet

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"bytes"
)

func TestConnect_Decode(t *testing.T) {
	b := []byte{0x0, 0x6, 'M', 'Q', 'I', 's', 'd', 'p', 0x3, 0xce, 0, 0xa, 0x0, 0x1, 'a', 0x0, 0x1, 'b', 0x0, 0x1, 'c', 0x0, 0x2, 'd', 'e', 0x0, 0x1, 'f'}
	h := Connect{
		header: Header{
			remLength: 6,
		},
	}
	err := h.Decode(bytes.NewReader(b))
	assert.Nil(t, err)
	assert.Equal(t, []byte{'M', 'Q', 'I', 's', 'd', 'p'}, h.protocolName.Bytes())
	assert.Equal(t, byte(3), h.protocolVersionNumber.Byte())
	assert.Equal(t, byte(0xce), h.connectFlags.Byte())
	assert.Equal(t, []byte{0, 0xa}, h.keepAliveTimer.Bytes())
	assert.True(t, h.Flags().Username())
	assert.True(t, h.Flags().Password())
	assert.True(t, h.Flags().Will())
	assert.False(t, h.Flags().WillRetain())
	assert.Equal(t, byte(1), h.Flags().WillQoS())
	assert.True(t, h.Flags().CleanSession())
	assert.Equal(t, "a", h.ClientId())
	assert.Equal(t, "b", h.WillTopic())
	assert.Equal(t, "c", h.WillMessageString())
	assert.Equal(t, "de", h.UserName())
	assert.Equal(t, "f", h.Password())
}
