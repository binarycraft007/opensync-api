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

// CustomerPrototypePutDppEnrollmentsReader is a Reader for the CustomerPrototypePutDppEnrollments structure.
type CustomerPrototypePutDppEnrollmentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypePutDppEnrollmentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewCustomerPrototypePutDppEnrollmentsAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypePutDppEnrollmentsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypePutDppEnrollmentsAccepted creates a CustomerPrototypePutDppEnrollmentsAccepted with default headers values
func NewCustomerPrototypePutDppEnrollmentsAccepted() *CustomerPrototypePutDppEnrollmentsAccepted {
	return &CustomerPrototypePutDppEnrollmentsAccepted{}
}

/*
CustomerPrototypePutDppEnrollmentsAccepted describes a response with status code 202, with default header values.

Request was successful
*/
type CustomerPrototypePutDppEnrollmentsAccepted struct {
	Payload interface{}
}

// IsSuccess returns true when this customer prototype put dpp enrollments accepted response has a 2xx status code
func (o *CustomerPrototypePutDppEnrollmentsAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype put dpp enrollments accepted response has a 3xx status code
func (o *CustomerPrototypePutDppEnrollmentsAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype put dpp enrollments accepted response has a 4xx status code
func (o *CustomerPrototypePutDppEnrollmentsAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype put dpp enrollments accepted response has a 5xx status code
func (o *CustomerPrototypePutDppEnrollmentsAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype put dpp enrollments accepted response a status code equal to that given
func (o *CustomerPrototypePutDppEnrollmentsAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the customer prototype put dpp enrollments accepted response
func (o *CustomerPrototypePutDppEnrollmentsAccepted) Code() int {
	return 202
}

func (o *CustomerPrototypePutDppEnrollmentsAccepted) Error() string {
	return fmt.Sprintf("[PUT /Customers/{id}/locations/{locationId}/dpp/enrollments][%d] customerPrototypePutDppEnrollmentsAccepted  %+v", 202, o.Payload)
}

func (o *CustomerPrototypePutDppEnrollmentsAccepted) String() string {
	return fmt.Sprintf("[PUT /Customers/{id}/locations/{locationId}/dpp/enrollments][%d] customerPrototypePutDppEnrollmentsAccepted  %+v", 202, o.Payload)
}

func (o *CustomerPrototypePutDppEnrollmentsAccepted) GetPayload() interface{} {
	return o.Payload
}

func (o *CustomerPrototypePutDppEnrollmentsAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypePutDppEnrollmentsDefault creates a CustomerPrototypePutDppEnrollmentsDefault with default headers values
func NewCustomerPrototypePutDppEnrollmentsDefault(code int) *CustomerPrototypePutDppEnrollmentsDefault {
	return &CustomerPrototypePutDppEnrollmentsDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypePutDppEnrollmentsDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypePutDppEnrollmentsDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype put dpp enrollments default response has a 2xx status code
func (o *CustomerPrototypePutDppEnrollmentsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype put dpp enrollments default response has a 3xx status code
func (o *CustomerPrototypePutDppEnrollmentsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype put dpp enrollments default response has a 4xx status code
func (o *CustomerPrototypePutDppEnrollmentsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype put dpp enrollments default response has a 5xx status code
func (o *CustomerPrototypePutDppEnrollmentsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype put dpp enrollments default response a status code equal to that given
func (o *CustomerPrototypePutDppEnrollmentsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype put dpp enrollments default response
func (o *CustomerPrototypePutDppEnrollmentsDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypePutDppEnrollmentsDefault) Error() string {
	return fmt.Sprintf("[PUT /Customers/{id}/locations/{locationId}/dpp/enrollments][%d] Customer.prototype.putDppEnrollments default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePutDppEnrollmentsDefault) String() string {
	return fmt.Sprintf("[PUT /Customers/{id}/locations/{locationId}/dpp/enrollments][%d] Customer.prototype.putDppEnrollments default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePutDppEnrollmentsDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypePutDppEnrollmentsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
