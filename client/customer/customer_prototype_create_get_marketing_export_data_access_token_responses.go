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

// CustomerPrototypeCreateGetMarketingExportDataAccessTokenReader is a Reader for the CustomerPrototypeCreateGetMarketingExportDataAccessToken structure.
type CustomerPrototypeCreateGetMarketingExportDataAccessTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypeCreateGetMarketingExportDataAccessTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypeCreateGetMarketingExportDataAccessTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypeCreateGetMarketingExportDataAccessTokenDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypeCreateGetMarketingExportDataAccessTokenOK creates a CustomerPrototypeCreateGetMarketingExportDataAccessTokenOK with default headers values
func NewCustomerPrototypeCreateGetMarketingExportDataAccessTokenOK() *CustomerPrototypeCreateGetMarketingExportDataAccessTokenOK {
	return &CustomerPrototypeCreateGetMarketingExportDataAccessTokenOK{}
}

/*
CustomerPrototypeCreateGetMarketingExportDataAccessTokenOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypeCreateGetMarketingExportDataAccessTokenOK struct {
	Payload *models.AccessToken
}

// IsSuccess returns true when this customer prototype create get marketing export data access token o k response has a 2xx status code
func (o *CustomerPrototypeCreateGetMarketingExportDataAccessTokenOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype create get marketing export data access token o k response has a 3xx status code
func (o *CustomerPrototypeCreateGetMarketingExportDataAccessTokenOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype create get marketing export data access token o k response has a 4xx status code
func (o *CustomerPrototypeCreateGetMarketingExportDataAccessTokenOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype create get marketing export data access token o k response has a 5xx status code
func (o *CustomerPrototypeCreateGetMarketingExportDataAccessTokenOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype create get marketing export data access token o k response a status code equal to that given
func (o *CustomerPrototypeCreateGetMarketingExportDataAccessTokenOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype create get marketing export data access token o k response
func (o *CustomerPrototypeCreateGetMarketingExportDataAccessTokenOK) Code() int {
	return 200
}

func (o *CustomerPrototypeCreateGetMarketingExportDataAccessTokenOK) Error() string {
	return fmt.Sprintf("[POST /Customers/{id}/createGetMarketingExportDataAccessToken][%d] customerPrototypeCreateGetMarketingExportDataAccessTokenOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeCreateGetMarketingExportDataAccessTokenOK) String() string {
	return fmt.Sprintf("[POST /Customers/{id}/createGetMarketingExportDataAccessToken][%d] customerPrototypeCreateGetMarketingExportDataAccessTokenOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeCreateGetMarketingExportDataAccessTokenOK) GetPayload() *models.AccessToken {
	return o.Payload
}

func (o *CustomerPrototypeCreateGetMarketingExportDataAccessTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccessToken)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypeCreateGetMarketingExportDataAccessTokenDefault creates a CustomerPrototypeCreateGetMarketingExportDataAccessTokenDefault with default headers values
func NewCustomerPrototypeCreateGetMarketingExportDataAccessTokenDefault(code int) *CustomerPrototypeCreateGetMarketingExportDataAccessTokenDefault {
	return &CustomerPrototypeCreateGetMarketingExportDataAccessTokenDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypeCreateGetMarketingExportDataAccessTokenDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypeCreateGetMarketingExportDataAccessTokenDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype create get marketing export data access token default response has a 2xx status code
func (o *CustomerPrototypeCreateGetMarketingExportDataAccessTokenDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype create get marketing export data access token default response has a 3xx status code
func (o *CustomerPrototypeCreateGetMarketingExportDataAccessTokenDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype create get marketing export data access token default response has a 4xx status code
func (o *CustomerPrototypeCreateGetMarketingExportDataAccessTokenDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype create get marketing export data access token default response has a 5xx status code
func (o *CustomerPrototypeCreateGetMarketingExportDataAccessTokenDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype create get marketing export data access token default response a status code equal to that given
func (o *CustomerPrototypeCreateGetMarketingExportDataAccessTokenDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype create get marketing export data access token default response
func (o *CustomerPrototypeCreateGetMarketingExportDataAccessTokenDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypeCreateGetMarketingExportDataAccessTokenDefault) Error() string {
	return fmt.Sprintf("[POST /Customers/{id}/createGetMarketingExportDataAccessToken][%d] Customer.prototype.createGetMarketingExportDataAccessToken default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeCreateGetMarketingExportDataAccessTokenDefault) String() string {
	return fmt.Sprintf("[POST /Customers/{id}/createGetMarketingExportDataAccessToken][%d] Customer.prototype.createGetMarketingExportDataAccessToken default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeCreateGetMarketingExportDataAccessTokenDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypeCreateGetMarketingExportDataAccessTokenDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
