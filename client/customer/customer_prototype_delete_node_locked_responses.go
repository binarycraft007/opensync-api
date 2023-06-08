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

// CustomerPrototypeDeleteNodeLockedReader is a Reader for the CustomerPrototypeDeleteNodeLocked structure.
type CustomerPrototypeDeleteNodeLockedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypeDeleteNodeLockedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypeDeleteNodeLockedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypeDeleteNodeLockedDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypeDeleteNodeLockedOK creates a CustomerPrototypeDeleteNodeLockedOK with default headers values
func NewCustomerPrototypeDeleteNodeLockedOK() *CustomerPrototypeDeleteNodeLockedOK {
	return &CustomerPrototypeDeleteNodeLockedOK{}
}

/*
CustomerPrototypeDeleteNodeLockedOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypeDeleteNodeLockedOK struct {
	Payload interface{}
}

// IsSuccess returns true when this customer prototype delete node locked o k response has a 2xx status code
func (o *CustomerPrototypeDeleteNodeLockedOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype delete node locked o k response has a 3xx status code
func (o *CustomerPrototypeDeleteNodeLockedOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype delete node locked o k response has a 4xx status code
func (o *CustomerPrototypeDeleteNodeLockedOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype delete node locked o k response has a 5xx status code
func (o *CustomerPrototypeDeleteNodeLockedOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype delete node locked o k response a status code equal to that given
func (o *CustomerPrototypeDeleteNodeLockedOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype delete node locked o k response
func (o *CustomerPrototypeDeleteNodeLockedOK) Code() int {
	return 200
}

func (o *CustomerPrototypeDeleteNodeLockedOK) Error() string {
	return fmt.Sprintf("[DELETE /Customers/{id}/nodes/{nodeId}][%d] customerPrototypeDeleteNodeLockedOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeDeleteNodeLockedOK) String() string {
	return fmt.Sprintf("[DELETE /Customers/{id}/nodes/{nodeId}][%d] customerPrototypeDeleteNodeLockedOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeDeleteNodeLockedOK) GetPayload() interface{} {
	return o.Payload
}

func (o *CustomerPrototypeDeleteNodeLockedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypeDeleteNodeLockedDefault creates a CustomerPrototypeDeleteNodeLockedDefault with default headers values
func NewCustomerPrototypeDeleteNodeLockedDefault(code int) *CustomerPrototypeDeleteNodeLockedDefault {
	return &CustomerPrototypeDeleteNodeLockedDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypeDeleteNodeLockedDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypeDeleteNodeLockedDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype delete node locked default response has a 2xx status code
func (o *CustomerPrototypeDeleteNodeLockedDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype delete node locked default response has a 3xx status code
func (o *CustomerPrototypeDeleteNodeLockedDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype delete node locked default response has a 4xx status code
func (o *CustomerPrototypeDeleteNodeLockedDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype delete node locked default response has a 5xx status code
func (o *CustomerPrototypeDeleteNodeLockedDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype delete node locked default response a status code equal to that given
func (o *CustomerPrototypeDeleteNodeLockedDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype delete node locked default response
func (o *CustomerPrototypeDeleteNodeLockedDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypeDeleteNodeLockedDefault) Error() string {
	return fmt.Sprintf("[DELETE /Customers/{id}/nodes/{nodeId}][%d] Customer.prototype.deleteNodeLocked default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeDeleteNodeLockedDefault) String() string {
	return fmt.Sprintf("[DELETE /Customers/{id}/nodes/{nodeId}][%d] Customer.prototype.deleteNodeLocked default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeDeleteNodeLockedDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypeDeleteNodeLockedDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}