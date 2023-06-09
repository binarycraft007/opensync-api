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

// CustomerPrototypePostPortForwardReader is a Reader for the CustomerPrototypePostPortForward structure.
type CustomerPrototypePostPortForwardReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypePostPortForwardReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypePostPortForwardOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypePostPortForwardDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypePostPortForwardOK creates a CustomerPrototypePostPortForwardOK with default headers values
func NewCustomerPrototypePostPortForwardOK() *CustomerPrototypePostPortForwardOK {
	return &CustomerPrototypePostPortForwardOK{}
}

/*
CustomerPrototypePostPortForwardOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypePostPortForwardOK struct {
	Payload *models.PortForward
}

// IsSuccess returns true when this customer prototype post port forward o k response has a 2xx status code
func (o *CustomerPrototypePostPortForwardOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype post port forward o k response has a 3xx status code
func (o *CustomerPrototypePostPortForwardOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype post port forward o k response has a 4xx status code
func (o *CustomerPrototypePostPortForwardOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype post port forward o k response has a 5xx status code
func (o *CustomerPrototypePostPortForwardOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype post port forward o k response a status code equal to that given
func (o *CustomerPrototypePostPortForwardOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype post port forward o k response
func (o *CustomerPrototypePostPortForwardOK) Code() int {
	return 200
}

func (o *CustomerPrototypePostPortForwardOK) Error() string {
	return fmt.Sprintf("[POST /Customers/{id}/locations/{locationId}/networkConfiguration/dhcpReservations/{mac}/portForward][%d] customerPrototypePostPortForwardOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypePostPortForwardOK) String() string {
	return fmt.Sprintf("[POST /Customers/{id}/locations/{locationId}/networkConfiguration/dhcpReservations/{mac}/portForward][%d] customerPrototypePostPortForwardOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypePostPortForwardOK) GetPayload() *models.PortForward {
	return o.Payload
}

func (o *CustomerPrototypePostPortForwardOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PortForward)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypePostPortForwardDefault creates a CustomerPrototypePostPortForwardDefault with default headers values
func NewCustomerPrototypePostPortForwardDefault(code int) *CustomerPrototypePostPortForwardDefault {
	return &CustomerPrototypePostPortForwardDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypePostPortForwardDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypePostPortForwardDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype post port forward default response has a 2xx status code
func (o *CustomerPrototypePostPortForwardDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype post port forward default response has a 3xx status code
func (o *CustomerPrototypePostPortForwardDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype post port forward default response has a 4xx status code
func (o *CustomerPrototypePostPortForwardDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype post port forward default response has a 5xx status code
func (o *CustomerPrototypePostPortForwardDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype post port forward default response a status code equal to that given
func (o *CustomerPrototypePostPortForwardDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype post port forward default response
func (o *CustomerPrototypePostPortForwardDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypePostPortForwardDefault) Error() string {
	return fmt.Sprintf("[POST /Customers/{id}/locations/{locationId}/networkConfiguration/dhcpReservations/{mac}/portForward][%d] Customer.prototype.postPortForward default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePostPortForwardDefault) String() string {
	return fmt.Sprintf("[POST /Customers/{id}/locations/{locationId}/networkConfiguration/dhcpReservations/{mac}/portForward][%d] Customer.prototype.postPortForward default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePostPortForwardDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypePostPortForwardDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
