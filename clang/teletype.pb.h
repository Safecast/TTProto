/* Automatically generated nanopb header */
/* Generated by nanopb-0.3.5 at Mon Jan 30 12:23:38 2017. */

#ifndef PB_TELETYPE_PB_H_INCLUDED
#define PB_TELETYPE_PB_H_INCLUDED
#include <pb.h>

/* @@protoc_insertion_point(includes) */
#if PB_PROTO_HEADER_VERSION != 30
#error Regenerate this file with the current version of nanopb generator.
#endif

#ifdef __cplusplus
extern "C" {
#endif

/* Enum definitions */
typedef enum _teletype_Telecast_deviceType {
    teletype_Telecast_deviceType_UNKNOWN_DEVICE_TYPE = 0,
    teletype_Telecast_deviceType_BGEIGIE_NANO = 1,
    teletype_Telecast_deviceType_SOLARCAST = 2,
    teletype_Telecast_deviceType_TTAPP = 3,
    teletype_Telecast_deviceType_TTNODE = 4,
    teletype_Telecast_deviceType_TTGATE = 5,
    teletype_Telecast_deviceType_TTSERVE = 6
} teletype_Telecast_deviceType;

typedef enum _teletype_Telecast_unit {
    teletype_Telecast_unit_UNKNOWN_UNIT = 0,
    teletype_Telecast_unit_CPM = 1
} teletype_Telecast_unit;

/* Struct definitions */
typedef struct _teletype_Telecast {
    teletype_Telecast_deviceType DeviceType;
    bool has_DeviceIDString;
    char DeviceIDString[32];
    bool has_DeviceIDNumber;
    uint32_t DeviceIDNumber;
    bool has_Message;
    char Message[250];
    bool has_CapturedAt;
    char CapturedAt[40];
    bool has_Unit;
    teletype_Telecast_unit Unit;
    bool has_Value;
    uint32_t Value;
    bool has_Latitude;
    float Latitude;
    bool has_Longitude;
    float Longitude;
    bool has_Altitude;
    uint32_t Altitude;
    bool has_BatteryVoltage;
    float BatteryVoltage;
    bool has_BatterySOC;
    float BatterySOC;
    bool has_WirelessSNR;
    float WirelessSNR;
    bool has_envTemperature;
    float envTemperature;
    bool has_envHumidity;
    float envHumidity;
    bool has_RelayDevice1;
    uint32_t RelayDevice1;
    bool has_RelayDevice2;
    uint32_t RelayDevice2;
    bool has_RelayDevice3;
    uint32_t RelayDevice3;
    bool has_RelayDevice4;
    uint32_t RelayDevice4;
    bool has_RelayDevice5;
    uint32_t RelayDevice5;
    bool has_cpm0;
    uint32_t cpm0;
    bool has_cpm1;
    uint32_t cpm1;
    bool has_stats_uptime_minutes;
    uint32_t stats_uptime_minutes;
    bool has_stats_app_version;
    char stats_app_version[64];
    bool has_stats_device_params;
    char stats_device_params[64];
    bool has_stats_transmitted_bytes;
    uint32_t stats_transmitted_bytes;
    bool has_stats_received_bytes;
    uint32_t stats_received_bytes;
    bool has_stats_oneshots;
    uint32_t stats_oneshots;
    bool has_stats_comms_resets;
    uint32_t stats_comms_resets;
    bool has_pms_pm01_0;
    uint32_t pms_pm01_0;
    bool has_pms_pm02_5;
    uint32_t pms_pm02_5;
    bool has_pms_pm10_0;
    uint32_t pms_pm10_0;
    bool has_pms_c00_30;
    uint32_t pms_c00_30;
    bool has_pms_c00_50;
    uint32_t pms_c00_50;
    bool has_pms_c01_00;
    uint32_t pms_c01_00;
    bool has_pms_c02_50;
    uint32_t pms_c02_50;
    bool has_pms_c05_00;
    uint32_t pms_c05_00;
    bool has_pms_c10_00;
    uint32_t pms_c10_00;
    bool has_pms_csecs;
    uint32_t pms_csecs;
    bool has_opc_pm01_0;
    float opc_pm01_0;
    bool has_opc_pm02_5;
    float opc_pm02_5;
    bool has_opc_pm10_0;
    float opc_pm10_0;
    bool has_opc_c00_38;
    uint32_t opc_c00_38;
    bool has_opc_c00_54;
    uint32_t opc_c00_54;
    bool has_opc_c01_00;
    uint32_t opc_c01_00;
    bool has_opc_c02_10;
    uint32_t opc_c02_10;
    bool has_opc_c05_00;
    uint32_t opc_c05_00;
    bool has_opc_c10_00;
    uint32_t opc_c10_00;
    bool has_opc_csecs;
    uint32_t opc_csecs;
    bool has_envPressure;
    float envPressure;
    bool has_stats_comms_power_fails;
    uint32_t stats_comms_power_fails;
    bool has_BatteryCurrent;
    float BatteryCurrent;
    bool has_stats_cell;
    char stats_cell[128];
    bool has_stats_motiondrops;
    uint32_t stats_motiondrops;
    bool has_stats_dfu;
    char stats_dfu[128];
    bool has_CapturedAtDate;
    uint32_t CapturedAtDate;
    bool has_CapturedAtTime;
    uint32_t CapturedAtTime;
    bool has_CapturedAtOffset;
    uint32_t CapturedAtOffset;
    bool has_stats_oneshot_seconds;
    uint32_t stats_oneshot_seconds;
    bool has_stamp;
    uint32_t stamp;
/* @@protoc_insertion_point(struct:teletype_Telecast) */
} teletype_Telecast;

/* Default values for struct fields */

/* Initializer values for message structs */
#define teletype_Telecast_init_default           {(teletype_Telecast_deviceType)0, false, "", false, 0, false, "", false, "", false, (teletype_Telecast_unit)0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, "", false, "", false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, "", false, 0, false, "", false, 0, false, 0, false, 0, false, 0, false, 0}
#define teletype_Telecast_init_zero              {(teletype_Telecast_deviceType)0, false, "", false, 0, false, "", false, "", false, (teletype_Telecast_unit)0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, "", false, "", false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, "", false, 0, false, "", false, 0, false, 0, false, 0, false, 0, false, 0}

/* Field tags (for use in manual encoding/decoding) */
#define teletype_Telecast_DeviceType_tag         1
#define teletype_Telecast_DeviceIDString_tag     2
#define teletype_Telecast_DeviceIDNumber_tag     3
#define teletype_Telecast_Message_tag            4
#define teletype_Telecast_CapturedAt_tag         5
#define teletype_Telecast_Unit_tag               6
#define teletype_Telecast_Value_tag              7
#define teletype_Telecast_Latitude_tag           8
#define teletype_Telecast_Longitude_tag          9
#define teletype_Telecast_Altitude_tag           10
#define teletype_Telecast_BatteryVoltage_tag     11
#define teletype_Telecast_BatterySOC_tag         12
#define teletype_Telecast_WirelessSNR_tag        13
#define teletype_Telecast_envTemperature_tag     14
#define teletype_Telecast_envHumidity_tag        15
#define teletype_Telecast_RelayDevice1_tag       16
#define teletype_Telecast_RelayDevice2_tag       17
#define teletype_Telecast_RelayDevice3_tag       18
#define teletype_Telecast_RelayDevice4_tag       19
#define teletype_Telecast_RelayDevice5_tag       20
#define teletype_Telecast_cpm0_tag               21
#define teletype_Telecast_cpm1_tag               22
#define teletype_Telecast_stats_uptime_minutes_tag 23
#define teletype_Telecast_stats_app_version_tag  24
#define teletype_Telecast_stats_device_params_tag 25
#define teletype_Telecast_stats_transmitted_bytes_tag 26
#define teletype_Telecast_stats_received_bytes_tag 27
#define teletype_Telecast_stats_oneshots_tag     28
#define teletype_Telecast_stats_comms_resets_tag 29
#define teletype_Telecast_pms_pm01_0_tag         30
#define teletype_Telecast_pms_pm02_5_tag         31
#define teletype_Telecast_pms_pm10_0_tag         32
#define teletype_Telecast_pms_c00_30_tag         33
#define teletype_Telecast_pms_c00_50_tag         34
#define teletype_Telecast_pms_c01_00_tag         35
#define teletype_Telecast_pms_c02_50_tag         36
#define teletype_Telecast_pms_c05_00_tag         37
#define teletype_Telecast_pms_c10_00_tag         38
#define teletype_Telecast_pms_csecs_tag          39
#define teletype_Telecast_opc_pm01_0_tag         40
#define teletype_Telecast_opc_pm02_5_tag         41
#define teletype_Telecast_opc_pm10_0_tag         42
#define teletype_Telecast_opc_c00_38_tag         43
#define teletype_Telecast_opc_c00_54_tag         44
#define teletype_Telecast_opc_c01_00_tag         45
#define teletype_Telecast_opc_c02_10_tag         46
#define teletype_Telecast_opc_c05_00_tag         47
#define teletype_Telecast_opc_c10_00_tag         48
#define teletype_Telecast_opc_csecs_tag          49
#define teletype_Telecast_envPressure_tag        50
#define teletype_Telecast_stats_comms_power_fails_tag 51
#define teletype_Telecast_BatteryCurrent_tag     52
#define teletype_Telecast_stats_cell_tag         53
#define teletype_Telecast_stats_motiondrops_tag  54
#define teletype_Telecast_stats_dfu_tag          55
#define teletype_Telecast_CapturedAtDate_tag     56
#define teletype_Telecast_CapturedAtTime_tag     57
#define teletype_Telecast_CapturedAtOffset_tag   58
#define teletype_Telecast_stats_oneshot_seconds_tag 59
#define teletype_Telecast_stamp_tag              60

/* Struct field encoding specification for nanopb */
extern const pb_field_t teletype_Telecast_fields[61];

/* Maximum encoded size of messages (where known) */
#define teletype_Telecast_size                   1066

/* Message IDs (where set with "msgid" option) */
#ifdef PB_MSGID

#define TELETYPE_MESSAGES \


#endif

#ifdef __cplusplus
} /* extern "C" */
#endif
/* @@protoc_insertion_point(eof) */

#endif
