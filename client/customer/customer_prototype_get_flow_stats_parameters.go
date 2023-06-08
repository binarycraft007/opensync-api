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

// NewCustomerPrototypeGetFlowStatsParams creates a new CustomerPrototypeGetFlowStatsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypeGetFlowStatsParams() *CustomerPrototypeGetFlowStatsParams {
	return &CustomerPrototypeGetFlowStatsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypeGetFlowStatsParamsWithTimeout creates a new CustomerPrototypeGetFlowStatsParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypeGetFlowStatsParamsWithTimeout(timeout time.Duration) *CustomerPrototypeGetFlowStatsParams {
	return &CustomerPrototypeGetFlowStatsParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypeGetFlowStatsParamsWithContext creates a new CustomerPrototypeGetFlowStatsParams object
// with the ability to set a context for a request.
func NewCustomerPrototypeGetFlowStatsParamsWithContext(ctx context.Context) *CustomerPrototypeGetFlowStatsParams {
	return &CustomerPrototypeGetFlowStatsParams{
		Context: ctx,
	}
}

// NewCustomerPrototypeGetFlowStatsParamsWithHTTPClient creates a new CustomerPrototypeGetFlowStatsParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypeGetFlowStatsParamsWithHTTPClient(client *http.Client) *CustomerPrototypeGetFlowStatsParams {
	return &CustomerPrototypeGetFlowStatsParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypeGetFlowStatsParams contains all the parameters to send to the API endpoint

	for the customer prototype get flow stats operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypeGetFlowStatsParams struct {

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

// WithDefaults hydrates default values in the customer prototype get flow stats params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeGetFlowStatsParams) WithDefaults() *CustomerPrototypeGetFlowStatsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype get flow stats params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeGetFlowStatsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype get flow stats params
func (o *CustomerPrototypeGetFlowStatsParams) WithTimeout(timeout time.Duration) *CustomerPrototypeGetFlowStatsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype get flow stats params
func (o *CustomerPrototypeGetFlowStatsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype get flow stats params
func (o *CustomerPrototypeGetFlowStatsParams) WithContext(ctx context.Context) *CustomerPrototypeGetFlowStatsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype get flow stats params
func (o *CustomerPrototypeGetFlowStatsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype get flow stats params
func (o *CustomerPrototypeGetFlowStatsParams) WithHTTPClient(client *http.Client) *CustomerPrototypeGetFlowStatsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype get flow stats params
func (o *CustomerPrototypeGetFlowStatsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the customer prototype get flow stats params
func (o *CustomerPrototypeGetFlowStatsParams) WithID(id string) *CustomerPrototypeGetFlowStatsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype get flow stats params
func (o *CustomerPrototypeGetFlowStatsParams) SetID(id string) {
	o.ID = id
}

// WithLocationID adds the locationID to the customer prototype get flow stats params
func (o *CustomerPrototypeGetFlowStatsParams) WithLocationID(locationID string) *CustomerPrototypeGetFlowStatsParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype get flow stats params
func (o *CustomerPrototypeGetFlowStatsParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypeGetFlowStatsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
