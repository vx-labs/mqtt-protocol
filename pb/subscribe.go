package pb

import (
	"encoding/binary"
)

func countSubscribeTopics(buff []byte) (count int, err error) {
	for {
		length := int(binary.BigEndian.Uint16(buff))
		next := 2 + length + 1
		count++
		if next == len(buff) {
			return
		}
		buff = buff[next:]
	}
}

func decodeSubscribe(p *MqttSubscribe, buff []byte) (int, error) {
	p.MessageId = int32(binary.BigEndian.Uint16(buff))
	total := 2
	length, err := countSubscribeTopics(buff[total:])
	if err != nil {
		return total, err
	}
	p.Qos = make([]int32, length)
	p.Topic = make([][]byte, length)
	idx := 0
	for total < len(buff) {
		topic, n, err := decodeLP(buff[total:])
		total += n
		if err != nil {
			return total, err
		}
		qos := int32(buff[total])
		total++
		p.Qos[idx] = qos
		p.Topic[idx] = topic
		idx++
	}
	return total, nil
}

type subscribeHandler func(*MqttSubscribe) error

func SubscribeDecoder(fn subscribeHandler) func(h *MqttHeader, buffer []byte) error {
	return func(h *MqttHeader, buffer []byte) error {
		packet := &MqttSubscribe{Header: h}
		_, err := decodeSubscribe(packet, buffer)
		if err != nil {
			return err
		}
		return fn(packet)
	}
}
