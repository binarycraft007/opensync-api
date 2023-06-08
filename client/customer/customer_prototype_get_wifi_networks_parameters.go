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

// NewCustomerPrototypeGetWifiNetworksParams creates a new CustomerPrototypeGetWifiNetworksParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypeGetWifiNetworksParams() *CustomerPrototypeGetWifiNetworksParams {
	return &CustomerPrototypeGetWifiNetworksParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypeGetWifiNetworksParamsWithTimeout creates a new CustomerPrototypeGetWifiNetworksParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypeGetWifiNetworksParamsWithTimeout(timeout time.Duration) *CustomerPrototypeGetWifiNetworksParams {
	return &CustomerPrototypeGetWifiNetworksParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypeGetWifiNetworksParamsWithContext creates a new CustomerPrototypeGetWifiNetworksParams object
// with the ability to set a context for a request.
func NewCustomerPrototypeGetWifiNetworksParamsWithContext(ctx context.Context) *CustomerPrototypeGetWifiNetworksParams {
	return &CustomerPrototypeGetWifiNetworksParams{
		Context: ctx,
	}
}

// NewCustomerPrototypeGetWifiNetworksParamsWithHTTPClient creates a new CustomerPrototypeGetWifiNetworksParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypeGetWifiNetworksParamsWithHTTPClient(client *http.Client) *CustomerPrototypeGetWifiNetworksParams {
	return &CustomerPrototypeGetWifiNetworksParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypeGetWifiNetworksParams contains all the parameters to send to the API endpoint

	for the customer prototype get wifi networks operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypeGetWifiNetworksParams struct {

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

// WithDefaults hydrates default values in the customer prototype get wifi networks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeGetWifiNetworksParams) WithDefaults() *CustomerPrototypeGetWifiNetworksParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype get wifi networks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeGetWifiNetworksParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype get wifi networks params
func (o *CustomerPrototypeGetWifiNetworksParams) WithTimeout(timeout time.Duration) *CustomerPrototypeGetWifiNetworksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype get wifi networks params
func (o *CustomerPrototypeGetWifiNetworksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype get wifi networks params
func (o *CustomerPrototypeGetWifiNetworksParams) WithContext(ctx context.Context) *CustomerPrototypeGetWifiNetworksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype get wifi networks params
func (o *CustomerPrototypeGetWifiNetworksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype get wifi networks params
func (o *CustomerPrototypeGetWifiNetworksParams) WithHTTPClient(client *http.Client) *CustomerPrototypeGetWifiNetworksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype get wifi networks params
func (o *CustomerPrototypeGetWifiNetworksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the customer prototype get wifi networks params
func (o *CustomerPrototypeGetWifiNetworksParams) WithID(id string) *CustomerPrototypeGetWifiNetworksParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype get wifi networks params
func (o *CustomerPrototypeGetWifiNetworksParams) SetID(id string) {
	o.ID = id
}

// WithLocationID adds the locationID to the customer prototype get wifi networks params
func (o *CustomerPrototypeGetWifiNetworksParams) WithLocationID(locationID string) *CustomerPrototypeGetWifiNetworksParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype get wifi networks params
func (o *CustomerPrototypeGetWifiNetworksParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypeGetWifiNetworksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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