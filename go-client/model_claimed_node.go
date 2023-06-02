/*
 * plume-customer
 *
 * Customer APIs for NOC, IOS, Android, QA scripts, and www.plume.com
 *
 * API version: 1.109.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ClaimedNode struct {
	// unique serialNumber
	SerialNumber string `json:"serialNumber"`
	// always returns the value of 'serialNumber'
	Id string `json:"id,omitempty"`
	// a cool nickname
	Nickname string `json:"nickname"`
	// defaultName is for code, display if empty nickname
	DefaultName string `json:"defaultName"`
	// connectionState
	Status string `json:"status"`
}
