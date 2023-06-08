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

// NewLocationPrototypeGetNodesParams creates a new LocationPrototypeGetNodesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLocationPrototypeGetNodesParams() *LocationPrototypeGetNodesParams {
	return &LocationPrototypeGetNodesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLocationPrototypeGetNodesParamsWithTimeout creates a new LocationPrototypeGetNodesParams object
// with the ability to set a timeout on a request.
func NewLocationPrototypeGetNodesParamsWithTimeout(timeout time.Duration) *LocationPrototypeGetNodesParams {
	return &LocationPrototypeGetNodesParams{
		timeout: timeout,
	}
}

// NewLocationPrototypeGetNodesParamsWithContext creates a new LocationPrototypeGetNodesParams object
// with the ability to set a context for a request.
func NewLocationPrototypeGetNodesParamsWithContext(ctx context.Context) *LocationPrototypeGetNodesParams {
	return &LocationPrototypeGetNodesParams{
		Context: ctx,
	}
}

// NewLocationPrototypeGetNodesParamsWithHTTPClient creates a new LocationPrototypeGetNodesParams object
// with the ability to set a custom HTTPClient for a request.
func NewLocationPrototypeGetNodesParamsWithHTTPClient(client *http.Client) *LocationPrototypeGetNodesParams {
	return &LocationPrototypeGetNodesParams{
		HTTPClient: client,
	}
}

/*
LocationPrototypeGetNodesParams contains all the parameters to send to the API endpoint

	for the location prototype get nodes operation.

	Typically these are written to a http.Request.
*/
type LocationPrototypeGetNodesParams struct {

	// Filter.
	//
	// Format: JSON
	Filter *string

	/* ID.

	   Location id

	   Format: JSON
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the location prototype get nodes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LocationPrototypeGetNodesParams) WithDefaults() *LocationPrototypeGetNodesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the location prototype get nodes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LocationPrototypeGetNodesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the location prototype get nodes params
func (o *LocationPrototypeGetNodesParams) WithTimeout(timeout time.Duration) *LocationPrototypeGetNodesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the location prototype get nodes params
func (o *LocationPrototypeGetNodesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the location prototype get nodes params
func (o *LocationPrototypeGetNodesParams) WithContext(ctx context.Context) *LocationPrototypeGetNodesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the location prototype get nodes params
func (o *LocationPrototypeGetNodesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the location prototype get nodes params
func (o *LocationPrototypeGetNodesParams) WithHTTPClient(client *http.Client) *LocationPrototypeGetNodesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the location prototype get nodes params
func (o *LocationPrototypeGetNodesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the location prototype get nodes params
func (o *LocationPrototypeGetNodesParams) WithFilter(filter *string) *LocationPrototypeGetNodesParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the location prototype get nodes params
func (o *LocationPrototypeGetNodesParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithID adds the id to the location prototype get nodes params
func (o *LocationPrototypeGetNodesParams) WithID(id string) *LocationPrototypeGetNodesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the location prototype get nodes params
func (o *LocationPrototypeGetNodesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *LocationPrototypeGetNodesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
