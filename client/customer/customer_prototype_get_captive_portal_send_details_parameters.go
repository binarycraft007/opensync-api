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

// NewCustomerPrototypeGetCaptivePortalSendDetailsParams creates a new CustomerPrototypeGetCaptivePortalSendDetailsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypeGetCaptivePortalSendDetailsParams() *CustomerPrototypeGetCaptivePortalSendDetailsParams {
	return &CustomerPrototypeGetCaptivePortalSendDetailsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypeGetCaptivePortalSendDetailsParamsWithTimeout creates a new CustomerPrototypeGetCaptivePortalSendDetailsParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypeGetCaptivePortalSendDetailsParamsWithTimeout(timeout time.Duration) *CustomerPrototypeGetCaptivePortalSendDetailsParams {
	return &CustomerPrototypeGetCaptivePortalSendDetailsParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypeGetCaptivePortalSendDetailsParamsWithContext creates a new CustomerPrototypeGetCaptivePortalSendDetailsParams object
// with the ability to set a context for a request.
func NewCustomerPrototypeGetCaptivePortalSendDetailsParamsWithContext(ctx context.Context) *CustomerPrototypeGetCaptivePortalSendDetailsParams {
	return &CustomerPrototypeGetCaptivePortalSendDetailsParams{
		Context: ctx,
	}
}

// NewCustomerPrototypeGetCaptivePortalSendDetailsParamsWithHTTPClient creates a new CustomerPrototypeGetCaptivePortalSendDetailsParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypeGetCaptivePortalSendDetailsParamsWithHTTPClient(client *http.Client) *CustomerPrototypeGetCaptivePortalSendDetailsParams {
	return &CustomerPrototypeGetCaptivePortalSendDetailsParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypeGetCaptivePortalSendDetailsParams contains all the parameters to send to the API endpoint

	for the customer prototype get captive portal send details operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypeGetCaptivePortalSendDetailsParams struct {

	/* ID.

	   Customer id

	   Format: JSON
	*/
	ID string

	// LocationID.
	//
	// Format: JSON
	LocationID string

	// NetworkID.
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer prototype get captive portal send details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeGetCaptivePortalSendDetailsParams) WithDefaults() *CustomerPrototypeGetCaptivePortalSendDetailsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype get captive portal send details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeGetCaptivePortalSendDetailsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype get captive portal send details params
func (o *CustomerPrototypeGetCaptivePortalSendDetailsParams) WithTimeout(timeout time.Duration) *CustomerPrototypeGetCaptivePortalSendDetailsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype get captive portal send details params
func (o *CustomerPrototypeGetCaptivePortalSendDetailsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype get captive portal send details params
func (o *CustomerPrototypeGetCaptivePortalSendDetailsParams) WithContext(ctx context.Context) *CustomerPrototypeGetCaptivePortalSendDetailsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype get captive portal send details params
func (o *CustomerPrototypeGetCaptivePortalSendDetailsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype get captive portal send details params
func (o *CustomerPrototypeGetCaptivePortalSendDetailsParams) WithHTTPClient(client *http.Client) *CustomerPrototypeGetCaptivePortalSendDetailsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype get captive portal send details params
func (o *CustomerPrototypeGetCaptivePortalSendDetailsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the customer prototype get captive portal send details params
func (o *CustomerPrototypeGetCaptivePortalSendDetailsParams) WithID(id string) *CustomerPrototypeGetCaptivePortalSendDetailsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype get captive portal send details params
func (o *CustomerPrototypeGetCaptivePortalSendDetailsParams) SetID(id string) {
	o.ID = id
}

// WithLocationID adds the locationID to the customer prototype get captive portal send details params
func (o *CustomerPrototypeGetCaptivePortalSendDetailsParams) WithLocationID(locationID string) *CustomerPrototypeGetCaptivePortalSendDetailsParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype get captive portal send details params
func (o *CustomerPrototypeGetCaptivePortalSendDetailsParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WithNetworkID adds the networkID to the customer prototype get captive portal send details params
func (o *CustomerPrototypeGetCaptivePortalSendDetailsParams) WithNetworkID(networkID string) *CustomerPrototypeGetCaptivePortalSendDetailsParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the customer prototype get captive portal send details params
func (o *CustomerPrototypeGetCaptivePortalSendDetailsParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypeGetCaptivePortalSendDetailsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
