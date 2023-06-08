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
	"github.com/go-openapi/swag"
)

// NewCustomerPrototypePostWifiNetworkDppParams creates a new CustomerPrototypePostWifiNetworkDppParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypePostWifiNetworkDppParams() *CustomerPrototypePostWifiNetworkDppParams {
	return &CustomerPrototypePostWifiNetworkDppParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypePostWifiNetworkDppParamsWithTimeout creates a new CustomerPrototypePostWifiNetworkDppParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypePostWifiNetworkDppParamsWithTimeout(timeout time.Duration) *CustomerPrototypePostWifiNetworkDppParams {
	return &CustomerPrototypePostWifiNetworkDppParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypePostWifiNetworkDppParamsWithContext creates a new CustomerPrototypePostWifiNetworkDppParams object
// with the ability to set a context for a request.
func NewCustomerPrototypePostWifiNetworkDppParamsWithContext(ctx context.Context) *CustomerPrototypePostWifiNetworkDppParams {
	return &CustomerPrototypePostWifiNetworkDppParams{
		Context: ctx,
	}
}

// NewCustomerPrototypePostWifiNetworkDppParamsWithHTTPClient creates a new CustomerPrototypePostWifiNetworkDppParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypePostWifiNetworkDppParamsWithHTTPClient(client *http.Client) *CustomerPrototypePostWifiNetworkDppParams {
	return &CustomerPrototypePostWifiNetworkDppParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypePostWifiNetworkDppParams contains all the parameters to send to the API endpoint

	for the customer prototype post wifi network dpp operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypePostWifiNetworkDppParams struct {

	/* Curve.

	   one of predefined elliptic curves, - optional,  if missing in request default to prime256v1
	*/
	Curve *string

	/* Enabled.

	   should we configure dpp for this network - defaults to true
	*/
	Enabled *bool

	/* ID.

	   Customer id

	   Format: JSON
	*/
	ID string

	// LocationID.
	//
	// Format: JSON
	LocationID string

	/* PrivateKey.

	   privateKey, must also provide public part if present, optional
	*/
	PrivateKey *string

	/* PublicKey.

	   publicKey
	*/
	PublicKey *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer prototype post wifi network dpp params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePostWifiNetworkDppParams) WithDefaults() *CustomerPrototypePostWifiNetworkDppParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype post wifi network dpp params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePostWifiNetworkDppParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype post wifi network dpp params
func (o *CustomerPrototypePostWifiNetworkDppParams) WithTimeout(timeout time.Duration) *CustomerPrototypePostWifiNetworkDppParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype post wifi network dpp params
func (o *CustomerPrototypePostWifiNetworkDppParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype post wifi network dpp params
func (o *CustomerPrototypePostWifiNetworkDppParams) WithContext(ctx context.Context) *CustomerPrototypePostWifiNetworkDppParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype post wifi network dpp params
func (o *CustomerPrototypePostWifiNetworkDppParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype post wifi network dpp params
func (o *CustomerPrototypePostWifiNetworkDppParams) WithHTTPClient(client *http.Client) *CustomerPrototypePostWifiNetworkDppParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype post wifi network dpp params
func (o *CustomerPrototypePostWifiNetworkDppParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCurve adds the curve to the customer prototype post wifi network dpp params
func (o *CustomerPrototypePostWifiNetworkDppParams) WithCurve(curve *string) *CustomerPrototypePostWifiNetworkDppParams {
	o.SetCurve(curve)
	return o
}

// SetCurve adds the curve to the customer prototype post wifi network dpp params
func (o *CustomerPrototypePostWifiNetworkDppParams) SetCurve(curve *string) {
	o.Curve = curve
}

// WithEnabled adds the enabled to the customer prototype post wifi network dpp params
func (o *CustomerPrototypePostWifiNetworkDppParams) WithEnabled(enabled *bool) *CustomerPrototypePostWifiNetworkDppParams {
	o.SetEnabled(enabled)
	return o
}

// SetEnabled adds the enabled to the customer prototype post wifi network dpp params
func (o *CustomerPrototypePostWifiNetworkDppParams) SetEnabled(enabled *bool) {
	o.Enabled = enabled
}

// WithID adds the id to the customer prototype post wifi network dpp params
func (o *CustomerPrototypePostWifiNetworkDppParams) WithID(id string) *CustomerPrototypePostWifiNetworkDppParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype post wifi network dpp params
func (o *CustomerPrototypePostWifiNetworkDppParams) SetID(id string) {
	o.ID = id
}

// WithLocationID adds the locationID to the customer prototype post wifi network dpp params
func (o *CustomerPrototypePostWifiNetworkDppParams) WithLocationID(locationID string) *CustomerPrototypePostWifiNetworkDppParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype post wifi network dpp params
func (o *CustomerPrototypePostWifiNetworkDppParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WithPrivateKey adds the privateKey to the customer prototype post wifi network dpp params
func (o *CustomerPrototypePostWifiNetworkDppParams) WithPrivateKey(privateKey *string) *CustomerPrototypePostWifiNetworkDppParams {
	o.SetPrivateKey(privateKey)
	return o
}

// SetPrivateKey adds the privateKey to the customer prototype post wifi network dpp params
func (o *CustomerPrototypePostWifiNetworkDppParams) SetPrivateKey(privateKey *string) {
	o.PrivateKey = privateKey
}

// WithPublicKey adds the publicKey to the customer prototype post wifi network dpp params
func (o *CustomerPrototypePostWifiNetworkDppParams) WithPublicKey(publicKey *string) *CustomerPrototypePostWifiNetworkDppParams {
	o.SetPublicKey(publicKey)
	return o
}

// SetPublicKey adds the publicKey to the customer prototype post wifi network dpp params
func (o *CustomerPrototypePostWifiNetworkDppParams) SetPublicKey(publicKey *string) {
	o.PublicKey = publicKey
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypePostWifiNetworkDppParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Curve != nil {

		// form param curve
		var frCurve string
		if o.Curve != nil {
			frCurve = *o.Curve
		}
		fCurve := frCurve
		if fCurve != "" {
			if err := r.SetFormParam("curve", fCurve); err != nil {
				return err
			}
		}
	}

	if o.Enabled != nil {

		// form param enabled
		var frEnabled bool
		if o.Enabled != nil {
			frEnabled = *o.Enabled
		}
		fEnabled := swag.FormatBool(frEnabled)
		if fEnabled != "" {
			if err := r.SetFormParam("enabled", fEnabled); err != nil {
				return err
			}
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

	if o.PrivateKey != nil {

		// form param privateKey
		var frPrivateKey string
		if o.PrivateKey != nil {
			frPrivateKey = *o.PrivateKey
		}
		fPrivateKey := frPrivateKey
		if fPrivateKey != "" {
			if err := r.SetFormParam("privateKey", fPrivateKey); err != nil {
				return err
			}
		}
	}

	if o.PublicKey != nil {

		// form param publicKey
		var frPublicKey string
		if o.PublicKey != nil {
			frPublicKey = *o.PublicKey
		}
		fPublicKey := frPublicKey
		if fPublicKey != "" {
			if err := r.SetFormParam("publicKey", fPublicKey); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
