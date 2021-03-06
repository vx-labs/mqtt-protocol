// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb.proto

package packet

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Header struct {
	Dup                  bool     `protobuf:"varint,1,opt,name=dup,proto3" json:"dup,omitempty"`
	Qos                  int32    `protobuf:"varint,2,opt,name=qos,proto3" json:"qos,omitempty"`
	Retain               bool     `protobuf:"varint,3,opt,name=retain,proto3" json:"retain,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Header) Reset()         { *m = Header{} }
func (m *Header) String() string { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()    {}
func (*Header) Descriptor() ([]byte, []int) {
	return fileDescriptor_f80abaa17e25ccc8, []int{0}
}

func (m *Header) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Header.Unmarshal(m, b)
}
func (m *Header) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Header.Marshal(b, m, deterministic)
}
func (m *Header) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Header.Merge(m, src)
}
func (m *Header) XXX_Size() int {
	return xxx_messageInfo_Header.Size(m)
}
func (m *Header) XXX_DiscardUnknown() {
	xxx_messageInfo_Header.DiscardUnknown(m)
}

var xxx_messageInfo_Header proto.InternalMessageInfo

func (m *Header) GetDup() bool {
	if m != nil {
		return m.Dup
	}
	return false
}

func (m *Header) GetQos() int32 {
	if m != nil {
		return m.Qos
	}
	return 0
}

func (m *Header) GetRetain() bool {
	if m != nil {
		return m.Retain
	}
	return false
}

type Connect struct {
	Header               *Header  `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Clean                bool     `protobuf:"varint,3,opt,name=clean,proto3" json:"clean,omitempty"`
	ClientId             []byte   `protobuf:"bytes,4,opt,name=clientId,proto3" json:"clientId,omitempty"`
	Username             []byte   `protobuf:"bytes,5,opt,name=username,proto3" json:"username,omitempty"`
	Password             []byte   `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty"`
	WillTopic            []byte   `protobuf:"bytes,7,opt,name=willTopic,proto3" json:"willTopic,omitempty"`
	WillPayload          []byte   `protobuf:"bytes,8,opt,name=willPayload,proto3" json:"willPayload,omitempty"`
	WillQos              int32    `protobuf:"varint,9,opt,name=willQos,proto3" json:"willQos,omitempty"`
	WillRetain           bool     `protobuf:"varint,10,opt,name=willRetain,proto3" json:"willRetain,omitempty"`
	KeepaliveTimer       int32    `protobuf:"varint,11,opt,name=keepaliveTimer,proto3" json:"keepaliveTimer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Connect) Reset()         { *m = Connect{} }
func (m *Connect) String() string { return proto.CompactTextString(m) }
func (*Connect) ProtoMessage()    {}
func (*Connect) Descriptor() ([]byte, []int) {
	return fileDescriptor_f80abaa17e25ccc8, []int{1}
}

func (m *Connect) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Connect.Unmarshal(m, b)
}
func (m *Connect) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Connect.Marshal(b, m, deterministic)
}
func (m *Connect) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Connect.Merge(m, src)
}
func (m *Connect) XXX_Size() int {
	return xxx_messageInfo_Connect.Size(m)
}
func (m *Connect) XXX_DiscardUnknown() {
	xxx_messageInfo_Connect.DiscardUnknown(m)
}

var xxx_messageInfo_Connect proto.InternalMessageInfo

func (m *Connect) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Connect) GetClean() bool {
	if m != nil {
		return m.Clean
	}
	return false
}

func (m *Connect) GetClientId() []byte {
	if m != nil {
		return m.ClientId
	}
	return nil
}

func (m *Connect) GetUsername() []byte {
	if m != nil {
		return m.Username
	}
	return nil
}

func (m *Connect) GetPassword() []byte {
	if m != nil {
		return m.Password
	}
	return nil
}

func (m *Connect) GetWillTopic() []byte {
	if m != nil {
		return m.WillTopic
	}
	return nil
}

func (m *Connect) GetWillPayload() []byte {
	if m != nil {
		return m.WillPayload
	}
	return nil
}

func (m *Connect) GetWillQos() int32 {
	if m != nil {
		return m.WillQos
	}
	return 0
}

func (m *Connect) GetWillRetain() bool {
	if m != nil {
		return m.WillRetain
	}
	return false
}

func (m *Connect) GetKeepaliveTimer() int32 {
	if m != nil {
		return m.KeepaliveTimer
	}
	return 0
}

type ConnAck struct {
	Header               *Header  `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	ReturnCode           int32    `protobuf:"varint,2,opt,name=returnCode,proto3" json:"returnCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConnAck) Reset()         { *m = ConnAck{} }
func (m *ConnAck) String() string { return proto.CompactTextString(m) }
func (*ConnAck) ProtoMessage()    {}
func (*ConnAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_f80abaa17e25ccc8, []int{2}
}

func (m *ConnAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnAck.Unmarshal(m, b)
}
func (m *ConnAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnAck.Marshal(b, m, deterministic)
}
func (m *ConnAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnAck.Merge(m, src)
}
func (m *ConnAck) XXX_Size() int {
	return xxx_messageInfo_ConnAck.Size(m)
}
func (m *ConnAck) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnAck.DiscardUnknown(m)
}

var xxx_messageInfo_ConnAck proto.InternalMessageInfo

func (m *ConnAck) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ConnAck) GetReturnCode() int32 {
	if m != nil {
		return m.ReturnCode
	}
	return 0
}

type Disconnect struct {
	Header               *Header  `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Disconnect) Reset()         { *m = Disconnect{} }
func (m *Disconnect) String() string { return proto.CompactTextString(m) }
func (*Disconnect) ProtoMessage()    {}
func (*Disconnect) Descriptor() ([]byte, []int) {
	return fileDescriptor_f80abaa17e25ccc8, []int{3}
}

func (m *Disconnect) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Disconnect.Unmarshal(m, b)
}
func (m *Disconnect) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Disconnect.Marshal(b, m, deterministic)
}
func (m *Disconnect) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Disconnect.Merge(m, src)
}
func (m *Disconnect) XXX_Size() int {
	return xxx_messageInfo_Disconnect.Size(m)
}
func (m *Disconnect) XXX_DiscardUnknown() {
	xxx_messageInfo_Disconnect.DiscardUnknown(m)
}

var xxx_messageInfo_Disconnect proto.InternalMessageInfo

func (m *Disconnect) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Disconnect) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Publish struct {
	Header               *Header  `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	MessageId            int32    `protobuf:"varint,2,opt,name=messageId,proto3" json:"messageId,omitempty"`
	Topic                []byte   `protobuf:"bytes,3,opt,name=topic,proto3" json:"topic,omitempty"`
	Payload              []byte   `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Publish) Reset()         { *m = Publish{} }
func (m *Publish) String() string { return proto.CompactTextString(m) }
func (*Publish) ProtoMessage()    {}
func (*Publish) Descriptor() ([]byte, []int) {
	return fileDescriptor_f80abaa17e25ccc8, []int{4}
}

func (m *Publish) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Publish.Unmarshal(m, b)
}
func (m *Publish) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Publish.Marshal(b, m, deterministic)
}
func (m *Publish) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Publish.Merge(m, src)
}
func (m *Publish) XXX_Size() int {
	return xxx_messageInfo_Publish.Size(m)
}
func (m *Publish) XXX_DiscardUnknown() {
	xxx_messageInfo_Publish.DiscardUnknown(m)
}

var xxx_messageInfo_Publish proto.InternalMessageInfo

func (m *Publish) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Publish) GetMessageId() int32 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

func (m *Publish) GetTopic() []byte {
	if m != nil {
		return m.Topic
	}
	return nil
}

func (m *Publish) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type PubAck struct {
	Header               *Header  `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	MessageId            int32    `protobuf:"varint,2,opt,name=messageId,proto3" json:"messageId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PubAck) Reset()         { *m = PubAck{} }
func (m *PubAck) String() string { return proto.CompactTextString(m) }
func (*PubAck) ProtoMessage()    {}
func (*PubAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_f80abaa17e25ccc8, []int{5}
}

func (m *PubAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PubAck.Unmarshal(m, b)
}
func (m *PubAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PubAck.Marshal(b, m, deterministic)
}
func (m *PubAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PubAck.Merge(m, src)
}
func (m *PubAck) XXX_Size() int {
	return xxx_messageInfo_PubAck.Size(m)
}
func (m *PubAck) XXX_DiscardUnknown() {
	xxx_messageInfo_PubAck.DiscardUnknown(m)
}

var xxx_messageInfo_PubAck proto.InternalMessageInfo

func (m *PubAck) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *PubAck) GetMessageId() int32 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

type PubRel struct {
	Header               *Header  `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	MessageId            int32    `protobuf:"varint,2,opt,name=messageId,proto3" json:"messageId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PubRel) Reset()         { *m = PubRel{} }
func (m *PubRel) String() string { return proto.CompactTextString(m) }
func (*PubRel) ProtoMessage()    {}
func (*PubRel) Descriptor() ([]byte, []int) {
	return fileDescriptor_f80abaa17e25ccc8, []int{6}
}

func (m *PubRel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PubRel.Unmarshal(m, b)
}
func (m *PubRel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PubRel.Marshal(b, m, deterministic)
}
func (m *PubRel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PubRel.Merge(m, src)
}
func (m *PubRel) XXX_Size() int {
	return xxx_messageInfo_PubRel.Size(m)
}
func (m *PubRel) XXX_DiscardUnknown() {
	xxx_messageInfo_PubRel.DiscardUnknown(m)
}

var xxx_messageInfo_PubRel proto.InternalMessageInfo

func (m *PubRel) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *PubRel) GetMessageId() int32 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

type PubComp struct {
	Header               *Header  `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	MessageId            int32    `protobuf:"varint,2,opt,name=messageId,proto3" json:"messageId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PubComp) Reset()         { *m = PubComp{} }
func (m *PubComp) String() string { return proto.CompactTextString(m) }
func (*PubComp) ProtoMessage()    {}
func (*PubComp) Descriptor() ([]byte, []int) {
	return fileDescriptor_f80abaa17e25ccc8, []int{7}
}

func (m *PubComp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PubComp.Unmarshal(m, b)
}
func (m *PubComp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PubComp.Marshal(b, m, deterministic)
}
func (m *PubComp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PubComp.Merge(m, src)
}
func (m *PubComp) XXX_Size() int {
	return xxx_messageInfo_PubComp.Size(m)
}
func (m *PubComp) XXX_DiscardUnknown() {
	xxx_messageInfo_PubComp.DiscardUnknown(m)
}

var xxx_messageInfo_PubComp proto.InternalMessageInfo

func (m *PubComp) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *PubComp) GetMessageId() int32 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

type PubRec struct {
	Header               *Header  `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	MessageId            int32    `protobuf:"varint,2,opt,name=messageId,proto3" json:"messageId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PubRec) Reset()         { *m = PubRec{} }
func (m *PubRec) String() string { return proto.CompactTextString(m) }
func (*PubRec) ProtoMessage()    {}
func (*PubRec) Descriptor() ([]byte, []int) {
	return fileDescriptor_f80abaa17e25ccc8, []int{8}
}

func (m *PubRec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PubRec.Unmarshal(m, b)
}
func (m *PubRec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PubRec.Marshal(b, m, deterministic)
}
func (m *PubRec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PubRec.Merge(m, src)
}
func (m *PubRec) XXX_Size() int {
	return xxx_messageInfo_PubRec.Size(m)
}
func (m *PubRec) XXX_DiscardUnknown() {
	xxx_messageInfo_PubRec.DiscardUnknown(m)
}

var xxx_messageInfo_PubRec proto.InternalMessageInfo

func (m *PubRec) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *PubRec) GetMessageId() int32 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

type Subscribe struct {
	Header               *Header  `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	MessageId            int32    `protobuf:"varint,2,opt,name=messageId,proto3" json:"messageId,omitempty"`
	Topic                [][]byte `protobuf:"bytes,3,rep,name=topic,proto3" json:"topic,omitempty"`
	Qos                  []int32  `protobuf:"varint,4,rep,packed,name=qos,proto3" json:"qos,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Subscribe) Reset()         { *m = Subscribe{} }
func (m *Subscribe) String() string { return proto.CompactTextString(m) }
func (*Subscribe) ProtoMessage()    {}
func (*Subscribe) Descriptor() ([]byte, []int) {
	return fileDescriptor_f80abaa17e25ccc8, []int{9}
}

func (m *Subscribe) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Subscribe.Unmarshal(m, b)
}
func (m *Subscribe) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Subscribe.Marshal(b, m, deterministic)
}
func (m *Subscribe) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Subscribe.Merge(m, src)
}
func (m *Subscribe) XXX_Size() int {
	return xxx_messageInfo_Subscribe.Size(m)
}
func (m *Subscribe) XXX_DiscardUnknown() {
	xxx_messageInfo_Subscribe.DiscardUnknown(m)
}

var xxx_messageInfo_Subscribe proto.InternalMessageInfo

func (m *Subscribe) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Subscribe) GetMessageId() int32 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

func (m *Subscribe) GetTopic() [][]byte {
	if m != nil {
		return m.Topic
	}
	return nil
}

func (m *Subscribe) GetQos() []int32 {
	if m != nil {
		return m.Qos
	}
	return nil
}

type SubAck struct {
	Header               *Header  `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	MessageId            int32    `protobuf:"varint,2,opt,name=messageId,proto3" json:"messageId,omitempty"`
	Qos                  []int32  `protobuf:"varint,3,rep,packed,name=qos,proto3" json:"qos,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubAck) Reset()         { *m = SubAck{} }
func (m *SubAck) String() string { return proto.CompactTextString(m) }
func (*SubAck) ProtoMessage()    {}
func (*SubAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_f80abaa17e25ccc8, []int{10}
}

func (m *SubAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubAck.Unmarshal(m, b)
}
func (m *SubAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubAck.Marshal(b, m, deterministic)
}
func (m *SubAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubAck.Merge(m, src)
}
func (m *SubAck) XXX_Size() int {
	return xxx_messageInfo_SubAck.Size(m)
}
func (m *SubAck) XXX_DiscardUnknown() {
	xxx_messageInfo_SubAck.DiscardUnknown(m)
}

var xxx_messageInfo_SubAck proto.InternalMessageInfo

func (m *SubAck) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SubAck) GetMessageId() int32 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

func (m *SubAck) GetQos() []int32 {
	if m != nil {
		return m.Qos
	}
	return nil
}

type Unsubscribe struct {
	Header               *Header  `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	MessageId            int32    `protobuf:"varint,2,opt,name=messageId,proto3" json:"messageId,omitempty"`
	Topic                [][]byte `protobuf:"bytes,3,rep,name=topic,proto3" json:"topic,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Unsubscribe) Reset()         { *m = Unsubscribe{} }
func (m *Unsubscribe) String() string { return proto.CompactTextString(m) }
func (*Unsubscribe) ProtoMessage()    {}
func (*Unsubscribe) Descriptor() ([]byte, []int) {
	return fileDescriptor_f80abaa17e25ccc8, []int{11}
}

func (m *Unsubscribe) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Unsubscribe.Unmarshal(m, b)
}
func (m *Unsubscribe) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Unsubscribe.Marshal(b, m, deterministic)
}
func (m *Unsubscribe) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Unsubscribe.Merge(m, src)
}
func (m *Unsubscribe) XXX_Size() int {
	return xxx_messageInfo_Unsubscribe.Size(m)
}
func (m *Unsubscribe) XXX_DiscardUnknown() {
	xxx_messageInfo_Unsubscribe.DiscardUnknown(m)
}

var xxx_messageInfo_Unsubscribe proto.InternalMessageInfo

func (m *Unsubscribe) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Unsubscribe) GetMessageId() int32 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

func (m *Unsubscribe) GetTopic() [][]byte {
	if m != nil {
		return m.Topic
	}
	return nil
}

type UnsubAck struct {
	Header               *Header  `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	MessageId            int32    `protobuf:"varint,2,opt,name=messageId,proto3" json:"messageId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnsubAck) Reset()         { *m = UnsubAck{} }
func (m *UnsubAck) String() string { return proto.CompactTextString(m) }
func (*UnsubAck) ProtoMessage()    {}
func (*UnsubAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_f80abaa17e25ccc8, []int{12}
}

func (m *UnsubAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnsubAck.Unmarshal(m, b)
}
func (m *UnsubAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnsubAck.Marshal(b, m, deterministic)
}
func (m *UnsubAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnsubAck.Merge(m, src)
}
func (m *UnsubAck) XXX_Size() int {
	return xxx_messageInfo_UnsubAck.Size(m)
}
func (m *UnsubAck) XXX_DiscardUnknown() {
	xxx_messageInfo_UnsubAck.DiscardUnknown(m)
}

var xxx_messageInfo_UnsubAck proto.InternalMessageInfo

func (m *UnsubAck) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *UnsubAck) GetMessageId() int32 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

type PingReq struct {
	Header               *Header  `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingReq) Reset()         { *m = PingReq{} }
func (m *PingReq) String() string { return proto.CompactTextString(m) }
func (*PingReq) ProtoMessage()    {}
func (*PingReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_f80abaa17e25ccc8, []int{13}
}

func (m *PingReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingReq.Unmarshal(m, b)
}
func (m *PingReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingReq.Marshal(b, m, deterministic)
}
func (m *PingReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingReq.Merge(m, src)
}
func (m *PingReq) XXX_Size() int {
	return xxx_messageInfo_PingReq.Size(m)
}
func (m *PingReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PingReq.DiscardUnknown(m)
}

var xxx_messageInfo_PingReq proto.InternalMessageInfo

func (m *PingReq) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

type PingResp struct {
	Header               *Header  `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingResp) Reset()         { *m = PingResp{} }
func (m *PingResp) String() string { return proto.CompactTextString(m) }
func (*PingResp) ProtoMessage()    {}
func (*PingResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_f80abaa17e25ccc8, []int{14}
}

func (m *PingResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingResp.Unmarshal(m, b)
}
func (m *PingResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingResp.Marshal(b, m, deterministic)
}
func (m *PingResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingResp.Merge(m, src)
}
func (m *PingResp) XXX_Size() int {
	return xxx_messageInfo_PingResp.Size(m)
}
func (m *PingResp) XXX_DiscardUnknown() {
	xxx_messageInfo_PingResp.DiscardUnknown(m)
}

var xxx_messageInfo_PingResp proto.InternalMessageInfo

func (m *PingResp) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func init() {
	proto.RegisterType((*Header)(nil), "packet.Header")
	proto.RegisterType((*Connect)(nil), "packet.Connect")
	proto.RegisterType((*ConnAck)(nil), "packet.ConnAck")
	proto.RegisterType((*Disconnect)(nil), "packet.Disconnect")
	proto.RegisterType((*Publish)(nil), "packet.Publish")
	proto.RegisterType((*PubAck)(nil), "packet.PubAck")
	proto.RegisterType((*PubRel)(nil), "packet.PubRel")
	proto.RegisterType((*PubComp)(nil), "packet.PubComp")
	proto.RegisterType((*PubRec)(nil), "packet.PubRec")
	proto.RegisterType((*Subscribe)(nil), "packet.Subscribe")
	proto.RegisterType((*SubAck)(nil), "packet.SubAck")
	proto.RegisterType((*Unsubscribe)(nil), "packet.Unsubscribe")
	proto.RegisterType((*UnsubAck)(nil), "packet.UnsubAck")
	proto.RegisterType((*PingReq)(nil), "packet.PingReq")
	proto.RegisterType((*PingResp)(nil), "packet.PingResp")
}

func init() { proto.RegisterFile("pb.proto", fileDescriptor_f80abaa17e25ccc8) }

var fileDescriptor_f80abaa17e25ccc8 = []byte{
	// 458 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0x41, 0x6b, 0xdb, 0x30,
	0x18, 0x25, 0x71, 0xe3, 0xd8, 0x5f, 0x46, 0x18, 0x62, 0x0c, 0x31, 0x4a, 0x31, 0x3e, 0x94, 0x9c,
	0x02, 0xeb, 0x7e, 0xc1, 0x48, 0x0e, 0xeb, 0x65, 0x73, 0xdd, 0xee, 0x3e, 0x59, 0xfa, 0x68, 0x45,
	0x1c, 0x4b, 0x95, 0xec, 0x95, 0x9e, 0xf6, 0xcb, 0xf6, 0xdf, 0x86, 0x24, 0x7b, 0x09, 0x3b, 0x8c,
	0x14, 0x4c, 0x6f, 0x7a, 0xef, 0xd9, 0xef, 0x93, 0xde, 0x13, 0x82, 0x44, 0x57, 0x6b, 0x6d, 0x54,
	0xab, 0x48, 0xac, 0x19, 0xdf, 0x61, 0x9b, 0x6f, 0x21, 0xfe, 0x82, 0x4c, 0xa0, 0x21, 0x6f, 0x21,
	0x12, 0x9d, 0xa6, 0x93, 0x6c, 0xb2, 0x4a, 0x4a, 0xb7, 0x74, 0xcc, 0xa3, 0xb2, 0x74, 0x9a, 0x4d,
	0x56, 0xb3, 0xd2, 0x2d, 0xc9, 0x7b, 0x88, 0x0d, 0xb6, 0x4c, 0x36, 0x34, 0xf2, 0x9f, 0xf5, 0x28,
	0xff, 0x3d, 0x85, 0xf9, 0x46, 0x35, 0x0d, 0xf2, 0x96, 0x5c, 0x42, 0xfc, 0xe0, 0x1d, 0xbd, 0xd5,
	0xe2, 0x6a, 0xb9, 0x0e, 0xa3, 0xd6, 0x61, 0x4e, 0xd9, 0xab, 0xe4, 0x1d, 0xcc, 0x78, 0x8d, 0x6c,
	0xb0, 0x0a, 0x80, 0x7c, 0x80, 0x84, 0xd7, 0x12, 0x9b, 0xf6, 0x5a, 0xd0, 0xb3, 0x6c, 0xb2, 0x7a,
	0x53, 0xfe, 0xc5, 0x4e, 0xeb, 0x2c, 0x9a, 0x86, 0xed, 0x91, 0xce, 0x82, 0x36, 0x60, 0xa7, 0x69,
	0x66, 0xed, 0x93, 0x32, 0x82, 0xc6, 0x41, 0x1b, 0x30, 0x39, 0x87, 0xf4, 0x49, 0xd6, 0xf5, 0x9d,
	0xd2, 0x92, 0xd3, 0xb9, 0x17, 0x0f, 0x04, 0xc9, 0x60, 0xe1, 0x40, 0xc1, 0x9e, 0x6b, 0xc5, 0x04,
	0x4d, 0xbc, 0x7e, 0x4c, 0x11, 0x0a, 0x73, 0x07, 0x6f, 0x94, 0xa5, 0xa9, 0xcf, 0x62, 0x80, 0xe4,
	0x02, 0xc0, 0x2d, 0xcb, 0x90, 0x09, 0xf8, 0x83, 0x1c, 0x31, 0xe4, 0x12, 0x96, 0x3b, 0x44, 0xcd,
	0x6a, 0xf9, 0x13, 0xef, 0xe4, 0x1e, 0x0d, 0x5d, 0x78, 0x83, 0x7f, 0xd8, 0xfc, 0x26, 0xc4, 0xf7,
	0x99, 0xef, 0x4e, 0x8e, 0xef, 0x02, 0xc0, 0x60, 0xdb, 0x99, 0x66, 0xa3, 0x04, 0xf6, 0x1d, 0x1d,
	0x31, 0xf9, 0x16, 0x60, 0x2b, 0x2d, 0x7f, 0x61, 0x29, 0x4b, 0x98, 0x4a, 0xe1, 0xdd, 0xd2, 0x72,
	0x2a, 0x45, 0xfe, 0x0b, 0xe6, 0x45, 0x57, 0xd5, 0xd2, 0x3e, 0x9c, 0x6c, 0x71, 0x0e, 0xe9, 0x1e,
	0xad, 0x65, 0xf7, 0x78, 0x2d, 0xfa, 0x7d, 0x1d, 0x08, 0xd7, 0x7a, 0xeb, 0x7b, 0x88, 0x7c, 0xce,
	0x01, 0xb8, 0x84, 0x75, 0x9f, 0x7f, 0x28, 0x7d, 0x80, 0xf9, 0x57, 0x88, 0x8b, 0xae, 0x7a, 0x49,
	0x30, 0xff, 0x9d, 0xdf, 0xfb, 0x95, 0x58, 0x8f, 0xe4, 0xf7, 0xcd, 0x07, 0xb4, 0x51, 0x7b, 0x3d,
	0xf2, 0x06, 0xf9, 0x48, 0x7e, 0xcf, 0x90, 0xde, 0x76, 0x95, 0xe5, 0x46, 0x56, 0x38, 0x7e, 0x87,
	0xd1, 0xa1, 0xc3, 0xfe, 0xb5, 0x38, 0xcb, 0xa2, 0xfe, 0xb5, 0xc8, 0x7f, 0x40, 0x7c, 0x3b, 0x62,
	0x77, 0xc3, 0x84, 0xe8, 0x30, 0x41, 0xc2, 0xe2, 0x7b, 0x63, 0x5f, 0xe3, 0x78, 0x79, 0x01, 0x89,
	0x1f, 0x35, 0xde, 0x55, 0xfc, 0x08, 0xf3, 0x42, 0x36, 0xf7, 0x25, 0x3e, 0x9e, 0x6a, 0x98, 0x5f,
	0x41, 0x12, 0x7e, 0xb1, 0x27, 0x5f, 0xb7, 0x2a, 0xf6, 0x0f, 0xfe, 0xa7, 0x3f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xc6, 0x5d, 0x51, 0xed, 0xfc, 0x05, 0x00, 0x00,
}
