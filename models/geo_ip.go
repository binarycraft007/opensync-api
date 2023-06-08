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

// GeoIP Location information recevied from geo IP service maxmind.com
//
// swagger:model GeoIp
type GeoIP struct {

	// ISP of location based from IP.
	ISP *bool `json:"ISP,omitempty"`

	// schema version # of a Mongo document
	// Required: true
	Version *string `json:"_version"`

	// city
	City string `json:"city,omitempty"`

	// country
	Country string `json:"country,omitempty"`

	// ip
	IP string `json:"ip,omitempty"`

	// latitude of location based from IP.
	Latitude float64 `json:"latitude,omitempty"`

	// longitude of location based from IP.
	Longitude float64 `json:"longitude,omitempty"`

	// province
	Province string `json:"province,omitempty"`

	// state
	State string `json:"state,omitempty"`

	// timezone of location based from IP.
	Timezone string `json:"timezone,omitempty"`

	// updated at
	// Required: true
	// Format: date-time
	UpdatedAt *strfmt.DateTime `json:"updatedAt"`
}

// Validate validates this geo Ip
func (m *GeoIP) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVersion(formats); err != nil {
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

func (m *GeoIP) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("_version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

func (m *GeoIP) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updatedAt", "body", m.UpdatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this geo Ip based on context it is used
func (m *GeoIP) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GeoIP) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GeoIP) UnmarshalBinary(b []byte) error {
	var res GeoIP
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
