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

// NewCustomerPrototypePostCaptivePortalAuthorizedClientsParams creates a new CustomerPrototypePostCaptivePortalAuthorizedClientsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypePostCaptivePortalAuthorizedClientsParams() *CustomerPrototypePostCaptivePortalAuthorizedClientsParams {
	return &CustomerPrototypePostCaptivePortalAuthorizedClientsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypePostCaptivePortalAuthorizedClientsParamsWithTimeout creates a new CustomerPrototypePostCaptivePortalAuthorizedClientsParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypePostCaptivePortalAuthorizedClientsParamsWithTimeout(timeout time.Duration) *CustomerPrototypePostCaptivePortalAuthorizedClientsParams {
	return &CustomerPrototypePostCaptivePortalAuthorizedClientsParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypePostCaptivePortalAuthorizedClientsParamsWithContext creates a new CustomerPrototypePostCaptivePortalAuthorizedClientsParams object
// with the ability to set a context for a request.
func NewCustomerPrototypePostCaptivePortalAuthorizedClientsParamsWithContext(ctx context.Context) *CustomerPrototypePostCaptivePortalAuthorizedClientsParams {
	return &CustomerPrototypePostCaptivePortalAuthorizedClientsParams{
		Context: ctx,
	}
}

// NewCustomerPrototypePostCaptivePortalAuthorizedClientsParamsWithHTTPClient creates a new CustomerPrototypePostCaptivePortalAuthorizedClientsParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypePostCaptivePortalAuthorizedClientsParamsWithHTTPClient(client *http.Client) *CustomerPrototypePostCaptivePortalAuthorizedClientsParams {
	return &CustomerPrototypePostCaptivePortalAuthorizedClientsParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypePostCaptivePortalAuthorizedClientsParams contains all the parameters to send to the API endpoint

	for the customer prototype post captive portal authorized clients operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypePostCaptivePortalAuthorizedClientsParams struct {

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

// WithDefaults hydrates default values in the customer prototype post captive portal authorized clients params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePostCaptivePortalAuthorizedClientsParams) WithDefaults() *CustomerPrototypePostCaptivePortalAuthorizedClientsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype post captive portal authorized clients params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePostCaptivePortalAuthorizedClientsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype post captive portal authorized clients params
func (o *CustomerPrototypePostCaptivePortalAuthorizedClientsParams) WithTimeout(timeout time.Duration) *CustomerPrototypePostCaptivePortalAuthorizedClientsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype post captive portal authorized clients params
func (o *CustomerPrototypePostCaptivePortalAuthorizedClientsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype post captive portal authorized clients params
func (o *CustomerPrototypePostCaptivePortalAuthorizedClientsParams) WithContext(ctx context.Context) *CustomerPrototypePostCaptivePortalAuthorizedClientsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype post captive portal authorized clients params
func (o *CustomerPrototypePostCaptivePortalAuthorizedClientsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype post captive portal authorized clients params
func (o *CustomerPrototypePostCaptivePortalAuthorizedClientsParams) WithHTTPClient(client *http.Client) *CustomerPrototypePostCaptivePortalAuthorizedClientsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype post captive portal authorized clients params
func (o *CustomerPrototypePostCaptivePortalAuthorizedClientsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExpireAt adds the expireAt to the customer prototype post captive portal authorized clients params
func (o *CustomerPrototypePostCaptivePortalAuthorizedClientsParams) WithExpireAt(expireAt string) *CustomerPrototypePostCaptivePortalAuthorizedClientsParams {
	o.SetExpireAt(expireAt)
	return o
}

// SetExpireAt adds the expireAt to the customer prototype post captive portal authorized clients params
func (o *CustomerPrototypePostCaptivePortalAuthorizedClientsParams) SetExpireAt(expireAt string) {
	o.ExpireAt = expireAt
}

// WithID adds the id to the customer prototype post captive portal authorized clients params
func (o *CustomerPrototypePostCaptivePortalAuthorizedClientsParams) WithID(id string) *CustomerPrototypePostCaptivePortalAuthorizedClientsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype post captive portal authorized clients params
func (o *CustomerPrototypePostCaptivePortalAuthorizedClientsParams) SetID(id string) {
	o.ID = id
}

// WithLocationID adds the locationID to the customer prototype post captive portal authorized clients params
func (o *CustomerPrototypePostCaptivePortalAuthorizedClientsParams) WithLocationID(locationID string) *CustomerPrototypePostCaptivePortalAuthorizedClientsParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype post captive portal authorized clients params
func (o *CustomerPrototypePostCaptivePortalAuthorizedClientsParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WithMac adds the mac to the customer prototype post captive portal authorized clients params
func (o *CustomerPrototypePostCaptivePortalAuthorizedClientsParams) WithMac(mac string) *CustomerPrototypePostCaptivePortalAuthorizedClientsParams {
	o.SetMac(mac)
	return o
}

// SetMac adds the mac to the customer prototype post captive portal authorized clients params
func (o *CustomerPrototypePostCaptivePortalAuthorizedClientsParams) SetMac(mac string) {
	o.Mac = mac
}

// WithNetworkID adds the networkID to the customer prototype post captive portal authorized clients params
func (o *CustomerPrototypePostCaptivePortalAuthorizedClientsParams) WithNetworkID(networkID string) *CustomerPrototypePostCaptivePortalAuthorizedClientsParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the customer prototype post captive portal authorized clients params
func (o *CustomerPrototypePostCaptivePortalAuthorizedClientsParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypePostCaptivePortalAuthorizedClientsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// form param mac
	frMac := o.Mac
	fMac := frMac
	if fMac != "" {
		if err := r.SetFormParam("mac", fMac); err != nil {
			return err
		}
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