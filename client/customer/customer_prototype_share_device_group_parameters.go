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

// NewCustomerPrototypeShareDeviceGroupParams creates a new CustomerPrototypeShareDeviceGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypeShareDeviceGroupParams() *CustomerPrototypeShareDeviceGroupParams {
	return &CustomerPrototypeShareDeviceGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypeShareDeviceGroupParamsWithTimeout creates a new CustomerPrototypeShareDeviceGroupParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypeShareDeviceGroupParamsWithTimeout(timeout time.Duration) *CustomerPrototypeShareDeviceGroupParams {
	return &CustomerPrototypeShareDeviceGroupParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypeShareDeviceGroupParamsWithContext creates a new CustomerPrototypeShareDeviceGroupParams object
// with the ability to set a context for a request.
func NewCustomerPrototypeShareDeviceGroupParamsWithContext(ctx context.Context) *CustomerPrototypeShareDeviceGroupParams {
	return &CustomerPrototypeShareDeviceGroupParams{
		Context: ctx,
	}
}

// NewCustomerPrototypeShareDeviceGroupParamsWithHTTPClient creates a new CustomerPrototypeShareDeviceGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypeShareDeviceGroupParamsWithHTTPClient(client *http.Client) *CustomerPrototypeShareDeviceGroupParams {
	return &CustomerPrototypeShareDeviceGroupParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypeShareDeviceGroupParams contains all the parameters to send to the API endpoint

	for the customer prototype share device group operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypeShareDeviceGroupParams struct {

	// Devices.
	//
	// Format: JSON
	Devices string

	// GroupID.
	GroupID string

	// Groups.
	//
	// Format: JSON
	Groups string

	/* ID.

	   Customer id

	   Format: JSON
	*/
	ID string

	// LocationID.
	LocationID string

	// NetworkID.
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer prototype share device group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeShareDeviceGroupParams) WithDefaults() *CustomerPrototypeShareDeviceGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype share device group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeShareDeviceGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype share device group params
func (o *CustomerPrototypeShareDeviceGroupParams) WithTimeout(timeout time.Duration) *CustomerPrototypeShareDeviceGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype share device group params
func (o *CustomerPrototypeShareDeviceGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype share device group params
func (o *CustomerPrototypeShareDeviceGroupParams) WithContext(ctx context.Context) *CustomerPrototypeShareDeviceGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype share device group params
func (o *CustomerPrototypeShareDeviceGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype share device group params
func (o *CustomerPrototypeShareDeviceGroupParams) WithHTTPClient(client *http.Client) *CustomerPrototypeShareDeviceGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype share device group params
func (o *CustomerPrototypeShareDeviceGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDevices adds the devices to the customer prototype share device group params
func (o *CustomerPrototypeShareDeviceGroupParams) WithDevices(devices string) *CustomerPrototypeShareDeviceGroupParams {
	o.SetDevices(devices)
	return o
}

// SetDevices adds the devices to the customer prototype share device group params
func (o *CustomerPrototypeShareDeviceGroupParams) SetDevices(devices string) {
	o.Devices = devices
}

// WithGroupID adds the groupID to the customer prototype share device group params
func (o *CustomerPrototypeShareDeviceGroupParams) WithGroupID(groupID string) *CustomerPrototypeShareDeviceGroupParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the customer prototype share device group params
func (o *CustomerPrototypeShareDeviceGroupParams) SetGroupID(groupID string) {
	o.GroupID = groupID
}

// WithGroups adds the groups to the customer prototype share device group params
func (o *CustomerPrototypeShareDeviceGroupParams) WithGroups(groups string) *CustomerPrototypeShareDeviceGroupParams {
	o.SetGroups(groups)
	return o
}

// SetGroups adds the groups to the customer prototype share device group params
func (o *CustomerPrototypeShareDeviceGroupParams) SetGroups(groups string) {
	o.Groups = groups
}

// WithID adds the id to the customer prototype share device group params
func (o *CustomerPrototypeShareDeviceGroupParams) WithID(id string) *CustomerPrototypeShareDeviceGroupParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype share device group params
func (o *CustomerPrototypeShareDeviceGroupParams) SetID(id string) {
	o.ID = id
}

// WithLocationID adds the locationID to the customer prototype share device group params
func (o *CustomerPrototypeShareDeviceGroupParams) WithLocationID(locationID string) *CustomerPrototypeShareDeviceGroupParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype share device group params
func (o *CustomerPrototypeShareDeviceGroupParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WithNetworkID adds the networkID to the customer prototype share device group params
func (o *CustomerPrototypeShareDeviceGroupParams) WithNetworkID(networkID string) *CustomerPrototypeShareDeviceGroupParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the customer prototype share device group params
func (o *CustomerPrototypeShareDeviceGroupParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypeShareDeviceGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// form param devices
	frDevices := o.Devices
	fDevices := frDevices
	if fDevices != "" {
		if err := r.SetFormParam("devices", fDevices); err != nil {
			return err
		}
	}

	// path param groupId
	if err := r.SetPathParam("groupId", o.GroupID); err != nil {
		return err
	}

	// form param groups
	frGroups := o.Groups
	fGroups := frGroups
	if fGroups != "" {
		if err := r.SetFormParam("groups", fGroups); err != nil {
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

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
