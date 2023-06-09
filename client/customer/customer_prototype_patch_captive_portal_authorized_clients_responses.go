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

// CustomerPrototypePatchCaptivePortalAuthorizedClientsReader is a Reader for the CustomerPrototypePatchCaptivePortalAuthorizedClients structure.
type CustomerPrototypePatchCaptivePortalAuthorizedClientsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypePatchCaptivePortalAuthorizedClientsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypePatchCaptivePortalAuthorizedClientsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypePatchCaptivePortalAuthorizedClientsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypePatchCaptivePortalAuthorizedClientsOK creates a CustomerPrototypePatchCaptivePortalAuthorizedClientsOK with default headers values
func NewCustomerPrototypePatchCaptivePortalAuthorizedClientsOK() *CustomerPrototypePatchCaptivePortalAuthorizedClientsOK {
	return &CustomerPrototypePatchCaptivePortalAuthorizedClientsOK{}
}

/*
CustomerPrototypePatchCaptivePortalAuthorizedClientsOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypePatchCaptivePortalAuthorizedClientsOK struct {
	Payload interface{}
}

// IsSuccess returns true when this customer prototype patch captive portal authorized clients o k response has a 2xx status code
func (o *CustomerPrototypePatchCaptivePortalAuthorizedClientsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype patch captive portal authorized clients o k response has a 3xx status code
func (o *CustomerPrototypePatchCaptivePortalAuthorizedClientsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype patch captive portal authorized clients o k response has a 4xx status code
func (o *CustomerPrototypePatchCaptivePortalAuthorizedClientsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype patch captive portal authorized clients o k response has a 5xx status code
func (o *CustomerPrototypePatchCaptivePortalAuthorizedClientsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype patch captive portal authorized clients o k response a status code equal to that given
func (o *CustomerPrototypePatchCaptivePortalAuthorizedClientsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype patch captive portal authorized clients o k response
func (o *CustomerPrototypePatchCaptivePortalAuthorizedClientsOK) Code() int {
	return 200
}

func (o *CustomerPrototypePatchCaptivePortalAuthorizedClientsOK) Error() string {
	return fmt.Sprintf("[PATCH /Customers/{id}/locations/{locationId}/secondaryNetworks/captivePortals/{networkId}/authorizedClients/{mac}][%d] customerPrototypePatchCaptivePortalAuthorizedClientsOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypePatchCaptivePortalAuthorizedClientsOK) String() string {
	return fmt.Sprintf("[PATCH /Customers/{id}/locations/{locationId}/secondaryNetworks/captivePortals/{networkId}/authorizedClients/{mac}][%d] customerPrototypePatchCaptivePortalAuthorizedClientsOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypePatchCaptivePortalAuthorizedClientsOK) GetPayload() interface{} {
	return o.Payload
}

func (o *CustomerPrototypePatchCaptivePortalAuthorizedClientsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypePatchCaptivePortalAuthorizedClientsDefault creates a CustomerPrototypePatchCaptivePortalAuthorizedClientsDefault with default headers values
func NewCustomerPrototypePatchCaptivePortalAuthorizedClientsDefault(code int) *CustomerPrototypePatchCaptivePortalAuthorizedClientsDefault {
	return &CustomerPrototypePatchCaptivePortalAuthorizedClientsDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypePatchCaptivePortalAuthorizedClientsDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypePatchCaptivePortalAuthorizedClientsDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype patch captive portal authorized clients default response has a 2xx status code
func (o *CustomerPrototypePatchCaptivePortalAuthorizedClientsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype patch captive portal authorized clients default response has a 3xx status code
func (o *CustomerPrototypePatchCaptivePortalAuthorizedClientsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype patch captive portal authorized clients default response has a 4xx status code
func (o *CustomerPrototypePatchCaptivePortalAuthorizedClientsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype patch captive portal authorized clients default response has a 5xx status code
func (o *CustomerPrototypePatchCaptivePortalAuthorizedClientsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype patch captive portal authorized clients default response a status code equal to that given
func (o *CustomerPrototypePatchCaptivePortalAuthorizedClientsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype patch captive portal authorized clients default response
func (o *CustomerPrototypePatchCaptivePortalAuthorizedClientsDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypePatchCaptivePortalAuthorizedClientsDefault) Error() string {
	return fmt.Sprintf("[PATCH /Customers/{id}/locations/{locationId}/secondaryNetworks/captivePortals/{networkId}/authorizedClients/{mac}][%d] Customer.prototype.patchCaptivePortalAuthorizedClients default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePatchCaptivePortalAuthorizedClientsDefault) String() string {
	return fmt.Sprintf("[PATCH /Customers/{id}/locations/{locationId}/secondaryNetworks/captivePortals/{networkId}/authorizedClients/{mac}][%d] Customer.prototype.patchCaptivePortalAuthorizedClients default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePatchCaptivePortalAuthorizedClientsDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypePatchCaptivePortalAuthorizedClientsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
