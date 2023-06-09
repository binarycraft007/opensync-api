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

// NewCustomerPrototypeGetDhcpReservationsParams creates a new CustomerPrototypeGetDhcpReservationsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypeGetDhcpReservationsParams() *CustomerPrototypeGetDhcpReservationsParams {
	return &CustomerPrototypeGetDhcpReservationsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypeGetDhcpReservationsParamsWithTimeout creates a new CustomerPrototypeGetDhcpReservationsParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypeGetDhcpReservationsParamsWithTimeout(timeout time.Duration) *CustomerPrototypeGetDhcpReservationsParams {
	return &CustomerPrototypeGetDhcpReservationsParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypeGetDhcpReservationsParamsWithContext creates a new CustomerPrototypeGetDhcpReservationsParams object
// with the ability to set a context for a request.
func NewCustomerPrototypeGetDhcpReservationsParamsWithContext(ctx context.Context) *CustomerPrototypeGetDhcpReservationsParams {
	return &CustomerPrototypeGetDhcpReservationsParams{
		Context: ctx,
	}
}

// NewCustomerPrototypeGetDhcpReservationsParamsWithHTTPClient creates a new CustomerPrototypeGetDhcpReservationsParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypeGetDhcpReservationsParamsWithHTTPClient(client *http.Client) *CustomerPrototypeGetDhcpReservationsParams {
	return &CustomerPrototypeGetDhcpReservationsParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypeGetDhcpReservationsParams contains all the parameters to send to the API endpoint

	for the customer prototype get dhcp reservations operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypeGetDhcpReservationsParams struct {

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

// WithDefaults hydrates default values in the customer prototype get dhcp reservations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeGetDhcpReservationsParams) WithDefaults() *CustomerPrototypeGetDhcpReservationsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype get dhcp reservations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeGetDhcpReservationsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype get dhcp reservations params
func (o *CustomerPrototypeGetDhcpReservationsParams) WithTimeout(timeout time.Duration) *CustomerPrototypeGetDhcpReservationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype get dhcp reservations params
func (o *CustomerPrototypeGetDhcpReservationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype get dhcp reservations params
func (o *CustomerPrototypeGetDhcpReservationsParams) WithContext(ctx context.Context) *CustomerPrototypeGetDhcpReservationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype get dhcp reservations params
func (o *CustomerPrototypeGetDhcpReservationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype get dhcp reservations params
func (o *CustomerPrototypeGetDhcpReservationsParams) WithHTTPClient(client *http.Client) *CustomerPrototypeGetDhcpReservationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype get dhcp reservations params
func (o *CustomerPrototypeGetDhcpReservationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the customer prototype get dhcp reservations params
func (o *CustomerPrototypeGetDhcpReservationsParams) WithID(id string) *CustomerPrototypeGetDhcpReservationsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype get dhcp reservations params
func (o *CustomerPrototypeGetDhcpReservationsParams) SetID(id string) {
	o.ID = id
}

// WithLocationID adds the locationID to the customer prototype get dhcp reservations params
func (o *CustomerPrototypeGetDhcpReservationsParams) WithLocationID(locationID string) *CustomerPrototypeGetDhcpReservationsParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype get dhcp reservations params
func (o *CustomerPrototypeGetDhcpReservationsParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypeGetDhcpReservationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
