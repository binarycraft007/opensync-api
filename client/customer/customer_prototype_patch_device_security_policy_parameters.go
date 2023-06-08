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

// NewCustomerPrototypePatchDeviceSecurityPolicyParams creates a new CustomerPrototypePatchDeviceSecurityPolicyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypePatchDeviceSecurityPolicyParams() *CustomerPrototypePatchDeviceSecurityPolicyParams {
	return &CustomerPrototypePatchDeviceSecurityPolicyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypePatchDeviceSecurityPolicyParamsWithTimeout creates a new CustomerPrototypePatchDeviceSecurityPolicyParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypePatchDeviceSecurityPolicyParamsWithTimeout(timeout time.Duration) *CustomerPrototypePatchDeviceSecurityPolicyParams {
	return &CustomerPrototypePatchDeviceSecurityPolicyParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypePatchDeviceSecurityPolicyParamsWithContext creates a new CustomerPrototypePatchDeviceSecurityPolicyParams object
// with the ability to set a context for a request.
func NewCustomerPrototypePatchDeviceSecurityPolicyParamsWithContext(ctx context.Context) *CustomerPrototypePatchDeviceSecurityPolicyParams {
	return &CustomerPrototypePatchDeviceSecurityPolicyParams{
		Context: ctx,
	}
}

// NewCustomerPrototypePatchDeviceSecurityPolicyParamsWithHTTPClient creates a new CustomerPrototypePatchDeviceSecurityPolicyParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypePatchDeviceSecurityPolicyParamsWithHTTPClient(client *http.Client) *CustomerPrototypePatchDeviceSecurityPolicyParams {
	return &CustomerPrototypePatchDeviceSecurityPolicyParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypePatchDeviceSecurityPolicyParams contains all the parameters to send to the API endpoint

	for the customer prototype patch device security policy operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypePatchDeviceSecurityPolicyParams struct {

	/* Content.

	   Valid values: 'kids || teenagers || adBlocking || adultAndSensitive || workAppropriate'

	   Format: JSON
	*/
	Content *string

	/* ID.

	   Customer id

	   Format: JSON
	*/
	ID string

	// IotProtect.
	IotProtect *bool

	// IotProtectReason.
	IotProtectReason *string

	// LocationID.
	LocationID string

	// Mac.
	Mac string

	// SecureAndProtect.
	SecureAndProtect *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer prototype patch device security policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePatchDeviceSecurityPolicyParams) WithDefaults() *CustomerPrototypePatchDeviceSecurityPolicyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype patch device security policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePatchDeviceSecurityPolicyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype patch device security policy params
func (o *CustomerPrototypePatchDeviceSecurityPolicyParams) WithTimeout(timeout time.Duration) *CustomerPrototypePatchDeviceSecurityPolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype patch device security policy params
func (o *CustomerPrototypePatchDeviceSecurityPolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype patch device security policy params
func (o *CustomerPrototypePatchDeviceSecurityPolicyParams) WithContext(ctx context.Context) *CustomerPrototypePatchDeviceSecurityPolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype patch device security policy params
func (o *CustomerPrototypePatchDeviceSecurityPolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype patch device security policy params
func (o *CustomerPrototypePatchDeviceSecurityPolicyParams) WithHTTPClient(client *http.Client) *CustomerPrototypePatchDeviceSecurityPolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype patch device security policy params
func (o *CustomerPrototypePatchDeviceSecurityPolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContent adds the content to the customer prototype patch device security policy params
func (o *CustomerPrototypePatchDeviceSecurityPolicyParams) WithContent(content *string) *CustomerPrototypePatchDeviceSecurityPolicyParams {
	o.SetContent(content)
	return o
}

// SetContent adds the content to the customer prototype patch device security policy params
func (o *CustomerPrototypePatchDeviceSecurityPolicyParams) SetContent(content *string) {
	o.Content = content
}

// WithID adds the id to the customer prototype patch device security policy params
func (o *CustomerPrototypePatchDeviceSecurityPolicyParams) WithID(id string) *CustomerPrototypePatchDeviceSecurityPolicyParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype patch device security policy params
func (o *CustomerPrototypePatchDeviceSecurityPolicyParams) SetID(id string) {
	o.ID = id
}

// WithIotProtect adds the iotProtect to the customer prototype patch device security policy params
func (o *CustomerPrototypePatchDeviceSecurityPolicyParams) WithIotProtect(iotProtect *bool) *CustomerPrototypePatchDeviceSecurityPolicyParams {
	o.SetIotProtect(iotProtect)
	return o
}

// SetIotProtect adds the iotProtect to the customer prototype patch device security policy params
func (o *CustomerPrototypePatchDeviceSecurityPolicyParams) SetIotProtect(iotProtect *bool) {
	o.IotProtect = iotProtect
}

// WithIotProtectReason adds the iotProtectReason to the customer prototype patch device security policy params
func (o *CustomerPrototypePatchDeviceSecurityPolicyParams) WithIotProtectReason(iotProtectReason *string) *CustomerPrototypePatchDeviceSecurityPolicyParams {
	o.SetIotProtectReason(iotProtectReason)
	return o
}

// SetIotProtectReason adds the iotProtectReason to the customer prototype patch device security policy params
func (o *CustomerPrototypePatchDeviceSecurityPolicyParams) SetIotProtectReason(iotProtectReason *string) {
	o.IotProtectReason = iotProtectReason
}

// WithLocationID adds the locationID to the customer prototype patch device security policy params
func (o *CustomerPrototypePatchDeviceSecurityPolicyParams) WithLocationID(locationID string) *CustomerPrototypePatchDeviceSecurityPolicyParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype patch device security policy params
func (o *CustomerPrototypePatchDeviceSecurityPolicyParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WithMac adds the mac to the customer prototype patch device security policy params
func (o *CustomerPrototypePatchDeviceSecurityPolicyParams) WithMac(mac string) *CustomerPrototypePatchDeviceSecurityPolicyParams {
	o.SetMac(mac)
	return o
}

// SetMac adds the mac to the customer prototype patch device security policy params
func (o *CustomerPrototypePatchDeviceSecurityPolicyParams) SetMac(mac string) {
	o.Mac = mac
}

// WithSecureAndProtect adds the secureAndProtect to the customer prototype patch device security policy params
func (o *CustomerPrototypePatchDeviceSecurityPolicyParams) WithSecureAndProtect(secureAndProtect *bool) *CustomerPrototypePatchDeviceSecurityPolicyParams {
	o.SetSecureAndProtect(secureAndProtect)
	return o
}

// SetSecureAndProtect adds the secureAndProtect to the customer prototype patch device security policy params
func (o *CustomerPrototypePatchDeviceSecurityPolicyParams) SetSecureAndProtect(secureAndProtect *bool) {
	o.SecureAndProtect = secureAndProtect
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypePatchDeviceSecurityPolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.IotProtect != nil {

		// form param iotProtect
		var frIotProtect bool
		if o.IotProtect != nil {
			frIotProtect = *o.IotProtect
		}
		fIotProtect := swag.FormatBool(frIotProtect)
		if fIotProtect != "" {
			if err := r.SetFormParam("iotProtect", fIotProtect); err != nil {
				return err
			}
		}
	}

	if o.IotProtectReason != nil {

		// form param iotProtectReason
		var frIotProtectReason string
		if o.IotProtectReason != nil {
			frIotProtectReason = *o.IotProtectReason
		}
		fIotProtectReason := frIotProtectReason
		if fIotProtectReason != "" {
			if err := r.SetFormParam("iotProtectReason", fIotProtectReason); err != nil {
				return err
			}
		}
	}

	// path param locationId
	if err := r.SetPathParam("locationId", o.LocationID); err != nil {
		return err
	}

	// path param mac
	if err := r.SetPathParam("mac", o.Mac); err != nil {
		return err
	}

	if o.SecureAndProtect != nil {

		// form param secureAndProtect
		var frSecureAndProtect bool
		if o.SecureAndProtect != nil {
			frSecureAndProtect = *o.SecureAndProtect
		}
		fSecureAndProtect := swag.FormatBool(frSecureAndProtect)
		if fSecureAndProtect != "" {
			if err := r.SetFormParam("secureAndProtect", fSecureAndProtect); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
