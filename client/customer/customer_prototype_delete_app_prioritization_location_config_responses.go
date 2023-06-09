// Code generated by go-swagger; DO NOT EDIT.

package customer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/binarycraft007/opensync-api/models"
)

// CustomerPrototypeDeleteAppPrioritizationLocationConfigReader is a Reader for the CustomerPrototypeDeleteAppPrioritizationLocationConfig structure.
type CustomerPrototypeDeleteAppPrioritizationLocationConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypeDeleteAppPrioritizationLocationConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerPrototypeDeleteAppPrioritizationLocationConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypeDeleteAppPrioritizationLocationConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypeDeleteAppPrioritizationLocationConfigOK creates a CustomerPrototypeDeleteAppPrioritizationLocationConfigOK with default headers values
func NewCustomerPrototypeDeleteAppPrioritizationLocationConfigOK() *CustomerPrototypeDeleteAppPrioritizationLocationConfigOK {
	return &CustomerPrototypeDeleteAppPrioritizationLocationConfigOK{}
}

/*
CustomerPrototypeDeleteAppPrioritizationLocationConfigOK describes a response with status code 200, with default header values.

Request was successful
*/
type CustomerPrototypeDeleteAppPrioritizationLocationConfigOK struct {
	Payload interface{}
}

// IsSuccess returns true when this customer prototype delete app prioritization location config o k response has a 2xx status code
func (o *CustomerPrototypeDeleteAppPrioritizationLocationConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype delete app prioritization location config o k response has a 3xx status code
func (o *CustomerPrototypeDeleteAppPrioritizationLocationConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype delete app prioritization location config o k response has a 4xx status code
func (o *CustomerPrototypeDeleteAppPrioritizationLocationConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype delete app prioritization location config o k response has a 5xx status code
func (o *CustomerPrototypeDeleteAppPrioritizationLocationConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype delete app prioritization location config o k response a status code equal to that given
func (o *CustomerPrototypeDeleteAppPrioritizationLocationConfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer prototype delete app prioritization location config o k response
func (o *CustomerPrototypeDeleteAppPrioritizationLocationConfigOK) Code() int {
	return 200
}

func (o *CustomerPrototypeDeleteAppPrioritizationLocationConfigOK) Error() string {
	return fmt.Sprintf("[DELETE /Customers/{id}/locations/{locationId}/qos/appPrioritization/customSetting][%d] customerPrototypeDeleteAppPrioritizationLocationConfigOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeDeleteAppPrioritizationLocationConfigOK) String() string {
	return fmt.Sprintf("[DELETE /Customers/{id}/locations/{locationId}/qos/appPrioritization/customSetting][%d] customerPrototypeDeleteAppPrioritizationLocationConfigOK  %+v", 200, o.Payload)
}

func (o *CustomerPrototypeDeleteAppPrioritizationLocationConfigOK) GetPayload() interface{} {
	return o.Payload
}

func (o *CustomerPrototypeDeleteAppPrioritizationLocationConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypeDeleteAppPrioritizationLocationConfigDefault creates a CustomerPrototypeDeleteAppPrioritizationLocationConfigDefault with default headers values
func NewCustomerPrototypeDeleteAppPrioritizationLocationConfigDefault(code int) *CustomerPrototypeDeleteAppPrioritizationLocationConfigDefault {
	return &CustomerPrototypeDeleteAppPrioritizationLocationConfigDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypeDeleteAppPrioritizationLocationConfigDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypeDeleteAppPrioritizationLocationConfigDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype delete app prioritization location config default response has a 2xx status code
func (o *CustomerPrototypeDeleteAppPrioritizationLocationConfigDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype delete app prioritization location config default response has a 3xx status code
func (o *CustomerPrototypeDeleteAppPrioritizationLocationConfigDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype delete app prioritization location config default response has a 4xx status code
func (o *CustomerPrototypeDeleteAppPrioritizationLocationConfigDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype delete app prioritization location config default response has a 5xx status code
func (o *CustomerPrototypeDeleteAppPrioritizationLocationConfigDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype delete app prioritization location config default response a status code equal to that given
func (o *CustomerPrototypeDeleteAppPrioritizationLocationConfigDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype delete app prioritization location config default response
func (o *CustomerPrototypeDeleteAppPrioritizationLocationConfigDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypeDeleteAppPrioritizationLocationConfigDefault) Error() string {
	return fmt.Sprintf("[DELETE /Customers/{id}/locations/{locationId}/qos/appPrioritization/customSetting][%d] Customer.prototype.deleteAppPrioritizationLocationConfig default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeDeleteAppPrioritizationLocationConfigDefault) String() string {
	return fmt.Sprintf("[DELETE /Customers/{id}/locations/{locationId}/qos/appPrioritization/customSetting][%d] Customer.prototype.deleteAppPrioritizationLocationConfig default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypeDeleteAppPrioritizationLocationConfigDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypeDeleteAppPrioritizationLocationConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
