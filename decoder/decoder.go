package decoder

import (
	fmt "fmt"
	"io"

	"github.com/vx-labs/mqtt-protocol/pb"
)

const (
	Default_Buffer_Size = 8192
)

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

func packetDecoders(d *Decoder) map[byte]packetDecoder {
	return map[byte]packetDecoder{
		PUBLISH:     pb.PublishDecoder(d.publishHandler),
		CONNECT:     pb.ConnectDecoder(d.connectHandler),
		SUBSCRIBE:   pb.SubscribeDecoder(d.subscribeHandler),
		UNSUBSCRIBE: pb.UnsubscribeDecoder(d.unsubscribeHandler),
		PINGREQ:     pb.PingReqDecoder(d.pingReqHandler),
		DISCONNECT:  pb.DisconnectDecoder(d.disconnectHandler),
	}
}

func defaultPacketHandler(t string) error {
	return fmt.Errorf("unhandled %s packet", t)
}

func NewDecoder(opts ...decoderCreateOp) *Decoder {
	d := Decoder{
		publishHandler: func(*pb.MqttPublish) error {
			return defaultPacketHandler("publish")
		},
		connectHandler: func(*pb.MqttConnect) error {
			return defaultPacketHandler("connect")
		},
		pubAckHandler: func(*pb.MqttPubAck) error {
			return defaultPacketHandler("puback")
		},
		pingReqHandler: func(*pb.MqttPingReq) error {
			return defaultPacketHandler("pingreq")
		},
		subscribeHandler: func(*pb.MqttSubscribe) error {
			return defaultPacketHandler("subscribe")
		},
		unsubscribeHandler: func(*pb.MqttUnsubscribe) error {
			return defaultPacketHandler("unsubcribe")
		},
		disconnectHandler: func(*pb.MqttDisconnect) error {
			return defaultPacketHandler("disconnect")
		},
	}
	for _, op := range opts {
		d = op(d)
	}
	if d.buffer == nil {
		d.buffer = make([]byte, Default_Buffer_Size)
	}
	d.packetDecoders = packetDecoders(&d)
	return &d
}

type packetDecoder func(header *pb.MqttHeader, buffer []byte) error

type Decoder struct {
	buffer             []byte
	packetDecoders     map[byte]packetDecoder
	publishHandler     func(*pb.MqttPublish) error
	connectHandler     func(*pb.MqttConnect) error
	pingReqHandler     func(*pb.MqttPingReq) error
	pubAckHandler      func(*pb.MqttPubAck) error
	subscribeHandler   func(*pb.MqttSubscribe) error
	unsubscribeHandler func(*pb.MqttUnsubscribe) error
	disconnectHandler  func(*pb.MqttDisconnect) error
}

func (d *Decoder) Decode(r io.Reader) error {
	h := &pb.MqttHeader{}
	packetType, buffer, err := d.readMessageBuffer(h, r)
	if err != nil {
		return err
	}
	p, ok := d.packetDecoders[packetType]
	if !ok {
		return fmt.Errorf("unknown packet type received: %v", packetType)
	}
	return p(h, buffer)
}
