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

// NewCustomerPrototypePostCaptivePortalCampaignAssetParams creates a new CustomerPrototypePostCaptivePortalCampaignAssetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypePostCaptivePortalCampaignAssetParams() *CustomerPrototypePostCaptivePortalCampaignAssetParams {
	return &CustomerPrototypePostCaptivePortalCampaignAssetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypePostCaptivePortalCampaignAssetParamsWithTimeout creates a new CustomerPrototypePostCaptivePortalCampaignAssetParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypePostCaptivePortalCampaignAssetParamsWithTimeout(timeout time.Duration) *CustomerPrototypePostCaptivePortalCampaignAssetParams {
	return &CustomerPrototypePostCaptivePortalCampaignAssetParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypePostCaptivePortalCampaignAssetParamsWithContext creates a new CustomerPrototypePostCaptivePortalCampaignAssetParams object
// with the ability to set a context for a request.
func NewCustomerPrototypePostCaptivePortalCampaignAssetParamsWithContext(ctx context.Context) *CustomerPrototypePostCaptivePortalCampaignAssetParams {
	return &CustomerPrototypePostCaptivePortalCampaignAssetParams{
		Context: ctx,
	}
}

// NewCustomerPrototypePostCaptivePortalCampaignAssetParamsWithHTTPClient creates a new CustomerPrototypePostCaptivePortalCampaignAssetParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypePostCaptivePortalCampaignAssetParamsWithHTTPClient(client *http.Client) *CustomerPrototypePostCaptivePortalCampaignAssetParams {
	return &CustomerPrototypePostCaptivePortalCampaignAssetParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypePostCaptivePortalCampaignAssetParams contains all the parameters to send to the API endpoint

	for the customer prototype post captive portal campaign asset operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypePostCaptivePortalCampaignAssetParams struct {

	/* ID.

	   Customer id

	   Format: JSON
	*/
	ID string

	// LocationID.
	//
	// Format: JSON
	LocationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer prototype post captive portal campaign asset params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePostCaptivePortalCampaignAssetParams) WithDefaults() *CustomerPrototypePostCaptivePortalCampaignAssetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype post captive portal campaign asset params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePostCaptivePortalCampaignAssetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype post captive portal campaign asset params
func (o *CustomerPrototypePostCaptivePortalCampaignAssetParams) WithTimeout(timeout time.Duration) *CustomerPrototypePostCaptivePortalCampaignAssetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype post captive portal campaign asset params
func (o *CustomerPrototypePostCaptivePortalCampaignAssetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype post captive portal campaign asset params
func (o *CustomerPrototypePostCaptivePortalCampaignAssetParams) WithContext(ctx context.Context) *CustomerPrototypePostCaptivePortalCampaignAssetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype post captive portal campaign asset params
func (o *CustomerPrototypePostCaptivePortalCampaignAssetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype post captive portal campaign asset params
func (o *CustomerPrototypePostCaptivePortalCampaignAssetParams) WithHTTPClient(client *http.Client) *CustomerPrototypePostCaptivePortalCampaignAssetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype post captive portal campaign asset params
func (o *CustomerPrototypePostCaptivePortalCampaignAssetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the customer prototype post captive portal campaign asset params
func (o *CustomerPrototypePostCaptivePortalCampaignAssetParams) WithID(id string) *CustomerPrototypePostCaptivePortalCampaignAssetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype post captive portal campaign asset params
func (o *CustomerPrototypePostCaptivePortalCampaignAssetParams) SetID(id string) {
	o.ID = id
}

// WithLocationID adds the locationID to the customer prototype post captive portal campaign asset params
func (o *CustomerPrototypePostCaptivePortalCampaignAssetParams) WithLocationID(locationID string) *CustomerPrototypePostCaptivePortalCampaignAssetParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype post captive portal campaign asset params
func (o *CustomerPrototypePostCaptivePortalCampaignAssetParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypePostCaptivePortalCampaignAssetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
