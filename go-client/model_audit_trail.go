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

type AuditTrail struct {
	Id *ObjectId `json:"id,omitempty"`
	// The customer id
	CustomerId string `json:"customerId"`
	// The location id
	LocationId string `json:"locationId,omitempty"`
	// The partner id
	PartnerId string `json:"partnerId,omitempty"`
	// The user who made the change
	Author string `json:"author"`
	// The date and time the change was made
	CreatedAt time.Time `json:"createdAt"`
	// The date and time the change will expire
	ExpiresAt time.Time `json:"expiresAt"`
	// The type of event that occurred
	Event string `json:"event"`
	// The details of the event
	Details interface{} `json:"details"`
	// The request id
	XRequestId string `json:"xRequestId,omitempty"`
}