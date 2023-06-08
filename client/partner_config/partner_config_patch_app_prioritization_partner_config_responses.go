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

// PartnerConfigPatchAppPrioritizationPartnerConfigReader is a Reader for the PartnerConfigPatchAppPrioritizationPartnerConfig structure.
type PartnerConfigPatchAppPrioritizationPartnerConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PartnerConfigPatchAppPrioritizationPartnerConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPartnerConfigPatchAppPrioritizationPartnerConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPartnerConfigPatchAppPrioritizationPartnerConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPartnerConfigPatchAppPrioritizationPartnerConfigOK creates a PartnerConfigPatchAppPrioritizationPartnerConfigOK with default headers values
func NewPartnerConfigPatchAppPrioritizationPartnerConfigOK() *PartnerConfigPatchAppPrioritizationPartnerConfigOK {
	return &PartnerConfigPatchAppPrioritizationPartnerConfigOK{}
}

/*
PartnerConfigPatchAppPrioritizationPartnerConfigOK describes a response with status code 200, with default header values.

Request was successful
*/
type PartnerConfigPatchAppPrioritizationPartnerConfigOK struct {
	Payload models.XAny
}

// IsSuccess returns true when this partner config patch app prioritization partner config o k response has a 2xx status code
func (o *PartnerConfigPatchAppPrioritizationPartnerConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this partner config patch app prioritization partner config o k response has a 3xx status code
func (o *PartnerConfigPatchAppPrioritizationPartnerConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this partner config patch app prioritization partner config o k response has a 4xx status code
func (o *PartnerConfigPatchAppPrioritizationPartnerConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this partner config patch app prioritization partner config o k response has a 5xx status code
func (o *PartnerConfigPatchAppPrioritizationPartnerConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this partner config patch app prioritization partner config o k response a status code equal to that given
func (o *PartnerConfigPatchAppPrioritizationPartnerConfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the partner config patch app prioritization partner config o k response
func (o *PartnerConfigPatchAppPrioritizationPartnerConfigOK) Code() int {
	return 200
}

func (o *PartnerConfigPatchAppPrioritizationPartnerConfigOK) Error() string {
	return fmt.Sprintf("[PATCH /partnerConfig/{partnerId}/qos/appPrioritization][%d] partnerConfigPatchAppPrioritizationPartnerConfigOK  %+v", 200, o.Payload)
}

func (o *PartnerConfigPatchAppPrioritizationPartnerConfigOK) String() string {
	return fmt.Sprintf("[PATCH /partnerConfig/{partnerId}/qos/appPrioritization][%d] partnerConfigPatchAppPrioritizationPartnerConfigOK  %+v", 200, o.Payload)
}

func (o *PartnerConfigPatchAppPrioritizationPartnerConfigOK) GetPayload() models.XAny {
	return o.Payload
}

func (o *PartnerConfigPatchAppPrioritizationPartnerConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPartnerConfigPatchAppPrioritizationPartnerConfigDefault creates a PartnerConfigPatchAppPrioritizationPartnerConfigDefault with default headers values
func NewPartnerConfigPatchAppPrioritizationPartnerConfigDefault(code int) *PartnerConfigPatchAppPrioritizationPartnerConfigDefault {
	return &PartnerConfigPatchAppPrioritizationPartnerConfigDefault{
		_statusCode: code,
	}
}

/*
PartnerConfigPatchAppPrioritizationPartnerConfigDefault describes a response with status code -1, with default header values.

unexpected error
*/
type PartnerConfigPatchAppPrioritizationPartnerConfigDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this partner config patch app prioritization partner config default response has a 2xx status code
func (o *PartnerConfigPatchAppPrioritizationPartnerConfigDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this partner config patch app prioritization partner config default response has a 3xx status code
func (o *PartnerConfigPatchAppPrioritizationPartnerConfigDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this partner config patch app prioritization partner config default response has a 4xx status code
func (o *PartnerConfigPatchAppPrioritizationPartnerConfigDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this partner config patch app prioritization partner config default response has a 5xx status code
func (o *PartnerConfigPatchAppPrioritizationPartnerConfigDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this partner config patch app prioritization partner config default response a status code equal to that given
func (o *PartnerConfigPatchAppPrioritizationPartnerConfigDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the partner config patch app prioritization partner config default response
func (o *PartnerConfigPatchAppPrioritizationPartnerConfigDefault) Code() int {
	return o._statusCode
}

func (o *PartnerConfigPatchAppPrioritizationPartnerConfigDefault) Error() string {
	return fmt.Sprintf("[PATCH /partnerConfig/{partnerId}/qos/appPrioritization][%d] PartnerConfig.patchAppPrioritizationPartnerConfig default  %+v", o._statusCode, o.Payload)
}

func (o *PartnerConfigPatchAppPrioritizationPartnerConfigDefault) String() string {
	return fmt.Sprintf("[PATCH /partnerConfig/{partnerId}/qos/appPrioritization][%d] PartnerConfig.patchAppPrioritizationPartnerConfig default  %+v", o._statusCode, o.Payload)
}

func (o *PartnerConfigPatchAppPrioritizationPartnerConfigDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *PartnerConfigPatchAppPrioritizationPartnerConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}