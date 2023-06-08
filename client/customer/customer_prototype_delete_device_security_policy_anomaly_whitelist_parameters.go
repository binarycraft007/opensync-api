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

// NewCustomerPrototypeDeleteDeviceSecurityPolicyAnomalyWhitelistParams creates a new CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyWhitelistParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypeDeleteDeviceSecurityPolicyAnomalyWhitelistParams() *CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyWhitelistParams {
	return &CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyWhitelistParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypeDeleteDeviceSecurityPolicyAnomalyWhitelistParamsWithTimeout creates a new CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyWhitelistParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypeDeleteDeviceSecurityPolicyAnomalyWhitelistParamsWithTimeout(timeout time.Duration) *CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyWhitelistParams {
	return &CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyWhitelistParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypeDeleteDeviceSecurityPolicyAnomalyWhitelistParamsWithContext creates a new CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyWhitelistParams object
// with the ability to set a context for a request.
func NewCustomerPrototypeDeleteDeviceSecurityPolicyAnomalyWhitelistParamsWithContext(ctx context.Context) *CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyWhitelistParams {
	return &CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyWhitelistParams{
		Context: ctx,
	}
}

// NewCustomerPrototypeDeleteDeviceSecurityPolicyAnomalyWhitelistParamsWithHTTPClient creates a new CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyWhitelistParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypeDeleteDeviceSecurityPolicyAnomalyWhitelistParamsWithHTTPClient(client *http.Client) *CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyWhitelistParams {
	return &CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyWhitelistParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyWhitelistParams contains all the parameters to send to the API endpoint

	for the customer prototype delete device security policy anomaly whitelist operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyWhitelistParams struct {

	// Fqdn.
	Fqdn string

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

// WithDefaults hydrates default values in the customer prototype delete device security policy anomaly whitelist params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyWhitelistParams) WithDefaults() *CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyWhitelistParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype delete device security policy anomaly whitelist params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyWhitelistParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype delete device security policy anomaly whitelist params
func (o *CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyWhitelistParams) WithTimeout(timeout time.Duration) *CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyWhitelistParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype delete device security policy anomaly whitelist params
func (o *CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyWhitelistParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype delete device security policy anomaly whitelist params
func (o *CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyWhitelistParams) WithContext(ctx context.Context) *CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyWhitelistParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype delete device security policy anomaly whitelist params
func (o *CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyWhitelistParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype delete device security policy anomaly whitelist params
func (o *CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyWhitelistParams) WithHTTPClient(client *http.Client) *CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyWhitelistParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype delete device security policy anomaly whitelist params
func (o *CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyWhitelistParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFqdn adds the fqdn to the customer prototype delete device security policy anomaly whitelist params
func (o *CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyWhitelistParams) WithFqdn(fqdn string) *CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyWhitelistParams {
	o.SetFqdn(fqdn)
	return o
}

// SetFqdn adds the fqdn to the customer prototype delete device security policy anomaly whitelist params
func (o *CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyWhitelistParams) SetFqdn(fqdn string) {
	o.Fqdn = fqdn
}

// WithID adds the id to the customer prototype delete device security policy anomaly whitelist params
func (o *CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyWhitelistParams) WithID(id string) *CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyWhitelistParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype delete device security policy anomaly whitelist params
func (o *CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyWhitelistParams) SetID(id string) {
	o.ID = id
}

// WithLocationID adds the locationID to the customer prototype delete device security policy anomaly whitelist params
func (o *CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyWhitelistParams) WithLocationID(locationID string) *CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyWhitelistParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype delete device security policy anomaly whitelist params
func (o *CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyWhitelistParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WithMac adds the mac to the customer prototype delete device security policy anomaly whitelist params
func (o *CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyWhitelistParams) WithMac(mac string) *CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyWhitelistParams {
	o.SetMac(mac)
	return o
}

// SetMac adds the mac to the customer prototype delete device security policy anomaly whitelist params
func (o *CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyWhitelistParams) SetMac(mac string) {
	o.Mac = mac
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyWhitelistParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param fqdn
	if err := r.SetPathParam("fqdn", o.Fqdn); err != nil {
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