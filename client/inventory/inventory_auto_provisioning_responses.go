// Code generated by go-swagger; DO NOT EDIT.

package inventory

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/binarycraft007/opensync-api/models"
)

// InventoryAutoProvisioningReader is a Reader for the InventoryAutoProvisioning structure.
type InventoryAutoProvisioningReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InventoryAutoProvisioningReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInventoryAutoProvisioningOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewInventoryAutoProvisioningDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewInventoryAutoProvisioningOK creates a InventoryAutoProvisioningOK with default headers values
func NewInventoryAutoProvisioningOK() *InventoryAutoProvisioningOK {
	return &InventoryAutoProvisioningOK{}
}

/*
InventoryAutoProvisioningOK describes a response with status code 200, with default header values.

Request was successful
*/
type InventoryAutoProvisioningOK struct {
	Payload interface{}
}

// IsSuccess returns true when this inventory auto provisioning o k response has a 2xx status code
func (o *InventoryAutoProvisioningOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this inventory auto provisioning o k response has a 3xx status code
func (o *InventoryAutoProvisioningOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this inventory auto provisioning o k response has a 4xx status code
func (o *InventoryAutoProvisioningOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this inventory auto provisioning o k response has a 5xx status code
func (o *InventoryAutoProvisioningOK) IsServerError() bool {
	return false
}

// IsCode returns true when this inventory auto provisioning o k response a status code equal to that given
func (o *InventoryAutoProvisioningOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the inventory auto provisioning o k response
func (o *InventoryAutoProvisioningOK) Code() int {
	return 200
}

func (o *InventoryAutoProvisioningOK) Error() string {
	return fmt.Sprintf("[PATCH /inventory/nodes/{nodeId}/autoProvisioning][%d] inventoryAutoProvisioningOK  %+v", 200, o.Payload)
}

func (o *InventoryAutoProvisioningOK) String() string {
	return fmt.Sprintf("[PATCH /inventory/nodes/{nodeId}/autoProvisioning][%d] inventoryAutoProvisioningOK  %+v", 200, o.Payload)
}

func (o *InventoryAutoProvisioningOK) GetPayload() interface{} {
	return o.Payload
}

func (o *InventoryAutoProvisioningOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInventoryAutoProvisioningDefault creates a InventoryAutoProvisioningDefault with default headers values
func NewInventoryAutoProvisioningDefault(code int) *InventoryAutoProvisioningDefault {
	return &InventoryAutoProvisioningDefault{
		_statusCode: code,
	}
}

/*
InventoryAutoProvisioningDefault describes a response with status code -1, with default header values.

unexpected error
*/
type InventoryAutoProvisioningDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// IsSuccess returns true when this inventory auto provisioning default response has a 2xx status code
func (o *InventoryAutoProvisioningDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this inventory auto provisioning default response has a 3xx status code
func (o *InventoryAutoProvisioningDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this inventory auto provisioning default response has a 4xx status code
func (o *InventoryAutoProvisioningDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this inventory auto provisioning default response has a 5xx status code
func (o *InventoryAutoProvisioningDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this inventory auto provisioning default response a status code equal to that given
func (o *InventoryAutoProvisioningDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the inventory auto provisioning default response
func (o *InventoryAutoProvisioningDefault) Code() int {
	return o._statusCode
}

func (o *InventoryAutoProvisioningDefault) Error() string {
	return fmt.Sprintf("[PATCH /inventory/nodes/{nodeId}/autoProvisioning][%d] Inventory.autoProvisioning default  %+v", o._statusCode, o.Payload)
}

func (o *InventoryAutoProvisioningDefault) String() string {
	return fmt.Sprintf("[PATCH /inventory/nodes/{nodeId}/autoProvisioning][%d] Inventory.autoProvisioning default  %+v", o._statusCode, o.Payload)
}

func (o *InventoryAutoProvisioningDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *InventoryAutoProvisioningDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
