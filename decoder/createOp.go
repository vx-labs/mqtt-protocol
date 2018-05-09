package decoder

import "github.com/vx-labs/mqtt-protocol/packet"

type decoderCreateOp func(d Decoder) Decoder

func OnPublish(fn func(*packet.Publish) error) decoderCreateOp {
	return func(d Decoder) Decoder {
		d.publishHandler = fn
		return d
	}
}

func OnDisconnect(fn func(*packet.Disconnect) error) decoderCreateOp {
	return func(d Decoder) Decoder {
		d.disconnectHandler = fn
		return d
	}
}
func OnConnect(fn func(*packet.Connect) error) decoderCreateOp {
	return func(d Decoder) Decoder {
		d.connectHandler = fn
		return d
	}
}
func OnSubscribe(fn func(*packet.Subscribe) error) decoderCreateOp {
	return func(d Decoder) Decoder {
		d.subscribeHandler = fn
		return d
	}
}
func OnUnsubscribe(fn func(*packet.Unsubscribe) error) decoderCreateOp {
	return func(d Decoder) Decoder {
		d.unsubscribeHandler = fn
		return d
	}
}
func OnPingReq(fn func(*packet.PingReq) error) decoderCreateOp {
	return func(d Decoder) Decoder {
		d.pingReqHandler = fn
		return d
	}
}
func OnPubAck(fn func(*packet.PubAck) error) decoderCreateOp {
	return func(d Decoder) Decoder {
		d.pubAckHandler = fn
		return d
	}
}
