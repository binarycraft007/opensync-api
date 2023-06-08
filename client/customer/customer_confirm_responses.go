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

// CustomerConfirmReader is a Reader for the CustomerConfirm structure.
type CustomerConfirmReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerConfirmReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewCustomerConfirmNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerConfirmDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerConfirmNoContent creates a CustomerConfirmNoContent with default headers values
func NewCustomerConfirmNoContent() *CustomerConfirmNoContent {
	return &CustomerConfirmNoContent{}
}

/*
CustomerConfirmNoContent describes a response with status code 204, with default header values.

Request was successful
*/
type CustomerConfirmNoContent struct {
}

// IsSuccess returns true when this customer confirm no content response has a 2xx status code
func (o *CustomerConfirmNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer confirm no content response has a 3xx status code
func (o *CustomerConfirmNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer confirm no content response has a 4xx status code
func (o *CustomerConfirmNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer confirm no content response has a 5xx status code
func (o *CustomerConfirmNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this customer confirm no content response a status code equal to that given
func (o *CustomerConfirmNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the customer confirm no content response
func (o *CustomerConfirmNoContent) Code() int {
	return 204
}

func (o *CustomerConfirmNoContent) Error() string {
	return fmt.Sprintf("[GET /Customers/confirm][%d] customerConfirmNoContent ", 204)
}

func (o *CustomerConfirmNoContent) String() string {
	return fmt.Sprintf("[GET /Customers/confirm][%d] customerConfirmNoContent ", 204)
}

func (o *CustomerConfirmNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCustomerConfirmDefault creates a CustomerConfirmDefault with default headers values
func NewCustomerConfirmDefault(code int) *CustomerConfirmDefault {
	return &CustomerConfirmDefault{
		_statusCode: code,
	}
}

/*
CustomerConfirmDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerConfirmDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer confirm default response has a 2xx status code
func (o *CustomerConfirmDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer confirm default response has a 3xx status code
func (o *CustomerConfirmDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer confirm default response has a 4xx status code
func (o *CustomerConfirmDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer confirm default response has a 5xx status code
func (o *CustomerConfirmDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer confirm default response a status code equal to that given
func (o *CustomerConfirmDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer confirm default response
func (o *CustomerConfirmDefault) Code() int {
	return o._statusCode
}

func (o *CustomerConfirmDefault) Error() string {
	return fmt.Sprintf("[GET /Customers/confirm][%d] Customer.confirm default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerConfirmDefault) String() string {
	return fmt.Sprintf("[GET /Customers/confirm][%d] Customer.confirm default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerConfirmDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerConfirmDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
