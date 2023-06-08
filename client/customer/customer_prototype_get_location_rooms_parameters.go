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

// NewCustomerPrototypeGetLocationRoomsParams creates a new CustomerPrototypeGetLocationRoomsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypeGetLocationRoomsParams() *CustomerPrototypeGetLocationRoomsParams {
	return &CustomerPrototypeGetLocationRoomsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypeGetLocationRoomsParamsWithTimeout creates a new CustomerPrototypeGetLocationRoomsParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypeGetLocationRoomsParamsWithTimeout(timeout time.Duration) *CustomerPrototypeGetLocationRoomsParams {
	return &CustomerPrototypeGetLocationRoomsParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypeGetLocationRoomsParamsWithContext creates a new CustomerPrototypeGetLocationRoomsParams object
// with the ability to set a context for a request.
func NewCustomerPrototypeGetLocationRoomsParamsWithContext(ctx context.Context) *CustomerPrototypeGetLocationRoomsParams {
	return &CustomerPrototypeGetLocationRoomsParams{
		Context: ctx,
	}
}

// NewCustomerPrototypeGetLocationRoomsParamsWithHTTPClient creates a new CustomerPrototypeGetLocationRoomsParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypeGetLocationRoomsParamsWithHTTPClient(client *http.Client) *CustomerPrototypeGetLocationRoomsParams {
	return &CustomerPrototypeGetLocationRoomsParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypeGetLocationRoomsParams contains all the parameters to send to the API endpoint

	for the customer prototype get location rooms operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypeGetLocationRoomsParams struct {

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

// WithDefaults hydrates default values in the customer prototype get location rooms params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeGetLocationRoomsParams) WithDefaults() *CustomerPrototypeGetLocationRoomsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype get location rooms params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeGetLocationRoomsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype get location rooms params
func (o *CustomerPrototypeGetLocationRoomsParams) WithTimeout(timeout time.Duration) *CustomerPrototypeGetLocationRoomsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype get location rooms params
func (o *CustomerPrototypeGetLocationRoomsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype get location rooms params
func (o *CustomerPrototypeGetLocationRoomsParams) WithContext(ctx context.Context) *CustomerPrototypeGetLocationRoomsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype get location rooms params
func (o *CustomerPrototypeGetLocationRoomsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype get location rooms params
func (o *CustomerPrototypeGetLocationRoomsParams) WithHTTPClient(client *http.Client) *CustomerPrototypeGetLocationRoomsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype get location rooms params
func (o *CustomerPrototypeGetLocationRoomsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the customer prototype get location rooms params
func (o *CustomerPrototypeGetLocationRoomsParams) WithID(id string) *CustomerPrototypeGetLocationRoomsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype get location rooms params
func (o *CustomerPrototypeGetLocationRoomsParams) SetID(id string) {
	o.ID = id
}

// WithLocationID adds the locationID to the customer prototype get location rooms params
func (o *CustomerPrototypeGetLocationRoomsParams) WithLocationID(locationID string) *CustomerPrototypeGetLocationRoomsParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype get location rooms params
func (o *CustomerPrototypeGetLocationRoomsParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypeGetLocationRoomsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
