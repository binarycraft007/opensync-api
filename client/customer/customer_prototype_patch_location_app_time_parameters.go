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
	"github.com/go-openapi/swag"
)

// NewCustomerPrototypePatchLocationAppTimeParams creates a new CustomerPrototypePatchLocationAppTimeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypePatchLocationAppTimeParams() *CustomerPrototypePatchLocationAppTimeParams {
	return &CustomerPrototypePatchLocationAppTimeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypePatchLocationAppTimeParamsWithTimeout creates a new CustomerPrototypePatchLocationAppTimeParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypePatchLocationAppTimeParamsWithTimeout(timeout time.Duration) *CustomerPrototypePatchLocationAppTimeParams {
	return &CustomerPrototypePatchLocationAppTimeParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypePatchLocationAppTimeParamsWithContext creates a new CustomerPrototypePatchLocationAppTimeParams object
// with the ability to set a context for a request.
func NewCustomerPrototypePatchLocationAppTimeParamsWithContext(ctx context.Context) *CustomerPrototypePatchLocationAppTimeParams {
	return &CustomerPrototypePatchLocationAppTimeParams{
		Context: ctx,
	}
}

// NewCustomerPrototypePatchLocationAppTimeParamsWithHTTPClient creates a new CustomerPrototypePatchLocationAppTimeParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypePatchLocationAppTimeParamsWithHTTPClient(client *http.Client) *CustomerPrototypePatchLocationAppTimeParams {
	return &CustomerPrototypePatchLocationAppTimeParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypePatchLocationAppTimeParams contains all the parameters to send to the API endpoint

	for the customer prototype patch location app time operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypePatchLocationAppTimeParams struct {

	// AppliesToAllDevices.
	AppliesToAllDevices *bool

	// Enable.
	Enable *bool

	/* ID.

	   Customer id

	   Format: JSON
	*/
	ID string

	// LocationID.
	LocationID string

	// SandboxSizeMb.
	//
	// Format: JSON
	SandboxSizeMb *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer prototype patch location app time params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePatchLocationAppTimeParams) WithDefaults() *CustomerPrototypePatchLocationAppTimeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype patch location app time params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePatchLocationAppTimeParams) SetDefaults() {
	var (
		appliesToAllDevicesDefault = bool(false)
	)

	val := CustomerPrototypePatchLocationAppTimeParams{
		AppliesToAllDevices: &appliesToAllDevicesDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the customer prototype patch location app time params
func (o *CustomerPrototypePatchLocationAppTimeParams) WithTimeout(timeout time.Duration) *CustomerPrototypePatchLocationAppTimeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype patch location app time params
func (o *CustomerPrototypePatchLocationAppTimeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype patch location app time params
func (o *CustomerPrototypePatchLocationAppTimeParams) WithContext(ctx context.Context) *CustomerPrototypePatchLocationAppTimeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype patch location app time params
func (o *CustomerPrototypePatchLocationAppTimeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype patch location app time params
func (o *CustomerPrototypePatchLocationAppTimeParams) WithHTTPClient(client *http.Client) *CustomerPrototypePatchLocationAppTimeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype patch location app time params
func (o *CustomerPrototypePatchLocationAppTimeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppliesToAllDevices adds the appliesToAllDevices to the customer prototype patch location app time params
func (o *CustomerPrototypePatchLocationAppTimeParams) WithAppliesToAllDevices(appliesToAllDevices *bool) *CustomerPrototypePatchLocationAppTimeParams {
	o.SetAppliesToAllDevices(appliesToAllDevices)
	return o
}

// SetAppliesToAllDevices adds the appliesToAllDevices to the customer prototype patch location app time params
func (o *CustomerPrototypePatchLocationAppTimeParams) SetAppliesToAllDevices(appliesToAllDevices *bool) {
	o.AppliesToAllDevices = appliesToAllDevices
}

// WithEnable adds the enable to the customer prototype patch location app time params
func (o *CustomerPrototypePatchLocationAppTimeParams) WithEnable(enable *bool) *CustomerPrototypePatchLocationAppTimeParams {
	o.SetEnable(enable)
	return o
}

// SetEnable adds the enable to the customer prototype patch location app time params
func (o *CustomerPrototypePatchLocationAppTimeParams) SetEnable(enable *bool) {
	o.Enable = enable
}

// WithID adds the id to the customer prototype patch location app time params
func (o *CustomerPrototypePatchLocationAppTimeParams) WithID(id string) *CustomerPrototypePatchLocationAppTimeParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype patch location app time params
func (o *CustomerPrototypePatchLocationAppTimeParams) SetID(id string) {
	o.ID = id
}

// WithLocationID adds the locationID to the customer prototype patch location app time params
func (o *CustomerPrototypePatchLocationAppTimeParams) WithLocationID(locationID string) *CustomerPrototypePatchLocationAppTimeParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype patch location app time params
func (o *CustomerPrototypePatchLocationAppTimeParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WithSandboxSizeMb adds the sandboxSizeMb to the customer prototype patch location app time params
func (o *CustomerPrototypePatchLocationAppTimeParams) WithSandboxSizeMb(sandboxSizeMb *string) *CustomerPrototypePatchLocationAppTimeParams {
	o.SetSandboxSizeMb(sandboxSizeMb)
	return o
}

// SetSandboxSizeMb adds the sandboxSizeMb to the customer prototype patch location app time params
func (o *CustomerPrototypePatchLocationAppTimeParams) SetSandboxSizeMb(sandboxSizeMb *string) {
	o.SandboxSizeMb = sandboxSizeMb
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypePatchLocationAppTimeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AppliesToAllDevices != nil {

		// form param appliesToAllDevices
		var frAppliesToAllDevices bool
		if o.AppliesToAllDevices != nil {
			frAppliesToAllDevices = *o.AppliesToAllDevices
		}
		fAppliesToAllDevices := swag.FormatBool(frAppliesToAllDevices)
		if fAppliesToAllDevices != "" {
			if err := r.SetFormParam("appliesToAllDevices", fAppliesToAllDevices); err != nil {
				return err
			}
		}
	}

	if o.Enable != nil {

		// form param enable
		var frEnable bool
		if o.Enable != nil {
			frEnable = *o.Enable
		}
		fEnable := swag.FormatBool(frEnable)
		if fEnable != "" {
			if err := r.SetFormParam("enable", fEnable); err != nil {
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

	if o.SandboxSizeMb != nil {

		// form param sandboxSizeMb
		var frSandboxSizeMb string
		if o.SandboxSizeMb != nil {
			frSandboxSizeMb = *o.SandboxSizeMb
		}
		fSandboxSizeMb := frSandboxSizeMb
		if fSandboxSizeMb != "" {
			if err := r.SetFormParam("sandboxSizeMb", fSandboxSizeMb); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
