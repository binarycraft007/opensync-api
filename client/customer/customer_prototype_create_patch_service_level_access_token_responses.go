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

// CustomerPrototypeCreatePatchServiceLevelAccessTokenReader is a Reader for the CustomerPrototypeCreatePatchServiceLevelAccessToken structure.
type CustomerPrototypeCreatePatchServiceLevelAccessTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypeCreatePatchServiceLevelAccessTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypeCreatePatchServiceLevelAccessTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypeCreatePatchServiceLevelAccessTokenDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypeCreatePatchServiceLevelAccessTokenOK creates a CustomerPrototypeCreatePatchServiceLevelAccessTokenOK with default headers values
func NewCustomerPrototypeCreatePatchServiceLevelAccessTokenOK() *CustomerPrototypeCreatePatchServiceLevelAccessTokenOK {
	return &CustomerPrototypeCreatePatchServiceLevelAccessTokenOK{}
}

/*
CustomerPrototypeCreatePatchServiceLevelAccessTokenOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypeCreatePatchServiceLevelAccessTokenOK struct {
	Payload *models.AccessToken
}

// IsSuccess returns true when this customer prototype create patch service level access token o k response has a 2xx status code
func (o *CustomerPrototypeCreatePatchServiceLevelAccessTokenOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype create patch service level access token o k response has a 3xx status code
func (o *CustomerPrototypeCreatePatchServiceLevelAccessTokenOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype create patch service level access token o k response has a 4xx status code
func (o *CustomerPrototypeCreatePatchServiceLevelAccessTokenOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype create patch service level access token o k response has a 5xx status code
func (o *CustomerPrototypeCreatePatchServiceLevelAccessTokenOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype create patch service level access token o k response a status code equal to that given
func (o *CustomerPrototypeCreatePatchServiceLevelAccessTokenOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype create patch service level access token o k response
func (o *CustomerPrototypeCreatePatchServiceLevelAccessTokenOK) Code() int {
	return 200
}

func (o *CustomerPrototypeCreatePatchServiceLevelAccessTokenOK) Error() string {
	return fmt.Sprintf("[POST /Customers/{id}/createPatchServiceLevelAccessToken][%d] customerPrototypeCreatePatchServiceLevelAccessTokenOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeCreatePatchServiceLevelAccessTokenOK) String() string {
	return fmt.Sprintf("[POST /Customers/{id}/createPatchServiceLevelAccessToken][%d] customerPrototypeCreatePatchServiceLevelAccessTokenOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeCreatePatchServiceLevelAccessTokenOK) GetPayload() *models.AccessToken {
	return o.Payload
}

func (o *CustomerPrototypeCreatePatchServiceLevelAccessTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccessToken)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypeCreatePatchServiceLevelAccessTokenDefault creates a CustomerPrototypeCreatePatchServiceLevelAccessTokenDefault with default headers values
func NewCustomerPrototypeCreatePatchServiceLevelAccessTokenDefault(code int) *CustomerPrototypeCreatePatchServiceLevelAccessTokenDefault {
	return &CustomerPrototypeCreatePatchServiceLevelAccessTokenDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypeCreatePatchServiceLevelAccessTokenDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypeCreatePatchServiceLevelAccessTokenDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype create patch service level access token default response has a 2xx status code
func (o *CustomerPrototypeCreatePatchServiceLevelAccessTokenDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype create patch service level access token default response has a 3xx status code
func (o *CustomerPrototypeCreatePatchServiceLevelAccessTokenDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype create patch service level access token default response has a 4xx status code
func (o *CustomerPrototypeCreatePatchServiceLevelAccessTokenDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype create patch service level access token default response has a 5xx status code
func (o *CustomerPrototypeCreatePatchServiceLevelAccessTokenDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype create patch service level access token default response a status code equal to that given
func (o *CustomerPrototypeCreatePatchServiceLevelAccessTokenDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype create patch service level access token default response
func (o *CustomerPrototypeCreatePatchServiceLevelAccessTokenDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypeCreatePatchServiceLevelAccessTokenDefault) Error() string {
	return fmt.Sprintf("[POST /Customers/{id}/createPatchServiceLevelAccessToken][%d] Customer.prototype.createPatchServiceLevelAccessToken default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeCreatePatchServiceLevelAccessTokenDefault) String() string {
	return fmt.Sprintf("[POST /Customers/{id}/createPatchServiceLevelAccessToken][%d] Customer.prototype.createPatchServiceLevelAccessToken default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeCreatePatchServiceLevelAccessTokenDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypeCreatePatchServiceLevelAccessTokenDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
