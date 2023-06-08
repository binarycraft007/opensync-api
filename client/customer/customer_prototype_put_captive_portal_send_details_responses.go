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

// CustomerPrototypePutCaptivePortalSendDetailsReader is a Reader for the CustomerPrototypePutCaptivePortalSendDetails structure.
type CustomerPrototypePutCaptivePortalSendDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypePutCaptivePortalSendDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewCustomerPrototypePutCaptivePortalSendDetailsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypePutCaptivePortalSendDetailsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypePutCaptivePortalSendDetailsNoContent creates a CustomerPrototypePutCaptivePortalSendDetailsNoContent with default headers values
func NewCustomerPrototypePutCaptivePortalSendDetailsNoContent() *CustomerPrototypePutCaptivePortalSendDetailsNoContent {
	return &CustomerPrototypePutCaptivePortalSendDetailsNoContent{}
}

/*
CustomerPrototypePutCaptivePortalSendDetailsNoContent describes a response with status code 204, with default header values.

Request was successful
*/
type CustomerPrototypePutCaptivePortalSendDetailsNoContent struct {
}

// IsSuccess returns true when this customer prototype put captive portal send details no content response has a 2xx status code
func (o *CustomerPrototypePutCaptivePortalSendDetailsNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype put captive portal send details no content response has a 3xx status code
func (o *CustomerPrototypePutCaptivePortalSendDetailsNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype put captive portal send details no content response has a 4xx status code
func (o *CustomerPrototypePutCaptivePortalSendDetailsNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype put captive portal send details no content response has a 5xx status code
func (o *CustomerPrototypePutCaptivePortalSendDetailsNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype put captive portal send details no content response a status code equal to that given
func (o *CustomerPrototypePutCaptivePortalSendDetailsNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the customer prototype put captive portal send details no content response
func (o *CustomerPrototypePutCaptivePortalSendDetailsNoContent) Code() int {
	return 204
}

func (o *CustomerPrototypePutCaptivePortalSendDetailsNoContent) Error() string {
	return fmt.Sprintf("[PUT /Customers/{id}/locations/{locationId}/secondaryNetworks/captivePortals/{networkId}/sendGuestDetails][%d] customerPrototypePutCaptivePortalSendDetailsNoContent ", 204)
}

func (o *CustomerPrototypePutCaptivePortalSendDetailsNoContent) String() string {
	return fmt.Sprintf("[PUT /Customers/{id}/locations/{locationId}/secondaryNetworks/captivePortals/{networkId}/sendGuestDetails][%d] customerPrototypePutCaptivePortalSendDetailsNoContent ", 204)
}

func (o *CustomerPrototypePutCaptivePortalSendDetailsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCustomerPrototypePutCaptivePortalSendDetailsDefault creates a CustomerPrototypePutCaptivePortalSendDetailsDefault with default headers values
func NewCustomerPrototypePutCaptivePortalSendDetailsDefault(code int) *CustomerPrototypePutCaptivePortalSendDetailsDefault {
	return &CustomerPrototypePutCaptivePortalSendDetailsDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypePutCaptivePortalSendDetailsDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypePutCaptivePortalSendDetailsDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype put captive portal send details default response has a 2xx status code
func (o *CustomerPrototypePutCaptivePortalSendDetailsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype put captive portal send details default response has a 3xx status code
func (o *CustomerPrototypePutCaptivePortalSendDetailsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype put captive portal send details default response has a 4xx status code
func (o *CustomerPrototypePutCaptivePortalSendDetailsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype put captive portal send details default response has a 5xx status code
func (o *CustomerPrototypePutCaptivePortalSendDetailsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype put captive portal send details default response a status code equal to that given
func (o *CustomerPrototypePutCaptivePortalSendDetailsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype put captive portal send details default response
func (o *CustomerPrototypePutCaptivePortalSendDetailsDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypePutCaptivePortalSendDetailsDefault) Error() string {
	return fmt.Sprintf("[PUT /Customers/{id}/locations/{locationId}/secondaryNetworks/captivePortals/{networkId}/sendGuestDetails][%d] Customer.prototype.putCaptivePortalSendDetails default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePutCaptivePortalSendDetailsDefault) String() string {
	return fmt.Sprintf("[PUT /Customers/{id}/locations/{locationId}/secondaryNetworks/captivePortals/{networkId}/sendGuestDetails][%d] Customer.prototype.putCaptivePortalSendDetails default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePutCaptivePortalSendDetailsDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypePutCaptivePortalSendDetailsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}