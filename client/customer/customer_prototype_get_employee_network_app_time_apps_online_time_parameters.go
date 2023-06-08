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

// NewCustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeParams creates a new CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeParams() *CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeParams {
	return &CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeParamsWithTimeout creates a new CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeParamsWithTimeout(timeout time.Duration) *CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeParams {
	return &CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeParamsWithContext creates a new CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeParams object
// with the ability to set a context for a request.
func NewCustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeParamsWithContext(ctx context.Context) *CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeParams {
	return &CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeParams{
		Context: ctx,
	}
}

// NewCustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeParamsWithHTTPClient creates a new CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeParamsWithHTTPClient(client *http.Client) *CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeParams {
	return &CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeParams contains all the parameters to send to the API endpoint

	for the customer prototype get employee network app time apps online time operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeParams struct {

	/* Grouping.

	   typing of Grouping for the purposes of applying the limit. Can be: 'perTimeSlot' ONLY

	   Default: "perTimeSlot"
	*/
	Grouping *string

	/* ID.

	   Customer id

	   Format: JSON
	*/
	ID string

	/* Limit.

	   Maximum number of apps to return. Defaults to 20

	   Format: double
	*/
	Limit *float64

	// LocationID.
	//
	// Format: JSON
	LocationID string

	// NetworkID.
	NetworkID string

	/* TimePeriod.

	   Any of "weekly","dailyToday","dailyYesterday","daily2DaysAgo","daily3DaysAgo","daily4DaysAgo","daily5DaysAgo","daily6DaysAgo","last30Days","last12Months"
	*/
	TimePeriod string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer prototype get employee network app time apps online time params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeParams) WithDefaults() *CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype get employee network app time apps online time params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeParams) SetDefaults() {
	var (
		groupingDefault = string("perTimeSlot")
	)

	val := CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeParams{
		Grouping: &groupingDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the customer prototype get employee network app time apps online time params
func (o *CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeParams) WithTimeout(timeout time.Duration) *CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype get employee network app time apps online time params
func (o *CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype get employee network app time apps online time params
func (o *CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeParams) WithContext(ctx context.Context) *CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype get employee network app time apps online time params
func (o *CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype get employee network app time apps online time params
func (o *CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeParams) WithHTTPClient(client *http.Client) *CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype get employee network app time apps online time params
func (o *CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGrouping adds the grouping to the customer prototype get employee network app time apps online time params
func (o *CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeParams) WithGrouping(grouping *string) *CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeParams {
	o.SetGrouping(grouping)
	return o
}

// SetGrouping adds the grouping to the customer prototype get employee network app time apps online time params
func (o *CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeParams) SetGrouping(grouping *string) {
	o.Grouping = grouping
}

// WithID adds the id to the customer prototype get employee network app time apps online time params
func (o *CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeParams) WithID(id string) *CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype get employee network app time apps online time params
func (o *CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeParams) SetID(id string) {
	o.ID = id
}

// WithLimit adds the limit to the customer prototype get employee network app time apps online time params
func (o *CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeParams) WithLimit(limit *float64) *CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the customer prototype get employee network app time apps online time params
func (o *CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeParams) SetLimit(limit *float64) {
	o.Limit = limit
}

// WithLocationID adds the locationID to the customer prototype get employee network app time apps online time params
func (o *CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeParams) WithLocationID(locationID string) *CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype get employee network app time apps online time params
func (o *CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WithNetworkID adds the networkID to the customer prototype get employee network app time apps online time params
func (o *CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeParams) WithNetworkID(networkID string) *CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the customer prototype get employee network app time apps online time params
func (o *CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithTimePeriod adds the timePeriod to the customer prototype get employee network app time apps online time params
func (o *CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeParams) WithTimePeriod(timePeriod string) *CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeParams {
	o.SetTimePeriod(timePeriod)
	return o
}

// SetTimePeriod adds the timePeriod to the customer prototype get employee network app time apps online time params
func (o *CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeParams) SetTimePeriod(timePeriod string) {
	o.TimePeriod = timePeriod
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Grouping != nil {

		// query param grouping
		var qrGrouping string

		if o.Grouping != nil {
			qrGrouping = *o.Grouping
		}
		qGrouping := qrGrouping
		if qGrouping != "" {

			if err := r.SetQueryParam("grouping", qGrouping); err != nil {
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

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	// query param timePeriod
	qrTimePeriod := o.TimePeriod
	qTimePeriod := qrTimePeriod
	if qTimePeriod != "" {

		if err := r.SetQueryParam("timePeriod", qTimePeriod); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
