// Code generated by go-swagger; DO NOT EDIT.

package customer

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

// CustomerPrototypePutLedModeReader is a Reader for the CustomerPrototypePutLedMode structure.
type CustomerPrototypePutLedModeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypePutLedModeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypePutLedModeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypePutLedModeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypePutLedModeOK creates a CustomerPrototypePutLedModeOK with default headers values
func NewCustomerPrototypePutLedModeOK() *CustomerPrototypePutLedModeOK {
	return &CustomerPrototypePutLedModeOK{}
}

/*
CustomerPrototypePutLedModeOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypePutLedModeOK struct {
	Payload *CustomerPrototypePutLedModeOKBody
}

// IsSuccess returns true when this customer prototype put led mode o k response has a 2xx status code
func (o *CustomerPrototypePutLedModeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype put led mode o k response has a 3xx status code
func (o *CustomerPrototypePutLedModeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype put led mode o k response has a 4xx status code
func (o *CustomerPrototypePutLedModeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype put led mode o k response has a 5xx status code
func (o *CustomerPrototypePutLedModeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype put led mode o k response a status code equal to that given
func (o *CustomerPrototypePutLedModeOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype put led mode o k response
func (o *CustomerPrototypePutLedModeOK) Code() int {
	return 200
}

func (o *CustomerPrototypePutLedModeOK) Error() string {
	return fmt.Sprintf("[PUT /Customers/{id}/locations/{locationId}/nodes/{nodeId}/ledMode][%d] customerPrototypePutLedModeOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypePutLedModeOK) String() string {
	return fmt.Sprintf("[PUT /Customers/{id}/locations/{locationId}/nodes/{nodeId}/ledMode][%d] customerPrototypePutLedModeOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypePutLedModeOK) GetPayload() *CustomerPrototypePutLedModeOKBody {
	return o.Payload
}

func (o *CustomerPrototypePutLedModeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CustomerPrototypePutLedModeOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypePutLedModeDefault creates a CustomerPrototypePutLedModeDefault with default headers values
func NewCustomerPrototypePutLedModeDefault(code int) *CustomerPrototypePutLedModeDefault {
	return &CustomerPrototypePutLedModeDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypePutLedModeDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypePutLedModeDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype put led mode default response has a 2xx status code
func (o *CustomerPrototypePutLedModeDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype put led mode default response has a 3xx status code
func (o *CustomerPrototypePutLedModeDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype put led mode default response has a 4xx status code
func (o *CustomerPrototypePutLedModeDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype put led mode default response has a 5xx status code
func (o *CustomerPrototypePutLedModeDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype put led mode default response a status code equal to that given
func (o *CustomerPrototypePutLedModeDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype put led mode default response
func (o *CustomerPrototypePutLedModeDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypePutLedModeDefault) Error() string {
	return fmt.Sprintf("[PUT /Customers/{id}/locations/{locationId}/nodes/{nodeId}/ledMode][%d] Customer.prototype.putLedMode default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePutLedModeDefault) String() string {
	return fmt.Sprintf("[PUT /Customers/{id}/locations/{locationId}/nodes/{nodeId}/ledMode][%d] Customer.prototype.putLedMode default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePutLedModeDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypePutLedModeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
CustomerPrototypePutLedModeOKBody customer prototype put led mode o k body
swagger:model CustomerPrototypePutLedModeOKBody
*/
type CustomerPrototypePutLedModeOKBody struct {

	// mode
	Mode string `json:"mode,omitempty"`
}

// Validate validates this customer prototype put led mode o k body
func (o *CustomerPrototypePutLedModeOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this customer prototype put led mode o k body based on context it is used
func (o *CustomerPrototypePutLedModeOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CustomerPrototypePutLedModeOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CustomerPrototypePutLedModeOKBody) UnmarshalBinary(b []byte) error {
	var res CustomerPrototypePutLedModeOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}