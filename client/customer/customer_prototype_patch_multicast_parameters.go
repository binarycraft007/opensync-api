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

	"github.com/binarycraft007/opensync-api/models"
)

// NewCustomerPrototypePatchMulticastParams creates a new CustomerPrototypePatchMulticastParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypePatchMulticastParams() *CustomerPrototypePatchMulticastParams {
	return &CustomerPrototypePatchMulticastParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypePatchMulticastParamsWithTimeout creates a new CustomerPrototypePatchMulticastParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypePatchMulticastParamsWithTimeout(timeout time.Duration) *CustomerPrototypePatchMulticastParams {
	return &CustomerPrototypePatchMulticastParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypePatchMulticastParamsWithContext creates a new CustomerPrototypePatchMulticastParams object
// with the ability to set a context for a request.
func NewCustomerPrototypePatchMulticastParamsWithContext(ctx context.Context) *CustomerPrototypePatchMulticastParams {
	return &CustomerPrototypePatchMulticastParams{
		Context: ctx,
	}
}

// NewCustomerPrototypePatchMulticastParamsWithHTTPClient creates a new CustomerPrototypePatchMulticastParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypePatchMulticastParamsWithHTTPClient(client *http.Client) *CustomerPrototypePatchMulticastParams {
	return &CustomerPrototypePatchMulticastParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypePatchMulticastParams contains all the parameters to send to the API endpoint

	for the customer prototype patch multicast operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypePatchMulticastParams struct {

	/* ID.

	   Customer id

	   Format: JSON
	*/
	ID string

	// LocationID.
	//
	// Format: JSON
	LocationID string

	/* Multicast.

	   multicast object
	*/
	Multicast *models.Multicast

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer prototype patch multicast params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePatchMulticastParams) WithDefaults() *CustomerPrototypePatchMulticastParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype patch multicast params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePatchMulticastParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype patch multicast params
func (o *CustomerPrototypePatchMulticastParams) WithTimeout(timeout time.Duration) *CustomerPrototypePatchMulticastParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype patch multicast params
func (o *CustomerPrototypePatchMulticastParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype patch multicast params
func (o *CustomerPrototypePatchMulticastParams) WithContext(ctx context.Context) *CustomerPrototypePatchMulticastParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype patch multicast params
func (o *CustomerPrototypePatchMulticastParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype patch multicast params
func (o *CustomerPrototypePatchMulticastParams) WithHTTPClient(client *http.Client) *CustomerPrototypePatchMulticastParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype patch multicast params
func (o *CustomerPrototypePatchMulticastParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the customer prototype patch multicast params
func (o *CustomerPrototypePatchMulticastParams) WithID(id string) *CustomerPrototypePatchMulticastParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype patch multicast params
func (o *CustomerPrototypePatchMulticastParams) SetID(id string) {
	o.ID = id
}

// WithLocationID adds the locationID to the customer prototype patch multicast params
func (o *CustomerPrototypePatchMulticastParams) WithLocationID(locationID string) *CustomerPrototypePatchMulticastParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype patch multicast params
func (o *CustomerPrototypePatchMulticastParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WithMulticast adds the multicast to the customer prototype patch multicast params
func (o *CustomerPrototypePatchMulticastParams) WithMulticast(multicast *models.Multicast) *CustomerPrototypePatchMulticastParams {
	o.SetMulticast(multicast)
	return o
}

// SetMulticast adds the multicast to the customer prototype patch multicast params
func (o *CustomerPrototypePatchMulticastParams) SetMulticast(multicast *models.Multicast) {
	o.Multicast = multicast
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypePatchMulticastParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if o.Multicast != nil {
		if err := r.SetBodyParam(o.Multicast); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
