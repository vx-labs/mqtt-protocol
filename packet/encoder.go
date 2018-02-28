package packet

import (
	"io"
	"github.com/vx-labs/mqtt-protocol/types"
	"bytes"
)

type Encoder struct {
}

func NewEncoder() *Encoder {
	e := &Encoder{
	}
	return e
}

func (e *Encoder) encodeBlock(writer io.Writer, i types.Encodable) error {
	b, err := i.Encode()
	if err != nil {
		return err
	}
	_, err = writer.Write(b)
	if err != nil {
		return err
	}
	return nil
}

func (e *Encoder) Encode(p types.Packet) ([]byte, error) {
	pktLen := 0
	for _, e := range p.VariableHeaders() {
		pktLen += e.Length()
	}
	for _, e := range p.Payload() {
		pktLen += e.Length()
	}
	pkt := make([]byte, pktLen+2)
	w := bytes.NewBuffer(pkt)
	w.Reset()
	e.encodeBlock(w, p.FixedHeader(pktLen))
	for _, t := range p.VariableHeaders() {
		err := e.encodeBlock(w, t)
		if err != nil {
			return nil, err
		}
	}
	for _, t := range p.Payload() {
		err := e.encodeBlock(w, t)
		if err != nil {
			return nil, err
		}
	}
	return w.Bytes(), nil
}
