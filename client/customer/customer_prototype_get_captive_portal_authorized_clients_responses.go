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

// CustomerPrototypeGetCaptivePortalAuthorizedClientsReader is a Reader for the CustomerPrototypeGetCaptivePortalAuthorizedClients structure.
type CustomerPrototypeGetCaptivePortalAuthorizedClientsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypeGetCaptivePortalAuthorizedClientsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypeGetCaptivePortalAuthorizedClientsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypeGetCaptivePortalAuthorizedClientsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypeGetCaptivePortalAuthorizedClientsOK creates a CustomerPrototypeGetCaptivePortalAuthorizedClientsOK with default headers values
func NewCustomerPrototypeGetCaptivePortalAuthorizedClientsOK() *CustomerPrototypeGetCaptivePortalAuthorizedClientsOK {
	return &CustomerPrototypeGetCaptivePortalAuthorizedClientsOK{}
}

/*
CustomerPrototypeGetCaptivePortalAuthorizedClientsOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypeGetCaptivePortalAuthorizedClientsOK struct {
	Payload interface{}
}

// IsSuccess returns true when this customer prototype get captive portal authorized clients o k response has a 2xx status code
func (o *CustomerPrototypeGetCaptivePortalAuthorizedClientsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype get captive portal authorized clients o k response has a 3xx status code
func (o *CustomerPrototypeGetCaptivePortalAuthorizedClientsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype get captive portal authorized clients o k response has a 4xx status code
func (o *CustomerPrototypeGetCaptivePortalAuthorizedClientsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype get captive portal authorized clients o k response has a 5xx status code
func (o *CustomerPrototypeGetCaptivePortalAuthorizedClientsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype get captive portal authorized clients o k response a status code equal to that given
func (o *CustomerPrototypeGetCaptivePortalAuthorizedClientsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype get captive portal authorized clients o k response
func (o *CustomerPrototypeGetCaptivePortalAuthorizedClientsOK) Code() int {
	return 200
}

func (o *CustomerPrototypeGetCaptivePortalAuthorizedClientsOK) Error() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/secondaryNetworks/captivePortals/{networkId}/authorizedClients][%d] customerPrototypeGetCaptivePortalAuthorizedClientsOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeGetCaptivePortalAuthorizedClientsOK) String() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/secondaryNetworks/captivePortals/{networkId}/authorizedClients][%d] customerPrototypeGetCaptivePortalAuthorizedClientsOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeGetCaptivePortalAuthorizedClientsOK) GetPayload() interface{} {
	return o.Payload
}

func (o *CustomerPrototypeGetCaptivePortalAuthorizedClientsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypeGetCaptivePortalAuthorizedClientsDefault creates a CustomerPrototypeGetCaptivePortalAuthorizedClientsDefault with default headers values
func NewCustomerPrototypeGetCaptivePortalAuthorizedClientsDefault(code int) *CustomerPrototypeGetCaptivePortalAuthorizedClientsDefault {
	return &CustomerPrototypeGetCaptivePortalAuthorizedClientsDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypeGetCaptivePortalAuthorizedClientsDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypeGetCaptivePortalAuthorizedClientsDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype get captive portal authorized clients default response has a 2xx status code
func (o *CustomerPrototypeGetCaptivePortalAuthorizedClientsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype get captive portal authorized clients default response has a 3xx status code
func (o *CustomerPrototypeGetCaptivePortalAuthorizedClientsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype get captive portal authorized clients default response has a 4xx status code
func (o *CustomerPrototypeGetCaptivePortalAuthorizedClientsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype get captive portal authorized clients default response has a 5xx status code
func (o *CustomerPrototypeGetCaptivePortalAuthorizedClientsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype get captive portal authorized clients default response a status code equal to that given
func (o *CustomerPrototypeGetCaptivePortalAuthorizedClientsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype get captive portal authorized clients default response
func (o *CustomerPrototypeGetCaptivePortalAuthorizedClientsDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypeGetCaptivePortalAuthorizedClientsDefault) Error() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/secondaryNetworks/captivePortals/{networkId}/authorizedClients][%d] Customer.prototype.getCaptivePortalAuthorizedClients default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeGetCaptivePortalAuthorizedClientsDefault) String() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/secondaryNetworks/captivePortals/{networkId}/authorizedClients][%d] Customer.prototype.getCaptivePortalAuthorizedClients default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeGetCaptivePortalAuthorizedClientsDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypeGetCaptivePortalAuthorizedClientsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
