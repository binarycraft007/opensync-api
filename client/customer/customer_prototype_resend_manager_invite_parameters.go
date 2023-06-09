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

// NewCustomerPrototypeResendManagerInviteParams creates a new CustomerPrototypeResendManagerInviteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypeResendManagerInviteParams() *CustomerPrototypeResendManagerInviteParams {
	return &CustomerPrototypeResendManagerInviteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypeResendManagerInviteParamsWithTimeout creates a new CustomerPrototypeResendManagerInviteParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypeResendManagerInviteParamsWithTimeout(timeout time.Duration) *CustomerPrototypeResendManagerInviteParams {
	return &CustomerPrototypeResendManagerInviteParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypeResendManagerInviteParamsWithContext creates a new CustomerPrototypeResendManagerInviteParams object
// with the ability to set a context for a request.
func NewCustomerPrototypeResendManagerInviteParamsWithContext(ctx context.Context) *CustomerPrototypeResendManagerInviteParams {
	return &CustomerPrototypeResendManagerInviteParams{
		Context: ctx,
	}
}

// NewCustomerPrototypeResendManagerInviteParamsWithHTTPClient creates a new CustomerPrototypeResendManagerInviteParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypeResendManagerInviteParamsWithHTTPClient(client *http.Client) *CustomerPrototypeResendManagerInviteParams {
	return &CustomerPrototypeResendManagerInviteParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypeResendManagerInviteParams contains all the parameters to send to the API endpoint

	for the customer prototype resend manager invite operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypeResendManagerInviteParams struct {

	/* ID.

	   Customer id

	   Format: JSON
	*/
	ID string

	// LocationID.
	LocationID string

	// ManagerID.
	ManagerID string

	// NotificationOptions.
	//
	// Format: JSON
	NotificationOptions *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer prototype resend manager invite params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeResendManagerInviteParams) WithDefaults() *CustomerPrototypeResendManagerInviteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype resend manager invite params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeResendManagerInviteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype resend manager invite params
func (o *CustomerPrototypeResendManagerInviteParams) WithTimeout(timeout time.Duration) *CustomerPrototypeResendManagerInviteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype resend manager invite params
func (o *CustomerPrototypeResendManagerInviteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype resend manager invite params
func (o *CustomerPrototypeResendManagerInviteParams) WithContext(ctx context.Context) *CustomerPrototypeResendManagerInviteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype resend manager invite params
func (o *CustomerPrototypeResendManagerInviteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype resend manager invite params
func (o *CustomerPrototypeResendManagerInviteParams) WithHTTPClient(client *http.Client) *CustomerPrototypeResendManagerInviteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype resend manager invite params
func (o *CustomerPrototypeResendManagerInviteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the customer prototype resend manager invite params
func (o *CustomerPrototypeResendManagerInviteParams) WithID(id string) *CustomerPrototypeResendManagerInviteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype resend manager invite params
func (o *CustomerPrototypeResendManagerInviteParams) SetID(id string) {
	o.ID = id
}

// WithLocationID adds the locationID to the customer prototype resend manager invite params
func (o *CustomerPrototypeResendManagerInviteParams) WithLocationID(locationID string) *CustomerPrototypeResendManagerInviteParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype resend manager invite params
func (o *CustomerPrototypeResendManagerInviteParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WithManagerID adds the managerID to the customer prototype resend manager invite params
func (o *CustomerPrototypeResendManagerInviteParams) WithManagerID(managerID string) *CustomerPrototypeResendManagerInviteParams {
	o.SetManagerID(managerID)
	return o
}

// SetManagerID adds the managerId to the customer prototype resend manager invite params
func (o *CustomerPrototypeResendManagerInviteParams) SetManagerID(managerID string) {
	o.ManagerID = managerID
}

// WithNotificationOptions adds the notificationOptions to the customer prototype resend manager invite params
func (o *CustomerPrototypeResendManagerInviteParams) WithNotificationOptions(notificationOptions *string) *CustomerPrototypeResendManagerInviteParams {
	o.SetNotificationOptions(notificationOptions)
	return o
}

// SetNotificationOptions adds the notificationOptions to the customer prototype resend manager invite params
func (o *CustomerPrototypeResendManagerInviteParams) SetNotificationOptions(notificationOptions *string) {
	o.NotificationOptions = notificationOptions
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypeResendManagerInviteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param managerId
	if err := r.SetPathParam("managerId", o.ManagerID); err != nil {
		return err
	}

	if o.NotificationOptions != nil {

		// form param notificationOptions
		var frNotificationOptions string
		if o.NotificationOptions != nil {
			frNotificationOptions = *o.NotificationOptions
		}
		fNotificationOptions := frNotificationOptions
		if fNotificationOptions != "" {
			if err := r.SetFormParam("notificationOptions", fNotificationOptions); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
