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
		packet.PINGRESP:    packet.PingRespDecoder(d.pingRespHandler),
		packet.DISCONNECT:  packet.DisconnectDecoder(d.disconnectHandler),
		packet.PUBACK:      packet.PubAckDecoder(d.pubAckHandler),
		packet.SUBACK:      packet.SubAckDecoder(d.subAckHandler),
	}
}

func defaultPacketHandler(t string) error {
	return fmt.Errorf("unhandled %s packet", t)
}

func New(opts ...decoderCreateOp) *Decoder {
	d := Decoder{
		publishHandler: func(*packet.Publish) error {
			return defaultPacketHandler("publish")
		},
		connectHandler: func(*packet.Connect) error {
			return defaultPacketHandler("connect")
		},
		pubAckHandler: func(*packet.PubAck) error {
			return defaultPacketHandler("puback")
		},
		subAckHandler: func(*packet.SubAck) error {
			return defaultPacketHandler("suback")
		},
		pingReqHandler: func(*packet.PingReq) error {
			return defaultPacketHandler("pingreq")
		},
		pingRespHandler: func(*packet.PingResp) error {
			return defaultPacketHandler("pingresp")
		},
		subscribeHandler: func(*packet.Subscribe) error {
			return defaultPacketHandler("subscribe")
		},
		unsubscribeHandler: func(*packet.Unsubscribe) error {
			return defaultPacketHandler("unsubcribe")
		},
		disconnectHandler: func(*packet.Disconnect) error {
			return defaultPacketHandler("disconnect")
		},
	}
	for _, op := range opts {
		d = op(d)
	}
	d.packetDecoders = packetDecoders(&d)
	return &d
}

type packetDecoder func(header *packet.Header, buffer []byte) error

type Decoder struct {
	packetDecoders     map[byte]packetDecoder
	publishHandler     func(*packet.Publish) error
	connectHandler     func(*packet.Connect) error
	pingReqHandler     func(*packet.PingReq) error
	pingRespHandler    func(*packet.PingResp) error
	pubAckHandler      func(*packet.PubAck) error
	subAckHandler      func(*packet.SubAck) error
	subscribeHandler   func(*packet.Subscribe) error
	unsubscribeHandler func(*packet.Unsubscribe) error
	disconnectHandler  func(*packet.Disconnect) error
}

func (d *Decoder) Decode(r io.Reader) error {
	h := &packet.Header{}
	packetType, buffer, _, err := readMessageBuffer(h, r)
	if err != nil {
		return err
	}
	p, ok := d.packetDecoders[packetType]
	if !ok {
		return fmt.Errorf("unknown packet type received: %v", packetType)
	}
	return p(h, buffer)
}
