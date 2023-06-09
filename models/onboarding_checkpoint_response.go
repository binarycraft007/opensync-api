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

// OnboardingCheckpointResponse onboarding checkpoint response
//
// swagger:model OnboardingCheckpointResponse
type OnboardingCheckpointResponse struct {

	// the last recorded checkpoint.
	// Required: true
	Checkpoint *OnboardingCheckpoint `json:"checkpoint"`
}

// Validate validates this onboarding checkpoint response
func (m *OnboardingCheckpointResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCheckpoint(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OnboardingCheckpointResponse) validateCheckpoint(formats strfmt.Registry) error {

	if err := validate.Required("checkpoint", "body", m.Checkpoint); err != nil {
		return err
	}

	if m.Checkpoint != nil {
		if err := m.Checkpoint.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("checkpoint")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("checkpoint")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this onboarding checkpoint response based on the context it is used
func (m *OnboardingCheckpointResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCheckpoint(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OnboardingCheckpointResponse) contextValidateCheckpoint(ctx context.Context, formats strfmt.Registry) error {

	if m.Checkpoint != nil {
		if err := m.Checkpoint.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("checkpoint")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("checkpoint")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OnboardingCheckpointResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OnboardingCheckpointResponse) UnmarshalBinary(b []byte) error {
	var res OnboardingCheckpointResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
