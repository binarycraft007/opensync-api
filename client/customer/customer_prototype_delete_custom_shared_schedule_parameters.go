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

// NewCustomerPrototypeDeleteCustomSharedScheduleParams creates a new CustomerPrototypeDeleteCustomSharedScheduleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypeDeleteCustomSharedScheduleParams() *CustomerPrototypeDeleteCustomSharedScheduleParams {
	return &CustomerPrototypeDeleteCustomSharedScheduleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypeDeleteCustomSharedScheduleParamsWithTimeout creates a new CustomerPrototypeDeleteCustomSharedScheduleParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypeDeleteCustomSharedScheduleParamsWithTimeout(timeout time.Duration) *CustomerPrototypeDeleteCustomSharedScheduleParams {
	return &CustomerPrototypeDeleteCustomSharedScheduleParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypeDeleteCustomSharedScheduleParamsWithContext creates a new CustomerPrototypeDeleteCustomSharedScheduleParams object
// with the ability to set a context for a request.
func NewCustomerPrototypeDeleteCustomSharedScheduleParamsWithContext(ctx context.Context) *CustomerPrototypeDeleteCustomSharedScheduleParams {
	return &CustomerPrototypeDeleteCustomSharedScheduleParams{
		Context: ctx,
	}
}

// NewCustomerPrototypeDeleteCustomSharedScheduleParamsWithHTTPClient creates a new CustomerPrototypeDeleteCustomSharedScheduleParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypeDeleteCustomSharedScheduleParamsWithHTTPClient(client *http.Client) *CustomerPrototypeDeleteCustomSharedScheduleParams {
	return &CustomerPrototypeDeleteCustomSharedScheduleParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypeDeleteCustomSharedScheduleParams contains all the parameters to send to the API endpoint

	for the customer prototype delete custom shared schedule operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypeDeleteCustomSharedScheduleParams struct {

	/* ID.

	   Customer id

	   Format: JSON
	*/
	ID string

	// LocationID.
	//
	// Format: JSON
	LocationID string

	// TemplateID.
	TemplateID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer prototype delete custom shared schedule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeDeleteCustomSharedScheduleParams) WithDefaults() *CustomerPrototypeDeleteCustomSharedScheduleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype delete custom shared schedule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeDeleteCustomSharedScheduleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype delete custom shared schedule params
func (o *CustomerPrototypeDeleteCustomSharedScheduleParams) WithTimeout(timeout time.Duration) *CustomerPrototypeDeleteCustomSharedScheduleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype delete custom shared schedule params
func (o *CustomerPrototypeDeleteCustomSharedScheduleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype delete custom shared schedule params
func (o *CustomerPrototypeDeleteCustomSharedScheduleParams) WithContext(ctx context.Context) *CustomerPrototypeDeleteCustomSharedScheduleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype delete custom shared schedule params
func (o *CustomerPrototypeDeleteCustomSharedScheduleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype delete custom shared schedule params
func (o *CustomerPrototypeDeleteCustomSharedScheduleParams) WithHTTPClient(client *http.Client) *CustomerPrototypeDeleteCustomSharedScheduleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype delete custom shared schedule params
func (o *CustomerPrototypeDeleteCustomSharedScheduleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the customer prototype delete custom shared schedule params
func (o *CustomerPrototypeDeleteCustomSharedScheduleParams) WithID(id string) *CustomerPrototypeDeleteCustomSharedScheduleParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype delete custom shared schedule params
func (o *CustomerPrototypeDeleteCustomSharedScheduleParams) SetID(id string) {
	o.ID = id
}

// WithLocationID adds the locationID to the customer prototype delete custom shared schedule params
func (o *CustomerPrototypeDeleteCustomSharedScheduleParams) WithLocationID(locationID string) *CustomerPrototypeDeleteCustomSharedScheduleParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype delete custom shared schedule params
func (o *CustomerPrototypeDeleteCustomSharedScheduleParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WithTemplateID adds the templateID to the customer prototype delete custom shared schedule params
func (o *CustomerPrototypeDeleteCustomSharedScheduleParams) WithTemplateID(templateID string) *CustomerPrototypeDeleteCustomSharedScheduleParams {
	o.SetTemplateID(templateID)
	return o
}

// SetTemplateID adds the templateId to the customer prototype delete custom shared schedule params
func (o *CustomerPrototypeDeleteCustomSharedScheduleParams) SetTemplateID(templateID string) {
	o.TemplateID = templateID
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypeDeleteCustomSharedScheduleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param templateId
	if err := r.SetPathParam("templateId", o.TemplateID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
