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

// NewCustomerPrototypePatchCustomDeviceTypeParams creates a new CustomerPrototypePatchCustomDeviceTypeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypePatchCustomDeviceTypeParams() *CustomerPrototypePatchCustomDeviceTypeParams {
	return &CustomerPrototypePatchCustomDeviceTypeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypePatchCustomDeviceTypeParamsWithTimeout creates a new CustomerPrototypePatchCustomDeviceTypeParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypePatchCustomDeviceTypeParamsWithTimeout(timeout time.Duration) *CustomerPrototypePatchCustomDeviceTypeParams {
	return &CustomerPrototypePatchCustomDeviceTypeParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypePatchCustomDeviceTypeParamsWithContext creates a new CustomerPrototypePatchCustomDeviceTypeParams object
// with the ability to set a context for a request.
func NewCustomerPrototypePatchCustomDeviceTypeParamsWithContext(ctx context.Context) *CustomerPrototypePatchCustomDeviceTypeParams {
	return &CustomerPrototypePatchCustomDeviceTypeParams{
		Context: ctx,
	}
}

// NewCustomerPrototypePatchCustomDeviceTypeParamsWithHTTPClient creates a new CustomerPrototypePatchCustomDeviceTypeParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypePatchCustomDeviceTypeParamsWithHTTPClient(client *http.Client) *CustomerPrototypePatchCustomDeviceTypeParams {
	return &CustomerPrototypePatchCustomDeviceTypeParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypePatchCustomDeviceTypeParams contains all the parameters to send to the API endpoint

	for the customer prototype patch custom device type operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypePatchCustomDeviceTypeParams struct {

	// Brand.
	//
	// Format: JSON
	Brand *string

	// Category.
	//
	// Format: JSON
	Category *string

	/* ID.

	   Customer id

	   Format: JSON
	*/
	ID string

	// LocationID.
	//
	// Format: JSON
	LocationID string

	// Mac.
	Mac string

	// Model.
	//
	// Format: JSON
	Model *string

	// OsName.
	//
	// Format: JSON
	OsName *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer prototype patch custom device type params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePatchCustomDeviceTypeParams) WithDefaults() *CustomerPrototypePatchCustomDeviceTypeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype patch custom device type params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePatchCustomDeviceTypeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype patch custom device type params
func (o *CustomerPrototypePatchCustomDeviceTypeParams) WithTimeout(timeout time.Duration) *CustomerPrototypePatchCustomDeviceTypeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype patch custom device type params
func (o *CustomerPrototypePatchCustomDeviceTypeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype patch custom device type params
func (o *CustomerPrototypePatchCustomDeviceTypeParams) WithContext(ctx context.Context) *CustomerPrototypePatchCustomDeviceTypeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype patch custom device type params
func (o *CustomerPrototypePatchCustomDeviceTypeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype patch custom device type params
func (o *CustomerPrototypePatchCustomDeviceTypeParams) WithHTTPClient(client *http.Client) *CustomerPrototypePatchCustomDeviceTypeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype patch custom device type params
func (o *CustomerPrototypePatchCustomDeviceTypeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBrand adds the brand to the customer prototype patch custom device type params
func (o *CustomerPrototypePatchCustomDeviceTypeParams) WithBrand(brand *string) *CustomerPrototypePatchCustomDeviceTypeParams {
	o.SetBrand(brand)
	return o
}

// SetBrand adds the brand to the customer prototype patch custom device type params
func (o *CustomerPrototypePatchCustomDeviceTypeParams) SetBrand(brand *string) {
	o.Brand = brand
}

// WithCategory adds the category to the customer prototype patch custom device type params
func (o *CustomerPrototypePatchCustomDeviceTypeParams) WithCategory(category *string) *CustomerPrototypePatchCustomDeviceTypeParams {
	o.SetCategory(category)
	return o
}

// SetCategory adds the category to the customer prototype patch custom device type params
func (o *CustomerPrototypePatchCustomDeviceTypeParams) SetCategory(category *string) {
	o.Category = category
}

// WithID adds the id to the customer prototype patch custom device type params
func (o *CustomerPrototypePatchCustomDeviceTypeParams) WithID(id string) *CustomerPrototypePatchCustomDeviceTypeParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype patch custom device type params
func (o *CustomerPrototypePatchCustomDeviceTypeParams) SetID(id string) {
	o.ID = id
}

// WithLocationID adds the locationID to the customer prototype patch custom device type params
func (o *CustomerPrototypePatchCustomDeviceTypeParams) WithLocationID(locationID string) *CustomerPrototypePatchCustomDeviceTypeParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype patch custom device type params
func (o *CustomerPrototypePatchCustomDeviceTypeParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WithMac adds the mac to the customer prototype patch custom device type params
func (o *CustomerPrototypePatchCustomDeviceTypeParams) WithMac(mac string) *CustomerPrototypePatchCustomDeviceTypeParams {
	o.SetMac(mac)
	return o
}

// SetMac adds the mac to the customer prototype patch custom device type params
func (o *CustomerPrototypePatchCustomDeviceTypeParams) SetMac(mac string) {
	o.Mac = mac
}

// WithModel adds the model to the customer prototype patch custom device type params
func (o *CustomerPrototypePatchCustomDeviceTypeParams) WithModel(model *string) *CustomerPrototypePatchCustomDeviceTypeParams {
	o.SetModel(model)
	return o
}

// SetModel adds the model to the customer prototype patch custom device type params
func (o *CustomerPrototypePatchCustomDeviceTypeParams) SetModel(model *string) {
	o.Model = model
}

// WithOsName adds the osName to the customer prototype patch custom device type params
func (o *CustomerPrototypePatchCustomDeviceTypeParams) WithOsName(osName *string) *CustomerPrototypePatchCustomDeviceTypeParams {
	o.SetOsName(osName)
	return o
}

// SetOsName adds the osName to the customer prototype patch custom device type params
func (o *CustomerPrototypePatchCustomDeviceTypeParams) SetOsName(osName *string) {
	o.OsName = osName
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypePatchCustomDeviceTypeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Brand != nil {

		// form param brand
		var frBrand string
		if o.Brand != nil {
			frBrand = *o.Brand
		}
		fBrand := frBrand
		if fBrand != "" {
			if err := r.SetFormParam("brand", fBrand); err != nil {
				return err
			}
		}
	}

	if o.Category != nil {

		// form param category
		var frCategory string
		if o.Category != nil {
			frCategory = *o.Category
		}
		fCategory := frCategory
		if fCategory != "" {
			if err := r.SetFormParam("category", fCategory); err != nil {
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

	if o.Model != nil {

		// form param model
		var frModel string
		if o.Model != nil {
			frModel = *o.Model
		}
		fModel := frModel
		if fModel != "" {
			if err := r.SetFormParam("model", fModel); err != nil {
				return err
			}
		}
	}

	if o.OsName != nil {

		// form param osName
		var frOsName string
		if o.OsName != nil {
			frOsName = *o.OsName
		}
		fOsName := frOsName
		if fOsName != "" {
			if err := r.SetFormParam("osName", fOsName); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
