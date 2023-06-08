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

// NewCustomerPrototypePatchKvConfigsParams creates a new CustomerPrototypePatchKvConfigsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypePatchKvConfigsParams() *CustomerPrototypePatchKvConfigsParams {
	return &CustomerPrototypePatchKvConfigsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypePatchKvConfigsParamsWithTimeout creates a new CustomerPrototypePatchKvConfigsParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypePatchKvConfigsParamsWithTimeout(timeout time.Duration) *CustomerPrototypePatchKvConfigsParams {
	return &CustomerPrototypePatchKvConfigsParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypePatchKvConfigsParamsWithContext creates a new CustomerPrototypePatchKvConfigsParams object
// with the ability to set a context for a request.
func NewCustomerPrototypePatchKvConfigsParamsWithContext(ctx context.Context) *CustomerPrototypePatchKvConfigsParams {
	return &CustomerPrototypePatchKvConfigsParams{
		Context: ctx,
	}
}

// NewCustomerPrototypePatchKvConfigsParamsWithHTTPClient creates a new CustomerPrototypePatchKvConfigsParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypePatchKvConfigsParamsWithHTTPClient(client *http.Client) *CustomerPrototypePatchKvConfigsParams {
	return &CustomerPrototypePatchKvConfigsParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypePatchKvConfigsParams contains all the parameters to send to the API endpoint

	for the customer prototype patch kv configs operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypePatchKvConfigsParams struct {

	/* ID.

	   Customer id

	   Format: JSON
	*/
	ID string

	// KvConfigs.
	//
	// Format: JSON
	KvConfigs string

	// LocationID.
	//
	// Format: JSON
	LocationID string

	// NodeID.
	NodeID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer prototype patch kv configs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePatchKvConfigsParams) WithDefaults() *CustomerPrototypePatchKvConfigsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype patch kv configs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePatchKvConfigsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype patch kv configs params
func (o *CustomerPrototypePatchKvConfigsParams) WithTimeout(timeout time.Duration) *CustomerPrototypePatchKvConfigsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype patch kv configs params
func (o *CustomerPrototypePatchKvConfigsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype patch kv configs params
func (o *CustomerPrototypePatchKvConfigsParams) WithContext(ctx context.Context) *CustomerPrototypePatchKvConfigsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype patch kv configs params
func (o *CustomerPrototypePatchKvConfigsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype patch kv configs params
func (o *CustomerPrototypePatchKvConfigsParams) WithHTTPClient(client *http.Client) *CustomerPrototypePatchKvConfigsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype patch kv configs params
func (o *CustomerPrototypePatchKvConfigsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the customer prototype patch kv configs params
func (o *CustomerPrototypePatchKvConfigsParams) WithID(id string) *CustomerPrototypePatchKvConfigsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype patch kv configs params
func (o *CustomerPrototypePatchKvConfigsParams) SetID(id string) {
	o.ID = id
}

// WithKvConfigs adds the kvConfigs to the customer prototype patch kv configs params
func (o *CustomerPrototypePatchKvConfigsParams) WithKvConfigs(kvConfigs string) *CustomerPrototypePatchKvConfigsParams {
	o.SetKvConfigs(kvConfigs)
	return o
}

// SetKvConfigs adds the kvConfigs to the customer prototype patch kv configs params
func (o *CustomerPrototypePatchKvConfigsParams) SetKvConfigs(kvConfigs string) {
	o.KvConfigs = kvConfigs
}

// WithLocationID adds the locationID to the customer prototype patch kv configs params
func (o *CustomerPrototypePatchKvConfigsParams) WithLocationID(locationID string) *CustomerPrototypePatchKvConfigsParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype patch kv configs params
func (o *CustomerPrototypePatchKvConfigsParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WithNodeID adds the nodeID to the customer prototype patch kv configs params
func (o *CustomerPrototypePatchKvConfigsParams) WithNodeID(nodeID string) *CustomerPrototypePatchKvConfigsParams {
	o.SetNodeID(nodeID)
	return o
}

// SetNodeID adds the nodeId to the customer prototype patch kv configs params
func (o *CustomerPrototypePatchKvConfigsParams) SetNodeID(nodeID string) {
	o.NodeID = nodeID
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypePatchKvConfigsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// form param kvConfigs
	frKvConfigs := o.KvConfigs
	fKvConfigs := frKvConfigs
	if fKvConfigs != "" {
		if err := r.SetFormParam("kvConfigs", fKvConfigs); err != nil {
			return err
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
