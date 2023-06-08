// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Invitation invitation
//
// swagger:model Invitation
type Invitation struct {

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// used by Mongo to auto delete the document
	// Required: true
	// Format: date-time
	ExpiresAt *strfmt.DateTime `json:"expiresAt"`

	// id
	// Required: true
	ID *string `json:"id"`

	// key Id
	KeyID float64 `json:"keyId,omitempty"`

	// location Id
	LocationID ObjectID `json:"locationId,omitempty"`

	// url
	URL string `json:"url,omitempty"`

	// wifi network Id
	WifiNetworkID ObjectID `json:"wifiNetworkId,omitempty"`
}

// Validate validates this invitation
func (m *Invitation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpiresAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWifiNetworkID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Invitation) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Invitation) validateExpiresAt(formats strfmt.Registry) error {

	if err := validate.Required("expiresAt", "body", m.ExpiresAt); err != nil {
		return err
	}

	if err := validate.FormatOf("expiresAt", "body", "date-time", m.ExpiresAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Invitation) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Invitation) validateLocationID(formats strfmt.Registry) error {
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

func (m *Invitation) validateWifiNetworkID(formats strfmt.Registry) error {
	if swag.IsZero(m.WifiNetworkID) { // not required
		return nil
	}

	if err := m.WifiNetworkID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("wifiNetworkId")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("wifiNetworkId")
		}
		return err
	}

	return nil
}

// ContextValidate validate this invitation based on the context it is used
func (m *Invitation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLocationID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWifiNetworkID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Invitation) contextValidateLocationID(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Invitation) contextValidateWifiNetworkID(ctx context.Context, formats strfmt.Registry) error {

	if err := m.WifiNetworkID.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("wifiNetworkId")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("wifiNetworkId")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Invitation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Invitation) UnmarshalBinary(b []byte) error {
	var res Invitation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}