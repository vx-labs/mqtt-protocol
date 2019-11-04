package packet

func UnmarshalPingReq(p *PingReq, buff []byte) (int, error) {
	return 0, nil
}

type pingReqHandler func(*PingReq) error

func PingReqDecoder(fn pingReqHandler) func(h *Header, buffer []byte) error {
	return func(h *Header, buffer []byte) error {
		packet := &PingReq{Header: h}
		_, err := UnmarshalPingReq(packet, buffer)
		if err != nil {
			return err
		}
		return fn(packet)
	}
}
