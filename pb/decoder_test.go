package pb

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecoder_ReadMessageBuffer(t *testing.T) {
	buff := []byte{0x32, 0x7, 0x0, 0x1, 'a', 0x0, 0x1, 'p', 'a'}
	reader := bytes.NewReader(buff)
	decoder := NewDecoder()
	p := &MqttHeader{}
	pType, buff, err := decoder.readMessageBuffer(p, reader)
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
	decoder := NewDecoder()
	p := &MqttHeader{}
	pType, buff, err := decoder.readMessageBuffer(p, reader)
	assert.Nil(t, err)
	assert.Equal(t, byte(3), pType)
	assert.Equal(t, 136, len(buff))
}
func BenchmarkDecoder_ReadMessageBuffer(b *testing.B) {
	buff := []byte{0x32, 0x7, 0x0, 0x1, 'a', 0x0, 0x1, 'p', 'a'}
	reader := bytes.NewReader(buff)
	d := NewDecoder()
	p := &MqttHeader{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.readMessageBuffer(p, reader)
		reader.Reset(buff)
	}
}
