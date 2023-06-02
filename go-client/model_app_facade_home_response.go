/*
 * plume-customer
 *
 * Customer APIs for NOC, IOS, Android, QA scripts, and www.plume.com
 *
 * API version: 1.109.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type AppFacadeHomeResponse struct {
	// schema version # of the API response
	Version string `json:"_version,omitempty"`
	Customer *Customer `json:"customer,omitempty"`
	Capabilities *CapabilitiesResponse `json:"capabilities,omitempty"`
	GeoIp *GeoIp `json:"geoIp,omitempty"`
	NetworkStatus *NetworkStatus `json:"networkStatus,omitempty"`
	Optimization *Optimizations `json:"optimization,omitempty"`
	WifiNetwork *WifiNetwork `json:"wifiNetwork,omitempty"`
	Devices *Devices `json:"devices,omitempty"`
	Nodes *Node `json:"nodes,omitempty"`
	Summary *SummaryResponse `json:"summary,omitempty"`
	Alerts interface{} `json:"alerts,omitempty"`
	Id float64 `json:"id,omitempty"`
}
