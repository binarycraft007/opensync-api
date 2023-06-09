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

// CustomerPrototypeGetLocationSecurityPolicyEventsReader is a Reader for the CustomerPrototypeGetLocationSecurityPolicyEvents structure.
type CustomerPrototypeGetLocationSecurityPolicyEventsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypeGetLocationSecurityPolicyEventsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypeGetLocationSecurityPolicyEventsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypeGetLocationSecurityPolicyEventsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypeGetLocationSecurityPolicyEventsOK creates a CustomerPrototypeGetLocationSecurityPolicyEventsOK with default headers values
func NewCustomerPrototypeGetLocationSecurityPolicyEventsOK() *CustomerPrototypeGetLocationSecurityPolicyEventsOK {
	return &CustomerPrototypeGetLocationSecurityPolicyEventsOK{}
}

/*
CustomerPrototypeGetLocationSecurityPolicyEventsOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypeGetLocationSecurityPolicyEventsOK struct {
	Payload models.SecurityEventsResponse
}

// IsSuccess returns true when this customer prototype get location security policy events o k response has a 2xx status code
func (o *CustomerPrototypeGetLocationSecurityPolicyEventsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype get location security policy events o k response has a 3xx status code
func (o *CustomerPrototypeGetLocationSecurityPolicyEventsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype get location security policy events o k response has a 4xx status code
func (o *CustomerPrototypeGetLocationSecurityPolicyEventsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype get location security policy events o k response has a 5xx status code
func (o *CustomerPrototypeGetLocationSecurityPolicyEventsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype get location security policy events o k response a status code equal to that given
func (o *CustomerPrototypeGetLocationSecurityPolicyEventsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype get location security policy events o k response
func (o *CustomerPrototypeGetLocationSecurityPolicyEventsOK) Code() int {
	return 200
}

func (o *CustomerPrototypeGetLocationSecurityPolicyEventsOK) Error() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/securityPolicy/events][%d] customerPrototypeGetLocationSecurityPolicyEventsOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeGetLocationSecurityPolicyEventsOK) String() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/securityPolicy/events][%d] customerPrototypeGetLocationSecurityPolicyEventsOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeGetLocationSecurityPolicyEventsOK) GetPayload() models.SecurityEventsResponse {
	return o.Payload
}

func (o *CustomerPrototypeGetLocationSecurityPolicyEventsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypeGetLocationSecurityPolicyEventsDefault creates a CustomerPrototypeGetLocationSecurityPolicyEventsDefault with default headers values
func NewCustomerPrototypeGetLocationSecurityPolicyEventsDefault(code int) *CustomerPrototypeGetLocationSecurityPolicyEventsDefault {
	return &CustomerPrototypeGetLocationSecurityPolicyEventsDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypeGetLocationSecurityPolicyEventsDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypeGetLocationSecurityPolicyEventsDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype get location security policy events default response has a 2xx status code
func (o *CustomerPrototypeGetLocationSecurityPolicyEventsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype get location security policy events default response has a 3xx status code
func (o *CustomerPrototypeGetLocationSecurityPolicyEventsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype get location security policy events default response has a 4xx status code
func (o *CustomerPrototypeGetLocationSecurityPolicyEventsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype get location security policy events default response has a 5xx status code
func (o *CustomerPrototypeGetLocationSecurityPolicyEventsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype get location security policy events default response a status code equal to that given
func (o *CustomerPrototypeGetLocationSecurityPolicyEventsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype get location security policy events default response
func (o *CustomerPrototypeGetLocationSecurityPolicyEventsDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypeGetLocationSecurityPolicyEventsDefault) Error() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/securityPolicy/events][%d] Customer.prototype.getLocationSecurityPolicyEvents default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeGetLocationSecurityPolicyEventsDefault) String() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/securityPolicy/events][%d] Customer.prototype.getLocationSecurityPolicyEvents default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeGetLocationSecurityPolicyEventsDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypeGetLocationSecurityPolicyEventsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
