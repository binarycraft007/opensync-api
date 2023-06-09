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

// CustomerPrototypeGetNetworkAccessNetworksReader is a Reader for the CustomerPrototypeGetNetworkAccessNetworks structure.
type CustomerPrototypeGetNetworkAccessNetworksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypeGetNetworkAccessNetworksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypeGetNetworkAccessNetworksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypeGetNetworkAccessNetworksDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypeGetNetworkAccessNetworksOK creates a CustomerPrototypeGetNetworkAccessNetworksOK with default headers values
func NewCustomerPrototypeGetNetworkAccessNetworksOK() *CustomerPrototypeGetNetworkAccessNetworksOK {
	return &CustomerPrototypeGetNetworkAccessNetworksOK{}
}

/*
CustomerPrototypeGetNetworkAccessNetworksOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypeGetNetworkAccessNetworksOK struct {
	Payload []*models.NetworkAccessNetwork
}

// IsSuccess returns true when this customer prototype get network access networks o k response has a 2xx status code
func (o *CustomerPrototypeGetNetworkAccessNetworksOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype get network access networks o k response has a 3xx status code
func (o *CustomerPrototypeGetNetworkAccessNetworksOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype get network access networks o k response has a 4xx status code
func (o *CustomerPrototypeGetNetworkAccessNetworksOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype get network access networks o k response has a 5xx status code
func (o *CustomerPrototypeGetNetworkAccessNetworksOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype get network access networks o k response a status code equal to that given
func (o *CustomerPrototypeGetNetworkAccessNetworksOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype get network access networks o k response
func (o *CustomerPrototypeGetNetworkAccessNetworksOK) Code() int {
	return 200
}

func (o *CustomerPrototypeGetNetworkAccessNetworksOK) Error() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/networkAccess/networks][%d] customerPrototypeGetNetworkAccessNetworksOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeGetNetworkAccessNetworksOK) String() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/networkAccess/networks][%d] customerPrototypeGetNetworkAccessNetworksOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeGetNetworkAccessNetworksOK) GetPayload() []*models.NetworkAccessNetwork {
	return o.Payload
}

func (o *CustomerPrototypeGetNetworkAccessNetworksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypeGetNetworkAccessNetworksDefault creates a CustomerPrototypeGetNetworkAccessNetworksDefault with default headers values
func NewCustomerPrototypeGetNetworkAccessNetworksDefault(code int) *CustomerPrototypeGetNetworkAccessNetworksDefault {
	return &CustomerPrototypeGetNetworkAccessNetworksDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypeGetNetworkAccessNetworksDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypeGetNetworkAccessNetworksDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype get network access networks default response has a 2xx status code
func (o *CustomerPrototypeGetNetworkAccessNetworksDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype get network access networks default response has a 3xx status code
func (o *CustomerPrototypeGetNetworkAccessNetworksDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype get network access networks default response has a 4xx status code
func (o *CustomerPrototypeGetNetworkAccessNetworksDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype get network access networks default response has a 5xx status code
func (o *CustomerPrototypeGetNetworkAccessNetworksDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype get network access networks default response a status code equal to that given
func (o *CustomerPrototypeGetNetworkAccessNetworksDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype get network access networks default response
func (o *CustomerPrototypeGetNetworkAccessNetworksDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypeGetNetworkAccessNetworksDefault) Error() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/networkAccess/networks][%d] Customer.prototype.getNetworkAccessNetworks default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeGetNetworkAccessNetworksDefault) String() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/networkAccess/networks][%d] Customer.prototype.getNetworkAccessNetworks default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeGetNetworkAccessNetworksDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypeGetNetworkAccessNetworksDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
