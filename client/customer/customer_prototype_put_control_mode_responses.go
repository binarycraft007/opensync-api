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

// CustomerPrototypePutControlModeReader is a Reader for the CustomerPrototypePutControlMode structure.
type CustomerPrototypePutControlModeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypePutControlModeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypePutControlModeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypePutControlModeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypePutControlModeOK creates a CustomerPrototypePutControlModeOK with default headers values
func NewCustomerPrototypePutControlModeOK() *CustomerPrototypePutControlModeOK {
	return &CustomerPrototypePutControlModeOK{}
}

/*
CustomerPrototypePutControlModeOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypePutControlModeOK struct {
	Payload *models.LocationControlMode
}

// IsSuccess returns true when this customer prototype put control mode o k response has a 2xx status code
func (o *CustomerPrototypePutControlModeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype put control mode o k response has a 3xx status code
func (o *CustomerPrototypePutControlModeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype put control mode o k response has a 4xx status code
func (o *CustomerPrototypePutControlModeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype put control mode o k response has a 5xx status code
func (o *CustomerPrototypePutControlModeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype put control mode o k response a status code equal to that given
func (o *CustomerPrototypePutControlModeOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype put control mode o k response
func (o *CustomerPrototypePutControlModeOK) Code() int {
	return 200
}

func (o *CustomerPrototypePutControlModeOK) Error() string {
	return fmt.Sprintf("[PUT /Customers/{id}/locations/{locationId}/controlMode][%d] customerPrototypePutControlModeOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypePutControlModeOK) String() string {
	return fmt.Sprintf("[PUT /Customers/{id}/locations/{locationId}/controlMode][%d] customerPrototypePutControlModeOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypePutControlModeOK) GetPayload() *models.LocationControlMode {
	return o.Payload
}

func (o *CustomerPrototypePutControlModeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LocationControlMode)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypePutControlModeDefault creates a CustomerPrototypePutControlModeDefault with default headers values
func NewCustomerPrototypePutControlModeDefault(code int) *CustomerPrototypePutControlModeDefault {
	return &CustomerPrototypePutControlModeDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypePutControlModeDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypePutControlModeDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype put control mode default response has a 2xx status code
func (o *CustomerPrototypePutControlModeDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype put control mode default response has a 3xx status code
func (o *CustomerPrototypePutControlModeDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype put control mode default response has a 4xx status code
func (o *CustomerPrototypePutControlModeDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype put control mode default response has a 5xx status code
func (o *CustomerPrototypePutControlModeDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype put control mode default response a status code equal to that given
func (o *CustomerPrototypePutControlModeDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype put control mode default response
func (o *CustomerPrototypePutControlModeDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypePutControlModeDefault) Error() string {
	return fmt.Sprintf("[PUT /Customers/{id}/locations/{locationId}/controlMode][%d] Customer.prototype.putControlMode default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePutControlModeDefault) String() string {
	return fmt.Sprintf("[PUT /Customers/{id}/locations/{locationId}/controlMode][%d] Customer.prototype.putControlMode default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePutControlModeDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypePutControlModeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
