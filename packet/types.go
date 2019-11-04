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

type Encoder interface {
	Encode(buff []byte) (int, error)
	Length() int
	GetHeader() *Header
	GetType() byte
}

type Decoder interface {
	Type() byte
	Unmarshal(buf []byte) error
}

func (*Connect) Type() byte { return CONNECT }
func (c *Connect) Unmarshal(buf []byte) error {
	_, err := unmarshalConnect(c, buf)
	return err
}
func (*ConnAck) Type() byte { return CONNACK }

func (*Publish) Type() byte { return PUBLISH }
func (c *Publish) Unmarshal(buf []byte) error {
	_, err := UnmarshalPublish(c, buf)
	return err
}
func (*PubAck) Type() byte { return PUBACK }
func (c *PubAck) Unmarshal(buf []byte) error {
	_, err := UnmarshalPubAck(c, buf)
	return err
}
func (*Subscribe) Type() byte { return SUBSCRIBE }
func (c *Subscribe) Unmarshal(buf []byte) error {
	_, err := UnmarshalSubscribe(c, buf)
	return err
}
func (*SubAck) Type() byte { return SUBACK }
func (c *SubAck) Unmarshal(buf []byte) error {
	_, err := UnmarshalSubAck(c, buf)
	return err
}
func (*UnsubAck) Type() byte { return UNSUBACK }
func (c *UnsubAck) Unmarshal(buf []byte) error {
	_, err := UnmarshalUnsubAck(c, buf)
	return err
}
func (*Unsubscribe) Type() byte { return UNSUBSCRIBE }
func (c *Unsubscribe) Unmarshal(buf []byte) error {
	_, err := UnmarshalUnsubscribe(c, buf)
	return err
}
func (*PingReq) Type() byte { return PINGREQ }
func (c *PingReq) Unmarshal(buf []byte) error {
	_, err := UnmarshalPingReq(c, buf)
	return err
}
func (*PingResp) Type() byte   { return PINGRESP }
func (*Disconnect) Type() byte { return DISCONNECT }
