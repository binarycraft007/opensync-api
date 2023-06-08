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

// CustomerPrototypePostLocationSecurityPolicyWebsitesWhitelistReader is a Reader for the CustomerPrototypePostLocationSecurityPolicyWebsitesWhitelist structure.
type CustomerPrototypePostLocationSecurityPolicyWebsitesWhitelistReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypePostLocationSecurityPolicyWebsitesWhitelistReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypePostLocationSecurityPolicyWebsitesWhitelistOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypePostLocationSecurityPolicyWebsitesWhitelistDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypePostLocationSecurityPolicyWebsitesWhitelistOK creates a CustomerPrototypePostLocationSecurityPolicyWebsitesWhitelistOK with default headers values
func NewCustomerPrototypePostLocationSecurityPolicyWebsitesWhitelistOK() *CustomerPrototypePostLocationSecurityPolicyWebsitesWhitelistOK {
	return &CustomerPrototypePostLocationSecurityPolicyWebsitesWhitelistOK{}
}

/*
CustomerPrototypePostLocationSecurityPolicyWebsitesWhitelistOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypePostLocationSecurityPolicyWebsitesWhitelistOK struct {
	Payload interface{}
}

// IsSuccess returns true when this customer prototype post location security policy websites whitelist o k response has a 2xx status code
func (o *CustomerPrototypePostLocationSecurityPolicyWebsitesWhitelistOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype post location security policy websites whitelist o k response has a 3xx status code
func (o *CustomerPrototypePostLocationSecurityPolicyWebsitesWhitelistOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype post location security policy websites whitelist o k response has a 4xx status code
func (o *CustomerPrototypePostLocationSecurityPolicyWebsitesWhitelistOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype post location security policy websites whitelist o k response has a 5xx status code
func (o *CustomerPrototypePostLocationSecurityPolicyWebsitesWhitelistOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype post location security policy websites whitelist o k response a status code equal to that given
func (o *CustomerPrototypePostLocationSecurityPolicyWebsitesWhitelistOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype post location security policy websites whitelist o k response
func (o *CustomerPrototypePostLocationSecurityPolicyWebsitesWhitelistOK) Code() int {
	return 200
}

func (o *CustomerPrototypePostLocationSecurityPolicyWebsitesWhitelistOK) Error() string {
	return fmt.Sprintf("[POST /Customers/{id}/locations/{locationId}/securityPolicy/websites/whitelist][%d] customerPrototypePostLocationSecurityPolicyWebsitesWhitelistOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypePostLocationSecurityPolicyWebsitesWhitelistOK) String() string {
	return fmt.Sprintf("[POST /Customers/{id}/locations/{locationId}/securityPolicy/websites/whitelist][%d] customerPrototypePostLocationSecurityPolicyWebsitesWhitelistOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypePostLocationSecurityPolicyWebsitesWhitelistOK) GetPayload() interface{} {
	return o.Payload
}

func (o *CustomerPrototypePostLocationSecurityPolicyWebsitesWhitelistOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypePostLocationSecurityPolicyWebsitesWhitelistDefault creates a CustomerPrototypePostLocationSecurityPolicyWebsitesWhitelistDefault with default headers values
func NewCustomerPrototypePostLocationSecurityPolicyWebsitesWhitelistDefault(code int) *CustomerPrototypePostLocationSecurityPolicyWebsitesWhitelistDefault {
	return &CustomerPrototypePostLocationSecurityPolicyWebsitesWhitelistDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypePostLocationSecurityPolicyWebsitesWhitelistDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypePostLocationSecurityPolicyWebsitesWhitelistDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype post location security policy websites whitelist default response has a 2xx status code
func (o *CustomerPrototypePostLocationSecurityPolicyWebsitesWhitelistDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype post location security policy websites whitelist default response has a 3xx status code
func (o *CustomerPrototypePostLocationSecurityPolicyWebsitesWhitelistDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype post location security policy websites whitelist default response has a 4xx status code
func (o *CustomerPrototypePostLocationSecurityPolicyWebsitesWhitelistDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype post location security policy websites whitelist default response has a 5xx status code
func (o *CustomerPrototypePostLocationSecurityPolicyWebsitesWhitelistDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype post location security policy websites whitelist default response a status code equal to that given
func (o *CustomerPrototypePostLocationSecurityPolicyWebsitesWhitelistDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype post location security policy websites whitelist default response
func (o *CustomerPrototypePostLocationSecurityPolicyWebsitesWhitelistDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypePostLocationSecurityPolicyWebsitesWhitelistDefault) Error() string {
	return fmt.Sprintf("[POST /Customers/{id}/locations/{locationId}/securityPolicy/websites/whitelist][%d] Customer.prototype.postLocationSecurityPolicyWebsitesWhitelist default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePostLocationSecurityPolicyWebsitesWhitelistDefault) String() string {
	return fmt.Sprintf("[POST /Customers/{id}/locations/{locationId}/securityPolicy/websites/whitelist][%d] Customer.prototype.postLocationSecurityPolicyWebsitesWhitelist default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePostLocationSecurityPolicyWebsitesWhitelistDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypePostLocationSecurityPolicyWebsitesWhitelistDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}