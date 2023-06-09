// Code generated by go-swagger; DO NOT EDIT.

package partner_config

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

// NewPartnerConfigPatchSecurityPolicyParams creates a new PartnerConfigPatchSecurityPolicyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPartnerConfigPatchSecurityPolicyParams() *PartnerConfigPatchSecurityPolicyParams {
	return &PartnerConfigPatchSecurityPolicyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPartnerConfigPatchSecurityPolicyParamsWithTimeout creates a new PartnerConfigPatchSecurityPolicyParams object
// with the ability to set a timeout on a request.
func NewPartnerConfigPatchSecurityPolicyParamsWithTimeout(timeout time.Duration) *PartnerConfigPatchSecurityPolicyParams {
	return &PartnerConfigPatchSecurityPolicyParams{
		timeout: timeout,
	}
}

// NewPartnerConfigPatchSecurityPolicyParamsWithContext creates a new PartnerConfigPatchSecurityPolicyParams object
// with the ability to set a context for a request.
func NewPartnerConfigPatchSecurityPolicyParamsWithContext(ctx context.Context) *PartnerConfigPatchSecurityPolicyParams {
	return &PartnerConfigPatchSecurityPolicyParams{
		Context: ctx,
	}
}

// NewPartnerConfigPatchSecurityPolicyParamsWithHTTPClient creates a new PartnerConfigPatchSecurityPolicyParams object
// with the ability to set a custom HTTPClient for a request.
func NewPartnerConfigPatchSecurityPolicyParamsWithHTTPClient(client *http.Client) *PartnerConfigPatchSecurityPolicyParams {
	return &PartnerConfigPatchSecurityPolicyParams{
		HTTPClient: client,
	}
}

/*
PartnerConfigPatchSecurityPolicyParams contains all the parameters to send to the API endpoint

	for the partner config patch security policy operation.

	Typically these are written to a http.Request.
*/
type PartnerConfigPatchSecurityPolicyParams struct {

	/* Blacklist.

	   List of blacklisted fqdn values - {type="fqdn", value}

	   Format: JSON
	*/
	Blacklist *string

	/* BlockedDNSRedirect.

	   Redirect "address" and corresponding "ttl" redirect in seconds

	   Format: JSON
	*/
	BlockedDNSRedirect *string

	/* ContentFilterOverrides.

	   Array of content filter overrides each with its own whitelist/blacklist

	   Format: JSON
	*/
	ContentFilterOverrides *string

	// DisableSafeSearch.
	DisableSafeSearch *bool

	// ID.
	ID string

	// RemoteConnectionsForceDisable.
	RemoteConnectionsForceDisable *bool

	/* RemoteConnectionsMode.

	   Any of "auto", "enabled", "disabled", "highRiskOnly"
	*/
	RemoteConnectionsMode *string

	/* Whitelist.

	   List of whitelisted fqdn values - {type="fqdn", value}"

	   Format: JSON
	*/
	Whitelist *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the partner config patch security policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PartnerConfigPatchSecurityPolicyParams) WithDefaults() *PartnerConfigPatchSecurityPolicyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the partner config patch security policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PartnerConfigPatchSecurityPolicyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the partner config patch security policy params
func (o *PartnerConfigPatchSecurityPolicyParams) WithTimeout(timeout time.Duration) *PartnerConfigPatchSecurityPolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the partner config patch security policy params
func (o *PartnerConfigPatchSecurityPolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the partner config patch security policy params
func (o *PartnerConfigPatchSecurityPolicyParams) WithContext(ctx context.Context) *PartnerConfigPatchSecurityPolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the partner config patch security policy params
func (o *PartnerConfigPatchSecurityPolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the partner config patch security policy params
func (o *PartnerConfigPatchSecurityPolicyParams) WithHTTPClient(client *http.Client) *PartnerConfigPatchSecurityPolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the partner config patch security policy params
func (o *PartnerConfigPatchSecurityPolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBlacklist adds the blacklist to the partner config patch security policy params
func (o *PartnerConfigPatchSecurityPolicyParams) WithBlacklist(blacklist *string) *PartnerConfigPatchSecurityPolicyParams {
	o.SetBlacklist(blacklist)
	return o
}

// SetBlacklist adds the blacklist to the partner config patch security policy params
func (o *PartnerConfigPatchSecurityPolicyParams) SetBlacklist(blacklist *string) {
	o.Blacklist = blacklist
}

// WithBlockedDNSRedirect adds the blockedDNSRedirect to the partner config patch security policy params
func (o *PartnerConfigPatchSecurityPolicyParams) WithBlockedDNSRedirect(blockedDNSRedirect *string) *PartnerConfigPatchSecurityPolicyParams {
	o.SetBlockedDNSRedirect(blockedDNSRedirect)
	return o
}

// SetBlockedDNSRedirect adds the blockedDnsRedirect to the partner config patch security policy params
func (o *PartnerConfigPatchSecurityPolicyParams) SetBlockedDNSRedirect(blockedDNSRedirect *string) {
	o.BlockedDNSRedirect = blockedDNSRedirect
}

// WithContentFilterOverrides adds the contentFilterOverrides to the partner config patch security policy params
func (o *PartnerConfigPatchSecurityPolicyParams) WithContentFilterOverrides(contentFilterOverrides *string) *PartnerConfigPatchSecurityPolicyParams {
	o.SetContentFilterOverrides(contentFilterOverrides)
	return o
}

// SetContentFilterOverrides adds the contentFilterOverrides to the partner config patch security policy params
func (o *PartnerConfigPatchSecurityPolicyParams) SetContentFilterOverrides(contentFilterOverrides *string) {
	o.ContentFilterOverrides = contentFilterOverrides
}

// WithDisableSafeSearch adds the disableSafeSearch to the partner config patch security policy params
func (o *PartnerConfigPatchSecurityPolicyParams) WithDisableSafeSearch(disableSafeSearch *bool) *PartnerConfigPatchSecurityPolicyParams {
	o.SetDisableSafeSearch(disableSafeSearch)
	return o
}

// SetDisableSafeSearch adds the disableSafeSearch to the partner config patch security policy params
func (o *PartnerConfigPatchSecurityPolicyParams) SetDisableSafeSearch(disableSafeSearch *bool) {
	o.DisableSafeSearch = disableSafeSearch
}

// WithID adds the id to the partner config patch security policy params
func (o *PartnerConfigPatchSecurityPolicyParams) WithID(id string) *PartnerConfigPatchSecurityPolicyParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the partner config patch security policy params
func (o *PartnerConfigPatchSecurityPolicyParams) SetID(id string) {
	o.ID = id
}

// WithRemoteConnectionsForceDisable adds the remoteConnectionsForceDisable to the partner config patch security policy params
func (o *PartnerConfigPatchSecurityPolicyParams) WithRemoteConnectionsForceDisable(remoteConnectionsForceDisable *bool) *PartnerConfigPatchSecurityPolicyParams {
	o.SetRemoteConnectionsForceDisable(remoteConnectionsForceDisable)
	return o
}

// SetRemoteConnectionsForceDisable adds the remoteConnectionsForceDisable to the partner config patch security policy params
func (o *PartnerConfigPatchSecurityPolicyParams) SetRemoteConnectionsForceDisable(remoteConnectionsForceDisable *bool) {
	o.RemoteConnectionsForceDisable = remoteConnectionsForceDisable
}

// WithRemoteConnectionsMode adds the remoteConnectionsMode to the partner config patch security policy params
func (o *PartnerConfigPatchSecurityPolicyParams) WithRemoteConnectionsMode(remoteConnectionsMode *string) *PartnerConfigPatchSecurityPolicyParams {
	o.SetRemoteConnectionsMode(remoteConnectionsMode)
	return o
}

// SetRemoteConnectionsMode adds the remoteConnectionsMode to the partner config patch security policy params
func (o *PartnerConfigPatchSecurityPolicyParams) SetRemoteConnectionsMode(remoteConnectionsMode *string) {
	o.RemoteConnectionsMode = remoteConnectionsMode
}

// WithWhitelist adds the whitelist to the partner config patch security policy params
func (o *PartnerConfigPatchSecurityPolicyParams) WithWhitelist(whitelist *string) *PartnerConfigPatchSecurityPolicyParams {
	o.SetWhitelist(whitelist)
	return o
}

// SetWhitelist adds the whitelist to the partner config patch security policy params
func (o *PartnerConfigPatchSecurityPolicyParams) SetWhitelist(whitelist *string) {
	o.Whitelist = whitelist
}

// WriteToRequest writes these params to a swagger request
func (o *PartnerConfigPatchSecurityPolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Blacklist != nil {

		// form param blacklist
		var frBlacklist string
		if o.Blacklist != nil {
			frBlacklist = *o.Blacklist
		}
		fBlacklist := frBlacklist
		if fBlacklist != "" {
			if err := r.SetFormParam("blacklist", fBlacklist); err != nil {
				return err
			}
		}
	}

	if o.BlockedDNSRedirect != nil {

		// form param blockedDnsRedirect
		var frBlockedDNSRedirect string
		if o.BlockedDNSRedirect != nil {
			frBlockedDNSRedirect = *o.BlockedDNSRedirect
		}
		fBlockedDNSRedirect := frBlockedDNSRedirect
		if fBlockedDNSRedirect != "" {
			if err := r.SetFormParam("blockedDnsRedirect", fBlockedDNSRedirect); err != nil {
				return err
			}
		}
	}

	if o.ContentFilterOverrides != nil {

		// form param contentFilterOverrides
		var frContentFilterOverrides string
		if o.ContentFilterOverrides != nil {
			frContentFilterOverrides = *o.ContentFilterOverrides
		}
		fContentFilterOverrides := frContentFilterOverrides
		if fContentFilterOverrides != "" {
			if err := r.SetFormParam("contentFilterOverrides", fContentFilterOverrides); err != nil {
				return err
			}
		}
	}

	if o.DisableSafeSearch != nil {

		// form param disableSafeSearch
		var frDisableSafeSearch bool
		if o.DisableSafeSearch != nil {
			frDisableSafeSearch = *o.DisableSafeSearch
		}
		fDisableSafeSearch := swag.FormatBool(frDisableSafeSearch)
		if fDisableSafeSearch != "" {
			if err := r.SetFormParam("disableSafeSearch", fDisableSafeSearch); err != nil {
				return err
			}
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.RemoteConnectionsForceDisable != nil {

		// form param remoteConnectionsForceDisable
		var frRemoteConnectionsForceDisable bool
		if o.RemoteConnectionsForceDisable != nil {
			frRemoteConnectionsForceDisable = *o.RemoteConnectionsForceDisable
		}
		fRemoteConnectionsForceDisable := swag.FormatBool(frRemoteConnectionsForceDisable)
		if fRemoteConnectionsForceDisable != "" {
			if err := r.SetFormParam("remoteConnectionsForceDisable", fRemoteConnectionsForceDisable); err != nil {
				return err
			}
		}
	}

	if o.RemoteConnectionsMode != nil {

		// form param remoteConnectionsMode
		var frRemoteConnectionsMode string
		if o.RemoteConnectionsMode != nil {
			frRemoteConnectionsMode = *o.RemoteConnectionsMode
		}
		fRemoteConnectionsMode := frRemoteConnectionsMode
		if fRemoteConnectionsMode != "" {
			if err := r.SetFormParam("remoteConnectionsMode", fRemoteConnectionsMode); err != nil {
				return err
			}
		}
	}

	if o.Whitelist != nil {

		// form param whitelist
		var frWhitelist string
		if o.Whitelist != nil {
			frWhitelist = *o.Whitelist
		}
		fWhitelist := frWhitelist
		if fWhitelist != "" {
			if err := r.SetFormParam("whitelist", fWhitelist); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
