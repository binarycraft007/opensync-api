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

// PartnerConfigUpdateSpeedTestReader is a Reader for the PartnerConfigUpdateSpeedTest structure.
type PartnerConfigUpdateSpeedTestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PartnerConfigUpdateSpeedTestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPartnerConfigUpdateSpeedTestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPartnerConfigUpdateSpeedTestDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPartnerConfigUpdateSpeedTestOK creates a PartnerConfigUpdateSpeedTestOK with default headers values
func NewPartnerConfigUpdateSpeedTestOK() *PartnerConfigUpdateSpeedTestOK {
	return &PartnerConfigUpdateSpeedTestOK{}
}

/*
PartnerConfigUpdateSpeedTestOK describes a response with status code 200, with default header values.

Request was successful
*/
type PartnerConfigUpdateSpeedTestOK struct {
	Payload interface{}
}

// IsSuccess returns true when this partner config update speed test o k response has a 2xx status code
func (o *PartnerConfigUpdateSpeedTestOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this partner config update speed test o k response has a 3xx status code
func (o *PartnerConfigUpdateSpeedTestOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this partner config update speed test o k response has a 4xx status code
func (o *PartnerConfigUpdateSpeedTestOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this partner config update speed test o k response has a 5xx status code
func (o *PartnerConfigUpdateSpeedTestOK) IsServerError() bool {
	return false
}

// IsCode returns true when this partner config update speed test o k response a status code equal to that given
func (o *PartnerConfigUpdateSpeedTestOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the partner config update speed test o k response
func (o *PartnerConfigUpdateSpeedTestOK) Code() int {
	return 200
}

func (o *PartnerConfigUpdateSpeedTestOK) Error() string {
	return fmt.Sprintf("[PUT /partnerConfig/{id}/platform/speedTest][%d] partnerConfigUpdateSpeedTestOK  %+v", 200, o.Payload)
}

func (o *PartnerConfigUpdateSpeedTestOK) String() string {
	return fmt.Sprintf("[PUT /partnerConfig/{id}/platform/speedTest][%d] partnerConfigUpdateSpeedTestOK  %+v", 200, o.Payload)
}

func (o *PartnerConfigUpdateSpeedTestOK) GetPayload() interface{} {
	return o.Payload
}

func (o *PartnerConfigUpdateSpeedTestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPartnerConfigUpdateSpeedTestDefault creates a PartnerConfigUpdateSpeedTestDefault with default headers values
func NewPartnerConfigUpdateSpeedTestDefault(code int) *PartnerConfigUpdateSpeedTestDefault {
	return &PartnerConfigUpdateSpeedTestDefault{
		_statusCode: code,
	}
}

/*
PartnerConfigUpdateSpeedTestDefault describes a response with status code -1, with default header values.

unexpected error
*/
type PartnerConfigUpdateSpeedTestDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this partner config update speed test default response has a 2xx status code
func (o *PartnerConfigUpdateSpeedTestDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this partner config update speed test default response has a 3xx status code
func (o *PartnerConfigUpdateSpeedTestDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this partner config update speed test default response has a 4xx status code
func (o *PartnerConfigUpdateSpeedTestDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this partner config update speed test default response has a 5xx status code
func (o *PartnerConfigUpdateSpeedTestDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this partner config update speed test default response a status code equal to that given
func (o *PartnerConfigUpdateSpeedTestDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the partner config update speed test default response
func (o *PartnerConfigUpdateSpeedTestDefault) Code() int {
	return o._statusCode
}

func (o *PartnerConfigUpdateSpeedTestDefault) Error() string {
	return fmt.Sprintf("[PUT /partnerConfig/{id}/platform/speedTest][%d] PartnerConfig.updateSpeedTest default  %+v", o._statusCode, o.Payload)
}

func (o *PartnerConfigUpdateSpeedTestDefault) String() string {
	return fmt.Sprintf("[PUT /partnerConfig/{id}/platform/speedTest][%d] PartnerConfig.updateSpeedTest default  %+v", o._statusCode, o.Payload)
}

func (o *PartnerConfigUpdateSpeedTestDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *PartnerConfigUpdateSpeedTestDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
