// Code generated by go-swagger; DO NOT EDIT.

package location

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

// NewLocationPrototypeCheckCustomerLoginToTurnOffWifiParams creates a new LocationPrototypeCheckCustomerLoginToTurnOffWifiParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLocationPrototypeCheckCustomerLoginToTurnOffWifiParams() *LocationPrototypeCheckCustomerLoginToTurnOffWifiParams {
	return &LocationPrototypeCheckCustomerLoginToTurnOffWifiParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLocationPrototypeCheckCustomerLoginToTurnOffWifiParamsWithTimeout creates a new LocationPrototypeCheckCustomerLoginToTurnOffWifiParams object
// with the ability to set a timeout on a request.
func NewLocationPrototypeCheckCustomerLoginToTurnOffWifiParamsWithTimeout(timeout time.Duration) *LocationPrototypeCheckCustomerLoginToTurnOffWifiParams {
	return &LocationPrototypeCheckCustomerLoginToTurnOffWifiParams{
		timeout: timeout,
	}
}

// NewLocationPrototypeCheckCustomerLoginToTurnOffWifiParamsWithContext creates a new LocationPrototypeCheckCustomerLoginToTurnOffWifiParams object
// with the ability to set a context for a request.
func NewLocationPrototypeCheckCustomerLoginToTurnOffWifiParamsWithContext(ctx context.Context) *LocationPrototypeCheckCustomerLoginToTurnOffWifiParams {
	return &LocationPrototypeCheckCustomerLoginToTurnOffWifiParams{
		Context: ctx,
	}
}

// NewLocationPrototypeCheckCustomerLoginToTurnOffWifiParamsWithHTTPClient creates a new LocationPrototypeCheckCustomerLoginToTurnOffWifiParams object
// with the ability to set a custom HTTPClient for a request.
func NewLocationPrototypeCheckCustomerLoginToTurnOffWifiParamsWithHTTPClient(client *http.Client) *LocationPrototypeCheckCustomerLoginToTurnOffWifiParams {
	return &LocationPrototypeCheckCustomerLoginToTurnOffWifiParams{
		HTTPClient: client,
	}
}

/*
LocationPrototypeCheckCustomerLoginToTurnOffWifiParams contains all the parameters to send to the API endpoint

	for the location prototype check customer login to turn off wifi operation.

	Typically these are written to a http.Request.
*/
type LocationPrototypeCheckCustomerLoginToTurnOffWifiParams struct {

	/* ID.

	   Location id

	   Format: JSON
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the location prototype check customer login to turn off wifi params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LocationPrototypeCheckCustomerLoginToTurnOffWifiParams) WithDefaults() *LocationPrototypeCheckCustomerLoginToTurnOffWifiParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the location prototype check customer login to turn off wifi params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LocationPrototypeCheckCustomerLoginToTurnOffWifiParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the location prototype check customer login to turn off wifi params
func (o *LocationPrototypeCheckCustomerLoginToTurnOffWifiParams) WithTimeout(timeout time.Duration) *LocationPrototypeCheckCustomerLoginToTurnOffWifiParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the location prototype check customer login to turn off wifi params
func (o *LocationPrototypeCheckCustomerLoginToTurnOffWifiParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the location prototype check customer login to turn off wifi params
func (o *LocationPrototypeCheckCustomerLoginToTurnOffWifiParams) WithContext(ctx context.Context) *LocationPrototypeCheckCustomerLoginToTurnOffWifiParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the location prototype check customer login to turn off wifi params
func (o *LocationPrototypeCheckCustomerLoginToTurnOffWifiParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the location prototype check customer login to turn off wifi params
func (o *LocationPrototypeCheckCustomerLoginToTurnOffWifiParams) WithHTTPClient(client *http.Client) *LocationPrototypeCheckCustomerLoginToTurnOffWifiParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the location prototype check customer login to turn off wifi params
func (o *LocationPrototypeCheckCustomerLoginToTurnOffWifiParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the location prototype check customer login to turn off wifi params
func (o *LocationPrototypeCheckCustomerLoginToTurnOffWifiParams) WithID(id string) *LocationPrototypeCheckCustomerLoginToTurnOffWifiParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the location prototype check customer login to turn off wifi params
func (o *LocationPrototypeCheckCustomerLoginToTurnOffWifiParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *LocationPrototypeCheckCustomerLoginToTurnOffWifiParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
