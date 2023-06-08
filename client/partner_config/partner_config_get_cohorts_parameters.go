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

// NewPartnerConfigGetCohortsParams creates a new PartnerConfigGetCohortsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPartnerConfigGetCohortsParams() *PartnerConfigGetCohortsParams {
	return &PartnerConfigGetCohortsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPartnerConfigGetCohortsParamsWithTimeout creates a new PartnerConfigGetCohortsParams object
// with the ability to set a timeout on a request.
func NewPartnerConfigGetCohortsParamsWithTimeout(timeout time.Duration) *PartnerConfigGetCohortsParams {
	return &PartnerConfigGetCohortsParams{
		timeout: timeout,
	}
}

// NewPartnerConfigGetCohortsParamsWithContext creates a new PartnerConfigGetCohortsParams object
// with the ability to set a context for a request.
func NewPartnerConfigGetCohortsParamsWithContext(ctx context.Context) *PartnerConfigGetCohortsParams {
	return &PartnerConfigGetCohortsParams{
		Context: ctx,
	}
}

// NewPartnerConfigGetCohortsParamsWithHTTPClient creates a new PartnerConfigGetCohortsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPartnerConfigGetCohortsParamsWithHTTPClient(client *http.Client) *PartnerConfigGetCohortsParams {
	return &PartnerConfigGetCohortsParams{
		HTTPClient: client,
	}
}

/*
PartnerConfigGetCohortsParams contains all the parameters to send to the API endpoint

	for the partner config get cohorts operation.

	Typically these are written to a http.Request.
*/
type PartnerConfigGetCohortsParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the partner config get cohorts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PartnerConfigGetCohortsParams) WithDefaults() *PartnerConfigGetCohortsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the partner config get cohorts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PartnerConfigGetCohortsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the partner config get cohorts params
func (o *PartnerConfigGetCohortsParams) WithTimeout(timeout time.Duration) *PartnerConfigGetCohortsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the partner config get cohorts params
func (o *PartnerConfigGetCohortsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the partner config get cohorts params
func (o *PartnerConfigGetCohortsParams) WithContext(ctx context.Context) *PartnerConfigGetCohortsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the partner config get cohorts params
func (o *PartnerConfigGetCohortsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the partner config get cohorts params
func (o *PartnerConfigGetCohortsParams) WithHTTPClient(client *http.Client) *PartnerConfigGetCohortsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the partner config get cohorts params
func (o *PartnerConfigGetCohortsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the partner config get cohorts params
func (o *PartnerConfigGetCohortsParams) WithID(id string) *PartnerConfigGetCohortsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the partner config get cohorts params
func (o *PartnerConfigGetCohortsParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PartnerConfigGetCohortsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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