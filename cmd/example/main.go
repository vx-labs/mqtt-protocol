package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
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
	meta := map[string]*session{}
	epoller, err := MkEpoll()
	if err != nil {
		panic(err)
	}
	go start(meta, epoller)
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
		go runSession(c, meta, epoller)
	}
}

type session struct {
	keepalive int32
	connect   *packet.Connect
	c         net.Conn
	published int
}

func runSession(c net.Conn, meta map[string]*session, epoller *epoll) {
	pkt, err := decoder.Decode(c, make([]byte, 4))
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

	encoder.New(c).Encode(&packet.ConnAck{
		Header:     p.Header,
		ReturnCode: packet.CONNACK_CONNECTION_ACCEPTED,
	})
	meta[c.RemoteAddr().String()] = &session{keepalive: p.KeepaliveTimer, c: c, published: 0, connect: p}
	if err := epoller.Add(c); err != nil {
		c.Close()
		log.Printf("Failed to add connection %v", err)
	}
}

func processSession(buf []byte, c net.Conn, meta map[string]*session) error {
	session, ok := meta[c.RemoteAddr().String()]
	if !ok {
		return errors.New("session not found")
	}
	enc := encoder.New(c)

	pkt, err := decoder.Decode(c, buf)

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
			enc.PubAck(&packet.PubAck{
				Header:    p.Header,
				MessageId: p.MessageId,
			})
		}
	case *packet.Subscribe:
		c.SetDeadline(
			time.Now().Add(time.Duration(session.keepalive) * time.Second),
		)
		enc.SubAck(&packet.SubAck{
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
		enc.PingResp(&packet.PingResp{
			Header: p.Header,
		})
	default:
		log.Printf("received unknown packet")
	}
	return nil
}

func start(meta map[string]*session, epoller *epoll) {
	buf := make([]byte, 4)
	for {
		connections, err := epoller.Wait()
		if err != nil {
			log.Printf("Failed to epoll wait %v", err)
			continue
		}
		for _, c := range connections {
			err := processSession(buf, c, meta)
			if err != nil {
				if err := epoller.Remove(c); err != nil {
					log.Printf("Failed to remove %v", err)
				}
				delete(meta, c.RemoteAddr().String())
			}
		}
	}
}
