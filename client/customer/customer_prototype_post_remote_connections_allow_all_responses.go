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

// CustomerPrototypePostRemoteConnectionsAllowAllReader is a Reader for the CustomerPrototypePostRemoteConnectionsAllowAll structure.
type CustomerPrototypePostRemoteConnectionsAllowAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypePostRemoteConnectionsAllowAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypePostRemoteConnectionsAllowAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypePostRemoteConnectionsAllowAllDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypePostRemoteConnectionsAllowAllOK creates a CustomerPrototypePostRemoteConnectionsAllowAllOK with default headers values
func NewCustomerPrototypePostRemoteConnectionsAllowAllOK() *CustomerPrototypePostRemoteConnectionsAllowAllOK {
	return &CustomerPrototypePostRemoteConnectionsAllowAllOK{}
}

/*
CustomerPrototypePostRemoteConnectionsAllowAllOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypePostRemoteConnectionsAllowAllOK struct {
	Payload interface{}
}

// IsSuccess returns true when this customer prototype post remote connections allow all o k response has a 2xx status code
func (o *CustomerPrototypePostRemoteConnectionsAllowAllOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype post remote connections allow all o k response has a 3xx status code
func (o *CustomerPrototypePostRemoteConnectionsAllowAllOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype post remote connections allow all o k response has a 4xx status code
func (o *CustomerPrototypePostRemoteConnectionsAllowAllOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype post remote connections allow all o k response has a 5xx status code
func (o *CustomerPrototypePostRemoteConnectionsAllowAllOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype post remote connections allow all o k response a status code equal to that given
func (o *CustomerPrototypePostRemoteConnectionsAllowAllOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype post remote connections allow all o k response
func (o *CustomerPrototypePostRemoteConnectionsAllowAllOK) Code() int {
	return 200
}

func (o *CustomerPrototypePostRemoteConnectionsAllowAllOK) Error() string {
	return fmt.Sprintf("[POST /Customers/{id}/locations/{locationId}/devices/{mac}/securityPolicy/remoteConnections/allowAll][%d] customerPrototypePostRemoteConnectionsAllowAllOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypePostRemoteConnectionsAllowAllOK) String() string {
	return fmt.Sprintf("[POST /Customers/{id}/locations/{locationId}/devices/{mac}/securityPolicy/remoteConnections/allowAll][%d] customerPrototypePostRemoteConnectionsAllowAllOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypePostRemoteConnectionsAllowAllOK) GetPayload() interface{} {
	return o.Payload
}

func (o *CustomerPrototypePostRemoteConnectionsAllowAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypePostRemoteConnectionsAllowAllDefault creates a CustomerPrototypePostRemoteConnectionsAllowAllDefault with default headers values
func NewCustomerPrototypePostRemoteConnectionsAllowAllDefault(code int) *CustomerPrototypePostRemoteConnectionsAllowAllDefault {
	return &CustomerPrototypePostRemoteConnectionsAllowAllDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypePostRemoteConnectionsAllowAllDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypePostRemoteConnectionsAllowAllDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype post remote connections allow all default response has a 2xx status code
func (o *CustomerPrototypePostRemoteConnectionsAllowAllDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype post remote connections allow all default response has a 3xx status code
func (o *CustomerPrototypePostRemoteConnectionsAllowAllDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype post remote connections allow all default response has a 4xx status code
func (o *CustomerPrototypePostRemoteConnectionsAllowAllDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype post remote connections allow all default response has a 5xx status code
func (o *CustomerPrototypePostRemoteConnectionsAllowAllDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype post remote connections allow all default response a status code equal to that given
func (o *CustomerPrototypePostRemoteConnectionsAllowAllDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype post remote connections allow all default response
func (o *CustomerPrototypePostRemoteConnectionsAllowAllDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypePostRemoteConnectionsAllowAllDefault) Error() string {
	return fmt.Sprintf("[POST /Customers/{id}/locations/{locationId}/devices/{mac}/securityPolicy/remoteConnections/allowAll][%d] Customer.prototype.postRemoteConnectionsAllowAll default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePostRemoteConnectionsAllowAllDefault) String() string {
	return fmt.Sprintf("[POST /Customers/{id}/locations/{locationId}/devices/{mac}/securityPolicy/remoteConnections/allowAll][%d] Customer.prototype.postRemoteConnectionsAllowAll default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePostRemoteConnectionsAllowAllDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypePostRemoteConnectionsAllowAllDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
