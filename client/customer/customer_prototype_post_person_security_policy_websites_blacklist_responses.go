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

// CustomerPrototypePostPersonSecurityPolicyWebsitesBlacklistReader is a Reader for the CustomerPrototypePostPersonSecurityPolicyWebsitesBlacklist structure.
type CustomerPrototypePostPersonSecurityPolicyWebsitesBlacklistReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypePostPersonSecurityPolicyWebsitesBlacklistReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypePostPersonSecurityPolicyWebsitesBlacklistOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypePostPersonSecurityPolicyWebsitesBlacklistDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypePostPersonSecurityPolicyWebsitesBlacklistOK creates a CustomerPrototypePostPersonSecurityPolicyWebsitesBlacklistOK with default headers values
func NewCustomerPrototypePostPersonSecurityPolicyWebsitesBlacklistOK() *CustomerPrototypePostPersonSecurityPolicyWebsitesBlacklistOK {
	return &CustomerPrototypePostPersonSecurityPolicyWebsitesBlacklistOK{}
}

/*
CustomerPrototypePostPersonSecurityPolicyWebsitesBlacklistOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypePostPersonSecurityPolicyWebsitesBlacklistOK struct {
	Payload interface{}
}

// IsSuccess returns true when this customer prototype post person security policy websites blacklist o k response has a 2xx status code
func (o *CustomerPrototypePostPersonSecurityPolicyWebsitesBlacklistOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype post person security policy websites blacklist o k response has a 3xx status code
func (o *CustomerPrototypePostPersonSecurityPolicyWebsitesBlacklistOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype post person security policy websites blacklist o k response has a 4xx status code
func (o *CustomerPrototypePostPersonSecurityPolicyWebsitesBlacklistOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype post person security policy websites blacklist o k response has a 5xx status code
func (o *CustomerPrototypePostPersonSecurityPolicyWebsitesBlacklistOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype post person security policy websites blacklist o k response a status code equal to that given
func (o *CustomerPrototypePostPersonSecurityPolicyWebsitesBlacklistOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype post person security policy websites blacklist o k response
func (o *CustomerPrototypePostPersonSecurityPolicyWebsitesBlacklistOK) Code() int {
	return 200
}

func (o *CustomerPrototypePostPersonSecurityPolicyWebsitesBlacklistOK) Error() string {
	return fmt.Sprintf("[POST /Customers/{id}/locations/{locationId}/persons/{personId}/securityPolicy/websites/blacklist][%d] customerPrototypePostPersonSecurityPolicyWebsitesBlacklistOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypePostPersonSecurityPolicyWebsitesBlacklistOK) String() string {
	return fmt.Sprintf("[POST /Customers/{id}/locations/{locationId}/persons/{personId}/securityPolicy/websites/blacklist][%d] customerPrototypePostPersonSecurityPolicyWebsitesBlacklistOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypePostPersonSecurityPolicyWebsitesBlacklistOK) GetPayload() interface{} {
	return o.Payload
}

func (o *CustomerPrototypePostPersonSecurityPolicyWebsitesBlacklistOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypePostPersonSecurityPolicyWebsitesBlacklistDefault creates a CustomerPrototypePostPersonSecurityPolicyWebsitesBlacklistDefault with default headers values
func NewCustomerPrototypePostPersonSecurityPolicyWebsitesBlacklistDefault(code int) *CustomerPrototypePostPersonSecurityPolicyWebsitesBlacklistDefault {
	return &CustomerPrototypePostPersonSecurityPolicyWebsitesBlacklistDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypePostPersonSecurityPolicyWebsitesBlacklistDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypePostPersonSecurityPolicyWebsitesBlacklistDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype post person security policy websites blacklist default response has a 2xx status code
func (o *CustomerPrototypePostPersonSecurityPolicyWebsitesBlacklistDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype post person security policy websites blacklist default response has a 3xx status code
func (o *CustomerPrototypePostPersonSecurityPolicyWebsitesBlacklistDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype post person security policy websites blacklist default response has a 4xx status code
func (o *CustomerPrototypePostPersonSecurityPolicyWebsitesBlacklistDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype post person security policy websites blacklist default response has a 5xx status code
func (o *CustomerPrototypePostPersonSecurityPolicyWebsitesBlacklistDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype post person security policy websites blacklist default response a status code equal to that given
func (o *CustomerPrototypePostPersonSecurityPolicyWebsitesBlacklistDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype post person security policy websites blacklist default response
func (o *CustomerPrototypePostPersonSecurityPolicyWebsitesBlacklistDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypePostPersonSecurityPolicyWebsitesBlacklistDefault) Error() string {
	return fmt.Sprintf("[POST /Customers/{id}/locations/{locationId}/persons/{personId}/securityPolicy/websites/blacklist][%d] Customer.prototype.postPersonSecurityPolicyWebsitesBlacklist default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePostPersonSecurityPolicyWebsitesBlacklistDefault) String() string {
	return fmt.Sprintf("[POST /Customers/{id}/locations/{locationId}/persons/{personId}/securityPolicy/websites/blacklist][%d] Customer.prototype.postPersonSecurityPolicyWebsitesBlacklist default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePostPersonSecurityPolicyWebsitesBlacklistDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypePostPersonSecurityPolicyWebsitesBlacklistDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
