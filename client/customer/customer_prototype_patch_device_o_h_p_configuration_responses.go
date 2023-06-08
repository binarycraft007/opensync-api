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

// CustomerPrototypePatchDeviceOHPConfigurationReader is a Reader for the CustomerPrototypePatchDeviceOHPConfiguration structure.
type CustomerPrototypePatchDeviceOHPConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypePatchDeviceOHPConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewCustomerPrototypePatchDeviceOHPConfigurationNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypePatchDeviceOHPConfigurationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypePatchDeviceOHPConfigurationNoContent creates a CustomerPrototypePatchDeviceOHPConfigurationNoContent with default headers values
func NewCustomerPrototypePatchDeviceOHPConfigurationNoContent() *CustomerPrototypePatchDeviceOHPConfigurationNoContent {
	return &CustomerPrototypePatchDeviceOHPConfigurationNoContent{}
}

/*
CustomerPrototypePatchDeviceOHPConfigurationNoContent describes a response with status code 204, with default header values.

Request was successful
*/
type CustomerPrototypePatchDeviceOHPConfigurationNoContent struct {
	Payload interface{}
}

// IsSuccess returns true when this customer prototype patch device o h p configuration no content response has a 2xx status code
func (o *CustomerPrototypePatchDeviceOHPConfigurationNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype patch device o h p configuration no content response has a 3xx status code
func (o *CustomerPrototypePatchDeviceOHPConfigurationNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype patch device o h p configuration no content response has a 4xx status code
func (o *CustomerPrototypePatchDeviceOHPConfigurationNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype patch device o h p configuration no content response has a 5xx status code
func (o *CustomerPrototypePatchDeviceOHPConfigurationNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype patch device o h p configuration no content response a status code equal to that given
func (o *CustomerPrototypePatchDeviceOHPConfigurationNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the customer prototype patch device o h p configuration no content response
func (o *CustomerPrototypePatchDeviceOHPConfigurationNoContent) Code() int {
	return 204
}

func (o *CustomerPrototypePatchDeviceOHPConfigurationNoContent) Error() string {
	return fmt.Sprintf("[PATCH /Customers/{id}/locations/{locationId}/devices/{mac}/securityPolicy/ohp][%d] customerPrototypePatchDeviceOHPConfigurationNoContent  %+v", 204, o.Payload)
}

func (o *CustomerPrototypePatchDeviceOHPConfigurationNoContent) String() string {
	return fmt.Sprintf("[PATCH /Customers/{id}/locations/{locationId}/devices/{mac}/securityPolicy/ohp][%d] customerPrototypePatchDeviceOHPConfigurationNoContent  %+v", 204, o.Payload)
}

func (o *CustomerPrototypePatchDeviceOHPConfigurationNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *CustomerPrototypePatchDeviceOHPConfigurationNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypePatchDeviceOHPConfigurationDefault creates a CustomerPrototypePatchDeviceOHPConfigurationDefault with default headers values
func NewCustomerPrototypePatchDeviceOHPConfigurationDefault(code int) *CustomerPrototypePatchDeviceOHPConfigurationDefault {
	return &CustomerPrototypePatchDeviceOHPConfigurationDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypePatchDeviceOHPConfigurationDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypePatchDeviceOHPConfigurationDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype patch device o h p configuration default response has a 2xx status code
func (o *CustomerPrototypePatchDeviceOHPConfigurationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype patch device o h p configuration default response has a 3xx status code
func (o *CustomerPrototypePatchDeviceOHPConfigurationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype patch device o h p configuration default response has a 4xx status code
func (o *CustomerPrototypePatchDeviceOHPConfigurationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype patch device o h p configuration default response has a 5xx status code
func (o *CustomerPrototypePatchDeviceOHPConfigurationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype patch device o h p configuration default response a status code equal to that given
func (o *CustomerPrototypePatchDeviceOHPConfigurationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype patch device o h p configuration default response
func (o *CustomerPrototypePatchDeviceOHPConfigurationDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypePatchDeviceOHPConfigurationDefault) Error() string {
	return fmt.Sprintf("[PATCH /Customers/{id}/locations/{locationId}/devices/{mac}/securityPolicy/ohp][%d] Customer.prototype.patchDeviceOHPConfiguration default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePatchDeviceOHPConfigurationDefault) String() string {
	return fmt.Sprintf("[PATCH /Customers/{id}/locations/{locationId}/devices/{mac}/securityPolicy/ohp][%d] Customer.prototype.patchDeviceOHPConfiguration default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePatchDeviceOHPConfigurationDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypePatchDeviceOHPConfigurationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
