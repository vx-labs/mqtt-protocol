package pb

import (
	"encoding/binary"
	"errors"
)

func decodePubAck(p *MqttPubAck, buff []byte) (int, error) {
	p.MessageId = int32(binary.BigEndian.Uint16(buff))
	return 2, nil
}

type pubAckHandler func(*MqttPubAck) error

func PubAckDecoder(fn pubAckHandler) func(h *MqttHeader, buffer []byte) error {
	return func(h *MqttHeader, buffer []byte) error {
		packet := &MqttPubAck{Header: h}
		_, err := decodePubAck(packet, buffer)
		if err != nil {
			return err
		}
		return fn(packet)
	}
}

func EncodePubAck(p *MqttPubAck, buff []byte) (int, error) {
	if len(buff) < 2 {
		return 0, errors.New("buffer to short to encode message id")
	}
	binary.BigEndian.PutUint16(buff, uint16(p.MessageId))
	return 2, nil
}
