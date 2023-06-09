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

// NewCustomerPrototypePutUpnpParams creates a new CustomerPrototypePutUpnpParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypePutUpnpParams() *CustomerPrototypePutUpnpParams {
	return &CustomerPrototypePutUpnpParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypePutUpnpParamsWithTimeout creates a new CustomerPrototypePutUpnpParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypePutUpnpParamsWithTimeout(timeout time.Duration) *CustomerPrototypePutUpnpParams {
	return &CustomerPrototypePutUpnpParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypePutUpnpParamsWithContext creates a new CustomerPrototypePutUpnpParams object
// with the ability to set a context for a request.
func NewCustomerPrototypePutUpnpParamsWithContext(ctx context.Context) *CustomerPrototypePutUpnpParams {
	return &CustomerPrototypePutUpnpParams{
		Context: ctx,
	}
}

// NewCustomerPrototypePutUpnpParamsWithHTTPClient creates a new CustomerPrototypePutUpnpParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypePutUpnpParamsWithHTTPClient(client *http.Client) *CustomerPrototypePutUpnpParams {
	return &CustomerPrototypePutUpnpParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypePutUpnpParams contains all the parameters to send to the API endpoint

	for the customer prototype put upnp operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypePutUpnpParams struct {

	/* Enabled.

	   DEPRECATED: boolean but marked as 'any' because our mobile app platforms mixed string and boolean primitive

	   Format: JSON
	*/
	Enabled *string

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

	   Possible values enable/disable/auto
	*/
	Mode *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer prototype put upnp params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePutUpnpParams) WithDefaults() *CustomerPrototypePutUpnpParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype put upnp params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePutUpnpParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype put upnp params
func (o *CustomerPrototypePutUpnpParams) WithTimeout(timeout time.Duration) *CustomerPrototypePutUpnpParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype put upnp params
func (o *CustomerPrototypePutUpnpParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype put upnp params
func (o *CustomerPrototypePutUpnpParams) WithContext(ctx context.Context) *CustomerPrototypePutUpnpParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype put upnp params
func (o *CustomerPrototypePutUpnpParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype put upnp params
func (o *CustomerPrototypePutUpnpParams) WithHTTPClient(client *http.Client) *CustomerPrototypePutUpnpParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype put upnp params
func (o *CustomerPrototypePutUpnpParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnabled adds the enabled to the customer prototype put upnp params
func (o *CustomerPrototypePutUpnpParams) WithEnabled(enabled *string) *CustomerPrototypePutUpnpParams {
	o.SetEnabled(enabled)
	return o
}

// SetEnabled adds the enabled to the customer prototype put upnp params
func (o *CustomerPrototypePutUpnpParams) SetEnabled(enabled *string) {
	o.Enabled = enabled
}

// WithID adds the id to the customer prototype put upnp params
func (o *CustomerPrototypePutUpnpParams) WithID(id string) *CustomerPrototypePutUpnpParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype put upnp params
func (o *CustomerPrototypePutUpnpParams) SetID(id string) {
	o.ID = id
}

// WithLocationID adds the locationID to the customer prototype put upnp params
func (o *CustomerPrototypePutUpnpParams) WithLocationID(locationID string) *CustomerPrototypePutUpnpParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype put upnp params
func (o *CustomerPrototypePutUpnpParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WithMode adds the mode to the customer prototype put upnp params
func (o *CustomerPrototypePutUpnpParams) WithMode(mode *string) *CustomerPrototypePutUpnpParams {
	o.SetMode(mode)
	return o
}

// SetMode adds the mode to the customer prototype put upnp params
func (o *CustomerPrototypePutUpnpParams) SetMode(mode *string) {
	o.Mode = mode
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypePutUpnpParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Enabled != nil {

		// form param enabled
		var frEnabled string
		if o.Enabled != nil {
			frEnabled = *o.Enabled
		}
		fEnabled := frEnabled
		if fEnabled != "" {
			if err := r.SetFormParam("enabled", fEnabled); err != nil {
				return err
			}
		}
	}

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
