// Code generated by protoc-gen-go.
// source: tt.proto
// DO NOT EDIT!

/*
Package ttproto is a generated protocol buffer package.

It is generated from these files:
	tt.proto

It has these top-level messages:
	Telecast
*/
package ttproto

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
	Telecast_SOLARCAST           TelecastDeviceType = 2
	Telecast_TTAPP               TelecastDeviceType = 3
	Telecast_TTNODE              TelecastDeviceType = 4
	Telecast_TTGATE              TelecastDeviceType = 5
	Telecast_TTSERVE             TelecastDeviceType = 6
)

var TelecastDeviceType_name = map[int32]string{
	0: "UNKNOWN_DEVICE_TYPE",
	1: "BGEIGIE_NANO",
	2: "SOLARCAST",
	3: "TTAPP",
	4: "TTNODE",
	5: "TTGATE",
	6: "TTSERVE",
}
var TelecastDeviceType_value = map[string]int32{
	"UNKNOWN_DEVICE_TYPE": 0,
	"BGEIGIE_NANO":        1,
	"SOLARCAST":           2,
	"TTAPP":               3,
	"TTNODE":              4,
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

type TelecastReplyType int32

const (
	Telecast_NO_REPLY       TelecastReplyType = 0
	Telecast_REPLY_EXPECTED TelecastReplyType = 1
)

var TelecastReplyType_name = map[int32]string{
	0: "NO_REPLY",
	1: "REPLY_EXPECTED",
}
var TelecastReplyType_value = map[string]int32{
	"NO_REPLY":       0,
	"REPLY_EXPECTED": 1,
}

func (x TelecastReplyType) Enum() *TelecastReplyType {
	p := new(TelecastReplyType)
	*p = x
	return p
}
func (x TelecastReplyType) String() string {
	return proto.EnumName(TelecastReplyType_name, int32(x))
}
func (x *TelecastReplyType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(TelecastReplyType_value, data, "TelecastReplyType")
	if err != nil {
		return err
	}
	*x = TelecastReplyType(value)
	return nil
}
func (TelecastReplyType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 1} }

type Telecast struct {
	DeviceType                      *TelecastDeviceType `protobuf:"varint,1,req,name=device_type,enum=ttproto.TelecastDeviceType" json:"device_type,omitempty"`
	DEPRECATED2017FEBDeviceIDString *string             `protobuf:"bytes,2,opt,name=DEPRECATED2017FEBDeviceIDString" json:"DEPRECATED2017FEBDeviceIDString,omitempty"`
	DeviceId                        *uint32             `protobuf:"varint,3,opt,name=device_id" json:"device_id,omitempty"`
	Message                         *string             `protobuf:"bytes,4,opt,name=message" json:"message,omitempty"`
	CapturedAt                      *string             `protobuf:"bytes,5,opt,name=captured_at" json:"captured_at,omitempty"`
	ReplyType                       *TelecastReplyType  `protobuf:"varint,6,opt,name=reply_type,enum=ttproto.TelecastReplyType" json:"reply_type,omitempty"`
	DEPRECATED2017FEBValue          *uint32             `protobuf:"varint,7,opt,name=DEPRECATED2017FEBValue" json:"DEPRECATED2017FEBValue,omitempty"`
	Latitude                        *float32            `protobuf:"fixed32,8,opt,name=latitude" json:"latitude,omitempty"`
	Longitude                       *float32            `protobuf:"fixed32,9,opt,name=longitude" json:"longitude,omitempty"`
	Altitude                        *int32              `protobuf:"varint,10,opt,name=altitude" json:"altitude,omitempty"`
	BatVoltage                      *float32            `protobuf:"fixed32,11,opt,name=bat_voltage" json:"bat_voltage,omitempty"`
	BatSoc                          *float32            `protobuf:"fixed32,12,opt,name=bat_soc" json:"bat_soc,omitempty"`
	WirelessSnr                     *float32            `protobuf:"fixed32,13,opt,name=wireless_snr" json:"wireless_snr,omitempty"`
	EnvTemp                         *float32            `protobuf:"fixed32,14,opt,name=env_temp" json:"env_temp,omitempty"`
	EnvHumid                        *float32            `protobuf:"fixed32,15,opt,name=env_humid" json:"env_humid,omitempty"`
	RelayDevice1                    *uint32             `protobuf:"varint,16,opt,name=relay_device1" json:"relay_device1,omitempty"`
	RelayDevice2                    *uint32             `protobuf:"varint,17,opt,name=relay_device2" json:"relay_device2,omitempty"`
	RelayDevice3                    *uint32             `protobuf:"varint,18,opt,name=relay_device3" json:"relay_device3,omitempty"`
	RelayDevice4                    *uint32             `protobuf:"varint,19,opt,name=relay_device4" json:"relay_device4,omitempty"`
	RelayDevice5                    *uint32             `protobuf:"varint,20,opt,name=relay_device5" json:"relay_device5,omitempty"`
	Lnd_7318U                       *uint32             `protobuf:"varint,21,opt,name=lnd_7318u" json:"lnd_7318u,omitempty"`
	Lnd_7128Ec                      *uint32             `protobuf:"varint,22,opt,name=lnd_7128ec" json:"lnd_7128ec,omitempty"`
	StatsUptimeMinutes              *uint32             `protobuf:"varint,23,opt,name=stats_uptime_minutes" json:"stats_uptime_minutes,omitempty"`
	StatsAppVersion                 *string             `protobuf:"bytes,24,opt,name=stats_app_version" json:"stats_app_version,omitempty"`
	StatsDeviceParams               *string             `protobuf:"bytes,25,opt,name=stats_device_params" json:"stats_device_params,omitempty"`
	StatsTransmittedBytes           *uint32             `protobuf:"varint,26,opt,name=stats_transmitted_bytes" json:"stats_transmitted_bytes,omitempty"`
	StatsReceivedBytes              *uint32             `protobuf:"varint,27,opt,name=stats_received_bytes" json:"stats_received_bytes,omitempty"`
	StatsOneshots                   *uint32             `protobuf:"varint,28,opt,name=stats_oneshots" json:"stats_oneshots,omitempty"`
	StatsCommsResets                *uint32             `protobuf:"varint,29,opt,name=stats_comms_resets" json:"stats_comms_resets,omitempty"`
	PmsPm01_0                       *uint32             `protobuf:"varint,30,opt,name=pms_pm01_0" json:"pms_pm01_0,omitempty"`
	PmsPm02_5                       *uint32             `protobuf:"varint,31,opt,name=pms_pm02_5" json:"pms_pm02_5,omitempty"`
	PmsPm10_0                       *uint32             `protobuf:"varint,32,opt,name=pms_pm10_0" json:"pms_pm10_0,omitempty"`
	PmsC00_30                       *uint32             `protobuf:"varint,33,opt,name=pms_c00_30" json:"pms_c00_30,omitempty"`
	PmsC00_50                       *uint32             `protobuf:"varint,34,opt,name=pms_c00_50" json:"pms_c00_50,omitempty"`
	PmsC01_00                       *uint32             `protobuf:"varint,35,opt,name=pms_c01_00" json:"pms_c01_00,omitempty"`
	PmsC02_50                       *uint32             `protobuf:"varint,36,opt,name=pms_c02_50" json:"pms_c02_50,omitempty"`
	PmsC05_00                       *uint32             `protobuf:"varint,37,opt,name=pms_c05_00" json:"pms_c05_00,omitempty"`
	PmsC10_00                       *uint32             `protobuf:"varint,38,opt,name=pms_c10_00" json:"pms_c10_00,omitempty"`
	PmsCsecs                        *uint32             `protobuf:"varint,39,opt,name=pms_csecs" json:"pms_csecs,omitempty"`
	OpcPm01_0                       *float32            `protobuf:"fixed32,40,opt,name=opc_pm01_0" json:"opc_pm01_0,omitempty"`
	OpcPm02_5                       *float32            `protobuf:"fixed32,41,opt,name=opc_pm02_5" json:"opc_pm02_5,omitempty"`
	OpcPm10_0                       *float32            `protobuf:"fixed32,42,opt,name=opc_pm10_0" json:"opc_pm10_0,omitempty"`
	OpcC00_38                       *uint32             `protobuf:"varint,43,opt,name=opc_c00_38" json:"opc_c00_38,omitempty"`
	OpcC00_54                       *uint32             `protobuf:"varint,44,opt,name=opc_c00_54" json:"opc_c00_54,omitempty"`
	OpcC01_00                       *uint32             `protobuf:"varint,45,opt,name=opc_c01_00" json:"opc_c01_00,omitempty"`
	OpcC02_10                       *uint32             `protobuf:"varint,46,opt,name=opc_c02_10" json:"opc_c02_10,omitempty"`
	OpcC05_00                       *uint32             `protobuf:"varint,47,opt,name=opc_c05_00" json:"opc_c05_00,omitempty"`
	OpcC10_00                       *uint32             `protobuf:"varint,48,opt,name=opc_c10_00" json:"opc_c10_00,omitempty"`
	OpcCsecs                        *uint32             `protobuf:"varint,49,opt,name=opc_csecs" json:"opc_csecs,omitempty"`
	EnvPressure                     *float32            `protobuf:"fixed32,50,opt,name=env_pressure" json:"env_pressure,omitempty"`
	StatsCommsPowerFails            *uint32             `protobuf:"varint,51,opt,name=stats_comms_power_fails" json:"stats_comms_power_fails,omitempty"`
	BatCurrent                      *float32            `protobuf:"fixed32,52,opt,name=bat_current" json:"bat_current,omitempty"`
	StatsIccid                      *string             `protobuf:"bytes,53,opt,name=stats_iccid" json:"stats_iccid,omitempty"`
	StatsMotiondrops                *uint32             `protobuf:"varint,54,opt,name=stats_motiondrops" json:"stats_motiondrops,omitempty"`
	StatsDfu                        *string             `protobuf:"bytes,55,opt,name=stats_dfu" json:"stats_dfu,omitempty"`
	CapturedAtDate                  *uint32             `protobuf:"varint,56,opt,name=captured_at_date" json:"captured_at_date,omitempty"`
	CapturedAtTime                  *uint32             `protobuf:"varint,57,opt,name=captured_at_time" json:"captured_at_time,omitempty"`
	CapturedAtOffset                *uint32             `protobuf:"varint,58,opt,name=captured_at_offset" json:"captured_at_offset,omitempty"`
	StatsOneshotSeconds             *uint32             `protobuf:"varint,59,opt,name=stats_oneshot_seconds" json:"stats_oneshot_seconds,omitempty"`
	Stamp                           *uint32             `protobuf:"varint,60,opt,name=stamp" json:"stamp,omitempty"`
	StampVersion                    *uint32             `protobuf:"varint,61,opt,name=stamp_version" json:"stamp_version,omitempty"`
	StatsCpsi                       *string             `protobuf:"bytes,62,opt,name=stats_cpsi" json:"stats_cpsi,omitempty"`
	StatsUptimeDays                 *uint32             `protobuf:"varint,63,opt,name=stats_uptime_days" json:"stats_uptime_days,omitempty"`
	StatsDeviceLabel                *string             `protobuf:"bytes,64,opt,name=stats_device_label" json:"stats_device_label,omitempty"`
	StatsGpsParams                  *string             `protobuf:"bytes,65,opt,name=stats_gps_params" json:"stats_gps_params,omitempty"`
	StatsServiceParams              *string             `protobuf:"bytes,66,opt,name=stats_service_params" json:"stats_service_params,omitempty"`
	StatsTtnParams                  *string             `protobuf:"bytes,67,opt,name=stats_ttn_params" json:"stats_ttn_params,omitempty"`
	StatsSensorParams               *string             `protobuf:"bytes,68,opt,name=stats_sensor_params" json:"stats_sensor_params,omitempty"`
	Lnd_7318C                       *uint32             `protobuf:"varint,69,opt,name=lnd_7318c" json:"lnd_7318c,omitempty"`
	StatsBattery                    *string             `protobuf:"bytes,70,opt,name=stats_battery" json:"stats_battery,omitempty"`
	XXX_unrecognized                []byte              `json:"-"`
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

func (m *Telecast) GetDEPRECATED2017FEBDeviceIDString() string {
	if m != nil && m.DEPRECATED2017FEBDeviceIDString != nil {
		return *m.DEPRECATED2017FEBDeviceIDString
	}
	return ""
}

func (m *Telecast) GetDeviceId() uint32 {
	if m != nil && m.DeviceId != nil {
		return *m.DeviceId
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

func (m *Telecast) GetReplyType() TelecastReplyType {
	if m != nil && m.ReplyType != nil {
		return *m.ReplyType
	}
	return Telecast_NO_REPLY
}

func (m *Telecast) GetDEPRECATED2017FEBValue() uint32 {
	if m != nil && m.DEPRECATED2017FEBValue != nil {
		return *m.DEPRECATED2017FEBValue
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

func (m *Telecast) GetAltitude() int32 {
	if m != nil && m.Altitude != nil {
		return *m.Altitude
	}
	return 0
}

func (m *Telecast) GetBatVoltage() float32 {
	if m != nil && m.BatVoltage != nil {
		return *m.BatVoltage
	}
	return 0
}

func (m *Telecast) GetBatSoc() float32 {
	if m != nil && m.BatSoc != nil {
		return *m.BatSoc
	}
	return 0
}

func (m *Telecast) GetWirelessSnr() float32 {
	if m != nil && m.WirelessSnr != nil {
		return *m.WirelessSnr
	}
	return 0
}

func (m *Telecast) GetEnvTemp() float32 {
	if m != nil && m.EnvTemp != nil {
		return *m.EnvTemp
	}
	return 0
}

func (m *Telecast) GetEnvHumid() float32 {
	if m != nil && m.EnvHumid != nil {
		return *m.EnvHumid
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

func (m *Telecast) GetLnd_7318U() uint32 {
	if m != nil && m.Lnd_7318U != nil {
		return *m.Lnd_7318U
	}
	return 0
}

func (m *Telecast) GetLnd_7128Ec() uint32 {
	if m != nil && m.Lnd_7128Ec != nil {
		return *m.Lnd_7128Ec
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

func (m *Telecast) GetPmsPm01_0() uint32 {
	if m != nil && m.PmsPm01_0 != nil {
		return *m.PmsPm01_0
	}
	return 0
}

func (m *Telecast) GetPmsPm02_5() uint32 {
	if m != nil && m.PmsPm02_5 != nil {
		return *m.PmsPm02_5
	}
	return 0
}

func (m *Telecast) GetPmsPm10_0() uint32 {
	if m != nil && m.PmsPm10_0 != nil {
		return *m.PmsPm10_0
	}
	return 0
}

func (m *Telecast) GetPmsC00_30() uint32 {
	if m != nil && m.PmsC00_30 != nil {
		return *m.PmsC00_30
	}
	return 0
}

func (m *Telecast) GetPmsC00_50() uint32 {
	if m != nil && m.PmsC00_50 != nil {
		return *m.PmsC00_50
	}
	return 0
}

func (m *Telecast) GetPmsC01_00() uint32 {
	if m != nil && m.PmsC01_00 != nil {
		return *m.PmsC01_00
	}
	return 0
}

func (m *Telecast) GetPmsC02_50() uint32 {
	if m != nil && m.PmsC02_50 != nil {
		return *m.PmsC02_50
	}
	return 0
}

func (m *Telecast) GetPmsC05_00() uint32 {
	if m != nil && m.PmsC05_00 != nil {
		return *m.PmsC05_00
	}
	return 0
}

func (m *Telecast) GetPmsC10_00() uint32 {
	if m != nil && m.PmsC10_00 != nil {
		return *m.PmsC10_00
	}
	return 0
}

func (m *Telecast) GetPmsCsecs() uint32 {
	if m != nil && m.PmsCsecs != nil {
		return *m.PmsCsecs
	}
	return 0
}

func (m *Telecast) GetOpcPm01_0() float32 {
	if m != nil && m.OpcPm01_0 != nil {
		return *m.OpcPm01_0
	}
	return 0
}

func (m *Telecast) GetOpcPm02_5() float32 {
	if m != nil && m.OpcPm02_5 != nil {
		return *m.OpcPm02_5
	}
	return 0
}

func (m *Telecast) GetOpcPm10_0() float32 {
	if m != nil && m.OpcPm10_0 != nil {
		return *m.OpcPm10_0
	}
	return 0
}

func (m *Telecast) GetOpcC00_38() uint32 {
	if m != nil && m.OpcC00_38 != nil {
		return *m.OpcC00_38
	}
	return 0
}

func (m *Telecast) GetOpcC00_54() uint32 {
	if m != nil && m.OpcC00_54 != nil {
		return *m.OpcC00_54
	}
	return 0
}

func (m *Telecast) GetOpcC01_00() uint32 {
	if m != nil && m.OpcC01_00 != nil {
		return *m.OpcC01_00
	}
	return 0
}

func (m *Telecast) GetOpcC02_10() uint32 {
	if m != nil && m.OpcC02_10 != nil {
		return *m.OpcC02_10
	}
	return 0
}

func (m *Telecast) GetOpcC05_00() uint32 {
	if m != nil && m.OpcC05_00 != nil {
		return *m.OpcC05_00
	}
	return 0
}

func (m *Telecast) GetOpcC10_00() uint32 {
	if m != nil && m.OpcC10_00 != nil {
		return *m.OpcC10_00
	}
	return 0
}

func (m *Telecast) GetOpcCsecs() uint32 {
	if m != nil && m.OpcCsecs != nil {
		return *m.OpcCsecs
	}
	return 0
}

func (m *Telecast) GetEnvPressure() float32 {
	if m != nil && m.EnvPressure != nil {
		return *m.EnvPressure
	}
	return 0
}

func (m *Telecast) GetStatsCommsPowerFails() uint32 {
	if m != nil && m.StatsCommsPowerFails != nil {
		return *m.StatsCommsPowerFails
	}
	return 0
}

func (m *Telecast) GetBatCurrent() float32 {
	if m != nil && m.BatCurrent != nil {
		return *m.BatCurrent
	}
	return 0
}

func (m *Telecast) GetStatsIccid() string {
	if m != nil && m.StatsIccid != nil {
		return *m.StatsIccid
	}
	return ""
}

func (m *Telecast) GetStatsMotiondrops() uint32 {
	if m != nil && m.StatsMotiondrops != nil {
		return *m.StatsMotiondrops
	}
	return 0
}

func (m *Telecast) GetStatsDfu() string {
	if m != nil && m.StatsDfu != nil {
		return *m.StatsDfu
	}
	return ""
}

func (m *Telecast) GetCapturedAtDate() uint32 {
	if m != nil && m.CapturedAtDate != nil {
		return *m.CapturedAtDate
	}
	return 0
}

func (m *Telecast) GetCapturedAtTime() uint32 {
	if m != nil && m.CapturedAtTime != nil {
		return *m.CapturedAtTime
	}
	return 0
}

func (m *Telecast) GetCapturedAtOffset() uint32 {
	if m != nil && m.CapturedAtOffset != nil {
		return *m.CapturedAtOffset
	}
	return 0
}

func (m *Telecast) GetStatsOneshotSeconds() uint32 {
	if m != nil && m.StatsOneshotSeconds != nil {
		return *m.StatsOneshotSeconds
	}
	return 0
}

func (m *Telecast) GetStamp() uint32 {
	if m != nil && m.Stamp != nil {
		return *m.Stamp
	}
	return 0
}

func (m *Telecast) GetStampVersion() uint32 {
	if m != nil && m.StampVersion != nil {
		return *m.StampVersion
	}
	return 0
}

func (m *Telecast) GetStatsCpsi() string {
	if m != nil && m.StatsCpsi != nil {
		return *m.StatsCpsi
	}
	return ""
}

func (m *Telecast) GetStatsUptimeDays() uint32 {
	if m != nil && m.StatsUptimeDays != nil {
		return *m.StatsUptimeDays
	}
	return 0
}

func (m *Telecast) GetStatsDeviceLabel() string {
	if m != nil && m.StatsDeviceLabel != nil {
		return *m.StatsDeviceLabel
	}
	return ""
}

func (m *Telecast) GetStatsGpsParams() string {
	if m != nil && m.StatsGpsParams != nil {
		return *m.StatsGpsParams
	}
	return ""
}

func (m *Telecast) GetStatsServiceParams() string {
	if m != nil && m.StatsServiceParams != nil {
		return *m.StatsServiceParams
	}
	return ""
}

func (m *Telecast) GetStatsTtnParams() string {
	if m != nil && m.StatsTtnParams != nil {
		return *m.StatsTtnParams
	}
	return ""
}

func (m *Telecast) GetStatsSensorParams() string {
	if m != nil && m.StatsSensorParams != nil {
		return *m.StatsSensorParams
	}
	return ""
}

func (m *Telecast) GetLnd_7318C() uint32 {
	if m != nil && m.Lnd_7318C != nil {
		return *m.Lnd_7318C
	}
	return 0
}

func (m *Telecast) GetStatsBattery() string {
	if m != nil && m.StatsBattery != nil {
		return *m.StatsBattery
	}
	return ""
}

func init() {
	proto.RegisterType((*Telecast)(nil), "ttproto.Telecast")
	proto.RegisterEnum("ttproto.TelecastDeviceType", TelecastDeviceType_name, TelecastDeviceType_value)
	proto.RegisterEnum("ttproto.TelecastReplyType", TelecastReplyType_name, TelecastReplyType_value)
}

var fileDescriptor0 = []byte{
	// 922 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x95, 0x7b, 0x53, 0xdb, 0x46,
	0x14, 0xc5, 0x0b, 0x09, 0x60, 0x2f, 0xe0, 0x08, 0xf1, 0xba, 0x09, 0x24, 0x50, 0xfa, 0x08, 0x7d,
	0xc4, 0x91, 0x05, 0x2e, 0xf4, 0x5d, 0x63, 0x2b, 0x0c, 0xd3, 0x8c, 0xed, 0x01, 0x95, 0x36, 0x7f,
	0xed, 0x2c, 0xf2, 0x42, 0x34, 0xa3, 0xd7, 0x68, 0x57, 0xce, 0xf8, 0x1b, 0xf5, 0x63, 0x66, 0x75,
	0xa5, 0x75, 0x2c, 0xc2, 0x7f, 0xf6, 0xef, 0x9c, 0xd5, 0xde, 0xb3, 0x7b, 0xaf, 0x44, 0x6a, 0x52,
	0x36, 0x93, 0x34, 0x96, 0xb1, 0xb9, 0x24, 0x25, 0xfe, 0x38, 0xf8, 0xbf, 0x41, 0x6a, 0x2e, 0x0f,
	0xb8, 0xc7, 0x84, 0x34, 0x5b, 0x64, 0x79, 0xc4, 0xc7, 0xbe, 0xc7, 0xa9, 0x9c, 0x24, 0x1c, 0xe6,
	0xf6, 0xe7, 0x0f, 0x1b, 0xf6, 0x6e, 0xb3, 0xf4, 0x36, 0xb5, 0xaf, 0x59, 0x98, 0x5c, 0xe5, 0x31,
	0x5f, 0x92, 0xbd, 0x9e, 0x33, 0xbc, 0x74, 0xba, 0x1d, 0xd7, 0xe9, 0xd9, 0x56, 0xeb, 0xe4, 0x8d,
	0x73, 0xd6, 0x43, 0xf9, 0xa2, 0x77, 0x25, 0x53, 0x3f, 0xba, 0x83, 0xf9, 0xfd, 0xb9, 0xc3, 0xba,
	0xb9, 0x46, 0xea, 0xe5, 0xb3, 0xfd, 0x11, 0x3c, 0x52, 0x68, 0xd5, 0x7c, 0x42, 0x96, 0x42, 0x2e,
	0x04, 0xbb, 0xe3, 0xf0, 0x18, 0x3d, 0xeb, 0x64, 0xd9, 0x63, 0x89, 0xcc, 0x52, 0x3e, 0xa2, 0x4c,
	0xc2, 0x02, 0xc2, 0xd7, 0x84, 0xa4, 0x3c, 0x09, 0x26, 0x45, 0x4d, 0x8b, 0x8a, 0x35, 0xec, 0x9d,
	0xcf, 0x6b, 0x42, 0x0f, 0x96, 0xf4, 0x82, 0x6c, 0x7d, 0x56, 0xd2, 0x35, 0x0b, 0x32, 0x0e, 0x4b,
	0xb8, 0xad, 0x41, 0x6a, 0x01, 0x93, 0xbe, 0xcc, 0x46, 0x1c, 0x6a, 0x8a, 0xcc, 0xe7, 0xb5, 0x05,
	0x71, 0x74, 0x57, 0xa0, 0x3a, 0x22, 0x65, 0x62, 0x41, 0x69, 0x22, 0x8a, 0x2c, 0xe4, 0xc5, 0xdd,
	0x30, 0x49, 0xc7, 0x71, 0x20, 0xf3, 0x8a, 0x97, 0xd1, 0xa6, 0x22, 0xe4, 0x50, 0xc4, 0x1e, 0xac,
	0x20, 0xd8, 0x20, 0x2b, 0x1f, 0xfc, 0x54, 0x15, 0x25, 0x04, 0x15, 0x51, 0x0a, 0xab, 0xfa, 0x69,
	0x3c, 0x1a, 0x53, 0xc9, 0xc3, 0x04, 0x1a, 0x7a, 0xcb, 0x9c, 0xbc, 0xcf, 0x42, 0x75, 0x1c, 0x4f,
	0x10, 0x6d, 0x92, 0x55, 0xb5, 0x90, 0x4d, 0x68, 0x71, 0x4e, 0x2d, 0x30, 0xb0, 0xdc, 0x7b, 0xd8,
	0x86, 0xb5, 0x87, 0xf0, 0x11, 0x98, 0x0f, 0xe1, 0x63, 0x58, 0x7f, 0x08, 0xb7, 0x61, 0x03, 0x71,
	0x1e, 0x3c, 0x1a, 0xd1, 0x93, 0xa3, 0xd6, 0x69, 0x06, 0x9b, 0x88, 0x4c, 0x42, 0x10, 0xb5, 0xec,
	0x53, 0xee, 0xc1, 0x16, 0xb2, 0x5d, 0xb2, 0x21, 0x24, 0x93, 0x82, 0x66, 0x89, 0xf4, 0x43, 0x4e,
	0x43, 0x3f, 0xca, 0x24, 0x17, 0xb0, 0x8d, 0xea, 0x53, 0xb2, 0x56, 0xa8, 0x2c, 0x49, 0xe8, 0x98,
	0xa7, 0xc2, 0x8f, 0x23, 0x00, 0xbc, 0xbb, 0x1d, 0xb2, 0x5e, 0x48, 0xe5, 0xd5, 0x27, 0x2c, 0x65,
	0xa1, 0x80, 0xa7, 0x28, 0xee, 0x91, 0xed, 0x42, 0x94, 0x29, 0x8b, 0x44, 0xe8, 0x4b, 0xa9, 0xae,
	0xfd, 0x66, 0x92, 0x3f, 0xf8, 0x59, 0x75, 0xdb, 0x94, 0x7b, 0xdc, 0x1f, 0x4f, 0xd5, 0x1d, 0x54,
	0xb7, 0x48, 0xa3, 0x50, 0xe3, 0x88, 0x8b, 0xf7, 0xb1, 0x14, 0xb0, 0x8b, 0xfc, 0x19, 0x31, 0x0b,
	0xee, 0xc5, 0x61, 0x98, 0xaf, 0x15, 0x5c, 0x69, 0xcf, 0x75, 0xb8, 0x44, 0xb1, 0x24, 0xb4, 0x5a,
	0xd4, 0x82, 0x17, 0xf7, 0x98, 0x4d, 0xdb, 0xb0, 0x57, 0x65, 0x2d, 0x4b, 0xf9, 0xf6, 0x67, 0x99,
	0x67, 0x59, 0xf4, 0xc8, 0x82, 0x2f, 0xef, 0xb3, 0xb6, 0x05, 0x07, 0x55, 0xa6, 0xb6, 0xb0, 0xe0,
	0xab, 0x2a, 0xb3, 0x73, 0xdf, 0xd7, 0x55, 0xd6, 0xce, 0x7d, 0xdf, 0x54, 0x58, 0xbe, 0xad, 0x05,
	0xdf, 0xea, 0x3b, 0x42, 0x26, 0xb8, 0x27, 0xe0, 0xa5, 0xb6, 0xc5, 0x89, 0xa7, 0x63, 0x1c, 0x62,
	0xf7, 0x7c, 0x62, 0x79, 0x8c, 0xef, 0xaa, 0x0c, 0x63, 0x7c, 0x3f, 0xcb, 0x30, 0xc6, 0x29, 0xfc,
	0x30, 0xfb, 0x3c, 0x8c, 0x71, 0x0c, 0x3f, 0x56, 0x19, 0xc6, 0x78, 0x55, 0x65, 0x36, 0x6d, 0x59,
	0xd0, 0xac, 0x32, 0x8c, 0xf1, 0xba, 0xc2, 0x8a, 0x18, 0x96, 0x8e, 0x81, 0x0c, 0x63, 0xb4, 0x10,
	0xa9, 0x59, 0xc9, 0x67, 0x20, 0x51, 0x57, 0x24, 0xd4, 0xcc, 0x83, 0x8d, 0x05, 0x4e, 0xdb, 0xa2,
	0xb8, 0xbf, 0x24, 0xfe, 0xc0, 0x53, 0x7a, 0xcb, 0xfc, 0x40, 0xc0, 0x11, 0x2e, 0x2b, 0x07, 0xd1,
	0xcb, 0xd2, 0x94, 0x47, 0x12, 0x8e, 0x71, 0x95, 0x82, 0xc5, 0x2a, 0xdf, 0xf3, 0xd4, 0x44, 0xb5,
	0xb1, 0xc3, 0xa6, 0x9d, 0x19, 0xc6, 0x52, 0x35, 0xe5, 0x28, 0x8d, 0x13, 0x01, 0x3f, 0xe9, 0x72,
	0xca, 0xce, 0xbc, 0xcd, 0xe0, 0x04, 0xdd, 0x40, 0x8c, 0x99, 0xb7, 0x0f, 0x1d, 0x31, 0xc9, 0xe1,
	0x14, 0xcd, 0xf7, 0x94, 0x7c, 0x06, 0xe0, 0x67, 0xdd, 0x6c, 0xb3, 0x4a, 0x7c, 0x7b, 0xab, 0xba,
	0x0d, 0x7e, 0x41, 0xed, 0x39, 0xd9, 0xac, 0x34, 0x28, 0x55, 0xd1, 0x55, 0x11, 0x02, 0x7e, 0x45,
	0x79, 0x95, 0x2c, 0x28, 0x59, 0xbd, 0x10, 0x7e, 0xd3, 0x13, 0x8a, 0x7f, 0xa7, 0x13, 0xf4, 0xbb,
	0x3e, 0xca, 0xf2, 0x34, 0x12, 0xe1, 0xc3, 0x1f, 0xd5, 0x58, 0xe5, 0x38, 0x8e, 0xd8, 0x44, 0xc0,
	0x9f, 0xd5, 0xe6, 0x2f, 0x07, 0x2e, 0x60, 0x37, 0x3c, 0x80, 0xbf, 0x74, 0xbe, 0x42, 0xbb, 0x4b,
	0x84, 0x9e, 0xc4, 0x0e, 0x2a, 0xd3, 0x41, 0x13, 0x3c, 0x9d, 0x9d, 0xd3, 0xb3, 0xea, 0x3a, 0x29,
	0x23, 0xad, 0x74, 0xab, 0xe3, 0x2d, 0x78, 0x24, 0xe2, 0x54, 0x8b, 0x3d, 0xfd, 0xc2, 0xd7, 0xef,
	0x16, 0x0f, 0x9c, 0x99, 0x8c, 0xca, 0xaf, 0xee, 0x4f, 0xf2, 0x74, 0x02, 0x6f, 0x72, 0xe7, 0xc1,
	0x98, 0x90, 0x99, 0x2f, 0xca, 0x36, 0x59, 0xff, 0xa7, 0xff, 0x77, 0x7f, 0xf0, 0x6f, 0x9f, 0xf6,
	0x9c, 0xeb, 0x8b, 0xae, 0x43, 0xdd, 0x77, 0x43, 0xc7, 0xf8, 0x42, 0xbd, 0x44, 0x57, 0xce, 0xce,
	0x9d, 0x8b, 0xf3, 0x0b, 0x87, 0xf6, 0x3b, 0xfd, 0x81, 0x31, 0xa7, 0x8e, 0xb0, 0x7e, 0x35, 0x78,
	0xdb, 0xb9, 0xec, 0x76, 0xae, 0x5c, 0x63, 0xde, 0xac, 0x93, 0x05, 0xd7, 0xed, 0x0c, 0x87, 0xc6,
	0x23, 0x75, 0x6a, 0x8b, 0xae, 0xdb, 0x1f, 0xf4, 0x1c, 0xe3, 0x71, 0xf1, 0xfb, 0x5c, 0x7d, 0x0c,
	0x8c, 0x05, 0x73, 0x99, 0x2c, 0xb9, 0xee, 0x95, 0x73, 0x79, 0xed, 0x18, 0x8b, 0x07, 0xaf, 0x48,
	0xfd, 0xd3, 0x57, 0x63, 0x85, 0xd4, 0xfa, 0x03, 0x7a, 0xe9, 0x0c, 0xdf, 0xbe, 0x53, 0x7b, 0x99,
	0xa4, 0x81, 0x3f, 0xa9, 0xf3, 0xdf, 0xd0, 0xe9, 0xaa, 0xef, 0x88, 0x31, 0xf7, 0x31, 0x00, 0x00,
	0xff, 0xff, 0x3c, 0x3f, 0xc3, 0x61, 0x3e, 0x07, 0x00, 0x00,
}
