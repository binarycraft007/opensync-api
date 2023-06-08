// Code generated by go-swagger; DO NOT EDIT.

package customer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/binarycraft007/opensync-api/models"
)

// CustomerPrototypeGetDeviceSecurityReader is a Reader for the CustomerPrototypeGetDeviceSecurity structure.
type CustomerPrototypeGetDeviceSecurityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypeGetDeviceSecurityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypeGetDeviceSecurityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypeGetDeviceSecurityDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypeGetDeviceSecurityOK creates a CustomerPrototypeGetDeviceSecurityOK with default headers values
func NewCustomerPrototypeGetDeviceSecurityOK() *CustomerPrototypeGetDeviceSecurityOK {
	return &CustomerPrototypeGetDeviceSecurityOK{}
}

/*
CustomerPrototypeGetDeviceSecurityOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypeGetDeviceSecurityOK struct {
	Payload *models.DeviceResponse
}

// IsSuccess returns true when this customer prototype get device security o k response has a 2xx status code
func (o *CustomerPrototypeGetDeviceSecurityOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype get device security o k response has a 3xx status code
func (o *CustomerPrototypeGetDeviceSecurityOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype get device security o k response has a 4xx status code
func (o *CustomerPrototypeGetDeviceSecurityOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype get device security o k response has a 5xx status code
func (o *CustomerPrototypeGetDeviceSecurityOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype get device security o k response a status code equal to that given
func (o *CustomerPrototypeGetDeviceSecurityOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype get device security o k response
func (o *CustomerPrototypeGetDeviceSecurityOK) Code() int {
	return 200
}

func (o *CustomerPrototypeGetDeviceSecurityOK) Error() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/devices/{mac}/securityPolicy][%d] customerPrototypeGetDeviceSecurityOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeGetDeviceSecurityOK) String() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/devices/{mac}/securityPolicy][%d] customerPrototypeGetDeviceSecurityOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeGetDeviceSecurityOK) GetPayload() *models.DeviceResponse {
	return o.Payload
}

func (o *CustomerPrototypeGetDeviceSecurityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypeGetDeviceSecurityDefault creates a CustomerPrototypeGetDeviceSecurityDefault with default headers values
func NewCustomerPrototypeGetDeviceSecurityDefault(code int) *CustomerPrototypeGetDeviceSecurityDefault {
	return &CustomerPrototypeGetDeviceSecurityDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypeGetDeviceSecurityDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypeGetDeviceSecurityDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype get device security default response has a 2xx status code
func (o *CustomerPrototypeGetDeviceSecurityDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype get device security default response has a 3xx status code
func (o *CustomerPrototypeGetDeviceSecurityDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype get device security default response has a 4xx status code
func (o *CustomerPrototypeGetDeviceSecurityDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype get device security default response has a 5xx status code
func (o *CustomerPrototypeGetDeviceSecurityDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype get device security default response a status code equal to that given
func (o *CustomerPrototypeGetDeviceSecurityDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype get device security default response
func (o *CustomerPrototypeGetDeviceSecurityDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypeGetDeviceSecurityDefault) Error() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/devices/{mac}/securityPolicy][%d] Customer.prototype.getDeviceSecurity default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeGetDeviceSecurityDefault) String() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/devices/{mac}/securityPolicy][%d] Customer.prototype.getDeviceSecurity default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeGetDeviceSecurityDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypeGetDeviceSecurityDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
