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
func encode(packetType byte, header *packet.Header, boundary, total int, buff []byte) error {
	n, err := encodeHeader(packetType, header, total, buff[:boundary])
	total += n
	if err != nil {
		return err
	}

	if n < boundary {
		copy(buff[boundary-n:total], buff[boundary:])
	}
	return nil
}

func encodeHeader(packetType byte, header *packet.Header, remLength int, buff []byte) (int, error) {
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
	if size <= 127 {
		return 1
	} else if size <= 16383 {
		return 2
	} else if size <= 2097151 {
		return 3
	} else {
		return 4
	}
}

func (e *Encoder) Publish(p *packet.Publish) error {
	return e.Encode(p)
}
func (e *Encoder) PubAck(p *packet.PubAck) error {
	return e.Encode(p)
}
func (e *Encoder) PingResp(p *packet.PingResp) error {
	return e.Encode(p)
}
func (e *Encoder) SubAck(p *packet.SubAck) error {
	return e.Encode(p)
}
func (e *Encoder) UnsubAck(p *packet.UnsubAck) error {
	return e.Encode(p)
}
func (e *Encoder) ConnAck(p *packet.ConnAck) error {
	return e.Encode(p)
}
func (e *Encoder) Encode(p packet.Encoder) error {
	buffer, err := Marshal(p)
	if err != nil {
		return err
	}
	return e.flush(buffer)
}

func MarshalPublish(p *packet.Publish) ([]byte, error) {
	return Marshal(p)
}

func MarshalPubAck(p *packet.PubAck) ([]byte, error) {
	return Marshal(p)
}
func MarshalPingResp(p *packet.PingResp) ([]byte, error) {
	return Marshal(p)
}
func MarshalSubAck(p *packet.SubAck) ([]byte, error) {
	return Marshal(p)
}
func MarshalUnsubAck(p *packet.UnsubAck) ([]byte, error) {
	return Marshal(p)
}
func MarshalConnAck(p *packet.ConnAck) ([]byte, error) {
	return Marshal(p)
}

func Marshal(p packet.Encoder) ([]byte, error) {
	length := p.Length()
	headerLength := 1 + remLengthBits(length)
	buffer := make([]byte, headerLength+length)

	total, err := p.Encode(buffer[headerLength:])
	if err != nil {
		return nil, err
	}
	return buffer[:total+headerLength], encode(p.GetType(), p.GetHeader(), headerLength, total, buffer)
}
