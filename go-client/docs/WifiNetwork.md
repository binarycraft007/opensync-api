# WifiNetwork

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EncryptionKey** | **string** |  | [default to null]
**Ssid** | **string** |  | [default to null]
**MasterKeyIndex** | **float64** | keys[] array index of the password to be used as the original home password for backwards compatibility for locations that don&#39;t use home pass (yet) | [default to 0.0]
**Encryption** | **string** |  | [optional] [default to null]
**WpaMode** | **string** | open || psk-mixed || sae-mixed || psk2 || sae | [optional] [default to null]
**EncryptionMode** | **string** |  | [optional] [default to null]
**RadiusServerIp** | **string** |  | [optional] [default to null]
**RadiusServerPort** | **float64** |  | [optional] [default to null]
**SsidBroadcast** | **bool** |  | [optional] [default to null]
**Uapsd** | **bool** |  | [optional] [default to null]
**Enabled** | **bool** | enables the VAPs and radios and backhauls | [optional] [default to null]
**DisableDefaultServiceNetwork** | **bool** | disables the primary network VAP | [optional] [default to null]
**GroupRekey** | **string** | auto || enable || disable | [optional] [default to null]
**FastTransition** | **string** | auto || enable || disable | [optional] [default to null]
**MinWifiMode24** | **string** | auto || 11b || 11g || 11n | [optional] [default to null]
**PrivateMode** | **bool** |  | [optional] [default to null]
**PrivateModeUpdatedAt** | [**time.Time**](time.Time.md) | Privacy Mode Visibility with timestamp in Frontline | [optional] [default to null]
**DevicesVisibleToGuests** | **[]string** | the list of the device mac addresses | [optional] [default to null]
**DefaultSecurityPolicy** | [***WifiSecurityPolicy**](WifiSecurityPolicy.md) | location&#39;s default security policy | [optional] [default to null]
**Content** | **[]string** | content filter of the master key[masterKeyIndex]: kids || teenagers || adBlocking || spamPhishMalware || adultAndSensitive || workAppropriate | [optional] [default to null]
**AppliesToAllDevicesSecurityPolicy** | [***WifiSecurityPolicy**](WifiSecurityPolicy.md) | location&#39;s default security policy | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**Version** | **string** | schema version # of a Mongo document | [default to null]
**WpaModeDeferredExpiresAt** | **string** | Hold off setting wpaMode on devices for WPA3 onboarding for a set amount of time. | [optional] [default to null]
**Id** | [***ObjectId**](ObjectID.md) |  | [optional] [default to null]
**LocationId** | [***ObjectId**](ObjectID.md) |  | [optional] [default to null]
**Keys** | [**[]WifiNetworkKey**](WifiNetworkKey.md) |  | [optional] [default to null]
**Persons** | [**[]WifiSecurityPolicy**](WifiSecurityPolicy.md) |  | [optional] [default to null]
**Devices** | [**[]WifiSecurityPolicy**](WifiSecurityPolicy.md) |  | [optional] [default to null]
**AccessZones** | [**[]WifiAccessZone**](WifiAccessZone.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


