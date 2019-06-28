package packet

import "encoding/binary"

func EncodeSubAck(p *SubAck, buff []byte) (int, error) {
	total := 0
	binary.BigEndian.PutUint16(buff, uint16(p.MessageId))
	total += 2
	for _, qos := range p.Qos {
		buff[total] = byte(qos)
		total++
	}
	return total, nil
}
func SubAckLength(p *SubAck) int {
	return 2 + len(p.Qos)
}

func (p *SubAck) Encode(buff []byte) (int, error) {
	return EncodeSubAck(p, buff)
}
func (p *SubAck) Length() int {
	return SubAckLength(p)
}
func (p *SubAck) GetType() byte {
	return SUBACK
}
