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

// CustomerPrototypeOverlordUpdateSipAlgConfigReader is a Reader for the CustomerPrototypeOverlordUpdateSipAlgConfig structure.
type CustomerPrototypeOverlordUpdateSipAlgConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypeOverlordUpdateSipAlgConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewCustomerPrototypeOverlordUpdateSipAlgConfigAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypeOverlordUpdateSipAlgConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypeOverlordUpdateSipAlgConfigAccepted creates a CustomerPrototypeOverlordUpdateSipAlgConfigAccepted with default headers values
func NewCustomerPrototypeOverlordUpdateSipAlgConfigAccepted() *CustomerPrototypeOverlordUpdateSipAlgConfigAccepted {
	return &CustomerPrototypeOverlordUpdateSipAlgConfigAccepted{}
}

/*
CustomerPrototypeOverlordUpdateSipAlgConfigAccepted describes a response with status code 202, with default header values.

Request was successful
*/
type CustomerPrototypeOverlordUpdateSipAlgConfigAccepted struct {
}

// IsSuccess returns true when this customer prototype overlord update sip alg config accepted response has a 2xx status code
func (o *CustomerPrototypeOverlordUpdateSipAlgConfigAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype overlord update sip alg config accepted response has a 3xx status code
func (o *CustomerPrototypeOverlordUpdateSipAlgConfigAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype overlord update sip alg config accepted response has a 4xx status code
func (o *CustomerPrototypeOverlordUpdateSipAlgConfigAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype overlord update sip alg config accepted response has a 5xx status code
func (o *CustomerPrototypeOverlordUpdateSipAlgConfigAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype overlord update sip alg config accepted response a status code equal to that given
func (o *CustomerPrototypeOverlordUpdateSipAlgConfigAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the customer prototype overlord update sip alg config accepted response
func (o *CustomerPrototypeOverlordUpdateSipAlgConfigAccepted) Code() int {
	return 202
}

func (o *CustomerPrototypeOverlordUpdateSipAlgConfigAccepted) Error() string {
	return fmt.Sprintf("[PATCH /Customers/{id}/locations/{locationId}/config/sipAlg][%d] customerPrototypeOverlordUpdateSipAlgConfigAccepted ", 202)
}

func (o *CustomerPrototypeOverlordUpdateSipAlgConfigAccepted) String() string {
	return fmt.Sprintf("[PATCH /Customers/{id}/locations/{locationId}/config/sipAlg][%d] customerPrototypeOverlordUpdateSipAlgConfigAccepted ", 202)
}

func (o *CustomerPrototypeOverlordUpdateSipAlgConfigAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCustomerPrototypeOverlordUpdateSipAlgConfigDefault creates a CustomerPrototypeOverlordUpdateSipAlgConfigDefault with default headers values
func NewCustomerPrototypeOverlordUpdateSipAlgConfigDefault(code int) *CustomerPrototypeOverlordUpdateSipAlgConfigDefault {
	return &CustomerPrototypeOverlordUpdateSipAlgConfigDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypeOverlordUpdateSipAlgConfigDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypeOverlordUpdateSipAlgConfigDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype overlord update sip alg config default response has a 2xx status code
func (o *CustomerPrototypeOverlordUpdateSipAlgConfigDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype overlord update sip alg config default response has a 3xx status code
func (o *CustomerPrototypeOverlordUpdateSipAlgConfigDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype overlord update sip alg config default response has a 4xx status code
func (o *CustomerPrototypeOverlordUpdateSipAlgConfigDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype overlord update sip alg config default response has a 5xx status code
func (o *CustomerPrototypeOverlordUpdateSipAlgConfigDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype overlord update sip alg config default response a status code equal to that given
func (o *CustomerPrototypeOverlordUpdateSipAlgConfigDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype overlord update sip alg config default response
func (o *CustomerPrototypeOverlordUpdateSipAlgConfigDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypeOverlordUpdateSipAlgConfigDefault) Error() string {
	return fmt.Sprintf("[PATCH /Customers/{id}/locations/{locationId}/config/sipAlg][%d] Customer.prototype.overlordUpdateSipAlgConfig default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeOverlordUpdateSipAlgConfigDefault) String() string {
	return fmt.Sprintf("[PATCH /Customers/{id}/locations/{locationId}/config/sipAlg][%d] Customer.prototype.overlordUpdateSipAlgConfig default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeOverlordUpdateSipAlgConfigDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypeOverlordUpdateSipAlgConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
