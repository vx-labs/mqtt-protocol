package packet

import (
	"encoding/binary"
	"errors"
)

func UnmarshalPubComp(p *PubComp, buff []byte) (int, error) {
	p.MessageId = int32(binary.BigEndian.Uint16(buff))
	return 2, nil
}

type pubCompHandler func(*PubComp) error

func PubCompDecoder(fn pubCompHandler) func(h *Header, buffer []byte) error {
	return func(h *Header, buffer []byte) error {
		pcompet := &PubComp{Header: h}
		_, err := UnmarshalPubComp(pcompet, buffer)
		if err != nil {
			return err
		}
		return fn(pcompet)
	}
}

func EncodePubComp(p *PubComp, buff []byte) (int, error) {
	if len(buff) < 2 {
		return 0, errors.New("buffer to short to encode message id")
	}
	binary.BigEndian.PutUint16(buff, uint16(p.MessageId))
	return 2, nil
}
func PubCompLength(p *PubComp) int {
	return 2
}

func (p *PubComp) Encode(buff []byte) (int, error) {
	return EncodePubComp(p, buff)
}
func (p *PubComp) Length() int {
	return PubCompLength(p)
}
func (p *PubComp) GetType() byte {
	return PUBCOMP
}
