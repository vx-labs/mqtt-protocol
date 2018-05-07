// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	pb.proto

It has these top-level messages:
	MqttHeader
	MqttConnect
	MqttConnack
	MqttDisconnect
	MqttPublish
	MqttPuback
	MqttSubscribe
	MqttSuback
	MqttUnsubscribe
	MqttUnsuback
	MqttPingReq
	MqttPingResp
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

type MqttHeader struct {
	Dup    bool  `protobuf:"varint,1,opt,name=dup" json:"dup,omitempty"`
	Qos    int32 `protobuf:"varint,2,opt,name=qos" json:"qos,omitempty"`
	Retain bool  `protobuf:"varint,3,opt,name=retain" json:"retain,omitempty"`
}

func (m *MqttHeader) Reset()                    { *m = MqttHeader{} }
func (m *MqttHeader) String() string            { return proto.CompactTextString(m) }
func (*MqttHeader) ProtoMessage()               {}
func (*MqttHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *MqttHeader) GetDup() bool {
	if m != nil {
		return m.Dup
	}
	return false
}

func (m *MqttHeader) GetQos() int32 {
	if m != nil {
		return m.Qos
	}
	return 0
}

func (m *MqttHeader) GetRetain() bool {
	if m != nil {
		return m.Retain
	}
	return false
}

type MqttConnect struct {
	Header         *MqttHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Id             string      `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	Clean          bool        `protobuf:"varint,3,opt,name=clean" json:"clean,omitempty"`
	ClientId       []byte      `protobuf:"bytes,4,opt,name=clientId,proto3" json:"clientId,omitempty"`
	Username       []byte      `protobuf:"bytes,5,opt,name=username,proto3" json:"username,omitempty"`
	Password       []byte      `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty"`
	WillTopic      []byte      `protobuf:"bytes,7,opt,name=willTopic,proto3" json:"willTopic,omitempty"`
	WillPayload    []byte      `protobuf:"bytes,8,opt,name=willPayload,proto3" json:"willPayload,omitempty"`
	WillQos        int32       `protobuf:"varint,9,opt,name=willQos" json:"willQos,omitempty"`
	WillRetain     bool        `protobuf:"varint,10,opt,name=willRetain" json:"willRetain,omitempty"`
	KeepaliveTimer int32       `protobuf:"varint,11,opt,name=keepaliveTimer" json:"keepaliveTimer,omitempty"`
}

func (m *MqttConnect) Reset()                    { *m = MqttConnect{} }
func (m *MqttConnect) String() string            { return proto.CompactTextString(m) }
func (*MqttConnect) ProtoMessage()               {}
func (*MqttConnect) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *MqttConnect) GetHeader() *MqttHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

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

func (m *MqttConnect) GetClientId() []byte {
	if m != nil {
		return m.ClientId
	}
	return nil
}

func (m *MqttConnect) GetUsername() []byte {
	if m != nil {
		return m.Username
	}
	return nil
}

func (m *MqttConnect) GetPassword() []byte {
	if m != nil {
		return m.Password
	}
	return nil
}

func (m *MqttConnect) GetWillTopic() []byte {
	if m != nil {
		return m.WillTopic
	}
	return nil
}

func (m *MqttConnect) GetWillPayload() []byte {
	if m != nil {
		return m.WillPayload
	}
	return nil
}

func (m *MqttConnect) GetWillQos() int32 {
	if m != nil {
		return m.WillQos
	}
	return 0
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
	Header     *MqttHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	ReturnCode []byte      `protobuf:"bytes,2,opt,name=returnCode,proto3" json:"returnCode,omitempty"`
}

func (m *MqttConnack) Reset()                    { *m = MqttConnack{} }
func (m *MqttConnack) String() string            { return proto.CompactTextString(m) }
func (*MqttConnack) ProtoMessage()               {}
func (*MqttConnack) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *MqttConnack) GetHeader() *MqttHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *MqttConnack) GetReturnCode() []byte {
	if m != nil {
		return m.ReturnCode
	}
	return nil
}

type MqttDisconnect struct {
	Header *MqttHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Id     string      `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
}

func (m *MqttDisconnect) Reset()                    { *m = MqttDisconnect{} }
func (m *MqttDisconnect) String() string            { return proto.CompactTextString(m) }
func (*MqttDisconnect) ProtoMessage()               {}
func (*MqttDisconnect) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *MqttDisconnect) GetHeader() *MqttHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *MqttDisconnect) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type MqttPublish struct {
	Header    *MqttHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	MessageId int32       `protobuf:"varint,2,opt,name=messageId" json:"messageId,omitempty"`
	Topic     []byte      `protobuf:"bytes,3,opt,name=topic,proto3" json:"topic,omitempty"`
	Payload   []byte      `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *MqttPublish) Reset()                    { *m = MqttPublish{} }
func (m *MqttPublish) String() string            { return proto.CompactTextString(m) }
func (*MqttPublish) ProtoMessage()               {}
func (*MqttPublish) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *MqttPublish) GetHeader() *MqttHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *MqttPublish) GetMessageId() int32 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

func (m *MqttPublish) GetTopic() []byte {
	if m != nil {
		return m.Topic
	}
	return nil
}

func (m *MqttPublish) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type MqttPuback struct {
	Header    *MqttHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	MessageId int32       `protobuf:"varint,2,opt,name=messageId" json:"messageId,omitempty"`
}

func (m *MqttPuback) Reset()                    { *m = MqttPuback{} }
func (m *MqttPuback) String() string            { return proto.CompactTextString(m) }
func (*MqttPuback) ProtoMessage()               {}
func (*MqttPuback) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *MqttPuback) GetHeader() *MqttHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *MqttPuback) GetMessageId() int32 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

type MqttSubscribe struct {
	Header    *MqttHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	MessageId int32       `protobuf:"varint,2,opt,name=messageId" json:"messageId,omitempty"`
	Topic     [][]byte    `protobuf:"bytes,3,rep,name=topic,proto3" json:"topic,omitempty"`
	Qos       []int32     `protobuf:"varint,4,rep,packed,name=qos" json:"qos,omitempty"`
}

func (m *MqttSubscribe) Reset()                    { *m = MqttSubscribe{} }
func (m *MqttSubscribe) String() string            { return proto.CompactTextString(m) }
func (*MqttSubscribe) ProtoMessage()               {}
func (*MqttSubscribe) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *MqttSubscribe) GetHeader() *MqttHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *MqttSubscribe) GetMessageId() int32 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

func (m *MqttSubscribe) GetTopic() [][]byte {
	if m != nil {
		return m.Topic
	}
	return nil
}

func (m *MqttSubscribe) GetQos() []int32 {
	if m != nil {
		return m.Qos
	}
	return nil
}

type MqttSuback struct {
	Header    *MqttHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	MessageId int32       `protobuf:"varint,2,opt,name=messageId" json:"messageId,omitempty"`
	Qos       []int32     `protobuf:"varint,3,rep,packed,name=qos" json:"qos,omitempty"`
}

func (m *MqttSuback) Reset()                    { *m = MqttSuback{} }
func (m *MqttSuback) String() string            { return proto.CompactTextString(m) }
func (*MqttSuback) ProtoMessage()               {}
func (*MqttSuback) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *MqttSuback) GetHeader() *MqttHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *MqttSuback) GetMessageId() int32 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

func (m *MqttSuback) GetQos() []int32 {
	if m != nil {
		return m.Qos
	}
	return nil
}

type MqttUnsubscribe struct {
	Header    *MqttHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	MessageId int32       `protobuf:"varint,2,opt,name=messageId" json:"messageId,omitempty"`
	Topic     [][]byte    `protobuf:"bytes,3,rep,name=topic,proto3" json:"topic,omitempty"`
}

func (m *MqttUnsubscribe) Reset()                    { *m = MqttUnsubscribe{} }
func (m *MqttUnsubscribe) String() string            { return proto.CompactTextString(m) }
func (*MqttUnsubscribe) ProtoMessage()               {}
func (*MqttUnsubscribe) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *MqttUnsubscribe) GetHeader() *MqttHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *MqttUnsubscribe) GetMessageId() int32 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

func (m *MqttUnsubscribe) GetTopic() [][]byte {
	if m != nil {
		return m.Topic
	}
	return nil
}

type MqttUnsuback struct {
	Header    *MqttHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	MessageId int32       `protobuf:"varint,2,opt,name=messageId" json:"messageId,omitempty"`
}

func (m *MqttUnsuback) Reset()                    { *m = MqttUnsuback{} }
func (m *MqttUnsuback) String() string            { return proto.CompactTextString(m) }
func (*MqttUnsuback) ProtoMessage()               {}
func (*MqttUnsuback) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *MqttUnsuback) GetHeader() *MqttHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *MqttUnsuback) GetMessageId() int32 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

type MqttPingReq struct {
	Header *MqttHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
}

func (m *MqttPingReq) Reset()                    { *m = MqttPingReq{} }
func (m *MqttPingReq) String() string            { return proto.CompactTextString(m) }
func (*MqttPingReq) ProtoMessage()               {}
func (*MqttPingReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *MqttPingReq) GetHeader() *MqttHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type MqttPingResp struct {
	Header *MqttHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
}

func (m *MqttPingResp) Reset()                    { *m = MqttPingResp{} }
func (m *MqttPingResp) String() string            { return proto.CompactTextString(m) }
func (*MqttPingResp) ProtoMessage()               {}
func (*MqttPingResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *MqttPingResp) GetHeader() *MqttHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func init() {
	proto.RegisterType((*MqttHeader)(nil), "MqttHeader")
	proto.RegisterType((*MqttConnect)(nil), "MqttConnect")
	proto.RegisterType((*MqttConnack)(nil), "MqttConnack")
	proto.RegisterType((*MqttDisconnect)(nil), "MqttDisconnect")
	proto.RegisterType((*MqttPublish)(nil), "MqttPublish")
	proto.RegisterType((*MqttPuback)(nil), "MqttPuback")
	proto.RegisterType((*MqttSubscribe)(nil), "MqttSubscribe")
	proto.RegisterType((*MqttSuback)(nil), "MqttSuback")
	proto.RegisterType((*MqttUnsubscribe)(nil), "MqttUnsubscribe")
	proto.RegisterType((*MqttUnsuback)(nil), "MqttUnsuback")
	proto.RegisterType((*MqttPingReq)(nil), "MqttPingReq")
	proto.RegisterType((*MqttPingResp)(nil), "MqttPingResp")
}

func init() { proto.RegisterFile("pb.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 452 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xc1, 0x6e, 0xd4, 0x30,
	0x10, 0x55, 0x36, 0xdd, 0xed, 0xee, 0x64, 0x59, 0x90, 0x85, 0x90, 0x85, 0xaa, 0x2a, 0x32, 0x12,
	0xda, 0x53, 0x0f, 0xed, 0x27, 0x14, 0xa4, 0xf6, 0x80, 0xd8, 0xba, 0xe5, 0x03, 0x9c, 0x78, 0xd4,
	0x9a, 0x66, 0xe3, 0xac, 0xed, 0x50, 0xc1, 0x89, 0x2f, 0xe4, 0x9b, 0x90, 0xed, 0xa4, 0x59, 0x71,
	0x40, 0x91, 0x58, 0x6e, 0x7e, 0xf3, 0x32, 0x6f, 0x66, 0xde, 0xd8, 0x81, 0x79, 0x53, 0x9c, 0x35,
	0x46, 0x3b, 0xcd, 0xae, 0x00, 0x3e, 0xed, 0x9c, 0xbb, 0x42, 0x21, 0xd1, 0x90, 0x57, 0x90, 0xca,
	0xb6, 0xa1, 0x49, 0x9e, 0xac, 0xe7, 0xdc, 0x1f, 0x7d, 0x64, 0xa7, 0x2d, 0x9d, 0xe4, 0xc9, 0x7a,
	0xca, 0xfd, 0x91, 0xbc, 0x81, 0x99, 0x41, 0x27, 0x54, 0x4d, 0xd3, 0xf0, 0x59, 0x87, 0xd8, 0xaf,
	0x09, 0x64, 0x5e, 0xea, 0x52, 0xd7, 0x35, 0x96, 0x8e, 0xbc, 0x83, 0xd9, 0x43, 0x50, 0x0d, 0x72,
	0xd9, 0x79, 0x76, 0x36, 0x14, 0xe2, 0x1d, 0x45, 0x56, 0x30, 0x51, 0x32, 0xa8, 0x2f, 0xf8, 0x44,
	0x49, 0xf2, 0x1a, 0xa6, 0x65, 0x85, 0xa2, 0xd7, 0x8e, 0x80, 0xbc, 0x85, 0x79, 0x59, 0x29, 0xac,
	0xdd, 0xb5, 0xa4, 0x47, 0x79, 0xb2, 0x5e, 0xf2, 0x67, 0xec, 0xb9, 0xd6, 0xa2, 0xa9, 0xc5, 0x16,
	0xe9, 0x34, 0x72, 0x3d, 0xf6, 0x5c, 0x23, 0xac, 0x7d, 0xd2, 0x46, 0xd2, 0x59, 0xe4, 0x7a, 0x4c,
	0x4e, 0x60, 0xf1, 0xa4, 0xaa, 0xea, 0x4e, 0x37, 0xaa, 0xa4, 0xc7, 0x81, 0x1c, 0x02, 0x24, 0x87,
	0xcc, 0x83, 0x8d, 0xf8, 0x5e, 0x69, 0x21, 0xe9, 0x3c, 0xf0, 0xfb, 0x21, 0x42, 0xe1, 0xd8, 0xc3,
	0x1b, 0x6d, 0xe9, 0x22, 0x98, 0xd3, 0x43, 0x72, 0x0a, 0xe0, 0x8f, 0x3c, 0x9a, 0x04, 0x61, 0x90,
	0xbd, 0x08, 0x79, 0x0f, 0xab, 0x47, 0xc4, 0x46, 0x54, 0xea, 0x1b, 0xde, 0xa9, 0x2d, 0x1a, 0x9a,
	0x05, 0x81, 0x3f, 0xa2, 0x8c, 0x0f, 0x7e, 0x8a, 0xf2, 0x71, 0x9c, 0x9f, 0xa7, 0x00, 0x06, 0x5d,
	0x6b, 0xea, 0x4b, 0x2d, 0x31, 0xf8, 0xba, 0xe4, 0x7b, 0x11, 0xf6, 0x11, 0x56, 0x3e, 0xeb, 0x83,
	0xb2, 0xe5, 0x3f, 0xac, 0x89, 0xfd, 0x4c, 0x62, 0x6f, 0x9b, 0xb6, 0xa8, 0x94, 0x7d, 0x18, 0x27,
	0x72, 0x02, 0x8b, 0x2d, 0x5a, 0x2b, 0xee, 0xf1, 0x5a, 0x76, 0x17, 0x6a, 0x08, 0xf8, 0xcd, 0xbb,
	0xb0, 0x8b, 0x34, 0x34, 0x1d, 0x81, 0x77, 0xb9, 0xe9, 0x76, 0x10, 0x17, 0xdf, 0x43, 0xf6, 0x39,
	0x5e, 0xdc, 0x4d, 0x5b, 0x8c, 0x36, 0xe7, 0xaf, 0x0d, 0xb0, 0x1f, 0xf0, 0xc2, 0xe7, 0xdc, 0xb6,
	0x85, 0x2d, 0x8d, 0x2a, 0xf0, 0xc0, 0x43, 0xa5, 0xc3, 0x50, 0xdd, 0x9b, 0x3a, 0xca, 0xd3, 0xee,
	0x4d, 0x31, 0x11, 0x87, 0xb9, 0x3d, 0xd4, 0x30, 0x7d, 0x89, 0x74, 0x28, 0xf1, 0x15, 0x5e, 0x7a,
	0x95, 0x2f, 0xb5, 0xfd, 0xef, 0x03, 0xb2, 0x1b, 0x58, 0x3e, 0xd7, 0x3a, 0xd0, 0x76, 0xce, 0xbb,
	0x0b, 0xa7, 0xea, 0x7b, 0x8e, 0xbb, 0x51, 0x8a, 0xec, 0x22, 0xb6, 0x11, 0x73, 0x6c, 0x33, 0x2a,
	0xa9, 0x98, 0x85, 0xff, 0xe2, 0xc5, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xca, 0xe6, 0xc7, 0x35,
	0x23, 0x05, 0x00, 0x00,
}
