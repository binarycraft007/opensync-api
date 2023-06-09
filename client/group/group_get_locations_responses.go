// Code generated by go-swagger; DO NOT EDIT.

package group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/binarycraft007/opensync-api/models"
)

// GroupGetLocationsReader is a Reader for the GroupGetLocations structure.
type GroupGetLocationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GroupGetLocationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGroupGetLocationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGroupGetLocationsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGroupGetLocationsOK creates a GroupGetLocationsOK with default headers values
func NewGroupGetLocationsOK() *GroupGetLocationsOK {
	return &GroupGetLocationsOK{}
}

/*
GroupGetLocationsOK describes a response with status code 200, with default header values.

Request was successful
*/
type GroupGetLocationsOK struct {
	Payload interface{}
}

// IsSuccess returns true when this group get locations o k response has a 2xx status code
func (o *GroupGetLocationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this group get locations o k response has a 3xx status code
func (o *GroupGetLocationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this group get locations o k response has a 4xx status code
func (o *GroupGetLocationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this group get locations o k response has a 5xx status code
func (o *GroupGetLocationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this group get locations o k response a status code equal to that given
func (o *GroupGetLocationsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the group get locations o k response
func (o *GroupGetLocationsOK) Code() int {
	return 200
}

func (o *GroupGetLocationsOK) Error() string {
	return fmt.Sprintf("[GET /Groups/locations/{keyword}][%d] groupGetLocationsOK  %+v", 200, o.Payload)
}

func (o *GroupGetLocationsOK) String() string {
	return fmt.Sprintf("[GET /Groups/locations/{keyword}][%d] groupGetLocationsOK  %+v", 200, o.Payload)
}

func (o *GroupGetLocationsOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GroupGetLocationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGroupGetLocationsDefault creates a GroupGetLocationsDefault with default headers values
func NewGroupGetLocationsDefault(code int) *GroupGetLocationsDefault {
	return &GroupGetLocationsDefault{
		_statusCode: code,
	}
}

/*
GroupGetLocationsDefault describes a response with status code -1, with default header values.

unexpected error
*/
type GroupGetLocationsDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this group get locations default response has a 2xx status code
func (o *GroupGetLocationsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this group get locations default response has a 3xx status code
func (o *GroupGetLocationsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this group get locations default response has a 4xx status code
func (o *GroupGetLocationsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this group get locations default response has a 5xx status code
func (o *GroupGetLocationsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this group get locations default response a status code equal to that given
func (o *GroupGetLocationsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the group get locations default response
func (o *GroupGetLocationsDefault) Code() int {
	return o._statusCode
}

func (o *GroupGetLocationsDefault) Error() string {
	return fmt.Sprintf("[GET /Groups/locations/{keyword}][%d] Group.getLocations default  %+v", o._statusCode, o.Payload)
}

func (o *GroupGetLocationsDefault) String() string {
	return fmt.Sprintf("[GET /Groups/locations/{keyword}][%d] Group.getLocations default  %+v", o._statusCode, o.Payload)
}

func (o *GroupGetLocationsDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *GroupGetLocationsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
