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

// CustomerPrototypeGetManagersListForLocationReader is a Reader for the CustomerPrototypeGetManagersListForLocation structure.
type CustomerPrototypeGetManagersListForLocationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypeGetManagersListForLocationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypeGetManagersListForLocationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypeGetManagersListForLocationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypeGetManagersListForLocationOK creates a CustomerPrototypeGetManagersListForLocationOK with default headers values
func NewCustomerPrototypeGetManagersListForLocationOK() *CustomerPrototypeGetManagersListForLocationOK {
	return &CustomerPrototypeGetManagersListForLocationOK{}
}

/*
CustomerPrototypeGetManagersListForLocationOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypeGetManagersListForLocationOK struct {
	Payload interface{}
}

// IsSuccess returns true when this customer prototype get managers list for location o k response has a 2xx status code
func (o *CustomerPrototypeGetManagersListForLocationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype get managers list for location o k response has a 3xx status code
func (o *CustomerPrototypeGetManagersListForLocationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype get managers list for location o k response has a 4xx status code
func (o *CustomerPrototypeGetManagersListForLocationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype get managers list for location o k response has a 5xx status code
func (o *CustomerPrototypeGetManagersListForLocationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype get managers list for location o k response a status code equal to that given
func (o *CustomerPrototypeGetManagersListForLocationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype get managers list for location o k response
func (o *CustomerPrototypeGetManagersListForLocationOK) Code() int {
	return 200
}

func (o *CustomerPrototypeGetManagersListForLocationOK) Error() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/managers][%d] customerPrototypeGetManagersListForLocationOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeGetManagersListForLocationOK) String() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/managers][%d] customerPrototypeGetManagersListForLocationOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeGetManagersListForLocationOK) GetPayload() interface{} {
	return o.Payload
}

func (o *CustomerPrototypeGetManagersListForLocationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypeGetManagersListForLocationDefault creates a CustomerPrototypeGetManagersListForLocationDefault with default headers values
func NewCustomerPrototypeGetManagersListForLocationDefault(code int) *CustomerPrototypeGetManagersListForLocationDefault {
	return &CustomerPrototypeGetManagersListForLocationDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypeGetManagersListForLocationDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypeGetManagersListForLocationDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype get managers list for location default response has a 2xx status code
func (o *CustomerPrototypeGetManagersListForLocationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype get managers list for location default response has a 3xx status code
func (o *CustomerPrototypeGetManagersListForLocationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype get managers list for location default response has a 4xx status code
func (o *CustomerPrototypeGetManagersListForLocationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype get managers list for location default response has a 5xx status code
func (o *CustomerPrototypeGetManagersListForLocationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype get managers list for location default response a status code equal to that given
func (o *CustomerPrototypeGetManagersListForLocationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype get managers list for location default response
func (o *CustomerPrototypeGetManagersListForLocationDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypeGetManagersListForLocationDefault) Error() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/managers][%d] Customer.prototype.getManagersListForLocation default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeGetManagersListForLocationDefault) String() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/managers][%d] Customer.prototype.getManagersListForLocation default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeGetManagersListForLocationDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypeGetManagersListForLocationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
