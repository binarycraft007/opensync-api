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

// NewCustomerPrototypeUnclaimAllNodesParams creates a new CustomerPrototypeUnclaimAllNodesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypeUnclaimAllNodesParams() *CustomerPrototypeUnclaimAllNodesParams {
	return &CustomerPrototypeUnclaimAllNodesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypeUnclaimAllNodesParamsWithTimeout creates a new CustomerPrototypeUnclaimAllNodesParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypeUnclaimAllNodesParamsWithTimeout(timeout time.Duration) *CustomerPrototypeUnclaimAllNodesParams {
	return &CustomerPrototypeUnclaimAllNodesParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypeUnclaimAllNodesParamsWithContext creates a new CustomerPrototypeUnclaimAllNodesParams object
// with the ability to set a context for a request.
func NewCustomerPrototypeUnclaimAllNodesParamsWithContext(ctx context.Context) *CustomerPrototypeUnclaimAllNodesParams {
	return &CustomerPrototypeUnclaimAllNodesParams{
		Context: ctx,
	}
}

// NewCustomerPrototypeUnclaimAllNodesParamsWithHTTPClient creates a new CustomerPrototypeUnclaimAllNodesParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypeUnclaimAllNodesParamsWithHTTPClient(client *http.Client) *CustomerPrototypeUnclaimAllNodesParams {
	return &CustomerPrototypeUnclaimAllNodesParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypeUnclaimAllNodesParams contains all the parameters to send to the API endpoint

	for the customer prototype unclaim all nodes operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypeUnclaimAllNodesParams struct {

	/* ID.

	   Customer id

	   Format: JSON
	*/
	ID string

	// LocationID.
	//
	// Format: JSON
	LocationID string

	/* PreservePackID.

	   packId should remain the same
	*/
	PreservePackID *bool

	/* RemoveAccountID.

	   delete account ids on the inventory nodes
	*/
	RemoveAccountID *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer prototype unclaim all nodes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeUnclaimAllNodesParams) WithDefaults() *CustomerPrototypeUnclaimAllNodesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype unclaim all nodes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeUnclaimAllNodesParams) SetDefaults() {
	var (
		preservePackIDDefault = bool(false)

		removeAccountIDDefault = bool(false)
	)

	val := CustomerPrototypeUnclaimAllNodesParams{
		PreservePackID:  &preservePackIDDefault,
		RemoveAccountID: &removeAccountIDDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the customer prototype unclaim all nodes params
func (o *CustomerPrototypeUnclaimAllNodesParams) WithTimeout(timeout time.Duration) *CustomerPrototypeUnclaimAllNodesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype unclaim all nodes params
func (o *CustomerPrototypeUnclaimAllNodesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype unclaim all nodes params
func (o *CustomerPrototypeUnclaimAllNodesParams) WithContext(ctx context.Context) *CustomerPrototypeUnclaimAllNodesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype unclaim all nodes params
func (o *CustomerPrototypeUnclaimAllNodesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype unclaim all nodes params
func (o *CustomerPrototypeUnclaimAllNodesParams) WithHTTPClient(client *http.Client) *CustomerPrototypeUnclaimAllNodesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype unclaim all nodes params
func (o *CustomerPrototypeUnclaimAllNodesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the customer prototype unclaim all nodes params
func (o *CustomerPrototypeUnclaimAllNodesParams) WithID(id string) *CustomerPrototypeUnclaimAllNodesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype unclaim all nodes params
func (o *CustomerPrototypeUnclaimAllNodesParams) SetID(id string) {
	o.ID = id
}

// WithLocationID adds the locationID to the customer prototype unclaim all nodes params
func (o *CustomerPrototypeUnclaimAllNodesParams) WithLocationID(locationID string) *CustomerPrototypeUnclaimAllNodesParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype unclaim all nodes params
func (o *CustomerPrototypeUnclaimAllNodesParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WithPreservePackID adds the preservePackID to the customer prototype unclaim all nodes params
func (o *CustomerPrototypeUnclaimAllNodesParams) WithPreservePackID(preservePackID *bool) *CustomerPrototypeUnclaimAllNodesParams {
	o.SetPreservePackID(preservePackID)
	return o
}

// SetPreservePackID adds the preservePackId to the customer prototype unclaim all nodes params
func (o *CustomerPrototypeUnclaimAllNodesParams) SetPreservePackID(preservePackID *bool) {
	o.PreservePackID = preservePackID
}

// WithRemoveAccountID adds the removeAccountID to the customer prototype unclaim all nodes params
func (o *CustomerPrototypeUnclaimAllNodesParams) WithRemoveAccountID(removeAccountID *bool) *CustomerPrototypeUnclaimAllNodesParams {
	o.SetRemoveAccountID(removeAccountID)
	return o
}

// SetRemoveAccountID adds the removeAccountId to the customer prototype unclaim all nodes params
func (o *CustomerPrototypeUnclaimAllNodesParams) SetRemoveAccountID(removeAccountID *bool) {
	o.RemoveAccountID = removeAccountID
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypeUnclaimAllNodesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.PreservePackID != nil {

		// form param preservePackId
		var frPreservePackID bool
		if o.PreservePackID != nil {
			frPreservePackID = *o.PreservePackID
		}
		fPreservePackID := swag.FormatBool(frPreservePackID)
		if fPreservePackID != "" {
			if err := r.SetFormParam("preservePackId", fPreservePackID); err != nil {
				return err
			}
		}
	}

	if o.RemoveAccountID != nil {

		// form param removeAccountId
		var frRemoveAccountID bool
		if o.RemoveAccountID != nil {
			frRemoveAccountID = *o.RemoveAccountID
		}
		fRemoveAccountID := swag.FormatBool(frRemoveAccountID)
		if fRemoveAccountID != "" {
			if err := r.SetFormParam("removeAccountId", fRemoveAccountID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}