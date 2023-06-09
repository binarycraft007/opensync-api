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

// CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesWhitelistReader is a Reader for the CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesWhitelist structure.
type CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesWhitelistReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesWhitelistReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewCustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesWhitelistNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesWhitelistDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesWhitelistNoContent creates a CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesWhitelistNoContent with default headers values
func NewCustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesWhitelistNoContent() *CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesWhitelistNoContent {
	return &CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesWhitelistNoContent{}
}

/*
CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesWhitelistNoContent describes a response with status code 204, with default header values.

Request was successful
*/
type CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesWhitelistNoContent struct {
}

// IsSuccess returns true when this customer prototype delete from location security policy websites whitelist no content response has a 2xx status code
func (o *CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesWhitelistNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype delete from location security policy websites whitelist no content response has a 3xx status code
func (o *CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesWhitelistNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype delete from location security policy websites whitelist no content response has a 4xx status code
func (o *CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesWhitelistNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype delete from location security policy websites whitelist no content response has a 5xx status code
func (o *CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesWhitelistNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype delete from location security policy websites whitelist no content response a status code equal to that given
func (o *CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesWhitelistNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the customer prototype delete from location security policy websites whitelist no content response
func (o *CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesWhitelistNoContent) Code() int {
	return 204
}

func (o *CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesWhitelistNoContent) Error() string {
	return fmt.Sprintf("[DELETE /Customers/{id}/locations/{locationId}/securityPolicy/websites/whitelist/{dns}][%d] customerPrototypeDeleteFromLocationSecurityPolicyWebsitesWhitelistNoContent ", 204)
}

func (o *CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesWhitelistNoContent) String() string {
	return fmt.Sprintf("[DELETE /Customers/{id}/locations/{locationId}/securityPolicy/websites/whitelist/{dns}][%d] customerPrototypeDeleteFromLocationSecurityPolicyWebsitesWhitelistNoContent ", 204)
}

func (o *CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesWhitelistNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesWhitelistDefault creates a CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesWhitelistDefault with default headers values
func NewCustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesWhitelistDefault(code int) *CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesWhitelistDefault {
	return &CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesWhitelistDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesWhitelistDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesWhitelistDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype delete from location security policy websites whitelist default response has a 2xx status code
func (o *CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesWhitelistDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype delete from location security policy websites whitelist default response has a 3xx status code
func (o *CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesWhitelistDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype delete from location security policy websites whitelist default response has a 4xx status code
func (o *CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesWhitelistDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype delete from location security policy websites whitelist default response has a 5xx status code
func (o *CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesWhitelistDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype delete from location security policy websites whitelist default response a status code equal to that given
func (o *CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesWhitelistDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype delete from location security policy websites whitelist default response
func (o *CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesWhitelistDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesWhitelistDefault) Error() string {
	return fmt.Sprintf("[DELETE /Customers/{id}/locations/{locationId}/securityPolicy/websites/whitelist/{dns}][%d] Customer.prototype.deleteFromLocationSecurityPolicyWebsitesWhitelist default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesWhitelistDefault) String() string {
	return fmt.Sprintf("[DELETE /Customers/{id}/locations/{locationId}/securityPolicy/websites/whitelist/{dns}][%d] Customer.prototype.deleteFromLocationSecurityPolicyWebsitesWhitelist default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesWhitelistDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesWhitelistDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
