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

// CustomerPrototypePostFrontHaulsDppReader is a Reader for the CustomerPrototypePostFrontHaulsDpp structure.
type CustomerPrototypePostFrontHaulsDppReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypePostFrontHaulsDppReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewCustomerPrototypePostFrontHaulsDppAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypePostFrontHaulsDppDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypePostFrontHaulsDppAccepted creates a CustomerPrototypePostFrontHaulsDppAccepted with default headers values
func NewCustomerPrototypePostFrontHaulsDppAccepted() *CustomerPrototypePostFrontHaulsDppAccepted {
	return &CustomerPrototypePostFrontHaulsDppAccepted{}
}

/*
CustomerPrototypePostFrontHaulsDppAccepted describes a response with status code 202, with default header values.

Request was successful
*/
type CustomerPrototypePostFrontHaulsDppAccepted struct {
	Payload interface{}
}

// IsSuccess returns true when this customer prototype post front hauls dpp accepted response has a 2xx status code
func (o *CustomerPrototypePostFrontHaulsDppAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype post front hauls dpp accepted response has a 3xx status code
func (o *CustomerPrototypePostFrontHaulsDppAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype post front hauls dpp accepted response has a 4xx status code
func (o *CustomerPrototypePostFrontHaulsDppAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype post front hauls dpp accepted response has a 5xx status code
func (o *CustomerPrototypePostFrontHaulsDppAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype post front hauls dpp accepted response a status code equal to that given
func (o *CustomerPrototypePostFrontHaulsDppAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the customer prototype post front hauls dpp accepted response
func (o *CustomerPrototypePostFrontHaulsDppAccepted) Code() int {
	return 202
}

func (o *CustomerPrototypePostFrontHaulsDppAccepted) Error() string {
	return fmt.Sprintf("[POST /Customers/{id}/locations/{locationId}/secondaryNetworks/fronthauls/{networkId}/dpp][%d] customerPrototypePostFrontHaulsDppAccepted  %+v", 202, o.Payload)
}

func (o *CustomerPrototypePostFrontHaulsDppAccepted) String() string {
	return fmt.Sprintf("[POST /Customers/{id}/locations/{locationId}/secondaryNetworks/fronthauls/{networkId}/dpp][%d] customerPrototypePostFrontHaulsDppAccepted  %+v", 202, o.Payload)
}

func (o *CustomerPrototypePostFrontHaulsDppAccepted) GetPayload() interface{} {
	return o.Payload
}

func (o *CustomerPrototypePostFrontHaulsDppAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypePostFrontHaulsDppDefault creates a CustomerPrototypePostFrontHaulsDppDefault with default headers values
func NewCustomerPrototypePostFrontHaulsDppDefault(code int) *CustomerPrototypePostFrontHaulsDppDefault {
	return &CustomerPrototypePostFrontHaulsDppDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypePostFrontHaulsDppDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypePostFrontHaulsDppDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype post front hauls dpp default response has a 2xx status code
func (o *CustomerPrototypePostFrontHaulsDppDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype post front hauls dpp default response has a 3xx status code
func (o *CustomerPrototypePostFrontHaulsDppDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype post front hauls dpp default response has a 4xx status code
func (o *CustomerPrototypePostFrontHaulsDppDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype post front hauls dpp default response has a 5xx status code
func (o *CustomerPrototypePostFrontHaulsDppDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype post front hauls dpp default response a status code equal to that given
func (o *CustomerPrototypePostFrontHaulsDppDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype post front hauls dpp default response
func (o *CustomerPrototypePostFrontHaulsDppDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypePostFrontHaulsDppDefault) Error() string {
	return fmt.Sprintf("[POST /Customers/{id}/locations/{locationId}/secondaryNetworks/fronthauls/{networkId}/dpp][%d] Customer.prototype.postFrontHaulsDpp default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePostFrontHaulsDppDefault) String() string {
	return fmt.Sprintf("[POST /Customers/{id}/locations/{locationId}/secondaryNetworks/fronthauls/{networkId}/dpp][%d] Customer.prototype.postFrontHaulsDpp default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePostFrontHaulsDppDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypePostFrontHaulsDppDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
