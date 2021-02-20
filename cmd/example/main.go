package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	_ "net/http/pprof"

	"github.com/vx-labs/mqtt-protocol/decoder"
	"github.com/vx-labs/mqtt-protocol/encoder"
	"github.com/vx-labs/mqtt-protocol/packet"
)

func main() {
	port := 1883
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		log.Println("profiling is running on :8080")
		log.Println(http.ListenAndServe(":8080", nil))
	}()
	log.Printf("server is running on :%d", port)
	acceptLoop(l)
}

func acceptLoop(l net.Listener) {

	worker := &worker{
		enc: encoder.New(),
		dec: decoder.New(4096),
	}
	var tempDelay time.Duration
	for {
		c, err := l.Accept()
		if err != nil {
			if err.Error() == fmt.Sprintf("accept tcp %v: use of closed network connection", l.Addr()) {
				return
			}
			if ne, ok := err.(net.Error); ok && ne.Temporary() {
				if tempDelay == 0 {
					tempDelay = 5 * time.Millisecond
				} else {
					tempDelay *= 2
				}
				if max := 1 * time.Second; tempDelay > max {
					tempDelay = max
				}
				log.Printf("accept error: %v; retrying in %v", err, tempDelay)
				time.Sleep(tempDelay)
				continue
			}
			log.Printf("connection handling failed: %v", err)
			l.Close()
			return
		}
		go worker.runSession(c)
	}
}

type sessions struct {
	data map[string]*session
	mtx  sync.RWMutex
}

type session struct {
	keepalive int32
	connect   *packet.Connect
	c         net.Conn
	published int
}

func (w *worker) runSession(c net.Conn) {
	defer c.Close()
	pkt, err := w.dec.Decode(c)
	if err != nil {
		return
	}
	if pkt.Type() != packet.CONNECT {
		return
	}
	p := pkt.(*packet.Connect)
	log.Printf("%s connected", string(p.ClientId))

	encoder.New().Encode(c, &packet.ConnAck{
		Header:     p.Header,
		ReturnCode: packet.CONNACK_CONNECTION_ACCEPTED,
	})
	session := &session{keepalive: p.KeepaliveTimer, c: c, published: 0, connect: p}

	for {
		c.SetDeadline(
			time.Now().Add(time.Duration(p.KeepaliveTimer) * 2 * time.Second),
		)
		err := w.processSession(session, c)
		if err != nil {
			return
		}
	}
}

type worker struct {
	enc *encoder.Encoder
	dec *decoder.Sync
}

func (w *worker) processSession(session *session, c net.Conn) error {
	pkt, err := w.dec.Decode(c)
	if err != nil {
		log.Printf("session lost: %v, %d message published", err, session.published)
		return err
	}
	c.SetDeadline(
		time.Now().Add(time.Duration(session.keepalive) * time.Second),
	)
	switch p := pkt.(type) {
	case *packet.PubRel:
		w.enc.Encode(c, &packet.PubComp{
			Header:    p.Header,
			MessageId: p.MessageId,
		})
	case *packet.Publish:
		session.published++
		if p.Header.Qos == 1 {
			w.enc.PubAck(c, &packet.PubAck{
				Header:    p.Header,
				MessageId: p.MessageId,
			})
		}
		if p.Header.Qos == 2 {
			w.enc.Encode(c, &packet.PubRec{
				Header:    p.Header,
				MessageId: p.MessageId,
			})
		}
	case *packet.Subscribe:
		w.enc.SubAck(c, &packet.SubAck{
			Header:    p.Header,
			MessageId: p.MessageId,
		})
	case *packet.Unsubscribe:
	case *packet.Disconnect:
		log.Printf("session closed, %d message published", session.published)
		return errors.New("session closed")
	case *packet.PingReq:
		w.enc.PingResp(c, &packet.PingResp{
			Header: p.Header,
		})
	default:
		log.Printf("received unknown packet")
	}
	return nil
}
