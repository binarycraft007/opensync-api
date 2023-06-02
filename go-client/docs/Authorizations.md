# Authorizations

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **string** | schema version # of a Mongo document | [default to null]
**NumPodsAuthorized** | **float64** | number of leaf pods that are authorized to be claimed and be a part of the Plume network | [optional] [default to 32.0]
**NumNodesAuthorized** | [**[]NodeAuthorization**](NodeAuthorization.md) |  | [optional] [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


