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

// CustomerPrototypeCreateNewPasswordlessTokenReader is a Reader for the CustomerPrototypeCreateNewPasswordlessToken structure.
type CustomerPrototypeCreateNewPasswordlessTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypeCreateNewPasswordlessTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypeCreateNewPasswordlessTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypeCreateNewPasswordlessTokenDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypeCreateNewPasswordlessTokenOK creates a CustomerPrototypeCreateNewPasswordlessTokenOK with default headers values
func NewCustomerPrototypeCreateNewPasswordlessTokenOK() *CustomerPrototypeCreateNewPasswordlessTokenOK {
	return &CustomerPrototypeCreateNewPasswordlessTokenOK{}
}

/*
CustomerPrototypeCreateNewPasswordlessTokenOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypeCreateNewPasswordlessTokenOK struct {
	Payload *models.AccessToken
}

// IsSuccess returns true when this customer prototype create new passwordless token o k response has a 2xx status code
func (o *CustomerPrototypeCreateNewPasswordlessTokenOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype create new passwordless token o k response has a 3xx status code
func (o *CustomerPrototypeCreateNewPasswordlessTokenOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype create new passwordless token o k response has a 4xx status code
func (o *CustomerPrototypeCreateNewPasswordlessTokenOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype create new passwordless token o k response has a 5xx status code
func (o *CustomerPrototypeCreateNewPasswordlessTokenOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype create new passwordless token o k response a status code equal to that given
func (o *CustomerPrototypeCreateNewPasswordlessTokenOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype create new passwordless token o k response
func (o *CustomerPrototypeCreateNewPasswordlessTokenOK) Code() int {
	return 200
}

func (o *CustomerPrototypeCreateNewPasswordlessTokenOK) Error() string {
	return fmt.Sprintf("[POST /Customers/{id}/accessToken][%d] customerPrototypeCreateNewPasswordlessTokenOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeCreateNewPasswordlessTokenOK) String() string {
	return fmt.Sprintf("[POST /Customers/{id}/accessToken][%d] customerPrototypeCreateNewPasswordlessTokenOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeCreateNewPasswordlessTokenOK) GetPayload() *models.AccessToken {
	return o.Payload
}

func (o *CustomerPrototypeCreateNewPasswordlessTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccessToken)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypeCreateNewPasswordlessTokenDefault creates a CustomerPrototypeCreateNewPasswordlessTokenDefault with default headers values
func NewCustomerPrototypeCreateNewPasswordlessTokenDefault(code int) *CustomerPrototypeCreateNewPasswordlessTokenDefault {
	return &CustomerPrototypeCreateNewPasswordlessTokenDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypeCreateNewPasswordlessTokenDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypeCreateNewPasswordlessTokenDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype create new passwordless token default response has a 2xx status code
func (o *CustomerPrototypeCreateNewPasswordlessTokenDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype create new passwordless token default response has a 3xx status code
func (o *CustomerPrototypeCreateNewPasswordlessTokenDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype create new passwordless token default response has a 4xx status code
func (o *CustomerPrototypeCreateNewPasswordlessTokenDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype create new passwordless token default response has a 5xx status code
func (o *CustomerPrototypeCreateNewPasswordlessTokenDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype create new passwordless token default response a status code equal to that given
func (o *CustomerPrototypeCreateNewPasswordlessTokenDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype create new passwordless token default response
func (o *CustomerPrototypeCreateNewPasswordlessTokenDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypeCreateNewPasswordlessTokenDefault) Error() string {
	return fmt.Sprintf("[POST /Customers/{id}/accessToken][%d] Customer.prototype.createNewPasswordlessToken default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeCreateNewPasswordlessTokenDefault) String() string {
	return fmt.Sprintf("[POST /Customers/{id}/accessToken][%d] Customer.prototype.createNewPasswordlessToken default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeCreateNewPasswordlessTokenDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypeCreateNewPasswordlessTokenDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
