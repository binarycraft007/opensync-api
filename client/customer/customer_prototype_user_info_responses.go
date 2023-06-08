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

// CustomerPrototypeUserInfoReader is a Reader for the CustomerPrototypeUserInfo structure.
type CustomerPrototypeUserInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypeUserInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypeUserInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypeUserInfoDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypeUserInfoOK creates a CustomerPrototypeUserInfoOK with default headers values
func NewCustomerPrototypeUserInfoOK() *CustomerPrototypeUserInfoOK {
	return &CustomerPrototypeUserInfoOK{}
}

/*
CustomerPrototypeUserInfoOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypeUserInfoOK struct {
	Payload *models.Customer
}

// IsSuccess returns true when this customer prototype user info o k response has a 2xx status code
func (o *CustomerPrototypeUserInfoOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype user info o k response has a 3xx status code
func (o *CustomerPrototypeUserInfoOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype user info o k response has a 4xx status code
func (o *CustomerPrototypeUserInfoOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype user info o k response has a 5xx status code
func (o *CustomerPrototypeUserInfoOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype user info o k response a status code equal to that given
func (o *CustomerPrototypeUserInfoOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype user info o k response
func (o *CustomerPrototypeUserInfoOK) Code() int {
	return 200
}

func (o *CustomerPrototypeUserInfoOK) Error() string {
	return fmt.Sprintf("[GET /Customers/{id}/userInfo][%d] customerPrototypeUserInfoOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeUserInfoOK) String() string {
	return fmt.Sprintf("[GET /Customers/{id}/userInfo][%d] customerPrototypeUserInfoOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeUserInfoOK) GetPayload() *models.Customer {
	return o.Payload
}

func (o *CustomerPrototypeUserInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Customer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypeUserInfoDefault creates a CustomerPrototypeUserInfoDefault with default headers values
func NewCustomerPrototypeUserInfoDefault(code int) *CustomerPrototypeUserInfoDefault {
	return &CustomerPrototypeUserInfoDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypeUserInfoDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypeUserInfoDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype user info default response has a 2xx status code
func (o *CustomerPrototypeUserInfoDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype user info default response has a 3xx status code
func (o *CustomerPrototypeUserInfoDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype user info default response has a 4xx status code
func (o *CustomerPrototypeUserInfoDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype user info default response has a 5xx status code
func (o *CustomerPrototypeUserInfoDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype user info default response a status code equal to that given
func (o *CustomerPrototypeUserInfoDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype user info default response
func (o *CustomerPrototypeUserInfoDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypeUserInfoDefault) Error() string {
	return fmt.Sprintf("[GET /Customers/{id}/userInfo][%d] Customer.prototype.userInfo default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeUserInfoDefault) String() string {
	return fmt.Sprintf("[GET /Customers/{id}/userInfo][%d] Customer.prototype.userInfo default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeUserInfoDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypeUserInfoDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}