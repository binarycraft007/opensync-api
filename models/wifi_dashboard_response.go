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

// WifiDashboardResponse wifi dashboard response
//
// swagger:model WifiDashboardResponse
type WifiDashboardResponse struct {

	// access zones
	AccessZones []*WifiAccessZoneResponse `json:"accessZones"`

	// ssid
	// Required: true
	Ssid *string `json:"ssid"`
}

// Validate validates this wifi dashboard response
func (m *WifiDashboardResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessZones(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSsid(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WifiDashboardResponse) validateAccessZones(formats strfmt.Registry) error {
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

func (m *WifiDashboardResponse) validateSsid(formats strfmt.Registry) error {

	if err := validate.Required("ssid", "body", m.Ssid); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this wifi dashboard response based on the context it is used
func (m *WifiDashboardResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAccessZones(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WifiDashboardResponse) contextValidateAccessZones(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *WifiDashboardResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WifiDashboardResponse) UnmarshalBinary(b []byte) error {
	var res WifiDashboardResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}