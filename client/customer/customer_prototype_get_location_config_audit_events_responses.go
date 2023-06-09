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

// CustomerPrototypeGetLocationConfigAuditEventsReader is a Reader for the CustomerPrototypeGetLocationConfigAuditEvents structure.
type CustomerPrototypeGetLocationConfigAuditEventsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypeGetLocationConfigAuditEventsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypeGetLocationConfigAuditEventsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypeGetLocationConfigAuditEventsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypeGetLocationConfigAuditEventsOK creates a CustomerPrototypeGetLocationConfigAuditEventsOK with default headers values
func NewCustomerPrototypeGetLocationConfigAuditEventsOK() *CustomerPrototypeGetLocationConfigAuditEventsOK {
	return &CustomerPrototypeGetLocationConfigAuditEventsOK{}
}

/*
CustomerPrototypeGetLocationConfigAuditEventsOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypeGetLocationConfigAuditEventsOK struct {
	Payload models.XAny
}

// IsSuccess returns true when this customer prototype get location config audit events o k response has a 2xx status code
func (o *CustomerPrototypeGetLocationConfigAuditEventsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype get location config audit events o k response has a 3xx status code
func (o *CustomerPrototypeGetLocationConfigAuditEventsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype get location config audit events o k response has a 4xx status code
func (o *CustomerPrototypeGetLocationConfigAuditEventsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype get location config audit events o k response has a 5xx status code
func (o *CustomerPrototypeGetLocationConfigAuditEventsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype get location config audit events o k response a status code equal to that given
func (o *CustomerPrototypeGetLocationConfigAuditEventsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype get location config audit events o k response
func (o *CustomerPrototypeGetLocationConfigAuditEventsOK) Code() int {
	return 200
}

func (o *CustomerPrototypeGetLocationConfigAuditEventsOK) Error() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/configAudit/events][%d] customerPrototypeGetLocationConfigAuditEventsOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeGetLocationConfigAuditEventsOK) String() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/configAudit/events][%d] customerPrototypeGetLocationConfigAuditEventsOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeGetLocationConfigAuditEventsOK) GetPayload() models.XAny {
	return o.Payload
}

func (o *CustomerPrototypeGetLocationConfigAuditEventsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypeGetLocationConfigAuditEventsDefault creates a CustomerPrototypeGetLocationConfigAuditEventsDefault with default headers values
func NewCustomerPrototypeGetLocationConfigAuditEventsDefault(code int) *CustomerPrototypeGetLocationConfigAuditEventsDefault {
	return &CustomerPrototypeGetLocationConfigAuditEventsDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypeGetLocationConfigAuditEventsDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypeGetLocationConfigAuditEventsDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype get location config audit events default response has a 2xx status code
func (o *CustomerPrototypeGetLocationConfigAuditEventsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype get location config audit events default response has a 3xx status code
func (o *CustomerPrototypeGetLocationConfigAuditEventsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype get location config audit events default response has a 4xx status code
func (o *CustomerPrototypeGetLocationConfigAuditEventsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype get location config audit events default response has a 5xx status code
func (o *CustomerPrototypeGetLocationConfigAuditEventsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype get location config audit events default response a status code equal to that given
func (o *CustomerPrototypeGetLocationConfigAuditEventsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype get location config audit events default response
func (o *CustomerPrototypeGetLocationConfigAuditEventsDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypeGetLocationConfigAuditEventsDefault) Error() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/configAudit/events][%d] Customer.prototype.getLocationConfigAuditEvents default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeGetLocationConfigAuditEventsDefault) String() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/configAudit/events][%d] Customer.prototype.getLocationConfigAuditEvents default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeGetLocationConfigAuditEventsDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypeGetLocationConfigAuditEventsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
