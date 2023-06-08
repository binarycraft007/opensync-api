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

// NewCustomerPrototypeGetCommandsParams creates a new CustomerPrototypeGetCommandsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypeGetCommandsParams() *CustomerPrototypeGetCommandsParams {
	return &CustomerPrototypeGetCommandsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypeGetCommandsParamsWithTimeout creates a new CustomerPrototypeGetCommandsParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypeGetCommandsParamsWithTimeout(timeout time.Duration) *CustomerPrototypeGetCommandsParams {
	return &CustomerPrototypeGetCommandsParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypeGetCommandsParamsWithContext creates a new CustomerPrototypeGetCommandsParams object
// with the ability to set a context for a request.
func NewCustomerPrototypeGetCommandsParamsWithContext(ctx context.Context) *CustomerPrototypeGetCommandsParams {
	return &CustomerPrototypeGetCommandsParams{
		Context: ctx,
	}
}

// NewCustomerPrototypeGetCommandsParamsWithHTTPClient creates a new CustomerPrototypeGetCommandsParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypeGetCommandsParamsWithHTTPClient(client *http.Client) *CustomerPrototypeGetCommandsParams {
	return &CustomerPrototypeGetCommandsParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypeGetCommandsParams contains all the parameters to send to the API endpoint

	for the customer prototype get commands operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypeGetCommandsParams struct {

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

// WithDefaults hydrates default values in the customer prototype get commands params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeGetCommandsParams) WithDefaults() *CustomerPrototypeGetCommandsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype get commands params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeGetCommandsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype get commands params
func (o *CustomerPrototypeGetCommandsParams) WithTimeout(timeout time.Duration) *CustomerPrototypeGetCommandsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype get commands params
func (o *CustomerPrototypeGetCommandsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype get commands params
func (o *CustomerPrototypeGetCommandsParams) WithContext(ctx context.Context) *CustomerPrototypeGetCommandsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype get commands params
func (o *CustomerPrototypeGetCommandsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype get commands params
func (o *CustomerPrototypeGetCommandsParams) WithHTTPClient(client *http.Client) *CustomerPrototypeGetCommandsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype get commands params
func (o *CustomerPrototypeGetCommandsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the customer prototype get commands params
func (o *CustomerPrototypeGetCommandsParams) WithID(id string) *CustomerPrototypeGetCommandsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype get commands params
func (o *CustomerPrototypeGetCommandsParams) SetID(id string) {
	o.ID = id
}

// WithLocationID adds the locationID to the customer prototype get commands params
func (o *CustomerPrototypeGetCommandsParams) WithLocationID(locationID string) *CustomerPrototypeGetCommandsParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype get commands params
func (o *CustomerPrototypeGetCommandsParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypeGetCommandsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
