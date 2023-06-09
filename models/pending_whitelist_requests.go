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

// PendingWhitelistRequests Requests for whitelisting of access
//
// swagger:model PendingWhitelistRequests
type PendingWhitelistRequests struct {

	// created at
	// Required: true
	// Format: date-time
	CreatedAt *strfmt.DateTime `json:"createdAt"`

	// location Id
	// Required: true
	LocationID *ObjectID `json:"locationId"`

	// PersonId of the requester
	PersonID string `json:"personId,omitempty"`

	// request Id
	RequestID ObjectID `json:"requestId,omitempty"`

	// CustomerId of the requester
	RequesterCustomerID string `json:"requesterCustomerId,omitempty"`

	// person's appTime config
	Status *string `json:"status,omitempty"`

	// mac address of the Person's primary device
	Type string `json:"type,omitempty"`

	// updated at
	// Required: true
	// Format: date-time
	UpdatedAt *strfmt.DateTime `json:"updatedAt"`

	// notification when person is away from home
	Value *bool `json:"value,omitempty"`
}

// Validate validates this pending whitelist requests
func (m *PendingWhitelistRequests) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestID(formats); err != nil {
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

func (m *PendingWhitelistRequests) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("createdAt", "body", m.CreatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PendingWhitelistRequests) validateLocationID(formats strfmt.Registry) error {

	if err := validate.Required("locationId", "body", m.LocationID); err != nil {
		return err
	}

	if err := validate.Required("locationId", "body", m.LocationID); err != nil {
		return err
	}

	if m.LocationID != nil {
		if err := m.LocationID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("locationId")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("locationId")
			}
			return err
		}
	}

	return nil
}

func (m *PendingWhitelistRequests) validateRequestID(formats strfmt.Registry) error {
	if swag.IsZero(m.RequestID) { // not required
		return nil
	}

	if err := m.RequestID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("requestId")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("requestId")
		}
		return err
	}

	return nil
}

func (m *PendingWhitelistRequests) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updatedAt", "body", m.UpdatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this pending whitelist requests based on the context it is used
func (m *PendingWhitelistRequests) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLocationID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRequestID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PendingWhitelistRequests) contextValidateLocationID(ctx context.Context, formats strfmt.Registry) error {

	if m.LocationID != nil {
		if err := m.LocationID.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("locationId")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("locationId")
			}
			return err
		}
	}

	return nil
}

func (m *PendingWhitelistRequests) contextValidateRequestID(ctx context.Context, formats strfmt.Registry) error {

	if err := m.RequestID.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("requestId")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("requestId")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PendingWhitelistRequests) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PendingWhitelistRequests) UnmarshalBinary(b []byte) error {
	var res PendingWhitelistRequests
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
