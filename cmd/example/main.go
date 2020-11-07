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

	epoller, err := MkEpoll()
	if err != nil {
		panic(err)
	}
	worker := &worker{
		enc: encoder.New(),
		dec: decoder.New(true),
		meta: &sessions{
			data: make(map[string]*session),
		},
	}
	go worker.start(epoller)
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
		go worker.runSession(c, epoller)
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

func (w *worker) runSession(c net.Conn, epoller *epoll) {
	pkt, err := decoder.Decode(c, make([]byte, 4), true)
	if err != nil {
		return
	}
	if pkt.Type() != packet.CONNECT {
		return
	}
	p := pkt.(*packet.Connect)
	c.SetDeadline(
		time.Now().Add(time.Duration(p.KeepaliveTimer) * 2 * time.Second),
	)
	log.Printf("%s connected", string(p.ClientId))

	encoder.New().Encode(c, &packet.ConnAck{
		Header:     p.Header,
		ReturnCode: packet.CONNACK_CONNECTION_ACCEPTED,
	})
	w.meta.mtx.Lock()
	if err := epoller.Add(c); err != nil {
		c.Close()
		w.meta.mtx.Unlock()
		log.Printf("Failed to add connection %v", err)
		return
	}
	w.meta.data[c.RemoteAddr().String()] = &session{keepalive: p.KeepaliveTimer, c: c, published: 0, connect: p}
	w.meta.mtx.Unlock()
}

type worker struct {
	enc  *encoder.Encoder
	dec  *decoder.Sync
	meta *sessions
}

func (w *worker) processSession(c net.Conn) error {
	w.meta.mtx.Lock()
	session, ok := w.meta.data[c.RemoteAddr().String()]
	w.meta.mtx.Unlock()
	if !ok {
		return errors.New("session not found")
	}

	pkt, err := w.dec.Decode(c)

	if err != nil {
		log.Printf("session lost: %v, %d message published", err, session.published)
		return err
	}
	switch p := pkt.(type) {
	case *packet.Publish:
		c.SetDeadline(
			time.Now().Add(time.Duration(session.keepalive) * time.Second),
		)
		session.published++
		if p.Header.Qos == 1 {
			w.enc.PubAck(c, &packet.PubAck{
				Header:    p.Header,
				MessageId: p.MessageId,
			})
		}
	case *packet.Subscribe:
		c.SetDeadline(
			time.Now().Add(time.Duration(session.keepalive) * time.Second),
		)
		w.enc.SubAck(c, &packet.SubAck{
			Header:    p.Header,
			MessageId: p.MessageId,
		})
	case *packet.Unsubscribe:
	case *packet.Disconnect:
		log.Printf("session closed, %d message published", session.published)
		return errors.New("session closed")
	case *packet.PingReq:
		c.SetDeadline(
			time.Now().Add(time.Duration(session.keepalive) * time.Second),
		)
		w.enc.PingResp(c, &packet.PingResp{
			Header: p.Header,
		})
	default:
		log.Printf("received unknown packet")
	}
	return nil
}

func (w *worker) start(epoller *epoll) {
	for {
		connections, err := epoller.Wait()
		if err != nil {
			log.Printf("Failed to epoll wait %v", err)
			continue
		}
		for _, c := range connections {
			err := w.processSession(c)
			if err != nil {
				if err := epoller.Remove(c); err != nil {
					log.Printf("Failed to remove %v", err)
				}
				w.meta.mtx.Lock()
				delete(w.meta.data, c.RemoteAddr().String())
				w.meta.mtx.Unlock()
			}
		}
	}
}
