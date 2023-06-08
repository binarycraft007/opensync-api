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

// NetworkStatus Location's network status information recevied from awlan.location.networkStatusChanged event
//
// swagger:model NetworkStatus
type NetworkStatus struct {

	// schema version # of a Mongo document
	// Required: true
	Version *string `json:"_version"`

	// Duration in seconds of previous status.
	DurationOfPreviousStatus float64 `json:"durationOfPreviousStatus,omitempty"`

	// Network status: [online | offline]
	Status string `json:"status,omitempty"`

	// Date time when status chenged.
	// Format: date-time
	StatusChangedAt strfmt.DateTime `json:"statusChangedAt,omitempty"`

	// updated at
	// Required: true
	// Format: date-time
	UpdatedAt *strfmt.DateTime `json:"updatedAt"`
}

// Validate validates this network status
func (m *NetworkStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatusChangedAt(formats); err != nil {
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

func (m *NetworkStatus) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("_version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

func (m *NetworkStatus) validateStatusChangedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.StatusChangedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("statusChangedAt", "body", "date-time", m.StatusChangedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NetworkStatus) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updatedAt", "body", m.UpdatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this network status based on context it is used
func (m *NetworkStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NetworkStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkStatus) UnmarshalBinary(b []byte) error {
	var res NetworkStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
