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

// PartnerConfigGetMachineToMachineReader is a Reader for the PartnerConfigGetMachineToMachine structure.
type PartnerConfigGetMachineToMachineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PartnerConfigGetMachineToMachineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPartnerConfigGetMachineToMachineOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPartnerConfigGetMachineToMachineDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPartnerConfigGetMachineToMachineOK creates a PartnerConfigGetMachineToMachineOK with default headers values
func NewPartnerConfigGetMachineToMachineOK() *PartnerConfigGetMachineToMachineOK {
	return &PartnerConfigGetMachineToMachineOK{}
}

/*
PartnerConfigGetMachineToMachineOK describes a response with status code 200, with default header values.

Request was successful
*/
type PartnerConfigGetMachineToMachineOK struct {
	Payload models.XAny
}

// IsSuccess returns true when this partner config get machine to machine o k response has a 2xx status code
func (o *PartnerConfigGetMachineToMachineOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this partner config get machine to machine o k response has a 3xx status code
func (o *PartnerConfigGetMachineToMachineOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this partner config get machine to machine o k response has a 4xx status code
func (o *PartnerConfigGetMachineToMachineOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this partner config get machine to machine o k response has a 5xx status code
func (o *PartnerConfigGetMachineToMachineOK) IsServerError() bool {
	return false
}

// IsCode returns true when this partner config get machine to machine o k response a status code equal to that given
func (o *PartnerConfigGetMachineToMachineOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the partner config get machine to machine o k response
func (o *PartnerConfigGetMachineToMachineOK) Code() int {
	return 200
}

func (o *PartnerConfigGetMachineToMachineOK) Error() string {
	return fmt.Sprintf("[GET /partnerConfig/{id}/machineToMachine][%d] partnerConfigGetMachineToMachineOK  %+v", 200, o.Payload)
}

func (o *PartnerConfigGetMachineToMachineOK) String() string {
	return fmt.Sprintf("[GET /partnerConfig/{id}/machineToMachine][%d] partnerConfigGetMachineToMachineOK  %+v", 200, o.Payload)
}

func (o *PartnerConfigGetMachineToMachineOK) GetPayload() models.XAny {
	return o.Payload
}

func (o *PartnerConfigGetMachineToMachineOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPartnerConfigGetMachineToMachineDefault creates a PartnerConfigGetMachineToMachineDefault with default headers values
func NewPartnerConfigGetMachineToMachineDefault(code int) *PartnerConfigGetMachineToMachineDefault {
	return &PartnerConfigGetMachineToMachineDefault{
		_statusCode: code,
	}
}

/*
PartnerConfigGetMachineToMachineDefault describes a response with status code -1, with default header values.

unexpected error
*/
type PartnerConfigGetMachineToMachineDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this partner config get machine to machine default response has a 2xx status code
func (o *PartnerConfigGetMachineToMachineDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this partner config get machine to machine default response has a 3xx status code
func (o *PartnerConfigGetMachineToMachineDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this partner config get machine to machine default response has a 4xx status code
func (o *PartnerConfigGetMachineToMachineDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this partner config get machine to machine default response has a 5xx status code
func (o *PartnerConfigGetMachineToMachineDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this partner config get machine to machine default response a status code equal to that given
func (o *PartnerConfigGetMachineToMachineDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the partner config get machine to machine default response
func (o *PartnerConfigGetMachineToMachineDefault) Code() int {
	return o._statusCode
}

func (o *PartnerConfigGetMachineToMachineDefault) Error() string {
	return fmt.Sprintf("[GET /partnerConfig/{id}/machineToMachine][%d] PartnerConfig.getMachineToMachine default  %+v", o._statusCode, o.Payload)
}

func (o *PartnerConfigGetMachineToMachineDefault) String() string {
	return fmt.Sprintf("[GET /partnerConfig/{id}/machineToMachine][%d] PartnerConfig.getMachineToMachine default  %+v", o._statusCode, o.Payload)
}

func (o *PartnerConfigGetMachineToMachineDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *PartnerConfigGetMachineToMachineDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
