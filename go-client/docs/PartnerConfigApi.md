# \PartnerConfigApi

All URIs are relative to *https://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PartnerConfigCreateAppEngagementTimerConfig**](PartnerConfigApi.md#PartnerConfigCreateAppEngagementTimerConfig) | **Post** /partnerConfig/{id}/appEngagementTimer | Create a partner config for app engagement timer
[**PartnerConfigDeleteAppPrioritizationPartnerAppPriority**](PartnerConfigApi.md#PartnerConfigDeleteAppPrioritizationPartnerAppPriority) | **Delete** /partnerConfig/{partnerId}/qos/appPrioritization/appPriority | Set app priority to default for app prioritization.
[**PartnerConfigDeleteAppPrioritizationPartnerCustomSetting**](PartnerConfigApi.md#PartnerConfigDeleteAppPrioritizationPartnerCustomSetting) | **Delete** /partnerConfig/{partnerId}/qos/appPrioritization/customSetting | Set custom setting to default for app prioritization.
[**PartnerConfigDeleteAppQoe**](PartnerConfigApi.md#PartnerConfigDeleteAppQoe) | **Delete** /partnerConfig/{id}/platform/appQoe | Delete cohort appQoe config.
[**PartnerConfigDeleteMachineToMachineToken**](PartnerConfigApi.md#PartnerConfigDeleteMachineToMachineToken) | **Delete** /partnerConfig/{id}/machineToMachine/tokens/{tokenId} | Delete one of the tokens for the PartnerId
[**PartnerConfigDeletePreCacScheduler**](PartnerConfigApi.md#PartnerConfigDeletePreCacScheduler) | **Delete** /partnerConfig/{id}/platform/pcs | Delete cohort preCacScheduler config.
[**PartnerConfigDeleteSamKnows**](PartnerConfigApi.md#PartnerConfigDeleteSamKnows) | **Delete** /partnerConfig/{id}/platform/samKnows | Delete cohort samKnows config.
[**PartnerConfigDeleteSipAlg**](PartnerConfigApi.md#PartnerConfigDeleteSipAlg) | **Delete** /partnerConfig/{id}/platform/sipAlg | Delete cohort sipAlg config.
[**PartnerConfigDeleteSpeedTest**](PartnerConfigApi.md#PartnerConfigDeleteSpeedTest) | **Delete** /partnerConfig/{id}/platform/speedTest | Delete partners speedTest config.
[**PartnerConfigDeleteVlanService**](PartnerConfigApi.md#PartnerConfigDeleteVlanService) | **Delete** /partnerConfig/{id}/platform/vlanServices | Delete cohort vlan service config.
[**PartnerConfigDisableMachineToMachine**](PartnerConfigApi.md#PartnerConfigDisableMachineToMachine) | **Delete** /partnerConfig/{id}/machineToMachine | Disable machine to machine token generation for partner
[**PartnerConfigEnableMachineToMachine**](PartnerConfigApi.md#PartnerConfigEnableMachineToMachine) | **Put** /partnerConfig/{id}/machineToMachine | Enable machine to machine token generation for partnerId
[**PartnerConfigGenerateMachineToMachineToken**](PartnerConfigApi.md#PartnerConfigGenerateMachineToMachineToken) | **Post** /partnerConfig/{id}/machineToMachine/tokens | Generate a new Machine to Machine token for PartnerId
[**PartnerConfigGetAppEngagementTimerConfig**](PartnerConfigApi.md#PartnerConfigGetAppEngagementTimerConfig) | **Get** /partnerConfig/{id}/appEngagementTimer | Get a partner config for app engagement timer
[**PartnerConfigGetAppPrioritizationPartnerConfig**](PartnerConfigApi.md#PartnerConfigGetAppPrioritizationPartnerConfig) | **Get** /partnerConfig/{partnerId}/qos/appPrioritization | Get status for app prioritization.
[**PartnerConfigGetAppQoe**](PartnerConfigApi.md#PartnerConfigGetAppQoe) | **Get** /partnerConfig/{id}/platform/appQoe | Get partners appQoe config.
[**PartnerConfigGetCaptivePortalConfig**](PartnerConfigApi.md#PartnerConfigGetCaptivePortalConfig) | **Get** /partnerConfig/{id}/captivePortal | Get partners captivePortal configs 
[**PartnerConfigGetCohorts**](PartnerConfigApi.md#PartnerConfigGetCohorts) | **Get** /partnerConfig/{id}/platform | Get all partners cohort configs.
[**PartnerConfigGetFeatureFlags**](PartnerConfigApi.md#PartnerConfigGetFeatureFlags) | **Get** /partnerConfig/{id}/featureFlags | Get partners feature flags
[**PartnerConfigGetHomepassCustomerSupportConfigurations**](PartnerConfigApi.md#PartnerConfigGetHomepassCustomerSupportConfigurations) | **Get** /partnerConfig/{id}/homepass/customerSupportConfigurations | get homepass customer support configurations
[**PartnerConfigGetMachineToMachine**](PartnerConfigApi.md#PartnerConfigGetMachineToMachine) | **Get** /partnerConfig/{id}/machineToMachine | Get partners machine to machine information
[**PartnerConfigGetPreCacScheduler**](PartnerConfigApi.md#PartnerConfigGetPreCacScheduler) | **Get** /partnerConfig/{id}/platform/pcs | Get Pre-CAC Scheduler Config
[**PartnerConfigGetSamKnows**](PartnerConfigApi.md#PartnerConfigGetSamKnows) | **Get** /partnerConfig/{id}/platform/samKnows | Get partners samKnows config.
[**PartnerConfigGetSecurityPolicy**](PartnerConfigApi.md#PartnerConfigGetSecurityPolicy) | **Get** /partnerConfig/{id}/securityPolicy | Get partner&#39;s default securityPolicy
[**PartnerConfigGetSipAlg**](PartnerConfigApi.md#PartnerConfigGetSipAlg) | **Get** /partnerConfig/{id}/platform/sipAlg | Get partners sipAlg config.
[**PartnerConfigGetSpeedTest**](PartnerConfigApi.md#PartnerConfigGetSpeedTest) | **Get** /partnerConfig/{id}/platform/speedTest | Get partners speedTest config.
[**PartnerConfigGetVlanServices**](PartnerConfigApi.md#PartnerConfigGetVlanServices) | **Get** /partnerConfig/{id}/platform/vlanServices | Get cohort vlan service configs.
[**PartnerConfigGetWorkpassCustomerSupportConfigurations**](PartnerConfigApi.md#PartnerConfigGetWorkpassCustomerSupportConfigurations) | **Get** /partnerConfig/{id}/workpass/customerSupportConfigurations | Get workpass customer support configurations
[**PartnerConfigPatchAppEngagementTimerConfig**](PartnerConfigApi.md#PartnerConfigPatchAppEngagementTimerConfig) | **Patch** /partnerConfig/{id}/appEngagementTimer | Patch a partner config for app engagement timer
[**PartnerConfigPatchAppPrioritizationPartnerConfig**](PartnerConfigApi.md#PartnerConfigPatchAppPrioritizationPartnerConfig) | **Patch** /partnerConfig/{partnerId}/qos/appPrioritization | Update app prioritization config.
[**PartnerConfigPatchHomepassCustomerSupportConfigurations**](PartnerConfigApi.md#PartnerConfigPatchHomepassCustomerSupportConfigurations) | **Patch** /partnerConfig/{id}/homepass/customerSupportConfigurations | Patch customer support configurations
[**PartnerConfigPatchSecurityPolicy**](PartnerConfigApi.md#PartnerConfigPatchSecurityPolicy) | **Patch** /partnerConfig/{id}/securityPolicy | Set partner&#39;s default securityPolicy
[**PartnerConfigPatchWorkpassCustomerSupportConfigurations**](PartnerConfigApi.md#PartnerConfigPatchWorkpassCustomerSupportConfigurations) | **Patch** /partnerConfig/{id}/workpass/customerSupportConfigurations | Patch customer support configurations
[**PartnerConfigPutVlanServices**](PartnerConfigApi.md#PartnerConfigPutVlanServices) | **Put** /partnerConfig/{id}/platform/vlanServices | Update cohort vlan service Config.
[**PartnerConfigUpdateAppQoe**](PartnerConfigApi.md#PartnerConfigUpdateAppQoe) | **Put** /partnerConfig/{id}/platform/appQoe | Update partners appQoe Config.
[**PartnerConfigUpdateCaptivePortalConfig**](PartnerConfigApi.md#PartnerConfigUpdateCaptivePortalConfig) | **Patch** /partnerConfig/{id}/captivePortal | Patch a partners captivePortal configs
[**PartnerConfigUpdateFeatureFlags**](PartnerConfigApi.md#PartnerConfigUpdateFeatureFlags) | **Patch** /partnerConfig/{id}/featureFlags | Patch a partners feature flags
[**PartnerConfigUpdatePreCacScheduler**](PartnerConfigApi.md#PartnerConfigUpdatePreCacScheduler) | **Put** /partnerConfig/{id}/platform/pcs | Update partners Pre CAC Scheduler config
[**PartnerConfigUpdateSamKnows**](PartnerConfigApi.md#PartnerConfigUpdateSamKnows) | **Put** /partnerConfig/{id}/platform/samKnows | Update partners samKnows Config.
[**PartnerConfigUpdateSipAlg**](PartnerConfigApi.md#PartnerConfigUpdateSipAlg) | **Put** /partnerConfig/{id}/platform/sipAlg | Update partners sipAlg Config.
[**PartnerConfigUpdateSpeedTest**](PartnerConfigApi.md#PartnerConfigUpdateSpeedTest) | **Put** /partnerConfig/{id}/platform/speedTest | Update partners speedTest Config.


# **PartnerConfigCreateAppEngagementTimerConfig**
> XAny PartnerConfigCreateAppEngagementTimerConfig(ctx, id, timerInSeconds, numberOfAllowedRetries, enable)
Create a partner config for app engagement timer

<div><strong>200</strong>: Success. Created partner config.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>422</strong>: \"illegal field\"</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **timerInSeconds** | **float64**|  | 
  **numberOfAllowedRetries** | **float64**|  | 
  **enable** | **bool**|  | 

### Return type

[**XAny**](x-any.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PartnerConfigDeleteAppPrioritizationPartnerAppPriority**
> interface{} PartnerConfigDeleteAppPrioritizationPartnerAppPriority(ctx, partnerId)
Set app priority to default for app prioritization.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **partnerId** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PartnerConfigDeleteAppPrioritizationPartnerCustomSetting**
> interface{} PartnerConfigDeleteAppPrioritizationPartnerCustomSetting(ctx, partnerId)
Set custom setting to default for app prioritization.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **partnerId** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PartnerConfigDeleteAppQoe**
> PartnerConfigDeleteAppQoe(ctx, id)
Delete cohort appQoe config.

<div><strong>204</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PartnerConfigDeleteMachineToMachineToken**
> XAny PartnerConfigDeleteMachineToMachineToken(ctx, id, tokenId)
Delete one of the tokens for the PartnerId

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required</div> <div><strong>404</strong>: There is not configuration for this partnerId</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **tokenId** | **string**|  | 

### Return type

[**XAny**](x-any.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PartnerConfigDeletePreCacScheduler**
> PartnerConfigDeletePreCacScheduler(ctx, id)
Delete cohort preCacScheduler config.

<div><strong>204</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PartnerConfigDeleteSamKnows**
> PartnerConfigDeleteSamKnows(ctx, id)
Delete cohort samKnows config.

<div><strong>204</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PartnerConfigDeleteSipAlg**
> PartnerConfigDeleteSipAlg(ctx, id)
Delete cohort sipAlg config.

<div><strong>204</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PartnerConfigDeleteSpeedTest**
> PartnerConfigDeleteSpeedTest(ctx, id)
Delete partners speedTest config.

<div><strong>204</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PartnerConfigDeleteVlanService**
> PartnerConfigDeleteVlanService(ctx, id)
Delete cohort vlan service config.

<div><strong>204</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PartnerConfigDisableMachineToMachine**
> XAny PartnerConfigDisableMachineToMachine(ctx, id)
Disable machine to machine token generation for partner

<div><strong>204</strong>: Success.</div> <div><strong>401</strong>: Authorization required</div> <div><strong>403</strong>: Not allowed to update configuration</div> <div><strong>404</strong>: Group of partnerId not found</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**XAny**](x-any.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PartnerConfigEnableMachineToMachine**
> XAny PartnerConfigEnableMachineToMachine(ctx, id)
Enable machine to machine token generation for partnerId

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required</div> <div><strong>403</strong>: Not allowed to update configuration</div> <div><strong>404</strong>: Group of partnerId not found</div> <div><strong>422</strong>: Illegal field</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**XAny**](x-any.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PartnerConfigGenerateMachineToMachineToken**
> XAny PartnerConfigGenerateMachineToMachineToken(ctx, id, tokenName, optional)
Generate a new Machine to Machine token for PartnerId

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required</div> <div><strong>404</strong>: There is not configuration for this partnerId</div> <div><strong>422</strong>: Illegal field</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **tokenName** | **string**|  | 
 **optional** | ***PartnerConfigApiPartnerConfigGenerateMachineToMachineTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PartnerConfigApiPartnerConfigGenerateMachineToMachineTokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tokenTTLSeconds** | **optional.Float64**|  | 

### Return type

[**XAny**](x-any.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PartnerConfigGetAppEngagementTimerConfig**
> XAny PartnerConfigGetAppEngagementTimerConfig(ctx, id)
Get a partner config for app engagement timer

<div><strong>200</strong>: Success. Retrieved partner config.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Config does not exist </div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**XAny**](x-any.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PartnerConfigGetAppPrioritizationPartnerConfig**
> interface{} PartnerConfigGetAppPrioritizationPartnerConfig(ctx, partnerId)
Get status for app prioritization.

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Partner id or WifiNetwork does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **partnerId** | **string**| partner Id | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PartnerConfigGetAppQoe**
> interface{} PartnerConfigGetAppQoe(ctx, id)
Get partners appQoe config.

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PartnerConfigGetCaptivePortalConfig**
> XAny PartnerConfigGetCaptivePortalConfig(ctx, id)
Get partners captivePortal configs 

<div><strong>200</strong>: Success. Patched partner config.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**XAny**](x-any.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PartnerConfigGetCohorts**
> interface{} PartnerConfigGetCohorts(ctx, id)
Get all partners cohort configs.

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required</div> <div><strong>404</strong>: PartnerId not found or no configurations exist for partner</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PartnerConfigGetFeatureFlags**
> XAny PartnerConfigGetFeatureFlags(ctx, id)
Get partners feature flags

<div><strong>200</strong>: Success. Patched partner config.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>403</strong>: Not allowed to update configuration</div> <div><strong>422</strong>: Illegal field</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**XAny**](x-any.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PartnerConfigGetHomepassCustomerSupportConfigurations**
> HomepassCustomerSupportConfigurations PartnerConfigGetHomepassCustomerSupportConfigurations(ctx, id)
get homepass customer support configurations

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required.</div> <div><strong>422</strong>: Input validation failed.</div> <div><strong>500</strong>: Internal server error.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**HomepassCustomerSupportConfigurations**](homepassCustomerSupportConfigurations.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PartnerConfigGetMachineToMachine**
> XAny PartnerConfigGetMachineToMachine(ctx, id)
Get partners machine to machine information

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required</div> <div><strong>404</strong>: There is not configuration for this partnerId</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**XAny**](x-any.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PartnerConfigGetPreCacScheduler**
> interface{} PartnerConfigGetPreCacScheduler(ctx, id)
Get Pre-CAC Scheduler Config

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PartnerConfigGetSamKnows**
> interface{} PartnerConfigGetSamKnows(ctx, id)
Get partners samKnows config.

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PartnerConfigGetSecurityPolicy**
> XAny PartnerConfigGetSecurityPolicy(ctx, id)
Get partner's default securityPolicy

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required</div> <div><strong>404</strong>: There is not configuration for this partnerId</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**XAny**](x-any.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PartnerConfigGetSipAlg**
> interface{} PartnerConfigGetSipAlg(ctx, id)
Get partners sipAlg config.

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PartnerConfigGetSpeedTest**
> interface{} PartnerConfigGetSpeedTest(ctx, id)
Get partners speedTest config.

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PartnerConfigGetVlanServices**
> interface{} PartnerConfigGetVlanServices(ctx, id)
Get cohort vlan service configs.

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required</div> <div><strong>404</strong>: PartnerId not found or no configurations exist for partner</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PartnerConfigGetWorkpassCustomerSupportConfigurations**
> WorkpassCustomerSupportConfigurations PartnerConfigGetWorkpassCustomerSupportConfigurations(ctx, id)
Get workpass customer support configurations

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required.</div> <div><strong>500</strong>: Internal server error.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**WorkpassCustomerSupportConfigurations**](workpassCustomerSupportConfigurations.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PartnerConfigPatchAppEngagementTimerConfig**
> XAny PartnerConfigPatchAppEngagementTimerConfig(ctx, id, optional)
Patch a partner config for app engagement timer

<div><strong>200</strong>: Success. Patched partner config.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>422</strong>: \"illegal field\"</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***PartnerConfigApiPartnerConfigPatchAppEngagementTimerConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PartnerConfigApiPartnerConfigPatchAppEngagementTimerConfigOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **timerInSeconds** | **optional.Float64**|  | 
 **numberOfAllowedRetries** | **optional.Float64**|  | 
 **enable** | **optional.Bool**|  | 

### Return type

[**XAny**](x-any.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PartnerConfigPatchAppPrioritizationPartnerConfig**
> XAny PartnerConfigPatchAppPrioritizationPartnerConfig(ctx, partnerId, optional)
Update app prioritization config.

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Partner id or WifiNetwork does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **partnerId** | **string**|  | 
 **optional** | ***PartnerConfigApiPartnerConfigPatchAppPrioritizationPartnerConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PartnerConfigApiPartnerConfigPatchAppPrioritizationPartnerConfigOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enabled** | **optional.Bool**| (deprecated) true if app prioritization is enabled | 
 **defaultMode** | **optional.String**| App Prioritization mode - any of enable | disable | 
 **initialLocationEnabled** | **optional.Bool**| true if app prioritization is enabled for all new locations | 
 **customSettingEnabled** | **optional.Bool**| true if custom setting is enabled | 
 **template** | **optional.String**| Template for app prioritization | 
 **appPriority** | **optional.String**| priority for apps | 
 **customSetting** | **optional.String**| Settings for app prioritization | 

### Return type

[**XAny**](x-any.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PartnerConfigPatchHomepassCustomerSupportConfigurations**
> XAny PartnerConfigPatchHomepassCustomerSupportConfigurations(ctx, id)
Patch customer support configurations

<div><strong>200</strong>: Success. Patched customer support configurations.</div> <div><strong>401</strong>: Authorization required.</div> <div><strong>422</strong>: Input validation failed.</div> <div><strong>500</strong>: Internal server error.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**XAny**](x-any.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PartnerConfigPatchSecurityPolicy**
> XAny PartnerConfigPatchSecurityPolicy(ctx, id, optional)
Set partner's default securityPolicy

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required</div> <div><strong>404</strong>: There is not configuration for this partnerId</div> <div><strong>422</strong>: Illegal field</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***PartnerConfigApiPartnerConfigPatchSecurityPolicyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PartnerConfigApiPartnerConfigPatchSecurityPolicyOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **whitelist** | **optional.String**| List of whitelisted fqdn values - {type&#x3D;\&quot;fqdn\&quot;, value}\&quot; | 
 **blacklist** | **optional.String**| List of blacklisted fqdn values - {type&#x3D;\&quot;fqdn\&quot;, value} | 
 **blockedDnsRedirect** | **optional.String**| Redirect \&quot;address\&quot; and corresponding \&quot;ttl\&quot; redirect in seconds | 
 **contentFilterOverrides** | **optional.String**| Array of content filter overrides each with its own whitelist/blacklist | 
 **remoteConnectionsMode** | **optional.String**| Any of \&quot;auto\&quot;, \&quot;enabled\&quot;, \&quot;disabled\&quot;, \&quot;highRiskOnly\&quot; | 
 **remoteConnectionsForceDisable** | **optional.Bool**|  | 
 **disableSafeSearch** | **optional.Bool**|  | 

### Return type

[**XAny**](x-any.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PartnerConfigPatchWorkpassCustomerSupportConfigurations**
> XAny PartnerConfigPatchWorkpassCustomerSupportConfigurations(ctx, id)
Patch customer support configurations

<div><strong>200</strong>: Success. Patched customer support configurations.</div> <div><strong>401</strong>: Authorization required.</div> <div><strong>422</strong>: Input validation failed.</div> <div><strong>500</strong>: Internal server error.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**XAny**](x-any.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PartnerConfigPutVlanServices**
> interface{} PartnerConfigPutVlanServices(ctx, id, config)
Update cohort vlan service Config.

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>422</strong>: Invalid schema.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **config** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PartnerConfigUpdateAppQoe**
> interface{} PartnerConfigUpdateAppQoe(ctx, id, config)
Update partners appQoe Config.

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>422</strong>: Invalid schema.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **config** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PartnerConfigUpdateCaptivePortalConfig**
> XAny PartnerConfigUpdateCaptivePortalConfig(ctx, id, optional)
Patch a partners captivePortal configs

<div><strong>200</strong>: Success. Patched partner config.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>422</strong>: Incorrect language type</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***PartnerConfigApiPartnerConfigUpdateCaptivePortalConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PartnerConfigApiPartnerConfigUpdateCaptivePortalConfigOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **defaultLanguage** | **optional.String**|  | 

### Return type

[**XAny**](x-any.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PartnerConfigUpdateFeatureFlags**
> XAny PartnerConfigUpdateFeatureFlags(ctx, id)
Patch a partners feature flags

<div><strong>200</strong>: Success. Patched partner config.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>403</strong>: Not allowed to update configuration</div> <div><strong>422</strong>: Illegal field</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**XAny**](x-any.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PartnerConfigUpdatePreCacScheduler**
> interface{} PartnerConfigUpdatePreCacScheduler(ctx, id, config)
Update partners Pre CAC Scheduler config

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>422</strong>: Invalid schema.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **config** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PartnerConfigUpdateSamKnows**
> interface{} PartnerConfigUpdateSamKnows(ctx, id, config)
Update partners samKnows Config.

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>422</strong>: Invalid schema.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **config** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PartnerConfigUpdateSipAlg**
> interface{} PartnerConfigUpdateSipAlg(ctx, id, config)
Update partners sipAlg Config.

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>422</strong>: Invalid schema.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **config** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PartnerConfigUpdateSpeedTest**
> interface{} PartnerConfigUpdateSpeedTest(ctx, id, config)
Update partners speedTest Config.

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>422</strong>: Invalid schema.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **config** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

