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

// NewLocationPrototypeMarketingExportParams creates a new LocationPrototypeMarketingExportParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLocationPrototypeMarketingExportParams() *LocationPrototypeMarketingExportParams {
	return &LocationPrototypeMarketingExportParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLocationPrototypeMarketingExportParamsWithTimeout creates a new LocationPrototypeMarketingExportParams object
// with the ability to set a timeout on a request.
func NewLocationPrototypeMarketingExportParamsWithTimeout(timeout time.Duration) *LocationPrototypeMarketingExportParams {
	return &LocationPrototypeMarketingExportParams{
		timeout: timeout,
	}
}

// NewLocationPrototypeMarketingExportParamsWithContext creates a new LocationPrototypeMarketingExportParams object
// with the ability to set a context for a request.
func NewLocationPrototypeMarketingExportParamsWithContext(ctx context.Context) *LocationPrototypeMarketingExportParams {
	return &LocationPrototypeMarketingExportParams{
		Context: ctx,
	}
}

// NewLocationPrototypeMarketingExportParamsWithHTTPClient creates a new LocationPrototypeMarketingExportParams object
// with the ability to set a custom HTTPClient for a request.
func NewLocationPrototypeMarketingExportParamsWithHTTPClient(client *http.Client) *LocationPrototypeMarketingExportParams {
	return &LocationPrototypeMarketingExportParams{
		HTTPClient: client,
	}
}

/*
LocationPrototypeMarketingExportParams contains all the parameters to send to the API endpoint

	for the location prototype marketing export operation.

	Typically these are written to a http.Request.
*/
type LocationPrototypeMarketingExportParams struct {

	/* ID.

	   Location id

	   Format: JSON
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the location prototype marketing export params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LocationPrototypeMarketingExportParams) WithDefaults() *LocationPrototypeMarketingExportParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the location prototype marketing export params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LocationPrototypeMarketingExportParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the location prototype marketing export params
func (o *LocationPrototypeMarketingExportParams) WithTimeout(timeout time.Duration) *LocationPrototypeMarketingExportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the location prototype marketing export params
func (o *LocationPrototypeMarketingExportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the location prototype marketing export params
func (o *LocationPrototypeMarketingExportParams) WithContext(ctx context.Context) *LocationPrototypeMarketingExportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the location prototype marketing export params
func (o *LocationPrototypeMarketingExportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the location prototype marketing export params
func (o *LocationPrototypeMarketingExportParams) WithHTTPClient(client *http.Client) *LocationPrototypeMarketingExportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the location prototype marketing export params
func (o *LocationPrototypeMarketingExportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the location prototype marketing export params
func (o *LocationPrototypeMarketingExportParams) WithID(id string) *LocationPrototypeMarketingExportParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the location prototype marketing export params
func (o *LocationPrototypeMarketingExportParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *LocationPrototypeMarketingExportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
