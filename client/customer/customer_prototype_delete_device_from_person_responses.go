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

// CustomerPrototypeDeleteDeviceFromPersonReader is a Reader for the CustomerPrototypeDeleteDeviceFromPerson structure.
type CustomerPrototypeDeleteDeviceFromPersonReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypeDeleteDeviceFromPersonReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewCustomerPrototypeDeleteDeviceFromPersonNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypeDeleteDeviceFromPersonDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypeDeleteDeviceFromPersonNoContent creates a CustomerPrototypeDeleteDeviceFromPersonNoContent with default headers values
func NewCustomerPrototypeDeleteDeviceFromPersonNoContent() *CustomerPrototypeDeleteDeviceFromPersonNoContent {
	return &CustomerPrototypeDeleteDeviceFromPersonNoContent{}
}

/*
CustomerPrototypeDeleteDeviceFromPersonNoContent describes a response with status code 204, with default header values.

Request was successful
*/
type CustomerPrototypeDeleteDeviceFromPersonNoContent struct {
}

// IsSuccess returns true when this customer prototype delete device from person no content response has a 2xx status code
func (o *CustomerPrototypeDeleteDeviceFromPersonNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype delete device from person no content response has a 3xx status code
func (o *CustomerPrototypeDeleteDeviceFromPersonNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype delete device from person no content response has a 4xx status code
func (o *CustomerPrototypeDeleteDeviceFromPersonNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype delete device from person no content response has a 5xx status code
func (o *CustomerPrototypeDeleteDeviceFromPersonNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype delete device from person no content response a status code equal to that given
func (o *CustomerPrototypeDeleteDeviceFromPersonNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the customer prototype delete device from person no content response
func (o *CustomerPrototypeDeleteDeviceFromPersonNoContent) Code() int {
	return 204
}

func (o *CustomerPrototypeDeleteDeviceFromPersonNoContent) Error() string {
	return fmt.Sprintf("[DELETE /Customers/{id}/locations/{locationId}/persons/{personId}/devices/{mac}][%d] customerPrototypeDeleteDeviceFromPersonNoContent ", 204)
}

func (o *CustomerPrototypeDeleteDeviceFromPersonNoContent) String() string {
	return fmt.Sprintf("[DELETE /Customers/{id}/locations/{locationId}/persons/{personId}/devices/{mac}][%d] customerPrototypeDeleteDeviceFromPersonNoContent ", 204)
}

func (o *CustomerPrototypeDeleteDeviceFromPersonNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCustomerPrototypeDeleteDeviceFromPersonDefault creates a CustomerPrototypeDeleteDeviceFromPersonDefault with default headers values
func NewCustomerPrototypeDeleteDeviceFromPersonDefault(code int) *CustomerPrototypeDeleteDeviceFromPersonDefault {
	return &CustomerPrototypeDeleteDeviceFromPersonDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypeDeleteDeviceFromPersonDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypeDeleteDeviceFromPersonDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype delete device from person default response has a 2xx status code
func (o *CustomerPrototypeDeleteDeviceFromPersonDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype delete device from person default response has a 3xx status code
func (o *CustomerPrototypeDeleteDeviceFromPersonDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype delete device from person default response has a 4xx status code
func (o *CustomerPrototypeDeleteDeviceFromPersonDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype delete device from person default response has a 5xx status code
func (o *CustomerPrototypeDeleteDeviceFromPersonDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype delete device from person default response a status code equal to that given
func (o *CustomerPrototypeDeleteDeviceFromPersonDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype delete device from person default response
func (o *CustomerPrototypeDeleteDeviceFromPersonDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypeDeleteDeviceFromPersonDefault) Error() string {
	return fmt.Sprintf("[DELETE /Customers/{id}/locations/{locationId}/persons/{personId}/devices/{mac}][%d] Customer.prototype.deleteDeviceFromPerson default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeDeleteDeviceFromPersonDefault) String() string {
	return fmt.Sprintf("[DELETE /Customers/{id}/locations/{locationId}/persons/{personId}/devices/{mac}][%d] Customer.prototype.deleteDeviceFromPerson default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeDeleteDeviceFromPersonDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypeDeleteDeviceFromPersonDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
