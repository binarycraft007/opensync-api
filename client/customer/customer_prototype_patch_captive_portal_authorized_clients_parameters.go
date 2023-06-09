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

// NewCustomerPrototypePatchCaptivePortalAuthorizedClientsParams creates a new CustomerPrototypePatchCaptivePortalAuthorizedClientsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypePatchCaptivePortalAuthorizedClientsParams() *CustomerPrototypePatchCaptivePortalAuthorizedClientsParams {
	return &CustomerPrototypePatchCaptivePortalAuthorizedClientsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypePatchCaptivePortalAuthorizedClientsParamsWithTimeout creates a new CustomerPrototypePatchCaptivePortalAuthorizedClientsParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypePatchCaptivePortalAuthorizedClientsParamsWithTimeout(timeout time.Duration) *CustomerPrototypePatchCaptivePortalAuthorizedClientsParams {
	return &CustomerPrototypePatchCaptivePortalAuthorizedClientsParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypePatchCaptivePortalAuthorizedClientsParamsWithContext creates a new CustomerPrototypePatchCaptivePortalAuthorizedClientsParams object
// with the ability to set a context for a request.
func NewCustomerPrototypePatchCaptivePortalAuthorizedClientsParamsWithContext(ctx context.Context) *CustomerPrototypePatchCaptivePortalAuthorizedClientsParams {
	return &CustomerPrototypePatchCaptivePortalAuthorizedClientsParams{
		Context: ctx,
	}
}

// NewCustomerPrototypePatchCaptivePortalAuthorizedClientsParamsWithHTTPClient creates a new CustomerPrototypePatchCaptivePortalAuthorizedClientsParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypePatchCaptivePortalAuthorizedClientsParamsWithHTTPClient(client *http.Client) *CustomerPrototypePatchCaptivePortalAuthorizedClientsParams {
	return &CustomerPrototypePatchCaptivePortalAuthorizedClientsParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypePatchCaptivePortalAuthorizedClientsParams contains all the parameters to send to the API endpoint

	for the customer prototype patch captive portal authorized clients operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypePatchCaptivePortalAuthorizedClientsParams struct {

	// ExpireAt.
	ExpireAt string

	/* ID.

	   Customer id

	   Format: JSON
	*/
	ID string

	// LocationID.
	//
	// Format: JSON
	LocationID string

	// Mac.
	Mac string

	// NetworkID.
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer prototype patch captive portal authorized clients params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePatchCaptivePortalAuthorizedClientsParams) WithDefaults() *CustomerPrototypePatchCaptivePortalAuthorizedClientsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype patch captive portal authorized clients params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePatchCaptivePortalAuthorizedClientsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype patch captive portal authorized clients params
func (o *CustomerPrototypePatchCaptivePortalAuthorizedClientsParams) WithTimeout(timeout time.Duration) *CustomerPrototypePatchCaptivePortalAuthorizedClientsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype patch captive portal authorized clients params
func (o *CustomerPrototypePatchCaptivePortalAuthorizedClientsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype patch captive portal authorized clients params
func (o *CustomerPrototypePatchCaptivePortalAuthorizedClientsParams) WithContext(ctx context.Context) *CustomerPrototypePatchCaptivePortalAuthorizedClientsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype patch captive portal authorized clients params
func (o *CustomerPrototypePatchCaptivePortalAuthorizedClientsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype patch captive portal authorized clients params
func (o *CustomerPrototypePatchCaptivePortalAuthorizedClientsParams) WithHTTPClient(client *http.Client) *CustomerPrototypePatchCaptivePortalAuthorizedClientsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype patch captive portal authorized clients params
func (o *CustomerPrototypePatchCaptivePortalAuthorizedClientsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExpireAt adds the expireAt to the customer prototype patch captive portal authorized clients params
func (o *CustomerPrototypePatchCaptivePortalAuthorizedClientsParams) WithExpireAt(expireAt string) *CustomerPrototypePatchCaptivePortalAuthorizedClientsParams {
	o.SetExpireAt(expireAt)
	return o
}

// SetExpireAt adds the expireAt to the customer prototype patch captive portal authorized clients params
func (o *CustomerPrototypePatchCaptivePortalAuthorizedClientsParams) SetExpireAt(expireAt string) {
	o.ExpireAt = expireAt
}

// WithID adds the id to the customer prototype patch captive portal authorized clients params
func (o *CustomerPrototypePatchCaptivePortalAuthorizedClientsParams) WithID(id string) *CustomerPrototypePatchCaptivePortalAuthorizedClientsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype patch captive portal authorized clients params
func (o *CustomerPrototypePatchCaptivePortalAuthorizedClientsParams) SetID(id string) {
	o.ID = id
}

// WithLocationID adds the locationID to the customer prototype patch captive portal authorized clients params
func (o *CustomerPrototypePatchCaptivePortalAuthorizedClientsParams) WithLocationID(locationID string) *CustomerPrototypePatchCaptivePortalAuthorizedClientsParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype patch captive portal authorized clients params
func (o *CustomerPrototypePatchCaptivePortalAuthorizedClientsParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WithMac adds the mac to the customer prototype patch captive portal authorized clients params
func (o *CustomerPrototypePatchCaptivePortalAuthorizedClientsParams) WithMac(mac string) *CustomerPrototypePatchCaptivePortalAuthorizedClientsParams {
	o.SetMac(mac)
	return o
}

// SetMac adds the mac to the customer prototype patch captive portal authorized clients params
func (o *CustomerPrototypePatchCaptivePortalAuthorizedClientsParams) SetMac(mac string) {
	o.Mac = mac
}

// WithNetworkID adds the networkID to the customer prototype patch captive portal authorized clients params
func (o *CustomerPrototypePatchCaptivePortalAuthorizedClientsParams) WithNetworkID(networkID string) *CustomerPrototypePatchCaptivePortalAuthorizedClientsParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the customer prototype patch captive portal authorized clients params
func (o *CustomerPrototypePatchCaptivePortalAuthorizedClientsParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypePatchCaptivePortalAuthorizedClientsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// form param expireAt
	frExpireAt := o.ExpireAt
	fExpireAt := frExpireAt
	if fExpireAt != "" {
		if err := r.SetFormParam("expireAt", fExpireAt); err != nil {
			return err
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

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
