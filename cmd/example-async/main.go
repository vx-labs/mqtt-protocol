package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
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
	if os.Getenv("API_ENABLE_PROFILING") == "true" {
		go func() {
			log.Println("profiling is running on :8080")
			log.Println(http.ListenAndServe(":8080", nil))
		}()
	}
	log.Printf("server is running on :%d", port)
	acceptLoop(l)
}

func acceptLoop(l net.Listener) {
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
		go runSession(c)
	}
}

func runSession(c net.Conn) {
	defer c.Close()
	enc := encoder.New(c)
	keepAlive := int32(30)
	dec := decoder.Async(bufio.NewReader(c))
	defer dec.Cancel()
	c.SetDeadline(
		time.Now().Add(10 * time.Second),
	)
	published := 0
	for pkt := range dec.Packet() {
		switch p := pkt.(type) {
		case *packet.Connect:
			log.Printf("%s connected", p.ClientId)
			keepAlive = p.KeepaliveTimer
			c.SetDeadline(
				time.Now().Add(time.Duration(keepAlive) * time.Second),
			)
			enc.ConnAck(&packet.ConnAck{
				Header:     p.Header,
				ReturnCode: packet.CONNACK_CONNECTION_ACCEPTED,
			})
		case *packet.Publish:
			c.SetDeadline(
				time.Now().Add(time.Duration(keepAlive) * time.Second),
			)
			published++
			if p.Header.Qos == 1 {
				enc.PubAck(&packet.PubAck{
					Header:    p.Header,
					MessageId: p.MessageId,
				})
			}
		case *packet.Subscribe:
			c.SetDeadline(
				time.Now().Add(time.Duration(keepAlive) * time.Second),
			)
			enc.SubAck(&packet.SubAck{
				Header:    p.Header,
				MessageId: p.MessageId,
			})
		case *packet.Unsubscribe:
		case *packet.Disconnect:
			log.Printf("session closed, %d message published", published)
			return
		case *packet.PingReq:
			c.SetDeadline(
				time.Now().Add(time.Duration(keepAlive) * time.Second),
			)
			enc.PingResp(&packet.PingResp{
				Header: p.Header,
			})
		default:
			log.Printf("received unknown packet")
		}
	}
	log.Printf("session lost: %v, %d message published", dec.Err(), published)
}
