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

// PartnerConfigDeleteAppPrioritizationPartnerCustomSettingReader is a Reader for the PartnerConfigDeleteAppPrioritizationPartnerCustomSetting structure.
type PartnerConfigDeleteAppPrioritizationPartnerCustomSettingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PartnerConfigDeleteAppPrioritizationPartnerCustomSettingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPartnerConfigDeleteAppPrioritizationPartnerCustomSettingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPartnerConfigDeleteAppPrioritizationPartnerCustomSettingDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPartnerConfigDeleteAppPrioritizationPartnerCustomSettingOK creates a PartnerConfigDeleteAppPrioritizationPartnerCustomSettingOK with default headers values
func NewPartnerConfigDeleteAppPrioritizationPartnerCustomSettingOK() *PartnerConfigDeleteAppPrioritizationPartnerCustomSettingOK {
	return &PartnerConfigDeleteAppPrioritizationPartnerCustomSettingOK{}
}

/*
PartnerConfigDeleteAppPrioritizationPartnerCustomSettingOK describes a response with status code 200, with default header values.

Request was successful
*/
type PartnerConfigDeleteAppPrioritizationPartnerCustomSettingOK struct {
	Payload interface{}
}

// IsSuccess returns true when this partner config delete app prioritization partner custom setting o k response has a 2xx status code
func (o *PartnerConfigDeleteAppPrioritizationPartnerCustomSettingOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this partner config delete app prioritization partner custom setting o k response has a 3xx status code
func (o *PartnerConfigDeleteAppPrioritizationPartnerCustomSettingOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this partner config delete app prioritization partner custom setting o k response has a 4xx status code
func (o *PartnerConfigDeleteAppPrioritizationPartnerCustomSettingOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this partner config delete app prioritization partner custom setting o k response has a 5xx status code
func (o *PartnerConfigDeleteAppPrioritizationPartnerCustomSettingOK) IsServerError() bool {
	return false
}

// IsCode returns true when this partner config delete app prioritization partner custom setting o k response a status code equal to that given
func (o *PartnerConfigDeleteAppPrioritizationPartnerCustomSettingOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the partner config delete app prioritization partner custom setting o k response
func (o *PartnerConfigDeleteAppPrioritizationPartnerCustomSettingOK) Code() int {
	return 200
}

func (o *PartnerConfigDeleteAppPrioritizationPartnerCustomSettingOK) Error() string {
	return fmt.Sprintf("[DELETE /partnerConfig/{partnerId}/qos/appPrioritization/customSetting][%d] partnerConfigDeleteAppPrioritizationPartnerCustomSettingOK  %+v", 200, o.Payload)
}

func (o *PartnerConfigDeleteAppPrioritizationPartnerCustomSettingOK) String() string {
	return fmt.Sprintf("[DELETE /partnerConfig/{partnerId}/qos/appPrioritization/customSetting][%d] partnerConfigDeleteAppPrioritizationPartnerCustomSettingOK  %+v", 200, o.Payload)
}

func (o *PartnerConfigDeleteAppPrioritizationPartnerCustomSettingOK) GetPayload() interface{} {
	return o.Payload
}

func (o *PartnerConfigDeleteAppPrioritizationPartnerCustomSettingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPartnerConfigDeleteAppPrioritizationPartnerCustomSettingDefault creates a PartnerConfigDeleteAppPrioritizationPartnerCustomSettingDefault with default headers values
func NewPartnerConfigDeleteAppPrioritizationPartnerCustomSettingDefault(code int) *PartnerConfigDeleteAppPrioritizationPartnerCustomSettingDefault {
	return &PartnerConfigDeleteAppPrioritizationPartnerCustomSettingDefault{
		_statusCode: code,
	}
}

/*
PartnerConfigDeleteAppPrioritizationPartnerCustomSettingDefault describes a response with status code -1, with default header values.

unexpected error
*/
type PartnerConfigDeleteAppPrioritizationPartnerCustomSettingDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this partner config delete app prioritization partner custom setting default response has a 2xx status code
func (o *PartnerConfigDeleteAppPrioritizationPartnerCustomSettingDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this partner config delete app prioritization partner custom setting default response has a 3xx status code
func (o *PartnerConfigDeleteAppPrioritizationPartnerCustomSettingDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this partner config delete app prioritization partner custom setting default response has a 4xx status code
func (o *PartnerConfigDeleteAppPrioritizationPartnerCustomSettingDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this partner config delete app prioritization partner custom setting default response has a 5xx status code
func (o *PartnerConfigDeleteAppPrioritizationPartnerCustomSettingDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this partner config delete app prioritization partner custom setting default response a status code equal to that given
func (o *PartnerConfigDeleteAppPrioritizationPartnerCustomSettingDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the partner config delete app prioritization partner custom setting default response
func (o *PartnerConfigDeleteAppPrioritizationPartnerCustomSettingDefault) Code() int {
	return o._statusCode
}

func (o *PartnerConfigDeleteAppPrioritizationPartnerCustomSettingDefault) Error() string {
	return fmt.Sprintf("[DELETE /partnerConfig/{partnerId}/qos/appPrioritization/customSetting][%d] PartnerConfig.deleteAppPrioritizationPartnerCustomSetting default  %+v", o._statusCode, o.Payload)
}

func (o *PartnerConfigDeleteAppPrioritizationPartnerCustomSettingDefault) String() string {
	return fmt.Sprintf("[DELETE /partnerConfig/{partnerId}/qos/appPrioritization/customSetting][%d] PartnerConfig.deleteAppPrioritizationPartnerCustomSetting default  %+v", o._statusCode, o.Payload)
}

func (o *PartnerConfigDeleteAppPrioritizationPartnerCustomSettingDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *PartnerConfigDeleteAppPrioritizationPartnerCustomSettingDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
