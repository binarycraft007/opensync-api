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

// CustomerPrototypeGetFirmwareUpgradeStatusReader is a Reader for the CustomerPrototypeGetFirmwareUpgradeStatus structure.
type CustomerPrototypeGetFirmwareUpgradeStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypeGetFirmwareUpgradeStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypeGetFirmwareUpgradeStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypeGetFirmwareUpgradeStatusDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypeGetFirmwareUpgradeStatusOK creates a CustomerPrototypeGetFirmwareUpgradeStatusOK with default headers values
func NewCustomerPrototypeGetFirmwareUpgradeStatusOK() *CustomerPrototypeGetFirmwareUpgradeStatusOK {
	return &CustomerPrototypeGetFirmwareUpgradeStatusOK{}
}

/*
CustomerPrototypeGetFirmwareUpgradeStatusOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypeGetFirmwareUpgradeStatusOK struct {
	Payload []interface{}
}

// IsSuccess returns true when this customer prototype get firmware upgrade status o k response has a 2xx status code
func (o *CustomerPrototypeGetFirmwareUpgradeStatusOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype get firmware upgrade status o k response has a 3xx status code
func (o *CustomerPrototypeGetFirmwareUpgradeStatusOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype get firmware upgrade status o k response has a 4xx status code
func (o *CustomerPrototypeGetFirmwareUpgradeStatusOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype get firmware upgrade status o k response has a 5xx status code
func (o *CustomerPrototypeGetFirmwareUpgradeStatusOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype get firmware upgrade status o k response a status code equal to that given
func (o *CustomerPrototypeGetFirmwareUpgradeStatusOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype get firmware upgrade status o k response
func (o *CustomerPrototypeGetFirmwareUpgradeStatusOK) Code() int {
	return 200
}

func (o *CustomerPrototypeGetFirmwareUpgradeStatusOK) Error() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/firmware][%d] customerPrototypeGetFirmwareUpgradeStatusOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeGetFirmwareUpgradeStatusOK) String() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/firmware][%d] customerPrototypeGetFirmwareUpgradeStatusOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeGetFirmwareUpgradeStatusOK) GetPayload() []interface{} {
	return o.Payload
}

func (o *CustomerPrototypeGetFirmwareUpgradeStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypeGetFirmwareUpgradeStatusDefault creates a CustomerPrototypeGetFirmwareUpgradeStatusDefault with default headers values
func NewCustomerPrototypeGetFirmwareUpgradeStatusDefault(code int) *CustomerPrototypeGetFirmwareUpgradeStatusDefault {
	return &CustomerPrototypeGetFirmwareUpgradeStatusDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypeGetFirmwareUpgradeStatusDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypeGetFirmwareUpgradeStatusDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype get firmware upgrade status default response has a 2xx status code
func (o *CustomerPrototypeGetFirmwareUpgradeStatusDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype get firmware upgrade status default response has a 3xx status code
func (o *CustomerPrototypeGetFirmwareUpgradeStatusDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype get firmware upgrade status default response has a 4xx status code
func (o *CustomerPrototypeGetFirmwareUpgradeStatusDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype get firmware upgrade status default response has a 5xx status code
func (o *CustomerPrototypeGetFirmwareUpgradeStatusDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype get firmware upgrade status default response a status code equal to that given
func (o *CustomerPrototypeGetFirmwareUpgradeStatusDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype get firmware upgrade status default response
func (o *CustomerPrototypeGetFirmwareUpgradeStatusDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypeGetFirmwareUpgradeStatusDefault) Error() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/firmware][%d] Customer.prototype.getFirmwareUpgradeStatus default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeGetFirmwareUpgradeStatusDefault) String() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/firmware][%d] Customer.prototype.getFirmwareUpgradeStatus default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeGetFirmwareUpgradeStatusDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypeGetFirmwareUpgradeStatusDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
