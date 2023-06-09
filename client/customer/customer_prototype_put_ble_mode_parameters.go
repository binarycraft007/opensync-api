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

// NewCustomerPrototypePutBleModeParams creates a new CustomerPrototypePutBleModeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypePutBleModeParams() *CustomerPrototypePutBleModeParams {
	return &CustomerPrototypePutBleModeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypePutBleModeParamsWithTimeout creates a new CustomerPrototypePutBleModeParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypePutBleModeParamsWithTimeout(timeout time.Duration) *CustomerPrototypePutBleModeParams {
	return &CustomerPrototypePutBleModeParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypePutBleModeParamsWithContext creates a new CustomerPrototypePutBleModeParams object
// with the ability to set a context for a request.
func NewCustomerPrototypePutBleModeParamsWithContext(ctx context.Context) *CustomerPrototypePutBleModeParams {
	return &CustomerPrototypePutBleModeParams{
		Context: ctx,
	}
}

// NewCustomerPrototypePutBleModeParamsWithHTTPClient creates a new CustomerPrototypePutBleModeParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypePutBleModeParamsWithHTTPClient(client *http.Client) *CustomerPrototypePutBleModeParams {
	return &CustomerPrototypePutBleModeParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypePutBleModeParams contains all the parameters to send to the API endpoint

	for the customer prototype put ble mode operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypePutBleModeParams struct {

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

	   on/off/wps/connectable
	*/
	Mode *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer prototype put ble mode params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePutBleModeParams) WithDefaults() *CustomerPrototypePutBleModeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype put ble mode params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePutBleModeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype put ble mode params
func (o *CustomerPrototypePutBleModeParams) WithTimeout(timeout time.Duration) *CustomerPrototypePutBleModeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype put ble mode params
func (o *CustomerPrototypePutBleModeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype put ble mode params
func (o *CustomerPrototypePutBleModeParams) WithContext(ctx context.Context) *CustomerPrototypePutBleModeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype put ble mode params
func (o *CustomerPrototypePutBleModeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype put ble mode params
func (o *CustomerPrototypePutBleModeParams) WithHTTPClient(client *http.Client) *CustomerPrototypePutBleModeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype put ble mode params
func (o *CustomerPrototypePutBleModeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the customer prototype put ble mode params
func (o *CustomerPrototypePutBleModeParams) WithID(id string) *CustomerPrototypePutBleModeParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype put ble mode params
func (o *CustomerPrototypePutBleModeParams) SetID(id string) {
	o.ID = id
}

// WithLocationID adds the locationID to the customer prototype put ble mode params
func (o *CustomerPrototypePutBleModeParams) WithLocationID(locationID string) *CustomerPrototypePutBleModeParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype put ble mode params
func (o *CustomerPrototypePutBleModeParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WithMode adds the mode to the customer prototype put ble mode params
func (o *CustomerPrototypePutBleModeParams) WithMode(mode *string) *CustomerPrototypePutBleModeParams {
	o.SetMode(mode)
	return o
}

// SetMode adds the mode to the customer prototype put ble mode params
func (o *CustomerPrototypePutBleModeParams) SetMode(mode *string) {
	o.Mode = mode
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypePutBleModeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Mode != nil {

		// form param mode
		var frMode string
		if o.Mode != nil {
			frMode = *o.Mode
		}
		fMode := frMode
		if fMode != "" {
			if err := r.SetFormParam("mode", fMode); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
