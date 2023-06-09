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

// CustomerPrototypePatchPersonSecurityPolicyReader is a Reader for the CustomerPrototypePatchPersonSecurityPolicy structure.
type CustomerPrototypePatchPersonSecurityPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypePatchPersonSecurityPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypePatchPersonSecurityPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypePatchPersonSecurityPolicyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypePatchPersonSecurityPolicyOK creates a CustomerPrototypePatchPersonSecurityPolicyOK with default headers values
func NewCustomerPrototypePatchPersonSecurityPolicyOK() *CustomerPrototypePatchPersonSecurityPolicyOK {
	return &CustomerPrototypePatchPersonSecurityPolicyOK{}
}

/*
CustomerPrototypePatchPersonSecurityPolicyOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypePatchPersonSecurityPolicyOK struct {
	Payload interface{}
}

// IsSuccess returns true when this customer prototype patch person security policy o k response has a 2xx status code
func (o *CustomerPrototypePatchPersonSecurityPolicyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype patch person security policy o k response has a 3xx status code
func (o *CustomerPrototypePatchPersonSecurityPolicyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype patch person security policy o k response has a 4xx status code
func (o *CustomerPrototypePatchPersonSecurityPolicyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype patch person security policy o k response has a 5xx status code
func (o *CustomerPrototypePatchPersonSecurityPolicyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype patch person security policy o k response a status code equal to that given
func (o *CustomerPrototypePatchPersonSecurityPolicyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype patch person security policy o k response
func (o *CustomerPrototypePatchPersonSecurityPolicyOK) Code() int {
	return 200
}

func (o *CustomerPrototypePatchPersonSecurityPolicyOK) Error() string {
	return fmt.Sprintf("[PATCH /Customers/{id}/locations/{locationId}/persons/{personId}/securityPolicy][%d] customerPrototypePatchPersonSecurityPolicyOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypePatchPersonSecurityPolicyOK) String() string {
	return fmt.Sprintf("[PATCH /Customers/{id}/locations/{locationId}/persons/{personId}/securityPolicy][%d] customerPrototypePatchPersonSecurityPolicyOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypePatchPersonSecurityPolicyOK) GetPayload() interface{} {
	return o.Payload
}

func (o *CustomerPrototypePatchPersonSecurityPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypePatchPersonSecurityPolicyDefault creates a CustomerPrototypePatchPersonSecurityPolicyDefault with default headers values
func NewCustomerPrototypePatchPersonSecurityPolicyDefault(code int) *CustomerPrototypePatchPersonSecurityPolicyDefault {
	return &CustomerPrototypePatchPersonSecurityPolicyDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypePatchPersonSecurityPolicyDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypePatchPersonSecurityPolicyDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype patch person security policy default response has a 2xx status code
func (o *CustomerPrototypePatchPersonSecurityPolicyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype patch person security policy default response has a 3xx status code
func (o *CustomerPrototypePatchPersonSecurityPolicyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype patch person security policy default response has a 4xx status code
func (o *CustomerPrototypePatchPersonSecurityPolicyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype patch person security policy default response has a 5xx status code
func (o *CustomerPrototypePatchPersonSecurityPolicyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype patch person security policy default response a status code equal to that given
func (o *CustomerPrototypePatchPersonSecurityPolicyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype patch person security policy default response
func (o *CustomerPrototypePatchPersonSecurityPolicyDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypePatchPersonSecurityPolicyDefault) Error() string {
	return fmt.Sprintf("[PATCH /Customers/{id}/locations/{locationId}/persons/{personId}/securityPolicy][%d] Customer.prototype.patchPersonSecurityPolicy default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePatchPersonSecurityPolicyDefault) String() string {
	return fmt.Sprintf("[PATCH /Customers/{id}/locations/{locationId}/persons/{personId}/securityPolicy][%d] Customer.prototype.patchPersonSecurityPolicy default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePatchPersonSecurityPolicyDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypePatchPersonSecurityPolicyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
