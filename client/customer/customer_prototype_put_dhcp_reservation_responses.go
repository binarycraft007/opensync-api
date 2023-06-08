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

// CustomerPrototypePutDhcpReservationReader is a Reader for the CustomerPrototypePutDhcpReservation structure.
type CustomerPrototypePutDhcpReservationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypePutDhcpReservationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypePutDhcpReservationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypePutDhcpReservationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypePutDhcpReservationOK creates a CustomerPrototypePutDhcpReservationOK with default headers values
func NewCustomerPrototypePutDhcpReservationOK() *CustomerPrototypePutDhcpReservationOK {
	return &CustomerPrototypePutDhcpReservationOK{}
}

/*
CustomerPrototypePutDhcpReservationOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypePutDhcpReservationOK struct {
	Payload *models.DhcpReservation
}

// IsSuccess returns true when this customer prototype put dhcp reservation o k response has a 2xx status code
func (o *CustomerPrototypePutDhcpReservationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype put dhcp reservation o k response has a 3xx status code
func (o *CustomerPrototypePutDhcpReservationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype put dhcp reservation o k response has a 4xx status code
func (o *CustomerPrototypePutDhcpReservationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype put dhcp reservation o k response has a 5xx status code
func (o *CustomerPrototypePutDhcpReservationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype put dhcp reservation o k response a status code equal to that given
func (o *CustomerPrototypePutDhcpReservationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype put dhcp reservation o k response
func (o *CustomerPrototypePutDhcpReservationOK) Code() int {
	return 200
}

func (o *CustomerPrototypePutDhcpReservationOK) Error() string {
	return fmt.Sprintf("[PUT /Customers/{id}/locations/{locationId}/networkConfiguration/dhcpReservations/{mac}][%d] customerPrototypePutDhcpReservationOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypePutDhcpReservationOK) String() string {
	return fmt.Sprintf("[PUT /Customers/{id}/locations/{locationId}/networkConfiguration/dhcpReservations/{mac}][%d] customerPrototypePutDhcpReservationOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypePutDhcpReservationOK) GetPayload() *models.DhcpReservation {
	return o.Payload
}

func (o *CustomerPrototypePutDhcpReservationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DhcpReservation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypePutDhcpReservationDefault creates a CustomerPrototypePutDhcpReservationDefault with default headers values
func NewCustomerPrototypePutDhcpReservationDefault(code int) *CustomerPrototypePutDhcpReservationDefault {
	return &CustomerPrototypePutDhcpReservationDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypePutDhcpReservationDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypePutDhcpReservationDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype put dhcp reservation default response has a 2xx status code
func (o *CustomerPrototypePutDhcpReservationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype put dhcp reservation default response has a 3xx status code
func (o *CustomerPrototypePutDhcpReservationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype put dhcp reservation default response has a 4xx status code
func (o *CustomerPrototypePutDhcpReservationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype put dhcp reservation default response has a 5xx status code
func (o *CustomerPrototypePutDhcpReservationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype put dhcp reservation default response a status code equal to that given
func (o *CustomerPrototypePutDhcpReservationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype put dhcp reservation default response
func (o *CustomerPrototypePutDhcpReservationDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypePutDhcpReservationDefault) Error() string {
	return fmt.Sprintf("[PUT /Customers/{id}/locations/{locationId}/networkConfiguration/dhcpReservations/{mac}][%d] Customer.prototype.putDhcpReservation default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePutDhcpReservationDefault) String() string {
	return fmt.Sprintf("[PUT /Customers/{id}/locations/{locationId}/networkConfiguration/dhcpReservations/{mac}][%d] Customer.prototype.putDhcpReservation default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePutDhcpReservationDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypePutDhcpReservationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}