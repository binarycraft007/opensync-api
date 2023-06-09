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

// NewCustomerPrototypePatchPersonParams creates a new CustomerPrototypePatchPersonParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypePatchPersonParams() *CustomerPrototypePatchPersonParams {
	return &CustomerPrototypePatchPersonParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypePatchPersonParamsWithTimeout creates a new CustomerPrototypePatchPersonParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypePatchPersonParamsWithTimeout(timeout time.Duration) *CustomerPrototypePatchPersonParams {
	return &CustomerPrototypePatchPersonParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypePatchPersonParamsWithContext creates a new CustomerPrototypePatchPersonParams object
// with the ability to set a context for a request.
func NewCustomerPrototypePatchPersonParamsWithContext(ctx context.Context) *CustomerPrototypePatchPersonParams {
	return &CustomerPrototypePatchPersonParams{
		Context: ctx,
	}
}

// NewCustomerPrototypePatchPersonParamsWithHTTPClient creates a new CustomerPrototypePatchPersonParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypePatchPersonParamsWithHTTPClient(client *http.Client) *CustomerPrototypePatchPersonParams {
	return &CustomerPrototypePatchPersonParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypePatchPersonParams contains all the parameters to send to the API endpoint

	for the customer prototype patch person operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypePatchPersonParams struct {

	/* AssignedDevices.

	   mac addresses assigned to this Person

	   Format: JSON
	*/
	AssignedDevices *string

	/* Email.

	   email for sending the manager invite
	*/
	Email *string

	/* HomeAwayNotification.

	   track person homeAway state
	*/
	HomeAwayNotification *bool

	/* ID.

	   Customer id

	   Format: JSON
	*/
	ID string

	/* ImageID.

	   unique identifier for referencing a Person's hosted profile image
	*/
	ImageID *string

	// LocationID.
	LocationID string

	// Nickname.
	Nickname *string

	/* Permission.

	   permission object for creating or deleting the manager

	   Format: JSON
	*/
	Permission *string

	// PersonID.
	PersonID string

	/* PrimaryDevice.

	   mac addresses of Person's primary device
	*/
	PrimaryDevice *string

	/* ServiceLinking.

	   serviceLinking object that links this Person object to a 3rd party's Person

	   Format: JSON
	*/
	ServiceLinking *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer prototype patch person params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePatchPersonParams) WithDefaults() *CustomerPrototypePatchPersonParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype patch person params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypePatchPersonParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer prototype patch person params
func (o *CustomerPrototypePatchPersonParams) WithTimeout(timeout time.Duration) *CustomerPrototypePatchPersonParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype patch person params
func (o *CustomerPrototypePatchPersonParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype patch person params
func (o *CustomerPrototypePatchPersonParams) WithContext(ctx context.Context) *CustomerPrototypePatchPersonParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype patch person params
func (o *CustomerPrototypePatchPersonParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype patch person params
func (o *CustomerPrototypePatchPersonParams) WithHTTPClient(client *http.Client) *CustomerPrototypePatchPersonParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype patch person params
func (o *CustomerPrototypePatchPersonParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAssignedDevices adds the assignedDevices to the customer prototype patch person params
func (o *CustomerPrototypePatchPersonParams) WithAssignedDevices(assignedDevices *string) *CustomerPrototypePatchPersonParams {
	o.SetAssignedDevices(assignedDevices)
	return o
}

// SetAssignedDevices adds the assignedDevices to the customer prototype patch person params
func (o *CustomerPrototypePatchPersonParams) SetAssignedDevices(assignedDevices *string) {
	o.AssignedDevices = assignedDevices
}

// WithEmail adds the email to the customer prototype patch person params
func (o *CustomerPrototypePatchPersonParams) WithEmail(email *string) *CustomerPrototypePatchPersonParams {
	o.SetEmail(email)
	return o
}

// SetEmail adds the email to the customer prototype patch person params
func (o *CustomerPrototypePatchPersonParams) SetEmail(email *string) {
	o.Email = email
}

// WithHomeAwayNotification adds the homeAwayNotification to the customer prototype patch person params
func (o *CustomerPrototypePatchPersonParams) WithHomeAwayNotification(homeAwayNotification *bool) *CustomerPrototypePatchPersonParams {
	o.SetHomeAwayNotification(homeAwayNotification)
	return o
}

// SetHomeAwayNotification adds the homeAwayNotification to the customer prototype patch person params
func (o *CustomerPrototypePatchPersonParams) SetHomeAwayNotification(homeAwayNotification *bool) {
	o.HomeAwayNotification = homeAwayNotification
}

// WithID adds the id to the customer prototype patch person params
func (o *CustomerPrototypePatchPersonParams) WithID(id string) *CustomerPrototypePatchPersonParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype patch person params
func (o *CustomerPrototypePatchPersonParams) SetID(id string) {
	o.ID = id
}

// WithImageID adds the imageID to the customer prototype patch person params
func (o *CustomerPrototypePatchPersonParams) WithImageID(imageID *string) *CustomerPrototypePatchPersonParams {
	o.SetImageID(imageID)
	return o
}

// SetImageID adds the imageId to the customer prototype patch person params
func (o *CustomerPrototypePatchPersonParams) SetImageID(imageID *string) {
	o.ImageID = imageID
}

// WithLocationID adds the locationID to the customer prototype patch person params
func (o *CustomerPrototypePatchPersonParams) WithLocationID(locationID string) *CustomerPrototypePatchPersonParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype patch person params
func (o *CustomerPrototypePatchPersonParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WithNickname adds the nickname to the customer prototype patch person params
func (o *CustomerPrototypePatchPersonParams) WithNickname(nickname *string) *CustomerPrototypePatchPersonParams {
	o.SetNickname(nickname)
	return o
}

// SetNickname adds the nickname to the customer prototype patch person params
func (o *CustomerPrototypePatchPersonParams) SetNickname(nickname *string) {
	o.Nickname = nickname
}

// WithPermission adds the permission to the customer prototype patch person params
func (o *CustomerPrototypePatchPersonParams) WithPermission(permission *string) *CustomerPrototypePatchPersonParams {
	o.SetPermission(permission)
	return o
}

// SetPermission adds the permission to the customer prototype patch person params
func (o *CustomerPrototypePatchPersonParams) SetPermission(permission *string) {
	o.Permission = permission
}

// WithPersonID adds the personID to the customer prototype patch person params
func (o *CustomerPrototypePatchPersonParams) WithPersonID(personID string) *CustomerPrototypePatchPersonParams {
	o.SetPersonID(personID)
	return o
}

// SetPersonID adds the personId to the customer prototype patch person params
func (o *CustomerPrototypePatchPersonParams) SetPersonID(personID string) {
	o.PersonID = personID
}

// WithPrimaryDevice adds the primaryDevice to the customer prototype patch person params
func (o *CustomerPrototypePatchPersonParams) WithPrimaryDevice(primaryDevice *string) *CustomerPrototypePatchPersonParams {
	o.SetPrimaryDevice(primaryDevice)
	return o
}

// SetPrimaryDevice adds the primaryDevice to the customer prototype patch person params
func (o *CustomerPrototypePatchPersonParams) SetPrimaryDevice(primaryDevice *string) {
	o.PrimaryDevice = primaryDevice
}

// WithServiceLinking adds the serviceLinking to the customer prototype patch person params
func (o *CustomerPrototypePatchPersonParams) WithServiceLinking(serviceLinking *string) *CustomerPrototypePatchPersonParams {
	o.SetServiceLinking(serviceLinking)
	return o
}

// SetServiceLinking adds the serviceLinking to the customer prototype patch person params
func (o *CustomerPrototypePatchPersonParams) SetServiceLinking(serviceLinking *string) {
	o.ServiceLinking = serviceLinking
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypePatchPersonParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AssignedDevices != nil {

		// form param assignedDevices
		var frAssignedDevices string
		if o.AssignedDevices != nil {
			frAssignedDevices = *o.AssignedDevices
		}
		fAssignedDevices := frAssignedDevices
		if fAssignedDevices != "" {
			if err := r.SetFormParam("assignedDevices", fAssignedDevices); err != nil {
				return err
			}
		}
	}

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

	if o.HomeAwayNotification != nil {

		// form param homeAwayNotification
		var frHomeAwayNotification bool
		if o.HomeAwayNotification != nil {
			frHomeAwayNotification = *o.HomeAwayNotification
		}
		fHomeAwayNotification := swag.FormatBool(frHomeAwayNotification)
		if fHomeAwayNotification != "" {
			if err := r.SetFormParam("homeAwayNotification", fHomeAwayNotification); err != nil {
				return err
			}
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.ImageID != nil {

		// form param imageId
		var frImageID string
		if o.ImageID != nil {
			frImageID = *o.ImageID
		}
		fImageID := frImageID
		if fImageID != "" {
			if err := r.SetFormParam("imageId", fImageID); err != nil {
				return err
			}
		}
	}

	// path param locationId
	if err := r.SetPathParam("locationId", o.LocationID); err != nil {
		return err
	}

	if o.Nickname != nil {

		// form param nickname
		var frNickname string
		if o.Nickname != nil {
			frNickname = *o.Nickname
		}
		fNickname := frNickname
		if fNickname != "" {
			if err := r.SetFormParam("nickname", fNickname); err != nil {
				return err
			}
		}
	}

	if o.Permission != nil {

		// form param permission
		var frPermission string
		if o.Permission != nil {
			frPermission = *o.Permission
		}
		fPermission := frPermission
		if fPermission != "" {
			if err := r.SetFormParam("permission", fPermission); err != nil {
				return err
			}
		}
	}

	// path param personId
	if err := r.SetPathParam("personId", o.PersonID); err != nil {
		return err
	}

	if o.PrimaryDevice != nil {

		// form param primaryDevice
		var frPrimaryDevice string
		if o.PrimaryDevice != nil {
			frPrimaryDevice = *o.PrimaryDevice
		}
		fPrimaryDevice := frPrimaryDevice
		if fPrimaryDevice != "" {
			if err := r.SetFormParam("primaryDevice", fPrimaryDevice); err != nil {
				return err
			}
		}
	}

	if o.ServiceLinking != nil {

		// form param serviceLinking
		var frServiceLinking string
		if o.ServiceLinking != nil {
			frServiceLinking = *o.ServiceLinking
		}
		fServiceLinking := frServiceLinking
		if fServiceLinking != "" {
			if err := r.SetFormParam("serviceLinking", fServiceLinking); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
