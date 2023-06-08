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

// LocationPrototypeGetDeviceSoundingStateReader is a Reader for the LocationPrototypeGetDeviceSoundingState structure.
type LocationPrototypeGetDeviceSoundingStateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LocationPrototypeGetDeviceSoundingStateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLocationPrototypeGetDeviceSoundingStateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLocationPrototypeGetDeviceSoundingStateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLocationPrototypeGetDeviceSoundingStateOK creates a LocationPrototypeGetDeviceSoundingStateOK with default headers values
func NewLocationPrototypeGetDeviceSoundingStateOK() *LocationPrototypeGetDeviceSoundingStateOK {
	return &LocationPrototypeGetDeviceSoundingStateOK{}
}

/*
LocationPrototypeGetDeviceSoundingStateOK describes a response with status code 200, with default header values.

Request was successful
*/
type LocationPrototypeGetDeviceSoundingStateOK struct {
	Payload interface{}
}

// IsSuccess returns true when this location prototype get device sounding state o k response has a 2xx status code
func (o *LocationPrototypeGetDeviceSoundingStateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this location prototype get device sounding state o k response has a 3xx status code
func (o *LocationPrototypeGetDeviceSoundingStateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this location prototype get device sounding state o k response has a 4xx status code
func (o *LocationPrototypeGetDeviceSoundingStateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this location prototype get device sounding state o k response has a 5xx status code
func (o *LocationPrototypeGetDeviceSoundingStateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this location prototype get device sounding state o k response a status code equal to that given
func (o *LocationPrototypeGetDeviceSoundingStateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the location prototype get device sounding state o k response
func (o *LocationPrototypeGetDeviceSoundingStateOK) Code() int {
	return 200
}

func (o *LocationPrototypeGetDeviceSoundingStateOK) Error() string {
	return fmt.Sprintf("[GET /Locations/{id}/homeSecurity/devices/sounding][%d] locationPrototypeGetDeviceSoundingStateOK  %+v", 200, o.Payload)
}

func (o *LocationPrototypeGetDeviceSoundingStateOK) String() string {
	return fmt.Sprintf("[GET /Locations/{id}/homeSecurity/devices/sounding][%d] locationPrototypeGetDeviceSoundingStateOK  %+v", 200, o.Payload)
}

func (o *LocationPrototypeGetDeviceSoundingStateOK) GetPayload() interface{} {
	return o.Payload
}

func (o *LocationPrototypeGetDeviceSoundingStateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLocationPrototypeGetDeviceSoundingStateDefault creates a LocationPrototypeGetDeviceSoundingStateDefault with default headers values
func NewLocationPrototypeGetDeviceSoundingStateDefault(code int) *LocationPrototypeGetDeviceSoundingStateDefault {
	return &LocationPrototypeGetDeviceSoundingStateDefault{
		_statusCode: code,
	}
}

/*
LocationPrototypeGetDeviceSoundingStateDefault describes a response with status code -1, with default header values.

unexpected error
*/
type LocationPrototypeGetDeviceSoundingStateDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this location prototype get device sounding state default response has a 2xx status code
func (o *LocationPrototypeGetDeviceSoundingStateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this location prototype get device sounding state default response has a 3xx status code
func (o *LocationPrototypeGetDeviceSoundingStateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this location prototype get device sounding state default response has a 4xx status code
func (o *LocationPrototypeGetDeviceSoundingStateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this location prototype get device sounding state default response has a 5xx status code
func (o *LocationPrototypeGetDeviceSoundingStateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this location prototype get device sounding state default response a status code equal to that given
func (o *LocationPrototypeGetDeviceSoundingStateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the location prototype get device sounding state default response
func (o *LocationPrototypeGetDeviceSoundingStateDefault) Code() int {
	return o._statusCode
}

func (o *LocationPrototypeGetDeviceSoundingStateDefault) Error() string {
	return fmt.Sprintf("[GET /Locations/{id}/homeSecurity/devices/sounding][%d] Location.prototype.getDeviceSoundingState default  %+v", o._statusCode, o.Payload)
}

func (o *LocationPrototypeGetDeviceSoundingStateDefault) String() string {
	return fmt.Sprintf("[GET /Locations/{id}/homeSecurity/devices/sounding][%d] Location.prototype.getDeviceSoundingState default  %+v", o._statusCode, o.Payload)
}

func (o *LocationPrototypeGetDeviceSoundingStateDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *LocationPrototypeGetDeviceSoundingStateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
