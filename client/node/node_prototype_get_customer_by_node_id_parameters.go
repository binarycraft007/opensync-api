// Code generated by go-swagger; DO NOT EDIT.

package node

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

// NewNodePrototypeGetCustomerByNodeIDParams creates a new NodePrototypeGetCustomerByNodeIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNodePrototypeGetCustomerByNodeIDParams() *NodePrototypeGetCustomerByNodeIDParams {
	return &NodePrototypeGetCustomerByNodeIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNodePrototypeGetCustomerByNodeIDParamsWithTimeout creates a new NodePrototypeGetCustomerByNodeIDParams object
// with the ability to set a timeout on a request.
func NewNodePrototypeGetCustomerByNodeIDParamsWithTimeout(timeout time.Duration) *NodePrototypeGetCustomerByNodeIDParams {
	return &NodePrototypeGetCustomerByNodeIDParams{
		timeout: timeout,
	}
}

// NewNodePrototypeGetCustomerByNodeIDParamsWithContext creates a new NodePrototypeGetCustomerByNodeIDParams object
// with the ability to set a context for a request.
func NewNodePrototypeGetCustomerByNodeIDParamsWithContext(ctx context.Context) *NodePrototypeGetCustomerByNodeIDParams {
	return &NodePrototypeGetCustomerByNodeIDParams{
		Context: ctx,
	}
}

// NewNodePrototypeGetCustomerByNodeIDParamsWithHTTPClient creates a new NodePrototypeGetCustomerByNodeIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewNodePrototypeGetCustomerByNodeIDParamsWithHTTPClient(client *http.Client) *NodePrototypeGetCustomerByNodeIDParams {
	return &NodePrototypeGetCustomerByNodeIDParams{
		HTTPClient: client,
	}
}

/*
NodePrototypeGetCustomerByNodeIDParams contains all the parameters to send to the API endpoint

	for the node prototype get customer by node Id operation.

	Typically these are written to a http.Request.
*/
type NodePrototypeGetCustomerByNodeIDParams struct {

	/* ID.

	   Node id

	   Format: JSON
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the node prototype get customer by node Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NodePrototypeGetCustomerByNodeIDParams) WithDefaults() *NodePrototypeGetCustomerByNodeIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the node prototype get customer by node Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NodePrototypeGetCustomerByNodeIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the node prototype get customer by node Id params
func (o *NodePrototypeGetCustomerByNodeIDParams) WithTimeout(timeout time.Duration) *NodePrototypeGetCustomerByNodeIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the node prototype get customer by node Id params
func (o *NodePrototypeGetCustomerByNodeIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the node prototype get customer by node Id params
func (o *NodePrototypeGetCustomerByNodeIDParams) WithContext(ctx context.Context) *NodePrototypeGetCustomerByNodeIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the node prototype get customer by node Id params
func (o *NodePrototypeGetCustomerByNodeIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the node prototype get customer by node Id params
func (o *NodePrototypeGetCustomerByNodeIDParams) WithHTTPClient(client *http.Client) *NodePrototypeGetCustomerByNodeIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the node prototype get customer by node Id params
func (o *NodePrototypeGetCustomerByNodeIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the node prototype get customer by node Id params
func (o *NodePrototypeGetCustomerByNodeIDParams) WithID(id string) *NodePrototypeGetCustomerByNodeIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the node prototype get customer by node Id params
func (o *NodePrototypeGetCustomerByNodeIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *NodePrototypeGetCustomerByNodeIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}