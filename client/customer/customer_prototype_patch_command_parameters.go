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

// NewCustomerPrototypePatchCommandParams creates a new CustomerPrototypePatchCommandParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypePatchCommandParams() *CustomerPrototypePatchCommandParams {
	return &CustomerPrototypePatchCommandParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypePatchCommandParamsWithTimeout creates a new CustomerPrototypePatchCommandParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypePatchCommandParamsWithTimeout(timeout time.Duration) *CustomerPrototypePatchCommandParams {
	return &CustomerPrototypePatchCommandParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypePatchCommandParamsWithContext creates a new CustomerPrototypePatchCommandParams object
// with the ability to set a context for a request.
func NewCustomerPrototypePatchCommandParamsWithContext(ctx context.Context) *CustomerPrototypePatchCommandParams {
	return &CustomerPrototypePatchCommandParams{
		Context: ctx,
	}
}

// NewCustomerPrototypePatchCommandParamsWithHTTPClient creates a new CustomerPrototypePatchCommandParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypePatchCommandParamsWithHTTPClient(client *http.Client) *CustomerPrototypePatchCommandParams {
	return &CustomerPrototypePatchCommandParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypePatchCommandParams contains all the parameters to send to the API endpoint

	for the customer prototype patch command operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypePatchCommandParams struct {

	/* ID.

	   Customer id

	   Format: JSON
	*/
	ID string

	// LocationID.
	//
	// Format: JSON
	LocationID string

	/* ProviderID.

	   enum to identify provider ex. commandAlexa
	*/
	ProviderID string

	/* ProviderUserID.

	   id of the user in provider's system
	*/
	ProviderUserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer prototype patch command params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePatchCommandParams) WithDefaults() *CustomerPrototypePatchCommandParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype patch command params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePatchCommandParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype patch command params
func (o *CustomerPrototypePatchCommandParams) WithTimeout(timeout time.Duration) *CustomerPrototypePatchCommandParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype patch command params
func (o *CustomerPrototypePatchCommandParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype patch command params
func (o *CustomerPrototypePatchCommandParams) WithContext(ctx context.Context) *CustomerPrototypePatchCommandParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype patch command params
func (o *CustomerPrototypePatchCommandParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype patch command params
func (o *CustomerPrototypePatchCommandParams) WithHTTPClient(client *http.Client) *CustomerPrototypePatchCommandParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype patch command params
func (o *CustomerPrototypePatchCommandParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the customer prototype patch command params
func (o *CustomerPrototypePatchCommandParams) WithID(id string) *CustomerPrototypePatchCommandParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype patch command params
func (o *CustomerPrototypePatchCommandParams) SetID(id string) {
	o.ID = id
}

// WithLocationID adds the locationID to the customer prototype patch command params
func (o *CustomerPrototypePatchCommandParams) WithLocationID(locationID string) *CustomerPrototypePatchCommandParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype patch command params
func (o *CustomerPrototypePatchCommandParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WithProviderID adds the providerID to the customer prototype patch command params
func (o *CustomerPrototypePatchCommandParams) WithProviderID(providerID string) *CustomerPrototypePatchCommandParams {
	o.SetProviderID(providerID)
	return o
}

// SetProviderID adds the providerId to the customer prototype patch command params
func (o *CustomerPrototypePatchCommandParams) SetProviderID(providerID string) {
	o.ProviderID = providerID
}

// WithProviderUserID adds the providerUserID to the customer prototype patch command params
func (o *CustomerPrototypePatchCommandParams) WithProviderUserID(providerUserID string) *CustomerPrototypePatchCommandParams {
	o.SetProviderUserID(providerUserID)
	return o
}

// SetProviderUserID adds the providerUserId to the customer prototype patch command params
func (o *CustomerPrototypePatchCommandParams) SetProviderUserID(providerUserID string) {
	o.ProviderUserID = providerUserID
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypePatchCommandParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// form param providerId
	frProviderID := o.ProviderID
	fProviderID := frProviderID
	if fProviderID != "" {
		if err := r.SetFormParam("providerId", fProviderID); err != nil {
			return err
		}
	}

	// form param providerUserId
	frProviderUserID := o.ProviderUserID
	fProviderUserID := frProviderUserID
	if fProviderUserID != "" {
		if err := r.SetFormParam("providerUserId", fProviderUserID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}