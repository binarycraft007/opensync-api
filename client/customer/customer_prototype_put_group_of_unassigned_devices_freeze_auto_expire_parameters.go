// Code generated by go-swagger; DO NOT EDIT.

package customer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewCustomerPrototypePutGroupOfUnassignedDevicesFreezeAutoExpireParams creates a new CustomerPrototypePutGroupOfUnassignedDevicesFreezeAutoExpireParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypePutGroupOfUnassignedDevicesFreezeAutoExpireParams() *CustomerPrototypePutGroupOfUnassignedDevicesFreezeAutoExpireParams {
	return &CustomerPrototypePutGroupOfUnassignedDevicesFreezeAutoExpireParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypePutGroupOfUnassignedDevicesFreezeAutoExpireParamsWithTimeout creates a new CustomerPrototypePutGroupOfUnassignedDevicesFreezeAutoExpireParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypePutGroupOfUnassignedDevicesFreezeAutoExpireParamsWithTimeout(timeout time.Duration) *CustomerPrototypePutGroupOfUnassignedDevicesFreezeAutoExpireParams {
	return &CustomerPrototypePutGroupOfUnassignedDevicesFreezeAutoExpireParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypePutGroupOfUnassignedDevicesFreezeAutoExpireParamsWithContext creates a new CustomerPrototypePutGroupOfUnassignedDevicesFreezeAutoExpireParams object
// with the ability to set a context for a request.
func NewCustomerPrototypePutGroupOfUnassignedDevicesFreezeAutoExpireParamsWithContext(ctx context.Context) *CustomerPrototypePutGroupOfUnassignedDevicesFreezeAutoExpireParams {
	return &CustomerPrototypePutGroupOfUnassignedDevicesFreezeAutoExpireParams{
		Context: ctx,
	}
}

// NewCustomerPrototypePutGroupOfUnassignedDevicesFreezeAutoExpireParamsWithHTTPClient creates a new CustomerPrototypePutGroupOfUnassignedDevicesFreezeAutoExpireParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypePutGroupOfUnassignedDevicesFreezeAutoExpireParamsWithHTTPClient(client *http.Client) *CustomerPrototypePutGroupOfUnassignedDevicesFreezeAutoExpireParams {
	return &CustomerPrototypePutGroupOfUnassignedDevicesFreezeAutoExpireParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypePutGroupOfUnassignedDevicesFreezeAutoExpireParams contains all the parameters to send to the API endpoint

	for the customer prototype put group of unassigned devices freeze auto expire operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypePutGroupOfUnassignedDevicesFreezeAutoExpireParams struct {

	// ExpiresAt.
	ExpiresAt string

	/* ID.

	   Customer id

	   Format: JSON
	*/
	ID string

	// LocationID.
	LocationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer prototype put group of unassigned devices freeze auto expire params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePutGroupOfUnassignedDevicesFreezeAutoExpireParams) WithDefaults() *CustomerPrototypePutGroupOfUnassignedDevicesFreezeAutoExpireParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype put group of unassigned devices freeze auto expire params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePutGroupOfUnassignedDevicesFreezeAutoExpireParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype put group of unassigned devices freeze auto expire params
func (o *CustomerPrototypePutGroupOfUnassignedDevicesFreezeAutoExpireParams) WithTimeout(timeout time.Duration) *CustomerPrototypePutGroupOfUnassignedDevicesFreezeAutoExpireParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype put group of unassigned devices freeze auto expire params
func (o *CustomerPrototypePutGroupOfUnassignedDevicesFreezeAutoExpireParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype put group of unassigned devices freeze auto expire params
func (o *CustomerPrototypePutGroupOfUnassignedDevicesFreezeAutoExpireParams) WithContext(ctx context.Context) *CustomerPrototypePutGroupOfUnassignedDevicesFreezeAutoExpireParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype put group of unassigned devices freeze auto expire params
func (o *CustomerPrototypePutGroupOfUnassignedDevicesFreezeAutoExpireParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype put group of unassigned devices freeze auto expire params
func (o *CustomerPrototypePutGroupOfUnassignedDevicesFreezeAutoExpireParams) WithHTTPClient(client *http.Client) *CustomerPrototypePutGroupOfUnassignedDevicesFreezeAutoExpireParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype put group of unassigned devices freeze auto expire params
func (o *CustomerPrototypePutGroupOfUnassignedDevicesFreezeAutoExpireParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExpiresAt adds the expiresAt to the customer prototype put group of unassigned devices freeze auto expire params
func (o *CustomerPrototypePutGroupOfUnassignedDevicesFreezeAutoExpireParams) WithExpiresAt(expiresAt string) *CustomerPrototypePutGroupOfUnassignedDevicesFreezeAutoExpireParams {
	o.SetExpiresAt(expiresAt)
	return o
}

// SetExpiresAt adds the expiresAt to the customer prototype put group of unassigned devices freeze auto expire params
func (o *CustomerPrototypePutGroupOfUnassignedDevicesFreezeAutoExpireParams) SetExpiresAt(expiresAt string) {
	o.ExpiresAt = expiresAt
}

// WithID adds the id to the customer prototype put group of unassigned devices freeze auto expire params
func (o *CustomerPrototypePutGroupOfUnassignedDevicesFreezeAutoExpireParams) WithID(id string) *CustomerPrototypePutGroupOfUnassignedDevicesFreezeAutoExpireParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype put group of unassigned devices freeze auto expire params
func (o *CustomerPrototypePutGroupOfUnassignedDevicesFreezeAutoExpireParams) SetID(id string) {
	o.ID = id
}

// WithLocationID adds the locationID to the customer prototype put group of unassigned devices freeze auto expire params
func (o *CustomerPrototypePutGroupOfUnassignedDevicesFreezeAutoExpireParams) WithLocationID(locationID string) *CustomerPrototypePutGroupOfUnassignedDevicesFreezeAutoExpireParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype put group of unassigned devices freeze auto expire params
func (o *CustomerPrototypePutGroupOfUnassignedDevicesFreezeAutoExpireParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypePutGroupOfUnassignedDevicesFreezeAutoExpireParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// form param expiresAt
	frExpiresAt := o.ExpiresAt
	fExpiresAt := frExpiresAt
	if fExpiresAt != "" {
		if err := r.SetFormParam("expiresAt", fExpiresAt); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param locationId
	if err := r.SetPathParam("locationId", o.LocationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}