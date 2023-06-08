// Code generated by go-swagger; DO NOT EDIT.

package partner_config

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

// NewPartnerConfigGetAppQoeParams creates a new PartnerConfigGetAppQoeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPartnerConfigGetAppQoeParams() *PartnerConfigGetAppQoeParams {
	return &PartnerConfigGetAppQoeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPartnerConfigGetAppQoeParamsWithTimeout creates a new PartnerConfigGetAppQoeParams object
// with the ability to set a timeout on a request.
func NewPartnerConfigGetAppQoeParamsWithTimeout(timeout time.Duration) *PartnerConfigGetAppQoeParams {
	return &PartnerConfigGetAppQoeParams{
		timeout: timeout,
	}
}

// NewPartnerConfigGetAppQoeParamsWithContext creates a new PartnerConfigGetAppQoeParams object
// with the ability to set a context for a request.
func NewPartnerConfigGetAppQoeParamsWithContext(ctx context.Context) *PartnerConfigGetAppQoeParams {
	return &PartnerConfigGetAppQoeParams{
		Context: ctx,
	}
}

// NewPartnerConfigGetAppQoeParamsWithHTTPClient creates a new PartnerConfigGetAppQoeParams object
// with the ability to set a custom HTTPClient for a request.
func NewPartnerConfigGetAppQoeParamsWithHTTPClient(client *http.Client) *PartnerConfigGetAppQoeParams {
	return &PartnerConfigGetAppQoeParams{
		HTTPClient: client,
	}
}

/*
PartnerConfigGetAppQoeParams contains all the parameters to send to the API endpoint

	for the partner config get app qoe operation.

	Typically these are written to a http.Request.
*/
type PartnerConfigGetAppQoeParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the partner config get app qoe params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PartnerConfigGetAppQoeParams) WithDefaults() *PartnerConfigGetAppQoeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the partner config get app qoe params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PartnerConfigGetAppQoeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the partner config get app qoe params
func (o *PartnerConfigGetAppQoeParams) WithTimeout(timeout time.Duration) *PartnerConfigGetAppQoeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the partner config get app qoe params
func (o *PartnerConfigGetAppQoeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the partner config get app qoe params
func (o *PartnerConfigGetAppQoeParams) WithContext(ctx context.Context) *PartnerConfigGetAppQoeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the partner config get app qoe params
func (o *PartnerConfigGetAppQoeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the partner config get app qoe params
func (o *PartnerConfigGetAppQoeParams) WithHTTPClient(client *http.Client) *PartnerConfigGetAppQoeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the partner config get app qoe params
func (o *PartnerConfigGetAppQoeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the partner config get app qoe params
func (o *PartnerConfigGetAppQoeParams) WithID(id string) *PartnerConfigGetAppQoeParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the partner config get app qoe params
func (o *PartnerConfigGetAppQoeParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PartnerConfigGetAppQoeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
