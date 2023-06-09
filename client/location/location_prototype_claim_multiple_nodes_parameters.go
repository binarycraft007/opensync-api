// Code generated by go-swagger; DO NOT EDIT.

package location

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

// NewLocationPrototypeClaimMultipleNodesParams creates a new LocationPrototypeClaimMultipleNodesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLocationPrototypeClaimMultipleNodesParams() *LocationPrototypeClaimMultipleNodesParams {
	return &LocationPrototypeClaimMultipleNodesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLocationPrototypeClaimMultipleNodesParamsWithTimeout creates a new LocationPrototypeClaimMultipleNodesParams object
// with the ability to set a timeout on a request.
func NewLocationPrototypeClaimMultipleNodesParamsWithTimeout(timeout time.Duration) *LocationPrototypeClaimMultipleNodesParams {
	return &LocationPrototypeClaimMultipleNodesParams{
		timeout: timeout,
	}
}

// NewLocationPrototypeClaimMultipleNodesParamsWithContext creates a new LocationPrototypeClaimMultipleNodesParams object
// with the ability to set a context for a request.
func NewLocationPrototypeClaimMultipleNodesParamsWithContext(ctx context.Context) *LocationPrototypeClaimMultipleNodesParams {
	return &LocationPrototypeClaimMultipleNodesParams{
		Context: ctx,
	}
}

// NewLocationPrototypeClaimMultipleNodesParamsWithHTTPClient creates a new LocationPrototypeClaimMultipleNodesParams object
// with the ability to set a custom HTTPClient for a request.
func NewLocationPrototypeClaimMultipleNodesParamsWithHTTPClient(client *http.Client) *LocationPrototypeClaimMultipleNodesParams {
	return &LocationPrototypeClaimMultipleNodesParams{
		HTTPClient: client,
	}
}

/*
LocationPrototypeClaimMultipleNodesParams contains all the parameters to send to the API endpoint

	for the location prototype claim multiple nodes operation.

	Typically these are written to a http.Request.
*/
type LocationPrototypeClaimMultipleNodesParams struct {

	/* ID.

	   Location id

	   Format: JSON
	*/
	ID string

	/* Nodes.

	   array of serialNumber/ids

	   Format: JSON
	*/
	Nodes *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the location prototype claim multiple nodes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LocationPrototypeClaimMultipleNodesParams) WithDefaults() *LocationPrototypeClaimMultipleNodesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the location prototype claim multiple nodes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LocationPrototypeClaimMultipleNodesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the location prototype claim multiple nodes params
func (o *LocationPrototypeClaimMultipleNodesParams) WithTimeout(timeout time.Duration) *LocationPrototypeClaimMultipleNodesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the location prototype claim multiple nodes params
func (o *LocationPrototypeClaimMultipleNodesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the location prototype claim multiple nodes params
func (o *LocationPrototypeClaimMultipleNodesParams) WithContext(ctx context.Context) *LocationPrototypeClaimMultipleNodesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the location prototype claim multiple nodes params
func (o *LocationPrototypeClaimMultipleNodesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the location prototype claim multiple nodes params
func (o *LocationPrototypeClaimMultipleNodesParams) WithHTTPClient(client *http.Client) *LocationPrototypeClaimMultipleNodesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the location prototype claim multiple nodes params
func (o *LocationPrototypeClaimMultipleNodesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the location prototype claim multiple nodes params
func (o *LocationPrototypeClaimMultipleNodesParams) WithID(id string) *LocationPrototypeClaimMultipleNodesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the location prototype claim multiple nodes params
func (o *LocationPrototypeClaimMultipleNodesParams) SetID(id string) {
	o.ID = id
}

// WithNodes adds the nodes to the location prototype claim multiple nodes params
func (o *LocationPrototypeClaimMultipleNodesParams) WithNodes(nodes *string) *LocationPrototypeClaimMultipleNodesParams {
	o.SetNodes(nodes)
	return o
}

// SetNodes adds the nodes to the location prototype claim multiple nodes params
func (o *LocationPrototypeClaimMultipleNodesParams) SetNodes(nodes *string) {
	o.Nodes = nodes
}

// WriteToRequest writes these params to a swagger request
func (o *LocationPrototypeClaimMultipleNodesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.Nodes != nil {

		// form param nodes
		var frNodes string
		if o.Nodes != nil {
			frNodes = *o.Nodes
		}
		fNodes := frNodes
		if fNodes != "" {
			if err := r.SetFormParam("nodes", fNodes); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
