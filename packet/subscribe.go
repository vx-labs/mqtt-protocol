package packet

import (
	"encoding/binary"
	"errors"
)

func countSubscribeTopics(buff []byte) (count int, err error) {
	for {
		if len(buff) < 3 {
			return 0, errors.New("buffer too short")
		}
		length := int(binary.BigEndian.Uint16(buff))
		next := 2 + length + 1
		count++
		if next == len(buff) {
			return
		}
		if len(buff) < next {
			return 0, errors.New("buffer too short")
		}
		buff = buff[next:]
	}
}

func UnmarshalSubscribe(p *Subscribe, buff []byte) (int, error) {
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

type subscribeHandler func(*Subscribe) error

func SubscribeDecoder(fn subscribeHandler) func(h *Header, buffer []byte) error {
	return func(h *Header, buffer []byte) error {
		packet := &Subscribe{Header: h}
		_, err := UnmarshalSubscribe(packet, buffer)
		if err != nil {
			return err
		}
		return fn(packet)
	}
}

func SubscribeLength(p *Subscribe) int {
	size := 2
	for idx := range p.Topic {
		size += 2 + len(p.Topic[idx]) + 1
	}
	return size
}

func EncodeSubscribe(p *Subscribe, buff []byte) (int, error) {
	total := 2
	binary.BigEndian.PutUint16(buff[0:], uint16(p.MessageId))
	for idx := range p.Topic {
		n, err := encodeLP(p.Topic[idx], buff[total:])
		total += n
		if err != nil {
			return total, err
		}
		buff[total] = byte(p.Qos[idx])
		total++
	}
	return total, nil
}
func (p *Subscribe) Encode(buff []byte) (int, error) {
	return EncodeSubscribe(p, buff)
}
func (p *Subscribe) Length() int {
	return SubscribeLength(p)
}
