package pb

import (
	"encoding/binary"
)

func countTopics(buff []byte) (count int, err error) {
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
	length, err := countTopics(buff[total:])
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
