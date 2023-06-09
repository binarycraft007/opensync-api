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

// NewCustomerPrototypePutPersonFreezeParams creates a new CustomerPrototypePutPersonFreezeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypePutPersonFreezeParams() *CustomerPrototypePutPersonFreezeParams {
	return &CustomerPrototypePutPersonFreezeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypePutPersonFreezeParamsWithTimeout creates a new CustomerPrototypePutPersonFreezeParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypePutPersonFreezeParamsWithTimeout(timeout time.Duration) *CustomerPrototypePutPersonFreezeParams {
	return &CustomerPrototypePutPersonFreezeParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypePutPersonFreezeParamsWithContext creates a new CustomerPrototypePutPersonFreezeParams object
// with the ability to set a context for a request.
func NewCustomerPrototypePutPersonFreezeParamsWithContext(ctx context.Context) *CustomerPrototypePutPersonFreezeParams {
	return &CustomerPrototypePutPersonFreezeParams{
		Context: ctx,
	}
}

// NewCustomerPrototypePutPersonFreezeParamsWithHTTPClient creates a new CustomerPrototypePutPersonFreezeParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypePutPersonFreezeParamsWithHTTPClient(client *http.Client) *CustomerPrototypePutPersonFreezeParams {
	return &CustomerPrototypePutPersonFreezeParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypePutPersonFreezeParams contains all the parameters to send to the API endpoint

	for the customer prototype put person freeze operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypePutPersonFreezeParams struct {

	// DeleteAllExceptSuspend.
	DeleteAllExceptSuspend *bool

	// Enable.
	Enable *bool

	/* FreezeTemplateID.

	   Valid templates are 'untilMidinight', 'schoolNights', etc.
	*/
	FreezeTemplateID string

	/* ID.

	   Customer id

	   Format: JSON
	*/
	ID string

	// LocationID.
	LocationID string

	// Name.
	//
	// Format: JSON
	Name *string

	// PersonID.
	PersonID string

	// Schedules.
	//
	// Format: JSON
	Schedules *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer prototype put person freeze params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePutPersonFreezeParams) WithDefaults() *CustomerPrototypePutPersonFreezeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype put person freeze params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePutPersonFreezeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype put person freeze params
func (o *CustomerPrototypePutPersonFreezeParams) WithTimeout(timeout time.Duration) *CustomerPrototypePutPersonFreezeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype put person freeze params
func (o *CustomerPrototypePutPersonFreezeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype put person freeze params
func (o *CustomerPrototypePutPersonFreezeParams) WithContext(ctx context.Context) *CustomerPrototypePutPersonFreezeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype put person freeze params
func (o *CustomerPrototypePutPersonFreezeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype put person freeze params
func (o *CustomerPrototypePutPersonFreezeParams) WithHTTPClient(client *http.Client) *CustomerPrototypePutPersonFreezeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype put person freeze params
func (o *CustomerPrototypePutPersonFreezeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeleteAllExceptSuspend adds the deleteAllExceptSuspend to the customer prototype put person freeze params
func (o *CustomerPrototypePutPersonFreezeParams) WithDeleteAllExceptSuspend(deleteAllExceptSuspend *bool) *CustomerPrototypePutPersonFreezeParams {
	o.SetDeleteAllExceptSuspend(deleteAllExceptSuspend)
	return o
}

// SetDeleteAllExceptSuspend adds the deleteAllExceptSuspend to the customer prototype put person freeze params
func (o *CustomerPrototypePutPersonFreezeParams) SetDeleteAllExceptSuspend(deleteAllExceptSuspend *bool) {
	o.DeleteAllExceptSuspend = deleteAllExceptSuspend
}

// WithEnable adds the enable to the customer prototype put person freeze params
func (o *CustomerPrototypePutPersonFreezeParams) WithEnable(enable *bool) *CustomerPrototypePutPersonFreezeParams {
	o.SetEnable(enable)
	return o
}

// SetEnable adds the enable to the customer prototype put person freeze params
func (o *CustomerPrototypePutPersonFreezeParams) SetEnable(enable *bool) {
	o.Enable = enable
}

// WithFreezeTemplateID adds the freezeTemplateID to the customer prototype put person freeze params
func (o *CustomerPrototypePutPersonFreezeParams) WithFreezeTemplateID(freezeTemplateID string) *CustomerPrototypePutPersonFreezeParams {
	o.SetFreezeTemplateID(freezeTemplateID)
	return o
}

// SetFreezeTemplateID adds the freezeTemplateId to the customer prototype put person freeze params
func (o *CustomerPrototypePutPersonFreezeParams) SetFreezeTemplateID(freezeTemplateID string) {
	o.FreezeTemplateID = freezeTemplateID
}

// WithID adds the id to the customer prototype put person freeze params
func (o *CustomerPrototypePutPersonFreezeParams) WithID(id string) *CustomerPrototypePutPersonFreezeParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype put person freeze params
func (o *CustomerPrototypePutPersonFreezeParams) SetID(id string) {
	o.ID = id
}

// WithLocationID adds the locationID to the customer prototype put person freeze params
func (o *CustomerPrototypePutPersonFreezeParams) WithLocationID(locationID string) *CustomerPrototypePutPersonFreezeParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype put person freeze params
func (o *CustomerPrototypePutPersonFreezeParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WithName adds the name to the customer prototype put person freeze params
func (o *CustomerPrototypePutPersonFreezeParams) WithName(name *string) *CustomerPrototypePutPersonFreezeParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the customer prototype put person freeze params
func (o *CustomerPrototypePutPersonFreezeParams) SetName(name *string) {
	o.Name = name
}

// WithPersonID adds the personID to the customer prototype put person freeze params
func (o *CustomerPrototypePutPersonFreezeParams) WithPersonID(personID string) *CustomerPrototypePutPersonFreezeParams {
	o.SetPersonID(personID)
	return o
}

// SetPersonID adds the personId to the customer prototype put person freeze params
func (o *CustomerPrototypePutPersonFreezeParams) SetPersonID(personID string) {
	o.PersonID = personID
}

// WithSchedules adds the schedules to the customer prototype put person freeze params
func (o *CustomerPrototypePutPersonFreezeParams) WithSchedules(schedules *string) *CustomerPrototypePutPersonFreezeParams {
	o.SetSchedules(schedules)
	return o
}

// SetSchedules adds the schedules to the customer prototype put person freeze params
func (o *CustomerPrototypePutPersonFreezeParams) SetSchedules(schedules *string) {
	o.Schedules = schedules
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypePutPersonFreezeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DeleteAllExceptSuspend != nil {

		// form param deleteAllExceptSuspend
		var frDeleteAllExceptSuspend bool
		if o.DeleteAllExceptSuspend != nil {
			frDeleteAllExceptSuspend = *o.DeleteAllExceptSuspend
		}
		fDeleteAllExceptSuspend := swag.FormatBool(frDeleteAllExceptSuspend)
		if fDeleteAllExceptSuspend != "" {
			if err := r.SetFormParam("deleteAllExceptSuspend", fDeleteAllExceptSuspend); err != nil {
				return err
			}
		}
	}

	if o.Enable != nil {

		// form param enable
		var frEnable bool
		if o.Enable != nil {
			frEnable = *o.Enable
		}
		fEnable := swag.FormatBool(frEnable)
		if fEnable != "" {
			if err := r.SetFormParam("enable", fEnable); err != nil {
				return err
			}
		}
	}

	// path param freezeTemplateId
	if err := r.SetPathParam("freezeTemplateId", o.FreezeTemplateID); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param locationId
	if err := r.SetPathParam("locationId", o.LocationID); err != nil {
		return err
	}

	if o.Name != nil {

		// form param name
		var frName string
		if o.Name != nil {
			frName = *o.Name
		}
		fName := frName
		if fName != "" {
			if err := r.SetFormParam("name", fName); err != nil {
				return err
			}
		}
	}

	// path param personId
	if err := r.SetPathParam("personId", o.PersonID); err != nil {
		return err
	}

	if o.Schedules != nil {

		// form param schedules
		var frSchedules string
		if o.Schedules != nil {
			frSchedules = *o.Schedules
		}
		fSchedules := frSchedules
		if fSchedules != "" {
			if err := r.SetFormParam("schedules", fSchedules); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
