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

// CustomerPrototypePostCustomSharedScheduleReader is a Reader for the CustomerPrototypePostCustomSharedSchedule structure.
type CustomerPrototypePostCustomSharedScheduleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypePostCustomSharedScheduleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypePostCustomSharedScheduleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypePostCustomSharedScheduleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypePostCustomSharedScheduleOK creates a CustomerPrototypePostCustomSharedScheduleOK with default headers values
func NewCustomerPrototypePostCustomSharedScheduleOK() *CustomerPrototypePostCustomSharedScheduleOK {
	return &CustomerPrototypePostCustomSharedScheduleOK{}
}

/*
CustomerPrototypePostCustomSharedScheduleOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypePostCustomSharedScheduleOK struct {
	Payload *models.LocationCustomSchedule
}

// IsSuccess returns true when this customer prototype post custom shared schedule o k response has a 2xx status code
func (o *CustomerPrototypePostCustomSharedScheduleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype post custom shared schedule o k response has a 3xx status code
func (o *CustomerPrototypePostCustomSharedScheduleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype post custom shared schedule o k response has a 4xx status code
func (o *CustomerPrototypePostCustomSharedScheduleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype post custom shared schedule o k response has a 5xx status code
func (o *CustomerPrototypePostCustomSharedScheduleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype post custom shared schedule o k response a status code equal to that given
func (o *CustomerPrototypePostCustomSharedScheduleOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype post custom shared schedule o k response
func (o *CustomerPrototypePostCustomSharedScheduleOK) Code() int {
	return 200
}

func (o *CustomerPrototypePostCustomSharedScheduleOK) Error() string {
	return fmt.Sprintf("[POST /Customers/{id}/locations/{locationId}/schedules][%d] customerPrototypePostCustomSharedScheduleOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypePostCustomSharedScheduleOK) String() string {
	return fmt.Sprintf("[POST /Customers/{id}/locations/{locationId}/schedules][%d] customerPrototypePostCustomSharedScheduleOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypePostCustomSharedScheduleOK) GetPayload() *models.LocationCustomSchedule {
	return o.Payload
}

func (o *CustomerPrototypePostCustomSharedScheduleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LocationCustomSchedule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypePostCustomSharedScheduleDefault creates a CustomerPrototypePostCustomSharedScheduleDefault with default headers values
func NewCustomerPrototypePostCustomSharedScheduleDefault(code int) *CustomerPrototypePostCustomSharedScheduleDefault {
	return &CustomerPrototypePostCustomSharedScheduleDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypePostCustomSharedScheduleDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypePostCustomSharedScheduleDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype post custom shared schedule default response has a 2xx status code
func (o *CustomerPrototypePostCustomSharedScheduleDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype post custom shared schedule default response has a 3xx status code
func (o *CustomerPrototypePostCustomSharedScheduleDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype post custom shared schedule default response has a 4xx status code
func (o *CustomerPrototypePostCustomSharedScheduleDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype post custom shared schedule default response has a 5xx status code
func (o *CustomerPrototypePostCustomSharedScheduleDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype post custom shared schedule default response a status code equal to that given
func (o *CustomerPrototypePostCustomSharedScheduleDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype post custom shared schedule default response
func (o *CustomerPrototypePostCustomSharedScheduleDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypePostCustomSharedScheduleDefault) Error() string {
	return fmt.Sprintf("[POST /Customers/{id}/locations/{locationId}/schedules][%d] Customer.prototype.postCustomSharedSchedule default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePostCustomSharedScheduleDefault) String() string {
	return fmt.Sprintf("[POST /Customers/{id}/locations/{locationId}/schedules][%d] Customer.prototype.postCustomSharedSchedule default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePostCustomSharedScheduleDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypePostCustomSharedScheduleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
