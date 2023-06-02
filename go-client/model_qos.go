/*
 * plume-customer
 *
 * Customer APIs for NOC, IOS, Android, QA scripts, and www.plume.com
 *
 * API version: 1.109.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Qos struct {
	Id *ObjectId `json:"id,omitempty"`
	Prioritization *Prioritization `json:"prioritization,omitempty"`
}
