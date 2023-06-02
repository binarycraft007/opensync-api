# LocationSpeedTest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **string** | schema version # of a Mongo document | [default to null]
**Enable** | **bool** | if network is idle, ISP speedtests will be run on a gateway node every 3 hours | [optional] [default to null]
**SpeedCapable** | **bool** | The first ISP Speed Test on a gateway node had speed results above a minimum threshold | [optional] [default to null]
**EnableAllNodes** | **bool** | if network is idle, ISP speedtests will be run on every node twice a day | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


