// Code generated by protoc-gen-go.
// source: teletype.proto
// DO NOT EDIT!

/*
Package teletype is a generated protocol buffer package.

It is generated from these files:
	teletype.proto

It has these top-level messages:
	Telecast
*/
package teletype

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

type TelecastDeviceType int32

const (
	Telecast_UNKNOWN_DEVICE_TYPE TelecastDeviceType = 0
	Telecast_BGEIGIE_NANO        TelecastDeviceType = 1
	Telecast_SIMPLECAST          TelecastDeviceType = 2
	Telecast_TTAPP               TelecastDeviceType = 3
	Telecast_TTRELAY             TelecastDeviceType = 4
	Telecast_TTGATE              TelecastDeviceType = 5
)

var TelecastDeviceType_name = map[int32]string{
	0: "UNKNOWN_DEVICE_TYPE",
	1: "BGEIGIE_NANO",
	2: "SIMPLECAST",
	3: "TTAPP",
	4: "TTRELAY",
	5: "TTGATE",
}
var TelecastDeviceType_value = map[string]int32{
	"UNKNOWN_DEVICE_TYPE": 0,
	"BGEIGIE_NANO":        1,
	"SIMPLECAST":          2,
	"TTAPP":               3,
	"TTRELAY":             4,
	"TTGATE":              5,
}

func (x TelecastDeviceType) Enum() *TelecastDeviceType {
	p := new(TelecastDeviceType)
	*p = x
	return p
}
func (x TelecastDeviceType) String() string {
	return proto.EnumName(TelecastDeviceType_name, int32(x))
}
func (x *TelecastDeviceType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(TelecastDeviceType_value, data, "TelecastDeviceType")
	if err != nil {
		return err
	}
	*x = TelecastDeviceType(value)
	return nil
}
func (TelecastDeviceType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type TelecastUnit int32

const (
	Telecast_UNKNOWN_UNIT TelecastUnit = 0
	Telecast_CPM          TelecastUnit = 1
)

var TelecastUnit_name = map[int32]string{
	0: "UNKNOWN_UNIT",
	1: "CPM",
}
var TelecastUnit_value = map[string]int32{
	"UNKNOWN_UNIT": 0,
	"CPM":          1,
}

func (x TelecastUnit) Enum() *TelecastUnit {
	p := new(TelecastUnit)
	*p = x
	return p
}
func (x TelecastUnit) String() string {
	return proto.EnumName(TelecastUnit_name, int32(x))
}
func (x *TelecastUnit) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(TelecastUnit_value, data, "TelecastUnit")
	if err != nil {
		return err
	}
	*x = TelecastUnit(value)
	return nil
}
func (TelecastUnit) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 1} }

type Telecast struct {
	DeviceType       *TelecastDeviceType `protobuf:"varint,1,req,name=DeviceType,enum=teletype.TelecastDeviceType" json:"DeviceType,omitempty"`
	DeviceIDString   *string             `protobuf:"bytes,2,opt,name=DeviceIDString" json:"DeviceIDString,omitempty"`
	DeviceIDNumber   *uint32             `protobuf:"varint,3,opt,name=DeviceIDNumber" json:"DeviceIDNumber,omitempty"`
	Message          *string             `protobuf:"bytes,4,opt,name=Message" json:"Message,omitempty"`
	CapturedAt       *string             `protobuf:"bytes,5,opt,name=CapturedAt" json:"CapturedAt,omitempty"`
	Unit             *TelecastUnit       `protobuf:"varint,6,opt,name=Unit,enum=teletype.TelecastUnit" json:"Unit,omitempty"`
	Value            *uint32             `protobuf:"varint,7,opt,name=Value" json:"Value,omitempty"`
	Latitude         *float32            `protobuf:"fixed32,8,opt,name=Latitude" json:"Latitude,omitempty"`
	Longitude        *float32            `protobuf:"fixed32,9,opt,name=Longitude" json:"Longitude,omitempty"`
	Altitude         *uint32             `protobuf:"varint,10,opt,name=Altitude" json:"Altitude,omitempty"`
	BatteryVoltage   *float32            `protobuf:"fixed32,11,opt,name=BatteryVoltage" json:"BatteryVoltage,omitempty"`
	BatterySOC       *float32            `protobuf:"fixed32,12,opt,name=BatterySOC" json:"BatterySOC,omitempty"`
	WirelessSNR      *float32            `protobuf:"fixed32,13,opt,name=WirelessSNR" json:"WirelessSNR,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *Telecast) Reset()                    { *m = Telecast{} }
func (m *Telecast) String() string            { return proto.CompactTextString(m) }
func (*Telecast) ProtoMessage()               {}
func (*Telecast) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Telecast) GetDeviceType() TelecastDeviceType {
	if m != nil && m.DeviceType != nil {
		return *m.DeviceType
	}
	return Telecast_UNKNOWN_DEVICE_TYPE
}

func (m *Telecast) GetDeviceIDString() string {
	if m != nil && m.DeviceIDString != nil {
		return *m.DeviceIDString
	}
	return ""
}

func (m *Telecast) GetDeviceIDNumber() uint32 {
	if m != nil && m.DeviceIDNumber != nil {
		return *m.DeviceIDNumber
	}
	return 0
}

func (m *Telecast) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func (m *Telecast) GetCapturedAt() string {
	if m != nil && m.CapturedAt != nil {
		return *m.CapturedAt
	}
	return ""
}

func (m *Telecast) GetUnit() TelecastUnit {
	if m != nil && m.Unit != nil {
		return *m.Unit
	}
	return Telecast_UNKNOWN_UNIT
}

func (m *Telecast) GetValue() uint32 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

func (m *Telecast) GetLatitude() float32 {
	if m != nil && m.Latitude != nil {
		return *m.Latitude
	}
	return 0
}

func (m *Telecast) GetLongitude() float32 {
	if m != nil && m.Longitude != nil {
		return *m.Longitude
	}
	return 0
}

func (m *Telecast) GetAltitude() uint32 {
	if m != nil && m.Altitude != nil {
		return *m.Altitude
	}
	return 0
}

func (m *Telecast) GetBatteryVoltage() float32 {
	if m != nil && m.BatteryVoltage != nil {
		return *m.BatteryVoltage
	}
	return 0
}

func (m *Telecast) GetBatterySOC() float32 {
	if m != nil && m.BatterySOC != nil {
		return *m.BatterySOC
	}
	return 0
}

func (m *Telecast) GetWirelessSNR() float32 {
	if m != nil && m.WirelessSNR != nil {
		return *m.WirelessSNR
	}
	return 0
}

func init() {
	proto.RegisterType((*Telecast)(nil), "teletype.Telecast")
	proto.RegisterEnum("teletype.TelecastDeviceType", TelecastDeviceType_name, TelecastDeviceType_value)
	proto.RegisterEnum("teletype.TelecastUnit", TelecastUnit_name, TelecastUnit_value)
}

var fileDescriptor0 = []byte{
	// 361 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x90, 0xdd, 0xaa, 0xda, 0x40,
	0x10, 0xc7, 0x9b, 0xaf, 0x93, 0x64, 0x3c, 0xa6, 0xdb, 0x3d, 0xd0, 0xb3, 0x37, 0x05, 0x2b, 0x14,
	0xbc, 0x12, 0xea, 0x1b, 0xc4, 0xb8, 0x48, 0x68, 0x5c, 0x83, 0x59, 0x15, 0xaf, 0x24, 0xd5, 0x45,
	0x42, 0x53, 0x23, 0xc9, 0xa6, 0xe0, 0x5b, 0xf6, 0x91, 0xba, 0xae, 0x1f, 0x78, 0xd1, 0xcb, 0xf9,
	0xcd, 0x7f, 0x66, 0x7e, 0xbb, 0x10, 0x48, 0x51, 0x0a, 0x79, 0x3e, 0x89, 0xe1, 0xa9, 0xae, 0x64,
	0x85, 0xbd, 0x7b, 0xdd, 0xff, 0x6b, 0x81, 0xc7, 0x55, 0xb1, 0xcb, 0x1b, 0x89, 0xbf, 0x03, 0x4c,
	0xc4, 0x9f, 0x62, 0x27, 0xb8, 0x6a, 0x11, 0xa3, 0x67, 0x0e, 0x82, 0xd1, 0x97, 0xe1, 0x63, 0xf6,
	0x9e, 0x1b, 0xee, 0x1f, 0x21, 0xfc, 0x19, 0x82, 0xeb, 0x48, 0x3c, 0xc9, 0x64, 0x5d, 0x1c, 0x0f,
	0xc4, 0xec, 0x19, 0x03, 0xff, 0x99, 0xb3, 0xf6, 0xf7, 0x4f, 0x51, 0x13, 0x4b, 0xf1, 0x2e, 0xfe,
	0x08, 0xee, 0x4c, 0x34, 0x4d, 0x7e, 0x10, 0xc4, 0xd6, 0x41, 0x0c, 0x10, 0xe5, 0x27, 0xd9, 0xd6,
	0x62, 0x1f, 0x4a, 0xe2, 0x68, 0xf6, 0x0d, 0xec, 0xe5, 0xb1, 0x90, 0xe4, 0x45, 0x55, 0xc1, 0xe8,
	0xfd, 0x3f, 0x06, 0xad, 0x6a, 0xe3, 0x2e, 0x38, 0xab, 0xbc, 0x6c, 0x05, 0x71, 0xf5, 0x6a, 0x04,
	0x5e, 0x92, 0xcb, 0x42, 0xb6, 0x7b, 0x41, 0x3c, 0x45, 0x4c, 0xfc, 0x09, 0xfc, 0xa4, 0x3a, 0x1e,
	0xae, 0xc8, 0xd7, 0x48, 0x85, 0xc2, 0xf2, 0x16, 0x02, 0x3d, 0xa6, 0x4c, 0xc7, 0xb9, 0x94, 0xa2,
	0x3e, 0xaf, 0xaa, 0x52, 0x5e, 0xc4, 0x3a, 0x3a, 0xa9, 0xc4, 0x6e, 0x3c, 0x9b, 0x47, 0xe4, 0x55,
	0xb3, 0x37, 0xe8, 0xac, 0x8b, 0x5a, 0x49, 0x34, 0x4d, 0xc6, 0x16, 0xa4, 0x7b, 0x81, 0xfd, 0x5f,
	0x00, 0x4f, 0x1f, 0xf2, 0x0e, 0x6f, 0x4b, 0xf6, 0x83, 0xcd, 0xd7, 0x6c, 0x3b, 0xa1, 0xab, 0x38,
	0xa2, 0x5b, 0xbe, 0x49, 0x29, 0xfa, 0xa0, 0x2e, 0xbf, 0x8e, 0xa7, 0x34, 0x9e, 0xc6, 0x74, 0xcb,
	0x42, 0x36, 0x47, 0x06, 0x0e, 0x00, 0xb2, 0x78, 0x96, 0x26, 0x34, 0x0a, 0x33, 0x8e, 0x4c, 0xec,
	0x83, 0xc3, 0x79, 0x98, 0xa6, 0xc8, 0xc2, 0x1d, 0x70, 0x39, 0x5f, 0xd0, 0x24, 0xdc, 0x20, 0x5b,
	0x89, 0xbc, 0x70, 0x3e, 0x0d, 0x39, 0x45, 0x4e, 0xff, 0x2b, 0xd8, 0xfa, 0xed, 0x6a, 0xdb, 0xfd,
	0xcc, 0x92, 0xc5, 0x5c, 0xed, 0x77, 0xc1, 0x8a, 0xd2, 0x19, 0x32, 0xfe, 0x05, 0x00, 0x00, 0xff,
	0xff, 0x55, 0x4f, 0xcb, 0xdb, 0xed, 0x01, 0x00, 0x00,
}
