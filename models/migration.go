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

// Migration migration
//
// swagger:model Migration
type Migration struct {

	// id
	ID ObjectID `json:"id,omitempty"`

	// Timestamp of the migration occurrence
	// Required: true
	// Format: date-time
	MigratedAt *strfmt.DateTime `json:"migratedAt"`

	// From which cloud was this customer migrated
	// Required: true
	MigratedFrom *string `json:"migratedFrom"`

	// To which cloud was this customer migrated
	// Required: true
	MigratedTo *string `json:"migratedTo"`

	// Reason for migration occurrence
	// Required: true
	MigrationReason *string `json:"migrationReason"`
}

// Validate validates this migration
func (m *Migration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMigratedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMigratedFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMigratedTo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMigrationReason(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Migration) validateID(formats strfmt.Registry) error {
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

func (m *Migration) validateMigratedAt(formats strfmt.Registry) error {

	if err := validate.Required("migratedAt", "body", m.MigratedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("migratedAt", "body", "date-time", m.MigratedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Migration) validateMigratedFrom(formats strfmt.Registry) error {

	if err := validate.Required("migratedFrom", "body", m.MigratedFrom); err != nil {
		return err
	}

	return nil
}

func (m *Migration) validateMigratedTo(formats strfmt.Registry) error {

	if err := validate.Required("migratedTo", "body", m.MigratedTo); err != nil {
		return err
	}

	return nil
}

func (m *Migration) validateMigrationReason(formats strfmt.Registry) error {

	if err := validate.Required("migrationReason", "body", m.MigrationReason); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this migration based on the context it is used
func (m *Migration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Migration) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *Migration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Migration) UnmarshalBinary(b []byte) error {
	var res Migration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}