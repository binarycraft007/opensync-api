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

// CustomerPrototypePostDeviceToAccessZoneReader is a Reader for the CustomerPrototypePostDeviceToAccessZone structure.
type CustomerPrototypePostDeviceToAccessZoneReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypePostDeviceToAccessZoneReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypePostDeviceToAccessZoneOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypePostDeviceToAccessZoneDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypePostDeviceToAccessZoneOK creates a CustomerPrototypePostDeviceToAccessZoneOK with default headers values
func NewCustomerPrototypePostDeviceToAccessZoneOK() *CustomerPrototypePostDeviceToAccessZoneOK {
	return &CustomerPrototypePostDeviceToAccessZoneOK{}
}

/*
CustomerPrototypePostDeviceToAccessZoneOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypePostDeviceToAccessZoneOK struct {
	Payload models.XAny
}

// IsSuccess returns true when this customer prototype post device to access zone o k response has a 2xx status code
func (o *CustomerPrototypePostDeviceToAccessZoneOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype post device to access zone o k response has a 3xx status code
func (o *CustomerPrototypePostDeviceToAccessZoneOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype post device to access zone o k response has a 4xx status code
func (o *CustomerPrototypePostDeviceToAccessZoneOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype post device to access zone o k response has a 5xx status code
func (o *CustomerPrototypePostDeviceToAccessZoneOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype post device to access zone o k response a status code equal to that given
func (o *CustomerPrototypePostDeviceToAccessZoneOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype post device to access zone o k response
func (o *CustomerPrototypePostDeviceToAccessZoneOK) Code() int {
	return 200
}

func (o *CustomerPrototypePostDeviceToAccessZoneOK) Error() string {
	return fmt.Sprintf("[POST /Customers/{id}/locations/{locationId}/wifiNetwork/accessZones/{zoneId}/accessibleDevices/{mac}][%d] customerPrototypePostDeviceToAccessZoneOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypePostDeviceToAccessZoneOK) String() string {
	return fmt.Sprintf("[POST /Customers/{id}/locations/{locationId}/wifiNetwork/accessZones/{zoneId}/accessibleDevices/{mac}][%d] customerPrototypePostDeviceToAccessZoneOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypePostDeviceToAccessZoneOK) GetPayload() models.XAny {
	return o.Payload
}

func (o *CustomerPrototypePostDeviceToAccessZoneOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypePostDeviceToAccessZoneDefault creates a CustomerPrototypePostDeviceToAccessZoneDefault with default headers values
func NewCustomerPrototypePostDeviceToAccessZoneDefault(code int) *CustomerPrototypePostDeviceToAccessZoneDefault {
	return &CustomerPrototypePostDeviceToAccessZoneDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypePostDeviceToAccessZoneDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypePostDeviceToAccessZoneDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype post device to access zone default response has a 2xx status code
func (o *CustomerPrototypePostDeviceToAccessZoneDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype post device to access zone default response has a 3xx status code
func (o *CustomerPrototypePostDeviceToAccessZoneDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype post device to access zone default response has a 4xx status code
func (o *CustomerPrototypePostDeviceToAccessZoneDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype post device to access zone default response has a 5xx status code
func (o *CustomerPrototypePostDeviceToAccessZoneDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype post device to access zone default response a status code equal to that given
func (o *CustomerPrototypePostDeviceToAccessZoneDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype post device to access zone default response
func (o *CustomerPrototypePostDeviceToAccessZoneDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypePostDeviceToAccessZoneDefault) Error() string {
	return fmt.Sprintf("[POST /Customers/{id}/locations/{locationId}/wifiNetwork/accessZones/{zoneId}/accessibleDevices/{mac}][%d] Customer.prototype.postDeviceToAccessZone default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePostDeviceToAccessZoneDefault) String() string {
	return fmt.Sprintf("[POST /Customers/{id}/locations/{locationId}/wifiNetwork/accessZones/{zoneId}/accessibleDevices/{mac}][%d] Customer.prototype.postDeviceToAccessZone default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePostDeviceToAccessZoneDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypePostDeviceToAccessZoneDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
