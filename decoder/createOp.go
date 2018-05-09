package decoder

import "github.com/vx-labs/mqtt-protocol/packet"

type decoderCreateOp func(d Decoder) Decoder

func OnPublish(fn func(*packet.MqttPublish) error) decoderCreateOp {
	return func(d Decoder) Decoder {
		d.publishHandler = fn
		return d
	}
}

func OnDisconnect(fn func(*packet.MqttDisconnect) error) decoderCreateOp {
	return func(d Decoder) Decoder {
		d.disconnectHandler = fn
		return d
	}
}
func OnConnect(fn func(*packet.MqttConnect) error) decoderCreateOp {
	return func(d Decoder) Decoder {
		d.connectHandler = fn
		return d
	}
}
func OnSubscribe(fn func(*packet.MqttSubscribe) error) decoderCreateOp {
	return func(d Decoder) Decoder {
		d.subscribeHandler = fn
		return d
	}
}
func OnUnsubscribe(fn func(*packet.MqttUnsubscribe) error) decoderCreateOp {
	return func(d Decoder) Decoder {
		d.unsubscribeHandler = fn
		return d
	}
}
func OnPingReq(fn func(*packet.MqttPingReq) error) decoderCreateOp {
	return func(d Decoder) Decoder {
		d.pingReqHandler = fn
		return d
	}
}
func OnPubAck(fn func(*packet.MqttPubAck) error) decoderCreateOp {
	return func(d Decoder) Decoder {
		d.pubAckHandler = fn
		return d
	}
}
