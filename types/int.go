package types

import (
	"fmt"
	"io"
)

type MqttInt struct {
	msb byte
	lsb byte
}

func (s *MqttInt) FromInt(i int) error {
	if i > 0xffff {
		return fmt.Errorf("value too big")
	}
	s.lsb = byte(i & 0xff)
	s.msb = byte(i >> 8)
	return nil
}
func (s *MqttInt) Decode(r io.Reader) error {
	b := make([]byte, 2)
	_, err := r.Read(b)
	if err != nil {
		return err
	}
	s.msb = b[0]
	s.lsb = b[1]
	return nil
}
func (s *MqttInt) Encode() ([]byte, error) {
	return s.Bytes(), nil
}

func (s *MqttInt) Length() int {
	return 2
}

func (s *MqttInt) Int() int {
	return int(s.lsb) + int(s.msb)<<8
}

func (s *MqttInt) Bytes() []byte {
	return []byte{s.msb, s.lsb}
}
