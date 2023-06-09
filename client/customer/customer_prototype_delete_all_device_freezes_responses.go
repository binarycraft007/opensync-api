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

// CustomerPrototypeDeleteAllDeviceFreezesReader is a Reader for the CustomerPrototypeDeleteAllDeviceFreezes structure.
type CustomerPrototypeDeleteAllDeviceFreezesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypeDeleteAllDeviceFreezesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewCustomerPrototypeDeleteAllDeviceFreezesNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypeDeleteAllDeviceFreezesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypeDeleteAllDeviceFreezesNoContent creates a CustomerPrototypeDeleteAllDeviceFreezesNoContent with default headers values
func NewCustomerPrototypeDeleteAllDeviceFreezesNoContent() *CustomerPrototypeDeleteAllDeviceFreezesNoContent {
	return &CustomerPrototypeDeleteAllDeviceFreezesNoContent{}
}

/*
CustomerPrototypeDeleteAllDeviceFreezesNoContent describes a response with status code 204, with default header values.

Request was successful
*/
type CustomerPrototypeDeleteAllDeviceFreezesNoContent struct {
}

// IsSuccess returns true when this customer prototype delete all device freezes no content response has a 2xx status code
func (o *CustomerPrototypeDeleteAllDeviceFreezesNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype delete all device freezes no content response has a 3xx status code
func (o *CustomerPrototypeDeleteAllDeviceFreezesNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype delete all device freezes no content response has a 4xx status code
func (o *CustomerPrototypeDeleteAllDeviceFreezesNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype delete all device freezes no content response has a 5xx status code
func (o *CustomerPrototypeDeleteAllDeviceFreezesNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype delete all device freezes no content response a status code equal to that given
func (o *CustomerPrototypeDeleteAllDeviceFreezesNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the customer prototype delete all device freezes no content response
func (o *CustomerPrototypeDeleteAllDeviceFreezesNoContent) Code() int {
	return 204
}

func (o *CustomerPrototypeDeleteAllDeviceFreezesNoContent) Error() string {
	return fmt.Sprintf("[DELETE /Customers/{id}/locations/{locationId}/devices/{mac}/freezes][%d] customerPrototypeDeleteAllDeviceFreezesNoContent ", 204)
}

func (o *CustomerPrototypeDeleteAllDeviceFreezesNoContent) String() string {
	return fmt.Sprintf("[DELETE /Customers/{id}/locations/{locationId}/devices/{mac}/freezes][%d] customerPrototypeDeleteAllDeviceFreezesNoContent ", 204)
}

func (o *CustomerPrototypeDeleteAllDeviceFreezesNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCustomerPrototypeDeleteAllDeviceFreezesDefault creates a CustomerPrototypeDeleteAllDeviceFreezesDefault with default headers values
func NewCustomerPrototypeDeleteAllDeviceFreezesDefault(code int) *CustomerPrototypeDeleteAllDeviceFreezesDefault {
	return &CustomerPrototypeDeleteAllDeviceFreezesDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypeDeleteAllDeviceFreezesDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypeDeleteAllDeviceFreezesDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype delete all device freezes default response has a 2xx status code
func (o *CustomerPrototypeDeleteAllDeviceFreezesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype delete all device freezes default response has a 3xx status code
func (o *CustomerPrototypeDeleteAllDeviceFreezesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype delete all device freezes default response has a 4xx status code
func (o *CustomerPrototypeDeleteAllDeviceFreezesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype delete all device freezes default response has a 5xx status code
func (o *CustomerPrototypeDeleteAllDeviceFreezesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype delete all device freezes default response a status code equal to that given
func (o *CustomerPrototypeDeleteAllDeviceFreezesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype delete all device freezes default response
func (o *CustomerPrototypeDeleteAllDeviceFreezesDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypeDeleteAllDeviceFreezesDefault) Error() string {
	return fmt.Sprintf("[DELETE /Customers/{id}/locations/{locationId}/devices/{mac}/freezes][%d] Customer.prototype.deleteAllDeviceFreezes default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeDeleteAllDeviceFreezesDefault) String() string {
	return fmt.Sprintf("[DELETE /Customers/{id}/locations/{locationId}/devices/{mac}/freezes][%d] Customer.prototype.deleteAllDeviceFreezes default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeDeleteAllDeviceFreezesDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypeDeleteAllDeviceFreezesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
