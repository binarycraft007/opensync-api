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

// NewCustomerPrototypeGetHomeSecurityParams creates a new CustomerPrototypeGetHomeSecurityParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypeGetHomeSecurityParams() *CustomerPrototypeGetHomeSecurityParams {
	return &CustomerPrototypeGetHomeSecurityParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypeGetHomeSecurityParamsWithTimeout creates a new CustomerPrototypeGetHomeSecurityParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypeGetHomeSecurityParamsWithTimeout(timeout time.Duration) *CustomerPrototypeGetHomeSecurityParams {
	return &CustomerPrototypeGetHomeSecurityParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypeGetHomeSecurityParamsWithContext creates a new CustomerPrototypeGetHomeSecurityParams object
// with the ability to set a context for a request.
func NewCustomerPrototypeGetHomeSecurityParamsWithContext(ctx context.Context) *CustomerPrototypeGetHomeSecurityParams {
	return &CustomerPrototypeGetHomeSecurityParams{
		Context: ctx,
	}
}

// NewCustomerPrototypeGetHomeSecurityParamsWithHTTPClient creates a new CustomerPrototypeGetHomeSecurityParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypeGetHomeSecurityParamsWithHTTPClient(client *http.Client) *CustomerPrototypeGetHomeSecurityParams {
	return &CustomerPrototypeGetHomeSecurityParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypeGetHomeSecurityParams contains all the parameters to send to the API endpoint

	for the customer prototype get home security operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypeGetHomeSecurityParams struct {

	/* ID.

	   Customer id

	   Format: JSON
	*/
	ID string

	// LocationID.
	//
	// Format: JSON
	LocationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer prototype get home security params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeGetHomeSecurityParams) WithDefaults() *CustomerPrototypeGetHomeSecurityParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype get home security params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeGetHomeSecurityParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype get home security params
func (o *CustomerPrototypeGetHomeSecurityParams) WithTimeout(timeout time.Duration) *CustomerPrototypeGetHomeSecurityParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype get home security params
func (o *CustomerPrototypeGetHomeSecurityParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype get home security params
func (o *CustomerPrototypeGetHomeSecurityParams) WithContext(ctx context.Context) *CustomerPrototypeGetHomeSecurityParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype get home security params
func (o *CustomerPrototypeGetHomeSecurityParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype get home security params
func (o *CustomerPrototypeGetHomeSecurityParams) WithHTTPClient(client *http.Client) *CustomerPrototypeGetHomeSecurityParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype get home security params
func (o *CustomerPrototypeGetHomeSecurityParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the customer prototype get home security params
func (o *CustomerPrototypeGetHomeSecurityParams) WithID(id string) *CustomerPrototypeGetHomeSecurityParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype get home security params
func (o *CustomerPrototypeGetHomeSecurityParams) SetID(id string) {
	o.ID = id
}

// WithLocationID adds the locationID to the customer prototype get home security params
func (o *CustomerPrototypeGetHomeSecurityParams) WithLocationID(locationID string) *CustomerPrototypeGetHomeSecurityParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype get home security params
func (o *CustomerPrototypeGetHomeSecurityParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypeGetHomeSecurityParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
