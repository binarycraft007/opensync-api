# \AuditTrailApi

All URIs are relative to *https://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuditTrailGetAuditTrail**](AuditTrailApi.md#AuditTrailGetAuditTrail) | **Get** /AuditTrails/getAuditTrail | Get Audit Trail for a customer and/or location


# **AuditTrailGetAuditTrail**
> []AuditTrail AuditTrailGetAuditTrail(ctx, customerId, optional)
Get Audit Trail for a customer and/or location

<div><strong>200</strong>: Ok.</div> <div><strong>401</strong>: Authorization required.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>500</strong>: Internal server error.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **string**| Customer Id | 
 **optional** | ***AuditTrailApiAuditTrailGetAuditTrailOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuditTrailApiAuditTrailGetAuditTrailOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **locationId** | **optional.String**| Location Id | 
 **partnerIds** | **optional.String**| Partner Id | 

### Return type

[**[]AuditTrail**](AuditTrail.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

