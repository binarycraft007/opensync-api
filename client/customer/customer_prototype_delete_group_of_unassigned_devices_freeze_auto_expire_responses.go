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

// CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireReader is a Reader for the CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpire structure.
type CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewCustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireNoContent creates a CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireNoContent with default headers values
func NewCustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireNoContent() *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireNoContent {
	return &CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireNoContent{}
}

/*
CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireNoContent describes a response with status code 204, with default header values.

Request was successful
*/
type CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireNoContent struct {
}

// IsSuccess returns true when this customer prototype delete group of unassigned devices freeze auto expire no content response has a 2xx status code
func (o *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype delete group of unassigned devices freeze auto expire no content response has a 3xx status code
func (o *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype delete group of unassigned devices freeze auto expire no content response has a 4xx status code
func (o *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype delete group of unassigned devices freeze auto expire no content response has a 5xx status code
func (o *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype delete group of unassigned devices freeze auto expire no content response a status code equal to that given
func (o *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the customer prototype delete group of unassigned devices freeze auto expire no content response
func (o *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireNoContent) Code() int {
	return 204
}

func (o *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireNoContent) Error() string {
	return fmt.Sprintf("[DELETE /Customers/{id}/locations/{locationId}/groupOfUnassignedDevices/freeze/autoExpire][%d] customerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireNoContent ", 204)
}

func (o *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireNoContent) String() string {
	return fmt.Sprintf("[DELETE /Customers/{id}/locations/{locationId}/groupOfUnassignedDevices/freeze/autoExpire][%d] customerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireNoContent ", 204)
}

func (o *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireDefault creates a CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireDefault with default headers values
func NewCustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireDefault(code int) *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireDefault {
	return &CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype delete group of unassigned devices freeze auto expire default response has a 2xx status code
func (o *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype delete group of unassigned devices freeze auto expire default response has a 3xx status code
func (o *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype delete group of unassigned devices freeze auto expire default response has a 4xx status code
func (o *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype delete group of unassigned devices freeze auto expire default response has a 5xx status code
func (o *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype delete group of unassigned devices freeze auto expire default response a status code equal to that given
func (o *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype delete group of unassigned devices freeze auto expire default response
func (o *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireDefault) Error() string {
	return fmt.Sprintf("[DELETE /Customers/{id}/locations/{locationId}/groupOfUnassignedDevices/freeze/autoExpire][%d] Customer.prototype.deleteGroupOfUnassignedDevicesFreezeAutoExpire default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireDefault) String() string {
	return fmt.Sprintf("[DELETE /Customers/{id}/locations/{locationId}/groupOfUnassignedDevices/freeze/autoExpire][%d] Customer.prototype.deleteGroupOfUnassignedDevicesFreezeAutoExpire default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpireDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
