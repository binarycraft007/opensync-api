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

// LocationPrototypeUpdateByIDInvitationsReader is a Reader for the LocationPrototypeUpdateByIDInvitations structure.
type LocationPrototypeUpdateByIDInvitationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LocationPrototypeUpdateByIDInvitationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLocationPrototypeUpdateByIDInvitationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLocationPrototypeUpdateByIDInvitationsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLocationPrototypeUpdateByIDInvitationsOK creates a LocationPrototypeUpdateByIDInvitationsOK with default headers values
func NewLocationPrototypeUpdateByIDInvitationsOK() *LocationPrototypeUpdateByIDInvitationsOK {
	return &LocationPrototypeUpdateByIDInvitationsOK{}
}

/*
LocationPrototypeUpdateByIDInvitationsOK describes a response with status code 200, with default header values.

Request was successful
*/
type LocationPrototypeUpdateByIDInvitationsOK struct {
	Payload *models.Invitation
}

// IsSuccess returns true when this location prototype update by Id invitations o k response has a 2xx status code
func (o *LocationPrototypeUpdateByIDInvitationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this location prototype update by Id invitations o k response has a 3xx status code
func (o *LocationPrototypeUpdateByIDInvitationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this location prototype update by Id invitations o k response has a 4xx status code
func (o *LocationPrototypeUpdateByIDInvitationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this location prototype update by Id invitations o k response has a 5xx status code
func (o *LocationPrototypeUpdateByIDInvitationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this location prototype update by Id invitations o k response a status code equal to that given
func (o *LocationPrototypeUpdateByIDInvitationsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the location prototype update by Id invitations o k response
func (o *LocationPrototypeUpdateByIDInvitationsOK) Code() int {
	return 200
}

func (o *LocationPrototypeUpdateByIDInvitationsOK) Error() string {
	return fmt.Sprintf("[PUT /Locations/{id}/invitations/{fk}][%d] locationPrototypeUpdateByIdInvitationsOK  %+v", 200, o.Payload)
}

func (o *LocationPrototypeUpdateByIDInvitationsOK) String() string {
	return fmt.Sprintf("[PUT /Locations/{id}/invitations/{fk}][%d] locationPrototypeUpdateByIdInvitationsOK  %+v", 200, o.Payload)
}

func (o *LocationPrototypeUpdateByIDInvitationsOK) GetPayload() *models.Invitation {
	return o.Payload
}

func (o *LocationPrototypeUpdateByIDInvitationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Invitation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLocationPrototypeUpdateByIDInvitationsDefault creates a LocationPrototypeUpdateByIDInvitationsDefault with default headers values
func NewLocationPrototypeUpdateByIDInvitationsDefault(code int) *LocationPrototypeUpdateByIDInvitationsDefault {
	return &LocationPrototypeUpdateByIDInvitationsDefault{
		_statusCode: code,
	}
}

/*
LocationPrototypeUpdateByIDInvitationsDefault describes a response with status code -1, with default header values.

unexpected error
*/
type LocationPrototypeUpdateByIDInvitationsDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this location prototype update by Id invitations default response has a 2xx status code
func (o *LocationPrototypeUpdateByIDInvitationsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this location prototype update by Id invitations default response has a 3xx status code
func (o *LocationPrototypeUpdateByIDInvitationsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this location prototype update by Id invitations default response has a 4xx status code
func (o *LocationPrototypeUpdateByIDInvitationsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this location prototype update by Id invitations default response has a 5xx status code
func (o *LocationPrototypeUpdateByIDInvitationsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this location prototype update by Id invitations default response a status code equal to that given
func (o *LocationPrototypeUpdateByIDInvitationsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the location prototype update by Id invitations default response
func (o *LocationPrototypeUpdateByIDInvitationsDefault) Code() int {
	return o._statusCode
}

func (o *LocationPrototypeUpdateByIDInvitationsDefault) Error() string {
	return fmt.Sprintf("[PUT /Locations/{id}/invitations/{fk}][%d] Location.prototype.__updateById__invitations default  %+v", o._statusCode, o.Payload)
}

func (o *LocationPrototypeUpdateByIDInvitationsDefault) String() string {
	return fmt.Sprintf("[PUT /Locations/{id}/invitations/{fk}][%d] Location.prototype.__updateById__invitations default  %+v", o._statusCode, o.Payload)
}

func (o *LocationPrototypeUpdateByIDInvitationsDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *LocationPrototypeUpdateByIDInvitationsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
