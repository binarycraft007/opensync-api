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

// NewCustomerPrototypeVlanServicesParams creates a new CustomerPrototypeVlanServicesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypeVlanServicesParams() *CustomerPrototypeVlanServicesParams {
	return &CustomerPrototypeVlanServicesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypeVlanServicesParamsWithTimeout creates a new CustomerPrototypeVlanServicesParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypeVlanServicesParamsWithTimeout(timeout time.Duration) *CustomerPrototypeVlanServicesParams {
	return &CustomerPrototypeVlanServicesParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypeVlanServicesParamsWithContext creates a new CustomerPrototypeVlanServicesParams object
// with the ability to set a context for a request.
func NewCustomerPrototypeVlanServicesParamsWithContext(ctx context.Context) *CustomerPrototypeVlanServicesParams {
	return &CustomerPrototypeVlanServicesParams{
		Context: ctx,
	}
}

// NewCustomerPrototypeVlanServicesParamsWithHTTPClient creates a new CustomerPrototypeVlanServicesParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypeVlanServicesParamsWithHTTPClient(client *http.Client) *CustomerPrototypeVlanServicesParams {
	return &CustomerPrototypeVlanServicesParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypeVlanServicesParams contains all the parameters to send to the API endpoint

	for the customer prototype vlan services operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypeVlanServicesParams struct {

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

// WithDefaults hydrates default values in the customer prototype vlan services params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeVlanServicesParams) WithDefaults() *CustomerPrototypeVlanServicesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype vlan services params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeVlanServicesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype vlan services params
func (o *CustomerPrototypeVlanServicesParams) WithTimeout(timeout time.Duration) *CustomerPrototypeVlanServicesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype vlan services params
func (o *CustomerPrototypeVlanServicesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype vlan services params
func (o *CustomerPrototypeVlanServicesParams) WithContext(ctx context.Context) *CustomerPrototypeVlanServicesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype vlan services params
func (o *CustomerPrototypeVlanServicesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype vlan services params
func (o *CustomerPrototypeVlanServicesParams) WithHTTPClient(client *http.Client) *CustomerPrototypeVlanServicesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype vlan services params
func (o *CustomerPrototypeVlanServicesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the customer prototype vlan services params
func (o *CustomerPrototypeVlanServicesParams) WithID(id string) *CustomerPrototypeVlanServicesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype vlan services params
func (o *CustomerPrototypeVlanServicesParams) SetID(id string) {
	o.ID = id
}

// WithLocationID adds the locationID to the customer prototype vlan services params
func (o *CustomerPrototypeVlanServicesParams) WithLocationID(locationID string) *CustomerPrototypeVlanServicesParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype vlan services params
func (o *CustomerPrototypeVlanServicesParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypeVlanServicesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
