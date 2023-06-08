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

// NetworkAccessDeviceGroup network access device group
//
// swagger:model NetworkAccessDeviceGroup
type NetworkAccessDeviceGroup struct {

	// devices
	Devices []XAny `json:"devices"`

	// group Id
	// Required: true
	GroupID *string `json:"groupId"`

	// name
	Name string `json:"name,omitempty"`

	// network Id
	NetworkID string `json:"networkId,omitempty"`
}

// Validate validates this network access device group
func (m *NetworkAccessDeviceGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGroupID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkAccessDeviceGroup) validateGroupID(formats strfmt.Registry) error {

	if err := validate.Required("groupId", "body", m.GroupID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this network access device group based on context it is used
func (m *NetworkAccessDeviceGroup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NetworkAccessDeviceGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkAccessDeviceGroup) UnmarshalBinary(b []byte) error {
	var res NetworkAccessDeviceGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
