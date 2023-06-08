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

// CustomerPrototypePatchKvConfigsReader is a Reader for the CustomerPrototypePatchKvConfigs structure.
type CustomerPrototypePatchKvConfigsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypePatchKvConfigsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypePatchKvConfigsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypePatchKvConfigsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypePatchKvConfigsOK creates a CustomerPrototypePatchKvConfigsOK with default headers values
func NewCustomerPrototypePatchKvConfigsOK() *CustomerPrototypePatchKvConfigsOK {
	return &CustomerPrototypePatchKvConfigsOK{}
}

/*
CustomerPrototypePatchKvConfigsOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypePatchKvConfigsOK struct {
	Payload []*models.KvConfig
}

// IsSuccess returns true when this customer prototype patch kv configs o k response has a 2xx status code
func (o *CustomerPrototypePatchKvConfigsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype patch kv configs o k response has a 3xx status code
func (o *CustomerPrototypePatchKvConfigsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype patch kv configs o k response has a 4xx status code
func (o *CustomerPrototypePatchKvConfigsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype patch kv configs o k response has a 5xx status code
func (o *CustomerPrototypePatchKvConfigsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype patch kv configs o k response a status code equal to that given
func (o *CustomerPrototypePatchKvConfigsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype patch kv configs o k response
func (o *CustomerPrototypePatchKvConfigsOK) Code() int {
	return 200
}

func (o *CustomerPrototypePatchKvConfigsOK) Error() string {
	return fmt.Sprintf("[PATCH /Customers/{id}/locations/{locationId}/nodes/{nodeId}/kvConfigs][%d] customerPrototypePatchKvConfigsOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypePatchKvConfigsOK) String() string {
	return fmt.Sprintf("[PATCH /Customers/{id}/locations/{locationId}/nodes/{nodeId}/kvConfigs][%d] customerPrototypePatchKvConfigsOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypePatchKvConfigsOK) GetPayload() []*models.KvConfig {
	return o.Payload
}

func (o *CustomerPrototypePatchKvConfigsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypePatchKvConfigsDefault creates a CustomerPrototypePatchKvConfigsDefault with default headers values
func NewCustomerPrototypePatchKvConfigsDefault(code int) *CustomerPrototypePatchKvConfigsDefault {
	return &CustomerPrototypePatchKvConfigsDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypePatchKvConfigsDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypePatchKvConfigsDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype patch kv configs default response has a 2xx status code
func (o *CustomerPrototypePatchKvConfigsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype patch kv configs default response has a 3xx status code
func (o *CustomerPrototypePatchKvConfigsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype patch kv configs default response has a 4xx status code
func (o *CustomerPrototypePatchKvConfigsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype patch kv configs default response has a 5xx status code
func (o *CustomerPrototypePatchKvConfigsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype patch kv configs default response a status code equal to that given
func (o *CustomerPrototypePatchKvConfigsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype patch kv configs default response
func (o *CustomerPrototypePatchKvConfigsDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypePatchKvConfigsDefault) Error() string {
	return fmt.Sprintf("[PATCH /Customers/{id}/locations/{locationId}/nodes/{nodeId}/kvConfigs][%d] Customer.prototype.patchKvConfigs default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePatchKvConfigsDefault) String() string {
	return fmt.Sprintf("[PATCH /Customers/{id}/locations/{locationId}/nodes/{nodeId}/kvConfigs][%d] Customer.prototype.patchKvConfigs default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePatchKvConfigsDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypePatchKvConfigsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
