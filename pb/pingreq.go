package pb

func decodePingReq(p *MqttPingReq, buff []byte) (int, error) {
	return 0, nil
}

type pingReqHandler func(*MqttPingReq) error

func PingReqDecoder(fn pingReqHandler) func(h *MqttHeader, buffer []byte) error {
	return func(h *MqttHeader, buffer []byte) error {
		packet := &MqttPingReq{Header: h}
		_, err := decodePingReq(packet, buffer)
		if err != nil {
			return err
		}
		return fn(packet)
	}
}
