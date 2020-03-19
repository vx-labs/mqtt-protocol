package encoder

import (
	"encoding/binary"
	"io"

	"github.com/vx-labs/mqtt-protocol/packet"
)

type defaultBufferProvider struct{}

func (d *defaultBufferProvider) New(size int) ([]byte, error) {
	return make([]byte, size), nil
}

type staticBufferProvider struct {
	b []byte
}

func (d *staticBufferProvider) New(size int) ([]byte, error) {
	if len(d.b) >= size {
		return d.b[:size], nil
	}
	d.b = make([]byte, size)
	return d.b, nil
}

type BufferProvider interface {
	New(size int) ([]byte, error)
}

type Encoder struct {
	w      io.Writer
	buffer BufferProvider
}

func New(w io.Writer) *Encoder {
	e := &Encoder{w: w, buffer: &defaultBufferProvider{}}
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
func (e *Encoder) Encode(p packet.Packet) error {
	buffer, err := e.Marshal(p)
	if err != nil {
		return err
	}
	return e.flush(buffer)
}

func (e *Encoder) Marshal(p packet.Packet) ([]byte, error) {
	length := p.Length()
	headerLength := 1 + remLengthBits(length)
	buffer, err := e.buffer.New(headerLength + length)
	if err != nil {
		return nil, err
	}
	total, err := p.Encode(buffer[headerLength:])
	if err != nil {
		return nil, err
	}
	return buffer[:total+headerLength], encode(p.Type(), p.GetHeader(), headerLength, total, buffer)
}
