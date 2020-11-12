package decoder

import (
	"errors"
	fmt "fmt"
	"io"

	"github.com/google/btree"

	"github.com/vx-labs/mqtt-protocol/packet"
)

type AsyncDecoder struct {
	stats     StatRecorder
	queue     chan packet.Packet
	done      chan struct{}
	cancel    chan struct{}
	err       error
	headerBuf []byte
	msgBuf    []byte
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

type asyncDecoderCreateOp func(*AsyncDecoder)

func WithStatRecorder(recorder StatRecorder) asyncDecoderCreateOp {
	return func(a *AsyncDecoder) {
		a.stats = recorder
	}
}

func Async(r io.Reader, opts ...asyncDecoderCreateOp) *AsyncDecoder {
	a := &AsyncDecoder{
		stats:     &noopStatRecorder{},
		queue:     make(chan packet.Packet, 20),
		done:      make(chan struct{}),
		cancel:    make(chan struct{}),
		headerBuf: make([]byte, 4),
	}
	for _, opt := range opts {
		opt(a)
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
	pkt, n, err := decodeEncodedPacket(a.headerBuf, r)
	a.stats.Add(float64(n))
	if pkt != nil {
		select {
		case a.queue <- pkt:
		case <-a.cancel:
			return errors.New("context cancelled")
		}
	}
	return err
}
func (a *AsyncDecoder) Packet() <-chan packet.Packet {
	return a.queue
}

func decodeEncodedPacket(headerBuf []byte, r io.Reader) (packet.Packet, int, error) {
	if len(headerBuf) != 4 {
		return nil, 0, errors.New("invalid header buffer size")
	}
	h := &packet.Header{}
	packetType, buffer, count, err := readMessageBuffer(h, headerBuf, r)
	if err != nil {
		return nil, count, err
	}
	pkt, pktRead, err := unmarshalPacket(packetType, h, buffer)
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
