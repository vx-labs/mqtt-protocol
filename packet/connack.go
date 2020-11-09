package packet

import (
	"errors"
)

const (
	CONNACK_CONNECTION_ACCEPTED int32 = iota
	CONNACK_REFUSED_BAD_PROTOCOL_VERSION
	CONNACK_REFUSED_IDENTIFIER_REJECTED
	CONNACK_REFUSED_SERVER_UNAVAILABLE
	CONNACK_REFUSED_BAD_USERNAME_OR_PASSWORD
	CONNACK_REFUSED_NOT_AUTHORIZED
)

func EncodeConnAck(p *ConnAck, buff []byte) (int, error) {
	if len(buff) < 2 {
		return 0, errors.New("buffer to short to encode connack")
	}
	buff[0] = 0x0
	buff[1] = byte(p.ReturnCode)
	return 2, nil
}
func ConnAckLength(p *ConnAck) int {
	return 2
}

func (p *ConnAck) Encode(buff []byte) (int, error) {
	return EncodeConnAck(p, buff)
}
func (p *ConnAck) Length() int {
	return ConnAckLength(p)
}
func (p *ConnAck) GetType() byte {
	return CONNACK
}
func unmarshalConnAck(p *ConnAck, buff []byte) (int, error) {
	if len(buff) < 2 {
		return 0, errors.New("buffer to short to decode connack")
	}
	p.ReturnCode = int32(buff[1])
	return 0, nil
}

type connAckHandler func(*ConnAck) error

func ConnAckDecoder(fn connAckHandler) func(h *Header, buffer []byte) error {
	return func(h *Header, buffer []byte) error {
		packet := &ConnAck{Header: h}
		_, err := unmarshalConnAck(packet, buffer)
		if err != nil {
			return err
		}
		return fn(packet)
	}
}
