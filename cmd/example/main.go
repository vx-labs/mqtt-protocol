package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	_ "net/http/pprof"

	"github.com/vx-labs/mqtt-protocol/decoder"
	"github.com/vx-labs/mqtt-protocol/encoder"
	"github.com/vx-labs/mqtt-protocol/pb"
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
	encBuffer := make([]byte, 50*1024*1024)
	enc := encoder.New(c)
	keepAlive := int32(30)
	dec := decoder.New(
		decoder.OnConnect(func(p *pb.MqttConnect) error {
			log.Printf("received CONNECT from %s", p.ClientId)
			keepAlive = p.KeepaliveTimer
			c.SetDeadline(
				time.Now().Add(time.Duration(keepAlive) * time.Second),
			)
			enc.ConnAck(&pb.MqttConnAck{
				Header:     p.Header,
				ReturnCode: pb.CONNACK_CONNECTION_ACCEPTED,
			}, encBuffer)
			return nil
		}),
		decoder.OnPublish(func(p *pb.MqttPublish) error {
			c.SetDeadline(
				time.Now().Add(time.Duration(keepAlive) * time.Second),
			)
			if p.Header.Qos == 1 {
				return enc.PubAck(&pb.MqttPubAck{
					Header:    p.Header,
					MessageId: p.MessageId,
				}, encBuffer)
			}
			return nil
		}),
		decoder.OnSubscribe(func(p *pb.MqttSubscribe) error {
			c.SetDeadline(
				time.Now().Add(time.Duration(keepAlive) * time.Second),
			)
			return enc.SubAck(&pb.MqttSubAck{
				Header:    p.Header,
				MessageId: p.MessageId,
			}, encBuffer)
		}),
		decoder.OnUnsubscribe(func(p *pb.MqttUnsubscribe) error { return nil }),
		decoder.OnPubAck(func(*pb.MqttPubAck) error { return nil }),
		decoder.OnPingReq(func(p *pb.MqttPingReq) error {
			c.SetDeadline(
				time.Now().Add(time.Duration(keepAlive) * time.Second),
			)
			return enc.PingResp(&pb.MqttPingResp{
				Header: p.Header,
			}, encBuffer)
		}),
		decoder.OnDisconnect(func(p *pb.MqttDisconnect) error {
			return io.EOF
		}),
	)
	c.SetDeadline(
		time.Now().Add(10 * time.Second),
	)
	decoderCh := make(chan struct{})
	go func() {
		defer close(decoderCh)
		var err error
		for {
			err = dec.Decode(c)
			if err == io.EOF {
				return
			}
			if err != nil {
				log.Printf("decoding failed: %v", err)
				return
			}
		}
	}()

	select {
	case <-decoderCh:
	}
}
