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

// NewCustomerPrototypeDeleteDeviceFreezeForeverParams creates a new CustomerPrototypeDeleteDeviceFreezeForeverParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypeDeleteDeviceFreezeForeverParams() *CustomerPrototypeDeleteDeviceFreezeForeverParams {
	return &CustomerPrototypeDeleteDeviceFreezeForeverParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypeDeleteDeviceFreezeForeverParamsWithTimeout creates a new CustomerPrototypeDeleteDeviceFreezeForeverParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypeDeleteDeviceFreezeForeverParamsWithTimeout(timeout time.Duration) *CustomerPrototypeDeleteDeviceFreezeForeverParams {
	return &CustomerPrototypeDeleteDeviceFreezeForeverParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypeDeleteDeviceFreezeForeverParamsWithContext creates a new CustomerPrototypeDeleteDeviceFreezeForeverParams object
// with the ability to set a context for a request.
func NewCustomerPrototypeDeleteDeviceFreezeForeverParamsWithContext(ctx context.Context) *CustomerPrototypeDeleteDeviceFreezeForeverParams {
	return &CustomerPrototypeDeleteDeviceFreezeForeverParams{
		Context: ctx,
	}
}

// NewCustomerPrototypeDeleteDeviceFreezeForeverParamsWithHTTPClient creates a new CustomerPrototypeDeleteDeviceFreezeForeverParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypeDeleteDeviceFreezeForeverParamsWithHTTPClient(client *http.Client) *CustomerPrototypeDeleteDeviceFreezeForeverParams {
	return &CustomerPrototypeDeleteDeviceFreezeForeverParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypeDeleteDeviceFreezeForeverParams contains all the parameters to send to the API endpoint

	for the customer prototype delete device freeze forever operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypeDeleteDeviceFreezeForeverParams struct {

	/* ID.

	   Customer id

	   Format: JSON
	*/
	ID string

	// LocationID.
	LocationID string

	// Mac.
	Mac string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer prototype delete device freeze forever params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeDeleteDeviceFreezeForeverParams) WithDefaults() *CustomerPrototypeDeleteDeviceFreezeForeverParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype delete device freeze forever params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeDeleteDeviceFreezeForeverParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype delete device freeze forever params
func (o *CustomerPrototypeDeleteDeviceFreezeForeverParams) WithTimeout(timeout time.Duration) *CustomerPrototypeDeleteDeviceFreezeForeverParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype delete device freeze forever params
func (o *CustomerPrototypeDeleteDeviceFreezeForeverParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype delete device freeze forever params
func (o *CustomerPrototypeDeleteDeviceFreezeForeverParams) WithContext(ctx context.Context) *CustomerPrototypeDeleteDeviceFreezeForeverParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype delete device freeze forever params
func (o *CustomerPrototypeDeleteDeviceFreezeForeverParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype delete device freeze forever params
func (o *CustomerPrototypeDeleteDeviceFreezeForeverParams) WithHTTPClient(client *http.Client) *CustomerPrototypeDeleteDeviceFreezeForeverParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype delete device freeze forever params
func (o *CustomerPrototypeDeleteDeviceFreezeForeverParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the customer prototype delete device freeze forever params
func (o *CustomerPrototypeDeleteDeviceFreezeForeverParams) WithID(id string) *CustomerPrototypeDeleteDeviceFreezeForeverParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype delete device freeze forever params
func (o *CustomerPrototypeDeleteDeviceFreezeForeverParams) SetID(id string) {
	o.ID = id
}

// WithLocationID adds the locationID to the customer prototype delete device freeze forever params
func (o *CustomerPrototypeDeleteDeviceFreezeForeverParams) WithLocationID(locationID string) *CustomerPrototypeDeleteDeviceFreezeForeverParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype delete device freeze forever params
func (o *CustomerPrototypeDeleteDeviceFreezeForeverParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WithMac adds the mac to the customer prototype delete device freeze forever params
func (o *CustomerPrototypeDeleteDeviceFreezeForeverParams) WithMac(mac string) *CustomerPrototypeDeleteDeviceFreezeForeverParams {
	o.SetMac(mac)
	return o
}

// SetMac adds the mac to the customer prototype delete device freeze forever params
func (o *CustomerPrototypeDeleteDeviceFreezeForeverParams) SetMac(mac string) {
	o.Mac = mac
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypeDeleteDeviceFreezeForeverParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
