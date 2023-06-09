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

	"github.com/binarycraft007/opensync-api/models"
)

// NewCustomerPrototypePutDppEnrollmentsParams creates a new CustomerPrototypePutDppEnrollmentsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypePutDppEnrollmentsParams() *CustomerPrototypePutDppEnrollmentsParams {
	return &CustomerPrototypePutDppEnrollmentsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypePutDppEnrollmentsParamsWithTimeout creates a new CustomerPrototypePutDppEnrollmentsParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypePutDppEnrollmentsParamsWithTimeout(timeout time.Duration) *CustomerPrototypePutDppEnrollmentsParams {
	return &CustomerPrototypePutDppEnrollmentsParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypePutDppEnrollmentsParamsWithContext creates a new CustomerPrototypePutDppEnrollmentsParams object
// with the ability to set a context for a request.
func NewCustomerPrototypePutDppEnrollmentsParamsWithContext(ctx context.Context) *CustomerPrototypePutDppEnrollmentsParams {
	return &CustomerPrototypePutDppEnrollmentsParams{
		Context: ctx,
	}
}

// NewCustomerPrototypePutDppEnrollmentsParamsWithHTTPClient creates a new CustomerPrototypePutDppEnrollmentsParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypePutDppEnrollmentsParamsWithHTTPClient(client *http.Client) *CustomerPrototypePutDppEnrollmentsParams {
	return &CustomerPrototypePutDppEnrollmentsParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypePutDppEnrollmentsParams contains all the parameters to send to the API endpoint

	for the customer prototype put dpp enrollments operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypePutDppEnrollmentsParams struct {

	// Enrollments.
	Enrollments []models.XAny

	/* ID.

	   Customer id

	   Format: JSON
	*/
	ID string

	// LocationID.
	//
	// Format: JSON
	LocationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer prototype put dpp enrollments params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePutDppEnrollmentsParams) WithDefaults() *CustomerPrototypePutDppEnrollmentsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype put dpp enrollments params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePutDppEnrollmentsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype put dpp enrollments params
func (o *CustomerPrototypePutDppEnrollmentsParams) WithTimeout(timeout time.Duration) *CustomerPrototypePutDppEnrollmentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype put dpp enrollments params
func (o *CustomerPrototypePutDppEnrollmentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype put dpp enrollments params
func (o *CustomerPrototypePutDppEnrollmentsParams) WithContext(ctx context.Context) *CustomerPrototypePutDppEnrollmentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype put dpp enrollments params
func (o *CustomerPrototypePutDppEnrollmentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype put dpp enrollments params
func (o *CustomerPrototypePutDppEnrollmentsParams) WithHTTPClient(client *http.Client) *CustomerPrototypePutDppEnrollmentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype put dpp enrollments params
func (o *CustomerPrototypePutDppEnrollmentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnrollments adds the enrollments to the customer prototype put dpp enrollments params
func (o *CustomerPrototypePutDppEnrollmentsParams) WithEnrollments(enrollments []models.XAny) *CustomerPrototypePutDppEnrollmentsParams {
	o.SetEnrollments(enrollments)
	return o
}

// SetEnrollments adds the enrollments to the customer prototype put dpp enrollments params
func (o *CustomerPrototypePutDppEnrollmentsParams) SetEnrollments(enrollments []models.XAny) {
	o.Enrollments = enrollments
}

// WithID adds the id to the customer prototype put dpp enrollments params
func (o *CustomerPrototypePutDppEnrollmentsParams) WithID(id string) *CustomerPrototypePutDppEnrollmentsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype put dpp enrollments params
func (o *CustomerPrototypePutDppEnrollmentsParams) SetID(id string) {
	o.ID = id
}

// WithLocationID adds the locationID to the customer prototype put dpp enrollments params
func (o *CustomerPrototypePutDppEnrollmentsParams) WithLocationID(locationID string) *CustomerPrototypePutDppEnrollmentsParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype put dpp enrollments params
func (o *CustomerPrototypePutDppEnrollmentsParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypePutDppEnrollmentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Enrollments != nil {
		if err := r.SetBodyParam(o.Enrollments); err != nil {
			return err
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
