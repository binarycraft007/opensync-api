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

// NewCustomerPrototypePostRoomsParams creates a new CustomerPrototypePostRoomsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypePostRoomsParams() *CustomerPrototypePostRoomsParams {
	return &CustomerPrototypePostRoomsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypePostRoomsParamsWithTimeout creates a new CustomerPrototypePostRoomsParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypePostRoomsParamsWithTimeout(timeout time.Duration) *CustomerPrototypePostRoomsParams {
	return &CustomerPrototypePostRoomsParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypePostRoomsParamsWithContext creates a new CustomerPrototypePostRoomsParams object
// with the ability to set a context for a request.
func NewCustomerPrototypePostRoomsParamsWithContext(ctx context.Context) *CustomerPrototypePostRoomsParams {
	return &CustomerPrototypePostRoomsParams{
		Context: ctx,
	}
}

// NewCustomerPrototypePostRoomsParamsWithHTTPClient creates a new CustomerPrototypePostRoomsParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypePostRoomsParamsWithHTTPClient(client *http.Client) *CustomerPrototypePostRoomsParams {
	return &CustomerPrototypePostRoomsParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypePostRoomsParams contains all the parameters to send to the API endpoint

	for the customer prototype post rooms operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypePostRoomsParams struct {

	/* Devices.

	   mac addresses of devices assigned to this Room

	   Format: JSON
	*/
	Devices *string

	/* ID.

	   Customer id

	   Format: JSON
	*/
	ID string

	// LocationID.
	LocationID string

	/* Name.

	   name of this Room
	*/
	Name string

	/* Nodes.

	   nodeIds assigned to this Room

	   Format: JSON
	*/
	Nodes *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer prototype post rooms params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePostRoomsParams) WithDefaults() *CustomerPrototypePostRoomsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype post rooms params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePostRoomsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype post rooms params
func (o *CustomerPrototypePostRoomsParams) WithTimeout(timeout time.Duration) *CustomerPrototypePostRoomsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype post rooms params
func (o *CustomerPrototypePostRoomsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype post rooms params
func (o *CustomerPrototypePostRoomsParams) WithContext(ctx context.Context) *CustomerPrototypePostRoomsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype post rooms params
func (o *CustomerPrototypePostRoomsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype post rooms params
func (o *CustomerPrototypePostRoomsParams) WithHTTPClient(client *http.Client) *CustomerPrototypePostRoomsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype post rooms params
func (o *CustomerPrototypePostRoomsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDevices adds the devices to the customer prototype post rooms params
func (o *CustomerPrototypePostRoomsParams) WithDevices(devices *string) *CustomerPrototypePostRoomsParams {
	o.SetDevices(devices)
	return o
}

// SetDevices adds the devices to the customer prototype post rooms params
func (o *CustomerPrototypePostRoomsParams) SetDevices(devices *string) {
	o.Devices = devices
}

// WithID adds the id to the customer prototype post rooms params
func (o *CustomerPrototypePostRoomsParams) WithID(id string) *CustomerPrototypePostRoomsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype post rooms params
func (o *CustomerPrototypePostRoomsParams) SetID(id string) {
	o.ID = id
}

// WithLocationID adds the locationID to the customer prototype post rooms params
func (o *CustomerPrototypePostRoomsParams) WithLocationID(locationID string) *CustomerPrototypePostRoomsParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype post rooms params
func (o *CustomerPrototypePostRoomsParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WithName adds the name to the customer prototype post rooms params
func (o *CustomerPrototypePostRoomsParams) WithName(name string) *CustomerPrototypePostRoomsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the customer prototype post rooms params
func (o *CustomerPrototypePostRoomsParams) SetName(name string) {
	o.Name = name
}

// WithNodes adds the nodes to the customer prototype post rooms params
func (o *CustomerPrototypePostRoomsParams) WithNodes(nodes *string) *CustomerPrototypePostRoomsParams {
	o.SetNodes(nodes)
	return o
}

// SetNodes adds the nodes to the customer prototype post rooms params
func (o *CustomerPrototypePostRoomsParams) SetNodes(nodes *string) {
	o.Nodes = nodes
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypePostRoomsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Devices != nil {

		// form param devices
		var frDevices string
		if o.Devices != nil {
			frDevices = *o.Devices
		}
		fDevices := frDevices
		if fDevices != "" {
			if err := r.SetFormParam("devices", fDevices); err != nil {
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

	// form param name
	frName := o.Name
	fName := frName
	if fName != "" {
		if err := r.SetFormParam("name", fName); err != nil {
			return err
		}
	}

	if o.Nodes != nil {

		// form param nodes
		var frNodes string
		if o.Nodes != nil {
			frNodes = *o.Nodes
		}
		fNodes := frNodes
		if fNodes != "" {
			if err := r.SetFormParam("nodes", fNodes); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}