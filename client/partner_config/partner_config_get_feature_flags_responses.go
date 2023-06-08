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

// PartnerConfigGetFeatureFlagsReader is a Reader for the PartnerConfigGetFeatureFlags structure.
type PartnerConfigGetFeatureFlagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PartnerConfigGetFeatureFlagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPartnerConfigGetFeatureFlagsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPartnerConfigGetFeatureFlagsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPartnerConfigGetFeatureFlagsOK creates a PartnerConfigGetFeatureFlagsOK with default headers values
func NewPartnerConfigGetFeatureFlagsOK() *PartnerConfigGetFeatureFlagsOK {
	return &PartnerConfigGetFeatureFlagsOK{}
}

/*
PartnerConfigGetFeatureFlagsOK describes a response with status code 200, with default header values.

Request was successful
*/
type PartnerConfigGetFeatureFlagsOK struct {
	Payload models.XAny
}

// IsSuccess returns true when this partner config get feature flags o k response has a 2xx status code
func (o *PartnerConfigGetFeatureFlagsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this partner config get feature flags o k response has a 3xx status code
func (o *PartnerConfigGetFeatureFlagsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this partner config get feature flags o k response has a 4xx status code
func (o *PartnerConfigGetFeatureFlagsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this partner config get feature flags o k response has a 5xx status code
func (o *PartnerConfigGetFeatureFlagsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this partner config get feature flags o k response a status code equal to that given
func (o *PartnerConfigGetFeatureFlagsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the partner config get feature flags o k response
func (o *PartnerConfigGetFeatureFlagsOK) Code() int {
	return 200
}

func (o *PartnerConfigGetFeatureFlagsOK) Error() string {
	return fmt.Sprintf("[GET /partnerConfig/{id}/featureFlags][%d] partnerConfigGetFeatureFlagsOK  %+v", 200, o.Payload)
}

func (o *PartnerConfigGetFeatureFlagsOK) String() string {
	return fmt.Sprintf("[GET /partnerConfig/{id}/featureFlags][%d] partnerConfigGetFeatureFlagsOK  %+v", 200, o.Payload)
}

func (o *PartnerConfigGetFeatureFlagsOK) GetPayload() models.XAny {
	return o.Payload
}

func (o *PartnerConfigGetFeatureFlagsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPartnerConfigGetFeatureFlagsDefault creates a PartnerConfigGetFeatureFlagsDefault with default headers values
func NewPartnerConfigGetFeatureFlagsDefault(code int) *PartnerConfigGetFeatureFlagsDefault {
	return &PartnerConfigGetFeatureFlagsDefault{
		_statusCode: code,
	}
}

/*
PartnerConfigGetFeatureFlagsDefault describes a response with status code -1, with default header values.

unexpected error
*/
type PartnerConfigGetFeatureFlagsDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this partner config get feature flags default response has a 2xx status code
func (o *PartnerConfigGetFeatureFlagsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this partner config get feature flags default response has a 3xx status code
func (o *PartnerConfigGetFeatureFlagsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this partner config get feature flags default response has a 4xx status code
func (o *PartnerConfigGetFeatureFlagsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this partner config get feature flags default response has a 5xx status code
func (o *PartnerConfigGetFeatureFlagsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this partner config get feature flags default response a status code equal to that given
func (o *PartnerConfigGetFeatureFlagsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the partner config get feature flags default response
func (o *PartnerConfigGetFeatureFlagsDefault) Code() int {
	return o._statusCode
}

func (o *PartnerConfigGetFeatureFlagsDefault) Error() string {
	return fmt.Sprintf("[GET /partnerConfig/{id}/featureFlags][%d] PartnerConfig.getFeatureFlags default  %+v", o._statusCode, o.Payload)
}

func (o *PartnerConfigGetFeatureFlagsDefault) String() string {
	return fmt.Sprintf("[GET /partnerConfig/{id}/featureFlags][%d] PartnerConfig.getFeatureFlags default  %+v", o._statusCode, o.Payload)
}

func (o *PartnerConfigGetFeatureFlagsDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *PartnerConfigGetFeatureFlagsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
