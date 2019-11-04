package packet

import (
	"encoding/binary"
	"errors"
)

func UnmarshalPubAck(p *PubAck, buff []byte) (int, error) {
	p.MessageId = int32(binary.BigEndian.Uint16(buff))
	return 2, nil
}

type pubAckHandler func(*PubAck) error

func PubAckDecoder(fn pubAckHandler) func(h *Header, buffer []byte) error {
	return func(h *Header, buffer []byte) error {
		packet := &PubAck{Header: h}
		_, err := UnmarshalPubAck(packet, buffer)
		if err != nil {
			return err
		}
		return fn(packet)
	}
}

func EncodePubAck(p *PubAck, buff []byte) (int, error) {
	if len(buff) < 2 {
		return 0, errors.New("buffer to short to encode message id")
	}
	binary.BigEndian.PutUint16(buff, uint16(p.MessageId))
	return 2, nil
}
func PubAckLength(p *PubAck) int {
	return 2
}

func (p *PubAck) Encode(buff []byte) (int, error) {
	return EncodePubAck(p, buff)
}
func (p *PubAck) Length() int {
	return PubAckLength(p)
}
func (p *PubAck) GetType() byte {
	return PUBACK
}
