# \InventoryApi

All URIs are relative to *https://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InventoryAutoProvisioning**](InventoryApi.md#InventoryAutoProvisioning) | **Patch** /inventory/nodes/{nodeId}/autoProvisioning | S node by ID from the Inventory service
[**InventoryGetNodeById**](InventoryApi.md#InventoryGetNodeById) | **Get** /inventory/nodes/{nodeId} | Get node by ID from the Inventory service


# **InventoryAutoProvisioning**
> interface{} InventoryAutoProvisioning(ctx, nodeId, autoProvisionToThisDeployment)
S node by ID from the Inventory service

<div><strong>200</strong>: Success, return the node object</div> <div><strong>401</strong>: Authorization required</div> <div><strong>404</strong>: Node not found.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeId** | **string**|  | 
  **autoProvisionToThisDeployment** | **bool**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InventoryGetNodeById**
> interface{} InventoryGetNodeById(ctx, nodeId)
Get node by ID from the Inventory service

<div><strong>200</strong>: Success, return the node object</div> <div><strong>401</strong>: Authorization required</div> <div><strong>404</strong>: Node not found.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeId** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

