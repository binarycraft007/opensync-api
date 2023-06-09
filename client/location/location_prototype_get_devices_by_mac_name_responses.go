// Code generated by go-swagger; DO NOT EDIT.

package location

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/binarycraft007/opensync-api/models"
)

// LocationPrototypeGetDevicesByMacNameReader is a Reader for the LocationPrototypeGetDevicesByMacName structure.
type LocationPrototypeGetDevicesByMacNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LocationPrototypeGetDevicesByMacNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLocationPrototypeGetDevicesByMacNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLocationPrototypeGetDevicesByMacNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLocationPrototypeGetDevicesByMacNameOK creates a LocationPrototypeGetDevicesByMacNameOK with default headers values
func NewLocationPrototypeGetDevicesByMacNameOK() *LocationPrototypeGetDevicesByMacNameOK {
	return &LocationPrototypeGetDevicesByMacNameOK{}
}

/*
LocationPrototypeGetDevicesByMacNameOK describes a response with status code 200, with default header values.

Request was successful
*/
type LocationPrototypeGetDevicesByMacNameOK struct {
	Payload interface{}
}

// IsSuccess returns true when this location prototype get devices by mac name o k response has a 2xx status code
func (o *LocationPrototypeGetDevicesByMacNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this location prototype get devices by mac name o k response has a 3xx status code
func (o *LocationPrototypeGetDevicesByMacNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this location prototype get devices by mac name o k response has a 4xx status code
func (o *LocationPrototypeGetDevicesByMacNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this location prototype get devices by mac name o k response has a 5xx status code
func (o *LocationPrototypeGetDevicesByMacNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this location prototype get devices by mac name o k response a status code equal to that given
func (o *LocationPrototypeGetDevicesByMacNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the location prototype get devices by mac name o k response
func (o *LocationPrototypeGetDevicesByMacNameOK) Code() int {
	return 200
}

func (o *LocationPrototypeGetDevicesByMacNameOK) Error() string {
	return fmt.Sprintf("[GET /Locations/{id}/devices/{mac}][%d] locationPrototypeGetDevicesByMacNameOK  %+v", 200, o.Payload)
}

func (o *LocationPrototypeGetDevicesByMacNameOK) String() string {
	return fmt.Sprintf("[GET /Locations/{id}/devices/{mac}][%d] locationPrototypeGetDevicesByMacNameOK  %+v", 200, o.Payload)
}

func (o *LocationPrototypeGetDevicesByMacNameOK) GetPayload() interface{} {
	return o.Payload
}

func (o *LocationPrototypeGetDevicesByMacNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLocationPrototypeGetDevicesByMacNameDefault creates a LocationPrototypeGetDevicesByMacNameDefault with default headers values
func NewLocationPrototypeGetDevicesByMacNameDefault(code int) *LocationPrototypeGetDevicesByMacNameDefault {
	return &LocationPrototypeGetDevicesByMacNameDefault{
		_statusCode: code,
	}
}

/*
LocationPrototypeGetDevicesByMacNameDefault describes a response with status code -1, with default header values.

unexpected error
*/
type LocationPrototypeGetDevicesByMacNameDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this location prototype get devices by mac name default response has a 2xx status code
func (o *LocationPrototypeGetDevicesByMacNameDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this location prototype get devices by mac name default response has a 3xx status code
func (o *LocationPrototypeGetDevicesByMacNameDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this location prototype get devices by mac name default response has a 4xx status code
func (o *LocationPrototypeGetDevicesByMacNameDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this location prototype get devices by mac name default response has a 5xx status code
func (o *LocationPrototypeGetDevicesByMacNameDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this location prototype get devices by mac name default response a status code equal to that given
func (o *LocationPrototypeGetDevicesByMacNameDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the location prototype get devices by mac name default response
func (o *LocationPrototypeGetDevicesByMacNameDefault) Code() int {
	return o._statusCode
}

func (o *LocationPrototypeGetDevicesByMacNameDefault) Error() string {
	return fmt.Sprintf("[GET /Locations/{id}/devices/{mac}][%d] Location.prototype.getDevicesByMacName default  %+v", o._statusCode, o.Payload)
}

func (o *LocationPrototypeGetDevicesByMacNameDefault) String() string {
	return fmt.Sprintf("[GET /Locations/{id}/devices/{mac}][%d] Location.prototype.getDevicesByMacName default  %+v", o._statusCode, o.Payload)
}

func (o *LocationPrototypeGetDevicesByMacNameDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *LocationPrototypeGetDevicesByMacNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
