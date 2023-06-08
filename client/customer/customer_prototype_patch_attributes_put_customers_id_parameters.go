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

	"github.com/binarycraft007/opensync-api/models"
)

// NewCustomerPrototypePatchAttributesPutCustomersIDParams creates a new CustomerPrototypePatchAttributesPutCustomersIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypePatchAttributesPutCustomersIDParams() *CustomerPrototypePatchAttributesPutCustomersIDParams {
	return &CustomerPrototypePatchAttributesPutCustomersIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypePatchAttributesPutCustomersIDParamsWithTimeout creates a new CustomerPrototypePatchAttributesPutCustomersIDParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypePatchAttributesPutCustomersIDParamsWithTimeout(timeout time.Duration) *CustomerPrototypePatchAttributesPutCustomersIDParams {
	return &CustomerPrototypePatchAttributesPutCustomersIDParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypePatchAttributesPutCustomersIDParamsWithContext creates a new CustomerPrototypePatchAttributesPutCustomersIDParams object
// with the ability to set a context for a request.
func NewCustomerPrototypePatchAttributesPutCustomersIDParamsWithContext(ctx context.Context) *CustomerPrototypePatchAttributesPutCustomersIDParams {
	return &CustomerPrototypePatchAttributesPutCustomersIDParams{
		Context: ctx,
	}
}

// NewCustomerPrototypePatchAttributesPutCustomersIDParamsWithHTTPClient creates a new CustomerPrototypePatchAttributesPutCustomersIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypePatchAttributesPutCustomersIDParamsWithHTTPClient(client *http.Client) *CustomerPrototypePatchAttributesPutCustomersIDParams {
	return &CustomerPrototypePatchAttributesPutCustomersIDParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypePatchAttributesPutCustomersIDParams contains all the parameters to send to the API endpoint

	for the customer prototype patch attributes put customers id operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypePatchAttributesPutCustomersIDParams struct {

	/* Data.

	   An object of model property name/value pairs
	*/
	Data *models.Customer

	/* ID.

	   Customer id

	   Format: JSON
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer prototype patch attributes put customers id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePatchAttributesPutCustomersIDParams) WithDefaults() *CustomerPrototypePatchAttributesPutCustomersIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype patch attributes put customers id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePatchAttributesPutCustomersIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype patch attributes put customers id params
func (o *CustomerPrototypePatchAttributesPutCustomersIDParams) WithTimeout(timeout time.Duration) *CustomerPrototypePatchAttributesPutCustomersIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype patch attributes put customers id params
func (o *CustomerPrototypePatchAttributesPutCustomersIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype patch attributes put customers id params
func (o *CustomerPrototypePatchAttributesPutCustomersIDParams) WithContext(ctx context.Context) *CustomerPrototypePatchAttributesPutCustomersIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype patch attributes put customers id params
func (o *CustomerPrototypePatchAttributesPutCustomersIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype patch attributes put customers id params
func (o *CustomerPrototypePatchAttributesPutCustomersIDParams) WithHTTPClient(client *http.Client) *CustomerPrototypePatchAttributesPutCustomersIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype patch attributes put customers id params
func (o *CustomerPrototypePatchAttributesPutCustomersIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the customer prototype patch attributes put customers id params
func (o *CustomerPrototypePatchAttributesPutCustomersIDParams) WithData(data *models.Customer) *CustomerPrototypePatchAttributesPutCustomersIDParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the customer prototype patch attributes put customers id params
func (o *CustomerPrototypePatchAttributesPutCustomersIDParams) SetData(data *models.Customer) {
	o.Data = data
}

// WithID adds the id to the customer prototype patch attributes put customers id params
func (o *CustomerPrototypePatchAttributesPutCustomersIDParams) WithID(id string) *CustomerPrototypePatchAttributesPutCustomersIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype patch attributes put customers id params
func (o *CustomerPrototypePatchAttributesPutCustomersIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypePatchAttributesPutCustomersIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
