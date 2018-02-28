package packet

import (
	"io"
)

const (
	_           byte = iota
	CONNECT
	CONNACK
	PUBLISH
	PUBACK
	PUBREC
	PUBREL
	PUBCOMP
	SUBSCRIBE
	SUBACK
	UNSUBSCRIBE
	UNSUBACK
	PINGREQ
	PINGRESP
	DISCONNECT
)

type Header struct {
	packetType byte
	dup        byte
	qos        byte
	retain     byte
	remLength  int
}

func (h *Header) PacketType() byte {
	return h.packetType
}
func (h *Header) WithPacketType(t byte) *Header {
	c := *h
	c.packetType = t
	return &c
}

func (h *Header) WithDUP(f bool) *Header {
	c := *h
	if f {
		c.dup = 0x1
	} else {
		c.dup = 0x0
	}
	return &c
}

func (h *Header) WithRETAIN(f bool) *Header {
	c := *h
	if f {
		c.retain = 0x1
	} else {
		c.retain = 0x0
	}
	return &c
}

func (h *Header) WithQoS(q byte) *Header {
	c := *h
	c.qos = q
	return &c
}

func (h *Header) WithRemainingLength(l int) *Header {
	c := *h
	c.remLength = l
	return &c
}
func (h *Header) RemainingLength() int {
	return h.remLength
}
func (h *Header) Type() byte {
	return h.packetType
}

func (h *Header) Dup() bool {
	return h.dup == 1
}

func (h *Header) QoS() byte {
	return h.qos
}

func (h *Header) Retain() bool {
	return h.retain == 1
}

func recurseRemainingLength(s []byte, value, multiplier int, r io.Reader) (int, error) {
	b := s[0]
	hasNext := (b>>7)&0x1 == 1
	value += int(b&0x7f) * multiplier
	if hasNext {
		_, err := r.Read(s)
		if err != nil {
			return 0, err
		}
		return recurseRemainingLength(s, value, multiplier*128, r)
	}
	return value, nil
}

func remainingLength(cur []byte, r io.Reader) (int, error) {
	return recurseRemainingLength(cur, 0, 1, r)
}

func (h *Header) Decode(r io.Reader) error {
	s := make([]byte, 2)
	read := 0
	for read < len(s) {
		n, err := r.Read(s[read:])
		if err != nil {
			return err
		}
		read += n
	}
	fixedHeader := s[0]
	h.retain = fixedHeader & 0x1
	h.qos = fixedHeader & (0x3 << 1) >> 1
	h.dup = fixedHeader & (0x1 << 3) >> 3
	h.packetType = fixedHeader & (0xf << 4) >> 4

	variableHeader, err := remainingLength(s[1:], r)
	if err != nil {
		return err
	}
	h.remLength = variableHeader
	return nil
}

func (h *Header) Length() int {
	return 1 + len(encodingRemainingLength(h.remLength))
}
func (h *Header) Encode() ([]byte, error) {
	rlength := encodingRemainingLength(h.remLength)
	size := 1 + len(rlength)
	b := make([]byte, size)
	b[0] = (h.packetType << 4) + (h.dup << 3) + (h.qos << 1) + h.retain
	for idx, bit := range rlength {
		b[idx+1] = bit
	}
	return b, nil
}

func encodingRemainingLength(remLength int) []byte {
	if remLength == 0 {
		return []byte{0}
	}
	var bag []byte
	for remLength > 0 {
		digit := remLength % 128
		remLength /= 128
		if remLength > 0 {
			digit = digit | 0x80
		}
		bag = append(bag, byte(digit))
	}
	return bag
}
