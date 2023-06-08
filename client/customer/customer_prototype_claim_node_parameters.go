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

// NewCustomerPrototypeClaimNodeParams creates a new CustomerPrototypeClaimNodeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerPrototypeClaimNodeParams() *CustomerPrototypeClaimNodeParams {
	return &CustomerPrototypeClaimNodeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerPrototypeClaimNodeParamsWithTimeout creates a new CustomerPrototypeClaimNodeParams object
// with the ability to set a timeout on a request.
func NewCustomerPrototypeClaimNodeParamsWithTimeout(timeout time.Duration) *CustomerPrototypeClaimNodeParams {
	return &CustomerPrototypeClaimNodeParams{
		timeout: timeout,
	}
}

// NewCustomerPrototypeClaimNodeParamsWithContext creates a new CustomerPrototypeClaimNodeParams object
// with the ability to set a context for a request.
func NewCustomerPrototypeClaimNodeParamsWithContext(ctx context.Context) *CustomerPrototypeClaimNodeParams {
	return &CustomerPrototypeClaimNodeParams{
		Context: ctx,
	}
}

// NewCustomerPrototypeClaimNodeParamsWithHTTPClient creates a new CustomerPrototypeClaimNodeParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerPrototypeClaimNodeParamsWithHTTPClient(client *http.Client) *CustomerPrototypeClaimNodeParams {
	return &CustomerPrototypeClaimNodeParams{
		HTTPClient: client,
	}
}

/*
CustomerPrototypeClaimNodeParams contains all the parameters to send to the API endpoint

	for the customer prototype claim node operation.

	Typically these are written to a http.Request.
*/
type CustomerPrototypeClaimNodeParams struct {

	/* BackhaulDhcpPoolIdx.

	   optional node backhaulDhcpPoolIdx

	   Format: double
	*/
	BackhaulDhcpPoolIdx *float64

	/* ClaimKey.

	   optional but required for auto-importing, must be a valid claimKey
	*/
	ClaimKey *string

	/* Ethernet1Mac.

	   optional but required for auto-importing, must be a valid mac address
	*/
	Ethernet1Mac *string

	/* EthernetMac.

	   optional but required for auto-importing, must be a valid mac address
	*/
	EthernetMac *string

	/* HybridCheck.

	   optional when auto-importing, ignored otherwise
	*/
	HybridCheck *bool

	/* ID.

	   Customer id

	   Format: JSON
	*/
	ID string

	/* LocationID.

	   location id
	*/
	LocationID string

	/* Model.

	   optional when auto-importing, ignored otherwise
	*/
	Model *string

	/* Nickname.

	   optional node nickname
	*/
	Nickname *string

	/* RadioMac24.

	   optional but required for auto-importing, must be a valid mac address
	*/
	RadioMac24 *string

	/* RadioMac50.

	   optional but required for auto-importing, must be a valid mac address
	*/
	RadioMac50 *string

	/* RadioMac60.

	   optional but required for auto-importing, must be a valid mac address
	*/
	RadioMac60 *string

	/* Room.

	   optional room identifier
	*/
	Room *string

	/* SerialNumber.

	   unique serial number or ID of Node
	*/
	SerialNumber *string

	/* SkipSubscription.

	   skip subscription update
	*/
	SkipSubscription *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer prototype claim node params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeClaimNodeParams) WithDefaults() *CustomerPrototypeClaimNodeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer prototype claim node params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerPrototypeClaimNodeParams) SetDefaults() {
	var (
		skipSubscriptionDefault = bool(false)
	)

	val := CustomerPrototypeClaimNodeParams{
		SkipSubscription: &skipSubscriptionDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the customer prototype claim node params
func (o *CustomerPrototypeClaimNodeParams) WithTimeout(timeout time.Duration) *CustomerPrototypeClaimNodeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer prototype claim node params
func (o *CustomerPrototypeClaimNodeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer prototype claim node params
func (o *CustomerPrototypeClaimNodeParams) WithContext(ctx context.Context) *CustomerPrototypeClaimNodeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer prototype claim node params
func (o *CustomerPrototypeClaimNodeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer prototype claim node params
func (o *CustomerPrototypeClaimNodeParams) WithHTTPClient(client *http.Client) *CustomerPrototypeClaimNodeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer prototype claim node params
func (o *CustomerPrototypeClaimNodeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBackhaulDhcpPoolIdx adds the backhaulDhcpPoolIdx to the customer prototype claim node params
func (o *CustomerPrototypeClaimNodeParams) WithBackhaulDhcpPoolIdx(backhaulDhcpPoolIdx *float64) *CustomerPrototypeClaimNodeParams {
	o.SetBackhaulDhcpPoolIdx(backhaulDhcpPoolIdx)
	return o
}

// SetBackhaulDhcpPoolIdx adds the backhaulDhcpPoolIdx to the customer prototype claim node params
func (o *CustomerPrototypeClaimNodeParams) SetBackhaulDhcpPoolIdx(backhaulDhcpPoolIdx *float64) {
	o.BackhaulDhcpPoolIdx = backhaulDhcpPoolIdx
}

// WithClaimKey adds the claimKey to the customer prototype claim node params
func (o *CustomerPrototypeClaimNodeParams) WithClaimKey(claimKey *string) *CustomerPrototypeClaimNodeParams {
	o.SetClaimKey(claimKey)
	return o
}

// SetClaimKey adds the claimKey to the customer prototype claim node params
func (o *CustomerPrototypeClaimNodeParams) SetClaimKey(claimKey *string) {
	o.ClaimKey = claimKey
}

// WithEthernet1Mac adds the ethernet1Mac to the customer prototype claim node params
func (o *CustomerPrototypeClaimNodeParams) WithEthernet1Mac(ethernet1Mac *string) *CustomerPrototypeClaimNodeParams {
	o.SetEthernet1Mac(ethernet1Mac)
	return o
}

// SetEthernet1Mac adds the ethernet1Mac to the customer prototype claim node params
func (o *CustomerPrototypeClaimNodeParams) SetEthernet1Mac(ethernet1Mac *string) {
	o.Ethernet1Mac = ethernet1Mac
}

// WithEthernetMac adds the ethernetMac to the customer prototype claim node params
func (o *CustomerPrototypeClaimNodeParams) WithEthernetMac(ethernetMac *string) *CustomerPrototypeClaimNodeParams {
	o.SetEthernetMac(ethernetMac)
	return o
}

// SetEthernetMac adds the ethernetMac to the customer prototype claim node params
func (o *CustomerPrototypeClaimNodeParams) SetEthernetMac(ethernetMac *string) {
	o.EthernetMac = ethernetMac
}

// WithHybridCheck adds the hybridCheck to the customer prototype claim node params
func (o *CustomerPrototypeClaimNodeParams) WithHybridCheck(hybridCheck *bool) *CustomerPrototypeClaimNodeParams {
	o.SetHybridCheck(hybridCheck)
	return o
}

// SetHybridCheck adds the hybridCheck to the customer prototype claim node params
func (o *CustomerPrototypeClaimNodeParams) SetHybridCheck(hybridCheck *bool) {
	o.HybridCheck = hybridCheck
}

// WithID adds the id to the customer prototype claim node params
func (o *CustomerPrototypeClaimNodeParams) WithID(id string) *CustomerPrototypeClaimNodeParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer prototype claim node params
func (o *CustomerPrototypeClaimNodeParams) SetID(id string) {
	o.ID = id
}

// WithLocationID adds the locationID to the customer prototype claim node params
func (o *CustomerPrototypeClaimNodeParams) WithLocationID(locationID string) *CustomerPrototypeClaimNodeParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the customer prototype claim node params
func (o *CustomerPrototypeClaimNodeParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WithModel adds the model to the customer prototype claim node params
func (o *CustomerPrototypeClaimNodeParams) WithModel(model *string) *CustomerPrototypeClaimNodeParams {
	o.SetModel(model)
	return o
}

// SetModel adds the model to the customer prototype claim node params
func (o *CustomerPrototypeClaimNodeParams) SetModel(model *string) {
	o.Model = model
}

// WithNickname adds the nickname to the customer prototype claim node params
func (o *CustomerPrototypeClaimNodeParams) WithNickname(nickname *string) *CustomerPrototypeClaimNodeParams {
	o.SetNickname(nickname)
	return o
}

// SetNickname adds the nickname to the customer prototype claim node params
func (o *CustomerPrototypeClaimNodeParams) SetNickname(nickname *string) {
	o.Nickname = nickname
}

// WithRadioMac24 adds the radioMac24 to the customer prototype claim node params
func (o *CustomerPrototypeClaimNodeParams) WithRadioMac24(radioMac24 *string) *CustomerPrototypeClaimNodeParams {
	o.SetRadioMac24(radioMac24)
	return o
}

// SetRadioMac24 adds the radioMac24 to the customer prototype claim node params
func (o *CustomerPrototypeClaimNodeParams) SetRadioMac24(radioMac24 *string) {
	o.RadioMac24 = radioMac24
}

// WithRadioMac50 adds the radioMac50 to the customer prototype claim node params
func (o *CustomerPrototypeClaimNodeParams) WithRadioMac50(radioMac50 *string) *CustomerPrototypeClaimNodeParams {
	o.SetRadioMac50(radioMac50)
	return o
}

// SetRadioMac50 adds the radioMac50 to the customer prototype claim node params
func (o *CustomerPrototypeClaimNodeParams) SetRadioMac50(radioMac50 *string) {
	o.RadioMac50 = radioMac50
}

// WithRadioMac60 adds the radioMac60 to the customer prototype claim node params
func (o *CustomerPrototypeClaimNodeParams) WithRadioMac60(radioMac60 *string) *CustomerPrototypeClaimNodeParams {
	o.SetRadioMac60(radioMac60)
	return o
}

// SetRadioMac60 adds the radioMac60 to the customer prototype claim node params
func (o *CustomerPrototypeClaimNodeParams) SetRadioMac60(radioMac60 *string) {
	o.RadioMac60 = radioMac60
}

// WithRoom adds the room to the customer prototype claim node params
func (o *CustomerPrototypeClaimNodeParams) WithRoom(room *string) *CustomerPrototypeClaimNodeParams {
	o.SetRoom(room)
	return o
}

// SetRoom adds the room to the customer prototype claim node params
func (o *CustomerPrototypeClaimNodeParams) SetRoom(room *string) {
	o.Room = room
}

// WithSerialNumber adds the serialNumber to the customer prototype claim node params
func (o *CustomerPrototypeClaimNodeParams) WithSerialNumber(serialNumber *string) *CustomerPrototypeClaimNodeParams {
	o.SetSerialNumber(serialNumber)
	return o
}

// SetSerialNumber adds the serialNumber to the customer prototype claim node params
func (o *CustomerPrototypeClaimNodeParams) SetSerialNumber(serialNumber *string) {
	o.SerialNumber = serialNumber
}

// WithSkipSubscription adds the skipSubscription to the customer prototype claim node params
func (o *CustomerPrototypeClaimNodeParams) WithSkipSubscription(skipSubscription *bool) *CustomerPrototypeClaimNodeParams {
	o.SetSkipSubscription(skipSubscription)
	return o
}

// SetSkipSubscription adds the skipSubscription to the customer prototype claim node params
func (o *CustomerPrototypeClaimNodeParams) SetSkipSubscription(skipSubscription *bool) {
	o.SkipSubscription = skipSubscription
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerPrototypeClaimNodeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.BackhaulDhcpPoolIdx != nil {

		// form param backhaulDhcpPoolIdx
		var frBackhaulDhcpPoolIdx float64
		if o.BackhaulDhcpPoolIdx != nil {
			frBackhaulDhcpPoolIdx = *o.BackhaulDhcpPoolIdx
		}
		fBackhaulDhcpPoolIdx := swag.FormatFloat64(frBackhaulDhcpPoolIdx)
		if fBackhaulDhcpPoolIdx != "" {
			if err := r.SetFormParam("backhaulDhcpPoolIdx", fBackhaulDhcpPoolIdx); err != nil {
				return err
			}
		}
	}

	if o.ClaimKey != nil {

		// form param claimKey
		var frClaimKey string
		if o.ClaimKey != nil {
			frClaimKey = *o.ClaimKey
		}
		fClaimKey := frClaimKey
		if fClaimKey != "" {
			if err := r.SetFormParam("claimKey", fClaimKey); err != nil {
				return err
			}
		}
	}

	if o.Ethernet1Mac != nil {

		// form param ethernet1Mac
		var frEthernet1Mac string
		if o.Ethernet1Mac != nil {
			frEthernet1Mac = *o.Ethernet1Mac
		}
		fEthernet1Mac := frEthernet1Mac
		if fEthernet1Mac != "" {
			if err := r.SetFormParam("ethernet1Mac", fEthernet1Mac); err != nil {
				return err
			}
		}
	}

	if o.EthernetMac != nil {

		// form param ethernetMac
		var frEthernetMac string
		if o.EthernetMac != nil {
			frEthernetMac = *o.EthernetMac
		}
		fEthernetMac := frEthernetMac
		if fEthernetMac != "" {
			if err := r.SetFormParam("ethernetMac", fEthernetMac); err != nil {
				return err
			}
		}
	}

	if o.HybridCheck != nil {

		// form param hybridCheck
		var frHybridCheck bool
		if o.HybridCheck != nil {
			frHybridCheck = *o.HybridCheck
		}
		fHybridCheck := swag.FormatBool(frHybridCheck)
		if fHybridCheck != "" {
			if err := r.SetFormParam("hybridCheck", fHybridCheck); err != nil {
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

	if o.RadioMac24 != nil {

		// form param radioMac24
		var frRadioMac24 string
		if o.RadioMac24 != nil {
			frRadioMac24 = *o.RadioMac24
		}
		fRadioMac24 := frRadioMac24
		if fRadioMac24 != "" {
			if err := r.SetFormParam("radioMac24", fRadioMac24); err != nil {
				return err
			}
		}
	}

	if o.RadioMac50 != nil {

		// form param radioMac50
		var frRadioMac50 string
		if o.RadioMac50 != nil {
			frRadioMac50 = *o.RadioMac50
		}
		fRadioMac50 := frRadioMac50
		if fRadioMac50 != "" {
			if err := r.SetFormParam("radioMac50", fRadioMac50); err != nil {
				return err
			}
		}
	}

	if o.RadioMac60 != nil {

		// form param radioMac60
		var frRadioMac60 string
		if o.RadioMac60 != nil {
			frRadioMac60 = *o.RadioMac60
		}
		fRadioMac60 := frRadioMac60
		if fRadioMac60 != "" {
			if err := r.SetFormParam("radioMac60", fRadioMac60); err != nil {
				return err
			}
		}
	}

	if o.Room != nil {

		// form param room
		var frRoom string
		if o.Room != nil {
			frRoom = *o.Room
		}
		fRoom := frRoom
		if fRoom != "" {
			if err := r.SetFormParam("room", fRoom); err != nil {
				return err
			}
		}
	}

	if o.SerialNumber != nil {

		// form param serialNumber
		var frSerialNumber string
		if o.SerialNumber != nil {
			frSerialNumber = *o.SerialNumber
		}
		fSerialNumber := frSerialNumber
		if fSerialNumber != "" {
			if err := r.SetFormParam("serialNumber", fSerialNumber); err != nil {
				return err
			}
		}
	}

	if o.SkipSubscription != nil {

		// query param skipSubscription
		var qrSkipSubscription bool

		if o.SkipSubscription != nil {
			qrSkipSubscription = *o.SkipSubscription
		}
		qSkipSubscription := swag.FormatBool(qrSkipSubscription)
		if qSkipSubscription != "" {

			if err := r.SetQueryParam("skipSubscription", qSkipSubscription); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}