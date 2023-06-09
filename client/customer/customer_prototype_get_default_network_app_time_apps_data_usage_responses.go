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

// CustomerPrototypeGetDefaultNetworkAppTimeAppsDataUsageReader is a Reader for the CustomerPrototypeGetDefaultNetworkAppTimeAppsDataUsage structure.
type CustomerPrototypeGetDefaultNetworkAppTimeAppsDataUsageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypeGetDefaultNetworkAppTimeAppsDataUsageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypeGetDefaultNetworkAppTimeAppsDataUsageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypeGetDefaultNetworkAppTimeAppsDataUsageDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypeGetDefaultNetworkAppTimeAppsDataUsageOK creates a CustomerPrototypeGetDefaultNetworkAppTimeAppsDataUsageOK with default headers values
func NewCustomerPrototypeGetDefaultNetworkAppTimeAppsDataUsageOK() *CustomerPrototypeGetDefaultNetworkAppTimeAppsDataUsageOK {
	return &CustomerPrototypeGetDefaultNetworkAppTimeAppsDataUsageOK{}
}

/*
CustomerPrototypeGetDefaultNetworkAppTimeAppsDataUsageOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypeGetDefaultNetworkAppTimeAppsDataUsageOK struct {
	Payload []interface{}
}

// IsSuccess returns true when this customer prototype get default network app time apps data usage o k response has a 2xx status code
func (o *CustomerPrototypeGetDefaultNetworkAppTimeAppsDataUsageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype get default network app time apps data usage o k response has a 3xx status code
func (o *CustomerPrototypeGetDefaultNetworkAppTimeAppsDataUsageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype get default network app time apps data usage o k response has a 4xx status code
func (o *CustomerPrototypeGetDefaultNetworkAppTimeAppsDataUsageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype get default network app time apps data usage o k response has a 5xx status code
func (o *CustomerPrototypeGetDefaultNetworkAppTimeAppsDataUsageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype get default network app time apps data usage o k response a status code equal to that given
func (o *CustomerPrototypeGetDefaultNetworkAppTimeAppsDataUsageOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype get default network app time apps data usage o k response
func (o *CustomerPrototypeGetDefaultNetworkAppTimeAppsDataUsageOK) Code() int {
	return 200
}

func (o *CustomerPrototypeGetDefaultNetworkAppTimeAppsDataUsageOK) Error() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/wifiNetwork/appTime/apps/dataUsage][%d] customerPrototypeGetDefaultNetworkAppTimeAppsDataUsageOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeGetDefaultNetworkAppTimeAppsDataUsageOK) String() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/wifiNetwork/appTime/apps/dataUsage][%d] customerPrototypeGetDefaultNetworkAppTimeAppsDataUsageOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeGetDefaultNetworkAppTimeAppsDataUsageOK) GetPayload() []interface{} {
	return o.Payload
}

func (o *CustomerPrototypeGetDefaultNetworkAppTimeAppsDataUsageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypeGetDefaultNetworkAppTimeAppsDataUsageDefault creates a CustomerPrototypeGetDefaultNetworkAppTimeAppsDataUsageDefault with default headers values
func NewCustomerPrototypeGetDefaultNetworkAppTimeAppsDataUsageDefault(code int) *CustomerPrototypeGetDefaultNetworkAppTimeAppsDataUsageDefault {
	return &CustomerPrototypeGetDefaultNetworkAppTimeAppsDataUsageDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypeGetDefaultNetworkAppTimeAppsDataUsageDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypeGetDefaultNetworkAppTimeAppsDataUsageDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype get default network app time apps data usage default response has a 2xx status code
func (o *CustomerPrototypeGetDefaultNetworkAppTimeAppsDataUsageDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype get default network app time apps data usage default response has a 3xx status code
func (o *CustomerPrototypeGetDefaultNetworkAppTimeAppsDataUsageDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype get default network app time apps data usage default response has a 4xx status code
func (o *CustomerPrototypeGetDefaultNetworkAppTimeAppsDataUsageDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype get default network app time apps data usage default response has a 5xx status code
func (o *CustomerPrototypeGetDefaultNetworkAppTimeAppsDataUsageDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype get default network app time apps data usage default response a status code equal to that given
func (o *CustomerPrototypeGetDefaultNetworkAppTimeAppsDataUsageDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype get default network app time apps data usage default response
func (o *CustomerPrototypeGetDefaultNetworkAppTimeAppsDataUsageDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypeGetDefaultNetworkAppTimeAppsDataUsageDefault) Error() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/wifiNetwork/appTime/apps/dataUsage][%d] Customer.prototype.getDefaultNetworkAppTimeAppsDataUsage default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeGetDefaultNetworkAppTimeAppsDataUsageDefault) String() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/wifiNetwork/appTime/apps/dataUsage][%d] Customer.prototype.getDefaultNetworkAppTimeAppsDataUsage default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeGetDefaultNetworkAppTimeAppsDataUsageDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypeGetDefaultNetworkAppTimeAppsDataUsageDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
