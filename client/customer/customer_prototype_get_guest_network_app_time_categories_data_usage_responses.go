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

// CustomerPrototypeGetGuestNetworkAppTimeCategoriesDataUsageReader is a Reader for the CustomerPrototypeGetGuestNetworkAppTimeCategoriesDataUsage structure.
type CustomerPrototypeGetGuestNetworkAppTimeCategoriesDataUsageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypeGetGuestNetworkAppTimeCategoriesDataUsageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypeGetGuestNetworkAppTimeCategoriesDataUsageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypeGetGuestNetworkAppTimeCategoriesDataUsageDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypeGetGuestNetworkAppTimeCategoriesDataUsageOK creates a CustomerPrototypeGetGuestNetworkAppTimeCategoriesDataUsageOK with default headers values
func NewCustomerPrototypeGetGuestNetworkAppTimeCategoriesDataUsageOK() *CustomerPrototypeGetGuestNetworkAppTimeCategoriesDataUsageOK {
	return &CustomerPrototypeGetGuestNetworkAppTimeCategoriesDataUsageOK{}
}

/*
CustomerPrototypeGetGuestNetworkAppTimeCategoriesDataUsageOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypeGetGuestNetworkAppTimeCategoriesDataUsageOK struct {
	Payload []interface{}
}

// IsSuccess returns true when this customer prototype get guest network app time categories data usage o k response has a 2xx status code
func (o *CustomerPrototypeGetGuestNetworkAppTimeCategoriesDataUsageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype get guest network app time categories data usage o k response has a 3xx status code
func (o *CustomerPrototypeGetGuestNetworkAppTimeCategoriesDataUsageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype get guest network app time categories data usage o k response has a 4xx status code
func (o *CustomerPrototypeGetGuestNetworkAppTimeCategoriesDataUsageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype get guest network app time categories data usage o k response has a 5xx status code
func (o *CustomerPrototypeGetGuestNetworkAppTimeCategoriesDataUsageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype get guest network app time categories data usage o k response a status code equal to that given
func (o *CustomerPrototypeGetGuestNetworkAppTimeCategoriesDataUsageOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype get guest network app time categories data usage o k response
func (o *CustomerPrototypeGetGuestNetworkAppTimeCategoriesDataUsageOK) Code() int {
	return 200
}

func (o *CustomerPrototypeGetGuestNetworkAppTimeCategoriesDataUsageOK) Error() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/secondaryNetworks/captivePortals/{networkId}/appTime/categories/dataUsage][%d] customerPrototypeGetGuestNetworkAppTimeCategoriesDataUsageOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeGetGuestNetworkAppTimeCategoriesDataUsageOK) String() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/secondaryNetworks/captivePortals/{networkId}/appTime/categories/dataUsage][%d] customerPrototypeGetGuestNetworkAppTimeCategoriesDataUsageOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeGetGuestNetworkAppTimeCategoriesDataUsageOK) GetPayload() []interface{} {
	return o.Payload
}

func (o *CustomerPrototypeGetGuestNetworkAppTimeCategoriesDataUsageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypeGetGuestNetworkAppTimeCategoriesDataUsageDefault creates a CustomerPrototypeGetGuestNetworkAppTimeCategoriesDataUsageDefault with default headers values
func NewCustomerPrototypeGetGuestNetworkAppTimeCategoriesDataUsageDefault(code int) *CustomerPrototypeGetGuestNetworkAppTimeCategoriesDataUsageDefault {
	return &CustomerPrototypeGetGuestNetworkAppTimeCategoriesDataUsageDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypeGetGuestNetworkAppTimeCategoriesDataUsageDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypeGetGuestNetworkAppTimeCategoriesDataUsageDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype get guest network app time categories data usage default response has a 2xx status code
func (o *CustomerPrototypeGetGuestNetworkAppTimeCategoriesDataUsageDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype get guest network app time categories data usage default response has a 3xx status code
func (o *CustomerPrototypeGetGuestNetworkAppTimeCategoriesDataUsageDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype get guest network app time categories data usage default response has a 4xx status code
func (o *CustomerPrototypeGetGuestNetworkAppTimeCategoriesDataUsageDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype get guest network app time categories data usage default response has a 5xx status code
func (o *CustomerPrototypeGetGuestNetworkAppTimeCategoriesDataUsageDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype get guest network app time categories data usage default response a status code equal to that given
func (o *CustomerPrototypeGetGuestNetworkAppTimeCategoriesDataUsageDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype get guest network app time categories data usage default response
func (o *CustomerPrototypeGetGuestNetworkAppTimeCategoriesDataUsageDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypeGetGuestNetworkAppTimeCategoriesDataUsageDefault) Error() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/secondaryNetworks/captivePortals/{networkId}/appTime/categories/dataUsage][%d] Customer.prototype.getGuestNetworkAppTimeCategoriesDataUsage default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeGetGuestNetworkAppTimeCategoriesDataUsageDefault) String() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/secondaryNetworks/captivePortals/{networkId}/appTime/categories/dataUsage][%d] Customer.prototype.getGuestNetworkAppTimeCategoriesDataUsage default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeGetGuestNetworkAppTimeCategoriesDataUsageDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypeGetGuestNetworkAppTimeCategoriesDataUsageDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}