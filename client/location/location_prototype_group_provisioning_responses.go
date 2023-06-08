// Code generated by go-swagger; DO NOT EDIT.

package location

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/binarycraft007/opensync-api/models"
)

// LocationPrototypeGroupProvisioningReader is a Reader for the LocationPrototypeGroupProvisioning structure.
type LocationPrototypeGroupProvisioningReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LocationPrototypeGroupProvisioningReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLocationPrototypeGroupProvisioningOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLocationPrototypeGroupProvisioningDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLocationPrototypeGroupProvisioningOK creates a LocationPrototypeGroupProvisioningOK with default headers values
func NewLocationPrototypeGroupProvisioningOK() *LocationPrototypeGroupProvisioningOK {
	return &LocationPrototypeGroupProvisioningOK{}
}

/*
LocationPrototypeGroupProvisioningOK describes a response with status code 200, with default header values.

Request was successful
*/
type LocationPrototypeGroupProvisioningOK struct {
	Payload interface{}
}

// IsSuccess returns true when this location prototype group provisioning o k response has a 2xx status code
func (o *LocationPrototypeGroupProvisioningOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this location prototype group provisioning o k response has a 3xx status code
func (o *LocationPrototypeGroupProvisioningOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this location prototype group provisioning o k response has a 4xx status code
func (o *LocationPrototypeGroupProvisioningOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this location prototype group provisioning o k response has a 5xx status code
func (o *LocationPrototypeGroupProvisioningOK) IsServerError() bool {
	return false
}

// IsCode returns true when this location prototype group provisioning o k response a status code equal to that given
func (o *LocationPrototypeGroupProvisioningOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the location prototype group provisioning o k response
func (o *LocationPrototypeGroupProvisioningOK) Code() int {
	return 200
}

func (o *LocationPrototypeGroupProvisioningOK) Error() string {
	return fmt.Sprintf("[GET /Locations/{id}/groupProvisioning][%d] locationPrototypeGroupProvisioningOK  %+v", 200, o.Payload)
}

func (o *LocationPrototypeGroupProvisioningOK) String() string {
	return fmt.Sprintf("[GET /Locations/{id}/groupProvisioning][%d] locationPrototypeGroupProvisioningOK  %+v", 200, o.Payload)
}

func (o *LocationPrototypeGroupProvisioningOK) GetPayload() interface{} {
	return o.Payload
}

func (o *LocationPrototypeGroupProvisioningOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLocationPrototypeGroupProvisioningDefault creates a LocationPrototypeGroupProvisioningDefault with default headers values
func NewLocationPrototypeGroupProvisioningDefault(code int) *LocationPrototypeGroupProvisioningDefault {
	return &LocationPrototypeGroupProvisioningDefault{
		_statusCode: code,
	}
}

/*
LocationPrototypeGroupProvisioningDefault describes a response with status code -1, with default header values.

unexpected error
*/
type LocationPrototypeGroupProvisioningDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this location prototype group provisioning default response has a 2xx status code
func (o *LocationPrototypeGroupProvisioningDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this location prototype group provisioning default response has a 3xx status code
func (o *LocationPrototypeGroupProvisioningDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this location prototype group provisioning default response has a 4xx status code
func (o *LocationPrototypeGroupProvisioningDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this location prototype group provisioning default response has a 5xx status code
func (o *LocationPrototypeGroupProvisioningDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this location prototype group provisioning default response a status code equal to that given
func (o *LocationPrototypeGroupProvisioningDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the location prototype group provisioning default response
func (o *LocationPrototypeGroupProvisioningDefault) Code() int {
	return o._statusCode
}

func (o *LocationPrototypeGroupProvisioningDefault) Error() string {
	return fmt.Sprintf("[GET /Locations/{id}/groupProvisioning][%d] Location.prototype.groupProvisioning default  %+v", o._statusCode, o.Payload)
}

func (o *LocationPrototypeGroupProvisioningDefault) String() string {
	return fmt.Sprintf("[GET /Locations/{id}/groupProvisioning][%d] Location.prototype.groupProvisioning default  %+v", o._statusCode, o.Payload)
}

func (o *LocationPrototypeGroupProvisioningDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *LocationPrototypeGroupProvisioningDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
