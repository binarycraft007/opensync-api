// Code generated by go-swagger; DO NOT EDIT.

package partner_config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/binarycraft007/opensync-api/models"
)

// PartnerConfigUpdatePreCacSchedulerReader is a Reader for the PartnerConfigUpdatePreCacScheduler structure.
type PartnerConfigUpdatePreCacSchedulerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PartnerConfigUpdatePreCacSchedulerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPartnerConfigUpdatePreCacSchedulerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPartnerConfigUpdatePreCacSchedulerDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPartnerConfigUpdatePreCacSchedulerOK creates a PartnerConfigUpdatePreCacSchedulerOK with default headers values
func NewPartnerConfigUpdatePreCacSchedulerOK() *PartnerConfigUpdatePreCacSchedulerOK {
	return &PartnerConfigUpdatePreCacSchedulerOK{}
}

/*
PartnerConfigUpdatePreCacSchedulerOK describes a response with status code 200, with default header values.

Request was successful
*/
type PartnerConfigUpdatePreCacSchedulerOK struct {
	Payload interface{}
}

// IsSuccess returns true when this partner config update pre cac scheduler o k response has a 2xx status code
func (o *PartnerConfigUpdatePreCacSchedulerOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this partner config update pre cac scheduler o k response has a 3xx status code
func (o *PartnerConfigUpdatePreCacSchedulerOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this partner config update pre cac scheduler o k response has a 4xx status code
func (o *PartnerConfigUpdatePreCacSchedulerOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this partner config update pre cac scheduler o k response has a 5xx status code
func (o *PartnerConfigUpdatePreCacSchedulerOK) IsServerError() bool {
	return false
}

// IsCode returns true when this partner config update pre cac scheduler o k response a status code equal to that given
func (o *PartnerConfigUpdatePreCacSchedulerOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the partner config update pre cac scheduler o k response
func (o *PartnerConfigUpdatePreCacSchedulerOK) Code() int {
	return 200
}

func (o *PartnerConfigUpdatePreCacSchedulerOK) Error() string {
	return fmt.Sprintf("[PUT /partnerConfig/{id}/platform/pcs][%d] partnerConfigUpdatePreCacSchedulerOK  %+v", 200, o.Payload)
}

func (o *PartnerConfigUpdatePreCacSchedulerOK) String() string {
	return fmt.Sprintf("[PUT /partnerConfig/{id}/platform/pcs][%d] partnerConfigUpdatePreCacSchedulerOK  %+v", 200, o.Payload)
}

func (o *PartnerConfigUpdatePreCacSchedulerOK) GetPayload() interface{} {
	return o.Payload
}

func (o *PartnerConfigUpdatePreCacSchedulerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPartnerConfigUpdatePreCacSchedulerDefault creates a PartnerConfigUpdatePreCacSchedulerDefault with default headers values
func NewPartnerConfigUpdatePreCacSchedulerDefault(code int) *PartnerConfigUpdatePreCacSchedulerDefault {
	return &PartnerConfigUpdatePreCacSchedulerDefault{
		_statusCode: code,
	}
}

/*
PartnerConfigUpdatePreCacSchedulerDefault describes a response with status code -1, with default header values.

unexpected error
*/
type PartnerConfigUpdatePreCacSchedulerDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this partner config update pre cac scheduler default response has a 2xx status code
func (o *PartnerConfigUpdatePreCacSchedulerDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this partner config update pre cac scheduler default response has a 3xx status code
func (o *PartnerConfigUpdatePreCacSchedulerDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this partner config update pre cac scheduler default response has a 4xx status code
func (o *PartnerConfigUpdatePreCacSchedulerDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this partner config update pre cac scheduler default response has a 5xx status code
func (o *PartnerConfigUpdatePreCacSchedulerDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this partner config update pre cac scheduler default response a status code equal to that given
func (o *PartnerConfigUpdatePreCacSchedulerDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the partner config update pre cac scheduler default response
func (o *PartnerConfigUpdatePreCacSchedulerDefault) Code() int {
	return o._statusCode
}

func (o *PartnerConfigUpdatePreCacSchedulerDefault) Error() string {
	return fmt.Sprintf("[PUT /partnerConfig/{id}/platform/pcs][%d] PartnerConfig.updatePreCacScheduler default  %+v", o._statusCode, o.Payload)
}

func (o *PartnerConfigUpdatePreCacSchedulerDefault) String() string {
	return fmt.Sprintf("[PUT /partnerConfig/{id}/platform/pcs][%d] PartnerConfig.updatePreCacScheduler default  %+v", o._statusCode, o.Payload)
}

func (o *PartnerConfigUpdatePreCacSchedulerDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *PartnerConfigUpdatePreCacSchedulerDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
