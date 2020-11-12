package packet

import (
	"encoding/binary"
	"errors"
)

func UnmarshalPubRel(p *PubRel, buff []byte) (int, error) {
	p.MessageId = int32(binary.BigEndian.Uint16(buff))
	return 2, nil
}

type pubRelHandler func(*PubRel) error

func PubRelDecoder(fn pubRelHandler) func(h *Header, buffer []byte) error {
	return func(h *Header, buffer []byte) error {
		prelet := &PubRel{Header: h}
		_, err := UnmarshalPubRel(prelet, buffer)
		if err != nil {
			return err
		}
		return fn(prelet)
	}
}

func EncodePubRel(p *PubRel, buff []byte) (int, error) {
	if len(buff) < 2 {
		return 0, errors.New("buffer to short to encode message id")
	}
	binary.BigEndian.PutUint16(buff, uint16(p.MessageId))
	return 2, nil
}
func PubRelLength(p *PubRel) int {
	return 2
}

func (p *PubRel) Encode(buff []byte) (int, error) {
	return EncodePubRel(p, buff)
}
func (p *PubRel) Length() int {
	return PubRelLength(p)
}
func (p *PubRel) GetType() byte {
	return PUBREL
}
