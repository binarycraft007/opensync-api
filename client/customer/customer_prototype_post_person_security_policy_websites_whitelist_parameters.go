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

// NewCustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParams creates a new CustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParams() *CustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParams {
	return &CustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParamsWithTimeout creates a new CustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParamsWithTimeout(timeout time.Duration) *CustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParams {
	return &CustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParamsWithContext creates a new CustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParams object
// with the ability to set a context for a request.
func NewCustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParamsWithContext(ctx context.Context) *CustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParams {
	return &CustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParams{
		Context: ctx,
	}
}

// NewCustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParamsWithHTTPClient creates a new CustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParamsWithHTTPClient(client *http.Client) *CustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParams {
	return &CustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParams contains all the parameters to send to the API endpoint

	for the customer prototype post person security policy websites whitelist operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParams struct {

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

	/* EventType.

	   EventType field from events response - can be 'kids', 'teenagers', 'secureAndProtect', etc
	*/
	EventType *string

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

	// PersonID.
	PersonID string

	/* Source.

	   Source field from events response - can be 'brightcloud', 'webpulse', 'gatekeeper', 'gatekeeper-ohp'
	*/
	Source *string

	// Type.
	Type *string

	// Value.
	Value *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer prototype post person security policy websites whitelist params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParams) WithDefaults() *CustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype post person security policy websites whitelist params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype post person security policy websites whitelist params
func (o *CustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParams) WithTimeout(timeout time.Duration) *CustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype post person security policy websites whitelist params
func (o *CustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype post person security policy websites whitelist params
func (o *CustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParams) WithContext(ctx context.Context) *CustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype post person security policy websites whitelist params
func (o *CustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype post person security policy websites whitelist params
func (o *CustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParams) WithHTTPClient(client *http.Client) *CustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype post person security policy websites whitelist params
func (o *CustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAkamaiCategoryID adds the akamaiCategoryID to the customer prototype post person security policy websites whitelist params
func (o *CustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParams) WithAkamaiCategoryID(akamaiCategoryID *float64) *CustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParams {
	o.SetAkamaiCategoryID(akamaiCategoryID)
	return o
}

// SetAkamaiCategoryID adds the akamaiCategoryId to the customer prototype post person security policy websites whitelist params
func (o *CustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParams) SetAkamaiCategoryID(akamaiCategoryID *float64) {
	o.AkamaiCategoryID = akamaiCategoryID
}

// WithDirection adds the direction to the customer prototype post person security policy websites whitelist params
func (o *CustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParams) WithDirection(direction *string) *CustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParams {
	o.SetDirection(direction)
	return o
}

// SetDirection adds the direction to the customer prototype post person security policy websites whitelist params
func (o *CustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParams) SetDirection(direction *string) {
	o.Direction = direction
}

// WithDNS adds the dns to the customer prototype post person security policy websites whitelist params
func (o *CustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParams) WithDNS(dns *string) *CustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParams {
	o.SetDNS(dns)
	return o
}

// SetDNS adds the dns to the customer prototype post person security policy websites whitelist params
func (o *CustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParams) SetDNS(dns *string) {
	o.DNS = dns
}

// WithEndTimestamp adds the endTimestamp to the customer prototype post person security policy websites whitelist params
func (o *CustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParams) WithEndTimestamp(endTimestamp *float64) *CustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParams {
	o.SetEndTimestamp(endTimestamp)
	return o
}

// SetEndTimestamp adds the endTimestamp to the customer prototype post person security policy websites whitelist params
func (o *CustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParams) SetEndTimestamp(endTimestamp *float64) {
	o.EndTimestamp = endTimestamp
}

// WithEventType adds the eventType to the customer prototype post person security policy websites whitelist params
func (o *CustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParams) WithEventType(eventType *string) *CustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParams {
	o.SetEventType(eventType)
	return o
}

// SetEventType adds the eventType to the customer prototype post person security policy websites whitelist params
func (o *CustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParams) SetEventType(eventType *string) {
	o.EventType = eventType
}

// WithGeoLocation adds the geoLocation to the customer prototype post person security policy websites whitelist params
func (o *CustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParams) WithGeoLocation(geoLocation *string) *CustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParams {
	o.SetGeoLocation(geoLocation)
	return o
}

// SetGeoLocation adds the geoLocation to the customer prototype post person security policy websites whitelist params
func (o *CustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParams) SetGeoLocation(geoLocation *string) {
	o.GeoLocation = geoLocation
}

// WithID adds the id to the customer prototype post person security policy websites whitelist params
func (o *CustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParams) WithID(id string) *CustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype post person security policy websites whitelist params
func (o *CustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParams) SetID(id string) {
	o.ID = id
}

// WithLocationID adds the locationID to the customer prototype post person security policy websites whitelist params
func (o *CustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParams) WithLocationID(locationID string) *CustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype post person security policy websites whitelist params
func (o *CustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WithPersonID adds the personID to the customer prototype post person security policy websites whitelist params
func (o *CustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParams) WithPersonID(personID string) *CustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParams {
	o.SetPersonID(personID)
	return o
}

// SetPersonID adds the personId to the customer prototype post person security policy websites whitelist params
func (o *CustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParams) SetPersonID(personID string) {
	o.PersonID = personID
}

// WithSource adds the source to the customer prototype post person security policy websites whitelist params
func (o *CustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParams) WithSource(source *string) *CustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParams {
	o.SetSource(source)
	return o
}

// SetSource adds the source to the customer prototype post person security policy websites whitelist params
func (o *CustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParams) SetSource(source *string) {
	o.Source = source
}

// WithType adds the typeVar to the customer prototype post person security policy websites whitelist params
func (o *CustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParams) WithType(typeVar *string) *CustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the customer prototype post person security policy websites whitelist params
func (o *CustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WithValue adds the value to the customer prototype post person security policy websites whitelist params
func (o *CustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParams) WithValue(value *string) *CustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParams {
	o.SetValue(value)
	return o
}

// SetValue adds the value to the customer prototype post person security policy websites whitelist params
func (o *CustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParams) SetValue(value *string) {
	o.Value = value
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.EventType != nil {

		// form param eventType
		var frEventType string
		if o.EventType != nil {
			frEventType = *o.EventType
		}
		fEventType := frEventType
		if fEventType != "" {
			if err := r.SetFormParam("eventType", fEventType); err != nil {
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

	// path param personId
	if err := r.SetPathParam("personId", o.PersonID); err != nil {
		return err
	}

	if o.Source != nil {

		// form param source
		var frSource string
		if o.Source != nil {
			frSource = *o.Source
		}
		fSource := frSource
		if fSource != "" {
			if err := r.SetFormParam("source", fSource); err != nil {
				return err
			}
		}
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
