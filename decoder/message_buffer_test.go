package decoder

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vx-labs/mqtt-protocol/packet"
)

func TestDecoder_ReadMessageBuffer(t *testing.T) {
	buff := []byte{0x32, 0x7, 0x0, 0x1, 'a', 0x0, 0x1, 'p', 'a'}
	reader := bytes.NewReader(buff)
	p := &packet.Header{}
	pType, buff, _, err := readMessageBuffer(p, make([]byte, 4), reader)
	assert.Nil(t, err)
	assert.Equal(t, byte(3), pType)
	assert.Equal(t, 7, len(buff))
}
func TestDecoder_ReadMessageBuffer_Long(t *testing.T) {
	buff := []byte{0x32, 0x81, 0x1,
		0x0, 0x1, 'a',
		0x0, 0x83}
	for i := 0; i < 0x83; i++ {
		buff = append(buff, 'a')
	}
	reader := bytes.NewReader(buff)
	p := &packet.Header{}
	pType, buff, _, err := readMessageBuffer(p, make([]byte, 4), reader)
	assert.Nil(t, err)
	assert.Equal(t, byte(3), pType)
	assert.Equal(t, 129, len(buff))
}
func BenchmarkDecoder_ReadMessageBuffer(b *testing.B) {
	buff := []byte{0x32, 0x7, 0x0, 0x1, 'a', 0x0, 0x1, 'p', 'a'}
	longBuff := []byte{0x32, 0x81, 0x1,
		0x0, 0x1, 'a',
		0x0, 0x83}
	for i := 0; i < 0x83; i++ {
		buff = append(longBuff, 'a')
	}
	reader := bytes.NewReader(buff)
	p := &packet.Header{}
	b.Run("short", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			reader.Reset(buff)
			readMessageBuffer(p, make([]byte, 4), reader)
		}
	})
	b.Run("long", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			reader.Reset(longBuff)
			readMessageBuffer(p, make([]byte, 4), reader)
		}
	})
}
