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

// NewPartnerConfigUpdateCaptivePortalConfigParams creates a new PartnerConfigUpdateCaptivePortalConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPartnerConfigUpdateCaptivePortalConfigParams() *PartnerConfigUpdateCaptivePortalConfigParams {
	return &PartnerConfigUpdateCaptivePortalConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPartnerConfigUpdateCaptivePortalConfigParamsWithTimeout creates a new PartnerConfigUpdateCaptivePortalConfigParams object
// with the ability to set a timeout on a request.
func NewPartnerConfigUpdateCaptivePortalConfigParamsWithTimeout(timeout time.Duration) *PartnerConfigUpdateCaptivePortalConfigParams {
	return &PartnerConfigUpdateCaptivePortalConfigParams{
		timeout: timeout,
	}
}

// NewPartnerConfigUpdateCaptivePortalConfigParamsWithContext creates a new PartnerConfigUpdateCaptivePortalConfigParams object
// with the ability to set a context for a request.
func NewPartnerConfigUpdateCaptivePortalConfigParamsWithContext(ctx context.Context) *PartnerConfigUpdateCaptivePortalConfigParams {
	return &PartnerConfigUpdateCaptivePortalConfigParams{
		Context: ctx,
	}
}

// NewPartnerConfigUpdateCaptivePortalConfigParamsWithHTTPClient creates a new PartnerConfigUpdateCaptivePortalConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewPartnerConfigUpdateCaptivePortalConfigParamsWithHTTPClient(client *http.Client) *PartnerConfigUpdateCaptivePortalConfigParams {
	return &PartnerConfigUpdateCaptivePortalConfigParams{
		HTTPClient: client,
	}
}

/*
PartnerConfigUpdateCaptivePortalConfigParams contains all the parameters to send to the API endpoint

	for the partner config update captive portal config operation.

	Typically these are written to a http.Request.
*/
type PartnerConfigUpdateCaptivePortalConfigParams struct {

	// DefaultLanguage.
	DefaultLanguage *string

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the partner config update captive portal config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PartnerConfigUpdateCaptivePortalConfigParams) WithDefaults() *PartnerConfigUpdateCaptivePortalConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the partner config update captive portal config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PartnerConfigUpdateCaptivePortalConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the partner config update captive portal config params
func (o *PartnerConfigUpdateCaptivePortalConfigParams) WithTimeout(timeout time.Duration) *PartnerConfigUpdateCaptivePortalConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the partner config update captive portal config params
func (o *PartnerConfigUpdateCaptivePortalConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the partner config update captive portal config params
func (o *PartnerConfigUpdateCaptivePortalConfigParams) WithContext(ctx context.Context) *PartnerConfigUpdateCaptivePortalConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the partner config update captive portal config params
func (o *PartnerConfigUpdateCaptivePortalConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the partner config update captive portal config params
func (o *PartnerConfigUpdateCaptivePortalConfigParams) WithHTTPClient(client *http.Client) *PartnerConfigUpdateCaptivePortalConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the partner config update captive portal config params
func (o *PartnerConfigUpdateCaptivePortalConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDefaultLanguage adds the defaultLanguage to the partner config update captive portal config params
func (o *PartnerConfigUpdateCaptivePortalConfigParams) WithDefaultLanguage(defaultLanguage *string) *PartnerConfigUpdateCaptivePortalConfigParams {
	o.SetDefaultLanguage(defaultLanguage)
	return o
}

// SetDefaultLanguage adds the defaultLanguage to the partner config update captive portal config params
func (o *PartnerConfigUpdateCaptivePortalConfigParams) SetDefaultLanguage(defaultLanguage *string) {
	o.DefaultLanguage = defaultLanguage
}

// WithID adds the id to the partner config update captive portal config params
func (o *PartnerConfigUpdateCaptivePortalConfigParams) WithID(id string) *PartnerConfigUpdateCaptivePortalConfigParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the partner config update captive portal config params
func (o *PartnerConfigUpdateCaptivePortalConfigParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PartnerConfigUpdateCaptivePortalConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DefaultLanguage != nil {

		// form param defaultLanguage
		var frDefaultLanguage string
		if o.DefaultLanguage != nil {
			frDefaultLanguage = *o.DefaultLanguage
		}
		fDefaultLanguage := frDefaultLanguage
		if fDefaultLanguage != "" {
			if err := r.SetFormParam("defaultLanguage", fDefaultLanguage); err != nil {
				return err
			}
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
