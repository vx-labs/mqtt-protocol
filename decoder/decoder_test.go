package decoder

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vx-labs/mqtt-protocol/packet"
)

func TestDecoder_Decode(t *testing.T) {
	buff := []byte{0x32, 0x81, 0x1,
		0x0, 0x1, 'a',
		0x0, 0x83}
	for i := 0; i < 0x83; i++ {
		buff = append(buff, 'a')
	}
	reader := bytes.NewReader(buff)
	decoder := New(OnPublish(func(m *packet.Publish) error { return nil }))
	err := decoder.Decode(reader)
	assert.Nil(t, err)
}
func BenchmarkDecoder_Decode(b *testing.B) {
	buff := []byte{0x32, 0x7, 0x0, 0x1, 'a', 0x0, 0x1, 'p', 'a'}
	longBuff := []byte{0x32, 0x81, 0x1,
		0x0, 0x1, 'a',
		0x0, 0x83}
	for i := 0; i < 0x83; i++ {
		buff = append(longBuff, 'a')
	}
	reader := bytes.NewReader(buff)
	d := New()
	b.Run("short", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			reader.Reset(buff)
			d.Decode(reader)
		}
	})
	b.Run("long", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			reader.Reset(longBuff)
			d.Decode(reader)
		}
	})
}

func TestDecoder_DecodePubAck(t *testing.T) {
	buff := []byte{0x40, 0x2, 0x0,
		0x1}
	for i := 0; i < 0x83; i++ {
		buff = append(buff, 'a')
	}
	reader := bytes.NewReader(buff)
	ok := false
	decoder := New(OnPubAck(func(m *packet.PubAck) error { ok = true; return nil }))
	err := decoder.Decode(reader)
	assert.Nil(t, err)
	assert.True(t, ok)
}
