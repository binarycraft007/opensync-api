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

// CustomerPrototypePutLocationSecurityPolicyOHPDeviceUUIDMappingReader is a Reader for the CustomerPrototypePutLocationSecurityPolicyOHPDeviceUUIDMapping structure.
type CustomerPrototypePutLocationSecurityPolicyOHPDeviceUUIDMappingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypePutLocationSecurityPolicyOHPDeviceUUIDMappingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypePutLocationSecurityPolicyOHPDeviceUUIDMappingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypePutLocationSecurityPolicyOHPDeviceUUIDMappingDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypePutLocationSecurityPolicyOHPDeviceUUIDMappingOK creates a CustomerPrototypePutLocationSecurityPolicyOHPDeviceUUIDMappingOK with default headers values
func NewCustomerPrototypePutLocationSecurityPolicyOHPDeviceUUIDMappingOK() *CustomerPrototypePutLocationSecurityPolicyOHPDeviceUUIDMappingOK {
	return &CustomerPrototypePutLocationSecurityPolicyOHPDeviceUUIDMappingOK{}
}

/*
CustomerPrototypePutLocationSecurityPolicyOHPDeviceUUIDMappingOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypePutLocationSecurityPolicyOHPDeviceUUIDMappingOK struct {
	Payload interface{}
}

// IsSuccess returns true when this customer prototype put location security policy o h p device Uuid mapping o k response has a 2xx status code
func (o *CustomerPrototypePutLocationSecurityPolicyOHPDeviceUUIDMappingOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype put location security policy o h p device Uuid mapping o k response has a 3xx status code
func (o *CustomerPrototypePutLocationSecurityPolicyOHPDeviceUUIDMappingOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype put location security policy o h p device Uuid mapping o k response has a 4xx status code
func (o *CustomerPrototypePutLocationSecurityPolicyOHPDeviceUUIDMappingOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype put location security policy o h p device Uuid mapping o k response has a 5xx status code
func (o *CustomerPrototypePutLocationSecurityPolicyOHPDeviceUUIDMappingOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype put location security policy o h p device Uuid mapping o k response a status code equal to that given
func (o *CustomerPrototypePutLocationSecurityPolicyOHPDeviceUUIDMappingOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype put location security policy o h p device Uuid mapping o k response
func (o *CustomerPrototypePutLocationSecurityPolicyOHPDeviceUUIDMappingOK) Code() int {
	return 200
}

func (o *CustomerPrototypePutLocationSecurityPolicyOHPDeviceUUIDMappingOK) Error() string {
	return fmt.Sprintf("[PUT /Customers/{id}/locations/{locationId}/securityPolicy/ohp/deviceUuid][%d] customerPrototypePutLocationSecurityPolicyOHPDeviceUuidMappingOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypePutLocationSecurityPolicyOHPDeviceUUIDMappingOK) String() string {
	return fmt.Sprintf("[PUT /Customers/{id}/locations/{locationId}/securityPolicy/ohp/deviceUuid][%d] customerPrototypePutLocationSecurityPolicyOHPDeviceUuidMappingOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypePutLocationSecurityPolicyOHPDeviceUUIDMappingOK) GetPayload() interface{} {
	return o.Payload
}

func (o *CustomerPrototypePutLocationSecurityPolicyOHPDeviceUUIDMappingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypePutLocationSecurityPolicyOHPDeviceUUIDMappingDefault creates a CustomerPrototypePutLocationSecurityPolicyOHPDeviceUUIDMappingDefault with default headers values
func NewCustomerPrototypePutLocationSecurityPolicyOHPDeviceUUIDMappingDefault(code int) *CustomerPrototypePutLocationSecurityPolicyOHPDeviceUUIDMappingDefault {
	return &CustomerPrototypePutLocationSecurityPolicyOHPDeviceUUIDMappingDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypePutLocationSecurityPolicyOHPDeviceUUIDMappingDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypePutLocationSecurityPolicyOHPDeviceUUIDMappingDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype put location security policy o h p device Uuid mapping default response has a 2xx status code
func (o *CustomerPrototypePutLocationSecurityPolicyOHPDeviceUUIDMappingDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype put location security policy o h p device Uuid mapping default response has a 3xx status code
func (o *CustomerPrototypePutLocationSecurityPolicyOHPDeviceUUIDMappingDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype put location security policy o h p device Uuid mapping default response has a 4xx status code
func (o *CustomerPrototypePutLocationSecurityPolicyOHPDeviceUUIDMappingDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype put location security policy o h p device Uuid mapping default response has a 5xx status code
func (o *CustomerPrototypePutLocationSecurityPolicyOHPDeviceUUIDMappingDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype put location security policy o h p device Uuid mapping default response a status code equal to that given
func (o *CustomerPrototypePutLocationSecurityPolicyOHPDeviceUUIDMappingDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype put location security policy o h p device Uuid mapping default response
func (o *CustomerPrototypePutLocationSecurityPolicyOHPDeviceUUIDMappingDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypePutLocationSecurityPolicyOHPDeviceUUIDMappingDefault) Error() string {
	return fmt.Sprintf("[PUT /Customers/{id}/locations/{locationId}/securityPolicy/ohp/deviceUuid][%d] Customer.prototype.putLocationSecurityPolicyOHPDeviceUuidMapping default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePutLocationSecurityPolicyOHPDeviceUUIDMappingDefault) String() string {
	return fmt.Sprintf("[PUT /Customers/{id}/locations/{locationId}/securityPolicy/ohp/deviceUuid][%d] Customer.prototype.putLocationSecurityPolicyOHPDeviceUuidMapping default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePutLocationSecurityPolicyOHPDeviceUUIDMappingDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypePutLocationSecurityPolicyOHPDeviceUUIDMappingDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}