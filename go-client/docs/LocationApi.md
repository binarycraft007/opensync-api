# \LocationApi

All URIs are relative to *https://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LocationCount**](LocationApi.md#LocationCount) | **Get** /Locations/count | Count instances of the model matched by where from the data source.
[**LocationFind**](LocationApi.md#LocationFind) | **Get** /Locations | Find all instances of the model matched by filter from the data source.
[**LocationFindById**](LocationApi.md#LocationFindById) | **Get** /Locations/{id} | Find a model instance by {{id}} from the data source.
[**LocationPrototypeCheckCustomerLoginToTurnOffWifi**](LocationApi.md#LocationPrototypeCheckCustomerLoginToTurnOffWifi) | **Post** /Locations/{id}/checkCustomerLoginToTurnOffWifi | Turn off WIfi Radio if no app login found for a customer
[**LocationPrototypeClaimMultipleNodes**](LocationApi.md#LocationPrototypeClaimMultipleNodes) | **Post** /Locations/{id}/nodes | Claim all nodes for a Location ID.
[**LocationPrototypeCountInvitations**](LocationApi.md#LocationPrototypeCountInvitations) | **Get** /Locations/{id}/invitations/count | Counts invitations of Location.
[**LocationPrototypeCountPendingWhitelistRequests**](LocationApi.md#LocationPrototypeCountPendingWhitelistRequests) | **Get** /Locations/{id}/_pendingWhitelistRequests/count | Counts _pendingWhitelistRequests of Location.
[**LocationPrototypeCreateInvitations**](LocationApi.md#LocationPrototypeCreateInvitations) | **Post** /Locations/{id}/invitations | Creates a new instance in invitations of this model.
[**LocationPrototypeCreatePendingWhitelistRequests**](LocationApi.md#LocationPrototypeCreatePendingWhitelistRequests) | **Post** /Locations/{id}/_pendingWhitelistRequests | Creates a new instance in _pendingWhitelistRequests of this model.
[**LocationPrototypeDeleteDeviceSecurityPolicyAnomalyBlacklist**](LocationApi.md#LocationPrototypeDeleteDeviceSecurityPolicyAnomalyBlacklist) | **Delete** /Locations/{id}/devices/{mac}/securityPolicy/anomaly/websites/blacklist/{fqdn} | Update a Location&#39;s Anomaly Security Policy for a location ID to remove a blacklisted DNS entry.
[**LocationPrototypeDeleteInvitations**](LocationApi.md#LocationPrototypeDeleteInvitations) | **Delete** /Locations/{id}/invitations | Deletes all invitations of this model.
[**LocationPrototypeDeletePendingWhitelistRequests**](LocationApi.md#LocationPrototypeDeletePendingWhitelistRequests) | **Delete** /Locations/{id}/_pendingWhitelistRequests | Deletes all _pendingWhitelistRequests of this model.
[**LocationPrototypeDestroyByIdInvitations**](LocationApi.md#LocationPrototypeDestroyByIdInvitations) | **Delete** /Locations/{id}/invitations/{fk} | Delete a related item by id for invitations.
[**LocationPrototypeDestroyByIdPendingWhitelistRequests**](LocationApi.md#LocationPrototypeDestroyByIdPendingWhitelistRequests) | **Delete** /Locations/{id}/_pendingWhitelistRequests/{fk} | Delete a related item by id for _pendingWhitelistRequests.
[**LocationPrototypeDeviceDetailsWithSsid**](LocationApi.md#LocationPrototypeDeviceDetailsWithSsid) | **Get** /Locations/{id}/devices/{mac}/detailsWithSsid | Get device parameters by mac id and the ssid by networkId
[**LocationPrototypeFindByIdInvitations**](LocationApi.md#LocationPrototypeFindByIdInvitations) | **Get** /Locations/{id}/invitations/{fk} | Find a related item by id for invitations.
[**LocationPrototypeFindByIdPendingWhitelistRequests**](LocationApi.md#LocationPrototypeFindByIdPendingWhitelistRequests) | **Get** /Locations/{id}/_pendingWhitelistRequests/{fk} | Find a related item by id for _pendingWhitelistRequests.
[**LocationPrototypeFindRoomByNodeIdNodeMacAndDeviceMac**](LocationApi.md#LocationPrototypeFindRoomByNodeIdNodeMacAndDeviceMac) | **Get** /Locations/{id}/rooms/search/{search} | Internal integration use only: Search to identify if the node or device is assigned to a room.
[**LocationPrototypeGetAppTime**](LocationApi.md#LocationPrototypeGetAppTime) | **Get** /Locations/{id}/appTime | Get the appTime configs to use for a location
[**LocationPrototypeGetBackhaul**](LocationApi.md#LocationPrototypeGetBackhaul) | **Get** /Locations/{id}/backhaul | 
[**LocationPrototypeGetCommandState**](LocationApi.md#LocationPrototypeGetCommandState) | **Get** /Locations/{id}/command/state | Get speed test result report state
[**LocationPrototypeGetDeviceSoundingState**](LocationApi.md#LocationPrototypeGetDeviceSoundingState) | **Get** /Locations/{id}/homeSecurity/devices/sounding | Fetch the sounding states for eligible devices in this location
[**LocationPrototypeGetDevicesByMacName**](LocationApi.md#LocationPrototypeGetDevicesByMacName) | **Get** /Locations/{id}/devices/{mac} | Get the name and icon of device by mac lookup
[**LocationPrototypeGetEventHistory**](LocationApi.md#LocationPrototypeGetEventHistory) | **Get** /Locations/{id}/homeSecurity/events/history | Fetch the event history for this location
[**LocationPrototypeGetForceGraph**](LocationApi.md#LocationPrototypeGetForceGraph) | **Get** /Locations/{id}/forceGraph | Vertices[] and edges[] used to display a Network Topology.
[**LocationPrototypeGetGatewayAccount**](LocationApi.md#LocationPrototypeGetGatewayAccount) | **Get** /Locations/{id}/gatewayAccount | Get AccountID and GatwewayID of a location
[**LocationPrototypeGetHomeAwayConfig**](LocationApi.md#LocationPrototypeGetHomeAwayConfig) | **Get** /Locations/{id}/homeAway | Get the homeAway configs to use for a location
[**LocationPrototypeGetHomeSecurity**](LocationApi.md#LocationPrototypeGetHomeSecurity) | **Get** /Locations/{id}/homeSecurity | Fetch the home security configuration for this location
[**LocationPrototypeGetInvitations**](LocationApi.md#LocationPrototypeGetInvitations) | **Get** /Locations/{id}/invitations | Queries invitations of Location.
[**LocationPrototypeGetLocationPartnerIdProfileInfo**](LocationApi.md#LocationPrototypeGetLocationPartnerIdProfileInfo) | **Get** /Locations/{id}/partnerIdProfileInfo | Get the partner id, location profile and other info for this location
[**LocationPrototypeGetLocationSummary**](LocationApi.md#LocationPrototypeGetLocationSummary) | **Get** /Locations/{id}/summary | Get the locationSummary for this location
[**LocationPrototypeGetMotionHistory**](LocationApi.md#LocationPrototypeGetMotionHistory) | **Get** /Locations/{id}/homeSecurity/motionHistory | Fetch the motion density history for this location
[**LocationPrototypeGetMotionStateHistory**](LocationApi.md#LocationPrototypeGetMotionStateHistory) | **Get** /Locations/{id}/homeSecurity/motionHistory/state | Fetch the motion state history for this location
[**LocationPrototypeGetNodes**](LocationApi.md#LocationPrototypeGetNodes) | **Get** /Locations/{id}/nodes | Queries nodes of Location.
[**LocationPrototypeGetPendingWhitelistRequests**](LocationApi.md#LocationPrototypeGetPendingWhitelistRequests) | **Get** /Locations/{id}/_pendingWhitelistRequests | Queries _pendingWhitelistRequests of Location.
[**LocationPrototypeGetWifiMotion**](LocationApi.md#LocationPrototypeGetWifiMotion) | **Get** /Locations/{id}/wifiMotion | Get WifiMotion config for this location
[**LocationPrototypeGroupProvisioning**](LocationApi.md#LocationPrototypeGroupProvisioning) | **Get** /Locations/{id}/groupProvisioning | Get Group Provisioning details
[**LocationPrototypeMarketingExport**](LocationApi.md#LocationPrototypeMarketingExport) | **Get** /Locations/{id}/marketingExport | Get detailed information of location for updating CRMs.
[**LocationPrototypeNasRedirect**](LocationApi.md#LocationPrototypeNasRedirect) | **Get** /Locations/{id}/nasRedirect | Handle proxy redirects from walled-garden networks requesting network access
[**LocationPrototypePatchDeviceSoundingState**](LocationApi.md#LocationPrototypePatchDeviceSoundingState) | **Patch** /Locations/{id}/homeSecurity/devices/sounding | Patch the sounding states for the given devices
[**LocationPrototypePatchHomeSecurity**](LocationApi.md#LocationPrototypePatchHomeSecurity) | **Patch** /Locations/{id}/homeSecurity | Enable/disable live motion streaming and/or motion events for this location
[**LocationPrototypePatchHomeSecuritySensitivity**](LocationApi.md#LocationPrototypePatchHomeSecuritySensitivity) | **Patch** /Locations/{id}/homeSecurity/sensitivity | Configure motion event configuration for this location
[**LocationPrototypePatchMobileAppSignInRequired**](LocationApi.md#LocationPrototypePatchMobileAppSignInRequired) | **Patch** /Locations/{id}/mobileAppSignInRequired | Patch mobileAppSignInRequired object handling the wifi disabling timer
[**LocationPrototypePatchUpriseOrFlex**](LocationApi.md#LocationPrototypePatchUpriseOrFlex) | **Patch** /Locations/{id} | Update uprise value for the location
[**LocationPrototypePatchWifiMotion**](LocationApi.md#LocationPrototypePatchWifiMotion) | **Patch** /Locations/{id}/wifiMotion | Enable/disable WifiMotion feature for this location
[**LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklist**](LocationApi.md#LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklist) | **Post** /Locations/{id}/devices/{mac}/securityPolicy/anomaly/websites/blacklist | Update a Device&#39;s Anomaly Security Policy for a location ID to include a blacklisted website.
[**LocationPrototypePutGeoIp**](LocationApi.md#LocationPrototypePutGeoIp) | **Put** /Locations/{id}/geoIp | Update raw Geo IP information for the location
[**LocationPrototypePutResyncLocation**](LocationApi.md#LocationPrototypePutResyncLocation) | **Put** /Locations/{id}/resyncLocation | Trigger Controller to refresh all Customer mongo data for a Location ID.
[**LocationPrototypeResyncWebconfig**](LocationApi.md#LocationPrototypeResyncWebconfig) | **Post** /Locations/{id}/resyncWebconfig | sync SSID data from location to all its claimed nodes.
[**LocationPrototypeUnclaimNode**](LocationApi.md#LocationPrototypeUnclaimNode) | **Delete** /Locations/{id}/nodes/{nodeId} | Unclaim a particular Node from a Location with the option of preserving the original Package ID.
[**LocationPrototypeUpdateByIdInvitations**](LocationApi.md#LocationPrototypeUpdateByIdInvitations) | **Put** /Locations/{id}/invitations/{fk} | Update a related item by id for invitations.
[**LocationPrototypeUpdateByIdPendingWhitelistRequests**](LocationApi.md#LocationPrototypeUpdateByIdPendingWhitelistRequests) | **Put** /Locations/{id}/_pendingWhitelistRequests/{fk} | Update a related item by id for _pendingWhitelistRequests.


# **LocationCount**
> InlineResponse2001 LocationCount(ctx, optional)
Count instances of the model matched by where from the data source.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LocationApiLocationCountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocationApiLocationCountOpts struct

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

# **LocationFind**
> []Location LocationFind(ctx, optional)
Find all instances of the model matched by filter from the data source.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LocationApiLocationFindOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocationApiLocationFindOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **optional.String**| Filter defining fields, where, include, order, offset, and limit - must be a JSON-encoded string (&#x60;{\&quot;where\&quot;:{\&quot;something\&quot;:\&quot;value\&quot;}}&#x60;).  See https://loopback.io/doc/en/lb3/Querying-data.html#using-stringified-json-in-rest-queries for more details. | 

### Return type

[**[]Location**](Location.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocationFindById**
> Location LocationFindById(ctx, id, optional)
Find a model instance by {{id}} from the data source.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Model id | 
 **optional** | ***LocationApiLocationFindByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocationApiLocationFindByIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **optional.String**| Filter defining fields and include - must be a JSON-encoded string ({\&quot;something\&quot;:\&quot;value\&quot;}) | 

### Return type

[**Location**](Location.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocationPrototypeCheckCustomerLoginToTurnOffWifi**
> interface{} LocationPrototypeCheckCustomerLoginToTurnOffWifi(ctx, id)
Turn off WIfi Radio if no app login found for a customer

<div><strong>200</strong>: Success.</div> <div><strong>404</strong>: location id does not exist and is not known to Plume</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Location id | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocationPrototypeClaimMultipleNodes**
> interface{} LocationPrototypeClaimMultipleNodes(ctx, id, optional)
Claim all nodes for a Location ID.

<div><strong>204</strong>: Success.</div> <div><strong>400</strong>: Required nodes field is missing.</div> <div><strong>422</strong>: Request contain wrong value or exceded the max number of 32 pods to claim.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Location id | 
 **optional** | ***LocationApiLocationPrototypeClaimMultipleNodesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocationApiLocationPrototypeClaimMultipleNodesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nodes** | **optional.String**| array of serialNumber/ids | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocationPrototypeCountInvitations**
> InlineResponse2001 LocationPrototypeCountInvitations(ctx, id, optional)
Counts invitations of Location.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Location id | 
 **optional** | ***LocationApiLocationPrototypeCountInvitationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocationApiLocationPrototypeCountInvitationsOpts struct

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

# **LocationPrototypeCountPendingWhitelistRequests**
> InlineResponse2001 LocationPrototypeCountPendingWhitelistRequests(ctx, id, optional)
Counts _pendingWhitelistRequests of Location.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Location id | 
 **optional** | ***LocationApiLocationPrototypeCountPendingWhitelistRequestsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocationApiLocationPrototypeCountPendingWhitelistRequestsOpts struct

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

# **LocationPrototypeCreateInvitations**
> Invitation LocationPrototypeCreateInvitations(ctx, id, optional)
Creates a new instance in invitations of this model.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Location id | 
 **optional** | ***LocationApiLocationPrototypeCreateInvitationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocationApiLocationPrototypeCreateInvitationsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**optional.Interface of Invitation**](Invitation.md)|  | 

### Return type

[**Invitation**](Invitation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocationPrototypeCreatePendingWhitelistRequests**
> PendingWhitelistRequests LocationPrototypeCreatePendingWhitelistRequests(ctx, id, optional)
Creates a new instance in _pendingWhitelistRequests of this model.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Location id | 
 **optional** | ***LocationApiLocationPrototypeCreatePendingWhitelistRequestsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocationApiLocationPrototypeCreatePendingWhitelistRequestsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**optional.Interface of PendingWhitelistRequests**](PendingWhitelistRequests.md)|  | 

### Return type

[**PendingWhitelistRequests**](PendingWhitelistRequests.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocationPrototypeDeleteDeviceSecurityPolicyAnomalyBlacklist**
> LocationPrototypeDeleteDeviceSecurityPolicyAnomalyBlacklist(ctx, id, mac, fqdn)
Update a Location's Anomaly Security Policy for a location ID to remove a blacklisted DNS entry.

<div><strong>204</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id, WifiNetwork, Device or DNS does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Location id | 
  **mac** | **string**|  | 
  **fqdn** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocationPrototypeDeleteInvitations**
> LocationPrototypeDeleteInvitations(ctx, id)
Deletes all invitations of this model.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Location id | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocationPrototypeDeletePendingWhitelistRequests**
> LocationPrototypeDeletePendingWhitelistRequests(ctx, id)
Deletes all _pendingWhitelistRequests of this model.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Location id | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocationPrototypeDestroyByIdInvitations**
> LocationPrototypeDestroyByIdInvitations(ctx, id, fk)
Delete a related item by id for invitations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Location id | 
  **fk** | **string**| Foreign key for invitations | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocationPrototypeDestroyByIdPendingWhitelistRequests**
> LocationPrototypeDestroyByIdPendingWhitelistRequests(ctx, id, fk)
Delete a related item by id for _pendingWhitelistRequests.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Location id | 
  **fk** | **string**| Foreign key for _pendingWhitelistRequests | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocationPrototypeDeviceDetailsWithSsid**
> interface{} LocationPrototypeDeviceDetailsWithSsid(ctx, id, mac, networkId)
Get device parameters by mac id and the ssid by networkId

<div>To be used by notification for fetching device name and icon.</div> <div><strong>200</strong>: Success, device details returned.</div> <div><strong>404</strong>: Location ID, Device mac or Network id not found</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Location id | 
  **mac** | **string**| mac of device | 
  **networkId** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocationPrototypeFindByIdInvitations**
> Invitation LocationPrototypeFindByIdInvitations(ctx, id, fk)
Find a related item by id for invitations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Location id | 
  **fk** | **string**| Foreign key for invitations | 

### Return type

[**Invitation**](Invitation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocationPrototypeFindByIdPendingWhitelistRequests**
> PendingWhitelistRequests LocationPrototypeFindByIdPendingWhitelistRequests(ctx, id, fk)
Find a related item by id for _pendingWhitelistRequests.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Location id | 
  **fk** | **string**| Foreign key for _pendingWhitelistRequests | 

### Return type

[**PendingWhitelistRequests**](PendingWhitelistRequests.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocationPrototypeFindRoomByNodeIdNodeMacAndDeviceMac**
> interface{} LocationPrototypeFindRoomByNodeIdNodeMacAndDeviceMac(ctx, id, search)
Internal integration use only: Search to identify if the node or device is assigned to a room.

<div>To be used by Notification API for fetching Room info.</div> <div><strong>200</strong>: Success, room details returned.</div> <div><strong>404</strong>: Location, device or node not found</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Location id | 
  **search** | **string**| Node ID (Serial Number), Node WiFi Radio Mac, or Device Mac | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocationPrototypeGetAppTime**
> interface{} LocationPrototypeGetAppTime(ctx, id)
Get the appTime configs to use for a location

<div><strong>200</strong>: Success, accountID and gatewayID returned.</div> <div><strong>404</strong>: Location does not exist</div> <div><strong>500</strong>: Internal Server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Location id | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocationPrototypeGetBackhaul**
> interface{} LocationPrototypeGetBackhaul(ctx, id, locationId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Location id | 
  **locationId** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocationPrototypeGetCommandState**
> interface{} LocationPrototypeGetCommandState(ctx, id)
Get speed test result report state

<div><strong>200</strong>: Success, return report state.</div> <div><strong>401</strong>: Authorization required.</div> <div><strong>500</strong>: Internal server error.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Location id | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocationPrototypeGetDeviceSoundingState**
> interface{} LocationPrototypeGetDeviceSoundingState(ctx, id, optional)
Fetch the sounding states for eligible devices in this location

<div><strong>200</strong>: Success, device sounding states returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id does not exist and is not known to Plume</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Location id | 
 **optional** | ***LocationApiLocationPrototypeGetDeviceSoundingStateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocationApiLocationPrototypeGetDeviceSoundingStateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mac** | **optional.String**| Optional mac address for single device lookup (fetches all devices by default) | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocationPrototypeGetDevicesByMacName**
> interface{} LocationPrototypeGetDevicesByMacName(ctx, id, mac)
Get the name and icon of device by mac lookup

<div>To be used by notification for fetching device name and icon.</div> <div><strong>200</strong>: Success, device details returned.</div> <div><strong>404</strong>: Location ID or Device mac not found</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Location id | 
  **mac** | **string**| mac of device | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocationPrototypeGetEventHistory**
> []interface{} LocationPrototypeGetEventHistory(ctx, id, optional)
Fetch the event history for this location

<div><strong>200</strong>: Success, event array returned.</div> <div><strong>401</strong>: Authorization required</div> <div><strong>404</strong>: location id does not exist and is not known to Plume</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Location id | 
 **optional** | ***LocationApiLocationPrototypeGetEventHistoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocationApiLocationPrototypeGetEventHistoryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **optional.Float64**| UTC unix ts | 
 **to** | **optional.Float64**| UTC unix ts, defaults to now | 
 **category** | **optional.String**| Filter events by category (Motion or Plume [config changes]). Multiple categories can be passed as a comma-separated string. Default is both. | 
 **limit** | **optional.Float64**| Maximum number of events to return; defaults to 10 | 
 **sort** | **optional.Bool**| whether the returned events will be post-sorted by timestamp | 

### Return type

**[]interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocationPrototypeGetForceGraph**
> interface{} LocationPrototypeGetForceGraph(ctx, id, optional)
Vertices[] and edges[] used to display a Network Topology.

<div>The data used to initialize and dynamically display and update a Topology.</div> <div>Can also be used to get a network's list of nodes + devices (a.k.a. vertices) and links (a.k.a., edges).</div><div>&nbsp;</div> <div><strong>200</strong>: Success, graph structure returned.</div> <div><strong>404</strong>: customer id or location id does not exist and is not known to Plume</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Location id | 
 **optional** | ***LocationApiLocationPrototypeGetForceGraphOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocationApiLocationPrototypeGetForceGraphOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **allSSIDs** | **optional.Bool**|  | 
 **showPartnerComponent** | **optional.Bool**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocationPrototypeGetGatewayAccount**
> interface{} LocationPrototypeGetGatewayAccount(ctx, id)
Get AccountID and GatwewayID of a location

<div><strong>200</strong>: Success, accountID and gatewayID returned.</div> <div><strong>404</strong>: Location does not exist</div> <div><strong>500</strong>: Internal Server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Location id | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocationPrototypeGetHomeAwayConfig**
> interface{} LocationPrototypeGetHomeAwayConfig(ctx, id)
Get the homeAway configs to use for a location

<div><strong>200</strong>: Success, accountID and gatewayID returned.</div> <div><strong>404</strong>: Location does not exist</div> <div><strong>500</strong>: Internal Server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Location id | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocationPrototypeGetHomeSecurity**
> HomeSecurity LocationPrototypeGetHomeSecurity(ctx, id)
Fetch the home security configuration for this location

<div><strong>200</strong>: Success, HomeSecurity object returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: location id does not exist and is not known to Plume</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Location id | 

### Return type

[**HomeSecurity**](HomeSecurity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocationPrototypeGetInvitations**
> []Invitation LocationPrototypeGetInvitations(ctx, id, optional)
Queries invitations of Location.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Location id | 
 **optional** | ***LocationApiLocationPrototypeGetInvitationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocationApiLocationPrototypeGetInvitationsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **optional.String**|  | 

### Return type

[**[]Invitation**](Invitation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocationPrototypeGetLocationPartnerIdProfileInfo**
> interface{} LocationPrototypeGetLocationPartnerIdProfileInfo(ctx, id)
Get the partner id, location profile and other info for this location

<div><strong>200</strong>: Success, wifiMotion object returned.</div> <div><strong>401</strong>: Authorization required</div> <div><strong>404</strong>: location id does not exist and is not known to Plume</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Location id | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocationPrototypeGetLocationSummary**
> interface{} LocationPrototypeGetLocationSummary(ctx, id)
Get the locationSummary for this location

<div><strong>200</strong>: Success, wifiMotion object returned.</div> <div><strong>401</strong>: Authorization required</div> <div><strong>404</strong>: location id does not exist and is not known to Plume</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Location id | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocationPrototypeGetMotionHistory**
> []interface{} LocationPrototypeGetMotionHistory(ctx, id, optional)
Fetch the motion density history for this location

<div><strong>200</strong>: Success, motion density array returned.</div> <div><strong>401</strong>: Authorization required</div> <div><strong>404</strong>: location id does not exist and is not known to Plume</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Location id | 
 **optional** | ***LocationApiLocationPrototypeGetMotionHistoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocationApiLocationPrototypeGetMotionHistoryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **optional.Float64**| UTC unix ts | 
 **to** | **optional.Float64**| UTC unix ts, defaults to now | 
 **bucket** | **optional.Float64**| number of seconds in density calculation window; returned data points represent % of non-zero intensity values in the window | [default to 3600]

### Return type

**[]interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocationPrototypeGetMotionStateHistory**
> []interface{} LocationPrototypeGetMotionStateHistory(ctx, id, optional)
Fetch the motion state history for this location

<div><strong>200</strong>: Success, motion state array returned (Each element of the array is in the form [\"val\", \"unix_ts\"], where \"val\" is one of:  <div>0 - Not armed, not tripped</div> <div>1 - Not armed, tripped</div> <div>2 - Armed, not tripped</div> <div>3 - Armed, tripped</div></div> <div><strong>401</strong>: Authorization required</div> <div><strong>404</strong>: location id does not exist and is not known to Plume</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Location id | 
 **optional** | ***LocationApiLocationPrototypeGetMotionStateHistoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocationApiLocationPrototypeGetMotionStateHistoryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **optional.Float64**| UTC unix ts | 
 **to** | **optional.Float64**| UTC unix ts, defaults to now | 
 **bucket** | **optional.Float64**| number of seconds in density calculation window; returned data points represent % of non-zero intensity values in the window | [default to 3600]

### Return type

**[]interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocationPrototypeGetNodes**
> []Node LocationPrototypeGetNodes(ctx, id, optional)
Queries nodes of Location.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Location id | 
 **optional** | ***LocationApiLocationPrototypeGetNodesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocationApiLocationPrototypeGetNodesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **optional.String**|  | 

### Return type

[**[]Node**](Node.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocationPrototypeGetPendingWhitelistRequests**
> []PendingWhitelistRequests LocationPrototypeGetPendingWhitelistRequests(ctx, id, optional)
Queries _pendingWhitelistRequests of Location.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Location id | 
 **optional** | ***LocationApiLocationPrototypeGetPendingWhitelistRequestsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocationApiLocationPrototypeGetPendingWhitelistRequestsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **optional.String**|  | 

### Return type

[**[]PendingWhitelistRequests**](PendingWhitelistRequests.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocationPrototypeGetWifiMotion**
> WifiMotion LocationPrototypeGetWifiMotion(ctx, id)
Get WifiMotion config for this location

<div><strong>200</strong>: Success, wifiMotion object returned.</div> <div><strong>401</strong>: Authorization required</div> <div><strong>404</strong>: location id does not exist and is not known to Plume</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Location id | 

### Return type

[**WifiMotion**](WifiMotion.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocationPrototypeGroupProvisioning**
> interface{} LocationPrototypeGroupProvisioning(ctx, id)
Get Group Provisioning details

<div><strong>200</strong>: Success.</div> <div><strong>404</strong>: Location does not exist</div> <div><strong>500</strong>: Internal Server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Location id | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocationPrototypeMarketingExport**
> interface{} LocationPrototypeMarketingExport(ctx, id)
Get detailed information of location for updating CRMs.

<div><strong>200</strong>: Success, location data in response.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Location id | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocationPrototypeNasRedirect**
> interface{} LocationPrototypeNasRedirect(ctx, id, proxy, proxyMac, nodeMac, ssid)
Handle proxy redirects from walled-garden networks requesting network access

<div><strong>200</strong>: Success, redirect URL returned.</div> <div><strong>404</strong>: Location does not exist</div> <div><strong>500</strong>: Internal Server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Location id | 
  **proxy** | **string**| client IP address | 
  **proxyMac** | **string**| client MAC address | 
  **nodeMac** | **string**| gateway MAC address | 
  **ssid** | **string**| guest SSID client is attempting to join | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocationPrototypePatchDeviceSoundingState**
> interface{} LocationPrototypePatchDeviceSoundingState(ctx, id, soundingStates)
Patch the sounding states for the given devices

<div><strong>200</strong>: Success, device sounding states returned.</div> <div><strong>401</strong>: Authorization required</div> <div><strong>404</strong>: Location id does not exist and is not known to Plume</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Location id | 
  **soundingStates** | **interface{}**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocationPrototypePatchHomeSecurity**
> HomeSecurity LocationPrototypePatchHomeSecurity(ctx, id, optional)
Enable/disable live motion streaming and/or motion events for this location

<div><strong>200</strong>: Success, updated HomeSecurity object returned.</div> <div><strong>401</strong>: Authorization required</div> <div><strong>404</strong>: location id does not exist and is not known to Plume</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Location id | 
 **optional** | ***LocationApiLocationPrototypePatchHomeSecurityOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocationApiLocationPrototypePatchHomeSecurityOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **source** | **optional.String**| Source of patch request; must be one of \&quot;user\&quot; or \&quot;geofence\&quot; | 
 **liveMotionEnabled** | **optional.Bool**|  | 
 **motionEventsEnabled** | **optional.Bool**|  | 
 **homeAwayActive** | **optional.Bool**| Enable/disable motion events based on location Homeaway state | 

### Return type

[**HomeSecurity**](HomeSecurity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocationPrototypePatchHomeSecuritySensitivity**
> HomeSecurity LocationPrototypePatchHomeSecuritySensitivity(ctx, id, optional)
Configure motion event configuration for this location

<div><strong>200</strong>: Success, updated HomeSecurity object returned.</div> <div><strong>401</strong>: Authorization required</div> <div><strong>404</strong>: location id does not exist and is not known to Plume</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Location id | 
 **optional** | ***LocationApiLocationPrototypePatchHomeSecuritySensitivityOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocationApiLocationPrototypePatchHomeSecuritySensitivityOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cooldown** | **optional.Float64**| sets required rest period for motion detected events to end, in seconds | 
 **petMode** | **optional.String**| adjusts sensitivity of motion detected events for pets; must be one of \&quot;none\&quot;, \&quot;under10\&quot;, \&quot;10to30\&quot;, \&quot;over30\&quot; and can only be set if sensitivity &#x3D; high | 
 **sensitivity** | **optional.String**| adjusts sensitivity of motion detected events; must be one of \&quot;low\&quot;, \&quot;medium\&quot;, \&quot;high\&quot; | 

### Return type

[**HomeSecurity**](HomeSecurity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocationPrototypePatchMobileAppSignInRequired**
> interface{} LocationPrototypePatchMobileAppSignInRequired(ctx, id, state)
Patch mobileAppSignInRequired object handling the wifi disabling timer

<div><strong>200</strong>: Success.</div> <div><strong>400</strong>: mobileAppSignInRequired is not configured for this location</div> <div><strong>404</strong>: location id does not exist and is not known to Plume</div> <div><strong>429</strong>: This location has already cancelled app engagement timer</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Location id | 
  **state** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocationPrototypePatchUpriseOrFlex**
> interface{} LocationPrototypePatchUpriseOrFlex(ctx, id, optional)
Update uprise value for the location

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required.</div> <div><strong>404</strong>: location id, does not exist.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Location id | 
 **optional** | ***LocationApiLocationPrototypePatchUpriseOrFlexOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocationApiLocationPrototypePatchUpriseOrFlexOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **uprise** | **optional.Bool**|  | 
 **flex** | **optional.Bool**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocationPrototypePatchWifiMotion**
> WifiMotion LocationPrototypePatchWifiMotion(ctx, id, auto)
Enable/disable WifiMotion feature for this location

<div><strong>200</strong>: Success, updated object returned.</div> <div><strong>401</strong>: Authorization required</div> <div><strong>404</strong>: location id does not exist and is not known to Plume</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Location id | 
  **auto** | **bool**|  | 

### Return type

[**WifiMotion**](WifiMotion.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklist**
> interface{} LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklist(ctx, id, mac, deviceType, fqdn, optional)
Update a Device's Anomaly Security Policy for a location ID to include a blacklisted website.

<div><strong>200</strong>: Success.</div> <div><strong>400</strong>: Required fields missing or field type is incorrect.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id, WifiNetwork, or Device does not exist and is not known to Plume</div> <div><strong>422</strong>: DNS value is invalid.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Location id | 
  **mac** | **string**|  | 
  **deviceType** | **string**|  | 
  **fqdn** | **string**|  | 
 **optional** | ***LocationApiLocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocationApiLocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **ipv4** | **optional.String**|  | 
 **ipv6** | **optional.String**|  | 
 **ttl** | **optional.Float64**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocationPrototypePutGeoIp**
> interface{} LocationPrototypePutGeoIp(ctx, id, optional)
Update raw Geo IP information for the location

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required.</div> <div><strong>404</strong>: location id, does not exist.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Location id | 
 **optional** | ***LocationApiLocationPrototypePutGeoIpOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocationApiLocationPrototypePutGeoIpOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **geoIp** | **optional.Interface{}**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocationPrototypePutResyncLocation**
> InlineResponse20012 LocationPrototypePutResyncLocation(ctx, id, id2)
Trigger Controller to refresh all Customer mongo data for a Location ID.

<div><strong>200</strong>: Success, triggered right way.</div> <div><strong>401</strong>: Authorization required.</div> <div><strong>404</strong>: location id, does not exist.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Location id | 
  **id2** | **string**|  | 

### Return type

[**InlineResponse20012**](inline_response_200_12.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocationPrototypeResyncWebconfig**
> interface{} LocationPrototypeResyncWebconfig(ctx, id, id2)
sync SSID data from location to all its claimed nodes.

<div><strong>204</strong>: Success.</div> <div><strong>404</strong>: Wifi network does not exist.</div> <div><strong>404</strong>: location id, does not exist.</div> <div><strong>422</strong>: Endpoint is not allowed in current deployment.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Location id | 
  **id2** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocationPrototypeUnclaimNode**
> LocationPrototypeUnclaimNode(ctx, id, nodeId, optional)
Unclaim a particular Node from a Location with the option of preserving the original Package ID.

<div><strong>204</strong>: Success, a job well done.</div> <div><strong>400</strong>: Pod already unclaimed.</div> <div><strong>401</strong>: Authorization required </div> <div><strong>403</strong>: the node is online, and can not be unclaimed.<br/> <div><strong>404</strong>: location id not found, nodeId missing from URL,<br/> or location has zero owned pods.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Location id | 
  **nodeId** | **string**|  | 
 **optional** | ***LocationApiLocationPrototypeUnclaimNodeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocationApiLocationPrototypeUnclaimNodeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **preservePackId** | **optional.Bool**| Whether or not packId should remain the same | 
 **forceUnclaim** | **optional.Bool**| Unclaim regardless of pod connectivity | 
 **purgeGroupIds** | **optional.Bool**| Whether or not groupIds should be kept on the node | 
 **removeAccountId** | **optional.Bool**| delete account id on the inventory node | [default to false]
 **unclaimReason** | **optional.String**| Used by controller to determine what actions to take for nodeClaimChanged | 
 **incrementFactoryResetCounter** | **optional.Bool**| Whether or not to increment the factory reset counter | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocationPrototypeUpdateByIdInvitations**
> Invitation LocationPrototypeUpdateByIdInvitations(ctx, id, fk, optional)
Update a related item by id for invitations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Location id | 
  **fk** | **string**| Foreign key for invitations | 
 **optional** | ***LocationApiLocationPrototypeUpdateByIdInvitationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocationApiLocationPrototypeUpdateByIdInvitationsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**optional.Interface of Invitation**](Invitation.md)|  | 

### Return type

[**Invitation**](Invitation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocationPrototypeUpdateByIdPendingWhitelistRequests**
> PendingWhitelistRequests LocationPrototypeUpdateByIdPendingWhitelistRequests(ctx, id, fk, optional)
Update a related item by id for _pendingWhitelistRequests.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Location id | 
  **fk** | **string**| Foreign key for _pendingWhitelistRequests | 
 **optional** | ***LocationApiLocationPrototypeUpdateByIdPendingWhitelistRequestsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocationApiLocationPrototypeUpdateByIdPendingWhitelistRequestsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**optional.Interface of PendingWhitelistRequests**](PendingWhitelistRequests.md)|  | 

### Return type

[**PendingWhitelistRequests**](PendingWhitelistRequests.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

