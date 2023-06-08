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

// NewCustomerPrototypeDeleteDeviceFromPersonParams creates a new CustomerPrototypeDeleteDeviceFromPersonParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypeDeleteDeviceFromPersonParams() *CustomerPrototypeDeleteDeviceFromPersonParams {
	return &CustomerPrototypeDeleteDeviceFromPersonParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypeDeleteDeviceFromPersonParamsWithTimeout creates a new CustomerPrototypeDeleteDeviceFromPersonParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypeDeleteDeviceFromPersonParamsWithTimeout(timeout time.Duration) *CustomerPrototypeDeleteDeviceFromPersonParams {
	return &CustomerPrototypeDeleteDeviceFromPersonParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypeDeleteDeviceFromPersonParamsWithContext creates a new CustomerPrototypeDeleteDeviceFromPersonParams object
// with the ability to set a context for a request.
func NewCustomerPrototypeDeleteDeviceFromPersonParamsWithContext(ctx context.Context) *CustomerPrototypeDeleteDeviceFromPersonParams {
	return &CustomerPrototypeDeleteDeviceFromPersonParams{
		Context: ctx,
	}
}

// NewCustomerPrototypeDeleteDeviceFromPersonParamsWithHTTPClient creates a new CustomerPrototypeDeleteDeviceFromPersonParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypeDeleteDeviceFromPersonParamsWithHTTPClient(client *http.Client) *CustomerPrototypeDeleteDeviceFromPersonParams {
	return &CustomerPrototypeDeleteDeviceFromPersonParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypeDeleteDeviceFromPersonParams contains all the parameters to send to the API endpoint

	for the customer prototype delete device from person operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypeDeleteDeviceFromPersonParams struct {

	/* ID.

	   Customer id

	   Format: JSON
	*/
	ID string

	// LocationID.
	LocationID string

	// Mac.
	Mac string

	// PersonID.
	PersonID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer prototype delete device from person params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeDeleteDeviceFromPersonParams) WithDefaults() *CustomerPrototypeDeleteDeviceFromPersonParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype delete device from person params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeDeleteDeviceFromPersonParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype delete device from person params
func (o *CustomerPrototypeDeleteDeviceFromPersonParams) WithTimeout(timeout time.Duration) *CustomerPrototypeDeleteDeviceFromPersonParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype delete device from person params
func (o *CustomerPrototypeDeleteDeviceFromPersonParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype delete device from person params
func (o *CustomerPrototypeDeleteDeviceFromPersonParams) WithContext(ctx context.Context) *CustomerPrototypeDeleteDeviceFromPersonParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype delete device from person params
func (o *CustomerPrototypeDeleteDeviceFromPersonParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype delete device from person params
func (o *CustomerPrototypeDeleteDeviceFromPersonParams) WithHTTPClient(client *http.Client) *CustomerPrototypeDeleteDeviceFromPersonParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype delete device from person params
func (o *CustomerPrototypeDeleteDeviceFromPersonParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the customer prototype delete device from person params
func (o *CustomerPrototypeDeleteDeviceFromPersonParams) WithID(id string) *CustomerPrototypeDeleteDeviceFromPersonParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype delete device from person params
func (o *CustomerPrototypeDeleteDeviceFromPersonParams) SetID(id string) {
	o.ID = id
}

// WithLocationID adds the locationID to the customer prototype delete device from person params
func (o *CustomerPrototypeDeleteDeviceFromPersonParams) WithLocationID(locationID string) *CustomerPrototypeDeleteDeviceFromPersonParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype delete device from person params
func (o *CustomerPrototypeDeleteDeviceFromPersonParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WithMac adds the mac to the customer prototype delete device from person params
func (o *CustomerPrototypeDeleteDeviceFromPersonParams) WithMac(mac string) *CustomerPrototypeDeleteDeviceFromPersonParams {
	o.SetMac(mac)
	return o
}

// SetMac adds the mac to the customer prototype delete device from person params
func (o *CustomerPrototypeDeleteDeviceFromPersonParams) SetMac(mac string) {
	o.Mac = mac
}

// WithPersonID adds the personID to the customer prototype delete device from person params
func (o *CustomerPrototypeDeleteDeviceFromPersonParams) WithPersonID(personID string) *CustomerPrototypeDeleteDeviceFromPersonParams {
	o.SetPersonID(personID)
	return o
}

// SetPersonID adds the personId to the customer prototype delete device from person params
func (o *CustomerPrototypeDeleteDeviceFromPersonParams) SetPersonID(personID string) {
	o.PersonID = personID
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypeDeleteDeviceFromPersonParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param mac
	if err := r.SetPathParam("mac", o.Mac); err != nil {
		return err
	}

	// path param personId
	if err := r.SetPathParam("personId", o.PersonID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}