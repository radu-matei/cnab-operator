// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/geo_targeting_restriction.proto

package enums // import "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A restriction used to determine if the request context's
// geo should be matched.
type GeoTargetingRestrictionEnum_GeoTargetingRestriction int32

const (
	// Not specified.
	GeoTargetingRestrictionEnum_UNSPECIFIED GeoTargetingRestrictionEnum_GeoTargetingRestriction = 0
	// Used for return value only. Represents value unknown in this version.
	GeoTargetingRestrictionEnum_UNKNOWN GeoTargetingRestrictionEnum_GeoTargetingRestriction = 1
	// Indicates that request context should match the physical location of
	// the user.
	GeoTargetingRestrictionEnum_LOCATION_OF_PRESENCE GeoTargetingRestrictionEnum_GeoTargetingRestriction = 2
)

var GeoTargetingRestrictionEnum_GeoTargetingRestriction_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "LOCATION_OF_PRESENCE",
}
var GeoTargetingRestrictionEnum_GeoTargetingRestriction_value = map[string]int32{
	"UNSPECIFIED":          0,
	"UNKNOWN":              1,
	"LOCATION_OF_PRESENCE": 2,
}

func (x GeoTargetingRestrictionEnum_GeoTargetingRestriction) String() string {
	return proto.EnumName(GeoTargetingRestrictionEnum_GeoTargetingRestriction_name, int32(x))
}
func (GeoTargetingRestrictionEnum_GeoTargetingRestriction) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_geo_targeting_restriction_40a83d06ddfa34b9, []int{0, 0}
}

// Message describing feed item geo targeting restriction.
type GeoTargetingRestrictionEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GeoTargetingRestrictionEnum) Reset()         { *m = GeoTargetingRestrictionEnum{} }
func (m *GeoTargetingRestrictionEnum) String() string { return proto.CompactTextString(m) }
func (*GeoTargetingRestrictionEnum) ProtoMessage()    {}
func (*GeoTargetingRestrictionEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_geo_targeting_restriction_40a83d06ddfa34b9, []int{0}
}
func (m *GeoTargetingRestrictionEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GeoTargetingRestrictionEnum.Unmarshal(m, b)
}
func (m *GeoTargetingRestrictionEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GeoTargetingRestrictionEnum.Marshal(b, m, deterministic)
}
func (dst *GeoTargetingRestrictionEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GeoTargetingRestrictionEnum.Merge(dst, src)
}
func (m *GeoTargetingRestrictionEnum) XXX_Size() int {
	return xxx_messageInfo_GeoTargetingRestrictionEnum.Size(m)
}
func (m *GeoTargetingRestrictionEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_GeoTargetingRestrictionEnum.DiscardUnknown(m)
}

var xxx_messageInfo_GeoTargetingRestrictionEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GeoTargetingRestrictionEnum)(nil), "google.ads.googleads.v1.enums.GeoTargetingRestrictionEnum")
	proto.RegisterEnum("google.ads.googleads.v1.enums.GeoTargetingRestrictionEnum_GeoTargetingRestriction", GeoTargetingRestrictionEnum_GeoTargetingRestriction_name, GeoTargetingRestrictionEnum_GeoTargetingRestriction_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/geo_targeting_restriction.proto", fileDescriptor_geo_targeting_restriction_40a83d06ddfa34b9)
}

var fileDescriptor_geo_targeting_restriction_40a83d06ddfa34b9 = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xcf, 0x4a, 0x2b, 0x31,
	0x18, 0xc5, 0x6f, 0xe7, 0x82, 0x42, 0xba, 0xb0, 0x0c, 0x82, 0xa2, 0xed, 0xa2, 0x7d, 0x80, 0x84,
	0xc1, 0x5d, 0xc4, 0x45, 0x5a, 0xd3, 0x52, 0x94, 0xcc, 0xd8, 0x7f, 0x82, 0x0c, 0x0c, 0xb1, 0x13,
	0xc2, 0x40, 0x9b, 0x0c, 0x93, 0xb4, 0x0f, 0xe4, 0xd2, 0x47, 0xf1, 0x51, 0xdc, 0xf9, 0x06, 0x32,
	0x89, 0x53, 0x57, 0x75, 0x13, 0x0e, 0x39, 0xdf, 0xef, 0xe4, 0xcb, 0x01, 0x77, 0x52, 0x6b, 0xb9,
	0x11, 0x88, 0xe7, 0x06, 0x79, 0x59, 0xab, 0x7d, 0x84, 0x84, 0xda, 0x6d, 0x0d, 0x92, 0x42, 0x67,
	0x96, 0x57, 0x52, 0xd8, 0x42, 0xc9, 0xac, 0x12, 0xc6, 0x56, 0xc5, 0xda, 0x16, 0x5a, 0xc1, 0xb2,
	0xd2, 0x56, 0x87, 0x3d, 0xcf, 0x40, 0x9e, 0x1b, 0x78, 0xc0, 0xe1, 0x3e, 0x82, 0x0e, 0xbf, 0xea,
	0x36, 0xe9, 0x65, 0x81, 0xb8, 0x52, 0xda, 0xf2, 0x9a, 0x35, 0x1e, 0x1e, 0x94, 0xe0, 0x7a, 0x22,
	0xf4, 0xa2, 0x89, 0x9f, 0xfd, 0xa6, 0x53, 0xb5, 0xdb, 0x0e, 0x9e, 0xc0, 0xc5, 0x11, 0x3b, 0x3c,
	0x03, 0xed, 0x25, 0x9b, 0x27, 0x74, 0x34, 0x1d, 0x4f, 0xe9, 0x7d, 0xe7, 0x5f, 0xd8, 0x06, 0xa7,
	0x4b, 0xf6, 0xc0, 0xe2, 0x67, 0xd6, 0x69, 0x85, 0x97, 0xe0, 0xfc, 0x31, 0x1e, 0x91, 0xc5, 0x34,
	0x66, 0x59, 0x3c, 0xce, 0x92, 0x19, 0x9d, 0x53, 0x36, 0xa2, 0x9d, 0x60, 0xf8, 0xd5, 0x02, 0xfd,
	0xb5, 0xde, 0xc2, 0x3f, 0xb7, 0x1e, 0x76, 0x8f, 0x3c, 0x9b, 0xd4, 0x5b, 0x27, 0xad, 0x97, 0xe1,
	0x0f, 0x2e, 0xf5, 0x86, 0x2b, 0x09, 0x75, 0x25, 0x91, 0x14, 0xca, 0xfd, 0xa9, 0xe9, 0xb0, 0x2c,
	0xcc, 0x91, 0x4a, 0x6f, 0xdd, 0xf9, 0x16, 0xfc, 0x9f, 0x10, 0xf2, 0x1e, 0xf4, 0x26, 0x3e, 0x8a,
	0xe4, 0x06, 0x7a, 0x59, 0xab, 0x55, 0x04, 0xeb, 0x06, 0xcc, 0x47, 0xe3, 0xa7, 0x24, 0x37, 0xe9,
	0xc1, 0x4f, 0x57, 0x51, 0xea, 0xfc, 0xcf, 0xa0, 0xef, 0x2f, 0x31, 0x26, 0xb9, 0xc1, 0xf8, 0x30,
	0x81, 0xf1, 0x2a, 0xc2, 0xd8, 0xcd, 0xbc, 0x9e, 0xb8, 0xc5, 0x6e, 0xbe, 0x03, 0x00, 0x00, 0xff,
	0xff, 0x50, 0xdf, 0x18, 0xa1, 0xea, 0x01, 0x00, 0x00,
}
