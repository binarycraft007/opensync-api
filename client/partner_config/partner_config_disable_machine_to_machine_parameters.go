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

// NewPartnerConfigDisableMachineToMachineParams creates a new PartnerConfigDisableMachineToMachineParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPartnerConfigDisableMachineToMachineParams() *PartnerConfigDisableMachineToMachineParams {
	return &PartnerConfigDisableMachineToMachineParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPartnerConfigDisableMachineToMachineParamsWithTimeout creates a new PartnerConfigDisableMachineToMachineParams object
// with the ability to set a timeout on a request.
func NewPartnerConfigDisableMachineToMachineParamsWithTimeout(timeout time.Duration) *PartnerConfigDisableMachineToMachineParams {
	return &PartnerConfigDisableMachineToMachineParams{
		timeout: timeout,
	}
}

// NewPartnerConfigDisableMachineToMachineParamsWithContext creates a new PartnerConfigDisableMachineToMachineParams object
// with the ability to set a context for a request.
func NewPartnerConfigDisableMachineToMachineParamsWithContext(ctx context.Context) *PartnerConfigDisableMachineToMachineParams {
	return &PartnerConfigDisableMachineToMachineParams{
		Context: ctx,
	}
}

// NewPartnerConfigDisableMachineToMachineParamsWithHTTPClient creates a new PartnerConfigDisableMachineToMachineParams object
// with the ability to set a custom HTTPClient for a request.
func NewPartnerConfigDisableMachineToMachineParamsWithHTTPClient(client *http.Client) *PartnerConfigDisableMachineToMachineParams {
	return &PartnerConfigDisableMachineToMachineParams{
		HTTPClient: client,
	}
}

/*
PartnerConfigDisableMachineToMachineParams contains all the parameters to send to the API endpoint

	for the partner config disable machine to machine operation.

	Typically these are written to a http.Request.
*/
type PartnerConfigDisableMachineToMachineParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the partner config disable machine to machine params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PartnerConfigDisableMachineToMachineParams) WithDefaults() *PartnerConfigDisableMachineToMachineParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the partner config disable machine to machine params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PartnerConfigDisableMachineToMachineParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the partner config disable machine to machine params
func (o *PartnerConfigDisableMachineToMachineParams) WithTimeout(timeout time.Duration) *PartnerConfigDisableMachineToMachineParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the partner config disable machine to machine params
func (o *PartnerConfigDisableMachineToMachineParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the partner config disable machine to machine params
func (o *PartnerConfigDisableMachineToMachineParams) WithContext(ctx context.Context) *PartnerConfigDisableMachineToMachineParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the partner config disable machine to machine params
func (o *PartnerConfigDisableMachineToMachineParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the partner config disable machine to machine params
func (o *PartnerConfigDisableMachineToMachineParams) WithHTTPClient(client *http.Client) *PartnerConfigDisableMachineToMachineParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the partner config disable machine to machine params
func (o *PartnerConfigDisableMachineToMachineParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the partner config disable machine to machine params
func (o *PartnerConfigDisableMachineToMachineParams) WithID(id string) *PartnerConfigDisableMachineToMachineParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the partner config disable machine to machine params
func (o *PartnerConfigDisableMachineToMachineParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PartnerConfigDisableMachineToMachineParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
