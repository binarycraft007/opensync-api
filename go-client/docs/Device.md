# Device

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | [***ObjectId**](ObjectID.md) |  | [optional] [default to null]
**Mac** | **string** |  | [default to null]
**KeyId** | **float64** | unique id of the WifiNetwork.keys[x] that the device is connected to or last connected to | [optional] [default to null]
**HostName** | **string** | host name from the device/user settings | [optional] [default to null]
**Model** | **string** | device model | [optional] [default to null]
**Medium** | **string** | wifi, ethernet, or moca | [optional] [default to null]
**Version** | **string** | schema version # of a Mongo document | [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**OnlineSince** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**OfflineSince** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**ConnectionState** | **string** |  | [optional] [default to null]
**ConnectionStateChangeAt** | [**time.Time**](time.Time.md) | time at which connectionStateChange last changed | [optional] [default to null]
**VapType** | **string** |  | [optional] [default to null]
**NetworkId** | **string** |  | [optional] [default to null]
**Favorite** | **bool** |  | [optional] [default to null]
**HiddenAt** | [**time.Time**](time.Time.md) | Time at which user removed device from device list | [optional] [default to null]
**AppTime** | **interface{}** | device&#39;s appTime config | [optional] [default to null]
**MobileAppDeviceUuid** | **string** | Unique identifier for mobile devices that the mobile app generates and controls | [optional] [default to null]
**LocationId** | [***ObjectId**](ObjectID.md) |  | [optional] [default to null]
**BandSteering** | [***DeviceBandSteering**](DeviceBandSteering.md) |  | [optional] [default to null]
**ClientSteering** | [***DeviceClientSteering**](DeviceClientSteering.md) |  | [optional] [default to null]
**Qos** | [***Qos**](Qos.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


