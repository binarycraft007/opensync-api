# SecurityConfiguration

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **string** | schema version # of a Mongo document | [default to null]
**PreferredIntelligence** | **string** | auto || BrightCloud || Webpulse || Gatekeeper | [default to null]
**IpThreatIntelligence** | **string** | auto || BrightCloud || Webpulse || Gatekeeper | [default to null]
**WcHealthTelemetry** | **string** | auto || enable || disable | [default to null]
**WcHeroTelemetry** | **string** | auto || enable || disable | [default to null]
**IpThreatProtect** | **string** | auto || enable || disable | [default to null]
**DpiContentFiltering** | **string** | auto || enable || disable | [default to null]
**InlineDpi** | **string** | auto || enable || disable | [default to null]
**DosProtection** | **interface{}** |  | [optional] [default to null]
**IpThreatProvider** | **string** | auto || plume || norton | [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


