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

// NewCustomerPrototypeGetDashboardParams creates a new CustomerPrototypeGetDashboardParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypeGetDashboardParams() *CustomerPrototypeGetDashboardParams {
	return &CustomerPrototypeGetDashboardParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypeGetDashboardParamsWithTimeout creates a new CustomerPrototypeGetDashboardParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypeGetDashboardParamsWithTimeout(timeout time.Duration) *CustomerPrototypeGetDashboardParams {
	return &CustomerPrototypeGetDashboardParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypeGetDashboardParamsWithContext creates a new CustomerPrototypeGetDashboardParams object
// with the ability to set a context for a request.
func NewCustomerPrototypeGetDashboardParamsWithContext(ctx context.Context) *CustomerPrototypeGetDashboardParams {
	return &CustomerPrototypeGetDashboardParams{
		Context: ctx,
	}
}

// NewCustomerPrototypeGetDashboardParamsWithHTTPClient creates a new CustomerPrototypeGetDashboardParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypeGetDashboardParamsWithHTTPClient(client *http.Client) *CustomerPrototypeGetDashboardParams {
	return &CustomerPrototypeGetDashboardParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypeGetDashboardParams contains all the parameters to send to the API endpoint

	for the customer prototype get dashboard operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypeGetDashboardParams struct {

	/* ID.

	   Customer id

	   Format: JSON
	*/
	ID string

	// LocationID.
	LocationID string

	/* Macs.

	   mac list of all devices in the location

	   Format: JSON
	*/
	Macs *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer prototype get dashboard params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeGetDashboardParams) WithDefaults() *CustomerPrototypeGetDashboardParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype get dashboard params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeGetDashboardParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype get dashboard params
func (o *CustomerPrototypeGetDashboardParams) WithTimeout(timeout time.Duration) *CustomerPrototypeGetDashboardParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype get dashboard params
func (o *CustomerPrototypeGetDashboardParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype get dashboard params
func (o *CustomerPrototypeGetDashboardParams) WithContext(ctx context.Context) *CustomerPrototypeGetDashboardParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype get dashboard params
func (o *CustomerPrototypeGetDashboardParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype get dashboard params
func (o *CustomerPrototypeGetDashboardParams) WithHTTPClient(client *http.Client) *CustomerPrototypeGetDashboardParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype get dashboard params
func (o *CustomerPrototypeGetDashboardParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the customer prototype get dashboard params
func (o *CustomerPrototypeGetDashboardParams) WithID(id string) *CustomerPrototypeGetDashboardParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype get dashboard params
func (o *CustomerPrototypeGetDashboardParams) SetID(id string) {
	o.ID = id
}

// WithLocationID adds the locationID to the customer prototype get dashboard params
func (o *CustomerPrototypeGetDashboardParams) WithLocationID(locationID string) *CustomerPrototypeGetDashboardParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype get dashboard params
func (o *CustomerPrototypeGetDashboardParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WithMacs adds the macs to the customer prototype get dashboard params
func (o *CustomerPrototypeGetDashboardParams) WithMacs(macs *string) *CustomerPrototypeGetDashboardParams {
	o.SetMacs(macs)
	return o
}

// SetMacs adds the macs to the customer prototype get dashboard params
func (o *CustomerPrototypeGetDashboardParams) SetMacs(macs *string) {
	o.Macs = macs
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypeGetDashboardParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Macs != nil {

		// query param macs
		var qrMacs string

		if o.Macs != nil {
			qrMacs = *o.Macs
		}
		qMacs := qrMacs
		if qMacs != "" {

			if err := r.SetQueryParam("macs", qMacs); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
