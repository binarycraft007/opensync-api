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

// LocationSpeedTest location speed test
//
// swagger:model LocationSpeedTest
type LocationSpeedTest struct {

	// schema version # of a Mongo document
	// Required: true
	Version *string `json:"_version"`

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// if network is idle, ISP speedtests will be run on a gateway node every 3 hours
	Enable bool `json:"enable,omitempty"`

	// if network is idle, ISP speedtests will be run on every node twice a day
	EnableAllNodes bool `json:"enableAllNodes,omitempty"`

	// The first ISP Speed Test on a gateway node had speed results above a minimum threshold
	SpeedCapable bool `json:"speedCapable,omitempty"`

	// updated at
	// Required: true
	// Format: date-time
	UpdatedAt *strfmt.DateTime `json:"updatedAt"`
}

// Validate validates this location speed test
func (m *LocationSpeedTest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
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

func (m *LocationSpeedTest) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("_version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

func (m *LocationSpeedTest) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *LocationSpeedTest) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updatedAt", "body", m.UpdatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this location speed test based on context it is used
func (m *LocationSpeedTest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LocationSpeedTest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LocationSpeedTest) UnmarshalBinary(b []byte) error {
	var res LocationSpeedTest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}