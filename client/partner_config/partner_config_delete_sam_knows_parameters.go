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

// NewPartnerConfigDeleteSamKnowsParams creates a new PartnerConfigDeleteSamKnowsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPartnerConfigDeleteSamKnowsParams() *PartnerConfigDeleteSamKnowsParams {
	return &PartnerConfigDeleteSamKnowsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPartnerConfigDeleteSamKnowsParamsWithTimeout creates a new PartnerConfigDeleteSamKnowsParams object
// with the ability to set a timeout on a request.
func NewPartnerConfigDeleteSamKnowsParamsWithTimeout(timeout time.Duration) *PartnerConfigDeleteSamKnowsParams {
	return &PartnerConfigDeleteSamKnowsParams{
		timeout: timeout,
	}
}

// NewPartnerConfigDeleteSamKnowsParamsWithContext creates a new PartnerConfigDeleteSamKnowsParams object
// with the ability to set a context for a request.
func NewPartnerConfigDeleteSamKnowsParamsWithContext(ctx context.Context) *PartnerConfigDeleteSamKnowsParams {
	return &PartnerConfigDeleteSamKnowsParams{
		Context: ctx,
	}
}

// NewPartnerConfigDeleteSamKnowsParamsWithHTTPClient creates a new PartnerConfigDeleteSamKnowsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPartnerConfigDeleteSamKnowsParamsWithHTTPClient(client *http.Client) *PartnerConfigDeleteSamKnowsParams {
	return &PartnerConfigDeleteSamKnowsParams{
		HTTPClient: client,
	}
}

/*
PartnerConfigDeleteSamKnowsParams contains all the parameters to send to the API endpoint

	for the partner config delete sam knows operation.

	Typically these are written to a http.Request.
*/
type PartnerConfigDeleteSamKnowsParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the partner config delete sam knows params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PartnerConfigDeleteSamKnowsParams) WithDefaults() *PartnerConfigDeleteSamKnowsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the partner config delete sam knows params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PartnerConfigDeleteSamKnowsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the partner config delete sam knows params
func (o *PartnerConfigDeleteSamKnowsParams) WithTimeout(timeout time.Duration) *PartnerConfigDeleteSamKnowsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the partner config delete sam knows params
func (o *PartnerConfigDeleteSamKnowsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the partner config delete sam knows params
func (o *PartnerConfigDeleteSamKnowsParams) WithContext(ctx context.Context) *PartnerConfigDeleteSamKnowsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the partner config delete sam knows params
func (o *PartnerConfigDeleteSamKnowsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the partner config delete sam knows params
func (o *PartnerConfigDeleteSamKnowsParams) WithHTTPClient(client *http.Client) *PartnerConfigDeleteSamKnowsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the partner config delete sam knows params
func (o *PartnerConfigDeleteSamKnowsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the partner config delete sam knows params
func (o *PartnerConfigDeleteSamKnowsParams) WithID(id string) *PartnerConfigDeleteSamKnowsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the partner config delete sam knows params
func (o *PartnerConfigDeleteSamKnowsParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PartnerConfigDeleteSamKnowsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
