package pb

func decodeDisconnect(p *MqttDisconnect, buff []byte) (int, error) {
	return 0, nil
}

type disconnectHandler func(*MqttDisconnect) error

func DisconnectDecoder(fn disconnectHandler) func(h *MqttHeader, buffer []byte) error {
	return func(h *MqttHeader, buffer []byte) error {
		packet := &MqttDisconnect{Header: h}
		_, err := decodeDisconnect(packet, buffer)
		if err != nil {
			return err
		}
		return fn(packet)
	}
}

func DisconnectLength(p *MqttDisconnect) int {
	return 0
}
