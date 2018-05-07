package pb

import "encoding/binary"

func decodePubAck(p *MqttPubAck, buff []byte) (int, error) {
	p.MessageId = int32(binary.BigEndian.Uint16(buff))
	return 2, nil
}
