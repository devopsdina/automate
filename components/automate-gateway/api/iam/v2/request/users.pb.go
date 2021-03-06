// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/automate-gateway/api/iam/v2/request/users.proto

package request

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
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

type CreateUserReq struct {
	// Unique ID. Cannot be changed. Used to sign in.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Display name for local user.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Password for user. Used to sign in.
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserReq) Reset()         { *m = CreateUserReq{} }
func (m *CreateUserReq) String() string { return proto.CompactTextString(m) }
func (*CreateUserReq) ProtoMessage()    {}
func (*CreateUserReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0d105d392afc5cf, []int{0}
}

func (m *CreateUserReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserReq.Unmarshal(m, b)
}
func (m *CreateUserReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserReq.Marshal(b, m, deterministic)
}
func (m *CreateUserReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserReq.Merge(m, src)
}
func (m *CreateUserReq) XXX_Size() int {
	return xxx_messageInfo_CreateUserReq.Size(m)
}
func (m *CreateUserReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserReq proto.InternalMessageInfo

func (m *CreateUserReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CreateUserReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateUserReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type ListUsersReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListUsersReq) Reset()         { *m = ListUsersReq{} }
func (m *ListUsersReq) String() string { return proto.CompactTextString(m) }
func (*ListUsersReq) ProtoMessage()    {}
func (*ListUsersReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0d105d392afc5cf, []int{1}
}

func (m *ListUsersReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUsersReq.Unmarshal(m, b)
}
func (m *ListUsersReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUsersReq.Marshal(b, m, deterministic)
}
func (m *ListUsersReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUsersReq.Merge(m, src)
}
func (m *ListUsersReq) XXX_Size() int {
	return xxx_messageInfo_ListUsersReq.Size(m)
}
func (m *ListUsersReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUsersReq.DiscardUnknown(m)
}

var xxx_messageInfo_ListUsersReq proto.InternalMessageInfo

type DeleteUserReq struct {
	// ID of the user.
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteUserReq) Reset()         { *m = DeleteUserReq{} }
func (m *DeleteUserReq) String() string { return proto.CompactTextString(m) }
func (*DeleteUserReq) ProtoMessage()    {}
func (*DeleteUserReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0d105d392afc5cf, []int{2}
}

func (m *DeleteUserReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteUserReq.Unmarshal(m, b)
}
func (m *DeleteUserReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteUserReq.Marshal(b, m, deterministic)
}
func (m *DeleteUserReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteUserReq.Merge(m, src)
}
func (m *DeleteUserReq) XXX_Size() int {
	return xxx_messageInfo_DeleteUserReq.Size(m)
}
func (m *DeleteUserReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteUserReq.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteUserReq proto.InternalMessageInfo

func (m *DeleteUserReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetUserReq struct {
	// ID of the user.
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserReq) Reset()         { *m = GetUserReq{} }
func (m *GetUserReq) String() string { return proto.CompactTextString(m) }
func (*GetUserReq) ProtoMessage()    {}
func (*GetUserReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0d105d392afc5cf, []int{3}
}

func (m *GetUserReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserReq.Unmarshal(m, b)
}
func (m *GetUserReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserReq.Marshal(b, m, deterministic)
}
func (m *GetUserReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserReq.Merge(m, src)
}
func (m *GetUserReq) XXX_Size() int {
	return xxx_messageInfo_GetUserReq.Size(m)
}
func (m *GetUserReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserReq proto.InternalMessageInfo

func (m *GetUserReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type UpdateUserReq struct {
	// ID of the user. Cannot be changed. Used to sign in.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Display name for local user.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Password used to log in. Will overwrite preexisting password.
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserReq) Reset()         { *m = UpdateUserReq{} }
func (m *UpdateUserReq) String() string { return proto.CompactTextString(m) }
func (*UpdateUserReq) ProtoMessage()    {}
func (*UpdateUserReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0d105d392afc5cf, []int{4}
}

func (m *UpdateUserReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserReq.Unmarshal(m, b)
}
func (m *UpdateUserReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserReq.Marshal(b, m, deterministic)
}
func (m *UpdateUserReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserReq.Merge(m, src)
}
func (m *UpdateUserReq) XXX_Size() int {
	return xxx_messageInfo_UpdateUserReq.Size(m)
}
func (m *UpdateUserReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserReq proto.InternalMessageInfo

func (m *UpdateUserReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateUserReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateUserReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type UpdateSelfReq struct {
	// ID of the user. Cannot be changed. Used to sign in.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Display name for local user.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// New password for user. Used to sign in.
	// Optional, but if included, previous_password is also required.
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	// Previous password for user.
	// Optional, but if included, password is also required.
	PreviousPassword     string   `protobuf:"bytes,4,opt,name=previous_password,json=previousPassword,proto3" json:"previous_password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateSelfReq) Reset()         { *m = UpdateSelfReq{} }
func (m *UpdateSelfReq) String() string { return proto.CompactTextString(m) }
func (*UpdateSelfReq) ProtoMessage()    {}
func (*UpdateSelfReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0d105d392afc5cf, []int{5}
}

func (m *UpdateSelfReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateSelfReq.Unmarshal(m, b)
}
func (m *UpdateSelfReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateSelfReq.Marshal(b, m, deterministic)
}
func (m *UpdateSelfReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateSelfReq.Merge(m, src)
}
func (m *UpdateSelfReq) XXX_Size() int {
	return xxx_messageInfo_UpdateSelfReq.Size(m)
}
func (m *UpdateSelfReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateSelfReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateSelfReq proto.InternalMessageInfo

func (m *UpdateSelfReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateSelfReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateSelfReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *UpdateSelfReq) GetPreviousPassword() string {
	if m != nil {
		return m.PreviousPassword
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateUserReq)(nil), "chef.automate.api.iam.v2.CreateUserReq")
	proto.RegisterType((*ListUsersReq)(nil), "chef.automate.api.iam.v2.ListUsersReq")
	proto.RegisterType((*DeleteUserReq)(nil), "chef.automate.api.iam.v2.DeleteUserReq")
	proto.RegisterType((*GetUserReq)(nil), "chef.automate.api.iam.v2.GetUserReq")
	proto.RegisterType((*UpdateUserReq)(nil), "chef.automate.api.iam.v2.UpdateUserReq")
	proto.RegisterType((*UpdateSelfReq)(nil), "chef.automate.api.iam.v2.UpdateSelfReq")
}

func init() {
	proto.RegisterFile("components/automate-gateway/api/iam/v2/request/users.proto", fileDescriptor_a0d105d392afc5cf)
}

var fileDescriptor_a0d105d392afc5cf = []byte{
	// 422 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xdf, 0x6a, 0xdb, 0x30,
	0x18, 0xc5, 0x71, 0x56, 0xf6, 0x47, 0x2c, 0x65, 0xd3, 0xcd, 0x42, 0x19, 0xcc, 0xe8, 0x6a, 0xb0,
	0xc6, 0x6a, 0xdd, 0x3b, 0x5f, 0x0c, 0xbc, 0x15, 0xca, 0x60, 0xcb, 0xba, 0x94, 0xde, 0xec, 0x66,
	0x7c, 0xb5, 0xbf, 0xb8, 0x02, 0xdb, 0x52, 0x24, 0x39, 0x26, 0x1b, 0x7b, 0x09, 0x3f, 0xd1, 0x9e,
	0x21, 0xaf, 0xb2, 0x17, 0x18, 0xd2, 0x6a, 0x87, 0x0e, 0x32, 0x28, 0xf4, 0x4e, 0x3a, 0xe7, 0xc8,
	0x3e, 0xbf, 0x0f, 0x89, 0x24, 0x99, 0xac, 0x94, 0xac, 0xb1, 0xb6, 0x86, 0x43, 0x63, 0x65, 0x05,
	0x16, 0xa7, 0x05, 0x58, 0x6c, 0x61, 0xcd, 0x41, 0x09, 0x2e, 0xa0, 0xe2, 0xab, 0x98, 0x6b, 0x5c,
	0x36, 0x68, 0x2c, 0x6f, 0x0c, 0x6a, 0x13, 0x29, 0x2d, 0xad, 0xa4, 0x93, 0xec, 0x1a, 0x17, 0x51,
	0x7f, 0x2a, 0x02, 0x25, 0x22, 0x01, 0x55, 0xb4, 0x8a, 0x0f, 0x0e, 0x7d, 0x20, 0x9b, 0x16, 0x58,
	0x4f, 0x4d, 0x0b, 0x45, 0x81, 0x9a, 0x4b, 0x65, 0x85, 0xac, 0x0d, 0x87, 0xba, 0x96, 0x16, 0xfc,
	0xfa, 0xef, 0x77, 0xd8, 0xaf, 0x80, 0x8c, 0xdf, 0x6b, 0x04, 0x8b, 0x97, 0x06, 0xf5, 0x1c, 0x97,
	0x74, 0x9f, 0x8c, 0x44, 0x3e, 0x09, 0xc2, 0xe0, 0xf5, 0x93, 0xf9, 0x48, 0xe4, 0x94, 0x92, 0xbd,
	0x1a, 0x2a, 0x9c, 0x8c, 0xbc, 0xe2, 0xd7, 0xf4, 0x80, 0x3c, 0x56, 0x60, 0x4c, 0x2b, 0x75, 0x3e,
	0x79, 0xe0, 0xf5, 0x61, 0x9f, 0xc8, 0x2e, 0x2d, 0xc9, 0x8b, 0x4d, 0x30, 0x12, 0xf9, 0x26, 0xf0,
	0xf1, 0x4d, 0x30, 0xb8, 0xf1, 0x17, 0xfa, 0xf9, 0x07, 0x73, 0x22, 0x4b, 0x42, 0x76, 0x2a, 0x8c,
	0x2a, 0x61, 0x1d, 0xce, 0xdc, 0xfe, 0x30, 0x64, 0x22, 0x77, 0xaa, 0xc3, 0x74, 0x89, 0xa3, 0xa3,
	0x63, 0xdd, 0x94, 0xf8, 0xdd, 0x39, 0xfd, 0x79, 0xe7, 0xc3, 0x05, 0x2c, 0xf0, 0xbc, 0x17, 0x7e,
	0xb2, 0x7d, 0xf2, 0xf4, 0xa3, 0x30, 0xd6, 0xf5, 0x37, 0x73, 0x5c, 0xb2, 0x57, 0x64, 0x7c, 0x8a,
	0x25, 0xee, 0x24, 0x62, 0x2f, 0x09, 0x39, 0x43, 0xbb, 0xcb, 0xed, 0x02, 0x32, 0xbe, 0x54, 0xf9,
	0x3d, 0x4e, 0x24, 0xed, 0xd2, 0xb7, 0xe4, 0xd1, 0xcd, 0x30, 0xe2, 0x13, 0x7a, 0xbc, 0x9d, 0xc0,
	0x0c, 0xdb, 0x81, 0xfe, 0x7f, 0x8c, 0xbf, 0x87, 0x52, 0x17, 0x58, 0x2e, 0xee, 0xa1, 0x14, 0x7d,
	0x43, 0x9e, 0x2b, 0x8d, 0x2b, 0x21, 0x1b, 0xf3, 0x6d, 0x08, 0xed, 0xf9, 0xd0, 0xb3, 0xde, 0xe8,
	0x0b, 0x24, 0x4d, 0x97, 0xea, 0x2d, 0x41, 0x41, 0x71, 0x4b, 0xf0, 0x69, 0x1d, 0xee, 0x84, 0x98,
	0x61, 0x7b, 0x8b, 0xc3, 0xf9, 0xff, 0xfe, 0xd3, 0x07, 0xcf, 0x6f, 0xd4, 0xdb, 0xd4, 0xef, 0x3e,
	0x7c, 0x3d, 0x2b, 0x84, 0xbd, 0x6e, 0xae, 0xa2, 0x4c, 0x56, 0xdc, 0xdd, 0xf8, 0xe1, 0x9d, 0xf0,
	0xbb, 0xbd, 0x9d, 0xab, 0x87, 0xfe, 0xba, 0x9f, 0xfc, 0x09, 0x00, 0x00, 0xff, 0xff, 0x47, 0xf1,
	0x30, 0x31, 0x74, 0x03, 0x00, 0x00,
}
