package types

import "io"

type Encodable interface {
	Length() int
	Encode() ([]byte, error)
}

type Decodable interface {
	Decode(io.Reader) (error)
	Packet() (Packet)
}

type Field interface {
	Length() int
	Encode() ([]byte, error)
	Decode(io.Reader) (error)
}

type Packet interface {
	Type() byte
	String() string
	Decode(io.Reader) (error)
	FixedHeader(remLength int) Field
	VariableHeaders() []Field
	Payload() []Field
}
