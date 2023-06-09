// Code generated by go-swagger; DO NOT EDIT.

package group

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

// NewGroupPrototypeGetCustomersParams creates a new GroupPrototypeGetCustomersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGroupPrototypeGetCustomersParams() *GroupPrototypeGetCustomersParams {
	return &GroupPrototypeGetCustomersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGroupPrototypeGetCustomersParamsWithTimeout creates a new GroupPrototypeGetCustomersParams object
// with the ability to set a timeout on a request.
func NewGroupPrototypeGetCustomersParamsWithTimeout(timeout time.Duration) *GroupPrototypeGetCustomersParams {
	return &GroupPrototypeGetCustomersParams{
		timeout: timeout,
	}
}

// NewGroupPrototypeGetCustomersParamsWithContext creates a new GroupPrototypeGetCustomersParams object
// with the ability to set a context for a request.
func NewGroupPrototypeGetCustomersParamsWithContext(ctx context.Context) *GroupPrototypeGetCustomersParams {
	return &GroupPrototypeGetCustomersParams{
		Context: ctx,
	}
}

// NewGroupPrototypeGetCustomersParamsWithHTTPClient creates a new GroupPrototypeGetCustomersParams object
// with the ability to set a custom HTTPClient for a request.
func NewGroupPrototypeGetCustomersParamsWithHTTPClient(client *http.Client) *GroupPrototypeGetCustomersParams {
	return &GroupPrototypeGetCustomersParams{
		HTTPClient: client,
	}
}

/*
GroupPrototypeGetCustomersParams contains all the parameters to send to the API endpoint

	for the group prototype get customers operation.

	Typically these are written to a http.Request.
*/
type GroupPrototypeGetCustomersParams struct {

	// Filter.
	//
	// Format: JSON
	Filter *string

	/* ID.

	   Group id
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the group prototype get customers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GroupPrototypeGetCustomersParams) WithDefaults() *GroupPrototypeGetCustomersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the group prototype get customers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GroupPrototypeGetCustomersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the group prototype get customers params
func (o *GroupPrototypeGetCustomersParams) WithTimeout(timeout time.Duration) *GroupPrototypeGetCustomersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the group prototype get customers params
func (o *GroupPrototypeGetCustomersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the group prototype get customers params
func (o *GroupPrototypeGetCustomersParams) WithContext(ctx context.Context) *GroupPrototypeGetCustomersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the group prototype get customers params
func (o *GroupPrototypeGetCustomersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the group prototype get customers params
func (o *GroupPrototypeGetCustomersParams) WithHTTPClient(client *http.Client) *GroupPrototypeGetCustomersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the group prototype get customers params
func (o *GroupPrototypeGetCustomersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the group prototype get customers params
func (o *GroupPrototypeGetCustomersParams) WithFilter(filter *string) *GroupPrototypeGetCustomersParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the group prototype get customers params
func (o *GroupPrototypeGetCustomersParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithID adds the id to the group prototype get customers params
func (o *GroupPrototypeGetCustomersParams) WithID(id string) *GroupPrototypeGetCustomersParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the group prototype get customers params
func (o *GroupPrototypeGetCustomersParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GroupPrototypeGetCustomersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
