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

// CustomerPrototypeGetPersonAppTimeAppsOnlineTimeReader is a Reader for the CustomerPrototypeGetPersonAppTimeAppsOnlineTime structure.
type CustomerPrototypeGetPersonAppTimeAppsOnlineTimeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypeGetPersonAppTimeAppsOnlineTimeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypeGetPersonAppTimeAppsOnlineTimeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypeGetPersonAppTimeAppsOnlineTimeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypeGetPersonAppTimeAppsOnlineTimeOK creates a CustomerPrototypeGetPersonAppTimeAppsOnlineTimeOK with default headers values
func NewCustomerPrototypeGetPersonAppTimeAppsOnlineTimeOK() *CustomerPrototypeGetPersonAppTimeAppsOnlineTimeOK {
	return &CustomerPrototypeGetPersonAppTimeAppsOnlineTimeOK{}
}

/*
CustomerPrototypeGetPersonAppTimeAppsOnlineTimeOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypeGetPersonAppTimeAppsOnlineTimeOK struct {
	Payload []interface{}
}

// IsSuccess returns true when this customer prototype get person app time apps online time o k response has a 2xx status code
func (o *CustomerPrototypeGetPersonAppTimeAppsOnlineTimeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype get person app time apps online time o k response has a 3xx status code
func (o *CustomerPrototypeGetPersonAppTimeAppsOnlineTimeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype get person app time apps online time o k response has a 4xx status code
func (o *CustomerPrototypeGetPersonAppTimeAppsOnlineTimeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype get person app time apps online time o k response has a 5xx status code
func (o *CustomerPrototypeGetPersonAppTimeAppsOnlineTimeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype get person app time apps online time o k response a status code equal to that given
func (o *CustomerPrototypeGetPersonAppTimeAppsOnlineTimeOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype get person app time apps online time o k response
func (o *CustomerPrototypeGetPersonAppTimeAppsOnlineTimeOK) Code() int {
	return 200
}

func (o *CustomerPrototypeGetPersonAppTimeAppsOnlineTimeOK) Error() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/persons/{personId}/appTime/apps/onlineTime][%d] customerPrototypeGetPersonAppTimeAppsOnlineTimeOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeGetPersonAppTimeAppsOnlineTimeOK) String() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/persons/{personId}/appTime/apps/onlineTime][%d] customerPrototypeGetPersonAppTimeAppsOnlineTimeOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeGetPersonAppTimeAppsOnlineTimeOK) GetPayload() []interface{} {
	return o.Payload
}

func (o *CustomerPrototypeGetPersonAppTimeAppsOnlineTimeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypeGetPersonAppTimeAppsOnlineTimeDefault creates a CustomerPrototypeGetPersonAppTimeAppsOnlineTimeDefault with default headers values
func NewCustomerPrototypeGetPersonAppTimeAppsOnlineTimeDefault(code int) *CustomerPrototypeGetPersonAppTimeAppsOnlineTimeDefault {
	return &CustomerPrototypeGetPersonAppTimeAppsOnlineTimeDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypeGetPersonAppTimeAppsOnlineTimeDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypeGetPersonAppTimeAppsOnlineTimeDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype get person app time apps online time default response has a 2xx status code
func (o *CustomerPrototypeGetPersonAppTimeAppsOnlineTimeDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype get person app time apps online time default response has a 3xx status code
func (o *CustomerPrototypeGetPersonAppTimeAppsOnlineTimeDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype get person app time apps online time default response has a 4xx status code
func (o *CustomerPrototypeGetPersonAppTimeAppsOnlineTimeDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype get person app time apps online time default response has a 5xx status code
func (o *CustomerPrototypeGetPersonAppTimeAppsOnlineTimeDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype get person app time apps online time default response a status code equal to that given
func (o *CustomerPrototypeGetPersonAppTimeAppsOnlineTimeDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype get person app time apps online time default response
func (o *CustomerPrototypeGetPersonAppTimeAppsOnlineTimeDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypeGetPersonAppTimeAppsOnlineTimeDefault) Error() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/persons/{personId}/appTime/apps/onlineTime][%d] Customer.prototype.getPersonAppTimeAppsOnlineTime default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeGetPersonAppTimeAppsOnlineTimeDefault) String() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/persons/{personId}/appTime/apps/onlineTime][%d] Customer.prototype.getPersonAppTimeAppsOnlineTime default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeGetPersonAppTimeAppsOnlineTimeDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypeGetPersonAppTimeAppsOnlineTimeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
