package types

import "io"

type TopicSubscription struct {
	Topic MqttString
	Qos   MqttByte
}

func (t *TopicSubscription) Decode(r io.Reader) error {
	err := t.Topic.Decode(r)
	if err != nil {
		return err
	}
	return t.Qos.Decode(r)
}

func (t *TopicSubscription) Length() int {
	return t.Topic.Length() + t.Qos.Length()
}
func (t *TopicSubscription) Encode() ([]byte, error) {
	result := make([]byte, t.Length())
	topic, err := t.Topic.Encode()
	if err != nil {
		return nil, err
	}
	qos, err := t.Qos.Encode()
	if err != nil {
		return nil, err
	}
	result =  append(result, topic...)
	return append(result, qos...), nil
}