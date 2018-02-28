package packet

import (
	"testing"
	"bytes"
	"github.com/stretchr/testify/assert"
)

func decoder() Decoder {
	return Decoder{}
}

func TestDecoder_Decode_Disconnect(t *testing.T) {
	b := bytes.NewReader([]byte{0xE0, 0})
	d := decoder()
	p, err := d.Decode(b)
	assert.Nil(t, err)
	assert.Equal(t, DISCONNECT, p.Type())
	assert.IsType(t, &Disconnect{}, p)
}
func TestDecoder_Decode_Publish(t *testing.T) {
	b := bytes.NewReader([]byte{0x32, 0x7, 0x0, 0x1, 'a', 0x0, 0x1, 'p', 'a'})
	d := decoder()
	p, err := d.Decode(b)
	assert.Nil(t, err)
	assert.Equal(t, PUBLISH, p.Type())
	assert.IsType(t, &Publish{}, p)
	publish := p.(*Publish)
	assert.Equal(t, 1, publish.MessageId())
	assert.Equal(t, "a", publish.TopicName())
	assert.Equal(t, "pa", publish.MessageString())
}
func TestDecoder_Decode_Pingreq(t *testing.T) {
	b := bytes.NewReader([]byte{0xC0, 0})
	d := decoder()
	p, err := d.Decode(b)
	assert.Nil(t, err)
	assert.Equal(t, PINGREQ, p.Type())
	assert.IsType(t, &Pingreq{}, p)
}
