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

// CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDReader is a Reader for the CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateID structure.
type CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDOK creates a CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDOK with default headers values
func NewCustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDOK() *CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDOK {
	return &CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDOK{}
}

/*
CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDOK struct {
	Payload models.XAny
}

// IsSuccess returns true when this customer prototype post group of unassigned devices freeze template Id o k response has a 2xx status code
func (o *CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype post group of unassigned devices freeze template Id o k response has a 3xx status code
func (o *CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype post group of unassigned devices freeze template Id o k response has a 4xx status code
func (o *CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype post group of unassigned devices freeze template Id o k response has a 5xx status code
func (o *CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype post group of unassigned devices freeze template Id o k response a status code equal to that given
func (o *CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype post group of unassigned devices freeze template Id o k response
func (o *CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDOK) Code() int {
	return 200
}

func (o *CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDOK) Error() string {
	return fmt.Sprintf("[POST /Customers/{id}/locations/{locationId}/groupOfUnassignedDevices/freeze/{freezeTemplateId}][%d] customerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIdOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDOK) String() string {
	return fmt.Sprintf("[POST /Customers/{id}/locations/{locationId}/groupOfUnassignedDevices/freeze/{freezeTemplateId}][%d] customerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIdOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDOK) GetPayload() models.XAny {
	return o.Payload
}

func (o *CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDDefault creates a CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDDefault with default headers values
func NewCustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDDefault(code int) *CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDDefault {
	return &CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype post group of unassigned devices freeze template Id default response has a 2xx status code
func (o *CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype post group of unassigned devices freeze template Id default response has a 3xx status code
func (o *CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype post group of unassigned devices freeze template Id default response has a 4xx status code
func (o *CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype post group of unassigned devices freeze template Id default response has a 5xx status code
func (o *CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype post group of unassigned devices freeze template Id default response a status code equal to that given
func (o *CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype post group of unassigned devices freeze template Id default response
func (o *CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDDefault) Error() string {
	return fmt.Sprintf("[POST /Customers/{id}/locations/{locationId}/groupOfUnassignedDevices/freeze/{freezeTemplateId}][%d] Customer.prototype.postGroupOfUnassignedDevicesFreezeTemplateId default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDDefault) String() string {
	return fmt.Sprintf("[POST /Customers/{id}/locations/{locationId}/groupOfUnassignedDevices/freeze/{freezeTemplateId}][%d] Customer.prototype.postGroupOfUnassignedDevicesFreezeTemplateId default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
