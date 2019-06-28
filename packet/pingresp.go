package packet

func decodePingResp(p *PingResp, buff []byte) (int, error) {
	return 0, nil
}

type pingRespHandler func(*PingResp) error

func PingRespDecoder(fn pingRespHandler) func(h *Header, buffer []byte) error {
	return func(h *Header, buffer []byte) error {
		packet := &PingResp{Header: h}
		_, err := decodePingResp(packet, buffer)
		if err != nil {
			return err
		}
		return fn(packet)
	}
}
func EncodePingResp(p *PingResp, buff []byte) (int, error) {
	return 0, nil
}
func PingRespLength(p *PingResp) int {
	return 0
}

func (p *PingResp) Encode(buff []byte) (int, error) {
	return EncodePingResp(p, buff)
}
func (p *PingResp) Length() int {
	return PingRespLength(p)
}
