package pb

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
	if string(protocolName) != "MQIsdp" {
		return total, fmt.Errorf("unsupported protocol")
	}
	protocolVersion := int(buff[total])
	total++
	if protocolVersion != 3 {
		return total, fmt.Errorf("unsupported mqtt version")
	}
	flags := buff[total]

	p.Clean = cleanSession(flags)
	total++

	keepalive := binary.BigEndian.Uint16(buff[total:])
	total += 2
	p.KeepaliveTimer = int32(keepalive)
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
