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

// CustomerPrototypeGetDeviceSecurityPolicyEventsReader is a Reader for the CustomerPrototypeGetDeviceSecurityPolicyEvents structure.
type CustomerPrototypeGetDeviceSecurityPolicyEventsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypeGetDeviceSecurityPolicyEventsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypeGetDeviceSecurityPolicyEventsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypeGetDeviceSecurityPolicyEventsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypeGetDeviceSecurityPolicyEventsOK creates a CustomerPrototypeGetDeviceSecurityPolicyEventsOK with default headers values
func NewCustomerPrototypeGetDeviceSecurityPolicyEventsOK() *CustomerPrototypeGetDeviceSecurityPolicyEventsOK {
	return &CustomerPrototypeGetDeviceSecurityPolicyEventsOK{}
}

/*
CustomerPrototypeGetDeviceSecurityPolicyEventsOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypeGetDeviceSecurityPolicyEventsOK struct {
	Payload models.SecurityEventsResponse
}

// IsSuccess returns true when this customer prototype get device security policy events o k response has a 2xx status code
func (o *CustomerPrototypeGetDeviceSecurityPolicyEventsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype get device security policy events o k response has a 3xx status code
func (o *CustomerPrototypeGetDeviceSecurityPolicyEventsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype get device security policy events o k response has a 4xx status code
func (o *CustomerPrototypeGetDeviceSecurityPolicyEventsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype get device security policy events o k response has a 5xx status code
func (o *CustomerPrototypeGetDeviceSecurityPolicyEventsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype get device security policy events o k response a status code equal to that given
func (o *CustomerPrototypeGetDeviceSecurityPolicyEventsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype get device security policy events o k response
func (o *CustomerPrototypeGetDeviceSecurityPolicyEventsOK) Code() int {
	return 200
}

func (o *CustomerPrototypeGetDeviceSecurityPolicyEventsOK) Error() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/devices/{mac}/securityPolicy/events][%d] customerPrototypeGetDeviceSecurityPolicyEventsOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeGetDeviceSecurityPolicyEventsOK) String() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/devices/{mac}/securityPolicy/events][%d] customerPrototypeGetDeviceSecurityPolicyEventsOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeGetDeviceSecurityPolicyEventsOK) GetPayload() models.SecurityEventsResponse {
	return o.Payload
}

func (o *CustomerPrototypeGetDeviceSecurityPolicyEventsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypeGetDeviceSecurityPolicyEventsDefault creates a CustomerPrototypeGetDeviceSecurityPolicyEventsDefault with default headers values
func NewCustomerPrototypeGetDeviceSecurityPolicyEventsDefault(code int) *CustomerPrototypeGetDeviceSecurityPolicyEventsDefault {
	return &CustomerPrototypeGetDeviceSecurityPolicyEventsDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypeGetDeviceSecurityPolicyEventsDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypeGetDeviceSecurityPolicyEventsDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype get device security policy events default response has a 2xx status code
func (o *CustomerPrototypeGetDeviceSecurityPolicyEventsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype get device security policy events default response has a 3xx status code
func (o *CustomerPrototypeGetDeviceSecurityPolicyEventsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype get device security policy events default response has a 4xx status code
func (o *CustomerPrototypeGetDeviceSecurityPolicyEventsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype get device security policy events default response has a 5xx status code
func (o *CustomerPrototypeGetDeviceSecurityPolicyEventsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype get device security policy events default response a status code equal to that given
func (o *CustomerPrototypeGetDeviceSecurityPolicyEventsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype get device security policy events default response
func (o *CustomerPrototypeGetDeviceSecurityPolicyEventsDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypeGetDeviceSecurityPolicyEventsDefault) Error() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/devices/{mac}/securityPolicy/events][%d] Customer.prototype.getDeviceSecurityPolicyEvents default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeGetDeviceSecurityPolicyEventsDefault) String() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/devices/{mac}/securityPolicy/events][%d] Customer.prototype.getDeviceSecurityPolicyEvents default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeGetDeviceSecurityPolicyEventsDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypeGetDeviceSecurityPolicyEventsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
