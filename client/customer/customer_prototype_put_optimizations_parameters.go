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

// NewCustomerPrototypePutOptimizationsParams creates a new CustomerPrototypePutOptimizationsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypePutOptimizationsParams() *CustomerPrototypePutOptimizationsParams {
	return &CustomerPrototypePutOptimizationsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypePutOptimizationsParamsWithTimeout creates a new CustomerPrototypePutOptimizationsParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypePutOptimizationsParamsWithTimeout(timeout time.Duration) *CustomerPrototypePutOptimizationsParams {
	return &CustomerPrototypePutOptimizationsParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypePutOptimizationsParamsWithContext creates a new CustomerPrototypePutOptimizationsParams object
// with the ability to set a context for a request.
func NewCustomerPrototypePutOptimizationsParamsWithContext(ctx context.Context) *CustomerPrototypePutOptimizationsParams {
	return &CustomerPrototypePutOptimizationsParams{
		Context: ctx,
	}
}

// NewCustomerPrototypePutOptimizationsParamsWithHTTPClient creates a new CustomerPrototypePutOptimizationsParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypePutOptimizationsParamsWithHTTPClient(client *http.Client) *CustomerPrototypePutOptimizationsParams {
	return &CustomerPrototypePutOptimizationsParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypePutOptimizationsParams contains all the parameters to send to the API endpoint

	for the customer prototype put optimizations operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypePutOptimizationsParams struct {

	/* Auto.

	   defaults to true

	   Format: JSON
	*/
	Auto string

	/* DfsMode.

	   enum of values include: auto, enable, disable, demo, HomeNonDFSChannels, usDfs, deviceAware

	   Default: "auto"
	*/
	DfsMode *string

	/* HopPenalty.

	   enum of values include: auto, low, medium, high

	   Default: "auto"
	*/
	HopPenalty *string

	/* ID.

	   Customer id

	   Format: JSON
	*/
	ID string

	// LocationID.
	LocationID string

	/* PreCACScheduler.

	   enum of values include: auto, enable, disable

	   Default: "auto"
	*/
	PreCACScheduler *string

	/* Prefer160MhzMode.

	   enum of values include: auto, enable, disable

	   Default: "auto"
	*/
	Prefer160MhzMode *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer prototype put optimizations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePutOptimizationsParams) WithDefaults() *CustomerPrototypePutOptimizationsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype put optimizations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePutOptimizationsParams) SetDefaults() {
	var (
		dfsModeDefault = string("auto")

		hopPenaltyDefault = string("auto")

		preCACSchedulerDefault = string("auto")

		prefer160MhzModeDefault = string("auto")
	)

	val := CustomerPrototypePutOptimizationsParams{
		DfsMode:          &dfsModeDefault,
		HopPenalty:       &hopPenaltyDefault,
		PreCACScheduler:  &preCACSchedulerDefault,
		Prefer160MhzMode: &prefer160MhzModeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the customer prototype put optimizations params
func (o *CustomerPrototypePutOptimizationsParams) WithTimeout(timeout time.Duration) *CustomerPrototypePutOptimizationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype put optimizations params
func (o *CustomerPrototypePutOptimizationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype put optimizations params
func (o *CustomerPrototypePutOptimizationsParams) WithContext(ctx context.Context) *CustomerPrototypePutOptimizationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype put optimizations params
func (o *CustomerPrototypePutOptimizationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype put optimizations params
func (o *CustomerPrototypePutOptimizationsParams) WithHTTPClient(client *http.Client) *CustomerPrototypePutOptimizationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype put optimizations params
func (o *CustomerPrototypePutOptimizationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuto adds the auto to the customer prototype put optimizations params
func (o *CustomerPrototypePutOptimizationsParams) WithAuto(auto string) *CustomerPrototypePutOptimizationsParams {
	o.SetAuto(auto)
	return o
}

// SetAuto adds the auto to the customer prototype put optimizations params
func (o *CustomerPrototypePutOptimizationsParams) SetAuto(auto string) {
	o.Auto = auto
}

// WithDfsMode adds the dfsMode to the customer prototype put optimizations params
func (o *CustomerPrototypePutOptimizationsParams) WithDfsMode(dfsMode *string) *CustomerPrototypePutOptimizationsParams {
	o.SetDfsMode(dfsMode)
	return o
}

// SetDfsMode adds the dfsMode to the customer prototype put optimizations params
func (o *CustomerPrototypePutOptimizationsParams) SetDfsMode(dfsMode *string) {
	o.DfsMode = dfsMode
}

// WithHopPenalty adds the hopPenalty to the customer prototype put optimizations params
func (o *CustomerPrototypePutOptimizationsParams) WithHopPenalty(hopPenalty *string) *CustomerPrototypePutOptimizationsParams {
	o.SetHopPenalty(hopPenalty)
	return o
}

// SetHopPenalty adds the hopPenalty to the customer prototype put optimizations params
func (o *CustomerPrototypePutOptimizationsParams) SetHopPenalty(hopPenalty *string) {
	o.HopPenalty = hopPenalty
}

// WithID adds the id to the customer prototype put optimizations params
func (o *CustomerPrototypePutOptimizationsParams) WithID(id string) *CustomerPrototypePutOptimizationsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype put optimizations params
func (o *CustomerPrototypePutOptimizationsParams) SetID(id string) {
	o.ID = id
}

// WithLocationID adds the locationID to the customer prototype put optimizations params
func (o *CustomerPrototypePutOptimizationsParams) WithLocationID(locationID string) *CustomerPrototypePutOptimizationsParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype put optimizations params
func (o *CustomerPrototypePutOptimizationsParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WithPreCACScheduler adds the preCACScheduler to the customer prototype put optimizations params
func (o *CustomerPrototypePutOptimizationsParams) WithPreCACScheduler(preCACScheduler *string) *CustomerPrototypePutOptimizationsParams {
	o.SetPreCACScheduler(preCACScheduler)
	return o
}

// SetPreCACScheduler adds the preCACScheduler to the customer prototype put optimizations params
func (o *CustomerPrototypePutOptimizationsParams) SetPreCACScheduler(preCACScheduler *string) {
	o.PreCACScheduler = preCACScheduler
}

// WithPrefer160MhzMode adds the prefer160MhzMode to the customer prototype put optimizations params
func (o *CustomerPrototypePutOptimizationsParams) WithPrefer160MhzMode(prefer160MhzMode *string) *CustomerPrototypePutOptimizationsParams {
	o.SetPrefer160MhzMode(prefer160MhzMode)
	return o
}

// SetPrefer160MhzMode adds the prefer160MhzMode to the customer prototype put optimizations params
func (o *CustomerPrototypePutOptimizationsParams) SetPrefer160MhzMode(prefer160MhzMode *string) {
	o.Prefer160MhzMode = prefer160MhzMode
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypePutOptimizationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// form param auto
	frAuto := o.Auto
	fAuto := frAuto
	if fAuto != "" {
		if err := r.SetFormParam("auto", fAuto); err != nil {
			return err
		}
	}

	if o.DfsMode != nil {

		// form param dfsMode
		var frDfsMode string
		if o.DfsMode != nil {
			frDfsMode = *o.DfsMode
		}
		fDfsMode := frDfsMode
		if fDfsMode != "" {
			if err := r.SetFormParam("dfsMode", fDfsMode); err != nil {
				return err
			}
		}
	}

	if o.HopPenalty != nil {

		// form param hopPenalty
		var frHopPenalty string
		if o.HopPenalty != nil {
			frHopPenalty = *o.HopPenalty
		}
		fHopPenalty := frHopPenalty
		if fHopPenalty != "" {
			if err := r.SetFormParam("hopPenalty", fHopPenalty); err != nil {
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

	if o.PreCACScheduler != nil {

		// form param preCACScheduler
		var frPreCACScheduler string
		if o.PreCACScheduler != nil {
			frPreCACScheduler = *o.PreCACScheduler
		}
		fPreCACScheduler := frPreCACScheduler
		if fPreCACScheduler != "" {
			if err := r.SetFormParam("preCACScheduler", fPreCACScheduler); err != nil {
				return err
			}
		}
	}

	if o.Prefer160MhzMode != nil {

		// form param prefer160MhzMode
		var frPrefer160MhzMode string
		if o.Prefer160MhzMode != nil {
			frPrefer160MhzMode = *o.Prefer160MhzMode
		}
		fPrefer160MhzMode := frPrefer160MhzMode
		if fPrefer160MhzMode != "" {
			if err := r.SetFormParam("prefer160MhzMode", fPrefer160MhzMode); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
