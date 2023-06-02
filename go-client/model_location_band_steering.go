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

type LocationBandSteering struct {
	// schema version # of a Mongo document
	Version string `json:"_version"`
	// deprecated
	Auto bool `json:"auto,omitempty"`
	// auto | enable | disable
	Mode string `json:"mode,omitempty"`
	UpdatedAt time.Time `json:"updatedAt"`
}
