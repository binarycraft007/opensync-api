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

// NewCustomerPrototypeGetLocationSecurityPolicyHourlyCountsParams creates a new CustomerPrototypeGetLocationSecurityPolicyHourlyCountsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypeGetLocationSecurityPolicyHourlyCountsParams() *CustomerPrototypeGetLocationSecurityPolicyHourlyCountsParams {
	return &CustomerPrototypeGetLocationSecurityPolicyHourlyCountsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypeGetLocationSecurityPolicyHourlyCountsParamsWithTimeout creates a new CustomerPrototypeGetLocationSecurityPolicyHourlyCountsParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypeGetLocationSecurityPolicyHourlyCountsParamsWithTimeout(timeout time.Duration) *CustomerPrototypeGetLocationSecurityPolicyHourlyCountsParams {
	return &CustomerPrototypeGetLocationSecurityPolicyHourlyCountsParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypeGetLocationSecurityPolicyHourlyCountsParamsWithContext creates a new CustomerPrototypeGetLocationSecurityPolicyHourlyCountsParams object
// with the ability to set a context for a request.
func NewCustomerPrototypeGetLocationSecurityPolicyHourlyCountsParamsWithContext(ctx context.Context) *CustomerPrototypeGetLocationSecurityPolicyHourlyCountsParams {
	return &CustomerPrototypeGetLocationSecurityPolicyHourlyCountsParams{
		Context: ctx,
	}
}

// NewCustomerPrototypeGetLocationSecurityPolicyHourlyCountsParamsWithHTTPClient creates a new CustomerPrototypeGetLocationSecurityPolicyHourlyCountsParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypeGetLocationSecurityPolicyHourlyCountsParamsWithHTTPClient(client *http.Client) *CustomerPrototypeGetLocationSecurityPolicyHourlyCountsParams {
	return &CustomerPrototypeGetLocationSecurityPolicyHourlyCountsParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypeGetLocationSecurityPolicyHourlyCountsParams contains all the parameters to send to the API endpoint

	for the customer prototype get location security policy hourly counts operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypeGetLocationSecurityPolicyHourlyCountsParams struct {

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

// WithDefaults hydrates default values in the customer prototype get location security policy hourly counts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeGetLocationSecurityPolicyHourlyCountsParams) WithDefaults() *CustomerPrototypeGetLocationSecurityPolicyHourlyCountsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype get location security policy hourly counts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeGetLocationSecurityPolicyHourlyCountsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype get location security policy hourly counts params
func (o *CustomerPrototypeGetLocationSecurityPolicyHourlyCountsParams) WithTimeout(timeout time.Duration) *CustomerPrototypeGetLocationSecurityPolicyHourlyCountsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype get location security policy hourly counts params
func (o *CustomerPrototypeGetLocationSecurityPolicyHourlyCountsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype get location security policy hourly counts params
func (o *CustomerPrototypeGetLocationSecurityPolicyHourlyCountsParams) WithContext(ctx context.Context) *CustomerPrototypeGetLocationSecurityPolicyHourlyCountsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype get location security policy hourly counts params
func (o *CustomerPrototypeGetLocationSecurityPolicyHourlyCountsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype get location security policy hourly counts params
func (o *CustomerPrototypeGetLocationSecurityPolicyHourlyCountsParams) WithHTTPClient(client *http.Client) *CustomerPrototypeGetLocationSecurityPolicyHourlyCountsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype get location security policy hourly counts params
func (o *CustomerPrototypeGetLocationSecurityPolicyHourlyCountsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the customer prototype get location security policy hourly counts params
func (o *CustomerPrototypeGetLocationSecurityPolicyHourlyCountsParams) WithID(id string) *CustomerPrototypeGetLocationSecurityPolicyHourlyCountsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype get location security policy hourly counts params
func (o *CustomerPrototypeGetLocationSecurityPolicyHourlyCountsParams) SetID(id string) {
	o.ID = id
}

// WithLocationID adds the locationID to the customer prototype get location security policy hourly counts params
func (o *CustomerPrototypeGetLocationSecurityPolicyHourlyCountsParams) WithLocationID(locationID string) *CustomerPrototypeGetLocationSecurityPolicyHourlyCountsParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype get location security policy hourly counts params
func (o *CustomerPrototypeGetLocationSecurityPolicyHourlyCountsParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypeGetLocationSecurityPolicyHourlyCountsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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