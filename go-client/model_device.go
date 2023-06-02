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

type Device struct {
	Id *ObjectId `json:"id,omitempty"`
	Mac string `json:"mac"`
	// unique id of the WifiNetwork.keys[x] that the device is connected to or last connected to
	KeyId float64 `json:"keyId,omitempty"`
	// host name from the device/user settings
	HostName string `json:"hostName,omitempty"`
	// device model
	Model string `json:"model,omitempty"`
	// wifi, ethernet, or moca
	Medium string `json:"medium,omitempty"`
	// schema version # of a Mongo document
	Version string `json:"_version"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	OnlineSince time.Time `json:"onlineSince,omitempty"`
	OfflineSince time.Time `json:"offlineSince,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	ConnectionState string `json:"connectionState,omitempty"`
	// time at which connectionStateChange last changed
	ConnectionStateChangeAt time.Time `json:"connectionStateChangeAt,omitempty"`
	VapType string `json:"vapType,omitempty"`
	NetworkId string `json:"networkId,omitempty"`
	Favorite bool `json:"favorite,omitempty"`
	// Time at which user removed device from device list
	HiddenAt time.Time `json:"hiddenAt,omitempty"`
	// device's appTime config
	AppTime interface{} `json:"appTime,omitempty"`
	// Unique identifier for mobile devices that the mobile app generates and controls
	MobileAppDeviceUuid string `json:"mobileAppDeviceUuid,omitempty"`
	LocationId *ObjectId `json:"locationId,omitempty"`
	BandSteering *DeviceBandSteering `json:"bandSteering,omitempty"`
	ClientSteering *DeviceClientSteering `json:"clientSteering,omitempty"`
	Qos *Qos `json:"qos,omitempty"`
}