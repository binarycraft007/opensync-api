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

// NewCustomerPrototypePatchDeviceSoundingStateParams creates a new CustomerPrototypePatchDeviceSoundingStateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypePatchDeviceSoundingStateParams() *CustomerPrototypePatchDeviceSoundingStateParams {
	return &CustomerPrototypePatchDeviceSoundingStateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypePatchDeviceSoundingStateParamsWithTimeout creates a new CustomerPrototypePatchDeviceSoundingStateParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypePatchDeviceSoundingStateParamsWithTimeout(timeout time.Duration) *CustomerPrototypePatchDeviceSoundingStateParams {
	return &CustomerPrototypePatchDeviceSoundingStateParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypePatchDeviceSoundingStateParamsWithContext creates a new CustomerPrototypePatchDeviceSoundingStateParams object
// with the ability to set a context for a request.
func NewCustomerPrototypePatchDeviceSoundingStateParamsWithContext(ctx context.Context) *CustomerPrototypePatchDeviceSoundingStateParams {
	return &CustomerPrototypePatchDeviceSoundingStateParams{
		Context: ctx,
	}
}

// NewCustomerPrototypePatchDeviceSoundingStateParamsWithHTTPClient creates a new CustomerPrototypePatchDeviceSoundingStateParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypePatchDeviceSoundingStateParamsWithHTTPClient(client *http.Client) *CustomerPrototypePatchDeviceSoundingStateParams {
	return &CustomerPrototypePatchDeviceSoundingStateParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypePatchDeviceSoundingStateParams contains all the parameters to send to the API endpoint

	for the customer prototype patch device sounding state operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypePatchDeviceSoundingStateParams struct {

	/* ID.

	   Customer id

	   Format: JSON
	*/
	ID string

	// LocationID.
	//
	// Format: JSON
	LocationID string

	// SoundingStates.
	SoundingStates interface{}

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer prototype patch device sounding state params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePatchDeviceSoundingStateParams) WithDefaults() *CustomerPrototypePatchDeviceSoundingStateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype patch device sounding state params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePatchDeviceSoundingStateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype patch device sounding state params
func (o *CustomerPrototypePatchDeviceSoundingStateParams) WithTimeout(timeout time.Duration) *CustomerPrototypePatchDeviceSoundingStateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype patch device sounding state params
func (o *CustomerPrototypePatchDeviceSoundingStateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype patch device sounding state params
func (o *CustomerPrototypePatchDeviceSoundingStateParams) WithContext(ctx context.Context) *CustomerPrototypePatchDeviceSoundingStateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype patch device sounding state params
func (o *CustomerPrototypePatchDeviceSoundingStateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype patch device sounding state params
func (o *CustomerPrototypePatchDeviceSoundingStateParams) WithHTTPClient(client *http.Client) *CustomerPrototypePatchDeviceSoundingStateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype patch device sounding state params
func (o *CustomerPrototypePatchDeviceSoundingStateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the customer prototype patch device sounding state params
func (o *CustomerPrototypePatchDeviceSoundingStateParams) WithID(id string) *CustomerPrototypePatchDeviceSoundingStateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype patch device sounding state params
func (o *CustomerPrototypePatchDeviceSoundingStateParams) SetID(id string) {
	o.ID = id
}

// WithLocationID adds the locationID to the customer prototype patch device sounding state params
func (o *CustomerPrototypePatchDeviceSoundingStateParams) WithLocationID(locationID string) *CustomerPrototypePatchDeviceSoundingStateParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype patch device sounding state params
func (o *CustomerPrototypePatchDeviceSoundingStateParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WithSoundingStates adds the soundingStates to the customer prototype patch device sounding state params
func (o *CustomerPrototypePatchDeviceSoundingStateParams) WithSoundingStates(soundingStates interface{}) *CustomerPrototypePatchDeviceSoundingStateParams {
	o.SetSoundingStates(soundingStates)
	return o
}

// SetSoundingStates adds the soundingStates to the customer prototype patch device sounding state params
func (o *CustomerPrototypePatchDeviceSoundingStateParams) SetSoundingStates(soundingStates interface{}) {
	o.SoundingStates = soundingStates
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypePatchDeviceSoundingStateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if o.SoundingStates != nil {
		if err := r.SetBodyParam(o.SoundingStates); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
