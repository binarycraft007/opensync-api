// Code generated by go-swagger; DO NOT EDIT.

package customer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/binarycraft007/opensync-api/models"
)

// CustomerPrototypePatchCampaignCaptivePortalBrandingReader is a Reader for the CustomerPrototypePatchCampaignCaptivePortalBranding structure.
type CustomerPrototypePatchCampaignCaptivePortalBrandingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypePatchCampaignCaptivePortalBrandingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypePatchCampaignCaptivePortalBrandingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypePatchCampaignCaptivePortalBrandingDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypePatchCampaignCaptivePortalBrandingOK creates a CustomerPrototypePatchCampaignCaptivePortalBrandingOK with default headers values
func NewCustomerPrototypePatchCampaignCaptivePortalBrandingOK() *CustomerPrototypePatchCampaignCaptivePortalBrandingOK {
	return &CustomerPrototypePatchCampaignCaptivePortalBrandingOK{}
}

/*
CustomerPrototypePatchCampaignCaptivePortalBrandingOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypePatchCampaignCaptivePortalBrandingOK struct {
	Payload models.NetworkConfig
}

// IsSuccess returns true when this customer prototype patch campaign captive portal branding o k response has a 2xx status code
func (o *CustomerPrototypePatchCampaignCaptivePortalBrandingOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype patch campaign captive portal branding o k response has a 3xx status code
func (o *CustomerPrototypePatchCampaignCaptivePortalBrandingOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype patch campaign captive portal branding o k response has a 4xx status code
func (o *CustomerPrototypePatchCampaignCaptivePortalBrandingOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype patch campaign captive portal branding o k response has a 5xx status code
func (o *CustomerPrototypePatchCampaignCaptivePortalBrandingOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype patch campaign captive portal branding o k response a status code equal to that given
func (o *CustomerPrototypePatchCampaignCaptivePortalBrandingOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype patch campaign captive portal branding o k response
func (o *CustomerPrototypePatchCampaignCaptivePortalBrandingOK) Code() int {
	return 200
}

func (o *CustomerPrototypePatchCampaignCaptivePortalBrandingOK) Error() string {
	return fmt.Sprintf("[PATCH /Customers/{id}/locations/{locationId}/secondaryNetworks/captivePortals/{networkId}/campaign/branding][%d] customerPrototypePatchCampaignCaptivePortalBrandingOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypePatchCampaignCaptivePortalBrandingOK) String() string {
	return fmt.Sprintf("[PATCH /Customers/{id}/locations/{locationId}/secondaryNetworks/captivePortals/{networkId}/campaign/branding][%d] customerPrototypePatchCampaignCaptivePortalBrandingOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypePatchCampaignCaptivePortalBrandingOK) GetPayload() models.NetworkConfig {
	return o.Payload
}

func (o *CustomerPrototypePatchCampaignCaptivePortalBrandingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypePatchCampaignCaptivePortalBrandingDefault creates a CustomerPrototypePatchCampaignCaptivePortalBrandingDefault with default headers values
func NewCustomerPrototypePatchCampaignCaptivePortalBrandingDefault(code int) *CustomerPrototypePatchCampaignCaptivePortalBrandingDefault {
	return &CustomerPrototypePatchCampaignCaptivePortalBrandingDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypePatchCampaignCaptivePortalBrandingDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypePatchCampaignCaptivePortalBrandingDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype patch campaign captive portal branding default response has a 2xx status code
func (o *CustomerPrototypePatchCampaignCaptivePortalBrandingDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype patch campaign captive portal branding default response has a 3xx status code
func (o *CustomerPrototypePatchCampaignCaptivePortalBrandingDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype patch campaign captive portal branding default response has a 4xx status code
func (o *CustomerPrototypePatchCampaignCaptivePortalBrandingDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype patch campaign captive portal branding default response has a 5xx status code
func (o *CustomerPrototypePatchCampaignCaptivePortalBrandingDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype patch campaign captive portal branding default response a status code equal to that given
func (o *CustomerPrototypePatchCampaignCaptivePortalBrandingDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype patch campaign captive portal branding default response
func (o *CustomerPrototypePatchCampaignCaptivePortalBrandingDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypePatchCampaignCaptivePortalBrandingDefault) Error() string {
	return fmt.Sprintf("[PATCH /Customers/{id}/locations/{locationId}/secondaryNetworks/captivePortals/{networkId}/campaign/branding][%d] Customer.prototype.patchCampaignCaptivePortalBranding default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePatchCampaignCaptivePortalBrandingDefault) String() string {
	return fmt.Sprintf("[PATCH /Customers/{id}/locations/{locationId}/secondaryNetworks/captivePortals/{networkId}/campaign/branding][%d] Customer.prototype.patchCampaignCaptivePortalBranding default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePatchCampaignCaptivePortalBrandingDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypePatchCampaignCaptivePortalBrandingDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}