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

// NewCustomerPrototypePostWifiKeyParams creates a new CustomerPrototypePostWifiKeyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypePostWifiKeyParams() *CustomerPrototypePostWifiKeyParams {
	return &CustomerPrototypePostWifiKeyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypePostWifiKeyParamsWithTimeout creates a new CustomerPrototypePostWifiKeyParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypePostWifiKeyParamsWithTimeout(timeout time.Duration) *CustomerPrototypePostWifiKeyParams {
	return &CustomerPrototypePostWifiKeyParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypePostWifiKeyParamsWithContext creates a new CustomerPrototypePostWifiKeyParams object
// with the ability to set a context for a request.
func NewCustomerPrototypePostWifiKeyParamsWithContext(ctx context.Context) *CustomerPrototypePostWifiKeyParams {
	return &CustomerPrototypePostWifiKeyParams{
		Context: ctx,
	}
}

// NewCustomerPrototypePostWifiKeyParamsWithHTTPClient creates a new CustomerPrototypePostWifiKeyParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypePostWifiKeyParamsWithHTTPClient(client *http.Client) *CustomerPrototypePostWifiKeyParams {
	return &CustomerPrototypePostWifiKeyParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypePostWifiKeyParams contains all the parameters to send to the API endpoint

	for the customer prototype post wifi key operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypePostWifiKeyParams struct {

	/* AccessZone.

	   home | guests | internetAccessOnly
	*/
	AccessZone string

	/* Content.

	   Valid values: 'adultAndSensitive'

	   Format: JSON
	*/
	Content *string

	/* Enable.

	   devices can connect using this encryptionKey
	*/
	Enable bool

	// EncryptionKey.
	EncryptionKey string

	/* ExpiresAt.

	   UTC in ISO 8601 String format

	   Format: date-time
	*/
	ExpiresAt *strfmt.DateTime

	/* Format.

	   encryptionKey | phoneNumber
	*/
	Format string

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

// WithDefaults hydrates default values in the customer prototype post wifi key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePostWifiKeyParams) WithDefaults() *CustomerPrototypePostWifiKeyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype post wifi key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePostWifiKeyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype post wifi key params
func (o *CustomerPrototypePostWifiKeyParams) WithTimeout(timeout time.Duration) *CustomerPrototypePostWifiKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype post wifi key params
func (o *CustomerPrototypePostWifiKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype post wifi key params
func (o *CustomerPrototypePostWifiKeyParams) WithContext(ctx context.Context) *CustomerPrototypePostWifiKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype post wifi key params
func (o *CustomerPrototypePostWifiKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype post wifi key params
func (o *CustomerPrototypePostWifiKeyParams) WithHTTPClient(client *http.Client) *CustomerPrototypePostWifiKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype post wifi key params
func (o *CustomerPrototypePostWifiKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccessZone adds the accessZone to the customer prototype post wifi key params
func (o *CustomerPrototypePostWifiKeyParams) WithAccessZone(accessZone string) *CustomerPrototypePostWifiKeyParams {
	o.SetAccessZone(accessZone)
	return o
}

// SetAccessZone adds the accessZone to the customer prototype post wifi key params
func (o *CustomerPrototypePostWifiKeyParams) SetAccessZone(accessZone string) {
	o.AccessZone = accessZone
}

// WithContent adds the content to the customer prototype post wifi key params
func (o *CustomerPrototypePostWifiKeyParams) WithContent(content *string) *CustomerPrototypePostWifiKeyParams {
	o.SetContent(content)
	return o
}

// SetContent adds the content to the customer prototype post wifi key params
func (o *CustomerPrototypePostWifiKeyParams) SetContent(content *string) {
	o.Content = content
}

// WithEnable adds the enable to the customer prototype post wifi key params
func (o *CustomerPrototypePostWifiKeyParams) WithEnable(enable bool) *CustomerPrototypePostWifiKeyParams {
	o.SetEnable(enable)
	return o
}

// SetEnable adds the enable to the customer prototype post wifi key params
func (o *CustomerPrototypePostWifiKeyParams) SetEnable(enable bool) {
	o.Enable = enable
}

// WithEncryptionKey adds the encryptionKey to the customer prototype post wifi key params
func (o *CustomerPrototypePostWifiKeyParams) WithEncryptionKey(encryptionKey string) *CustomerPrototypePostWifiKeyParams {
	o.SetEncryptionKey(encryptionKey)
	return o
}

// SetEncryptionKey adds the encryptionKey to the customer prototype post wifi key params
func (o *CustomerPrototypePostWifiKeyParams) SetEncryptionKey(encryptionKey string) {
	o.EncryptionKey = encryptionKey
}

// WithExpiresAt adds the expiresAt to the customer prototype post wifi key params
func (o *CustomerPrototypePostWifiKeyParams) WithExpiresAt(expiresAt *strfmt.DateTime) *CustomerPrototypePostWifiKeyParams {
	o.SetExpiresAt(expiresAt)
	return o
}

// SetExpiresAt adds the expiresAt to the customer prototype post wifi key params
func (o *CustomerPrototypePostWifiKeyParams) SetExpiresAt(expiresAt *strfmt.DateTime) {
	o.ExpiresAt = expiresAt
}

// WithFormat adds the format to the customer prototype post wifi key params
func (o *CustomerPrototypePostWifiKeyParams) WithFormat(format string) *CustomerPrototypePostWifiKeyParams {
	o.SetFormat(format)
	return o
}

// SetFormat adds the format to the customer prototype post wifi key params
func (o *CustomerPrototypePostWifiKeyParams) SetFormat(format string) {
	o.Format = format
}

// WithID adds the id to the customer prototype post wifi key params
func (o *CustomerPrototypePostWifiKeyParams) WithID(id string) *CustomerPrototypePostWifiKeyParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype post wifi key params
func (o *CustomerPrototypePostWifiKeyParams) SetID(id string) {
	o.ID = id
}

// WithLocationID adds the locationID to the customer prototype post wifi key params
func (o *CustomerPrototypePostWifiKeyParams) WithLocationID(locationID string) *CustomerPrototypePostWifiKeyParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype post wifi key params
func (o *CustomerPrototypePostWifiKeyParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypePostWifiKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param accessZone
	if err := r.SetPathParam("accessZone", o.AccessZone); err != nil {
		return err
	}

	if o.Content != nil {

		// form param content
		var frContent string
		if o.Content != nil {
			frContent = *o.Content
		}
		fContent := frContent
		if fContent != "" {
			if err := r.SetFormParam("content", fContent); err != nil {
				return err
			}
		}
	}

	// form param enable
	frEnable := o.Enable
	fEnable := swag.FormatBool(frEnable)
	if fEnable != "" {
		if err := r.SetFormParam("enable", fEnable); err != nil {
			return err
		}
	}

	// form param encryptionKey
	frEncryptionKey := o.EncryptionKey
	fEncryptionKey := frEncryptionKey
	if fEncryptionKey != "" {
		if err := r.SetFormParam("encryptionKey", fEncryptionKey); err != nil {
			return err
		}
	}

	if o.ExpiresAt != nil {

		// form param expiresAt
		var frExpiresAt strfmt.DateTime
		if o.ExpiresAt != nil {
			frExpiresAt = *o.ExpiresAt
		}
		fExpiresAt := frExpiresAt.String()
		if fExpiresAt != "" {
			if err := r.SetFormParam("expiresAt", fExpiresAt); err != nil {
				return err
			}
		}
	}

	// form param format
	frFormat := o.Format
	fFormat := frFormat
	if fFormat != "" {
		if err := r.SetFormParam("format", fFormat); err != nil {
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
