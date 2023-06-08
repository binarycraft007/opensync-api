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

// PartnerGetLocationsReader is a Reader for the PartnerGetLocations structure.
type PartnerGetLocationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PartnerGetLocationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPartnerGetLocationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPartnerGetLocationsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPartnerGetLocationsOK creates a PartnerGetLocationsOK with default headers values
func NewPartnerGetLocationsOK() *PartnerGetLocationsOK {
	return &PartnerGetLocationsOK{}
}

/*
PartnerGetLocationsOK describes a response with status code 200, with default header values.

Request was successful
*/
type PartnerGetLocationsOK struct {
	Payload interface{}
}

// IsSuccess returns true when this partner get locations o k response has a 2xx status code
func (o *PartnerGetLocationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this partner get locations o k response has a 3xx status code
func (o *PartnerGetLocationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this partner get locations o k response has a 4xx status code
func (o *PartnerGetLocationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this partner get locations o k response has a 5xx status code
func (o *PartnerGetLocationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this partner get locations o k response a status code equal to that given
func (o *PartnerGetLocationsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the partner get locations o k response
func (o *PartnerGetLocationsOK) Code() int {
	return 200
}

func (o *PartnerGetLocationsOK) Error() string {
	return fmt.Sprintf("[GET /partners/locations/{keyword}][%d] partnerGetLocationsOK  %+v", 200, o.Payload)
}

func (o *PartnerGetLocationsOK) String() string {
	return fmt.Sprintf("[GET /partners/locations/{keyword}][%d] partnerGetLocationsOK  %+v", 200, o.Payload)
}

func (o *PartnerGetLocationsOK) GetPayload() interface{} {
	return o.Payload
}

func (o *PartnerGetLocationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPartnerGetLocationsDefault creates a PartnerGetLocationsDefault with default headers values
func NewPartnerGetLocationsDefault(code int) *PartnerGetLocationsDefault {
	return &PartnerGetLocationsDefault{
		_statusCode: code,
	}
}

/*
PartnerGetLocationsDefault describes a response with status code -1, with default header values.

unexpected error
*/
type PartnerGetLocationsDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this partner get locations default response has a 2xx status code
func (o *PartnerGetLocationsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this partner get locations default response has a 3xx status code
func (o *PartnerGetLocationsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this partner get locations default response has a 4xx status code
func (o *PartnerGetLocationsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this partner get locations default response has a 5xx status code
func (o *PartnerGetLocationsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this partner get locations default response a status code equal to that given
func (o *PartnerGetLocationsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the partner get locations default response
func (o *PartnerGetLocationsDefault) Code() int {
	return o._statusCode
}

func (o *PartnerGetLocationsDefault) Error() string {
	return fmt.Sprintf("[GET /partners/locations/{keyword}][%d] Partner.getLocations default  %+v", o._statusCode, o.Payload)
}

func (o *PartnerGetLocationsDefault) String() string {
	return fmt.Sprintf("[GET /partners/locations/{keyword}][%d] Partner.getLocations default  %+v", o._statusCode, o.Payload)
}

func (o *PartnerGetLocationsDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *PartnerGetLocationsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}