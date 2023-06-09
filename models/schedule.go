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

// Schedule schedule
//
// swagger:model Schedule
type Schedule struct {

	// schema version # of a Mongo document
	// Required: true
	Version *string `json:"_version"`

	// days of week
	// Required: true
	DaysOfWeek []float64 `json:"daysOfWeek"`

	// end time
	EndTime string `json:"endTime,omitempty"`

	// id
	ID float64 `json:"id,omitempty"`

	// start time
	// Required: true
	StartTime *string `json:"startTime"`
}

// Validate validates this schedule
func (m *Schedule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDaysOfWeek(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Schedule) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("_version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

func (m *Schedule) validateDaysOfWeek(formats strfmt.Registry) error {

	if err := validate.Required("daysOfWeek", "body", m.DaysOfWeek); err != nil {
		return err
	}

	return nil
}

func (m *Schedule) validateStartTime(formats strfmt.Registry) error {

	if err := validate.Required("startTime", "body", m.StartTime); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this schedule based on context it is used
func (m *Schedule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Schedule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Schedule) UnmarshalBinary(b []byte) error {
	var res Schedule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
