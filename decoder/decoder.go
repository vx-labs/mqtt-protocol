package decoder

import (
	fmt "fmt"
	"io"

	"github.com/vx-labs/mqtt-protocol/packet"
	"github.com/vx-labs/mqtt-protocol/pb"
)

const (
	Default_Buffer_Size = 8192
)

func packetDecoders(d *Decoder) map[byte]packetDecoder {
	return map[byte]packetDecoder{
		packet.PUBLISH:     pb.PublishDecoder(d.publishHandler),
		packet.CONNECT:     pb.ConnectDecoder(d.connectHandler),
		packet.SUBSCRIBE:   pb.SubscribeDecoder(d.subscribeHandler),
		packet.UNSUBSCRIBE: pb.UnsubscribeDecoder(d.unsubscribeHandler),
		packet.PINGREQ:     pb.PingReqDecoder(d.pingReqHandler),
		packet.DISCONNECT:  pb.DisconnectDecoder(d.disconnectHandler),
	}
}

func defaultPacketHandler(t string) error {
	return fmt.Errorf("unhandled %s packet", t)
}

func New(opts ...decoderCreateOp) *Decoder {
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
	d.packetDecoders = packetDecoders(&d)
	return &d
}

type packetDecoder func(header *pb.MqttHeader, buffer []byte) error

type Decoder struct {
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
