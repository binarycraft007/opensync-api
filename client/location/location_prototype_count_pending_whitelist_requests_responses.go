// Code generated by go-swagger; DO NOT EDIT.

package location

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/binarycraft007/opensync-api/models"
)

// LocationPrototypeCountPendingWhitelistRequestsReader is a Reader for the LocationPrototypeCountPendingWhitelistRequests structure.
type LocationPrototypeCountPendingWhitelistRequestsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LocationPrototypeCountPendingWhitelistRequestsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLocationPrototypeCountPendingWhitelistRequestsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLocationPrototypeCountPendingWhitelistRequestsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLocationPrototypeCountPendingWhitelistRequestsOK creates a LocationPrototypeCountPendingWhitelistRequestsOK with default headers values
func NewLocationPrototypeCountPendingWhitelistRequestsOK() *LocationPrototypeCountPendingWhitelistRequestsOK {
	return &LocationPrototypeCountPendingWhitelistRequestsOK{}
}

/*
LocationPrototypeCountPendingWhitelistRequestsOK describes a response with status code 200, with default header values.

Request was successful
*/
type LocationPrototypeCountPendingWhitelistRequestsOK struct {
	Payload *LocationPrototypeCountPendingWhitelistRequestsOKBody
}

// IsSuccess returns true when this location prototype count pending whitelist requests o k response has a 2xx status code
func (o *LocationPrototypeCountPendingWhitelistRequestsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this location prototype count pending whitelist requests o k response has a 3xx status code
func (o *LocationPrototypeCountPendingWhitelistRequestsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this location prototype count pending whitelist requests o k response has a 4xx status code
func (o *LocationPrototypeCountPendingWhitelistRequestsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this location prototype count pending whitelist requests o k response has a 5xx status code
func (o *LocationPrototypeCountPendingWhitelistRequestsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this location prototype count pending whitelist requests o k response a status code equal to that given
func (o *LocationPrototypeCountPendingWhitelistRequestsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the location prototype count pending whitelist requests o k response
func (o *LocationPrototypeCountPendingWhitelistRequestsOK) Code() int {
	return 200
}

func (o *LocationPrototypeCountPendingWhitelistRequestsOK) Error() string {
	return fmt.Sprintf("[GET /Locations/{id}/_pendingWhitelistRequests/count][%d] locationPrototypeCountPendingWhitelistRequestsOK  %+v", 200, o.Payload)
}

func (o *LocationPrototypeCountPendingWhitelistRequestsOK) String() string {
	return fmt.Sprintf("[GET /Locations/{id}/_pendingWhitelistRequests/count][%d] locationPrototypeCountPendingWhitelistRequestsOK  %+v", 200, o.Payload)
}

func (o *LocationPrototypeCountPendingWhitelistRequestsOK) GetPayload() *LocationPrototypeCountPendingWhitelistRequestsOKBody {
	return o.Payload
}

func (o *LocationPrototypeCountPendingWhitelistRequestsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LocationPrototypeCountPendingWhitelistRequestsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLocationPrototypeCountPendingWhitelistRequestsDefault creates a LocationPrototypeCountPendingWhitelistRequestsDefault with default headers values
func NewLocationPrototypeCountPendingWhitelistRequestsDefault(code int) *LocationPrototypeCountPendingWhitelistRequestsDefault {
	return &LocationPrototypeCountPendingWhitelistRequestsDefault{
		_statusCode: code,
	}
}

/*
LocationPrototypeCountPendingWhitelistRequestsDefault describes a response with status code -1, with default header values.

unexpected error
*/
type LocationPrototypeCountPendingWhitelistRequestsDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this location prototype count pending whitelist requests default response has a 2xx status code
func (o *LocationPrototypeCountPendingWhitelistRequestsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this location prototype count pending whitelist requests default response has a 3xx status code
func (o *LocationPrototypeCountPendingWhitelistRequestsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this location prototype count pending whitelist requests default response has a 4xx status code
func (o *LocationPrototypeCountPendingWhitelistRequestsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this location prototype count pending whitelist requests default response has a 5xx status code
func (o *LocationPrototypeCountPendingWhitelistRequestsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this location prototype count pending whitelist requests default response a status code equal to that given
func (o *LocationPrototypeCountPendingWhitelistRequestsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the location prototype count pending whitelist requests default response
func (o *LocationPrototypeCountPendingWhitelistRequestsDefault) Code() int {
	return o._statusCode
}

func (o *LocationPrototypeCountPendingWhitelistRequestsDefault) Error() string {
	return fmt.Sprintf("[GET /Locations/{id}/_pendingWhitelistRequests/count][%d] Location.prototype.__count___pendingWhitelistRequests default  %+v", o._statusCode, o.Payload)
}

func (o *LocationPrototypeCountPendingWhitelistRequestsDefault) String() string {
	return fmt.Sprintf("[GET /Locations/{id}/_pendingWhitelistRequests/count][%d] Location.prototype.__count___pendingWhitelistRequests default  %+v", o._statusCode, o.Payload)
}

func (o *LocationPrototypeCountPendingWhitelistRequestsDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *LocationPrototypeCountPendingWhitelistRequestsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
LocationPrototypeCountPendingWhitelistRequestsOKBody location prototype count pending whitelist requests o k body
swagger:model LocationPrototypeCountPendingWhitelistRequestsOKBody
*/
type LocationPrototypeCountPendingWhitelistRequestsOKBody struct {

	// count
	Count float64 `json:"count,omitempty"`
}

// Validate validates this location prototype count pending whitelist requests o k body
func (o *LocationPrototypeCountPendingWhitelistRequestsOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this location prototype count pending whitelist requests o k body based on context it is used
func (o *LocationPrototypeCountPendingWhitelistRequestsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LocationPrototypeCountPendingWhitelistRequestsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LocationPrototypeCountPendingWhitelistRequestsOKBody) UnmarshalBinary(b []byte) error {
	var res LocationPrototypeCountPendingWhitelistRequestsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
