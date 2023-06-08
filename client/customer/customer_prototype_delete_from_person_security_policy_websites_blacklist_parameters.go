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

// NewCustomerPrototypeDeleteFromPersonSecurityPolicyWebsitesBlacklistParams creates a new CustomerPrototypeDeleteFromPersonSecurityPolicyWebsitesBlacklistParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypeDeleteFromPersonSecurityPolicyWebsitesBlacklistParams() *CustomerPrototypeDeleteFromPersonSecurityPolicyWebsitesBlacklistParams {
	return &CustomerPrototypeDeleteFromPersonSecurityPolicyWebsitesBlacklistParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypeDeleteFromPersonSecurityPolicyWebsitesBlacklistParamsWithTimeout creates a new CustomerPrototypeDeleteFromPersonSecurityPolicyWebsitesBlacklistParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypeDeleteFromPersonSecurityPolicyWebsitesBlacklistParamsWithTimeout(timeout time.Duration) *CustomerPrototypeDeleteFromPersonSecurityPolicyWebsitesBlacklistParams {
	return &CustomerPrototypeDeleteFromPersonSecurityPolicyWebsitesBlacklistParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypeDeleteFromPersonSecurityPolicyWebsitesBlacklistParamsWithContext creates a new CustomerPrototypeDeleteFromPersonSecurityPolicyWebsitesBlacklistParams object
// with the ability to set a context for a request.
func NewCustomerPrototypeDeleteFromPersonSecurityPolicyWebsitesBlacklistParamsWithContext(ctx context.Context) *CustomerPrototypeDeleteFromPersonSecurityPolicyWebsitesBlacklistParams {
	return &CustomerPrototypeDeleteFromPersonSecurityPolicyWebsitesBlacklistParams{
		Context: ctx,
	}
}

// NewCustomerPrototypeDeleteFromPersonSecurityPolicyWebsitesBlacklistParamsWithHTTPClient creates a new CustomerPrototypeDeleteFromPersonSecurityPolicyWebsitesBlacklistParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypeDeleteFromPersonSecurityPolicyWebsitesBlacklistParamsWithHTTPClient(client *http.Client) *CustomerPrototypeDeleteFromPersonSecurityPolicyWebsitesBlacklistParams {
	return &CustomerPrototypeDeleteFromPersonSecurityPolicyWebsitesBlacklistParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypeDeleteFromPersonSecurityPolicyWebsitesBlacklistParams contains all the parameters to send to the API endpoint

	for the customer prototype delete from person security policy websites blacklist operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypeDeleteFromPersonSecurityPolicyWebsitesBlacklistParams struct {

	// DNS.
	DNS string

	/* ID.

	   Customer id

	   Format: JSON
	*/
	ID string

	// LocationID.
	LocationID string

	// PersonID.
	PersonID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer prototype delete from person security policy websites blacklist params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeDeleteFromPersonSecurityPolicyWebsitesBlacklistParams) WithDefaults() *CustomerPrototypeDeleteFromPersonSecurityPolicyWebsitesBlacklistParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype delete from person security policy websites blacklist params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeDeleteFromPersonSecurityPolicyWebsitesBlacklistParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype delete from person security policy websites blacklist params
func (o *CustomerPrototypeDeleteFromPersonSecurityPolicyWebsitesBlacklistParams) WithTimeout(timeout time.Duration) *CustomerPrototypeDeleteFromPersonSecurityPolicyWebsitesBlacklistParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype delete from person security policy websites blacklist params
func (o *CustomerPrototypeDeleteFromPersonSecurityPolicyWebsitesBlacklistParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype delete from person security policy websites blacklist params
func (o *CustomerPrototypeDeleteFromPersonSecurityPolicyWebsitesBlacklistParams) WithContext(ctx context.Context) *CustomerPrototypeDeleteFromPersonSecurityPolicyWebsitesBlacklistParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype delete from person security policy websites blacklist params
func (o *CustomerPrototypeDeleteFromPersonSecurityPolicyWebsitesBlacklistParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype delete from person security policy websites blacklist params
func (o *CustomerPrototypeDeleteFromPersonSecurityPolicyWebsitesBlacklistParams) WithHTTPClient(client *http.Client) *CustomerPrototypeDeleteFromPersonSecurityPolicyWebsitesBlacklistParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype delete from person security policy websites blacklist params
func (o *CustomerPrototypeDeleteFromPersonSecurityPolicyWebsitesBlacklistParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDNS adds the dns to the customer prototype delete from person security policy websites blacklist params
func (o *CustomerPrototypeDeleteFromPersonSecurityPolicyWebsitesBlacklistParams) WithDNS(dns string) *CustomerPrototypeDeleteFromPersonSecurityPolicyWebsitesBlacklistParams {
	o.SetDNS(dns)
	return o
}

// SetDNS adds the dns to the customer prototype delete from person security policy websites blacklist params
func (o *CustomerPrototypeDeleteFromPersonSecurityPolicyWebsitesBlacklistParams) SetDNS(dns string) {
	o.DNS = dns
}

// WithID adds the id to the customer prototype delete from person security policy websites blacklist params
func (o *CustomerPrototypeDeleteFromPersonSecurityPolicyWebsitesBlacklistParams) WithID(id string) *CustomerPrototypeDeleteFromPersonSecurityPolicyWebsitesBlacklistParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype delete from person security policy websites blacklist params
func (o *CustomerPrototypeDeleteFromPersonSecurityPolicyWebsitesBlacklistParams) SetID(id string) {
	o.ID = id
}

// WithLocationID adds the locationID to the customer prototype delete from person security policy websites blacklist params
func (o *CustomerPrototypeDeleteFromPersonSecurityPolicyWebsitesBlacklistParams) WithLocationID(locationID string) *CustomerPrototypeDeleteFromPersonSecurityPolicyWebsitesBlacklistParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype delete from person security policy websites blacklist params
func (o *CustomerPrototypeDeleteFromPersonSecurityPolicyWebsitesBlacklistParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WithPersonID adds the personID to the customer prototype delete from person security policy websites blacklist params
func (o *CustomerPrototypeDeleteFromPersonSecurityPolicyWebsitesBlacklistParams) WithPersonID(personID string) *CustomerPrototypeDeleteFromPersonSecurityPolicyWebsitesBlacklistParams {
	o.SetPersonID(personID)
	return o
}

// SetPersonID adds the personId to the customer prototype delete from person security policy websites blacklist params
func (o *CustomerPrototypeDeleteFromPersonSecurityPolicyWebsitesBlacklistParams) SetPersonID(personID string) {
	o.PersonID = personID
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypeDeleteFromPersonSecurityPolicyWebsitesBlacklistParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param personId
	if err := r.SetPathParam("personId", o.PersonID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
