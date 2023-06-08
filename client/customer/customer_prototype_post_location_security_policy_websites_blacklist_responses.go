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

// CustomerPrototypePostLocationSecurityPolicyWebsitesBlacklistReader is a Reader for the CustomerPrototypePostLocationSecurityPolicyWebsitesBlacklist structure.
type CustomerPrototypePostLocationSecurityPolicyWebsitesBlacklistReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypePostLocationSecurityPolicyWebsitesBlacklistReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypePostLocationSecurityPolicyWebsitesBlacklistOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypePostLocationSecurityPolicyWebsitesBlacklistDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypePostLocationSecurityPolicyWebsitesBlacklistOK creates a CustomerPrototypePostLocationSecurityPolicyWebsitesBlacklistOK with default headers values
func NewCustomerPrototypePostLocationSecurityPolicyWebsitesBlacklistOK() *CustomerPrototypePostLocationSecurityPolicyWebsitesBlacklistOK {
	return &CustomerPrototypePostLocationSecurityPolicyWebsitesBlacklistOK{}
}

/*
CustomerPrototypePostLocationSecurityPolicyWebsitesBlacklistOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypePostLocationSecurityPolicyWebsitesBlacklistOK struct {
	Payload interface{}
}

// IsSuccess returns true when this customer prototype post location security policy websites blacklist o k response has a 2xx status code
func (o *CustomerPrototypePostLocationSecurityPolicyWebsitesBlacklistOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype post location security policy websites blacklist o k response has a 3xx status code
func (o *CustomerPrototypePostLocationSecurityPolicyWebsitesBlacklistOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype post location security policy websites blacklist o k response has a 4xx status code
func (o *CustomerPrototypePostLocationSecurityPolicyWebsitesBlacklistOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype post location security policy websites blacklist o k response has a 5xx status code
func (o *CustomerPrototypePostLocationSecurityPolicyWebsitesBlacklistOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype post location security policy websites blacklist o k response a status code equal to that given
func (o *CustomerPrototypePostLocationSecurityPolicyWebsitesBlacklistOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype post location security policy websites blacklist o k response
func (o *CustomerPrototypePostLocationSecurityPolicyWebsitesBlacklistOK) Code() int {
	return 200
}

func (o *CustomerPrototypePostLocationSecurityPolicyWebsitesBlacklistOK) Error() string {
	return fmt.Sprintf("[POST /Customers/{id}/locations/{locationId}/securityPolicy/websites/blacklist][%d] customerPrototypePostLocationSecurityPolicyWebsitesBlacklistOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypePostLocationSecurityPolicyWebsitesBlacklistOK) String() string {
	return fmt.Sprintf("[POST /Customers/{id}/locations/{locationId}/securityPolicy/websites/blacklist][%d] customerPrototypePostLocationSecurityPolicyWebsitesBlacklistOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypePostLocationSecurityPolicyWebsitesBlacklistOK) GetPayload() interface{} {
	return o.Payload
}

func (o *CustomerPrototypePostLocationSecurityPolicyWebsitesBlacklistOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypePostLocationSecurityPolicyWebsitesBlacklistDefault creates a CustomerPrototypePostLocationSecurityPolicyWebsitesBlacklistDefault with default headers values
func NewCustomerPrototypePostLocationSecurityPolicyWebsitesBlacklistDefault(code int) *CustomerPrototypePostLocationSecurityPolicyWebsitesBlacklistDefault {
	return &CustomerPrototypePostLocationSecurityPolicyWebsitesBlacklistDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypePostLocationSecurityPolicyWebsitesBlacklistDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypePostLocationSecurityPolicyWebsitesBlacklistDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype post location security policy websites blacklist default response has a 2xx status code
func (o *CustomerPrototypePostLocationSecurityPolicyWebsitesBlacklistDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype post location security policy websites blacklist default response has a 3xx status code
func (o *CustomerPrototypePostLocationSecurityPolicyWebsitesBlacklistDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype post location security policy websites blacklist default response has a 4xx status code
func (o *CustomerPrototypePostLocationSecurityPolicyWebsitesBlacklistDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype post location security policy websites blacklist default response has a 5xx status code
func (o *CustomerPrototypePostLocationSecurityPolicyWebsitesBlacklistDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype post location security policy websites blacklist default response a status code equal to that given
func (o *CustomerPrototypePostLocationSecurityPolicyWebsitesBlacklistDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype post location security policy websites blacklist default response
func (o *CustomerPrototypePostLocationSecurityPolicyWebsitesBlacklistDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypePostLocationSecurityPolicyWebsitesBlacklistDefault) Error() string {
	return fmt.Sprintf("[POST /Customers/{id}/locations/{locationId}/securityPolicy/websites/blacklist][%d] Customer.prototype.postLocationSecurityPolicyWebsitesBlacklist default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePostLocationSecurityPolicyWebsitesBlacklistDefault) String() string {
	return fmt.Sprintf("[POST /Customers/{id}/locations/{locationId}/securityPolicy/websites/blacklist][%d] Customer.prototype.postLocationSecurityPolicyWebsitesBlacklist default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePostLocationSecurityPolicyWebsitesBlacklistDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypePostLocationSecurityPolicyWebsitesBlacklistDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
