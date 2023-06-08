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

// CustomerPrototypeRebootNodeReader is a Reader for the CustomerPrototypeRebootNode structure.
type CustomerPrototypeRebootNodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypeRebootNodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewCustomerPrototypeRebootNodeNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypeRebootNodeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypeRebootNodeNoContent creates a CustomerPrototypeRebootNodeNoContent with default headers values
func NewCustomerPrototypeRebootNodeNoContent() *CustomerPrototypeRebootNodeNoContent {
	return &CustomerPrototypeRebootNodeNoContent{}
}

/*
CustomerPrototypeRebootNodeNoContent describes a response with status code 204, with default header values.

Request was successful
*/
type CustomerPrototypeRebootNodeNoContent struct {
}

// IsSuccess returns true when this customer prototype reboot node no content response has a 2xx status code
func (o *CustomerPrototypeRebootNodeNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype reboot node no content response has a 3xx status code
func (o *CustomerPrototypeRebootNodeNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype reboot node no content response has a 4xx status code
func (o *CustomerPrototypeRebootNodeNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype reboot node no content response has a 5xx status code
func (o *CustomerPrototypeRebootNodeNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype reboot node no content response a status code equal to that given
func (o *CustomerPrototypeRebootNodeNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the customer prototype reboot node no content response
func (o *CustomerPrototypeRebootNodeNoContent) Code() int {
	return 204
}

func (o *CustomerPrototypeRebootNodeNoContent) Error() string {
	return fmt.Sprintf("[PUT /Customers/{id}/locations/{locationId}/nodes/{nodeId}/reboot][%d] customerPrototypeRebootNodeNoContent ", 204)
}

func (o *CustomerPrototypeRebootNodeNoContent) String() string {
	return fmt.Sprintf("[PUT /Customers/{id}/locations/{locationId}/nodes/{nodeId}/reboot][%d] customerPrototypeRebootNodeNoContent ", 204)
}

func (o *CustomerPrototypeRebootNodeNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCustomerPrototypeRebootNodeDefault creates a CustomerPrototypeRebootNodeDefault with default headers values
func NewCustomerPrototypeRebootNodeDefault(code int) *CustomerPrototypeRebootNodeDefault {
	return &CustomerPrototypeRebootNodeDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypeRebootNodeDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypeRebootNodeDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype reboot node default response has a 2xx status code
func (o *CustomerPrototypeRebootNodeDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype reboot node default response has a 3xx status code
func (o *CustomerPrototypeRebootNodeDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype reboot node default response has a 4xx status code
func (o *CustomerPrototypeRebootNodeDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype reboot node default response has a 5xx status code
func (o *CustomerPrototypeRebootNodeDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype reboot node default response a status code equal to that given
func (o *CustomerPrototypeRebootNodeDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype reboot node default response
func (o *CustomerPrototypeRebootNodeDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypeRebootNodeDefault) Error() string {
	return fmt.Sprintf("[PUT /Customers/{id}/locations/{locationId}/nodes/{nodeId}/reboot][%d] Customer.prototype.rebootNode default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeRebootNodeDefault) String() string {
	return fmt.Sprintf("[PUT /Customers/{id}/locations/{locationId}/nodes/{nodeId}/reboot][%d] Customer.prototype.rebootNode default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeRebootNodeDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypeRebootNodeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
