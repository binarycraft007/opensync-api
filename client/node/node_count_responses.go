// Code generated by go-swagger; DO NOT EDIT.

package node

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

// NodeCountReader is a Reader for the NodeCount structure.
type NodeCountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NodeCountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNodeCountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNodeCountDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNodeCountOK creates a NodeCountOK with default headers values
func NewNodeCountOK() *NodeCountOK {
	return &NodeCountOK{}
}

/*
NodeCountOK describes a response with status code 200, with default header values.

Request was successful
*/
type NodeCountOK struct {
	Payload *NodeCountOKBody
}

// IsSuccess returns true when this node count o k response has a 2xx status code
func (o *NodeCountOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this node count o k response has a 3xx status code
func (o *NodeCountOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this node count o k response has a 4xx status code
func (o *NodeCountOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this node count o k response has a 5xx status code
func (o *NodeCountOK) IsServerError() bool {
	return false
}

// IsCode returns true when this node count o k response a status code equal to that given
func (o *NodeCountOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the node count o k response
func (o *NodeCountOK) Code() int {
	return 200
}

func (o *NodeCountOK) Error() string {
	return fmt.Sprintf("[GET /Nodes/count][%d] nodeCountOK  %+v", 200, o.Payload)
}

func (o *NodeCountOK) String() string {
	return fmt.Sprintf("[GET /Nodes/count][%d] nodeCountOK  %+v", 200, o.Payload)
}

func (o *NodeCountOK) GetPayload() *NodeCountOKBody {
	return o.Payload
}

func (o *NodeCountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(NodeCountOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNodeCountDefault creates a NodeCountDefault with default headers values
func NewNodeCountDefault(code int) *NodeCountDefault {
	return &NodeCountDefault{
		_statusCode: code,
	}
}

/*
NodeCountDefault describes a response with status code -1, with default header values.

unexpected error
*/
type NodeCountDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this node count default response has a 2xx status code
func (o *NodeCountDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this node count default response has a 3xx status code
func (o *NodeCountDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this node count default response has a 4xx status code
func (o *NodeCountDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this node count default response has a 5xx status code
func (o *NodeCountDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this node count default response a status code equal to that given
func (o *NodeCountDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the node count default response
func (o *NodeCountDefault) Code() int {
	return o._statusCode
}

func (o *NodeCountDefault) Error() string {
	return fmt.Sprintf("[GET /Nodes/count][%d] Node.count default  %+v", o._statusCode, o.Payload)
}

func (o *NodeCountDefault) String() string {
	return fmt.Sprintf("[GET /Nodes/count][%d] Node.count default  %+v", o._statusCode, o.Payload)
}

func (o *NodeCountDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *NodeCountDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
NodeCountOKBody node count o k body
swagger:model NodeCountOKBody
*/
type NodeCountOKBody struct {

	// count
	Count float64 `json:"count,omitempty"`
}

// Validate validates this node count o k body
func (o *NodeCountOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this node count o k body based on context it is used
func (o *NodeCountOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *NodeCountOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NodeCountOKBody) UnmarshalBinary(b []byte) error {
	var res NodeCountOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
