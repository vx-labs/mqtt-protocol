package pb

func decodePingResp(p *MqttPingResp, buff []byte) (int, error) {
	return 0, nil
}

type pingRespHandler func(*MqttPingResp) error

func PingRespDecoder(fn pingRespHandler) func(h *MqttHeader, buffer []byte) error {
	return func(h *MqttHeader, buffer []byte) error {
		packet := &MqttPingResp{Header: h}
		_, err := decodePingResp(packet, buffer)
		if err != nil {
			return err
		}
		return fn(packet)
	}
}
func EncodePingResp(p *MqttPingResp, buff []byte) (int, error) {
	return 0, nil
}
func PingRespLength(p *MqttPingResp) int {
	return 0
}
