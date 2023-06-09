// Code generated by go-swagger; DO NOT EDIT.

package group

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

// NewGroupCustomCreateParams creates a new GroupCustomCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGroupCustomCreateParams() *GroupCustomCreateParams {
	return &GroupCustomCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGroupCustomCreateParamsWithTimeout creates a new GroupCustomCreateParams object
// with the ability to set a timeout on a request.
func NewGroupCustomCreateParamsWithTimeout(timeout time.Duration) *GroupCustomCreateParams {
	return &GroupCustomCreateParams{
		timeout: timeout,
	}
}

// NewGroupCustomCreateParamsWithContext creates a new GroupCustomCreateParams object
// with the ability to set a context for a request.
func NewGroupCustomCreateParamsWithContext(ctx context.Context) *GroupCustomCreateParams {
	return &GroupCustomCreateParams{
		Context: ctx,
	}
}

// NewGroupCustomCreateParamsWithHTTPClient creates a new GroupCustomCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewGroupCustomCreateParamsWithHTTPClient(client *http.Client) *GroupCustomCreateParams {
	return &GroupCustomCreateParams{
		HTTPClient: client,
	}
}

/*
GroupCustomCreateParams contains all the parameters to send to the API endpoint

	for the group custom create operation.

	Typically these are written to a http.Request.
*/
type GroupCustomCreateParams struct {

	// Description.
	Description string

	/* ID.

	   an ObjectID id for porting group ids
	*/
	ID *string

	// ImportGroupIDAsPartnerIDIntoInventory.
	ImportGroupIDAsPartnerIDIntoInventory *bool

	// Name.
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the group custom create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GroupCustomCreateParams) WithDefaults() *GroupCustomCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the group custom create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GroupCustomCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the group custom create params
func (o *GroupCustomCreateParams) WithTimeout(timeout time.Duration) *GroupCustomCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the group custom create params
func (o *GroupCustomCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the group custom create params
func (o *GroupCustomCreateParams) WithContext(ctx context.Context) *GroupCustomCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the group custom create params
func (o *GroupCustomCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the group custom create params
func (o *GroupCustomCreateParams) WithHTTPClient(client *http.Client) *GroupCustomCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the group custom create params
func (o *GroupCustomCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDescription adds the description to the group custom create params
func (o *GroupCustomCreateParams) WithDescription(description string) *GroupCustomCreateParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the group custom create params
func (o *GroupCustomCreateParams) SetDescription(description string) {
	o.Description = description
}

// WithID adds the id to the group custom create params
func (o *GroupCustomCreateParams) WithID(id *string) *GroupCustomCreateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the group custom create params
func (o *GroupCustomCreateParams) SetID(id *string) {
	o.ID = id
}

// WithImportGroupIDAsPartnerIDIntoInventory adds the importGroupIDAsPartnerIDIntoInventory to the group custom create params
func (o *GroupCustomCreateParams) WithImportGroupIDAsPartnerIDIntoInventory(importGroupIDAsPartnerIDIntoInventory *bool) *GroupCustomCreateParams {
	o.SetImportGroupIDAsPartnerIDIntoInventory(importGroupIDAsPartnerIDIntoInventory)
	return o
}

// SetImportGroupIDAsPartnerIDIntoInventory adds the importGroupIdAsPartnerIdIntoInventory to the group custom create params
func (o *GroupCustomCreateParams) SetImportGroupIDAsPartnerIDIntoInventory(importGroupIDAsPartnerIDIntoInventory *bool) {
	o.ImportGroupIDAsPartnerIDIntoInventory = importGroupIDAsPartnerIDIntoInventory
}

// WithName adds the name to the group custom create params
func (o *GroupCustomCreateParams) WithName(name string) *GroupCustomCreateParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the group custom create params
func (o *GroupCustomCreateParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GroupCustomCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// form param description
	frDescription := o.Description
	fDescription := frDescription
	if fDescription != "" {
		if err := r.SetFormParam("description", fDescription); err != nil {
			return err
		}
	}

	if o.ID != nil {

		// form param id
		var frID string
		if o.ID != nil {
			frID = *o.ID
		}
		fID := frID
		if fID != "" {
			if err := r.SetFormParam("id", fID); err != nil {
				return err
			}
		}
	}

	if o.ImportGroupIDAsPartnerIDIntoInventory != nil {

		// form param importGroupIdAsPartnerIdIntoInventory
		var frImportGroupIDAsPartnerIDIntoInventory bool
		if o.ImportGroupIDAsPartnerIDIntoInventory != nil {
			frImportGroupIDAsPartnerIDIntoInventory = *o.ImportGroupIDAsPartnerIDIntoInventory
		}
		fImportGroupIDAsPartnerIDIntoInventory := swag.FormatBool(frImportGroupIDAsPartnerIDIntoInventory)
		if fImportGroupIDAsPartnerIDIntoInventory != "" {
			if err := r.SetFormParam("importGroupIdAsPartnerIdIntoInventory", fImportGroupIDAsPartnerIDIntoInventory); err != nil {
				return err
			}
		}
	}

	// form param name
	frName := o.Name
	fName := frName
	if fName != "" {
		if err := r.SetFormParam("name", fName); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
