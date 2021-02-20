package decoder

import (
	"errors"
	fmt "fmt"
	"io"

	"github.com/google/btree"

	"github.com/vx-labs/mqtt-protocol/packet"
)

type Packet struct {
	timestamp int64
	err       error
	value     interface{}
}

func (p Packet) Err() error {
	return p.err
}
func (p Packet) Value() interface{} {
	return p.value
}

func (local Packet) Less(remote btree.Item) bool {
	return local.timestamp < remote.(Packet).timestamp
}

func decodeEncodedPacket(buf []byte, r io.Reader) (packet.Packet, int, error) {
	if len(buf) < 4 {
		return nil, 0, errors.New("buffer too short")
	}
	h := &packet.Header{}
	packetType, count, err := readMessageBuffer(h, buf, r)
	if err != nil {
		return nil, count, err
	}
	pkt, pktRead, err := unmarshalPacket(packetType, h, buf)
	return pkt, pktRead + count, err
}

func unmarshalPacket(packetType byte, header *packet.Header, buffer []byte) (packet.Packet, int, error) {
	var p packet.Decoder
	switch packetType {
	case packet.CONNECT:
		p = &packet.Connect{Header: header}
	case packet.CONNACK:
		p = &packet.ConnAck{Header: header}
	case packet.SUBSCRIBE:
		p = &packet.Subscribe{Header: header}
	case packet.PUBLISH:
		p = &packet.Publish{Header: header}
	case packet.PINGREQ:
		p = &packet.PingReq{Header: header}
	case packet.PINGRESP:
		p = &packet.PingResp{Header: header}
	case packet.PUBREC:
		p = &packet.PubRec{Header: header}
	case packet.PUBREL:
		p = &packet.PubRel{Header: header}
	case packet.PUBCOMP:
		p = &packet.PubComp{Header: header}
	case packet.PUBACK:
		p = &packet.PubAck{Header: header}
	case packet.SUBACK:
		p = &packet.SubAck{Header: header}
	case packet.UNSUBSCRIBE:
		p = &packet.Unsubscribe{Header: header}
	case packet.DISCONNECT:
		p = &packet.Disconnect{Header: header}
	default:
		err := fmt.Errorf("received unsuported packet type %v", packetType)
		return nil, 0, err
	}
	n, err := p.UnmarshalMQTT(buffer)
	return p, n, err
}

/*
	publishHandler     func(*packet.Publish) error
	connectHandler     func(*packet.Connect) error
	pingReqHandler     func(*packet.PingReq) error
	pingRespHandler    func(*packet.PingResp) error
	pubAckHandler      func(*packet.PubAck) error
	subAckHandler      func(*packet.SubAck) error
	subscribeHandler   func(*packet.Subscribe) error
	unsubscribeHandler func(*packet.Unsubscribe) error
	disconnectHandler  func(*packet.Disconnect) error
*/
