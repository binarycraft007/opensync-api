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

// CustomerPrototypeCreateIPLimitedAccessTokenReader is a Reader for the CustomerPrototypeCreateIPLimitedAccessToken structure.
type CustomerPrototypeCreateIPLimitedAccessTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypeCreateIPLimitedAccessTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypeCreateIPLimitedAccessTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypeCreateIPLimitedAccessTokenDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypeCreateIPLimitedAccessTokenOK creates a CustomerPrototypeCreateIPLimitedAccessTokenOK with default headers values
func NewCustomerPrototypeCreateIPLimitedAccessTokenOK() *CustomerPrototypeCreateIPLimitedAccessTokenOK {
	return &CustomerPrototypeCreateIPLimitedAccessTokenOK{}
}

/*
CustomerPrototypeCreateIPLimitedAccessTokenOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypeCreateIPLimitedAccessTokenOK struct {
	Payload *models.AccessToken
}

// IsSuccess returns true when this customer prototype create Ip limited access token o k response has a 2xx status code
func (o *CustomerPrototypeCreateIPLimitedAccessTokenOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype create Ip limited access token o k response has a 3xx status code
func (o *CustomerPrototypeCreateIPLimitedAccessTokenOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype create Ip limited access token o k response has a 4xx status code
func (o *CustomerPrototypeCreateIPLimitedAccessTokenOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype create Ip limited access token o k response has a 5xx status code
func (o *CustomerPrototypeCreateIPLimitedAccessTokenOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype create Ip limited access token o k response a status code equal to that given
func (o *CustomerPrototypeCreateIPLimitedAccessTokenOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype create Ip limited access token o k response
func (o *CustomerPrototypeCreateIPLimitedAccessTokenOK) Code() int {
	return 200
}

func (o *CustomerPrototypeCreateIPLimitedAccessTokenOK) Error() string {
	return fmt.Sprintf("[POST /Customers/{id}/createIpLimitedAccessToken][%d] customerPrototypeCreateIpLimitedAccessTokenOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeCreateIPLimitedAccessTokenOK) String() string {
	return fmt.Sprintf("[POST /Customers/{id}/createIpLimitedAccessToken][%d] customerPrototypeCreateIpLimitedAccessTokenOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeCreateIPLimitedAccessTokenOK) GetPayload() *models.AccessToken {
	return o.Payload
}

func (o *CustomerPrototypeCreateIPLimitedAccessTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccessToken)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypeCreateIPLimitedAccessTokenDefault creates a CustomerPrototypeCreateIPLimitedAccessTokenDefault with default headers values
func NewCustomerPrototypeCreateIPLimitedAccessTokenDefault(code int) *CustomerPrototypeCreateIPLimitedAccessTokenDefault {
	return &CustomerPrototypeCreateIPLimitedAccessTokenDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypeCreateIPLimitedAccessTokenDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypeCreateIPLimitedAccessTokenDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype create Ip limited access token default response has a 2xx status code
func (o *CustomerPrototypeCreateIPLimitedAccessTokenDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype create Ip limited access token default response has a 3xx status code
func (o *CustomerPrototypeCreateIPLimitedAccessTokenDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype create Ip limited access token default response has a 4xx status code
func (o *CustomerPrototypeCreateIPLimitedAccessTokenDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype create Ip limited access token default response has a 5xx status code
func (o *CustomerPrototypeCreateIPLimitedAccessTokenDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype create Ip limited access token default response a status code equal to that given
func (o *CustomerPrototypeCreateIPLimitedAccessTokenDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype create Ip limited access token default response
func (o *CustomerPrototypeCreateIPLimitedAccessTokenDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypeCreateIPLimitedAccessTokenDefault) Error() string {
	return fmt.Sprintf("[POST /Customers/{id}/createIpLimitedAccessToken][%d] Customer.prototype.createIpLimitedAccessToken default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeCreateIPLimitedAccessTokenDefault) String() string {
	return fmt.Sprintf("[POST /Customers/{id}/createIpLimitedAccessToken][%d] Customer.prototype.createIpLimitedAccessToken default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeCreateIPLimitedAccessTokenDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypeCreateIPLimitedAccessTokenDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
