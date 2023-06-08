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

// NewCustomerPrototypeOverlordDeleteStatsConfigParams creates a new CustomerPrototypeOverlordDeleteStatsConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypeOverlordDeleteStatsConfigParams() *CustomerPrototypeOverlordDeleteStatsConfigParams {
	return &CustomerPrototypeOverlordDeleteStatsConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypeOverlordDeleteStatsConfigParamsWithTimeout creates a new CustomerPrototypeOverlordDeleteStatsConfigParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypeOverlordDeleteStatsConfigParamsWithTimeout(timeout time.Duration) *CustomerPrototypeOverlordDeleteStatsConfigParams {
	return &CustomerPrototypeOverlordDeleteStatsConfigParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypeOverlordDeleteStatsConfigParamsWithContext creates a new CustomerPrototypeOverlordDeleteStatsConfigParams object
// with the ability to set a context for a request.
func NewCustomerPrototypeOverlordDeleteStatsConfigParamsWithContext(ctx context.Context) *CustomerPrototypeOverlordDeleteStatsConfigParams {
	return &CustomerPrototypeOverlordDeleteStatsConfigParams{
		Context: ctx,
	}
}

// NewCustomerPrototypeOverlordDeleteStatsConfigParamsWithHTTPClient creates a new CustomerPrototypeOverlordDeleteStatsConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypeOverlordDeleteStatsConfigParamsWithHTTPClient(client *http.Client) *CustomerPrototypeOverlordDeleteStatsConfigParams {
	return &CustomerPrototypeOverlordDeleteStatsConfigParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypeOverlordDeleteStatsConfigParams contains all the parameters to send to the API endpoint

	for the customer prototype overlord delete stats config operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypeOverlordDeleteStatsConfigParams struct {

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

// WithDefaults hydrates default values in the customer prototype overlord delete stats config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeOverlordDeleteStatsConfigParams) WithDefaults() *CustomerPrototypeOverlordDeleteStatsConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype overlord delete stats config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeOverlordDeleteStatsConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype overlord delete stats config params
func (o *CustomerPrototypeOverlordDeleteStatsConfigParams) WithTimeout(timeout time.Duration) *CustomerPrototypeOverlordDeleteStatsConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype overlord delete stats config params
func (o *CustomerPrototypeOverlordDeleteStatsConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype overlord delete stats config params
func (o *CustomerPrototypeOverlordDeleteStatsConfigParams) WithContext(ctx context.Context) *CustomerPrototypeOverlordDeleteStatsConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype overlord delete stats config params
func (o *CustomerPrototypeOverlordDeleteStatsConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype overlord delete stats config params
func (o *CustomerPrototypeOverlordDeleteStatsConfigParams) WithHTTPClient(client *http.Client) *CustomerPrototypeOverlordDeleteStatsConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype overlord delete stats config params
func (o *CustomerPrototypeOverlordDeleteStatsConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the customer prototype overlord delete stats config params
func (o *CustomerPrototypeOverlordDeleteStatsConfigParams) WithID(id string) *CustomerPrototypeOverlordDeleteStatsConfigParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype overlord delete stats config params
func (o *CustomerPrototypeOverlordDeleteStatsConfigParams) SetID(id string) {
	o.ID = id
}

// WithLocationID adds the locationID to the customer prototype overlord delete stats config params
func (o *CustomerPrototypeOverlordDeleteStatsConfigParams) WithLocationID(locationID string) *CustomerPrototypeOverlordDeleteStatsConfigParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype overlord delete stats config params
func (o *CustomerPrototypeOverlordDeleteStatsConfigParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypeOverlordDeleteStatsConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
