package pb

import "encoding/binary"

func EncodeSubAck(p *MqttSubAck, buff []byte) (int, error) {
	total := 0
	binary.BigEndian.PutUint16(buff, uint16(p.MessageId))
	total += 2
	for _, qos := range p.Qos {
		buff[total] = byte(qos)
		total++
	}
	return total, nil
}
func SubAckLength(p *MqttSubAck) int {
	return 2 + len(p.Qos)
}
