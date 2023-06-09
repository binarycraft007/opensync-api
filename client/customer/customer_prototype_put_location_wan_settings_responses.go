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

// CustomerPrototypePutLocationWanSettingsReader is a Reader for the CustomerPrototypePutLocationWanSettings structure.
type CustomerPrototypePutLocationWanSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypePutLocationWanSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypePutLocationWanSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypePutLocationWanSettingsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypePutLocationWanSettingsOK creates a CustomerPrototypePutLocationWanSettingsOK with default headers values
func NewCustomerPrototypePutLocationWanSettingsOK() *CustomerPrototypePutLocationWanSettingsOK {
	return &CustomerPrototypePutLocationWanSettingsOK{}
}

/*
CustomerPrototypePutLocationWanSettingsOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypePutLocationWanSettingsOK struct {
	Payload *models.LocationWanSettings
}

// IsSuccess returns true when this customer prototype put location wan settings o k response has a 2xx status code
func (o *CustomerPrototypePutLocationWanSettingsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype put location wan settings o k response has a 3xx status code
func (o *CustomerPrototypePutLocationWanSettingsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype put location wan settings o k response has a 4xx status code
func (o *CustomerPrototypePutLocationWanSettingsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype put location wan settings o k response has a 5xx status code
func (o *CustomerPrototypePutLocationWanSettingsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype put location wan settings o k response a status code equal to that given
func (o *CustomerPrototypePutLocationWanSettingsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype put location wan settings o k response
func (o *CustomerPrototypePutLocationWanSettingsOK) Code() int {
	return 200
}

func (o *CustomerPrototypePutLocationWanSettingsOK) Error() string {
	return fmt.Sprintf("[PUT /Customers/{id}/locations/{locationId}/wanSettings][%d] customerPrototypePutLocationWanSettingsOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypePutLocationWanSettingsOK) String() string {
	return fmt.Sprintf("[PUT /Customers/{id}/locations/{locationId}/wanSettings][%d] customerPrototypePutLocationWanSettingsOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypePutLocationWanSettingsOK) GetPayload() *models.LocationWanSettings {
	return o.Payload
}

func (o *CustomerPrototypePutLocationWanSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LocationWanSettings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypePutLocationWanSettingsDefault creates a CustomerPrototypePutLocationWanSettingsDefault with default headers values
func NewCustomerPrototypePutLocationWanSettingsDefault(code int) *CustomerPrototypePutLocationWanSettingsDefault {
	return &CustomerPrototypePutLocationWanSettingsDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypePutLocationWanSettingsDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypePutLocationWanSettingsDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype put location wan settings default response has a 2xx status code
func (o *CustomerPrototypePutLocationWanSettingsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype put location wan settings default response has a 3xx status code
func (o *CustomerPrototypePutLocationWanSettingsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype put location wan settings default response has a 4xx status code
func (o *CustomerPrototypePutLocationWanSettingsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype put location wan settings default response has a 5xx status code
func (o *CustomerPrototypePutLocationWanSettingsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype put location wan settings default response a status code equal to that given
func (o *CustomerPrototypePutLocationWanSettingsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype put location wan settings default response
func (o *CustomerPrototypePutLocationWanSettingsDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypePutLocationWanSettingsDefault) Error() string {
	return fmt.Sprintf("[PUT /Customers/{id}/locations/{locationId}/wanSettings][%d] Customer.prototype.putLocationWanSettings default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePutLocationWanSettingsDefault) String() string {
	return fmt.Sprintf("[PUT /Customers/{id}/locations/{locationId}/wanSettings][%d] Customer.prototype.putLocationWanSettings default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePutLocationWanSettingsDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypePutLocationWanSettingsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
