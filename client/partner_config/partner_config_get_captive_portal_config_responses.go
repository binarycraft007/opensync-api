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

// PartnerConfigGetCaptivePortalConfigReader is a Reader for the PartnerConfigGetCaptivePortalConfig structure.
type PartnerConfigGetCaptivePortalConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PartnerConfigGetCaptivePortalConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPartnerConfigGetCaptivePortalConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPartnerConfigGetCaptivePortalConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPartnerConfigGetCaptivePortalConfigOK creates a PartnerConfigGetCaptivePortalConfigOK with default headers values
func NewPartnerConfigGetCaptivePortalConfigOK() *PartnerConfigGetCaptivePortalConfigOK {
	return &PartnerConfigGetCaptivePortalConfigOK{}
}

/*
PartnerConfigGetCaptivePortalConfigOK describes a response with status code 200, with default header values.

Request was successful
*/
type PartnerConfigGetCaptivePortalConfigOK struct {
	Payload models.XAny
}

// IsSuccess returns true when this partner config get captive portal config o k response has a 2xx status code
func (o *PartnerConfigGetCaptivePortalConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this partner config get captive portal config o k response has a 3xx status code
func (o *PartnerConfigGetCaptivePortalConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this partner config get captive portal config o k response has a 4xx status code
func (o *PartnerConfigGetCaptivePortalConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this partner config get captive portal config o k response has a 5xx status code
func (o *PartnerConfigGetCaptivePortalConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this partner config get captive portal config o k response a status code equal to that given
func (o *PartnerConfigGetCaptivePortalConfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the partner config get captive portal config o k response
func (o *PartnerConfigGetCaptivePortalConfigOK) Code() int {
	return 200
}

func (o *PartnerConfigGetCaptivePortalConfigOK) Error() string {
	return fmt.Sprintf("[GET /partnerConfig/{id}/captivePortal][%d] partnerConfigGetCaptivePortalConfigOK  %+v", 200, o.Payload)
}

func (o *PartnerConfigGetCaptivePortalConfigOK) String() string {
	return fmt.Sprintf("[GET /partnerConfig/{id}/captivePortal][%d] partnerConfigGetCaptivePortalConfigOK  %+v", 200, o.Payload)
}

func (o *PartnerConfigGetCaptivePortalConfigOK) GetPayload() models.XAny {
	return o.Payload
}

func (o *PartnerConfigGetCaptivePortalConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPartnerConfigGetCaptivePortalConfigDefault creates a PartnerConfigGetCaptivePortalConfigDefault with default headers values
func NewPartnerConfigGetCaptivePortalConfigDefault(code int) *PartnerConfigGetCaptivePortalConfigDefault {
	return &PartnerConfigGetCaptivePortalConfigDefault{
		_statusCode: code,
	}
}

/*
PartnerConfigGetCaptivePortalConfigDefault describes a response with status code -1, with default header values.

unexpected error
*/
type PartnerConfigGetCaptivePortalConfigDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this partner config get captive portal config default response has a 2xx status code
func (o *PartnerConfigGetCaptivePortalConfigDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this partner config get captive portal config default response has a 3xx status code
func (o *PartnerConfigGetCaptivePortalConfigDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this partner config get captive portal config default response has a 4xx status code
func (o *PartnerConfigGetCaptivePortalConfigDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this partner config get captive portal config default response has a 5xx status code
func (o *PartnerConfigGetCaptivePortalConfigDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this partner config get captive portal config default response a status code equal to that given
func (o *PartnerConfigGetCaptivePortalConfigDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the partner config get captive portal config default response
func (o *PartnerConfigGetCaptivePortalConfigDefault) Code() int {
	return o._statusCode
}

func (o *PartnerConfigGetCaptivePortalConfigDefault) Error() string {
	return fmt.Sprintf("[GET /partnerConfig/{id}/captivePortal][%d] PartnerConfig.getCaptivePortalConfig default  %+v", o._statusCode, o.Payload)
}

func (o *PartnerConfigGetCaptivePortalConfigDefault) String() string {
	return fmt.Sprintf("[GET /partnerConfig/{id}/captivePortal][%d] PartnerConfig.getCaptivePortalConfig default  %+v", o._statusCode, o.Payload)
}

func (o *PartnerConfigGetCaptivePortalConfigDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *PartnerConfigGetCaptivePortalConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
