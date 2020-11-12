package packet

import (
	"encoding/binary"
	"errors"
)

func UnmarshalPubRec(p *PubRec, buff []byte) (int, error) {
	p.MessageId = int32(binary.BigEndian.Uint16(buff))
	return 2, nil
}

type pubRecHandler func(*PubRec) error

func PubRecDecoder(fn pubRecHandler) func(h *Header, buffer []byte) error {
	return func(h *Header, buffer []byte) error {
		precet := &PubRec{Header: h}
		_, err := UnmarshalPubRec(precet, buffer)
		if err != nil {
			return err
		}
		return fn(precet)
	}
}

func EncodePubRec(p *PubRec, buff []byte) (int, error) {
	if len(buff) < 2 {
		return 0, errors.New("buffer to short to encode message id")
	}
	binary.BigEndian.PutUint16(buff, uint16(p.MessageId))
	return 2, nil
}
func PubRecLength(p *PubRec) int {
	return 2
}

func (p *PubRec) Encode(buff []byte) (int, error) {
	return EncodePubRec(p, buff)
}
func (p *PubRec) Length() int {
	return PubRecLength(p)
}
func (p *PubRec) GetType() byte {
	return PUBREC
}
