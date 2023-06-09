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

// NewPartnerConfigDeleteAppQoeParams creates a new PartnerConfigDeleteAppQoeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPartnerConfigDeleteAppQoeParams() *PartnerConfigDeleteAppQoeParams {
	return &PartnerConfigDeleteAppQoeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPartnerConfigDeleteAppQoeParamsWithTimeout creates a new PartnerConfigDeleteAppQoeParams object
// with the ability to set a timeout on a request.
func NewPartnerConfigDeleteAppQoeParamsWithTimeout(timeout time.Duration) *PartnerConfigDeleteAppQoeParams {
	return &PartnerConfigDeleteAppQoeParams{
		timeout: timeout,
	}
}

// NewPartnerConfigDeleteAppQoeParamsWithContext creates a new PartnerConfigDeleteAppQoeParams object
// with the ability to set a context for a request.
func NewPartnerConfigDeleteAppQoeParamsWithContext(ctx context.Context) *PartnerConfigDeleteAppQoeParams {
	return &PartnerConfigDeleteAppQoeParams{
		Context: ctx,
	}
}

// NewPartnerConfigDeleteAppQoeParamsWithHTTPClient creates a new PartnerConfigDeleteAppQoeParams object
// with the ability to set a custom HTTPClient for a request.
func NewPartnerConfigDeleteAppQoeParamsWithHTTPClient(client *http.Client) *PartnerConfigDeleteAppQoeParams {
	return &PartnerConfigDeleteAppQoeParams{
		HTTPClient: client,
	}
}

/*
PartnerConfigDeleteAppQoeParams contains all the parameters to send to the API endpoint

	for the partner config delete app qoe operation.

	Typically these are written to a http.Request.
*/
type PartnerConfigDeleteAppQoeParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the partner config delete app qoe params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PartnerConfigDeleteAppQoeParams) WithDefaults() *PartnerConfigDeleteAppQoeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the partner config delete app qoe params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PartnerConfigDeleteAppQoeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the partner config delete app qoe params
func (o *PartnerConfigDeleteAppQoeParams) WithTimeout(timeout time.Duration) *PartnerConfigDeleteAppQoeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the partner config delete app qoe params
func (o *PartnerConfigDeleteAppQoeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the partner config delete app qoe params
func (o *PartnerConfigDeleteAppQoeParams) WithContext(ctx context.Context) *PartnerConfigDeleteAppQoeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the partner config delete app qoe params
func (o *PartnerConfigDeleteAppQoeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the partner config delete app qoe params
func (o *PartnerConfigDeleteAppQoeParams) WithHTTPClient(client *http.Client) *PartnerConfigDeleteAppQoeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the partner config delete app qoe params
func (o *PartnerConfigDeleteAppQoeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the partner config delete app qoe params
func (o *PartnerConfigDeleteAppQoeParams) WithID(id string) *PartnerConfigDeleteAppQoeParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the partner config delete app qoe params
func (o *PartnerConfigDeleteAppQoeParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PartnerConfigDeleteAppQoeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
