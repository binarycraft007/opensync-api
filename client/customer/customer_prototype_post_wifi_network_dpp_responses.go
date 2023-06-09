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

// CustomerPrototypePostWifiNetworkDppReader is a Reader for the CustomerPrototypePostWifiNetworkDpp structure.
type CustomerPrototypePostWifiNetworkDppReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerPrototypePostWifiNetworkDppReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewCustomerPrototypePostWifiNetworkDppAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerPrototypePostWifiNetworkDppDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerPrototypePostWifiNetworkDppAccepted creates a CustomerPrototypePostWifiNetworkDppAccepted with default headers values
func NewCustomerPrototypePostWifiNetworkDppAccepted() *CustomerPrototypePostWifiNetworkDppAccepted {
	return &CustomerPrototypePostWifiNetworkDppAccepted{}
}

/*
CustomerPrototypePostWifiNetworkDppAccepted describes a response with status code 202, with default header values.

Request was successful
*/
type CustomerPrototypePostWifiNetworkDppAccepted struct {
	Payload interface{}
}

// IsSuccess returns true when this customer prototype post wifi network dpp accepted response has a 2xx status code
func (o *CustomerPrototypePostWifiNetworkDppAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer prototype post wifi network dpp accepted response has a 3xx status code
func (o *CustomerPrototypePostWifiNetworkDppAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer prototype post wifi network dpp accepted response has a 4xx status code
func (o *CustomerPrototypePostWifiNetworkDppAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer prototype post wifi network dpp accepted response has a 5xx status code
func (o *CustomerPrototypePostWifiNetworkDppAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this customer prototype post wifi network dpp accepted response a status code equal to that given
func (o *CustomerPrototypePostWifiNetworkDppAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the customer prototype post wifi network dpp accepted response
func (o *CustomerPrototypePostWifiNetworkDppAccepted) Code() int {
	return 202
}

func (o *CustomerPrototypePostWifiNetworkDppAccepted) Error() string {
	return fmt.Sprintf("[POST /Customers/{id}/locations/{locationId}/wifiNetwork/dpp][%d] customerPrototypePostWifiNetworkDppAccepted  %+v", 202, o.Payload)
}

func (o *CustomerPrototypePostWifiNetworkDppAccepted) String() string {
	return fmt.Sprintf("[POST /Customers/{id}/locations/{locationId}/wifiNetwork/dpp][%d] customerPrototypePostWifiNetworkDppAccepted  %+v", 202, o.Payload)
}

func (o *CustomerPrototypePostWifiNetworkDppAccepted) GetPayload() interface{} {
	return o.Payload
}

func (o *CustomerPrototypePostWifiNetworkDppAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerPrototypePostWifiNetworkDppDefault creates a CustomerPrototypePostWifiNetworkDppDefault with default headers values
func NewCustomerPrototypePostWifiNetworkDppDefault(code int) *CustomerPrototypePostWifiNetworkDppDefault {
	return &CustomerPrototypePostWifiNetworkDppDefault{
		_statusCode: code,
	}
}

/*
CustomerPrototypePostWifiNetworkDppDefault describes a response with status code -1, with default header values.

unexpected error
*/
type CustomerPrototypePostWifiNetworkDppDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this customer prototype post wifi network dpp default response has a 2xx status code
func (o *CustomerPrototypePostWifiNetworkDppDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer prototype post wifi network dpp default response has a 3xx status code
func (o *CustomerPrototypePostWifiNetworkDppDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer prototype post wifi network dpp default response has a 4xx status code
func (o *CustomerPrototypePostWifiNetworkDppDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer prototype post wifi network dpp default response has a 5xx status code
func (o *CustomerPrototypePostWifiNetworkDppDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer prototype post wifi network dpp default response a status code equal to that given
func (o *CustomerPrototypePostWifiNetworkDppDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer prototype post wifi network dpp default response
func (o *CustomerPrototypePostWifiNetworkDppDefault) Code() int {
	return o._statusCode
}

func (o *CustomerPrototypePostWifiNetworkDppDefault) Error() string {
	return fmt.Sprintf("[POST /Customers/{id}/locations/{locationId}/wifiNetwork/dpp][%d] Customer.prototype.postWifiNetworkDpp default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePostWifiNetworkDppDefault) String() string {
	return fmt.Sprintf("[POST /Customers/{id}/locations/{locationId}/wifiNetwork/dpp][%d] Customer.prototype.postWifiNetworkDpp default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerPrototypePostWifiNetworkDppDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CustomerPrototypePostWifiNetworkDppDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
