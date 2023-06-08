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

// NewCustomerPrototypeApproveWhitelistRequestParams creates a new CustomerPrototypeApproveWhitelistRequestParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypeApproveWhitelistRequestParams() *CustomerPrototypeApproveWhitelistRequestParams {
	return &CustomerPrototypeApproveWhitelistRequestParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypeApproveWhitelistRequestParamsWithTimeout creates a new CustomerPrototypeApproveWhitelistRequestParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypeApproveWhitelistRequestParamsWithTimeout(timeout time.Duration) *CustomerPrototypeApproveWhitelistRequestParams {
	return &CustomerPrototypeApproveWhitelistRequestParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypeApproveWhitelistRequestParamsWithContext creates a new CustomerPrototypeApproveWhitelistRequestParams object
// with the ability to set a context for a request.
func NewCustomerPrototypeApproveWhitelistRequestParamsWithContext(ctx context.Context) *CustomerPrototypeApproveWhitelistRequestParams {
	return &CustomerPrototypeApproveWhitelistRequestParams{
		Context: ctx,
	}
}

// NewCustomerPrototypeApproveWhitelistRequestParamsWithHTTPClient creates a new CustomerPrototypeApproveWhitelistRequestParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypeApproveWhitelistRequestParamsWithHTTPClient(client *http.Client) *CustomerPrototypeApproveWhitelistRequestParams {
	return &CustomerPrototypeApproveWhitelistRequestParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypeApproveWhitelistRequestParams contains all the parameters to send to the API endpoint

	for the customer prototype approve whitelist request operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypeApproveWhitelistRequestParams struct {

	/* ID.

	   Customer id

	   Format: JSON
	*/
	ID string

	// LocationID.
	LocationID string

	// Persons.
	//
	// Format: JSON
	Persons string

	// RequestID.
	RequestID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer prototype approve whitelist request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeApproveWhitelistRequestParams) WithDefaults() *CustomerPrototypeApproveWhitelistRequestParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype approve whitelist request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeApproveWhitelistRequestParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype approve whitelist request params
func (o *CustomerPrototypeApproveWhitelistRequestParams) WithTimeout(timeout time.Duration) *CustomerPrototypeApproveWhitelistRequestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype approve whitelist request params
func (o *CustomerPrototypeApproveWhitelistRequestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype approve whitelist request params
func (o *CustomerPrototypeApproveWhitelistRequestParams) WithContext(ctx context.Context) *CustomerPrototypeApproveWhitelistRequestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype approve whitelist request params
func (o *CustomerPrototypeApproveWhitelistRequestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype approve whitelist request params
func (o *CustomerPrototypeApproveWhitelistRequestParams) WithHTTPClient(client *http.Client) *CustomerPrototypeApproveWhitelistRequestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype approve whitelist request params
func (o *CustomerPrototypeApproveWhitelistRequestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the customer prototype approve whitelist request params
func (o *CustomerPrototypeApproveWhitelistRequestParams) WithID(id string) *CustomerPrototypeApproveWhitelistRequestParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype approve whitelist request params
func (o *CustomerPrototypeApproveWhitelistRequestParams) SetID(id string) {
	o.ID = id
}

// WithLocationID adds the locationID to the customer prototype approve whitelist request params
func (o *CustomerPrototypeApproveWhitelistRequestParams) WithLocationID(locationID string) *CustomerPrototypeApproveWhitelistRequestParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype approve whitelist request params
func (o *CustomerPrototypeApproveWhitelistRequestParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WithPersons adds the persons to the customer prototype approve whitelist request params
func (o *CustomerPrototypeApproveWhitelistRequestParams) WithPersons(persons string) *CustomerPrototypeApproveWhitelistRequestParams {
	o.SetPersons(persons)
	return o
}

// SetPersons adds the persons to the customer prototype approve whitelist request params
func (o *CustomerPrototypeApproveWhitelistRequestParams) SetPersons(persons string) {
	o.Persons = persons
}

// WithRequestID adds the requestID to the customer prototype approve whitelist request params
func (o *CustomerPrototypeApproveWhitelistRequestParams) WithRequestID(requestID string) *CustomerPrototypeApproveWhitelistRequestParams {
	o.SetRequestID(requestID)
	return o
}

// SetRequestID adds the requestId to the customer prototype approve whitelist request params
func (o *CustomerPrototypeApproveWhitelistRequestParams) SetRequestID(requestID string) {
	o.RequestID = requestID
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypeApproveWhitelistRequestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// form param persons
	frPersons := o.Persons
	fPersons := frPersons
	if fPersons != "" {
		if err := r.SetFormParam("persons", fPersons); err != nil {
			return err
		}
	}

	// path param requestId
	if err := r.SetPathParam("requestId", o.RequestID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
