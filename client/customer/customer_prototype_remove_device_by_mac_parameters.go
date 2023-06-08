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

// NewCustomerPrototypeRemoveDeviceByMacParams creates a new CustomerPrototypeRemoveDeviceByMacParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypeRemoveDeviceByMacParams() *CustomerPrototypeRemoveDeviceByMacParams {
	return &CustomerPrototypeRemoveDeviceByMacParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypeRemoveDeviceByMacParamsWithTimeout creates a new CustomerPrototypeRemoveDeviceByMacParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypeRemoveDeviceByMacParamsWithTimeout(timeout time.Duration) *CustomerPrototypeRemoveDeviceByMacParams {
	return &CustomerPrototypeRemoveDeviceByMacParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypeRemoveDeviceByMacParamsWithContext creates a new CustomerPrototypeRemoveDeviceByMacParams object
// with the ability to set a context for a request.
func NewCustomerPrototypeRemoveDeviceByMacParamsWithContext(ctx context.Context) *CustomerPrototypeRemoveDeviceByMacParams {
	return &CustomerPrototypeRemoveDeviceByMacParams{
		Context: ctx,
	}
}

// NewCustomerPrototypeRemoveDeviceByMacParamsWithHTTPClient creates a new CustomerPrototypeRemoveDeviceByMacParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypeRemoveDeviceByMacParamsWithHTTPClient(client *http.Client) *CustomerPrototypeRemoveDeviceByMacParams {
	return &CustomerPrototypeRemoveDeviceByMacParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypeRemoveDeviceByMacParams contains all the parameters to send to the API endpoint

	for the customer prototype remove device by mac operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypeRemoveDeviceByMacParams struct {

	/* DaysOffline.

	   exclude devices disconnected longer than daysOffline. if not set, it will be 31. for older devices, it will return 404, "not found"

	   Format: double
	*/
	DaysOffline *float64

	/* ID.

	   Customer id

	   Format: JSON
	*/
	ID string

	// LocationID.
	LocationID string

	/* Mac.

	   mac of device
	*/
	Mac string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer prototype remove device by mac params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeRemoveDeviceByMacParams) WithDefaults() *CustomerPrototypeRemoveDeviceByMacParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype remove device by mac params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeRemoveDeviceByMacParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype remove device by mac params
func (o *CustomerPrototypeRemoveDeviceByMacParams) WithTimeout(timeout time.Duration) *CustomerPrototypeRemoveDeviceByMacParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype remove device by mac params
func (o *CustomerPrototypeRemoveDeviceByMacParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype remove device by mac params
func (o *CustomerPrototypeRemoveDeviceByMacParams) WithContext(ctx context.Context) *CustomerPrototypeRemoveDeviceByMacParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype remove device by mac params
func (o *CustomerPrototypeRemoveDeviceByMacParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype remove device by mac params
func (o *CustomerPrototypeRemoveDeviceByMacParams) WithHTTPClient(client *http.Client) *CustomerPrototypeRemoveDeviceByMacParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype remove device by mac params
func (o *CustomerPrototypeRemoveDeviceByMacParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDaysOffline adds the daysOffline to the customer prototype remove device by mac params
func (o *CustomerPrototypeRemoveDeviceByMacParams) WithDaysOffline(daysOffline *float64) *CustomerPrototypeRemoveDeviceByMacParams {
	o.SetDaysOffline(daysOffline)
	return o
}

// SetDaysOffline adds the daysOffline to the customer prototype remove device by mac params
func (o *CustomerPrototypeRemoveDeviceByMacParams) SetDaysOffline(daysOffline *float64) {
	o.DaysOffline = daysOffline
}

// WithID adds the id to the customer prototype remove device by mac params
func (o *CustomerPrototypeRemoveDeviceByMacParams) WithID(id string) *CustomerPrototypeRemoveDeviceByMacParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype remove device by mac params
func (o *CustomerPrototypeRemoveDeviceByMacParams) SetID(id string) {
	o.ID = id
}

// WithLocationID adds the locationID to the customer prototype remove device by mac params
func (o *CustomerPrototypeRemoveDeviceByMacParams) WithLocationID(locationID string) *CustomerPrototypeRemoveDeviceByMacParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype remove device by mac params
func (o *CustomerPrototypeRemoveDeviceByMacParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WithMac adds the mac to the customer prototype remove device by mac params
func (o *CustomerPrototypeRemoveDeviceByMacParams) WithMac(mac string) *CustomerPrototypeRemoveDeviceByMacParams {
	o.SetMac(mac)
	return o
}

// SetMac adds the mac to the customer prototype remove device by mac params
func (o *CustomerPrototypeRemoveDeviceByMacParams) SetMac(mac string) {
	o.Mac = mac
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypeRemoveDeviceByMacParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DaysOffline != nil {

		// form param daysOffline
		var frDaysOffline float64
		if o.DaysOffline != nil {
			frDaysOffline = *o.DaysOffline
		}
		fDaysOffline := swag.FormatFloat64(frDaysOffline)
		if fDaysOffline != "" {
			if err := r.SetFormParam("daysOffline", fDaysOffline); err != nil {
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

	// path param mac
	if err := r.SetPathParam("mac", o.Mac); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}