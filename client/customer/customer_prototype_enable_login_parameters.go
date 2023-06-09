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

// NewCustomerPrototypeEnableLoginParams creates a new CustomerPrototypeEnableLoginParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypeEnableLoginParams() *CustomerPrototypeEnableLoginParams {
	return &CustomerPrototypeEnableLoginParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypeEnableLoginParamsWithTimeout creates a new CustomerPrototypeEnableLoginParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypeEnableLoginParamsWithTimeout(timeout time.Duration) *CustomerPrototypeEnableLoginParams {
	return &CustomerPrototypeEnableLoginParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypeEnableLoginParamsWithContext creates a new CustomerPrototypeEnableLoginParams object
// with the ability to set a context for a request.
func NewCustomerPrototypeEnableLoginParamsWithContext(ctx context.Context) *CustomerPrototypeEnableLoginParams {
	return &CustomerPrototypeEnableLoginParams{
		Context: ctx,
	}
}

// NewCustomerPrototypeEnableLoginParamsWithHTTPClient creates a new CustomerPrototypeEnableLoginParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypeEnableLoginParamsWithHTTPClient(client *http.Client) *CustomerPrototypeEnableLoginParams {
	return &CustomerPrototypeEnableLoginParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypeEnableLoginParams contains all the parameters to send to the API endpoint

	for the customer prototype enable login operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypeEnableLoginParams struct {

	/* ID.

	   Customer id

	   Format: JSON
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer prototype enable login params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeEnableLoginParams) WithDefaults() *CustomerPrototypeEnableLoginParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype enable login params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeEnableLoginParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype enable login params
func (o *CustomerPrototypeEnableLoginParams) WithTimeout(timeout time.Duration) *CustomerPrototypeEnableLoginParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype enable login params
func (o *CustomerPrototypeEnableLoginParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype enable login params
func (o *CustomerPrototypeEnableLoginParams) WithContext(ctx context.Context) *CustomerPrototypeEnableLoginParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype enable login params
func (o *CustomerPrototypeEnableLoginParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype enable login params
func (o *CustomerPrototypeEnableLoginParams) WithHTTPClient(client *http.Client) *CustomerPrototypeEnableLoginParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype enable login params
func (o *CustomerPrototypeEnableLoginParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the customer prototype enable login params
func (o *CustomerPrototypeEnableLoginParams) WithID(id string) *CustomerPrototypeEnableLoginParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype enable login params
func (o *CustomerPrototypeEnableLoginParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypeEnableLoginParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
