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

// CustomerPrototypeGetLocationGuardPersonEventsSummaryReader is a Reader for the CustomerPrototypeGetLocationGuardPersonEventsSummary structure.
type CustomerPrototypeGetLocationGuardPersonEventsSummaryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypeGetLocationGuardPersonEventsSummaryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypeGetLocationGuardPersonEventsSummaryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypeGetLocationGuardPersonEventsSummaryDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypeGetLocationGuardPersonEventsSummaryOK creates a CustomerPrototypeGetLocationGuardPersonEventsSummaryOK with default headers values
func NewCustomerPrototypeGetLocationGuardPersonEventsSummaryOK() *CustomerPrototypeGetLocationGuardPersonEventsSummaryOK {
	return &CustomerPrototypeGetLocationGuardPersonEventsSummaryOK{}
}

/*
CustomerPrototypeGetLocationGuardPersonEventsSummaryOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypeGetLocationGuardPersonEventsSummaryOK struct {
	Payload models.XAny
}

// IsSuccess returns true when this customer prototype get location guard person events summary o k response has a 2xx status code
func (o *CustomerPrototypeGetLocationGuardPersonEventsSummaryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype get location guard person events summary o k response has a 3xx status code
func (o *CustomerPrototypeGetLocationGuardPersonEventsSummaryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype get location guard person events summary o k response has a 4xx status code
func (o *CustomerPrototypeGetLocationGuardPersonEventsSummaryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype get location guard person events summary o k response has a 5xx status code
func (o *CustomerPrototypeGetLocationGuardPersonEventsSummaryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype get location guard person events summary o k response a status code equal to that given
func (o *CustomerPrototypeGetLocationGuardPersonEventsSummaryOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype get location guard person events summary o k response
func (o *CustomerPrototypeGetLocationGuardPersonEventsSummaryOK) Code() int {
	return 200
}

func (o *CustomerPrototypeGetLocationGuardPersonEventsSummaryOK) Error() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/securityPolicy/guard/personEventsSummary][%d] customerPrototypeGetLocationGuardPersonEventsSummaryOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeGetLocationGuardPersonEventsSummaryOK) String() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/securityPolicy/guard/personEventsSummary][%d] customerPrototypeGetLocationGuardPersonEventsSummaryOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeGetLocationGuardPersonEventsSummaryOK) GetPayload() models.XAny {
	return o.Payload
}

func (o *CustomerPrototypeGetLocationGuardPersonEventsSummaryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypeGetLocationGuardPersonEventsSummaryDefault creates a CustomerPrototypeGetLocationGuardPersonEventsSummaryDefault with default headers values
func NewCustomerPrototypeGetLocationGuardPersonEventsSummaryDefault(code int) *CustomerPrototypeGetLocationGuardPersonEventsSummaryDefault {
	return &CustomerPrototypeGetLocationGuardPersonEventsSummaryDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypeGetLocationGuardPersonEventsSummaryDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypeGetLocationGuardPersonEventsSummaryDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype get location guard person events summary default response has a 2xx status code
func (o *CustomerPrototypeGetLocationGuardPersonEventsSummaryDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype get location guard person events summary default response has a 3xx status code
func (o *CustomerPrototypeGetLocationGuardPersonEventsSummaryDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype get location guard person events summary default response has a 4xx status code
func (o *CustomerPrototypeGetLocationGuardPersonEventsSummaryDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype get location guard person events summary default response has a 5xx status code
func (o *CustomerPrototypeGetLocationGuardPersonEventsSummaryDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype get location guard person events summary default response a status code equal to that given
func (o *CustomerPrototypeGetLocationGuardPersonEventsSummaryDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype get location guard person events summary default response
func (o *CustomerPrototypeGetLocationGuardPersonEventsSummaryDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypeGetLocationGuardPersonEventsSummaryDefault) Error() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/securityPolicy/guard/personEventsSummary][%d] Customer.prototype.getLocationGuardPersonEventsSummary default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeGetLocationGuardPersonEventsSummaryDefault) String() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/securityPolicy/guard/personEventsSummary][%d] Customer.prototype.getLocationGuardPersonEventsSummary default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeGetLocationGuardPersonEventsSummaryDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypeGetLocationGuardPersonEventsSummaryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
