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

// CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesOnlineTimeReader is a Reader for the CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesOnlineTime structure.
type CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesOnlineTimeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesOnlineTimeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypeGetEmployeeNetworkAppTimeCategoriesOnlineTimeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypeGetEmployeeNetworkAppTimeCategoriesOnlineTimeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypeGetEmployeeNetworkAppTimeCategoriesOnlineTimeOK creates a CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesOnlineTimeOK with default headers values
func NewCustomerPrototypeGetEmployeeNetworkAppTimeCategoriesOnlineTimeOK() *CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesOnlineTimeOK {
	return &CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesOnlineTimeOK{}
}

/*
CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesOnlineTimeOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesOnlineTimeOK struct {
	Payload []interface{}
}

// IsSuccess returns true when this customer prototype get employee network app time categories online time o k response has a 2xx status code
func (o *CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesOnlineTimeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype get employee network app time categories online time o k response has a 3xx status code
func (o *CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesOnlineTimeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype get employee network app time categories online time o k response has a 4xx status code
func (o *CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesOnlineTimeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype get employee network app time categories online time o k response has a 5xx status code
func (o *CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesOnlineTimeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype get employee network app time categories online time o k response a status code equal to that given
func (o *CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesOnlineTimeOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype get employee network app time categories online time o k response
func (o *CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesOnlineTimeOK) Code() int {
	return 200
}

func (o *CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesOnlineTimeOK) Error() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/secondaryNetworks/fronthauls/{networkId}/appTime/categories/onlineTime][%d] customerPrototypeGetEmployeeNetworkAppTimeCategoriesOnlineTimeOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesOnlineTimeOK) String() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/secondaryNetworks/fronthauls/{networkId}/appTime/categories/onlineTime][%d] customerPrototypeGetEmployeeNetworkAppTimeCategoriesOnlineTimeOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesOnlineTimeOK) GetPayload() []interface{} {
	return o.Payload
}

func (o *CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesOnlineTimeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypeGetEmployeeNetworkAppTimeCategoriesOnlineTimeDefault creates a CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesOnlineTimeDefault with default headers values
func NewCustomerPrototypeGetEmployeeNetworkAppTimeCategoriesOnlineTimeDefault(code int) *CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesOnlineTimeDefault {
	return &CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesOnlineTimeDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesOnlineTimeDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesOnlineTimeDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype get employee network app time categories online time default response has a 2xx status code
func (o *CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesOnlineTimeDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype get employee network app time categories online time default response has a 3xx status code
func (o *CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesOnlineTimeDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype get employee network app time categories online time default response has a 4xx status code
func (o *CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesOnlineTimeDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype get employee network app time categories online time default response has a 5xx status code
func (o *CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesOnlineTimeDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype get employee network app time categories online time default response a status code equal to that given
func (o *CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesOnlineTimeDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype get employee network app time categories online time default response
func (o *CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesOnlineTimeDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesOnlineTimeDefault) Error() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/secondaryNetworks/fronthauls/{networkId}/appTime/categories/onlineTime][%d] Customer.prototype.getEmployeeNetworkAppTimeCategoriesOnlineTime default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesOnlineTimeDefault) String() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/secondaryNetworks/fronthauls/{networkId}/appTime/categories/onlineTime][%d] Customer.prototype.getEmployeeNetworkAppTimeCategoriesOnlineTime default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesOnlineTimeDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesOnlineTimeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
