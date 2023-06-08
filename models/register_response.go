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

// RegisterResponse register response
//
// swagger:model RegisterResponse
type RegisterResponse struct {

	// schema version # of a Mongo document
	// Required: true
	Version *string `json:"_version"`

	// account Id
	AccountID string `json:"accountId,omitempty"`

	// ISP registered accounts are anonymous and do not include real emails and passwords
	Anonymous *bool `json:"anonymous,omitempty"`

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// Fake email created on behalf of anonymous account.
	Email string `json:"email,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// ID of default location created during registration API call.
	LocationID string `json:"locationId,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this register response
func (m *RegisterResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RegisterResponse) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("_version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

func (m *RegisterResponse) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this register response based on context it is used
func (m *RegisterResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RegisterResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RegisterResponse) UnmarshalBinary(b []byte) error {
	var res RegisterResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
