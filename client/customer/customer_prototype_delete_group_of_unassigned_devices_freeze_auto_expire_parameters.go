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

// NewCustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireParams creates a new CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireParams() *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireParams {
	return &CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireParamsWithTimeout creates a new CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireParamsWithTimeout(timeout time.Duration) *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireParams {
	return &CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireParamsWithContext creates a new CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireParams object
// with the ability to set a context for a request.
func NewCustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireParamsWithContext(ctx context.Context) *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireParams {
	return &CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireParams{
		Context: ctx,
	}
}

// NewCustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireParamsWithHTTPClient creates a new CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireParamsWithHTTPClient(client *http.Client) *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireParams {
	return &CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireParams contains all the parameters to send to the API endpoint

	for the customer prototype delete group of unassigned devices freeze auto expire operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireParams struct {

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

// WithDefaults hydrates default values in the customer prototype delete group of unassigned devices freeze auto expire params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireParams) WithDefaults() *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype delete group of unassigned devices freeze auto expire params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype delete group of unassigned devices freeze auto expire params
func (o *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireParams) WithTimeout(timeout time.Duration) *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype delete group of unassigned devices freeze auto expire params
func (o *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype delete group of unassigned devices freeze auto expire params
func (o *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireParams) WithContext(ctx context.Context) *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype delete group of unassigned devices freeze auto expire params
func (o *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype delete group of unassigned devices freeze auto expire params
func (o *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireParams) WithHTTPClient(client *http.Client) *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype delete group of unassigned devices freeze auto expire params
func (o *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the customer prototype delete group of unassigned devices freeze auto expire params
func (o *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireParams) WithID(id string) *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype delete group of unassigned devices freeze auto expire params
func (o *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireParams) SetID(id string) {
	o.ID = id
}

// WithLocationID adds the locationID to the customer prototype delete group of unassigned devices freeze auto expire params
func (o *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireParams) WithLocationID(locationID string) *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype delete group of unassigned devices freeze auto expire params
func (o *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
