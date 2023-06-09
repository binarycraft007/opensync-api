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

// CustomerPrototypePostWhitelistApprovalRequestReader is a Reader for the CustomerPrototypePostWhitelistApprovalRequest structure.
type CustomerPrototypePostWhitelistApprovalRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypePostWhitelistApprovalRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypePostWhitelistApprovalRequestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypePostWhitelistApprovalRequestDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypePostWhitelistApprovalRequestOK creates a CustomerPrototypePostWhitelistApprovalRequestOK with default headers values
func NewCustomerPrototypePostWhitelistApprovalRequestOK() *CustomerPrototypePostWhitelistApprovalRequestOK {
	return &CustomerPrototypePostWhitelistApprovalRequestOK{}
}

/*
CustomerPrototypePostWhitelistApprovalRequestOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypePostWhitelistApprovalRequestOK struct {
	Payload interface{}
}

// IsSuccess returns true when this customer prototype post whitelist approval request o k response has a 2xx status code
func (o *CustomerPrototypePostWhitelistApprovalRequestOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype post whitelist approval request o k response has a 3xx status code
func (o *CustomerPrototypePostWhitelistApprovalRequestOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype post whitelist approval request o k response has a 4xx status code
func (o *CustomerPrototypePostWhitelistApprovalRequestOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype post whitelist approval request o k response has a 5xx status code
func (o *CustomerPrototypePostWhitelistApprovalRequestOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype post whitelist approval request o k response a status code equal to that given
func (o *CustomerPrototypePostWhitelistApprovalRequestOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype post whitelist approval request o k response
func (o *CustomerPrototypePostWhitelistApprovalRequestOK) Code() int {
	return 200
}

func (o *CustomerPrototypePostWhitelistApprovalRequestOK) Error() string {
	return fmt.Sprintf("[POST /Customers/{id}/locations/{locationId}/securityPolicy/websites/whitelist/approvalRequests][%d] customerPrototypePostWhitelistApprovalRequestOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypePostWhitelistApprovalRequestOK) String() string {
	return fmt.Sprintf("[POST /Customers/{id}/locations/{locationId}/securityPolicy/websites/whitelist/approvalRequests][%d] customerPrototypePostWhitelistApprovalRequestOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypePostWhitelistApprovalRequestOK) GetPayload() interface{} {
	return o.Payload
}

func (o *CustomerPrototypePostWhitelistApprovalRequestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypePostWhitelistApprovalRequestDefault creates a CustomerPrototypePostWhitelistApprovalRequestDefault with default headers values
func NewCustomerPrototypePostWhitelistApprovalRequestDefault(code int) *CustomerPrototypePostWhitelistApprovalRequestDefault {
	return &CustomerPrototypePostWhitelistApprovalRequestDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypePostWhitelistApprovalRequestDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypePostWhitelistApprovalRequestDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype post whitelist approval request default response has a 2xx status code
func (o *CustomerPrototypePostWhitelistApprovalRequestDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype post whitelist approval request default response has a 3xx status code
func (o *CustomerPrototypePostWhitelistApprovalRequestDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype post whitelist approval request default response has a 4xx status code
func (o *CustomerPrototypePostWhitelistApprovalRequestDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype post whitelist approval request default response has a 5xx status code
func (o *CustomerPrototypePostWhitelistApprovalRequestDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype post whitelist approval request default response a status code equal to that given
func (o *CustomerPrototypePostWhitelistApprovalRequestDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype post whitelist approval request default response
func (o *CustomerPrototypePostWhitelistApprovalRequestDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypePostWhitelistApprovalRequestDefault) Error() string {
	return fmt.Sprintf("[POST /Customers/{id}/locations/{locationId}/securityPolicy/websites/whitelist/approvalRequests][%d] Customer.prototype.postWhitelistApprovalRequest default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePostWhitelistApprovalRequestDefault) String() string {
	return fmt.Sprintf("[POST /Customers/{id}/locations/{locationId}/securityPolicy/websites/whitelist/approvalRequests][%d] Customer.prototype.postWhitelistApprovalRequest default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePostWhitelistApprovalRequestDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypePostWhitelistApprovalRequestDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
