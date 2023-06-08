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

// NewLocationPrototypeGetWifiMotionParams creates a new LocationPrototypeGetWifiMotionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLocationPrototypeGetWifiMotionParams() *LocationPrototypeGetWifiMotionParams {
	return &LocationPrototypeGetWifiMotionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLocationPrototypeGetWifiMotionParamsWithTimeout creates a new LocationPrototypeGetWifiMotionParams object
// with the ability to set a timeout on a request.
func NewLocationPrototypeGetWifiMotionParamsWithTimeout(timeout time.Duration) *LocationPrototypeGetWifiMotionParams {
	return &LocationPrototypeGetWifiMotionParams{
		timeout: timeout,
	}
}

// NewLocationPrototypeGetWifiMotionParamsWithContext creates a new LocationPrototypeGetWifiMotionParams object
// with the ability to set a context for a request.
func NewLocationPrototypeGetWifiMotionParamsWithContext(ctx context.Context) *LocationPrototypeGetWifiMotionParams {
	return &LocationPrototypeGetWifiMotionParams{
		Context: ctx,
	}
}

// NewLocationPrototypeGetWifiMotionParamsWithHTTPClient creates a new LocationPrototypeGetWifiMotionParams object
// with the ability to set a custom HTTPClient for a request.
func NewLocationPrototypeGetWifiMotionParamsWithHTTPClient(client *http.Client) *LocationPrototypeGetWifiMotionParams {
	return &LocationPrototypeGetWifiMotionParams{
		HTTPClient: client,
	}
}

/*
LocationPrototypeGetWifiMotionParams contains all the parameters to send to the API endpoint

	for the location prototype get wifi motion operation.

	Typically these are written to a http.Request.
*/
type LocationPrototypeGetWifiMotionParams struct {

	/* ID.

	   Location id

	   Format: JSON
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the location prototype get wifi motion params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LocationPrototypeGetWifiMotionParams) WithDefaults() *LocationPrototypeGetWifiMotionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the location prototype get wifi motion params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LocationPrototypeGetWifiMotionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the location prototype get wifi motion params
func (o *LocationPrototypeGetWifiMotionParams) WithTimeout(timeout time.Duration) *LocationPrototypeGetWifiMotionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the location prototype get wifi motion params
func (o *LocationPrototypeGetWifiMotionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the location prototype get wifi motion params
func (o *LocationPrototypeGetWifiMotionParams) WithContext(ctx context.Context) *LocationPrototypeGetWifiMotionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the location prototype get wifi motion params
func (o *LocationPrototypeGetWifiMotionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the location prototype get wifi motion params
func (o *LocationPrototypeGetWifiMotionParams) WithHTTPClient(client *http.Client) *LocationPrototypeGetWifiMotionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the location prototype get wifi motion params
func (o *LocationPrototypeGetWifiMotionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the location prototype get wifi motion params
func (o *LocationPrototypeGetWifiMotionParams) WithID(id string) *LocationPrototypeGetWifiMotionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the location prototype get wifi motion params
func (o *LocationPrototypeGetWifiMotionParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *LocationPrototypeGetWifiMotionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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