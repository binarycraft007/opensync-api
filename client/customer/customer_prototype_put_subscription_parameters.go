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

// NewCustomerPrototypePutSubscriptionParams creates a new CustomerPrototypePutSubscriptionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypePutSubscriptionParams() *CustomerPrototypePutSubscriptionParams {
	return &CustomerPrototypePutSubscriptionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypePutSubscriptionParamsWithTimeout creates a new CustomerPrototypePutSubscriptionParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypePutSubscriptionParamsWithTimeout(timeout time.Duration) *CustomerPrototypePutSubscriptionParams {
	return &CustomerPrototypePutSubscriptionParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypePutSubscriptionParamsWithContext creates a new CustomerPrototypePutSubscriptionParams object
// with the ability to set a context for a request.
func NewCustomerPrototypePutSubscriptionParamsWithContext(ctx context.Context) *CustomerPrototypePutSubscriptionParams {
	return &CustomerPrototypePutSubscriptionParams{
		Context: ctx,
	}
}

// NewCustomerPrototypePutSubscriptionParamsWithHTTPClient creates a new CustomerPrototypePutSubscriptionParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypePutSubscriptionParamsWithHTTPClient(client *http.Client) *CustomerPrototypePutSubscriptionParams {
	return &CustomerPrototypePutSubscriptionParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypePutSubscriptionParams contains all the parameters to send to the API endpoint

	for the customer prototype put subscription operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypePutSubscriptionParams struct {

	/* ID.

	   Customer id

	   Format: JSON
	*/
	ID string

	// LocationID.
	LocationID string

	// RatePlanID.
	RatePlanID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer prototype put subscription params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePutSubscriptionParams) WithDefaults() *CustomerPrototypePutSubscriptionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype put subscription params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePutSubscriptionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype put subscription params
func (o *CustomerPrototypePutSubscriptionParams) WithTimeout(timeout time.Duration) *CustomerPrototypePutSubscriptionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype put subscription params
func (o *CustomerPrototypePutSubscriptionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype put subscription params
func (o *CustomerPrototypePutSubscriptionParams) WithContext(ctx context.Context) *CustomerPrototypePutSubscriptionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype put subscription params
func (o *CustomerPrototypePutSubscriptionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype put subscription params
func (o *CustomerPrototypePutSubscriptionParams) WithHTTPClient(client *http.Client) *CustomerPrototypePutSubscriptionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype put subscription params
func (o *CustomerPrototypePutSubscriptionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the customer prototype put subscription params
func (o *CustomerPrototypePutSubscriptionParams) WithID(id string) *CustomerPrototypePutSubscriptionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype put subscription params
func (o *CustomerPrototypePutSubscriptionParams) SetID(id string) {
	o.ID = id
}

// WithLocationID adds the locationID to the customer prototype put subscription params
func (o *CustomerPrototypePutSubscriptionParams) WithLocationID(locationID string) *CustomerPrototypePutSubscriptionParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype put subscription params
func (o *CustomerPrototypePutSubscriptionParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WithRatePlanID adds the ratePlanID to the customer prototype put subscription params
func (o *CustomerPrototypePutSubscriptionParams) WithRatePlanID(ratePlanID string) *CustomerPrototypePutSubscriptionParams {
	o.SetRatePlanID(ratePlanID)
	return o
}

// SetRatePlanID adds the ratePlanId to the customer prototype put subscription params
func (o *CustomerPrototypePutSubscriptionParams) SetRatePlanID(ratePlanID string) {
	o.RatePlanID = ratePlanID
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypePutSubscriptionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// form param ratePlanId
	frRatePlanID := o.RatePlanID
	fRatePlanID := frRatePlanID
	if fRatePlanID != "" {
		if err := r.SetFormParam("ratePlanId", fRatePlanID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
