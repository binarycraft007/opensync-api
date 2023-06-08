// Code generated by go-swagger; DO NOT EDIT.

package location

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

// NewLocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistParams creates a new LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistParams() *LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistParams {
	return &LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistParamsWithTimeout creates a new LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistParams object
// with the ability to set a timeout on a request.
func NewLocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistParamsWithTimeout(timeout time.Duration) *LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistParams {
	return &LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistParams{
		timeout: timeout,
	}
}

// NewLocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistParamsWithContext creates a new LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistParams object
// with the ability to set a context for a request.
func NewLocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistParamsWithContext(ctx context.Context) *LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistParams {
	return &LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistParams{
		Context: ctx,
	}
}

// NewLocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistParamsWithHTTPClient creates a new LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistParams object
// with the ability to set a custom HTTPClient for a request.
func NewLocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistParamsWithHTTPClient(client *http.Client) *LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistParams {
	return &LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistParams{
		HTTPClient: client,
	}
}

/*
LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistParams contains all the parameters to send to the API endpoint

	for the location prototype post device security policy anomaly blacklist operation.

	Typically these are written to a http.Request.
*/
type LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistParams struct {

	// DeviceType.
	DeviceType string

	// Fqdn.
	Fqdn string

	/* ID.

	   Location id

	   Format: JSON
	*/
	ID string

	// IPV4.
	//
	// Format: JSON
	IPV4 *string

	// IPV6.
	//
	// Format: JSON
	IPV6 *string

	// Mac.
	Mac string

	// TTL.
	//
	// Format: double
	TTL *float64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the location prototype post device security policy anomaly blacklist params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistParams) WithDefaults() *LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the location prototype post device security policy anomaly blacklist params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the location prototype post device security policy anomaly blacklist params
func (o *LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistParams) WithTimeout(timeout time.Duration) *LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the location prototype post device security policy anomaly blacklist params
func (o *LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the location prototype post device security policy anomaly blacklist params
func (o *LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistParams) WithContext(ctx context.Context) *LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the location prototype post device security policy anomaly blacklist params
func (o *LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the location prototype post device security policy anomaly blacklist params
func (o *LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistParams) WithHTTPClient(client *http.Client) *LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the location prototype post device security policy anomaly blacklist params
func (o *LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeviceType adds the deviceType to the location prototype post device security policy anomaly blacklist params
func (o *LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistParams) WithDeviceType(deviceType string) *LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistParams {
	o.SetDeviceType(deviceType)
	return o
}

// SetDeviceType adds the deviceType to the location prototype post device security policy anomaly blacklist params
func (o *LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistParams) SetDeviceType(deviceType string) {
	o.DeviceType = deviceType
}

// WithFqdn adds the fqdn to the location prototype post device security policy anomaly blacklist params
func (o *LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistParams) WithFqdn(fqdn string) *LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistParams {
	o.SetFqdn(fqdn)
	return o
}

// SetFqdn adds the fqdn to the location prototype post device security policy anomaly blacklist params
func (o *LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistParams) SetFqdn(fqdn string) {
	o.Fqdn = fqdn
}

// WithID adds the id to the location prototype post device security policy anomaly blacklist params
func (o *LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistParams) WithID(id string) *LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the location prototype post device security policy anomaly blacklist params
func (o *LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistParams) SetID(id string) {
	o.ID = id
}

// WithIPV4 adds the iPV4 to the location prototype post device security policy anomaly blacklist params
func (o *LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistParams) WithIPV4(iPV4 *string) *LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistParams {
	o.SetIPV4(iPV4)
	return o
}

// SetIPV4 adds the ipv4 to the location prototype post device security policy anomaly blacklist params
func (o *LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistParams) SetIPV4(iPV4 *string) {
	o.IPV4 = iPV4
}

// WithIPV6 adds the iPV6 to the location prototype post device security policy anomaly blacklist params
func (o *LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistParams) WithIPV6(iPV6 *string) *LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistParams {
	o.SetIPV6(iPV6)
	return o
}

// SetIPV6 adds the ipv6 to the location prototype post device security policy anomaly blacklist params
func (o *LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistParams) SetIPV6(iPV6 *string) {
	o.IPV6 = iPV6
}

// WithMac adds the mac to the location prototype post device security policy anomaly blacklist params
func (o *LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistParams) WithMac(mac string) *LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistParams {
	o.SetMac(mac)
	return o
}

// SetMac adds the mac to the location prototype post device security policy anomaly blacklist params
func (o *LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistParams) SetMac(mac string) {
	o.Mac = mac
}

// WithTTL adds the ttl to the location prototype post device security policy anomaly blacklist params
func (o *LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistParams) WithTTL(ttl *float64) *LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistParams {
	o.SetTTL(ttl)
	return o
}

// SetTTL adds the ttl to the location prototype post device security policy anomaly blacklist params
func (o *LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistParams) SetTTL(ttl *float64) {
	o.TTL = ttl
}

// WriteToRequest writes these params to a swagger request
func (o *LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// form param deviceType
	frDeviceType := o.DeviceType
	fDeviceType := frDeviceType
	if fDeviceType != "" {
		if err := r.SetFormParam("deviceType", fDeviceType); err != nil {
			return err
		}
	}

	// form param fqdn
	frFqdn := o.Fqdn
	fFqdn := frFqdn
	if fFqdn != "" {
		if err := r.SetFormParam("fqdn", fFqdn); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.IPV4 != nil {

		// form param ipv4
		var frIPV4 string
		if o.IPV4 != nil {
			frIPV4 = *o.IPV4
		}
		fIPV4 := frIPV4
		if fIPV4 != "" {
			if err := r.SetFormParam("ipv4", fIPV4); err != nil {
				return err
			}
		}
	}

	if o.IPV6 != nil {

		// form param ipv6
		var frIPV6 string
		if o.IPV6 != nil {
			frIPV6 = *o.IPV6
		}
		fIPV6 := frIPV6
		if fIPV6 != "" {
			if err := r.SetFormParam("ipv6", fIPV6); err != nil {
				return err
			}
		}
	}

	// path param mac
	if err := r.SetPathParam("mac", o.Mac); err != nil {
		return err
	}

	if o.TTL != nil {

		// form param ttl
		var frTTL float64
		if o.TTL != nil {
			frTTL = *o.TTL
		}
		fTTL := swag.FormatFloat64(frTTL)
		if fTTL != "" {
			if err := r.SetFormParam("ttl", fTTL); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
