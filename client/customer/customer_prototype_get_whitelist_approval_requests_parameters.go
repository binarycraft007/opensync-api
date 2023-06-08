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

// NewCustomerPrototypeGetWhitelistApprovalRequestsParams creates a new CustomerPrototypeGetWhitelistApprovalRequestsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypeGetWhitelistApprovalRequestsParams() *CustomerPrototypeGetWhitelistApprovalRequestsParams {
	return &CustomerPrototypeGetWhitelistApprovalRequestsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypeGetWhitelistApprovalRequestsParamsWithTimeout creates a new CustomerPrototypeGetWhitelistApprovalRequestsParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypeGetWhitelistApprovalRequestsParamsWithTimeout(timeout time.Duration) *CustomerPrototypeGetWhitelistApprovalRequestsParams {
	return &CustomerPrototypeGetWhitelistApprovalRequestsParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypeGetWhitelistApprovalRequestsParamsWithContext creates a new CustomerPrototypeGetWhitelistApprovalRequestsParams object
// with the ability to set a context for a request.
func NewCustomerPrototypeGetWhitelistApprovalRequestsParamsWithContext(ctx context.Context) *CustomerPrototypeGetWhitelistApprovalRequestsParams {
	return &CustomerPrototypeGetWhitelistApprovalRequestsParams{
		Context: ctx,
	}
}

// NewCustomerPrototypeGetWhitelistApprovalRequestsParamsWithHTTPClient creates a new CustomerPrototypeGetWhitelistApprovalRequestsParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypeGetWhitelistApprovalRequestsParamsWithHTTPClient(client *http.Client) *CustomerPrototypeGetWhitelistApprovalRequestsParams {
	return &CustomerPrototypeGetWhitelistApprovalRequestsParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypeGetWhitelistApprovalRequestsParams contains all the parameters to send to the API endpoint

	for the customer prototype get whitelist approval requests operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypeGetWhitelistApprovalRequestsParams struct {

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

// WithDefaults hydrates default values in the customer prototype get whitelist approval requests params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeGetWhitelistApprovalRequestsParams) WithDefaults() *CustomerPrototypeGetWhitelistApprovalRequestsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype get whitelist approval requests params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeGetWhitelistApprovalRequestsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype get whitelist approval requests params
func (o *CustomerPrototypeGetWhitelistApprovalRequestsParams) WithTimeout(timeout time.Duration) *CustomerPrototypeGetWhitelistApprovalRequestsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype get whitelist approval requests params
func (o *CustomerPrototypeGetWhitelistApprovalRequestsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype get whitelist approval requests params
func (o *CustomerPrototypeGetWhitelistApprovalRequestsParams) WithContext(ctx context.Context) *CustomerPrototypeGetWhitelistApprovalRequestsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype get whitelist approval requests params
func (o *CustomerPrototypeGetWhitelistApprovalRequestsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype get whitelist approval requests params
func (o *CustomerPrototypeGetWhitelistApprovalRequestsParams) WithHTTPClient(client *http.Client) *CustomerPrototypeGetWhitelistApprovalRequestsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype get whitelist approval requests params
func (o *CustomerPrototypeGetWhitelistApprovalRequestsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the customer prototype get whitelist approval requests params
func (o *CustomerPrototypeGetWhitelistApprovalRequestsParams) WithID(id string) *CustomerPrototypeGetWhitelistApprovalRequestsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype get whitelist approval requests params
func (o *CustomerPrototypeGetWhitelistApprovalRequestsParams) SetID(id string) {
	o.ID = id
}

// WithLocationID adds the locationID to the customer prototype get whitelist approval requests params
func (o *CustomerPrototypeGetWhitelistApprovalRequestsParams) WithLocationID(locationID string) *CustomerPrototypeGetWhitelistApprovalRequestsParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype get whitelist approval requests params
func (o *CustomerPrototypeGetWhitelistApprovalRequestsParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypeGetWhitelistApprovalRequestsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
