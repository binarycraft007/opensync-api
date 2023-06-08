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

// LocationFindByIDReader is a Reader for the LocationFindByID structure.
type LocationFindByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LocationFindByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLocationFindByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLocationFindByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLocationFindByIDOK creates a LocationFindByIDOK with default headers values
func NewLocationFindByIDOK() *LocationFindByIDOK {
	return &LocationFindByIDOK{}
}

/*
LocationFindByIDOK describes a response with status code 200, with default header values.

Request was successful
*/
type LocationFindByIDOK struct {
	Payload *models.Location
}

// IsSuccess returns true when this location find by Id o k response has a 2xx status code
func (o *LocationFindByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this location find by Id o k response has a 3xx status code
func (o *LocationFindByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this location find by Id o k response has a 4xx status code
func (o *LocationFindByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this location find by Id o k response has a 5xx status code
func (o *LocationFindByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this location find by Id o k response a status code equal to that given
func (o *LocationFindByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the location find by Id o k response
func (o *LocationFindByIDOK) Code() int {
	return 200
}

func (o *LocationFindByIDOK) Error() string {
	return fmt.Sprintf("[GET /Locations/{id}][%d] locationFindByIdOK  %+v", 200, o.Payload)
}

func (o *LocationFindByIDOK) String() string {
	return fmt.Sprintf("[GET /Locations/{id}][%d] locationFindByIdOK  %+v", 200, o.Payload)
}

func (o *LocationFindByIDOK) GetPayload() *models.Location {
	return o.Payload
}

func (o *LocationFindByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Location)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLocationFindByIDDefault creates a LocationFindByIDDefault with default headers values
func NewLocationFindByIDDefault(code int) *LocationFindByIDDefault {
	return &LocationFindByIDDefault{
		_statusCode: code,
	}
}

/*
LocationFindByIDDefault describes a response with status code -1, with default header values.

unexpected error
*/
type LocationFindByIDDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this location find by Id default response has a 2xx status code
func (o *LocationFindByIDDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this location find by Id default response has a 3xx status code
func (o *LocationFindByIDDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this location find by Id default response has a 4xx status code
func (o *LocationFindByIDDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this location find by Id default response has a 5xx status code
func (o *LocationFindByIDDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this location find by Id default response a status code equal to that given
func (o *LocationFindByIDDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the location find by Id default response
func (o *LocationFindByIDDefault) Code() int {
	return o._statusCode
}

func (o *LocationFindByIDDefault) Error() string {
	return fmt.Sprintf("[GET /Locations/{id}][%d] Location.findById default  %+v", o._statusCode, o.Payload)
}

func (o *LocationFindByIDDefault) String() string {
	return fmt.Sprintf("[GET /Locations/{id}][%d] Location.findById default  %+v", o._statusCode, o.Payload)
}

func (o *LocationFindByIDDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *LocationFindByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}