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

// NewCustomerPrototypePostCaptivePortalParams creates a new CustomerPrototypePostCaptivePortalParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypePostCaptivePortalParams() *CustomerPrototypePostCaptivePortalParams {
	return &CustomerPrototypePostCaptivePortalParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypePostCaptivePortalParamsWithTimeout creates a new CustomerPrototypePostCaptivePortalParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypePostCaptivePortalParamsWithTimeout(timeout time.Duration) *CustomerPrototypePostCaptivePortalParams {
	return &CustomerPrototypePostCaptivePortalParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypePostCaptivePortalParamsWithContext creates a new CustomerPrototypePostCaptivePortalParams object
// with the ability to set a context for a request.
func NewCustomerPrototypePostCaptivePortalParamsWithContext(ctx context.Context) *CustomerPrototypePostCaptivePortalParams {
	return &CustomerPrototypePostCaptivePortalParams{
		Context: ctx,
	}
}

// NewCustomerPrototypePostCaptivePortalParamsWithHTTPClient creates a new CustomerPrototypePostCaptivePortalParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypePostCaptivePortalParamsWithHTTPClient(client *http.Client) *CustomerPrototypePostCaptivePortalParams {
	return &CustomerPrototypePostCaptivePortalParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypePostCaptivePortalParams contains all the parameters to send to the API endpoint

	for the customer prototype post captive portal operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypePostCaptivePortalParams struct {

	/* BandwidthLimit.

	   attributes: "enabled" boolean, "type": "absolute"|"percentage", "upload"/"download" - either as percentage or absolute (Mbps)

	   Format: JSON
	*/
	BandwidthLimit *string

	// Enable.
	//
	// Default: true
	Enable *bool

	// EncryptionKey.
	EncryptionKey *string

	/* ID.

	   Customer id

	   Format: JSON
	*/
	ID string

	// Language.
	Language *string

	// LocationID.
	LocationID string

	// NetworkID.
	NetworkID *string

	// SessionTimeLimitSec.
	//
	// Format: double
	SessionTimeLimitSec *float64

	// Ssid.
	Ssid string

	// WpaMode.
	WpaMode *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer prototype post captive portal params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePostCaptivePortalParams) WithDefaults() *CustomerPrototypePostCaptivePortalParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype post captive portal params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePostCaptivePortalParams) SetDefaults() {
	var (
		enableDefault = bool(true)
	)

	val := CustomerPrototypePostCaptivePortalParams{
		Enable: &enableDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the customer prototype post captive portal params
func (o *CustomerPrototypePostCaptivePortalParams) WithTimeout(timeout time.Duration) *CustomerPrototypePostCaptivePortalParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype post captive portal params
func (o *CustomerPrototypePostCaptivePortalParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype post captive portal params
func (o *CustomerPrototypePostCaptivePortalParams) WithContext(ctx context.Context) *CustomerPrototypePostCaptivePortalParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype post captive portal params
func (o *CustomerPrototypePostCaptivePortalParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype post captive portal params
func (o *CustomerPrototypePostCaptivePortalParams) WithHTTPClient(client *http.Client) *CustomerPrototypePostCaptivePortalParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype post captive portal params
func (o *CustomerPrototypePostCaptivePortalParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBandwidthLimit adds the bandwidthLimit to the customer prototype post captive portal params
func (o *CustomerPrototypePostCaptivePortalParams) WithBandwidthLimit(bandwidthLimit *string) *CustomerPrototypePostCaptivePortalParams {
	o.SetBandwidthLimit(bandwidthLimit)
	return o
}

// SetBandwidthLimit adds the bandwidthLimit to the customer prototype post captive portal params
func (o *CustomerPrototypePostCaptivePortalParams) SetBandwidthLimit(bandwidthLimit *string) {
	o.BandwidthLimit = bandwidthLimit
}

// WithEnable adds the enable to the customer prototype post captive portal params
func (o *CustomerPrototypePostCaptivePortalParams) WithEnable(enable *bool) *CustomerPrototypePostCaptivePortalParams {
	o.SetEnable(enable)
	return o
}

// SetEnable adds the enable to the customer prototype post captive portal params
func (o *CustomerPrototypePostCaptivePortalParams) SetEnable(enable *bool) {
	o.Enable = enable
}

// WithEncryptionKey adds the encryptionKey to the customer prototype post captive portal params
func (o *CustomerPrototypePostCaptivePortalParams) WithEncryptionKey(encryptionKey *string) *CustomerPrototypePostCaptivePortalParams {
	o.SetEncryptionKey(encryptionKey)
	return o
}

// SetEncryptionKey adds the encryptionKey to the customer prototype post captive portal params
func (o *CustomerPrototypePostCaptivePortalParams) SetEncryptionKey(encryptionKey *string) {
	o.EncryptionKey = encryptionKey
}

// WithID adds the id to the customer prototype post captive portal params
func (o *CustomerPrototypePostCaptivePortalParams) WithID(id string) *CustomerPrototypePostCaptivePortalParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype post captive portal params
func (o *CustomerPrototypePostCaptivePortalParams) SetID(id string) {
	o.ID = id
}

// WithLanguage adds the language to the customer prototype post captive portal params
func (o *CustomerPrototypePostCaptivePortalParams) WithLanguage(language *string) *CustomerPrototypePostCaptivePortalParams {
	o.SetLanguage(language)
	return o
}

// SetLanguage adds the language to the customer prototype post captive portal params
func (o *CustomerPrototypePostCaptivePortalParams) SetLanguage(language *string) {
	o.Language = language
}

// WithLocationID adds the locationID to the customer prototype post captive portal params
func (o *CustomerPrototypePostCaptivePortalParams) WithLocationID(locationID string) *CustomerPrototypePostCaptivePortalParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype post captive portal params
func (o *CustomerPrototypePostCaptivePortalParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WithNetworkID adds the networkID to the customer prototype post captive portal params
func (o *CustomerPrototypePostCaptivePortalParams) WithNetworkID(networkID *string) *CustomerPrototypePostCaptivePortalParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the customer prototype post captive portal params
func (o *CustomerPrototypePostCaptivePortalParams) SetNetworkID(networkID *string) {
	o.NetworkID = networkID
}

// WithSessionTimeLimitSec adds the sessionTimeLimitSec to the customer prototype post captive portal params
func (o *CustomerPrototypePostCaptivePortalParams) WithSessionTimeLimitSec(sessionTimeLimitSec *float64) *CustomerPrototypePostCaptivePortalParams {
	o.SetSessionTimeLimitSec(sessionTimeLimitSec)
	return o
}

// SetSessionTimeLimitSec adds the sessionTimeLimitSec to the customer prototype post captive portal params
func (o *CustomerPrototypePostCaptivePortalParams) SetSessionTimeLimitSec(sessionTimeLimitSec *float64) {
	o.SessionTimeLimitSec = sessionTimeLimitSec
}

// WithSsid adds the ssid to the customer prototype post captive portal params
func (o *CustomerPrototypePostCaptivePortalParams) WithSsid(ssid string) *CustomerPrototypePostCaptivePortalParams {
	o.SetSsid(ssid)
	return o
}

// SetSsid adds the ssid to the customer prototype post captive portal params
func (o *CustomerPrototypePostCaptivePortalParams) SetSsid(ssid string) {
	o.Ssid = ssid
}

// WithWpaMode adds the wpaMode to the customer prototype post captive portal params
func (o *CustomerPrototypePostCaptivePortalParams) WithWpaMode(wpaMode *string) *CustomerPrototypePostCaptivePortalParams {
	o.SetWpaMode(wpaMode)
	return o
}

// SetWpaMode adds the wpaMode to the customer prototype post captive portal params
func (o *CustomerPrototypePostCaptivePortalParams) SetWpaMode(wpaMode *string) {
	o.WpaMode = wpaMode
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypePostCaptivePortalParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.BandwidthLimit != nil {

		// form param bandwidthLimit
		var frBandwidthLimit string
		if o.BandwidthLimit != nil {
			frBandwidthLimit = *o.BandwidthLimit
		}
		fBandwidthLimit := frBandwidthLimit
		if fBandwidthLimit != "" {
			if err := r.SetFormParam("bandwidthLimit", fBandwidthLimit); err != nil {
				return err
			}
		}
	}

	if o.Enable != nil {

		// form param enable
		var frEnable bool
		if o.Enable != nil {
			frEnable = *o.Enable
		}
		fEnable := swag.FormatBool(frEnable)
		if fEnable != "" {
			if err := r.SetFormParam("enable", fEnable); err != nil {
				return err
			}
		}
	}

	if o.EncryptionKey != nil {

		// form param encryptionKey
		var frEncryptionKey string
		if o.EncryptionKey != nil {
			frEncryptionKey = *o.EncryptionKey
		}
		fEncryptionKey := frEncryptionKey
		if fEncryptionKey != "" {
			if err := r.SetFormParam("encryptionKey", fEncryptionKey); err != nil {
				return err
			}
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.Language != nil {

		// form param language
		var frLanguage string
		if o.Language != nil {
			frLanguage = *o.Language
		}
		fLanguage := frLanguage
		if fLanguage != "" {
			if err := r.SetFormParam("language", fLanguage); err != nil {
				return err
			}
		}
	}

	// path param locationId
	if err := r.SetPathParam("locationId", o.LocationID); err != nil {
		return err
	}

	if o.NetworkID != nil {

		// form param networkId
		var frNetworkID string
		if o.NetworkID != nil {
			frNetworkID = *o.NetworkID
		}
		fNetworkID := frNetworkID
		if fNetworkID != "" {
			if err := r.SetFormParam("networkId", fNetworkID); err != nil {
				return err
			}
		}
	}

	if o.SessionTimeLimitSec != nil {

		// form param sessionTimeLimitSec
		var frSessionTimeLimitSec float64
		if o.SessionTimeLimitSec != nil {
			frSessionTimeLimitSec = *o.SessionTimeLimitSec
		}
		fSessionTimeLimitSec := swag.FormatFloat64(frSessionTimeLimitSec)
		if fSessionTimeLimitSec != "" {
			if err := r.SetFormParam("sessionTimeLimitSec", fSessionTimeLimitSec); err != nil {
				return err
			}
		}
	}

	// form param ssid
	frSsid := o.Ssid
	fSsid := frSsid
	if fSsid != "" {
		if err := r.SetFormParam("ssid", fSsid); err != nil {
			return err
		}
	}

	if o.WpaMode != nil {

		// form param wpaMode
		var frWpaMode string
		if o.WpaMode != nil {
			frWpaMode = *o.WpaMode
		}
		fWpaMode := frWpaMode
		if fWpaMode != "" {
			if err := r.SetFormParam("wpaMode", fWpaMode); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
