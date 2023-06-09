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

// NewCustomerPrototypeGetSpeedTestResultsParams creates a new CustomerPrototypeGetSpeedTestResultsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypeGetSpeedTestResultsParams() *CustomerPrototypeGetSpeedTestResultsParams {
	return &CustomerPrototypeGetSpeedTestResultsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypeGetSpeedTestResultsParamsWithTimeout creates a new CustomerPrototypeGetSpeedTestResultsParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypeGetSpeedTestResultsParamsWithTimeout(timeout time.Duration) *CustomerPrototypeGetSpeedTestResultsParams {
	return &CustomerPrototypeGetSpeedTestResultsParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypeGetSpeedTestResultsParamsWithContext creates a new CustomerPrototypeGetSpeedTestResultsParams object
// with the ability to set a context for a request.
func NewCustomerPrototypeGetSpeedTestResultsParamsWithContext(ctx context.Context) *CustomerPrototypeGetSpeedTestResultsParams {
	return &CustomerPrototypeGetSpeedTestResultsParams{
		Context: ctx,
	}
}

// NewCustomerPrototypeGetSpeedTestResultsParamsWithHTTPClient creates a new CustomerPrototypeGetSpeedTestResultsParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypeGetSpeedTestResultsParamsWithHTTPClient(client *http.Client) *CustomerPrototypeGetSpeedTestResultsParams {
	return &CustomerPrototypeGetSpeedTestResultsParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypeGetSpeedTestResultsParams contains all the parameters to send to the API endpoint

	for the customer prototype get speed test results operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypeGetSpeedTestResultsParams struct {

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

	// LocationID.
	LocationID string

	// NodeID.
	NodeID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer prototype get speed test results params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeGetSpeedTestResultsParams) WithDefaults() *CustomerPrototypeGetSpeedTestResultsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype get speed test results params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeGetSpeedTestResultsParams) SetDefaults() {
	var (
		granularityDefault = string("days")

		limitDefault = float64(7)
	)

	val := CustomerPrototypeGetSpeedTestResultsParams{
		Granularity: &granularityDefault,
		Limit:       &limitDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the customer prototype get speed test results params
func (o *CustomerPrototypeGetSpeedTestResultsParams) WithTimeout(timeout time.Duration) *CustomerPrototypeGetSpeedTestResultsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype get speed test results params
func (o *CustomerPrototypeGetSpeedTestResultsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype get speed test results params
func (o *CustomerPrototypeGetSpeedTestResultsParams) WithContext(ctx context.Context) *CustomerPrototypeGetSpeedTestResultsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype get speed test results params
func (o *CustomerPrototypeGetSpeedTestResultsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype get speed test results params
func (o *CustomerPrototypeGetSpeedTestResultsParams) WithHTTPClient(client *http.Client) *CustomerPrototypeGetSpeedTestResultsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype get speed test results params
func (o *CustomerPrototypeGetSpeedTestResultsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGranularity adds the granularity to the customer prototype get speed test results params
func (o *CustomerPrototypeGetSpeedTestResultsParams) WithGranularity(granularity *string) *CustomerPrototypeGetSpeedTestResultsParams {
	o.SetGranularity(granularity)
	return o
}

// SetGranularity adds the granularity to the customer prototype get speed test results params
func (o *CustomerPrototypeGetSpeedTestResultsParams) SetGranularity(granularity *string) {
	o.Granularity = granularity
}

// WithID adds the id to the customer prototype get speed test results params
func (o *CustomerPrototypeGetSpeedTestResultsParams) WithID(id string) *CustomerPrototypeGetSpeedTestResultsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype get speed test results params
func (o *CustomerPrototypeGetSpeedTestResultsParams) SetID(id string) {
	o.ID = id
}

// WithLimit adds the limit to the customer prototype get speed test results params
func (o *CustomerPrototypeGetSpeedTestResultsParams) WithLimit(limit *float64) *CustomerPrototypeGetSpeedTestResultsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the customer prototype get speed test results params
func (o *CustomerPrototypeGetSpeedTestResultsParams) SetLimit(limit *float64) {
	o.Limit = limit
}

// WithLocationID adds the locationID to the customer prototype get speed test results params
func (o *CustomerPrototypeGetSpeedTestResultsParams) WithLocationID(locationID string) *CustomerPrototypeGetSpeedTestResultsParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype get speed test results params
func (o *CustomerPrototypeGetSpeedTestResultsParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WithNodeID adds the nodeID to the customer prototype get speed test results params
func (o *CustomerPrototypeGetSpeedTestResultsParams) WithNodeID(nodeID string) *CustomerPrototypeGetSpeedTestResultsParams {
	o.SetNodeID(nodeID)
	return o
}

// SetNodeID adds the nodeId to the customer prototype get speed test results params
func (o *CustomerPrototypeGetSpeedTestResultsParams) SetNodeID(nodeID string) {
	o.NodeID = nodeID
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypeGetSpeedTestResultsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param nodeId
	if err := r.SetPathParam("nodeId", o.NodeID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
