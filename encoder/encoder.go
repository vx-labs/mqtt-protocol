package encoder

import (
	"encoding/binary"
	"io"

	"github.com/vx-labs/mqtt-protocol/packet"
	"github.com/vx-labs/mqtt-protocol/pb"
)

type Encoder struct {
	w io.Writer
}

func New(w io.Writer) *Encoder {
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
func (e *Encoder) encode(packetType byte, header *pb.MqttHeader, boundary, total int, buff []byte) error {
	n, err := encodeHeader(packetType, header, total, buff[:boundary])
	total += n
	if err != nil {
		return err
	}

	if n < boundary {
		copy(buff[boundary-n:total], buff[boundary:])
	}
	return e.flush(buff[:total])
}

func encodeHeader(packetType byte, header *pb.MqttHeader, remLength int, buff []byte) (int, error) {
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
	return e.encode(packet.PUBLISH, p.Header, 4, total, buff)
}
func (e *Encoder) PubAck(p *pb.MqttPubAck, buff []byte) error {
	total, err := pb.EncodePubAck(p, buff[2:])
	if err != nil {
		return err
	}
	return e.encode(packet.PUBACK, p.Header, 2, total, buff)
}
func (e *Encoder) PingResp(p *pb.MqttPingResp, buff []byte) error {
	total, err := pb.EncodePingResp(p, buff[2:])
	if err != nil {
		return err
	}
	return e.encode(packet.PINGRESP, p.Header, 2, total, buff)
}
func (e *Encoder) SubAck(p *pb.MqttSubAck, buff []byte) error {
	total, err := pb.EncodeSubAck(p, buff[4:])
	if err != nil {
		return err
	}
	return e.encode(packet.SUBACK, p.Header, 4, total, buff)
}
func (e *Encoder) UnsubAck(p *pb.MqttUnsubAck, buff []byte) error {
	total, err := pb.EncodeUnsubAck(p, buff[4:])
	if err != nil {
		return err
	}
	return e.encode(packet.UNSUBACK, p.Header, 4, total, buff)
}
func (e *Encoder) ConnAck(p *pb.MqttConnAck, buff []byte) error {
	total, err := pb.EncodeConnAck(p, buff[2:])
	if err != nil {
		return err
	}
	return e.encode(packet.UNSUBACK, p.Header, 2, total, buff)
}
