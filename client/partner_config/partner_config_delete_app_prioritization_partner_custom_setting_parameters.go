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

// NewPartnerConfigDeleteAppPrioritizationPartnerCustomSettingParams creates a new PartnerConfigDeleteAppPrioritizationPartnerCustomSettingParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPartnerConfigDeleteAppPrioritizationPartnerCustomSettingParams() *PartnerConfigDeleteAppPrioritizationPartnerCustomSettingParams {
	return &PartnerConfigDeleteAppPrioritizationPartnerCustomSettingParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPartnerConfigDeleteAppPrioritizationPartnerCustomSettingParamsWithTimeout creates a new PartnerConfigDeleteAppPrioritizationPartnerCustomSettingParams object
// with the ability to set a timeout on a request.
func NewPartnerConfigDeleteAppPrioritizationPartnerCustomSettingParamsWithTimeout(timeout time.Duration) *PartnerConfigDeleteAppPrioritizationPartnerCustomSettingParams {
	return &PartnerConfigDeleteAppPrioritizationPartnerCustomSettingParams{
		timeout: timeout,
	}
}

// NewPartnerConfigDeleteAppPrioritizationPartnerCustomSettingParamsWithContext creates a new PartnerConfigDeleteAppPrioritizationPartnerCustomSettingParams object
// with the ability to set a context for a request.
func NewPartnerConfigDeleteAppPrioritizationPartnerCustomSettingParamsWithContext(ctx context.Context) *PartnerConfigDeleteAppPrioritizationPartnerCustomSettingParams {
	return &PartnerConfigDeleteAppPrioritizationPartnerCustomSettingParams{
		Context: ctx,
	}
}

// NewPartnerConfigDeleteAppPrioritizationPartnerCustomSettingParamsWithHTTPClient creates a new PartnerConfigDeleteAppPrioritizationPartnerCustomSettingParams object
// with the ability to set a custom HTTPClient for a request.
func NewPartnerConfigDeleteAppPrioritizationPartnerCustomSettingParamsWithHTTPClient(client *http.Client) *PartnerConfigDeleteAppPrioritizationPartnerCustomSettingParams {
	return &PartnerConfigDeleteAppPrioritizationPartnerCustomSettingParams{
		HTTPClient: client,
	}
}

/*
PartnerConfigDeleteAppPrioritizationPartnerCustomSettingParams contains all the parameters to send to the API endpoint

	for the partner config delete app prioritization partner custom setting operation.

	Typically these are written to a http.Request.
*/
type PartnerConfigDeleteAppPrioritizationPartnerCustomSettingParams struct {

	// PartnerID.
	PartnerID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the partner config delete app prioritization partner custom setting params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PartnerConfigDeleteAppPrioritizationPartnerCustomSettingParams) WithDefaults() *PartnerConfigDeleteAppPrioritizationPartnerCustomSettingParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the partner config delete app prioritization partner custom setting params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PartnerConfigDeleteAppPrioritizationPartnerCustomSettingParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the partner config delete app prioritization partner custom setting params
func (o *PartnerConfigDeleteAppPrioritizationPartnerCustomSettingParams) WithTimeout(timeout time.Duration) *PartnerConfigDeleteAppPrioritizationPartnerCustomSettingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the partner config delete app prioritization partner custom setting params
func (o *PartnerConfigDeleteAppPrioritizationPartnerCustomSettingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the partner config delete app prioritization partner custom setting params
func (o *PartnerConfigDeleteAppPrioritizationPartnerCustomSettingParams) WithContext(ctx context.Context) *PartnerConfigDeleteAppPrioritizationPartnerCustomSettingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the partner config delete app prioritization partner custom setting params
func (o *PartnerConfigDeleteAppPrioritizationPartnerCustomSettingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the partner config delete app prioritization partner custom setting params
func (o *PartnerConfigDeleteAppPrioritizationPartnerCustomSettingParams) WithHTTPClient(client *http.Client) *PartnerConfigDeleteAppPrioritizationPartnerCustomSettingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the partner config delete app prioritization partner custom setting params
func (o *PartnerConfigDeleteAppPrioritizationPartnerCustomSettingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPartnerID adds the partnerID to the partner config delete app prioritization partner custom setting params
func (o *PartnerConfigDeleteAppPrioritizationPartnerCustomSettingParams) WithPartnerID(partnerID string) *PartnerConfigDeleteAppPrioritizationPartnerCustomSettingParams {
	o.SetPartnerID(partnerID)
	return o
}

// SetPartnerID adds the partnerId to the partner config delete app prioritization partner custom setting params
func (o *PartnerConfigDeleteAppPrioritizationPartnerCustomSettingParams) SetPartnerID(partnerID string) {
	o.PartnerID = partnerID
}

// WriteToRequest writes these params to a swagger request
func (o *PartnerConfigDeleteAppPrioritizationPartnerCustomSettingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param partnerId
	if err := r.SetPathParam("partnerId", o.PartnerID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}