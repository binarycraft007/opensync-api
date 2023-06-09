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

// NewCustomerPrototypePutDeviceFreezeForeverParams creates a new CustomerPrototypePutDeviceFreezeForeverParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypePutDeviceFreezeForeverParams() *CustomerPrototypePutDeviceFreezeForeverParams {
	return &CustomerPrototypePutDeviceFreezeForeverParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypePutDeviceFreezeForeverParamsWithTimeout creates a new CustomerPrototypePutDeviceFreezeForeverParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypePutDeviceFreezeForeverParamsWithTimeout(timeout time.Duration) *CustomerPrototypePutDeviceFreezeForeverParams {
	return &CustomerPrototypePutDeviceFreezeForeverParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypePutDeviceFreezeForeverParamsWithContext creates a new CustomerPrototypePutDeviceFreezeForeverParams object
// with the ability to set a context for a request.
func NewCustomerPrototypePutDeviceFreezeForeverParamsWithContext(ctx context.Context) *CustomerPrototypePutDeviceFreezeForeverParams {
	return &CustomerPrototypePutDeviceFreezeForeverParams{
		Context: ctx,
	}
}

// NewCustomerPrototypePutDeviceFreezeForeverParamsWithHTTPClient creates a new CustomerPrototypePutDeviceFreezeForeverParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypePutDeviceFreezeForeverParamsWithHTTPClient(client *http.Client) *CustomerPrototypePutDeviceFreezeForeverParams {
	return &CustomerPrototypePutDeviceFreezeForeverParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypePutDeviceFreezeForeverParams contains all the parameters to send to the API endpoint

	for the customer prototype put device freeze forever operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypePutDeviceFreezeForeverParams struct {

	// DeleteAllExceptSuspend.
	DeleteAllExceptSuspend *bool

	// Enable.
	Enable *bool

	/* ID.

	   Customer id

	   Format: JSON
	*/
	ID string

	// LocationID.
	LocationID string

	// Mac.
	Mac string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer prototype put device freeze forever params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePutDeviceFreezeForeverParams) WithDefaults() *CustomerPrototypePutDeviceFreezeForeverParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype put device freeze forever params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePutDeviceFreezeForeverParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype put device freeze forever params
func (o *CustomerPrototypePutDeviceFreezeForeverParams) WithTimeout(timeout time.Duration) *CustomerPrototypePutDeviceFreezeForeverParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype put device freeze forever params
func (o *CustomerPrototypePutDeviceFreezeForeverParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype put device freeze forever params
func (o *CustomerPrototypePutDeviceFreezeForeverParams) WithContext(ctx context.Context) *CustomerPrototypePutDeviceFreezeForeverParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype put device freeze forever params
func (o *CustomerPrototypePutDeviceFreezeForeverParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype put device freeze forever params
func (o *CustomerPrototypePutDeviceFreezeForeverParams) WithHTTPClient(client *http.Client) *CustomerPrototypePutDeviceFreezeForeverParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype put device freeze forever params
func (o *CustomerPrototypePutDeviceFreezeForeverParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeleteAllExceptSuspend adds the deleteAllExceptSuspend to the customer prototype put device freeze forever params
func (o *CustomerPrototypePutDeviceFreezeForeverParams) WithDeleteAllExceptSuspend(deleteAllExceptSuspend *bool) *CustomerPrototypePutDeviceFreezeForeverParams {
	o.SetDeleteAllExceptSuspend(deleteAllExceptSuspend)
	return o
}

// SetDeleteAllExceptSuspend adds the deleteAllExceptSuspend to the customer prototype put device freeze forever params
func (o *CustomerPrototypePutDeviceFreezeForeverParams) SetDeleteAllExceptSuspend(deleteAllExceptSuspend *bool) {
	o.DeleteAllExceptSuspend = deleteAllExceptSuspend
}

// WithEnable adds the enable to the customer prototype put device freeze forever params
func (o *CustomerPrototypePutDeviceFreezeForeverParams) WithEnable(enable *bool) *CustomerPrototypePutDeviceFreezeForeverParams {
	o.SetEnable(enable)
	return o
}

// SetEnable adds the enable to the customer prototype put device freeze forever params
func (o *CustomerPrototypePutDeviceFreezeForeverParams) SetEnable(enable *bool) {
	o.Enable = enable
}

// WithID adds the id to the customer prototype put device freeze forever params
func (o *CustomerPrototypePutDeviceFreezeForeverParams) WithID(id string) *CustomerPrototypePutDeviceFreezeForeverParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype put device freeze forever params
func (o *CustomerPrototypePutDeviceFreezeForeverParams) SetID(id string) {
	o.ID = id
}

// WithLocationID adds the locationID to the customer prototype put device freeze forever params
func (o *CustomerPrototypePutDeviceFreezeForeverParams) WithLocationID(locationID string) *CustomerPrototypePutDeviceFreezeForeverParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype put device freeze forever params
func (o *CustomerPrototypePutDeviceFreezeForeverParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WithMac adds the mac to the customer prototype put device freeze forever params
func (o *CustomerPrototypePutDeviceFreezeForeverParams) WithMac(mac string) *CustomerPrototypePutDeviceFreezeForeverParams {
	o.SetMac(mac)
	return o
}

// SetMac adds the mac to the customer prototype put device freeze forever params
func (o *CustomerPrototypePutDeviceFreezeForeverParams) SetMac(mac string) {
	o.Mac = mac
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypePutDeviceFreezeForeverParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DeleteAllExceptSuspend != nil {

		// form param deleteAllExceptSuspend
		var frDeleteAllExceptSuspend bool
		if o.DeleteAllExceptSuspend != nil {
			frDeleteAllExceptSuspend = *o.DeleteAllExceptSuspend
		}
		fDeleteAllExceptSuspend := swag.FormatBool(frDeleteAllExceptSuspend)
		if fDeleteAllExceptSuspend != "" {
			if err := r.SetFormParam("deleteAllExceptSuspend", fDeleteAllExceptSuspend); err != nil {
				return err
			}
		}
	}

	if o.Enable != nil {

		// form param enable
		var frEnable bool
		if o.Enable != nil {
			frEnable = *o.Enable
		}
		fEnable := swag.FormatBool(frEnable)
		if fEnable != "" {
			if err := r.SetFormParam("enable", fEnable); err != nil {
				return err
			}
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param locationId
	if err := r.SetPathParam("locationId", o.LocationID); err != nil {
		return err
	}

	// path param mac
	if err := r.SetPathParam("mac", o.Mac); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
