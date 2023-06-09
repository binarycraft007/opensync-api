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

// NewCustomerPrototypePostSpeedTestParams creates a new CustomerPrototypePostSpeedTestParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypePostSpeedTestParams() *CustomerPrototypePostSpeedTestParams {
	return &CustomerPrototypePostSpeedTestParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypePostSpeedTestParamsWithTimeout creates a new CustomerPrototypePostSpeedTestParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypePostSpeedTestParamsWithTimeout(timeout time.Duration) *CustomerPrototypePostSpeedTestParams {
	return &CustomerPrototypePostSpeedTestParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypePostSpeedTestParamsWithContext creates a new CustomerPrototypePostSpeedTestParams object
// with the ability to set a context for a request.
func NewCustomerPrototypePostSpeedTestParamsWithContext(ctx context.Context) *CustomerPrototypePostSpeedTestParams {
	return &CustomerPrototypePostSpeedTestParams{
		Context: ctx,
	}
}

// NewCustomerPrototypePostSpeedTestParamsWithHTTPClient creates a new CustomerPrototypePostSpeedTestParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypePostSpeedTestParamsWithHTTPClient(client *http.Client) *CustomerPrototypePostSpeedTestParams {
	return &CustomerPrototypePostSpeedTestParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypePostSpeedTestParams contains all the parameters to send to the API endpoint

	for the customer prototype post speed test operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypePostSpeedTestParams struct {

	/* ID.

	   Customer id

	   Format: JSON
	*/
	ID string

	// LocationID.
	LocationID string

	// NodeID.
	NodeID string

	// ServerID.
	//
	// Format: double
	ServerID *float64

	// TestType.
	//
	// Default: "OOKLA"
	TestType string

	// UplinkType.
	//
	// Default: "wire"
	UplinkType *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer prototype post speed test params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePostSpeedTestParams) WithDefaults() *CustomerPrototypePostSpeedTestParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype post speed test params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePostSpeedTestParams) SetDefaults() {
	var (
		testTypeDefault = string("OOKLA")

		uplinkTypeDefault = string("wire")
	)

	val := CustomerPrototypePostSpeedTestParams{
		TestType:   testTypeDefault,
		UplinkType: &uplinkTypeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the customer prototype post speed test params
func (o *CustomerPrototypePostSpeedTestParams) WithTimeout(timeout time.Duration) *CustomerPrototypePostSpeedTestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype post speed test params
func (o *CustomerPrototypePostSpeedTestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype post speed test params
func (o *CustomerPrototypePostSpeedTestParams) WithContext(ctx context.Context) *CustomerPrototypePostSpeedTestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype post speed test params
func (o *CustomerPrototypePostSpeedTestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype post speed test params
func (o *CustomerPrototypePostSpeedTestParams) WithHTTPClient(client *http.Client) *CustomerPrototypePostSpeedTestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype post speed test params
func (o *CustomerPrototypePostSpeedTestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the customer prototype post speed test params
func (o *CustomerPrototypePostSpeedTestParams) WithID(id string) *CustomerPrototypePostSpeedTestParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype post speed test params
func (o *CustomerPrototypePostSpeedTestParams) SetID(id string) {
	o.ID = id
}

// WithLocationID adds the locationID to the customer prototype post speed test params
func (o *CustomerPrototypePostSpeedTestParams) WithLocationID(locationID string) *CustomerPrototypePostSpeedTestParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype post speed test params
func (o *CustomerPrototypePostSpeedTestParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WithNodeID adds the nodeID to the customer prototype post speed test params
func (o *CustomerPrototypePostSpeedTestParams) WithNodeID(nodeID string) *CustomerPrototypePostSpeedTestParams {
	o.SetNodeID(nodeID)
	return o
}

// SetNodeID adds the nodeId to the customer prototype post speed test params
func (o *CustomerPrototypePostSpeedTestParams) SetNodeID(nodeID string) {
	o.NodeID = nodeID
}

// WithServerID adds the serverID to the customer prototype post speed test params
func (o *CustomerPrototypePostSpeedTestParams) WithServerID(serverID *float64) *CustomerPrototypePostSpeedTestParams {
	o.SetServerID(serverID)
	return o
}

// SetServerID adds the serverId to the customer prototype post speed test params
func (o *CustomerPrototypePostSpeedTestParams) SetServerID(serverID *float64) {
	o.ServerID = serverID
}

// WithTestType adds the testType to the customer prototype post speed test params
func (o *CustomerPrototypePostSpeedTestParams) WithTestType(testType string) *CustomerPrototypePostSpeedTestParams {
	o.SetTestType(testType)
	return o
}

// SetTestType adds the testType to the customer prototype post speed test params
func (o *CustomerPrototypePostSpeedTestParams) SetTestType(testType string) {
	o.TestType = testType
}

// WithUplinkType adds the uplinkType to the customer prototype post speed test params
func (o *CustomerPrototypePostSpeedTestParams) WithUplinkType(uplinkType *string) *CustomerPrototypePostSpeedTestParams {
	o.SetUplinkType(uplinkType)
	return o
}

// SetUplinkType adds the uplinkType to the customer prototype post speed test params
func (o *CustomerPrototypePostSpeedTestParams) SetUplinkType(uplinkType *string) {
	o.UplinkType = uplinkType
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypePostSpeedTestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param nodeId
	if err := r.SetPathParam("nodeId", o.NodeID); err != nil {
		return err
	}

	if o.ServerID != nil {

		// form param serverId
		var frServerID float64
		if o.ServerID != nil {
			frServerID = *o.ServerID
		}
		fServerID := swag.FormatFloat64(frServerID)
		if fServerID != "" {
			if err := r.SetFormParam("serverId", fServerID); err != nil {
				return err
			}
		}
	}

	// form param testType
	frTestType := o.TestType
	fTestType := frTestType
	if fTestType != "" {
		if err := r.SetFormParam("testType", fTestType); err != nil {
			return err
		}
	}

	if o.UplinkType != nil {

		// form param uplinkType
		var frUplinkType string
		if o.UplinkType != nil {
			frUplinkType = *o.UplinkType
		}
		fUplinkType := frUplinkType
		if fUplinkType != "" {
			if err := r.SetFormParam("uplinkType", fUplinkType); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
