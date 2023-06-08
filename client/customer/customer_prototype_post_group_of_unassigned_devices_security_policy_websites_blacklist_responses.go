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

// CustomerPrototypePostGroupOfUnassignedDevicesSecurityPolicyWebsitesBlacklistReader is a Reader for the CustomerPrototypePostGroupOfUnassignedDevicesSecurityPolicyWebsitesBlacklist structure.
type CustomerPrototypePostGroupOfUnassignedDevicesSecurityPolicyWebsitesBlacklistReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypePostGroupOfUnassignedDevicesSecurityPolicyWebsitesBlacklistReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewCustomerPrototypePostGroupOfUnassignedDevicesSecurityPolicyWebsitesBlacklistAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypePostGroupOfUnassignedDevicesSecurityPolicyWebsitesBlacklistDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypePostGroupOfUnassignedDevicesSecurityPolicyWebsitesBlacklistAccepted creates a CustomerPrototypePostGroupOfUnassignedDevicesSecurityPolicyWebsitesBlacklistAccepted with default headers values
func NewCustomerPrototypePostGroupOfUnassignedDevicesSecurityPolicyWebsitesBlacklistAccepted() *CustomerPrototypePostGroupOfUnassignedDevicesSecurityPolicyWebsitesBlacklistAccepted {
	return &CustomerPrototypePostGroupOfUnassignedDevicesSecurityPolicyWebsitesBlacklistAccepted{}
}

/*
CustomerPrototypePostGroupOfUnassignedDevicesSecurityPolicyWebsitesBlacklistAccepted describes a response with status code 202, with default header values.

Request was successful
*/
type CustomerPrototypePostGroupOfUnassignedDevicesSecurityPolicyWebsitesBlacklistAccepted struct {
	Payload interface{}
}

// IsSuccess returns true when this customer prototype post group of unassigned devices security policy websites blacklist accepted response has a 2xx status code
func (o *CustomerPrototypePostGroupOfUnassignedDevicesSecurityPolicyWebsitesBlacklistAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype post group of unassigned devices security policy websites blacklist accepted response has a 3xx status code
func (o *CustomerPrototypePostGroupOfUnassignedDevicesSecurityPolicyWebsitesBlacklistAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype post group of unassigned devices security policy websites blacklist accepted response has a 4xx status code
func (o *CustomerPrototypePostGroupOfUnassignedDevicesSecurityPolicyWebsitesBlacklistAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype post group of unassigned devices security policy websites blacklist accepted response has a 5xx status code
func (o *CustomerPrototypePostGroupOfUnassignedDevicesSecurityPolicyWebsitesBlacklistAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype post group of unassigned devices security policy websites blacklist accepted response a status code equal to that given
func (o *CustomerPrototypePostGroupOfUnassignedDevicesSecurityPolicyWebsitesBlacklistAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the customer prototype post group of unassigned devices security policy websites blacklist accepted response
func (o *CustomerPrototypePostGroupOfUnassignedDevicesSecurityPolicyWebsitesBlacklistAccepted) Code() int {
	return 202
}

func (o *CustomerPrototypePostGroupOfUnassignedDevicesSecurityPolicyWebsitesBlacklistAccepted) Error() string {
	return fmt.Sprintf("[POST /Customers/{id}/locations/{locationId}/groupOfUnassignedDevices/securityPolicy/websites/blacklist][%d] customerPrototypePostGroupOfUnassignedDevicesSecurityPolicyWebsitesBlacklistAccepted  %+v", 202, o.Payload)
}

func (o *CustomerPrototypePostGroupOfUnassignedDevicesSecurityPolicyWebsitesBlacklistAccepted) String() string {
	return fmt.Sprintf("[POST /Customers/{id}/locations/{locationId}/groupOfUnassignedDevices/securityPolicy/websites/blacklist][%d] customerPrototypePostGroupOfUnassignedDevicesSecurityPolicyWebsitesBlacklistAccepted  %+v", 202, o.Payload)
}

func (o *CustomerPrototypePostGroupOfUnassignedDevicesSecurityPolicyWebsitesBlacklistAccepted) GetPayload() interface{} {
	return o.Payload
}

func (o *CustomerPrototypePostGroupOfUnassignedDevicesSecurityPolicyWebsitesBlacklistAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypePostGroupOfUnassignedDevicesSecurityPolicyWebsitesBlacklistDefault creates a CustomerPrototypePostGroupOfUnassignedDevicesSecurityPolicyWebsitesBlacklistDefault with default headers values
func NewCustomerPrototypePostGroupOfUnassignedDevicesSecurityPolicyWebsitesBlacklistDefault(code int) *CustomerPrototypePostGroupOfUnassignedDevicesSecurityPolicyWebsitesBlacklistDefault {
	return &CustomerPrototypePostGroupOfUnassignedDevicesSecurityPolicyWebsitesBlacklistDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypePostGroupOfUnassignedDevicesSecurityPolicyWebsitesBlacklistDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypePostGroupOfUnassignedDevicesSecurityPolicyWebsitesBlacklistDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype post group of unassigned devices security policy websites blacklist default response has a 2xx status code
func (o *CustomerPrototypePostGroupOfUnassignedDevicesSecurityPolicyWebsitesBlacklistDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype post group of unassigned devices security policy websites blacklist default response has a 3xx status code
func (o *CustomerPrototypePostGroupOfUnassignedDevicesSecurityPolicyWebsitesBlacklistDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype post group of unassigned devices security policy websites blacklist default response has a 4xx status code
func (o *CustomerPrototypePostGroupOfUnassignedDevicesSecurityPolicyWebsitesBlacklistDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype post group of unassigned devices security policy websites blacklist default response has a 5xx status code
func (o *CustomerPrototypePostGroupOfUnassignedDevicesSecurityPolicyWebsitesBlacklistDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype post group of unassigned devices security policy websites blacklist default response a status code equal to that given
func (o *CustomerPrototypePostGroupOfUnassignedDevicesSecurityPolicyWebsitesBlacklistDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype post group of unassigned devices security policy websites blacklist default response
func (o *CustomerPrototypePostGroupOfUnassignedDevicesSecurityPolicyWebsitesBlacklistDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypePostGroupOfUnassignedDevicesSecurityPolicyWebsitesBlacklistDefault) Error() string {
	return fmt.Sprintf("[POST /Customers/{id}/locations/{locationId}/groupOfUnassignedDevices/securityPolicy/websites/blacklist][%d] Customer.prototype.postGroupOfUnassignedDevicesSecurityPolicyWebsitesBlacklist default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePostGroupOfUnassignedDevicesSecurityPolicyWebsitesBlacklistDefault) String() string {
	return fmt.Sprintf("[POST /Customers/{id}/locations/{locationId}/groupOfUnassignedDevices/securityPolicy/websites/blacklist][%d] Customer.prototype.postGroupOfUnassignedDevicesSecurityPolicyWebsitesBlacklist default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePostGroupOfUnassignedDevicesSecurityPolicyWebsitesBlacklistDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypePostGroupOfUnassignedDevicesSecurityPolicyWebsitesBlacklistDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
