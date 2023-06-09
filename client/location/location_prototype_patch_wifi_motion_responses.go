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

// LocationPrototypePatchWifiMotionReader is a Reader for the LocationPrototypePatchWifiMotion structure.
type LocationPrototypePatchWifiMotionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LocationPrototypePatchWifiMotionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLocationPrototypePatchWifiMotionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLocationPrototypePatchWifiMotionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLocationPrototypePatchWifiMotionOK creates a LocationPrototypePatchWifiMotionOK with default headers values
func NewLocationPrototypePatchWifiMotionOK() *LocationPrototypePatchWifiMotionOK {
	return &LocationPrototypePatchWifiMotionOK{}
}

/*
LocationPrototypePatchWifiMotionOK describes a response with status code 200, with default header values.

Request was successful
*/
type LocationPrototypePatchWifiMotionOK struct {
	Payload *models.WifiMotion
}

// IsSuccess returns true when this location prototype patch wifi motion o k response has a 2xx status code
func (o *LocationPrototypePatchWifiMotionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this location prototype patch wifi motion o k response has a 3xx status code
func (o *LocationPrototypePatchWifiMotionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this location prototype patch wifi motion o k response has a 4xx status code
func (o *LocationPrototypePatchWifiMotionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this location prototype patch wifi motion o k response has a 5xx status code
func (o *LocationPrototypePatchWifiMotionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this location prototype patch wifi motion o k response a status code equal to that given
func (o *LocationPrototypePatchWifiMotionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the location prototype patch wifi motion o k response
func (o *LocationPrototypePatchWifiMotionOK) Code() int {
	return 200
}

func (o *LocationPrototypePatchWifiMotionOK) Error() string {
	return fmt.Sprintf("[PATCH /Locations/{id}/wifiMotion][%d] locationPrototypePatchWifiMotionOK  %+v", 200, o.Payload)
}

func (o *LocationPrototypePatchWifiMotionOK) String() string {
	return fmt.Sprintf("[PATCH /Locations/{id}/wifiMotion][%d] locationPrototypePatchWifiMotionOK  %+v", 200, o.Payload)
}

func (o *LocationPrototypePatchWifiMotionOK) GetPayload() *models.WifiMotion {
	return o.Payload
}

func (o *LocationPrototypePatchWifiMotionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WifiMotion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLocationPrototypePatchWifiMotionDefault creates a LocationPrototypePatchWifiMotionDefault with default headers values
func NewLocationPrototypePatchWifiMotionDefault(code int) *LocationPrototypePatchWifiMotionDefault {
	return &LocationPrototypePatchWifiMotionDefault{
		_statusCode: code,
	}
}

/*
LocationPrototypePatchWifiMotionDefault describes a response with status code -1, with default header values.

unexpected error
*/
type LocationPrototypePatchWifiMotionDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this location prototype patch wifi motion default response has a 2xx status code
func (o *LocationPrototypePatchWifiMotionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this location prototype patch wifi motion default response has a 3xx status code
func (o *LocationPrototypePatchWifiMotionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this location prototype patch wifi motion default response has a 4xx status code
func (o *LocationPrototypePatchWifiMotionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this location prototype patch wifi motion default response has a 5xx status code
func (o *LocationPrototypePatchWifiMotionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this location prototype patch wifi motion default response a status code equal to that given
func (o *LocationPrototypePatchWifiMotionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the location prototype patch wifi motion default response
func (o *LocationPrototypePatchWifiMotionDefault) Code() int {
	return o._statusCode
}

func (o *LocationPrototypePatchWifiMotionDefault) Error() string {
	return fmt.Sprintf("[PATCH /Locations/{id}/wifiMotion][%d] Location.prototype.patchWifiMotion default  %+v", o._statusCode, o.Payload)
}

func (o *LocationPrototypePatchWifiMotionDefault) String() string {
	return fmt.Sprintf("[PATCH /Locations/{id}/wifiMotion][%d] Location.prototype.patchWifiMotion default  %+v", o._statusCode, o.Payload)
}

func (o *LocationPrototypePatchWifiMotionDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *LocationPrototypePatchWifiMotionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
