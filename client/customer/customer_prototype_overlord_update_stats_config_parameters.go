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

// NewCustomerPrototypeOverlordUpdateStatsConfigParams creates a new CustomerPrototypeOverlordUpdateStatsConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypeOverlordUpdateStatsConfigParams() *CustomerPrototypeOverlordUpdateStatsConfigParams {
	return &CustomerPrototypeOverlordUpdateStatsConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypeOverlordUpdateStatsConfigParamsWithTimeout creates a new CustomerPrototypeOverlordUpdateStatsConfigParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypeOverlordUpdateStatsConfigParamsWithTimeout(timeout time.Duration) *CustomerPrototypeOverlordUpdateStatsConfigParams {
	return &CustomerPrototypeOverlordUpdateStatsConfigParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypeOverlordUpdateStatsConfigParamsWithContext creates a new CustomerPrototypeOverlordUpdateStatsConfigParams object
// with the ability to set a context for a request.
func NewCustomerPrototypeOverlordUpdateStatsConfigParamsWithContext(ctx context.Context) *CustomerPrototypeOverlordUpdateStatsConfigParams {
	return &CustomerPrototypeOverlordUpdateStatsConfigParams{
		Context: ctx,
	}
}

// NewCustomerPrototypeOverlordUpdateStatsConfigParamsWithHTTPClient creates a new CustomerPrototypeOverlordUpdateStatsConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypeOverlordUpdateStatsConfigParamsWithHTTPClient(client *http.Client) *CustomerPrototypeOverlordUpdateStatsConfigParams {
	return &CustomerPrototypeOverlordUpdateStatsConfigParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypeOverlordUpdateStatsConfigParams contains all the parameters to send to the API endpoint

	for the customer prototype overlord update stats config operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypeOverlordUpdateStatsConfigParams struct {

	/* ClientAuthFails.

	   string enum: [ AUTO, ENABLE, DISABLE ]
	*/
	ClientAuthFails *string

	/* ID.

	   Customer id

	   Format: JSON
	*/
	ID string

	// LocationID.
	LocationID string

	/* OffChannelScan24.

	   string enum: [ AUTO, ENABLE, DISABLE ]
	*/
	OffChannelScan24 *string

	/* OffChannelScan50.

	   string enum: [ AUTO, ENABLE, DISABLE ]
	*/
	OffChannelScan50 *string

	/* OffChannelScan60.

	   string enum: [ AUTO, ENABLE, DISABLE ]
	*/
	OffChannelScan60 *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer prototype overlord update stats config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeOverlordUpdateStatsConfigParams) WithDefaults() *CustomerPrototypeOverlordUpdateStatsConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype overlord update stats config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeOverlordUpdateStatsConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype overlord update stats config params
func (o *CustomerPrototypeOverlordUpdateStatsConfigParams) WithTimeout(timeout time.Duration) *CustomerPrototypeOverlordUpdateStatsConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype overlord update stats config params
func (o *CustomerPrototypeOverlordUpdateStatsConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype overlord update stats config params
func (o *CustomerPrototypeOverlordUpdateStatsConfigParams) WithContext(ctx context.Context) *CustomerPrototypeOverlordUpdateStatsConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype overlord update stats config params
func (o *CustomerPrototypeOverlordUpdateStatsConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype overlord update stats config params
func (o *CustomerPrototypeOverlordUpdateStatsConfigParams) WithHTTPClient(client *http.Client) *CustomerPrototypeOverlordUpdateStatsConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype overlord update stats config params
func (o *CustomerPrototypeOverlordUpdateStatsConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClientAuthFails adds the clientAuthFails to the customer prototype overlord update stats config params
func (o *CustomerPrototypeOverlordUpdateStatsConfigParams) WithClientAuthFails(clientAuthFails *string) *CustomerPrototypeOverlordUpdateStatsConfigParams {
	o.SetClientAuthFails(clientAuthFails)
	return o
}

// SetClientAuthFails adds the clientAuthFails to the customer prototype overlord update stats config params
func (o *CustomerPrototypeOverlordUpdateStatsConfigParams) SetClientAuthFails(clientAuthFails *string) {
	o.ClientAuthFails = clientAuthFails
}

// WithID adds the id to the customer prototype overlord update stats config params
func (o *CustomerPrototypeOverlordUpdateStatsConfigParams) WithID(id string) *CustomerPrototypeOverlordUpdateStatsConfigParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype overlord update stats config params
func (o *CustomerPrototypeOverlordUpdateStatsConfigParams) SetID(id string) {
	o.ID = id
}

// WithLocationID adds the locationID to the customer prototype overlord update stats config params
func (o *CustomerPrototypeOverlordUpdateStatsConfigParams) WithLocationID(locationID string) *CustomerPrototypeOverlordUpdateStatsConfigParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype overlord update stats config params
func (o *CustomerPrototypeOverlordUpdateStatsConfigParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WithOffChannelScan24 adds the offChannelScan24 to the customer prototype overlord update stats config params
func (o *CustomerPrototypeOverlordUpdateStatsConfigParams) WithOffChannelScan24(offChannelScan24 *string) *CustomerPrototypeOverlordUpdateStatsConfigParams {
	o.SetOffChannelScan24(offChannelScan24)
	return o
}

// SetOffChannelScan24 adds the offChannelScan24 to the customer prototype overlord update stats config params
func (o *CustomerPrototypeOverlordUpdateStatsConfigParams) SetOffChannelScan24(offChannelScan24 *string) {
	o.OffChannelScan24 = offChannelScan24
}

// WithOffChannelScan50 adds the offChannelScan50 to the customer prototype overlord update stats config params
func (o *CustomerPrototypeOverlordUpdateStatsConfigParams) WithOffChannelScan50(offChannelScan50 *string) *CustomerPrototypeOverlordUpdateStatsConfigParams {
	o.SetOffChannelScan50(offChannelScan50)
	return o
}

// SetOffChannelScan50 adds the offChannelScan50 to the customer prototype overlord update stats config params
func (o *CustomerPrototypeOverlordUpdateStatsConfigParams) SetOffChannelScan50(offChannelScan50 *string) {
	o.OffChannelScan50 = offChannelScan50
}

// WithOffChannelScan60 adds the offChannelScan60 to the customer prototype overlord update stats config params
func (o *CustomerPrototypeOverlordUpdateStatsConfigParams) WithOffChannelScan60(offChannelScan60 *string) *CustomerPrototypeOverlordUpdateStatsConfigParams {
	o.SetOffChannelScan60(offChannelScan60)
	return o
}

// SetOffChannelScan60 adds the offChannelScan60 to the customer prototype overlord update stats config params
func (o *CustomerPrototypeOverlordUpdateStatsConfigParams) SetOffChannelScan60(offChannelScan60 *string) {
	o.OffChannelScan60 = offChannelScan60
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypeOverlordUpdateStatsConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ClientAuthFails != nil {

		// form param clientAuthFails
		var frClientAuthFails string
		if o.ClientAuthFails != nil {
			frClientAuthFails = *o.ClientAuthFails
		}
		fClientAuthFails := frClientAuthFails
		if fClientAuthFails != "" {
			if err := r.SetFormParam("clientAuthFails", fClientAuthFails); err != nil {
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

	if o.OffChannelScan24 != nil {

		// form param offChannelScan24
		var frOffChannelScan24 string
		if o.OffChannelScan24 != nil {
			frOffChannelScan24 = *o.OffChannelScan24
		}
		fOffChannelScan24 := frOffChannelScan24
		if fOffChannelScan24 != "" {
			if err := r.SetFormParam("offChannelScan24", fOffChannelScan24); err != nil {
				return err
			}
		}
	}

	if o.OffChannelScan50 != nil {

		// form param offChannelScan50
		var frOffChannelScan50 string
		if o.OffChannelScan50 != nil {
			frOffChannelScan50 = *o.OffChannelScan50
		}
		fOffChannelScan50 := frOffChannelScan50
		if fOffChannelScan50 != "" {
			if err := r.SetFormParam("offChannelScan50", fOffChannelScan50); err != nil {
				return err
			}
		}
	}

	if o.OffChannelScan60 != nil {

		// form param offChannelScan60
		var frOffChannelScan60 string
		if o.OffChannelScan60 != nil {
			frOffChannelScan60 = *o.OffChannelScan60
		}
		fOffChannelScan60 := frOffChannelScan60
		if fOffChannelScan60 != "" {
			if err := r.SetFormParam("offChannelScan60", fOffChannelScan60); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
