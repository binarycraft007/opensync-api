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

// NewCustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDParams creates a new CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDParams() *CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDParams {
	return &CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDParamsWithTimeout creates a new CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDParamsWithTimeout(timeout time.Duration) *CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDParams {
	return &CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDParamsWithContext creates a new CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDParams object
// with the ability to set a context for a request.
func NewCustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDParamsWithContext(ctx context.Context) *CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDParams {
	return &CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDParams{
		Context: ctx,
	}
}

// NewCustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDParamsWithHTTPClient creates a new CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDParamsWithHTTPClient(client *http.Client) *CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDParams {
	return &CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDParams contains all the parameters to send to the API endpoint

	for the customer prototype post group of unassigned devices freeze template Id operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDParams struct {

	/* FreezeTemplateID.

	   Valid templates are uuids
	*/
	FreezeTemplateID string

	/* ID.

	   Customer id

	   Format: JSON
	*/
	ID string

	// LocationID.
	LocationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer prototype post group of unassigned devices freeze template Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDParams) WithDefaults() *CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype post group of unassigned devices freeze template Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype post group of unassigned devices freeze template Id params
func (o *CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDParams) WithTimeout(timeout time.Duration) *CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype post group of unassigned devices freeze template Id params
func (o *CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype post group of unassigned devices freeze template Id params
func (o *CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDParams) WithContext(ctx context.Context) *CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype post group of unassigned devices freeze template Id params
func (o *CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype post group of unassigned devices freeze template Id params
func (o *CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDParams) WithHTTPClient(client *http.Client) *CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype post group of unassigned devices freeze template Id params
func (o *CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFreezeTemplateID adds the freezeTemplateID to the customer prototype post group of unassigned devices freeze template Id params
func (o *CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDParams) WithFreezeTemplateID(freezeTemplateID string) *CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDParams {
	o.SetFreezeTemplateID(freezeTemplateID)
	return o
}

// SetFreezeTemplateID adds the freezeTemplateId to the customer prototype post group of unassigned devices freeze template Id params
func (o *CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDParams) SetFreezeTemplateID(freezeTemplateID string) {
	o.FreezeTemplateID = freezeTemplateID
}

// WithID adds the id to the customer prototype post group of unassigned devices freeze template Id params
func (o *CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDParams) WithID(id string) *CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype post group of unassigned devices freeze template Id params
func (o *CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDParams) SetID(id string) {
	o.ID = id
}

// WithLocationID adds the locationID to the customer prototype post group of unassigned devices freeze template Id params
func (o *CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDParams) WithLocationID(locationID string) *CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype post group of unassigned devices freeze template Id params
func (o *CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param freezeTemplateId
	if err := r.SetPathParam("freezeTemplateId", o.FreezeTemplateID); err != nil {
		return err
	}

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
