# \NodeApi

All URIs are relative to *https://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NodeCount**](NodeApi.md#NodeCount) | **Get** /Nodes/count | Count instances of the model matched by where from the data source.
[**NodeCustomCreate**](NodeApi.md#NodeCustomCreate) | **Post** /Nodes | Import a node into the global/shared inventory (does NOT claim).
[**NodeFind**](NodeApi.md#NodeFind) | **Get** /Nodes | Find all instances of the model matched by filter from the data source.
[**NodeFindById**](NodeApi.md#NodeFindById) | **Get** /Nodes/{id} | Find a model instance by {{id}} from the data source.
[**NodePrototypeGetCustomerByNodeId**](NodeApi.md#NodePrototypeGetCustomerByNodeId) | **Get** /Nodes/{id}/customer | Get the customer info with the node Id.
[**NodePrototypeGetMqttBroker**](NodeApi.md#NodePrototypeGetMqttBroker) | **Get** /Nodes/{id}/mqtt | Get the MQTT broker address of the node.
[**NodePrototypeUnclaim**](NodeApi.md#NodePrototypeUnclaim) | **Delete** /Nodes/{id} | Unclaim a node by a groupAdmin or admin
[**NodePrototypeVerifyEmailPasswordlessToken**](NodeApi.md#NodePrototypeVerifyEmailPasswordlessToken) | **Post** /Nodes/{id}/passwordLessToken | Update the name and email for customer and generate emailToken and appToken.
[**NodeUpdatePackId**](NodeApi.md#NodeUpdatePackId) | **Put** /Nodes/{id}/packId | Rename an unclaimed node/pod&#39;s packId in Plume&#39;s global inventory.


# **NodeCount**
> InlineResponse2001 NodeCount(ctx, optional)
Count instances of the model matched by where from the data source.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NodeApiNodeCountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodeApiNodeCountOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **where** | **optional.String**| Criteria to match model instances | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NodeCustomCreate**
> Node NodeCustomCreate(ctx, id, residentialGateway, optional)
Import a node into the global/shared inventory (does NOT claim).

<div><strong>200</strong>: Success, node imported.</div> <div><strong>401</strong>: Authorization required.</div> <div><strong>409</strong>: NodeId already exists in shared inventory.</div> <div><strong>422</strong>: Input validation failed.</div> <div><strong>500</strong>: Internal server error.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **residentialGateway** | **bool**|  | 
 **optional** | ***NodeApiNodeCustomCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodeApiNodeCustomCreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **model** | **optional.String**| Node model ID value is required unless a Partner ID exemption has been configured | 
 **packId** | **optional.String**| optional packId to group nodes | 
 **partnerId** | **optional.String**| Partner ID required on Plume production clouds | 
 **radioMac24** | **optional.String**| optional radioMac24, must be a valid mac address | 
 **radioMac50** | **optional.String**| optional radioMac50, must be a valid mac address | 
 **radioMac60** | **optional.String**| optional radioMac60, must be a valid mac address | 
 **ethernetMac** | **optional.String**| optional ethernetMac, must be a valid mac address | 
 **ethernet1Mac** | **optional.String**| optional ethernet1Mac, must be a valid mac address | 
 **claimKeyRequired** | **optional.Bool**| optional claimKeyRequired, default is false | 
 **radioMac50L** | **optional.String**| optional radioMac50L, must be a valid mac address | 
 **radioMac50U** | **optional.String**| optional radioMac50U, must be a valid mac address | 
 **subscriptionRequired** | **optional.Bool**| optional subscriptionRequired, default is false | 
 **thread** | **optional.String**| optional Thread/Matter MAC addres | 

### Return type

[**Node**](Node.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NodeFind**
> []Node NodeFind(ctx, optional)
Find all instances of the model matched by filter from the data source.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NodeApiNodeFindOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodeApiNodeFindOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **optional.String**| Filter defining fields, where, include, order, offset, and limit - must be a JSON-encoded string (&#x60;{\&quot;where\&quot;:{\&quot;something\&quot;:\&quot;value\&quot;}}&#x60;).  See https://loopback.io/doc/en/lb3/Querying-data.html#using-stringified-json-in-rest-queries for more details. | 

### Return type

[**[]Node**](Node.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NodeFindById**
> Node NodeFindById(ctx, id, optional)
Find a model instance by {{id}} from the data source.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Model id | 
 **optional** | ***NodeApiNodeFindByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodeApiNodeFindByIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **optional.String**| Filter defining fields and include - must be a JSON-encoded string ({\&quot;something\&quot;:\&quot;value\&quot;}) | 

### Return type

[**Node**](Node.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NodePrototypeGetCustomerByNodeId**
> Customer NodePrototypeGetCustomerByNodeId(ctx, id)
Get the customer info with the node Id.

<div><strong>200</strong>: Success, return the customer info.</div> <div><strong>403</strong>: Public ip not matched.</div> <div><strong>404</strong>: NodeId not found.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Node id | 

### Return type

[**Customer**](Customer.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NodePrototypeGetMqttBroker**
> interface{} NodePrototypeGetMqttBroker(ctx, id)
Get the MQTT broker address of the node.

<div><strong>200</strong>: Success, return the customer info.</div> <div><strong>404</strong>: NodeId not found.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Node id | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NodePrototypeUnclaim**
> NodePrototypeUnclaim(ctx, id, id2, optional)
Unclaim a node by a groupAdmin or admin

<div><strong>204</strong>: Success, node changed.</div> <div><strong>401</strong>: Authorization required.</div> <div><strong>500</strong>: Internal server error.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Node id | 
  **id2** | **string**|  | 
 **optional** | ***NodeApiNodePrototypeUnclaimOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodeApiNodePrototypeUnclaimOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **forceUnclaim** | **optional.Bool**|  | [default to false]
 **preservePackId** | **optional.Bool**|  | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NodePrototypeVerifyEmailPasswordlessToken**
> Customer NodePrototypeVerifyEmailPasswordlessToken(ctx, id, name, email)
Update the name and email for customer and generate emailToken and appToken.

<div><strong>200</strong>: Success, return the customer info.</div> <div><strong>403</strong>: Public ip not matched.</div> <div><strong>404</strong>: NodeId not found.</div> <div><strong>422</strong>: Email must be defined and valid.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Node id | 
  **name** | **string**|  | 
  **email** | **string**|  | 

### Return type

[**Customer**](Customer.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NodeUpdatePackId**
> Node NodeUpdatePackId(ctx, id, packId)
Rename an unclaimed node/pod's packId in Plume's global inventory.

<div><strong>200</strong>: Success, a job well done.</div> <div><strong>400</strong>: Bad request, packId is undefined or empty string.</div> <div><strong>404</strong>: NodeId not found.</div> <div><strong>422</strong>: PackId is invalid (too long).</div> <div><strong>423</strong>: PackId cannot be changed for a claimed pod.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **packId** | **string**|  | 

### Return type

[**Node**](Node.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

