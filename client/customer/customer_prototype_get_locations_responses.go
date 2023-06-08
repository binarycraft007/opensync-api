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

// CustomerPrototypeGetLocationsReader is a Reader for the CustomerPrototypeGetLocations structure.
type CustomerPrototypeGetLocationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypeGetLocationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypeGetLocationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypeGetLocationsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypeGetLocationsOK creates a CustomerPrototypeGetLocationsOK with default headers values
func NewCustomerPrototypeGetLocationsOK() *CustomerPrototypeGetLocationsOK {
	return &CustomerPrototypeGetLocationsOK{}
}

/*
CustomerPrototypeGetLocationsOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypeGetLocationsOK struct {
	Payload *models.Location
}

// IsSuccess returns true when this customer prototype get locations o k response has a 2xx status code
func (o *CustomerPrototypeGetLocationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype get locations o k response has a 3xx status code
func (o *CustomerPrototypeGetLocationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype get locations o k response has a 4xx status code
func (o *CustomerPrototypeGetLocationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype get locations o k response has a 5xx status code
func (o *CustomerPrototypeGetLocationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype get locations o k response a status code equal to that given
func (o *CustomerPrototypeGetLocationsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype get locations o k response
func (o *CustomerPrototypeGetLocationsOK) Code() int {
	return 200
}

func (o *CustomerPrototypeGetLocationsOK) Error() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations][%d] customerPrototypeGetLocationsOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeGetLocationsOK) String() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations][%d] customerPrototypeGetLocationsOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeGetLocationsOK) GetPayload() *models.Location {
	return o.Payload
}

func (o *CustomerPrototypeGetLocationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Location)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypeGetLocationsDefault creates a CustomerPrototypeGetLocationsDefault with default headers values
func NewCustomerPrototypeGetLocationsDefault(code int) *CustomerPrototypeGetLocationsDefault {
	return &CustomerPrototypeGetLocationsDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypeGetLocationsDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypeGetLocationsDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype get locations default response has a 2xx status code
func (o *CustomerPrototypeGetLocationsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype get locations default response has a 3xx status code
func (o *CustomerPrototypeGetLocationsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype get locations default response has a 4xx status code
func (o *CustomerPrototypeGetLocationsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype get locations default response has a 5xx status code
func (o *CustomerPrototypeGetLocationsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype get locations default response a status code equal to that given
func (o *CustomerPrototypeGetLocationsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype get locations default response
func (o *CustomerPrototypeGetLocationsDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypeGetLocationsDefault) Error() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations][%d] Customer.prototype.getLocations default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeGetLocationsDefault) String() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations][%d] Customer.prototype.getLocations default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeGetLocationsDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypeGetLocationsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
