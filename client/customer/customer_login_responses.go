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

// CustomerLoginReader is a Reader for the CustomerLogin structure.
type CustomerLoginReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerLoginReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerLoginOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerLoginDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerLoginOK creates a CustomerLoginOK with default headers values
func NewCustomerLoginOK() *CustomerLoginOK {
	return &CustomerLoginOK{}
}

/*
CustomerLoginOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerLoginOK struct {
	Payload interface{}
}

// IsSuccess returns true when this customer login o k response has a 2xx status code
func (o *CustomerLoginOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer login o k response has a 3xx status code
func (o *CustomerLoginOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer login o k response has a 4xx status code
func (o *CustomerLoginOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer login o k response has a 5xx status code
func (o *CustomerLoginOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer login o k response a status code equal to that given
func (o *CustomerLoginOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer login o k response
func (o *CustomerLoginOK) Code() int {
	return 200
}

func (o *CustomerLoginOK) Error() string {
	return fmt.Sprintf("[POST /Customers/login][%d] customerLoginOK  %+v", 200, o.Payload)
}

func (o *CustomerLoginOK) String() string {
	return fmt.Sprintf("[POST /Customers/login][%d] customerLoginOK  %+v", 200, o.Payload)
}

func (o *CustomerLoginOK) GetPayload() interface{} {
	return o.Payload
}

func (o *CustomerLoginOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerLoginDefault creates a CustomerLoginDefault with default headers values
func NewCustomerLoginDefault(code int) *CustomerLoginDefault {
	return &CustomerLoginDefault{
		_statusCode: code,
	}
}

/*
CustomerLoginDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerLoginDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer login default response has a 2xx status code
func (o *CustomerLoginDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer login default response has a 3xx status code
func (o *CustomerLoginDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer login default response has a 4xx status code
func (o *CustomerLoginDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer login default response has a 5xx status code
func (o *CustomerLoginDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer login default response a status code equal to that given
func (o *CustomerLoginDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer login default response
func (o *CustomerLoginDefault) Code() int {
	return o._statusCode
}

func (o *CustomerLoginDefault) Error() string {
	return fmt.Sprintf("[POST /Customers/login][%d] Customer.login default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerLoginDefault) String() string {
	return fmt.Sprintf("[POST /Customers/login][%d] Customer.login default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerLoginDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerLoginDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
