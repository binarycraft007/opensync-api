// Code generated by go-swagger; DO NOT EDIT.

package location

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

// NewLocationPrototypeGetDevicesByMacNameParams creates a new LocationPrototypeGetDevicesByMacNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLocationPrototypeGetDevicesByMacNameParams() *LocationPrototypeGetDevicesByMacNameParams {
	return &LocationPrototypeGetDevicesByMacNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLocationPrototypeGetDevicesByMacNameParamsWithTimeout creates a new LocationPrototypeGetDevicesByMacNameParams object
// with the ability to set a timeout on a request.
func NewLocationPrototypeGetDevicesByMacNameParamsWithTimeout(timeout time.Duration) *LocationPrototypeGetDevicesByMacNameParams {
	return &LocationPrototypeGetDevicesByMacNameParams{
		timeout: timeout,
	}
}

// NewLocationPrototypeGetDevicesByMacNameParamsWithContext creates a new LocationPrototypeGetDevicesByMacNameParams object
// with the ability to set a context for a request.
func NewLocationPrototypeGetDevicesByMacNameParamsWithContext(ctx context.Context) *LocationPrototypeGetDevicesByMacNameParams {
	return &LocationPrototypeGetDevicesByMacNameParams{
		Context: ctx,
	}
}

// NewLocationPrototypeGetDevicesByMacNameParamsWithHTTPClient creates a new LocationPrototypeGetDevicesByMacNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewLocationPrototypeGetDevicesByMacNameParamsWithHTTPClient(client *http.Client) *LocationPrototypeGetDevicesByMacNameParams {
	return &LocationPrototypeGetDevicesByMacNameParams{
		HTTPClient: client,
	}
}

/*
LocationPrototypeGetDevicesByMacNameParams contains all the parameters to send to the API endpoint

	for the location prototype get devices by mac name operation.

	Typically these are written to a http.Request.
*/
type LocationPrototypeGetDevicesByMacNameParams struct {

	/* ID.

	   Location id

	   Format: JSON
	*/
	ID string

	/* Mac.

	   mac of device
	*/
	Mac string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the location prototype get devices by mac name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LocationPrototypeGetDevicesByMacNameParams) WithDefaults() *LocationPrototypeGetDevicesByMacNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the location prototype get devices by mac name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LocationPrototypeGetDevicesByMacNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the location prototype get devices by mac name params
func (o *LocationPrototypeGetDevicesByMacNameParams) WithTimeout(timeout time.Duration) *LocationPrototypeGetDevicesByMacNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the location prototype get devices by mac name params
func (o *LocationPrototypeGetDevicesByMacNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the location prototype get devices by mac name params
func (o *LocationPrototypeGetDevicesByMacNameParams) WithContext(ctx context.Context) *LocationPrototypeGetDevicesByMacNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the location prototype get devices by mac name params
func (o *LocationPrototypeGetDevicesByMacNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the location prototype get devices by mac name params
func (o *LocationPrototypeGetDevicesByMacNameParams) WithHTTPClient(client *http.Client) *LocationPrototypeGetDevicesByMacNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the location prototype get devices by mac name params
func (o *LocationPrototypeGetDevicesByMacNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the location prototype get devices by mac name params
func (o *LocationPrototypeGetDevicesByMacNameParams) WithID(id string) *LocationPrototypeGetDevicesByMacNameParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the location prototype get devices by mac name params
func (o *LocationPrototypeGetDevicesByMacNameParams) SetID(id string) {
	o.ID = id
}

// WithMac adds the mac to the location prototype get devices by mac name params
func (o *LocationPrototypeGetDevicesByMacNameParams) WithMac(mac string) *LocationPrototypeGetDevicesByMacNameParams {
	o.SetMac(mac)
	return o
}

// SetMac adds the mac to the location prototype get devices by mac name params
func (o *LocationPrototypeGetDevicesByMacNameParams) SetMac(mac string) {
	o.Mac = mac
}

// WriteToRequest writes these params to a swagger request
func (o *LocationPrototypeGetDevicesByMacNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
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