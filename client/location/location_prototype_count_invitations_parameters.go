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

// NewLocationPrototypeCountInvitationsParams creates a new LocationPrototypeCountInvitationsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLocationPrototypeCountInvitationsParams() *LocationPrototypeCountInvitationsParams {
	return &LocationPrototypeCountInvitationsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLocationPrototypeCountInvitationsParamsWithTimeout creates a new LocationPrototypeCountInvitationsParams object
// with the ability to set a timeout on a request.
func NewLocationPrototypeCountInvitationsParamsWithTimeout(timeout time.Duration) *LocationPrototypeCountInvitationsParams {
	return &LocationPrototypeCountInvitationsParams{
		timeout: timeout,
	}
}

// NewLocationPrototypeCountInvitationsParamsWithContext creates a new LocationPrototypeCountInvitationsParams object
// with the ability to set a context for a request.
func NewLocationPrototypeCountInvitationsParamsWithContext(ctx context.Context) *LocationPrototypeCountInvitationsParams {
	return &LocationPrototypeCountInvitationsParams{
		Context: ctx,
	}
}

// NewLocationPrototypeCountInvitationsParamsWithHTTPClient creates a new LocationPrototypeCountInvitationsParams object
// with the ability to set a custom HTTPClient for a request.
func NewLocationPrototypeCountInvitationsParamsWithHTTPClient(client *http.Client) *LocationPrototypeCountInvitationsParams {
	return &LocationPrototypeCountInvitationsParams{
		HTTPClient: client,
	}
}

/*
LocationPrototypeCountInvitationsParams contains all the parameters to send to the API endpoint

	for the location prototype count invitations operation.

	Typically these are written to a http.Request.
*/
type LocationPrototypeCountInvitationsParams struct {

	/* ID.

	   Location id

	   Format: JSON
	*/
	ID string

	/* Where.

	   Criteria to match model instances

	   Format: JSON
	*/
	Where *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the location prototype count invitations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LocationPrototypeCountInvitationsParams) WithDefaults() *LocationPrototypeCountInvitationsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the location prototype count invitations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LocationPrototypeCountInvitationsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the location prototype count invitations params
func (o *LocationPrototypeCountInvitationsParams) WithTimeout(timeout time.Duration) *LocationPrototypeCountInvitationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the location prototype count invitations params
func (o *LocationPrototypeCountInvitationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the location prototype count invitations params
func (o *LocationPrototypeCountInvitationsParams) WithContext(ctx context.Context) *LocationPrototypeCountInvitationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the location prototype count invitations params
func (o *LocationPrototypeCountInvitationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the location prototype count invitations params
func (o *LocationPrototypeCountInvitationsParams) WithHTTPClient(client *http.Client) *LocationPrototypeCountInvitationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the location prototype count invitations params
func (o *LocationPrototypeCountInvitationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the location prototype count invitations params
func (o *LocationPrototypeCountInvitationsParams) WithID(id string) *LocationPrototypeCountInvitationsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the location prototype count invitations params
func (o *LocationPrototypeCountInvitationsParams) SetID(id string) {
	o.ID = id
}

// WithWhere adds the where to the location prototype count invitations params
func (o *LocationPrototypeCountInvitationsParams) WithWhere(where *string) *LocationPrototypeCountInvitationsParams {
	o.SetWhere(where)
	return o
}

// SetWhere adds the where to the location prototype count invitations params
func (o *LocationPrototypeCountInvitationsParams) SetWhere(where *string) {
	o.Where = where
}

// WriteToRequest writes these params to a swagger request
func (o *LocationPrototypeCountInvitationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.Where != nil {

		// query param where
		var qrWhere string

		if o.Where != nil {
			qrWhere = *o.Where
		}
		qWhere := qrWhere
		if qWhere != "" {

			if err := r.SetQueryParam("where", qWhere); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}