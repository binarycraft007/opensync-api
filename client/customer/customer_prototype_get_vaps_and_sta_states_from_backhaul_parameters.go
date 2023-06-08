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

// NewCustomerPrototypeGetVapsAndStaStatesFromBackhaulParams creates a new CustomerPrototypeGetVapsAndStaStatesFromBackhaulParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypeGetVapsAndStaStatesFromBackhaulParams() *CustomerPrototypeGetVapsAndStaStatesFromBackhaulParams {
	return &CustomerPrototypeGetVapsAndStaStatesFromBackhaulParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypeGetVapsAndStaStatesFromBackhaulParamsWithTimeout creates a new CustomerPrototypeGetVapsAndStaStatesFromBackhaulParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypeGetVapsAndStaStatesFromBackhaulParamsWithTimeout(timeout time.Duration) *CustomerPrototypeGetVapsAndStaStatesFromBackhaulParams {
	return &CustomerPrototypeGetVapsAndStaStatesFromBackhaulParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypeGetVapsAndStaStatesFromBackhaulParamsWithContext creates a new CustomerPrototypeGetVapsAndStaStatesFromBackhaulParams object
// with the ability to set a context for a request.
func NewCustomerPrototypeGetVapsAndStaStatesFromBackhaulParamsWithContext(ctx context.Context) *CustomerPrototypeGetVapsAndStaStatesFromBackhaulParams {
	return &CustomerPrototypeGetVapsAndStaStatesFromBackhaulParams{
		Context: ctx,
	}
}

// NewCustomerPrototypeGetVapsAndStaStatesFromBackhaulParamsWithHTTPClient creates a new CustomerPrototypeGetVapsAndStaStatesFromBackhaulParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypeGetVapsAndStaStatesFromBackhaulParamsWithHTTPClient(client *http.Client) *CustomerPrototypeGetVapsAndStaStatesFromBackhaulParams {
	return &CustomerPrototypeGetVapsAndStaStatesFromBackhaulParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypeGetVapsAndStaStatesFromBackhaulParams contains all the parameters to send to the API endpoint

	for the customer prototype get vaps and sta states from backhaul operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypeGetVapsAndStaStatesFromBackhaulParams struct {

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

// WithDefaults hydrates default values in the customer prototype get vaps and sta states from backhaul params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeGetVapsAndStaStatesFromBackhaulParams) WithDefaults() *CustomerPrototypeGetVapsAndStaStatesFromBackhaulParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype get vaps and sta states from backhaul params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeGetVapsAndStaStatesFromBackhaulParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype get vaps and sta states from backhaul params
func (o *CustomerPrototypeGetVapsAndStaStatesFromBackhaulParams) WithTimeout(timeout time.Duration) *CustomerPrototypeGetVapsAndStaStatesFromBackhaulParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype get vaps and sta states from backhaul params
func (o *CustomerPrototypeGetVapsAndStaStatesFromBackhaulParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype get vaps and sta states from backhaul params
func (o *CustomerPrototypeGetVapsAndStaStatesFromBackhaulParams) WithContext(ctx context.Context) *CustomerPrototypeGetVapsAndStaStatesFromBackhaulParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype get vaps and sta states from backhaul params
func (o *CustomerPrototypeGetVapsAndStaStatesFromBackhaulParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype get vaps and sta states from backhaul params
func (o *CustomerPrototypeGetVapsAndStaStatesFromBackhaulParams) WithHTTPClient(client *http.Client) *CustomerPrototypeGetVapsAndStaStatesFromBackhaulParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype get vaps and sta states from backhaul params
func (o *CustomerPrototypeGetVapsAndStaStatesFromBackhaulParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the customer prototype get vaps and sta states from backhaul params
func (o *CustomerPrototypeGetVapsAndStaStatesFromBackhaulParams) WithID(id string) *CustomerPrototypeGetVapsAndStaStatesFromBackhaulParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype get vaps and sta states from backhaul params
func (o *CustomerPrototypeGetVapsAndStaStatesFromBackhaulParams) SetID(id string) {
	o.ID = id
}

// WithLocationID adds the locationID to the customer prototype get vaps and sta states from backhaul params
func (o *CustomerPrototypeGetVapsAndStaStatesFromBackhaulParams) WithLocationID(locationID string) *CustomerPrototypeGetVapsAndStaStatesFromBackhaulParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype get vaps and sta states from backhaul params
func (o *CustomerPrototypeGetVapsAndStaStatesFromBackhaulParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypeGetVapsAndStaStatesFromBackhaulParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
