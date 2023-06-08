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

// CustomerPrototypeUnapproveDeviceReader is a Reader for the CustomerPrototypeUnapproveDevice structure.
type CustomerPrototypeUnapproveDeviceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypeUnapproveDeviceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewCustomerPrototypeUnapproveDeviceNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypeUnapproveDeviceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypeUnapproveDeviceNoContent creates a CustomerPrototypeUnapproveDeviceNoContent with default headers values
func NewCustomerPrototypeUnapproveDeviceNoContent() *CustomerPrototypeUnapproveDeviceNoContent {
	return &CustomerPrototypeUnapproveDeviceNoContent{}
}

/*
CustomerPrototypeUnapproveDeviceNoContent describes a response with status code 204, with default header values.

Request was successful
*/
type CustomerPrototypeUnapproveDeviceNoContent struct {
}

// IsSuccess returns true when this customer prototype unapprove device no content response has a 2xx status code
func (o *CustomerPrototypeUnapproveDeviceNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype unapprove device no content response has a 3xx status code
func (o *CustomerPrototypeUnapproveDeviceNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype unapprove device no content response has a 4xx status code
func (o *CustomerPrototypeUnapproveDeviceNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype unapprove device no content response has a 5xx status code
func (o *CustomerPrototypeUnapproveDeviceNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype unapprove device no content response a status code equal to that given
func (o *CustomerPrototypeUnapproveDeviceNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the customer prototype unapprove device no content response
func (o *CustomerPrototypeUnapproveDeviceNoContent) Code() int {
	return 204
}

func (o *CustomerPrototypeUnapproveDeviceNoContent) Error() string {
	return fmt.Sprintf("[DELETE /Customers/{id}/locations/{locationId}/networkAccess/networks/{networkId}/approved/{mac}][%d] customerPrototypeUnapproveDeviceNoContent ", 204)
}

func (o *CustomerPrototypeUnapproveDeviceNoContent) String() string {
	return fmt.Sprintf("[DELETE /Customers/{id}/locations/{locationId}/networkAccess/networks/{networkId}/approved/{mac}][%d] customerPrototypeUnapproveDeviceNoContent ", 204)
}

func (o *CustomerPrototypeUnapproveDeviceNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCustomerPrototypeUnapproveDeviceDefault creates a CustomerPrototypeUnapproveDeviceDefault with default headers values
func NewCustomerPrototypeUnapproveDeviceDefault(code int) *CustomerPrototypeUnapproveDeviceDefault {
	return &CustomerPrototypeUnapproveDeviceDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypeUnapproveDeviceDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypeUnapproveDeviceDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype unapprove device default response has a 2xx status code
func (o *CustomerPrototypeUnapproveDeviceDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype unapprove device default response has a 3xx status code
func (o *CustomerPrototypeUnapproveDeviceDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype unapprove device default response has a 4xx status code
func (o *CustomerPrototypeUnapproveDeviceDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype unapprove device default response has a 5xx status code
func (o *CustomerPrototypeUnapproveDeviceDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype unapprove device default response a status code equal to that given
func (o *CustomerPrototypeUnapproveDeviceDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype unapprove device default response
func (o *CustomerPrototypeUnapproveDeviceDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypeUnapproveDeviceDefault) Error() string {
	return fmt.Sprintf("[DELETE /Customers/{id}/locations/{locationId}/networkAccess/networks/{networkId}/approved/{mac}][%d] Customer.prototype.unapproveDevice default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeUnapproveDeviceDefault) String() string {
	return fmt.Sprintf("[DELETE /Customers/{id}/locations/{locationId}/networkAccess/networks/{networkId}/approved/{mac}][%d] Customer.prototype.unapproveDevice default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeUnapproveDeviceDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypeUnapproveDeviceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
