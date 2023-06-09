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

// CustomerPrototypePutOptimizationsReader is a Reader for the CustomerPrototypePutOptimizations structure.
type CustomerPrototypePutOptimizationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypePutOptimizationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypePutOptimizationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypePutOptimizationsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypePutOptimizationsOK creates a CustomerPrototypePutOptimizationsOK with default headers values
func NewCustomerPrototypePutOptimizationsOK() *CustomerPrototypePutOptimizationsOK {
	return &CustomerPrototypePutOptimizationsOK{}
}

/*
CustomerPrototypePutOptimizationsOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypePutOptimizationsOK struct {
	Payload *models.Optimizations
}

// IsSuccess returns true when this customer prototype put optimizations o k response has a 2xx status code
func (o *CustomerPrototypePutOptimizationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype put optimizations o k response has a 3xx status code
func (o *CustomerPrototypePutOptimizationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype put optimizations o k response has a 4xx status code
func (o *CustomerPrototypePutOptimizationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype put optimizations o k response has a 5xx status code
func (o *CustomerPrototypePutOptimizationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype put optimizations o k response a status code equal to that given
func (o *CustomerPrototypePutOptimizationsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype put optimizations o k response
func (o *CustomerPrototypePutOptimizationsOK) Code() int {
	return 200
}

func (o *CustomerPrototypePutOptimizationsOK) Error() string {
	return fmt.Sprintf("[PUT /Customers/{id}/locations/{locationId}/optimizations][%d] customerPrototypePutOptimizationsOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypePutOptimizationsOK) String() string {
	return fmt.Sprintf("[PUT /Customers/{id}/locations/{locationId}/optimizations][%d] customerPrototypePutOptimizationsOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypePutOptimizationsOK) GetPayload() *models.Optimizations {
	return o.Payload
}

func (o *CustomerPrototypePutOptimizationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Optimizations)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypePutOptimizationsDefault creates a CustomerPrototypePutOptimizationsDefault with default headers values
func NewCustomerPrototypePutOptimizationsDefault(code int) *CustomerPrototypePutOptimizationsDefault {
	return &CustomerPrototypePutOptimizationsDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypePutOptimizationsDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypePutOptimizationsDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype put optimizations default response has a 2xx status code
func (o *CustomerPrototypePutOptimizationsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype put optimizations default response has a 3xx status code
func (o *CustomerPrototypePutOptimizationsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype put optimizations default response has a 4xx status code
func (o *CustomerPrototypePutOptimizationsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype put optimizations default response has a 5xx status code
func (o *CustomerPrototypePutOptimizationsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype put optimizations default response a status code equal to that given
func (o *CustomerPrototypePutOptimizationsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype put optimizations default response
func (o *CustomerPrototypePutOptimizationsDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypePutOptimizationsDefault) Error() string {
	return fmt.Sprintf("[PUT /Customers/{id}/locations/{locationId}/optimizations][%d] Customer.prototype.putOptimizations default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePutOptimizationsDefault) String() string {
	return fmt.Sprintf("[PUT /Customers/{id}/locations/{locationId}/optimizations][%d] Customer.prototype.putOptimizations default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePutOptimizationsDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypePutOptimizationsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
