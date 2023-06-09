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

// NewCustomerPrototypeGetGdprCaptivePortalsDataParams creates a new CustomerPrototypeGetGdprCaptivePortalsDataParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypeGetGdprCaptivePortalsDataParams() *CustomerPrototypeGetGdprCaptivePortalsDataParams {
	return &CustomerPrototypeGetGdprCaptivePortalsDataParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypeGetGdprCaptivePortalsDataParamsWithTimeout creates a new CustomerPrototypeGetGdprCaptivePortalsDataParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypeGetGdprCaptivePortalsDataParamsWithTimeout(timeout time.Duration) *CustomerPrototypeGetGdprCaptivePortalsDataParams {
	return &CustomerPrototypeGetGdprCaptivePortalsDataParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypeGetGdprCaptivePortalsDataParamsWithContext creates a new CustomerPrototypeGetGdprCaptivePortalsDataParams object
// with the ability to set a context for a request.
func NewCustomerPrototypeGetGdprCaptivePortalsDataParamsWithContext(ctx context.Context) *CustomerPrototypeGetGdprCaptivePortalsDataParams {
	return &CustomerPrototypeGetGdprCaptivePortalsDataParams{
		Context: ctx,
	}
}

// NewCustomerPrototypeGetGdprCaptivePortalsDataParamsWithHTTPClient creates a new CustomerPrototypeGetGdprCaptivePortalsDataParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypeGetGdprCaptivePortalsDataParamsWithHTTPClient(client *http.Client) *CustomerPrototypeGetGdprCaptivePortalsDataParams {
	return &CustomerPrototypeGetGdprCaptivePortalsDataParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypeGetGdprCaptivePortalsDataParams contains all the parameters to send to the API endpoint

	for the customer prototype get gdpr captive portals data operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypeGetGdprCaptivePortalsDataParams struct {

	// Email.
	Email string

	/* ID.

	   Customer id

	   Format: JSON
	*/
	ID string

	// LocalEndDate.
	LocalEndDate string

	// LocalStartDate.
	LocalStartDate string

	// LocationID.
	//
	// Format: JSON
	LocationID string

	// NetworkID.
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer prototype get gdpr captive portals data params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeGetGdprCaptivePortalsDataParams) WithDefaults() *CustomerPrototypeGetGdprCaptivePortalsDataParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype get gdpr captive portals data params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeGetGdprCaptivePortalsDataParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype get gdpr captive portals data params
func (o *CustomerPrototypeGetGdprCaptivePortalsDataParams) WithTimeout(timeout time.Duration) *CustomerPrototypeGetGdprCaptivePortalsDataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype get gdpr captive portals data params
func (o *CustomerPrototypeGetGdprCaptivePortalsDataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype get gdpr captive portals data params
func (o *CustomerPrototypeGetGdprCaptivePortalsDataParams) WithContext(ctx context.Context) *CustomerPrototypeGetGdprCaptivePortalsDataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype get gdpr captive portals data params
func (o *CustomerPrototypeGetGdprCaptivePortalsDataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype get gdpr captive portals data params
func (o *CustomerPrototypeGetGdprCaptivePortalsDataParams) WithHTTPClient(client *http.Client) *CustomerPrototypeGetGdprCaptivePortalsDataParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype get gdpr captive portals data params
func (o *CustomerPrototypeGetGdprCaptivePortalsDataParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEmail adds the email to the customer prototype get gdpr captive portals data params
func (o *CustomerPrototypeGetGdprCaptivePortalsDataParams) WithEmail(email string) *CustomerPrototypeGetGdprCaptivePortalsDataParams {
	o.SetEmail(email)
	return o
}

// SetEmail adds the email to the customer prototype get gdpr captive portals data params
func (o *CustomerPrototypeGetGdprCaptivePortalsDataParams) SetEmail(email string) {
	o.Email = email
}

// WithID adds the id to the customer prototype get gdpr captive portals data params
func (o *CustomerPrototypeGetGdprCaptivePortalsDataParams) WithID(id string) *CustomerPrototypeGetGdprCaptivePortalsDataParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype get gdpr captive portals data params
func (o *CustomerPrototypeGetGdprCaptivePortalsDataParams) SetID(id string) {
	o.ID = id
}

// WithLocalEndDate adds the localEndDate to the customer prototype get gdpr captive portals data params
func (o *CustomerPrototypeGetGdprCaptivePortalsDataParams) WithLocalEndDate(localEndDate string) *CustomerPrototypeGetGdprCaptivePortalsDataParams {
	o.SetLocalEndDate(localEndDate)
	return o
}

// SetLocalEndDate adds the localEndDate to the customer prototype get gdpr captive portals data params
func (o *CustomerPrototypeGetGdprCaptivePortalsDataParams) SetLocalEndDate(localEndDate string) {
	o.LocalEndDate = localEndDate
}

// WithLocalStartDate adds the localStartDate to the customer prototype get gdpr captive portals data params
func (o *CustomerPrototypeGetGdprCaptivePortalsDataParams) WithLocalStartDate(localStartDate string) *CustomerPrototypeGetGdprCaptivePortalsDataParams {
	o.SetLocalStartDate(localStartDate)
	return o
}

// SetLocalStartDate adds the localStartDate to the customer prototype get gdpr captive portals data params
func (o *CustomerPrototypeGetGdprCaptivePortalsDataParams) SetLocalStartDate(localStartDate string) {
	o.LocalStartDate = localStartDate
}

// WithLocationID adds the locationID to the customer prototype get gdpr captive portals data params
func (o *CustomerPrototypeGetGdprCaptivePortalsDataParams) WithLocationID(locationID string) *CustomerPrototypeGetGdprCaptivePortalsDataParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype get gdpr captive portals data params
func (o *CustomerPrototypeGetGdprCaptivePortalsDataParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WithNetworkID adds the networkID to the customer prototype get gdpr captive portals data params
func (o *CustomerPrototypeGetGdprCaptivePortalsDataParams) WithNetworkID(networkID string) *CustomerPrototypeGetGdprCaptivePortalsDataParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the customer prototype get gdpr captive portals data params
func (o *CustomerPrototypeGetGdprCaptivePortalsDataParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypeGetGdprCaptivePortalsDataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param email
	qrEmail := o.Email
	qEmail := qrEmail
	if qEmail != "" {

		if err := r.SetQueryParam("email", qEmail); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// query param localEndDate
	qrLocalEndDate := o.LocalEndDate
	qLocalEndDate := qrLocalEndDate
	if qLocalEndDate != "" {

		if err := r.SetQueryParam("localEndDate", qLocalEndDate); err != nil {
			return err
		}
	}

	// query param localStartDate
	qrLocalStartDate := o.LocalStartDate
	qLocalStartDate := qrLocalStartDate
	if qLocalStartDate != "" {

		if err := r.SetQueryParam("localStartDate", qLocalStartDate); err != nil {
			return err
		}
	}

	// path param locationId
	if err := r.SetPathParam("locationId", o.LocationID); err != nil {
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
