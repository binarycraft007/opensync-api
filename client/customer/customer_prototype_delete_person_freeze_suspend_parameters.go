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

// NewCustomerPrototypeDeletePersonFreezeSuspendParams creates a new CustomerPrototypeDeletePersonFreezeSuspendParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypeDeletePersonFreezeSuspendParams() *CustomerPrototypeDeletePersonFreezeSuspendParams {
	return &CustomerPrototypeDeletePersonFreezeSuspendParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypeDeletePersonFreezeSuspendParamsWithTimeout creates a new CustomerPrototypeDeletePersonFreezeSuspendParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypeDeletePersonFreezeSuspendParamsWithTimeout(timeout time.Duration) *CustomerPrototypeDeletePersonFreezeSuspendParams {
	return &CustomerPrototypeDeletePersonFreezeSuspendParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypeDeletePersonFreezeSuspendParamsWithContext creates a new CustomerPrototypeDeletePersonFreezeSuspendParams object
// with the ability to set a context for a request.
func NewCustomerPrototypeDeletePersonFreezeSuspendParamsWithContext(ctx context.Context) *CustomerPrototypeDeletePersonFreezeSuspendParams {
	return &CustomerPrototypeDeletePersonFreezeSuspendParams{
		Context: ctx,
	}
}

// NewCustomerPrototypeDeletePersonFreezeSuspendParamsWithHTTPClient creates a new CustomerPrototypeDeletePersonFreezeSuspendParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypeDeletePersonFreezeSuspendParamsWithHTTPClient(client *http.Client) *CustomerPrototypeDeletePersonFreezeSuspendParams {
	return &CustomerPrototypeDeletePersonFreezeSuspendParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypeDeletePersonFreezeSuspendParams contains all the parameters to send to the API endpoint

	for the customer prototype delete person freeze suspend operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypeDeletePersonFreezeSuspendParams struct {

	/* ID.

	   Customer id

	   Format: JSON
	*/
	ID string

	// LocationID.
	LocationID string

	// PersonID.
	PersonID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer prototype delete person freeze suspend params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeDeletePersonFreezeSuspendParams) WithDefaults() *CustomerPrototypeDeletePersonFreezeSuspendParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype delete person freeze suspend params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeDeletePersonFreezeSuspendParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype delete person freeze suspend params
func (o *CustomerPrototypeDeletePersonFreezeSuspendParams) WithTimeout(timeout time.Duration) *CustomerPrototypeDeletePersonFreezeSuspendParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype delete person freeze suspend params
func (o *CustomerPrototypeDeletePersonFreezeSuspendParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype delete person freeze suspend params
func (o *CustomerPrototypeDeletePersonFreezeSuspendParams) WithContext(ctx context.Context) *CustomerPrototypeDeletePersonFreezeSuspendParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype delete person freeze suspend params
func (o *CustomerPrototypeDeletePersonFreezeSuspendParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype delete person freeze suspend params
func (o *CustomerPrototypeDeletePersonFreezeSuspendParams) WithHTTPClient(client *http.Client) *CustomerPrototypeDeletePersonFreezeSuspendParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype delete person freeze suspend params
func (o *CustomerPrototypeDeletePersonFreezeSuspendParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the customer prototype delete person freeze suspend params
func (o *CustomerPrototypeDeletePersonFreezeSuspendParams) WithID(id string) *CustomerPrototypeDeletePersonFreezeSuspendParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype delete person freeze suspend params
func (o *CustomerPrototypeDeletePersonFreezeSuspendParams) SetID(id string) {
	o.ID = id
}

// WithLocationID adds the locationID to the customer prototype delete person freeze suspend params
func (o *CustomerPrototypeDeletePersonFreezeSuspendParams) WithLocationID(locationID string) *CustomerPrototypeDeletePersonFreezeSuspendParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype delete person freeze suspend params
func (o *CustomerPrototypeDeletePersonFreezeSuspendParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WithPersonID adds the personID to the customer prototype delete person freeze suspend params
func (o *CustomerPrototypeDeletePersonFreezeSuspendParams) WithPersonID(personID string) *CustomerPrototypeDeletePersonFreezeSuspendParams {
	o.SetPersonID(personID)
	return o
}

// SetPersonID adds the personId to the customer prototype delete person freeze suspend params
func (o *CustomerPrototypeDeletePersonFreezeSuspendParams) SetPersonID(personID string) {
	o.PersonID = personID
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypeDeletePersonFreezeSuspendParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param personId
	if err := r.SetPathParam("personId", o.PersonID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
