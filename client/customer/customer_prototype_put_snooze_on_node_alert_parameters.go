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

// NewCustomerPrototypePutSnoozeOnNodeAlertParams creates a new CustomerPrototypePutSnoozeOnNodeAlertParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypePutSnoozeOnNodeAlertParams() *CustomerPrototypePutSnoozeOnNodeAlertParams {
	return &CustomerPrototypePutSnoozeOnNodeAlertParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypePutSnoozeOnNodeAlertParamsWithTimeout creates a new CustomerPrototypePutSnoozeOnNodeAlertParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypePutSnoozeOnNodeAlertParamsWithTimeout(timeout time.Duration) *CustomerPrototypePutSnoozeOnNodeAlertParams {
	return &CustomerPrototypePutSnoozeOnNodeAlertParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypePutSnoozeOnNodeAlertParamsWithContext creates a new CustomerPrototypePutSnoozeOnNodeAlertParams object
// with the ability to set a context for a request.
func NewCustomerPrototypePutSnoozeOnNodeAlertParamsWithContext(ctx context.Context) *CustomerPrototypePutSnoozeOnNodeAlertParams {
	return &CustomerPrototypePutSnoozeOnNodeAlertParams{
		Context: ctx,
	}
}

// NewCustomerPrototypePutSnoozeOnNodeAlertParamsWithHTTPClient creates a new CustomerPrototypePutSnoozeOnNodeAlertParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypePutSnoozeOnNodeAlertParamsWithHTTPClient(client *http.Client) *CustomerPrototypePutSnoozeOnNodeAlertParams {
	return &CustomerPrototypePutSnoozeOnNodeAlertParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypePutSnoozeOnNodeAlertParams contains all the parameters to send to the API endpoint

	for the customer prototype put snooze on node alert operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypePutSnoozeOnNodeAlertParams struct {

	/* ID.

	   Customer id

	   Format: JSON
	*/
	ID string

	// LocationID.
	LocationID string

	/* NodeID.

	   id of node
	*/
	NodeID string

	/* State.

	   enum of values include: snooze, ignore, performanceAcceptable, reset
	*/
	State *string

	/* Type.

	   enum of values include: poorHealth
	*/
	Type string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer prototype put snooze on node alert params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePutSnoozeOnNodeAlertParams) WithDefaults() *CustomerPrototypePutSnoozeOnNodeAlertParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype put snooze on node alert params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePutSnoozeOnNodeAlertParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype put snooze on node alert params
func (o *CustomerPrototypePutSnoozeOnNodeAlertParams) WithTimeout(timeout time.Duration) *CustomerPrototypePutSnoozeOnNodeAlertParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype put snooze on node alert params
func (o *CustomerPrototypePutSnoozeOnNodeAlertParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype put snooze on node alert params
func (o *CustomerPrototypePutSnoozeOnNodeAlertParams) WithContext(ctx context.Context) *CustomerPrototypePutSnoozeOnNodeAlertParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype put snooze on node alert params
func (o *CustomerPrototypePutSnoozeOnNodeAlertParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype put snooze on node alert params
func (o *CustomerPrototypePutSnoozeOnNodeAlertParams) WithHTTPClient(client *http.Client) *CustomerPrototypePutSnoozeOnNodeAlertParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype put snooze on node alert params
func (o *CustomerPrototypePutSnoozeOnNodeAlertParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the customer prototype put snooze on node alert params
func (o *CustomerPrototypePutSnoozeOnNodeAlertParams) WithID(id string) *CustomerPrototypePutSnoozeOnNodeAlertParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype put snooze on node alert params
func (o *CustomerPrototypePutSnoozeOnNodeAlertParams) SetID(id string) {
	o.ID = id
}

// WithLocationID adds the locationID to the customer prototype put snooze on node alert params
func (o *CustomerPrototypePutSnoozeOnNodeAlertParams) WithLocationID(locationID string) *CustomerPrototypePutSnoozeOnNodeAlertParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype put snooze on node alert params
func (o *CustomerPrototypePutSnoozeOnNodeAlertParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WithNodeID adds the nodeID to the customer prototype put snooze on node alert params
func (o *CustomerPrototypePutSnoozeOnNodeAlertParams) WithNodeID(nodeID string) *CustomerPrototypePutSnoozeOnNodeAlertParams {
	o.SetNodeID(nodeID)
	return o
}

// SetNodeID adds the nodeId to the customer prototype put snooze on node alert params
func (o *CustomerPrototypePutSnoozeOnNodeAlertParams) SetNodeID(nodeID string) {
	o.NodeID = nodeID
}

// WithState adds the state to the customer prototype put snooze on node alert params
func (o *CustomerPrototypePutSnoozeOnNodeAlertParams) WithState(state *string) *CustomerPrototypePutSnoozeOnNodeAlertParams {
	o.SetState(state)
	return o
}

// SetState adds the state to the customer prototype put snooze on node alert params
func (o *CustomerPrototypePutSnoozeOnNodeAlertParams) SetState(state *string) {
	o.State = state
}

// WithType adds the typeVar to the customer prototype put snooze on node alert params
func (o *CustomerPrototypePutSnoozeOnNodeAlertParams) WithType(typeVar string) *CustomerPrototypePutSnoozeOnNodeAlertParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the customer prototype put snooze on node alert params
func (o *CustomerPrototypePutSnoozeOnNodeAlertParams) SetType(typeVar string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypePutSnoozeOnNodeAlertParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.State != nil {

		// form param state
		var frState string
		if o.State != nil {
			frState = *o.State
		}
		fState := frState
		if fState != "" {
			if err := r.SetFormParam("state", fState); err != nil {
				return err
			}
		}
	}

	// path param type
	if err := r.SetPathParam("type", o.Type); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}