package packet

func decodeDisconnect(p *Disconnect, buff []byte) (int, error) {
	return 0, nil
}

type disconnectHandler func(*Disconnect) error

func DisconnectDecoder(fn disconnectHandler) func(h *Header, buffer []byte) error {
	return func(h *Header, buffer []byte) error {
		packet := &Disconnect{Header: h}
		_, err := decodeDisconnect(packet, buffer)
		if err != nil {
			return err
		}
		return fn(packet)
	}
}

func DisconnectLength(p *Disconnect) int {
	return 0
}
