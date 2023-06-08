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

// NewCustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesWhitelistParams creates a new CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesWhitelistParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesWhitelistParams() *CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesWhitelistParams {
	return &CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesWhitelistParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesWhitelistParamsWithTimeout creates a new CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesWhitelistParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesWhitelistParamsWithTimeout(timeout time.Duration) *CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesWhitelistParams {
	return &CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesWhitelistParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesWhitelistParamsWithContext creates a new CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesWhitelistParams object
// with the ability to set a context for a request.
func NewCustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesWhitelistParamsWithContext(ctx context.Context) *CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesWhitelistParams {
	return &CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesWhitelistParams{
		Context: ctx,
	}
}

// NewCustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesWhitelistParamsWithHTTPClient creates a new CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesWhitelistParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesWhitelistParamsWithHTTPClient(client *http.Client) *CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesWhitelistParams {
	return &CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesWhitelistParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesWhitelistParams contains all the parameters to send to the API endpoint

	for the customer prototype delete from device security policy websites whitelist operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesWhitelistParams struct {

	// DNS.
	DNS string

	/* ID.

	   Customer id

	   Format: JSON
	*/
	ID string

	// LocationID.
	LocationID string

	// Mac.
	Mac string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer prototype delete from device security policy websites whitelist params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesWhitelistParams) WithDefaults() *CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesWhitelistParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype delete from device security policy websites whitelist params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesWhitelistParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype delete from device security policy websites whitelist params
func (o *CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesWhitelistParams) WithTimeout(timeout time.Duration) *CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesWhitelistParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype delete from device security policy websites whitelist params
func (o *CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesWhitelistParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype delete from device security policy websites whitelist params
func (o *CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesWhitelistParams) WithContext(ctx context.Context) *CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesWhitelistParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype delete from device security policy websites whitelist params
func (o *CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesWhitelistParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype delete from device security policy websites whitelist params
func (o *CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesWhitelistParams) WithHTTPClient(client *http.Client) *CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesWhitelistParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype delete from device security policy websites whitelist params
func (o *CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesWhitelistParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDNS adds the dns to the customer prototype delete from device security policy websites whitelist params
func (o *CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesWhitelistParams) WithDNS(dns string) *CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesWhitelistParams {
	o.SetDNS(dns)
	return o
}

// SetDNS adds the dns to the customer prototype delete from device security policy websites whitelist params
func (o *CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesWhitelistParams) SetDNS(dns string) {
	o.DNS = dns
}

// WithID adds the id to the customer prototype delete from device security policy websites whitelist params
func (o *CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesWhitelistParams) WithID(id string) *CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesWhitelistParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype delete from device security policy websites whitelist params
func (o *CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesWhitelistParams) SetID(id string) {
	o.ID = id
}

// WithLocationID adds the locationID to the customer prototype delete from device security policy websites whitelist params
func (o *CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesWhitelistParams) WithLocationID(locationID string) *CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesWhitelistParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype delete from device security policy websites whitelist params
func (o *CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesWhitelistParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WithMac adds the mac to the customer prototype delete from device security policy websites whitelist params
func (o *CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesWhitelistParams) WithMac(mac string) *CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesWhitelistParams {
	o.SetMac(mac)
	return o
}

// SetMac adds the mac to the customer prototype delete from device security policy websites whitelist params
func (o *CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesWhitelistParams) SetMac(mac string) {
	o.Mac = mac
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesWhitelistParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param dns
	if err := r.SetPathParam("dns", o.DNS); err != nil {
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

	// path param mac
	if err := r.SetPathParam("mac", o.Mac); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}