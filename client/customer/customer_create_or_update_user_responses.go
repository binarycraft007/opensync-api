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

// CustomerCreateOrUpdateUserReader is a Reader for the CustomerCreateOrUpdateUser structure.
type CustomerCreateOrUpdateUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerCreateOrUpdateUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerCreateOrUpdateUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerCreateOrUpdateUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerCreateOrUpdateUserOK creates a CustomerCreateOrUpdateUserOK with default headers values
func NewCustomerCreateOrUpdateUserOK() *CustomerCreateOrUpdateUserOK {
	return &CustomerCreateOrUpdateUserOK{}
}

/*
CustomerCreateOrUpdateUserOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerCreateOrUpdateUserOK struct {
	Payload *models.Customer
}

// IsSuccess returns true when this customer create or update user o k response has a 2xx status code
func (o *CustomerCreateOrUpdateUserOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer create or update user o k response has a 3xx status code
func (o *CustomerCreateOrUpdateUserOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer create or update user o k response has a 4xx status code
func (o *CustomerCreateOrUpdateUserOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer create or update user o k response has a 5xx status code
func (o *CustomerCreateOrUpdateUserOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer create or update user o k response a status code equal to that given
func (o *CustomerCreateOrUpdateUserOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer create or update user o k response
func (o *CustomerCreateOrUpdateUserOK) Code() int {
	return 200
}

func (o *CustomerCreateOrUpdateUserOK) Error() string {
	return fmt.Sprintf("[PUT /Customers/createOrUpdateUser][%d] customerCreateOrUpdateUserOK  %+v", 200, o.Payload)
}

func (o *CustomerCreateOrUpdateUserOK) String() string {
	return fmt.Sprintf("[PUT /Customers/createOrUpdateUser][%d] customerCreateOrUpdateUserOK  %+v", 200, o.Payload)
}

func (o *CustomerCreateOrUpdateUserOK) GetPayload() *models.Customer {
	return o.Payload
}

func (o *CustomerCreateOrUpdateUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Customer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerCreateOrUpdateUserDefault creates a CustomerCreateOrUpdateUserDefault with default headers values
func NewCustomerCreateOrUpdateUserDefault(code int) *CustomerCreateOrUpdateUserDefault {
	return &CustomerCreateOrUpdateUserDefault{
		_statusCode: code,
	}
}

/*
CustomerCreateOrUpdateUserDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerCreateOrUpdateUserDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer create or update user default response has a 2xx status code
func (o *CustomerCreateOrUpdateUserDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer create or update user default response has a 3xx status code
func (o *CustomerCreateOrUpdateUserDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer create or update user default response has a 4xx status code
func (o *CustomerCreateOrUpdateUserDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer create or update user default response has a 5xx status code
func (o *CustomerCreateOrUpdateUserDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer create or update user default response a status code equal to that given
func (o *CustomerCreateOrUpdateUserDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer create or update user default response
func (o *CustomerCreateOrUpdateUserDefault) Code() int {
	return o._statusCode
}

func (o *CustomerCreateOrUpdateUserDefault) Error() string {
	return fmt.Sprintf("[PUT /Customers/createOrUpdateUser][%d] Customer.createOrUpdateUser default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerCreateOrUpdateUserDefault) String() string {
	return fmt.Sprintf("[PUT /Customers/createOrUpdateUser][%d] Customer.createOrUpdateUser default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerCreateOrUpdateUserDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerCreateOrUpdateUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
