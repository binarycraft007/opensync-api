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

// NewCustomerPrototypePatchCampaignCaptivePortalNetworkParams creates a new CustomerPrototypePatchCampaignCaptivePortalNetworkParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypePatchCampaignCaptivePortalNetworkParams() *CustomerPrototypePatchCampaignCaptivePortalNetworkParams {
	return &CustomerPrototypePatchCampaignCaptivePortalNetworkParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypePatchCampaignCaptivePortalNetworkParamsWithTimeout creates a new CustomerPrototypePatchCampaignCaptivePortalNetworkParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypePatchCampaignCaptivePortalNetworkParamsWithTimeout(timeout time.Duration) *CustomerPrototypePatchCampaignCaptivePortalNetworkParams {
	return &CustomerPrototypePatchCampaignCaptivePortalNetworkParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypePatchCampaignCaptivePortalNetworkParamsWithContext creates a new CustomerPrototypePatchCampaignCaptivePortalNetworkParams object
// with the ability to set a context for a request.
func NewCustomerPrototypePatchCampaignCaptivePortalNetworkParamsWithContext(ctx context.Context) *CustomerPrototypePatchCampaignCaptivePortalNetworkParams {
	return &CustomerPrototypePatchCampaignCaptivePortalNetworkParams{
		Context: ctx,
	}
}

// NewCustomerPrototypePatchCampaignCaptivePortalNetworkParamsWithHTTPClient creates a new CustomerPrototypePatchCampaignCaptivePortalNetworkParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypePatchCampaignCaptivePortalNetworkParamsWithHTTPClient(client *http.Client) *CustomerPrototypePatchCampaignCaptivePortalNetworkParams {
	return &CustomerPrototypePatchCampaignCaptivePortalNetworkParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypePatchCampaignCaptivePortalNetworkParams contains all the parameters to send to the API endpoint

	for the customer prototype patch campaign captive portal network operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypePatchCampaignCaptivePortalNetworkParams struct {

	// CampaignPayload.
	//
	// Format: JSON
	CampaignPayload *string

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

// WithDefaults hydrates default values in the customer prototype patch campaign captive portal network params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePatchCampaignCaptivePortalNetworkParams) WithDefaults() *CustomerPrototypePatchCampaignCaptivePortalNetworkParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype patch campaign captive portal network params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePatchCampaignCaptivePortalNetworkParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype patch campaign captive portal network params
func (o *CustomerPrototypePatchCampaignCaptivePortalNetworkParams) WithTimeout(timeout time.Duration) *CustomerPrototypePatchCampaignCaptivePortalNetworkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype patch campaign captive portal network params
func (o *CustomerPrototypePatchCampaignCaptivePortalNetworkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype patch campaign captive portal network params
func (o *CustomerPrototypePatchCampaignCaptivePortalNetworkParams) WithContext(ctx context.Context) *CustomerPrototypePatchCampaignCaptivePortalNetworkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype patch campaign captive portal network params
func (o *CustomerPrototypePatchCampaignCaptivePortalNetworkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype patch campaign captive portal network params
func (o *CustomerPrototypePatchCampaignCaptivePortalNetworkParams) WithHTTPClient(client *http.Client) *CustomerPrototypePatchCampaignCaptivePortalNetworkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype patch campaign captive portal network params
func (o *CustomerPrototypePatchCampaignCaptivePortalNetworkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCampaignPayload adds the campaignPayload to the customer prototype patch campaign captive portal network params
func (o *CustomerPrototypePatchCampaignCaptivePortalNetworkParams) WithCampaignPayload(campaignPayload *string) *CustomerPrototypePatchCampaignCaptivePortalNetworkParams {
	o.SetCampaignPayload(campaignPayload)
	return o
}

// SetCampaignPayload adds the campaignPayload to the customer prototype patch campaign captive portal network params
func (o *CustomerPrototypePatchCampaignCaptivePortalNetworkParams) SetCampaignPayload(campaignPayload *string) {
	o.CampaignPayload = campaignPayload
}

// WithID adds the id to the customer prototype patch campaign captive portal network params
func (o *CustomerPrototypePatchCampaignCaptivePortalNetworkParams) WithID(id string) *CustomerPrototypePatchCampaignCaptivePortalNetworkParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype patch campaign captive portal network params
func (o *CustomerPrototypePatchCampaignCaptivePortalNetworkParams) SetID(id string) {
	o.ID = id
}

// WithLocationID adds the locationID to the customer prototype patch campaign captive portal network params
func (o *CustomerPrototypePatchCampaignCaptivePortalNetworkParams) WithLocationID(locationID string) *CustomerPrototypePatchCampaignCaptivePortalNetworkParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype patch campaign captive portal network params
func (o *CustomerPrototypePatchCampaignCaptivePortalNetworkParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WithNetworkID adds the networkID to the customer prototype patch campaign captive portal network params
func (o *CustomerPrototypePatchCampaignCaptivePortalNetworkParams) WithNetworkID(networkID string) *CustomerPrototypePatchCampaignCaptivePortalNetworkParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the customer prototype patch campaign captive portal network params
func (o *CustomerPrototypePatchCampaignCaptivePortalNetworkParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypePatchCampaignCaptivePortalNetworkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CampaignPayload != nil {

		// form param campaignPayload
		var frCampaignPayload string
		if o.CampaignPayload != nil {
			frCampaignPayload = *o.CampaignPayload
		}
		fCampaignPayload := frCampaignPayload
		if fCampaignPayload != "" {
			if err := r.SetFormParam("campaignPayload", fCampaignPayload); err != nil {
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

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}