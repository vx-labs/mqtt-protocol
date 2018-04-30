// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	pb.proto

It has these top-level messages:
	MqttConnect
	MqttConnack
	MqttDisconnect
	MqttPublish
	MqttPuback
	MqttSubscribe
	MqttSuback
	MqttUnsubscribe
	MqttUnsuback
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MqttConnect struct {
	Id             string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Clean          bool   `protobuf:"varint,2,opt,name=clean" json:"clean,omitempty"`
	ClientId       string `protobuf:"bytes,3,opt,name=clientId" json:"clientId,omitempty"`
	Username       string `protobuf:"bytes,4,opt,name=username" json:"username,omitempty"`
	Password       string `protobuf:"bytes,5,opt,name=password" json:"password,omitempty"`
	WillTopic      string `protobuf:"bytes,6,opt,name=willTopic" json:"willTopic,omitempty"`
	WillPayload    string `protobuf:"bytes,7,opt,name=willPayload" json:"willPayload,omitempty"`
	WillQos        []byte `protobuf:"bytes,8,opt,name=willQos,proto3" json:"willQos,omitempty"`
	WillRetain     bool   `protobuf:"varint,9,opt,name=willRetain" json:"willRetain,omitempty"`
	KeepaliveTimer int32  `protobuf:"varint,10,opt,name=keepaliveTimer" json:"keepaliveTimer,omitempty"`
}

func (m *MqttConnect) Reset()                    { *m = MqttConnect{} }
func (m *MqttConnect) String() string            { return proto.CompactTextString(m) }
func (*MqttConnect) ProtoMessage()               {}
func (*MqttConnect) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *MqttConnect) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MqttConnect) GetClean() bool {
	if m != nil {
		return m.Clean
	}
	return false
}

func (m *MqttConnect) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *MqttConnect) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *MqttConnect) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *MqttConnect) GetWillTopic() string {
	if m != nil {
		return m.WillTopic
	}
	return ""
}

func (m *MqttConnect) GetWillPayload() string {
	if m != nil {
		return m.WillPayload
	}
	return ""
}

func (m *MqttConnect) GetWillQos() []byte {
	if m != nil {
		return m.WillQos
	}
	return nil
}

func (m *MqttConnect) GetWillRetain() bool {
	if m != nil {
		return m.WillRetain
	}
	return false
}

func (m *MqttConnect) GetKeepaliveTimer() int32 {
	if m != nil {
		return m.KeepaliveTimer
	}
	return 0
}

type MqttConnack struct {
	ReturnCode []byte `protobuf:"bytes,1,opt,name=returnCode,proto3" json:"returnCode,omitempty"`
}

func (m *MqttConnack) Reset()                    { *m = MqttConnack{} }
func (m *MqttConnack) String() string            { return proto.CompactTextString(m) }
func (*MqttConnack) ProtoMessage()               {}
func (*MqttConnack) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *MqttConnack) GetReturnCode() []byte {
	if m != nil {
		return m.ReturnCode
	}
	return nil
}

type MqttDisconnect struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *MqttDisconnect) Reset()                    { *m = MqttDisconnect{} }
func (m *MqttDisconnect) String() string            { return proto.CompactTextString(m) }
func (*MqttDisconnect) ProtoMessage()               {}
func (*MqttDisconnect) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *MqttDisconnect) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type MqttPublish struct {
	MessageId int32  `protobuf:"varint,1,opt,name=messageId" json:"messageId,omitempty"`
	Topic     string `protobuf:"bytes,2,opt,name=topic" json:"topic,omitempty"`
	Qos       []byte `protobuf:"bytes,3,opt,name=qos,proto3" json:"qos,omitempty"`
	Payload   string `protobuf:"bytes,4,opt,name=payload" json:"payload,omitempty"`
}

func (m *MqttPublish) Reset()                    { *m = MqttPublish{} }
func (m *MqttPublish) String() string            { return proto.CompactTextString(m) }
func (*MqttPublish) ProtoMessage()               {}
func (*MqttPublish) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *MqttPublish) GetMessageId() int32 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

func (m *MqttPublish) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *MqttPublish) GetQos() []byte {
	if m != nil {
		return m.Qos
	}
	return nil
}

func (m *MqttPublish) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

type MqttPuback struct {
	MessageId int32 `protobuf:"varint,1,opt,name=messageId" json:"messageId,omitempty"`
}

func (m *MqttPuback) Reset()                    { *m = MqttPuback{} }
func (m *MqttPuback) String() string            { return proto.CompactTextString(m) }
func (*MqttPuback) ProtoMessage()               {}
func (*MqttPuback) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *MqttPuback) GetMessageId() int32 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

type MqttSubscribe struct {
	MessageId int32    `protobuf:"varint,1,opt,name=messageId" json:"messageId,omitempty"`
	Topic     []string `protobuf:"bytes,2,rep,name=topic" json:"topic,omitempty"`
	Qos       [][]byte `protobuf:"bytes,3,rep,name=qos,proto3" json:"qos,omitempty"`
}

func (m *MqttSubscribe) Reset()                    { *m = MqttSubscribe{} }
func (m *MqttSubscribe) String() string            { return proto.CompactTextString(m) }
func (*MqttSubscribe) ProtoMessage()               {}
func (*MqttSubscribe) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *MqttSubscribe) GetMessageId() int32 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

func (m *MqttSubscribe) GetTopic() []string {
	if m != nil {
		return m.Topic
	}
	return nil
}

func (m *MqttSubscribe) GetQos() [][]byte {
	if m != nil {
		return m.Qos
	}
	return nil
}

type MqttSuback struct {
	MessageId int32    `protobuf:"varint,1,opt,name=messageId" json:"messageId,omitempty"`
	Qos       [][]byte `protobuf:"bytes,3,rep,name=qos,proto3" json:"qos,omitempty"`
}

func (m *MqttSuback) Reset()                    { *m = MqttSuback{} }
func (m *MqttSuback) String() string            { return proto.CompactTextString(m) }
func (*MqttSuback) ProtoMessage()               {}
func (*MqttSuback) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *MqttSuback) GetMessageId() int32 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

func (m *MqttSuback) GetQos() [][]byte {
	if m != nil {
		return m.Qos
	}
	return nil
}

type MqttUnsubscribe struct {
	MessageId int32    `protobuf:"varint,1,opt,name=messageId" json:"messageId,omitempty"`
	Topic     []string `protobuf:"bytes,2,rep,name=topic" json:"topic,omitempty"`
}

func (m *MqttUnsubscribe) Reset()                    { *m = MqttUnsubscribe{} }
func (m *MqttUnsubscribe) String() string            { return proto.CompactTextString(m) }
func (*MqttUnsubscribe) ProtoMessage()               {}
func (*MqttUnsubscribe) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *MqttUnsubscribe) GetMessageId() int32 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

func (m *MqttUnsubscribe) GetTopic() []string {
	if m != nil {
		return m.Topic
	}
	return nil
}

type MqttUnsuback struct {
	MessageId int32 `protobuf:"varint,1,opt,name=messageId" json:"messageId,omitempty"`
}

func (m *MqttUnsuback) Reset()                    { *m = MqttUnsuback{} }
func (m *MqttUnsuback) String() string            { return proto.CompactTextString(m) }
func (*MqttUnsuback) ProtoMessage()               {}
func (*MqttUnsuback) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *MqttUnsuback) GetMessageId() int32 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

func init() {
	proto.RegisterType((*MqttConnect)(nil), "MqttConnect")
	proto.RegisterType((*MqttConnack)(nil), "MqttConnack")
	proto.RegisterType((*MqttDisconnect)(nil), "MqttDisconnect")
	proto.RegisterType((*MqttPublish)(nil), "MqttPublish")
	proto.RegisterType((*MqttPuback)(nil), "MqttPuback")
	proto.RegisterType((*MqttSubscribe)(nil), "MqttSubscribe")
	proto.RegisterType((*MqttSuback)(nil), "MqttSuback")
	proto.RegisterType((*MqttUnsubscribe)(nil), "MqttUnsubscribe")
	proto.RegisterType((*MqttUnsuback)(nil), "MqttUnsuback")
}

func init() { proto.RegisterFile("pb.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 374 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0x4d, 0x6b, 0xdb, 0x30,
	0x18, 0xc6, 0xf6, 0x9c, 0x38, 0x6f, 0xbc, 0x6c, 0x88, 0x1d, 0xc4, 0x18, 0xc3, 0xe8, 0x30, 0xcc,
	0xd8, 0x76, 0xd9, 0x75, 0xb7, 0x6c, 0x87, 0x1c, 0x06, 0xa9, 0x92, 0xfc, 0x00, 0xd9, 0x16, 0xad,
	0x88, 0x2d, 0x39, 0x96, 0xdc, 0xd0, 0x9f, 0xd4, 0x7f, 0x59, 0x24, 0xc5, 0x49, 0x28, 0xa1, 0x0d,
	0xf4, 0xa6, 0xe7, 0xe3, 0x7d, 0xf5, 0xf8, 0x11, 0x86, 0xa4, 0x2d, 0x7e, 0xb5, 0x9d, 0x32, 0x8a,
	0x3c, 0x86, 0x30, 0xfd, 0xbf, 0x33, 0x66, 0xae, 0xa4, 0xe4, 0xa5, 0x41, 0x33, 0x08, 0x45, 0x85,
	0x83, 0x2c, 0xc8, 0x27, 0x34, 0x14, 0x15, 0xfa, 0x04, 0x71, 0x59, 0x73, 0x26, 0x71, 0x98, 0x05,
	0x79, 0x42, 0x3d, 0x40, 0x9f, 0x21, 0x29, 0x6b, 0xc1, 0xa5, 0x59, 0x54, 0x38, 0x72, 0xde, 0x23,
	0xb6, 0x5a, 0xaf, 0x79, 0x27, 0x59, 0xc3, 0xf1, 0x3b, 0xaf, 0x0d, 0xd8, 0x6a, 0x2d, 0xd3, 0x7a,
	0xaf, 0xba, 0x0a, 0xc7, 0x5e, 0x1b, 0x30, 0xfa, 0x02, 0x93, 0xbd, 0xa8, 0xeb, 0xb5, 0x6a, 0x45,
	0x89, 0x47, 0x4e, 0x3c, 0x11, 0x28, 0x83, 0xa9, 0x05, 0x4b, 0xf6, 0x50, 0x2b, 0x56, 0xe1, 0xb1,
	0xd3, 0xcf, 0x29, 0x84, 0x61, 0x6c, 0xe1, 0x8d, 0xd2, 0x38, 0xc9, 0x82, 0x3c, 0xa5, 0x03, 0x44,
	0x5f, 0x01, 0xec, 0x91, 0x72, 0xc3, 0x84, 0xc4, 0x13, 0xf7, 0x21, 0x67, 0x0c, 0xfa, 0x06, 0xb3,
	0x2d, 0xe7, 0x2d, 0xab, 0xc5, 0x3d, 0x5f, 0x8b, 0x86, 0x77, 0x18, 0xb2, 0x20, 0x8f, 0xe9, 0x33,
	0x96, 0xfc, 0x3c, 0x55, 0xc5, 0xca, 0xad, 0x5d, 0xdb, 0x71, 0xd3, 0x77, 0x72, 0xae, 0x2a, 0xee,
	0x2a, 0x4b, 0xe9, 0x19, 0x43, 0x32, 0x98, 0x59, 0xfb, 0x5f, 0xa1, 0xcb, 0xcb, 0xe5, 0x92, 0xc6,
	0x2f, 0x5c, 0xf6, 0x45, 0x2d, 0xf4, 0x9d, 0x6d, 0xa0, 0xe1, 0x5a, 0xb3, 0x5b, 0xbe, 0xf0, 0xae,
	0x98, 0x9e, 0x08, 0xfb, 0x12, 0xc6, 0x75, 0x13, 0xba, 0x79, 0x0f, 0xd0, 0x47, 0x88, 0x76, 0x4a,
	0xbb, 0x47, 0x48, 0xa9, 0x3d, 0xda, 0x1e, 0xda, 0x43, 0x4b, 0xbe, 0xfe, 0x01, 0x92, 0xef, 0x00,
	0x87, 0xeb, 0x6c, 0xfc, 0x17, 0x6f, 0x23, 0x1b, 0x78, 0x6f, 0xbd, 0xab, 0xbe, 0xd0, 0x65, 0x27,
	0x0a, 0x7e, 0x7d, 0xb8, 0xe8, 0x42, 0xb8, 0xe8, 0x10, 0x8e, 0xfc, 0xf1, 0x11, 0x56, 0x57, 0x44,
	0xb8, 0x30, 0xfd, 0x0f, 0x3e, 0xd8, 0xe9, 0x8d, 0xd4, 0x6f, 0x89, 0x45, 0x7e, 0x40, 0x7a, 0x5c,
	0xf3, 0x6a, 0x8c, 0x62, 0xe4, 0x7e, 0x94, 0xdf, 0x4f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd5, 0x2c,
	0x83, 0xc9, 0x34, 0x03, 0x00, 0x00,
}