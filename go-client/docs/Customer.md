# Customer

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | [***ObjectId**](ObjectID.md) |  | [optional] [default to null]
**AccountId** | **string** |  | [optional] [default to null]
**Anonymous** | **bool** | ISP registered accounts are anonymous and do not include real emails and passwords | [optional] [default to null]
**AutoProvisioned** | **bool** | ISP auto registered accounts created using gatewayId instead of accountId need to eventually be updated a real accountId | [optional] [default to null]
**Contact** | **string** |  | [optional] [default to null]
**Name** | **string** |  | [default to null]
**Locked** | **bool** | The groupAdmin account needs this property to protect it from being changed by other groupAdmins unintentionally. | [optional] [default to null]
**PartnerId** | **string** |  | [optional] [default to null]
**AcceptLanguage** | **string** | The acceptable language for this user | [optional] [default to null]
**PreferredLanguage** | **string** | The language set for this user based on API and/or it&#39;s language header | [optional] [default to null]
**NocSettings** | **interface{}** |  | [optional] [default to null]
**LinkedAccounts** | **[]interface{}** | array of the outside/non-plume accounts to be linked | [optional] [default to null]
**Source** | **string** |  | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**ProvisioningSsoAuditTrail** | **bool** | If true, the user has provisioning SSO audit trail | [optional] [default to null]
**Version** | **string** | schema version # of a Mongo document | [default to null]
**Realm** | **string** |  | [optional] [default to null]
**Username** | **string** |  | [optional] [default to null]
**Email** | **string** |  | [default to null]
**EmailVerified** | **bool** |  | [optional] [default to null]
**TermsAndPrivacy** | [***TermsAndPrivacy**](TermsAndPrivacy.md) |  | [optional] [default to null]
**Migration** | [***Migration**](Migration.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


