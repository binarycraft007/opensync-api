// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NodeVendor node vendor
//
// swagger:model NodeVendor
type NodeVendor struct {

	// factory
	Factory string `json:"factory,omitempty"`

	// manufacturer
	Manufacturer string `json:"manufacturer,omitempty"`

	// mfg date
	MfgDate string `json:"mfgDate,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// part number
	PartNumber string `json:"partNumber,omitempty"`
}

// Validate validates this node vendor
func (m *NodeVendor) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this node vendor based on context it is used
func (m *NodeVendor) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NodeVendor) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NodeVendor) UnmarshalBinary(b []byte) error {
	var res NodeVendor
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
