// Code generated by protoc-gen-go.
// source: networking_envelopes/networking_envelopes.proto
// DO NOT EDIT!

/*
Package networking_envelopes is a generated protocol buffer package.

It is generated from these files:
	networking_envelopes/networking_envelopes.proto

It has these top-level messages:
	AuthTicket
	RequestEnvelope
	ResponseEnvelope
	Unknown6
	Unknown6Response
*/
package networking_envelopes

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import POGOProtos_Networking_Requests "github.com/zeeraw/pogo-protos/networking_requests"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Request from public import networking_requests/networking_requests.proto
type Request POGOProtos_Networking_Requests.Request

func (m *Request) Reset()         { (*POGOProtos_Networking_Requests.Request)(m).Reset() }
func (m *Request) String() string { return (*POGOProtos_Networking_Requests.Request)(m).String() }
func (*Request) ProtoMessage()    {}

// RequestType from public import networking_requests/networking_requests.proto
type RequestType POGOProtos_Networking_Requests.RequestType

var RequestType_name = POGOProtos_Networking_Requests.RequestType_name
var RequestType_value = POGOProtos_Networking_Requests.RequestType_value

func (x RequestType) String() string { return (POGOProtos_Networking_Requests.RequestType)(x).String() }

const RequestType_METHOD_UNSET = RequestType(POGOProtos_Networking_Requests.RequestType_METHOD_UNSET)
const RequestType_PLAYER_UPDATE = RequestType(POGOProtos_Networking_Requests.RequestType_PLAYER_UPDATE)
const RequestType_GET_PLAYER = RequestType(POGOProtos_Networking_Requests.RequestType_GET_PLAYER)
const RequestType_GET_INVENTORY = RequestType(POGOProtos_Networking_Requests.RequestType_GET_INVENTORY)
const RequestType_DOWNLOAD_SETTINGS = RequestType(POGOProtos_Networking_Requests.RequestType_DOWNLOAD_SETTINGS)
const RequestType_DOWNLOAD_ITEM_TEMPLATES = RequestType(POGOProtos_Networking_Requests.RequestType_DOWNLOAD_ITEM_TEMPLATES)
const RequestType_DOWNLOAD_REMOTE_CONFIG_VERSION = RequestType(POGOProtos_Networking_Requests.RequestType_DOWNLOAD_REMOTE_CONFIG_VERSION)
const RequestType_FORT_SEARCH = RequestType(POGOProtos_Networking_Requests.RequestType_FORT_SEARCH)
const RequestType_ENCOUNTER = RequestType(POGOProtos_Networking_Requests.RequestType_ENCOUNTER)
const RequestType_CATCH_POKEMON = RequestType(POGOProtos_Networking_Requests.RequestType_CATCH_POKEMON)
const RequestType_FORT_DETAILS = RequestType(POGOProtos_Networking_Requests.RequestType_FORT_DETAILS)
const RequestType_ITEM_USE = RequestType(POGOProtos_Networking_Requests.RequestType_ITEM_USE)
const RequestType_GET_MAP_OBJECTS = RequestType(POGOProtos_Networking_Requests.RequestType_GET_MAP_OBJECTS)
const RequestType_FORT_DEPLOY_POKEMON = RequestType(POGOProtos_Networking_Requests.RequestType_FORT_DEPLOY_POKEMON)
const RequestType_FORT_RECALL_POKEMON = RequestType(POGOProtos_Networking_Requests.RequestType_FORT_RECALL_POKEMON)
const RequestType_RELEASE_POKEMON = RequestType(POGOProtos_Networking_Requests.RequestType_RELEASE_POKEMON)
const RequestType_USE_ITEM_POTION = RequestType(POGOProtos_Networking_Requests.RequestType_USE_ITEM_POTION)
const RequestType_USE_ITEM_CAPTURE = RequestType(POGOProtos_Networking_Requests.RequestType_USE_ITEM_CAPTURE)
const RequestType_USE_ITEM_FLEE = RequestType(POGOProtos_Networking_Requests.RequestType_USE_ITEM_FLEE)
const RequestType_USE_ITEM_REVIVE = RequestType(POGOProtos_Networking_Requests.RequestType_USE_ITEM_REVIVE)
const RequestType_TRADE_SEARCH = RequestType(POGOProtos_Networking_Requests.RequestType_TRADE_SEARCH)
const RequestType_TRADE_OFFER = RequestType(POGOProtos_Networking_Requests.RequestType_TRADE_OFFER)
const RequestType_TRADE_RESPONSE = RequestType(POGOProtos_Networking_Requests.RequestType_TRADE_RESPONSE)
const RequestType_TRADE_RESULT = RequestType(POGOProtos_Networking_Requests.RequestType_TRADE_RESULT)
const RequestType_GET_PLAYER_PROFILE = RequestType(POGOProtos_Networking_Requests.RequestType_GET_PLAYER_PROFILE)
const RequestType_GET_ITEM_PACK = RequestType(POGOProtos_Networking_Requests.RequestType_GET_ITEM_PACK)
const RequestType_BUY_ITEM_PACK = RequestType(POGOProtos_Networking_Requests.RequestType_BUY_ITEM_PACK)
const RequestType_BUY_GEM_PACK = RequestType(POGOProtos_Networking_Requests.RequestType_BUY_GEM_PACK)
const RequestType_EVOLVE_POKEMON = RequestType(POGOProtos_Networking_Requests.RequestType_EVOLVE_POKEMON)
const RequestType_GET_HATCHED_EGGS = RequestType(POGOProtos_Networking_Requests.RequestType_GET_HATCHED_EGGS)
const RequestType_ENCOUNTER_TUTORIAL_COMPLETE = RequestType(POGOProtos_Networking_Requests.RequestType_ENCOUNTER_TUTORIAL_COMPLETE)
const RequestType_LEVEL_UP_REWARDS = RequestType(POGOProtos_Networking_Requests.RequestType_LEVEL_UP_REWARDS)
const RequestType_CHECK_AWARDED_BADGES = RequestType(POGOProtos_Networking_Requests.RequestType_CHECK_AWARDED_BADGES)
const RequestType_USE_ITEM_GYM = RequestType(POGOProtos_Networking_Requests.RequestType_USE_ITEM_GYM)
const RequestType_GET_GYM_DETAILS = RequestType(POGOProtos_Networking_Requests.RequestType_GET_GYM_DETAILS)
const RequestType_START_GYM_BATTLE = RequestType(POGOProtos_Networking_Requests.RequestType_START_GYM_BATTLE)
const RequestType_ATTACK_GYM = RequestType(POGOProtos_Networking_Requests.RequestType_ATTACK_GYM)
const RequestType_RECYCLE_INVENTORY_ITEM = RequestType(POGOProtos_Networking_Requests.RequestType_RECYCLE_INVENTORY_ITEM)
const RequestType_COLLECT_DAILY_BONUS = RequestType(POGOProtos_Networking_Requests.RequestType_COLLECT_DAILY_BONUS)
const RequestType_USE_ITEM_XP_BOOST = RequestType(POGOProtos_Networking_Requests.RequestType_USE_ITEM_XP_BOOST)
const RequestType_USE_ITEM_EGG_INCUBATOR = RequestType(POGOProtos_Networking_Requests.RequestType_USE_ITEM_EGG_INCUBATOR)
const RequestType_USE_INCENSE = RequestType(POGOProtos_Networking_Requests.RequestType_USE_INCENSE)
const RequestType_GET_INCENSE_POKEMON = RequestType(POGOProtos_Networking_Requests.RequestType_GET_INCENSE_POKEMON)
const RequestType_INCENSE_ENCOUNTER = RequestType(POGOProtos_Networking_Requests.RequestType_INCENSE_ENCOUNTER)
const RequestType_ADD_FORT_MODIFIER = RequestType(POGOProtos_Networking_Requests.RequestType_ADD_FORT_MODIFIER)
const RequestType_DISK_ENCOUNTER = RequestType(POGOProtos_Networking_Requests.RequestType_DISK_ENCOUNTER)
const RequestType_COLLECT_DAILY_DEFENDER_BONUS = RequestType(POGOProtos_Networking_Requests.RequestType_COLLECT_DAILY_DEFENDER_BONUS)
const RequestType_UPGRADE_POKEMON = RequestType(POGOProtos_Networking_Requests.RequestType_UPGRADE_POKEMON)
const RequestType_SET_FAVORITE_POKEMON = RequestType(POGOProtos_Networking_Requests.RequestType_SET_FAVORITE_POKEMON)
const RequestType_NICKNAME_POKEMON = RequestType(POGOProtos_Networking_Requests.RequestType_NICKNAME_POKEMON)
const RequestType_EQUIP_BADGE = RequestType(POGOProtos_Networking_Requests.RequestType_EQUIP_BADGE)
const RequestType_SET_CONTACT_SETTINGS = RequestType(POGOProtos_Networking_Requests.RequestType_SET_CONTACT_SETTINGS)
const RequestType_GET_ASSET_DIGEST = RequestType(POGOProtos_Networking_Requests.RequestType_GET_ASSET_DIGEST)
const RequestType_GET_DOWNLOAD_URLS = RequestType(POGOProtos_Networking_Requests.RequestType_GET_DOWNLOAD_URLS)
const RequestType_GET_SUGGESTED_CODENAMES = RequestType(POGOProtos_Networking_Requests.RequestType_GET_SUGGESTED_CODENAMES)
const RequestType_CHECK_CODENAME_AVAILABLE = RequestType(POGOProtos_Networking_Requests.RequestType_CHECK_CODENAME_AVAILABLE)
const RequestType_CLAIM_CODENAME = RequestType(POGOProtos_Networking_Requests.RequestType_CLAIM_CODENAME)
const RequestType_SET_AVATAR = RequestType(POGOProtos_Networking_Requests.RequestType_SET_AVATAR)
const RequestType_SET_PLAYER_TEAM = RequestType(POGOProtos_Networking_Requests.RequestType_SET_PLAYER_TEAM)
const RequestType_MARK_TUTORIAL_COMPLETE = RequestType(POGOProtos_Networking_Requests.RequestType_MARK_TUTORIAL_COMPLETE)
const RequestType_LOAD_SPAWN_POINTS = RequestType(POGOProtos_Networking_Requests.RequestType_LOAD_SPAWN_POINTS)
const RequestType_ECHO = RequestType(POGOProtos_Networking_Requests.RequestType_ECHO)
const RequestType_DEBUG_UPDATE_INVENTORY = RequestType(POGOProtos_Networking_Requests.RequestType_DEBUG_UPDATE_INVENTORY)
const RequestType_DEBUG_DELETE_PLAYER = RequestType(POGOProtos_Networking_Requests.RequestType_DEBUG_DELETE_PLAYER)
const RequestType_SFIDA_REGISTRATION = RequestType(POGOProtos_Networking_Requests.RequestType_SFIDA_REGISTRATION)
const RequestType_SFIDA_ACTION_LOG = RequestType(POGOProtos_Networking_Requests.RequestType_SFIDA_ACTION_LOG)
const RequestType_SFIDA_CERTIFICATION = RequestType(POGOProtos_Networking_Requests.RequestType_SFIDA_CERTIFICATION)
const RequestType_SFIDA_UPDATE = RequestType(POGOProtos_Networking_Requests.RequestType_SFIDA_UPDATE)
const RequestType_SFIDA_ACTION = RequestType(POGOProtos_Networking_Requests.RequestType_SFIDA_ACTION)
const RequestType_SFIDA_DOWSER = RequestType(POGOProtos_Networking_Requests.RequestType_SFIDA_DOWSER)
const RequestType_SFIDA_CAPTURE = RequestType(POGOProtos_Networking_Requests.RequestType_SFIDA_CAPTURE)

type AuthTicket struct {
	Start             []byte `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	ExpireTimestampMs uint64 `protobuf:"varint,2,opt,name=expire_timestamp_ms,json=expireTimestampMs" json:"expire_timestamp_ms,omitempty"`
	End               []byte `protobuf:"bytes,3,opt,name=end,proto3" json:"end,omitempty"`
}

func (m *AuthTicket) Reset()                    { *m = AuthTicket{} }
func (m *AuthTicket) String() string            { return proto.CompactTextString(m) }
func (*AuthTicket) ProtoMessage()               {}
func (*AuthTicket) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type RequestEnvelope struct {
	StatusCode int32                                     `protobuf:"varint,1,opt,name=status_code,json=statusCode" json:"status_code,omitempty"`
	RequestId  uint64                                    `protobuf:"varint,3,opt,name=request_id,json=requestId" json:"request_id,omitempty"`
	Requests   []*POGOProtos_Networking_Requests.Request `protobuf:"bytes,4,rep,name=requests" json:"requests,omitempty"`
	Unknown6   *Unknown6                                 `protobuf:"bytes,6,opt,name=unknown6" json:"unknown6,omitempty"`
	Latitude   float64                                   `protobuf:"fixed64,7,opt,name=latitude" json:"latitude,omitempty"`
	Longitude  float64                                   `protobuf:"fixed64,8,opt,name=longitude" json:"longitude,omitempty"`
	Altitude   float64                                   `protobuf:"fixed64,9,opt,name=altitude" json:"altitude,omitempty"`
	AuthInfo   *RequestEnvelope_AuthInfo                 `protobuf:"bytes,10,opt,name=auth_info,json=authInfo" json:"auth_info,omitempty"`
	AuthTicket *AuthTicket                               `protobuf:"bytes,11,opt,name=auth_ticket,json=authTicket" json:"auth_ticket,omitempty"`
	Unknown12  int64                                     `protobuf:"varint,12,opt,name=unknown12" json:"unknown12,omitempty"`
}

func (m *RequestEnvelope) Reset()                    { *m = RequestEnvelope{} }
func (m *RequestEnvelope) String() string            { return proto.CompactTextString(m) }
func (*RequestEnvelope) ProtoMessage()               {}
func (*RequestEnvelope) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RequestEnvelope) GetRequests() []*POGOProtos_Networking_Requests.Request {
	if m != nil {
		return m.Requests
	}
	return nil
}

func (m *RequestEnvelope) GetUnknown6() *Unknown6 {
	if m != nil {
		return m.Unknown6
	}
	return nil
}

func (m *RequestEnvelope) GetAuthInfo() *RequestEnvelope_AuthInfo {
	if m != nil {
		return m.AuthInfo
	}
	return nil
}

func (m *RequestEnvelope) GetAuthTicket() *AuthTicket {
	if m != nil {
		return m.AuthTicket
	}
	return nil
}

type RequestEnvelope_AuthInfo struct {
	Provider string                        `protobuf:"bytes,1,opt,name=provider" json:"provider,omitempty"`
	Token    *RequestEnvelope_AuthInfo_JWT `protobuf:"bytes,2,opt,name=token" json:"token,omitempty"`
}

func (m *RequestEnvelope_AuthInfo) Reset()                    { *m = RequestEnvelope_AuthInfo{} }
func (m *RequestEnvelope_AuthInfo) String() string            { return proto.CompactTextString(m) }
func (*RequestEnvelope_AuthInfo) ProtoMessage()               {}
func (*RequestEnvelope_AuthInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *RequestEnvelope_AuthInfo) GetToken() *RequestEnvelope_AuthInfo_JWT {
	if m != nil {
		return m.Token
	}
	return nil
}

type RequestEnvelope_AuthInfo_JWT struct {
	Contents string `protobuf:"bytes,1,opt,name=contents" json:"contents,omitempty"`
	Unknown2 int32  `protobuf:"varint,2,opt,name=unknown2" json:"unknown2,omitempty"`
}

func (m *RequestEnvelope_AuthInfo_JWT) Reset()         { *m = RequestEnvelope_AuthInfo_JWT{} }
func (m *RequestEnvelope_AuthInfo_JWT) String() string { return proto.CompactTextString(m) }
func (*RequestEnvelope_AuthInfo_JWT) ProtoMessage()    {}
func (*RequestEnvelope_AuthInfo_JWT) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{1, 0, 0}
}

type ResponseEnvelope struct {
	StatusCode int32             `protobuf:"varint,1,opt,name=status_code,json=statusCode" json:"status_code,omitempty"`
	RequestId  uint64            `protobuf:"varint,2,opt,name=request_id,json=requestId" json:"request_id,omitempty"`
	ApiUrl     string            `protobuf:"bytes,3,opt,name=api_url,json=apiUrl" json:"api_url,omitempty"`
	Unknown6   *Unknown6Response `protobuf:"bytes,6,opt,name=unknown6" json:"unknown6,omitempty"`
	AuthTicket *AuthTicket       `protobuf:"bytes,7,opt,name=auth_ticket,json=authTicket" json:"auth_ticket,omitempty"`
	Returns    [][]byte          `protobuf:"bytes,100,rep,name=returns,proto3" json:"returns,omitempty"`
	Error      string            `protobuf:"bytes,101,opt,name=error" json:"error,omitempty"`
}

func (m *ResponseEnvelope) Reset()                    { *m = ResponseEnvelope{} }
func (m *ResponseEnvelope) String() string            { return proto.CompactTextString(m) }
func (*ResponseEnvelope) ProtoMessage()               {}
func (*ResponseEnvelope) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ResponseEnvelope) GetUnknown6() *Unknown6Response {
	if m != nil {
		return m.Unknown6
	}
	return nil
}

func (m *ResponseEnvelope) GetAuthTicket() *AuthTicket {
	if m != nil {
		return m.AuthTicket
	}
	return nil
}

type ResponseEnvelope_Unknown7 struct {
	Unknown71 []byte `protobuf:"bytes,1,opt,name=unknown71,proto3" json:"unknown71,omitempty"`
	Unknown72 int64  `protobuf:"varint,2,opt,name=unknown72" json:"unknown72,omitempty"`
	Unknown73 []byte `protobuf:"bytes,3,opt,name=unknown73,proto3" json:"unknown73,omitempty"`
}

func (m *ResponseEnvelope_Unknown7) Reset()                    { *m = ResponseEnvelope_Unknown7{} }
func (m *ResponseEnvelope_Unknown7) String() string            { return proto.CompactTextString(m) }
func (*ResponseEnvelope_Unknown7) ProtoMessage()               {}
func (*ResponseEnvelope_Unknown7) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

type Unknown6 struct {
	Unknown1 int32              `protobuf:"varint,1,opt,name=unknown1" json:"unknown1,omitempty"`
	Unknown2 *Unknown6_Unknown2 `protobuf:"bytes,2,opt,name=unknown2" json:"unknown2,omitempty"`
}

func (m *Unknown6) Reset()                    { *m = Unknown6{} }
func (m *Unknown6) String() string            { return proto.CompactTextString(m) }
func (*Unknown6) ProtoMessage()               {}
func (*Unknown6) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Unknown6) GetUnknown2() *Unknown6_Unknown2 {
	if m != nil {
		return m.Unknown2
	}
	return nil
}

type Unknown6_Unknown2 struct {
	Unknown1 []byte `protobuf:"bytes,1,opt,name=unknown1,proto3" json:"unknown1,omitempty"`
}

func (m *Unknown6_Unknown2) Reset()                    { *m = Unknown6_Unknown2{} }
func (m *Unknown6_Unknown2) String() string            { return proto.CompactTextString(m) }
func (*Unknown6_Unknown2) ProtoMessage()               {}
func (*Unknown6_Unknown2) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

type Unknown6Response struct {
	Unknown1 int32                      `protobuf:"varint,1,opt,name=unknown1" json:"unknown1,omitempty"`
	Unknown2 *Unknown6Response_Unknown2 `protobuf:"bytes,2,opt,name=unknown2" json:"unknown2,omitempty"`
}

func (m *Unknown6Response) Reset()                    { *m = Unknown6Response{} }
func (m *Unknown6Response) String() string            { return proto.CompactTextString(m) }
func (*Unknown6Response) ProtoMessage()               {}
func (*Unknown6Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Unknown6Response) GetUnknown2() *Unknown6Response_Unknown2 {
	if m != nil {
		return m.Unknown2
	}
	return nil
}

type Unknown6Response_Unknown2 struct {
	Unknown1 uint64 `protobuf:"varint,1,opt,name=unknown1" json:"unknown1,omitempty"`
}

func (m *Unknown6Response_Unknown2) Reset()                    { *m = Unknown6Response_Unknown2{} }
func (m *Unknown6Response_Unknown2) String() string            { return proto.CompactTextString(m) }
func (*Unknown6Response_Unknown2) ProtoMessage()               {}
func (*Unknown6Response_Unknown2) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4, 0} }

func init() {
	proto.RegisterType((*AuthTicket)(nil), "POGOProtos.Networking.Envelopes.AuthTicket")
	proto.RegisterType((*RequestEnvelope)(nil), "POGOProtos.Networking.Envelopes.RequestEnvelope")
	proto.RegisterType((*RequestEnvelope_AuthInfo)(nil), "POGOProtos.Networking.Envelopes.RequestEnvelope.AuthInfo")
	proto.RegisterType((*RequestEnvelope_AuthInfo_JWT)(nil), "POGOProtos.Networking.Envelopes.RequestEnvelope.AuthInfo.JWT")
	proto.RegisterType((*ResponseEnvelope)(nil), "POGOProtos.Networking.Envelopes.ResponseEnvelope")
	proto.RegisterType((*ResponseEnvelope_Unknown7)(nil), "POGOProtos.Networking.Envelopes.ResponseEnvelope.Unknown7")
	proto.RegisterType((*Unknown6)(nil), "POGOProtos.Networking.Envelopes.Unknown6")
	proto.RegisterType((*Unknown6_Unknown2)(nil), "POGOProtos.Networking.Envelopes.Unknown6.Unknown2")
	proto.RegisterType((*Unknown6Response)(nil), "POGOProtos.Networking.Envelopes.Unknown6Response")
	proto.RegisterType((*Unknown6Response_Unknown2)(nil), "POGOProtos.Networking.Envelopes.Unknown6Response.Unknown2")
}

func init() { proto.RegisterFile("networking_envelopes/networking_envelopes.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 629 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x55, 0x5d, 0x4f, 0xd4, 0x40,
	0x14, 0xb5, 0x74, 0x97, 0x6d, 0x6f, 0x37, 0x11, 0x47, 0xa2, 0x4d, 0xa3, 0x81, 0xec, 0x83, 0x62,
	0x8c, 0x35, 0x5b, 0x12, 0x88, 0x26, 0x3c, 0x28, 0x21, 0x06, 0x23, 0x1f, 0x19, 0x01, 0x13, 0x5f,
	0x9a, 0x4a, 0x47, 0x6c, 0xb6, 0xcc, 0xd4, 0xe9, 0x14, 0xfc, 0x35, 0xfc, 0x04, 0x9f, 0x7c, 0xf2,
	0xd7, 0x39, 0x9d, 0x4e, 0x3f, 0x5c, 0x88, 0x0b, 0xf8, 0x36, 0xf7, 0xde, 0xde, 0x73, 0xcf, 0x9d,
	0x73, 0x76, 0x16, 0x5e, 0x52, 0x22, 0xce, 0x19, 0x9f, 0x24, 0xf4, 0x24, 0x24, 0xf4, 0x8c, 0xa4,
	0x2c, 0x23, 0xf9, 0x95, 0x49, 0x3f, 0xe3, 0x4c, 0x30, 0xb4, 0xb4, 0xbf, 0xf7, 0x6e, 0x6f, 0xbf,
	0x3c, 0xe6, 0xfe, 0x6e, 0xf3, 0x99, 0xbf, 0x55, 0x7f, 0xe6, 0xbd, 0xe8, 0x34, 0x73, 0xf2, 0xbd,
	0x20, 0xb9, 0xf8, 0x0b, 0xb0, 0xce, 0x55, 0x78, 0xa3, 0x18, 0xe0, 0x4d, 0x21, 0xbe, 0x1d, 0x24,
	0xc7, 0x13, 0x22, 0xd0, 0x22, 0xf4, 0x73, 0x11, 0x71, 0xe1, 0x1a, 0xcb, 0xc6, 0xca, 0x10, 0x57,
	0x01, 0xf2, 0xe1, 0x3e, 0xf9, 0x91, 0x25, 0x9c, 0x84, 0x22, 0x39, 0x95, 0xbd, 0xd1, 0x69, 0x16,
	0x9e, 0xe6, 0xee, 0x9c, 0xfc, 0xa6, 0x87, 0xef, 0x55, 0xa5, 0x83, 0xba, 0xb2, 0x93, 0xa3, 0x05,
	0x30, 0x09, 0x8d, 0x5d, 0x53, 0x61, 0x94, 0xc7, 0xd1, 0xaf, 0x3e, 0xdc, 0xc5, 0xd5, 0xe0, 0x9a,
	0x29, 0x5a, 0x02, 0x47, 0x36, 0x88, 0x22, 0x0f, 0x8f, 0x59, 0x4c, 0xd4, 0xc4, 0x3e, 0x86, 0x2a,
	0xb5, 0x29, 0x33, 0xe8, 0x31, 0x80, 0x26, 0x1b, 0x26, 0x15, 0x5a, 0x0f, 0xdb, 0x3a, 0xb3, 0x1d,
	0xa3, 0x4d, 0xb0, 0xea, 0x5d, 0xdc, 0xde, 0xb2, 0xb9, 0xe2, 0x04, 0x4f, 0xfd, 0xab, 0x2f, 0x07,
	0xd7, 0x2b, 0xeb, 0x03, 0x6e, 0x1a, 0xd1, 0x16, 0x58, 0x05, 0x9d, 0x50, 0x76, 0x4e, 0xd7, 0xdc,
	0x79, 0x39, 0xc1, 0x09, 0x9e, 0xf9, 0x33, 0x6e, 0xd8, 0x3f, 0xd4, 0x0d, 0xb8, 0x69, 0x45, 0x1e,
	0x58, 0x69, 0x24, 0x12, 0x51, 0xc8, 0x45, 0x06, 0x12, 0xc6, 0xc0, 0x4d, 0x8c, 0x1e, 0x81, 0x9d,
	0x32, 0x7a, 0x52, 0x15, 0x2d, 0x55, 0x6c, 0x13, 0x65, 0x67, 0x94, 0xea, 0x4e, 0xbb, 0xea, 0xac,
	0x63, 0x74, 0x04, 0x76, 0x24, 0xb5, 0x09, 0x13, 0xfa, 0x95, 0xb9, 0xa0, 0xd8, 0xbd, 0x9a, 0xc9,
	0x6e, 0xea, 0x9a, 0xfd, 0x52, 0xdd, 0x6d, 0x09, 0x20, 0x71, 0xf5, 0x09, 0x7d, 0x00, 0x47, 0xe1,
	0x0a, 0x25, 0xba, 0xeb, 0x28, 0xe4, 0xe7, 0x33, 0x91, 0x5b, 0x9f, 0x60, 0x88, 0x5a, 0xcf, 0xc8,
	0xfd, 0xf4, 0x3d, 0x8c, 0x03, 0x77, 0x28, 0xb1, 0x4c, 0xdc, 0x26, 0xbc, 0xdf, 0x06, 0x58, 0x35,
	0x85, 0x72, 0x59, 0xe9, 0xba, 0xb3, 0x24, 0x26, 0x5c, 0xe9, 0x6d, 0xe3, 0x26, 0x46, 0x1f, 0xa1,
	0x2f, 0xd8, 0x84, 0x50, 0x65, 0x2b, 0x27, 0xd8, 0xb8, 0xf5, 0xa2, 0xfe, 0xfb, 0x4f, 0x07, 0xb8,
	0xc2, 0xf2, 0x36, 0xc0, 0x94, 0x51, 0x39, 0xf7, 0x98, 0x51, 0x41, 0xa8, 0xb4, 0x8a, 0x9e, 0x5b,
	0xc7, 0x65, 0x4d, 0xb3, 0x0d, 0xd4, 0xe8, 0x7e, 0x23, 0x6b, 0x30, 0xba, 0x30, 0x61, 0x01, 0x93,
	0x3c, 0x63, 0x34, 0x27, 0xb7, 0xf5, 0xed, 0xdc, 0xb4, 0x6f, 0x1f, 0xc2, 0x20, 0xca, 0x92, 0xb0,
	0xe0, 0xa9, 0xf2, 0xb4, 0x8d, 0xe7, 0x65, 0x78, 0xc8, 0x53, 0xb4, 0x73, 0xc9, 0x8b, 0xe3, 0xeb,
	0x7b, 0x51, 0xb3, 0xec, 0x78, 0x72, 0x4a, 0xe5, 0xc1, 0xff, 0xa9, 0xec, 0xc2, 0x80, 0x13, 0x51,
	0x70, 0x9a, 0xbb, 0xb1, 0xfc, 0xb1, 0x0d, 0x71, 0x1d, 0x96, 0x6f, 0x06, 0xe1, 0x9c, 0x71, 0x97,
	0xa8, 0x6d, 0xaa, 0xc0, 0x8b, 0xc1, 0xd2, 0xdc, 0xd6, 0x3b, 0x0e, 0x59, 0x1f, 0xeb, 0x97, 0xa5,
	0x4d, 0x74, 0xab, 0x95, 0x02, 0xad, 0x7f, 0xd6, 0x83, 0x6e, 0x75, 0x55, 0xbf, 0x28, 0x6d, 0x62,
	0x74, 0x61, 0x34, 0x63, 0xd6, 0x3a, 0x4a, 0x8e, 0xb5, 0x2a, 0x4d, 0x8c, 0x76, 0xa7, 0x54, 0x76,
	0x82, 0xe0, 0xda, 0x77, 0x5b, 0x1f, 0x82, 0xd6, 0x19, 0xde, 0x93, 0x66, 0x6e, 0x70, 0x69, 0xee,
	0xb0, 0x9d, 0x3b, 0xfa, 0x69, 0xc0, 0xc2, 0xb4, 0x46, 0xff, 0x24, 0x7a, 0x74, 0x89, 0xe8, 0xeb,
	0x1b, 0x9b, 0xe0, 0xa6, 0x84, 0x7b, 0xed, 0xfc, 0xb7, 0x0f, 0x3e, 0x2f, 0x5e, 0xf5, 0xef, 0xb3,
	0x7f, 0xe7, 0xcb, 0xbc, 0xfa, 0xc3, 0x58, 0xfd, 0x13, 0x00, 0x00, 0xff, 0xff, 0xde, 0xac, 0xf8,
	0xfe, 0xb3, 0x06, 0x00, 0x00,
}
