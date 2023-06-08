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

// CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesBlacklistReader is a Reader for the CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesBlacklist structure.
type CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesBlacklistReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesBlacklistReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewCustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesBlacklistNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesBlacklistDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesBlacklistNoContent creates a CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesBlacklistNoContent with default headers values
func NewCustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesBlacklistNoContent() *CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesBlacklistNoContent {
	return &CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesBlacklistNoContent{}
}

/*
CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesBlacklistNoContent describes a response with status code 204, with default header values.

Request was successful
*/
type CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesBlacklistNoContent struct {
}

// IsSuccess returns true when this customer prototype delete from location security policy websites blacklist no content response has a 2xx status code
func (o *CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesBlacklistNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype delete from location security policy websites blacklist no content response has a 3xx status code
func (o *CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesBlacklistNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype delete from location security policy websites blacklist no content response has a 4xx status code
func (o *CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesBlacklistNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype delete from location security policy websites blacklist no content response has a 5xx status code
func (o *CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesBlacklistNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype delete from location security policy websites blacklist no content response a status code equal to that given
func (o *CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesBlacklistNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the customer prototype delete from location security policy websites blacklist no content response
func (o *CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesBlacklistNoContent) Code() int {
	return 204
}

func (o *CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesBlacklistNoContent) Error() string {
	return fmt.Sprintf("[DELETE /Customers/{id}/locations/{locationId}/securityPolicy/websites/blacklist/{dns}][%d] customerPrototypeDeleteFromLocationSecurityPolicyWebsitesBlacklistNoContent ", 204)
}

func (o *CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesBlacklistNoContent) String() string {
	return fmt.Sprintf("[DELETE /Customers/{id}/locations/{locationId}/securityPolicy/websites/blacklist/{dns}][%d] customerPrototypeDeleteFromLocationSecurityPolicyWebsitesBlacklistNoContent ", 204)
}

func (o *CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesBlacklistNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesBlacklistDefault creates a CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesBlacklistDefault with default headers values
func NewCustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesBlacklistDefault(code int) *CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesBlacklistDefault {
	return &CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesBlacklistDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesBlacklistDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesBlacklistDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype delete from location security policy websites blacklist default response has a 2xx status code
func (o *CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesBlacklistDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype delete from location security policy websites blacklist default response has a 3xx status code
func (o *CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesBlacklistDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype delete from location security policy websites blacklist default response has a 4xx status code
func (o *CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesBlacklistDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype delete from location security policy websites blacklist default response has a 5xx status code
func (o *CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesBlacklistDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype delete from location security policy websites blacklist default response a status code equal to that given
func (o *CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesBlacklistDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype delete from location security policy websites blacklist default response
func (o *CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesBlacklistDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesBlacklistDefault) Error() string {
	return fmt.Sprintf("[DELETE /Customers/{id}/locations/{locationId}/securityPolicy/websites/blacklist/{dns}][%d] Customer.prototype.deleteFromLocationSecurityPolicyWebsitesBlacklist default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesBlacklistDefault) String() string {
	return fmt.Sprintf("[DELETE /Customers/{id}/locations/{locationId}/securityPolicy/websites/blacklist/{dns}][%d] Customer.prototype.deleteFromLocationSecurityPolicyWebsitesBlacklist default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesBlacklistDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesBlacklistDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
