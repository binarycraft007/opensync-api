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

// NewPartnerConfigUpdateFeatureFlagsParams creates a new PartnerConfigUpdateFeatureFlagsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPartnerConfigUpdateFeatureFlagsParams() *PartnerConfigUpdateFeatureFlagsParams {
	return &PartnerConfigUpdateFeatureFlagsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPartnerConfigUpdateFeatureFlagsParamsWithTimeout creates a new PartnerConfigUpdateFeatureFlagsParams object
// with the ability to set a timeout on a request.
func NewPartnerConfigUpdateFeatureFlagsParamsWithTimeout(timeout time.Duration) *PartnerConfigUpdateFeatureFlagsParams {
	return &PartnerConfigUpdateFeatureFlagsParams{
		timeout: timeout,
	}
}

// NewPartnerConfigUpdateFeatureFlagsParamsWithContext creates a new PartnerConfigUpdateFeatureFlagsParams object
// with the ability to set a context for a request.
func NewPartnerConfigUpdateFeatureFlagsParamsWithContext(ctx context.Context) *PartnerConfigUpdateFeatureFlagsParams {
	return &PartnerConfigUpdateFeatureFlagsParams{
		Context: ctx,
	}
}

// NewPartnerConfigUpdateFeatureFlagsParamsWithHTTPClient creates a new PartnerConfigUpdateFeatureFlagsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPartnerConfigUpdateFeatureFlagsParamsWithHTTPClient(client *http.Client) *PartnerConfigUpdateFeatureFlagsParams {
	return &PartnerConfigUpdateFeatureFlagsParams{
		HTTPClient: client,
	}
}

/*
PartnerConfigUpdateFeatureFlagsParams contains all the parameters to send to the API endpoint

	for the partner config update feature flags operation.

	Typically these are written to a http.Request.
*/
type PartnerConfigUpdateFeatureFlagsParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the partner config update feature flags params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PartnerConfigUpdateFeatureFlagsParams) WithDefaults() *PartnerConfigUpdateFeatureFlagsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the partner config update feature flags params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PartnerConfigUpdateFeatureFlagsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the partner config update feature flags params
func (o *PartnerConfigUpdateFeatureFlagsParams) WithTimeout(timeout time.Duration) *PartnerConfigUpdateFeatureFlagsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the partner config update feature flags params
func (o *PartnerConfigUpdateFeatureFlagsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the partner config update feature flags params
func (o *PartnerConfigUpdateFeatureFlagsParams) WithContext(ctx context.Context) *PartnerConfigUpdateFeatureFlagsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the partner config update feature flags params
func (o *PartnerConfigUpdateFeatureFlagsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the partner config update feature flags params
func (o *PartnerConfigUpdateFeatureFlagsParams) WithHTTPClient(client *http.Client) *PartnerConfigUpdateFeatureFlagsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the partner config update feature flags params
func (o *PartnerConfigUpdateFeatureFlagsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the partner config update feature flags params
func (o *PartnerConfigUpdateFeatureFlagsParams) WithID(id string) *PartnerConfigUpdateFeatureFlagsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the partner config update feature flags params
func (o *PartnerConfigUpdateFeatureFlagsParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PartnerConfigUpdateFeatureFlagsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
