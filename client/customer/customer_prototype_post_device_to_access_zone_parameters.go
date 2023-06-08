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

// NewCustomerPrototypePostDeviceToAccessZoneParams creates a new CustomerPrototypePostDeviceToAccessZoneParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypePostDeviceToAccessZoneParams() *CustomerPrototypePostDeviceToAccessZoneParams {
	return &CustomerPrototypePostDeviceToAccessZoneParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypePostDeviceToAccessZoneParamsWithTimeout creates a new CustomerPrototypePostDeviceToAccessZoneParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypePostDeviceToAccessZoneParamsWithTimeout(timeout time.Duration) *CustomerPrototypePostDeviceToAccessZoneParams {
	return &CustomerPrototypePostDeviceToAccessZoneParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypePostDeviceToAccessZoneParamsWithContext creates a new CustomerPrototypePostDeviceToAccessZoneParams object
// with the ability to set a context for a request.
func NewCustomerPrototypePostDeviceToAccessZoneParamsWithContext(ctx context.Context) *CustomerPrototypePostDeviceToAccessZoneParams {
	return &CustomerPrototypePostDeviceToAccessZoneParams{
		Context: ctx,
	}
}

// NewCustomerPrototypePostDeviceToAccessZoneParamsWithHTTPClient creates a new CustomerPrototypePostDeviceToAccessZoneParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypePostDeviceToAccessZoneParamsWithHTTPClient(client *http.Client) *CustomerPrototypePostDeviceToAccessZoneParams {
	return &CustomerPrototypePostDeviceToAccessZoneParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypePostDeviceToAccessZoneParams contains all the parameters to send to the API endpoint

	for the customer prototype post device to access zone operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypePostDeviceToAccessZoneParams struct {

	/* ID.

	   Customer id

	   Format: JSON
	*/
	ID string

	// LocationID.
	//
	// Format: JSON
	LocationID string

	/* Mac.

	   the device mac to be added to the access zone
	*/
	Mac string

	/* ZoneID.

	   id of access zone

	   Format: double
	*/
	ZoneID float64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer prototype post device to access zone params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePostDeviceToAccessZoneParams) WithDefaults() *CustomerPrototypePostDeviceToAccessZoneParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype post device to access zone params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePostDeviceToAccessZoneParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype post device to access zone params
func (o *CustomerPrototypePostDeviceToAccessZoneParams) WithTimeout(timeout time.Duration) *CustomerPrototypePostDeviceToAccessZoneParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype post device to access zone params
func (o *CustomerPrototypePostDeviceToAccessZoneParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype post device to access zone params
func (o *CustomerPrototypePostDeviceToAccessZoneParams) WithContext(ctx context.Context) *CustomerPrototypePostDeviceToAccessZoneParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype post device to access zone params
func (o *CustomerPrototypePostDeviceToAccessZoneParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype post device to access zone params
func (o *CustomerPrototypePostDeviceToAccessZoneParams) WithHTTPClient(client *http.Client) *CustomerPrototypePostDeviceToAccessZoneParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype post device to access zone params
func (o *CustomerPrototypePostDeviceToAccessZoneParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the customer prototype post device to access zone params
func (o *CustomerPrototypePostDeviceToAccessZoneParams) WithID(id string) *CustomerPrototypePostDeviceToAccessZoneParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype post device to access zone params
func (o *CustomerPrototypePostDeviceToAccessZoneParams) SetID(id string) {
	o.ID = id
}

// WithLocationID adds the locationID to the customer prototype post device to access zone params
func (o *CustomerPrototypePostDeviceToAccessZoneParams) WithLocationID(locationID string) *CustomerPrototypePostDeviceToAccessZoneParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype post device to access zone params
func (o *CustomerPrototypePostDeviceToAccessZoneParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WithMac adds the mac to the customer prototype post device to access zone params
func (o *CustomerPrototypePostDeviceToAccessZoneParams) WithMac(mac string) *CustomerPrototypePostDeviceToAccessZoneParams {
	o.SetMac(mac)
	return o
}

// SetMac adds the mac to the customer prototype post device to access zone params
func (o *CustomerPrototypePostDeviceToAccessZoneParams) SetMac(mac string) {
	o.Mac = mac
}

// WithZoneID adds the zoneID to the customer prototype post device to access zone params
func (o *CustomerPrototypePostDeviceToAccessZoneParams) WithZoneID(zoneID float64) *CustomerPrototypePostDeviceToAccessZoneParams {
	o.SetZoneID(zoneID)
	return o
}

// SetZoneID adds the zoneId to the customer prototype post device to access zone params
func (o *CustomerPrototypePostDeviceToAccessZoneParams) SetZoneID(zoneID float64) {
	o.ZoneID = zoneID
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypePostDeviceToAccessZoneParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param mac
	if err := r.SetPathParam("mac", o.Mac); err != nil {
		return err
	}

	// path param zoneId
	if err := r.SetPathParam("zoneId", swag.FormatFloat64(o.ZoneID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
