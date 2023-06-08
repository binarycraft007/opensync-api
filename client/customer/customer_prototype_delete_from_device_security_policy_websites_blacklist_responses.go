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

// CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesBlacklistReader is a Reader for the CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesBlacklist structure.
type CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesBlacklistReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesBlacklistReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewCustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesBlacklistNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesBlacklistDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesBlacklistNoContent creates a CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesBlacklistNoContent with default headers values
func NewCustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesBlacklistNoContent() *CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesBlacklistNoContent {
	return &CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesBlacklistNoContent{}
}

/*
CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesBlacklistNoContent describes a response with status code 204, with default header values.

Request was successful
*/
type CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesBlacklistNoContent struct {
}

// IsSuccess returns true when this customer prototype delete from device security policy websites blacklist no content response has a 2xx status code
func (o *CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesBlacklistNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype delete from device security policy websites blacklist no content response has a 3xx status code
func (o *CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesBlacklistNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype delete from device security policy websites blacklist no content response has a 4xx status code
func (o *CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesBlacklistNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype delete from device security policy websites blacklist no content response has a 5xx status code
func (o *CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesBlacklistNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype delete from device security policy websites blacklist no content response a status code equal to that given
func (o *CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesBlacklistNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the customer prototype delete from device security policy websites blacklist no content response
func (o *CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesBlacklistNoContent) Code() int {
	return 204
}

func (o *CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesBlacklistNoContent) Error() string {
	return fmt.Sprintf("[DELETE /Customers/{id}/locations/{locationId}/devices/{mac}/securityPolicy/websites/blacklist/{dns}][%d] customerPrototypeDeleteFromDeviceSecurityPolicyWebsitesBlacklistNoContent ", 204)
}

func (o *CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesBlacklistNoContent) String() string {
	return fmt.Sprintf("[DELETE /Customers/{id}/locations/{locationId}/devices/{mac}/securityPolicy/websites/blacklist/{dns}][%d] customerPrototypeDeleteFromDeviceSecurityPolicyWebsitesBlacklistNoContent ", 204)
}

func (o *CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesBlacklistNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesBlacklistDefault creates a CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesBlacklistDefault with default headers values
func NewCustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesBlacklistDefault(code int) *CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesBlacklistDefault {
	return &CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesBlacklistDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesBlacklistDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesBlacklistDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype delete from device security policy websites blacklist default response has a 2xx status code
func (o *CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesBlacklistDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype delete from device security policy websites blacklist default response has a 3xx status code
func (o *CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesBlacklistDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype delete from device security policy websites blacklist default response has a 4xx status code
func (o *CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesBlacklistDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype delete from device security policy websites blacklist default response has a 5xx status code
func (o *CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesBlacklistDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype delete from device security policy websites blacklist default response a status code equal to that given
func (o *CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesBlacklistDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype delete from device security policy websites blacklist default response
func (o *CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesBlacklistDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesBlacklistDefault) Error() string {
	return fmt.Sprintf("[DELETE /Customers/{id}/locations/{locationId}/devices/{mac}/securityPolicy/websites/blacklist/{dns}][%d] Customer.prototype.deleteFromDeviceSecurityPolicyWebsitesBlacklist default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesBlacklistDefault) String() string {
	return fmt.Sprintf("[DELETE /Customers/{id}/locations/{locationId}/devices/{mac}/securityPolicy/websites/blacklist/{dns}][%d] Customer.prototype.deleteFromDeviceSecurityPolicyWebsitesBlacklist default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesBlacklistDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesBlacklistDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
