package packet

import (
	"encoding/binary"
)

func countUnsubscribeTopics(buff []byte) (count int, err error) {
	for {
		length := int(binary.BigEndian.Uint16(buff))
		next := 2 + length
		count++
		if next == len(buff) {
			return
		}
		buff = buff[next:]
	}
}

func decodeUnsubscribe(p *MqttUnsubscribe, buff []byte) (int, error) {
	p.MessageId = int32(binary.BigEndian.Uint16(buff))
	total := 2
	length, err := countUnsubscribeTopics(buff[total:])
	if err != nil {
		return total, err
	}
	p.Topic = make([][]byte, length)
	idx := 0
	for total < len(buff) {
		topic, n, err := decodeLP(buff[total:])
		total += n
		if err != nil {
			return total, err
		}
		p.Topic[idx] = topic
		idx++
	}
	return total, nil
}

type unsubscribeHandler func(*MqttUnsubscribe) error

func UnsubscribeDecoder(fn unsubscribeHandler) func(h *MqttHeader, buffer []byte) error {
	return func(h *MqttHeader, buffer []byte) error {
		packet := &MqttUnsubscribe{Header: h}
		_, err := decodeUnsubscribe(packet, buffer)
		if err != nil {
			return err
		}
		return fn(packet)
	}
}
