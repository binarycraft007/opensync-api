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

// NewCustomerPrototypeBlockDevicesParams creates a new CustomerPrototypeBlockDevicesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypeBlockDevicesParams() *CustomerPrototypeBlockDevicesParams {
	return &CustomerPrototypeBlockDevicesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypeBlockDevicesParamsWithTimeout creates a new CustomerPrototypeBlockDevicesParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypeBlockDevicesParamsWithTimeout(timeout time.Duration) *CustomerPrototypeBlockDevicesParams {
	return &CustomerPrototypeBlockDevicesParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypeBlockDevicesParamsWithContext creates a new CustomerPrototypeBlockDevicesParams object
// with the ability to set a context for a request.
func NewCustomerPrototypeBlockDevicesParamsWithContext(ctx context.Context) *CustomerPrototypeBlockDevicesParams {
	return &CustomerPrototypeBlockDevicesParams{
		Context: ctx,
	}
}

// NewCustomerPrototypeBlockDevicesParamsWithHTTPClient creates a new CustomerPrototypeBlockDevicesParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypeBlockDevicesParamsWithHTTPClient(client *http.Client) *CustomerPrototypeBlockDevicesParams {
	return &CustomerPrototypeBlockDevicesParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypeBlockDevicesParams contains all the parameters to send to the API endpoint

	for the customer prototype block devices operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypeBlockDevicesParams struct {

	/* ID.

	   Customer id

	   Format: JSON
	*/
	ID string

	// LocationID.
	LocationID string

	// Macs.
	Macs []models.XAny

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer prototype block devices params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeBlockDevicesParams) WithDefaults() *CustomerPrototypeBlockDevicesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype block devices params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeBlockDevicesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype block devices params
func (o *CustomerPrototypeBlockDevicesParams) WithTimeout(timeout time.Duration) *CustomerPrototypeBlockDevicesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype block devices params
func (o *CustomerPrototypeBlockDevicesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype block devices params
func (o *CustomerPrototypeBlockDevicesParams) WithContext(ctx context.Context) *CustomerPrototypeBlockDevicesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype block devices params
func (o *CustomerPrototypeBlockDevicesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype block devices params
func (o *CustomerPrototypeBlockDevicesParams) WithHTTPClient(client *http.Client) *CustomerPrototypeBlockDevicesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype block devices params
func (o *CustomerPrototypeBlockDevicesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the customer prototype block devices params
func (o *CustomerPrototypeBlockDevicesParams) WithID(id string) *CustomerPrototypeBlockDevicesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype block devices params
func (o *CustomerPrototypeBlockDevicesParams) SetID(id string) {
	o.ID = id
}

// WithLocationID adds the locationID to the customer prototype block devices params
func (o *CustomerPrototypeBlockDevicesParams) WithLocationID(locationID string) *CustomerPrototypeBlockDevicesParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype block devices params
func (o *CustomerPrototypeBlockDevicesParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WithMacs adds the macs to the customer prototype block devices params
func (o *CustomerPrototypeBlockDevicesParams) WithMacs(macs []models.XAny) *CustomerPrototypeBlockDevicesParams {
	o.SetMacs(macs)
	return o
}

// SetMacs adds the macs to the customer prototype block devices params
func (o *CustomerPrototypeBlockDevicesParams) SetMacs(macs []models.XAny) {
	o.Macs = macs
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypeBlockDevicesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if o.Macs != nil {
		if err := r.SetBodyParam(o.Macs); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}