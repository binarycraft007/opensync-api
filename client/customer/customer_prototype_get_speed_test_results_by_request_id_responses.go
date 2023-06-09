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

// CustomerPrototypeGetSpeedTestResultsByRequestIDReader is a Reader for the CustomerPrototypeGetSpeedTestResultsByRequestID structure.
type CustomerPrototypeGetSpeedTestResultsByRequestIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypeGetSpeedTestResultsByRequestIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypeGetSpeedTestResultsByRequestIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypeGetSpeedTestResultsByRequestIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypeGetSpeedTestResultsByRequestIDOK creates a CustomerPrototypeGetSpeedTestResultsByRequestIDOK with default headers values
func NewCustomerPrototypeGetSpeedTestResultsByRequestIDOK() *CustomerPrototypeGetSpeedTestResultsByRequestIDOK {
	return &CustomerPrototypeGetSpeedTestResultsByRequestIDOK{}
}

/*
CustomerPrototypeGetSpeedTestResultsByRequestIDOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypeGetSpeedTestResultsByRequestIDOK struct {
	Payload interface{}
}

// IsSuccess returns true when this customer prototype get speed test results by request Id o k response has a 2xx status code
func (o *CustomerPrototypeGetSpeedTestResultsByRequestIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype get speed test results by request Id o k response has a 3xx status code
func (o *CustomerPrototypeGetSpeedTestResultsByRequestIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype get speed test results by request Id o k response has a 4xx status code
func (o *CustomerPrototypeGetSpeedTestResultsByRequestIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype get speed test results by request Id o k response has a 5xx status code
func (o *CustomerPrototypeGetSpeedTestResultsByRequestIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype get speed test results by request Id o k response a status code equal to that given
func (o *CustomerPrototypeGetSpeedTestResultsByRequestIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype get speed test results by request Id o k response
func (o *CustomerPrototypeGetSpeedTestResultsByRequestIDOK) Code() int {
	return 200
}

func (o *CustomerPrototypeGetSpeedTestResultsByRequestIDOK) Error() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/nodes/{nodeId}/speedTestResults/{requestId}][%d] customerPrototypeGetSpeedTestResultsByRequestIdOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeGetSpeedTestResultsByRequestIDOK) String() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/nodes/{nodeId}/speedTestResults/{requestId}][%d] customerPrototypeGetSpeedTestResultsByRequestIdOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeGetSpeedTestResultsByRequestIDOK) GetPayload() interface{} {
	return o.Payload
}

func (o *CustomerPrototypeGetSpeedTestResultsByRequestIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypeGetSpeedTestResultsByRequestIDDefault creates a CustomerPrototypeGetSpeedTestResultsByRequestIDDefault with default headers values
func NewCustomerPrototypeGetSpeedTestResultsByRequestIDDefault(code int) *CustomerPrototypeGetSpeedTestResultsByRequestIDDefault {
	return &CustomerPrototypeGetSpeedTestResultsByRequestIDDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypeGetSpeedTestResultsByRequestIDDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypeGetSpeedTestResultsByRequestIDDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype get speed test results by request Id default response has a 2xx status code
func (o *CustomerPrototypeGetSpeedTestResultsByRequestIDDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype get speed test results by request Id default response has a 3xx status code
func (o *CustomerPrototypeGetSpeedTestResultsByRequestIDDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype get speed test results by request Id default response has a 4xx status code
func (o *CustomerPrototypeGetSpeedTestResultsByRequestIDDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype get speed test results by request Id default response has a 5xx status code
func (o *CustomerPrototypeGetSpeedTestResultsByRequestIDDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype get speed test results by request Id default response a status code equal to that given
func (o *CustomerPrototypeGetSpeedTestResultsByRequestIDDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype get speed test results by request Id default response
func (o *CustomerPrototypeGetSpeedTestResultsByRequestIDDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypeGetSpeedTestResultsByRequestIDDefault) Error() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/nodes/{nodeId}/speedTestResults/{requestId}][%d] Customer.prototype.getSpeedTestResultsByRequestId default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeGetSpeedTestResultsByRequestIDDefault) String() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/nodes/{nodeId}/speedTestResults/{requestId}][%d] Customer.prototype.getSpeedTestResultsByRequestId default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeGetSpeedTestResultsByRequestIDDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypeGetSpeedTestResultsByRequestIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
