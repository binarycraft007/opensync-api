# WifiNetworkKey

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float64** |  | [default to null]
**EncryptionKey** | **string** | a password for parent WifiNetwork.ssid | [default to null]
**AccessZone** | **string** | home | guests | internetAccessOnly | [default to null]
**AccessZoneId** | **float64** | home:0 | internetAccessOnly:1 | guests:2 | guests:3 | guests:4 | [optional] [default to null]
**Format** | **string** | encryptionKey | phoneNumber | [default to null]
**Enable** | **bool** | devices can connect using this encryptionKey | [optional] [default to null]
**Content** | **[]string** | content filter IDs: kids || teenagers || adBlocking || spamPhishMalware || adultAndSensitive || workAppropriate | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**ExpiresAt** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**Version** | **string** | schema version # of a Mongo document | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


