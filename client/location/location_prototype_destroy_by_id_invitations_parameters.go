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

// NewLocationPrototypeDestroyByIDInvitationsParams creates a new LocationPrototypeDestroyByIDInvitationsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLocationPrototypeDestroyByIDInvitationsParams() *LocationPrototypeDestroyByIDInvitationsParams {
	return &LocationPrototypeDestroyByIDInvitationsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLocationPrototypeDestroyByIDInvitationsParamsWithTimeout creates a new LocationPrototypeDestroyByIDInvitationsParams object
// with the ability to set a timeout on a request.
func NewLocationPrototypeDestroyByIDInvitationsParamsWithTimeout(timeout time.Duration) *LocationPrototypeDestroyByIDInvitationsParams {
	return &LocationPrototypeDestroyByIDInvitationsParams{
		timeout: timeout,
	}
}

// NewLocationPrototypeDestroyByIDInvitationsParamsWithContext creates a new LocationPrototypeDestroyByIDInvitationsParams object
// with the ability to set a context for a request.
func NewLocationPrototypeDestroyByIDInvitationsParamsWithContext(ctx context.Context) *LocationPrototypeDestroyByIDInvitationsParams {
	return &LocationPrototypeDestroyByIDInvitationsParams{
		Context: ctx,
	}
}

// NewLocationPrototypeDestroyByIDInvitationsParamsWithHTTPClient creates a new LocationPrototypeDestroyByIDInvitationsParams object
// with the ability to set a custom HTTPClient for a request.
func NewLocationPrototypeDestroyByIDInvitationsParamsWithHTTPClient(client *http.Client) *LocationPrototypeDestroyByIDInvitationsParams {
	return &LocationPrototypeDestroyByIDInvitationsParams{
		HTTPClient: client,
	}
}

/*
LocationPrototypeDestroyByIDInvitationsParams contains all the parameters to send to the API endpoint

	for the location prototype destroy by Id invitations operation.

	Typically these are written to a http.Request.
*/
type LocationPrototypeDestroyByIDInvitationsParams struct {

	/* Fk.

	   Foreign key for invitations

	   Format: JSON
	*/
	Fk string

	/* ID.

	   Location id

	   Format: JSON
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the location prototype destroy by Id invitations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LocationPrototypeDestroyByIDInvitationsParams) WithDefaults() *LocationPrototypeDestroyByIDInvitationsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the location prototype destroy by Id invitations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LocationPrototypeDestroyByIDInvitationsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the location prototype destroy by Id invitations params
func (o *LocationPrototypeDestroyByIDInvitationsParams) WithTimeout(timeout time.Duration) *LocationPrototypeDestroyByIDInvitationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the location prototype destroy by Id invitations params
func (o *LocationPrototypeDestroyByIDInvitationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the location prototype destroy by Id invitations params
func (o *LocationPrototypeDestroyByIDInvitationsParams) WithContext(ctx context.Context) *LocationPrototypeDestroyByIDInvitationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the location prototype destroy by Id invitations params
func (o *LocationPrototypeDestroyByIDInvitationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the location prototype destroy by Id invitations params
func (o *LocationPrototypeDestroyByIDInvitationsParams) WithHTTPClient(client *http.Client) *LocationPrototypeDestroyByIDInvitationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the location prototype destroy by Id invitations params
func (o *LocationPrototypeDestroyByIDInvitationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFk adds the fk to the location prototype destroy by Id invitations params
func (o *LocationPrototypeDestroyByIDInvitationsParams) WithFk(fk string) *LocationPrototypeDestroyByIDInvitationsParams {
	o.SetFk(fk)
	return o
}

// SetFk adds the fk to the location prototype destroy by Id invitations params
func (o *LocationPrototypeDestroyByIDInvitationsParams) SetFk(fk string) {
	o.Fk = fk
}

// WithID adds the id to the location prototype destroy by Id invitations params
func (o *LocationPrototypeDestroyByIDInvitationsParams) WithID(id string) *LocationPrototypeDestroyByIDInvitationsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the location prototype destroy by Id invitations params
func (o *LocationPrototypeDestroyByIDInvitationsParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *LocationPrototypeDestroyByIDInvitationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param fk
	if err := r.SetPathParam("fk", o.Fk); err != nil {
		return err
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
