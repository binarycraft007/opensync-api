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

// NewCustomerPrototypeGetSecurityRealizedStatesParams creates a new CustomerPrototypeGetSecurityRealizedStatesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypeGetSecurityRealizedStatesParams() *CustomerPrototypeGetSecurityRealizedStatesParams {
	return &CustomerPrototypeGetSecurityRealizedStatesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypeGetSecurityRealizedStatesParamsWithTimeout creates a new CustomerPrototypeGetSecurityRealizedStatesParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypeGetSecurityRealizedStatesParamsWithTimeout(timeout time.Duration) *CustomerPrototypeGetSecurityRealizedStatesParams {
	return &CustomerPrototypeGetSecurityRealizedStatesParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypeGetSecurityRealizedStatesParamsWithContext creates a new CustomerPrototypeGetSecurityRealizedStatesParams object
// with the ability to set a context for a request.
func NewCustomerPrototypeGetSecurityRealizedStatesParamsWithContext(ctx context.Context) *CustomerPrototypeGetSecurityRealizedStatesParams {
	return &CustomerPrototypeGetSecurityRealizedStatesParams{
		Context: ctx,
	}
}

// NewCustomerPrototypeGetSecurityRealizedStatesParamsWithHTTPClient creates a new CustomerPrototypeGetSecurityRealizedStatesParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypeGetSecurityRealizedStatesParamsWithHTTPClient(client *http.Client) *CustomerPrototypeGetSecurityRealizedStatesParams {
	return &CustomerPrototypeGetSecurityRealizedStatesParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypeGetSecurityRealizedStatesParams contains all the parameters to send to the API endpoint

	for the customer prototype get security realized states operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypeGetSecurityRealizedStatesParams struct {

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

// WithDefaults hydrates default values in the customer prototype get security realized states params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeGetSecurityRealizedStatesParams) WithDefaults() *CustomerPrototypeGetSecurityRealizedStatesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype get security realized states params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeGetSecurityRealizedStatesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype get security realized states params
func (o *CustomerPrototypeGetSecurityRealizedStatesParams) WithTimeout(timeout time.Duration) *CustomerPrototypeGetSecurityRealizedStatesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype get security realized states params
func (o *CustomerPrototypeGetSecurityRealizedStatesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype get security realized states params
func (o *CustomerPrototypeGetSecurityRealizedStatesParams) WithContext(ctx context.Context) *CustomerPrototypeGetSecurityRealizedStatesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype get security realized states params
func (o *CustomerPrototypeGetSecurityRealizedStatesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype get security realized states params
func (o *CustomerPrototypeGetSecurityRealizedStatesParams) WithHTTPClient(client *http.Client) *CustomerPrototypeGetSecurityRealizedStatesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype get security realized states params
func (o *CustomerPrototypeGetSecurityRealizedStatesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the customer prototype get security realized states params
func (o *CustomerPrototypeGetSecurityRealizedStatesParams) WithID(id string) *CustomerPrototypeGetSecurityRealizedStatesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype get security realized states params
func (o *CustomerPrototypeGetSecurityRealizedStatesParams) SetID(id string) {
	o.ID = id
}

// WithLocationID adds the locationID to the customer prototype get security realized states params
func (o *CustomerPrototypeGetSecurityRealizedStatesParams) WithLocationID(locationID string) *CustomerPrototypeGetSecurityRealizedStatesParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype get security realized states params
func (o *CustomerPrototypeGetSecurityRealizedStatesParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypeGetSecurityRealizedStatesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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