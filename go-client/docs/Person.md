# Person

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | [***ObjectId**](ObjectID.md) |  | [optional] [default to null]
**Nickname** | **string** |  | [default to null]
**ImageId** | **string** | unique identifier for referencing a Person&#39;s hosted profile image | [optional] [default to null]
**PrimaryDevice** | **string** | mac address of the Person&#39;s primary device | [optional] [default to null]
**HomeAwayNotification** | **bool** | notification when person is away from home | [optional] [default to null]
**AppTime** | **interface{}** | person&#39;s appTime config | [optional] [default to null]
**Permission** | **interface{}** | person&#39;s access permissions | [optional] [default to null]
**Profile** | [***PersonProfile**](PersonProfile.md) | person specific profile | [optional] [default to null]
**ServiceLinking** | **interface{}** | A link to a 3rd party&#39;s Person object | [optional] [default to null]
**NetworkId** | **string** | person and their devices behave as a group on this fronthaul network as it pertains to sharing access | [optional] [default to null]
**Email** | **string** | email address for creating co-manager access | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**LocationId** | [***ObjectId**](ObjectID.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


