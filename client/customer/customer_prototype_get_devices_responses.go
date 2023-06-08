// Code generated by go-swagger; DO NOT EDIT.

package customer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/binarycraft007/opensync-api/models"
)

// CustomerPrototypeGetDevicesReader is a Reader for the CustomerPrototypeGetDevices structure.
type CustomerPrototypeGetDevicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypeGetDevicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypeGetDevicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypeGetDevicesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypeGetDevicesOK creates a CustomerPrototypeGetDevicesOK with default headers values
func NewCustomerPrototypeGetDevicesOK() *CustomerPrototypeGetDevicesOK {
	return &CustomerPrototypeGetDevicesOK{}
}

/*
CustomerPrototypeGetDevicesOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypeGetDevicesOK struct {
	Payload *CustomerPrototypeGetDevicesOKBody
}

// IsSuccess returns true when this customer prototype get devices o k response has a 2xx status code
func (o *CustomerPrototypeGetDevicesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype get devices o k response has a 3xx status code
func (o *CustomerPrototypeGetDevicesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype get devices o k response has a 4xx status code
func (o *CustomerPrototypeGetDevicesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype get devices o k response has a 5xx status code
func (o *CustomerPrototypeGetDevicesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype get devices o k response a status code equal to that given
func (o *CustomerPrototypeGetDevicesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype get devices o k response
func (o *CustomerPrototypeGetDevicesOK) Code() int {
	return 200
}

func (o *CustomerPrototypeGetDevicesOK) Error() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/devices][%d] customerPrototypeGetDevicesOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeGetDevicesOK) String() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/devices][%d] customerPrototypeGetDevicesOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeGetDevicesOK) GetPayload() *CustomerPrototypeGetDevicesOKBody {
	return o.Payload
}

func (o *CustomerPrototypeGetDevicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CustomerPrototypeGetDevicesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypeGetDevicesDefault creates a CustomerPrototypeGetDevicesDefault with default headers values
func NewCustomerPrototypeGetDevicesDefault(code int) *CustomerPrototypeGetDevicesDefault {
	return &CustomerPrototypeGetDevicesDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypeGetDevicesDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypeGetDevicesDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype get devices default response has a 2xx status code
func (o *CustomerPrototypeGetDevicesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype get devices default response has a 3xx status code
func (o *CustomerPrototypeGetDevicesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype get devices default response has a 4xx status code
func (o *CustomerPrototypeGetDevicesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype get devices default response has a 5xx status code
func (o *CustomerPrototypeGetDevicesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype get devices default response a status code equal to that given
func (o *CustomerPrototypeGetDevicesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype get devices default response
func (o *CustomerPrototypeGetDevicesDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypeGetDevicesDefault) Error() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/devices][%d] Customer.prototype.getDevices default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeGetDevicesDefault) String() string {
	return fmt.Sprintf("[GET /Customers/{id}/locations/{locationId}/devices][%d] Customer.prototype.getDevices default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeGetDevicesDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypeGetDevicesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
CustomerPrototypeGetDevicesOKBody customer prototype get devices o k body
swagger:model CustomerPrototypeGetDevicesOKBody
*/
type CustomerPrototypeGetDevicesOKBody struct {

	// devices
	Devices *models.Devices `json:"devices,omitempty"`
}

// Validate validates this customer prototype get devices o k body
func (o *CustomerPrototypeGetDevicesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDevices(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CustomerPrototypeGetDevicesOKBody) validateDevices(formats strfmt.Registry) error {
	if swag.IsZero(o.Devices) { // not required
		return nil
	}

	if o.Devices != nil {
		if err := o.Devices.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("customerPrototypeGetDevicesOK" + "." + "devices")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("customerPrototypeGetDevicesOK" + "." + "devices")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this customer prototype get devices o k body based on the context it is used
func (o *CustomerPrototypeGetDevicesOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDevices(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CustomerPrototypeGetDevicesOKBody) contextValidateDevices(ctx context.Context, formats strfmt.Registry) error {

	if o.Devices != nil {
		if err := o.Devices.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("customerPrototypeGetDevicesOK" + "." + "devices")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("customerPrototypeGetDevicesOK" + "." + "devices")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CustomerPrototypeGetDevicesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CustomerPrototypeGetDevicesOKBody) UnmarshalBinary(b []byte) error {
	var res CustomerPrototypeGetDevicesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
