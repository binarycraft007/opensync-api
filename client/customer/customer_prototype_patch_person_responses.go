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

// CustomerPrototypePatchPersonReader is a Reader for the CustomerPrototypePatchPerson structure.
type CustomerPrototypePatchPersonReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypePatchPersonReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypePatchPersonOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypePatchPersonDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypePatchPersonOK creates a CustomerPrototypePatchPersonOK with default headers values
func NewCustomerPrototypePatchPersonOK() *CustomerPrototypePatchPersonOK {
	return &CustomerPrototypePatchPersonOK{}
}

/*
CustomerPrototypePatchPersonOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypePatchPersonOK struct {
	Payload *models.Person
}

// IsSuccess returns true when this customer prototype patch person o k response has a 2xx status code
func (o *CustomerPrototypePatchPersonOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype patch person o k response has a 3xx status code
func (o *CustomerPrototypePatchPersonOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype patch person o k response has a 4xx status code
func (o *CustomerPrototypePatchPersonOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype patch person o k response has a 5xx status code
func (o *CustomerPrototypePatchPersonOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype patch person o k response a status code equal to that given
func (o *CustomerPrototypePatchPersonOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype patch person o k response
func (o *CustomerPrototypePatchPersonOK) Code() int {
	return 200
}

func (o *CustomerPrototypePatchPersonOK) Error() string {
	return fmt.Sprintf("[PATCH /Customers/{id}/locations/{locationId}/persons/{personId}][%d] customerPrototypePatchPersonOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypePatchPersonOK) String() string {
	return fmt.Sprintf("[PATCH /Customers/{id}/locations/{locationId}/persons/{personId}][%d] customerPrototypePatchPersonOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypePatchPersonOK) GetPayload() *models.Person {
	return o.Payload
}

func (o *CustomerPrototypePatchPersonOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Person)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypePatchPersonDefault creates a CustomerPrototypePatchPersonDefault with default headers values
func NewCustomerPrototypePatchPersonDefault(code int) *CustomerPrototypePatchPersonDefault {
	return &CustomerPrototypePatchPersonDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypePatchPersonDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypePatchPersonDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype patch person default response has a 2xx status code
func (o *CustomerPrototypePatchPersonDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype patch person default response has a 3xx status code
func (o *CustomerPrototypePatchPersonDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype patch person default response has a 4xx status code
func (o *CustomerPrototypePatchPersonDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype patch person default response has a 5xx status code
func (o *CustomerPrototypePatchPersonDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype patch person default response a status code equal to that given
func (o *CustomerPrototypePatchPersonDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype patch person default response
func (o *CustomerPrototypePatchPersonDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypePatchPersonDefault) Error() string {
	return fmt.Sprintf("[PATCH /Customers/{id}/locations/{locationId}/persons/{personId}][%d] Customer.prototype.patchPerson default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePatchPersonDefault) String() string {
	return fmt.Sprintf("[PATCH /Customers/{id}/locations/{locationId}/persons/{personId}][%d] Customer.prototype.patchPerson default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePatchPersonDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypePatchPersonDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
