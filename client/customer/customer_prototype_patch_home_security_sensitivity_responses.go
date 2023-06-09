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

// CustomerPrototypePatchHomeSecuritySensitivityReader is a Reader for the CustomerPrototypePatchHomeSecuritySensitivity structure.
type CustomerPrototypePatchHomeSecuritySensitivityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypePatchHomeSecuritySensitivityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypePatchHomeSecuritySensitivityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypePatchHomeSecuritySensitivityDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypePatchHomeSecuritySensitivityOK creates a CustomerPrototypePatchHomeSecuritySensitivityOK with default headers values
func NewCustomerPrototypePatchHomeSecuritySensitivityOK() *CustomerPrototypePatchHomeSecuritySensitivityOK {
	return &CustomerPrototypePatchHomeSecuritySensitivityOK{}
}

/*
CustomerPrototypePatchHomeSecuritySensitivityOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypePatchHomeSecuritySensitivityOK struct {
	Payload *models.HomeSecurity
}

// IsSuccess returns true when this customer prototype patch home security sensitivity o k response has a 2xx status code
func (o *CustomerPrototypePatchHomeSecuritySensitivityOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype patch home security sensitivity o k response has a 3xx status code
func (o *CustomerPrototypePatchHomeSecuritySensitivityOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype patch home security sensitivity o k response has a 4xx status code
func (o *CustomerPrototypePatchHomeSecuritySensitivityOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype patch home security sensitivity o k response has a 5xx status code
func (o *CustomerPrototypePatchHomeSecuritySensitivityOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype patch home security sensitivity o k response a status code equal to that given
func (o *CustomerPrototypePatchHomeSecuritySensitivityOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype patch home security sensitivity o k response
func (o *CustomerPrototypePatchHomeSecuritySensitivityOK) Code() int {
	return 200
}

func (o *CustomerPrototypePatchHomeSecuritySensitivityOK) Error() string {
	return fmt.Sprintf("[PATCH /Customers/{id}/locations/{locationId}/homeSecurity/sensitivity][%d] customerPrototypePatchHomeSecuritySensitivityOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypePatchHomeSecuritySensitivityOK) String() string {
	return fmt.Sprintf("[PATCH /Customers/{id}/locations/{locationId}/homeSecurity/sensitivity][%d] customerPrototypePatchHomeSecuritySensitivityOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypePatchHomeSecuritySensitivityOK) GetPayload() *models.HomeSecurity {
	return o.Payload
}

func (o *CustomerPrototypePatchHomeSecuritySensitivityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HomeSecurity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypePatchHomeSecuritySensitivityDefault creates a CustomerPrototypePatchHomeSecuritySensitivityDefault with default headers values
func NewCustomerPrototypePatchHomeSecuritySensitivityDefault(code int) *CustomerPrototypePatchHomeSecuritySensitivityDefault {
	return &CustomerPrototypePatchHomeSecuritySensitivityDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypePatchHomeSecuritySensitivityDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypePatchHomeSecuritySensitivityDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype patch home security sensitivity default response has a 2xx status code
func (o *CustomerPrototypePatchHomeSecuritySensitivityDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype patch home security sensitivity default response has a 3xx status code
func (o *CustomerPrototypePatchHomeSecuritySensitivityDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype patch home security sensitivity default response has a 4xx status code
func (o *CustomerPrototypePatchHomeSecuritySensitivityDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype patch home security sensitivity default response has a 5xx status code
func (o *CustomerPrototypePatchHomeSecuritySensitivityDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype patch home security sensitivity default response a status code equal to that given
func (o *CustomerPrototypePatchHomeSecuritySensitivityDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype patch home security sensitivity default response
func (o *CustomerPrototypePatchHomeSecuritySensitivityDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypePatchHomeSecuritySensitivityDefault) Error() string {
	return fmt.Sprintf("[PATCH /Customers/{id}/locations/{locationId}/homeSecurity/sensitivity][%d] Customer.prototype.patchHomeSecuritySensitivity default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePatchHomeSecuritySensitivityDefault) String() string {
	return fmt.Sprintf("[PATCH /Customers/{id}/locations/{locationId}/homeSecurity/sensitivity][%d] Customer.prototype.patchHomeSecuritySensitivity default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePatchHomeSecuritySensitivityDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypePatchHomeSecuritySensitivityDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
