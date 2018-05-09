package encoder

import (
	"encoding/binary"
	"io"

	"github.com/vx-labs/mqtt-protocol/packet"
)

type Encoder struct {
	w io.Writer
}

func New(w io.Writer) *Encoder {
	e := &Encoder{w: w}
	return e
}
func (e *Encoder) flush(buff []byte) error {
	total := 0
	for total < len(buff) {
		n, err := e.w.Write(buff[total:])
		total += n
		if err != nil {
			return err
		}
	}
	return nil
}
func (e *Encoder) encode(packetType byte, header *packet.MqttHeader, boundary, total int, buff []byte) error {
	n, err := encodeHeader(packetType, header, total, buff[:boundary])
	total += n
	if err != nil {
		return err
	}

	if n < boundary {
		copy(buff[boundary-n:total], buff[boundary:])
	}
	return e.flush(buff[:total])
}

func encodeHeader(packetType byte, header *packet.MqttHeader, remLength int, buff []byte) (int, error) {
	var dup, qos, retain byte
	if header.Dup {
		dup = 1
	}
	if header.Retain {
		retain = 1
	}
	qos = byte(header.Qos)

	buff[0] = (packetType << 4) + (dup << 3) + (qos << 1) + retain
	total := 1
	n := binary.PutUvarint(buff[1:], uint64(remLength))
	total += n
	return total, nil
}

func remLengthBits(size int) int {
	value := 1
	for size > 0x80 {
		size = size >> 8
		value++
	}
	return value
}

func (e *Encoder) Publish(p *packet.MqttPublish) error {
	length := packet.PublishLength(p)
	headerLength := 1 + remLengthBits(length)
	buffer := make([]byte, headerLength+length)
	total, err := packet.EncodePublish(p, buffer[headerLength:])
	if err != nil {
		return err
	}
	return e.encode(packet.PUBLISH, p.Header, headerLength, total, buffer)
}
func (e *Encoder) PubAck(p *packet.MqttPubAck) error {
	length := packet.PubAckLength(p)
	headerLength := 1 + remLengthBits(length)
	buffer := make([]byte, headerLength+length)

	total, err := packet.EncodePubAck(p, buffer[headerLength:])
	if err != nil {
		return err
	}
	return e.encode(packet.PUBACK, p.Header, headerLength, total, buffer)
}
func (e *Encoder) PingResp(p *packet.MqttPingResp) error {
	length := packet.PingRespLength(p)
	headerLength := 1 + remLengthBits(length)
	buffer := make([]byte, headerLength+length)

	total, err := packet.EncodePingResp(p, buffer[headerLength:])
	if err != nil {
		return err
	}
	return e.encode(packet.PINGRESP, p.Header, headerLength, total, buffer)
}
func (e *Encoder) SubAck(p *packet.MqttSubAck) error {
	length := packet.SubAckLength(p)
	headerLength := 1 + remLengthBits(length)
	buffer := make([]byte, headerLength+length)

	total, err := packet.EncodeSubAck(p, buffer[headerLength:])
	if err != nil {
		return err
	}
	return e.encode(packet.SUBACK, p.Header, headerLength, total, buffer)
}
func (e *Encoder) UnsubAck(p *packet.MqttUnsubAck) error {
	length := packet.UnsubAckLength(p)
	headerLength := 1 + remLengthBits(length)
	buffer := make([]byte, headerLength+length)

	total, err := packet.EncodeUnsubAck(p, buffer[headerLength:])
	if err != nil {
		return err
	}
	return e.encode(packet.UNSUBACK, p.Header, headerLength, total, buffer)
}
func (e *Encoder) ConnAck(p *packet.MqttConnAck) error {
	length := packet.ConnAckLength(p)
	headerLength := 1 + remLengthBits(length)
	buffer := make([]byte, headerLength+length)

	total, err := packet.EncodeConnAck(p, buffer[headerLength:])
	if err != nil {
		return err
	}
	return e.encode(packet.CONNACK, p.Header, headerLength, total, buffer)
}
