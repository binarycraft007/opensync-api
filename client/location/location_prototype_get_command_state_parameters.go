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

// NewLocationPrototypeGetCommandStateParams creates a new LocationPrototypeGetCommandStateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLocationPrototypeGetCommandStateParams() *LocationPrototypeGetCommandStateParams {
	return &LocationPrototypeGetCommandStateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLocationPrototypeGetCommandStateParamsWithTimeout creates a new LocationPrototypeGetCommandStateParams object
// with the ability to set a timeout on a request.
func NewLocationPrototypeGetCommandStateParamsWithTimeout(timeout time.Duration) *LocationPrototypeGetCommandStateParams {
	return &LocationPrototypeGetCommandStateParams{
		timeout: timeout,
	}
}

// NewLocationPrototypeGetCommandStateParamsWithContext creates a new LocationPrototypeGetCommandStateParams object
// with the ability to set a context for a request.
func NewLocationPrototypeGetCommandStateParamsWithContext(ctx context.Context) *LocationPrototypeGetCommandStateParams {
	return &LocationPrototypeGetCommandStateParams{
		Context: ctx,
	}
}

// NewLocationPrototypeGetCommandStateParamsWithHTTPClient creates a new LocationPrototypeGetCommandStateParams object
// with the ability to set a custom HTTPClient for a request.
func NewLocationPrototypeGetCommandStateParamsWithHTTPClient(client *http.Client) *LocationPrototypeGetCommandStateParams {
	return &LocationPrototypeGetCommandStateParams{
		HTTPClient: client,
	}
}

/*
LocationPrototypeGetCommandStateParams contains all the parameters to send to the API endpoint

	for the location prototype get command state operation.

	Typically these are written to a http.Request.
*/
type LocationPrototypeGetCommandStateParams struct {

	/* ID.

	   Location id

	   Format: JSON
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the location prototype get command state params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LocationPrototypeGetCommandStateParams) WithDefaults() *LocationPrototypeGetCommandStateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the location prototype get command state params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LocationPrototypeGetCommandStateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the location prototype get command state params
func (o *LocationPrototypeGetCommandStateParams) WithTimeout(timeout time.Duration) *LocationPrototypeGetCommandStateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the location prototype get command state params
func (o *LocationPrototypeGetCommandStateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the location prototype get command state params
func (o *LocationPrototypeGetCommandStateParams) WithContext(ctx context.Context) *LocationPrototypeGetCommandStateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the location prototype get command state params
func (o *LocationPrototypeGetCommandStateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the location prototype get command state params
func (o *LocationPrototypeGetCommandStateParams) WithHTTPClient(client *http.Client) *LocationPrototypeGetCommandStateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the location prototype get command state params
func (o *LocationPrototypeGetCommandStateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the location prototype get command state params
func (o *LocationPrototypeGetCommandStateParams) WithID(id string) *LocationPrototypeGetCommandStateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the location prototype get command state params
func (o *LocationPrototypeGetCommandStateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *LocationPrototypeGetCommandStateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
