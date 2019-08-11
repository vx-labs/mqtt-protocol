package main

import (
	"bytes"
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/vx-labs/mqtt-protocol/decoder"
	"github.com/vx-labs/mqtt-protocol/packet"
)

func logConnect(p *packet.Connect) error {
	fmt.Printf("connect: %v\n", p)
	return nil
}
func logPublish(p *packet.Publish) error {
	fmt.Printf("publish: %v\n", p)
	return nil
}
func logSubscribe(p *packet.Subscribe) error {
	fmt.Printf("subscribe: %v\n", p)
	return nil
}
func logPuback(p *packet.PubAck) error {
	fmt.Printf("puback: %v\n", p)
	return nil
}
func logDisconnect(p *packet.Disconnect) error {
	fmt.Printf("disconnect: %v\n", p)
	return nil
}
func logPingReq(p *packet.PingReq) error {
	fmt.Printf("ping-req: %v\n", p)
	return nil
}
func logPingResp(p *packet.PingResp) error {
	fmt.Printf("ping-resp: %v\n", p)
	return nil
}
func logUnsubscribe(p *packet.Unsubscribe) error {
	fmt.Printf("unsubscribe: %v\n", p)
	return nil
}

func main() {
	rootCmd := &cobra.Command{
		Use:     "mqtt-decoder",
		Example: "mqtt-decoder 64 2 0 20",
		Run: func(cmd *cobra.Command, args []string) {
			encoded := []byte{}
			for _, elt := range args {
				b, err := strconv.ParseUint(elt, 10, 8)
				if err != nil {
					log.Fatal(err)
				}
				encoded = append(encoded, byte(b))
			}
			buf := bytes.NewBuffer(encoded)
			dec := decoder.New(
				decoder.OnConnect(logConnect),
				decoder.OnPublish(logPublish),
				decoder.OnSubscribe(logSubscribe),
				decoder.OnPubAck(logPuback),
				decoder.OnDisconnect(logDisconnect),
				decoder.OnPingReq(logPingReq),
				decoder.OnPingResp(logPingResp),
				decoder.OnUnsubscribe(logUnsubscribe),
			)
			err := dec.Decode(buf)
			if err != nil {
				log.Printf("ERR: decoding failed: %v", err)
			}
		},
	}
	rootCmd.Execute()
}
