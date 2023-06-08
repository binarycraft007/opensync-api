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

// NetworkAccessGroupShare network access group share
//
// swagger:model NetworkAccessGroupShare
type NetworkAccessGroupShare struct {

	// group ids
	// Required: true
	GroupIds []XAny `json:"groupIds"`
}

// Validate validates this network access group share
func (m *NetworkAccessGroupShare) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGroupIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkAccessGroupShare) validateGroupIds(formats strfmt.Registry) error {

	if err := validate.Required("groupIds", "body", m.GroupIds); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this network access group share based on context it is used
func (m *NetworkAccessGroupShare) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NetworkAccessGroupShare) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkAccessGroupShare) UnmarshalBinary(b []byte) error {
	var res NetworkAccessGroupShare
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
