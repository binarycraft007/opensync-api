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

// CustomerPrototypeMigrationStatusReader is a Reader for the CustomerPrototypeMigrationStatus structure.
type CustomerPrototypeMigrationStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypeMigrationStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypeMigrationStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypeMigrationStatusDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypeMigrationStatusOK creates a CustomerPrototypeMigrationStatusOK with default headers values
func NewCustomerPrototypeMigrationStatusOK() *CustomerPrototypeMigrationStatusOK {
	return &CustomerPrototypeMigrationStatusOK{}
}

/*
CustomerPrototypeMigrationStatusOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypeMigrationStatusOK struct {
	Payload interface{}
}

// IsSuccess returns true when this customer prototype migration status o k response has a 2xx status code
func (o *CustomerPrototypeMigrationStatusOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype migration status o k response has a 3xx status code
func (o *CustomerPrototypeMigrationStatusOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype migration status o k response has a 4xx status code
func (o *CustomerPrototypeMigrationStatusOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype migration status o k response has a 5xx status code
func (o *CustomerPrototypeMigrationStatusOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype migration status o k response a status code equal to that given
func (o *CustomerPrototypeMigrationStatusOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype migration status o k response
func (o *CustomerPrototypeMigrationStatusOK) Code() int {
	return 200
}

func (o *CustomerPrototypeMigrationStatusOK) Error() string {
	return fmt.Sprintf("[GET /Customers/{id}/migration][%d] customerPrototypeMigrationStatusOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeMigrationStatusOK) String() string {
	return fmt.Sprintf("[GET /Customers/{id}/migration][%d] customerPrototypeMigrationStatusOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeMigrationStatusOK) GetPayload() interface{} {
	return o.Payload
}

func (o *CustomerPrototypeMigrationStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypeMigrationStatusDefault creates a CustomerPrototypeMigrationStatusDefault with default headers values
func NewCustomerPrototypeMigrationStatusDefault(code int) *CustomerPrototypeMigrationStatusDefault {
	return &CustomerPrototypeMigrationStatusDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypeMigrationStatusDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypeMigrationStatusDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype migration status default response has a 2xx status code
func (o *CustomerPrototypeMigrationStatusDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype migration status default response has a 3xx status code
func (o *CustomerPrototypeMigrationStatusDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype migration status default response has a 4xx status code
func (o *CustomerPrototypeMigrationStatusDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype migration status default response has a 5xx status code
func (o *CustomerPrototypeMigrationStatusDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype migration status default response a status code equal to that given
func (o *CustomerPrototypeMigrationStatusDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype migration status default response
func (o *CustomerPrototypeMigrationStatusDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypeMigrationStatusDefault) Error() string {
	return fmt.Sprintf("[GET /Customers/{id}/migration][%d] Customer.prototype.migrationStatus default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeMigrationStatusDefault) String() string {
	return fmt.Sprintf("[GET /Customers/{id}/migration][%d] Customer.prototype.migrationStatus default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeMigrationStatusDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypeMigrationStatusDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}