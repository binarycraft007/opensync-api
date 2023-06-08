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

// NewCustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklistParams creates a new CustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklistParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklistParams() *CustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklistParams {
	return &CustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklistParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklistParamsWithTimeout creates a new CustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklistParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklistParamsWithTimeout(timeout time.Duration) *CustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklistParams {
	return &CustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklistParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklistParamsWithContext creates a new CustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklistParams object
// with the ability to set a context for a request.
func NewCustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklistParamsWithContext(ctx context.Context) *CustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklistParams {
	return &CustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklistParams{
		Context: ctx,
	}
}

// NewCustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklistParamsWithHTTPClient creates a new CustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklistParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklistParamsWithHTTPClient(client *http.Client) *CustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklistParams {
	return &CustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklistParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklistParams contains all the parameters to send to the API endpoint

	for the customer prototype post device security policy websites blacklist operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklistParams struct {

	/* AkamaiCategoryID.

	   the akamai category id, number

	   Format: double
	*/
	AkamaiCategoryID *float64

	// Direction.
	Direction *string

	// DNS.
	DNS *string

	/* EndTimestamp.

	   the end time stamp,  UTC unix epoch timestamp in ms

	   Format: double
	*/
	EndTimestamp *float64

	// GeoLocation.
	//
	// Format: JSON
	GeoLocation *string

	/* ID.

	   Customer id

	   Format: JSON
	*/
	ID string

	// LocationID.
	LocationID string

	// Mac.
	Mac string

	// Type.
	Type *string

	// Value.
	Value *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer prototype post device security policy websites blacklist params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklistParams) WithDefaults() *CustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklistParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype post device security policy websites blacklist params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklistParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype post device security policy websites blacklist params
func (o *CustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklistParams) WithTimeout(timeout time.Duration) *CustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklistParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype post device security policy websites blacklist params
func (o *CustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklistParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype post device security policy websites blacklist params
func (o *CustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklistParams) WithContext(ctx context.Context) *CustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklistParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype post device security policy websites blacklist params
func (o *CustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklistParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype post device security policy websites blacklist params
func (o *CustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklistParams) WithHTTPClient(client *http.Client) *CustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklistParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype post device security policy websites blacklist params
func (o *CustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklistParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAkamaiCategoryID adds the akamaiCategoryID to the customer prototype post device security policy websites blacklist params
func (o *CustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklistParams) WithAkamaiCategoryID(akamaiCategoryID *float64) *CustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklistParams {
	o.SetAkamaiCategoryID(akamaiCategoryID)
	return o
}

// SetAkamaiCategoryID adds the akamaiCategoryId to the customer prototype post device security policy websites blacklist params
func (o *CustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklistParams) SetAkamaiCategoryID(akamaiCategoryID *float64) {
	o.AkamaiCategoryID = akamaiCategoryID
}

// WithDirection adds the direction to the customer prototype post device security policy websites blacklist params
func (o *CustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklistParams) WithDirection(direction *string) *CustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklistParams {
	o.SetDirection(direction)
	return o
}

// SetDirection adds the direction to the customer prototype post device security policy websites blacklist params
func (o *CustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklistParams) SetDirection(direction *string) {
	o.Direction = direction
}

// WithDNS adds the dns to the customer prototype post device security policy websites blacklist params
func (o *CustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklistParams) WithDNS(dns *string) *CustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklistParams {
	o.SetDNS(dns)
	return o
}

// SetDNS adds the dns to the customer prototype post device security policy websites blacklist params
func (o *CustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklistParams) SetDNS(dns *string) {
	o.DNS = dns
}

// WithEndTimestamp adds the endTimestamp to the customer prototype post device security policy websites blacklist params
func (o *CustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklistParams) WithEndTimestamp(endTimestamp *float64) *CustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklistParams {
	o.SetEndTimestamp(endTimestamp)
	return o
}

// SetEndTimestamp adds the endTimestamp to the customer prototype post device security policy websites blacklist params
func (o *CustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklistParams) SetEndTimestamp(endTimestamp *float64) {
	o.EndTimestamp = endTimestamp
}

// WithGeoLocation adds the geoLocation to the customer prototype post device security policy websites blacklist params
func (o *CustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklistParams) WithGeoLocation(geoLocation *string) *CustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklistParams {
	o.SetGeoLocation(geoLocation)
	return o
}

// SetGeoLocation adds the geoLocation to the customer prototype post device security policy websites blacklist params
func (o *CustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklistParams) SetGeoLocation(geoLocation *string) {
	o.GeoLocation = geoLocation
}

// WithID adds the id to the customer prototype post device security policy websites blacklist params
func (o *CustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklistParams) WithID(id string) *CustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklistParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype post device security policy websites blacklist params
func (o *CustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklistParams) SetID(id string) {
	o.ID = id
}

// WithLocationID adds the locationID to the customer prototype post device security policy websites blacklist params
func (o *CustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklistParams) WithLocationID(locationID string) *CustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklistParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype post device security policy websites blacklist params
func (o *CustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklistParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WithMac adds the mac to the customer prototype post device security policy websites blacklist params
func (o *CustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklistParams) WithMac(mac string) *CustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklistParams {
	o.SetMac(mac)
	return o
}

// SetMac adds the mac to the customer prototype post device security policy websites blacklist params
func (o *CustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklistParams) SetMac(mac string) {
	o.Mac = mac
}

// WithType adds the typeVar to the customer prototype post device security policy websites blacklist params
func (o *CustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklistParams) WithType(typeVar *string) *CustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklistParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the customer prototype post device security policy websites blacklist params
func (o *CustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklistParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WithValue adds the value to the customer prototype post device security policy websites blacklist params
func (o *CustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklistParams) WithValue(value *string) *CustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklistParams {
	o.SetValue(value)
	return o
}

// SetValue adds the value to the customer prototype post device security policy websites blacklist params
func (o *CustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklistParams) SetValue(value *string) {
	o.Value = value
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklistParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AkamaiCategoryID != nil {

		// form param akamaiCategoryId
		var frAkamaiCategoryID float64
		if o.AkamaiCategoryID != nil {
			frAkamaiCategoryID = *o.AkamaiCategoryID
		}
		fAkamaiCategoryID := swag.FormatFloat64(frAkamaiCategoryID)
		if fAkamaiCategoryID != "" {
			if err := r.SetFormParam("akamaiCategoryId", fAkamaiCategoryID); err != nil {
				return err
			}
		}
	}

	if o.Direction != nil {

		// form param direction
		var frDirection string
		if o.Direction != nil {
			frDirection = *o.Direction
		}
		fDirection := frDirection
		if fDirection != "" {
			if err := r.SetFormParam("direction", fDirection); err != nil {
				return err
			}
		}
	}

	if o.DNS != nil {

		// form param dns
		var frDNS string
		if o.DNS != nil {
			frDNS = *o.DNS
		}
		fDNS := frDNS
		if fDNS != "" {
			if err := r.SetFormParam("dns", fDNS); err != nil {
				return err
			}
		}
	}

	if o.EndTimestamp != nil {

		// form param endTimestamp
		var frEndTimestamp float64
		if o.EndTimestamp != nil {
			frEndTimestamp = *o.EndTimestamp
		}
		fEndTimestamp := swag.FormatFloat64(frEndTimestamp)
		if fEndTimestamp != "" {
			if err := r.SetFormParam("endTimestamp", fEndTimestamp); err != nil {
				return err
			}
		}
	}

	if o.GeoLocation != nil {

		// form param geoLocation
		var frGeoLocation string
		if o.GeoLocation != nil {
			frGeoLocation = *o.GeoLocation
		}
		fGeoLocation := frGeoLocation
		if fGeoLocation != "" {
			if err := r.SetFormParam("geoLocation", fGeoLocation); err != nil {
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

	// path param mac
	if err := r.SetPathParam("mac", o.Mac); err != nil {
		return err
	}

	if o.Type != nil {

		// form param type
		var frType string
		if o.Type != nil {
			frType = *o.Type
		}
		fType := frType
		if fType != "" {
			if err := r.SetFormParam("type", fType); err != nil {
				return err
			}
		}
	}

	if o.Value != nil {

		// form param value
		var frValue string
		if o.Value != nil {
			frValue = *o.Value
		}
		fValue := frValue
		if fValue != "" {
			if err := r.SetFormParam("value", fValue); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
