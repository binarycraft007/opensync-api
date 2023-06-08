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

// CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeForeverReader is a Reader for the CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeForever structure.
type CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeForeverReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeForeverReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewCustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeForeverNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeForeverDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeForeverNoContent creates a CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeForeverNoContent with default headers values
func NewCustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeForeverNoContent() *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeForeverNoContent {
	return &CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeForeverNoContent{}
}

/*
CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeForeverNoContent describes a response with status code 204, with default header values.

Request was successful
*/
type CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeForeverNoContent struct {
}

// IsSuccess returns true when this customer prototype delete group of unassigned devices freeze forever no content response has a 2xx status code
func (o *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeForeverNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype delete group of unassigned devices freeze forever no content response has a 3xx status code
func (o *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeForeverNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype delete group of unassigned devices freeze forever no content response has a 4xx status code
func (o *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeForeverNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype delete group of unassigned devices freeze forever no content response has a 5xx status code
func (o *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeForeverNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype delete group of unassigned devices freeze forever no content response a status code equal to that given
func (o *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeForeverNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the customer prototype delete group of unassigned devices freeze forever no content response
func (o *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeForeverNoContent) Code() int {
	return 204
}

func (o *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeForeverNoContent) Error() string {
	return fmt.Sprintf("[DELETE /Customers/{id}/locations/{locationId}/groupOfUnassignedDevices/freeze/forever][%d] customerPrototypeDeleteGroupOfUnassignedDevicesFreezeForeverNoContent ", 204)
}

func (o *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeForeverNoContent) String() string {
	return fmt.Sprintf("[DELETE /Customers/{id}/locations/{locationId}/groupOfUnassignedDevices/freeze/forever][%d] customerPrototypeDeleteGroupOfUnassignedDevicesFreezeForeverNoContent ", 204)
}

func (o *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeForeverNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeForeverDefault creates a CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeForeverDefault with default headers values
func NewCustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeForeverDefault(code int) *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeForeverDefault {
	return &CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeForeverDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeForeverDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeForeverDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype delete group of unassigned devices freeze forever default response has a 2xx status code
func (o *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeForeverDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype delete group of unassigned devices freeze forever default response has a 3xx status code
func (o *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeForeverDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype delete group of unassigned devices freeze forever default response has a 4xx status code
func (o *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeForeverDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype delete group of unassigned devices freeze forever default response has a 5xx status code
func (o *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeForeverDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype delete group of unassigned devices freeze forever default response a status code equal to that given
func (o *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeForeverDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype delete group of unassigned devices freeze forever default response
func (o *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeForeverDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeForeverDefault) Error() string {
	return fmt.Sprintf("[DELETE /Customers/{id}/locations/{locationId}/groupOfUnassignedDevices/freeze/forever][%d] Customer.prototype.deleteGroupOfUnassignedDevicesFreezeForever default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeForeverDefault) String() string {
	return fmt.Sprintf("[DELETE /Customers/{id}/locations/{locationId}/groupOfUnassignedDevices/freeze/forever][%d] Customer.prototype.deleteGroupOfUnassignedDevicesFreezeForever default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeForeverDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeForeverDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}