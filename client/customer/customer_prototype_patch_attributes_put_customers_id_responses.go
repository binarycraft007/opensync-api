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

// CustomerPrototypePatchAttributesPutCustomersIDReader is a Reader for the CustomerPrototypePatchAttributesPutCustomersID structure.
type CustomerPrototypePatchAttributesPutCustomersIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypePatchAttributesPutCustomersIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypePatchAttributesPutCustomersIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypePatchAttributesPutCustomersIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypePatchAttributesPutCustomersIDOK creates a CustomerPrototypePatchAttributesPutCustomersIDOK with default headers values
func NewCustomerPrototypePatchAttributesPutCustomersIDOK() *CustomerPrototypePatchAttributesPutCustomersIDOK {
	return &CustomerPrototypePatchAttributesPutCustomersIDOK{}
}

/*
CustomerPrototypePatchAttributesPutCustomersIDOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypePatchAttributesPutCustomersIDOK struct {
	Payload *models.Customer
}

// IsSuccess returns true when this customer prototype patch attributes put customers Id o k response has a 2xx status code
func (o *CustomerPrototypePatchAttributesPutCustomersIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype patch attributes put customers Id o k response has a 3xx status code
func (o *CustomerPrototypePatchAttributesPutCustomersIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype patch attributes put customers Id o k response has a 4xx status code
func (o *CustomerPrototypePatchAttributesPutCustomersIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype patch attributes put customers Id o k response has a 5xx status code
func (o *CustomerPrototypePatchAttributesPutCustomersIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype patch attributes put customers Id o k response a status code equal to that given
func (o *CustomerPrototypePatchAttributesPutCustomersIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype patch attributes put customers Id o k response
func (o *CustomerPrototypePatchAttributesPutCustomersIDOK) Code() int {
	return 200
}

func (o *CustomerPrototypePatchAttributesPutCustomersIDOK) Error() string {
	return fmt.Sprintf("[PUT /Customers/{id}][%d] customerPrototypePatchAttributesPutCustomersIdOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypePatchAttributesPutCustomersIDOK) String() string {
	return fmt.Sprintf("[PUT /Customers/{id}][%d] customerPrototypePatchAttributesPutCustomersIdOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypePatchAttributesPutCustomersIDOK) GetPayload() *models.Customer {
	return o.Payload
}

func (o *CustomerPrototypePatchAttributesPutCustomersIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Customer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypePatchAttributesPutCustomersIDDefault creates a CustomerPrototypePatchAttributesPutCustomersIDDefault with default headers values
func NewCustomerPrototypePatchAttributesPutCustomersIDDefault(code int) *CustomerPrototypePatchAttributesPutCustomersIDDefault {
	return &CustomerPrototypePatchAttributesPutCustomersIDDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypePatchAttributesPutCustomersIDDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypePatchAttributesPutCustomersIDDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype patch attributes put customers id default response has a 2xx status code
func (o *CustomerPrototypePatchAttributesPutCustomersIDDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype patch attributes put customers id default response has a 3xx status code
func (o *CustomerPrototypePatchAttributesPutCustomersIDDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype patch attributes put customers id default response has a 4xx status code
func (o *CustomerPrototypePatchAttributesPutCustomersIDDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype patch attributes put customers id default response has a 5xx status code
func (o *CustomerPrototypePatchAttributesPutCustomersIDDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype patch attributes put customers id default response a status code equal to that given
func (o *CustomerPrototypePatchAttributesPutCustomersIDDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype patch attributes put customers id default response
func (o *CustomerPrototypePatchAttributesPutCustomersIDDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypePatchAttributesPutCustomersIDDefault) Error() string {
	return fmt.Sprintf("[PUT /Customers/{id}][%d] Customer.prototype.patchAttributes__put_Customers_{id} default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePatchAttributesPutCustomersIDDefault) String() string {
	return fmt.Sprintf("[PUT /Customers/{id}][%d] Customer.prototype.patchAttributes__put_Customers_{id} default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePatchAttributesPutCustomersIDDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypePatchAttributesPutCustomersIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
