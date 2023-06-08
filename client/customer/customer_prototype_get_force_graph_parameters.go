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

// NewCustomerPrototypeGetForceGraphParams creates a new CustomerPrototypeGetForceGraphParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypeGetForceGraphParams() *CustomerPrototypeGetForceGraphParams {
	return &CustomerPrototypeGetForceGraphParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypeGetForceGraphParamsWithTimeout creates a new CustomerPrototypeGetForceGraphParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypeGetForceGraphParamsWithTimeout(timeout time.Duration) *CustomerPrototypeGetForceGraphParams {
	return &CustomerPrototypeGetForceGraphParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypeGetForceGraphParamsWithContext creates a new CustomerPrototypeGetForceGraphParams object
// with the ability to set a context for a request.
func NewCustomerPrototypeGetForceGraphParamsWithContext(ctx context.Context) *CustomerPrototypeGetForceGraphParams {
	return &CustomerPrototypeGetForceGraphParams{
		Context: ctx,
	}
}

// NewCustomerPrototypeGetForceGraphParamsWithHTTPClient creates a new CustomerPrototypeGetForceGraphParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypeGetForceGraphParamsWithHTTPClient(client *http.Client) *CustomerPrototypeGetForceGraphParams {
	return &CustomerPrototypeGetForceGraphParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypeGetForceGraphParams contains all the parameters to send to the API endpoint

	for the customer prototype get force graph operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypeGetForceGraphParams struct {

	// AllSSIDs.
	AllSSIDs *bool

	/* AuthKey.

	   PubNub authKey
	*/
	AuthKey *string

	/* ID.

	   Customer id

	   Format: JSON
	*/
	ID string

	/* IP.

	   deprecated and optional IP address of client displaying the topology
	*/
	IP *string

	// LocationID.
	LocationID string

	/* Mac.

	   optional mac address of client displaying the topology
	*/
	Mac *string

	// ShowPartnerComponent.
	ShowPartnerComponent *bool

	/* SubscribeKey.

	   PubNub subscribeKey
	*/
	SubscribeKey *string

	/* View.

	   view template override (e.g., iguana)
	*/
	View *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer prototype get force graph params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeGetForceGraphParams) WithDefaults() *CustomerPrototypeGetForceGraphParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype get force graph params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeGetForceGraphParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype get force graph params
func (o *CustomerPrototypeGetForceGraphParams) WithTimeout(timeout time.Duration) *CustomerPrototypeGetForceGraphParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype get force graph params
func (o *CustomerPrototypeGetForceGraphParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype get force graph params
func (o *CustomerPrototypeGetForceGraphParams) WithContext(ctx context.Context) *CustomerPrototypeGetForceGraphParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype get force graph params
func (o *CustomerPrototypeGetForceGraphParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype get force graph params
func (o *CustomerPrototypeGetForceGraphParams) WithHTTPClient(client *http.Client) *CustomerPrototypeGetForceGraphParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype get force graph params
func (o *CustomerPrototypeGetForceGraphParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAllSSIDs adds the allSSIDs to the customer prototype get force graph params
func (o *CustomerPrototypeGetForceGraphParams) WithAllSSIDs(allSSIDs *bool) *CustomerPrototypeGetForceGraphParams {
	o.SetAllSSIDs(allSSIDs)
	return o
}

// SetAllSSIDs adds the allSSIDs to the customer prototype get force graph params
func (o *CustomerPrototypeGetForceGraphParams) SetAllSSIDs(allSSIDs *bool) {
	o.AllSSIDs = allSSIDs
}

// WithAuthKey adds the authKey to the customer prototype get force graph params
func (o *CustomerPrototypeGetForceGraphParams) WithAuthKey(authKey *string) *CustomerPrototypeGetForceGraphParams {
	o.SetAuthKey(authKey)
	return o
}

// SetAuthKey adds the authKey to the customer prototype get force graph params
func (o *CustomerPrototypeGetForceGraphParams) SetAuthKey(authKey *string) {
	o.AuthKey = authKey
}

// WithID adds the id to the customer prototype get force graph params
func (o *CustomerPrototypeGetForceGraphParams) WithID(id string) *CustomerPrototypeGetForceGraphParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype get force graph params
func (o *CustomerPrototypeGetForceGraphParams) SetID(id string) {
	o.ID = id
}

// WithIP adds the ip to the customer prototype get force graph params
func (o *CustomerPrototypeGetForceGraphParams) WithIP(ip *string) *CustomerPrototypeGetForceGraphParams {
	o.SetIP(ip)
	return o
}

// SetIP adds the ip to the customer prototype get force graph params
func (o *CustomerPrototypeGetForceGraphParams) SetIP(ip *string) {
	o.IP = ip
}

// WithLocationID adds the locationID to the customer prototype get force graph params
func (o *CustomerPrototypeGetForceGraphParams) WithLocationID(locationID string) *CustomerPrototypeGetForceGraphParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype get force graph params
func (o *CustomerPrototypeGetForceGraphParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WithMac adds the mac to the customer prototype get force graph params
func (o *CustomerPrototypeGetForceGraphParams) WithMac(mac *string) *CustomerPrototypeGetForceGraphParams {
	o.SetMac(mac)
	return o
}

// SetMac adds the mac to the customer prototype get force graph params
func (o *CustomerPrototypeGetForceGraphParams) SetMac(mac *string) {
	o.Mac = mac
}

// WithShowPartnerComponent adds the showPartnerComponent to the customer prototype get force graph params
func (o *CustomerPrototypeGetForceGraphParams) WithShowPartnerComponent(showPartnerComponent *bool) *CustomerPrototypeGetForceGraphParams {
	o.SetShowPartnerComponent(showPartnerComponent)
	return o
}

// SetShowPartnerComponent adds the showPartnerComponent to the customer prototype get force graph params
func (o *CustomerPrototypeGetForceGraphParams) SetShowPartnerComponent(showPartnerComponent *bool) {
	o.ShowPartnerComponent = showPartnerComponent
}

// WithSubscribeKey adds the subscribeKey to the customer prototype get force graph params
func (o *CustomerPrototypeGetForceGraphParams) WithSubscribeKey(subscribeKey *string) *CustomerPrototypeGetForceGraphParams {
	o.SetSubscribeKey(subscribeKey)
	return o
}

// SetSubscribeKey adds the subscribeKey to the customer prototype get force graph params
func (o *CustomerPrototypeGetForceGraphParams) SetSubscribeKey(subscribeKey *string) {
	o.SubscribeKey = subscribeKey
}

// WithView adds the view to the customer prototype get force graph params
func (o *CustomerPrototypeGetForceGraphParams) WithView(view *string) *CustomerPrototypeGetForceGraphParams {
	o.SetView(view)
	return o
}

// SetView adds the view to the customer prototype get force graph params
func (o *CustomerPrototypeGetForceGraphParams) SetView(view *string) {
	o.View = view
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypeGetForceGraphParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AllSSIDs != nil {

		// query param allSSIDs
		var qrAllSSIDs bool

		if o.AllSSIDs != nil {
			qrAllSSIDs = *o.AllSSIDs
		}
		qAllSSIDs := swag.FormatBool(qrAllSSIDs)
		if qAllSSIDs != "" {

			if err := r.SetQueryParam("allSSIDs", qAllSSIDs); err != nil {
				return err
			}
		}
	}

	if o.AuthKey != nil {

		// query param authKey
		var qrAuthKey string

		if o.AuthKey != nil {
			qrAuthKey = *o.AuthKey
		}
		qAuthKey := qrAuthKey
		if qAuthKey != "" {

			if err := r.SetQueryParam("authKey", qAuthKey); err != nil {
				return err
			}
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.IP != nil {

		// query param ip
		var qrIP string

		if o.IP != nil {
			qrIP = *o.IP
		}
		qIP := qrIP
		if qIP != "" {

			if err := r.SetQueryParam("ip", qIP); err != nil {
				return err
			}
		}
	}

	// path param locationId
	if err := r.SetPathParam("locationId", o.LocationID); err != nil {
		return err
	}

	if o.Mac != nil {

		// query param mac
		var qrMac string

		if o.Mac != nil {
			qrMac = *o.Mac
		}
		qMac := qrMac
		if qMac != "" {

			if err := r.SetQueryParam("mac", qMac); err != nil {
				return err
			}
		}
	}

	if o.ShowPartnerComponent != nil {

		// query param showPartnerComponent
		var qrShowPartnerComponent bool

		if o.ShowPartnerComponent != nil {
			qrShowPartnerComponent = *o.ShowPartnerComponent
		}
		qShowPartnerComponent := swag.FormatBool(qrShowPartnerComponent)
		if qShowPartnerComponent != "" {

			if err := r.SetQueryParam("showPartnerComponent", qShowPartnerComponent); err != nil {
				return err
			}
		}
	}

	if o.SubscribeKey != nil {

		// query param subscribeKey
		var qrSubscribeKey string

		if o.SubscribeKey != nil {
			qrSubscribeKey = *o.SubscribeKey
		}
		qSubscribeKey := qrSubscribeKey
		if qSubscribeKey != "" {

			if err := r.SetQueryParam("subscribeKey", qSubscribeKey); err != nil {
				return err
			}
		}
	}

	if o.View != nil {

		// query param view
		var qrView string

		if o.View != nil {
			qrView = *o.View
		}
		qView := qrView
		if qView != "" {

			if err := r.SetQueryParam("view", qView); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}