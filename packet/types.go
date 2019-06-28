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
