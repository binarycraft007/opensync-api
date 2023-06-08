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

	"github.com/binarycraft007/opensync-api/models"
)

// NewLocationPrototypeCreateInvitationsParams creates a new LocationPrototypeCreateInvitationsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLocationPrototypeCreateInvitationsParams() *LocationPrototypeCreateInvitationsParams {
	return &LocationPrototypeCreateInvitationsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLocationPrototypeCreateInvitationsParamsWithTimeout creates a new LocationPrototypeCreateInvitationsParams object
// with the ability to set a timeout on a request.
func NewLocationPrototypeCreateInvitationsParamsWithTimeout(timeout time.Duration) *LocationPrototypeCreateInvitationsParams {
	return &LocationPrototypeCreateInvitationsParams{
		timeout: timeout,
	}
}

// NewLocationPrototypeCreateInvitationsParamsWithContext creates a new LocationPrototypeCreateInvitationsParams object
// with the ability to set a context for a request.
func NewLocationPrototypeCreateInvitationsParamsWithContext(ctx context.Context) *LocationPrototypeCreateInvitationsParams {
	return &LocationPrototypeCreateInvitationsParams{
		Context: ctx,
	}
}

// NewLocationPrototypeCreateInvitationsParamsWithHTTPClient creates a new LocationPrototypeCreateInvitationsParams object
// with the ability to set a custom HTTPClient for a request.
func NewLocationPrototypeCreateInvitationsParamsWithHTTPClient(client *http.Client) *LocationPrototypeCreateInvitationsParams {
	return &LocationPrototypeCreateInvitationsParams{
		HTTPClient: client,
	}
}

/*
LocationPrototypeCreateInvitationsParams contains all the parameters to send to the API endpoint

	for the location prototype create invitations operation.

	Typically these are written to a http.Request.
*/
type LocationPrototypeCreateInvitationsParams struct {

	// Data.
	Data *models.Invitation

	/* ID.

	   Location id

	   Format: JSON
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the location prototype create invitations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LocationPrototypeCreateInvitationsParams) WithDefaults() *LocationPrototypeCreateInvitationsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the location prototype create invitations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LocationPrototypeCreateInvitationsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the location prototype create invitations params
func (o *LocationPrototypeCreateInvitationsParams) WithTimeout(timeout time.Duration) *LocationPrototypeCreateInvitationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the location prototype create invitations params
func (o *LocationPrototypeCreateInvitationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the location prototype create invitations params
func (o *LocationPrototypeCreateInvitationsParams) WithContext(ctx context.Context) *LocationPrototypeCreateInvitationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the location prototype create invitations params
func (o *LocationPrototypeCreateInvitationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the location prototype create invitations params
func (o *LocationPrototypeCreateInvitationsParams) WithHTTPClient(client *http.Client) *LocationPrototypeCreateInvitationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the location prototype create invitations params
func (o *LocationPrototypeCreateInvitationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the location prototype create invitations params
func (o *LocationPrototypeCreateInvitationsParams) WithData(data *models.Invitation) *LocationPrototypeCreateInvitationsParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the location prototype create invitations params
func (o *LocationPrototypeCreateInvitationsParams) SetData(data *models.Invitation) {
	o.Data = data
}

// WithID adds the id to the location prototype create invitations params
func (o *LocationPrototypeCreateInvitationsParams) WithID(id string) *LocationPrototypeCreateInvitationsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the location prototype create invitations params
func (o *LocationPrototypeCreateInvitationsParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *LocationPrototypeCreateInvitationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
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
