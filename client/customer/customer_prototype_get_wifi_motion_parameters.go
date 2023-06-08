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

// NewCustomerPrototypeGetWifiMotionParams creates a new CustomerPrototypeGetWifiMotionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypeGetWifiMotionParams() *CustomerPrototypeGetWifiMotionParams {
	return &CustomerPrototypeGetWifiMotionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypeGetWifiMotionParamsWithTimeout creates a new CustomerPrototypeGetWifiMotionParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypeGetWifiMotionParamsWithTimeout(timeout time.Duration) *CustomerPrototypeGetWifiMotionParams {
	return &CustomerPrototypeGetWifiMotionParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypeGetWifiMotionParamsWithContext creates a new CustomerPrototypeGetWifiMotionParams object
// with the ability to set a context for a request.
func NewCustomerPrototypeGetWifiMotionParamsWithContext(ctx context.Context) *CustomerPrototypeGetWifiMotionParams {
	return &CustomerPrototypeGetWifiMotionParams{
		Context: ctx,
	}
}

// NewCustomerPrototypeGetWifiMotionParamsWithHTTPClient creates a new CustomerPrototypeGetWifiMotionParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypeGetWifiMotionParamsWithHTTPClient(client *http.Client) *CustomerPrototypeGetWifiMotionParams {
	return &CustomerPrototypeGetWifiMotionParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypeGetWifiMotionParams contains all the parameters to send to the API endpoint

	for the customer prototype get wifi motion operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypeGetWifiMotionParams struct {

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

// WithDefaults hydrates default values in the customer prototype get wifi motion params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeGetWifiMotionParams) WithDefaults() *CustomerPrototypeGetWifiMotionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype get wifi motion params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeGetWifiMotionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype get wifi motion params
func (o *CustomerPrototypeGetWifiMotionParams) WithTimeout(timeout time.Duration) *CustomerPrototypeGetWifiMotionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype get wifi motion params
func (o *CustomerPrototypeGetWifiMotionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype get wifi motion params
func (o *CustomerPrototypeGetWifiMotionParams) WithContext(ctx context.Context) *CustomerPrototypeGetWifiMotionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype get wifi motion params
func (o *CustomerPrototypeGetWifiMotionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype get wifi motion params
func (o *CustomerPrototypeGetWifiMotionParams) WithHTTPClient(client *http.Client) *CustomerPrototypeGetWifiMotionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype get wifi motion params
func (o *CustomerPrototypeGetWifiMotionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the customer prototype get wifi motion params
func (o *CustomerPrototypeGetWifiMotionParams) WithID(id string) *CustomerPrototypeGetWifiMotionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype get wifi motion params
func (o *CustomerPrototypeGetWifiMotionParams) SetID(id string) {
	o.ID = id
}

// WithLocationID adds the locationID to the customer prototype get wifi motion params
func (o *CustomerPrototypeGetWifiMotionParams) WithLocationID(locationID string) *CustomerPrototypeGetWifiMotionParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype get wifi motion params
func (o *CustomerPrototypeGetWifiMotionParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypeGetWifiMotionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
