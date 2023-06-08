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

// NewCustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeSuspendParams creates a new CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeSuspendParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeSuspendParams() *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeSuspendParams {
	return &CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeSuspendParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeSuspendParamsWithTimeout creates a new CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeSuspendParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeSuspendParamsWithTimeout(timeout time.Duration) *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeSuspendParams {
	return &CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeSuspendParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeSuspendParamsWithContext creates a new CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeSuspendParams object
// with the ability to set a context for a request.
func NewCustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeSuspendParamsWithContext(ctx context.Context) *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeSuspendParams {
	return &CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeSuspendParams{
		Context: ctx,
	}
}

// NewCustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeSuspendParamsWithHTTPClient creates a new CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeSuspendParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeSuspendParamsWithHTTPClient(client *http.Client) *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeSuspendParams {
	return &CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeSuspendParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeSuspendParams contains all the parameters to send to the API endpoint

	for the customer prototype delete group of unassigned devices freeze suspend operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeSuspendParams struct {

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

// WithDefaults hydrates default values in the customer prototype delete group of unassigned devices freeze suspend params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeSuspendParams) WithDefaults() *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeSuspendParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype delete group of unassigned devices freeze suspend params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeSuspendParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype delete group of unassigned devices freeze suspend params
func (o *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeSuspendParams) WithTimeout(timeout time.Duration) *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeSuspendParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype delete group of unassigned devices freeze suspend params
func (o *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeSuspendParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype delete group of unassigned devices freeze suspend params
func (o *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeSuspendParams) WithContext(ctx context.Context) *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeSuspendParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype delete group of unassigned devices freeze suspend params
func (o *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeSuspendParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype delete group of unassigned devices freeze suspend params
func (o *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeSuspendParams) WithHTTPClient(client *http.Client) *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeSuspendParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype delete group of unassigned devices freeze suspend params
func (o *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeSuspendParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the customer prototype delete group of unassigned devices freeze suspend params
func (o *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeSuspendParams) WithID(id string) *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeSuspendParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype delete group of unassigned devices freeze suspend params
func (o *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeSuspendParams) SetID(id string) {
	o.ID = id
}

// WithLocationID adds the locationID to the customer prototype delete group of unassigned devices freeze suspend params
func (o *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeSuspendParams) WithLocationID(locationID string) *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeSuspendParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype delete group of unassigned devices freeze suspend params
func (o *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeSuspendParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeSuspendParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
