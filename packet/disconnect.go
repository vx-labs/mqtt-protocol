package packet

func UnmarshalDisconnect(p *Disconnect, buff []byte) (int, error) {
	return 0, nil
}

type disconnectHandler func(*Disconnect) error

func DisconnectDecoder(fn disconnectHandler) func(h *Header, buffer []byte) error {
	return func(h *Header, buffer []byte) error {
		packet := &Disconnect{Header: h}
		_, err := UnmarshalDisconnect(packet, buffer)
		if err != nil {
			return err
		}
		return fn(packet)
	}
}

func EncodeDisconnect(p *Disconnect, buff []byte) (int, error) {
	return 0, nil
}
func (p *Disconnect) Encode(buff []byte) (int, error) {
	return EncodeDisconnect(p, buff)
}
func (p *Disconnect) Length() int {
	return 0
}
