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

// NewCustomerPrototypeGetPersonAppTimeCategoriesOnlineTimeParams creates a new CustomerPrototypeGetPersonAppTimeCategoriesOnlineTimeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypeGetPersonAppTimeCategoriesOnlineTimeParams() *CustomerPrototypeGetPersonAppTimeCategoriesOnlineTimeParams {
	return &CustomerPrototypeGetPersonAppTimeCategoriesOnlineTimeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypeGetPersonAppTimeCategoriesOnlineTimeParamsWithTimeout creates a new CustomerPrototypeGetPersonAppTimeCategoriesOnlineTimeParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypeGetPersonAppTimeCategoriesOnlineTimeParamsWithTimeout(timeout time.Duration) *CustomerPrototypeGetPersonAppTimeCategoriesOnlineTimeParams {
	return &CustomerPrototypeGetPersonAppTimeCategoriesOnlineTimeParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypeGetPersonAppTimeCategoriesOnlineTimeParamsWithContext creates a new CustomerPrototypeGetPersonAppTimeCategoriesOnlineTimeParams object
// with the ability to set a context for a request.
func NewCustomerPrototypeGetPersonAppTimeCategoriesOnlineTimeParamsWithContext(ctx context.Context) *CustomerPrototypeGetPersonAppTimeCategoriesOnlineTimeParams {
	return &CustomerPrototypeGetPersonAppTimeCategoriesOnlineTimeParams{
		Context: ctx,
	}
}

// NewCustomerPrototypeGetPersonAppTimeCategoriesOnlineTimeParamsWithHTTPClient creates a new CustomerPrototypeGetPersonAppTimeCategoriesOnlineTimeParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypeGetPersonAppTimeCategoriesOnlineTimeParamsWithHTTPClient(client *http.Client) *CustomerPrototypeGetPersonAppTimeCategoriesOnlineTimeParams {
	return &CustomerPrototypeGetPersonAppTimeCategoriesOnlineTimeParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypeGetPersonAppTimeCategoriesOnlineTimeParams contains all the parameters to send to the API endpoint

	for the customer prototype get person app time categories online time operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypeGetPersonAppTimeCategoriesOnlineTimeParams struct {

	/* Grouping.

	   typing of Grouping for the purposes of applying the limit. Can be: 'overall'|'perTimeSlot'

	   Default: "perTimeSlot"
	*/
	Grouping *string

	/* ID.

	   Customer id

	   Format: JSON
	*/
	ID string

	/* Limit.

	   Maximum number of categories to return. Defaults to 20

	   Format: double
	*/
	Limit *float64

	// LocationID.
	//
	// Format: JSON
	LocationID string

	// PersonID.
	PersonID string

	/* TimePeriod.

	   Any of "weekly","dailyToday","dailyYesterday","daily2DaysAgo","daily3DaysAgo","daily4DaysAgo","daily5DaysAgo","daily6DaysAgo","last30Days","last12Months"
	*/
	TimePeriod string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer prototype get person app time categories online time params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeGetPersonAppTimeCategoriesOnlineTimeParams) WithDefaults() *CustomerPrototypeGetPersonAppTimeCategoriesOnlineTimeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype get person app time categories online time params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeGetPersonAppTimeCategoriesOnlineTimeParams) SetDefaults() {
	var (
		groupingDefault = string("perTimeSlot")
	)

	val := CustomerPrototypeGetPersonAppTimeCategoriesOnlineTimeParams{
		Grouping: &groupingDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the customer prototype get person app time categories online time params
func (o *CustomerPrototypeGetPersonAppTimeCategoriesOnlineTimeParams) WithTimeout(timeout time.Duration) *CustomerPrototypeGetPersonAppTimeCategoriesOnlineTimeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype get person app time categories online time params
func (o *CustomerPrototypeGetPersonAppTimeCategoriesOnlineTimeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype get person app time categories online time params
func (o *CustomerPrototypeGetPersonAppTimeCategoriesOnlineTimeParams) WithContext(ctx context.Context) *CustomerPrototypeGetPersonAppTimeCategoriesOnlineTimeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype get person app time categories online time params
func (o *CustomerPrototypeGetPersonAppTimeCategoriesOnlineTimeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype get person app time categories online time params
func (o *CustomerPrototypeGetPersonAppTimeCategoriesOnlineTimeParams) WithHTTPClient(client *http.Client) *CustomerPrototypeGetPersonAppTimeCategoriesOnlineTimeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype get person app time categories online time params
func (o *CustomerPrototypeGetPersonAppTimeCategoriesOnlineTimeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGrouping adds the grouping to the customer prototype get person app time categories online time params
func (o *CustomerPrototypeGetPersonAppTimeCategoriesOnlineTimeParams) WithGrouping(grouping *string) *CustomerPrototypeGetPersonAppTimeCategoriesOnlineTimeParams {
	o.SetGrouping(grouping)
	return o
}

// SetGrouping adds the grouping to the customer prototype get person app time categories online time params
func (o *CustomerPrototypeGetPersonAppTimeCategoriesOnlineTimeParams) SetGrouping(grouping *string) {
	o.Grouping = grouping
}

// WithID adds the id to the customer prototype get person app time categories online time params
func (o *CustomerPrototypeGetPersonAppTimeCategoriesOnlineTimeParams) WithID(id string) *CustomerPrototypeGetPersonAppTimeCategoriesOnlineTimeParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype get person app time categories online time params
func (o *CustomerPrototypeGetPersonAppTimeCategoriesOnlineTimeParams) SetID(id string) {
	o.ID = id
}

// WithLimit adds the limit to the customer prototype get person app time categories online time params
func (o *CustomerPrototypeGetPersonAppTimeCategoriesOnlineTimeParams) WithLimit(limit *float64) *CustomerPrototypeGetPersonAppTimeCategoriesOnlineTimeParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the customer prototype get person app time categories online time params
func (o *CustomerPrototypeGetPersonAppTimeCategoriesOnlineTimeParams) SetLimit(limit *float64) {
	o.Limit = limit
}

// WithLocationID adds the locationID to the customer prototype get person app time categories online time params
func (o *CustomerPrototypeGetPersonAppTimeCategoriesOnlineTimeParams) WithLocationID(locationID string) *CustomerPrototypeGetPersonAppTimeCategoriesOnlineTimeParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype get person app time categories online time params
func (o *CustomerPrototypeGetPersonAppTimeCategoriesOnlineTimeParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WithPersonID adds the personID to the customer prototype get person app time categories online time params
func (o *CustomerPrototypeGetPersonAppTimeCategoriesOnlineTimeParams) WithPersonID(personID string) *CustomerPrototypeGetPersonAppTimeCategoriesOnlineTimeParams {
	o.SetPersonID(personID)
	return o
}

// SetPersonID adds the personId to the customer prototype get person app time categories online time params
func (o *CustomerPrototypeGetPersonAppTimeCategoriesOnlineTimeParams) SetPersonID(personID string) {
	o.PersonID = personID
}

// WithTimePeriod adds the timePeriod to the customer prototype get person app time categories online time params
func (o *CustomerPrototypeGetPersonAppTimeCategoriesOnlineTimeParams) WithTimePeriod(timePeriod string) *CustomerPrototypeGetPersonAppTimeCategoriesOnlineTimeParams {
	o.SetTimePeriod(timePeriod)
	return o
}

// SetTimePeriod adds the timePeriod to the customer prototype get person app time categories online time params
func (o *CustomerPrototypeGetPersonAppTimeCategoriesOnlineTimeParams) SetTimePeriod(timePeriod string) {
	o.TimePeriod = timePeriod
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypeGetPersonAppTimeCategoriesOnlineTimeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Grouping != nil {

		// query param grouping
		var qrGrouping string

		if o.Grouping != nil {
			qrGrouping = *o.Grouping
		}
		qGrouping := qrGrouping
		if qGrouping != "" {

			if err := r.SetQueryParam("grouping", qGrouping); err != nil {
				return err
			}
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit float64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatFloat64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	// path param locationId
	if err := r.SetPathParam("locationId", o.LocationID); err != nil {
		return err
	}

	// path param personId
	if err := r.SetPathParam("personId", o.PersonID); err != nil {
		return err
	}

	// query param timePeriod
	qrTimePeriod := o.TimePeriod
	qTimePeriod := qrTimePeriod
	if qTimePeriod != "" {

		if err := r.SetQueryParam("timePeriod", qTimePeriod); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
