# \CommandApi

All URIs are relative to *https://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CommandDeleteCommand**](CommandApi.md#CommandDeleteCommand) | **Delete** /command/{providerUserId} | Delete account linking
[**CommandGetCommand**](CommandApi.md#CommandGetCommand) | **Get** /command/{providerUserId} | Get account linking
[**CommandGetHealth**](CommandApi.md#CommandGetHealth) | **Get** /command/health | return a 200 response when your server is healthy, else send a 500 response


# **CommandDeleteCommand**
> XAny CommandDeleteCommand(ctx, providerUserId)
Delete account linking

<div><strong>204</strong>: Success.</div> <div><strong>400</strong>: Missing providerUserId body parameter</div> <div><strong>401</strong>: Authorization required.</div> <div><strong>404</strong>: Command not found.</div> <div><strong>500</strong>: internal server error.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **providerUserId** | **string**| id of the user in provider&#39;s system | 

### Return type

[**XAny**](x-any.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CommandGetCommand**
> interface{} CommandGetCommand(ctx, providerUserId)
Get account linking

<div><strong>200</strong>: Success.</div> <div><strong>400</strong>: Missing providerUserId body parameter</div> <div><strong>401</strong>: Authorization required.</div> <div><strong>404</strong>: Command not found.</div> <div><strong>500</strong>: internal server error.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **providerUserId** | **string**| id of the user in provider&#39;s system | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CommandGetHealth**
> XAny CommandGetHealth(ctx, )
return a 200 response when your server is healthy, else send a 500 response

<div><strong>200</strong>: Success, return health data.</div> <div><strong>401</strong>: Authorization required.</div> <div><strong>500</strong>: internal server error.</div>

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**XAny**](x-any.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

