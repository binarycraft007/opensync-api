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

// NewCustomerPrototypeDeleteNodeLockedParams creates a new CustomerPrototypeDeleteNodeLockedParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypeDeleteNodeLockedParams() *CustomerPrototypeDeleteNodeLockedParams {
	return &CustomerPrototypeDeleteNodeLockedParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypeDeleteNodeLockedParamsWithTimeout creates a new CustomerPrototypeDeleteNodeLockedParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypeDeleteNodeLockedParamsWithTimeout(timeout time.Duration) *CustomerPrototypeDeleteNodeLockedParams {
	return &CustomerPrototypeDeleteNodeLockedParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypeDeleteNodeLockedParamsWithContext creates a new CustomerPrototypeDeleteNodeLockedParams object
// with the ability to set a context for a request.
func NewCustomerPrototypeDeleteNodeLockedParamsWithContext(ctx context.Context) *CustomerPrototypeDeleteNodeLockedParams {
	return &CustomerPrototypeDeleteNodeLockedParams{
		Context: ctx,
	}
}

// NewCustomerPrototypeDeleteNodeLockedParamsWithHTTPClient creates a new CustomerPrototypeDeleteNodeLockedParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypeDeleteNodeLockedParamsWithHTTPClient(client *http.Client) *CustomerPrototypeDeleteNodeLockedParams {
	return &CustomerPrototypeDeleteNodeLockedParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypeDeleteNodeLockedParams contains all the parameters to send to the API endpoint

	for the customer prototype delete node locked operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypeDeleteNodeLockedParams struct {

	/* ID.

	   Customer id

	   Format: JSON
	*/
	ID string

	// NodeID.
	NodeID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer prototype delete node locked params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeDeleteNodeLockedParams) WithDefaults() *CustomerPrototypeDeleteNodeLockedParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype delete node locked params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeDeleteNodeLockedParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype delete node locked params
func (o *CustomerPrototypeDeleteNodeLockedParams) WithTimeout(timeout time.Duration) *CustomerPrototypeDeleteNodeLockedParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype delete node locked params
func (o *CustomerPrototypeDeleteNodeLockedParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype delete node locked params
func (o *CustomerPrototypeDeleteNodeLockedParams) WithContext(ctx context.Context) *CustomerPrototypeDeleteNodeLockedParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype delete node locked params
func (o *CustomerPrototypeDeleteNodeLockedParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype delete node locked params
func (o *CustomerPrototypeDeleteNodeLockedParams) WithHTTPClient(client *http.Client) *CustomerPrototypeDeleteNodeLockedParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype delete node locked params
func (o *CustomerPrototypeDeleteNodeLockedParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the customer prototype delete node locked params
func (o *CustomerPrototypeDeleteNodeLockedParams) WithID(id string) *CustomerPrototypeDeleteNodeLockedParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype delete node locked params
func (o *CustomerPrototypeDeleteNodeLockedParams) SetID(id string) {
	o.ID = id
}

// WithNodeID adds the nodeID to the customer prototype delete node locked params
func (o *CustomerPrototypeDeleteNodeLockedParams) WithNodeID(nodeID string) *CustomerPrototypeDeleteNodeLockedParams {
	o.SetNodeID(nodeID)
	return o
}

// SetNodeID adds the nodeId to the customer prototype delete node locked params
func (o *CustomerPrototypeDeleteNodeLockedParams) SetNodeID(nodeID string) {
	o.NodeID = nodeID
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypeDeleteNodeLockedParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
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
