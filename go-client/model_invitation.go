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

type Invitation struct {
	Id string `json:"id"`
	KeyId float64 `json:"keyId,omitempty"`
	Url string `json:"url,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// used by Mongo to auto delete the document
	ExpiresAt time.Time `json:"expiresAt"`
	WifiNetworkId *ObjectId `json:"wifiNetworkId,omitempty"`
	LocationId *ObjectId `json:"locationId,omitempty"`
}