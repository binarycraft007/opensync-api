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

// CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyExperienceReader is a Reader for the CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyExperience structure.
type CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyExperienceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyExperienceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewCustomerPrototypeDeleteDeviceSecurityPolicyAnomalyExperienceNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypeDeleteDeviceSecurityPolicyAnomalyExperienceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypeDeleteDeviceSecurityPolicyAnomalyExperienceNoContent creates a CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyExperienceNoContent with default headers values
func NewCustomerPrototypeDeleteDeviceSecurityPolicyAnomalyExperienceNoContent() *CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyExperienceNoContent {
	return &CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyExperienceNoContent{}
}

/*
CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyExperienceNoContent describes a response with status code 204, with default header values.

Request was successful
*/
type CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyExperienceNoContent struct {
}

// IsSuccess returns true when this customer prototype delete device security policy anomaly experience no content response has a 2xx status code
func (o *CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyExperienceNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype delete device security policy anomaly experience no content response has a 3xx status code
func (o *CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyExperienceNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype delete device security policy anomaly experience no content response has a 4xx status code
func (o *CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyExperienceNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype delete device security policy anomaly experience no content response has a 5xx status code
func (o *CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyExperienceNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype delete device security policy anomaly experience no content response a status code equal to that given
func (o *CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyExperienceNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the customer prototype delete device security policy anomaly experience no content response
func (o *CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyExperienceNoContent) Code() int {
	return 204
}

func (o *CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyExperienceNoContent) Error() string {
	return fmt.Sprintf("[DELETE /Customers/{id}/locations/{locationId}/devices/{mac}/securityPolicy/anomaly/experience][%d] customerPrototypeDeleteDeviceSecurityPolicyAnomalyExperienceNoContent ", 204)
}

func (o *CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyExperienceNoContent) String() string {
	return fmt.Sprintf("[DELETE /Customers/{id}/locations/{locationId}/devices/{mac}/securityPolicy/anomaly/experience][%d] customerPrototypeDeleteDeviceSecurityPolicyAnomalyExperienceNoContent ", 204)
}

func (o *CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyExperienceNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCustomerPrototypeDeleteDeviceSecurityPolicyAnomalyExperienceDefault creates a CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyExperienceDefault with default headers values
func NewCustomerPrototypeDeleteDeviceSecurityPolicyAnomalyExperienceDefault(code int) *CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyExperienceDefault {
	return &CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyExperienceDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyExperienceDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyExperienceDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype delete device security policy anomaly experience default response has a 2xx status code
func (o *CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyExperienceDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype delete device security policy anomaly experience default response has a 3xx status code
func (o *CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyExperienceDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype delete device security policy anomaly experience default response has a 4xx status code
func (o *CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyExperienceDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype delete device security policy anomaly experience default response has a 5xx status code
func (o *CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyExperienceDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype delete device security policy anomaly experience default response a status code equal to that given
func (o *CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyExperienceDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype delete device security policy anomaly experience default response
func (o *CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyExperienceDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyExperienceDefault) Error() string {
	return fmt.Sprintf("[DELETE /Customers/{id}/locations/{locationId}/devices/{mac}/securityPolicy/anomaly/experience][%d] Customer.prototype.deleteDeviceSecurityPolicyAnomalyExperience default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyExperienceDefault) String() string {
	return fmt.Sprintf("[DELETE /Customers/{id}/locations/{locationId}/devices/{mac}/securityPolicy/anomaly/experience][%d] Customer.prototype.deleteDeviceSecurityPolicyAnomalyExperience default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyExperienceDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyExperienceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
