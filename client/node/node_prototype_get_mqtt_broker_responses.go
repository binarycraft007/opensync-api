// Code generated by go-swagger; DO NOT EDIT.

package node

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/binarycraft007/opensync-api/models"
)

// NodePrototypeGetMqttBrokerReader is a Reader for the NodePrototypeGetMqttBroker structure.
type NodePrototypeGetMqttBrokerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NodePrototypeGetMqttBrokerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNodePrototypeGetMqttBrokerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNodePrototypeGetMqttBrokerDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNodePrototypeGetMqttBrokerOK creates a NodePrototypeGetMqttBrokerOK with default headers values
func NewNodePrototypeGetMqttBrokerOK() *NodePrototypeGetMqttBrokerOK {
	return &NodePrototypeGetMqttBrokerOK{}
}

/*
NodePrototypeGetMqttBrokerOK describes a response with status code 200, with default header values.

Request was successful
*/
type NodePrototypeGetMqttBrokerOK struct {
	Payload interface{}
}

// IsSuccess returns true when this node prototype get mqtt broker o k response has a 2xx status code
func (o *NodePrototypeGetMqttBrokerOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this node prototype get mqtt broker o k response has a 3xx status code
func (o *NodePrototypeGetMqttBrokerOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this node prototype get mqtt broker o k response has a 4xx status code
func (o *NodePrototypeGetMqttBrokerOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this node prototype get mqtt broker o k response has a 5xx status code
func (o *NodePrototypeGetMqttBrokerOK) IsServerError() bool {
	return false
}

// IsCode returns true when this node prototype get mqtt broker o k response a status code equal to that given
func (o *NodePrototypeGetMqttBrokerOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the node prototype get mqtt broker o k response
func (o *NodePrototypeGetMqttBrokerOK) Code() int {
	return 200
}

func (o *NodePrototypeGetMqttBrokerOK) Error() string {
	return fmt.Sprintf("[GET /Nodes/{id}/mqtt][%d] nodePrototypeGetMqttBrokerOK  %+v", 200, o.Payload)
}

func (o *NodePrototypeGetMqttBrokerOK) String() string {
	return fmt.Sprintf("[GET /Nodes/{id}/mqtt][%d] nodePrototypeGetMqttBrokerOK  %+v", 200, o.Payload)
}

func (o *NodePrototypeGetMqttBrokerOK) GetPayload() interface{} {
	return o.Payload
}

func (o *NodePrototypeGetMqttBrokerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNodePrototypeGetMqttBrokerDefault creates a NodePrototypeGetMqttBrokerDefault with default headers values
func NewNodePrototypeGetMqttBrokerDefault(code int) *NodePrototypeGetMqttBrokerDefault {
	return &NodePrototypeGetMqttBrokerDefault{
		_statusCode: code,
	}
}

/*
NodePrototypeGetMqttBrokerDefault describes a response with status code -1, with default header values.

unexpected error
*/
type NodePrototypeGetMqttBrokerDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this node prototype get mqtt broker default response has a 2xx status code
func (o *NodePrototypeGetMqttBrokerDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this node prototype get mqtt broker default response has a 3xx status code
func (o *NodePrototypeGetMqttBrokerDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this node prototype get mqtt broker default response has a 4xx status code
func (o *NodePrototypeGetMqttBrokerDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this node prototype get mqtt broker default response has a 5xx status code
func (o *NodePrototypeGetMqttBrokerDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this node prototype get mqtt broker default response a status code equal to that given
func (o *NodePrototypeGetMqttBrokerDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the node prototype get mqtt broker default response
func (o *NodePrototypeGetMqttBrokerDefault) Code() int {
	return o._statusCode
}

func (o *NodePrototypeGetMqttBrokerDefault) Error() string {
	return fmt.Sprintf("[GET /Nodes/{id}/mqtt][%d] Node.prototype.getMqttBroker default  %+v", o._statusCode, o.Payload)
}

func (o *NodePrototypeGetMqttBrokerDefault) String() string {
	return fmt.Sprintf("[GET /Nodes/{id}/mqtt][%d] Node.prototype.getMqttBroker default  %+v", o._statusCode, o.Payload)
}

func (o *NodePrototypeGetMqttBrokerDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *NodePrototypeGetMqttBrokerDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
