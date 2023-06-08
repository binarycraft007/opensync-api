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

// LocationFindReader is a Reader for the LocationFind structure.
type LocationFindReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LocationFindReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLocationFindOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLocationFindDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLocationFindOK creates a LocationFindOK with default headers values
func NewLocationFindOK() *LocationFindOK {
	return &LocationFindOK{}
}

/*
LocationFindOK describes a response with status code 200, with default header values.

Request was successful
*/
type LocationFindOK struct {
	Payload []*models.Location
}

// IsSuccess returns true when this location find o k response has a 2xx status code
func (o *LocationFindOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this location find o k response has a 3xx status code
func (o *LocationFindOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this location find o k response has a 4xx status code
func (o *LocationFindOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this location find o k response has a 5xx status code
func (o *LocationFindOK) IsServerError() bool {
	return false
}

// IsCode returns true when this location find o k response a status code equal to that given
func (o *LocationFindOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the location find o k response
func (o *LocationFindOK) Code() int {
	return 200
}

func (o *LocationFindOK) Error() string {
	return fmt.Sprintf("[GET /Locations][%d] locationFindOK  %+v", 200, o.Payload)
}

func (o *LocationFindOK) String() string {
	return fmt.Sprintf("[GET /Locations][%d] locationFindOK  %+v", 200, o.Payload)
}

func (o *LocationFindOK) GetPayload() []*models.Location {
	return o.Payload
}

func (o *LocationFindOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLocationFindDefault creates a LocationFindDefault with default headers values
func NewLocationFindDefault(code int) *LocationFindDefault {
	return &LocationFindDefault{
		_statusCode: code,
	}
}

/*
LocationFindDefault describes a response with status code -1, with default header values.

unexpected error
*/
type LocationFindDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this location find default response has a 2xx status code
func (o *LocationFindDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this location find default response has a 3xx status code
func (o *LocationFindDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this location find default response has a 4xx status code
func (o *LocationFindDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this location find default response has a 5xx status code
func (o *LocationFindDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this location find default response a status code equal to that given
func (o *LocationFindDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the location find default response
func (o *LocationFindDefault) Code() int {
	return o._statusCode
}

func (o *LocationFindDefault) Error() string {
	return fmt.Sprintf("[GET /Locations][%d] Location.find default  %+v", o._statusCode, o.Payload)
}

func (o *LocationFindDefault) String() string {
	return fmt.Sprintf("[GET /Locations][%d] Location.find default  %+v", o._statusCode, o.Payload)
}

func (o *LocationFindDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *LocationFindDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}