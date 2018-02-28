package types

import (
	"io"
	"fmt"
)

type MqttByteArray struct {
	b   []byte
	len int
}

func (q *MqttByteArray) FromBytes(b []byte) error {
	q.b = b
	return nil
}
func (q *MqttByteArray) SetLength(i int){
	q.len = i
}
func (q *MqttByteArray) Length() int {
	return len(q.b)
}
func (q *MqttByteArray) Bytes() []byte {
	if q.b == nil {
		return []byte{}
	}
	return q.b
}
func (q *MqttByteArray) Encode() ([]byte, error) {
	return q.b, nil
}
func (q *MqttByteArray) Decode(r io.Reader) error {
	if q.len < 1 {
		return fmt.Errorf("refusing to decode an empty array")
	}
	b := make([]byte, q.len)
	_, err := r.Read(b)
	if err != nil {
		return err
	}
	q.b = b
	return nil
}
