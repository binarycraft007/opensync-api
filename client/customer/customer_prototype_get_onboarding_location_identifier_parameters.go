// Code generated by go-swagger; DO NOT EDIT.

package customer

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

// NewCustomerPrototypeGetOnboardingLocationIdentifierParams creates a new CustomerPrototypeGetOnboardingLocationIdentifierParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypeGetOnboardingLocationIdentifierParams() *CustomerPrototypeGetOnboardingLocationIdentifierParams {
	return &CustomerPrototypeGetOnboardingLocationIdentifierParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypeGetOnboardingLocationIdentifierParamsWithTimeout creates a new CustomerPrototypeGetOnboardingLocationIdentifierParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypeGetOnboardingLocationIdentifierParamsWithTimeout(timeout time.Duration) *CustomerPrototypeGetOnboardingLocationIdentifierParams {
	return &CustomerPrototypeGetOnboardingLocationIdentifierParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypeGetOnboardingLocationIdentifierParamsWithContext creates a new CustomerPrototypeGetOnboardingLocationIdentifierParams object
// with the ability to set a context for a request.
func NewCustomerPrototypeGetOnboardingLocationIdentifierParamsWithContext(ctx context.Context) *CustomerPrototypeGetOnboardingLocationIdentifierParams {
	return &CustomerPrototypeGetOnboardingLocationIdentifierParams{
		Context: ctx,
	}
}

// NewCustomerPrototypeGetOnboardingLocationIdentifierParamsWithHTTPClient creates a new CustomerPrototypeGetOnboardingLocationIdentifierParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypeGetOnboardingLocationIdentifierParamsWithHTTPClient(client *http.Client) *CustomerPrototypeGetOnboardingLocationIdentifierParams {
	return &CustomerPrototypeGetOnboardingLocationIdentifierParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypeGetOnboardingLocationIdentifierParams contains all the parameters to send to the API endpoint

	for the customer prototype get onboarding location identifier operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypeGetOnboardingLocationIdentifierParams struct {

	/* ID.

	   Customer id

	   Format: JSON
	*/
	ID string

	// LocationID.
	LocationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer prototype get onboarding location identifier params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeGetOnboardingLocationIdentifierParams) WithDefaults() *CustomerPrototypeGetOnboardingLocationIdentifierParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype get onboarding location identifier params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeGetOnboardingLocationIdentifierParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype get onboarding location identifier params
func (o *CustomerPrototypeGetOnboardingLocationIdentifierParams) WithTimeout(timeout time.Duration) *CustomerPrototypeGetOnboardingLocationIdentifierParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype get onboarding location identifier params
func (o *CustomerPrototypeGetOnboardingLocationIdentifierParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype get onboarding location identifier params
func (o *CustomerPrototypeGetOnboardingLocationIdentifierParams) WithContext(ctx context.Context) *CustomerPrototypeGetOnboardingLocationIdentifierParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype get onboarding location identifier params
func (o *CustomerPrototypeGetOnboardingLocationIdentifierParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype get onboarding location identifier params
func (o *CustomerPrototypeGetOnboardingLocationIdentifierParams) WithHTTPClient(client *http.Client) *CustomerPrototypeGetOnboardingLocationIdentifierParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype get onboarding location identifier params
func (o *CustomerPrototypeGetOnboardingLocationIdentifierParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the customer prototype get onboarding location identifier params
func (o *CustomerPrototypeGetOnboardingLocationIdentifierParams) WithID(id string) *CustomerPrototypeGetOnboardingLocationIdentifierParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype get onboarding location identifier params
func (o *CustomerPrototypeGetOnboardingLocationIdentifierParams) SetID(id string) {
	o.ID = id
}

// WithLocationID adds the locationID to the customer prototype get onboarding location identifier params
func (o *CustomerPrototypeGetOnboardingLocationIdentifierParams) WithLocationID(locationID string) *CustomerPrototypeGetOnboardingLocationIdentifierParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype get onboarding location identifier params
func (o *CustomerPrototypeGetOnboardingLocationIdentifierParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypeGetOnboardingLocationIdentifierParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param locationId
	if err := r.SetPathParam("locationId", o.LocationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
