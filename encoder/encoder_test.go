package encoder

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vx-labs/mqtt-protocol/packet"
	"github.com/vx-labs/mqtt-protocol/pb"
)

func TestEncoder_EncodeHeader(t *testing.T) {
	buff := make([]byte, 5)
	n, err := encodeHeader(packet.PUBLISH, &pb.MqttHeader{}, 10, buff)
	assert.Nil(t, err)
	assert.Equal(t, 2, n)
	assert.Equal(t, packet.PUBLISH<<4, buff[0])
	assert.Equal(t, byte(0xa), buff[1])
}
func TestEncoder_EncodeHeader_Long(t *testing.T) {
	buff := make([]byte, 5)
	n, err := encodeHeader(packet.PUBLISH, &pb.MqttHeader{}, 129, buff)
	assert.Nil(t, err)
	assert.Equal(t, 3, n)
	assert.Equal(t, packet.PUBLISH<<4, buff[0])
	assert.Equal(t, byte(0x81), buff[1])
	assert.Equal(t, byte(0x1), buff[2])
}
func BenchmarkEncoder_EncodeHeader_Long(b *testing.B) {
	buff := make([]byte, 5)
	for i := 0; i < b.N; i++ {
		encodeHeader(packet.PUBLISH, &pb.MqttHeader{}, 129, buff)
	}
}

func TestEncoder_Publish(t *testing.T) {
	buff := make([]byte, 12)
	writer := bytes.NewBuffer([]byte{})
	e := NewEncoder(writer)
	err := e.Publish(&pb.MqttPublish{
		Header: &pb.MqttHeader{
			Qos: 1,
		},
		MessageId: 1,
		Topic:     []byte("a"),
		Payload:   []byte("pa"),
	}, buff)
	assert.Nil(t, err)
	assert.Equal(t,
		[]byte{0x32, 0x7, 0x0, 0x1, 'a', 0x0, 0x1, 'p', 'a'},
		writer.Bytes())
}

func BenchmarkEncoder_Publish(b *testing.B) {
	buff := make([]byte, 12)
	writer := bytes.NewBuffer([]byte{})
	e := NewEncoder(writer)
	p := &pb.MqttPublish{
		Header: &pb.MqttHeader{
			Qos: 1,
		},
		MessageId: 1,
		Topic:     []byte("a"),
		Payload:   []byte("pa"),
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		e.Publish(p, buff)
	}
}
func TestEncoder_PubAck(t *testing.T) {
	buff := make([]byte, 12)
	writer := bytes.NewBuffer([]byte{})
	e := NewEncoder(writer)
	err := e.PubAck(&pb.MqttPubAck{
		Header: &pb.MqttHeader{
			Qos: 1,
		},
		MessageId: 9,
	}, buff)
	assert.Nil(t, err)
	assert.Equal(t,
		[]byte{0x42, 0x2, 0x0, 0x9},
		writer.Bytes())
}
func BenchmarkEncoder_PubAck(b *testing.B) {
	buff := make([]byte, 12)
	writer := bytes.NewBuffer([]byte{})
	e := NewEncoder(writer)
	p := &pb.MqttPubAck{
		Header: &pb.MqttHeader{
			Qos: 1,
		},
		MessageId: 9,
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		e.PubAck(p, buff)
	}
}
func TestEncoder_PingResp(t *testing.T) {
	buff := make([]byte, 12)
	writer := bytes.NewBuffer([]byte{})
	e := NewEncoder(writer)
	err := e.PingResp(&pb.MqttPingResp{
		Header: &pb.MqttHeader{},
	}, buff)
	assert.Nil(t, err)
	assert.Equal(t,
		[]byte{0xd0, 0x0},
		writer.Bytes())
}
func TestEncoder_SubAck(t *testing.T) {
	buff := make([]byte, 12)
	writer := bytes.NewBuffer([]byte{})
	e := NewEncoder(writer)
	err := e.SubAck(&pb.MqttSubAck{
		Header:    &pb.MqttHeader{},
		MessageId: 12,
		Qos:       []int32{1, 2, 1},
	}, buff)
	assert.Nil(t, err)
	assert.Equal(t,
		[]byte{0x90, 0x5, 0x0, 0xc, 0x1, 0x2, 0x1},
		writer.Bytes())
}
func TestEncoder_UnsubAck(t *testing.T) {
	buff := make([]byte, 12)
	writer := bytes.NewBuffer([]byte{})
	e := NewEncoder(writer)
	err := e.UnsubAck(&pb.MqttUnsubAck{
		Header:    &pb.MqttHeader{},
		MessageId: 12,
	}, buff)
	assert.Nil(t, err)
	assert.Equal(t,
		[]byte{0xb0, 0x2, 0x0, 0xc},
		writer.Bytes())
}
