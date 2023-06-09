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

// CustomerRegisterWithGroupsReader is a Reader for the CustomerRegisterWithGroups structure.
type CustomerRegisterWithGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerRegisterWithGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerRegisterWithGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerRegisterWithGroupsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerRegisterWithGroupsOK creates a CustomerRegisterWithGroupsOK with default headers values
func NewCustomerRegisterWithGroupsOK() *CustomerRegisterWithGroupsOK {
	return &CustomerRegisterWithGroupsOK{}
}

/*
CustomerRegisterWithGroupsOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerRegisterWithGroupsOK struct {
	Payload *models.RegisterResponse
}

// IsSuccess returns true when this customer register with groups o k response has a 2xx status code
func (o *CustomerRegisterWithGroupsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer register with groups o k response has a 3xx status code
func (o *CustomerRegisterWithGroupsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer register with groups o k response has a 4xx status code
func (o *CustomerRegisterWithGroupsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer register with groups o k response has a 5xx status code
func (o *CustomerRegisterWithGroupsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer register with groups o k response a status code equal to that given
func (o *CustomerRegisterWithGroupsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer register with groups o k response
func (o *CustomerRegisterWithGroupsOK) Code() int {
	return 200
}

func (o *CustomerRegisterWithGroupsOK) Error() string {
	return fmt.Sprintf("[POST /Customers/registerWithGroups][%d] customerRegisterWithGroupsOK  %+v", 200, o.Payload)
}

func (o *CustomerRegisterWithGroupsOK) String() string {
	return fmt.Sprintf("[POST /Customers/registerWithGroups][%d] customerRegisterWithGroupsOK  %+v", 200, o.Payload)
}

func (o *CustomerRegisterWithGroupsOK) GetPayload() *models.RegisterResponse {
	return o.Payload
}

func (o *CustomerRegisterWithGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RegisterResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerRegisterWithGroupsDefault creates a CustomerRegisterWithGroupsDefault with default headers values
func NewCustomerRegisterWithGroupsDefault(code int) *CustomerRegisterWithGroupsDefault {
	return &CustomerRegisterWithGroupsDefault{
		_statusCode: code,
	}
}

/*
CustomerRegisterWithGroupsDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerRegisterWithGroupsDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer register with groups default response has a 2xx status code
func (o *CustomerRegisterWithGroupsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer register with groups default response has a 3xx status code
func (o *CustomerRegisterWithGroupsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer register with groups default response has a 4xx status code
func (o *CustomerRegisterWithGroupsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer register with groups default response has a 5xx status code
func (o *CustomerRegisterWithGroupsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer register with groups default response a status code equal to that given
func (o *CustomerRegisterWithGroupsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer register with groups default response
func (o *CustomerRegisterWithGroupsDefault) Code() int {
	return o._statusCode
}

func (o *CustomerRegisterWithGroupsDefault) Error() string {
	return fmt.Sprintf("[POST /Customers/registerWithGroups][%d] Customer.registerWithGroups default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerRegisterWithGroupsDefault) String() string {
	return fmt.Sprintf("[POST /Customers/registerWithGroups][%d] Customer.registerWithGroups default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerRegisterWithGroupsDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerRegisterWithGroupsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
