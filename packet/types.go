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

func TypeString(p Packet) string {
	switch p.Type() {
	case CONNECT:
		return "CONNECT"
	case SUBSCRIBE:
		return "SUBSCRIBE"
	case PUBLISH:
		return "PUBLISH"
	case PINGREQ:
		return "PINGREQ"
	case PUBACK:
		return "PUBACK"
	case SUBACK:
		return "SUBACK"
	case UNSUBSCRIBE:
		return "UNSUBSCRIBE"
	case UNSUBACK:
		return "UNSUBACK"
	case DISCONNECT:
		return "DISCONNECT"
	default:
		return "Unsupported"
	}
}

type Packet interface {
	Type() byte
	Length() int
	Encode(buff []byte) (int, error)
	GetHeader() *Header
}
type Decoder interface {
	Packet
	UnmarshalMQTT(buf []byte) (int, error)
}

func (*Connect) Type() byte { return CONNECT }
func (c *Connect) UnmarshalMQTT(buf []byte) (int, error) {
	return unmarshalConnect(c, buf)
}
func (*ConnAck) Type() byte { return CONNACK }

func (*Publish) Type() byte { return PUBLISH }
func (c *Publish) UnmarshalMQTT(buf []byte) (int, error) {
	return UnmarshalPublish(c, buf)
}
func (*PubAck) Type() byte { return PUBACK }
func (c *PubAck) UnmarshalMQTT(buf []byte) (int, error) {
	return UnmarshalPubAck(c, buf)
}
func (*Subscribe) Type() byte { return SUBSCRIBE }
func (c *Subscribe) UnmarshalMQTT(buf []byte) (int, error) {
	return UnmarshalSubscribe(c, buf)
}
func (*SubAck) Type() byte { return SUBACK }
func (c *SubAck) UnmarshalMQTT(buf []byte) (int, error) {
	return UnmarshalSubAck(c, buf)
}
func (*UnsubAck) Type() byte { return UNSUBACK }
func (c *UnsubAck) UnmarshalMQTT(buf []byte) (int, error) {
	return UnmarshalUnsubAck(c, buf)
}
func (*Unsubscribe) Type() byte { return UNSUBSCRIBE }
func (c *Unsubscribe) UnmarshalMQTT(buf []byte) (int, error) {
	return UnmarshalUnsubscribe(c, buf)
}
func (*PingReq) Type() byte { return PINGREQ }
func (c *PingReq) UnmarshalMQTT(buf []byte) (int, error) {
	return UnmarshalPingReq(c, buf)
}
func (*PingResp) Type() byte   { return PINGRESP }
func (*Disconnect) Type() byte { return DISCONNECT }
func (c *Disconnect) UnmarshalMQTT(buf []byte) (int, error) {
	return UnmarshalDisconnect(c, buf)
}
