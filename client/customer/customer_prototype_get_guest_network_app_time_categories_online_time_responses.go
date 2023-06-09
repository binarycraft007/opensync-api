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

// CustomerPrototypeGetGuestNetworkAppTimeCategoriesOnlineTimeReader is a Reader for the CustomerPrototypeGetGuestNetworkAppTimeCategoriesOnlineTime structure.
type CustomerPrototypeGetGuestNetworkAppTimeCategoriesOnlineTimeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypeGetGuestNetworkAppTimeCategoriesOnlineTimeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypeGetGuestNetworkAppTimeCategoriesOnlineTimeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypeGetGuestNetworkAppTimeCategoriesOnlineTimeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypeGetGuestNetworkAppTimeCategoriesOnlineTimeOK creates a CustomerPrototypeGetGuestNetworkAppTimeCategoriesOnlineTimeOK with default headers values
func NewCustomerPrototypeGetGuestNetworkAppTimeCategoriesOnlineTimeOK() *CustomerPrototypeGetGuestNetworkAppTimeCategoriesOnlineTimeOK {
	return &CustomerPrototypeGetGuestNetworkAppTimeCategoriesOnlineTimeOK{}
}

/*
CustomerPrototypeGetGuestNetworkAppTimeCategoriesOnlineTimeOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypeGetGuestNetworkAppTimeCategoriesOnlineTimeOK struct {
	Payload []interface{}
}

// IsSuccess returns true when this customer prototype get guest network app time categories online time o k response has a 2xx status code
func (o *CustomerPrototypeGetGuestNetworkAppTimeCategoriesOnlineTimeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype get guest network app time categories online time o k response has a 3xx status code
func (o *CustomerPrototypeGetGuestNetworkAppTimeCategoriesOnlineTimeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype get guest network app time categories online time o k response has a 4xx status code
func (o *CustomerPrototypeGetGuestNetworkAppTimeCategoriesOnlineTimeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype get guest network app time categories online time o k response has a 5xx status code
func (o *CustomerPrototypeGetGuestNetworkAppTimeCategoriesOnlineTimeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype get guest network app time categories online time o k response a status code equal to that given
func (o *CustomerPrototypeGetGuestNetworkAppTimeCategoriesOnlineTimeOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype get guest network app time categories online time o k response
func (o *CustomerPrototypeGetGuestNetworkAppTimeCategoriesOnlineTimeOK) Code() int {
	return 200
}

func (o *CustomerPrototypeGetGuestNetworkAppTimeCategoriesOnlineTimeOK) Error() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/secondaryNetworks/captivePortals/{networkId}/appTime/categories/onlineTime][%d] customerPrototypeGetGuestNetworkAppTimeCategoriesOnlineTimeOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeGetGuestNetworkAppTimeCategoriesOnlineTimeOK) String() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/secondaryNetworks/captivePortals/{networkId}/appTime/categories/onlineTime][%d] customerPrototypeGetGuestNetworkAppTimeCategoriesOnlineTimeOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeGetGuestNetworkAppTimeCategoriesOnlineTimeOK) GetPayload() []interface{} {
	return o.Payload
}

func (o *CustomerPrototypeGetGuestNetworkAppTimeCategoriesOnlineTimeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypeGetGuestNetworkAppTimeCategoriesOnlineTimeDefault creates a CustomerPrototypeGetGuestNetworkAppTimeCategoriesOnlineTimeDefault with default headers values
func NewCustomerPrototypeGetGuestNetworkAppTimeCategoriesOnlineTimeDefault(code int) *CustomerPrototypeGetGuestNetworkAppTimeCategoriesOnlineTimeDefault {
	return &CustomerPrototypeGetGuestNetworkAppTimeCategoriesOnlineTimeDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypeGetGuestNetworkAppTimeCategoriesOnlineTimeDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypeGetGuestNetworkAppTimeCategoriesOnlineTimeDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype get guest network app time categories online time default response has a 2xx status code
func (o *CustomerPrototypeGetGuestNetworkAppTimeCategoriesOnlineTimeDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype get guest network app time categories online time default response has a 3xx status code
func (o *CustomerPrototypeGetGuestNetworkAppTimeCategoriesOnlineTimeDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype get guest network app time categories online time default response has a 4xx status code
func (o *CustomerPrototypeGetGuestNetworkAppTimeCategoriesOnlineTimeDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype get guest network app time categories online time default response has a 5xx status code
func (o *CustomerPrototypeGetGuestNetworkAppTimeCategoriesOnlineTimeDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype get guest network app time categories online time default response a status code equal to that given
func (o *CustomerPrototypeGetGuestNetworkAppTimeCategoriesOnlineTimeDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype get guest network app time categories online time default response
func (o *CustomerPrototypeGetGuestNetworkAppTimeCategoriesOnlineTimeDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypeGetGuestNetworkAppTimeCategoriesOnlineTimeDefault) Error() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/secondaryNetworks/captivePortals/{networkId}/appTime/categories/onlineTime][%d] Customer.prototype.getGuestNetworkAppTimeCategoriesOnlineTime default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeGetGuestNetworkAppTimeCategoriesOnlineTimeDefault) String() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/secondaryNetworks/captivePortals/{networkId}/appTime/categories/onlineTime][%d] Customer.prototype.getGuestNetworkAppTimeCategoriesOnlineTime default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeGetGuestNetworkAppTimeCategoriesOnlineTimeDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypeGetGuestNetworkAppTimeCategoriesOnlineTimeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
