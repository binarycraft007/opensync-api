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

// NewLocationPrototypeDestroyByIDPendingWhitelistRequestsParams creates a new LocationPrototypeDestroyByIDPendingWhitelistRequestsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLocationPrototypeDestroyByIDPendingWhitelistRequestsParams() *LocationPrototypeDestroyByIDPendingWhitelistRequestsParams {
	return &LocationPrototypeDestroyByIDPendingWhitelistRequestsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLocationPrototypeDestroyByIDPendingWhitelistRequestsParamsWithTimeout creates a new LocationPrototypeDestroyByIDPendingWhitelistRequestsParams object
// with the ability to set a timeout on a request.
func NewLocationPrototypeDestroyByIDPendingWhitelistRequestsParamsWithTimeout(timeout time.Duration) *LocationPrototypeDestroyByIDPendingWhitelistRequestsParams {
	return &LocationPrototypeDestroyByIDPendingWhitelistRequestsParams{
		timeout: timeout,
	}
}

// NewLocationPrototypeDestroyByIDPendingWhitelistRequestsParamsWithContext creates a new LocationPrototypeDestroyByIDPendingWhitelistRequestsParams object
// with the ability to set a context for a request.
func NewLocationPrototypeDestroyByIDPendingWhitelistRequestsParamsWithContext(ctx context.Context) *LocationPrototypeDestroyByIDPendingWhitelistRequestsParams {
	return &LocationPrototypeDestroyByIDPendingWhitelistRequestsParams{
		Context: ctx,
	}
}

// NewLocationPrototypeDestroyByIDPendingWhitelistRequestsParamsWithHTTPClient creates a new LocationPrototypeDestroyByIDPendingWhitelistRequestsParams object
// with the ability to set a custom HTTPClient for a request.
func NewLocationPrototypeDestroyByIDPendingWhitelistRequestsParamsWithHTTPClient(client *http.Client) *LocationPrototypeDestroyByIDPendingWhitelistRequestsParams {
	return &LocationPrototypeDestroyByIDPendingWhitelistRequestsParams{
		HTTPClient: client,
	}
}

/*
LocationPrototypeDestroyByIDPendingWhitelistRequestsParams contains all the parameters to send to the API endpoint

	for the location prototype destroy by Id pending whitelist requests operation.

	Typically these are written to a http.Request.
*/
type LocationPrototypeDestroyByIDPendingWhitelistRequestsParams struct {

	/* Fk.

	   Foreign key for _pendingWhitelistRequests

	   Format: JSON
	*/
	Fk string

	/* ID.

	   Location id

	   Format: JSON
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the location prototype destroy by Id pending whitelist requests params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LocationPrototypeDestroyByIDPendingWhitelistRequestsParams) WithDefaults() *LocationPrototypeDestroyByIDPendingWhitelistRequestsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the location prototype destroy by Id pending whitelist requests params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LocationPrototypeDestroyByIDPendingWhitelistRequestsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the location prototype destroy by Id pending whitelist requests params
func (o *LocationPrototypeDestroyByIDPendingWhitelistRequestsParams) WithTimeout(timeout time.Duration) *LocationPrototypeDestroyByIDPendingWhitelistRequestsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the location prototype destroy by Id pending whitelist requests params
func (o *LocationPrototypeDestroyByIDPendingWhitelistRequestsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the location prototype destroy by Id pending whitelist requests params
func (o *LocationPrototypeDestroyByIDPendingWhitelistRequestsParams) WithContext(ctx context.Context) *LocationPrototypeDestroyByIDPendingWhitelistRequestsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the location prototype destroy by Id pending whitelist requests params
func (o *LocationPrototypeDestroyByIDPendingWhitelistRequestsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the location prototype destroy by Id pending whitelist requests params
func (o *LocationPrototypeDestroyByIDPendingWhitelistRequestsParams) WithHTTPClient(client *http.Client) *LocationPrototypeDestroyByIDPendingWhitelistRequestsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the location prototype destroy by Id pending whitelist requests params
func (o *LocationPrototypeDestroyByIDPendingWhitelistRequestsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFk adds the fk to the location prototype destroy by Id pending whitelist requests params
func (o *LocationPrototypeDestroyByIDPendingWhitelistRequestsParams) WithFk(fk string) *LocationPrototypeDestroyByIDPendingWhitelistRequestsParams {
	o.SetFk(fk)
	return o
}

// SetFk adds the fk to the location prototype destroy by Id pending whitelist requests params
func (o *LocationPrototypeDestroyByIDPendingWhitelistRequestsParams) SetFk(fk string) {
	o.Fk = fk
}

// WithID adds the id to the location prototype destroy by Id pending whitelist requests params
func (o *LocationPrototypeDestroyByIDPendingWhitelistRequestsParams) WithID(id string) *LocationPrototypeDestroyByIDPendingWhitelistRequestsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the location prototype destroy by Id pending whitelist requests params
func (o *LocationPrototypeDestroyByIDPendingWhitelistRequestsParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *LocationPrototypeDestroyByIDPendingWhitelistRequestsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param fk
	if err := r.SetPathParam("fk", o.Fk); err != nil {
		return err
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