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

// PartnerConfigUpdateSipAlgReader is a Reader for the PartnerConfigUpdateSipAlg structure.
type PartnerConfigUpdateSipAlgReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PartnerConfigUpdateSipAlgReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPartnerConfigUpdateSipAlgOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPartnerConfigUpdateSipAlgDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPartnerConfigUpdateSipAlgOK creates a PartnerConfigUpdateSipAlgOK with default headers values
func NewPartnerConfigUpdateSipAlgOK() *PartnerConfigUpdateSipAlgOK {
	return &PartnerConfigUpdateSipAlgOK{}
}

/*
PartnerConfigUpdateSipAlgOK describes a response with status code 200, with default header values.

Request was successful
*/
type PartnerConfigUpdateSipAlgOK struct {
	Payload interface{}
}

// IsSuccess returns true when this partner config update sip alg o k response has a 2xx status code
func (o *PartnerConfigUpdateSipAlgOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this partner config update sip alg o k response has a 3xx status code
func (o *PartnerConfigUpdateSipAlgOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this partner config update sip alg o k response has a 4xx status code
func (o *PartnerConfigUpdateSipAlgOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this partner config update sip alg o k response has a 5xx status code
func (o *PartnerConfigUpdateSipAlgOK) IsServerError() bool {
	return false
}

// IsCode returns true when this partner config update sip alg o k response a status code equal to that given
func (o *PartnerConfigUpdateSipAlgOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the partner config update sip alg o k response
func (o *PartnerConfigUpdateSipAlgOK) Code() int {
	return 200
}

func (o *PartnerConfigUpdateSipAlgOK) Error() string {
	return fmt.Sprintf("[PUT /partnerConfig/{id}/platform/sipAlg][%d] partnerConfigUpdateSipAlgOK  %+v", 200, o.Payload)
}

func (o *PartnerConfigUpdateSipAlgOK) String() string {
	return fmt.Sprintf("[PUT /partnerConfig/{id}/platform/sipAlg][%d] partnerConfigUpdateSipAlgOK  %+v", 200, o.Payload)
}

func (o *PartnerConfigUpdateSipAlgOK) GetPayload() interface{} {
	return o.Payload
}

func (o *PartnerConfigUpdateSipAlgOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPartnerConfigUpdateSipAlgDefault creates a PartnerConfigUpdateSipAlgDefault with default headers values
func NewPartnerConfigUpdateSipAlgDefault(code int) *PartnerConfigUpdateSipAlgDefault {
	return &PartnerConfigUpdateSipAlgDefault{
		_statusCode: code,
	}
}

/*
PartnerConfigUpdateSipAlgDefault describes a response with status code -1, with default header values.

unexpected error
*/
type PartnerConfigUpdateSipAlgDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this partner config update sip alg default response has a 2xx status code
func (o *PartnerConfigUpdateSipAlgDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this partner config update sip alg default response has a 3xx status code
func (o *PartnerConfigUpdateSipAlgDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this partner config update sip alg default response has a 4xx status code
func (o *PartnerConfigUpdateSipAlgDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this partner config update sip alg default response has a 5xx status code
func (o *PartnerConfigUpdateSipAlgDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this partner config update sip alg default response a status code equal to that given
func (o *PartnerConfigUpdateSipAlgDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the partner config update sip alg default response
func (o *PartnerConfigUpdateSipAlgDefault) Code() int {
	return o._statusCode
}

func (o *PartnerConfigUpdateSipAlgDefault) Error() string {
	return fmt.Sprintf("[PUT /partnerConfig/{id}/platform/sipAlg][%d] PartnerConfig.updateSipAlg default  %+v", o._statusCode, o.Payload)
}

func (o *PartnerConfigUpdateSipAlgDefault) String() string {
	return fmt.Sprintf("[PUT /partnerConfig/{id}/platform/sipAlg][%d] PartnerConfig.updateSipAlg default  %+v", o._statusCode, o.Payload)
}

func (o *PartnerConfigUpdateSipAlgDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *PartnerConfigUpdateSipAlgDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
