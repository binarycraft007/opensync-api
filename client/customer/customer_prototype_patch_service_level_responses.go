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

// CustomerPrototypePatchServiceLevelReader is a Reader for the CustomerPrototypePatchServiceLevel structure.
type CustomerPrototypePatchServiceLevelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypePatchServiceLevelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypePatchServiceLevelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypePatchServiceLevelDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypePatchServiceLevelOK creates a CustomerPrototypePatchServiceLevelOK with default headers values
func NewCustomerPrototypePatchServiceLevelOK() *CustomerPrototypePatchServiceLevelOK {
	return &CustomerPrototypePatchServiceLevelOK{}
}

/*
CustomerPrototypePatchServiceLevelOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypePatchServiceLevelOK struct {
	Payload interface{}
}

// IsSuccess returns true when this customer prototype patch service level o k response has a 2xx status code
func (o *CustomerPrototypePatchServiceLevelOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype patch service level o k response has a 3xx status code
func (o *CustomerPrototypePatchServiceLevelOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype patch service level o k response has a 4xx status code
func (o *CustomerPrototypePatchServiceLevelOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype patch service level o k response has a 5xx status code
func (o *CustomerPrototypePatchServiceLevelOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype patch service level o k response a status code equal to that given
func (o *CustomerPrototypePatchServiceLevelOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype patch service level o k response
func (o *CustomerPrototypePatchServiceLevelOK) Code() int {
	return 200
}

func (o *CustomerPrototypePatchServiceLevelOK) Error() string {
	return fmt.Sprintf("[PATCH /Customers/{id}/locations/{locationId}/serviceLevel][%d] customerPrototypePatchServiceLevelOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypePatchServiceLevelOK) String() string {
	return fmt.Sprintf("[PATCH /Customers/{id}/locations/{locationId}/serviceLevel][%d] customerPrototypePatchServiceLevelOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypePatchServiceLevelOK) GetPayload() interface{} {
	return o.Payload
}

func (o *CustomerPrototypePatchServiceLevelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypePatchServiceLevelDefault creates a CustomerPrototypePatchServiceLevelDefault with default headers values
func NewCustomerPrototypePatchServiceLevelDefault(code int) *CustomerPrototypePatchServiceLevelDefault {
	return &CustomerPrototypePatchServiceLevelDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypePatchServiceLevelDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypePatchServiceLevelDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype patch service level default response has a 2xx status code
func (o *CustomerPrototypePatchServiceLevelDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype patch service level default response has a 3xx status code
func (o *CustomerPrototypePatchServiceLevelDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype patch service level default response has a 4xx status code
func (o *CustomerPrototypePatchServiceLevelDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype patch service level default response has a 5xx status code
func (o *CustomerPrototypePatchServiceLevelDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype patch service level default response a status code equal to that given
func (o *CustomerPrototypePatchServiceLevelDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype patch service level default response
func (o *CustomerPrototypePatchServiceLevelDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypePatchServiceLevelDefault) Error() string {
	return fmt.Sprintf("[PATCH /Customers/{id}/locations/{locationId}/serviceLevel][%d] Customer.prototype.patchServiceLevel default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePatchServiceLevelDefault) String() string {
	return fmt.Sprintf("[PATCH /Customers/{id}/locations/{locationId}/serviceLevel][%d] Customer.prototype.patchServiceLevel default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePatchServiceLevelDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypePatchServiceLevelDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
