package main

import (
	"log"
	"net"
	"time"

	"github.com/vx-labs/mqtt-protocol/encoder"
	"github.com/vx-labs/mqtt-protocol/packet"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:1883")
	if err != nil {
		panic(err)
	}
	connect := packet.Connect{Header: &packet.Header{}, ClientId: []byte("mqtt-benchmarker"), KeepaliveTimer: 3600}
	connectBuf := make([]byte, connect.Length()+4)
	c, err := encoder.EncodeHeader(packet.CONNECT, &packet.Header{}, connect.Length(), connectBuf)
	if err != nil {
		panic(err)
	}
	_, err = connect.Encode(connectBuf[c : c+connect.Length()])
	if err != nil {
		panic(err)
	}
	_, err = conn.Write(connectBuf[:c+connect.Length()])
	if err != nil {
		panic(err)
	}

	publish := packet.Publish{Header: &packet.Header{}, MessageId: 1, Payload: []byte("test"), Topic: []byte("test")}
	publishBuf := make([]byte, publish.Length()+4)
	c, err = encoder.EncodeHeader(packet.PUBLISH, &packet.Header{}, publish.Length(), publishBuf)
	if err != nil {
		panic(err)
	}
	_, err = publish.Encode(publishBuf[c : c+publish.Length()])
	if err != nil {
		panic(err)
	}
	start := time.Now()
	count := 1000000
	for i := 0; i < count; i++ {
		_, err := conn.Write(publishBuf[:c+publish.Length()])
		if err != nil {
			break
		}
	}
	duration := time.Since(start)
	conn.Close()
	log.Printf("written %d messages in %s", count, duration.String())
	log.Printf("rate: %f messages per second", float64(count)/duration.Seconds())
}
