# RegisterResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [optional] [default to null]
**Email** | **string** | Fake email created on behalf of anonymous account. | [optional] [default to null]
**AccountId** | **string** |  | [optional] [default to null]
**LocationId** | **string** | ID of default location created during registration API call. | [optional] [default to null]
**Anonymous** | **bool** | ISP registered accounts are anonymous and do not include real emails and passwords | [optional] [default to null]
**Name** | **string** |  | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**Version** | **string** | schema version # of a Mongo document | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


