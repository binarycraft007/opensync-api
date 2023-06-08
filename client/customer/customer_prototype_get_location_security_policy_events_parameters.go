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

// NewCustomerPrototypeGetLocationSecurityPolicyEventsParams creates a new CustomerPrototypeGetLocationSecurityPolicyEventsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypeGetLocationSecurityPolicyEventsParams() *CustomerPrototypeGetLocationSecurityPolicyEventsParams {
	return &CustomerPrototypeGetLocationSecurityPolicyEventsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypeGetLocationSecurityPolicyEventsParamsWithTimeout creates a new CustomerPrototypeGetLocationSecurityPolicyEventsParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypeGetLocationSecurityPolicyEventsParamsWithTimeout(timeout time.Duration) *CustomerPrototypeGetLocationSecurityPolicyEventsParams {
	return &CustomerPrototypeGetLocationSecurityPolicyEventsParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypeGetLocationSecurityPolicyEventsParamsWithContext creates a new CustomerPrototypeGetLocationSecurityPolicyEventsParams object
// with the ability to set a context for a request.
func NewCustomerPrototypeGetLocationSecurityPolicyEventsParamsWithContext(ctx context.Context) *CustomerPrototypeGetLocationSecurityPolicyEventsParams {
	return &CustomerPrototypeGetLocationSecurityPolicyEventsParams{
		Context: ctx,
	}
}

// NewCustomerPrototypeGetLocationSecurityPolicyEventsParamsWithHTTPClient creates a new CustomerPrototypeGetLocationSecurityPolicyEventsParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypeGetLocationSecurityPolicyEventsParamsWithHTTPClient(client *http.Client) *CustomerPrototypeGetLocationSecurityPolicyEventsParams {
	return &CustomerPrototypeGetLocationSecurityPolicyEventsParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypeGetLocationSecurityPolicyEventsParams contains all the parameters to send to the API endpoint

	for the customer prototype get location security policy events operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypeGetLocationSecurityPolicyEventsParams struct {

	// Direction.
	Direction *string

	/* ID.

	   Customer id

	   Format: JSON
	*/
	ID string

	// Includes.
	//
	// Format: JSON
	Includes *string

	// Limit.
	//
	// Format: double
	Limit *float64

	// LocationID.
	LocationID string

	// ProtectionType.
	//
	// Default: "ihp"
	ProtectionType *string

	// StartTime.
	//
	// Format: date-time
	StartTime strfmt.DateTime

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer prototype get location security policy events params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeGetLocationSecurityPolicyEventsParams) WithDefaults() *CustomerPrototypeGetLocationSecurityPolicyEventsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype get location security policy events params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeGetLocationSecurityPolicyEventsParams) SetDefaults() {
	var (
		protectionTypeDefault = string("ihp")
	)

	val := CustomerPrototypeGetLocationSecurityPolicyEventsParams{
		ProtectionType: &protectionTypeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the customer prototype get location security policy events params
func (o *CustomerPrototypeGetLocationSecurityPolicyEventsParams) WithTimeout(timeout time.Duration) *CustomerPrototypeGetLocationSecurityPolicyEventsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype get location security policy events params
func (o *CustomerPrototypeGetLocationSecurityPolicyEventsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype get location security policy events params
func (o *CustomerPrototypeGetLocationSecurityPolicyEventsParams) WithContext(ctx context.Context) *CustomerPrototypeGetLocationSecurityPolicyEventsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype get location security policy events params
func (o *CustomerPrototypeGetLocationSecurityPolicyEventsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype get location security policy events params
func (o *CustomerPrototypeGetLocationSecurityPolicyEventsParams) WithHTTPClient(client *http.Client) *CustomerPrototypeGetLocationSecurityPolicyEventsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype get location security policy events params
func (o *CustomerPrototypeGetLocationSecurityPolicyEventsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDirection adds the direction to the customer prototype get location security policy events params
func (o *CustomerPrototypeGetLocationSecurityPolicyEventsParams) WithDirection(direction *string) *CustomerPrototypeGetLocationSecurityPolicyEventsParams {
	o.SetDirection(direction)
	return o
}

// SetDirection adds the direction to the customer prototype get location security policy events params
func (o *CustomerPrototypeGetLocationSecurityPolicyEventsParams) SetDirection(direction *string) {
	o.Direction = direction
}

// WithID adds the id to the customer prototype get location security policy events params
func (o *CustomerPrototypeGetLocationSecurityPolicyEventsParams) WithID(id string) *CustomerPrototypeGetLocationSecurityPolicyEventsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype get location security policy events params
func (o *CustomerPrototypeGetLocationSecurityPolicyEventsParams) SetID(id string) {
	o.ID = id
}

// WithIncludes adds the includes to the customer prototype get location security policy events params
func (o *CustomerPrototypeGetLocationSecurityPolicyEventsParams) WithIncludes(includes *string) *CustomerPrototypeGetLocationSecurityPolicyEventsParams {
	o.SetIncludes(includes)
	return o
}

// SetIncludes adds the includes to the customer prototype get location security policy events params
func (o *CustomerPrototypeGetLocationSecurityPolicyEventsParams) SetIncludes(includes *string) {
	o.Includes = includes
}

// WithLimit adds the limit to the customer prototype get location security policy events params
func (o *CustomerPrototypeGetLocationSecurityPolicyEventsParams) WithLimit(limit *float64) *CustomerPrototypeGetLocationSecurityPolicyEventsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the customer prototype get location security policy events params
func (o *CustomerPrototypeGetLocationSecurityPolicyEventsParams) SetLimit(limit *float64) {
	o.Limit = limit
}

// WithLocationID adds the locationID to the customer prototype get location security policy events params
func (o *CustomerPrototypeGetLocationSecurityPolicyEventsParams) WithLocationID(locationID string) *CustomerPrototypeGetLocationSecurityPolicyEventsParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype get location security policy events params
func (o *CustomerPrototypeGetLocationSecurityPolicyEventsParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WithProtectionType adds the protectionType to the customer prototype get location security policy events params
func (o *CustomerPrototypeGetLocationSecurityPolicyEventsParams) WithProtectionType(protectionType *string) *CustomerPrototypeGetLocationSecurityPolicyEventsParams {
	o.SetProtectionType(protectionType)
	return o
}

// SetProtectionType adds the protectionType to the customer prototype get location security policy events params
func (o *CustomerPrototypeGetLocationSecurityPolicyEventsParams) SetProtectionType(protectionType *string) {
	o.ProtectionType = protectionType
}

// WithStartTime adds the startTime to the customer prototype get location security policy events params
func (o *CustomerPrototypeGetLocationSecurityPolicyEventsParams) WithStartTime(startTime strfmt.DateTime) *CustomerPrototypeGetLocationSecurityPolicyEventsParams {
	o.SetStartTime(startTime)
	return o
}

// SetStartTime adds the startTime to the customer prototype get location security policy events params
func (o *CustomerPrototypeGetLocationSecurityPolicyEventsParams) SetStartTime(startTime strfmt.DateTime) {
	o.StartTime = startTime
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypeGetLocationSecurityPolicyEventsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Direction != nil {

		// query param direction
		var qrDirection string

		if o.Direction != nil {
			qrDirection = *o.Direction
		}
		qDirection := qrDirection
		if qDirection != "" {

			if err := r.SetQueryParam("direction", qDirection); err != nil {
				return err
			}
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.Includes != nil {

		// query param includes
		var qrIncludes string

		if o.Includes != nil {
			qrIncludes = *o.Includes
		}
		qIncludes := qrIncludes
		if qIncludes != "" {

			if err := r.SetQueryParam("includes", qIncludes); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit float64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatFloat64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	// path param locationId
	if err := r.SetPathParam("locationId", o.LocationID); err != nil {
		return err
	}

	if o.ProtectionType != nil {

		// query param protectionType
		var qrProtectionType string

		if o.ProtectionType != nil {
			qrProtectionType = *o.ProtectionType
		}
		qProtectionType := qrProtectionType
		if qProtectionType != "" {

			if err := r.SetQueryParam("protectionType", qProtectionType); err != nil {
				return err
			}
		}
	}

	// query param startTime
	qrStartTime := o.StartTime
	qStartTime := qrStartTime.String()
	if qStartTime != "" {

		if err := r.SetQueryParam("startTime", qStartTime); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}