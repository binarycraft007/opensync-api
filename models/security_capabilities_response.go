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

// SecurityCapabilitiesResponse security capabilities response
//
// swagger:model SecurityCapabilitiesResponse
type SecurityCapabilitiesResponse struct {

	// feature's availability for a location
	// Required: true
	DNSCategoryBlocking *Capability `json:"dnsCategoryBlocking"`
}

// Validate validates this security capabilities response
func (m *SecurityCapabilitiesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDNSCategoryBlocking(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityCapabilitiesResponse) validateDNSCategoryBlocking(formats strfmt.Registry) error {

	if err := validate.Required("dnsCategoryBlocking", "body", m.DNSCategoryBlocking); err != nil {
		return err
	}

	if m.DNSCategoryBlocking != nil {
		if err := m.DNSCategoryBlocking.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dnsCategoryBlocking")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dnsCategoryBlocking")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this security capabilities response based on the context it is used
func (m *SecurityCapabilitiesResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDNSCategoryBlocking(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityCapabilitiesResponse) contextValidateDNSCategoryBlocking(ctx context.Context, formats strfmt.Registry) error {

	if m.DNSCategoryBlocking != nil {
		if err := m.DNSCategoryBlocking.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dnsCategoryBlocking")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dnsCategoryBlocking")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SecurityCapabilitiesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecurityCapabilitiesResponse) UnmarshalBinary(b []byte) error {
	var res SecurityCapabilitiesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
