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
	Telecast_TTSERVE             TelecastDeviceType = 6
)

var TelecastDeviceType_name = map[int32]string{
	0: "UNKNOWN_DEVICE_TYPE",
	1: "BGEIGIE_NANO",
	2: "SIMPLECAST",
	3: "TTAPP",
	4: "TTRELAY",
	5: "TTGATE",
	6: "TTSERVE",
}
var TelecastDeviceType_value = map[string]int32{
	"UNKNOWN_DEVICE_TYPE": 0,
	"BGEIGIE_NANO":        1,
	"SIMPLECAST":          2,
	"TTAPP":               3,
	"TTRELAY":             4,
	"TTGATE":              5,
	"TTSERVE":             6,
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
	DeviceType            *TelecastDeviceType `protobuf:"varint,1,req,name=DeviceType,enum=teletype.TelecastDeviceType" json:"DeviceType,omitempty"`
	DeviceIDString        *string             `protobuf:"bytes,2,opt,name=DeviceIDString" json:"DeviceIDString,omitempty"`
	DeviceIDNumber        *uint32             `protobuf:"varint,3,opt,name=DeviceIDNumber" json:"DeviceIDNumber,omitempty"`
	Message               *string             `protobuf:"bytes,4,opt,name=Message" json:"Message,omitempty"`
	CapturedAt            *string             `protobuf:"bytes,5,opt,name=CapturedAt" json:"CapturedAt,omitempty"`
	Unit                  *TelecastUnit       `protobuf:"varint,6,opt,name=Unit,enum=teletype.TelecastUnit" json:"Unit,omitempty"`
	Value                 *uint32             `protobuf:"varint,7,opt,name=Value" json:"Value,omitempty"`
	Latitude              *float32            `protobuf:"fixed32,8,opt,name=Latitude" json:"Latitude,omitempty"`
	Longitude             *float32            `protobuf:"fixed32,9,opt,name=Longitude" json:"Longitude,omitempty"`
	Altitude              *uint32             `protobuf:"varint,10,opt,name=Altitude" json:"Altitude,omitempty"`
	BatteryVoltage        *float32            `protobuf:"fixed32,11,opt,name=BatteryVoltage" json:"BatteryVoltage,omitempty"`
	BatterySOC            *float32            `protobuf:"fixed32,12,opt,name=BatterySOC" json:"BatterySOC,omitempty"`
	WirelessSNR           *float32            `protobuf:"fixed32,13,opt,name=WirelessSNR" json:"WirelessSNR,omitempty"`
	EnvTemperature        *float32            `protobuf:"fixed32,14,opt,name=envTemperature" json:"envTemperature,omitempty"`
	EnvHumidity           *float32            `protobuf:"fixed32,15,opt,name=envHumidity" json:"envHumidity,omitempty"`
	RelayDevice1          *uint32             `protobuf:"varint,16,opt,name=RelayDevice1" json:"RelayDevice1,omitempty"`
	RelayDevice2          *uint32             `protobuf:"varint,17,opt,name=RelayDevice2" json:"RelayDevice2,omitempty"`
	RelayDevice3          *uint32             `protobuf:"varint,18,opt,name=RelayDevice3" json:"RelayDevice3,omitempty"`
	RelayDevice4          *uint32             `protobuf:"varint,19,opt,name=RelayDevice4" json:"RelayDevice4,omitempty"`
	RelayDevice5          *uint32             `protobuf:"varint,20,opt,name=RelayDevice5" json:"RelayDevice5,omitempty"`
	PmsTsi_01_0           *uint32             `protobuf:"varint,21,opt,name=pms_tsi_01_0" json:"pms_tsi_01_0,omitempty"`
	PmsTsi_02_5           *uint32             `protobuf:"varint,22,opt,name=pms_tsi_02_5" json:"pms_tsi_02_5,omitempty"`
	PmsTsi_10_0           *uint32             `protobuf:"varint,23,opt,name=pms_tsi_10_0" json:"pms_tsi_10_0,omitempty"`
	PmsStd_01_0           *uint32             `protobuf:"varint,24,opt,name=pms_std_01_0" json:"pms_std_01_0,omitempty"`
	PmsStd_02_5           *uint32             `protobuf:"varint,25,opt,name=pms_std_02_5" json:"pms_std_02_5,omitempty"`
	PmsStd_10_0           *uint32             `protobuf:"varint,26,opt,name=pms_std_10_0" json:"pms_std_10_0,omitempty"`
	PmsCount_00_3         *uint32             `protobuf:"varint,27,opt,name=pms_count_00_3" json:"pms_count_00_3,omitempty"`
	PmsCount_00_5         *uint32             `protobuf:"varint,28,opt,name=pms_count_00_5" json:"pms_count_00_5,omitempty"`
	PmsCount_01_0         *uint32             `protobuf:"varint,29,opt,name=pms_count_01_0" json:"pms_count_01_0,omitempty"`
	PmsCount_02_5         *uint32             `protobuf:"varint,30,opt,name=pms_count_02_5" json:"pms_count_02_5,omitempty"`
	PmsCount_05_0         *uint32             `protobuf:"varint,31,opt,name=pms_count_05_0" json:"pms_count_05_0,omitempty"`
	PmsCount_10_0         *uint32             `protobuf:"varint,32,opt,name=pms_count_10_0" json:"pms_count_10_0,omitempty"`
	StatsUptimeMinutes    *uint32             `protobuf:"varint,33,opt,name=stats_uptime_minutes" json:"stats_uptime_minutes,omitempty"`
	StatsAppVersion       *string             `protobuf:"bytes,34,opt,name=stats_app_version" json:"stats_app_version,omitempty"`
	StatsDeviceParams     *string             `protobuf:"bytes,35,opt,name=stats_device_params" json:"stats_device_params,omitempty"`
	StatsTransmittedBytes *uint32             `protobuf:"varint,36,opt,name=stats_transmitted_bytes" json:"stats_transmitted_bytes,omitempty"`
	StatsReceivedBytes    *uint32             `protobuf:"varint,37,opt,name=stats_received_bytes" json:"stats_received_bytes,omitempty"`
	StatsOneshots         *uint32             `protobuf:"varint,38,opt,name=stats_oneshots" json:"stats_oneshots,omitempty"`
	StatsCommsResets      *uint32             `protobuf:"varint,39,opt,name=stats_comms_resets" json:"stats_comms_resets,omitempty"`
	Opc_01_0              *float32            `protobuf:"fixed32,40,opt,name=opc_01_0" json:"opc_01_0,omitempty"`
	Opc_02_5              *float32            `protobuf:"fixed32,41,opt,name=opc_02_5" json:"opc_02_5,omitempty"`
	Opc_10_0              *float32            `protobuf:"fixed32,42,opt,name=opc_10_0" json:"opc_10_0,omitempty"`
	XXX_unrecognized      []byte              `json:"-"`
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

func (m *Telecast) GetEnvTemperature() float32 {
	if m != nil && m.EnvTemperature != nil {
		return *m.EnvTemperature
	}
	return 0
}

func (m *Telecast) GetEnvHumidity() float32 {
	if m != nil && m.EnvHumidity != nil {
		return *m.EnvHumidity
	}
	return 0
}

func (m *Telecast) GetRelayDevice1() uint32 {
	if m != nil && m.RelayDevice1 != nil {
		return *m.RelayDevice1
	}
	return 0
}

func (m *Telecast) GetRelayDevice2() uint32 {
	if m != nil && m.RelayDevice2 != nil {
		return *m.RelayDevice2
	}
	return 0
}

func (m *Telecast) GetRelayDevice3() uint32 {
	if m != nil && m.RelayDevice3 != nil {
		return *m.RelayDevice3
	}
	return 0
}

func (m *Telecast) GetRelayDevice4() uint32 {
	if m != nil && m.RelayDevice4 != nil {
		return *m.RelayDevice4
	}
	return 0
}

func (m *Telecast) GetRelayDevice5() uint32 {
	if m != nil && m.RelayDevice5 != nil {
		return *m.RelayDevice5
	}
	return 0
}

func (m *Telecast) GetPmsTsi_01_0() uint32 {
	if m != nil && m.PmsTsi_01_0 != nil {
		return *m.PmsTsi_01_0
	}
	return 0
}

func (m *Telecast) GetPmsTsi_02_5() uint32 {
	if m != nil && m.PmsTsi_02_5 != nil {
		return *m.PmsTsi_02_5
	}
	return 0
}

func (m *Telecast) GetPmsTsi_10_0() uint32 {
	if m != nil && m.PmsTsi_10_0 != nil {
		return *m.PmsTsi_10_0
	}
	return 0
}

func (m *Telecast) GetPmsStd_01_0() uint32 {
	if m != nil && m.PmsStd_01_0 != nil {
		return *m.PmsStd_01_0
	}
	return 0
}

func (m *Telecast) GetPmsStd_02_5() uint32 {
	if m != nil && m.PmsStd_02_5 != nil {
		return *m.PmsStd_02_5
	}
	return 0
}

func (m *Telecast) GetPmsStd_10_0() uint32 {
	if m != nil && m.PmsStd_10_0 != nil {
		return *m.PmsStd_10_0
	}
	return 0
}

func (m *Telecast) GetPmsCount_00_3() uint32 {
	if m != nil && m.PmsCount_00_3 != nil {
		return *m.PmsCount_00_3
	}
	return 0
}

func (m *Telecast) GetPmsCount_00_5() uint32 {
	if m != nil && m.PmsCount_00_5 != nil {
		return *m.PmsCount_00_5
	}
	return 0
}

func (m *Telecast) GetPmsCount_01_0() uint32 {
	if m != nil && m.PmsCount_01_0 != nil {
		return *m.PmsCount_01_0
	}
	return 0
}

func (m *Telecast) GetPmsCount_02_5() uint32 {
	if m != nil && m.PmsCount_02_5 != nil {
		return *m.PmsCount_02_5
	}
	return 0
}

func (m *Telecast) GetPmsCount_05_0() uint32 {
	if m != nil && m.PmsCount_05_0 != nil {
		return *m.PmsCount_05_0
	}
	return 0
}

func (m *Telecast) GetPmsCount_10_0() uint32 {
	if m != nil && m.PmsCount_10_0 != nil {
		return *m.PmsCount_10_0
	}
	return 0
}

func (m *Telecast) GetStatsUptimeMinutes() uint32 {
	if m != nil && m.StatsUptimeMinutes != nil {
		return *m.StatsUptimeMinutes
	}
	return 0
}

func (m *Telecast) GetStatsAppVersion() string {
	if m != nil && m.StatsAppVersion != nil {
		return *m.StatsAppVersion
	}
	return ""
}

func (m *Telecast) GetStatsDeviceParams() string {
	if m != nil && m.StatsDeviceParams != nil {
		return *m.StatsDeviceParams
	}
	return ""
}

func (m *Telecast) GetStatsTransmittedBytes() uint32 {
	if m != nil && m.StatsTransmittedBytes != nil {
		return *m.StatsTransmittedBytes
	}
	return 0
}

func (m *Telecast) GetStatsReceivedBytes() uint32 {
	if m != nil && m.StatsReceivedBytes != nil {
		return *m.StatsReceivedBytes
	}
	return 0
}

func (m *Telecast) GetStatsOneshots() uint32 {
	if m != nil && m.StatsOneshots != nil {
		return *m.StatsOneshots
	}
	return 0
}

func (m *Telecast) GetStatsCommsResets() uint32 {
	if m != nil && m.StatsCommsResets != nil {
		return *m.StatsCommsResets
	}
	return 0
}

func (m *Telecast) GetOpc_01_0() float32 {
	if m != nil && m.Opc_01_0 != nil {
		return *m.Opc_01_0
	}
	return 0
}

func (m *Telecast) GetOpc_02_5() float32 {
	if m != nil && m.Opc_02_5 != nil {
		return *m.Opc_02_5
	}
	return 0
}

func (m *Telecast) GetOpc_10_0() float32 {
	if m != nil && m.Opc_10_0 != nil {
		return *m.Opc_10_0
	}
	return 0
}

func init() {
	proto.RegisterType((*Telecast)(nil), "teletype.Telecast")
	proto.RegisterEnum("teletype.TelecastDeviceType", TelecastDeviceType_name, TelecastDeviceType_value)
	proto.RegisterEnum("teletype.TelecastUnit", TelecastUnit_name, TelecastUnit_value)
}

var fileDescriptor0 = []byte{
	// 634 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x93, 0x6f, 0x53, 0xda, 0x40,
	0x10, 0xc6, 0x0b, 0x82, 0xc0, 0xaa, 0x18, 0x83, 0x95, 0xf5, 0x5f, 0x55, 0x5a, 0x5b, 0xdb, 0x17,
	0x0e, 0x60, 0xf9, 0x00, 0x11, 0x33, 0x96, 0x29, 0x46, 0x06, 0x22, 0x8e, 0xaf, 0x6e, 0x22, 0xdc,
	0xd8, 0x9b, 0x21, 0x7f, 0x26, 0x77, 0x61, 0xca, 0x67, 0xe8, 0x97, 0xee, 0xe5, 0x62, 0x1c, 0xd4,
	0xbc, 0xbc, 0xdf, 0xf3, 0xec, 0xee, 0x73, 0x37, 0x7b, 0x50, 0x15, 0x74, 0x46, 0xc5, 0x22, 0xa0,
	0xe7, 0x41, 0xe8, 0x0b, 0x5f, 0x2f, 0xa7, 0xe7, 0xc6, 0xbf, 0x0a, 0x94, 0x6d, 0x79, 0x98, 0x38,
	0x5c, 0xe8, 0x2d, 0x80, 0x2b, 0x3a, 0x67, 0x13, 0x6a, 0x4b, 0x09, 0x73, 0xc7, 0xf9, 0xb3, 0x6a,
	0xfb, 0xf0, 0xfc, 0xa5, 0x36, 0xf5, 0x9d, 0x4f, 0x5f, 0x4c, 0xfa, 0x0e, 0x54, 0x93, 0x92, 0xde,
	0xd5, 0x48, 0x84, 0xcc, 0x7b, 0xc2, 0xfc, 0x71, 0xee, 0xac, 0xb2, 0xcc, 0xad, 0xc8, 0x7d, 0xa4,
	0x21, 0xae, 0x48, 0xbe, 0xa1, 0x6f, 0x42, 0xe9, 0x86, 0x72, 0xee, 0x3c, 0x51, 0x2c, 0x28, 0xa3,
	0x0e, 0xd0, 0x75, 0x02, 0x11, 0x85, 0x74, 0x6a, 0x08, 0x2c, 0x2a, 0x76, 0x0a, 0x85, 0x3b, 0x8f,
	0x09, 0x5c, 0x95, 0xa7, 0x6a, 0xbb, 0x9e, 0x91, 0x20, 0x92, 0xb2, 0xbe, 0x01, 0xc5, 0xb1, 0x33,
	0x8b, 0x28, 0x96, 0x54, 0x6b, 0x0d, 0xca, 0x7d, 0x47, 0x30, 0x11, 0x4d, 0x29, 0x96, 0x25, 0xc9,
	0xeb, 0x5b, 0x50, 0xe9, 0xfb, 0xde, 0x53, 0x82, 0x2a, 0x0a, 0x49, 0x93, 0x31, 0x7b, 0x36, 0x81,
	0x2a, 0x93, 0x49, 0x2f, 0x1d, 0x21, 0x68, 0xb8, 0x18, 0xfb, 0x33, 0x11, 0x07, 0x5b, 0x53, 0x4e,
	0x19, 0xec, 0x99, 0x8f, 0x6e, 0xbb, 0xb8, 0xae, 0x58, 0x0d, 0xd6, 0xee, 0x59, 0x28, 0x43, 0x70,
	0x3e, 0xb2, 0x86, 0xb8, 0xa1, 0xa0, 0x6c, 0x40, 0xbd, 0xb9, 0x4d, 0xdd, 0x80, 0x86, 0x4e, 0x7c,
	0x11, 0xac, 0xa6, 0x66, 0xc9, 0x7f, 0x45, 0x2e, 0x9b, 0x32, 0xb1, 0xc0, 0x4d, 0x05, 0xb7, 0x61,
	0x7d, 0x48, 0x67, 0xce, 0x22, 0x79, 0x9c, 0x16, 0x6a, 0x2a, 0xc3, 0x6b, 0xda, 0xc6, 0xad, 0x0c,
	0x7a, 0x81, 0x7a, 0x06, 0xfd, 0x89, 0xb5, 0x0c, 0xda, 0xc1, 0xed, 0x94, 0x06, 0x2e, 0x27, 0x82,
	0x33, 0xd2, 0x6c, 0x91, 0x26, 0x7e, 0x7c, 0x47, 0xdb, 0xa4, 0x83, 0x3b, 0x6f, 0x69, 0xab, 0x29,
	0xbd, 0xf5, 0x65, 0xca, 0xc5, 0x34, 0xe9, 0x80, 0xef, 0x68, 0xdc, 0x61, 0xf7, 0x2d, 0x55, 0x1d,
	0xf6, 0xd2, 0xf7, 0x8d, 0xe9, 0xc4, 0x8f, 0x3c, 0x41, 0x9a, 0x4d, 0x72, 0x81, 0xfb, 0x99, 0xbc,
	0x83, 0x07, 0x19, 0x3c, 0x9e, 0x79, 0x98, 0xc1, 0xe3, 0xa9, 0x9f, 0x32, 0x78, 0x47, 0xfa, 0x8f,
	0xde, 0x73, 0x95, 0xe7, 0x58, 0xf1, 0x03, 0xd8, 0xe6, 0xc2, 0x11, 0x9c, 0x44, 0x81, 0x60, 0x2e,
	0x25, 0x2e, 0xf3, 0x22, 0x41, 0x39, 0x9e, 0x28, 0x75, 0x17, 0xb6, 0x12, 0xd5, 0x09, 0x02, 0x32,
	0xa7, 0x21, 0x67, 0xbe, 0x87, 0x0d, 0xb5, 0x95, 0xfb, 0x50, 0x4b, 0xa4, 0x64, 0xfd, 0x49, 0xe0,
	0x84, 0x8e, 0xcb, 0xf1, 0xb3, 0x12, 0x8f, 0xa0, 0x9e, 0x88, 0x22, 0x74, 0x3c, 0xee, 0x32, 0xb9,
	0x38, 0x53, 0xf2, 0xb8, 0x88, 0x1b, 0x7f, 0x79, 0x3d, 0x36, 0xa4, 0x13, 0xca, 0xe6, 0x2f, 0xea,
	0x69, 0x1a, 0x36, 0x51, 0x7d, 0x8f, 0xf2, 0x3f, 0xbe, 0xe0, 0xf8, 0x55, 0xf1, 0x3d, 0xd0, 0x13,
	0x3e, 0xf1, 0x5d, 0x37, 0xae, 0xe5, 0x54, 0x6a, 0xdf, 0xd2, 0x7d, 0xf7, 0x83, 0x49, 0xf2, 0x44,
	0x67, 0xe9, 0x72, 0x2b, 0x12, 0x3f, 0xce, 0xf7, 0x65, 0xa2, 0xae, 0xff, 0x23, 0x26, 0x8d, 0xbf,
	0x00, 0x4b, 0xdf, 0xb7, 0x0e, 0xb5, 0x3b, 0xeb, 0xb7, 0x75, 0x7b, 0x6f, 0x91, 0x2b, 0x73, 0xdc,
	0xeb, 0x9a, 0xc4, 0x7e, 0x18, 0x98, 0xda, 0x07, 0x59, 0xb8, 0x7e, 0x79, 0x6d, 0xf6, 0xae, 0x7b,
	0x26, 0xb1, 0x0c, 0xeb, 0x56, 0xcb, 0xe9, 0x55, 0x80, 0x51, 0xef, 0x66, 0xd0, 0x37, 0xbb, 0xc6,
	0xc8, 0xd6, 0xf2, 0x7a, 0x05, 0x8a, 0xb6, 0x6d, 0x0c, 0x06, 0xda, 0x8a, 0xbe, 0x06, 0x25, 0xdb,
	0x1e, 0x9a, 0x7d, 0xe3, 0x41, 0x2b, 0xc8, 0x6f, 0xb3, 0x6a, 0xdb, 0xd7, 0x86, 0x6d, 0x6a, 0xc5,
	0x44, 0x18, 0x99, 0xc3, 0xb1, 0xa9, 0xad, 0x36, 0x4e, 0xa0, 0xa0, 0xbe, 0xad, 0x6c, 0x9d, 0xce,
	0xbc, 0xb3, 0x7a, 0xb6, 0x1c, 0x56, 0x82, 0x95, 0xee, 0xe0, 0x46, 0xcb, 0xfd, 0x0f, 0x00, 0x00,
	0xff, 0xff, 0x93, 0x74, 0x50, 0x49, 0xa8, 0x04, 0x00, 0x00,
}
