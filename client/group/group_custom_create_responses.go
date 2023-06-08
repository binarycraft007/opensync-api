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

// GroupCustomCreateReader is a Reader for the GroupCustomCreate structure.
type GroupCustomCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GroupCustomCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGroupCustomCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGroupCustomCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGroupCustomCreateOK creates a GroupCustomCreateOK with default headers values
func NewGroupCustomCreateOK() *GroupCustomCreateOK {
	return &GroupCustomCreateOK{}
}

/*
GroupCustomCreateOK describes a response with status code 200, with default header values.

Request was successful
*/
type GroupCustomCreateOK struct {
	Payload *models.Group
}

// IsSuccess returns true when this group custom create o k response has a 2xx status code
func (o *GroupCustomCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this group custom create o k response has a 3xx status code
func (o *GroupCustomCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this group custom create o k response has a 4xx status code
func (o *GroupCustomCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this group custom create o k response has a 5xx status code
func (o *GroupCustomCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this group custom create o k response a status code equal to that given
func (o *GroupCustomCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the group custom create o k response
func (o *GroupCustomCreateOK) Code() int {
	return 200
}

func (o *GroupCustomCreateOK) Error() string {
	return fmt.Sprintf("[POST /Groups][%d] groupCustomCreateOK  %+v", 200, o.Payload)
}

func (o *GroupCustomCreateOK) String() string {
	return fmt.Sprintf("[POST /Groups][%d] groupCustomCreateOK  %+v", 200, o.Payload)
}

func (o *GroupCustomCreateOK) GetPayload() *models.Group {
	return o.Payload
}

func (o *GroupCustomCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Group)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGroupCustomCreateDefault creates a GroupCustomCreateDefault with default headers values
func NewGroupCustomCreateDefault(code int) *GroupCustomCreateDefault {
	return &GroupCustomCreateDefault{
		_statusCode: code,
	}
}

/*
GroupCustomCreateDefault describes a response with status code -1, with default header values.

unexpected error
*/
type GroupCustomCreateDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this group custom create default response has a 2xx status code
func (o *GroupCustomCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this group custom create default response has a 3xx status code
func (o *GroupCustomCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this group custom create default response has a 4xx status code
func (o *GroupCustomCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this group custom create default response has a 5xx status code
func (o *GroupCustomCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this group custom create default response a status code equal to that given
func (o *GroupCustomCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the group custom create default response
func (o *GroupCustomCreateDefault) Code() int {
	return o._statusCode
}

func (o *GroupCustomCreateDefault) Error() string {
	return fmt.Sprintf("[POST /Groups][%d] Group.customCreate default  %+v", o._statusCode, o.Payload)
}

func (o *GroupCustomCreateDefault) String() string {
	return fmt.Sprintf("[POST /Groups][%d] Group.customCreate default  %+v", o._statusCode, o.Payload)
}

func (o *GroupCustomCreateDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *GroupCustomCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
