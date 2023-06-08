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

// CustomerPrototypeSetPrimarySecondaryNetworksReader is a Reader for the CustomerPrototypeSetPrimarySecondaryNetworks structure.
type CustomerPrototypeSetPrimarySecondaryNetworksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypeSetPrimarySecondaryNetworksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewCustomerPrototypeSetPrimarySecondaryNetworksAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypeSetPrimarySecondaryNetworksDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypeSetPrimarySecondaryNetworksAccepted creates a CustomerPrototypeSetPrimarySecondaryNetworksAccepted with default headers values
func NewCustomerPrototypeSetPrimarySecondaryNetworksAccepted() *CustomerPrototypeSetPrimarySecondaryNetworksAccepted {
	return &CustomerPrototypeSetPrimarySecondaryNetworksAccepted{}
}

/*
CustomerPrototypeSetPrimarySecondaryNetworksAccepted describes a response with status code 202, with default header values.

Request was successful
*/
type CustomerPrototypeSetPrimarySecondaryNetworksAccepted struct {
}

// IsSuccess returns true when this customer prototype set primary secondary networks accepted response has a 2xx status code
func (o *CustomerPrototypeSetPrimarySecondaryNetworksAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype set primary secondary networks accepted response has a 3xx status code
func (o *CustomerPrototypeSetPrimarySecondaryNetworksAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype set primary secondary networks accepted response has a 4xx status code
func (o *CustomerPrototypeSetPrimarySecondaryNetworksAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype set primary secondary networks accepted response has a 5xx status code
func (o *CustomerPrototypeSetPrimarySecondaryNetworksAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype set primary secondary networks accepted response a status code equal to that given
func (o *CustomerPrototypeSetPrimarySecondaryNetworksAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the customer prototype set primary secondary networks accepted response
func (o *CustomerPrototypeSetPrimarySecondaryNetworksAccepted) Code() int {
	return 202
}

func (o *CustomerPrototypeSetPrimarySecondaryNetworksAccepted) Error() string {
	return fmt.Sprintf("[PUT /Customers/{id}/locations/{locationId}/primarySecondaryNetworks][%d] customerPrototypeSetPrimarySecondaryNetworksAccepted ", 202)
}

func (o *CustomerPrototypeSetPrimarySecondaryNetworksAccepted) String() string {
	return fmt.Sprintf("[PUT /Customers/{id}/locations/{locationId}/primarySecondaryNetworks][%d] customerPrototypeSetPrimarySecondaryNetworksAccepted ", 202)
}

func (o *CustomerPrototypeSetPrimarySecondaryNetworksAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCustomerPrototypeSetPrimarySecondaryNetworksDefault creates a CustomerPrototypeSetPrimarySecondaryNetworksDefault with default headers values
func NewCustomerPrototypeSetPrimarySecondaryNetworksDefault(code int) *CustomerPrototypeSetPrimarySecondaryNetworksDefault {
	return &CustomerPrototypeSetPrimarySecondaryNetworksDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypeSetPrimarySecondaryNetworksDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypeSetPrimarySecondaryNetworksDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype set primary secondary networks default response has a 2xx status code
func (o *CustomerPrototypeSetPrimarySecondaryNetworksDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype set primary secondary networks default response has a 3xx status code
func (o *CustomerPrototypeSetPrimarySecondaryNetworksDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype set primary secondary networks default response has a 4xx status code
func (o *CustomerPrototypeSetPrimarySecondaryNetworksDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype set primary secondary networks default response has a 5xx status code
func (o *CustomerPrototypeSetPrimarySecondaryNetworksDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype set primary secondary networks default response a status code equal to that given
func (o *CustomerPrototypeSetPrimarySecondaryNetworksDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype set primary secondary networks default response
func (o *CustomerPrototypeSetPrimarySecondaryNetworksDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypeSetPrimarySecondaryNetworksDefault) Error() string {
	return fmt.Sprintf("[PUT /Customers/{id}/locations/{locationId}/primarySecondaryNetworks][%d] Customer.prototype.setPrimarySecondaryNetworks default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeSetPrimarySecondaryNetworksDefault) String() string {
	return fmt.Sprintf("[PUT /Customers/{id}/locations/{locationId}/primarySecondaryNetworks][%d] Customer.prototype.setPrimarySecondaryNetworks default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeSetPrimarySecondaryNetworksDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypeSetPrimarySecondaryNetworksDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
