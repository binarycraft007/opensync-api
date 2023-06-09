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

// NewCustomerPrototypeDeleteAppPrioritizationLocationConfigParams creates a new CustomerPrototypeDeleteAppPrioritizationLocationConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypeDeleteAppPrioritizationLocationConfigParams() *CustomerPrototypeDeleteAppPrioritizationLocationConfigParams {
	return &CustomerPrototypeDeleteAppPrioritizationLocationConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypeDeleteAppPrioritizationLocationConfigParamsWithTimeout creates a new CustomerPrototypeDeleteAppPrioritizationLocationConfigParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypeDeleteAppPrioritizationLocationConfigParamsWithTimeout(timeout time.Duration) *CustomerPrototypeDeleteAppPrioritizationLocationConfigParams {
	return &CustomerPrototypeDeleteAppPrioritizationLocationConfigParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypeDeleteAppPrioritizationLocationConfigParamsWithContext creates a new CustomerPrototypeDeleteAppPrioritizationLocationConfigParams object
// with the ability to set a context for a request.
func NewCustomerPrototypeDeleteAppPrioritizationLocationConfigParamsWithContext(ctx context.Context) *CustomerPrototypeDeleteAppPrioritizationLocationConfigParams {
	return &CustomerPrototypeDeleteAppPrioritizationLocationConfigParams{
		Context: ctx,
	}
}

// NewCustomerPrototypeDeleteAppPrioritizationLocationConfigParamsWithHTTPClient creates a new CustomerPrototypeDeleteAppPrioritizationLocationConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypeDeleteAppPrioritizationLocationConfigParamsWithHTTPClient(client *http.Client) *CustomerPrototypeDeleteAppPrioritizationLocationConfigParams {
	return &CustomerPrototypeDeleteAppPrioritizationLocationConfigParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypeDeleteAppPrioritizationLocationConfigParams contains all the parameters to send to the API endpoint

	for the customer prototype delete app prioritization location config operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypeDeleteAppPrioritizationLocationConfigParams struct {

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

// WithDefaults hydrates default values in the customer prototype delete app prioritization location config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeDeleteAppPrioritizationLocationConfigParams) WithDefaults() *CustomerPrototypeDeleteAppPrioritizationLocationConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype delete app prioritization location config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeDeleteAppPrioritizationLocationConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype delete app prioritization location config params
func (o *CustomerPrototypeDeleteAppPrioritizationLocationConfigParams) WithTimeout(timeout time.Duration) *CustomerPrototypeDeleteAppPrioritizationLocationConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype delete app prioritization location config params
func (o *CustomerPrototypeDeleteAppPrioritizationLocationConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype delete app prioritization location config params
func (o *CustomerPrototypeDeleteAppPrioritizationLocationConfigParams) WithContext(ctx context.Context) *CustomerPrototypeDeleteAppPrioritizationLocationConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype delete app prioritization location config params
func (o *CustomerPrototypeDeleteAppPrioritizationLocationConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype delete app prioritization location config params
func (o *CustomerPrototypeDeleteAppPrioritizationLocationConfigParams) WithHTTPClient(client *http.Client) *CustomerPrototypeDeleteAppPrioritizationLocationConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype delete app prioritization location config params
func (o *CustomerPrototypeDeleteAppPrioritizationLocationConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the customer prototype delete app prioritization location config params
func (o *CustomerPrototypeDeleteAppPrioritizationLocationConfigParams) WithID(id string) *CustomerPrototypeDeleteAppPrioritizationLocationConfigParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype delete app prioritization location config params
func (o *CustomerPrototypeDeleteAppPrioritizationLocationConfigParams) SetID(id string) {
	o.ID = id
}

// WithLocationID adds the locationID to the customer prototype delete app prioritization location config params
func (o *CustomerPrototypeDeleteAppPrioritizationLocationConfigParams) WithLocationID(locationID string) *CustomerPrototypeDeleteAppPrioritizationLocationConfigParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype delete app prioritization location config params
func (o *CustomerPrototypeDeleteAppPrioritizationLocationConfigParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypeDeleteAppPrioritizationLocationConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
