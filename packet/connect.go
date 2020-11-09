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
func set(b byte, mask byte) byte {
	return b & mask
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

func unmarshalConnect(p *Connect, buff []byte) (int, error) {
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

type connectHandler func(*Connect) error

func ConnectDecoder(fn connectHandler) func(h *Header, buffer []byte) error {
	return func(h *Header, buffer []byte) error {
		packet := &Connect{Header: h}
		_, err := unmarshalConnect(packet, buffer)
		if err != nil {
			return err
		}
		return fn(packet)
	}
}

func EncodeConnect(p *Connect, buff []byte) (int, error) {
	total, err := encodeLP([]byte(supportedProtocolVersions[4]), buff[0:])
	buff[total] = 4
	total++
	var flag byte = 0
	if p.Clean {
		flag = set(CONNECT_FLAG_CLEAN_SESSION, flag)
	}
	if len(p.WillTopic) > 0 {
		flag = set(CONNECT_FLAG_WILL_FLAG, flag)
	}
	qosFlags := byte(p.WillQos) << 3
	flag = set(CONNECT_FLAG_WILL_QOS, qosFlags)
	if p.WillRetain {
		flag = set(CONNECT_FLAG_WILL_RETAIN, flag)
	}
	if len(p.Username) > 0 {
		flag = set(CONNECT_FLAG_USERNAME, flag)
	}
	if len(p.Password) > 0 {
		flag = set(CONNECT_FLAG_PASSWORD, flag)
	}
	buff[total+1] = flag
	total++
	binary.BigEndian.PutUint16(buff[total:], uint16(p.KeepaliveTimer))
	total += 2
	n, err := encodeLP(p.ClientId, buff[total:])
	total += n
	if err != nil {
		return total, err
	}
	if len(p.WillTopic) > 0 {
		n, err := encodeLP(p.WillTopic, buff[total:])
		total += n
		if err != nil {
			return total, err
		}
		n, err = encodeLP(p.WillPayload, buff[total:])
		total += n
		if err != nil {
			return total, err
		}
	}
	if len(p.Username) > 0 {
		n, err = encodeLP(p.Username, buff[total:])
		total += n
		if err != nil {
			return total, err
		}
	}
	if len(p.Password) > 0 {
		n, err = encodeLP(p.Password, buff[total:])
		total += n
		if err != nil {
			return total, err
		}
	}
	return total, err
}
func ConnectLength(p *Connect) int {
	size := 6 + len(supportedProtocolVersions[4]) + 2 + len(p.ClientId)
	if len(p.WillTopic) > 0 {
		size += 2 + len(p.WillTopic) + 2 + len(p.WillPayload)
	}
	if len(p.Username) > 0 {
		size += 2 + len(p.Username)
	}
	if len(p.Password) > 0 {
		size += 2 + len(p.Password)
	}
	return size
}

func (p *Connect) Encode(buff []byte) (int, error) {
	return EncodeConnect(p, buff)
}
func (p *Connect) Length() int {
	return ConnectLength(p)
}
