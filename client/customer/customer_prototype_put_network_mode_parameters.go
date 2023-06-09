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

// NewCustomerPrototypePutNetworkModeParams creates a new CustomerPrototypePutNetworkModeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypePutNetworkModeParams() *CustomerPrototypePutNetworkModeParams {
	return &CustomerPrototypePutNetworkModeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypePutNetworkModeParamsWithTimeout creates a new CustomerPrototypePutNetworkModeParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypePutNetworkModeParamsWithTimeout(timeout time.Duration) *CustomerPrototypePutNetworkModeParams {
	return &CustomerPrototypePutNetworkModeParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypePutNetworkModeParamsWithContext creates a new CustomerPrototypePutNetworkModeParams object
// with the ability to set a context for a request.
func NewCustomerPrototypePutNetworkModeParamsWithContext(ctx context.Context) *CustomerPrototypePutNetworkModeParams {
	return &CustomerPrototypePutNetworkModeParams{
		Context: ctx,
	}
}

// NewCustomerPrototypePutNetworkModeParamsWithHTTPClient creates a new CustomerPrototypePutNetworkModeParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypePutNetworkModeParamsWithHTTPClient(client *http.Client) *CustomerPrototypePutNetworkModeParams {
	return &CustomerPrototypePutNetworkModeParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypePutNetworkModeParams contains all the parameters to send to the API endpoint

	for the customer prototype put network mode operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypePutNetworkModeParams struct {

	/* ID.

	   Customer id

	   Format: JSON
	*/
	ID string

	// LocationID.
	//
	// Format: JSON
	LocationID string

	// NetworkMode.
	NetworkMode *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer prototype put network mode params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePutNetworkModeParams) WithDefaults() *CustomerPrototypePutNetworkModeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype put network mode params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePutNetworkModeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype put network mode params
func (o *CustomerPrototypePutNetworkModeParams) WithTimeout(timeout time.Duration) *CustomerPrototypePutNetworkModeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype put network mode params
func (o *CustomerPrototypePutNetworkModeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype put network mode params
func (o *CustomerPrototypePutNetworkModeParams) WithContext(ctx context.Context) *CustomerPrototypePutNetworkModeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype put network mode params
func (o *CustomerPrototypePutNetworkModeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype put network mode params
func (o *CustomerPrototypePutNetworkModeParams) WithHTTPClient(client *http.Client) *CustomerPrototypePutNetworkModeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype put network mode params
func (o *CustomerPrototypePutNetworkModeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the customer prototype put network mode params
func (o *CustomerPrototypePutNetworkModeParams) WithID(id string) *CustomerPrototypePutNetworkModeParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype put network mode params
func (o *CustomerPrototypePutNetworkModeParams) SetID(id string) {
	o.ID = id
}

// WithLocationID adds the locationID to the customer prototype put network mode params
func (o *CustomerPrototypePutNetworkModeParams) WithLocationID(locationID string) *CustomerPrototypePutNetworkModeParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype put network mode params
func (o *CustomerPrototypePutNetworkModeParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WithNetworkMode adds the networkMode to the customer prototype put network mode params
func (o *CustomerPrototypePutNetworkModeParams) WithNetworkMode(networkMode *string) *CustomerPrototypePutNetworkModeParams {
	o.SetNetworkMode(networkMode)
	return o
}

// SetNetworkMode adds the networkMode to the customer prototype put network mode params
func (o *CustomerPrototypePutNetworkModeParams) SetNetworkMode(networkMode *string) {
	o.NetworkMode = networkMode
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypePutNetworkModeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.NetworkMode != nil {

		// form param networkMode
		var frNetworkMode string
		if o.NetworkMode != nil {
			frNetworkMode = *o.NetworkMode
		}
		fNetworkMode := frNetworkMode
		if fNetworkMode != "" {
			if err := r.SetFormParam("networkMode", fNetworkMode); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
