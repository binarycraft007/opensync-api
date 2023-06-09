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

// CustomerPrototypeGetDhcpReader is a Reader for the CustomerPrototypeGetDhcp structure.
type CustomerPrototypeGetDhcpReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypeGetDhcpReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypeGetDhcpOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypeGetDhcpDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypeGetDhcpOK creates a CustomerPrototypeGetDhcpOK with default headers values
func NewCustomerPrototypeGetDhcpOK() *CustomerPrototypeGetDhcpOK {
	return &CustomerPrototypeGetDhcpOK{}
}

/*
CustomerPrototypeGetDhcpOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypeGetDhcpOK struct {
	Payload interface{}
}

// IsSuccess returns true when this customer prototype get dhcp o k response has a 2xx status code
func (o *CustomerPrototypeGetDhcpOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype get dhcp o k response has a 3xx status code
func (o *CustomerPrototypeGetDhcpOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype get dhcp o k response has a 4xx status code
func (o *CustomerPrototypeGetDhcpOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype get dhcp o k response has a 5xx status code
func (o *CustomerPrototypeGetDhcpOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype get dhcp o k response a status code equal to that given
func (o *CustomerPrototypeGetDhcpOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype get dhcp o k response
func (o *CustomerPrototypeGetDhcpOK) Code() int {
	return 200
}

func (o *CustomerPrototypeGetDhcpOK) Error() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/networkConfiguration/dhcp][%d] customerPrototypeGetDhcpOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeGetDhcpOK) String() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/networkConfiguration/dhcp][%d] customerPrototypeGetDhcpOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeGetDhcpOK) GetPayload() interface{} {
	return o.Payload
}

func (o *CustomerPrototypeGetDhcpOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypeGetDhcpDefault creates a CustomerPrototypeGetDhcpDefault with default headers values
func NewCustomerPrototypeGetDhcpDefault(code int) *CustomerPrototypeGetDhcpDefault {
	return &CustomerPrototypeGetDhcpDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypeGetDhcpDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypeGetDhcpDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype get dhcp default response has a 2xx status code
func (o *CustomerPrototypeGetDhcpDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype get dhcp default response has a 3xx status code
func (o *CustomerPrototypeGetDhcpDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype get dhcp default response has a 4xx status code
func (o *CustomerPrototypeGetDhcpDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype get dhcp default response has a 5xx status code
func (o *CustomerPrototypeGetDhcpDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype get dhcp default response a status code equal to that given
func (o *CustomerPrototypeGetDhcpDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype get dhcp default response
func (o *CustomerPrototypeGetDhcpDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypeGetDhcpDefault) Error() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/networkConfiguration/dhcp][%d] Customer.prototype.getDhcp default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeGetDhcpDefault) String() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/networkConfiguration/dhcp][%d] Customer.prototype.getDhcp default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeGetDhcpDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypeGetDhcpDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
