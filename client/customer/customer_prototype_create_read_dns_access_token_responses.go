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

// CustomerPrototypeCreateReadDNSAccessTokenReader is a Reader for the CustomerPrototypeCreateReadDNSAccessToken structure.
type CustomerPrototypeCreateReadDNSAccessTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypeCreateReadDNSAccessTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypeCreateReadDNSAccessTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypeCreateReadDNSAccessTokenDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypeCreateReadDNSAccessTokenOK creates a CustomerPrototypeCreateReadDNSAccessTokenOK with default headers values
func NewCustomerPrototypeCreateReadDNSAccessTokenOK() *CustomerPrototypeCreateReadDNSAccessTokenOK {
	return &CustomerPrototypeCreateReadDNSAccessTokenOK{}
}

/*
CustomerPrototypeCreateReadDNSAccessTokenOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypeCreateReadDNSAccessTokenOK struct {
	Payload *models.AccessToken
}

// IsSuccess returns true when this customer prototype create read Dns access token o k response has a 2xx status code
func (o *CustomerPrototypeCreateReadDNSAccessTokenOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype create read Dns access token o k response has a 3xx status code
func (o *CustomerPrototypeCreateReadDNSAccessTokenOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype create read Dns access token o k response has a 4xx status code
func (o *CustomerPrototypeCreateReadDNSAccessTokenOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype create read Dns access token o k response has a 5xx status code
func (o *CustomerPrototypeCreateReadDNSAccessTokenOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype create read Dns access token o k response a status code equal to that given
func (o *CustomerPrototypeCreateReadDNSAccessTokenOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype create read Dns access token o k response
func (o *CustomerPrototypeCreateReadDNSAccessTokenOK) Code() int {
	return 200
}

func (o *CustomerPrototypeCreateReadDNSAccessTokenOK) Error() string {
	return fmt.Sprintf("[POST /Customers/{id}/createReadDnsAccessToken][%d] customerPrototypeCreateReadDnsAccessTokenOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeCreateReadDNSAccessTokenOK) String() string {
	return fmt.Sprintf("[POST /Customers/{id}/createReadDnsAccessToken][%d] customerPrototypeCreateReadDnsAccessTokenOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeCreateReadDNSAccessTokenOK) GetPayload() *models.AccessToken {
	return o.Payload
}

func (o *CustomerPrototypeCreateReadDNSAccessTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccessToken)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypeCreateReadDNSAccessTokenDefault creates a CustomerPrototypeCreateReadDNSAccessTokenDefault with default headers values
func NewCustomerPrototypeCreateReadDNSAccessTokenDefault(code int) *CustomerPrototypeCreateReadDNSAccessTokenDefault {
	return &CustomerPrototypeCreateReadDNSAccessTokenDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypeCreateReadDNSAccessTokenDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypeCreateReadDNSAccessTokenDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype create read Dns access token default response has a 2xx status code
func (o *CustomerPrototypeCreateReadDNSAccessTokenDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype create read Dns access token default response has a 3xx status code
func (o *CustomerPrototypeCreateReadDNSAccessTokenDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype create read Dns access token default response has a 4xx status code
func (o *CustomerPrototypeCreateReadDNSAccessTokenDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype create read Dns access token default response has a 5xx status code
func (o *CustomerPrototypeCreateReadDNSAccessTokenDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype create read Dns access token default response a status code equal to that given
func (o *CustomerPrototypeCreateReadDNSAccessTokenDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype create read Dns access token default response
func (o *CustomerPrototypeCreateReadDNSAccessTokenDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypeCreateReadDNSAccessTokenDefault) Error() string {
	return fmt.Sprintf("[POST /Customers/{id}/createReadDnsAccessToken][%d] Customer.prototype.createReadDnsAccessToken default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeCreateReadDNSAccessTokenDefault) String() string {
	return fmt.Sprintf("[POST /Customers/{id}/createReadDnsAccessToken][%d] Customer.prototype.createReadDnsAccessToken default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeCreateReadDNSAccessTokenDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypeCreateReadDNSAccessTokenDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
