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

// NewCustomerPrototypePatchAccessZoneParams creates a new CustomerPrototypePatchAccessZoneParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypePatchAccessZoneParams() *CustomerPrototypePatchAccessZoneParams {
	return &CustomerPrototypePatchAccessZoneParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypePatchAccessZoneParamsWithTimeout creates a new CustomerPrototypePatchAccessZoneParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypePatchAccessZoneParamsWithTimeout(timeout time.Duration) *CustomerPrototypePatchAccessZoneParams {
	return &CustomerPrototypePatchAccessZoneParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypePatchAccessZoneParamsWithContext creates a new CustomerPrototypePatchAccessZoneParams object
// with the ability to set a context for a request.
func NewCustomerPrototypePatchAccessZoneParamsWithContext(ctx context.Context) *CustomerPrototypePatchAccessZoneParams {
	return &CustomerPrototypePatchAccessZoneParams{
		Context: ctx,
	}
}

// NewCustomerPrototypePatchAccessZoneParamsWithHTTPClient creates a new CustomerPrototypePatchAccessZoneParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypePatchAccessZoneParamsWithHTTPClient(client *http.Client) *CustomerPrototypePatchAccessZoneParams {
	return &CustomerPrototypePatchAccessZoneParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypePatchAccessZoneParams contains all the parameters to send to the API endpoint

	for the customer prototype patch access zone operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypePatchAccessZoneParams struct {

	/* AccessibleDevices.

	   array of home macs[] visible to this access zone

	   Format: JSON
	*/
	AccessibleDevices *string

	// Description.
	Description *string

	/* ID.

	   Customer id

	   Format: JSON
	*/
	ID string

	// LocationID.
	//
	// Format: JSON
	LocationID string

	/* ZoneID.

	   id of access zone

	   Format: double
	*/
	ZoneID float64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer prototype patch access zone params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePatchAccessZoneParams) WithDefaults() *CustomerPrototypePatchAccessZoneParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype patch access zone params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePatchAccessZoneParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype patch access zone params
func (o *CustomerPrototypePatchAccessZoneParams) WithTimeout(timeout time.Duration) *CustomerPrototypePatchAccessZoneParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype patch access zone params
func (o *CustomerPrototypePatchAccessZoneParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype patch access zone params
func (o *CustomerPrototypePatchAccessZoneParams) WithContext(ctx context.Context) *CustomerPrototypePatchAccessZoneParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype patch access zone params
func (o *CustomerPrototypePatchAccessZoneParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype patch access zone params
func (o *CustomerPrototypePatchAccessZoneParams) WithHTTPClient(client *http.Client) *CustomerPrototypePatchAccessZoneParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype patch access zone params
func (o *CustomerPrototypePatchAccessZoneParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccessibleDevices adds the accessibleDevices to the customer prototype patch access zone params
func (o *CustomerPrototypePatchAccessZoneParams) WithAccessibleDevices(accessibleDevices *string) *CustomerPrototypePatchAccessZoneParams {
	o.SetAccessibleDevices(accessibleDevices)
	return o
}

// SetAccessibleDevices adds the accessibleDevices to the customer prototype patch access zone params
func (o *CustomerPrototypePatchAccessZoneParams) SetAccessibleDevices(accessibleDevices *string) {
	o.AccessibleDevices = accessibleDevices
}

// WithDescription adds the description to the customer prototype patch access zone params
func (o *CustomerPrototypePatchAccessZoneParams) WithDescription(description *string) *CustomerPrototypePatchAccessZoneParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the customer prototype patch access zone params
func (o *CustomerPrototypePatchAccessZoneParams) SetDescription(description *string) {
	o.Description = description
}

// WithID adds the id to the customer prototype patch access zone params
func (o *CustomerPrototypePatchAccessZoneParams) WithID(id string) *CustomerPrototypePatchAccessZoneParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype patch access zone params
func (o *CustomerPrototypePatchAccessZoneParams) SetID(id string) {
	o.ID = id
}

// WithLocationID adds the locationID to the customer prototype patch access zone params
func (o *CustomerPrototypePatchAccessZoneParams) WithLocationID(locationID string) *CustomerPrototypePatchAccessZoneParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype patch access zone params
func (o *CustomerPrototypePatchAccessZoneParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WithZoneID adds the zoneID to the customer prototype patch access zone params
func (o *CustomerPrototypePatchAccessZoneParams) WithZoneID(zoneID float64) *CustomerPrototypePatchAccessZoneParams {
	o.SetZoneID(zoneID)
	return o
}

// SetZoneID adds the zoneId to the customer prototype patch access zone params
func (o *CustomerPrototypePatchAccessZoneParams) SetZoneID(zoneID float64) {
	o.ZoneID = zoneID
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypePatchAccessZoneParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AccessibleDevices != nil {

		// form param accessibleDevices
		var frAccessibleDevices string
		if o.AccessibleDevices != nil {
			frAccessibleDevices = *o.AccessibleDevices
		}
		fAccessibleDevices := frAccessibleDevices
		if fAccessibleDevices != "" {
			if err := r.SetFormParam("accessibleDevices", fAccessibleDevices); err != nil {
				return err
			}
		}
	}

	if o.Description != nil {

		// form param description
		var frDescription string
		if o.Description != nil {
			frDescription = *o.Description
		}
		fDescription := frDescription
		if fDescription != "" {
			if err := r.SetFormParam("description", fDescription); err != nil {
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

	// path param zoneId
	if err := r.SetPathParam("zoneId", swag.FormatFloat64(o.ZoneID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
