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

// NewCustomerResendEmailVerificationParams creates a new CustomerResendEmailVerificationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerResendEmailVerificationParams() *CustomerResendEmailVerificationParams {
	return &CustomerResendEmailVerificationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerResendEmailVerificationParamsWithTimeout creates a new CustomerResendEmailVerificationParams object
// with the ability to set a timeout on a request.
func NewCustomerResendEmailVerificationParamsWithTimeout(timeout time.Duration) *CustomerResendEmailVerificationParams {
	return &CustomerResendEmailVerificationParams{
		timeout: timeout,
	}
}

// NewCustomerResendEmailVerificationParamsWithContext creates a new CustomerResendEmailVerificationParams object
// with the ability to set a context for a request.
func NewCustomerResendEmailVerificationParamsWithContext(ctx context.Context) *CustomerResendEmailVerificationParams {
	return &CustomerResendEmailVerificationParams{
		Context: ctx,
	}
}

// NewCustomerResendEmailVerificationParamsWithHTTPClient creates a new CustomerResendEmailVerificationParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerResendEmailVerificationParamsWithHTTPClient(client *http.Client) *CustomerResendEmailVerificationParams {
	return &CustomerResendEmailVerificationParams{
		HTTPClient: client,
	}
}

/*
CustomerResendEmailVerificationParams contains all the parameters to send to the API endpoint

	for the customer resend email verification operation.

	Typically these are written to a http.Request.
*/
type CustomerResendEmailVerificationParams struct {

	/* Email.

	   Email address that verification email will be sent to.
	*/
	Email *string

	// NotificationOptions.
	//
	// Format: JSON
	NotificationOptions *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer resend email verification params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerResendEmailVerificationParams) WithDefaults() *CustomerResendEmailVerificationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer resend email verification params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerResendEmailVerificationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer resend email verification params
func (o *CustomerResendEmailVerificationParams) WithTimeout(timeout time.Duration) *CustomerResendEmailVerificationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer resend email verification params
func (o *CustomerResendEmailVerificationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer resend email verification params
func (o *CustomerResendEmailVerificationParams) WithContext(ctx context.Context) *CustomerResendEmailVerificationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer resend email verification params
func (o *CustomerResendEmailVerificationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer resend email verification params
func (o *CustomerResendEmailVerificationParams) WithHTTPClient(client *http.Client) *CustomerResendEmailVerificationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer resend email verification params
func (o *CustomerResendEmailVerificationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEmail adds the email to the customer resend email verification params
func (o *CustomerResendEmailVerificationParams) WithEmail(email *string) *CustomerResendEmailVerificationParams {
	o.SetEmail(email)
	return o
}

// SetEmail adds the email to the customer resend email verification params
func (o *CustomerResendEmailVerificationParams) SetEmail(email *string) {
	o.Email = email
}

// WithNotificationOptions adds the notificationOptions to the customer resend email verification params
func (o *CustomerResendEmailVerificationParams) WithNotificationOptions(notificationOptions *string) *CustomerResendEmailVerificationParams {
	o.SetNotificationOptions(notificationOptions)
	return o
}

// SetNotificationOptions adds the notificationOptions to the customer resend email verification params
func (o *CustomerResendEmailVerificationParams) SetNotificationOptions(notificationOptions *string) {
	o.NotificationOptions = notificationOptions
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerResendEmailVerificationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Email != nil {

		// form param email
		var frEmail string
		if o.Email != nil {
			frEmail = *o.Email
		}
		fEmail := frEmail
		if fEmail != "" {
			if err := r.SetFormParam("email", fEmail); err != nil {
				return err
			}
		}
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