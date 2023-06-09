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

// GroupPrototypePatchAttributesPutGroupsIDReader is a Reader for the GroupPrototypePatchAttributesPutGroupsID structure.
type GroupPrototypePatchAttributesPutGroupsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GroupPrototypePatchAttributesPutGroupsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGroupPrototypePatchAttributesPutGroupsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGroupPrototypePatchAttributesPutGroupsIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGroupPrototypePatchAttributesPutGroupsIDOK creates a GroupPrototypePatchAttributesPutGroupsIDOK with default headers values
func NewGroupPrototypePatchAttributesPutGroupsIDOK() *GroupPrototypePatchAttributesPutGroupsIDOK {
	return &GroupPrototypePatchAttributesPutGroupsIDOK{}
}

/*
GroupPrototypePatchAttributesPutGroupsIDOK describes a response with status code 200, with default header values.

Request was successful
*/
type GroupPrototypePatchAttributesPutGroupsIDOK struct {
	Payload *models.Group
}

// IsSuccess returns true when this group prototype patch attributes put groups Id o k response has a 2xx status code
func (o *GroupPrototypePatchAttributesPutGroupsIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this group prototype patch attributes put groups Id o k response has a 3xx status code
func (o *GroupPrototypePatchAttributesPutGroupsIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this group prototype patch attributes put groups Id o k response has a 4xx status code
func (o *GroupPrototypePatchAttributesPutGroupsIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this group prototype patch attributes put groups Id o k response has a 5xx status code
func (o *GroupPrototypePatchAttributesPutGroupsIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this group prototype patch attributes put groups Id o k response a status code equal to that given
func (o *GroupPrototypePatchAttributesPutGroupsIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the group prototype patch attributes put groups Id o k response
func (o *GroupPrototypePatchAttributesPutGroupsIDOK) Code() int {
	return 200
}

func (o *GroupPrototypePatchAttributesPutGroupsIDOK) Error() string {
	return fmt.Sprintf("[PUT /Groups/{id}][%d] groupPrototypePatchAttributesPutGroupsIdOK  %+v", 200, o.Payload)
}

func (o *GroupPrototypePatchAttributesPutGroupsIDOK) String() string {
	return fmt.Sprintf("[PUT /Groups/{id}][%d] groupPrototypePatchAttributesPutGroupsIdOK  %+v", 200, o.Payload)
}

func (o *GroupPrototypePatchAttributesPutGroupsIDOK) GetPayload() *models.Group {
	return o.Payload
}

func (o *GroupPrototypePatchAttributesPutGroupsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Group)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGroupPrototypePatchAttributesPutGroupsIDDefault creates a GroupPrototypePatchAttributesPutGroupsIDDefault with default headers values
func NewGroupPrototypePatchAttributesPutGroupsIDDefault(code int) *GroupPrototypePatchAttributesPutGroupsIDDefault {
	return &GroupPrototypePatchAttributesPutGroupsIDDefault{
		_statusCode: code,
	}
}

/*
GroupPrototypePatchAttributesPutGroupsIDDefault describes a response with status code -1, with default header values.

unexpected error
*/
type GroupPrototypePatchAttributesPutGroupsIDDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this group prototype patch attributes put groups id default response has a 2xx status code
func (o *GroupPrototypePatchAttributesPutGroupsIDDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this group prototype patch attributes put groups id default response has a 3xx status code
func (o *GroupPrototypePatchAttributesPutGroupsIDDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this group prototype patch attributes put groups id default response has a 4xx status code
func (o *GroupPrototypePatchAttributesPutGroupsIDDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this group prototype patch attributes put groups id default response has a 5xx status code
func (o *GroupPrototypePatchAttributesPutGroupsIDDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this group prototype patch attributes put groups id default response a status code equal to that given
func (o *GroupPrototypePatchAttributesPutGroupsIDDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the group prototype patch attributes put groups id default response
func (o *GroupPrototypePatchAttributesPutGroupsIDDefault) Code() int {
	return o._statusCode
}

func (o *GroupPrototypePatchAttributesPutGroupsIDDefault) Error() string {
	return fmt.Sprintf("[PUT /Groups/{id}][%d] Group.prototype.patchAttributes__put_Groups_{id} default  %+v", o._statusCode, o.Payload)
}

func (o *GroupPrototypePatchAttributesPutGroupsIDDefault) String() string {
	return fmt.Sprintf("[PUT /Groups/{id}][%d] Group.prototype.patchAttributes__put_Groups_{id} default  %+v", o._statusCode, o.Payload)
}

func (o *GroupPrototypePatchAttributesPutGroupsIDDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *GroupPrototypePatchAttributesPutGroupsIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
