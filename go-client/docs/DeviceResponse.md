# DeviceResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [optional] [default to null]
**Icon** | **string** | The mapping is: &#39;mediaSTB&#39;: [&#39;appletv&#39;, &#39;chromecast&#39;, &#39;tivo&#39;, &#39;np-&#39;], &#39;camera&#39; : [&#39;cam&#39;], &#39;gamingConsole&#39; : [&#39;xbox&#39;, &#39;ps3&#39;, &#39;ps4&#39;], &#39;smartPhone&#39; : [&#39;iphone&#39;, &#39;ipod&#39;, &#39;phone&#39;, &#39;android&#39;], &#39;tablet&#39; : [&#39;ipad&#39;, &#39;kindle&#39;], &#39;imac&#39; : [&#39;imac&#39;, &#39;macintosh&#39;], &#39;desktop&#39; : [&#39;desktop&#39;], &#39;routerbox&#39; : [&#39;linksys&#39;, &#39;netgear&#39;, &#39;asus&#39;, &#39;belkin&#39;, &#39;airport&#39;], &#39;laptop&#39; : [&#39;mbp&#39;, &#39;air&#39;, &#39;macbook&#39;, &#39;pc&#39;, &#39;thinkpad&#39;], &#39;printer&#39; : [&#39;printer&#39;, &#39;epson&#39;], &#39;tv&#39; : [&#39;tv&#39;, &#39;vizo&#39;], &#39;voipPhone&#39; : [&#39;voip&#39;], &#39;speaker&#39; : [&#39;sonos&#39;], &#39;lightbulb&#39; : [&#39;lifx&#39;] | [optional] [default to null]
**Mac** | **string** |  | [default to null]
**AccessZone** | **string** | home | guests | internetAccessOnly | [optional] [default to null]
**KeyId** | **float64** | unique id of the WifiNetwork.keys[x] that the device is connected to or last connected to | [optional] [default to null]
**Medium** | **string** | wifi or ethernet | [optional] [default to null]
**Ip** | **string** |  | [default to null]
**FreqBand** | **string** | 2.4G or 5G or 6G, undefined if medium ethernet | [optional] [default to null]
**Channel** | **float64** | undefined if medium ethernet | [optional] [default to null]
**Name** | **string** | host name else mac | [optional] [default to null]
**ConnectionState** | **string** | connected, disconnected, or unavailable | [default to null]
**ConnectionStateChangeAt** | [**time.Time**](time.Time.md) | time at which connectionStateChange last changed | [optional] [default to null]
**Health** | **interface{}** | healthy, poor, or degraded | [optional] [default to null]
**LeafToRoot** | **[]interface{}** | the list member is the parentId and freqBand,  from its pod to the root/gateway. | [default to null]
**MobileAppDeviceUuid** | **string** | Unique identifier for mobile devices that the mobile app generates and controls | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


