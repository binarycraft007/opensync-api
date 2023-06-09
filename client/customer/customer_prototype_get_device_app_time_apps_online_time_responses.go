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

// CustomerPrototypeGetDeviceAppTimeAppsOnlineTimeReader is a Reader for the CustomerPrototypeGetDeviceAppTimeAppsOnlineTime structure.
type CustomerPrototypeGetDeviceAppTimeAppsOnlineTimeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypeGetDeviceAppTimeAppsOnlineTimeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypeGetDeviceAppTimeAppsOnlineTimeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypeGetDeviceAppTimeAppsOnlineTimeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypeGetDeviceAppTimeAppsOnlineTimeOK creates a CustomerPrototypeGetDeviceAppTimeAppsOnlineTimeOK with default headers values
func NewCustomerPrototypeGetDeviceAppTimeAppsOnlineTimeOK() *CustomerPrototypeGetDeviceAppTimeAppsOnlineTimeOK {
	return &CustomerPrototypeGetDeviceAppTimeAppsOnlineTimeOK{}
}

/*
CustomerPrototypeGetDeviceAppTimeAppsOnlineTimeOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypeGetDeviceAppTimeAppsOnlineTimeOK struct {
	Payload []interface{}
}

// IsSuccess returns true when this customer prototype get device app time apps online time o k response has a 2xx status code
func (o *CustomerPrototypeGetDeviceAppTimeAppsOnlineTimeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype get device app time apps online time o k response has a 3xx status code
func (o *CustomerPrototypeGetDeviceAppTimeAppsOnlineTimeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype get device app time apps online time o k response has a 4xx status code
func (o *CustomerPrototypeGetDeviceAppTimeAppsOnlineTimeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype get device app time apps online time o k response has a 5xx status code
func (o *CustomerPrototypeGetDeviceAppTimeAppsOnlineTimeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype get device app time apps online time o k response a status code equal to that given
func (o *CustomerPrototypeGetDeviceAppTimeAppsOnlineTimeOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype get device app time apps online time o k response
func (o *CustomerPrototypeGetDeviceAppTimeAppsOnlineTimeOK) Code() int {
	return 200
}

func (o *CustomerPrototypeGetDeviceAppTimeAppsOnlineTimeOK) Error() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/devices/{mac}/appTime/apps/onlineTime][%d] customerPrototypeGetDeviceAppTimeAppsOnlineTimeOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeGetDeviceAppTimeAppsOnlineTimeOK) String() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/devices/{mac}/appTime/apps/onlineTime][%d] customerPrototypeGetDeviceAppTimeAppsOnlineTimeOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeGetDeviceAppTimeAppsOnlineTimeOK) GetPayload() []interface{} {
	return o.Payload
}

func (o *CustomerPrototypeGetDeviceAppTimeAppsOnlineTimeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypeGetDeviceAppTimeAppsOnlineTimeDefault creates a CustomerPrototypeGetDeviceAppTimeAppsOnlineTimeDefault with default headers values
func NewCustomerPrototypeGetDeviceAppTimeAppsOnlineTimeDefault(code int) *CustomerPrototypeGetDeviceAppTimeAppsOnlineTimeDefault {
	return &CustomerPrototypeGetDeviceAppTimeAppsOnlineTimeDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypeGetDeviceAppTimeAppsOnlineTimeDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypeGetDeviceAppTimeAppsOnlineTimeDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype get device app time apps online time default response has a 2xx status code
func (o *CustomerPrototypeGetDeviceAppTimeAppsOnlineTimeDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype get device app time apps online time default response has a 3xx status code
func (o *CustomerPrototypeGetDeviceAppTimeAppsOnlineTimeDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype get device app time apps online time default response has a 4xx status code
func (o *CustomerPrototypeGetDeviceAppTimeAppsOnlineTimeDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype get device app time apps online time default response has a 5xx status code
func (o *CustomerPrototypeGetDeviceAppTimeAppsOnlineTimeDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype get device app time apps online time default response a status code equal to that given
func (o *CustomerPrototypeGetDeviceAppTimeAppsOnlineTimeDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype get device app time apps online time default response
func (o *CustomerPrototypeGetDeviceAppTimeAppsOnlineTimeDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypeGetDeviceAppTimeAppsOnlineTimeDefault) Error() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/devices/{mac}/appTime/apps/onlineTime][%d] Customer.prototype.getDeviceAppTimeAppsOnlineTime default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeGetDeviceAppTimeAppsOnlineTimeDefault) String() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/devices/{mac}/appTime/apps/onlineTime][%d] Customer.prototype.getDeviceAppTimeAppsOnlineTime default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeGetDeviceAppTimeAppsOnlineTimeDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypeGetDeviceAppTimeAppsOnlineTimeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
