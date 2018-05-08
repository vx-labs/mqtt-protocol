package encoder

import (
	"encoding/binary"
	"errors"
	"io"

	"github.com/vx-labs/mqtt-protocol/packet"
	"github.com/vx-labs/mqtt-protocol/pb"
)

type Encoder struct {
	w io.Writer
}

func NewEncoder(w io.Writer) *Encoder {
	e := &Encoder{w: w}
	return e
}
func (e *Encoder) flush(buff []byte) error {
	total := 0
	for total < len(buff) {
		n, err := e.w.Write(buff[total:])
		total += n
		if err != nil {
			return err
		}
	}
	return nil
}
func (e *Encoder) encode(packetType byte, header *pb.MqttHeader, total int, buff []byte) error {
	n, err := encodeHeader(packet.PUBLISH, header, total, buff[:4])
	total += n
	if err != nil {
		return err
	}

	if n < 4 {
		copy(buff[4-n:total], buff[4:])
	}
	return e.flush(buff[:total])
}

func encodeHeader(packetType byte, header *pb.MqttHeader, remLength int, buff []byte) (int, error) {
	if len(buff) < 4 {
		return 0, errors.New("buffer too small")
	}
	var dup, qos, retain byte
	if header.Dup {
		dup = 1
	}
	if header.Retain {
		retain = 1
	}
	qos = byte(header.Qos)

	buff[0] = (packetType << 4) + (dup << 3) + (qos << 1) + retain
	total := 1
	n := binary.PutUvarint(buff[1:], uint64(remLength))
	total += n
	return total, nil
}

func (e *Encoder) Publish(p *pb.MqttPublish, buff []byte) error {
	total, err := pb.EncodePublish(p, buff[4:])
	if err != nil {
		return err
	}
	return e.encode(packet.PUBLISH, p.Header, total, buff)
}
