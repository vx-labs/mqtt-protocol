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

func UnmarshalUnsubscribe(p *Unsubscribe, buff []byte) (int, error) {
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

type unsubscribeHandler func(*Unsubscribe) error

func UnsubscribeDecoder(fn unsubscribeHandler) func(h *Header, buffer []byte) error {
	return func(h *Header, buffer []byte) error {
		packet := &Unsubscribe{Header: h}
		_, err := UnmarshalUnsubscribe(packet, buffer)
		if err != nil {
			return err
		}
		return fn(packet)
	}
}

func UnsubscribeLength(p *Subscribe) int {
	size := 2
	for idx := range p.Topic {
		size += 2 + len(p.Topic[idx])
	}
	return size
}

func EncodeUnsubscribe(p *Unsubscribe, buff []byte) (int, error) {
	total := 2
	binary.BigEndian.PutUint16(buff[0:], uint16(p.MessageId))
	for idx := range p.Topic {
		n, err := encodeLP(p.Topic[idx], buff[total:])
		total += n
		if err != nil {
			return total, err
		}
	}
	return total, nil
}
func (p *Unsubscribe) Encode(buff []byte) (int, error) {
	return EncodeUnsubscribe(p, buff)
}
