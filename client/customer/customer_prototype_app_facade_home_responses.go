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

// CustomerPrototypeAppFacadeHomeReader is a Reader for the CustomerPrototypeAppFacadeHome structure.
type CustomerPrototypeAppFacadeHomeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypeAppFacadeHomeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypeAppFacadeHomeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypeAppFacadeHomeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypeAppFacadeHomeOK creates a CustomerPrototypeAppFacadeHomeOK with default headers values
func NewCustomerPrototypeAppFacadeHomeOK() *CustomerPrototypeAppFacadeHomeOK {
	return &CustomerPrototypeAppFacadeHomeOK{}
}

/*
CustomerPrototypeAppFacadeHomeOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypeAppFacadeHomeOK struct {
	Payload *models.AppFacadeHomeResponse
}

// IsSuccess returns true when this customer prototype app facade home o k response has a 2xx status code
func (o *CustomerPrototypeAppFacadeHomeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype app facade home o k response has a 3xx status code
func (o *CustomerPrototypeAppFacadeHomeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype app facade home o k response has a 4xx status code
func (o *CustomerPrototypeAppFacadeHomeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype app facade home o k response has a 5xx status code
func (o *CustomerPrototypeAppFacadeHomeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype app facade home o k response a status code equal to that given
func (o *CustomerPrototypeAppFacadeHomeOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype app facade home o k response
func (o *CustomerPrototypeAppFacadeHomeOK) Code() int {
	return 200
}

func (o *CustomerPrototypeAppFacadeHomeOK) Error() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/appFacade/home][%d] customerPrototypeAppFacadeHomeOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeAppFacadeHomeOK) String() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/appFacade/home][%d] customerPrototypeAppFacadeHomeOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeAppFacadeHomeOK) GetPayload() *models.AppFacadeHomeResponse {
	return o.Payload
}

func (o *CustomerPrototypeAppFacadeHomeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppFacadeHomeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypeAppFacadeHomeDefault creates a CustomerPrototypeAppFacadeHomeDefault with default headers values
func NewCustomerPrototypeAppFacadeHomeDefault(code int) *CustomerPrototypeAppFacadeHomeDefault {
	return &CustomerPrototypeAppFacadeHomeDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypeAppFacadeHomeDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypeAppFacadeHomeDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype app facade home default response has a 2xx status code
func (o *CustomerPrototypeAppFacadeHomeDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype app facade home default response has a 3xx status code
func (o *CustomerPrototypeAppFacadeHomeDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype app facade home default response has a 4xx status code
func (o *CustomerPrototypeAppFacadeHomeDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype app facade home default response has a 5xx status code
func (o *CustomerPrototypeAppFacadeHomeDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype app facade home default response a status code equal to that given
func (o *CustomerPrototypeAppFacadeHomeDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype app facade home default response
func (o *CustomerPrototypeAppFacadeHomeDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypeAppFacadeHomeDefault) Error() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/appFacade/home][%d] Customer.prototype.appFacadeHome default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeAppFacadeHomeDefault) String() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/appFacade/home][%d] Customer.prototype.appFacadeHome default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeAppFacadeHomeDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypeAppFacadeHomeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
