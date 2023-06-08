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

// CustomerPrototypeGetServiceLevelReader is a Reader for the CustomerPrototypeGetServiceLevel structure.
type CustomerPrototypeGetServiceLevelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypeGetServiceLevelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypeGetServiceLevelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypeGetServiceLevelDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypeGetServiceLevelOK creates a CustomerPrototypeGetServiceLevelOK with default headers values
func NewCustomerPrototypeGetServiceLevelOK() *CustomerPrototypeGetServiceLevelOK {
	return &CustomerPrototypeGetServiceLevelOK{}
}

/*
CustomerPrototypeGetServiceLevelOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypeGetServiceLevelOK struct {
	Payload interface{}
}

// IsSuccess returns true when this customer prototype get service level o k response has a 2xx status code
func (o *CustomerPrototypeGetServiceLevelOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype get service level o k response has a 3xx status code
func (o *CustomerPrototypeGetServiceLevelOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype get service level o k response has a 4xx status code
func (o *CustomerPrototypeGetServiceLevelOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype get service level o k response has a 5xx status code
func (o *CustomerPrototypeGetServiceLevelOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype get service level o k response a status code equal to that given
func (o *CustomerPrototypeGetServiceLevelOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype get service level o k response
func (o *CustomerPrototypeGetServiceLevelOK) Code() int {
	return 200
}

func (o *CustomerPrototypeGetServiceLevelOK) Error() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/serviceLevel][%d] customerPrototypeGetServiceLevelOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeGetServiceLevelOK) String() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/serviceLevel][%d] customerPrototypeGetServiceLevelOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeGetServiceLevelOK) GetPayload() interface{} {
	return o.Payload
}

func (o *CustomerPrototypeGetServiceLevelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypeGetServiceLevelDefault creates a CustomerPrototypeGetServiceLevelDefault with default headers values
func NewCustomerPrototypeGetServiceLevelDefault(code int) *CustomerPrototypeGetServiceLevelDefault {
	return &CustomerPrototypeGetServiceLevelDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypeGetServiceLevelDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypeGetServiceLevelDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype get service level default response has a 2xx status code
func (o *CustomerPrototypeGetServiceLevelDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype get service level default response has a 3xx status code
func (o *CustomerPrototypeGetServiceLevelDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype get service level default response has a 4xx status code
func (o *CustomerPrototypeGetServiceLevelDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype get service level default response has a 5xx status code
func (o *CustomerPrototypeGetServiceLevelDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype get service level default response a status code equal to that given
func (o *CustomerPrototypeGetServiceLevelDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype get service level default response
func (o *CustomerPrototypeGetServiceLevelDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypeGetServiceLevelDefault) Error() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/serviceLevel][%d] Customer.prototype.getServiceLevel default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeGetServiceLevelDefault) String() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/serviceLevel][%d] Customer.prototype.getServiceLevel default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeGetServiceLevelDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypeGetServiceLevelDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}