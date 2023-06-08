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

// CustomerPrototypePostOnboardingCheckpointReader is a Reader for the CustomerPrototypePostOnboardingCheckpoint structure.
type CustomerPrototypePostOnboardingCheckpointReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypePostOnboardingCheckpointReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypePostOnboardingCheckpointOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypePostOnboardingCheckpointDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypePostOnboardingCheckpointOK creates a CustomerPrototypePostOnboardingCheckpointOK with default headers values
func NewCustomerPrototypePostOnboardingCheckpointOK() *CustomerPrototypePostOnboardingCheckpointOK {
	return &CustomerPrototypePostOnboardingCheckpointOK{}
}

/*
CustomerPrototypePostOnboardingCheckpointOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypePostOnboardingCheckpointOK struct {
	Payload *CustomerPrototypePostOnboardingCheckpointOKBody
}

// IsSuccess returns true when this customer prototype post onboarding checkpoint o k response has a 2xx status code
func (o *CustomerPrototypePostOnboardingCheckpointOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype post onboarding checkpoint o k response has a 3xx status code
func (o *CustomerPrototypePostOnboardingCheckpointOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype post onboarding checkpoint o k response has a 4xx status code
func (o *CustomerPrototypePostOnboardingCheckpointOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype post onboarding checkpoint o k response has a 5xx status code
func (o *CustomerPrototypePostOnboardingCheckpointOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype post onboarding checkpoint o k response a status code equal to that given
func (o *CustomerPrototypePostOnboardingCheckpointOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype post onboarding checkpoint o k response
func (o *CustomerPrototypePostOnboardingCheckpointOK) Code() int {
	return 200
}

func (o *CustomerPrototypePostOnboardingCheckpointOK) Error() string {
	return fmt.Sprintf("[POST /Customers/{id}/locations/{locationId}/onboardingCheckpoint][%d] customerPrototypePostOnboardingCheckpointOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypePostOnboardingCheckpointOK) String() string {
	return fmt.Sprintf("[POST /Customers/{id}/locations/{locationId}/onboardingCheckpoint][%d] customerPrototypePostOnboardingCheckpointOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypePostOnboardingCheckpointOK) GetPayload() *CustomerPrototypePostOnboardingCheckpointOKBody {
	return o.Payload
}

func (o *CustomerPrototypePostOnboardingCheckpointOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CustomerPrototypePostOnboardingCheckpointOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypePostOnboardingCheckpointDefault creates a CustomerPrototypePostOnboardingCheckpointDefault with default headers values
func NewCustomerPrototypePostOnboardingCheckpointDefault(code int) *CustomerPrototypePostOnboardingCheckpointDefault {
	return &CustomerPrototypePostOnboardingCheckpointDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypePostOnboardingCheckpointDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypePostOnboardingCheckpointDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype post onboarding checkpoint default response has a 2xx status code
func (o *CustomerPrototypePostOnboardingCheckpointDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype post onboarding checkpoint default response has a 3xx status code
func (o *CustomerPrototypePostOnboardingCheckpointDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype post onboarding checkpoint default response has a 4xx status code
func (o *CustomerPrototypePostOnboardingCheckpointDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype post onboarding checkpoint default response has a 5xx status code
func (o *CustomerPrototypePostOnboardingCheckpointDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype post onboarding checkpoint default response a status code equal to that given
func (o *CustomerPrototypePostOnboardingCheckpointDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype post onboarding checkpoint default response
func (o *CustomerPrototypePostOnboardingCheckpointDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypePostOnboardingCheckpointDefault) Error() string {
	return fmt.Sprintf("[POST /Customers/{id}/locations/{locationId}/onboardingCheckpoint][%d] Customer.prototype.postOnboardingCheckpoint default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePostOnboardingCheckpointDefault) String() string {
	return fmt.Sprintf("[POST /Customers/{id}/locations/{locationId}/onboardingCheckpoint][%d] Customer.prototype.postOnboardingCheckpoint default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePostOnboardingCheckpointDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypePostOnboardingCheckpointDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
CustomerPrototypePostOnboardingCheckpointOKBody customer prototype post onboarding checkpoint o k body
swagger:model CustomerPrototypePostOnboardingCheckpointOKBody
*/
type CustomerPrototypePostOnboardingCheckpointOKBody struct {

	// checkpoint
	Checkpoint *models.OnboardingCheckpointResponse `json:"checkpoint,omitempty"`
}

// Validate validates this customer prototype post onboarding checkpoint o k body
func (o *CustomerPrototypePostOnboardingCheckpointOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCheckpoint(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CustomerPrototypePostOnboardingCheckpointOKBody) validateCheckpoint(formats strfmt.Registry) error {
	if swag.IsZero(o.Checkpoint) { // not required
		return nil
	}

	if o.Checkpoint != nil {
		if err := o.Checkpoint.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("customerPrototypePostOnboardingCheckpointOK" + "." + "checkpoint")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("customerPrototypePostOnboardingCheckpointOK" + "." + "checkpoint")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this customer prototype post onboarding checkpoint o k body based on the context it is used
func (o *CustomerPrototypePostOnboardingCheckpointOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateCheckpoint(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CustomerPrototypePostOnboardingCheckpointOKBody) contextValidateCheckpoint(ctx context.Context, formats strfmt.Registry) error {

	if o.Checkpoint != nil {
		if err := o.Checkpoint.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("customerPrototypePostOnboardingCheckpointOK" + "." + "checkpoint")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("customerPrototypePostOnboardingCheckpointOK" + "." + "checkpoint")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CustomerPrototypePostOnboardingCheckpointOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CustomerPrototypePostOnboardingCheckpointOKBody) UnmarshalBinary(b []byte) error {
	var res CustomerPrototypePostOnboardingCheckpointOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
