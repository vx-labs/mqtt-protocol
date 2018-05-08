package pb

import "encoding/binary"

func decodePublish(p *MqttPublish, buff []byte) (int, error) {
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

type publishHandler func(*MqttPublish) error

func PublishDecoder(fn publishHandler) func(h *MqttHeader, buffer []byte) error {
	return func(h *MqttHeader, buffer []byte) error {
		packet := &MqttPublish{Header: h}
		_, err := decodePublish(packet, buffer)
		if err != nil {
			return err
		}
		return fn(packet)
	}
}
