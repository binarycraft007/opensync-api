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

// CustomerPrototypeGetDashboardReader is a Reader for the CustomerPrototypeGetDashboard structure.
type CustomerPrototypeGetDashboardReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypeGetDashboardReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypeGetDashboardOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypeGetDashboardDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypeGetDashboardOK creates a CustomerPrototypeGetDashboardOK with default headers values
func NewCustomerPrototypeGetDashboardOK() *CustomerPrototypeGetDashboardOK {
	return &CustomerPrototypeGetDashboardOK{}
}

/*
CustomerPrototypeGetDashboardOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypeGetDashboardOK struct {
	Payload models.XAny
}

// IsSuccess returns true when this customer prototype get dashboard o k response has a 2xx status code
func (o *CustomerPrototypeGetDashboardOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype get dashboard o k response has a 3xx status code
func (o *CustomerPrototypeGetDashboardOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype get dashboard o k response has a 4xx status code
func (o *CustomerPrototypeGetDashboardOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype get dashboard o k response has a 5xx status code
func (o *CustomerPrototypeGetDashboardOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype get dashboard o k response a status code equal to that given
func (o *CustomerPrototypeGetDashboardOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype get dashboard o k response
func (o *CustomerPrototypeGetDashboardOK) Code() int {
	return 200
}

func (o *CustomerPrototypeGetDashboardOK) Error() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/flex/dashboard][%d] customerPrototypeGetDashboardOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeGetDashboardOK) String() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/flex/dashboard][%d] customerPrototypeGetDashboardOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeGetDashboardOK) GetPayload() models.XAny {
	return o.Payload
}

func (o *CustomerPrototypeGetDashboardOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypeGetDashboardDefault creates a CustomerPrototypeGetDashboardDefault with default headers values
func NewCustomerPrototypeGetDashboardDefault(code int) *CustomerPrototypeGetDashboardDefault {
	return &CustomerPrototypeGetDashboardDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypeGetDashboardDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypeGetDashboardDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype get dashboard default response has a 2xx status code
func (o *CustomerPrototypeGetDashboardDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype get dashboard default response has a 3xx status code
func (o *CustomerPrototypeGetDashboardDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype get dashboard default response has a 4xx status code
func (o *CustomerPrototypeGetDashboardDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype get dashboard default response has a 5xx status code
func (o *CustomerPrototypeGetDashboardDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype get dashboard default response a status code equal to that given
func (o *CustomerPrototypeGetDashboardDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype get dashboard default response
func (o *CustomerPrototypeGetDashboardDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypeGetDashboardDefault) Error() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/flex/dashboard][%d] Customer.prototype.getDashboard default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeGetDashboardDefault) String() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/flex/dashboard][%d] Customer.prototype.getDashboard default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeGetDashboardDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypeGetDashboardDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}