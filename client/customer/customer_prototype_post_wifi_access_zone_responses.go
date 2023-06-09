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

// CustomerPrototypePostWifiAccessZoneReader is a Reader for the CustomerPrototypePostWifiAccessZone structure.
type CustomerPrototypePostWifiAccessZoneReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypePostWifiAccessZoneReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypePostWifiAccessZoneOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypePostWifiAccessZoneDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypePostWifiAccessZoneOK creates a CustomerPrototypePostWifiAccessZoneOK with default headers values
func NewCustomerPrototypePostWifiAccessZoneOK() *CustomerPrototypePostWifiAccessZoneOK {
	return &CustomerPrototypePostWifiAccessZoneOK{}
}

/*
CustomerPrototypePostWifiAccessZoneOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypePostWifiAccessZoneOK struct {
	Payload []*models.WifiAccessZone
}

// IsSuccess returns true when this customer prototype post wifi access zone o k response has a 2xx status code
func (o *CustomerPrototypePostWifiAccessZoneOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype post wifi access zone o k response has a 3xx status code
func (o *CustomerPrototypePostWifiAccessZoneOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype post wifi access zone o k response has a 4xx status code
func (o *CustomerPrototypePostWifiAccessZoneOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype post wifi access zone o k response has a 5xx status code
func (o *CustomerPrototypePostWifiAccessZoneOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype post wifi access zone o k response a status code equal to that given
func (o *CustomerPrototypePostWifiAccessZoneOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype post wifi access zone o k response
func (o *CustomerPrototypePostWifiAccessZoneOK) Code() int {
	return 200
}

func (o *CustomerPrototypePostWifiAccessZoneOK) Error() string {
	return fmt.Sprintf("[POST /Customers/{id}/locations/{locationId}/wifiNetwork/accessZones][%d] customerPrototypePostWifiAccessZoneOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypePostWifiAccessZoneOK) String() string {
	return fmt.Sprintf("[POST /Customers/{id}/locations/{locationId}/wifiNetwork/accessZones][%d] customerPrototypePostWifiAccessZoneOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypePostWifiAccessZoneOK) GetPayload() []*models.WifiAccessZone {
	return o.Payload
}

func (o *CustomerPrototypePostWifiAccessZoneOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypePostWifiAccessZoneDefault creates a CustomerPrototypePostWifiAccessZoneDefault with default headers values
func NewCustomerPrototypePostWifiAccessZoneDefault(code int) *CustomerPrototypePostWifiAccessZoneDefault {
	return &CustomerPrototypePostWifiAccessZoneDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypePostWifiAccessZoneDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypePostWifiAccessZoneDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype post wifi access zone default response has a 2xx status code
func (o *CustomerPrototypePostWifiAccessZoneDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype post wifi access zone default response has a 3xx status code
func (o *CustomerPrototypePostWifiAccessZoneDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype post wifi access zone default response has a 4xx status code
func (o *CustomerPrototypePostWifiAccessZoneDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype post wifi access zone default response has a 5xx status code
func (o *CustomerPrototypePostWifiAccessZoneDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype post wifi access zone default response a status code equal to that given
func (o *CustomerPrototypePostWifiAccessZoneDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype post wifi access zone default response
func (o *CustomerPrototypePostWifiAccessZoneDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypePostWifiAccessZoneDefault) Error() string {
	return fmt.Sprintf("[POST /Customers/{id}/locations/{locationId}/wifiNetwork/accessZones][%d] Customer.prototype.postWifiAccessZone default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePostWifiAccessZoneDefault) String() string {
	return fmt.Sprintf("[POST /Customers/{id}/locations/{locationId}/wifiNetwork/accessZones][%d] Customer.prototype.postWifiAccessZone default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePostWifiAccessZoneDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypePostWifiAccessZoneDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
