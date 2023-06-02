# \GroupApi

All URIs are relative to *https://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupCount**](GroupApi.md#GroupCount) | **Get** /Groups/count | Count instances of the model matched by where from the data source.
[**GroupCustomCreate**](GroupApi.md#GroupCustomCreate) | **Post** /Groups | Create a group.
[**GroupFind**](GroupApi.md#GroupFind) | **Get** /Groups | Find all instances of the model matched by filter from the data source.
[**GroupFindById**](GroupApi.md#GroupFindById) | **Get** /Groups/{id} | Find a model instance by {{id}} from the data source.
[**GroupFindCustomers**](GroupApi.md#GroupFindCustomers) | **Get** /Groups/customers/search/{keyword} | Queries Customers with caller&#39;s groups.
[**GroupGetCustomerCount**](GroupApi.md#GroupGetCustomerCount) | **Get** /Groups/{id}/customers/count | Queries Customers/locations/count with caller&#39;s groups.
[**GroupGetCustomers**](GroupApi.md#GroupGetCustomers) | **Get** /Groups/customers/{keyword} | Queries Customers with caller&#39;s groups.
[**GroupGetLocations**](GroupApi.md#GroupGetLocations) | **Get** /Groups/locations/{keyword} | Queries Locations with serviceId or locationId within the caller&#39;s groups.
[**GroupGetNodesById**](GroupApi.md#GroupGetNodesById) | **Get** /Groups/nodes/{id} | Queries Customers/locations/nodes with caller&#39;s groups.
[**GroupGetRecentCustomers**](GroupApi.md#GroupGetRecentCustomers) | **Get** /Groups/{id}/customers/recent | Queries Customers/locations/count with caller&#39;s groups.
[**GroupPatchNodesById**](GroupApi.md#GroupPatchNodesById) | **Patch** /Groups/nodes/{id} | Queries Customers/locations/nodes with caller&#39;s groups, and update it.
[**GroupPrototypeDelete**](GroupApi.md#GroupPrototypeDelete) | **Delete** /Groups/{id} | Delete a group.
[**GroupPrototypeGetCustomers**](GroupApi.md#GroupPrototypeGetCustomers) | **Get** /Groups/{id}/customers | Queries customers of Group.
[**GroupPrototypeLinkCustomers**](GroupApi.md#GroupPrototypeLinkCustomers) | **Put** /Groups/{id}/customers/rel/{fk} | Add a related item by id for customers.
[**GroupPrototypePatchAttributesPatchGroupsid**](GroupApi.md#GroupPrototypePatchAttributesPatchGroupsid) | **Patch** /Groups/{id} | Patch attributes for a model instance and persist it into the data source.
[**GroupPrototypePatchAttributesPutGroupsid**](GroupApi.md#GroupPrototypePatchAttributesPutGroupsid) | **Put** /Groups/{id} | Patch attributes for a model instance and persist it into the data source.
[**GroupPrototypeUnlinkCustomers**](GroupApi.md#GroupPrototypeUnlinkCustomers) | **Delete** /Groups/{id}/customers/rel/{fk} | Remove the customers relation to an item by id.


# **GroupCount**
> InlineResponse2001 GroupCount(ctx, optional)
Count instances of the model matched by where from the data source.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GroupApiGroupCountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupApiGroupCountOpts struct

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

# **GroupCustomCreate**
> Group GroupCustomCreate(ctx, name, description, optional)
Create a group.

<div><strong>200</strong>: Success, group created.</div> <div><strong>422</strong>: Input validation failed.</div> <div><strong>500</strong>: Internal server error.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 
  **description** | **string**|  | 
 **optional** | ***GroupApiGroupCustomCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupApiGroupCustomCreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **id** | **optional.String**| an ObjectID id for porting group ids | 
 **importGroupIdAsPartnerIdIntoInventory** | **optional.Bool**|  | 

### Return type

[**Group**](Group.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupFind**
> []Group GroupFind(ctx, optional)
Find all instances of the model matched by filter from the data source.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GroupApiGroupFindOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupApiGroupFindOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **optional.String**| Filter defining fields, where, include, order, offset, and limit - must be a JSON-encoded string (&#x60;{\&quot;where\&quot;:{\&quot;something\&quot;:\&quot;value\&quot;}}&#x60;).  See https://loopback.io/doc/en/lb3/Querying-data.html#using-stringified-json-in-rest-queries for more details. | 

### Return type

[**[]Group**](Group.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupFindById**
> Group GroupFindById(ctx, id, optional)
Find a model instance by {{id}} from the data source.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Model id | 
 **optional** | ***GroupApiGroupFindByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupApiGroupFindByIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **optional.String**| Filter defining fields and include - must be a JSON-encoded string ({\&quot;something\&quot;:\&quot;value\&quot;}) | 

### Return type

[**Group**](Group.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupFindCustomers**
> interface{} GroupFindCustomers(ctx, keyword, field, optional)
Queries Customers with caller's groups.

<div><strong>200</strong>: Success, full object returned.</div> <div><strong>401</strong>: Authorization required.</div> <div><strong>404</strong>: LocationId not found.</div> <div><strong>500</strong>: Internal server error.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **keyword** | **string**|  | 
  **field** | **string**|  | 
 **optional** | ***GroupApiGroupFindCustomersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupApiGroupFindCustomersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **exactMatch** | **optional.Bool**|  | 
 **startsWith** | **optional.Bool**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupGetCustomerCount**
> interface{} GroupGetCustomerCount(ctx, id)
Queries Customers/locations/count with caller's groups.

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

# **GroupGetCustomers**
> interface{} GroupGetCustomers(ctx, keyword)
Queries Customers with caller's groups.

<div><strong>200</strong>: Success, full object returned.</div> <div><strong>401</strong>: Authorization required.</div> <div><strong>404</strong>: LocationId not found.</div> <div><strong>500</strong>: Internal server error.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **keyword** | **string**| could be name, email, or accountId, even partial | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupGetLocations**
> interface{} GroupGetLocations(ctx, keyword)
Queries Locations with serviceId or locationId within the caller's groups.

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

# **GroupGetNodesById**
> interface{} GroupGetNodesById(ctx, id, optional)
Queries Customers/locations/nodes with caller's groups.

<div><strong>200</strong>: Success, full object returned.</div> <div><strong>401</strong>: Authorization required.</div> <div><strong>403</strong>: No right to access the node.</div> <div><strong>404</strong>: LocationId not found.</div> <div><strong>500</strong>: Internal server error.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| node Id | 
 **optional** | ***GroupApiGroupGetNodesByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupApiGroupGetNodesByIdOpts struct

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

# **GroupGetRecentCustomers**
> interface{} GroupGetRecentCustomers(ctx, id)
Queries Customers/locations/count with caller's groups.

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

# **GroupPatchNodesById**
> interface{} GroupPatchNodesById(ctx, id, optional)
Queries Customers/locations/nodes with caller's groups, and update it.

<div><strong>200</strong>: Success, full object returned.</div> <div><strong>401</strong>: Authorization required.</div> <div><strong>404</strong>: LocationId not found.</div> <div><strong>500</strong>: Internal server error.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| node Id | 
 **optional** | ***GroupApiGroupPatchNodesByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupApiGroupPatchNodesByIdOpts struct

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

# **GroupPrototypeDelete**
> InlineResponse200 GroupPrototypeDelete(ctx, id)
Delete a group.

<div><strong>200</strong>: Returns count of deleted groups, if any were deleted <div><strong>423</strong>: Locked, group cannot be deleted because it is used as a partnerId in Inventory

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Group id | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupPrototypeGetCustomers**
> []Customer GroupPrototypeGetCustomers(ctx, id, optional)
Queries customers of Group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Group id | 
 **optional** | ***GroupApiGroupPrototypeGetCustomersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupApiGroupPrototypeGetCustomersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **optional.String**|  | 

### Return type

[**[]Customer**](Customer.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupPrototypeLinkCustomers**
> CustomerGroup GroupPrototypeLinkCustomers(ctx, id, fk, optional)
Add a related item by id for customers.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Group id | 
  **fk** | **string**| Foreign key for customers | 
 **optional** | ***GroupApiGroupPrototypeLinkCustomersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupApiGroupPrototypeLinkCustomersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**optional.Interface of CustomerGroup**](CustomerGroup.md)|  | 

### Return type

[**CustomerGroup**](CustomerGroup.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupPrototypePatchAttributesPatchGroupsid**
> Group GroupPrototypePatchAttributesPatchGroupsid(ctx, id, optional)
Patch attributes for a model instance and persist it into the data source.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Group id | 
 **optional** | ***GroupApiGroupPrototypePatchAttributesPatchGroupsidOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupApiGroupPrototypePatchAttributesPatchGroupsidOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**optional.Interface of Group**](Group.md)| An object of model property name/value pairs | 

### Return type

[**Group**](Group.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupPrototypePatchAttributesPutGroupsid**
> Group GroupPrototypePatchAttributesPutGroupsid(ctx, id, optional)
Patch attributes for a model instance and persist it into the data source.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Group id | 
 **optional** | ***GroupApiGroupPrototypePatchAttributesPutGroupsidOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupApiGroupPrototypePatchAttributesPutGroupsidOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**optional.Interface of Group**](Group.md)| An object of model property name/value pairs | 

### Return type

[**Group**](Group.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupPrototypeUnlinkCustomers**
> GroupPrototypeUnlinkCustomers(ctx, id, fk)
Remove the customers relation to an item by id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Group id | 
  **fk** | **string**| Foreign key for customers | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

