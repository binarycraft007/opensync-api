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

// PartnerConfigUpdateAppQoeReader is a Reader for the PartnerConfigUpdateAppQoe structure.
type PartnerConfigUpdateAppQoeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PartnerConfigUpdateAppQoeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPartnerConfigUpdateAppQoeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPartnerConfigUpdateAppQoeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPartnerConfigUpdateAppQoeOK creates a PartnerConfigUpdateAppQoeOK with default headers values
func NewPartnerConfigUpdateAppQoeOK() *PartnerConfigUpdateAppQoeOK {
	return &PartnerConfigUpdateAppQoeOK{}
}

/*
PartnerConfigUpdateAppQoeOK describes a response with status code 200, with default header values.

Request was successful
*/
type PartnerConfigUpdateAppQoeOK struct {
	Payload interface{}
}

// IsSuccess returns true when this partner config update app qoe o k response has a 2xx status code
func (o *PartnerConfigUpdateAppQoeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this partner config update app qoe o k response has a 3xx status code
func (o *PartnerConfigUpdateAppQoeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this partner config update app qoe o k response has a 4xx status code
func (o *PartnerConfigUpdateAppQoeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this partner config update app qoe o k response has a 5xx status code
func (o *PartnerConfigUpdateAppQoeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this partner config update app qoe o k response a status code equal to that given
func (o *PartnerConfigUpdateAppQoeOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the partner config update app qoe o k response
func (o *PartnerConfigUpdateAppQoeOK) Code() int {
	return 200
}

func (o *PartnerConfigUpdateAppQoeOK) Error() string {
	return fmt.Sprintf("[PUT /partnerConfig/{id}/platform/appQoe][%d] partnerConfigUpdateAppQoeOK  %+v", 200, o.Payload)
}

func (o *PartnerConfigUpdateAppQoeOK) String() string {
	return fmt.Sprintf("[PUT /partnerConfig/{id}/platform/appQoe][%d] partnerConfigUpdateAppQoeOK  %+v", 200, o.Payload)
}

func (o *PartnerConfigUpdateAppQoeOK) GetPayload() interface{} {
	return o.Payload
}

func (o *PartnerConfigUpdateAppQoeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPartnerConfigUpdateAppQoeDefault creates a PartnerConfigUpdateAppQoeDefault with default headers values
func NewPartnerConfigUpdateAppQoeDefault(code int) *PartnerConfigUpdateAppQoeDefault {
	return &PartnerConfigUpdateAppQoeDefault{
		_statusCode: code,
	}
}

/*
PartnerConfigUpdateAppQoeDefault describes a response with status code -1, with default header values.

unexpected error
*/
type PartnerConfigUpdateAppQoeDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this partner config update app qoe default response has a 2xx status code
func (o *PartnerConfigUpdateAppQoeDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this partner config update app qoe default response has a 3xx status code
func (o *PartnerConfigUpdateAppQoeDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this partner config update app qoe default response has a 4xx status code
func (o *PartnerConfigUpdateAppQoeDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this partner config update app qoe default response has a 5xx status code
func (o *PartnerConfigUpdateAppQoeDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this partner config update app qoe default response a status code equal to that given
func (o *PartnerConfigUpdateAppQoeDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the partner config update app qoe default response
func (o *PartnerConfigUpdateAppQoeDefault) Code() int {
	return o._statusCode
}

func (o *PartnerConfigUpdateAppQoeDefault) Error() string {
	return fmt.Sprintf("[PUT /partnerConfig/{id}/platform/appQoe][%d] PartnerConfig.updateAppQoe default  %+v", o._statusCode, o.Payload)
}

func (o *PartnerConfigUpdateAppQoeDefault) String() string {
	return fmt.Sprintf("[PUT /partnerConfig/{id}/platform/appQoe][%d] PartnerConfig.updateAppQoe default  %+v", o._statusCode, o.Payload)
}

func (o *PartnerConfigUpdateAppQoeDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *PartnerConfigUpdateAppQoeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
