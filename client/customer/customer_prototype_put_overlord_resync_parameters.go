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

// NewCustomerPrototypePutOverlordResyncParams creates a new CustomerPrototypePutOverlordResyncParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypePutOverlordResyncParams() *CustomerPrototypePutOverlordResyncParams {
	return &CustomerPrototypePutOverlordResyncParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypePutOverlordResyncParamsWithTimeout creates a new CustomerPrototypePutOverlordResyncParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypePutOverlordResyncParamsWithTimeout(timeout time.Duration) *CustomerPrototypePutOverlordResyncParams {
	return &CustomerPrototypePutOverlordResyncParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypePutOverlordResyncParamsWithContext creates a new CustomerPrototypePutOverlordResyncParams object
// with the ability to set a context for a request.
func NewCustomerPrototypePutOverlordResyncParamsWithContext(ctx context.Context) *CustomerPrototypePutOverlordResyncParams {
	return &CustomerPrototypePutOverlordResyncParams{
		Context: ctx,
	}
}

// NewCustomerPrototypePutOverlordResyncParamsWithHTTPClient creates a new CustomerPrototypePutOverlordResyncParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypePutOverlordResyncParamsWithHTTPClient(client *http.Client) *CustomerPrototypePutOverlordResyncParams {
	return &CustomerPrototypePutOverlordResyncParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypePutOverlordResyncParams contains all the parameters to send to the API endpoint

	for the customer prototype put overlord resync operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypePutOverlordResyncParams struct {

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

// WithDefaults hydrates default values in the customer prototype put overlord resync params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePutOverlordResyncParams) WithDefaults() *CustomerPrototypePutOverlordResyncParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype put overlord resync params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePutOverlordResyncParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype put overlord resync params
func (o *CustomerPrototypePutOverlordResyncParams) WithTimeout(timeout time.Duration) *CustomerPrototypePutOverlordResyncParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype put overlord resync params
func (o *CustomerPrototypePutOverlordResyncParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype put overlord resync params
func (o *CustomerPrototypePutOverlordResyncParams) WithContext(ctx context.Context) *CustomerPrototypePutOverlordResyncParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype put overlord resync params
func (o *CustomerPrototypePutOverlordResyncParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype put overlord resync params
func (o *CustomerPrototypePutOverlordResyncParams) WithHTTPClient(client *http.Client) *CustomerPrototypePutOverlordResyncParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype put overlord resync params
func (o *CustomerPrototypePutOverlordResyncParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the customer prototype put overlord resync params
func (o *CustomerPrototypePutOverlordResyncParams) WithID(id string) *CustomerPrototypePutOverlordResyncParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype put overlord resync params
func (o *CustomerPrototypePutOverlordResyncParams) SetID(id string) {
	o.ID = id
}

// WithLocationID adds the locationID to the customer prototype put overlord resync params
func (o *CustomerPrototypePutOverlordResyncParams) WithLocationID(locationID string) *CustomerPrototypePutOverlordResyncParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype put overlord resync params
func (o *CustomerPrototypePutOverlordResyncParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypePutOverlordResyncParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
