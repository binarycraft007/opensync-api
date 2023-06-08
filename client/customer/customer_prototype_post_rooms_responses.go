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

// CustomerPrototypePostRoomsReader is a Reader for the CustomerPrototypePostRooms structure.
type CustomerPrototypePostRoomsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypePostRoomsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypePostRoomsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypePostRoomsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypePostRoomsOK creates a CustomerPrototypePostRoomsOK with default headers values
func NewCustomerPrototypePostRoomsOK() *CustomerPrototypePostRoomsOK {
	return &CustomerPrototypePostRoomsOK{}
}

/*
CustomerPrototypePostRoomsOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypePostRoomsOK struct {
	Payload *models.Room
}

// IsSuccess returns true when this customer prototype post rooms o k response has a 2xx status code
func (o *CustomerPrototypePostRoomsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype post rooms o k response has a 3xx status code
func (o *CustomerPrototypePostRoomsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype post rooms o k response has a 4xx status code
func (o *CustomerPrototypePostRoomsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype post rooms o k response has a 5xx status code
func (o *CustomerPrototypePostRoomsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype post rooms o k response a status code equal to that given
func (o *CustomerPrototypePostRoomsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype post rooms o k response
func (o *CustomerPrototypePostRoomsOK) Code() int {
	return 200
}

func (o *CustomerPrototypePostRoomsOK) Error() string {
	return fmt.Sprintf("[POST /Customers/{id}/locations/{locationId}/rooms][%d] customerPrototypePostRoomsOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypePostRoomsOK) String() string {
	return fmt.Sprintf("[POST /Customers/{id}/locations/{locationId}/rooms][%d] customerPrototypePostRoomsOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypePostRoomsOK) GetPayload() *models.Room {
	return o.Payload
}

func (o *CustomerPrototypePostRoomsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Room)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypePostRoomsDefault creates a CustomerPrototypePostRoomsDefault with default headers values
func NewCustomerPrototypePostRoomsDefault(code int) *CustomerPrototypePostRoomsDefault {
	return &CustomerPrototypePostRoomsDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypePostRoomsDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypePostRoomsDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype post rooms default response has a 2xx status code
func (o *CustomerPrototypePostRoomsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype post rooms default response has a 3xx status code
func (o *CustomerPrototypePostRoomsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype post rooms default response has a 4xx status code
func (o *CustomerPrototypePostRoomsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype post rooms default response has a 5xx status code
func (o *CustomerPrototypePostRoomsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype post rooms default response a status code equal to that given
func (o *CustomerPrototypePostRoomsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype post rooms default response
func (o *CustomerPrototypePostRoomsDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypePostRoomsDefault) Error() string {
	return fmt.Sprintf("[POST /Customers/{id}/locations/{locationId}/rooms][%d] Customer.prototype.postRooms default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePostRoomsDefault) String() string {
	return fmt.Sprintf("[POST /Customers/{id}/locations/{locationId}/rooms][%d] Customer.prototype.postRooms default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePostRoomsDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypePostRoomsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}