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

	"github.com/binarycraft007/opensync-api/models"
)

// NewLocationPrototypeUpdateByIDPendingWhitelistRequestsParams creates a new LocationPrototypeUpdateByIDPendingWhitelistRequestsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLocationPrototypeUpdateByIDPendingWhitelistRequestsParams() *LocationPrototypeUpdateByIDPendingWhitelistRequestsParams {
	return &LocationPrototypeUpdateByIDPendingWhitelistRequestsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLocationPrototypeUpdateByIDPendingWhitelistRequestsParamsWithTimeout creates a new LocationPrototypeUpdateByIDPendingWhitelistRequestsParams object
// with the ability to set a timeout on a request.
func NewLocationPrototypeUpdateByIDPendingWhitelistRequestsParamsWithTimeout(timeout time.Duration) *LocationPrototypeUpdateByIDPendingWhitelistRequestsParams {
	return &LocationPrototypeUpdateByIDPendingWhitelistRequestsParams{
		timeout: timeout,
	}
}

// NewLocationPrototypeUpdateByIDPendingWhitelistRequestsParamsWithContext creates a new LocationPrototypeUpdateByIDPendingWhitelistRequestsParams object
// with the ability to set a context for a request.
func NewLocationPrototypeUpdateByIDPendingWhitelistRequestsParamsWithContext(ctx context.Context) *LocationPrototypeUpdateByIDPendingWhitelistRequestsParams {
	return &LocationPrototypeUpdateByIDPendingWhitelistRequestsParams{
		Context: ctx,
	}
}

// NewLocationPrototypeUpdateByIDPendingWhitelistRequestsParamsWithHTTPClient creates a new LocationPrototypeUpdateByIDPendingWhitelistRequestsParams object
// with the ability to set a custom HTTPClient for a request.
func NewLocationPrototypeUpdateByIDPendingWhitelistRequestsParamsWithHTTPClient(client *http.Client) *LocationPrototypeUpdateByIDPendingWhitelistRequestsParams {
	return &LocationPrototypeUpdateByIDPendingWhitelistRequestsParams{
		HTTPClient: client,
	}
}

/*
LocationPrototypeUpdateByIDPendingWhitelistRequestsParams contains all the parameters to send to the API endpoint

	for the location prototype update by Id pending whitelist requests operation.

	Typically these are written to a http.Request.
*/
type LocationPrototypeUpdateByIDPendingWhitelistRequestsParams struct {

	// Data.
	Data *models.PendingWhitelistRequests

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

// WithDefaults hydrates default values in the location prototype update by Id pending whitelist requests params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LocationPrototypeUpdateByIDPendingWhitelistRequestsParams) WithDefaults() *LocationPrototypeUpdateByIDPendingWhitelistRequestsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the location prototype update by Id pending whitelist requests params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LocationPrototypeUpdateByIDPendingWhitelistRequestsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the location prototype update by Id pending whitelist requests params
func (o *LocationPrototypeUpdateByIDPendingWhitelistRequestsParams) WithTimeout(timeout time.Duration) *LocationPrototypeUpdateByIDPendingWhitelistRequestsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the location prototype update by Id pending whitelist requests params
func (o *LocationPrototypeUpdateByIDPendingWhitelistRequestsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the location prototype update by Id pending whitelist requests params
func (o *LocationPrototypeUpdateByIDPendingWhitelistRequestsParams) WithContext(ctx context.Context) *LocationPrototypeUpdateByIDPendingWhitelistRequestsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the location prototype update by Id pending whitelist requests params
func (o *LocationPrototypeUpdateByIDPendingWhitelistRequestsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the location prototype update by Id pending whitelist requests params
func (o *LocationPrototypeUpdateByIDPendingWhitelistRequestsParams) WithHTTPClient(client *http.Client) *LocationPrototypeUpdateByIDPendingWhitelistRequestsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the location prototype update by Id pending whitelist requests params
func (o *LocationPrototypeUpdateByIDPendingWhitelistRequestsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the location prototype update by Id pending whitelist requests params
func (o *LocationPrototypeUpdateByIDPendingWhitelistRequestsParams) WithData(data *models.PendingWhitelistRequests) *LocationPrototypeUpdateByIDPendingWhitelistRequestsParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the location prototype update by Id pending whitelist requests params
func (o *LocationPrototypeUpdateByIDPendingWhitelistRequestsParams) SetData(data *models.PendingWhitelistRequests) {
	o.Data = data
}

// WithFk adds the fk to the location prototype update by Id pending whitelist requests params
func (o *LocationPrototypeUpdateByIDPendingWhitelistRequestsParams) WithFk(fk string) *LocationPrototypeUpdateByIDPendingWhitelistRequestsParams {
	o.SetFk(fk)
	return o
}

// SetFk adds the fk to the location prototype update by Id pending whitelist requests params
func (o *LocationPrototypeUpdateByIDPendingWhitelistRequestsParams) SetFk(fk string) {
	o.Fk = fk
}

// WithID adds the id to the location prototype update by Id pending whitelist requests params
func (o *LocationPrototypeUpdateByIDPendingWhitelistRequestsParams) WithID(id string) *LocationPrototypeUpdateByIDPendingWhitelistRequestsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the location prototype update by Id pending whitelist requests params
func (o *LocationPrototypeUpdateByIDPendingWhitelistRequestsParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *LocationPrototypeUpdateByIDPendingWhitelistRequestsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

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
