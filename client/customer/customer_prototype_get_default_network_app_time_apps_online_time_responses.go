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

// CustomerPrototypeGetDefaultNetworkAppTimeAppsOnlineTimeReader is a Reader for the CustomerPrototypeGetDefaultNetworkAppTimeAppsOnlineTime structure.
type CustomerPrototypeGetDefaultNetworkAppTimeAppsOnlineTimeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypeGetDefaultNetworkAppTimeAppsOnlineTimeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypeGetDefaultNetworkAppTimeAppsOnlineTimeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypeGetDefaultNetworkAppTimeAppsOnlineTimeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypeGetDefaultNetworkAppTimeAppsOnlineTimeOK creates a CustomerPrototypeGetDefaultNetworkAppTimeAppsOnlineTimeOK with default headers values
func NewCustomerPrototypeGetDefaultNetworkAppTimeAppsOnlineTimeOK() *CustomerPrototypeGetDefaultNetworkAppTimeAppsOnlineTimeOK {
	return &CustomerPrototypeGetDefaultNetworkAppTimeAppsOnlineTimeOK{}
}

/*
CustomerPrototypeGetDefaultNetworkAppTimeAppsOnlineTimeOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypeGetDefaultNetworkAppTimeAppsOnlineTimeOK struct {
	Payload []interface{}
}

// IsSuccess returns true when this customer prototype get default network app time apps online time o k response has a 2xx status code
func (o *CustomerPrototypeGetDefaultNetworkAppTimeAppsOnlineTimeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype get default network app time apps online time o k response has a 3xx status code
func (o *CustomerPrototypeGetDefaultNetworkAppTimeAppsOnlineTimeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype get default network app time apps online time o k response has a 4xx status code
func (o *CustomerPrototypeGetDefaultNetworkAppTimeAppsOnlineTimeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype get default network app time apps online time o k response has a 5xx status code
func (o *CustomerPrototypeGetDefaultNetworkAppTimeAppsOnlineTimeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype get default network app time apps online time o k response a status code equal to that given
func (o *CustomerPrototypeGetDefaultNetworkAppTimeAppsOnlineTimeOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype get default network app time apps online time o k response
func (o *CustomerPrototypeGetDefaultNetworkAppTimeAppsOnlineTimeOK) Code() int {
	return 200
}

func (o *CustomerPrototypeGetDefaultNetworkAppTimeAppsOnlineTimeOK) Error() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/wifiNetwork/appTime/apps/onlineTime][%d] customerPrototypeGetDefaultNetworkAppTimeAppsOnlineTimeOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeGetDefaultNetworkAppTimeAppsOnlineTimeOK) String() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/wifiNetwork/appTime/apps/onlineTime][%d] customerPrototypeGetDefaultNetworkAppTimeAppsOnlineTimeOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeGetDefaultNetworkAppTimeAppsOnlineTimeOK) GetPayload() []interface{} {
	return o.Payload
}

func (o *CustomerPrototypeGetDefaultNetworkAppTimeAppsOnlineTimeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypeGetDefaultNetworkAppTimeAppsOnlineTimeDefault creates a CustomerPrototypeGetDefaultNetworkAppTimeAppsOnlineTimeDefault with default headers values
func NewCustomerPrototypeGetDefaultNetworkAppTimeAppsOnlineTimeDefault(code int) *CustomerPrototypeGetDefaultNetworkAppTimeAppsOnlineTimeDefault {
	return &CustomerPrototypeGetDefaultNetworkAppTimeAppsOnlineTimeDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypeGetDefaultNetworkAppTimeAppsOnlineTimeDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypeGetDefaultNetworkAppTimeAppsOnlineTimeDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype get default network app time apps online time default response has a 2xx status code
func (o *CustomerPrototypeGetDefaultNetworkAppTimeAppsOnlineTimeDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype get default network app time apps online time default response has a 3xx status code
func (o *CustomerPrototypeGetDefaultNetworkAppTimeAppsOnlineTimeDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype get default network app time apps online time default response has a 4xx status code
func (o *CustomerPrototypeGetDefaultNetworkAppTimeAppsOnlineTimeDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype get default network app time apps online time default response has a 5xx status code
func (o *CustomerPrototypeGetDefaultNetworkAppTimeAppsOnlineTimeDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype get default network app time apps online time default response a status code equal to that given
func (o *CustomerPrototypeGetDefaultNetworkAppTimeAppsOnlineTimeDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype get default network app time apps online time default response
func (o *CustomerPrototypeGetDefaultNetworkAppTimeAppsOnlineTimeDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypeGetDefaultNetworkAppTimeAppsOnlineTimeDefault) Error() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/wifiNetwork/appTime/apps/onlineTime][%d] Customer.prototype.getDefaultNetworkAppTimeAppsOnlineTime default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeGetDefaultNetworkAppTimeAppsOnlineTimeDefault) String() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/wifiNetwork/appTime/apps/onlineTime][%d] Customer.prototype.getDefaultNetworkAppTimeAppsOnlineTime default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeGetDefaultNetworkAppTimeAppsOnlineTimeDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypeGetDefaultNetworkAppTimeAppsOnlineTimeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}