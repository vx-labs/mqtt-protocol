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
