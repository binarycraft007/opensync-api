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

// LocationPrototypeGetNodesReader is a Reader for the LocationPrototypeGetNodes structure.
type LocationPrototypeGetNodesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LocationPrototypeGetNodesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLocationPrototypeGetNodesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLocationPrototypeGetNodesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLocationPrototypeGetNodesOK creates a LocationPrototypeGetNodesOK with default headers values
func NewLocationPrototypeGetNodesOK() *LocationPrototypeGetNodesOK {
	return &LocationPrototypeGetNodesOK{}
}

/*
LocationPrototypeGetNodesOK describes a response with status code 200, with default header values.

Request was successful
*/
type LocationPrototypeGetNodesOK struct {
	Payload []*models.Node
}

// IsSuccess returns true when this location prototype get nodes o k response has a 2xx status code
func (o *LocationPrototypeGetNodesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this location prototype get nodes o k response has a 3xx status code
func (o *LocationPrototypeGetNodesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this location prototype get nodes o k response has a 4xx status code
func (o *LocationPrototypeGetNodesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this location prototype get nodes o k response has a 5xx status code
func (o *LocationPrototypeGetNodesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this location prototype get nodes o k response a status code equal to that given
func (o *LocationPrototypeGetNodesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the location prototype get nodes o k response
func (o *LocationPrototypeGetNodesOK) Code() int {
	return 200
}

func (o *LocationPrototypeGetNodesOK) Error() string {
	return fmt.Sprintf("[GET /Locations/{id}/nodes][%d] locationPrototypeGetNodesOK  %+v", 200, o.Payload)
}

func (o *LocationPrototypeGetNodesOK) String() string {
	return fmt.Sprintf("[GET /Locations/{id}/nodes][%d] locationPrototypeGetNodesOK  %+v", 200, o.Payload)
}

func (o *LocationPrototypeGetNodesOK) GetPayload() []*models.Node {
	return o.Payload
}

func (o *LocationPrototypeGetNodesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLocationPrototypeGetNodesDefault creates a LocationPrototypeGetNodesDefault with default headers values
func NewLocationPrototypeGetNodesDefault(code int) *LocationPrototypeGetNodesDefault {
	return &LocationPrototypeGetNodesDefault{
		_statusCode: code,
	}
}

/*
LocationPrototypeGetNodesDefault describes a response with status code -1, with default header values.

unexpected error
*/
type LocationPrototypeGetNodesDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this location prototype get nodes default response has a 2xx status code
func (o *LocationPrototypeGetNodesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this location prototype get nodes default response has a 3xx status code
func (o *LocationPrototypeGetNodesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this location prototype get nodes default response has a 4xx status code
func (o *LocationPrototypeGetNodesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this location prototype get nodes default response has a 5xx status code
func (o *LocationPrototypeGetNodesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this location prototype get nodes default response a status code equal to that given
func (o *LocationPrototypeGetNodesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the location prototype get nodes default response
func (o *LocationPrototypeGetNodesDefault) Code() int {
	return o._statusCode
}

func (o *LocationPrototypeGetNodesDefault) Error() string {
	return fmt.Sprintf("[GET /Locations/{id}/nodes][%d] Location.prototype.__get__nodes default  %+v", o._statusCode, o.Payload)
}

func (o *LocationPrototypeGetNodesDefault) String() string {
	return fmt.Sprintf("[GET /Locations/{id}/nodes][%d] Location.prototype.__get__nodes default  %+v", o._statusCode, o.Payload)
}

func (o *LocationPrototypeGetNodesDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *LocationPrototypeGetNodesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
