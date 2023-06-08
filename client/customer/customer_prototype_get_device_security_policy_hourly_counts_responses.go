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

// CustomerPrototypeGetDeviceSecurityPolicyHourlyCountsReader is a Reader for the CustomerPrototypeGetDeviceSecurityPolicyHourlyCounts structure.
type CustomerPrototypeGetDeviceSecurityPolicyHourlyCountsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypeGetDeviceSecurityPolicyHourlyCountsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypeGetDeviceSecurityPolicyHourlyCountsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypeGetDeviceSecurityPolicyHourlyCountsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypeGetDeviceSecurityPolicyHourlyCountsOK creates a CustomerPrototypeGetDeviceSecurityPolicyHourlyCountsOK with default headers values
func NewCustomerPrototypeGetDeviceSecurityPolicyHourlyCountsOK() *CustomerPrototypeGetDeviceSecurityPolicyHourlyCountsOK {
	return &CustomerPrototypeGetDeviceSecurityPolicyHourlyCountsOK{}
}

/*
CustomerPrototypeGetDeviceSecurityPolicyHourlyCountsOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypeGetDeviceSecurityPolicyHourlyCountsOK struct {
	Payload *models.Counts
}

// IsSuccess returns true when this customer prototype get device security policy hourly counts o k response has a 2xx status code
func (o *CustomerPrototypeGetDeviceSecurityPolicyHourlyCountsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype get device security policy hourly counts o k response has a 3xx status code
func (o *CustomerPrototypeGetDeviceSecurityPolicyHourlyCountsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype get device security policy hourly counts o k response has a 4xx status code
func (o *CustomerPrototypeGetDeviceSecurityPolicyHourlyCountsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype get device security policy hourly counts o k response has a 5xx status code
func (o *CustomerPrototypeGetDeviceSecurityPolicyHourlyCountsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype get device security policy hourly counts o k response a status code equal to that given
func (o *CustomerPrototypeGetDeviceSecurityPolicyHourlyCountsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype get device security policy hourly counts o k response
func (o *CustomerPrototypeGetDeviceSecurityPolicyHourlyCountsOK) Code() int {
	return 200
}

func (o *CustomerPrototypeGetDeviceSecurityPolicyHourlyCountsOK) Error() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/devices/{mac}/securityPolicy/hourlyBlockedCounts][%d] customerPrototypeGetDeviceSecurityPolicyHourlyCountsOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeGetDeviceSecurityPolicyHourlyCountsOK) String() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/devices/{mac}/securityPolicy/hourlyBlockedCounts][%d] customerPrototypeGetDeviceSecurityPolicyHourlyCountsOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeGetDeviceSecurityPolicyHourlyCountsOK) GetPayload() *models.Counts {
	return o.Payload
}

func (o *CustomerPrototypeGetDeviceSecurityPolicyHourlyCountsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Counts)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypeGetDeviceSecurityPolicyHourlyCountsDefault creates a CustomerPrototypeGetDeviceSecurityPolicyHourlyCountsDefault with default headers values
func NewCustomerPrototypeGetDeviceSecurityPolicyHourlyCountsDefault(code int) *CustomerPrototypeGetDeviceSecurityPolicyHourlyCountsDefault {
	return &CustomerPrototypeGetDeviceSecurityPolicyHourlyCountsDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypeGetDeviceSecurityPolicyHourlyCountsDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypeGetDeviceSecurityPolicyHourlyCountsDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype get device security policy hourly counts default response has a 2xx status code
func (o *CustomerPrototypeGetDeviceSecurityPolicyHourlyCountsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype get device security policy hourly counts default response has a 3xx status code
func (o *CustomerPrototypeGetDeviceSecurityPolicyHourlyCountsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype get device security policy hourly counts default response has a 4xx status code
func (o *CustomerPrototypeGetDeviceSecurityPolicyHourlyCountsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype get device security policy hourly counts default response has a 5xx status code
func (o *CustomerPrototypeGetDeviceSecurityPolicyHourlyCountsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype get device security policy hourly counts default response a status code equal to that given
func (o *CustomerPrototypeGetDeviceSecurityPolicyHourlyCountsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype get device security policy hourly counts default response
func (o *CustomerPrototypeGetDeviceSecurityPolicyHourlyCountsDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypeGetDeviceSecurityPolicyHourlyCountsDefault) Error() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/devices/{mac}/securityPolicy/hourlyBlockedCounts][%d] Customer.prototype.getDeviceSecurityPolicyHourlyCounts default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeGetDeviceSecurityPolicyHourlyCountsDefault) String() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/devices/{mac}/securityPolicy/hourlyBlockedCounts][%d] Customer.prototype.getDeviceSecurityPolicyHourlyCounts default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeGetDeviceSecurityPolicyHourlyCountsDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypeGetDeviceSecurityPolicyHourlyCountsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}