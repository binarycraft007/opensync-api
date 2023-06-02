# LocationBackhaul

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **string** | schema version # of a Mongo document | [default to null]
**Mode** | **string** | auto || enable || disable | [default to null]
**WpaKey** | **string** | 63 characters, all CAPS, Hexadecimal | [default to null]
**DynamicBeacon** | **string** | Represents whether or not dynamic beaconing is available for this location. Can be either auto, enable, or disable | [optional] [default to null]
**Wds** | **string** | auto || enable || disable | [optional] [default to null]
**WpaMode** | **string** | auto || psk2 || sae-mixed | [optional] [default to null]
**WpaModeRealized** | **string** | psk2 || sae-mixed | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


