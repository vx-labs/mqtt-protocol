package decoder

import "github.com/vx-labs/mqtt-protocol/pb"

type decoderCreateOp func(d Decoder) Decoder

func OnPublish(fn func(*pb.MqttPublish) error) decoderCreateOp {
	return func(d Decoder) Decoder {
		d.publishHandler = fn
		return d
	}
}

func OnConnect(fn func(*pb.MqttConnect) error) decoderCreateOp {
	return func(d Decoder) Decoder {
		d.connectHandler = fn
		return d
	}
}
func OnSubscribe(fn func(*pb.MqttSubscribe) error) decoderCreateOp {
	return func(d Decoder) Decoder {
		d.subscribeHandler = fn
		return d
	}
}
func OnUnsubscribe(fn func(*pb.MqttUnsubscribe) error) decoderCreateOp {
	return func(d Decoder) Decoder {
		d.unsubscribeHandler = fn
		return d
	}
}
func OnPingReq(fn func(*pb.MqttPingReq) error) decoderCreateOp {
	return func(d Decoder) Decoder {
		d.pingReqHandler = fn
		return d
	}
}
func OnPubAck(fn func(*pb.MqttPubAck) error) decoderCreateOp {
	return func(d Decoder) Decoder {
		d.pubAckHandler = fn
		return d
	}
}
