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
	"github.com/go-openapi/swag"
)

// NewPartnerConfigGenerateMachineToMachineTokenParams creates a new PartnerConfigGenerateMachineToMachineTokenParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPartnerConfigGenerateMachineToMachineTokenParams() *PartnerConfigGenerateMachineToMachineTokenParams {
	return &PartnerConfigGenerateMachineToMachineTokenParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPartnerConfigGenerateMachineToMachineTokenParamsWithTimeout creates a new PartnerConfigGenerateMachineToMachineTokenParams object
// with the ability to set a timeout on a request.
func NewPartnerConfigGenerateMachineToMachineTokenParamsWithTimeout(timeout time.Duration) *PartnerConfigGenerateMachineToMachineTokenParams {
	return &PartnerConfigGenerateMachineToMachineTokenParams{
		timeout: timeout,
	}
}

// NewPartnerConfigGenerateMachineToMachineTokenParamsWithContext creates a new PartnerConfigGenerateMachineToMachineTokenParams object
// with the ability to set a context for a request.
func NewPartnerConfigGenerateMachineToMachineTokenParamsWithContext(ctx context.Context) *PartnerConfigGenerateMachineToMachineTokenParams {
	return &PartnerConfigGenerateMachineToMachineTokenParams{
		Context: ctx,
	}
}

// NewPartnerConfigGenerateMachineToMachineTokenParamsWithHTTPClient creates a new PartnerConfigGenerateMachineToMachineTokenParams object
// with the ability to set a custom HTTPClient for a request.
func NewPartnerConfigGenerateMachineToMachineTokenParamsWithHTTPClient(client *http.Client) *PartnerConfigGenerateMachineToMachineTokenParams {
	return &PartnerConfigGenerateMachineToMachineTokenParams{
		HTTPClient: client,
	}
}

/*
PartnerConfigGenerateMachineToMachineTokenParams contains all the parameters to send to the API endpoint

	for the partner config generate machine to machine token operation.

	Typically these are written to a http.Request.
*/
type PartnerConfigGenerateMachineToMachineTokenParams struct {

	// ID.
	ID string

	// TokenName.
	TokenName string

	// TokenTTLSeconds.
	//
	// Format: double
	TokenTTLSeconds *float64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the partner config generate machine to machine token params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PartnerConfigGenerateMachineToMachineTokenParams) WithDefaults() *PartnerConfigGenerateMachineToMachineTokenParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the partner config generate machine to machine token params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PartnerConfigGenerateMachineToMachineTokenParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the partner config generate machine to machine token params
func (o *PartnerConfigGenerateMachineToMachineTokenParams) WithTimeout(timeout time.Duration) *PartnerConfigGenerateMachineToMachineTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the partner config generate machine to machine token params
func (o *PartnerConfigGenerateMachineToMachineTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the partner config generate machine to machine token params
func (o *PartnerConfigGenerateMachineToMachineTokenParams) WithContext(ctx context.Context) *PartnerConfigGenerateMachineToMachineTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the partner config generate machine to machine token params
func (o *PartnerConfigGenerateMachineToMachineTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the partner config generate machine to machine token params
func (o *PartnerConfigGenerateMachineToMachineTokenParams) WithHTTPClient(client *http.Client) *PartnerConfigGenerateMachineToMachineTokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the partner config generate machine to machine token params
func (o *PartnerConfigGenerateMachineToMachineTokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the partner config generate machine to machine token params
func (o *PartnerConfigGenerateMachineToMachineTokenParams) WithID(id string) *PartnerConfigGenerateMachineToMachineTokenParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the partner config generate machine to machine token params
func (o *PartnerConfigGenerateMachineToMachineTokenParams) SetID(id string) {
	o.ID = id
}

// WithTokenName adds the tokenName to the partner config generate machine to machine token params
func (o *PartnerConfigGenerateMachineToMachineTokenParams) WithTokenName(tokenName string) *PartnerConfigGenerateMachineToMachineTokenParams {
	o.SetTokenName(tokenName)
	return o
}

// SetTokenName adds the tokenName to the partner config generate machine to machine token params
func (o *PartnerConfigGenerateMachineToMachineTokenParams) SetTokenName(tokenName string) {
	o.TokenName = tokenName
}

// WithTokenTTLSeconds adds the tokenTTLSeconds to the partner config generate machine to machine token params
func (o *PartnerConfigGenerateMachineToMachineTokenParams) WithTokenTTLSeconds(tokenTTLSeconds *float64) *PartnerConfigGenerateMachineToMachineTokenParams {
	o.SetTokenTTLSeconds(tokenTTLSeconds)
	return o
}

// SetTokenTTLSeconds adds the tokenTtlSeconds to the partner config generate machine to machine token params
func (o *PartnerConfigGenerateMachineToMachineTokenParams) SetTokenTTLSeconds(tokenTTLSeconds *float64) {
	o.TokenTTLSeconds = tokenTTLSeconds
}

// WriteToRequest writes these params to a swagger request
func (o *PartnerConfigGenerateMachineToMachineTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// form param tokenName
	frTokenName := o.TokenName
	fTokenName := frTokenName
	if fTokenName != "" {
		if err := r.SetFormParam("tokenName", fTokenName); err != nil {
			return err
		}
	}

	if o.TokenTTLSeconds != nil {

		// form param tokenTTLSeconds
		var frTokenTTLSeconds float64
		if o.TokenTTLSeconds != nil {
			frTokenTTLSeconds = *o.TokenTTLSeconds
		}
		fTokenTTLSeconds := swag.FormatFloat64(frTokenTTLSeconds)
		if fTokenTTLSeconds != "" {
			if err := r.SetFormParam("tokenTTLSeconds", fTokenTTLSeconds); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
