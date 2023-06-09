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

// LocationPrototypeCreateInvitationsReader is a Reader for the LocationPrototypeCreateInvitations structure.
type LocationPrototypeCreateInvitationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LocationPrototypeCreateInvitationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLocationPrototypeCreateInvitationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLocationPrototypeCreateInvitationsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLocationPrototypeCreateInvitationsOK creates a LocationPrototypeCreateInvitationsOK with default headers values
func NewLocationPrototypeCreateInvitationsOK() *LocationPrototypeCreateInvitationsOK {
	return &LocationPrototypeCreateInvitationsOK{}
}

/*
LocationPrototypeCreateInvitationsOK describes a response with status code 200, with default header values.

Request was successful
*/
type LocationPrototypeCreateInvitationsOK struct {
	Payload *models.Invitation
}

// IsSuccess returns true when this location prototype create invitations o k response has a 2xx status code
func (o *LocationPrototypeCreateInvitationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this location prototype create invitations o k response has a 3xx status code
func (o *LocationPrototypeCreateInvitationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this location prototype create invitations o k response has a 4xx status code
func (o *LocationPrototypeCreateInvitationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this location prototype create invitations o k response has a 5xx status code
func (o *LocationPrototypeCreateInvitationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this location prototype create invitations o k response a status code equal to that given
func (o *LocationPrototypeCreateInvitationsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the location prototype create invitations o k response
func (o *LocationPrototypeCreateInvitationsOK) Code() int {
	return 200
}

func (o *LocationPrototypeCreateInvitationsOK) Error() string {
	return fmt.Sprintf("[POST /Locations/{id}/invitations][%d] locationPrototypeCreateInvitationsOK  %+v", 200, o.Payload)
}

func (o *LocationPrototypeCreateInvitationsOK) String() string {
	return fmt.Sprintf("[POST /Locations/{id}/invitations][%d] locationPrototypeCreateInvitationsOK  %+v", 200, o.Payload)
}

func (o *LocationPrototypeCreateInvitationsOK) GetPayload() *models.Invitation {
	return o.Payload
}

func (o *LocationPrototypeCreateInvitationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Invitation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLocationPrototypeCreateInvitationsDefault creates a LocationPrototypeCreateInvitationsDefault with default headers values
func NewLocationPrototypeCreateInvitationsDefault(code int) *LocationPrototypeCreateInvitationsDefault {
	return &LocationPrototypeCreateInvitationsDefault{
		_statusCode: code,
	}
}

/*
LocationPrototypeCreateInvitationsDefault describes a response with status code -1, with default header values.

unexpected error
*/
type LocationPrototypeCreateInvitationsDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this location prototype create invitations default response has a 2xx status code
func (o *LocationPrototypeCreateInvitationsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this location prototype create invitations default response has a 3xx status code
func (o *LocationPrototypeCreateInvitationsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this location prototype create invitations default response has a 4xx status code
func (o *LocationPrototypeCreateInvitationsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this location prototype create invitations default response has a 5xx status code
func (o *LocationPrototypeCreateInvitationsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this location prototype create invitations default response a status code equal to that given
func (o *LocationPrototypeCreateInvitationsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the location prototype create invitations default response
func (o *LocationPrototypeCreateInvitationsDefault) Code() int {
	return o._statusCode
}

func (o *LocationPrototypeCreateInvitationsDefault) Error() string {
	return fmt.Sprintf("[POST /Locations/{id}/invitations][%d] Location.prototype.__create__invitations default  %+v", o._statusCode, o.Payload)
}

func (o *LocationPrototypeCreateInvitationsDefault) String() string {
	return fmt.Sprintf("[POST /Locations/{id}/invitations][%d] Location.prototype.__create__invitations default  %+v", o._statusCode, o.Payload)
}

func (o *LocationPrototypeCreateInvitationsDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *LocationPrototypeCreateInvitationsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
