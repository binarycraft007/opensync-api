# TopologyResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **string** | schema version # of the API response | [default to null]
**Timestamp** | **string** | ISO Date of returned topology, optional if in included in client API call | [optional] [default to null]
**Topology** | [**[]NodeResponse**](NodeResponse.md) | List of Nodes | [default to null]
**GatewayConnectionState** | **string** | connected or disconnected | [default to null]
**NodeConnectedCount** | **float64** | count total nodes with &#39;connected&#39; connectionState | [default to null]
**NodeClaimedCount** | **float64** | nodeArray length | [default to null]
**ConnectedDeviceCount** | **float64** | count total devices with &#39;connected&#39; connectionState | [default to null]
**NetworkModeRealized** | **string** | default value is &#39;unknow&#39;, &#39;bridge&#39; if there is no router for all connected pods, &#39;router&#39; if there is a connected pod with &#39;router&#39; | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


