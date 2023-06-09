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

// CustomerPrototypePutSnoozeOnNodeAlertReader is a Reader for the CustomerPrototypePutSnoozeOnNodeAlert structure.
type CustomerPrototypePutSnoozeOnNodeAlertReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypePutSnoozeOnNodeAlertReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypePutSnoozeOnNodeAlertOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypePutSnoozeOnNodeAlertDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypePutSnoozeOnNodeAlertOK creates a CustomerPrototypePutSnoozeOnNodeAlertOK with default headers values
func NewCustomerPrototypePutSnoozeOnNodeAlertOK() *CustomerPrototypePutSnoozeOnNodeAlertOK {
	return &CustomerPrototypePutSnoozeOnNodeAlertOK{}
}

/*
CustomerPrototypePutSnoozeOnNodeAlertOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypePutSnoozeOnNodeAlertOK struct {
	Payload *models.Node
}

// IsSuccess returns true when this customer prototype put snooze on node alert o k response has a 2xx status code
func (o *CustomerPrototypePutSnoozeOnNodeAlertOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype put snooze on node alert o k response has a 3xx status code
func (o *CustomerPrototypePutSnoozeOnNodeAlertOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype put snooze on node alert o k response has a 4xx status code
func (o *CustomerPrototypePutSnoozeOnNodeAlertOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype put snooze on node alert o k response has a 5xx status code
func (o *CustomerPrototypePutSnoozeOnNodeAlertOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype put snooze on node alert o k response a status code equal to that given
func (o *CustomerPrototypePutSnoozeOnNodeAlertOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype put snooze on node alert o k response
func (o *CustomerPrototypePutSnoozeOnNodeAlertOK) Code() int {
	return 200
}

func (o *CustomerPrototypePutSnoozeOnNodeAlertOK) Error() string {
	return fmt.Sprintf("[PUT /Customers/{id}/locations/{locationId}/nodes/{nodeId}/alerts/{type}][%d] customerPrototypePutSnoozeOnNodeAlertOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypePutSnoozeOnNodeAlertOK) String() string {
	return fmt.Sprintf("[PUT /Customers/{id}/locations/{locationId}/nodes/{nodeId}/alerts/{type}][%d] customerPrototypePutSnoozeOnNodeAlertOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypePutSnoozeOnNodeAlertOK) GetPayload() *models.Node {
	return o.Payload
}

func (o *CustomerPrototypePutSnoozeOnNodeAlertOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Node)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypePutSnoozeOnNodeAlertDefault creates a CustomerPrototypePutSnoozeOnNodeAlertDefault with default headers values
func NewCustomerPrototypePutSnoozeOnNodeAlertDefault(code int) *CustomerPrototypePutSnoozeOnNodeAlertDefault {
	return &CustomerPrototypePutSnoozeOnNodeAlertDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypePutSnoozeOnNodeAlertDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypePutSnoozeOnNodeAlertDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype put snooze on node alert default response has a 2xx status code
func (o *CustomerPrototypePutSnoozeOnNodeAlertDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype put snooze on node alert default response has a 3xx status code
func (o *CustomerPrototypePutSnoozeOnNodeAlertDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype put snooze on node alert default response has a 4xx status code
func (o *CustomerPrototypePutSnoozeOnNodeAlertDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype put snooze on node alert default response has a 5xx status code
func (o *CustomerPrototypePutSnoozeOnNodeAlertDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype put snooze on node alert default response a status code equal to that given
func (o *CustomerPrototypePutSnoozeOnNodeAlertDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype put snooze on node alert default response
func (o *CustomerPrototypePutSnoozeOnNodeAlertDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypePutSnoozeOnNodeAlertDefault) Error() string {
	return fmt.Sprintf("[PUT /Customers/{id}/locations/{locationId}/nodes/{nodeId}/alerts/{type}][%d] Customer.prototype.putSnoozeOnNodeAlert default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePutSnoozeOnNodeAlertDefault) String() string {
	return fmt.Sprintf("[PUT /Customers/{id}/locations/{locationId}/nodes/{nodeId}/alerts/{type}][%d] Customer.prototype.putSnoozeOnNodeAlert default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePutSnoozeOnNodeAlertDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypePutSnoozeOnNodeAlertDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
