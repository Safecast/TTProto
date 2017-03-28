/* Automatically generated nanopb header */
/* Generated by nanopb-0.3.5 at Mon Mar 27 22:17:47 2017. */

#ifndef PB_TT_PB_H_INCLUDED
#define PB_TT_PB_H_INCLUDED
#include <pb.h>

/* @@protoc_insertion_point(includes) */
#if PB_PROTO_HEADER_VERSION != 30
#error Regenerate this file with the current version of nanopb generator.
#endif

#ifdef __cplusplus
extern "C" {
#endif

/* Enum definitions */
typedef enum _ttproto_Telecast_deviceType {
    ttproto_Telecast_deviceType_UNKNOWN_DEVICE_TYPE = 0,
    ttproto_Telecast_deviceType_BGEIGIE_NANO = 1,
    ttproto_Telecast_deviceType_SOLARCAST = 2,
    ttproto_Telecast_deviceType_TTAPP = 3,
    ttproto_Telecast_deviceType_TTNODE = 4,
    ttproto_Telecast_deviceType_TTGATE = 5,
    ttproto_Telecast_deviceType_TTSERVE = 6
} ttproto_Telecast_deviceType;

typedef enum _ttproto_Telecast_replyType {
    ttproto_Telecast_replyType_NO_REPLY = 0,
    ttproto_Telecast_replyType_ALLOWED = 1
} ttproto_Telecast_replyType;

/* Struct definitions */
typedef struct _ttproto_Telecast {
    bool has_device_type;
    ttproto_Telecast_deviceType device_type;
    pb_callback_t DEPRECATED2017FEBDeviceIDString;
    bool has_device_id;
    uint32_t device_id;
    bool has_message;
    char message[250];
    bool has_captured_at;
    char captured_at[40];
    bool has_reply_type;
    ttproto_Telecast_replyType reply_type;
    bool has_DEPRECATED2017FEBValue;
    uint32_t DEPRECATED2017FEBValue;
    bool has_latitude;
    float latitude;
    bool has_longitude;
    float longitude;
    bool has_altitude;
    int32_t altitude;
    bool has_bat_voltage;
    float bat_voltage;
    bool has_bat_soc;
    float bat_soc;
    bool has_wireless_snr;
    float wireless_snr;
    bool has_env_temp;
    float env_temp;
    bool has_env_humid;
    float env_humid;
    bool has_relay_device1;
    uint32_t relay_device1;
    bool has_relay_device2;
    uint32_t relay_device2;
    bool has_relay_device3;
    uint32_t relay_device3;
    bool has_relay_device4;
    uint32_t relay_device4;
    bool has_relay_device5;
    uint32_t relay_device5;
    bool has_lnd_7318u;
    uint32_t lnd_7318u;
    bool has_lnd_7128ec;
    uint32_t lnd_7128ec;
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
    bool has_env_pressure;
    float env_pressure;
    bool has_stats_comms_power_fails;
    uint32_t stats_comms_power_fails;
    bool has_bat_current;
    float bat_current;
    bool has_stats_iccid;
    char stats_iccid[128];
    bool has_stats_motiondrops;
    uint32_t stats_motiondrops;
    bool has_stats_dfu;
    char stats_dfu[128];
    bool has_captured_at_date;
    uint32_t captured_at_date;
    bool has_captured_at_time;
    uint32_t captured_at_time;
    bool has_captured_at_offset;
    uint32_t captured_at_offset;
    bool has_stats_oneshot_seconds;
    uint32_t stats_oneshot_seconds;
    bool has_stamp;
    uint32_t stamp;
    bool has_stamp_version;
    uint32_t stamp_version;
    bool has_stats_cpsi;
    char stats_cpsi[128];
    bool has_stats_uptime_days;
    uint32_t stats_uptime_days;
    bool has_stats_device_label;
    char stats_device_label[64];
    bool has_stats_gps_params;
    char stats_gps_params[64];
    bool has_stats_service_params;
    char stats_service_params[64];
    bool has_stats_ttn_params;
    char stats_ttn_params[64];
    bool has_stats_sensor_params;
    char stats_sensor_params[64];
    bool has_lnd_7318c;
    uint32_t lnd_7318c;
    bool has_stats_battery;
    char stats_battery[64];
    bool has_stats_module_fona;
    char stats_module_fona[64];
    bool has_stats_module_lora;
    char stats_module_lora[64];
    bool has_motion;
    bool motion;
    bool has_test;
    bool test;
    bool has_enc_temp;
    float enc_temp;
    bool has_enc_humid;
    float enc_humid;
    bool has_enc_pressure;
    float enc_pressure;
    bool has_errors_opc;
    uint32_t errors_opc;
    bool has_errors_pms;
    uint32_t errors_pms;
    bool has_errors_bme0;
    uint32_t errors_bme0;
    bool has_errors_bme1;
    uint32_t errors_bme1;
    bool has_errors_lora;
    uint32_t errors_lora;
    bool has_errors_fona;
    uint32_t errors_fona;
    bool has_errors_geiger;
    uint32_t errors_geiger;
    bool has_errors_max01;
    uint32_t errors_max01;
    bool has_errors_ugps;
    uint32_t errors_ugps;
    bool has_errors_twi;
    uint32_t errors_twi;
    bool has_errors_twi_info;
    char errors_twi_info[128];
    bool has_errors_lis;
    uint32_t errors_lis;
    bool has_errors_spi;
    uint32_t errors_spi;
    bool has_errors_connect_lora;
    uint32_t errors_connect_lora;
    bool has_errors_connect_fona;
    uint32_t errors_connect_fona;
    bool has_errors_connect_wireless;
    uint32_t errors_connect_wireless;
    bool has_errors_connect_data;
    uint32_t errors_connect_data;
    bool has_errors_connect_service;
    uint32_t errors_connect_service;
/* @@protoc_insertion_point(struct:ttproto_Telecast) */
} ttproto_Telecast;

/* Default values for struct fields */

/* Initializer values for message structs */
#define ttproto_Telecast_init_default            {false, (ttproto_Telecast_deviceType)0, {{NULL}, NULL}, false, 0, false, "", false, "", false, (ttproto_Telecast_replyType)0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, "", false, "", false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, "", false, 0, false, "", false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, "", false, 0, false, "", false, "", false, "", false, "", false, "", false, 0, false, "", false, "", false, "", false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, "", false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0}
#define ttproto_Telecast_init_zero               {false, (ttproto_Telecast_deviceType)0, {{NULL}, NULL}, false, 0, false, "", false, "", false, (ttproto_Telecast_replyType)0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, "", false, "", false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, "", false, 0, false, "", false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, "", false, 0, false, "", false, "", false, "", false, "", false, "", false, 0, false, "", false, "", false, "", false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, "", false, 0, false, 0, false, 0, false, 0, false, 0, false, 0, false, 0}

/* Field tags (for use in manual encoding/decoding) */
#define ttproto_Telecast_device_type_tag         1
#define ttproto_Telecast_DEPRECATED2017FEBDeviceIDString_tag 2
#define ttproto_Telecast_device_id_tag           3
#define ttproto_Telecast_message_tag             4
#define ttproto_Telecast_captured_at_tag         5
#define ttproto_Telecast_reply_type_tag          6
#define ttproto_Telecast_DEPRECATED2017FEBValue_tag 7
#define ttproto_Telecast_latitude_tag            8
#define ttproto_Telecast_longitude_tag           9
#define ttproto_Telecast_altitude_tag            10
#define ttproto_Telecast_bat_voltage_tag         11
#define ttproto_Telecast_bat_soc_tag             12
#define ttproto_Telecast_wireless_snr_tag        13
#define ttproto_Telecast_env_temp_tag            14
#define ttproto_Telecast_env_humid_tag           15
#define ttproto_Telecast_relay_device1_tag       16
#define ttproto_Telecast_relay_device2_tag       17
#define ttproto_Telecast_relay_device3_tag       18
#define ttproto_Telecast_relay_device4_tag       19
#define ttproto_Telecast_relay_device5_tag       20
#define ttproto_Telecast_lnd_7318u_tag           21
#define ttproto_Telecast_lnd_7128ec_tag          22
#define ttproto_Telecast_stats_uptime_minutes_tag 23
#define ttproto_Telecast_stats_app_version_tag   24
#define ttproto_Telecast_stats_device_params_tag 25
#define ttproto_Telecast_stats_transmitted_bytes_tag 26
#define ttproto_Telecast_stats_received_bytes_tag 27
#define ttproto_Telecast_stats_oneshots_tag      28
#define ttproto_Telecast_stats_comms_resets_tag  29
#define ttproto_Telecast_pms_pm01_0_tag          30
#define ttproto_Telecast_pms_pm02_5_tag          31
#define ttproto_Telecast_pms_pm10_0_tag          32
#define ttproto_Telecast_pms_c00_30_tag          33
#define ttproto_Telecast_pms_c00_50_tag          34
#define ttproto_Telecast_pms_c01_00_tag          35
#define ttproto_Telecast_pms_c02_50_tag          36
#define ttproto_Telecast_pms_c05_00_tag          37
#define ttproto_Telecast_pms_c10_00_tag          38
#define ttproto_Telecast_pms_csecs_tag           39
#define ttproto_Telecast_opc_pm01_0_tag          40
#define ttproto_Telecast_opc_pm02_5_tag          41
#define ttproto_Telecast_opc_pm10_0_tag          42
#define ttproto_Telecast_opc_c00_38_tag          43
#define ttproto_Telecast_opc_c00_54_tag          44
#define ttproto_Telecast_opc_c01_00_tag          45
#define ttproto_Telecast_opc_c02_10_tag          46
#define ttproto_Telecast_opc_c05_00_tag          47
#define ttproto_Telecast_opc_c10_00_tag          48
#define ttproto_Telecast_opc_csecs_tag           49
#define ttproto_Telecast_env_pressure_tag        50
#define ttproto_Telecast_stats_comms_power_fails_tag 51
#define ttproto_Telecast_bat_current_tag         52
#define ttproto_Telecast_stats_iccid_tag         53
#define ttproto_Telecast_stats_motiondrops_tag   54
#define ttproto_Telecast_stats_dfu_tag           55
#define ttproto_Telecast_captured_at_date_tag    56
#define ttproto_Telecast_captured_at_time_tag    57
#define ttproto_Telecast_captured_at_offset_tag  58
#define ttproto_Telecast_stats_oneshot_seconds_tag 59
#define ttproto_Telecast_stamp_tag               60
#define ttproto_Telecast_stamp_version_tag       61
#define ttproto_Telecast_stats_cpsi_tag          62
#define ttproto_Telecast_stats_uptime_days_tag   63
#define ttproto_Telecast_stats_device_label_tag  64
#define ttproto_Telecast_stats_gps_params_tag    65
#define ttproto_Telecast_stats_service_params_tag 66
#define ttproto_Telecast_stats_ttn_params_tag    67
#define ttproto_Telecast_stats_sensor_params_tag 68
#define ttproto_Telecast_lnd_7318c_tag           69
#define ttproto_Telecast_stats_battery_tag       70
#define ttproto_Telecast_stats_module_fona_tag   71
#define ttproto_Telecast_stats_module_lora_tag   72
#define ttproto_Telecast_motion_tag              73
#define ttproto_Telecast_test_tag                74
#define ttproto_Telecast_enc_temp_tag            75
#define ttproto_Telecast_enc_humid_tag           76
#define ttproto_Telecast_enc_pressure_tag        77
#define ttproto_Telecast_errors_opc_tag          78
#define ttproto_Telecast_errors_pms_tag          79
#define ttproto_Telecast_errors_bme0_tag         80
#define ttproto_Telecast_errors_bme1_tag         81
#define ttproto_Telecast_errors_lora_tag         82
#define ttproto_Telecast_errors_fona_tag         83
#define ttproto_Telecast_errors_geiger_tag       84
#define ttproto_Telecast_errors_max01_tag        85
#define ttproto_Telecast_errors_ugps_tag         86
#define ttproto_Telecast_errors_twi_tag          87
#define ttproto_Telecast_errors_twi_info_tag     88
#define ttproto_Telecast_errors_lis_tag          89
#define ttproto_Telecast_errors_spi_tag          90
#define ttproto_Telecast_errors_connect_lora_tag 91
#define ttproto_Telecast_errors_connect_fona_tag 92
#define ttproto_Telecast_errors_connect_wireless_tag 93
#define ttproto_Telecast_errors_connect_data_tag 94
#define ttproto_Telecast_errors_connect_service_tag 95

/* Struct field encoding specification for nanopb */
extern const pb_field_t ttproto_Telecast_fields[96];

/* Maximum encoded size of messages (where known) */

/* Message IDs (where set with "msgid" option) */
#ifdef PB_MSGID

#define TT_MESSAGES \


#endif

#ifdef __cplusplus
} /* extern "C" */
#endif
/* @@protoc_insertion_point(eof) */

#endif
