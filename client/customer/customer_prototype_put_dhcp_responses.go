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

// CustomerPrototypePutDhcpReader is a Reader for the CustomerPrototypePutDhcp structure.
type CustomerPrototypePutDhcpReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypePutDhcpReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypePutDhcpOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypePutDhcpDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypePutDhcpOK creates a CustomerPrototypePutDhcpOK with default headers values
func NewCustomerPrototypePutDhcpOK() *CustomerPrototypePutDhcpOK {
	return &CustomerPrototypePutDhcpOK{}
}

/*
CustomerPrototypePutDhcpOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypePutDhcpOK struct {
	Payload interface{}
}

// IsSuccess returns true when this customer prototype put dhcp o k response has a 2xx status code
func (o *CustomerPrototypePutDhcpOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype put dhcp o k response has a 3xx status code
func (o *CustomerPrototypePutDhcpOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype put dhcp o k response has a 4xx status code
func (o *CustomerPrototypePutDhcpOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype put dhcp o k response has a 5xx status code
func (o *CustomerPrototypePutDhcpOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype put dhcp o k response a status code equal to that given
func (o *CustomerPrototypePutDhcpOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype put dhcp o k response
func (o *CustomerPrototypePutDhcpOK) Code() int {
	return 200
}

func (o *CustomerPrototypePutDhcpOK) Error() string {
	return fmt.Sprintf("[PUT /Customers/{id}/locations/{locationId}/networkConfiguration/dhcp][%d] customerPrototypePutDhcpOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypePutDhcpOK) String() string {
	return fmt.Sprintf("[PUT /Customers/{id}/locations/{locationId}/networkConfiguration/dhcp][%d] customerPrototypePutDhcpOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypePutDhcpOK) GetPayload() interface{} {
	return o.Payload
}

func (o *CustomerPrototypePutDhcpOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypePutDhcpDefault creates a CustomerPrototypePutDhcpDefault with default headers values
func NewCustomerPrototypePutDhcpDefault(code int) *CustomerPrototypePutDhcpDefault {
	return &CustomerPrototypePutDhcpDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypePutDhcpDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypePutDhcpDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype put dhcp default response has a 2xx status code
func (o *CustomerPrototypePutDhcpDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype put dhcp default response has a 3xx status code
func (o *CustomerPrototypePutDhcpDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype put dhcp default response has a 4xx status code
func (o *CustomerPrototypePutDhcpDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype put dhcp default response has a 5xx status code
func (o *CustomerPrototypePutDhcpDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype put dhcp default response a status code equal to that given
func (o *CustomerPrototypePutDhcpDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype put dhcp default response
func (o *CustomerPrototypePutDhcpDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypePutDhcpDefault) Error() string {
	return fmt.Sprintf("[PUT /Customers/{id}/locations/{locationId}/networkConfiguration/dhcp][%d] Customer.prototype.putDhcp default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePutDhcpDefault) String() string {
	return fmt.Sprintf("[PUT /Customers/{id}/locations/{locationId}/networkConfiguration/dhcp][%d] Customer.prototype.putDhcp default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePutDhcpDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypePutDhcpDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
