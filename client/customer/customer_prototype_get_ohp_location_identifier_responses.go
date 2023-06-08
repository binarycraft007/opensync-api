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

// CustomerPrototypeGetOhpLocationIdentifierReader is a Reader for the CustomerPrototypeGetOhpLocationIdentifier structure.
type CustomerPrototypeGetOhpLocationIdentifierReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypeGetOhpLocationIdentifierReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypeGetOhpLocationIdentifierOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypeGetOhpLocationIdentifierDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypeGetOhpLocationIdentifierOK creates a CustomerPrototypeGetOhpLocationIdentifierOK with default headers values
func NewCustomerPrototypeGetOhpLocationIdentifierOK() *CustomerPrototypeGetOhpLocationIdentifierOK {
	return &CustomerPrototypeGetOhpLocationIdentifierOK{}
}

/*
CustomerPrototypeGetOhpLocationIdentifierOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypeGetOhpLocationIdentifierOK struct {
	Payload interface{}
}

// IsSuccess returns true when this customer prototype get ohp location identifier o k response has a 2xx status code
func (o *CustomerPrototypeGetOhpLocationIdentifierOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype get ohp location identifier o k response has a 3xx status code
func (o *CustomerPrototypeGetOhpLocationIdentifierOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype get ohp location identifier o k response has a 4xx status code
func (o *CustomerPrototypeGetOhpLocationIdentifierOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype get ohp location identifier o k response has a 5xx status code
func (o *CustomerPrototypeGetOhpLocationIdentifierOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype get ohp location identifier o k response a status code equal to that given
func (o *CustomerPrototypeGetOhpLocationIdentifierOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype get ohp location identifier o k response
func (o *CustomerPrototypeGetOhpLocationIdentifierOK) Code() int {
	return 200
}

func (o *CustomerPrototypeGetOhpLocationIdentifierOK) Error() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/securityPolicy/ohp/locationIdentifier][%d] customerPrototypeGetOhpLocationIdentifierOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeGetOhpLocationIdentifierOK) String() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/securityPolicy/ohp/locationIdentifier][%d] customerPrototypeGetOhpLocationIdentifierOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeGetOhpLocationIdentifierOK) GetPayload() interface{} {
	return o.Payload
}

func (o *CustomerPrototypeGetOhpLocationIdentifierOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypeGetOhpLocationIdentifierDefault creates a CustomerPrototypeGetOhpLocationIdentifierDefault with default headers values
func NewCustomerPrototypeGetOhpLocationIdentifierDefault(code int) *CustomerPrototypeGetOhpLocationIdentifierDefault {
	return &CustomerPrototypeGetOhpLocationIdentifierDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypeGetOhpLocationIdentifierDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypeGetOhpLocationIdentifierDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype get ohp location identifier default response has a 2xx status code
func (o *CustomerPrototypeGetOhpLocationIdentifierDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype get ohp location identifier default response has a 3xx status code
func (o *CustomerPrototypeGetOhpLocationIdentifierDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype get ohp location identifier default response has a 4xx status code
func (o *CustomerPrototypeGetOhpLocationIdentifierDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype get ohp location identifier default response has a 5xx status code
func (o *CustomerPrototypeGetOhpLocationIdentifierDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype get ohp location identifier default response a status code equal to that given
func (o *CustomerPrototypeGetOhpLocationIdentifierDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype get ohp location identifier default response
func (o *CustomerPrototypeGetOhpLocationIdentifierDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypeGetOhpLocationIdentifierDefault) Error() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/securityPolicy/ohp/locationIdentifier][%d] Customer.prototype.getOhpLocationIdentifier default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeGetOhpLocationIdentifierDefault) String() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/securityPolicy/ohp/locationIdentifier][%d] Customer.prototype.getOhpLocationIdentifier default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeGetOhpLocationIdentifierDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypeGetOhpLocationIdentifierDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
