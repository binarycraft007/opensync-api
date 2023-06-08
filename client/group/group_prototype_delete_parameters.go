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

// NewGroupPrototypeDeleteParams creates a new GroupPrototypeDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGroupPrototypeDeleteParams() *GroupPrototypeDeleteParams {
	return &GroupPrototypeDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGroupPrototypeDeleteParamsWithTimeout creates a new GroupPrototypeDeleteParams object
// with the ability to set a timeout on a request.
func NewGroupPrototypeDeleteParamsWithTimeout(timeout time.Duration) *GroupPrototypeDeleteParams {
	return &GroupPrototypeDeleteParams{
		timeout: timeout,
	}
}

// NewGroupPrototypeDeleteParamsWithContext creates a new GroupPrototypeDeleteParams object
// with the ability to set a context for a request.
func NewGroupPrototypeDeleteParamsWithContext(ctx context.Context) *GroupPrototypeDeleteParams {
	return &GroupPrototypeDeleteParams{
		Context: ctx,
	}
}

// NewGroupPrototypeDeleteParamsWithHTTPClient creates a new GroupPrototypeDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewGroupPrototypeDeleteParamsWithHTTPClient(client *http.Client) *GroupPrototypeDeleteParams {
	return &GroupPrototypeDeleteParams{
		HTTPClient: client,
	}
}

/*
GroupPrototypeDeleteParams contains all the parameters to send to the API endpoint

	for the group prototype delete operation.

	Typically these are written to a http.Request.
*/
type GroupPrototypeDeleteParams struct {

	/* ID.

	   Group id
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the group prototype delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GroupPrototypeDeleteParams) WithDefaults() *GroupPrototypeDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the group prototype delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GroupPrototypeDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the group prototype delete params
func (o *GroupPrototypeDeleteParams) WithTimeout(timeout time.Duration) *GroupPrototypeDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the group prototype delete params
func (o *GroupPrototypeDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the group prototype delete params
func (o *GroupPrototypeDeleteParams) WithContext(ctx context.Context) *GroupPrototypeDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the group prototype delete params
func (o *GroupPrototypeDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the group prototype delete params
func (o *GroupPrototypeDeleteParams) WithHTTPClient(client *http.Client) *GroupPrototypeDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the group prototype delete params
func (o *GroupPrototypeDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the group prototype delete params
func (o *GroupPrototypeDeleteParams) WithID(id string) *GroupPrototypeDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the group prototype delete params
func (o *GroupPrototypeDeleteParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GroupPrototypeDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
