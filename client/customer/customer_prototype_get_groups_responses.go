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

// CustomerPrototypeGetGroupsReader is a Reader for the CustomerPrototypeGetGroups structure.
type CustomerPrototypeGetGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypeGetGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypeGetGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypeGetGroupsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypeGetGroupsOK creates a CustomerPrototypeGetGroupsOK with default headers values
func NewCustomerPrototypeGetGroupsOK() *CustomerPrototypeGetGroupsOK {
	return &CustomerPrototypeGetGroupsOK{}
}

/*
CustomerPrototypeGetGroupsOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypeGetGroupsOK struct {
	Payload []models.XAny
}

// IsSuccess returns true when this customer prototype get groups o k response has a 2xx status code
func (o *CustomerPrototypeGetGroupsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype get groups o k response has a 3xx status code
func (o *CustomerPrototypeGetGroupsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype get groups o k response has a 4xx status code
func (o *CustomerPrototypeGetGroupsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype get groups o k response has a 5xx status code
func (o *CustomerPrototypeGetGroupsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype get groups o k response a status code equal to that given
func (o *CustomerPrototypeGetGroupsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype get groups o k response
func (o *CustomerPrototypeGetGroupsOK) Code() int {
	return 200
}

func (o *CustomerPrototypeGetGroupsOK) Error() string {
	return fmt.Sprintf("[GET /Customers/{id}/groups][%d] customerPrototypeGetGroupsOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeGetGroupsOK) String() string {
	return fmt.Sprintf("[GET /Customers/{id}/groups][%d] customerPrototypeGetGroupsOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeGetGroupsOK) GetPayload() []models.XAny {
	return o.Payload
}

func (o *CustomerPrototypeGetGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypeGetGroupsDefault creates a CustomerPrototypeGetGroupsDefault with default headers values
func NewCustomerPrototypeGetGroupsDefault(code int) *CustomerPrototypeGetGroupsDefault {
	return &CustomerPrototypeGetGroupsDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypeGetGroupsDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypeGetGroupsDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype get groups default response has a 2xx status code
func (o *CustomerPrototypeGetGroupsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype get groups default response has a 3xx status code
func (o *CustomerPrototypeGetGroupsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype get groups default response has a 4xx status code
func (o *CustomerPrototypeGetGroupsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype get groups default response has a 5xx status code
func (o *CustomerPrototypeGetGroupsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype get groups default response a status code equal to that given
func (o *CustomerPrototypeGetGroupsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype get groups default response
func (o *CustomerPrototypeGetGroupsDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypeGetGroupsDefault) Error() string {
	return fmt.Sprintf("[GET /Customers/{id}/groups][%d] Customer.prototype.getGroups default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeGetGroupsDefault) String() string {
	return fmt.Sprintf("[GET /Customers/{id}/groups][%d] Customer.prototype.getGroups default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeGetGroupsDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypeGetGroupsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}