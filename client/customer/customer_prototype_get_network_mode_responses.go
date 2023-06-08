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

// CustomerPrototypeGetNetworkModeReader is a Reader for the CustomerPrototypeGetNetworkMode structure.
type CustomerPrototypeGetNetworkModeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypeGetNetworkModeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypeGetNetworkModeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypeGetNetworkModeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypeGetNetworkModeOK creates a CustomerPrototypeGetNetworkModeOK with default headers values
func NewCustomerPrototypeGetNetworkModeOK() *CustomerPrototypeGetNetworkModeOK {
	return &CustomerPrototypeGetNetworkModeOK{}
}

/*
CustomerPrototypeGetNetworkModeOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypeGetNetworkModeOK struct {
	Payload *CustomerPrototypeGetNetworkModeOKBody
}

// IsSuccess returns true when this customer prototype get network mode o k response has a 2xx status code
func (o *CustomerPrototypeGetNetworkModeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype get network mode o k response has a 3xx status code
func (o *CustomerPrototypeGetNetworkModeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype get network mode o k response has a 4xx status code
func (o *CustomerPrototypeGetNetworkModeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype get network mode o k response has a 5xx status code
func (o *CustomerPrototypeGetNetworkModeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype get network mode o k response a status code equal to that given
func (o *CustomerPrototypeGetNetworkModeOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype get network mode o k response
func (o *CustomerPrototypeGetNetworkModeOK) Code() int {
	return 200
}

func (o *CustomerPrototypeGetNetworkModeOK) Error() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/networkMode][%d] customerPrototypeGetNetworkModeOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeGetNetworkModeOK) String() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/networkMode][%d] customerPrototypeGetNetworkModeOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeGetNetworkModeOK) GetPayload() *CustomerPrototypeGetNetworkModeOKBody {
	return o.Payload
}

func (o *CustomerPrototypeGetNetworkModeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CustomerPrototypeGetNetworkModeOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypeGetNetworkModeDefault creates a CustomerPrototypeGetNetworkModeDefault with default headers values
func NewCustomerPrototypeGetNetworkModeDefault(code int) *CustomerPrototypeGetNetworkModeDefault {
	return &CustomerPrototypeGetNetworkModeDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypeGetNetworkModeDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypeGetNetworkModeDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype get network mode default response has a 2xx status code
func (o *CustomerPrototypeGetNetworkModeDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype get network mode default response has a 3xx status code
func (o *CustomerPrototypeGetNetworkModeDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype get network mode default response has a 4xx status code
func (o *CustomerPrototypeGetNetworkModeDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype get network mode default response has a 5xx status code
func (o *CustomerPrototypeGetNetworkModeDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype get network mode default response a status code equal to that given
func (o *CustomerPrototypeGetNetworkModeDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype get network mode default response
func (o *CustomerPrototypeGetNetworkModeDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypeGetNetworkModeDefault) Error() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/networkMode][%d] Customer.prototype.getNetworkMode default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeGetNetworkModeDefault) String() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/networkMode][%d] Customer.prototype.getNetworkMode default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeGetNetworkModeDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypeGetNetworkModeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
CustomerPrototypeGetNetworkModeOKBody customer prototype get network mode o k body
swagger:model CustomerPrototypeGetNetworkModeOKBody
*/
type CustomerPrototypeGetNetworkModeOKBody struct {

	// network mode
	NetworkMode string `json:"networkMode,omitempty"`
}

// Validate validates this customer prototype get network mode o k body
func (o *CustomerPrototypeGetNetworkModeOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this customer prototype get network mode o k body based on context it is used
func (o *CustomerPrototypeGetNetworkModeOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CustomerPrototypeGetNetworkModeOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CustomerPrototypeGetNetworkModeOKBody) UnmarshalBinary(b []byte) error {
	var res CustomerPrototypeGetNetworkModeOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}