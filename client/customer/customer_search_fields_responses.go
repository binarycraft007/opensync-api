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

// CustomerSearchFieldsReader is a Reader for the CustomerSearchFields structure.
type CustomerSearchFieldsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerSearchFieldsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerSearchFieldsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerSearchFieldsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerSearchFieldsOK creates a CustomerSearchFieldsOK with default headers values
func NewCustomerSearchFieldsOK() *CustomerSearchFieldsOK {
	return &CustomerSearchFieldsOK{}
}

/*
CustomerSearchFieldsOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerSearchFieldsOK struct {
	Payload models.XAny
}

// IsSuccess returns true when this customer search fields o k response has a 2xx status code
func (o *CustomerSearchFieldsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer search fields o k response has a 3xx status code
func (o *CustomerSearchFieldsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer search fields o k response has a 4xx status code
func (o *CustomerSearchFieldsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer search fields o k response has a 5xx status code
func (o *CustomerSearchFieldsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer search fields o k response a status code equal to that given
func (o *CustomerSearchFieldsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer search fields o k response
func (o *CustomerSearchFieldsOK) Code() int {
	return 200
}

func (o *CustomerSearchFieldsOK) Error() string {
	return fmt.Sprintf("[GET /Customers/search/{keyword}][%d] customerSearchFieldsOK  %+v", 200, o.Payload)
}

func (o *CustomerSearchFieldsOK) String() string {
	return fmt.Sprintf("[GET /Customers/search/{keyword}][%d] customerSearchFieldsOK  %+v", 200, o.Payload)
}

func (o *CustomerSearchFieldsOK) GetPayload() models.XAny {
	return o.Payload
}

func (o *CustomerSearchFieldsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerSearchFieldsDefault creates a CustomerSearchFieldsDefault with default headers values
func NewCustomerSearchFieldsDefault(code int) *CustomerSearchFieldsDefault {
	return &CustomerSearchFieldsDefault{
		_statusCode: code,
	}
}

/*
CustomerSearchFieldsDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerSearchFieldsDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer search fields default response has a 2xx status code
func (o *CustomerSearchFieldsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer search fields default response has a 3xx status code
func (o *CustomerSearchFieldsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer search fields default response has a 4xx status code
func (o *CustomerSearchFieldsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer search fields default response has a 5xx status code
func (o *CustomerSearchFieldsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer search fields default response a status code equal to that given
func (o *CustomerSearchFieldsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer search fields default response
func (o *CustomerSearchFieldsDefault) Code() int {
	return o._statusCode
}

func (o *CustomerSearchFieldsDefault) Error() string {
	return fmt.Sprintf("[GET /Customers/search/{keyword}][%d] Customer.searchFields default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerSearchFieldsDefault) String() string {
	return fmt.Sprintf("[GET /Customers/search/{keyword}][%d] Customer.searchFields default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerSearchFieldsDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerSearchFieldsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
