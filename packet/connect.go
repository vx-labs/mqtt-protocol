package packet

import (
	"fmt"
	"github.com/vx-labs/mqtt-protocol/types"
	"io"
)

func NewConnect() *Connect {
	return &Connect{}
}

func (p *Connect) WithFixedHeader(h Header) *Connect{
	c := *p
	c.header = h
	return &c
}

type Connect struct {
	header                Header
	protocolName          types.MqttString
	protocolVersionNumber types.MqttByte
	connectFlags          types.MqttByte
	keepAliveTimer        types.MqttInt
	clientId              types.MqttString
	willTopic             types.MqttString
	willMessage           types.MqttString
	userName              types.MqttString
	password              types.MqttString
}

func (c *Connect) KeepAliveTimer() int {
	return c.keepAliveTimer.Int()
}
func (c *Connect) ClientId() string {
	return c.clientId.String()
}
func (c *Connect) WillTopic() string {
	return c.willTopic.String()
}

func (c *Connect) WillMessageString() string {
	return c.willMessage.String()
}
func (c *Connect) WillMessage() []byte {
	return c.willMessage.Bytes()
}
func (c *Connect) UserName() string {
	return c.userName.String()
}

func (c *Connect) Password() string {
	return c.password.String()
}

type ConnectFlags struct {
	mask byte
}

func (c *ConnectFlags) isSet(mask byte) bool {
	return c.mask&mask == mask
}
func (c *ConnectFlags) Username() bool {
	return c.isSet(CONNECT_FLAG_USERNAME)
}
func (c *ConnectFlags) Password() bool {
	return c.isSet(CONNECT_FLAG_PASSWORD)
}
func (c *ConnectFlags) CleanSession() bool {
	return c.isSet(CONNECT_FLAG_CLEAN_SESSION)
}
func (c *ConnectFlags) Will() bool {
	return c.isSet(CONNECT_FLAG_WILL_FLAG)
}
func (c *ConnectFlags) WillRetain() bool {
	return c.isSet(CONNECT_FLAG_WILL_RETAIN)
}
func (c *ConnectFlags) WillQoS() byte {
	return (c.mask & CONNECT_FLAG_WILL_QOS) >> 3
}

const (
	_                          byte = 1 << iota
	CONNECT_FLAG_CLEAN_SESSION
	CONNECT_FLAG_WILL_FLAG
	CONNECT_FLAG_WILL_QOS      byte = 3 << iota
	_
	CONNECT_FLAG_WILL_RETAIN   byte = 1 << iota
	CONNECT_FLAG_PASSWORD
	CONNECT_FLAG_USERNAME
)

func (h *Connect) FixedHeader(l int) types.Field {
	return h.header.WithRemainingLength(l)
}
func (h *Connect) VariableHeaders() []types.Field {
	return []types.Field{
		&h.protocolName,
		&h.protocolVersionNumber,
		&h.connectFlags,
		&h.keepAliveTimer,
	}
}
func (h *Connect) Type() byte{
	return CONNECT
}
func (h *Connect) String() string{
	return "connect"
}
func (h *Connect) Payload() []types.Field {
	p := []types.Field{
		&h.clientId,
	}
	if h.Flags().Will() {
		p = append(p, &h.willTopic, &h.willMessage)
	}
	if h.Flags().Username() {
		p = append(p, &h.userName)
	}
	if h.Flags().Password() {
		p = append(p, &h.password)
	}
	return p
}

func (h *Connect) Decode(r io.Reader) error {
	if h.header.remLength < 3 {
		return fmt.Errorf("malformed packet: remaining length is too low")
	}
	for _, f := range h.VariableHeaders() {
		err := f.Decode(r)
		if err != nil {
			return err
		}
	}
	for _, f := range h.Payload() {
		err := f.Decode(r)
		if err != nil {
			return err
		}
	}
	return nil
}

func (h *Connect) Flags() *ConnectFlags {
	return &ConnectFlags{mask: h.connectFlags.Byte()}
}

func (c *Connect) Packet() types.Packet {
	return c
}
