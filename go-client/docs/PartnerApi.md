# \PartnerApi

All URIs are relative to *https://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PartnerFindCustomers**](PartnerApi.md#PartnerFindCustomers) | **Get** /partners/customers/search/{keyword} | Queries Customers with caller&#39;s partnerId.
[**PartnerFindCustomersForIntegrationUser**](PartnerApi.md#PartnerFindCustomersForIntegrationUser) | **Get** /partners/{partnerId}/customers/search/{keyword} | Queries Customers with caller&#39;s partnerId.
[**PartnerGetCustomerCount**](PartnerApi.md#PartnerGetCustomerCount) | **Get** /partners/{id}/customers/count | Queries Customers/locations/count with caller&#39;s groups.
[**PartnerGetLocations**](PartnerApi.md#PartnerGetLocations) | **Get** /partners/locations/{keyword} | Queries Locations with serviceId or locationId within the caller&#39;s partnerId.
[**PartnerGetNodesById**](PartnerApi.md#PartnerGetNodesById) | **Get** /partners/nodes/{id} | Queries Customers/locations/nodes with caller&#39;s partnerId.
[**PartnerGetNodesByIdForIntegrationUser**](PartnerApi.md#PartnerGetNodesByIdForIntegrationUser) | **Get** /partners/{partnerId}/nodes/{nodeId} | Queries Customers/locations/nodes with caller&#39;s partnerId.
[**PartnerGetRecentCustomers**](PartnerApi.md#PartnerGetRecentCustomers) | **Get** /partners/{id}/customers/recent | Queries Customers/locations/count with caller&#39;s partnerId.
[**PartnerPatchNodesById**](PartnerApi.md#PartnerPatchNodesById) | **Patch** /partners/nodes/{id} | Queries Customers/locations/nodes with caller&#39;s partnerId, and update it.


# **PartnerFindCustomers**
> interface{} PartnerFindCustomers(ctx, keyword, field, optional)
Queries Customers with caller's partnerId.

<div><strong>200</strong>: Success, full object returned.</div> <div><strong>401</strong>: Authorization required.</div> <div><strong>404</strong>: LocationId not found.</div> <div><strong>500</strong>: Internal server error.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **keyword** | **string**|  | 
  **field** | **string**|  | 
 **optional** | ***PartnerApiPartnerFindCustomersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PartnerApiPartnerFindCustomersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **exactMatch** | **optional.Bool**|  | 
 **startsWith** | **optional.Bool**|  | 
 **limit** | **optional.Float64**|  | [default to 10]
 **skip** | **optional.Float64**|  | [default to 0]

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PartnerFindCustomersForIntegrationUser**
> interface{} PartnerFindCustomersForIntegrationUser(ctx, partnerId, keyword, field, optional)
Queries Customers with caller's partnerId.

<div><strong>200</strong>: Success, full object returned.</div> <div><strong>401</strong>: Authorization required.</div> <div><strong>403</strong>: Not allowed to access partner.</div> <div><strong>404</strong>: partnerId or nodeId not found.</div> <div><strong>500</strong>: Internal server error.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **partnerId** | **string**| partner Id | 
  **keyword** | **string**|  | 
  **field** | **string**|  | 
 **optional** | ***PartnerApiPartnerFindCustomersForIntegrationUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PartnerApiPartnerFindCustomersForIntegrationUserOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **exactMatch** | **optional.Bool**|  | 
 **startsWith** | **optional.Bool**|  | 
 **limit** | **optional.Float64**|  | [default to 10]
 **skip** | **optional.Float64**|  | [default to 0]

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PartnerGetCustomerCount**
> interface{} PartnerGetCustomerCount(ctx, id)
Queries Customers/locations/count with caller's groups.

<div><strong>200</strong>: Success, full object returned.</div> <div><strong>401</strong>: Authorization required.</div> <div><strong>404</strong>: Group id not found.</div> <div><strong>500</strong>: Internal server error.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| partner Id | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PartnerGetLocations**
> interface{} PartnerGetLocations(ctx, keyword)
Queries Locations with serviceId or locationId within the caller's partnerId.

<div><strong>200</strong>: Success, full object returned.</div> <div><strong>401</strong>: Authorization required.</div> <div><strong>404</strong>: LocationId not found.</div> <div><strong>500</strong>: Internal server error.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **keyword** | **string**| could be locationId, or serviceId. | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PartnerGetNodesById**
> interface{} PartnerGetNodesById(ctx, id, optional)
Queries Customers/locations/nodes with caller's partnerId.

<div><strong>200</strong>: Success, full object returned.</div> <div><strong>401</strong>: Authorization required.</div> <div><strong>403</strong>: No right to access the node.</div> <div><strong>404</strong>: LocationId not found.</div> <div><strong>500</strong>: Internal server error.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| node Id | 
 **optional** | ***PartnerApiPartnerGetNodesByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PartnerApiPartnerGetNodesByIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **excludeUnclaimed** | **optional.Bool**| whether to filter out unclaimed nodes | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PartnerGetNodesByIdForIntegrationUser**
> interface{} PartnerGetNodesByIdForIntegrationUser(ctx, partnerId, nodeId, optional)
Queries Customers/locations/nodes with caller's partnerId.

<div><strong>200</strong>: Success, full object returned.</div> <div><strong>401</strong>: Authorization required.</div> <div><strong>403</strong>: No right to access the node.</div> <div><strong>404</strong>: partnerId or nodeId not found.</div> <div><strong>500</strong>: Internal server error.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **partnerId** | **string**| partner Id | 
  **nodeId** | **string**| node Id | 
 **optional** | ***PartnerApiPartnerGetNodesByIdForIntegrationUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PartnerApiPartnerGetNodesByIdForIntegrationUserOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **excludeUnclaimed** | **optional.Bool**| whether to filter out unclaimed nodes | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PartnerGetRecentCustomers**
> interface{} PartnerGetRecentCustomers(ctx, id)
Queries Customers/locations/count with caller's partnerId.

<div><strong>200</strong>: Success, full object returned.</div> <div><strong>401</strong>: Authorization required.</div> <div><strong>404</strong>: Group id not found.</div> <div><strong>500</strong>: Internal server error.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| group Id | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PartnerPatchNodesById**
> interface{} PartnerPatchNodesById(ctx, id, optional)
Queries Customers/locations/nodes with caller's partnerId, and update it.

<div><strong>200</strong>: Success, full object returned.</div> <div><strong>401</strong>: Authorization required.</div> <div><strong>404</strong>: LocationId not found.</div> <div><strong>500</strong>: Internal server error.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| node Id | 
 **optional** | ***PartnerApiPartnerPatchNodesByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PartnerApiPartnerPatchNodesByIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountId** | **optional.String**| accountId | 
 **unclaimable** | **optional.String**| unclaimable | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

