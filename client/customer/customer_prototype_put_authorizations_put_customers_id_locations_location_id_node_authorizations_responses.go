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

// CustomerPrototypePutAuthorizationsPutCustomersIDLocationsLocationIDNodeAuthorizationsReader is a Reader for the CustomerPrototypePutAuthorizationsPutCustomersIDLocationsLocationIDNodeAuthorizations structure.
type CustomerPrototypePutAuthorizationsPutCustomersIDLocationsLocationIDNodeAuthorizationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypePutAuthorizationsPutCustomersIDLocationsLocationIDNodeAuthorizationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypePutAuthorizationsPutCustomersIDLocationsLocationIDNodeAuthorizationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypePutAuthorizationsPutCustomersIDLocationsLocationIDNodeAuthorizationsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypePutAuthorizationsPutCustomersIDLocationsLocationIDNodeAuthorizationsOK creates a CustomerPrototypePutAuthorizationsPutCustomersIDLocationsLocationIDNodeAuthorizationsOK with default headers values
func NewCustomerPrototypePutAuthorizationsPutCustomersIDLocationsLocationIDNodeAuthorizationsOK() *CustomerPrototypePutAuthorizationsPutCustomersIDLocationsLocationIDNodeAuthorizationsOK {
	return &CustomerPrototypePutAuthorizationsPutCustomersIDLocationsLocationIDNodeAuthorizationsOK{}
}

/*
CustomerPrototypePutAuthorizationsPutCustomersIDLocationsLocationIDNodeAuthorizationsOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypePutAuthorizationsPutCustomersIDLocationsLocationIDNodeAuthorizationsOK struct {
	Payload *models.Authorizations
}

// IsSuccess returns true when this customer prototype put authorizations put customers Id locations location Id node authorizations o k response has a 2xx status code
func (o *CustomerPrototypePutAuthorizationsPutCustomersIDLocationsLocationIDNodeAuthorizationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype put authorizations put customers Id locations location Id node authorizations o k response has a 3xx status code
func (o *CustomerPrototypePutAuthorizationsPutCustomersIDLocationsLocationIDNodeAuthorizationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype put authorizations put customers Id locations location Id node authorizations o k response has a 4xx status code
func (o *CustomerPrototypePutAuthorizationsPutCustomersIDLocationsLocationIDNodeAuthorizationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype put authorizations put customers Id locations location Id node authorizations o k response has a 5xx status code
func (o *CustomerPrototypePutAuthorizationsPutCustomersIDLocationsLocationIDNodeAuthorizationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype put authorizations put customers Id locations location Id node authorizations o k response a status code equal to that given
func (o *CustomerPrototypePutAuthorizationsPutCustomersIDLocationsLocationIDNodeAuthorizationsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype put authorizations put customers Id locations location Id node authorizations o k response
func (o *CustomerPrototypePutAuthorizationsPutCustomersIDLocationsLocationIDNodeAuthorizationsOK) Code() int {
	return 200
}

func (o *CustomerPrototypePutAuthorizationsPutCustomersIDLocationsLocationIDNodeAuthorizationsOK) Error() string {
	return fmt.Sprintf("[PUT /Customers/{id}/locations/{locationId}/nodeAuthorizations][%d] customerPrototypePutAuthorizationsPutCustomersIdLocationsLocationIdNodeAuthorizationsOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypePutAuthorizationsPutCustomersIDLocationsLocationIDNodeAuthorizationsOK) String() string {
	return fmt.Sprintf("[PUT /Customers/{id}/locations/{locationId}/nodeAuthorizations][%d] customerPrototypePutAuthorizationsPutCustomersIdLocationsLocationIdNodeAuthorizationsOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypePutAuthorizationsPutCustomersIDLocationsLocationIDNodeAuthorizationsOK) GetPayload() *models.Authorizations {
	return o.Payload
}

func (o *CustomerPrototypePutAuthorizationsPutCustomersIDLocationsLocationIDNodeAuthorizationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Authorizations)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypePutAuthorizationsPutCustomersIDLocationsLocationIDNodeAuthorizationsDefault creates a CustomerPrototypePutAuthorizationsPutCustomersIDLocationsLocationIDNodeAuthorizationsDefault with default headers values
func NewCustomerPrototypePutAuthorizationsPutCustomersIDLocationsLocationIDNodeAuthorizationsDefault(code int) *CustomerPrototypePutAuthorizationsPutCustomersIDLocationsLocationIDNodeAuthorizationsDefault {
	return &CustomerPrototypePutAuthorizationsPutCustomersIDLocationsLocationIDNodeAuthorizationsDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypePutAuthorizationsPutCustomersIDLocationsLocationIDNodeAuthorizationsDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypePutAuthorizationsPutCustomersIDLocationsLocationIDNodeAuthorizationsDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype put authorizations put customers id locations location Id node authorizations default response has a 2xx status code
func (o *CustomerPrototypePutAuthorizationsPutCustomersIDLocationsLocationIDNodeAuthorizationsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype put authorizations put customers id locations location Id node authorizations default response has a 3xx status code
func (o *CustomerPrototypePutAuthorizationsPutCustomersIDLocationsLocationIDNodeAuthorizationsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype put authorizations put customers id locations location Id node authorizations default response has a 4xx status code
func (o *CustomerPrototypePutAuthorizationsPutCustomersIDLocationsLocationIDNodeAuthorizationsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype put authorizations put customers id locations location Id node authorizations default response has a 5xx status code
func (o *CustomerPrototypePutAuthorizationsPutCustomersIDLocationsLocationIDNodeAuthorizationsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype put authorizations put customers id locations location Id node authorizations default response a status code equal to that given
func (o *CustomerPrototypePutAuthorizationsPutCustomersIDLocationsLocationIDNodeAuthorizationsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype put authorizations put customers id locations location Id node authorizations default response
func (o *CustomerPrototypePutAuthorizationsPutCustomersIDLocationsLocationIDNodeAuthorizationsDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypePutAuthorizationsPutCustomersIDLocationsLocationIDNodeAuthorizationsDefault) Error() string {
	return fmt.Sprintf("[PUT /Customers/{id}/locations/{locationId}/nodeAuthorizations][%d] Customer.prototype.putAuthorizations__put_Customers_{id}_locations_{locationId}_nodeAuthorizations default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePutAuthorizationsPutCustomersIDLocationsLocationIDNodeAuthorizationsDefault) String() string {
	return fmt.Sprintf("[PUT /Customers/{id}/locations/{locationId}/nodeAuthorizations][%d] Customer.prototype.putAuthorizations__put_Customers_{id}_locations_{locationId}_nodeAuthorizations default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePutAuthorizationsPutCustomersIDLocationsLocationIDNodeAuthorizationsDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypePutAuthorizationsPutCustomersIDLocationsLocationIDNodeAuthorizationsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
