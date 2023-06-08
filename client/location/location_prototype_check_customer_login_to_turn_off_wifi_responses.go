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

// LocationPrototypeCheckCustomerLoginToTurnOffWifiReader is a Reader for the LocationPrototypeCheckCustomerLoginToTurnOffWifi structure.
type LocationPrototypeCheckCustomerLoginToTurnOffWifiReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LocationPrototypeCheckCustomerLoginToTurnOffWifiReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLocationPrototypeCheckCustomerLoginToTurnOffWifiOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLocationPrototypeCheckCustomerLoginToTurnOffWifiDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLocationPrototypeCheckCustomerLoginToTurnOffWifiOK creates a LocationPrototypeCheckCustomerLoginToTurnOffWifiOK with default headers values
func NewLocationPrototypeCheckCustomerLoginToTurnOffWifiOK() *LocationPrototypeCheckCustomerLoginToTurnOffWifiOK {
	return &LocationPrototypeCheckCustomerLoginToTurnOffWifiOK{}
}

/*
LocationPrototypeCheckCustomerLoginToTurnOffWifiOK describes a response with status code 200, with default header values.

Request was successful
*/
type LocationPrototypeCheckCustomerLoginToTurnOffWifiOK struct {
	Payload interface{}
}

// IsSuccess returns true when this location prototype check customer login to turn off wifi o k response has a 2xx status code
func (o *LocationPrototypeCheckCustomerLoginToTurnOffWifiOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this location prototype check customer login to turn off wifi o k response has a 3xx status code
func (o *LocationPrototypeCheckCustomerLoginToTurnOffWifiOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this location prototype check customer login to turn off wifi o k response has a 4xx status code
func (o *LocationPrototypeCheckCustomerLoginToTurnOffWifiOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this location prototype check customer login to turn off wifi o k response has a 5xx status code
func (o *LocationPrototypeCheckCustomerLoginToTurnOffWifiOK) IsServerError() bool {
	return false
}

// IsCode returns true when this location prototype check customer login to turn off wifi o k response a status code equal to that given
func (o *LocationPrototypeCheckCustomerLoginToTurnOffWifiOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the location prototype check customer login to turn off wifi o k response
func (o *LocationPrototypeCheckCustomerLoginToTurnOffWifiOK) Code() int {
	return 200
}

func (o *LocationPrototypeCheckCustomerLoginToTurnOffWifiOK) Error() string {
	return fmt.Sprintf("[POST /Locations/{id}/checkCustomerLoginToTurnOffWifi][%d] locationPrototypeCheckCustomerLoginToTurnOffWifiOK  %+v", 200, o.Payload)
}

func (o *LocationPrototypeCheckCustomerLoginToTurnOffWifiOK) String() string {
	return fmt.Sprintf("[POST /Locations/{id}/checkCustomerLoginToTurnOffWifi][%d] locationPrototypeCheckCustomerLoginToTurnOffWifiOK  %+v", 200, o.Payload)
}

func (o *LocationPrototypeCheckCustomerLoginToTurnOffWifiOK) GetPayload() interface{} {
	return o.Payload
}

func (o *LocationPrototypeCheckCustomerLoginToTurnOffWifiOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLocationPrototypeCheckCustomerLoginToTurnOffWifiDefault creates a LocationPrototypeCheckCustomerLoginToTurnOffWifiDefault with default headers values
func NewLocationPrototypeCheckCustomerLoginToTurnOffWifiDefault(code int) *LocationPrototypeCheckCustomerLoginToTurnOffWifiDefault {
	return &LocationPrototypeCheckCustomerLoginToTurnOffWifiDefault{
		_statusCode: code,
	}
}

/*
LocationPrototypeCheckCustomerLoginToTurnOffWifiDefault describes a response with status code -1, with default header values.

unexpected error
*/
type LocationPrototypeCheckCustomerLoginToTurnOffWifiDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this location prototype check customer login to turn off wifi default response has a 2xx status code
func (o *LocationPrototypeCheckCustomerLoginToTurnOffWifiDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this location prototype check customer login to turn off wifi default response has a 3xx status code
func (o *LocationPrototypeCheckCustomerLoginToTurnOffWifiDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this location prototype check customer login to turn off wifi default response has a 4xx status code
func (o *LocationPrototypeCheckCustomerLoginToTurnOffWifiDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this location prototype check customer login to turn off wifi default response has a 5xx status code
func (o *LocationPrototypeCheckCustomerLoginToTurnOffWifiDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this location prototype check customer login to turn off wifi default response a status code equal to that given
func (o *LocationPrototypeCheckCustomerLoginToTurnOffWifiDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the location prototype check customer login to turn off wifi default response
func (o *LocationPrototypeCheckCustomerLoginToTurnOffWifiDefault) Code() int {
	return o._statusCode
}

func (o *LocationPrototypeCheckCustomerLoginToTurnOffWifiDefault) Error() string {
	return fmt.Sprintf("[POST /Locations/{id}/checkCustomerLoginToTurnOffWifi][%d] Location.prototype.checkCustomerLoginToTurnOffWifi default  %+v", o._statusCode, o.Payload)
}

func (o *LocationPrototypeCheckCustomerLoginToTurnOffWifiDefault) String() string {
	return fmt.Sprintf("[POST /Locations/{id}/checkCustomerLoginToTurnOffWifi][%d] Location.prototype.checkCustomerLoginToTurnOffWifi default  %+v", o._statusCode, o.Payload)
}

func (o *LocationPrototypeCheckCustomerLoginToTurnOffWifiDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *LocationPrototypeCheckCustomerLoginToTurnOffWifiDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
