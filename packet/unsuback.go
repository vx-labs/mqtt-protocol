package packet

import (
	"encoding/binary"
	"errors"
)

func UnmarshalUnsubAck(p *UnsubAck, buff []byte) (int, error) {
	p.MessageId = int32(binary.BigEndian.Uint16(buff))
	return 2, nil
}

type unsubAckHandler func(*UnsubAck) error

func UnsubAckDecoder(fn unsubAckHandler) func(h *Header, buffer []byte) error {
	return func(h *Header, buffer []byte) error {
		packet := &UnsubAck{Header: h}
		_, err := UnmarshalUnsubAck(packet, buffer)
		if err != nil {
			return err
		}
		return fn(packet)
	}
}

func EncodeUnsubAck(p *UnsubAck, buff []byte) (int, error) {
	if len(buff) < 2 {
		return 0, errors.New("buffer to short to encode message id")
	}
	binary.BigEndian.PutUint16(buff, uint16(p.MessageId))
	return 2, nil
}

func UnsubAckLength(p *UnsubAck) int {
	return 2
}

func (p *UnsubAck) Encode(buff []byte) (int, error) {
	return EncodeUnsubAck(p, buff)
}
func (p *UnsubAck) Length() int {
	return UnsubAckLength(p)
}
func (p *UnsubAck) GetType() byte {
	return UNSUBACK
}
