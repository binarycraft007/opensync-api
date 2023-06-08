// Code generated by go-swagger; DO NOT EDIT.

package group

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

// NewGroupPatchNodesByIDParams creates a new GroupPatchNodesByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGroupPatchNodesByIDParams() *GroupPatchNodesByIDParams {
	return &GroupPatchNodesByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGroupPatchNodesByIDParamsWithTimeout creates a new GroupPatchNodesByIDParams object
// with the ability to set a timeout on a request.
func NewGroupPatchNodesByIDParamsWithTimeout(timeout time.Duration) *GroupPatchNodesByIDParams {
	return &GroupPatchNodesByIDParams{
		timeout: timeout,
	}
}

// NewGroupPatchNodesByIDParamsWithContext creates a new GroupPatchNodesByIDParams object
// with the ability to set a context for a request.
func NewGroupPatchNodesByIDParamsWithContext(ctx context.Context) *GroupPatchNodesByIDParams {
	return &GroupPatchNodesByIDParams{
		Context: ctx,
	}
}

// NewGroupPatchNodesByIDParamsWithHTTPClient creates a new GroupPatchNodesByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGroupPatchNodesByIDParamsWithHTTPClient(client *http.Client) *GroupPatchNodesByIDParams {
	return &GroupPatchNodesByIDParams{
		HTTPClient: client,
	}
}

/*
GroupPatchNodesByIDParams contains all the parameters to send to the API endpoint

	for the group patch nodes by Id operation.

	Typically these are written to a http.Request.
*/
type GroupPatchNodesByIDParams struct {

	/* AccountID.

	   accountId
	*/
	AccountID *string

	/* ID.

	   node Id
	*/
	ID string

	/* Unclaimable.

	   unclaimable
	*/
	Unclaimable *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the group patch nodes by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GroupPatchNodesByIDParams) WithDefaults() *GroupPatchNodesByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the group patch nodes by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GroupPatchNodesByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the group patch nodes by Id params
func (o *GroupPatchNodesByIDParams) WithTimeout(timeout time.Duration) *GroupPatchNodesByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the group patch nodes by Id params
func (o *GroupPatchNodesByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the group patch nodes by Id params
func (o *GroupPatchNodesByIDParams) WithContext(ctx context.Context) *GroupPatchNodesByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the group patch nodes by Id params
func (o *GroupPatchNodesByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the group patch nodes by Id params
func (o *GroupPatchNodesByIDParams) WithHTTPClient(client *http.Client) *GroupPatchNodesByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the group patch nodes by Id params
func (o *GroupPatchNodesByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the group patch nodes by Id params
func (o *GroupPatchNodesByIDParams) WithAccountID(accountID *string) *GroupPatchNodesByIDParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the group patch nodes by Id params
func (o *GroupPatchNodesByIDParams) SetAccountID(accountID *string) {
	o.AccountID = accountID
}

// WithID adds the id to the group patch nodes by Id params
func (o *GroupPatchNodesByIDParams) WithID(id string) *GroupPatchNodesByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the group patch nodes by Id params
func (o *GroupPatchNodesByIDParams) SetID(id string) {
	o.ID = id
}

// WithUnclaimable adds the unclaimable to the group patch nodes by Id params
func (o *GroupPatchNodesByIDParams) WithUnclaimable(unclaimable *string) *GroupPatchNodesByIDParams {
	o.SetUnclaimable(unclaimable)
	return o
}

// SetUnclaimable adds the unclaimable to the group patch nodes by Id params
func (o *GroupPatchNodesByIDParams) SetUnclaimable(unclaimable *string) {
	o.Unclaimable = unclaimable
}

// WriteToRequest writes these params to a swagger request
func (o *GroupPatchNodesByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AccountID != nil {

		// form param accountId
		var frAccountID string
		if o.AccountID != nil {
			frAccountID = *o.AccountID
		}
		fAccountID := frAccountID
		if fAccountID != "" {
			if err := r.SetFormParam("accountId", fAccountID); err != nil {
				return err
			}
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.Unclaimable != nil {

		// form param unclaimable
		var frUnclaimable string
		if o.Unclaimable != nil {
			frUnclaimable = *o.Unclaimable
		}
		fUnclaimable := frUnclaimable
		if fUnclaimable != "" {
			if err := r.SetFormParam("unclaimable", fUnclaimable); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}