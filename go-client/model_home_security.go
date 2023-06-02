/*
 * plume-customer
 *
 * Customer APIs for NOC, IOS, Android, QA scripts, and www.plume.com
 *
 * API version: 1.109.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type HomeSecurity struct {
	LiveMotionEnabled bool `json:"liveMotionEnabled,omitempty"`
	MotionEventsEnabled bool `json:"motionEventsEnabled,omitempty"`
	Cooldown float64 `json:"cooldown,omitempty"`
	PetMode string `json:"petMode,omitempty"`
	Sensitivity string `json:"sensitivity,omitempty"`
	Id float64 `json:"id,omitempty"`
}
