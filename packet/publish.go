package packet

import (
	"encoding/binary"
	"errors"
)

func UnmarshalPublish(p *Publish, buff []byte) (int, error) {
	topic, n, err := decodeLP(buff)
	total := n
	if err != nil {
		return total, err
	}
	p.Topic = topic
	if p.Header != nil && p.Header.Qos > 0 {
		messageID := binary.BigEndian.Uint16(buff[total:])
		total += 2
		p.MessageId = int32(messageID)
	}
	p.Payload = buff[total:]
	return len(buff), nil
}
func PublishLength(p *Publish) int {
	length := len(p.Payload) + 2 + len(p.Topic)
	if p.Header.Qos > 0 {
		length += 2
	}
	return length
}
func EncodePublish(p *Publish, buff []byte) (int, error) {
	total := 0
	n, err := encodeLP(p.Topic, buff)
	total += n
	if err != nil {
		return total, err
	}
	if p.Header != nil && p.Header.Qos > 0 {
		if len(buff[total:]) < 2 {
			return total, errors.New("buffer to short to encode message id")
		}
		binary.BigEndian.PutUint16(buff[total:], uint16(p.MessageId))
		total += 2
	}
	if len(buff[total:]) < len(p.Payload) {
		return total, errors.New("buffer to short to encode message payload")
	}
	copy(buff[total:], p.Payload)
	total += len(p.Payload)
	return total, nil
}

func (p *Publish) Encode(buff []byte) (int, error) {
	return EncodePublish(p, buff)
}
func (p *Publish) Length() int {
	return PublishLength(p)
}
func (p *Publish) GetType() byte {
	return PUBLISH
}

type publishHandler func(*Publish) error

func PublishDecoder(fn publishHandler) func(h *Header, buffer []byte) error {
	return func(h *Header, buffer []byte) error {
		packet := &Publish{Header: h}
		_, err := UnmarshalPublish(packet, buffer)
		if err != nil {
			return err
		}
		return fn(packet)
	}
}
