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

// NewPartnerConfigPatchWorkpassCustomerSupportConfigurationsParams creates a new PartnerConfigPatchWorkpassCustomerSupportConfigurationsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPartnerConfigPatchWorkpassCustomerSupportConfigurationsParams() *PartnerConfigPatchWorkpassCustomerSupportConfigurationsParams {
	return &PartnerConfigPatchWorkpassCustomerSupportConfigurationsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPartnerConfigPatchWorkpassCustomerSupportConfigurationsParamsWithTimeout creates a new PartnerConfigPatchWorkpassCustomerSupportConfigurationsParams object
// with the ability to set a timeout on a request.
func NewPartnerConfigPatchWorkpassCustomerSupportConfigurationsParamsWithTimeout(timeout time.Duration) *PartnerConfigPatchWorkpassCustomerSupportConfigurationsParams {
	return &PartnerConfigPatchWorkpassCustomerSupportConfigurationsParams{
		timeout: timeout,
	}
}

// NewPartnerConfigPatchWorkpassCustomerSupportConfigurationsParamsWithContext creates a new PartnerConfigPatchWorkpassCustomerSupportConfigurationsParams object
// with the ability to set a context for a request.
func NewPartnerConfigPatchWorkpassCustomerSupportConfigurationsParamsWithContext(ctx context.Context) *PartnerConfigPatchWorkpassCustomerSupportConfigurationsParams {
	return &PartnerConfigPatchWorkpassCustomerSupportConfigurationsParams{
		Context: ctx,
	}
}

// NewPartnerConfigPatchWorkpassCustomerSupportConfigurationsParamsWithHTTPClient creates a new PartnerConfigPatchWorkpassCustomerSupportConfigurationsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPartnerConfigPatchWorkpassCustomerSupportConfigurationsParamsWithHTTPClient(client *http.Client) *PartnerConfigPatchWorkpassCustomerSupportConfigurationsParams {
	return &PartnerConfigPatchWorkpassCustomerSupportConfigurationsParams{
		HTTPClient: client,
	}
}

/*
PartnerConfigPatchWorkpassCustomerSupportConfigurationsParams contains all the parameters to send to the API endpoint

	for the partner config patch workpass customer support configurations operation.

	Typically these are written to a http.Request.
*/
type PartnerConfigPatchWorkpassCustomerSupportConfigurationsParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the partner config patch workpass customer support configurations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PartnerConfigPatchWorkpassCustomerSupportConfigurationsParams) WithDefaults() *PartnerConfigPatchWorkpassCustomerSupportConfigurationsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the partner config patch workpass customer support configurations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PartnerConfigPatchWorkpassCustomerSupportConfigurationsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the partner config patch workpass customer support configurations params
func (o *PartnerConfigPatchWorkpassCustomerSupportConfigurationsParams) WithTimeout(timeout time.Duration) *PartnerConfigPatchWorkpassCustomerSupportConfigurationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the partner config patch workpass customer support configurations params
func (o *PartnerConfigPatchWorkpassCustomerSupportConfigurationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the partner config patch workpass customer support configurations params
func (o *PartnerConfigPatchWorkpassCustomerSupportConfigurationsParams) WithContext(ctx context.Context) *PartnerConfigPatchWorkpassCustomerSupportConfigurationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the partner config patch workpass customer support configurations params
func (o *PartnerConfigPatchWorkpassCustomerSupportConfigurationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the partner config patch workpass customer support configurations params
func (o *PartnerConfigPatchWorkpassCustomerSupportConfigurationsParams) WithHTTPClient(client *http.Client) *PartnerConfigPatchWorkpassCustomerSupportConfigurationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the partner config patch workpass customer support configurations params
func (o *PartnerConfigPatchWorkpassCustomerSupportConfigurationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the partner config patch workpass customer support configurations params
func (o *PartnerConfigPatchWorkpassCustomerSupportConfigurationsParams) WithID(id string) *PartnerConfigPatchWorkpassCustomerSupportConfigurationsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the partner config patch workpass customer support configurations params
func (o *PartnerConfigPatchWorkpassCustomerSupportConfigurationsParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PartnerConfigPatchWorkpassCustomerSupportConfigurationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
