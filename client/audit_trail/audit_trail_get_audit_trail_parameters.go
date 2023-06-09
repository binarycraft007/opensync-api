// Code generated by go-swagger; DO NOT EDIT.

package audit_trail

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

// NewAuditTrailGetAuditTrailParams creates a new AuditTrailGetAuditTrailParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAuditTrailGetAuditTrailParams() *AuditTrailGetAuditTrailParams {
	return &AuditTrailGetAuditTrailParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAuditTrailGetAuditTrailParamsWithTimeout creates a new AuditTrailGetAuditTrailParams object
// with the ability to set a timeout on a request.
func NewAuditTrailGetAuditTrailParamsWithTimeout(timeout time.Duration) *AuditTrailGetAuditTrailParams {
	return &AuditTrailGetAuditTrailParams{
		timeout: timeout,
	}
}

// NewAuditTrailGetAuditTrailParamsWithContext creates a new AuditTrailGetAuditTrailParams object
// with the ability to set a context for a request.
func NewAuditTrailGetAuditTrailParamsWithContext(ctx context.Context) *AuditTrailGetAuditTrailParams {
	return &AuditTrailGetAuditTrailParams{
		Context: ctx,
	}
}

// NewAuditTrailGetAuditTrailParamsWithHTTPClient creates a new AuditTrailGetAuditTrailParams object
// with the ability to set a custom HTTPClient for a request.
func NewAuditTrailGetAuditTrailParamsWithHTTPClient(client *http.Client) *AuditTrailGetAuditTrailParams {
	return &AuditTrailGetAuditTrailParams{
		HTTPClient: client,
	}
}

/*
AuditTrailGetAuditTrailParams contains all the parameters to send to the API endpoint

	for the audit trail get audit trail operation.

	Typically these are written to a http.Request.
*/
type AuditTrailGetAuditTrailParams struct {

	/* CustomerID.

	   Customer Id
	*/
	CustomerID string

	/* LocationID.

	   Location Id
	*/
	LocationID *string

	/* PartnerIds.

	   Partner Id

	   Format: JSON
	*/
	PartnerIds *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the audit trail get audit trail params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuditTrailGetAuditTrailParams) WithDefaults() *AuditTrailGetAuditTrailParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the audit trail get audit trail params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuditTrailGetAuditTrailParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the audit trail get audit trail params
func (o *AuditTrailGetAuditTrailParams) WithTimeout(timeout time.Duration) *AuditTrailGetAuditTrailParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the audit trail get audit trail params
func (o *AuditTrailGetAuditTrailParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the audit trail get audit trail params
func (o *AuditTrailGetAuditTrailParams) WithContext(ctx context.Context) *AuditTrailGetAuditTrailParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the audit trail get audit trail params
func (o *AuditTrailGetAuditTrailParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the audit trail get audit trail params
func (o *AuditTrailGetAuditTrailParams) WithHTTPClient(client *http.Client) *AuditTrailGetAuditTrailParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the audit trail get audit trail params
func (o *AuditTrailGetAuditTrailParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerID adds the customerID to the audit trail get audit trail params
func (o *AuditTrailGetAuditTrailParams) WithCustomerID(customerID string) *AuditTrailGetAuditTrailParams {
	o.SetCustomerID(customerID)
	return o
}

// SetCustomerID adds the customerId to the audit trail get audit trail params
func (o *AuditTrailGetAuditTrailParams) SetCustomerID(customerID string) {
	o.CustomerID = customerID
}

// WithLocationID adds the locationID to the audit trail get audit trail params
func (o *AuditTrailGetAuditTrailParams) WithLocationID(locationID *string) *AuditTrailGetAuditTrailParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the audit trail get audit trail params
func (o *AuditTrailGetAuditTrailParams) SetLocationID(locationID *string) {
	o.LocationID = locationID
}

// WithPartnerIds adds the partnerIds to the audit trail get audit trail params
func (o *AuditTrailGetAuditTrailParams) WithPartnerIds(partnerIds *string) *AuditTrailGetAuditTrailParams {
	o.SetPartnerIds(partnerIds)
	return o
}

// SetPartnerIds adds the partnerIds to the audit trail get audit trail params
func (o *AuditTrailGetAuditTrailParams) SetPartnerIds(partnerIds *string) {
	o.PartnerIds = partnerIds
}

// WriteToRequest writes these params to a swagger request
func (o *AuditTrailGetAuditTrailParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param customerId
	qrCustomerID := o.CustomerID
	qCustomerID := qrCustomerID
	if qCustomerID != "" {

		if err := r.SetQueryParam("customerId", qCustomerID); err != nil {
			return err
		}
	}

	if o.LocationID != nil {

		// query param locationId
		var qrLocationID string

		if o.LocationID != nil {
			qrLocationID = *o.LocationID
		}
		qLocationID := qrLocationID
		if qLocationID != "" {

			if err := r.SetQueryParam("locationId", qLocationID); err != nil {
				return err
			}
		}
	}

	if o.PartnerIds != nil {

		// query param partnerIds
		var qrPartnerIds string

		if o.PartnerIds != nil {
			qrPartnerIds = *o.PartnerIds
		}
		qPartnerIds := qrPartnerIds
		if qPartnerIds != "" {

			if err := r.SetQueryParam("partnerIds", qPartnerIds); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
