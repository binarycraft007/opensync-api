# NodeResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [optional] [default to null]
**Ip** | **string** |  | [default to null]
**Nickname** | **string** |  | [optional] [default to null]
**DefaultName** | **string** |  | [optional] [default to null]
**FirmwareVersion** | **string** |  | [optional] [default to null]
**WifiConfig** | [**[]XAny**](x-any.md) |  | [optional] [default to null]
**Channel** | **string** |  | [optional] [default to null]
**ConnectionState** | **string** | connected, disconnected, or unavailable | [default to null]
**Var2gChannel** | **string** |  | [default to null]
**Var5gChannel** | **string** |  | [default to null]
**Var6gChannel** | **string** |  | [optional] [default to null]
**BackhaulType** | **string** | ethernet, wifi, or unknown | [default to null]
**ConnectedDeviceCount** | **float64** | count total devices with &#39;connected&#39; connectionState | [default to null]
**LeafToRoot** | **[]interface{}** | the list member is the parentId and freqBand, from itself to the root/gateway. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


