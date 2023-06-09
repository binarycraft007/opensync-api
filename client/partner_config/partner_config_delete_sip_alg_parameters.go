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

// NewPartnerConfigDeleteSipAlgParams creates a new PartnerConfigDeleteSipAlgParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPartnerConfigDeleteSipAlgParams() *PartnerConfigDeleteSipAlgParams {
	return &PartnerConfigDeleteSipAlgParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPartnerConfigDeleteSipAlgParamsWithTimeout creates a new PartnerConfigDeleteSipAlgParams object
// with the ability to set a timeout on a request.
func NewPartnerConfigDeleteSipAlgParamsWithTimeout(timeout time.Duration) *PartnerConfigDeleteSipAlgParams {
	return &PartnerConfigDeleteSipAlgParams{
		timeout: timeout,
	}
}

// NewPartnerConfigDeleteSipAlgParamsWithContext creates a new PartnerConfigDeleteSipAlgParams object
// with the ability to set a context for a request.
func NewPartnerConfigDeleteSipAlgParamsWithContext(ctx context.Context) *PartnerConfigDeleteSipAlgParams {
	return &PartnerConfigDeleteSipAlgParams{
		Context: ctx,
	}
}

// NewPartnerConfigDeleteSipAlgParamsWithHTTPClient creates a new PartnerConfigDeleteSipAlgParams object
// with the ability to set a custom HTTPClient for a request.
func NewPartnerConfigDeleteSipAlgParamsWithHTTPClient(client *http.Client) *PartnerConfigDeleteSipAlgParams {
	return &PartnerConfigDeleteSipAlgParams{
		HTTPClient: client,
	}
}

/*
PartnerConfigDeleteSipAlgParams contains all the parameters to send to the API endpoint

	for the partner config delete sip alg operation.

	Typically these are written to a http.Request.
*/
type PartnerConfigDeleteSipAlgParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the partner config delete sip alg params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PartnerConfigDeleteSipAlgParams) WithDefaults() *PartnerConfigDeleteSipAlgParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the partner config delete sip alg params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PartnerConfigDeleteSipAlgParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the partner config delete sip alg params
func (o *PartnerConfigDeleteSipAlgParams) WithTimeout(timeout time.Duration) *PartnerConfigDeleteSipAlgParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the partner config delete sip alg params
func (o *PartnerConfigDeleteSipAlgParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the partner config delete sip alg params
func (o *PartnerConfigDeleteSipAlgParams) WithContext(ctx context.Context) *PartnerConfigDeleteSipAlgParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the partner config delete sip alg params
func (o *PartnerConfigDeleteSipAlgParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the partner config delete sip alg params
func (o *PartnerConfigDeleteSipAlgParams) WithHTTPClient(client *http.Client) *PartnerConfigDeleteSipAlgParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the partner config delete sip alg params
func (o *PartnerConfigDeleteSipAlgParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the partner config delete sip alg params
func (o *PartnerConfigDeleteSipAlgParams) WithID(id string) *PartnerConfigDeleteSipAlgParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the partner config delete sip alg params
func (o *PartnerConfigDeleteSipAlgParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PartnerConfigDeleteSipAlgParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
