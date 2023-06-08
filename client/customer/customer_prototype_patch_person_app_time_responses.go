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

// CustomerPrototypePatchPersonAppTimeReader is a Reader for the CustomerPrototypePatchPersonAppTime structure.
type CustomerPrototypePatchPersonAppTimeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypePatchPersonAppTimeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypePatchPersonAppTimeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypePatchPersonAppTimeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypePatchPersonAppTimeOK creates a CustomerPrototypePatchPersonAppTimeOK with default headers values
func NewCustomerPrototypePatchPersonAppTimeOK() *CustomerPrototypePatchPersonAppTimeOK {
	return &CustomerPrototypePatchPersonAppTimeOK{}
}

/*
CustomerPrototypePatchPersonAppTimeOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypePatchPersonAppTimeOK struct {
	Payload interface{}
}

// IsSuccess returns true when this customer prototype patch person app time o k response has a 2xx status code
func (o *CustomerPrototypePatchPersonAppTimeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype patch person app time o k response has a 3xx status code
func (o *CustomerPrototypePatchPersonAppTimeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype patch person app time o k response has a 4xx status code
func (o *CustomerPrototypePatchPersonAppTimeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype patch person app time o k response has a 5xx status code
func (o *CustomerPrototypePatchPersonAppTimeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype patch person app time o k response a status code equal to that given
func (o *CustomerPrototypePatchPersonAppTimeOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype patch person app time o k response
func (o *CustomerPrototypePatchPersonAppTimeOK) Code() int {
	return 200
}

func (o *CustomerPrototypePatchPersonAppTimeOK) Error() string {
	return fmt.Sprintf("[PATCH /Customers/{id}/locations/{locationId}/persons/{personId}/appTime][%d] customerPrototypePatchPersonAppTimeOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypePatchPersonAppTimeOK) String() string {
	return fmt.Sprintf("[PATCH /Customers/{id}/locations/{locationId}/persons/{personId}/appTime][%d] customerPrototypePatchPersonAppTimeOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypePatchPersonAppTimeOK) GetPayload() interface{} {
	return o.Payload
}

func (o *CustomerPrototypePatchPersonAppTimeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypePatchPersonAppTimeDefault creates a CustomerPrototypePatchPersonAppTimeDefault with default headers values
func NewCustomerPrototypePatchPersonAppTimeDefault(code int) *CustomerPrototypePatchPersonAppTimeDefault {
	return &CustomerPrototypePatchPersonAppTimeDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypePatchPersonAppTimeDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypePatchPersonAppTimeDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype patch person app time default response has a 2xx status code
func (o *CustomerPrototypePatchPersonAppTimeDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype patch person app time default response has a 3xx status code
func (o *CustomerPrototypePatchPersonAppTimeDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype patch person app time default response has a 4xx status code
func (o *CustomerPrototypePatchPersonAppTimeDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype patch person app time default response has a 5xx status code
func (o *CustomerPrototypePatchPersonAppTimeDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype patch person app time default response a status code equal to that given
func (o *CustomerPrototypePatchPersonAppTimeDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype patch person app time default response
func (o *CustomerPrototypePatchPersonAppTimeDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypePatchPersonAppTimeDefault) Error() string {
	return fmt.Sprintf("[PATCH /Customers/{id}/locations/{locationId}/persons/{personId}/appTime][%d] Customer.prototype.patchPersonAppTime default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePatchPersonAppTimeDefault) String() string {
	return fmt.Sprintf("[PATCH /Customers/{id}/locations/{locationId}/persons/{personId}/appTime][%d] Customer.prototype.patchPersonAppTime default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePatchPersonAppTimeDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypePatchPersonAppTimeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
