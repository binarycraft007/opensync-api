# WifiSecurityPolicy

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Person Mongo ID || device mac | [default to null]
**Content** | **[]string** | content filter IDs: kids || teenagers || adBlocking || spamPhishMalware || adultAndSensitive || workAppropriate | [optional] [default to null]
**IotProtect** | **bool** | iotProtect enable/disable (enable requires secureAndProtect already enabled) | [optional] [default to null]
**IotProtectUpdatedAt** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**IotProtectReason** | **string** | optional reason used to identify why it is being disabled (device level only) | [optional] [default to null]
**Websites** | **interface{}** | has two nested properties: whitelist[dnsObject] and blacklist[dnsObject] | [optional] [default to null]
**SecureWebsites** | **interface{}** | has two nested properties: whitelist[dnsObject] and blacklist[dnsObject] | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


