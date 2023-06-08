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

// CustomerPrototypeGetWifiMotionReader is a Reader for the CustomerPrototypeGetWifiMotion structure.
type CustomerPrototypeGetWifiMotionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypeGetWifiMotionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypeGetWifiMotionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypeGetWifiMotionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypeGetWifiMotionOK creates a CustomerPrototypeGetWifiMotionOK with default headers values
func NewCustomerPrototypeGetWifiMotionOK() *CustomerPrototypeGetWifiMotionOK {
	return &CustomerPrototypeGetWifiMotionOK{}
}

/*
CustomerPrototypeGetWifiMotionOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypeGetWifiMotionOK struct {
	Payload *models.WifiMotion
}

// IsSuccess returns true when this customer prototype get wifi motion o k response has a 2xx status code
func (o *CustomerPrototypeGetWifiMotionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype get wifi motion o k response has a 3xx status code
func (o *CustomerPrototypeGetWifiMotionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype get wifi motion o k response has a 4xx status code
func (o *CustomerPrototypeGetWifiMotionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype get wifi motion o k response has a 5xx status code
func (o *CustomerPrototypeGetWifiMotionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype get wifi motion o k response a status code equal to that given
func (o *CustomerPrototypeGetWifiMotionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype get wifi motion o k response
func (o *CustomerPrototypeGetWifiMotionOK) Code() int {
	return 200
}

func (o *CustomerPrototypeGetWifiMotionOK) Error() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/wifiMotion][%d] customerPrototypeGetWifiMotionOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeGetWifiMotionOK) String() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/wifiMotion][%d] customerPrototypeGetWifiMotionOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeGetWifiMotionOK) GetPayload() *models.WifiMotion {
	return o.Payload
}

func (o *CustomerPrototypeGetWifiMotionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WifiMotion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypeGetWifiMotionDefault creates a CustomerPrototypeGetWifiMotionDefault with default headers values
func NewCustomerPrototypeGetWifiMotionDefault(code int) *CustomerPrototypeGetWifiMotionDefault {
	return &CustomerPrototypeGetWifiMotionDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypeGetWifiMotionDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypeGetWifiMotionDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype get wifi motion default response has a 2xx status code
func (o *CustomerPrototypeGetWifiMotionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype get wifi motion default response has a 3xx status code
func (o *CustomerPrototypeGetWifiMotionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype get wifi motion default response has a 4xx status code
func (o *CustomerPrototypeGetWifiMotionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype get wifi motion default response has a 5xx status code
func (o *CustomerPrototypeGetWifiMotionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype get wifi motion default response a status code equal to that given
func (o *CustomerPrototypeGetWifiMotionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype get wifi motion default response
func (o *CustomerPrototypeGetWifiMotionDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypeGetWifiMotionDefault) Error() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/wifiMotion][%d] Customer.prototype.getWifiMotion default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeGetWifiMotionDefault) String() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/wifiMotion][%d] Customer.prototype.getWifiMotion default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeGetWifiMotionDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypeGetWifiMotionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
