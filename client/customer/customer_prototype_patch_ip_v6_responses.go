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

// CustomerPrototypePatchIPV6Reader is a Reader for the CustomerPrototypePatchIPV6 structure.
type CustomerPrototypePatchIPV6Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypePatchIPV6Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypePatchIPV6OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypePatchIPV6Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypePatchIPV6OK creates a CustomerPrototypePatchIPV6OK with default headers values
func NewCustomerPrototypePatchIPV6OK() *CustomerPrototypePatchIPV6OK {
	return &CustomerPrototypePatchIPV6OK{}
}

/*
CustomerPrototypePatchIPV6OK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypePatchIPV6OK struct {
	Payload interface{}
}

// IsSuccess returns true when this customer prototype patch Ipv6 o k response has a 2xx status code
func (o *CustomerPrototypePatchIPV6OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype patch Ipv6 o k response has a 3xx status code
func (o *CustomerPrototypePatchIPV6OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype patch Ipv6 o k response has a 4xx status code
func (o *CustomerPrototypePatchIPV6OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype patch Ipv6 o k response has a 5xx status code
func (o *CustomerPrototypePatchIPV6OK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype patch Ipv6 o k response a status code equal to that given
func (o *CustomerPrototypePatchIPV6OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype patch Ipv6 o k response
func (o *CustomerPrototypePatchIPV6OK) Code() int {
	return 200
}

func (o *CustomerPrototypePatchIPV6OK) Error() string {
	return fmt.Sprintf("[PATCH /Customers/{id}/locations/{locationId}/ipv6][%d] customerPrototypePatchIpv6OK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypePatchIPV6OK) String() string {
	return fmt.Sprintf("[PATCH /Customers/{id}/locations/{locationId}/ipv6][%d] customerPrototypePatchIpv6OK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypePatchIPV6OK) GetPayload() interface{} {
	return o.Payload
}

func (o *CustomerPrototypePatchIPV6OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypePatchIPV6Default creates a CustomerPrototypePatchIPV6Default with default headers values
func NewCustomerPrototypePatchIPV6Default(code int) *CustomerPrototypePatchIPV6Default {
	return &CustomerPrototypePatchIPV6Default{
		_statusCode: code,
	}
}

/*
CustomerPrototypePatchIPV6Default describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypePatchIPV6Default struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype patch IPv6 default response has a 2xx status code
func (o *CustomerPrototypePatchIPV6Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype patch IPv6 default response has a 3xx status code
func (o *CustomerPrototypePatchIPV6Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype patch IPv6 default response has a 4xx status code
func (o *CustomerPrototypePatchIPV6Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype patch IPv6 default response has a 5xx status code
func (o *CustomerPrototypePatchIPV6Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype patch IPv6 default response a status code equal to that given
func (o *CustomerPrototypePatchIPV6Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype patch IPv6 default response
func (o *CustomerPrototypePatchIPV6Default) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypePatchIPV6Default) Error() string {
	return fmt.Sprintf("[PATCH /Customers/{id}/locations/{locationId}/ipv6][%d] Customer.prototype.patchIPv6 default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePatchIPV6Default) String() string {
	return fmt.Sprintf("[PATCH /Customers/{id}/locations/{locationId}/ipv6][%d] Customer.prototype.patchIPv6 default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePatchIPV6Default) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypePatchIPV6Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
