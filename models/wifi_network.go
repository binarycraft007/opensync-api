// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WifiNetwork wifi network
//
// swagger:model WifiNetwork
type WifiNetwork struct {

	// schema version # of a Mongo document
	// Required: true
	Version *string `json:"_version"`

	// access zones
	AccessZones []*WifiAccessZone `json:"accessZones"`

	// location's default security policy
	AppliesToAllDevicesSecurityPolicy *WifiSecurityPolicy `json:"appliesToAllDevicesSecurityPolicy,omitempty"`

	// content filter of the master key[masterKeyIndex]: kids || teenagers || adBlocking || spamPhishMalware || adultAndSensitive || workAppropriate
	Content []string `json:"content"`

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// location's default security policy
	DefaultSecurityPolicy *WifiSecurityPolicy `json:"defaultSecurityPolicy,omitempty"`

	// devices
	Devices []*WifiSecurityPolicy `json:"devices"`

	// the list of the device mac addresses
	DevicesVisibleToGuests []string `json:"devicesVisibleToGuests"`

	// disables the primary network VAP
	DisableDefaultServiceNetwork *bool `json:"disableDefaultServiceNetwork,omitempty"`

	// enables the VAPs and radios and backhauls
	Enabled *bool `json:"enabled,omitempty"`

	// encryption
	// Enum: [OPEN WEP WPA-EAP WPA-PSK]
	Encryption string `json:"encryption,omitempty"`

	// encryption key
	// Required: true
	EncryptionKey *string `json:"encryptionKey"`

	// encryption mode
	// Enum: [1 2 mixed 64 128]
	EncryptionMode string `json:"encryptionMode,omitempty"`

	// auto || enable || disable
	FastTransition *string `json:"fastTransition,omitempty"`

	// auto || enable || disable
	GroupRekey *string `json:"groupRekey,omitempty"`

	// id
	ID ObjectID `json:"id,omitempty"`

	// keys
	Keys []*WifiNetworkKey `json:"keys"`

	// location Id
	LocationID ObjectID `json:"locationId,omitempty"`

	// keys[] array index of the password to be used as the original home password for backwards compatibility for locations that don't use home pass (yet)
	// Required: true
	MasterKeyIndex *float64 `json:"masterKeyIndex"`

	// auto || 11b || 11g || 11n
	MinWifiMode24 *string `json:"minWifiMode24,omitempty"`

	// persons
	Persons []*WifiSecurityPolicy `json:"persons"`

	// private mode
	PrivateMode *bool `json:"privateMode,omitempty"`

	// Privacy Mode Visibility with timestamp in Frontline
	// Format: date-time
	PrivateModeUpdatedAt strfmt.DateTime `json:"privateModeUpdatedAt,omitempty"`

	// radius server Ip
	RadiusServerIP string `json:"radiusServerIp,omitempty"`

	// radius server port
	RadiusServerPort float64 `json:"radiusServerPort,omitempty"`

	// ssid
	// Required: true
	Ssid *string `json:"ssid"`

	// ssid broadcast
	SsidBroadcast bool `json:"ssidBroadcast,omitempty"`

	// uapsd
	Uapsd *bool `json:"uapsd,omitempty"`

	// updated at
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updatedAt,omitempty"`

	// open || psk-mixed || sae-mixed || psk2 || sae
	WpaMode *string `json:"wpaMode,omitempty"`

	// Hold off setting wpaMode on devices for WPA3 onboarding for a set amount of time.
	WpaModeDeferredExpiresAt string `json:"wpaModeDeferredExpiresAt,omitempty"`
}

// Validate validates this wifi network
func (m *WifiNetwork) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccessZones(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAppliesToAllDevicesSecurityPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefaultSecurityPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDevices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEncryption(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEncryptionKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEncryptionMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKeys(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMasterKeyIndex(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePersons(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrivateModeUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSsid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WifiNetwork) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("_version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

func (m *WifiNetwork) validateAccessZones(formats strfmt.Registry) error {
	if swag.IsZero(m.AccessZones) { // not required
		return nil
	}

	for i := 0; i < len(m.AccessZones); i++ {
		if swag.IsZero(m.AccessZones[i]) { // not required
			continue
		}

		if m.AccessZones[i] != nil {
			if err := m.AccessZones[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("accessZones" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("accessZones" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WifiNetwork) validateAppliesToAllDevicesSecurityPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.AppliesToAllDevicesSecurityPolicy) { // not required
		return nil
	}

	if m.AppliesToAllDevicesSecurityPolicy != nil {
		if err := m.AppliesToAllDevicesSecurityPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("appliesToAllDevicesSecurityPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("appliesToAllDevicesSecurityPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *WifiNetwork) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WifiNetwork) validateDefaultSecurityPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.DefaultSecurityPolicy) { // not required
		return nil
	}

	if m.DefaultSecurityPolicy != nil {
		if err := m.DefaultSecurityPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("defaultSecurityPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("defaultSecurityPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *WifiNetwork) validateDevices(formats strfmt.Registry) error {
	if swag.IsZero(m.Devices) { // not required
		return nil
	}

	for i := 0; i < len(m.Devices); i++ {
		if swag.IsZero(m.Devices[i]) { // not required
			continue
		}

		if m.Devices[i] != nil {
			if err := m.Devices[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("devices" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("devices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var wifiNetworkTypeEncryptionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["OPEN","WEP","WPA-EAP","WPA-PSK"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		wifiNetworkTypeEncryptionPropEnum = append(wifiNetworkTypeEncryptionPropEnum, v)
	}
}

const (

	// WifiNetworkEncryptionOPEN captures enum value "OPEN"
	WifiNetworkEncryptionOPEN string = "OPEN"

	// WifiNetworkEncryptionWEP captures enum value "WEP"
	WifiNetworkEncryptionWEP string = "WEP"

	// WifiNetworkEncryptionWPADashEAP captures enum value "WPA-EAP"
	WifiNetworkEncryptionWPADashEAP string = "WPA-EAP"

	// WifiNetworkEncryptionWPADashPSK captures enum value "WPA-PSK"
	WifiNetworkEncryptionWPADashPSK string = "WPA-PSK"
)

// prop value enum
func (m *WifiNetwork) validateEncryptionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, wifiNetworkTypeEncryptionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WifiNetwork) validateEncryption(formats strfmt.Registry) error {
	if swag.IsZero(m.Encryption) { // not required
		return nil
	}

	// value enum
	if err := m.validateEncryptionEnum("encryption", "body", m.Encryption); err != nil {
		return err
	}

	return nil
}

func (m *WifiNetwork) validateEncryptionKey(formats strfmt.Registry) error {

	if err := validate.Required("encryptionKey", "body", m.EncryptionKey); err != nil {
		return err
	}

	return nil
}

var wifiNetworkTypeEncryptionModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["1","2","mixed","64","128"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		wifiNetworkTypeEncryptionModePropEnum = append(wifiNetworkTypeEncryptionModePropEnum, v)
	}
}

const (

	// WifiNetworkEncryptionModeNr1 captures enum value "1"
	WifiNetworkEncryptionModeNr1 string = "1"

	// WifiNetworkEncryptionModeNr2 captures enum value "2"
	WifiNetworkEncryptionModeNr2 string = "2"

	// WifiNetworkEncryptionModeMixed captures enum value "mixed"
	WifiNetworkEncryptionModeMixed string = "mixed"

	// WifiNetworkEncryptionModeNr64 captures enum value "64"
	WifiNetworkEncryptionModeNr64 string = "64"

	// WifiNetworkEncryptionModeNr128 captures enum value "128"
	WifiNetworkEncryptionModeNr128 string = "128"
)

// prop value enum
func (m *WifiNetwork) validateEncryptionModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, wifiNetworkTypeEncryptionModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WifiNetwork) validateEncryptionMode(formats strfmt.Registry) error {
	if swag.IsZero(m.EncryptionMode) { // not required
		return nil
	}

	// value enum
	if err := m.validateEncryptionModeEnum("encryptionMode", "body", m.EncryptionMode); err != nil {
		return err
	}

	return nil
}

func (m *WifiNetwork) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := m.ID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("id")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("id")
		}
		return err
	}

	return nil
}

func (m *WifiNetwork) validateKeys(formats strfmt.Registry) error {
	if swag.IsZero(m.Keys) { // not required
		return nil
	}

	for i := 0; i < len(m.Keys); i++ {
		if swag.IsZero(m.Keys[i]) { // not required
			continue
		}

		if m.Keys[i] != nil {
			if err := m.Keys[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("keys" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("keys" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WifiNetwork) validateLocationID(formats strfmt.Registry) error {
	if swag.IsZero(m.LocationID) { // not required
		return nil
	}

	if err := m.LocationID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("locationId")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("locationId")
		}
		return err
	}

	return nil
}

func (m *WifiNetwork) validateMasterKeyIndex(formats strfmt.Registry) error {

	if err := validate.Required("masterKeyIndex", "body", m.MasterKeyIndex); err != nil {
		return err
	}

	return nil
}

func (m *WifiNetwork) validatePersons(formats strfmt.Registry) error {
	if swag.IsZero(m.Persons) { // not required
		return nil
	}

	for i := 0; i < len(m.Persons); i++ {
		if swag.IsZero(m.Persons[i]) { // not required
			continue
		}

		if m.Persons[i] != nil {
			if err := m.Persons[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("persons" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("persons" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WifiNetwork) validatePrivateModeUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.PrivateModeUpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("privateModeUpdatedAt", "body", "date-time", m.PrivateModeUpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WifiNetwork) validateSsid(formats strfmt.Registry) error {

	if err := validate.Required("ssid", "body", m.Ssid); err != nil {
		return err
	}

	return nil
}

func (m *WifiNetwork) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this wifi network based on the context it is used
func (m *WifiNetwork) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAccessZones(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAppliesToAllDevicesSecurityPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDefaultSecurityPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDevices(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateKeys(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLocationID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePersons(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WifiNetwork) contextValidateAccessZones(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AccessZones); i++ {

		if m.AccessZones[i] != nil {
			if err := m.AccessZones[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("accessZones" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("accessZones" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WifiNetwork) contextValidateAppliesToAllDevicesSecurityPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.AppliesToAllDevicesSecurityPolicy != nil {
		if err := m.AppliesToAllDevicesSecurityPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("appliesToAllDevicesSecurityPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("appliesToAllDevicesSecurityPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *WifiNetwork) contextValidateDefaultSecurityPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.DefaultSecurityPolicy != nil {
		if err := m.DefaultSecurityPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("defaultSecurityPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("defaultSecurityPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *WifiNetwork) contextValidateDevices(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Devices); i++ {

		if m.Devices[i] != nil {
			if err := m.Devices[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("devices" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("devices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WifiNetwork) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ID.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("id")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("id")
		}
		return err
	}

	return nil
}

func (m *WifiNetwork) contextValidateKeys(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Keys); i++ {

		if m.Keys[i] != nil {
			if err := m.Keys[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("keys" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("keys" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WifiNetwork) contextValidateLocationID(ctx context.Context, formats strfmt.Registry) error {

	if err := m.LocationID.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("locationId")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("locationId")
		}
		return err
	}

	return nil
}

func (m *WifiNetwork) contextValidatePersons(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Persons); i++ {

		if m.Persons[i] != nil {
			if err := m.Persons[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("persons" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("persons" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *WifiNetwork) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WifiNetwork) UnmarshalBinary(b []byte) error {
	var res WifiNetwork
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
