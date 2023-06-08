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

// LocationPrototypePatchHomeSecurityReader is a Reader for the LocationPrototypePatchHomeSecurity structure.
type LocationPrototypePatchHomeSecurityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LocationPrototypePatchHomeSecurityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLocationPrototypePatchHomeSecurityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLocationPrototypePatchHomeSecurityDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLocationPrototypePatchHomeSecurityOK creates a LocationPrototypePatchHomeSecurityOK with default headers values
func NewLocationPrototypePatchHomeSecurityOK() *LocationPrototypePatchHomeSecurityOK {
	return &LocationPrototypePatchHomeSecurityOK{}
}

/*
LocationPrototypePatchHomeSecurityOK describes a response with status code 200, with default header values.

Request was successful
*/
type LocationPrototypePatchHomeSecurityOK struct {
	Payload *models.HomeSecurity
}

// IsSuccess returns true when this location prototype patch home security o k response has a 2xx status code
func (o *LocationPrototypePatchHomeSecurityOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this location prototype patch home security o k response has a 3xx status code
func (o *LocationPrototypePatchHomeSecurityOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this location prototype patch home security o k response has a 4xx status code
func (o *LocationPrototypePatchHomeSecurityOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this location prototype patch home security o k response has a 5xx status code
func (o *LocationPrototypePatchHomeSecurityOK) IsServerError() bool {
	return false
}

// IsCode returns true when this location prototype patch home security o k response a status code equal to that given
func (o *LocationPrototypePatchHomeSecurityOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the location prototype patch home security o k response
func (o *LocationPrototypePatchHomeSecurityOK) Code() int {
	return 200
}

func (o *LocationPrototypePatchHomeSecurityOK) Error() string {
	return fmt.Sprintf("[PATCH /Locations/{id}/homeSecurity][%d] locationPrototypePatchHomeSecurityOK  %+v", 200, o.Payload)
}

func (o *LocationPrototypePatchHomeSecurityOK) String() string {
	return fmt.Sprintf("[PATCH /Locations/{id}/homeSecurity][%d] locationPrototypePatchHomeSecurityOK  %+v", 200, o.Payload)
}

func (o *LocationPrototypePatchHomeSecurityOK) GetPayload() *models.HomeSecurity {
	return o.Payload
}

func (o *LocationPrototypePatchHomeSecurityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HomeSecurity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLocationPrototypePatchHomeSecurityDefault creates a LocationPrototypePatchHomeSecurityDefault with default headers values
func NewLocationPrototypePatchHomeSecurityDefault(code int) *LocationPrototypePatchHomeSecurityDefault {
	return &LocationPrototypePatchHomeSecurityDefault{
		_statusCode: code,
	}
}

/*
LocationPrototypePatchHomeSecurityDefault describes a response with status code -1, with default header values.

unexpected error
*/
type LocationPrototypePatchHomeSecurityDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this location prototype patch home security default response has a 2xx status code
func (o *LocationPrototypePatchHomeSecurityDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this location prototype patch home security default response has a 3xx status code
func (o *LocationPrototypePatchHomeSecurityDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this location prototype patch home security default response has a 4xx status code
func (o *LocationPrototypePatchHomeSecurityDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this location prototype patch home security default response has a 5xx status code
func (o *LocationPrototypePatchHomeSecurityDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this location prototype patch home security default response a status code equal to that given
func (o *LocationPrototypePatchHomeSecurityDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the location prototype patch home security default response
func (o *LocationPrototypePatchHomeSecurityDefault) Code() int {
	return o._statusCode
}

func (o *LocationPrototypePatchHomeSecurityDefault) Error() string {
	return fmt.Sprintf("[PATCH /Locations/{id}/homeSecurity][%d] Location.prototype.patchHomeSecurity default  %+v", o._statusCode, o.Payload)
}

func (o *LocationPrototypePatchHomeSecurityDefault) String() string {
	return fmt.Sprintf("[PATCH /Locations/{id}/homeSecurity][%d] Location.prototype.patchHomeSecurity default  %+v", o._statusCode, o.Payload)
}

func (o *LocationPrototypePatchHomeSecurityDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *LocationPrototypePatchHomeSecurityDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
