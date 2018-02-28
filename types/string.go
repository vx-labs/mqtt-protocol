package types

import (
	"fmt"
	"io"
)

type MqttString struct {
	len MqttInt
	b   []byte
}

func (s *MqttString) Decode(r io.Reader) error {
	err := s.len.Decode(r)
	if err != nil {
		return err
	}
	if s.len.Int() < 1 {
		return nil
	}
	b := make([]byte, s.len.Int())
	_, err = r.Read(b)
	if err != nil {
		return err
	}
	s.b = b
	return nil
}
func (s *MqttString) Encode() ([]byte, error) {
	if (s.len.Int()) < 1 {
		return nil, fmt.Errorf("refusing to encode null string")
	}
	buffer := make([]byte, s.Length())

	b, err := s.len.Encode()
	if err != nil {
		return nil, err
	}
	for i := 0; i < s.len.Length(); i++ {
		buffer[i] = b[i]
	}
	for i := 2; i < s.Length(); i ++ {
		buffer[i] = s.b[i-2]
	}
	return buffer, nil
}

func (s *MqttString) Length() int {
	return s.len.Length() + len(s.b)
}

func (s *MqttString) String() string {
	return string(s.b)
}

func (s *MqttString) FromString(n string) {
	s.b = []byte(n)
	s.len.FromInt(len(s.b))
}

func (s *MqttString) Bytes() []byte {
	return s.b
}
