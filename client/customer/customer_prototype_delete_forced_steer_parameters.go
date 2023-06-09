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

// NewCustomerPrototypeDeleteForcedSteerParams creates a new CustomerPrototypeDeleteForcedSteerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypeDeleteForcedSteerParams() *CustomerPrototypeDeleteForcedSteerParams {
	return &CustomerPrototypeDeleteForcedSteerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypeDeleteForcedSteerParamsWithTimeout creates a new CustomerPrototypeDeleteForcedSteerParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypeDeleteForcedSteerParamsWithTimeout(timeout time.Duration) *CustomerPrototypeDeleteForcedSteerParams {
	return &CustomerPrototypeDeleteForcedSteerParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypeDeleteForcedSteerParamsWithContext creates a new CustomerPrototypeDeleteForcedSteerParams object
// with the ability to set a context for a request.
func NewCustomerPrototypeDeleteForcedSteerParamsWithContext(ctx context.Context) *CustomerPrototypeDeleteForcedSteerParams {
	return &CustomerPrototypeDeleteForcedSteerParams{
		Context: ctx,
	}
}

// NewCustomerPrototypeDeleteForcedSteerParamsWithHTTPClient creates a new CustomerPrototypeDeleteForcedSteerParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypeDeleteForcedSteerParamsWithHTTPClient(client *http.Client) *CustomerPrototypeDeleteForcedSteerParams {
	return &CustomerPrototypeDeleteForcedSteerParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypeDeleteForcedSteerParams contains all the parameters to send to the API endpoint

	for the customer prototype delete forced steer operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypeDeleteForcedSteerParams struct {

	/* ID.

	   Customer id

	   Format: JSON
	*/
	ID string

	/* LocationID.

	   locationId
	*/
	LocationID string

	/* Mac.

	   MAC address of the target device. Must have been online in the last 31 days.
	*/
	Mac string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer prototype delete forced steer params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeDeleteForcedSteerParams) WithDefaults() *CustomerPrototypeDeleteForcedSteerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype delete forced steer params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeDeleteForcedSteerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype delete forced steer params
func (o *CustomerPrototypeDeleteForcedSteerParams) WithTimeout(timeout time.Duration) *CustomerPrototypeDeleteForcedSteerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype delete forced steer params
func (o *CustomerPrototypeDeleteForcedSteerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype delete forced steer params
func (o *CustomerPrototypeDeleteForcedSteerParams) WithContext(ctx context.Context) *CustomerPrototypeDeleteForcedSteerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype delete forced steer params
func (o *CustomerPrototypeDeleteForcedSteerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype delete forced steer params
func (o *CustomerPrototypeDeleteForcedSteerParams) WithHTTPClient(client *http.Client) *CustomerPrototypeDeleteForcedSteerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype delete forced steer params
func (o *CustomerPrototypeDeleteForcedSteerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the customer prototype delete forced steer params
func (o *CustomerPrototypeDeleteForcedSteerParams) WithID(id string) *CustomerPrototypeDeleteForcedSteerParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype delete forced steer params
func (o *CustomerPrototypeDeleteForcedSteerParams) SetID(id string) {
	o.ID = id
}

// WithLocationID adds the locationID to the customer prototype delete forced steer params
func (o *CustomerPrototypeDeleteForcedSteerParams) WithLocationID(locationID string) *CustomerPrototypeDeleteForcedSteerParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype delete forced steer params
func (o *CustomerPrototypeDeleteForcedSteerParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WithMac adds the mac to the customer prototype delete forced steer params
func (o *CustomerPrototypeDeleteForcedSteerParams) WithMac(mac string) *CustomerPrototypeDeleteForcedSteerParams {
	o.SetMac(mac)
	return o
}

// SetMac adds the mac to the customer prototype delete forced steer params
func (o *CustomerPrototypeDeleteForcedSteerParams) SetMac(mac string) {
	o.Mac = mac
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypeDeleteForcedSteerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param mac
	if err := r.SetPathParam("mac", o.Mac); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
