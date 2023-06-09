// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Qos qos
//
// swagger:model Qos
type Qos struct {

	// id
	ID ObjectID `json:"id,omitempty"`

	// prioritization
	Prioritization *Prioritization `json:"prioritization,omitempty"`
}

// Validate validates this qos
func (m *Qos) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrioritization(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Qos) validateID(formats strfmt.Registry) error {
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

func (m *Qos) validatePrioritization(formats strfmt.Registry) error {
	if swag.IsZero(m.Prioritization) { // not required
		return nil
	}

	if m.Prioritization != nil {
		if err := m.Prioritization.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("prioritization")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("prioritization")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this qos based on the context it is used
func (m *Qos) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePrioritization(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Qos) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Qos) contextValidatePrioritization(ctx context.Context, formats strfmt.Registry) error {

	if m.Prioritization != nil {
		if err := m.Prioritization.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("prioritization")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("prioritization")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Qos) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Qos) UnmarshalBinary(b []byte) error {
	var res Qos
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
