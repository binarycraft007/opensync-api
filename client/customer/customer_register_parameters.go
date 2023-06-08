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

// NewCustomerRegisterParams creates a new CustomerRegisterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerRegisterParams() *CustomerRegisterParams {
	return &CustomerRegisterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerRegisterParamsWithTimeout creates a new CustomerRegisterParams object
// with the ability to set a timeout on a request.
func NewCustomerRegisterParamsWithTimeout(timeout time.Duration) *CustomerRegisterParams {
	return &CustomerRegisterParams{
		timeout: timeout,
	}
}

// NewCustomerRegisterParamsWithContext creates a new CustomerRegisterParams object
// with the ability to set a context for a request.
func NewCustomerRegisterParamsWithContext(ctx context.Context) *CustomerRegisterParams {
	return &CustomerRegisterParams{
		Context: ctx,
	}
}

// NewCustomerRegisterParamsWithHTTPClient creates a new CustomerRegisterParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerRegisterParamsWithHTTPClient(client *http.Client) *CustomerRegisterParams {
	return &CustomerRegisterParams{
		HTTPClient: client,
	}
}

/*
CustomerRegisterParams contains all the parameters to send to the API endpoint

	for the customer register operation.

	Typically these are written to a http.Request.
*/
type CustomerRegisterParams struct {

	/* AcceptLanguage.

	   acceptable language
	*/
	AcceptLanguage *string

	/* AccountID.

	   must be unique, a UUID is recommended, min length is 6 characters.
	*/
	AccountID string

	// Email.
	Email *string

	/* Name.

	   Full name of customer, defaults to value of accountId
	*/
	Name *string

	/* OnboardingCheckpoint.

	   is the last passed onboarding step by the customer: 'PodsAdded' or 'OnboardingComplete';
	*/
	OnboardingCheckpoint *string

	/* PartnerID.

	   PartnerId of customer for accountId
	*/
	PartnerID *string

	/* Profile.

	   location profile
	*/
	Profile *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer register params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerRegisterParams) WithDefaults() *CustomerRegisterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer register params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerRegisterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer register params
func (o *CustomerRegisterParams) WithTimeout(timeout time.Duration) *CustomerRegisterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer register params
func (o *CustomerRegisterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer register params
func (o *CustomerRegisterParams) WithContext(ctx context.Context) *CustomerRegisterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer register params
func (o *CustomerRegisterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer register params
func (o *CustomerRegisterParams) WithHTTPClient(client *http.Client) *CustomerRegisterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer register params
func (o *CustomerRegisterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAcceptLanguage adds the acceptLanguage to the customer register params
func (o *CustomerRegisterParams) WithAcceptLanguage(acceptLanguage *string) *CustomerRegisterParams {
	o.SetAcceptLanguage(acceptLanguage)
	return o
}

// SetAcceptLanguage adds the acceptLanguage to the customer register params
func (o *CustomerRegisterParams) SetAcceptLanguage(acceptLanguage *string) {
	o.AcceptLanguage = acceptLanguage
}

// WithAccountID adds the accountID to the customer register params
func (o *CustomerRegisterParams) WithAccountID(accountID string) *CustomerRegisterParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the customer register params
func (o *CustomerRegisterParams) SetAccountID(accountID string) {
	o.AccountID = accountID
}

// WithEmail adds the email to the customer register params
func (o *CustomerRegisterParams) WithEmail(email *string) *CustomerRegisterParams {
	o.SetEmail(email)
	return o
}

// SetEmail adds the email to the customer register params
func (o *CustomerRegisterParams) SetEmail(email *string) {
	o.Email = email
}

// WithName adds the name to the customer register params
func (o *CustomerRegisterParams) WithName(name *string) *CustomerRegisterParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the customer register params
func (o *CustomerRegisterParams) SetName(name *string) {
	o.Name = name
}

// WithOnboardingCheckpoint adds the onboardingCheckpoint to the customer register params
func (o *CustomerRegisterParams) WithOnboardingCheckpoint(onboardingCheckpoint *string) *CustomerRegisterParams {
	o.SetOnboardingCheckpoint(onboardingCheckpoint)
	return o
}

// SetOnboardingCheckpoint adds the onboardingCheckpoint to the customer register params
func (o *CustomerRegisterParams) SetOnboardingCheckpoint(onboardingCheckpoint *string) {
	o.OnboardingCheckpoint = onboardingCheckpoint
}

// WithPartnerID adds the partnerID to the customer register params
func (o *CustomerRegisterParams) WithPartnerID(partnerID *string) *CustomerRegisterParams {
	o.SetPartnerID(partnerID)
	return o
}

// SetPartnerID adds the partnerId to the customer register params
func (o *CustomerRegisterParams) SetPartnerID(partnerID *string) {
	o.PartnerID = partnerID
}

// WithProfile adds the profile to the customer register params
func (o *CustomerRegisterParams) WithProfile(profile *string) *CustomerRegisterParams {
	o.SetProfile(profile)
	return o
}

// SetProfile adds the profile to the customer register params
func (o *CustomerRegisterParams) SetProfile(profile *string) {
	o.Profile = profile
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerRegisterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AcceptLanguage != nil {

		// form param acceptLanguage
		var frAcceptLanguage string
		if o.AcceptLanguage != nil {
			frAcceptLanguage = *o.AcceptLanguage
		}
		fAcceptLanguage := frAcceptLanguage
		if fAcceptLanguage != "" {
			if err := r.SetFormParam("acceptLanguage", fAcceptLanguage); err != nil {
				return err
			}
		}
	}

	// form param accountId
	frAccountID := o.AccountID
	fAccountID := frAccountID
	if fAccountID != "" {
		if err := r.SetFormParam("accountId", fAccountID); err != nil {
			return err
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

	if o.OnboardingCheckpoint != nil {

		// form param onboardingCheckpoint
		var frOnboardingCheckpoint string
		if o.OnboardingCheckpoint != nil {
			frOnboardingCheckpoint = *o.OnboardingCheckpoint
		}
		fOnboardingCheckpoint := frOnboardingCheckpoint
		if fOnboardingCheckpoint != "" {
			if err := r.SetFormParam("onboardingCheckpoint", fOnboardingCheckpoint); err != nil {
				return err
			}
		}
	}

	if o.PartnerID != nil {

		// form param partnerId
		var frPartnerID string
		if o.PartnerID != nil {
			frPartnerID = *o.PartnerID
		}
		fPartnerID := frPartnerID
		if fPartnerID != "" {
			if err := r.SetFormParam("partnerId", fPartnerID); err != nil {
				return err
			}
		}
	}

	if o.Profile != nil {

		// form param profile
		var frProfile string
		if o.Profile != nil {
			frProfile = *o.Profile
		}
		fProfile := frProfile
		if fProfile != "" {
			if err := r.SetFormParam("profile", fProfile); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
