// Code generated by protoc-gen-go.
// source: token.proto
// DO NOT EDIT!

package igrpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type OfficialAccountPlatform struct {
	Appid                 string `protobuf:"bytes,1,opt,name=Appid" json:"Appid,omitempty"`
	ComponentVerifyTicket string `protobuf:"bytes,2,opt,name=ComponentVerifyTicket" json:"ComponentVerifyTicket,omitempty"`
	EncodingAesKey        string `protobuf:"bytes,3,opt,name=EncodingAesKey" json:"EncodingAesKey,omitempty"`
	AppSecret             string `protobuf:"bytes,4,opt,name=AppSecret" json:"AppSecret,omitempty"`
	ComponentAccessToken  string `protobuf:"bytes,5,opt,name=ComponentAccessToken" json:"ComponentAccessToken,omitempty"`
	PreAuthCode           string `protobuf:"bytes,6,opt,name=PreAuthCode" json:"PreAuthCode,omitempty"`
	Token                 string `protobuf:"bytes,7,opt,name=Token" json:"Token,omitempty"`
}

func (m *OfficialAccountPlatform) Reset()                    { *m = OfficialAccountPlatform{} }
func (m *OfficialAccountPlatform) String() string            { return proto.CompactTextString(m) }
func (*OfficialAccountPlatform) ProtoMessage()               {}
func (*OfficialAccountPlatform) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{0} }

func (m *OfficialAccountPlatform) GetAppid() string {
	if m != nil {
		return m.Appid
	}
	return ""
}

func (m *OfficialAccountPlatform) GetComponentVerifyTicket() string {
	if m != nil {
		return m.ComponentVerifyTicket
	}
	return ""
}

func (m *OfficialAccountPlatform) GetEncodingAesKey() string {
	if m != nil {
		return m.EncodingAesKey
	}
	return ""
}

func (m *OfficialAccountPlatform) GetAppSecret() string {
	if m != nil {
		return m.AppSecret
	}
	return ""
}

func (m *OfficialAccountPlatform) GetComponentAccessToken() string {
	if m != nil {
		return m.ComponentAccessToken
	}
	return ""
}

func (m *OfficialAccountPlatform) GetPreAuthCode() string {
	if m != nil {
		return m.PreAuthCode
	}
	return ""
}

func (m *OfficialAccountPlatform) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type OfficialAccount struct {
	Appid                          string `protobuf:"bytes,1,opt,name=Appid" json:"Appid,omitempty"`
	AuthorizerAccessToken          string `protobuf:"bytes,2,opt,name=AuthorizerAccessToken" json:"AuthorizerAccessToken,omitempty"`
	AuthorizerAccessTokenExpiresIn int64  `protobuf:"varint,3,opt,name=AuthorizerAccessTokenExpiresIn" json:"AuthorizerAccessTokenExpiresIn,omitempty"`
	AuthorizerRefreshToken         string `protobuf:"bytes,4,opt,name=AuthorizerRefreshToken" json:"AuthorizerRefreshToken,omitempty"`
}

func (m *OfficialAccount) Reset()                    { *m = OfficialAccount{} }
func (m *OfficialAccount) String() string            { return proto.CompactTextString(m) }
func (*OfficialAccount) ProtoMessage()               {}
func (*OfficialAccount) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{1} }

func (m *OfficialAccount) GetAppid() string {
	if m != nil {
		return m.Appid
	}
	return ""
}

func (m *OfficialAccount) GetAuthorizerAccessToken() string {
	if m != nil {
		return m.AuthorizerAccessToken
	}
	return ""
}

func (m *OfficialAccount) GetAuthorizerAccessTokenExpiresIn() int64 {
	if m != nil {
		return m.AuthorizerAccessTokenExpiresIn
	}
	return 0
}

func (m *OfficialAccount) GetAuthorizerRefreshToken() string {
	if m != nil {
		return m.AuthorizerRefreshToken
	}
	return ""
}

type ComponentVerifyTicket struct {
	TimeStamp         string `protobuf:"bytes,1,opt,name=TimeStamp" json:"TimeStamp,omitempty"`
	Nonce             string `protobuf:"bytes,2,opt,name=Nonce" json:"Nonce,omitempty"`
	MsgSign           string `protobuf:"bytes,3,opt,name=MsgSign" json:"MsgSign,omitempty"`
	Bts               []byte `protobuf:"bytes,4,opt,name=bts,proto3" json:"bts,omitempty"`
	AuthorizationCode string `protobuf:"bytes,5,opt,name=AuthorizationCode" json:"AuthorizationCode,omitempty"`
}

func (m *ComponentVerifyTicket) Reset()                    { *m = ComponentVerifyTicket{} }
func (m *ComponentVerifyTicket) String() string            { return proto.CompactTextString(m) }
func (*ComponentVerifyTicket) ProtoMessage()               {}
func (*ComponentVerifyTicket) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{2} }

func (m *ComponentVerifyTicket) GetTimeStamp() string {
	if m != nil {
		return m.TimeStamp
	}
	return ""
}

func (m *ComponentVerifyTicket) GetNonce() string {
	if m != nil {
		return m.Nonce
	}
	return ""
}

func (m *ComponentVerifyTicket) GetMsgSign() string {
	if m != nil {
		return m.MsgSign
	}
	return ""
}

func (m *ComponentVerifyTicket) GetBts() []byte {
	if m != nil {
		return m.Bts
	}
	return nil
}

func (m *ComponentVerifyTicket) GetAuthorizationCode() string {
	if m != nil {
		return m.AuthorizationCode
	}
	return ""
}

func init() {
	proto.RegisterType((*OfficialAccountPlatform)(nil), "igrpc.OfficialAccountPlatform")
	proto.RegisterType((*OfficialAccount)(nil), "igrpc.OfficialAccount")
	proto.RegisterType((*ComponentVerifyTicket)(nil), "igrpc.ComponentVerifyTicket")
}

func init() { proto.RegisterFile("token.proto", fileDescriptor11) }

var fileDescriptor11 = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x52, 0xcd, 0x4e, 0x32, 0x31,
	0x14, 0xcd, 0xc0, 0x07, 0x84, 0xcb, 0x17, 0x7f, 0x1a, 0xd4, 0x2e, 0x8c, 0x21, 0x2c, 0x8c, 0x0b,
	0xe3, 0x42, 0x8d, 0xfb, 0x09, 0xc1, 0xc4, 0x18, 0x95, 0x0c, 0xc4, 0xfd, 0x50, 0xee, 0x0c, 0x0d,
	0x4c, 0xdb, 0xb4, 0x25, 0x11, 0x5f, 0xc4, 0x27, 0xf0, 0xb1, 0x7c, 0x17, 0x33, 0xed, 0x20, 0x04,
	0x87, 0xb8, 0x9b, 0x7b, 0xce, 0xc9, 0x9d, 0x7b, 0xce, 0x29, 0xb4, 0xac, 0x9c, 0xa1, 0xb8, 0x52,
	0x5a, 0x5a, 0x49, 0x6a, 0x3c, 0xd5, 0x8a, 0x75, 0x3f, 0x2a, 0x70, 0xf2, 0x92, 0x24, 0x9c, 0xf1,
	0x78, 0x1e, 0x32, 0x26, 0x17, 0xc2, 0x0e, 0xe6, 0xb1, 0x4d, 0xa4, 0xce, 0x48, 0x1b, 0x6a, 0xa1,
	0x52, 0x7c, 0x42, 0x83, 0x4e, 0x70, 0xd1, 0x8c, 0xfc, 0x40, 0x6e, 0xe1, 0xa8, 0x27, 0x33, 0x25,
	0x05, 0x0a, 0xfb, 0x8a, 0x9a, 0x27, 0xcb, 0x11, 0x67, 0x33, 0xb4, 0xb4, 0xe2, 0x54, 0xe5, 0x24,
	0x39, 0x87, 0xbd, 0xbe, 0x60, 0x72, 0xc2, 0x45, 0x1a, 0xa2, 0x79, 0xc4, 0x25, 0xad, 0x3a, 0xf9,
	0x16, 0x4a, 0x4e, 0xa1, 0x19, 0x2a, 0x35, 0x44, 0xa6, 0xd1, 0xd2, 0x7f, 0x4e, 0xb2, 0x06, 0xc8,
	0x35, 0xb4, 0x7f, 0xd6, 0x87, 0x8c, 0xa1, 0x31, 0xa3, 0xdc, 0x12, 0xad, 0x39, 0x61, 0x29, 0x47,
	0x3a, 0xd0, 0x1a, 0x68, 0x0c, 0x17, 0x76, 0xda, 0x93, 0x13, 0xa4, 0x75, 0x27, 0xdd, 0x84, 0x72,
	0x9f, 0x7e, 0x4d, 0xc3, 0xfb, 0x74, 0x43, 0xf7, 0x2b, 0x80, 0xfd, 0xad, 0x64, 0x76, 0x27, 0x92,
	0xef, 0x92, 0x9a, 0xbf, 0xa3, 0xde, 0x3c, 0xab, 0x48, 0xa4, 0x94, 0x24, 0xf7, 0x70, 0x56, 0x4a,
	0xf4, 0xdf, 0x14, 0xd7, 0x68, 0x1e, 0x84, 0x4b, 0xa8, 0x1a, 0xfd, 0xa1, 0x22, 0x77, 0x70, 0xbc,
	0x56, 0x44, 0x98, 0x68, 0x34, 0x53, 0xff, 0x7b, 0x1f, 0xdf, 0x0e, 0xb6, 0xfb, 0x19, 0xec, 0x28,
	0x32, 0xef, 0x60, 0xc4, 0x33, 0x1c, 0xda, 0x38, 0x53, 0x85, 0xd3, 0x35, 0x90, 0x67, 0xf0, 0x2c,
	0x05, 0xc3, 0xc2, 0x9d, 0x1f, 0x08, 0x85, 0xc6, 0x93, 0x49, 0x87, 0x3c, 0x15, 0x45, 0xb1, 0xab,
	0x91, 0x1c, 0x40, 0x75, 0x6c, 0x8d, 0x3b, 0xe6, 0x7f, 0x94, 0x7f, 0x92, 0x4b, 0x38, 0x5c, 0xdd,
	0x14, 0x5b, 0x2e, 0x85, 0xeb, 0xc5, 0x57, 0xf8, 0x9b, 0x18, 0xd7, 0xdd, 0x7b, 0xbd, 0xf9, 0x0e,
	0x00, 0x00, 0xff, 0xff, 0x6c, 0xf3, 0xeb, 0xfa, 0xbe, 0x02, 0x00, 0x00,
}