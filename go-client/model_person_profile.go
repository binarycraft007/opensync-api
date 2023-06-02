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

type PersonProfile struct {
	// schema version # of a Mongo document
	Version string `json:"_version"`
	// employee only for now
	Type_ string `json:"type"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	Id *ObjectId `json:"id,omitempty"`
}
