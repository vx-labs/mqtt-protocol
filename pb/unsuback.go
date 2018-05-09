package pb

import (
	"encoding/binary"
	"errors"
)

func decodeUnsubAck(p *MqttUnsubAck, buff []byte) (int, error) {
	p.MessageId = int32(binary.BigEndian.Uint16(buff))
	return 2, nil
}

type unsubAckHandler func(*MqttUnsubAck) error

func UnsubAckDecoder(fn unsubAckHandler) func(h *MqttHeader, buffer []byte) error {
	return func(h *MqttHeader, buffer []byte) error {
		packet := &MqttUnsubAck{Header: h}
		_, err := decodeUnsubAck(packet, buffer)
		if err != nil {
			return err
		}
		return fn(packet)
	}
}

func EncodeUnsubAck(p *MqttUnsubAck, buff []byte) (int, error) {
	if len(buff) < 2 {
		return 0, errors.New("buffer to short to encode message id")
	}
	binary.BigEndian.PutUint16(buff, uint16(p.MessageId))
	return 2, nil
}

func UnsubAckLength(p *MqttUnsubAck) int {
	return 2
}
