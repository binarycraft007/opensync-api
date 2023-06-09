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

// NewCustomerPrototypeGetCaptivePortalNetworksParams creates a new CustomerPrototypeGetCaptivePortalNetworksParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypeGetCaptivePortalNetworksParams() *CustomerPrototypeGetCaptivePortalNetworksParams {
	return &CustomerPrototypeGetCaptivePortalNetworksParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypeGetCaptivePortalNetworksParamsWithTimeout creates a new CustomerPrototypeGetCaptivePortalNetworksParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypeGetCaptivePortalNetworksParamsWithTimeout(timeout time.Duration) *CustomerPrototypeGetCaptivePortalNetworksParams {
	return &CustomerPrototypeGetCaptivePortalNetworksParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypeGetCaptivePortalNetworksParamsWithContext creates a new CustomerPrototypeGetCaptivePortalNetworksParams object
// with the ability to set a context for a request.
func NewCustomerPrototypeGetCaptivePortalNetworksParamsWithContext(ctx context.Context) *CustomerPrototypeGetCaptivePortalNetworksParams {
	return &CustomerPrototypeGetCaptivePortalNetworksParams{
		Context: ctx,
	}
}

// NewCustomerPrototypeGetCaptivePortalNetworksParamsWithHTTPClient creates a new CustomerPrototypeGetCaptivePortalNetworksParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypeGetCaptivePortalNetworksParamsWithHTTPClient(client *http.Client) *CustomerPrototypeGetCaptivePortalNetworksParams {
	return &CustomerPrototypeGetCaptivePortalNetworksParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypeGetCaptivePortalNetworksParams contains all the parameters to send to the API endpoint

	for the customer prototype get captive portal networks operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypeGetCaptivePortalNetworksParams struct {

	/* ID.

	   Customer id

	   Format: JSON
	*/
	ID string

	// LocationID.
	//
	// Format: JSON
	LocationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer prototype get captive portal networks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeGetCaptivePortalNetworksParams) WithDefaults() *CustomerPrototypeGetCaptivePortalNetworksParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype get captive portal networks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeGetCaptivePortalNetworksParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype get captive portal networks params
func (o *CustomerPrototypeGetCaptivePortalNetworksParams) WithTimeout(timeout time.Duration) *CustomerPrototypeGetCaptivePortalNetworksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype get captive portal networks params
func (o *CustomerPrototypeGetCaptivePortalNetworksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype get captive portal networks params
func (o *CustomerPrototypeGetCaptivePortalNetworksParams) WithContext(ctx context.Context) *CustomerPrototypeGetCaptivePortalNetworksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype get captive portal networks params
func (o *CustomerPrototypeGetCaptivePortalNetworksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype get captive portal networks params
func (o *CustomerPrototypeGetCaptivePortalNetworksParams) WithHTTPClient(client *http.Client) *CustomerPrototypeGetCaptivePortalNetworksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype get captive portal networks params
func (o *CustomerPrototypeGetCaptivePortalNetworksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the customer prototype get captive portal networks params
func (o *CustomerPrototypeGetCaptivePortalNetworksParams) WithID(id string) *CustomerPrototypeGetCaptivePortalNetworksParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype get captive portal networks params
func (o *CustomerPrototypeGetCaptivePortalNetworksParams) SetID(id string) {
	o.ID = id
}

// WithLocationID adds the locationID to the customer prototype get captive portal networks params
func (o *CustomerPrototypeGetCaptivePortalNetworksParams) WithLocationID(locationID string) *CustomerPrototypeGetCaptivePortalNetworksParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype get captive portal networks params
func (o *CustomerPrototypeGetCaptivePortalNetworksParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypeGetCaptivePortalNetworksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
