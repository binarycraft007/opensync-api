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

// CustomerPrototypePatchLocationManagerReader is a Reader for the CustomerPrototypePatchLocationManager structure.
type CustomerPrototypePatchLocationManagerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypePatchLocationManagerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypePatchLocationManagerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypePatchLocationManagerDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypePatchLocationManagerOK creates a CustomerPrototypePatchLocationManagerOK with default headers values
func NewCustomerPrototypePatchLocationManagerOK() *CustomerPrototypePatchLocationManagerOK {
	return &CustomerPrototypePatchLocationManagerOK{}
}

/*
CustomerPrototypePatchLocationManagerOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypePatchLocationManagerOK struct {
	Payload *models.Location
}

// IsSuccess returns true when this customer prototype patch location manager o k response has a 2xx status code
func (o *CustomerPrototypePatchLocationManagerOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype patch location manager o k response has a 3xx status code
func (o *CustomerPrototypePatchLocationManagerOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype patch location manager o k response has a 4xx status code
func (o *CustomerPrototypePatchLocationManagerOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype patch location manager o k response has a 5xx status code
func (o *CustomerPrototypePatchLocationManagerOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype patch location manager o k response a status code equal to that given
func (o *CustomerPrototypePatchLocationManagerOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype patch location manager o k response
func (o *CustomerPrototypePatchLocationManagerOK) Code() int {
	return 200
}

func (o *CustomerPrototypePatchLocationManagerOK) Error() string {
	return fmt.Sprintf("[PATCH /Customers/{id}/locations/{locationId}/managers/{managerId}][%d] customerPrototypePatchLocationManagerOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypePatchLocationManagerOK) String() string {
	return fmt.Sprintf("[PATCH /Customers/{id}/locations/{locationId}/managers/{managerId}][%d] customerPrototypePatchLocationManagerOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypePatchLocationManagerOK) GetPayload() *models.Location {
	return o.Payload
}

func (o *CustomerPrototypePatchLocationManagerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Location)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypePatchLocationManagerDefault creates a CustomerPrototypePatchLocationManagerDefault with default headers values
func NewCustomerPrototypePatchLocationManagerDefault(code int) *CustomerPrototypePatchLocationManagerDefault {
	return &CustomerPrototypePatchLocationManagerDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypePatchLocationManagerDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypePatchLocationManagerDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype patch location manager default response has a 2xx status code
func (o *CustomerPrototypePatchLocationManagerDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype patch location manager default response has a 3xx status code
func (o *CustomerPrototypePatchLocationManagerDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype patch location manager default response has a 4xx status code
func (o *CustomerPrototypePatchLocationManagerDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype patch location manager default response has a 5xx status code
func (o *CustomerPrototypePatchLocationManagerDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype patch location manager default response a status code equal to that given
func (o *CustomerPrototypePatchLocationManagerDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype patch location manager default response
func (o *CustomerPrototypePatchLocationManagerDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypePatchLocationManagerDefault) Error() string {
	return fmt.Sprintf("[PATCH /Customers/{id}/locations/{locationId}/managers/{managerId}][%d] Customer.prototype.patchLocationManager default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePatchLocationManagerDefault) String() string {
	return fmt.Sprintf("[PATCH /Customers/{id}/locations/{locationId}/managers/{managerId}][%d] Customer.prototype.patchLocationManager default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePatchLocationManagerDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypePatchLocationManagerDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}