package pb

import (
	"encoding/binary"
	fmt "fmt"
)

func decodeString(buff []byte) (string, int, error) {
	b, n, err := decodeLP(buff)
	return string(b), n, err
}
func decodeLP(buff []byte) ([]byte, int, error) {
	if len(buff) < 2 {
		return nil, 0, fmt.Errorf("buffer too short to decode size")
	}
	size := int(binary.BigEndian.Uint16(buff))
	total := 2
	if len(buff) < size+2 {
		return nil, 0, fmt.Errorf("buffer too short to decode payload")
	}
	total += size
	return buff[2 : size+2], total, nil
}
func encodeLP(value []byte, buff []byte) (int, error) {
	size := len(value)
	if len(buff) < size+2 {
		return 0, fmt.Errorf("buffer too short to encode payload")
	}
	binary.BigEndian.PutUint16(buff, uint16(size))
	total := 2
	copy(buff[2:], value)
	total += size
	return total, nil
}
