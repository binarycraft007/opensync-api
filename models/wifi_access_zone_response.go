// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WifiAccessZoneResponse wifi access zone response
//
// swagger:model WifiAccessZoneResponse
type WifiAccessZoneResponse struct {

	// accessible devices
	AccessibleDevices []*DeviceResponse `json:"accessibleDevices"`

	// name of accessZone
	Description string `json:"description,omitempty"`

	// accessZoneId
	// Required: true
	ID *float64 `json:"id"`

	// keys
	Keys []*WifiNetworkKey `json:"keys"`

	// home | guests | internetAccessOnly
	Type string `json:"type,omitempty"`
}

// Validate validates this wifi access zone response
func (m *WifiAccessZoneResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessibleDevices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKeys(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WifiAccessZoneResponse) validateAccessibleDevices(formats strfmt.Registry) error {
	if swag.IsZero(m.AccessibleDevices) { // not required
		return nil
	}

	for i := 0; i < len(m.AccessibleDevices); i++ {
		if swag.IsZero(m.AccessibleDevices[i]) { // not required
			continue
		}

		if m.AccessibleDevices[i] != nil {
			if err := m.AccessibleDevices[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("accessibleDevices" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("accessibleDevices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WifiAccessZoneResponse) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *WifiAccessZoneResponse) validateKeys(formats strfmt.Registry) error {
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

// ContextValidate validate this wifi access zone response based on the context it is used
func (m *WifiAccessZoneResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAccessibleDevices(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateKeys(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WifiAccessZoneResponse) contextValidateAccessibleDevices(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AccessibleDevices); i++ {

		if m.AccessibleDevices[i] != nil {
			if err := m.AccessibleDevices[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("accessibleDevices" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("accessibleDevices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WifiAccessZoneResponse) contextValidateKeys(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *WifiAccessZoneResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WifiAccessZoneResponse) UnmarshalBinary(b []byte) error {
	var res WifiAccessZoneResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
