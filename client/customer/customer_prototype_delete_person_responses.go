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

// CustomerPrototypeDeletePersonReader is a Reader for the CustomerPrototypeDeletePerson structure.
type CustomerPrototypeDeletePersonReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypeDeletePersonReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewCustomerPrototypeDeletePersonNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypeDeletePersonDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypeDeletePersonNoContent creates a CustomerPrototypeDeletePersonNoContent with default headers values
func NewCustomerPrototypeDeletePersonNoContent() *CustomerPrototypeDeletePersonNoContent {
	return &CustomerPrototypeDeletePersonNoContent{}
}

/*
CustomerPrototypeDeletePersonNoContent describes a response with status code 204, with default header values.

Request was successful
*/
type CustomerPrototypeDeletePersonNoContent struct {
}

// IsSuccess returns true when this customer prototype delete person no content response has a 2xx status code
func (o *CustomerPrototypeDeletePersonNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype delete person no content response has a 3xx status code
func (o *CustomerPrototypeDeletePersonNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype delete person no content response has a 4xx status code
func (o *CustomerPrototypeDeletePersonNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype delete person no content response has a 5xx status code
func (o *CustomerPrototypeDeletePersonNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype delete person no content response a status code equal to that given
func (o *CustomerPrototypeDeletePersonNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the customer prototype delete person no content response
func (o *CustomerPrototypeDeletePersonNoContent) Code() int {
	return 204
}

func (o *CustomerPrototypeDeletePersonNoContent) Error() string {
	return fmt.Sprintf("[DELETE /Customers/{id}/locations/{locationId}/persons/{personId}][%d] customerPrototypeDeletePersonNoContent ", 204)
}

func (o *CustomerPrototypeDeletePersonNoContent) String() string {
	return fmt.Sprintf("[DELETE /Customers/{id}/locations/{locationId}/persons/{personId}][%d] customerPrototypeDeletePersonNoContent ", 204)
}

func (o *CustomerPrototypeDeletePersonNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCustomerPrototypeDeletePersonDefault creates a CustomerPrototypeDeletePersonDefault with default headers values
func NewCustomerPrototypeDeletePersonDefault(code int) *CustomerPrototypeDeletePersonDefault {
	return &CustomerPrototypeDeletePersonDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypeDeletePersonDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypeDeletePersonDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype delete person default response has a 2xx status code
func (o *CustomerPrototypeDeletePersonDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype delete person default response has a 3xx status code
func (o *CustomerPrototypeDeletePersonDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype delete person default response has a 4xx status code
func (o *CustomerPrototypeDeletePersonDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype delete person default response has a 5xx status code
func (o *CustomerPrototypeDeletePersonDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype delete person default response a status code equal to that given
func (o *CustomerPrototypeDeletePersonDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype delete person default response
func (o *CustomerPrototypeDeletePersonDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypeDeletePersonDefault) Error() string {
	return fmt.Sprintf("[DELETE /Customers/{id}/locations/{locationId}/persons/{personId}][%d] Customer.prototype.deletePerson default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeDeletePersonDefault) String() string {
	return fmt.Sprintf("[DELETE /Customers/{id}/locations/{locationId}/persons/{personId}][%d] Customer.prototype.deletePerson default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeDeletePersonDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypeDeletePersonDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}