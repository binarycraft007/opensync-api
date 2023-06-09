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

// NewLocationFindParams creates a new LocationFindParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLocationFindParams() *LocationFindParams {
	return &LocationFindParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLocationFindParamsWithTimeout creates a new LocationFindParams object
// with the ability to set a timeout on a request.
func NewLocationFindParamsWithTimeout(timeout time.Duration) *LocationFindParams {
	return &LocationFindParams{
		timeout: timeout,
	}
}

// NewLocationFindParamsWithContext creates a new LocationFindParams object
// with the ability to set a context for a request.
func NewLocationFindParamsWithContext(ctx context.Context) *LocationFindParams {
	return &LocationFindParams{
		Context: ctx,
	}
}

// NewLocationFindParamsWithHTTPClient creates a new LocationFindParams object
// with the ability to set a custom HTTPClient for a request.
func NewLocationFindParamsWithHTTPClient(client *http.Client) *LocationFindParams {
	return &LocationFindParams{
		HTTPClient: client,
	}
}

/*
LocationFindParams contains all the parameters to send to the API endpoint

	for the location find operation.

	Typically these are written to a http.Request.
*/
type LocationFindParams struct {

	/* Filter.

	   Filter defining fields, where, include, order, offset, and limit - must be a JSON-encoded string (`{"where":{"something":"value"}}`).  See https://loopback.io/doc/en/lb3/Querying-data.html#using-stringified-json-in-rest-queries for more details.

	   Format: JSON
	*/
	Filter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the location find params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LocationFindParams) WithDefaults() *LocationFindParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the location find params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LocationFindParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the location find params
func (o *LocationFindParams) WithTimeout(timeout time.Duration) *LocationFindParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the location find params
func (o *LocationFindParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the location find params
func (o *LocationFindParams) WithContext(ctx context.Context) *LocationFindParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the location find params
func (o *LocationFindParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the location find params
func (o *LocationFindParams) WithHTTPClient(client *http.Client) *LocationFindParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the location find params
func (o *LocationFindParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the location find params
func (o *LocationFindParams) WithFilter(filter *string) *LocationFindParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the location find params
func (o *LocationFindParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WriteToRequest writes these params to a swagger request
func (o *LocationFindParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Filter != nil {

		// query param filter
		var qrFilter string

		if o.Filter != nil {
			qrFilter = *o.Filter
		}
		qFilter := qrFilter
		if qFilter != "" {

			if err := r.SetQueryParam("filter", qFilter); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
