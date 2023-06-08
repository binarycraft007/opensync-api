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

// NewCustomerPrototypePatchDppParams creates a new CustomerPrototypePatchDppParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypePatchDppParams() *CustomerPrototypePatchDppParams {
	return &CustomerPrototypePatchDppParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypePatchDppParamsWithTimeout creates a new CustomerPrototypePatchDppParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypePatchDppParamsWithTimeout(timeout time.Duration) *CustomerPrototypePatchDppParams {
	return &CustomerPrototypePatchDppParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypePatchDppParamsWithContext creates a new CustomerPrototypePatchDppParams object
// with the ability to set a context for a request.
func NewCustomerPrototypePatchDppParamsWithContext(ctx context.Context) *CustomerPrototypePatchDppParams {
	return &CustomerPrototypePatchDppParams{
		Context: ctx,
	}
}

// NewCustomerPrototypePatchDppParamsWithHTTPClient creates a new CustomerPrototypePatchDppParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypePatchDppParamsWithHTTPClient(client *http.Client) *CustomerPrototypePatchDppParams {
	return &CustomerPrototypePatchDppParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypePatchDppParams contains all the parameters to send to the API endpoint

	for the customer prototype patch dpp operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypePatchDppParams struct {

	/* ID.

	   Customer id

	   Format: JSON
	*/
	ID string

	// LocationID.
	//
	// Format: JSON
	LocationID string

	/* Mode.

	   auto || enable || disable
	*/
	Mode string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer prototype patch dpp params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePatchDppParams) WithDefaults() *CustomerPrototypePatchDppParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype patch dpp params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePatchDppParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype patch dpp params
func (o *CustomerPrototypePatchDppParams) WithTimeout(timeout time.Duration) *CustomerPrototypePatchDppParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype patch dpp params
func (o *CustomerPrototypePatchDppParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype patch dpp params
func (o *CustomerPrototypePatchDppParams) WithContext(ctx context.Context) *CustomerPrototypePatchDppParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype patch dpp params
func (o *CustomerPrototypePatchDppParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype patch dpp params
func (o *CustomerPrototypePatchDppParams) WithHTTPClient(client *http.Client) *CustomerPrototypePatchDppParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype patch dpp params
func (o *CustomerPrototypePatchDppParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the customer prototype patch dpp params
func (o *CustomerPrototypePatchDppParams) WithID(id string) *CustomerPrototypePatchDppParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype patch dpp params
func (o *CustomerPrototypePatchDppParams) SetID(id string) {
	o.ID = id
}

// WithLocationID adds the locationID to the customer prototype patch dpp params
func (o *CustomerPrototypePatchDppParams) WithLocationID(locationID string) *CustomerPrototypePatchDppParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype patch dpp params
func (o *CustomerPrototypePatchDppParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WithMode adds the mode to the customer prototype patch dpp params
func (o *CustomerPrototypePatchDppParams) WithMode(mode string) *CustomerPrototypePatchDppParams {
	o.SetMode(mode)
	return o
}

// SetMode adds the mode to the customer prototype patch dpp params
func (o *CustomerPrototypePatchDppParams) SetMode(mode string) {
	o.Mode = mode
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypePatchDppParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// form param mode
	frMode := o.Mode
	fMode := frMode
	if fMode != "" {
		if err := r.SetFormParam("mode", fMode); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
