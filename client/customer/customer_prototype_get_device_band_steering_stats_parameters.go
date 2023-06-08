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

// NewCustomerPrototypeGetDeviceBandSteeringStatsParams creates a new CustomerPrototypeGetDeviceBandSteeringStatsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypeGetDeviceBandSteeringStatsParams() *CustomerPrototypeGetDeviceBandSteeringStatsParams {
	return &CustomerPrototypeGetDeviceBandSteeringStatsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypeGetDeviceBandSteeringStatsParamsWithTimeout creates a new CustomerPrototypeGetDeviceBandSteeringStatsParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypeGetDeviceBandSteeringStatsParamsWithTimeout(timeout time.Duration) *CustomerPrototypeGetDeviceBandSteeringStatsParams {
	return &CustomerPrototypeGetDeviceBandSteeringStatsParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypeGetDeviceBandSteeringStatsParamsWithContext creates a new CustomerPrototypeGetDeviceBandSteeringStatsParams object
// with the ability to set a context for a request.
func NewCustomerPrototypeGetDeviceBandSteeringStatsParamsWithContext(ctx context.Context) *CustomerPrototypeGetDeviceBandSteeringStatsParams {
	return &CustomerPrototypeGetDeviceBandSteeringStatsParams{
		Context: ctx,
	}
}

// NewCustomerPrototypeGetDeviceBandSteeringStatsParamsWithHTTPClient creates a new CustomerPrototypeGetDeviceBandSteeringStatsParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypeGetDeviceBandSteeringStatsParamsWithHTTPClient(client *http.Client) *CustomerPrototypeGetDeviceBandSteeringStatsParams {
	return &CustomerPrototypeGetDeviceBandSteeringStatsParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypeGetDeviceBandSteeringStatsParams contains all the parameters to send to the API endpoint

	for the customer prototype get device band steering stats operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypeGetDeviceBandSteeringStatsParams struct {

	/* Granularity.

	   days/hours/minutes

	   Default: "days"
	*/
	Granularity *string

	/* ID.

	   Customer id

	   Format: JSON
	*/
	ID string

	/* Limit.

	   X # of days/hours/minutes

	   Format: double
	   Default: 7
	*/
	Limit *float64

	/* LocationID.

	   location id of devices and nodes
	*/
	LocationID string

	/* Mac.

	   mac id of device
	*/
	Mac string

	/* Start.

	   number of milliseconds elapsed since 1 January 1970 00:00:00 UTC. Defaults to now - (limit * granularity)

	   Format: double
	*/
	Start *float64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer prototype get device band steering stats params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeGetDeviceBandSteeringStatsParams) WithDefaults() *CustomerPrototypeGetDeviceBandSteeringStatsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype get device band steering stats params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeGetDeviceBandSteeringStatsParams) SetDefaults() {
	var (
		granularityDefault = string("days")

		limitDefault = float64(7)
	)

	val := CustomerPrototypeGetDeviceBandSteeringStatsParams{
		Granularity: &granularityDefault,
		Limit:       &limitDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the customer prototype get device band steering stats params
func (o *CustomerPrototypeGetDeviceBandSteeringStatsParams) WithTimeout(timeout time.Duration) *CustomerPrototypeGetDeviceBandSteeringStatsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype get device band steering stats params
func (o *CustomerPrototypeGetDeviceBandSteeringStatsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype get device band steering stats params
func (o *CustomerPrototypeGetDeviceBandSteeringStatsParams) WithContext(ctx context.Context) *CustomerPrototypeGetDeviceBandSteeringStatsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype get device band steering stats params
func (o *CustomerPrototypeGetDeviceBandSteeringStatsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype get device band steering stats params
func (o *CustomerPrototypeGetDeviceBandSteeringStatsParams) WithHTTPClient(client *http.Client) *CustomerPrototypeGetDeviceBandSteeringStatsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype get device band steering stats params
func (o *CustomerPrototypeGetDeviceBandSteeringStatsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGranularity adds the granularity to the customer prototype get device band steering stats params
func (o *CustomerPrototypeGetDeviceBandSteeringStatsParams) WithGranularity(granularity *string) *CustomerPrototypeGetDeviceBandSteeringStatsParams {
	o.SetGranularity(granularity)
	return o
}

// SetGranularity adds the granularity to the customer prototype get device band steering stats params
func (o *CustomerPrototypeGetDeviceBandSteeringStatsParams) SetGranularity(granularity *string) {
	o.Granularity = granularity
}

// WithID adds the id to the customer prototype get device band steering stats params
func (o *CustomerPrototypeGetDeviceBandSteeringStatsParams) WithID(id string) *CustomerPrototypeGetDeviceBandSteeringStatsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype get device band steering stats params
func (o *CustomerPrototypeGetDeviceBandSteeringStatsParams) SetID(id string) {
	o.ID = id
}

// WithLimit adds the limit to the customer prototype get device band steering stats params
func (o *CustomerPrototypeGetDeviceBandSteeringStatsParams) WithLimit(limit *float64) *CustomerPrototypeGetDeviceBandSteeringStatsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the customer prototype get device band steering stats params
func (o *CustomerPrototypeGetDeviceBandSteeringStatsParams) SetLimit(limit *float64) {
	o.Limit = limit
}

// WithLocationID adds the locationID to the customer prototype get device band steering stats params
func (o *CustomerPrototypeGetDeviceBandSteeringStatsParams) WithLocationID(locationID string) *CustomerPrototypeGetDeviceBandSteeringStatsParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype get device band steering stats params
func (o *CustomerPrototypeGetDeviceBandSteeringStatsParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WithMac adds the mac to the customer prototype get device band steering stats params
func (o *CustomerPrototypeGetDeviceBandSteeringStatsParams) WithMac(mac string) *CustomerPrototypeGetDeviceBandSteeringStatsParams {
	o.SetMac(mac)
	return o
}

// SetMac adds the mac to the customer prototype get device band steering stats params
func (o *CustomerPrototypeGetDeviceBandSteeringStatsParams) SetMac(mac string) {
	o.Mac = mac
}

// WithStart adds the start to the customer prototype get device band steering stats params
func (o *CustomerPrototypeGetDeviceBandSteeringStatsParams) WithStart(start *float64) *CustomerPrototypeGetDeviceBandSteeringStatsParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the customer prototype get device band steering stats params
func (o *CustomerPrototypeGetDeviceBandSteeringStatsParams) SetStart(start *float64) {
	o.Start = start
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypeGetDeviceBandSteeringStatsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Granularity != nil {

		// query param granularity
		var qrGranularity string

		if o.Granularity != nil {
			qrGranularity = *o.Granularity
		}
		qGranularity := qrGranularity
		if qGranularity != "" {

			if err := r.SetQueryParam("granularity", qGranularity); err != nil {
				return err
			}
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit float64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatFloat64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	// path param locationId
	if err := r.SetPathParam("locationId", o.LocationID); err != nil {
		return err
	}

	// path param mac
	if err := r.SetPathParam("mac", o.Mac); err != nil {
		return err
	}

	if o.Start != nil {

		// query param start
		var qrStart float64

		if o.Start != nil {
			qrStart = *o.Start
		}
		qStart := swag.FormatFloat64(qrStart)
		if qStart != "" {

			if err := r.SetQueryParam("start", qStart); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
