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

// CustomerPrototypeOverlordDeleteSipAlgConfigReader is a Reader for the CustomerPrototypeOverlordDeleteSipAlgConfig structure.
type CustomerPrototypeOverlordDeleteSipAlgConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypeOverlordDeleteSipAlgConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewCustomerPrototypeOverlordDeleteSipAlgConfigAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypeOverlordDeleteSipAlgConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypeOverlordDeleteSipAlgConfigAccepted creates a CustomerPrototypeOverlordDeleteSipAlgConfigAccepted with default headers values
func NewCustomerPrototypeOverlordDeleteSipAlgConfigAccepted() *CustomerPrototypeOverlordDeleteSipAlgConfigAccepted {
	return &CustomerPrototypeOverlordDeleteSipAlgConfigAccepted{}
}

/*
CustomerPrototypeOverlordDeleteSipAlgConfigAccepted describes a response with status code 202, with default header values.

Request was successful
*/
type CustomerPrototypeOverlordDeleteSipAlgConfigAccepted struct {
}

// IsSuccess returns true when this customer prototype overlord delete sip alg config accepted response has a 2xx status code
func (o *CustomerPrototypeOverlordDeleteSipAlgConfigAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype overlord delete sip alg config accepted response has a 3xx status code
func (o *CustomerPrototypeOverlordDeleteSipAlgConfigAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype overlord delete sip alg config accepted response has a 4xx status code
func (o *CustomerPrototypeOverlordDeleteSipAlgConfigAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype overlord delete sip alg config accepted response has a 5xx status code
func (o *CustomerPrototypeOverlordDeleteSipAlgConfigAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype overlord delete sip alg config accepted response a status code equal to that given
func (o *CustomerPrototypeOverlordDeleteSipAlgConfigAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the customer prototype overlord delete sip alg config accepted response
func (o *CustomerPrototypeOverlordDeleteSipAlgConfigAccepted) Code() int {
	return 202
}

func (o *CustomerPrototypeOverlordDeleteSipAlgConfigAccepted) Error() string {
	return fmt.Sprintf("[DELETE /Customers/{id}/locations/{locationId}/config/sipAlg][%d] customerPrototypeOverlordDeleteSipAlgConfigAccepted ", 202)
}

func (o *CustomerPrototypeOverlordDeleteSipAlgConfigAccepted) String() string {
	return fmt.Sprintf("[DELETE /Customers/{id}/locations/{locationId}/config/sipAlg][%d] customerPrototypeOverlordDeleteSipAlgConfigAccepted ", 202)
}

func (o *CustomerPrototypeOverlordDeleteSipAlgConfigAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCustomerPrototypeOverlordDeleteSipAlgConfigDefault creates a CustomerPrototypeOverlordDeleteSipAlgConfigDefault with default headers values
func NewCustomerPrototypeOverlordDeleteSipAlgConfigDefault(code int) *CustomerPrototypeOverlordDeleteSipAlgConfigDefault {
	return &CustomerPrototypeOverlordDeleteSipAlgConfigDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypeOverlordDeleteSipAlgConfigDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypeOverlordDeleteSipAlgConfigDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype overlord delete sip alg config default response has a 2xx status code
func (o *CustomerPrototypeOverlordDeleteSipAlgConfigDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype overlord delete sip alg config default response has a 3xx status code
func (o *CustomerPrototypeOverlordDeleteSipAlgConfigDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype overlord delete sip alg config default response has a 4xx status code
func (o *CustomerPrototypeOverlordDeleteSipAlgConfigDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype overlord delete sip alg config default response has a 5xx status code
func (o *CustomerPrototypeOverlordDeleteSipAlgConfigDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype overlord delete sip alg config default response a status code equal to that given
func (o *CustomerPrototypeOverlordDeleteSipAlgConfigDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype overlord delete sip alg config default response
func (o *CustomerPrototypeOverlordDeleteSipAlgConfigDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypeOverlordDeleteSipAlgConfigDefault) Error() string {
	return fmt.Sprintf("[DELETE /Customers/{id}/locations/{locationId}/config/sipAlg][%d] Customer.prototype.overlordDeleteSipAlgConfig default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeOverlordDeleteSipAlgConfigDefault) String() string {
	return fmt.Sprintf("[DELETE /Customers/{id}/locations/{locationId}/config/sipAlg][%d] Customer.prototype.overlordDeleteSipAlgConfig default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeOverlordDeleteSipAlgConfigDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypeOverlordDeleteSipAlgConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
