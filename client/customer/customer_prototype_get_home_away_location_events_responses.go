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

// CustomerPrototypeGetHomeAwayLocationEventsReader is a Reader for the CustomerPrototypeGetHomeAwayLocationEvents structure.
type CustomerPrototypeGetHomeAwayLocationEventsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypeGetHomeAwayLocationEventsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypeGetHomeAwayLocationEventsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypeGetHomeAwayLocationEventsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypeGetHomeAwayLocationEventsOK creates a CustomerPrototypeGetHomeAwayLocationEventsOK with default headers values
func NewCustomerPrototypeGetHomeAwayLocationEventsOK() *CustomerPrototypeGetHomeAwayLocationEventsOK {
	return &CustomerPrototypeGetHomeAwayLocationEventsOK{}
}

/*
CustomerPrototypeGetHomeAwayLocationEventsOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypeGetHomeAwayLocationEventsOK struct {
	Payload []interface{}
}

// IsSuccess returns true when this customer prototype get home away location events o k response has a 2xx status code
func (o *CustomerPrototypeGetHomeAwayLocationEventsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype get home away location events o k response has a 3xx status code
func (o *CustomerPrototypeGetHomeAwayLocationEventsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype get home away location events o k response has a 4xx status code
func (o *CustomerPrototypeGetHomeAwayLocationEventsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype get home away location events o k response has a 5xx status code
func (o *CustomerPrototypeGetHomeAwayLocationEventsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype get home away location events o k response a status code equal to that given
func (o *CustomerPrototypeGetHomeAwayLocationEventsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype get home away location events o k response
func (o *CustomerPrototypeGetHomeAwayLocationEventsOK) Code() int {
	return 200
}

func (o *CustomerPrototypeGetHomeAwayLocationEventsOK) Error() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/homeAway/events][%d] customerPrototypeGetHomeAwayLocationEventsOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeGetHomeAwayLocationEventsOK) String() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/homeAway/events][%d] customerPrototypeGetHomeAwayLocationEventsOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeGetHomeAwayLocationEventsOK) GetPayload() []interface{} {
	return o.Payload
}

func (o *CustomerPrototypeGetHomeAwayLocationEventsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypeGetHomeAwayLocationEventsDefault creates a CustomerPrototypeGetHomeAwayLocationEventsDefault with default headers values
func NewCustomerPrototypeGetHomeAwayLocationEventsDefault(code int) *CustomerPrototypeGetHomeAwayLocationEventsDefault {
	return &CustomerPrototypeGetHomeAwayLocationEventsDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypeGetHomeAwayLocationEventsDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypeGetHomeAwayLocationEventsDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype get home away location events default response has a 2xx status code
func (o *CustomerPrototypeGetHomeAwayLocationEventsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype get home away location events default response has a 3xx status code
func (o *CustomerPrototypeGetHomeAwayLocationEventsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype get home away location events default response has a 4xx status code
func (o *CustomerPrototypeGetHomeAwayLocationEventsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype get home away location events default response has a 5xx status code
func (o *CustomerPrototypeGetHomeAwayLocationEventsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype get home away location events default response a status code equal to that given
func (o *CustomerPrototypeGetHomeAwayLocationEventsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype get home away location events default response
func (o *CustomerPrototypeGetHomeAwayLocationEventsDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypeGetHomeAwayLocationEventsDefault) Error() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/homeAway/events][%d] Customer.prototype.getHomeAwayLocationEvents default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeGetHomeAwayLocationEventsDefault) String() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/homeAway/events][%d] Customer.prototype.getHomeAwayLocationEvents default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeGetHomeAwayLocationEventsDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypeGetHomeAwayLocationEventsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
