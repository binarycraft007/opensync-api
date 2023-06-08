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
	"github.com/go-openapi/swag"
)

// NewCustomerPrototypeGetCaptivePortalSendDetailsDirectParams creates a new CustomerPrototypeGetCaptivePortalSendDetailsDirectParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypeGetCaptivePortalSendDetailsDirectParams() *CustomerPrototypeGetCaptivePortalSendDetailsDirectParams {
	return &CustomerPrototypeGetCaptivePortalSendDetailsDirectParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypeGetCaptivePortalSendDetailsDirectParamsWithTimeout creates a new CustomerPrototypeGetCaptivePortalSendDetailsDirectParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypeGetCaptivePortalSendDetailsDirectParamsWithTimeout(timeout time.Duration) *CustomerPrototypeGetCaptivePortalSendDetailsDirectParams {
	return &CustomerPrototypeGetCaptivePortalSendDetailsDirectParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypeGetCaptivePortalSendDetailsDirectParamsWithContext creates a new CustomerPrototypeGetCaptivePortalSendDetailsDirectParams object
// with the ability to set a context for a request.
func NewCustomerPrototypeGetCaptivePortalSendDetailsDirectParamsWithContext(ctx context.Context) *CustomerPrototypeGetCaptivePortalSendDetailsDirectParams {
	return &CustomerPrototypeGetCaptivePortalSendDetailsDirectParams{
		Context: ctx,
	}
}

// NewCustomerPrototypeGetCaptivePortalSendDetailsDirectParamsWithHTTPClient creates a new CustomerPrototypeGetCaptivePortalSendDetailsDirectParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypeGetCaptivePortalSendDetailsDirectParamsWithHTTPClient(client *http.Client) *CustomerPrototypeGetCaptivePortalSendDetailsDirectParams {
	return &CustomerPrototypeGetCaptivePortalSendDetailsDirectParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypeGetCaptivePortalSendDetailsDirectParams contains all the parameters to send to the API endpoint

	for the customer prototype get captive portal send details direct operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypeGetCaptivePortalSendDetailsDirectParams struct {

	// Duration.
	//
	// Format: double
	Duration *float64

	/* ID.

	   Customer id

	   Format: JSON
	*/
	ID string

	// Limit.
	//
	// Format: double
	Limit *float64

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

// WithDefaults hydrates default values in the customer prototype get captive portal send details direct params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeGetCaptivePortalSendDetailsDirectParams) WithDefaults() *CustomerPrototypeGetCaptivePortalSendDetailsDirectParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype get captive portal send details direct params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeGetCaptivePortalSendDetailsDirectParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype get captive portal send details direct params
func (o *CustomerPrototypeGetCaptivePortalSendDetailsDirectParams) WithTimeout(timeout time.Duration) *CustomerPrototypeGetCaptivePortalSendDetailsDirectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype get captive portal send details direct params
func (o *CustomerPrototypeGetCaptivePortalSendDetailsDirectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype get captive portal send details direct params
func (o *CustomerPrototypeGetCaptivePortalSendDetailsDirectParams) WithContext(ctx context.Context) *CustomerPrototypeGetCaptivePortalSendDetailsDirectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype get captive portal send details direct params
func (o *CustomerPrototypeGetCaptivePortalSendDetailsDirectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype get captive portal send details direct params
func (o *CustomerPrototypeGetCaptivePortalSendDetailsDirectParams) WithHTTPClient(client *http.Client) *CustomerPrototypeGetCaptivePortalSendDetailsDirectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype get captive portal send details direct params
func (o *CustomerPrototypeGetCaptivePortalSendDetailsDirectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDuration adds the duration to the customer prototype get captive portal send details direct params
func (o *CustomerPrototypeGetCaptivePortalSendDetailsDirectParams) WithDuration(duration *float64) *CustomerPrototypeGetCaptivePortalSendDetailsDirectParams {
	o.SetDuration(duration)
	return o
}

// SetDuration adds the duration to the customer prototype get captive portal send details direct params
func (o *CustomerPrototypeGetCaptivePortalSendDetailsDirectParams) SetDuration(duration *float64) {
	o.Duration = duration
}

// WithID adds the id to the customer prototype get captive portal send details direct params
func (o *CustomerPrototypeGetCaptivePortalSendDetailsDirectParams) WithID(id string) *CustomerPrototypeGetCaptivePortalSendDetailsDirectParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype get captive portal send details direct params
func (o *CustomerPrototypeGetCaptivePortalSendDetailsDirectParams) SetID(id string) {
	o.ID = id
}

// WithLimit adds the limit to the customer prototype get captive portal send details direct params
func (o *CustomerPrototypeGetCaptivePortalSendDetailsDirectParams) WithLimit(limit *float64) *CustomerPrototypeGetCaptivePortalSendDetailsDirectParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the customer prototype get captive portal send details direct params
func (o *CustomerPrototypeGetCaptivePortalSendDetailsDirectParams) SetLimit(limit *float64) {
	o.Limit = limit
}

// WithLocationID adds the locationID to the customer prototype get captive portal send details direct params
func (o *CustomerPrototypeGetCaptivePortalSendDetailsDirectParams) WithLocationID(locationID string) *CustomerPrototypeGetCaptivePortalSendDetailsDirectParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype get captive portal send details direct params
func (o *CustomerPrototypeGetCaptivePortalSendDetailsDirectParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WithNetworkID adds the networkID to the customer prototype get captive portal send details direct params
func (o *CustomerPrototypeGetCaptivePortalSendDetailsDirectParams) WithNetworkID(networkID string) *CustomerPrototypeGetCaptivePortalSendDetailsDirectParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the customer prototype get captive portal send details direct params
func (o *CustomerPrototypeGetCaptivePortalSendDetailsDirectParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypeGetCaptivePortalSendDetailsDirectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Duration != nil {

		// query param duration
		var qrDuration float64

		if o.Duration != nil {
			qrDuration = *o.Duration
		}
		qDuration := swag.FormatFloat64(qrDuration)
		if qDuration != "" {

			if err := r.SetQueryParam("duration", qDuration); err != nil {
				return err
			}
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit float64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatFloat64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
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