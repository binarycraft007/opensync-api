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

// NewCustomerPrototypeGetFirmwareUpgradeStatusParams creates a new CustomerPrototypeGetFirmwareUpgradeStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypeGetFirmwareUpgradeStatusParams() *CustomerPrototypeGetFirmwareUpgradeStatusParams {
	return &CustomerPrototypeGetFirmwareUpgradeStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypeGetFirmwareUpgradeStatusParamsWithTimeout creates a new CustomerPrototypeGetFirmwareUpgradeStatusParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypeGetFirmwareUpgradeStatusParamsWithTimeout(timeout time.Duration) *CustomerPrototypeGetFirmwareUpgradeStatusParams {
	return &CustomerPrototypeGetFirmwareUpgradeStatusParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypeGetFirmwareUpgradeStatusParamsWithContext creates a new CustomerPrototypeGetFirmwareUpgradeStatusParams object
// with the ability to set a context for a request.
func NewCustomerPrototypeGetFirmwareUpgradeStatusParamsWithContext(ctx context.Context) *CustomerPrototypeGetFirmwareUpgradeStatusParams {
	return &CustomerPrototypeGetFirmwareUpgradeStatusParams{
		Context: ctx,
	}
}

// NewCustomerPrototypeGetFirmwareUpgradeStatusParamsWithHTTPClient creates a new CustomerPrototypeGetFirmwareUpgradeStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypeGetFirmwareUpgradeStatusParamsWithHTTPClient(client *http.Client) *CustomerPrototypeGetFirmwareUpgradeStatusParams {
	return &CustomerPrototypeGetFirmwareUpgradeStatusParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypeGetFirmwareUpgradeStatusParams contains all the parameters to send to the API endpoint

	for the customer prototype get firmware upgrade status operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypeGetFirmwareUpgradeStatusParams struct {

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

// WithDefaults hydrates default values in the customer prototype get firmware upgrade status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeGetFirmwareUpgradeStatusParams) WithDefaults() *CustomerPrototypeGetFirmwareUpgradeStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype get firmware upgrade status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeGetFirmwareUpgradeStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype get firmware upgrade status params
func (o *CustomerPrototypeGetFirmwareUpgradeStatusParams) WithTimeout(timeout time.Duration) *CustomerPrototypeGetFirmwareUpgradeStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype get firmware upgrade status params
func (o *CustomerPrototypeGetFirmwareUpgradeStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype get firmware upgrade status params
func (o *CustomerPrototypeGetFirmwareUpgradeStatusParams) WithContext(ctx context.Context) *CustomerPrototypeGetFirmwareUpgradeStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype get firmware upgrade status params
func (o *CustomerPrototypeGetFirmwareUpgradeStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype get firmware upgrade status params
func (o *CustomerPrototypeGetFirmwareUpgradeStatusParams) WithHTTPClient(client *http.Client) *CustomerPrototypeGetFirmwareUpgradeStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype get firmware upgrade status params
func (o *CustomerPrototypeGetFirmwareUpgradeStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the customer prototype get firmware upgrade status params
func (o *CustomerPrototypeGetFirmwareUpgradeStatusParams) WithID(id string) *CustomerPrototypeGetFirmwareUpgradeStatusParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype get firmware upgrade status params
func (o *CustomerPrototypeGetFirmwareUpgradeStatusParams) SetID(id string) {
	o.ID = id
}

// WithLocationID adds the locationID to the customer prototype get firmware upgrade status params
func (o *CustomerPrototypeGetFirmwareUpgradeStatusParams) WithLocationID(locationID string) *CustomerPrototypeGetFirmwareUpgradeStatusParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype get firmware upgrade status params
func (o *CustomerPrototypeGetFirmwareUpgradeStatusParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypeGetFirmwareUpgradeStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
