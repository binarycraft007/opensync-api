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
)

// NewLocationPrototypePatchDeviceSoundingStateParams creates a new LocationPrototypePatchDeviceSoundingStateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLocationPrototypePatchDeviceSoundingStateParams() *LocationPrototypePatchDeviceSoundingStateParams {
	return &LocationPrototypePatchDeviceSoundingStateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLocationPrototypePatchDeviceSoundingStateParamsWithTimeout creates a new LocationPrototypePatchDeviceSoundingStateParams object
// with the ability to set a timeout on a request.
func NewLocationPrototypePatchDeviceSoundingStateParamsWithTimeout(timeout time.Duration) *LocationPrototypePatchDeviceSoundingStateParams {
	return &LocationPrototypePatchDeviceSoundingStateParams{
		timeout: timeout,
	}
}

// NewLocationPrototypePatchDeviceSoundingStateParamsWithContext creates a new LocationPrototypePatchDeviceSoundingStateParams object
// with the ability to set a context for a request.
func NewLocationPrototypePatchDeviceSoundingStateParamsWithContext(ctx context.Context) *LocationPrototypePatchDeviceSoundingStateParams {
	return &LocationPrototypePatchDeviceSoundingStateParams{
		Context: ctx,
	}
}

// NewLocationPrototypePatchDeviceSoundingStateParamsWithHTTPClient creates a new LocationPrototypePatchDeviceSoundingStateParams object
// with the ability to set a custom HTTPClient for a request.
func NewLocationPrototypePatchDeviceSoundingStateParamsWithHTTPClient(client *http.Client) *LocationPrototypePatchDeviceSoundingStateParams {
	return &LocationPrototypePatchDeviceSoundingStateParams{
		HTTPClient: client,
	}
}

/*
LocationPrototypePatchDeviceSoundingStateParams contains all the parameters to send to the API endpoint

	for the location prototype patch device sounding state operation.

	Typically these are written to a http.Request.
*/
type LocationPrototypePatchDeviceSoundingStateParams struct {

	/* ID.

	   Location id

	   Format: JSON
	*/
	ID string

	// SoundingStates.
	SoundingStates interface{}

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the location prototype patch device sounding state params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LocationPrototypePatchDeviceSoundingStateParams) WithDefaults() *LocationPrototypePatchDeviceSoundingStateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the location prototype patch device sounding state params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LocationPrototypePatchDeviceSoundingStateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the location prototype patch device sounding state params
func (o *LocationPrototypePatchDeviceSoundingStateParams) WithTimeout(timeout time.Duration) *LocationPrototypePatchDeviceSoundingStateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the location prototype patch device sounding state params
func (o *LocationPrototypePatchDeviceSoundingStateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the location prototype patch device sounding state params
func (o *LocationPrototypePatchDeviceSoundingStateParams) WithContext(ctx context.Context) *LocationPrototypePatchDeviceSoundingStateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the location prototype patch device sounding state params
func (o *LocationPrototypePatchDeviceSoundingStateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the location prototype patch device sounding state params
func (o *LocationPrototypePatchDeviceSoundingStateParams) WithHTTPClient(client *http.Client) *LocationPrototypePatchDeviceSoundingStateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the location prototype patch device sounding state params
func (o *LocationPrototypePatchDeviceSoundingStateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the location prototype patch device sounding state params
func (o *LocationPrototypePatchDeviceSoundingStateParams) WithID(id string) *LocationPrototypePatchDeviceSoundingStateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the location prototype patch device sounding state params
func (o *LocationPrototypePatchDeviceSoundingStateParams) SetID(id string) {
	o.ID = id
}

// WithSoundingStates adds the soundingStates to the location prototype patch device sounding state params
func (o *LocationPrototypePatchDeviceSoundingStateParams) WithSoundingStates(soundingStates interface{}) *LocationPrototypePatchDeviceSoundingStateParams {
	o.SetSoundingStates(soundingStates)
	return o
}

// SetSoundingStates adds the soundingStates to the location prototype patch device sounding state params
func (o *LocationPrototypePatchDeviceSoundingStateParams) SetSoundingStates(soundingStates interface{}) {
	o.SoundingStates = soundingStates
}

// WriteToRequest writes these params to a swagger request
func (o *LocationPrototypePatchDeviceSoundingStateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
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