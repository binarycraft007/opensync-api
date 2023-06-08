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

// NewCustomerPrototypeMigrationStatusParams creates a new CustomerPrototypeMigrationStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypeMigrationStatusParams() *CustomerPrototypeMigrationStatusParams {
	return &CustomerPrototypeMigrationStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypeMigrationStatusParamsWithTimeout creates a new CustomerPrototypeMigrationStatusParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypeMigrationStatusParamsWithTimeout(timeout time.Duration) *CustomerPrototypeMigrationStatusParams {
	return &CustomerPrototypeMigrationStatusParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypeMigrationStatusParamsWithContext creates a new CustomerPrototypeMigrationStatusParams object
// with the ability to set a context for a request.
func NewCustomerPrototypeMigrationStatusParamsWithContext(ctx context.Context) *CustomerPrototypeMigrationStatusParams {
	return &CustomerPrototypeMigrationStatusParams{
		Context: ctx,
	}
}

// NewCustomerPrototypeMigrationStatusParamsWithHTTPClient creates a new CustomerPrototypeMigrationStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypeMigrationStatusParamsWithHTTPClient(client *http.Client) *CustomerPrototypeMigrationStatusParams {
	return &CustomerPrototypeMigrationStatusParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypeMigrationStatusParams contains all the parameters to send to the API endpoint

	for the customer prototype migration status operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypeMigrationStatusParams struct {

	/* ID.

	   Customer id

	   Format: JSON
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer prototype migration status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeMigrationStatusParams) WithDefaults() *CustomerPrototypeMigrationStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype migration status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeMigrationStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype migration status params
func (o *CustomerPrototypeMigrationStatusParams) WithTimeout(timeout time.Duration) *CustomerPrototypeMigrationStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype migration status params
func (o *CustomerPrototypeMigrationStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype migration status params
func (o *CustomerPrototypeMigrationStatusParams) WithContext(ctx context.Context) *CustomerPrototypeMigrationStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype migration status params
func (o *CustomerPrototypeMigrationStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype migration status params
func (o *CustomerPrototypeMigrationStatusParams) WithHTTPClient(client *http.Client) *CustomerPrototypeMigrationStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype migration status params
func (o *CustomerPrototypeMigrationStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the customer prototype migration status params
func (o *CustomerPrototypeMigrationStatusParams) WithID(id string) *CustomerPrototypeMigrationStatusParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype migration status params
func (o *CustomerPrototypeMigrationStatusParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypeMigrationStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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