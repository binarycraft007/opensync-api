# AuditTrail

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | [***ObjectId**](ObjectID.md) |  | [optional] [default to null]
**CustomerId** | **string** | The customer id | [default to null]
**LocationId** | **string** | The location id | [optional] [default to null]
**PartnerId** | **string** | The partner id | [optional] [default to null]
**Author** | **string** | The user who made the change | [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | The date and time the change was made | [default to null]
**ExpiresAt** | [**time.Time**](time.Time.md) | The date and time the change will expire | [default to null]
**Event** | **string** | The type of event that occurred | [default to null]
**Details** | **interface{}** | The details of the event | [default to null]
**XRequestId** | **string** | The request id | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


