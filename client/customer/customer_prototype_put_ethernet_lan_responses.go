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

// CustomerPrototypePutEthernetLanReader is a Reader for the CustomerPrototypePutEthernetLan structure.
type CustomerPrototypePutEthernetLanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypePutEthernetLanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewCustomerPrototypePutEthernetLanAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypePutEthernetLanDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypePutEthernetLanAccepted creates a CustomerPrototypePutEthernetLanAccepted with default headers values
func NewCustomerPrototypePutEthernetLanAccepted() *CustomerPrototypePutEthernetLanAccepted {
	return &CustomerPrototypePutEthernetLanAccepted{}
}

/*
CustomerPrototypePutEthernetLanAccepted describes a response with status code 202, with default header values.

Request was successful
*/
type CustomerPrototypePutEthernetLanAccepted struct {
}

// IsSuccess returns true when this customer prototype put ethernet lan accepted response has a 2xx status code
func (o *CustomerPrototypePutEthernetLanAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype put ethernet lan accepted response has a 3xx status code
func (o *CustomerPrototypePutEthernetLanAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype put ethernet lan accepted response has a 4xx status code
func (o *CustomerPrototypePutEthernetLanAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype put ethernet lan accepted response has a 5xx status code
func (o *CustomerPrototypePutEthernetLanAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype put ethernet lan accepted response a status code equal to that given
func (o *CustomerPrototypePutEthernetLanAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the customer prototype put ethernet lan accepted response
func (o *CustomerPrototypePutEthernetLanAccepted) Code() int {
	return 202
}

func (o *CustomerPrototypePutEthernetLanAccepted) Error() string {
	return fmt.Sprintf("[PUT /Customers/{id}/locations/{locationId}/nodes/{nodeId}/ethernetLan][%d] customerPrototypePutEthernetLanAccepted ", 202)
}

func (o *CustomerPrototypePutEthernetLanAccepted) String() string {
	return fmt.Sprintf("[PUT /Customers/{id}/locations/{locationId}/nodes/{nodeId}/ethernetLan][%d] customerPrototypePutEthernetLanAccepted ", 202)
}

func (o *CustomerPrototypePutEthernetLanAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCustomerPrototypePutEthernetLanDefault creates a CustomerPrototypePutEthernetLanDefault with default headers values
func NewCustomerPrototypePutEthernetLanDefault(code int) *CustomerPrototypePutEthernetLanDefault {
	return &CustomerPrototypePutEthernetLanDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypePutEthernetLanDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypePutEthernetLanDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype put ethernet lan default response has a 2xx status code
func (o *CustomerPrototypePutEthernetLanDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype put ethernet lan default response has a 3xx status code
func (o *CustomerPrototypePutEthernetLanDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype put ethernet lan default response has a 4xx status code
func (o *CustomerPrototypePutEthernetLanDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype put ethernet lan default response has a 5xx status code
func (o *CustomerPrototypePutEthernetLanDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype put ethernet lan default response a status code equal to that given
func (o *CustomerPrototypePutEthernetLanDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype put ethernet lan default response
func (o *CustomerPrototypePutEthernetLanDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypePutEthernetLanDefault) Error() string {
	return fmt.Sprintf("[PUT /Customers/{id}/locations/{locationId}/nodes/{nodeId}/ethernetLan][%d] Customer.prototype.putEthernetLan default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePutEthernetLanDefault) String() string {
	return fmt.Sprintf("[PUT /Customers/{id}/locations/{locationId}/nodes/{nodeId}/ethernetLan][%d] Customer.prototype.putEthernetLan default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePutEthernetLanDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypePutEthernetLanDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
