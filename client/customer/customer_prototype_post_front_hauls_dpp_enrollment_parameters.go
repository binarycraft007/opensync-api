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

// NewCustomerPrototypePostFrontHaulsDppEnrollmentParams creates a new CustomerPrototypePostFrontHaulsDppEnrollmentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypePostFrontHaulsDppEnrollmentParams() *CustomerPrototypePostFrontHaulsDppEnrollmentParams {
	return &CustomerPrototypePostFrontHaulsDppEnrollmentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypePostFrontHaulsDppEnrollmentParamsWithTimeout creates a new CustomerPrototypePostFrontHaulsDppEnrollmentParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypePostFrontHaulsDppEnrollmentParamsWithTimeout(timeout time.Duration) *CustomerPrototypePostFrontHaulsDppEnrollmentParams {
	return &CustomerPrototypePostFrontHaulsDppEnrollmentParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypePostFrontHaulsDppEnrollmentParamsWithContext creates a new CustomerPrototypePostFrontHaulsDppEnrollmentParams object
// with the ability to set a context for a request.
func NewCustomerPrototypePostFrontHaulsDppEnrollmentParamsWithContext(ctx context.Context) *CustomerPrototypePostFrontHaulsDppEnrollmentParams {
	return &CustomerPrototypePostFrontHaulsDppEnrollmentParams{
		Context: ctx,
	}
}

// NewCustomerPrototypePostFrontHaulsDppEnrollmentParamsWithHTTPClient creates a new CustomerPrototypePostFrontHaulsDppEnrollmentParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypePostFrontHaulsDppEnrollmentParamsWithHTTPClient(client *http.Client) *CustomerPrototypePostFrontHaulsDppEnrollmentParams {
	return &CustomerPrototypePostFrontHaulsDppEnrollmentParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypePostFrontHaulsDppEnrollmentParams contains all the parameters to send to the API endpoint

	for the customer prototype post front hauls dpp enrollment operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypePostFrontHaulsDppEnrollmentParams struct {

	// BootstrapURI.
	BootstrapURI string

	/* ID.

	   Customer id

	   Format: JSON
	*/
	ID string

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

// WithDefaults hydrates default values in the customer prototype post front hauls dpp enrollment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePostFrontHaulsDppEnrollmentParams) WithDefaults() *CustomerPrototypePostFrontHaulsDppEnrollmentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype post front hauls dpp enrollment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePostFrontHaulsDppEnrollmentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype post front hauls dpp enrollment params
func (o *CustomerPrototypePostFrontHaulsDppEnrollmentParams) WithTimeout(timeout time.Duration) *CustomerPrototypePostFrontHaulsDppEnrollmentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype post front hauls dpp enrollment params
func (o *CustomerPrototypePostFrontHaulsDppEnrollmentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype post front hauls dpp enrollment params
func (o *CustomerPrototypePostFrontHaulsDppEnrollmentParams) WithContext(ctx context.Context) *CustomerPrototypePostFrontHaulsDppEnrollmentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype post front hauls dpp enrollment params
func (o *CustomerPrototypePostFrontHaulsDppEnrollmentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype post front hauls dpp enrollment params
func (o *CustomerPrototypePostFrontHaulsDppEnrollmentParams) WithHTTPClient(client *http.Client) *CustomerPrototypePostFrontHaulsDppEnrollmentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype post front hauls dpp enrollment params
func (o *CustomerPrototypePostFrontHaulsDppEnrollmentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBootstrapURI adds the bootstrapURI to the customer prototype post front hauls dpp enrollment params
func (o *CustomerPrototypePostFrontHaulsDppEnrollmentParams) WithBootstrapURI(bootstrapURI string) *CustomerPrototypePostFrontHaulsDppEnrollmentParams {
	o.SetBootstrapURI(bootstrapURI)
	return o
}

// SetBootstrapURI adds the bootstrapUri to the customer prototype post front hauls dpp enrollment params
func (o *CustomerPrototypePostFrontHaulsDppEnrollmentParams) SetBootstrapURI(bootstrapURI string) {
	o.BootstrapURI = bootstrapURI
}

// WithID adds the id to the customer prototype post front hauls dpp enrollment params
func (o *CustomerPrototypePostFrontHaulsDppEnrollmentParams) WithID(id string) *CustomerPrototypePostFrontHaulsDppEnrollmentParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype post front hauls dpp enrollment params
func (o *CustomerPrototypePostFrontHaulsDppEnrollmentParams) SetID(id string) {
	o.ID = id
}

// WithLocationID adds the locationID to the customer prototype post front hauls dpp enrollment params
func (o *CustomerPrototypePostFrontHaulsDppEnrollmentParams) WithLocationID(locationID string) *CustomerPrototypePostFrontHaulsDppEnrollmentParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype post front hauls dpp enrollment params
func (o *CustomerPrototypePostFrontHaulsDppEnrollmentParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WithNetworkID adds the networkID to the customer prototype post front hauls dpp enrollment params
func (o *CustomerPrototypePostFrontHaulsDppEnrollmentParams) WithNetworkID(networkID string) *CustomerPrototypePostFrontHaulsDppEnrollmentParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the customer prototype post front hauls dpp enrollment params
func (o *CustomerPrototypePostFrontHaulsDppEnrollmentParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypePostFrontHaulsDppEnrollmentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// form param bootstrapUri
	frBootstrapURI := o.BootstrapURI
	fBootstrapURI := frBootstrapURI
	if fBootstrapURI != "" {
		if err := r.SetFormParam("bootstrapUri", fBootstrapURI); err != nil {
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

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
