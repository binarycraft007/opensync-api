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

// CustomerPrototypePutSubscriptionReader is a Reader for the CustomerPrototypePutSubscription structure.
type CustomerPrototypePutSubscriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypePutSubscriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypePutSubscriptionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypePutSubscriptionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypePutSubscriptionOK creates a CustomerPrototypePutSubscriptionOK with default headers values
func NewCustomerPrototypePutSubscriptionOK() *CustomerPrototypePutSubscriptionOK {
	return &CustomerPrototypePutSubscriptionOK{}
}

/*
CustomerPrototypePutSubscriptionOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypePutSubscriptionOK struct {
	Payload interface{}
}

// IsSuccess returns true when this customer prototype put subscription o k response has a 2xx status code
func (o *CustomerPrototypePutSubscriptionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype put subscription o k response has a 3xx status code
func (o *CustomerPrototypePutSubscriptionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype put subscription o k response has a 4xx status code
func (o *CustomerPrototypePutSubscriptionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype put subscription o k response has a 5xx status code
func (o *CustomerPrototypePutSubscriptionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype put subscription o k response a status code equal to that given
func (o *CustomerPrototypePutSubscriptionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype put subscription o k response
func (o *CustomerPrototypePutSubscriptionOK) Code() int {
	return 200
}

func (o *CustomerPrototypePutSubscriptionOK) Error() string {
	return fmt.Sprintf("[PUT /Customers/{id}/locations/{locationId}/subscription][%d] customerPrototypePutSubscriptionOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypePutSubscriptionOK) String() string {
	return fmt.Sprintf("[PUT /Customers/{id}/locations/{locationId}/subscription][%d] customerPrototypePutSubscriptionOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypePutSubscriptionOK) GetPayload() interface{} {
	return o.Payload
}

func (o *CustomerPrototypePutSubscriptionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypePutSubscriptionDefault creates a CustomerPrototypePutSubscriptionDefault with default headers values
func NewCustomerPrototypePutSubscriptionDefault(code int) *CustomerPrototypePutSubscriptionDefault {
	return &CustomerPrototypePutSubscriptionDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypePutSubscriptionDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypePutSubscriptionDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype put subscription default response has a 2xx status code
func (o *CustomerPrototypePutSubscriptionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype put subscription default response has a 3xx status code
func (o *CustomerPrototypePutSubscriptionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype put subscription default response has a 4xx status code
func (o *CustomerPrototypePutSubscriptionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype put subscription default response has a 5xx status code
func (o *CustomerPrototypePutSubscriptionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype put subscription default response a status code equal to that given
func (o *CustomerPrototypePutSubscriptionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype put subscription default response
func (o *CustomerPrototypePutSubscriptionDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypePutSubscriptionDefault) Error() string {
	return fmt.Sprintf("[PUT /Customers/{id}/locations/{locationId}/subscription][%d] Customer.prototype.putSubscription default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePutSubscriptionDefault) String() string {
	return fmt.Sprintf("[PUT /Customers/{id}/locations/{locationId}/subscription][%d] Customer.prototype.putSubscription default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePutSubscriptionDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypePutSubscriptionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
