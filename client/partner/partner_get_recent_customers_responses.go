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

// PartnerGetRecentCustomersReader is a Reader for the PartnerGetRecentCustomers structure.
type PartnerGetRecentCustomersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PartnerGetRecentCustomersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPartnerGetRecentCustomersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPartnerGetRecentCustomersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPartnerGetRecentCustomersOK creates a PartnerGetRecentCustomersOK with default headers values
func NewPartnerGetRecentCustomersOK() *PartnerGetRecentCustomersOK {
	return &PartnerGetRecentCustomersOK{}
}

/*
PartnerGetRecentCustomersOK describes a response with status code 200, with default header values.

Request was successful
*/
type PartnerGetRecentCustomersOK struct {
	Payload interface{}
}

// IsSuccess returns true when this partner get recent customers o k response has a 2xx status code
func (o *PartnerGetRecentCustomersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this partner get recent customers o k response has a 3xx status code
func (o *PartnerGetRecentCustomersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this partner get recent customers o k response has a 4xx status code
func (o *PartnerGetRecentCustomersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this partner get recent customers o k response has a 5xx status code
func (o *PartnerGetRecentCustomersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this partner get recent customers o k response a status code equal to that given
func (o *PartnerGetRecentCustomersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the partner get recent customers o k response
func (o *PartnerGetRecentCustomersOK) Code() int {
	return 200
}

func (o *PartnerGetRecentCustomersOK) Error() string {
	return fmt.Sprintf("[GET /partners/{id}/customers/recent][%d] partnerGetRecentCustomersOK  %+v", 200, o.Payload)
}

func (o *PartnerGetRecentCustomersOK) String() string {
	return fmt.Sprintf("[GET /partners/{id}/customers/recent][%d] partnerGetRecentCustomersOK  %+v", 200, o.Payload)
}

func (o *PartnerGetRecentCustomersOK) GetPayload() interface{} {
	return o.Payload
}

func (o *PartnerGetRecentCustomersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPartnerGetRecentCustomersDefault creates a PartnerGetRecentCustomersDefault with default headers values
func NewPartnerGetRecentCustomersDefault(code int) *PartnerGetRecentCustomersDefault {
	return &PartnerGetRecentCustomersDefault{
		_statusCode: code,
	}
}

/*
PartnerGetRecentCustomersDefault describes a response with status code -1, with default header values.

unexpected error
*/
type PartnerGetRecentCustomersDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this partner get recent customers default response has a 2xx status code
func (o *PartnerGetRecentCustomersDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this partner get recent customers default response has a 3xx status code
func (o *PartnerGetRecentCustomersDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this partner get recent customers default response has a 4xx status code
func (o *PartnerGetRecentCustomersDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this partner get recent customers default response has a 5xx status code
func (o *PartnerGetRecentCustomersDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this partner get recent customers default response a status code equal to that given
func (o *PartnerGetRecentCustomersDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the partner get recent customers default response
func (o *PartnerGetRecentCustomersDefault) Code() int {
	return o._statusCode
}

func (o *PartnerGetRecentCustomersDefault) Error() string {
	return fmt.Sprintf("[GET /partners/{id}/customers/recent][%d] Partner.getRecentCustomers default  %+v", o._statusCode, o.Payload)
}

func (o *PartnerGetRecentCustomersDefault) String() string {
	return fmt.Sprintf("[GET /partners/{id}/customers/recent][%d] Partner.getRecentCustomers default  %+v", o._statusCode, o.Payload)
}

func (o *PartnerGetRecentCustomersDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *PartnerGetRecentCustomersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
