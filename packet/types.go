package packet

//go:generate protoc --go_out=plugins=grpc:. pb.proto
const (
	_ byte = iota
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

type Packet interface {
	Type() byte
	Encode(buff []byte) (int, error)
}
type Encoder interface {
	Packet
	Encode(buff []byte) (int, error)
	Length() int
	GetHeader() *Header
}

type Decoder interface {
	Packet
	UnmarshalMQTT(buf []byte) error
}

func (*Connect) Type() byte { return CONNECT }
func (c *Connect) UnmarshalMQTT(buf []byte) error {
	_, err := unmarshalConnect(c, buf)
	return err
}
func (*ConnAck) Type() byte { return CONNACK }

func (*Publish) Type() byte { return PUBLISH }
func (c *Publish) UnmarshalMQTT(buf []byte) error {
	_, err := UnmarshalPublish(c, buf)
	return err
}
func (*PubAck) Type() byte { return PUBACK }
func (c *PubAck) UnmarshalMQTT(buf []byte) error {
	_, err := UnmarshalPubAck(c, buf)
	return err
}
func (*Subscribe) Type() byte { return SUBSCRIBE }
func (c *Subscribe) UnmarshalMQTT(buf []byte) error {
	_, err := UnmarshalSubscribe(c, buf)
	return err
}
func (*SubAck) Type() byte { return SUBACK }
func (c *SubAck) UnmarshalMQTT(buf []byte) error {
	_, err := UnmarshalSubAck(c, buf)
	return err
}
func (*UnsubAck) Type() byte { return UNSUBACK }
func (c *UnsubAck) UnmarshalMQTT(buf []byte) error {
	_, err := UnmarshalUnsubAck(c, buf)
	return err
}
func (*Unsubscribe) Type() byte { return UNSUBSCRIBE }
func (c *Unsubscribe) UnmarshalMQTT(buf []byte) error {
	_, err := UnmarshalUnsubscribe(c, buf)
	return err
}
func (*PingReq) Type() byte { return PINGREQ }
func (c *PingReq) UnmarshalMQTT(buf []byte) error {
	_, err := UnmarshalPingReq(c, buf)
	return err
}
func (*PingResp) Type() byte   { return PINGRESP }
func (*Disconnect) Type() byte { return DISCONNECT }
func (c *Disconnect) UnmarshalMQTT(buf []byte) error {
	_, err := UnmarshalDisconnect(c, buf)
	return err
}
