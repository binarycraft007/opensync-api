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

// NewPartnerConfigUpdatePreCacSchedulerParams creates a new PartnerConfigUpdatePreCacSchedulerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPartnerConfigUpdatePreCacSchedulerParams() *PartnerConfigUpdatePreCacSchedulerParams {
	return &PartnerConfigUpdatePreCacSchedulerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPartnerConfigUpdatePreCacSchedulerParamsWithTimeout creates a new PartnerConfigUpdatePreCacSchedulerParams object
// with the ability to set a timeout on a request.
func NewPartnerConfigUpdatePreCacSchedulerParamsWithTimeout(timeout time.Duration) *PartnerConfigUpdatePreCacSchedulerParams {
	return &PartnerConfigUpdatePreCacSchedulerParams{
		timeout: timeout,
	}
}

// NewPartnerConfigUpdatePreCacSchedulerParamsWithContext creates a new PartnerConfigUpdatePreCacSchedulerParams object
// with the ability to set a context for a request.
func NewPartnerConfigUpdatePreCacSchedulerParamsWithContext(ctx context.Context) *PartnerConfigUpdatePreCacSchedulerParams {
	return &PartnerConfigUpdatePreCacSchedulerParams{
		Context: ctx,
	}
}

// NewPartnerConfigUpdatePreCacSchedulerParamsWithHTTPClient creates a new PartnerConfigUpdatePreCacSchedulerParams object
// with the ability to set a custom HTTPClient for a request.
func NewPartnerConfigUpdatePreCacSchedulerParamsWithHTTPClient(client *http.Client) *PartnerConfigUpdatePreCacSchedulerParams {
	return &PartnerConfigUpdatePreCacSchedulerParams{
		HTTPClient: client,
	}
}

/*
PartnerConfigUpdatePreCacSchedulerParams contains all the parameters to send to the API endpoint

	for the partner config update pre cac scheduler operation.

	Typically these are written to a http.Request.
*/
type PartnerConfigUpdatePreCacSchedulerParams struct {

	// Config.
	//
	// Format: JSON
	Config string

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the partner config update pre cac scheduler params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PartnerConfigUpdatePreCacSchedulerParams) WithDefaults() *PartnerConfigUpdatePreCacSchedulerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the partner config update pre cac scheduler params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PartnerConfigUpdatePreCacSchedulerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the partner config update pre cac scheduler params
func (o *PartnerConfigUpdatePreCacSchedulerParams) WithTimeout(timeout time.Duration) *PartnerConfigUpdatePreCacSchedulerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the partner config update pre cac scheduler params
func (o *PartnerConfigUpdatePreCacSchedulerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the partner config update pre cac scheduler params
func (o *PartnerConfigUpdatePreCacSchedulerParams) WithContext(ctx context.Context) *PartnerConfigUpdatePreCacSchedulerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the partner config update pre cac scheduler params
func (o *PartnerConfigUpdatePreCacSchedulerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the partner config update pre cac scheduler params
func (o *PartnerConfigUpdatePreCacSchedulerParams) WithHTTPClient(client *http.Client) *PartnerConfigUpdatePreCacSchedulerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the partner config update pre cac scheduler params
func (o *PartnerConfigUpdatePreCacSchedulerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfig adds the config to the partner config update pre cac scheduler params
func (o *PartnerConfigUpdatePreCacSchedulerParams) WithConfig(config string) *PartnerConfigUpdatePreCacSchedulerParams {
	o.SetConfig(config)
	return o
}

// SetConfig adds the config to the partner config update pre cac scheduler params
func (o *PartnerConfigUpdatePreCacSchedulerParams) SetConfig(config string) {
	o.Config = config
}

// WithID adds the id to the partner config update pre cac scheduler params
func (o *PartnerConfigUpdatePreCacSchedulerParams) WithID(id string) *PartnerConfigUpdatePreCacSchedulerParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the partner config update pre cac scheduler params
func (o *PartnerConfigUpdatePreCacSchedulerParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PartnerConfigUpdatePreCacSchedulerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// form param config
	frConfig := o.Config
	fConfig := frConfig
	if fConfig != "" {
		if err := r.SetFormParam("config", fConfig); err != nil {
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
