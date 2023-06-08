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

// LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistReader is a Reader for the LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklist structure.
type LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistOK creates a LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistOK with default headers values
func NewLocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistOK() *LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistOK {
	return &LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistOK{}
}

/*
LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistOK describes a response with status code 200, with default header values.

Request was successful
*/
type LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistOK struct {
	Payload interface{}
}

// IsSuccess returns true when this location prototype post device security policy anomaly blacklist o k response has a 2xx status code
func (o *LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this location prototype post device security policy anomaly blacklist o k response has a 3xx status code
func (o *LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this location prototype post device security policy anomaly blacklist o k response has a 4xx status code
func (o *LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this location prototype post device security policy anomaly blacklist o k response has a 5xx status code
func (o *LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistOK) IsServerError() bool {
	return false
}

// IsCode returns true when this location prototype post device security policy anomaly blacklist o k response a status code equal to that given
func (o *LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the location prototype post device security policy anomaly blacklist o k response
func (o *LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistOK) Code() int {
	return 200
}

func (o *LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistOK) Error() string {
	return fmt.Sprintf("[POST /Locations/{id}/devices/{mac}/securityPolicy/anomaly/websites/blacklist][%d] locationPrototypePostDeviceSecurityPolicyAnomalyBlacklistOK  %+v", 200, o.Payload)
}

func (o *LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistOK) String() string {
	return fmt.Sprintf("[POST /Locations/{id}/devices/{mac}/securityPolicy/anomaly/websites/blacklist][%d] locationPrototypePostDeviceSecurityPolicyAnomalyBlacklistOK  %+v", 200, o.Payload)
}

func (o *LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistOK) GetPayload() interface{} {
	return o.Payload
}

func (o *LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistDefault creates a LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistDefault with default headers values
func NewLocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistDefault(code int) *LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistDefault {
	return &LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistDefault{
		_statusCode: code,
	}
}

/*
LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistDefault describes a response with status code -1, with default header values.

unexpected error
*/
type LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this location prototype post device security policy anomaly blacklist default response has a 2xx status code
func (o *LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this location prototype post device security policy anomaly blacklist default response has a 3xx status code
func (o *LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this location prototype post device security policy anomaly blacklist default response has a 4xx status code
func (o *LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this location prototype post device security policy anomaly blacklist default response has a 5xx status code
func (o *LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this location prototype post device security policy anomaly blacklist default response a status code equal to that given
func (o *LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the location prototype post device security policy anomaly blacklist default response
func (o *LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistDefault) Code() int {
	return o._statusCode
}

func (o *LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistDefault) Error() string {
	return fmt.Sprintf("[POST /Locations/{id}/devices/{mac}/securityPolicy/anomaly/websites/blacklist][%d] Location.prototype.postDeviceSecurityPolicyAnomalyBlacklist default  %+v", o._statusCode, o.Payload)
}

func (o *LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistDefault) String() string {
	return fmt.Sprintf("[POST /Locations/{id}/devices/{mac}/securityPolicy/anomaly/websites/blacklist][%d] Location.prototype.postDeviceSecurityPolicyAnomalyBlacklist default  %+v", o._statusCode, o.Payload)
}

func (o *LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *LocationPrototypePostDeviceSecurityPolicyAnomalyBlacklistDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
