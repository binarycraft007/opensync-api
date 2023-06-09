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

// NewNodeFindByIDParams creates a new NodeFindByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNodeFindByIDParams() *NodeFindByIDParams {
	return &NodeFindByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNodeFindByIDParamsWithTimeout creates a new NodeFindByIDParams object
// with the ability to set a timeout on a request.
func NewNodeFindByIDParamsWithTimeout(timeout time.Duration) *NodeFindByIDParams {
	return &NodeFindByIDParams{
		timeout: timeout,
	}
}

// NewNodeFindByIDParamsWithContext creates a new NodeFindByIDParams object
// with the ability to set a context for a request.
func NewNodeFindByIDParamsWithContext(ctx context.Context) *NodeFindByIDParams {
	return &NodeFindByIDParams{
		Context: ctx,
	}
}

// NewNodeFindByIDParamsWithHTTPClient creates a new NodeFindByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewNodeFindByIDParamsWithHTTPClient(client *http.Client) *NodeFindByIDParams {
	return &NodeFindByIDParams{
		HTTPClient: client,
	}
}

/*
NodeFindByIDParams contains all the parameters to send to the API endpoint

	for the node find by Id operation.

	Typically these are written to a http.Request.
*/
type NodeFindByIDParams struct {

	/* Filter.

	   Filter defining fields and include - must be a JSON-encoded string ({"something":"value"})

	   Format: JSON
	*/
	Filter *string

	/* ID.

	   Model id

	   Format: JSON
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the node find by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NodeFindByIDParams) WithDefaults() *NodeFindByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the node find by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NodeFindByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the node find by Id params
func (o *NodeFindByIDParams) WithTimeout(timeout time.Duration) *NodeFindByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the node find by Id params
func (o *NodeFindByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the node find by Id params
func (o *NodeFindByIDParams) WithContext(ctx context.Context) *NodeFindByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the node find by Id params
func (o *NodeFindByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the node find by Id params
func (o *NodeFindByIDParams) WithHTTPClient(client *http.Client) *NodeFindByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the node find by Id params
func (o *NodeFindByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the node find by Id params
func (o *NodeFindByIDParams) WithFilter(filter *string) *NodeFindByIDParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the node find by Id params
func (o *NodeFindByIDParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithID adds the id to the node find by Id params
func (o *NodeFindByIDParams) WithID(id string) *NodeFindByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the node find by Id params
func (o *NodeFindByIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *NodeFindByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Filter != nil {

		// query param filter
		var qrFilter string

		if o.Filter != nil {
			qrFilter = *o.Filter
		}
		qFilter := qrFilter
		if qFilter != "" {

			if err := r.SetQueryParam("filter", qFilter); err != nil {
				return err
			}
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
