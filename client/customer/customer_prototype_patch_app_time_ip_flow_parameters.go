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

// NewCustomerPrototypePatchAppTimeIPFlowParams creates a new CustomerPrototypePatchAppTimeIPFlowParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypePatchAppTimeIPFlowParams() *CustomerPrototypePatchAppTimeIPFlowParams {
	return &CustomerPrototypePatchAppTimeIPFlowParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypePatchAppTimeIPFlowParamsWithTimeout creates a new CustomerPrototypePatchAppTimeIPFlowParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypePatchAppTimeIPFlowParamsWithTimeout(timeout time.Duration) *CustomerPrototypePatchAppTimeIPFlowParams {
	return &CustomerPrototypePatchAppTimeIPFlowParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypePatchAppTimeIPFlowParamsWithContext creates a new CustomerPrototypePatchAppTimeIPFlowParams object
// with the ability to set a context for a request.
func NewCustomerPrototypePatchAppTimeIPFlowParamsWithContext(ctx context.Context) *CustomerPrototypePatchAppTimeIPFlowParams {
	return &CustomerPrototypePatchAppTimeIPFlowParams{
		Context: ctx,
	}
}

// NewCustomerPrototypePatchAppTimeIPFlowParamsWithHTTPClient creates a new CustomerPrototypePatchAppTimeIPFlowParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypePatchAppTimeIPFlowParamsWithHTTPClient(client *http.Client) *CustomerPrototypePatchAppTimeIPFlowParams {
	return &CustomerPrototypePatchAppTimeIPFlowParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypePatchAppTimeIPFlowParams contains all the parameters to send to the API endpoint

	for the customer prototype patch app time Ip flow operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypePatchAppTimeIPFlowParams struct {

	// Enable.
	Enable bool

	// ExpiresAt.
	ExpiresAt *string

	/* ID.

	   Customer id

	   Format: JSON
	*/
	ID string

	// LocationID.
	LocationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer prototype patch app time Ip flow params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePatchAppTimeIPFlowParams) WithDefaults() *CustomerPrototypePatchAppTimeIPFlowParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype patch app time Ip flow params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePatchAppTimeIPFlowParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype patch app time Ip flow params
func (o *CustomerPrototypePatchAppTimeIPFlowParams) WithTimeout(timeout time.Duration) *CustomerPrototypePatchAppTimeIPFlowParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype patch app time Ip flow params
func (o *CustomerPrototypePatchAppTimeIPFlowParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype patch app time Ip flow params
func (o *CustomerPrototypePatchAppTimeIPFlowParams) WithContext(ctx context.Context) *CustomerPrototypePatchAppTimeIPFlowParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype patch app time Ip flow params
func (o *CustomerPrototypePatchAppTimeIPFlowParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype patch app time Ip flow params
func (o *CustomerPrototypePatchAppTimeIPFlowParams) WithHTTPClient(client *http.Client) *CustomerPrototypePatchAppTimeIPFlowParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype patch app time Ip flow params
func (o *CustomerPrototypePatchAppTimeIPFlowParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnable adds the enable to the customer prototype patch app time Ip flow params
func (o *CustomerPrototypePatchAppTimeIPFlowParams) WithEnable(enable bool) *CustomerPrototypePatchAppTimeIPFlowParams {
	o.SetEnable(enable)
	return o
}

// SetEnable adds the enable to the customer prototype patch app time Ip flow params
func (o *CustomerPrototypePatchAppTimeIPFlowParams) SetEnable(enable bool) {
	o.Enable = enable
}

// WithExpiresAt adds the expiresAt to the customer prototype patch app time Ip flow params
func (o *CustomerPrototypePatchAppTimeIPFlowParams) WithExpiresAt(expiresAt *string) *CustomerPrototypePatchAppTimeIPFlowParams {
	o.SetExpiresAt(expiresAt)
	return o
}

// SetExpiresAt adds the expiresAt to the customer prototype patch app time Ip flow params
func (o *CustomerPrototypePatchAppTimeIPFlowParams) SetExpiresAt(expiresAt *string) {
	o.ExpiresAt = expiresAt
}

// WithID adds the id to the customer prototype patch app time Ip flow params
func (o *CustomerPrototypePatchAppTimeIPFlowParams) WithID(id string) *CustomerPrototypePatchAppTimeIPFlowParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype patch app time Ip flow params
func (o *CustomerPrototypePatchAppTimeIPFlowParams) SetID(id string) {
	o.ID = id
}

// WithLocationID adds the locationID to the customer prototype patch app time Ip flow params
func (o *CustomerPrototypePatchAppTimeIPFlowParams) WithLocationID(locationID string) *CustomerPrototypePatchAppTimeIPFlowParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype patch app time Ip flow params
func (o *CustomerPrototypePatchAppTimeIPFlowParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypePatchAppTimeIPFlowParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// form param enable
	frEnable := o.Enable
	fEnable := swag.FormatBool(frEnable)
	if fEnable != "" {
		if err := r.SetFormParam("enable", fEnable); err != nil {
			return err
		}
	}

	if o.ExpiresAt != nil {

		// form param expiresAt
		var frExpiresAt string
		if o.ExpiresAt != nil {
			frExpiresAt = *o.ExpiresAt
		}
		fExpiresAt := frExpiresAt
		if fExpiresAt != "" {
			if err := r.SetFormParam("expiresAt", fExpiresAt); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
