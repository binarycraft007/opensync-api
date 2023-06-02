# CapabilitiesResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkConfiguration** | [***NetworkConfigurationCapabilitiesResponse**](NetworkConfigurationCapabilitiesResponse.md) |  | [default to null]
**Security** | [***SecurityCapabilitiesResponse**](SecurityCapabilitiesResponse.md) | Firmware capabilities for security features. | [default to null]
**IspSpeedTest** | [***Capability**](Capability.md) | ISP speed test availability for a location | [default to null]
**DeviceFreeze** | [***Capability**](Capability.md) | support for device freeze templates | [default to null]
**MultiPasswordSSID** | [***Capability**](Capability.md) | support for non-host access (a.k.a., Single SSID, Multi Password onboarding) | [default to null]
**WifiMotion** | [***Capability**](Capability.md) | support for wifi motion features | [default to null]
**Version** | **string** | schema version # of a Mongo document | [default to null]
**CapabilitiesChangedAt** | [**time.Time**](time.Time.md) |  | [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


