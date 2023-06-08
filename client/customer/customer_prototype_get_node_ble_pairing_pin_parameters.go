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

// NewCustomerPrototypeGetNodeBlePairingPinParams creates a new CustomerPrototypeGetNodeBlePairingPinParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypeGetNodeBlePairingPinParams() *CustomerPrototypeGetNodeBlePairingPinParams {
	return &CustomerPrototypeGetNodeBlePairingPinParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypeGetNodeBlePairingPinParamsWithTimeout creates a new CustomerPrototypeGetNodeBlePairingPinParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypeGetNodeBlePairingPinParamsWithTimeout(timeout time.Duration) *CustomerPrototypeGetNodeBlePairingPinParams {
	return &CustomerPrototypeGetNodeBlePairingPinParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypeGetNodeBlePairingPinParamsWithContext creates a new CustomerPrototypeGetNodeBlePairingPinParams object
// with the ability to set a context for a request.
func NewCustomerPrototypeGetNodeBlePairingPinParamsWithContext(ctx context.Context) *CustomerPrototypeGetNodeBlePairingPinParams {
	return &CustomerPrototypeGetNodeBlePairingPinParams{
		Context: ctx,
	}
}

// NewCustomerPrototypeGetNodeBlePairingPinParamsWithHTTPClient creates a new CustomerPrototypeGetNodeBlePairingPinParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypeGetNodeBlePairingPinParamsWithHTTPClient(client *http.Client) *CustomerPrototypeGetNodeBlePairingPinParams {
	return &CustomerPrototypeGetNodeBlePairingPinParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypeGetNodeBlePairingPinParams contains all the parameters to send to the API endpoint

	for the customer prototype get node ble pairing pin operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypeGetNodeBlePairingPinParams struct {

	/* ID.

	   Customer id

	   Format: JSON
	*/
	ID string

	// LocationID.
	LocationID string

	// NodeID.
	NodeID string

	// Token.
	Token string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer prototype get node ble pairing pin params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeGetNodeBlePairingPinParams) WithDefaults() *CustomerPrototypeGetNodeBlePairingPinParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype get node ble pairing pin params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeGetNodeBlePairingPinParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype get node ble pairing pin params
func (o *CustomerPrototypeGetNodeBlePairingPinParams) WithTimeout(timeout time.Duration) *CustomerPrototypeGetNodeBlePairingPinParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype get node ble pairing pin params
func (o *CustomerPrototypeGetNodeBlePairingPinParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype get node ble pairing pin params
func (o *CustomerPrototypeGetNodeBlePairingPinParams) WithContext(ctx context.Context) *CustomerPrototypeGetNodeBlePairingPinParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype get node ble pairing pin params
func (o *CustomerPrototypeGetNodeBlePairingPinParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype get node ble pairing pin params
func (o *CustomerPrototypeGetNodeBlePairingPinParams) WithHTTPClient(client *http.Client) *CustomerPrototypeGetNodeBlePairingPinParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype get node ble pairing pin params
func (o *CustomerPrototypeGetNodeBlePairingPinParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the customer prototype get node ble pairing pin params
func (o *CustomerPrototypeGetNodeBlePairingPinParams) WithID(id string) *CustomerPrototypeGetNodeBlePairingPinParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype get node ble pairing pin params
func (o *CustomerPrototypeGetNodeBlePairingPinParams) SetID(id string) {
	o.ID = id
}

// WithLocationID adds the locationID to the customer prototype get node ble pairing pin params
func (o *CustomerPrototypeGetNodeBlePairingPinParams) WithLocationID(locationID string) *CustomerPrototypeGetNodeBlePairingPinParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype get node ble pairing pin params
func (o *CustomerPrototypeGetNodeBlePairingPinParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WithNodeID adds the nodeID to the customer prototype get node ble pairing pin params
func (o *CustomerPrototypeGetNodeBlePairingPinParams) WithNodeID(nodeID string) *CustomerPrototypeGetNodeBlePairingPinParams {
	o.SetNodeID(nodeID)
	return o
}

// SetNodeID adds the nodeId to the customer prototype get node ble pairing pin params
func (o *CustomerPrototypeGetNodeBlePairingPinParams) SetNodeID(nodeID string) {
	o.NodeID = nodeID
}

// WithToken adds the token to the customer prototype get node ble pairing pin params
func (o *CustomerPrototypeGetNodeBlePairingPinParams) WithToken(token string) *CustomerPrototypeGetNodeBlePairingPinParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the customer prototype get node ble pairing pin params
func (o *CustomerPrototypeGetNodeBlePairingPinParams) SetToken(token string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypeGetNodeBlePairingPinParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param nodeId
	if err := r.SetPathParam("nodeId", o.NodeID); err != nil {
		return err
	}

	// query param token
	qrToken := o.Token
	qToken := qrToken
	if qToken != "" {

		if err := r.SetQueryParam("token", qToken); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}