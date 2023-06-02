/*
 * plume-customer
 *
 * Customer APIs for NOC, IOS, Android, QA scripts, and www.plume.com
 *
 * API version: 1.109.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"time"
)

type LocationMonitorMode struct {
	// schema version # of a Mongo document
	Version string `json:"_version"`
	Enable bool `json:"enable,omitempty"`
	UpdatedAt time.Time `json:"updatedAt"`
}