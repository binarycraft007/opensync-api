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

// CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesDataUsageReader is a Reader for the CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesDataUsage structure.
type CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesDataUsageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesDataUsageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypeGetEmployeeNetworkAppTimeCategoriesDataUsageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypeGetEmployeeNetworkAppTimeCategoriesDataUsageDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypeGetEmployeeNetworkAppTimeCategoriesDataUsageOK creates a CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesDataUsageOK with default headers values
func NewCustomerPrototypeGetEmployeeNetworkAppTimeCategoriesDataUsageOK() *CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesDataUsageOK {
	return &CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesDataUsageOK{}
}

/*
CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesDataUsageOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesDataUsageOK struct {
	Payload []interface{}
}

// IsSuccess returns true when this customer prototype get employee network app time categories data usage o k response has a 2xx status code
func (o *CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesDataUsageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype get employee network app time categories data usage o k response has a 3xx status code
func (o *CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesDataUsageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype get employee network app time categories data usage o k response has a 4xx status code
func (o *CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesDataUsageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype get employee network app time categories data usage o k response has a 5xx status code
func (o *CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesDataUsageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype get employee network app time categories data usage o k response a status code equal to that given
func (o *CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesDataUsageOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype get employee network app time categories data usage o k response
func (o *CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesDataUsageOK) Code() int {
	return 200
}

func (o *CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesDataUsageOK) Error() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/secondaryNetworks/fronthauls/{networkId}/appTime/categories/dataUsage][%d] customerPrototypeGetEmployeeNetworkAppTimeCategoriesDataUsageOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesDataUsageOK) String() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/secondaryNetworks/fronthauls/{networkId}/appTime/categories/dataUsage][%d] customerPrototypeGetEmployeeNetworkAppTimeCategoriesDataUsageOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesDataUsageOK) GetPayload() []interface{} {
	return o.Payload
}

func (o *CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesDataUsageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypeGetEmployeeNetworkAppTimeCategoriesDataUsageDefault creates a CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesDataUsageDefault with default headers values
func NewCustomerPrototypeGetEmployeeNetworkAppTimeCategoriesDataUsageDefault(code int) *CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesDataUsageDefault {
	return &CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesDataUsageDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesDataUsageDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesDataUsageDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype get employee network app time categories data usage default response has a 2xx status code
func (o *CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesDataUsageDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype get employee network app time categories data usage default response has a 3xx status code
func (o *CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesDataUsageDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype get employee network app time categories data usage default response has a 4xx status code
func (o *CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesDataUsageDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype get employee network app time categories data usage default response has a 5xx status code
func (o *CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesDataUsageDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype get employee network app time categories data usage default response a status code equal to that given
func (o *CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesDataUsageDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype get employee network app time categories data usage default response
func (o *CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesDataUsageDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesDataUsageDefault) Error() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/secondaryNetworks/fronthauls/{networkId}/appTime/categories/dataUsage][%d] Customer.prototype.getEmployeeNetworkAppTimeCategoriesDataUsage default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesDataUsageDefault) String() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/secondaryNetworks/fronthauls/{networkId}/appTime/categories/dataUsage][%d] Customer.prototype.getEmployeeNetworkAppTimeCategoriesDataUsage default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesDataUsageDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesDataUsageDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
