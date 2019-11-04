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

func UnmarshalSubAck(p *SubAck, buff []byte) (int, error) {
	p.MessageId = int32(binary.BigEndian.Uint16(buff))
	total := 2
	qosSlice := buff[total:]
	p.Qos = make([]int32, len(qosSlice))
	for idx := range qosSlice {
		p.Qos[idx] = int32(qosSlice[idx])
		total++
	}
	return total, nil
}

type subAckHandler func(*SubAck) error

func SubAckDecoder(fn subAckHandler) func(h *Header, buffer []byte) error {
	return func(h *Header, buffer []byte) error {
		packet := &SubAck{Header: h}
		_, err := UnmarshalSubAck(packet, buffer)
		if err != nil {
			return err
		}
		return fn(packet)
	}
}
