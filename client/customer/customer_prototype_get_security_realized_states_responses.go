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

// CustomerPrototypeGetSecurityRealizedStatesReader is a Reader for the CustomerPrototypeGetSecurityRealizedStates structure.
type CustomerPrototypeGetSecurityRealizedStatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypeGetSecurityRealizedStatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypeGetSecurityRealizedStatesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypeGetSecurityRealizedStatesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypeGetSecurityRealizedStatesOK creates a CustomerPrototypeGetSecurityRealizedStatesOK with default headers values
func NewCustomerPrototypeGetSecurityRealizedStatesOK() *CustomerPrototypeGetSecurityRealizedStatesOK {
	return &CustomerPrototypeGetSecurityRealizedStatesOK{}
}

/*
CustomerPrototypeGetSecurityRealizedStatesOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypeGetSecurityRealizedStatesOK struct {
	Payload interface{}
}

// IsSuccess returns true when this customer prototype get security realized states o k response has a 2xx status code
func (o *CustomerPrototypeGetSecurityRealizedStatesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype get security realized states o k response has a 3xx status code
func (o *CustomerPrototypeGetSecurityRealizedStatesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype get security realized states o k response has a 4xx status code
func (o *CustomerPrototypeGetSecurityRealizedStatesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype get security realized states o k response has a 5xx status code
func (o *CustomerPrototypeGetSecurityRealizedStatesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype get security realized states o k response a status code equal to that given
func (o *CustomerPrototypeGetSecurityRealizedStatesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype get security realized states o k response
func (o *CustomerPrototypeGetSecurityRealizedStatesOK) Code() int {
	return 200
}

func (o *CustomerPrototypeGetSecurityRealizedStatesOK) Error() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/securityPolicy/realizedState][%d] customerPrototypeGetSecurityRealizedStatesOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeGetSecurityRealizedStatesOK) String() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/securityPolicy/realizedState][%d] customerPrototypeGetSecurityRealizedStatesOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeGetSecurityRealizedStatesOK) GetPayload() interface{} {
	return o.Payload
}

func (o *CustomerPrototypeGetSecurityRealizedStatesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypeGetSecurityRealizedStatesDefault creates a CustomerPrototypeGetSecurityRealizedStatesDefault with default headers values
func NewCustomerPrototypeGetSecurityRealizedStatesDefault(code int) *CustomerPrototypeGetSecurityRealizedStatesDefault {
	return &CustomerPrototypeGetSecurityRealizedStatesDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypeGetSecurityRealizedStatesDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypeGetSecurityRealizedStatesDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype get security realized states default response has a 2xx status code
func (o *CustomerPrototypeGetSecurityRealizedStatesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype get security realized states default response has a 3xx status code
func (o *CustomerPrototypeGetSecurityRealizedStatesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype get security realized states default response has a 4xx status code
func (o *CustomerPrototypeGetSecurityRealizedStatesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype get security realized states default response has a 5xx status code
func (o *CustomerPrototypeGetSecurityRealizedStatesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype get security realized states default response a status code equal to that given
func (o *CustomerPrototypeGetSecurityRealizedStatesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype get security realized states default response
func (o *CustomerPrototypeGetSecurityRealizedStatesDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypeGetSecurityRealizedStatesDefault) Error() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/securityPolicy/realizedState][%d] Customer.prototype.getSecurityRealizedStates default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeGetSecurityRealizedStatesDefault) String() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/securityPolicy/realizedState][%d] Customer.prototype.getSecurityRealizedStates default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeGetSecurityRealizedStatesDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypeGetSecurityRealizedStatesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
