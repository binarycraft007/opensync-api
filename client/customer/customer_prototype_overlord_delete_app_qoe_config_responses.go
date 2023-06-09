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

// CustomerPrototypeOverlordDeleteAppQoeConfigReader is a Reader for the CustomerPrototypeOverlordDeleteAppQoeConfig structure.
type CustomerPrototypeOverlordDeleteAppQoeConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypeOverlordDeleteAppQoeConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewCustomerPrototypeOverlordDeleteAppQoeConfigAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypeOverlordDeleteAppQoeConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypeOverlordDeleteAppQoeConfigAccepted creates a CustomerPrototypeOverlordDeleteAppQoeConfigAccepted with default headers values
func NewCustomerPrototypeOverlordDeleteAppQoeConfigAccepted() *CustomerPrototypeOverlordDeleteAppQoeConfigAccepted {
	return &CustomerPrototypeOverlordDeleteAppQoeConfigAccepted{}
}

/*
CustomerPrototypeOverlordDeleteAppQoeConfigAccepted describes a response with status code 202, with default header values.

Request was successful
*/
type CustomerPrototypeOverlordDeleteAppQoeConfigAccepted struct {
}

// IsSuccess returns true when this customer prototype overlord delete app qoe config accepted response has a 2xx status code
func (o *CustomerPrototypeOverlordDeleteAppQoeConfigAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype overlord delete app qoe config accepted response has a 3xx status code
func (o *CustomerPrototypeOverlordDeleteAppQoeConfigAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype overlord delete app qoe config accepted response has a 4xx status code
func (o *CustomerPrototypeOverlordDeleteAppQoeConfigAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype overlord delete app qoe config accepted response has a 5xx status code
func (o *CustomerPrototypeOverlordDeleteAppQoeConfigAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype overlord delete app qoe config accepted response a status code equal to that given
func (o *CustomerPrototypeOverlordDeleteAppQoeConfigAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the customer prototype overlord delete app qoe config accepted response
func (o *CustomerPrototypeOverlordDeleteAppQoeConfigAccepted) Code() int {
	return 202
}

func (o *CustomerPrototypeOverlordDeleteAppQoeConfigAccepted) Error() string {
	return fmt.Sprintf("[DELETE /Customers/{id}/locations/{locationId}/config/appQoe][%d] customerPrototypeOverlordDeleteAppQoeConfigAccepted ", 202)
}

func (o *CustomerPrototypeOverlordDeleteAppQoeConfigAccepted) String() string {
	return fmt.Sprintf("[DELETE /Customers/{id}/locations/{locationId}/config/appQoe][%d] customerPrototypeOverlordDeleteAppQoeConfigAccepted ", 202)
}

func (o *CustomerPrototypeOverlordDeleteAppQoeConfigAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCustomerPrototypeOverlordDeleteAppQoeConfigDefault creates a CustomerPrototypeOverlordDeleteAppQoeConfigDefault with default headers values
func NewCustomerPrototypeOverlordDeleteAppQoeConfigDefault(code int) *CustomerPrototypeOverlordDeleteAppQoeConfigDefault {
	return &CustomerPrototypeOverlordDeleteAppQoeConfigDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypeOverlordDeleteAppQoeConfigDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypeOverlordDeleteAppQoeConfigDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype overlord delete app qoe config default response has a 2xx status code
func (o *CustomerPrototypeOverlordDeleteAppQoeConfigDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype overlord delete app qoe config default response has a 3xx status code
func (o *CustomerPrototypeOverlordDeleteAppQoeConfigDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype overlord delete app qoe config default response has a 4xx status code
func (o *CustomerPrototypeOverlordDeleteAppQoeConfigDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype overlord delete app qoe config default response has a 5xx status code
func (o *CustomerPrototypeOverlordDeleteAppQoeConfigDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype overlord delete app qoe config default response a status code equal to that given
func (o *CustomerPrototypeOverlordDeleteAppQoeConfigDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype overlord delete app qoe config default response
func (o *CustomerPrototypeOverlordDeleteAppQoeConfigDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypeOverlordDeleteAppQoeConfigDefault) Error() string {
	return fmt.Sprintf("[DELETE /Customers/{id}/locations/{locationId}/config/appQoe][%d] Customer.prototype.overlordDeleteAppQoeConfig default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeOverlordDeleteAppQoeConfigDefault) String() string {
	return fmt.Sprintf("[DELETE /Customers/{id}/locations/{locationId}/config/appQoe][%d] Customer.prototype.overlordDeleteAppQoeConfig default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeOverlordDeleteAppQoeConfigDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypeOverlordDeleteAppQoeConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
