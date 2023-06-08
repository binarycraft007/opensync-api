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

// CustomerPrototypeGetLocationSecurityPolicyReader is a Reader for the CustomerPrototypeGetLocationSecurityPolicy structure.
type CustomerPrototypeGetLocationSecurityPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypeGetLocationSecurityPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypeGetLocationSecurityPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypeGetLocationSecurityPolicyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypeGetLocationSecurityPolicyOK creates a CustomerPrototypeGetLocationSecurityPolicyOK with default headers values
func NewCustomerPrototypeGetLocationSecurityPolicyOK() *CustomerPrototypeGetLocationSecurityPolicyOK {
	return &CustomerPrototypeGetLocationSecurityPolicyOK{}
}

/*
CustomerPrototypeGetLocationSecurityPolicyOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypeGetLocationSecurityPolicyOK struct {
	Payload *models.Person
}

// IsSuccess returns true when this customer prototype get location security policy o k response has a 2xx status code
func (o *CustomerPrototypeGetLocationSecurityPolicyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype get location security policy o k response has a 3xx status code
func (o *CustomerPrototypeGetLocationSecurityPolicyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype get location security policy o k response has a 4xx status code
func (o *CustomerPrototypeGetLocationSecurityPolicyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype get location security policy o k response has a 5xx status code
func (o *CustomerPrototypeGetLocationSecurityPolicyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype get location security policy o k response a status code equal to that given
func (o *CustomerPrototypeGetLocationSecurityPolicyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype get location security policy o k response
func (o *CustomerPrototypeGetLocationSecurityPolicyOK) Code() int {
	return 200
}

func (o *CustomerPrototypeGetLocationSecurityPolicyOK) Error() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/securityPolicy][%d] customerPrototypeGetLocationSecurityPolicyOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeGetLocationSecurityPolicyOK) String() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/securityPolicy][%d] customerPrototypeGetLocationSecurityPolicyOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeGetLocationSecurityPolicyOK) GetPayload() *models.Person {
	return o.Payload
}

func (o *CustomerPrototypeGetLocationSecurityPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Person)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypeGetLocationSecurityPolicyDefault creates a CustomerPrototypeGetLocationSecurityPolicyDefault with default headers values
func NewCustomerPrototypeGetLocationSecurityPolicyDefault(code int) *CustomerPrototypeGetLocationSecurityPolicyDefault {
	return &CustomerPrototypeGetLocationSecurityPolicyDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypeGetLocationSecurityPolicyDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypeGetLocationSecurityPolicyDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype get location security policy default response has a 2xx status code
func (o *CustomerPrototypeGetLocationSecurityPolicyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype get location security policy default response has a 3xx status code
func (o *CustomerPrototypeGetLocationSecurityPolicyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype get location security policy default response has a 4xx status code
func (o *CustomerPrototypeGetLocationSecurityPolicyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype get location security policy default response has a 5xx status code
func (o *CustomerPrototypeGetLocationSecurityPolicyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype get location security policy default response a status code equal to that given
func (o *CustomerPrototypeGetLocationSecurityPolicyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype get location security policy default response
func (o *CustomerPrototypeGetLocationSecurityPolicyDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypeGetLocationSecurityPolicyDefault) Error() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/securityPolicy][%d] Customer.prototype.getLocationSecurityPolicy default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeGetLocationSecurityPolicyDefault) String() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/securityPolicy][%d] Customer.prototype.getLocationSecurityPolicy default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeGetLocationSecurityPolicyDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypeGetLocationSecurityPolicyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}