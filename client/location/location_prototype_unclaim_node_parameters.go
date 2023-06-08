// Code generated by go-swagger; DO NOT EDIT.

package location

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

// NewLocationPrototypeUnclaimNodeParams creates a new LocationPrototypeUnclaimNodeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLocationPrototypeUnclaimNodeParams() *LocationPrototypeUnclaimNodeParams {
	return &LocationPrototypeUnclaimNodeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLocationPrototypeUnclaimNodeParamsWithTimeout creates a new LocationPrototypeUnclaimNodeParams object
// with the ability to set a timeout on a request.
func NewLocationPrototypeUnclaimNodeParamsWithTimeout(timeout time.Duration) *LocationPrototypeUnclaimNodeParams {
	return &LocationPrototypeUnclaimNodeParams{
		timeout: timeout,
	}
}

// NewLocationPrototypeUnclaimNodeParamsWithContext creates a new LocationPrototypeUnclaimNodeParams object
// with the ability to set a context for a request.
func NewLocationPrototypeUnclaimNodeParamsWithContext(ctx context.Context) *LocationPrototypeUnclaimNodeParams {
	return &LocationPrototypeUnclaimNodeParams{
		Context: ctx,
	}
}

// NewLocationPrototypeUnclaimNodeParamsWithHTTPClient creates a new LocationPrototypeUnclaimNodeParams object
// with the ability to set a custom HTTPClient for a request.
func NewLocationPrototypeUnclaimNodeParamsWithHTTPClient(client *http.Client) *LocationPrototypeUnclaimNodeParams {
	return &LocationPrototypeUnclaimNodeParams{
		HTTPClient: client,
	}
}

/*
LocationPrototypeUnclaimNodeParams contains all the parameters to send to the API endpoint

	for the location prototype unclaim node operation.

	Typically these are written to a http.Request.
*/
type LocationPrototypeUnclaimNodeParams struct {

	/* ForceUnclaim.

	   Unclaim regardless of pod connectivity
	*/
	ForceUnclaim *bool

	/* ID.

	   Location id

	   Format: JSON
	*/
	ID string

	/* IncrementFactoryResetCounter.

	   Whether or not to increment the factory reset counter
	*/
	IncrementFactoryResetCounter *bool

	// NodeID.
	NodeID string

	/* PreservePackID.

	   Whether or not packId should remain the same
	*/
	PreservePackID *bool

	/* PurgeGroupIds.

	   Whether or not groupIds should be kept on the node
	*/
	PurgeGroupIds *bool

	/* RemoveAccountID.

	   delete account id on the inventory node
	*/
	RemoveAccountID *bool

	/* UnclaimReason.

	   Used by controller to determine what actions to take for nodeClaimChanged
	*/
	UnclaimReason *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the location prototype unclaim node params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LocationPrototypeUnclaimNodeParams) WithDefaults() *LocationPrototypeUnclaimNodeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the location prototype unclaim node params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LocationPrototypeUnclaimNodeParams) SetDefaults() {
	var (
		removeAccountIDDefault = bool(false)
	)

	val := LocationPrototypeUnclaimNodeParams{
		RemoveAccountID: &removeAccountIDDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the location prototype unclaim node params
func (o *LocationPrototypeUnclaimNodeParams) WithTimeout(timeout time.Duration) *LocationPrototypeUnclaimNodeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the location prototype unclaim node params
func (o *LocationPrototypeUnclaimNodeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the location prototype unclaim node params
func (o *LocationPrototypeUnclaimNodeParams) WithContext(ctx context.Context) *LocationPrototypeUnclaimNodeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the location prototype unclaim node params
func (o *LocationPrototypeUnclaimNodeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the location prototype unclaim node params
func (o *LocationPrototypeUnclaimNodeParams) WithHTTPClient(client *http.Client) *LocationPrototypeUnclaimNodeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the location prototype unclaim node params
func (o *LocationPrototypeUnclaimNodeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithForceUnclaim adds the forceUnclaim to the location prototype unclaim node params
func (o *LocationPrototypeUnclaimNodeParams) WithForceUnclaim(forceUnclaim *bool) *LocationPrototypeUnclaimNodeParams {
	o.SetForceUnclaim(forceUnclaim)
	return o
}

// SetForceUnclaim adds the forceUnclaim to the location prototype unclaim node params
func (o *LocationPrototypeUnclaimNodeParams) SetForceUnclaim(forceUnclaim *bool) {
	o.ForceUnclaim = forceUnclaim
}

// WithID adds the id to the location prototype unclaim node params
func (o *LocationPrototypeUnclaimNodeParams) WithID(id string) *LocationPrototypeUnclaimNodeParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the location prototype unclaim node params
func (o *LocationPrototypeUnclaimNodeParams) SetID(id string) {
	o.ID = id
}

// WithIncrementFactoryResetCounter adds the incrementFactoryResetCounter to the location prototype unclaim node params
func (o *LocationPrototypeUnclaimNodeParams) WithIncrementFactoryResetCounter(incrementFactoryResetCounter *bool) *LocationPrototypeUnclaimNodeParams {
	o.SetIncrementFactoryResetCounter(incrementFactoryResetCounter)
	return o
}

// SetIncrementFactoryResetCounter adds the incrementFactoryResetCounter to the location prototype unclaim node params
func (o *LocationPrototypeUnclaimNodeParams) SetIncrementFactoryResetCounter(incrementFactoryResetCounter *bool) {
	o.IncrementFactoryResetCounter = incrementFactoryResetCounter
}

// WithNodeID adds the nodeID to the location prototype unclaim node params
func (o *LocationPrototypeUnclaimNodeParams) WithNodeID(nodeID string) *LocationPrototypeUnclaimNodeParams {
	o.SetNodeID(nodeID)
	return o
}

// SetNodeID adds the nodeId to the location prototype unclaim node params
func (o *LocationPrototypeUnclaimNodeParams) SetNodeID(nodeID string) {
	o.NodeID = nodeID
}

// WithPreservePackID adds the preservePackID to the location prototype unclaim node params
func (o *LocationPrototypeUnclaimNodeParams) WithPreservePackID(preservePackID *bool) *LocationPrototypeUnclaimNodeParams {
	o.SetPreservePackID(preservePackID)
	return o
}

// SetPreservePackID adds the preservePackId to the location prototype unclaim node params
func (o *LocationPrototypeUnclaimNodeParams) SetPreservePackID(preservePackID *bool) {
	o.PreservePackID = preservePackID
}

// WithPurgeGroupIds adds the purgeGroupIds to the location prototype unclaim node params
func (o *LocationPrototypeUnclaimNodeParams) WithPurgeGroupIds(purgeGroupIds *bool) *LocationPrototypeUnclaimNodeParams {
	o.SetPurgeGroupIds(purgeGroupIds)
	return o
}

// SetPurgeGroupIds adds the purgeGroupIds to the location prototype unclaim node params
func (o *LocationPrototypeUnclaimNodeParams) SetPurgeGroupIds(purgeGroupIds *bool) {
	o.PurgeGroupIds = purgeGroupIds
}

// WithRemoveAccountID adds the removeAccountID to the location prototype unclaim node params
func (o *LocationPrototypeUnclaimNodeParams) WithRemoveAccountID(removeAccountID *bool) *LocationPrototypeUnclaimNodeParams {
	o.SetRemoveAccountID(removeAccountID)
	return o
}

// SetRemoveAccountID adds the removeAccountId to the location prototype unclaim node params
func (o *LocationPrototypeUnclaimNodeParams) SetRemoveAccountID(removeAccountID *bool) {
	o.RemoveAccountID = removeAccountID
}

// WithUnclaimReason adds the unclaimReason to the location prototype unclaim node params
func (o *LocationPrototypeUnclaimNodeParams) WithUnclaimReason(unclaimReason *string) *LocationPrototypeUnclaimNodeParams {
	o.SetUnclaimReason(unclaimReason)
	return o
}

// SetUnclaimReason adds the unclaimReason to the location prototype unclaim node params
func (o *LocationPrototypeUnclaimNodeParams) SetUnclaimReason(unclaimReason *string) {
	o.UnclaimReason = unclaimReason
}

// WriteToRequest writes these params to a swagger request
func (o *LocationPrototypeUnclaimNodeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ForceUnclaim != nil {

		// form param forceUnclaim
		var frForceUnclaim bool
		if o.ForceUnclaim != nil {
			frForceUnclaim = *o.ForceUnclaim
		}
		fForceUnclaim := swag.FormatBool(frForceUnclaim)
		if fForceUnclaim != "" {
			if err := r.SetFormParam("forceUnclaim", fForceUnclaim); err != nil {
				return err
			}
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.IncrementFactoryResetCounter != nil {

		// form param incrementFactoryResetCounter
		var frIncrementFactoryResetCounter bool
		if o.IncrementFactoryResetCounter != nil {
			frIncrementFactoryResetCounter = *o.IncrementFactoryResetCounter
		}
		fIncrementFactoryResetCounter := swag.FormatBool(frIncrementFactoryResetCounter)
		if fIncrementFactoryResetCounter != "" {
			if err := r.SetFormParam("incrementFactoryResetCounter", fIncrementFactoryResetCounter); err != nil {
				return err
			}
		}
	}

	// path param nodeId
	if err := r.SetPathParam("nodeId", o.NodeID); err != nil {
		return err
	}

	if o.PreservePackID != nil {

		// form param preservePackId
		var frPreservePackID bool
		if o.PreservePackID != nil {
			frPreservePackID = *o.PreservePackID
		}
		fPreservePackID := swag.FormatBool(frPreservePackID)
		if fPreservePackID != "" {
			if err := r.SetFormParam("preservePackId", fPreservePackID); err != nil {
				return err
			}
		}
	}

	if o.PurgeGroupIds != nil {

		// form param purgeGroupIds
		var frPurgeGroupIds bool
		if o.PurgeGroupIds != nil {
			frPurgeGroupIds = *o.PurgeGroupIds
		}
		fPurgeGroupIds := swag.FormatBool(frPurgeGroupIds)
		if fPurgeGroupIds != "" {
			if err := r.SetFormParam("purgeGroupIds", fPurgeGroupIds); err != nil {
				return err
			}
		}
	}

	if o.RemoveAccountID != nil {

		// form param removeAccountId
		var frRemoveAccountID bool
		if o.RemoveAccountID != nil {
			frRemoveAccountID = *o.RemoveAccountID
		}
		fRemoveAccountID := swag.FormatBool(frRemoveAccountID)
		if fRemoveAccountID != "" {
			if err := r.SetFormParam("removeAccountId", fRemoveAccountID); err != nil {
				return err
			}
		}
	}

	if o.UnclaimReason != nil {

		// form param unclaimReason
		var frUnclaimReason string
		if o.UnclaimReason != nil {
			frUnclaimReason = *o.UnclaimReason
		}
		fUnclaimReason := frUnclaimReason
		if fUnclaimReason != "" {
			if err := r.SetFormParam("unclaimReason", fUnclaimReason); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}