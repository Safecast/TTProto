/* Automatically generated nanopb constant definitions */
/* Generated by nanopb-0.3.5 at Mon Jun 12 23:00:04 2017. */

#include "tt.pb.h"

/* @@protoc_insertion_point(includes) */
#if PB_PROTO_HEADER_VERSION != 30
#error Regenerate this file with the current version of nanopb generator.
#endif



const pb_field_t ttproto_Telecast_fields[110] = {
    PB_FIELD(  1, UENUM   , OPTIONAL, STATIC  , FIRST, ttproto_Telecast, device_type, device_type, 0),
    PB_FIELD(  2, STRING  , OPTIONAL, CALLBACK, OTHER, ttproto_Telecast, DEPRECATED2017FEBDeviceIDString, device_type, 0),
    PB_FIELD(  3, UINT32  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, device_id, DEPRECATED2017FEBDeviceIDString, 0),
    PB_FIELD(  4, STRING  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, message, device_id, 0),
    PB_FIELD(  5, STRING  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, captured_at, message, 0),
    PB_FIELD(  6, UENUM   , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, reply_type, captured_at, 0),
    PB_FIELD(  7, UINT32  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, DEPRECATED2017FEBValue, reply_type, 0),
    PB_FIELD(  8, FLOAT   , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, latitude, DEPRECATED2017FEBValue, 0),
    PB_FIELD(  9, FLOAT   , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, longitude, latitude, 0),
    PB_FIELD( 10, INT32   , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, altitude, longitude, 0),
    PB_FIELD( 11, FLOAT   , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, bat_voltage, altitude, 0),
    PB_FIELD( 12, FLOAT   , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, bat_soc, bat_voltage, 0),
    PB_FIELD( 13, FLOAT   , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, wireless_snr, bat_soc, 0),
    PB_FIELD( 14, FLOAT   , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, env_temp, wireless_snr, 0),
    PB_FIELD( 15, FLOAT   , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, env_humid, env_temp, 0),
    PB_FIELD( 16, UINT32  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, relay_device1, env_humid, 0),
    PB_FIELD( 17, UINT32  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, relay_device2, relay_device1, 0),
    PB_FIELD( 18, UINT32  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, relay_device3, relay_device2, 0),
    PB_FIELD( 19, UINT32  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, relay_device4, relay_device3, 0),
    PB_FIELD( 20, UINT32  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, relay_device5, relay_device4, 0),
    PB_FIELD( 21, UINT32  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, lnd_7318u, relay_device5, 0),
    PB_FIELD( 22, UINT32  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, lnd_7128ec, lnd_7318u, 0),
    PB_FIELD( 23, UINT32  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, stats_uptime_minutes, lnd_7128ec, 0),
    PB_FIELD( 24, STRING  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, stats_app_version, stats_uptime_minutes, 0),
    PB_FIELD( 25, STRING  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, stats_device_params, stats_app_version, 0),
    PB_FIELD( 26, UINT32  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, stats_transmitted_bytes, stats_device_params, 0),
    PB_FIELD( 27, UINT32  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, stats_received_bytes, stats_transmitted_bytes, 0),
    PB_FIELD( 28, UINT32  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, stats_oneshots, stats_received_bytes, 0),
    PB_FIELD( 29, UINT32  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, stats_comms_resets, stats_oneshots, 0),
    PB_FIELD( 30, UINT32  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, pms_pm01_0, stats_comms_resets, 0),
    PB_FIELD( 31, UINT32  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, pms_pm02_5, pms_pm01_0, 0),
    PB_FIELD( 32, UINT32  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, pms_pm10_0, pms_pm02_5, 0),
    PB_FIELD( 33, UINT32  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, pms_c00_30, pms_pm10_0, 0),
    PB_FIELD( 34, UINT32  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, pms_c00_50, pms_c00_30, 0),
    PB_FIELD( 35, UINT32  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, pms_c01_00, pms_c00_50, 0),
    PB_FIELD( 36, UINT32  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, pms_c02_50, pms_c01_00, 0),
    PB_FIELD( 37, UINT32  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, pms_c05_00, pms_c02_50, 0),
    PB_FIELD( 38, UINT32  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, pms_c10_00, pms_c05_00, 0),
    PB_FIELD( 39, UINT32  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, pms_csecs, pms_c10_00, 0),
    PB_FIELD( 40, FLOAT   , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, opc_pm01_0, pms_csecs, 0),
    PB_FIELD( 41, FLOAT   , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, opc_pm02_5, opc_pm01_0, 0),
    PB_FIELD( 42, FLOAT   , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, opc_pm10_0, opc_pm02_5, 0),
    PB_FIELD( 43, UINT32  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, opc_c00_38, opc_pm10_0, 0),
    PB_FIELD( 44, UINT32  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, opc_c00_54, opc_c00_38, 0),
    PB_FIELD( 45, UINT32  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, opc_c01_00, opc_c00_54, 0),
    PB_FIELD( 46, UINT32  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, opc_c02_10, opc_c01_00, 0),
    PB_FIELD( 47, UINT32  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, opc_c05_00, opc_c02_10, 0),
    PB_FIELD( 48, UINT32  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, opc_c10_00, opc_c05_00, 0),
    PB_FIELD( 49, UINT32  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, opc_csecs, opc_c10_00, 0),
    PB_FIELD( 50, FLOAT   , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, env_pressure, opc_csecs, 0),
    PB_FIELD( 51, UINT32  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, stats_comms_power_fails, env_pressure, 0),
    PB_FIELD( 52, FLOAT   , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, bat_current, stats_comms_power_fails, 0),
    PB_FIELD( 53, STRING  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, stats_iccid, bat_current, 0),
    PB_FIELD( 54, UINT32  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, stats_motion_events, stats_iccid, 0),
    PB_FIELD( 55, STRING  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, stats_dfu, stats_motion_events, 0),
    PB_FIELD( 56, UINT32  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, captured_at_date, stats_dfu, 0),
    PB_FIELD( 57, UINT32  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, captured_at_time, captured_at_date, 0),
    PB_FIELD( 58, UINT32  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, captured_at_offset, captured_at_time, 0),
    PB_FIELD( 59, UINT32  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, stats_oneshot_seconds, captured_at_offset, 0),
    PB_FIELD( 60, UINT32  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, stamp, stats_oneshot_seconds, 0),
    PB_FIELD( 61, UINT32  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, stamp_version, stamp, 0),
    PB_FIELD( 62, STRING  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, stats_cpsi, stamp_version, 0),
    PB_FIELD( 63, UINT32  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, stats_uptime_days, stats_cpsi, 0),
    PB_FIELD( 64, STRING  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, stats_device_label, stats_uptime_days, 0),
    PB_FIELD( 65, STRING  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, stats_gps_params, stats_device_label, 0),
    PB_FIELD( 66, STRING  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, stats_service_params, stats_gps_params, 0),
    PB_FIELD( 67, STRING  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, stats_ttn_params, stats_service_params, 0),
    PB_FIELD( 68, STRING  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, stats_sensor_params, stats_ttn_params, 0),
    PB_FIELD( 69, UINT32  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, lnd_7318c, stats_sensor_params, 0),
    PB_FIELD( 70, STRING  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, stats_battery, lnd_7318c, 0),
    PB_FIELD( 71, STRING  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, stats_module_fona, stats_battery, 0),
    PB_FIELD( 72, STRING  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, stats_module_lora, stats_module_fona, 0),
    PB_FIELD( 73, BOOL    , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, motion, stats_module_lora, 0),
    PB_FIELD( 74, BOOL    , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, test, motion, 0),
    PB_FIELD( 75, FLOAT   , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, enc_temp, test, 0),
    PB_FIELD( 76, FLOAT   , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, enc_humid, enc_temp, 0),
    PB_FIELD( 77, FLOAT   , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, enc_pressure, enc_humid, 0),
    PB_FIELD( 78, UINT32  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, errors_opc, enc_pressure, 0),
    PB_FIELD( 79, UINT32  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, errors_pms, errors_opc, 0),
    PB_FIELD( 80, UINT32  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, errors_bme0, errors_pms, 0),
    PB_FIELD( 81, UINT32  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, errors_bme1, errors_bme0, 0),
    PB_FIELD( 82, UINT32  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, errors_lora, errors_bme1, 0),
    PB_FIELD( 83, UINT32  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, errors_fona, errors_lora, 0),
    PB_FIELD( 84, UINT32  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, errors_geiger, errors_fona, 0),
    PB_FIELD( 85, UINT32  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, errors_max01, errors_geiger, 0),
    PB_FIELD( 86, UINT32  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, errors_ugps, errors_max01, 0),
    PB_FIELD( 87, UINT32  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, errors_twi, errors_ugps, 0),
    PB_FIELD( 88, STRING  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, errors_twi_info, errors_twi, 0),
    PB_FIELD( 89, UINT32  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, errors_lis, errors_twi_info, 0),
    PB_FIELD( 90, UINT32  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, errors_spi, errors_lis, 0),
    PB_FIELD( 91, UINT32  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, errors_connect_lora, errors_spi, 0),
    PB_FIELD( 92, UINT32  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, errors_connect_fona, errors_connect_lora, 0),
    PB_FIELD( 93, UINT32  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, errors_connect_wireless, errors_connect_fona, 0),
    PB_FIELD( 94, UINT32  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, errors_connect_data, errors_connect_wireless, 0),
    PB_FIELD( 95, UINT32  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, errors_connect_service, errors_connect_data, 0),
    PB_FIELD( 96, UINT32  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, motion_began_offset, errors_connect_service, 0),
    PB_FIELD( 97, UINT32  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, errors_connect_gateway, motion_began_offset, 0),
    PB_FIELD( 98, UINT32  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, stats_comms_ant_fails, errors_connect_gateway, 0),
    PB_FIELD( 99, UINT32  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, lnd_712u, stats_comms_ant_fails, 0),
    PB_FIELD(100, UINT32  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, lnd_78017w, lnd_712u, 0),
    PB_FIELD(101, UINT32  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, stats_overcurrent_events, lnd_78017w, 0),
    PB_FIELD(102, FLOAT   , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, pms_std01_0, stats_overcurrent_events, 0),
    PB_FIELD(103, FLOAT   , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, pms_std02_5, pms_std01_0, 0),
    PB_FIELD(104, FLOAT   , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, pms_std10_0, pms_std02_5, 0),
    PB_FIELD(105, FLOAT   , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, opc_std01_0, pms_std10_0, 0),
    PB_FIELD(106, FLOAT   , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, opc_std02_5, opc_std01_0, 0),
    PB_FIELD(107, FLOAT   , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, opc_std10_0, opc_std02_5, 0),
    PB_FIELD(108, UINT32  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, errors_mtu, opc_std10_0, 0),
    PB_FIELD(109, UINT32  , OPTIONAL, STATIC  , OTHER, ttproto_Telecast, stats_seqno, errors_mtu, 0),
    PB_LAST_FIELD
};


/* @@protoc_insertion_point(eof) */
