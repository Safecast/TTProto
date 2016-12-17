/* Automatically generated nanopb constant definitions */
/* Generated by nanopb-0.3.5 at Sat Dec 17 15:20:06 2016. */

#include "teletype.pb.h"

/* @@protoc_insertion_point(includes) */
#if PB_PROTO_HEADER_VERSION != 30
#error Regenerate this file with the current version of nanopb generator.
#endif



const pb_field_t teletype_Telecast_fields[56] = {
    PB_FIELD(  1, UENUM   , REQUIRED, STATIC  , FIRST, teletype_Telecast, DeviceType, DeviceType, 0),
    PB_FIELD(  2, STRING  , OPTIONAL, STATIC  , OTHER, teletype_Telecast, DeviceIDString, DeviceType, 0),
    PB_FIELD(  3, UINT32  , OPTIONAL, STATIC  , OTHER, teletype_Telecast, DeviceIDNumber, DeviceIDString, 0),
    PB_FIELD(  4, STRING  , OPTIONAL, STATIC  , OTHER, teletype_Telecast, Message, DeviceIDNumber, 0),
    PB_FIELD(  5, STRING  , OPTIONAL, STATIC  , OTHER, teletype_Telecast, CapturedAt, Message, 0),
    PB_FIELD(  6, UENUM   , OPTIONAL, STATIC  , OTHER, teletype_Telecast, Unit, CapturedAt, 0),
    PB_FIELD(  7, UINT32  , OPTIONAL, STATIC  , OTHER, teletype_Telecast, Value, Unit, 0),
    PB_FIELD(  8, FLOAT   , OPTIONAL, STATIC  , OTHER, teletype_Telecast, Latitude, Value, 0),
    PB_FIELD(  9, FLOAT   , OPTIONAL, STATIC  , OTHER, teletype_Telecast, Longitude, Latitude, 0),
    PB_FIELD( 10, UINT32  , OPTIONAL, STATIC  , OTHER, teletype_Telecast, Altitude, Longitude, 0),
    PB_FIELD( 11, FLOAT   , OPTIONAL, STATIC  , OTHER, teletype_Telecast, BatteryVoltage, Altitude, 0),
    PB_FIELD( 12, FLOAT   , OPTIONAL, STATIC  , OTHER, teletype_Telecast, BatterySOC, BatteryVoltage, 0),
    PB_FIELD( 13, FLOAT   , OPTIONAL, STATIC  , OTHER, teletype_Telecast, WirelessSNR, BatterySOC, 0),
    PB_FIELD( 14, FLOAT   , OPTIONAL, STATIC  , OTHER, teletype_Telecast, envTemperature, WirelessSNR, 0),
    PB_FIELD( 15, FLOAT   , OPTIONAL, STATIC  , OTHER, teletype_Telecast, envHumidity, envTemperature, 0),
    PB_FIELD( 16, UINT32  , OPTIONAL, STATIC  , OTHER, teletype_Telecast, RelayDevice1, envHumidity, 0),
    PB_FIELD( 17, UINT32  , OPTIONAL, STATIC  , OTHER, teletype_Telecast, RelayDevice2, RelayDevice1, 0),
    PB_FIELD( 18, UINT32  , OPTIONAL, STATIC  , OTHER, teletype_Telecast, RelayDevice3, RelayDevice2, 0),
    PB_FIELD( 19, UINT32  , OPTIONAL, STATIC  , OTHER, teletype_Telecast, RelayDevice4, RelayDevice3, 0),
    PB_FIELD( 20, UINT32  , OPTIONAL, STATIC  , OTHER, teletype_Telecast, RelayDevice5, RelayDevice4, 0),
    PB_FIELD( 21, UINT32  , OPTIONAL, STATIC  , OTHER, teletype_Telecast, cpm0, RelayDevice5, 0),
    PB_FIELD( 22, UINT32  , OPTIONAL, STATIC  , OTHER, teletype_Telecast, cpm1, cpm0, 0),
    PB_FIELD( 23, UINT32  , OPTIONAL, STATIC  , OTHER, teletype_Telecast, stats_uptime_minutes, cpm1, 0),
    PB_FIELD( 24, STRING  , OPTIONAL, STATIC  , OTHER, teletype_Telecast, stats_app_version, stats_uptime_minutes, 0),
    PB_FIELD( 25, STRING  , OPTIONAL, STATIC  , OTHER, teletype_Telecast, stats_device_params, stats_app_version, 0),
    PB_FIELD( 26, UINT32  , OPTIONAL, STATIC  , OTHER, teletype_Telecast, stats_transmitted_bytes, stats_device_params, 0),
    PB_FIELD( 27, UINT32  , OPTIONAL, STATIC  , OTHER, teletype_Telecast, stats_received_bytes, stats_transmitted_bytes, 0),
    PB_FIELD( 28, UINT32  , OPTIONAL, STATIC  , OTHER, teletype_Telecast, stats_oneshots, stats_received_bytes, 0),
    PB_FIELD( 29, UINT32  , OPTIONAL, STATIC  , OTHER, teletype_Telecast, stats_comms_resets, stats_oneshots, 0),
    PB_FIELD( 30, UINT32  , OPTIONAL, STATIC  , OTHER, teletype_Telecast, pms_pm01_0, stats_comms_resets, 0),
    PB_FIELD( 31, UINT32  , OPTIONAL, STATIC  , OTHER, teletype_Telecast, pms_pm02_5, pms_pm01_0, 0),
    PB_FIELD( 32, UINT32  , OPTIONAL, STATIC  , OTHER, teletype_Telecast, pms_pm10_0, pms_pm02_5, 0),
    PB_FIELD( 33, UINT32  , OPTIONAL, STATIC  , OTHER, teletype_Telecast, pms_c00_30, pms_pm10_0, 0),
    PB_FIELD( 34, UINT32  , OPTIONAL, STATIC  , OTHER, teletype_Telecast, pms_c00_50, pms_c00_30, 0),
    PB_FIELD( 35, UINT32  , OPTIONAL, STATIC  , OTHER, teletype_Telecast, pms_c01_00, pms_c00_50, 0),
    PB_FIELD( 36, UINT32  , OPTIONAL, STATIC  , OTHER, teletype_Telecast, pms_c02_50, pms_c01_00, 0),
    PB_FIELD( 37, UINT32  , OPTIONAL, STATIC  , OTHER, teletype_Telecast, pms_c05_00, pms_c02_50, 0),
    PB_FIELD( 38, UINT32  , OPTIONAL, STATIC  , OTHER, teletype_Telecast, pms_c10_00, pms_c05_00, 0),
    PB_FIELD( 39, UINT32  , OPTIONAL, STATIC  , OTHER, teletype_Telecast, pms_csecs, pms_c10_00, 0),
    PB_FIELD( 40, FLOAT   , OPTIONAL, STATIC  , OTHER, teletype_Telecast, opc_pm01_0, pms_csecs, 0),
    PB_FIELD( 41, FLOAT   , OPTIONAL, STATIC  , OTHER, teletype_Telecast, opc_pm02_5, opc_pm01_0, 0),
    PB_FIELD( 42, FLOAT   , OPTIONAL, STATIC  , OTHER, teletype_Telecast, opc_pm10_0, opc_pm02_5, 0),
    PB_FIELD( 43, UINT32  , OPTIONAL, STATIC  , OTHER, teletype_Telecast, opc_c00_38, opc_pm10_0, 0),
    PB_FIELD( 44, UINT32  , OPTIONAL, STATIC  , OTHER, teletype_Telecast, opc_c00_54, opc_c00_38, 0),
    PB_FIELD( 45, UINT32  , OPTIONAL, STATIC  , OTHER, teletype_Telecast, opc_c01_00, opc_c00_54, 0),
    PB_FIELD( 46, UINT32  , OPTIONAL, STATIC  , OTHER, teletype_Telecast, opc_c02_10, opc_c01_00, 0),
    PB_FIELD( 47, UINT32  , OPTIONAL, STATIC  , OTHER, teletype_Telecast, opc_c05_00, opc_c02_10, 0),
    PB_FIELD( 48, UINT32  , OPTIONAL, STATIC  , OTHER, teletype_Telecast, opc_c10_00, opc_c05_00, 0),
    PB_FIELD( 49, UINT32  , OPTIONAL, STATIC  , OTHER, teletype_Telecast, opc_csecs, opc_c10_00, 0),
    PB_FIELD( 50, FLOAT   , OPTIONAL, STATIC  , OTHER, teletype_Telecast, envPressure, opc_csecs, 0),
    PB_FIELD( 51, UINT32  , OPTIONAL, STATIC  , OTHER, teletype_Telecast, stats_comms_power_fails, envPressure, 0),
    PB_FIELD( 52, FLOAT   , OPTIONAL, STATIC  , OTHER, teletype_Telecast, BatteryCurrent, stats_comms_power_fails, 0),
    PB_FIELD( 53, STRING  , OPTIONAL, STATIC  , OTHER, teletype_Telecast, stats_cell, BatteryCurrent, 0),
    PB_FIELD( 54, UINT32  , OPTIONAL, STATIC  , OTHER, teletype_Telecast, stats_motiondrops, stats_cell, 0),
    PB_FIELD( 55, STRING  , OPTIONAL, STATIC  , OTHER, teletype_Telecast, stats_dfu, stats_motiondrops, 0),
    PB_LAST_FIELD
};


/* @@protoc_insertion_point(eof) */
