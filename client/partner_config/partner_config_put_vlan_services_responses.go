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

// PartnerConfigPutVlanServicesReader is a Reader for the PartnerConfigPutVlanServices structure.
type PartnerConfigPutVlanServicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PartnerConfigPutVlanServicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPartnerConfigPutVlanServicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPartnerConfigPutVlanServicesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPartnerConfigPutVlanServicesOK creates a PartnerConfigPutVlanServicesOK with default headers values
func NewPartnerConfigPutVlanServicesOK() *PartnerConfigPutVlanServicesOK {
	return &PartnerConfigPutVlanServicesOK{}
}

/*
PartnerConfigPutVlanServicesOK describes a response with status code 200, with default header values.

Request was successful
*/
type PartnerConfigPutVlanServicesOK struct {
	Payload interface{}
}

// IsSuccess returns true when this partner config put vlan services o k response has a 2xx status code
func (o *PartnerConfigPutVlanServicesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this partner config put vlan services o k response has a 3xx status code
func (o *PartnerConfigPutVlanServicesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this partner config put vlan services o k response has a 4xx status code
func (o *PartnerConfigPutVlanServicesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this partner config put vlan services o k response has a 5xx status code
func (o *PartnerConfigPutVlanServicesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this partner config put vlan services o k response a status code equal to that given
func (o *PartnerConfigPutVlanServicesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the partner config put vlan services o k response
func (o *PartnerConfigPutVlanServicesOK) Code() int {
	return 200
}

func (o *PartnerConfigPutVlanServicesOK) Error() string {
	return fmt.Sprintf("[PUT /partnerConfig/{id}/platform/vlanServices][%d] partnerConfigPutVlanServicesOK  %+v", 200, o.Payload)
}

func (o *PartnerConfigPutVlanServicesOK) String() string {
	return fmt.Sprintf("[PUT /partnerConfig/{id}/platform/vlanServices][%d] partnerConfigPutVlanServicesOK  %+v", 200, o.Payload)
}

func (o *PartnerConfigPutVlanServicesOK) GetPayload() interface{} {
	return o.Payload
}

func (o *PartnerConfigPutVlanServicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPartnerConfigPutVlanServicesDefault creates a PartnerConfigPutVlanServicesDefault with default headers values
func NewPartnerConfigPutVlanServicesDefault(code int) *PartnerConfigPutVlanServicesDefault {
	return &PartnerConfigPutVlanServicesDefault{
		_statusCode: code,
	}
}

/*
PartnerConfigPutVlanServicesDefault describes a response with status code -1, with default header values.

unexpected error
*/
type PartnerConfigPutVlanServicesDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this partner config put vlan services default response has a 2xx status code
func (o *PartnerConfigPutVlanServicesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this partner config put vlan services default response has a 3xx status code
func (o *PartnerConfigPutVlanServicesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this partner config put vlan services default response has a 4xx status code
func (o *PartnerConfigPutVlanServicesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this partner config put vlan services default response has a 5xx status code
func (o *PartnerConfigPutVlanServicesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this partner config put vlan services default response a status code equal to that given
func (o *PartnerConfigPutVlanServicesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the partner config put vlan services default response
func (o *PartnerConfigPutVlanServicesDefault) Code() int {
	return o._statusCode
}

func (o *PartnerConfigPutVlanServicesDefault) Error() string {
	return fmt.Sprintf("[PUT /partnerConfig/{id}/platform/vlanServices][%d] PartnerConfig.putVlanServices default  %+v", o._statusCode, o.Payload)
}

func (o *PartnerConfigPutVlanServicesDefault) String() string {
	return fmt.Sprintf("[PUT /partnerConfig/{id}/platform/vlanServices][%d] PartnerConfig.putVlanServices default  %+v", o._statusCode, o.Payload)
}

func (o *PartnerConfigPutVlanServicesDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *PartnerConfigPutVlanServicesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
