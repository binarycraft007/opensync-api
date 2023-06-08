// Code generated by go-swagger; DO NOT EDIT.

package location

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/binarycraft007/opensync-api/models"
)

// LocationPrototypeFindByIDPendingWhitelistRequestsReader is a Reader for the LocationPrototypeFindByIDPendingWhitelistRequests structure.
type LocationPrototypeFindByIDPendingWhitelistRequestsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LocationPrototypeFindByIDPendingWhitelistRequestsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLocationPrototypeFindByIDPendingWhitelistRequestsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLocationPrototypeFindByIDPendingWhitelistRequestsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLocationPrototypeFindByIDPendingWhitelistRequestsOK creates a LocationPrototypeFindByIDPendingWhitelistRequestsOK with default headers values
func NewLocationPrototypeFindByIDPendingWhitelistRequestsOK() *LocationPrototypeFindByIDPendingWhitelistRequestsOK {
	return &LocationPrototypeFindByIDPendingWhitelistRequestsOK{}
}

/*
LocationPrototypeFindByIDPendingWhitelistRequestsOK describes a response with status code 200, with default header values.

Request was successful
*/
type LocationPrototypeFindByIDPendingWhitelistRequestsOK struct {
	Payload *models.PendingWhitelistRequests
}

// IsSuccess returns true when this location prototype find by Id pending whitelist requests o k response has a 2xx status code
func (o *LocationPrototypeFindByIDPendingWhitelistRequestsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this location prototype find by Id pending whitelist requests o k response has a 3xx status code
func (o *LocationPrototypeFindByIDPendingWhitelistRequestsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this location prototype find by Id pending whitelist requests o k response has a 4xx status code
func (o *LocationPrototypeFindByIDPendingWhitelistRequestsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this location prototype find by Id pending whitelist requests o k response has a 5xx status code
func (o *LocationPrototypeFindByIDPendingWhitelistRequestsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this location prototype find by Id pending whitelist requests o k response a status code equal to that given
func (o *LocationPrototypeFindByIDPendingWhitelistRequestsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the location prototype find by Id pending whitelist requests o k response
func (o *LocationPrototypeFindByIDPendingWhitelistRequestsOK) Code() int {
	return 200
}

func (o *LocationPrototypeFindByIDPendingWhitelistRequestsOK) Error() string {
	return fmt.Sprintf("[GET /Locations/{id}/_pendingWhitelistRequests/{fk}][%d] locationPrototypeFindByIdPendingWhitelistRequestsOK  %+v", 200, o.Payload)
}

func (o *LocationPrototypeFindByIDPendingWhitelistRequestsOK) String() string {
	return fmt.Sprintf("[GET /Locations/{id}/_pendingWhitelistRequests/{fk}][%d] locationPrototypeFindByIdPendingWhitelistRequestsOK  %+v", 200, o.Payload)
}

func (o *LocationPrototypeFindByIDPendingWhitelistRequestsOK) GetPayload() *models.PendingWhitelistRequests {
	return o.Payload
}

func (o *LocationPrototypeFindByIDPendingWhitelistRequestsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PendingWhitelistRequests)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLocationPrototypeFindByIDPendingWhitelistRequestsDefault creates a LocationPrototypeFindByIDPendingWhitelistRequestsDefault with default headers values
func NewLocationPrototypeFindByIDPendingWhitelistRequestsDefault(code int) *LocationPrototypeFindByIDPendingWhitelistRequestsDefault {
	return &LocationPrototypeFindByIDPendingWhitelistRequestsDefault{
		_statusCode: code,
	}
}

/*
LocationPrototypeFindByIDPendingWhitelistRequestsDefault describes a response with status code -1, with default header values.

unexpected error
*/
type LocationPrototypeFindByIDPendingWhitelistRequestsDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this location prototype find by Id pending whitelist requests default response has a 2xx status code
func (o *LocationPrototypeFindByIDPendingWhitelistRequestsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this location prototype find by Id pending whitelist requests default response has a 3xx status code
func (o *LocationPrototypeFindByIDPendingWhitelistRequestsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this location prototype find by Id pending whitelist requests default response has a 4xx status code
func (o *LocationPrototypeFindByIDPendingWhitelistRequestsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this location prototype find by Id pending whitelist requests default response has a 5xx status code
func (o *LocationPrototypeFindByIDPendingWhitelistRequestsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this location prototype find by Id pending whitelist requests default response a status code equal to that given
func (o *LocationPrototypeFindByIDPendingWhitelistRequestsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the location prototype find by Id pending whitelist requests default response
func (o *LocationPrototypeFindByIDPendingWhitelistRequestsDefault) Code() int {
	return o._statusCode
}

func (o *LocationPrototypeFindByIDPendingWhitelistRequestsDefault) Error() string {
	return fmt.Sprintf("[GET /Locations/{id}/_pendingWhitelistRequests/{fk}][%d] Location.prototype.__findById___pendingWhitelistRequests default  %+v", o._statusCode, o.Payload)
}

func (o *LocationPrototypeFindByIDPendingWhitelistRequestsDefault) String() string {
	return fmt.Sprintf("[GET /Locations/{id}/_pendingWhitelistRequests/{fk}][%d] Location.prototype.__findById___pendingWhitelistRequests default  %+v", o._statusCode, o.Payload)
}

func (o *LocationPrototypeFindByIDPendingWhitelistRequestsDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *LocationPrototypeFindByIDPendingWhitelistRequestsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
