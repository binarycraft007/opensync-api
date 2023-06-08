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

// CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeReader is a Reader for the CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTime structure.
type CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeOK creates a CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeOK with default headers values
func NewCustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeOK() *CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeOK {
	return &CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeOK{}
}

/*
CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeOK struct {
	Payload []interface{}
}

// IsSuccess returns true when this customer prototype get employee network app time apps online time o k response has a 2xx status code
func (o *CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype get employee network app time apps online time o k response has a 3xx status code
func (o *CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype get employee network app time apps online time o k response has a 4xx status code
func (o *CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype get employee network app time apps online time o k response has a 5xx status code
func (o *CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype get employee network app time apps online time o k response a status code equal to that given
func (o *CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype get employee network app time apps online time o k response
func (o *CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeOK) Code() int {
	return 200
}

func (o *CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeOK) Error() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/secondaryNetworks/fronthauls/{networkId}/appTime/apps/onlineTime][%d] customerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeOK) String() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/secondaryNetworks/fronthauls/{networkId}/appTime/apps/onlineTime][%d] customerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeOK) GetPayload() []interface{} {
	return o.Payload
}

func (o *CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeDefault creates a CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeDefault with default headers values
func NewCustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeDefault(code int) *CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeDefault {
	return &CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype get employee network app time apps online time default response has a 2xx status code
func (o *CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype get employee network app time apps online time default response has a 3xx status code
func (o *CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype get employee network app time apps online time default response has a 4xx status code
func (o *CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype get employee network app time apps online time default response has a 5xx status code
func (o *CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype get employee network app time apps online time default response a status code equal to that given
func (o *CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype get employee network app time apps online time default response
func (o *CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeDefault) Error() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/secondaryNetworks/fronthauls/{networkId}/appTime/apps/onlineTime][%d] Customer.prototype.getEmployeeNetworkAppTimeAppsOnlineTime default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeDefault) String() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/secondaryNetworks/fronthauls/{networkId}/appTime/apps/onlineTime][%d] Customer.prototype.getEmployeeNetworkAppTimeAppsOnlineTime default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}