// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Node Plume internal-only APIs.
//
// swagger:model Node
type Node struct {

	// schema version # of a Mongo document
	// Required: true
	Version *string `json:"_version"`

	// backhaul dhcp pool idx
	BackhaulDhcpPoolIdx float64 `json:"backhaulDhcpPoolIdx,omitempty"`

	// blacklists
	Blacklists []XAny `json:"blacklists"`

	// bluetooth mac
	BluetoothMac string `json:"bluetoothMac,omitempty"`

	// box serial number
	BoxSerialNumber string `json:"boxSerialNumber,omitempty"`

	// certificates
	Certificates *NodeCertificates `json:"certificates,omitempty"`

	// connection state
	ConnectionState string `json:"connectionState,omitempty"`

	// time at which connectionStateChange last changed
	// Format: date-time
	ConnectionStateChangeAt strfmt.DateTime `json:"connectionStateChangeAt,omitempty"`

	// country code
	CountryCode string `json:"countryCode,omitempty"`

	// customer Id
	CustomerID ObjectID `json:"customerId,omitempty"`

	// used by redirector to route pods
	Deployment string `json:"deployment,omitempty"`

	// gives the ability to enable or disable individual pods ethernetLan
	EthernetLan interface{} `json:"ethernetLan,omitempty"`

	// ethernet mac
	EthernetMac string `json:"ethernetMac,omitempty"`

	// firmware version
	FirmwareVersion string `json:"firmwareVersion,omitempty"`

	// always returns the value of 'serialNumber'
	ID string `json:"id,omitempty"`

	// kv configs
	KvConfigs []*KvConfig `json:"kvConfigs"`

	// location Id
	LocationID ObjectID `json:"locationId,omitempty"`

	// macs
	Macs []XAny `json:"macs"`

	// 1 = Arcadyan Kunshan
	ManufactureLocation string `json:"manufactureLocation,omitempty"`

	// 48 (week 48), 01,02,03,- 52
	ManufactureWeek string `json:"manufactureWeek,omitempty"`

	// 6 (2016), 7 (2017)
	ManufactureYear string `json:"manufactureYear,omitempty"`

	// value from the firmware's OVS AWLAN_Node.serial_number column
	ManufacturerSerialNumber string `json:"manufacturerSerialNumber,omitempty"`

	// reported by OVS schema
	Model string `json:"model,omitempty"`

	// the default value is 'auto', users can choose 'router' or 'auto'
	NetworkMode string `json:"networkMode,omitempty"`

	// nickname
	Nickname string `json:"nickname,omitempty"`

	// pack Id
	PackID string `json:"packId,omitempty"`

	// platform version
	PlatformVersion string `json:"platformVersion,omitempty"`

	// 1 = Pod, 2x2 DBDC, 1xGbE Ethernet, BLE, All Plug Types
	ProductDescriptor string `json:"productDescriptor,omitempty"`

	// Manufacturer Product Family (Alpha)
	ProductFamily string `json:"productFamily,omitempty"`

	// radio mac24
	RadioMac24 string `json:"radioMac24,omitempty"`

	// radio mac50
	RadioMac50 string `json:"radioMac50,omitempty"`

	// radio mac60
	RadioMac60 string `json:"radioMac60,omitempty"`

	// residential gateways (mostly auto imported) require special treatment and searchability
	ResidentialGateway *bool `json:"residentialGateway,omitempty"`

	// serial number
	// Required: true
	SerialNumber *string `json:"serialNumber"`

	// used by redirector to route pods
	ShardNumber string `json:"shardNumber,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// if true, Plume mobile app customers will need to activate a subscription to use our service
	SubscriptionRequired *bool `json:"subscriptionRequired,omitempty"`

	// synchronized key
	SynchronizedKey string `json:"synchronizedKey,omitempty"`

	// thread
	Thread string `json:"thread,omitempty"`

	// date/time node was claimed
	// Format: date-time
	Timestamp strfmt.DateTime `json:"timestamp,omitempty"`

	// updated at
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updatedAt,omitempty"`

	// vendor
	Vendor *NodeVendor `json:"vendor,omitempty"`
}

// Validate validates this node
func (m *Node) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCertificates(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnectionStateChangeAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomerID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKvConfigs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSerialNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVendor(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Node) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("_version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

func (m *Node) validateCertificates(formats strfmt.Registry) error {
	if swag.IsZero(m.Certificates) { // not required
		return nil
	}

	if m.Certificates != nil {
		if err := m.Certificates.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("certificates")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("certificates")
			}
			return err
		}
	}

	return nil
}

func (m *Node) validateConnectionStateChangeAt(formats strfmt.Registry) error {
	if swag.IsZero(m.ConnectionStateChangeAt) { // not required
		return nil
	}

	if err := validate.FormatOf("connectionStateChangeAt", "body", "date-time", m.ConnectionStateChangeAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Node) validateCustomerID(formats strfmt.Registry) error {
	if swag.IsZero(m.CustomerID) { // not required
		return nil
	}

	if err := m.CustomerID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("customerId")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("customerId")
		}
		return err
	}

	return nil
}

func (m *Node) validateKvConfigs(formats strfmt.Registry) error {
	if swag.IsZero(m.KvConfigs) { // not required
		return nil
	}

	for i := 0; i < len(m.KvConfigs); i++ {
		if swag.IsZero(m.KvConfigs[i]) { // not required
			continue
		}

		if m.KvConfigs[i] != nil {
			if err := m.KvConfigs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("kvConfigs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("kvConfigs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Node) validateLocationID(formats strfmt.Registry) error {
	if swag.IsZero(m.LocationID) { // not required
		return nil
	}

	if err := m.LocationID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("locationId")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("locationId")
		}
		return err
	}

	return nil
}

func (m *Node) validateSerialNumber(formats strfmt.Registry) error {

	if err := validate.Required("serialNumber", "body", m.SerialNumber); err != nil {
		return err
	}

	return nil
}

func (m *Node) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Node) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Node) validateVendor(formats strfmt.Registry) error {
	if swag.IsZero(m.Vendor) { // not required
		return nil
	}

	if m.Vendor != nil {
		if err := m.Vendor.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vendor")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vendor")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this node based on the context it is used
func (m *Node) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCertificates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCustomerID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateKvConfigs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLocationID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVendor(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Node) contextValidateCertificates(ctx context.Context, formats strfmt.Registry) error {

	if m.Certificates != nil {
		if err := m.Certificates.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("certificates")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("certificates")
			}
			return err
		}
	}

	return nil
}

func (m *Node) contextValidateCustomerID(ctx context.Context, formats strfmt.Registry) error {

	if err := m.CustomerID.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("customerId")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("customerId")
		}
		return err
	}

	return nil
}

func (m *Node) contextValidateKvConfigs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.KvConfigs); i++ {

		if m.KvConfigs[i] != nil {
			if err := m.KvConfigs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("kvConfigs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("kvConfigs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Node) contextValidateLocationID(ctx context.Context, formats strfmt.Registry) error {

	if err := m.LocationID.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("locationId")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("locationId")
		}
		return err
	}

	return nil
}

func (m *Node) contextValidateVendor(ctx context.Context, formats strfmt.Registry) error {

	if m.Vendor != nil {
		if err := m.Vendor.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vendor")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vendor")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Node) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Node) UnmarshalBinary(b []byte) error {
	var res Node
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
