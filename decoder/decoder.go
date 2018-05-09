package decoder

import (
	fmt "fmt"
	"io"

	"github.com/vx-labs/mqtt-protocol/packet"
)

func packetDecoders(d *Decoder) map[byte]packetDecoder {
	return map[byte]packetDecoder{
		packet.PUBLISH:     packet.PublishDecoder(d.publishHandler),
		packet.CONNECT:     packet.ConnectDecoder(d.connectHandler),
		packet.SUBSCRIBE:   packet.SubscribeDecoder(d.subscribeHandler),
		packet.UNSUBSCRIBE: packet.UnsubscribeDecoder(d.unsubscribeHandler),
		packet.PINGREQ:     packet.PingReqDecoder(d.pingReqHandler),
		packet.DISCONNECT:  packet.DisconnectDecoder(d.disconnectHandler),
		packet.PUBACK:      packet.PubAckDecoder(d.pubAckHandler),
	}
}

func defaultPacketHandler(t string) error {
	return fmt.Errorf("unhandled %s packet", t)
}

func New(opts ...decoderCreateOp) *Decoder {
	d := Decoder{
		publishHandler: func(*packet.MqttPublish) error {
			return defaultPacketHandler("publish")
		},
		connectHandler: func(*packet.MqttConnect) error {
			return defaultPacketHandler("connect")
		},
		pubAckHandler: func(*packet.MqttPubAck) error {
			return defaultPacketHandler("puback")
		},
		pingReqHandler: func(*packet.MqttPingReq) error {
			return defaultPacketHandler("pingreq")
		},
		subscribeHandler: func(*packet.MqttSubscribe) error {
			return defaultPacketHandler("subscribe")
		},
		unsubscribeHandler: func(*packet.MqttUnsubscribe) error {
			return defaultPacketHandler("unsubcribe")
		},
		disconnectHandler: func(*packet.MqttDisconnect) error {
			return defaultPacketHandler("disconnect")
		},
	}
	for _, op := range opts {
		d = op(d)
	}
	d.packetDecoders = packetDecoders(&d)
	return &d
}

type packetDecoder func(header *packet.MqttHeader, buffer []byte) error

type Decoder struct {
	packetDecoders     map[byte]packetDecoder
	publishHandler     func(*packet.MqttPublish) error
	connectHandler     func(*packet.MqttConnect) error
	pingReqHandler     func(*packet.MqttPingReq) error
	pubAckHandler      func(*packet.MqttPubAck) error
	subscribeHandler   func(*packet.MqttSubscribe) error
	unsubscribeHandler func(*packet.MqttUnsubscribe) error
	disconnectHandler  func(*packet.MqttDisconnect) error
}

func (d *Decoder) Decode(r io.Reader) error {
	h := &packet.MqttHeader{}
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
