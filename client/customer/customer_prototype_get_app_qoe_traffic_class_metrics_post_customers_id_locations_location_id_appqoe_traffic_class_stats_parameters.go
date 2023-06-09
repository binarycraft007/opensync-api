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

// NewCustomerPrototypeGetAppQoeTrafficClassMetricsPostCustomersIDLocationsLocationIDAppqoeTrafficClassStatsParams creates a new CustomerPrototypeGetAppQoeTrafficClassMetricsPostCustomersIDLocationsLocationIDAppqoeTrafficClassStatsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypeGetAppQoeTrafficClassMetricsPostCustomersIDLocationsLocationIDAppqoeTrafficClassStatsParams() *CustomerPrototypeGetAppQoeTrafficClassMetricsPostCustomersIDLocationsLocationIDAppqoeTrafficClassStatsParams {
	return &CustomerPrototypeGetAppQoeTrafficClassMetricsPostCustomersIDLocationsLocationIDAppqoeTrafficClassStatsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypeGetAppQoeTrafficClassMetricsPostCustomersIDLocationsLocationIDAppqoeTrafficClassStatsParamsWithTimeout creates a new CustomerPrototypeGetAppQoeTrafficClassMetricsPostCustomersIDLocationsLocationIDAppqoeTrafficClassStatsParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypeGetAppQoeTrafficClassMetricsPostCustomersIDLocationsLocationIDAppqoeTrafficClassStatsParamsWithTimeout(timeout time.Duration) *CustomerPrototypeGetAppQoeTrafficClassMetricsPostCustomersIDLocationsLocationIDAppqoeTrafficClassStatsParams {
	return &CustomerPrototypeGetAppQoeTrafficClassMetricsPostCustomersIDLocationsLocationIDAppqoeTrafficClassStatsParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypeGetAppQoeTrafficClassMetricsPostCustomersIDLocationsLocationIDAppqoeTrafficClassStatsParamsWithContext creates a new CustomerPrototypeGetAppQoeTrafficClassMetricsPostCustomersIDLocationsLocationIDAppqoeTrafficClassStatsParams object
// with the ability to set a context for a request.
func NewCustomerPrototypeGetAppQoeTrafficClassMetricsPostCustomersIDLocationsLocationIDAppqoeTrafficClassStatsParamsWithContext(ctx context.Context) *CustomerPrototypeGetAppQoeTrafficClassMetricsPostCustomersIDLocationsLocationIDAppqoeTrafficClassStatsParams {
	return &CustomerPrototypeGetAppQoeTrafficClassMetricsPostCustomersIDLocationsLocationIDAppqoeTrafficClassStatsParams{
		Context: ctx,
	}
}

// NewCustomerPrototypeGetAppQoeTrafficClassMetricsPostCustomersIDLocationsLocationIDAppqoeTrafficClassStatsParamsWithHTTPClient creates a new CustomerPrototypeGetAppQoeTrafficClassMetricsPostCustomersIDLocationsLocationIDAppqoeTrafficClassStatsParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypeGetAppQoeTrafficClassMetricsPostCustomersIDLocationsLocationIDAppqoeTrafficClassStatsParamsWithHTTPClient(client *http.Client) *CustomerPrototypeGetAppQoeTrafficClassMetricsPostCustomersIDLocationsLocationIDAppqoeTrafficClassStatsParams {
	return &CustomerPrototypeGetAppQoeTrafficClassMetricsPostCustomersIDLocationsLocationIDAppqoeTrafficClassStatsParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypeGetAppQoeTrafficClassMetricsPostCustomersIDLocationsLocationIDAppqoeTrafficClassStatsParams contains all the parameters to send to the API endpoint

	for the customer prototype get app qoe traffic class metrics post customers id locations location Id appqoe traffic class stats operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypeGetAppQoeTrafficClassMetricsPostCustomersIDLocationsLocationIDAppqoeTrafficClassStatsParams struct {

	/* EndTime.

	   format yyyy-mm-ddThh:MM:ss.nnnZ, 24 hours time specified in UTC

	   Format: date-time
	*/
	EndTime strfmt.DateTime

	/* Granularity.

	   any of the values - total/1 minute/15 minutes/1 hour/1 day
	*/
	Granularity string

	/* ID.

	   Customer id

	   Format: JSON
	*/
	ID string

	/* Limit.

	   Maximum number of devices to return.

	   Format: double
	*/
	Limit *float64

	// LocationID.
	LocationID string

	/* Macs.

	   array of macs

	   Format: JSON
	*/
	Macs *string

	/* StartTime.

	   format yyyy-mm-ddThh:MM:ss.nnnZ, 24 hours time specified in UTC

	   Format: date-time
	*/
	StartTime strfmt.DateTime

	/* TrafficClasses.

	   array of trafficClasses

	   Format: JSON
	*/
	TrafficClasses *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer prototype get app qoe traffic class metrics post customers id locations location Id appqoe traffic class stats params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeGetAppQoeTrafficClassMetricsPostCustomersIDLocationsLocationIDAppqoeTrafficClassStatsParams) WithDefaults() *CustomerPrototypeGetAppQoeTrafficClassMetricsPostCustomersIDLocationsLocationIDAppqoeTrafficClassStatsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype get app qoe traffic class metrics post customers id locations location Id appqoe traffic class stats params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeGetAppQoeTrafficClassMetricsPostCustomersIDLocationsLocationIDAppqoeTrafficClassStatsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype get app qoe traffic class metrics post customers id locations location Id appqoe traffic class stats params
func (o *CustomerPrototypeGetAppQoeTrafficClassMetricsPostCustomersIDLocationsLocationIDAppqoeTrafficClassStatsParams) WithTimeout(timeout time.Duration) *CustomerPrototypeGetAppQoeTrafficClassMetricsPostCustomersIDLocationsLocationIDAppqoeTrafficClassStatsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype get app qoe traffic class metrics post customers id locations location Id appqoe traffic class stats params
func (o *CustomerPrototypeGetAppQoeTrafficClassMetricsPostCustomersIDLocationsLocationIDAppqoeTrafficClassStatsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype get app qoe traffic class metrics post customers id locations location Id appqoe traffic class stats params
func (o *CustomerPrototypeGetAppQoeTrafficClassMetricsPostCustomersIDLocationsLocationIDAppqoeTrafficClassStatsParams) WithContext(ctx context.Context) *CustomerPrototypeGetAppQoeTrafficClassMetricsPostCustomersIDLocationsLocationIDAppqoeTrafficClassStatsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype get app qoe traffic class metrics post customers id locations location Id appqoe traffic class stats params
func (o *CustomerPrototypeGetAppQoeTrafficClassMetricsPostCustomersIDLocationsLocationIDAppqoeTrafficClassStatsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype get app qoe traffic class metrics post customers id locations location Id appqoe traffic class stats params
func (o *CustomerPrototypeGetAppQoeTrafficClassMetricsPostCustomersIDLocationsLocationIDAppqoeTrafficClassStatsParams) WithHTTPClient(client *http.Client) *CustomerPrototypeGetAppQoeTrafficClassMetricsPostCustomersIDLocationsLocationIDAppqoeTrafficClassStatsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype get app qoe traffic class metrics post customers id locations location Id appqoe traffic class stats params
func (o *CustomerPrototypeGetAppQoeTrafficClassMetricsPostCustomersIDLocationsLocationIDAppqoeTrafficClassStatsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndTime adds the endTime to the customer prototype get app qoe traffic class metrics post customers id locations location Id appqoe traffic class stats params
func (o *CustomerPrototypeGetAppQoeTrafficClassMetricsPostCustomersIDLocationsLocationIDAppqoeTrafficClassStatsParams) WithEndTime(endTime strfmt.DateTime) *CustomerPrototypeGetAppQoeTrafficClassMetricsPostCustomersIDLocationsLocationIDAppqoeTrafficClassStatsParams {
	o.SetEndTime(endTime)
	return o
}

// SetEndTime adds the endTime to the customer prototype get app qoe traffic class metrics post customers id locations location Id appqoe traffic class stats params
func (o *CustomerPrototypeGetAppQoeTrafficClassMetricsPostCustomersIDLocationsLocationIDAppqoeTrafficClassStatsParams) SetEndTime(endTime strfmt.DateTime) {
	o.EndTime = endTime
}

// WithGranularity adds the granularity to the customer prototype get app qoe traffic class metrics post customers id locations location Id appqoe traffic class stats params
func (o *CustomerPrototypeGetAppQoeTrafficClassMetricsPostCustomersIDLocationsLocationIDAppqoeTrafficClassStatsParams) WithGranularity(granularity string) *CustomerPrototypeGetAppQoeTrafficClassMetricsPostCustomersIDLocationsLocationIDAppqoeTrafficClassStatsParams {
	o.SetGranularity(granularity)
	return o
}

// SetGranularity adds the granularity to the customer prototype get app qoe traffic class metrics post customers id locations location Id appqoe traffic class stats params
func (o *CustomerPrototypeGetAppQoeTrafficClassMetricsPostCustomersIDLocationsLocationIDAppqoeTrafficClassStatsParams) SetGranularity(granularity string) {
	o.Granularity = granularity
}

// WithID adds the id to the customer prototype get app qoe traffic class metrics post customers id locations location Id appqoe traffic class stats params
func (o *CustomerPrototypeGetAppQoeTrafficClassMetricsPostCustomersIDLocationsLocationIDAppqoeTrafficClassStatsParams) WithID(id string) *CustomerPrototypeGetAppQoeTrafficClassMetricsPostCustomersIDLocationsLocationIDAppqoeTrafficClassStatsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype get app qoe traffic class metrics post customers id locations location Id appqoe traffic class stats params
func (o *CustomerPrototypeGetAppQoeTrafficClassMetricsPostCustomersIDLocationsLocationIDAppqoeTrafficClassStatsParams) SetID(id string) {
	o.ID = id
}

// WithLimit adds the limit to the customer prototype get app qoe traffic class metrics post customers id locations location Id appqoe traffic class stats params
func (o *CustomerPrototypeGetAppQoeTrafficClassMetricsPostCustomersIDLocationsLocationIDAppqoeTrafficClassStatsParams) WithLimit(limit *float64) *CustomerPrototypeGetAppQoeTrafficClassMetricsPostCustomersIDLocationsLocationIDAppqoeTrafficClassStatsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the customer prototype get app qoe traffic class metrics post customers id locations location Id appqoe traffic class stats params
func (o *CustomerPrototypeGetAppQoeTrafficClassMetricsPostCustomersIDLocationsLocationIDAppqoeTrafficClassStatsParams) SetLimit(limit *float64) {
	o.Limit = limit
}

// WithLocationID adds the locationID to the customer prototype get app qoe traffic class metrics post customers id locations location Id appqoe traffic class stats params
func (o *CustomerPrototypeGetAppQoeTrafficClassMetricsPostCustomersIDLocationsLocationIDAppqoeTrafficClassStatsParams) WithLocationID(locationID string) *CustomerPrototypeGetAppQoeTrafficClassMetricsPostCustomersIDLocationsLocationIDAppqoeTrafficClassStatsParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype get app qoe traffic class metrics post customers id locations location Id appqoe traffic class stats params
func (o *CustomerPrototypeGetAppQoeTrafficClassMetricsPostCustomersIDLocationsLocationIDAppqoeTrafficClassStatsParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WithMacs adds the macs to the customer prototype get app qoe traffic class metrics post customers id locations location Id appqoe traffic class stats params
func (o *CustomerPrototypeGetAppQoeTrafficClassMetricsPostCustomersIDLocationsLocationIDAppqoeTrafficClassStatsParams) WithMacs(macs *string) *CustomerPrototypeGetAppQoeTrafficClassMetricsPostCustomersIDLocationsLocationIDAppqoeTrafficClassStatsParams {
	o.SetMacs(macs)
	return o
}

// SetMacs adds the macs to the customer prototype get app qoe traffic class metrics post customers id locations location Id appqoe traffic class stats params
func (o *CustomerPrototypeGetAppQoeTrafficClassMetricsPostCustomersIDLocationsLocationIDAppqoeTrafficClassStatsParams) SetMacs(macs *string) {
	o.Macs = macs
}

// WithStartTime adds the startTime to the customer prototype get app qoe traffic class metrics post customers id locations location Id appqoe traffic class stats params
func (o *CustomerPrototypeGetAppQoeTrafficClassMetricsPostCustomersIDLocationsLocationIDAppqoeTrafficClassStatsParams) WithStartTime(startTime strfmt.DateTime) *CustomerPrototypeGetAppQoeTrafficClassMetricsPostCustomersIDLocationsLocationIDAppqoeTrafficClassStatsParams {
	o.SetStartTime(startTime)
	return o
}

// SetStartTime adds the startTime to the customer prototype get app qoe traffic class metrics post customers id locations location Id appqoe traffic class stats params
func (o *CustomerPrototypeGetAppQoeTrafficClassMetricsPostCustomersIDLocationsLocationIDAppqoeTrafficClassStatsParams) SetStartTime(startTime strfmt.DateTime) {
	o.StartTime = startTime
}

// WithTrafficClasses adds the trafficClasses to the customer prototype get app qoe traffic class metrics post customers id locations location Id appqoe traffic class stats params
func (o *CustomerPrototypeGetAppQoeTrafficClassMetricsPostCustomersIDLocationsLocationIDAppqoeTrafficClassStatsParams) WithTrafficClasses(trafficClasses *string) *CustomerPrototypeGetAppQoeTrafficClassMetricsPostCustomersIDLocationsLocationIDAppqoeTrafficClassStatsParams {
	o.SetTrafficClasses(trafficClasses)
	return o
}

// SetTrafficClasses adds the trafficClasses to the customer prototype get app qoe traffic class metrics post customers id locations location Id appqoe traffic class stats params
func (o *CustomerPrototypeGetAppQoeTrafficClassMetricsPostCustomersIDLocationsLocationIDAppqoeTrafficClassStatsParams) SetTrafficClasses(trafficClasses *string) {
	o.TrafficClasses = trafficClasses
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypeGetAppQoeTrafficClassMetricsPostCustomersIDLocationsLocationIDAppqoeTrafficClassStatsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// form param endTime
	frEndTime := o.EndTime
	fEndTime := frEndTime.String()
	if fEndTime != "" {
		if err := r.SetFormParam("endTime", fEndTime); err != nil {
			return err
		}
	}

	// form param granularity
	frGranularity := o.Granularity
	fGranularity := frGranularity
	if fGranularity != "" {
		if err := r.SetFormParam("granularity", fGranularity); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.Limit != nil {

		// form param limit
		var frLimit float64
		if o.Limit != nil {
			frLimit = *o.Limit
		}
		fLimit := swag.FormatFloat64(frLimit)
		if fLimit != "" {
			if err := r.SetFormParam("limit", fLimit); err != nil {
				return err
			}
		}
	}

	// path param locationId
	if err := r.SetPathParam("locationId", o.LocationID); err != nil {
		return err
	}

	if o.Macs != nil {

		// form param macs
		var frMacs string
		if o.Macs != nil {
			frMacs = *o.Macs
		}
		fMacs := frMacs
		if fMacs != "" {
			if err := r.SetFormParam("macs", fMacs); err != nil {
				return err
			}
		}
	}

	// form param startTime
	frStartTime := o.StartTime
	fStartTime := frStartTime.String()
	if fStartTime != "" {
		if err := r.SetFormParam("startTime", fStartTime); err != nil {
			return err
		}
	}

	if o.TrafficClasses != nil {

		// form param trafficClasses
		var frTrafficClasses string
		if o.TrafficClasses != nil {
			frTrafficClasses = *o.TrafficClasses
		}
		fTrafficClasses := frTrafficClasses
		if fTrafficClasses != "" {
			if err := r.SetFormParam("trafficClasses", fTrafficClasses); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
