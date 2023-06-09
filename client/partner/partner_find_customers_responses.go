// Code generated by go-swagger; DO NOT EDIT.

package partner

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/binarycraft007/opensync-api/models"
)

// PartnerFindCustomersReader is a Reader for the PartnerFindCustomers structure.
type PartnerFindCustomersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PartnerFindCustomersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPartnerFindCustomersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPartnerFindCustomersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPartnerFindCustomersOK creates a PartnerFindCustomersOK with default headers values
func NewPartnerFindCustomersOK() *PartnerFindCustomersOK {
	return &PartnerFindCustomersOK{}
}

/*
PartnerFindCustomersOK describes a response with status code 200, with default header values.

Request was successful
*/
type PartnerFindCustomersOK struct {
	Payload interface{}
}

// IsSuccess returns true when this partner find customers o k response has a 2xx status code
func (o *PartnerFindCustomersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this partner find customers o k response has a 3xx status code
func (o *PartnerFindCustomersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this partner find customers o k response has a 4xx status code
func (o *PartnerFindCustomersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this partner find customers o k response has a 5xx status code
func (o *PartnerFindCustomersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this partner find customers o k response a status code equal to that given
func (o *PartnerFindCustomersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the partner find customers o k response
func (o *PartnerFindCustomersOK) Code() int {
	return 200
}

func (o *PartnerFindCustomersOK) Error() string {
	return fmt.Sprintf("[GET /partners/customers/search/{keyword}][%d] partnerFindCustomersOK  %+v", 200, o.Payload)
}

func (o *PartnerFindCustomersOK) String() string {
	return fmt.Sprintf("[GET /partners/customers/search/{keyword}][%d] partnerFindCustomersOK  %+v", 200, o.Payload)
}

func (o *PartnerFindCustomersOK) GetPayload() interface{} {
	return o.Payload
}

func (o *PartnerFindCustomersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPartnerFindCustomersDefault creates a PartnerFindCustomersDefault with default headers values
func NewPartnerFindCustomersDefault(code int) *PartnerFindCustomersDefault {
	return &PartnerFindCustomersDefault{
		_statusCode: code,
	}
}

/*
PartnerFindCustomersDefault describes a response with status code -1, with default header values.

unexpected error
*/
type PartnerFindCustomersDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this partner find customers default response has a 2xx status code
func (o *PartnerFindCustomersDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this partner find customers default response has a 3xx status code
func (o *PartnerFindCustomersDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this partner find customers default response has a 4xx status code
func (o *PartnerFindCustomersDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this partner find customers default response has a 5xx status code
func (o *PartnerFindCustomersDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this partner find customers default response a status code equal to that given
func (o *PartnerFindCustomersDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the partner find customers default response
func (o *PartnerFindCustomersDefault) Code() int {
	return o._statusCode
}

func (o *PartnerFindCustomersDefault) Error() string {
	return fmt.Sprintf("[GET /partners/customers/search/{keyword}][%d] Partner.findCustomers default  %+v", o._statusCode, o.Payload)
}

func (o *PartnerFindCustomersDefault) String() string {
	return fmt.Sprintf("[GET /partners/customers/search/{keyword}][%d] Partner.findCustomers default  %+v", o._statusCode, o.Payload)
}

func (o *PartnerFindCustomersDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *PartnerFindCustomersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
