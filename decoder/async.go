package decoder

import (
	"errors"
	fmt "fmt"
	"io"

	"github.com/google/btree"

	"github.com/vx-labs/mqtt-protocol/packet"
)

type AsyncDecoder struct {
	queue  chan packet.Packet
	done   chan struct{}
	cancel chan struct{}
	err    error
}

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

func Async(r io.Reader, opts ...decoderCreateOp) *AsyncDecoder {
	a := &AsyncDecoder{
		queue:  make(chan packet.Packet, 20),
		done:   make(chan struct{}),
		cancel: make(chan struct{}),
	}
	go func() {
		defer func() {
			close(a.queue)
			close(a.done)
		}()
		for {
			select {
			case <-a.cancel:
				a.err = errors.New("context cancelled")
				return
			default:
				err := a.loop(r)
				if err != nil {
					a.err = err
					return
				}
			}
		}
	}()
	return a
}

func (a *AsyncDecoder) Cancel() {
	close(a.cancel)
}
func (a *AsyncDecoder) Done() <-chan struct{} {
	return a.done
}
func (a *AsyncDecoder) Err() error {
	return a.err
}
func (a *AsyncDecoder) loop(r io.Reader) error {
	pkt, err := decodeEncodedPacket(r)
	if pkt != nil {
		a.queue <- pkt
	}
	return err
}
func (a *AsyncDecoder) Packet() <-chan packet.Packet {
	return a.queue
}

func decodeEncodedPacket(r io.Reader) (packet.Packet, error) {
	h := &packet.Header{}
	packetType, buffer, err := readMessageBuffer(h, r)
	if err != nil {
		return nil, err
	}
	return unmarshalPacket(packetType, h, buffer)
}

func unmarshalPacket(packetType byte, header *packet.Header, buffer []byte) (packet.Packet, error) {
	var p packet.Decoder
	switch packetType {
	case packet.CONNECT:
		p = &packet.Connect{Header: header}
	case packet.SUBSCRIBE:
		p = &packet.Subscribe{Header: header}
	case packet.PUBLISH:
		p = &packet.Publish{Header: header}
	case packet.PINGREQ:
		p = &packet.PingReq{Header: header}
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
		return nil, err
	}
	err := p.UnmarshalMQTT(buffer)
	return p, err
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
