package types

import (
	"fmt"
	"io"
)

const (
	QoS0 byte = iota
	QoS1
	QoS2
)

type MqttQos struct {
	b byte
}

func (q *MqttQos) FromByte(b byte) error {
	if q.b > 2 {
		return fmt.Errorf("invalid QoS value")
	}
	q.b = b
	return nil
}
func (q *MqttQos) Byte() byte {
	return q.b
}
func (q *MqttQos) Encode() ([]byte, error) {
	if q.b > 2 {
		return nil, fmt.Errorf("invalid QoS value")
	}
	return []byte{q.b}, nil
}
func (q *MqttQos) Decode(r io.Reader) error {
	b := make([]byte, 1)
	_, err := r.Read(b)
	if err != nil {
		return err
	}
	if b[0] > 2 {
		return fmt.Errorf("invalid QoS value")
	}
	q.b = b[0]
	return nil
}
