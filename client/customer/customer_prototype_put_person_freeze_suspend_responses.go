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

// CustomerPrototypePutPersonFreezeSuspendReader is a Reader for the CustomerPrototypePutPersonFreezeSuspend structure.
type CustomerPrototypePutPersonFreezeSuspendReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypePutPersonFreezeSuspendReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypePutPersonFreezeSuspendOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypePutPersonFreezeSuspendDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypePutPersonFreezeSuspendOK creates a CustomerPrototypePutPersonFreezeSuspendOK with default headers values
func NewCustomerPrototypePutPersonFreezeSuspendOK() *CustomerPrototypePutPersonFreezeSuspendOK {
	return &CustomerPrototypePutPersonFreezeSuspendOK{}
}

/*
CustomerPrototypePutPersonFreezeSuspendOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypePutPersonFreezeSuspendOK struct {
	Payload models.XAny
}

// IsSuccess returns true when this customer prototype put person freeze suspend o k response has a 2xx status code
func (o *CustomerPrototypePutPersonFreezeSuspendOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype put person freeze suspend o k response has a 3xx status code
func (o *CustomerPrototypePutPersonFreezeSuspendOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype put person freeze suspend o k response has a 4xx status code
func (o *CustomerPrototypePutPersonFreezeSuspendOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype put person freeze suspend o k response has a 5xx status code
func (o *CustomerPrototypePutPersonFreezeSuspendOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype put person freeze suspend o k response a status code equal to that given
func (o *CustomerPrototypePutPersonFreezeSuspendOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype put person freeze suspend o k response
func (o *CustomerPrototypePutPersonFreezeSuspendOK) Code() int {
	return 200
}

func (o *CustomerPrototypePutPersonFreezeSuspendOK) Error() string {
	return fmt.Sprintf("[PUT /Customers/{id}/locations/{locationId}/persons/{personId}/freeze/suspend][%d] customerPrototypePutPersonFreezeSuspendOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypePutPersonFreezeSuspendOK) String() string {
	return fmt.Sprintf("[PUT /Customers/{id}/locations/{locationId}/persons/{personId}/freeze/suspend][%d] customerPrototypePutPersonFreezeSuspendOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypePutPersonFreezeSuspendOK) GetPayload() models.XAny {
	return o.Payload
}

func (o *CustomerPrototypePutPersonFreezeSuspendOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypePutPersonFreezeSuspendDefault creates a CustomerPrototypePutPersonFreezeSuspendDefault with default headers values
func NewCustomerPrototypePutPersonFreezeSuspendDefault(code int) *CustomerPrototypePutPersonFreezeSuspendDefault {
	return &CustomerPrototypePutPersonFreezeSuspendDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypePutPersonFreezeSuspendDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypePutPersonFreezeSuspendDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype put person freeze suspend default response has a 2xx status code
func (o *CustomerPrototypePutPersonFreezeSuspendDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype put person freeze suspend default response has a 3xx status code
func (o *CustomerPrototypePutPersonFreezeSuspendDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype put person freeze suspend default response has a 4xx status code
func (o *CustomerPrototypePutPersonFreezeSuspendDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype put person freeze suspend default response has a 5xx status code
func (o *CustomerPrototypePutPersonFreezeSuspendDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype put person freeze suspend default response a status code equal to that given
func (o *CustomerPrototypePutPersonFreezeSuspendDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype put person freeze suspend default response
func (o *CustomerPrototypePutPersonFreezeSuspendDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypePutPersonFreezeSuspendDefault) Error() string {
	return fmt.Sprintf("[PUT /Customers/{id}/locations/{locationId}/persons/{personId}/freeze/suspend][%d] Customer.prototype.putPersonFreezeSuspend default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePutPersonFreezeSuspendDefault) String() string {
	return fmt.Sprintf("[PUT /Customers/{id}/locations/{locationId}/persons/{personId}/freeze/suspend][%d] Customer.prototype.putPersonFreezeSuspend default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePutPersonFreezeSuspendDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypePutPersonFreezeSuspendDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
