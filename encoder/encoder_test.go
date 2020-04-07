package encoder

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vx-labs/mqtt-protocol/packet"
)

func TestEncoder_EncodeHeader(t *testing.T) {
	buff := make([]byte, 4)
	n, err := encodeHeader(packet.PUBLISH, &packet.Header{}, 10, buff)
	assert.Nil(t, err)
	assert.Equal(t, 2, n)
	assert.Equal(t, packet.PUBLISH<<4, buff[0])
	assert.Equal(t, byte(0xa), buff[1])
}
func TestEncoder_EncodeHeader_Long(t *testing.T) {
	buff := make([]byte, 4)
	n, err := encodeHeader(packet.PUBLISH, &packet.Header{}, 129, buff)
	assert.Nil(t, err)
	assert.Equal(t, 3, n)
	assert.Equal(t, packet.PUBLISH<<4, buff[0])
	assert.Equal(t, byte(0x81), buff[1])
	assert.Equal(t, byte(0x1), buff[2])
}
func BenchmarkEncoder_EncodeHeader_Long(b *testing.B) {
	buff := make([]byte, 4)
	for i := 0; i < b.N; i++ {
		encodeHeader(packet.PUBLISH, &packet.Header{}, 129, buff)
	}
}

func TestEncoder_Publish(t *testing.T) {
	writer := bytes.NewBuffer([]byte{})
	e := New(writer)
	err := e.Publish(&packet.Publish{
		Header: &packet.Header{
			Qos: 1,
		},
		MessageId: 1,
		Topic:     []byte("a"),
		Payload:   []byte("pa"),
	})
	assert.Nil(t, err)
	assert.Equal(t,
		[]byte{0x32, 0x7, 0x0, 0x1, 'a', 0x0, 0x1, 'p', 'a'},
		writer.Bytes())
}

func BenchmarkEncoder_Publish(b *testing.B) {
	writer := ioutil.Discard
	e := New(writer)
	p := &packet.Publish{
		Header: &packet.Header{
			Qos: 1,
		},
		MessageId: 1,
		Topic:     []byte("a"),
		Payload:   []byte("pa"),
	}
	for i := 0; i < b.N; i++ {
		e.Publish(p)
	}
}
func TestEncoder_PubAck(t *testing.T) {
	writer := bytes.NewBuffer([]byte{})
	e := New(writer)
	err := e.PubAck(&packet.PubAck{
		Header: &packet.Header{
			Qos: 1,
		},
		MessageId: 9,
	})
	assert.Nil(t, err)
	assert.Equal(t,
		[]byte{0x42, 0x2, 0x0, 0x9},
		writer.Bytes())
}
func BenchmarkEncoder_PubAck(b *testing.B) {
	writer := bytes.NewBuffer([]byte{})
	e := New(writer)
	p := &packet.PubAck{
		Header: &packet.Header{
			Qos: 1,
		},
		MessageId: 9,
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		e.PubAck(p)
	}
}
func TestEncoder_PingResp(t *testing.T) {
	writer := bytes.NewBuffer([]byte{})
	e := New(writer)
	err := e.PingResp(&packet.PingResp{
		Header: &packet.Header{},
	})
	assert.Nil(t, err)
	assert.Equal(t,
		[]byte{0xd0, 0x0},
		writer.Bytes())
}
func TestEncoder_SubAck(t *testing.T) {
	writer := bytes.NewBuffer([]byte{})
	e := New(writer)
	err := e.SubAck(&packet.SubAck{
		Header:    &packet.Header{},
		MessageId: 12,
		Qos:       []int32{1, 2, 1},
	})
	assert.Nil(t, err)
	assert.Equal(t,
		[]byte{0x90, 0x5, 0x0, 0xc, 0x1, 0x2, 0x1},
		writer.Bytes())
}
func TestEncoder_UnsubAck(t *testing.T) {
	writer := bytes.NewBuffer([]byte{})
	e := New(writer)
	err := e.UnsubAck(&packet.UnsubAck{
		Header:    &packet.Header{},
		MessageId: 12,
	})
	assert.Nil(t, err)
	assert.Equal(t,
		[]byte{0xb0, 0x2, 0x0, 0xc},
		writer.Bytes())
}
func TestEncoder_ConnAck(t *testing.T) {
	writer := bytes.NewBuffer([]byte{})
	e := New(writer)
	err := e.ConnAck(&packet.ConnAck{
		Header:     &packet.Header{},
		ReturnCode: packet.CONNACK_REFUSED_BAD_USERNAME_OR_PASSWORD,
	})
	assert.Nil(t, err)
	assert.Equal(t,
		[]byte{0x20, 0x2, 0x0, 0x4},
		writer.Bytes())
}

func TestEncoder_RemLength(t *testing.T) {
	assert.Equal(t, 1, remLengthBits(0))
	assert.Equal(t, 2, remLengthBits(129))
}
func BenchmarkEncoder_RemLength(b *testing.B) {
	for i := 0; i < b.N; i++ {
		remLengthBits(129)
	}
}
