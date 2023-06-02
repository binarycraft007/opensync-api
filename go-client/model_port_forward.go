/*
 * plume-customer
 *
 * Customer APIs for NOC, IOS, Android, QA scripts, and www.plume.com
 *
 * API version: 1.109.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type PortForward struct {
	ExternalPort float64 `json:"externalPort"`
	InternalPort float64 `json:"internalPort"`
	Protocol string `json:"protocol"`
	Name string `json:"name,omitempty"`
}