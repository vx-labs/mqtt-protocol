package types

import (
	"io"
)

type MqttByte struct {
	b byte
}

func (q *MqttByte) FromByte(b byte) error {
	q.b = b
	return nil
}
func (q *MqttByte) Length() int {
	return 1
}
func (q *MqttByte) Byte() byte {
	return q.b
}
func (q *MqttByte) Encode() ([]byte, error) {
	return []byte{q.b}, nil
}
func (q *MqttByte) Decode(r io.Reader) error {
	b := make([]byte, 1)
	_, err := r.Read(b)
	if err != nil {
		return err
	}
	q.b = b[0]
	return nil
}
