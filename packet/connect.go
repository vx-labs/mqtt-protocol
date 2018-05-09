package packet

import (
	"encoding/binary"
	"fmt"
)

const (
	_ byte = 1 << iota
	CONNECT_FLAG_CLEAN_SESSION
	CONNECT_FLAG_WILL_FLAG
	CONNECT_FLAG_WILL_QOS byte = 3 << iota
	_
	CONNECT_FLAG_WILL_RETAIN byte = 1 << iota
	CONNECT_FLAG_PASSWORD
	CONNECT_FLAG_USERNAME
)

var supportedProtocolVersions = map[int]string{
	3: "MQIsdp",
	4: "MQTT",
}

const DefaultKeepalive int32 = 30

func isSet(b byte, mask byte) bool {
	return b&mask == mask
}
func username(b byte) bool {
	return isSet(b, CONNECT_FLAG_USERNAME)
}
func password(b byte) bool {
	return isSet(b, CONNECT_FLAG_PASSWORD)
}
func cleanSession(b byte) bool {
	return isSet(b, CONNECT_FLAG_CLEAN_SESSION)
}
func will(b byte) bool {
	return isSet(b, CONNECT_FLAG_WILL_FLAG)
}
func willRetain(b byte) bool {
	return isSet(b, CONNECT_FLAG_WILL_RETAIN)
}
func willQoS(b byte) byte {
	return (b & CONNECT_FLAG_WILL_QOS) >> 3
}

func decodeConnect(p *MqttConnect, buff []byte) (int, error) {
	total := 0
	protocolName, n, err := decodeLP(buff)
	if err != nil {
		return n, err
	}
	total += n
	protocolVersion := int(buff[total])
	total++

	if name, ok := supportedProtocolVersions[protocolVersion]; !ok {
		return total, fmt.Errorf("unsupported mqtt version")
		if name != string(protocolName) {
			return total, fmt.Errorf("unsupported protocol name")
		}
	}
	flags := buff[total]

	p.Clean = cleanSession(flags)
	total++

	keepalive := binary.BigEndian.Uint16(buff[total:])
	total += 2
	if keepalive > 0 {
		p.KeepaliveTimer = int32(keepalive)
	} else {
		p.KeepaliveTimer = DefaultKeepalive
	}
	clientId, n, err := decodeLP(buff[total:])
	total += n
	if err != nil {
		return total, err
	}
	p.ClientId = clientId

	if will(flags) {
		p.WillRetain = willRetain(flags)
		p.WillQos = int32(willQoS(flags))
		willTopic, n, err := decodeLP(buff[total:])
		total += n
		if err != nil {
			return total, err
		}
		p.WillTopic = willTopic
		willPayload, n, err := decodeLP(buff[total:])
		total += n
		if err != nil {
			return total, err
		}
		p.WillPayload = willPayload
	}

	if username(flags) {
		username, n, err := decodeLP(buff[total:])
		total += n
		if err != nil {
			return total, err
		}
		p.Username = username
	}
	if password(flags) {
		password, n, err := decodeLP(buff[total:])
		total += n
		if err != nil {
			return total, err
		}
		p.Password = password
	}
	return total, nil
}

type connectHandler func(*MqttConnect) error

func ConnectDecoder(fn connectHandler) func(h *MqttHeader, buffer []byte) error {
	return func(h *MqttHeader, buffer []byte) error {
		packet := &MqttConnect{Header: h}
		_, err := decodeConnect(packet, buffer)
		if err != nil {
			return err
		}
		return fn(packet)
	}
}
