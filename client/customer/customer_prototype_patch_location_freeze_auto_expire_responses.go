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

// CustomerPrototypePatchLocationFreezeAutoExpireReader is a Reader for the CustomerPrototypePatchLocationFreezeAutoExpire structure.
type CustomerPrototypePatchLocationFreezeAutoExpireReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypePatchLocationFreezeAutoExpireReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewCustomerPrototypePatchLocationFreezeAutoExpireNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypePatchLocationFreezeAutoExpireDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypePatchLocationFreezeAutoExpireNoContent creates a CustomerPrototypePatchLocationFreezeAutoExpireNoContent with default headers values
func NewCustomerPrototypePatchLocationFreezeAutoExpireNoContent() *CustomerPrototypePatchLocationFreezeAutoExpireNoContent {
	return &CustomerPrototypePatchLocationFreezeAutoExpireNoContent{}
}

/*
CustomerPrototypePatchLocationFreezeAutoExpireNoContent describes a response with status code 204, with default header values.

Request was successful
*/
type CustomerPrototypePatchLocationFreezeAutoExpireNoContent struct {
}

// IsSuccess returns true when this customer prototype patch location freeze auto expire no content response has a 2xx status code
func (o *CustomerPrototypePatchLocationFreezeAutoExpireNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype patch location freeze auto expire no content response has a 3xx status code
func (o *CustomerPrototypePatchLocationFreezeAutoExpireNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype patch location freeze auto expire no content response has a 4xx status code
func (o *CustomerPrototypePatchLocationFreezeAutoExpireNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype patch location freeze auto expire no content response has a 5xx status code
func (o *CustomerPrototypePatchLocationFreezeAutoExpireNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype patch location freeze auto expire no content response a status code equal to that given
func (o *CustomerPrototypePatchLocationFreezeAutoExpireNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the customer prototype patch location freeze auto expire no content response
func (o *CustomerPrototypePatchLocationFreezeAutoExpireNoContent) Code() int {
	return 204
}

func (o *CustomerPrototypePatchLocationFreezeAutoExpireNoContent) Error() string {
	return fmt.Sprintf("[PATCH /Customers/{id}/locations/{locationId}/freeze/autoExpire][%d] customerPrototypePatchLocationFreezeAutoExpireNoContent ", 204)
}

func (o *CustomerPrototypePatchLocationFreezeAutoExpireNoContent) String() string {
	return fmt.Sprintf("[PATCH /Customers/{id}/locations/{locationId}/freeze/autoExpire][%d] customerPrototypePatchLocationFreezeAutoExpireNoContent ", 204)
}

func (o *CustomerPrototypePatchLocationFreezeAutoExpireNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCustomerPrototypePatchLocationFreezeAutoExpireDefault creates a CustomerPrototypePatchLocationFreezeAutoExpireDefault with default headers values
func NewCustomerPrototypePatchLocationFreezeAutoExpireDefault(code int) *CustomerPrototypePatchLocationFreezeAutoExpireDefault {
	return &CustomerPrototypePatchLocationFreezeAutoExpireDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypePatchLocationFreezeAutoExpireDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypePatchLocationFreezeAutoExpireDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype patch location freeze auto expire default response has a 2xx status code
func (o *CustomerPrototypePatchLocationFreezeAutoExpireDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype patch location freeze auto expire default response has a 3xx status code
func (o *CustomerPrototypePatchLocationFreezeAutoExpireDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype patch location freeze auto expire default response has a 4xx status code
func (o *CustomerPrototypePatchLocationFreezeAutoExpireDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype patch location freeze auto expire default response has a 5xx status code
func (o *CustomerPrototypePatchLocationFreezeAutoExpireDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype patch location freeze auto expire default response a status code equal to that given
func (o *CustomerPrototypePatchLocationFreezeAutoExpireDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype patch location freeze auto expire default response
func (o *CustomerPrototypePatchLocationFreezeAutoExpireDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypePatchLocationFreezeAutoExpireDefault) Error() string {
	return fmt.Sprintf("[PATCH /Customers/{id}/locations/{locationId}/freeze/autoExpire][%d] Customer.prototype.patchLocationFreezeAutoExpire default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePatchLocationFreezeAutoExpireDefault) String() string {
	return fmt.Sprintf("[PATCH /Customers/{id}/locations/{locationId}/freeze/autoExpire][%d] Customer.prototype.patchLocationFreezeAutoExpire default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePatchLocationFreezeAutoExpireDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypePatchLocationFreezeAutoExpireDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}