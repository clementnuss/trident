// Code generated by go-swagger; DO NOT EDIT.

package networking

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// SwitchModifyCollectionReader is a Reader for the SwitchModifyCollection structure.
type SwitchModifyCollectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SwitchModifyCollectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSwitchModifyCollectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewSwitchModifyCollectionAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSwitchModifyCollectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSwitchModifyCollectionOK creates a SwitchModifyCollectionOK with default headers values
func NewSwitchModifyCollectionOK() *SwitchModifyCollectionOK {
	return &SwitchModifyCollectionOK{}
}

/*
SwitchModifyCollectionOK describes a response with status code 200, with default header values.

OK
*/
type SwitchModifyCollectionOK struct {
	Payload *models.SwitchJobLinkResponse
}

// IsSuccess returns true when this switch modify collection o k response has a 2xx status code
func (o *SwitchModifyCollectionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this switch modify collection o k response has a 3xx status code
func (o *SwitchModifyCollectionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this switch modify collection o k response has a 4xx status code
func (o *SwitchModifyCollectionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this switch modify collection o k response has a 5xx status code
func (o *SwitchModifyCollectionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this switch modify collection o k response a status code equal to that given
func (o *SwitchModifyCollectionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the switch modify collection o k response
func (o *SwitchModifyCollectionOK) Code() int {
	return 200
}

func (o *SwitchModifyCollectionOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /network/ethernet/switches][%d] switchModifyCollectionOK %s", 200, payload)
}

func (o *SwitchModifyCollectionOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /network/ethernet/switches][%d] switchModifyCollectionOK %s", 200, payload)
}

func (o *SwitchModifyCollectionOK) GetPayload() *models.SwitchJobLinkResponse {
	return o.Payload
}

func (o *SwitchModifyCollectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SwitchJobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSwitchModifyCollectionAccepted creates a SwitchModifyCollectionAccepted with default headers values
func NewSwitchModifyCollectionAccepted() *SwitchModifyCollectionAccepted {
	return &SwitchModifyCollectionAccepted{}
}

/*
SwitchModifyCollectionAccepted describes a response with status code 202, with default header values.

Accepted
*/
type SwitchModifyCollectionAccepted struct {
	Payload *models.SwitchJobLinkResponse
}

// IsSuccess returns true when this switch modify collection accepted response has a 2xx status code
func (o *SwitchModifyCollectionAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this switch modify collection accepted response has a 3xx status code
func (o *SwitchModifyCollectionAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this switch modify collection accepted response has a 4xx status code
func (o *SwitchModifyCollectionAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this switch modify collection accepted response has a 5xx status code
func (o *SwitchModifyCollectionAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this switch modify collection accepted response a status code equal to that given
func (o *SwitchModifyCollectionAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the switch modify collection accepted response
func (o *SwitchModifyCollectionAccepted) Code() int {
	return 202
}

func (o *SwitchModifyCollectionAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /network/ethernet/switches][%d] switchModifyCollectionAccepted %s", 202, payload)
}

func (o *SwitchModifyCollectionAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /network/ethernet/switches][%d] switchModifyCollectionAccepted %s", 202, payload)
}

func (o *SwitchModifyCollectionAccepted) GetPayload() *models.SwitchJobLinkResponse {
	return o.Payload
}

func (o *SwitchModifyCollectionAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SwitchJobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSwitchModifyCollectionDefault creates a SwitchModifyCollectionDefault with default headers values
func NewSwitchModifyCollectionDefault(code int) *SwitchModifyCollectionDefault {
	return &SwitchModifyCollectionDefault{
		_statusCode: code,
	}
}

/*
	SwitchModifyCollectionDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 12517379 | Settings updated, but the IP address \"{address}\" is not reachable. Verify that the address is valid or check the network path. |
| 12517381 | Settings updated, but the SNMP validation request timed out. Verify that the \"snmp.user\" parameter is valid. |
| 12517383 | Settings updated, but the SNMP validation request timed out. Verify that the \"snmp.user\" parameter is valid (i.e., the SNMPv3 user exists in ONTAP and on the remote switch). If the \"snmp.user\" parameter is valid, verify that the SNMPv3 user's credentials are the same both in ONTAP as well as in the remote switch. If a custom engine-id was provided for the SNMPv3 user, ensure it is the same as that of the remote switch. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type SwitchModifyCollectionDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this switch modify collection default response has a 2xx status code
func (o *SwitchModifyCollectionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this switch modify collection default response has a 3xx status code
func (o *SwitchModifyCollectionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this switch modify collection default response has a 4xx status code
func (o *SwitchModifyCollectionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this switch modify collection default response has a 5xx status code
func (o *SwitchModifyCollectionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this switch modify collection default response a status code equal to that given
func (o *SwitchModifyCollectionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the switch modify collection default response
func (o *SwitchModifyCollectionDefault) Code() int {
	return o._statusCode
}

func (o *SwitchModifyCollectionDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /network/ethernet/switches][%d] switch_modify_collection default %s", o._statusCode, payload)
}

func (o *SwitchModifyCollectionDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /network/ethernet/switches][%d] switch_modify_collection default %s", o._statusCode, payload)
}

func (o *SwitchModifyCollectionDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SwitchModifyCollectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
SwitchModifyCollectionBody switch modify collection body
swagger:model SwitchModifyCollectionBody
*/
type SwitchModifyCollectionBody struct {

	// links
	Links *models.SelfLink `json:"_links,omitempty"`

	// IP Address.
	Address *string `json:"address,omitempty"`

	// Discovered By ONTAP CDP/LLDP
	// Read Only: true
	Discovered *bool `json:"discovered,omitempty"`

	// Model Number.
	Model *string `json:"model,omitempty"`

	// monitoring
	Monitoring *models.SwitchInlineMonitoring `json:"monitoring,omitempty"`

	// Name.
	Name *string `json:"name,omitempty"`

	// Switch Network.
	// Enum: ["cluster","storage"]
	Network *string `json:"network,omitempty"`

	// Serial Number.
	// Read Only: true
	SerialNumber *string `json:"serial_number,omitempty"`

	// snmp
	Snmp *models.SwitchInlineSnmp `json:"snmp,omitempty"`

	// switch response inline records
	SwitchResponseInlineRecords []*models.Switch `json:"records,omitempty"`

	// Software Version.
	// Read Only: true
	Version *string `json:"version,omitempty"`
}

// Validate validates this switch modify collection body
func (o *SwitchModifyCollectionBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMonitoring(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNetwork(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSnmp(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSwitchResponseInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SwitchModifyCollectionBody) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(o.Links) { // not required
		return nil
	}

	if o.Links != nil {
		if err := o.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

func (o *SwitchModifyCollectionBody) validateMonitoring(formats strfmt.Registry) error {
	if swag.IsZero(o.Monitoring) { // not required
		return nil
	}

	if o.Monitoring != nil {
		if err := o.Monitoring.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "monitoring")
			}
			return err
		}
	}

	return nil
}

var switchModifyCollectionBodyTypeNetworkPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["cluster","storage"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		switchModifyCollectionBodyTypeNetworkPropEnum = append(switchModifyCollectionBodyTypeNetworkPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// SwitchModifyCollectionBody
	// SwitchModifyCollectionBody
	// network
	// Network
	// cluster
	// END DEBUGGING
	// SwitchModifyCollectionBodyNetworkCluster captures enum value "cluster"
	SwitchModifyCollectionBodyNetworkCluster string = "cluster"

	// BEGIN DEBUGGING
	// SwitchModifyCollectionBody
	// SwitchModifyCollectionBody
	// network
	// Network
	// storage
	// END DEBUGGING
	// SwitchModifyCollectionBodyNetworkStorage captures enum value "storage"
	SwitchModifyCollectionBodyNetworkStorage string = "storage"
)

// prop value enum
func (o *SwitchModifyCollectionBody) validateNetworkEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, switchModifyCollectionBodyTypeNetworkPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *SwitchModifyCollectionBody) validateNetwork(formats strfmt.Registry) error {
	if swag.IsZero(o.Network) { // not required
		return nil
	}

	// value enum
	if err := o.validateNetworkEnum("info"+"."+"network", "body", *o.Network); err != nil {
		return err
	}

	return nil
}

func (o *SwitchModifyCollectionBody) validateSnmp(formats strfmt.Registry) error {
	if swag.IsZero(o.Snmp) { // not required
		return nil
	}

	if o.Snmp != nil {
		if err := o.Snmp.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "snmp")
			}
			return err
		}
	}

	return nil
}

func (o *SwitchModifyCollectionBody) validateSwitchResponseInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(o.SwitchResponseInlineRecords) { // not required
		return nil
	}

	for i := 0; i < len(o.SwitchResponseInlineRecords); i++ {
		if swag.IsZero(o.SwitchResponseInlineRecords[i]) { // not required
			continue
		}

		if o.SwitchResponseInlineRecords[i] != nil {
			if err := o.SwitchResponseInlineRecords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this switch modify collection body based on the context it is used
func (o *SwitchModifyCollectionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateDiscovered(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateMonitoring(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateSerialNumber(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateSnmp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateSwitchResponseInlineRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateVersion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SwitchModifyCollectionBody) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if o.Links != nil {
		if err := o.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

func (o *SwitchModifyCollectionBody) contextValidateDiscovered(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "info"+"."+"discovered", "body", o.Discovered); err != nil {
		return err
	}

	return nil
}

func (o *SwitchModifyCollectionBody) contextValidateMonitoring(ctx context.Context, formats strfmt.Registry) error {

	if o.Monitoring != nil {
		if err := o.Monitoring.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "monitoring")
			}
			return err
		}
	}

	return nil
}

func (o *SwitchModifyCollectionBody) contextValidateSerialNumber(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "info"+"."+"serial_number", "body", o.SerialNumber); err != nil {
		return err
	}

	return nil
}

func (o *SwitchModifyCollectionBody) contextValidateSnmp(ctx context.Context, formats strfmt.Registry) error {

	if o.Snmp != nil {
		if err := o.Snmp.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "snmp")
			}
			return err
		}
	}

	return nil
}

func (o *SwitchModifyCollectionBody) contextValidateSwitchResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.SwitchResponseInlineRecords); i++ {

		if o.SwitchResponseInlineRecords[i] != nil {
			if err := o.SwitchResponseInlineRecords[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *SwitchModifyCollectionBody) contextValidateVersion(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "info"+"."+"version", "body", o.Version); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SwitchModifyCollectionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SwitchModifyCollectionBody) UnmarshalBinary(b []byte) error {
	var res SwitchModifyCollectionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SwitchInlineMonitoring switch inline monitoring
swagger:model switch_inline_monitoring
*/
type SwitchInlineMonitoring struct {

	// Enable Health Monitoring.
	Enabled *bool `json:"enabled,omitempty"`

	// Is Monitored.
	// Read Only: true
	Monitored *bool `json:"monitored,omitempty"`

	// Reason For Not Monitoring.
	// Read Only: true
	// Enum: ["none","unsupported_model","user_deleted","bad_ip_address","invalid_snmp_settings","bad_model","invalid_software_version","user_disabled","unknown"]
	Reason *string `json:"reason,omitempty"`
}

// Validate validates this switch inline monitoring
func (o *SwitchInlineMonitoring) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateReason(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var switchInlineMonitoringTypeReasonPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["none","unsupported_model","user_deleted","bad_ip_address","invalid_snmp_settings","bad_model","invalid_software_version","user_disabled","unknown"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		switchInlineMonitoringTypeReasonPropEnum = append(switchInlineMonitoringTypeReasonPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// switch_inline_monitoring
	// SwitchInlineMonitoring
	// reason
	// Reason
	// none
	// END DEBUGGING
	// SwitchInlineMonitoringReasonNone captures enum value "none"
	SwitchInlineMonitoringReasonNone string = "none"

	// BEGIN DEBUGGING
	// switch_inline_monitoring
	// SwitchInlineMonitoring
	// reason
	// Reason
	// unsupported_model
	// END DEBUGGING
	// SwitchInlineMonitoringReasonUnsupportedModel captures enum value "unsupported_model"
	SwitchInlineMonitoringReasonUnsupportedModel string = "unsupported_model"

	// BEGIN DEBUGGING
	// switch_inline_monitoring
	// SwitchInlineMonitoring
	// reason
	// Reason
	// user_deleted
	// END DEBUGGING
	// SwitchInlineMonitoringReasonUserDeleted captures enum value "user_deleted"
	SwitchInlineMonitoringReasonUserDeleted string = "user_deleted"

	// BEGIN DEBUGGING
	// switch_inline_monitoring
	// SwitchInlineMonitoring
	// reason
	// Reason
	// bad_ip_address
	// END DEBUGGING
	// SwitchInlineMonitoringReasonBadIPAddress captures enum value "bad_ip_address"
	SwitchInlineMonitoringReasonBadIPAddress string = "bad_ip_address"

	// BEGIN DEBUGGING
	// switch_inline_monitoring
	// SwitchInlineMonitoring
	// reason
	// Reason
	// invalid_snmp_settings
	// END DEBUGGING
	// SwitchInlineMonitoringReasonInvalidSnmpSettings captures enum value "invalid_snmp_settings"
	SwitchInlineMonitoringReasonInvalidSnmpSettings string = "invalid_snmp_settings"

	// BEGIN DEBUGGING
	// switch_inline_monitoring
	// SwitchInlineMonitoring
	// reason
	// Reason
	// bad_model
	// END DEBUGGING
	// SwitchInlineMonitoringReasonBadModel captures enum value "bad_model"
	SwitchInlineMonitoringReasonBadModel string = "bad_model"

	// BEGIN DEBUGGING
	// switch_inline_monitoring
	// SwitchInlineMonitoring
	// reason
	// Reason
	// invalid_software_version
	// END DEBUGGING
	// SwitchInlineMonitoringReasonInvalidSoftwareVersion captures enum value "invalid_software_version"
	SwitchInlineMonitoringReasonInvalidSoftwareVersion string = "invalid_software_version"

	// BEGIN DEBUGGING
	// switch_inline_monitoring
	// SwitchInlineMonitoring
	// reason
	// Reason
	// user_disabled
	// END DEBUGGING
	// SwitchInlineMonitoringReasonUserDisabled captures enum value "user_disabled"
	SwitchInlineMonitoringReasonUserDisabled string = "user_disabled"

	// BEGIN DEBUGGING
	// switch_inline_monitoring
	// SwitchInlineMonitoring
	// reason
	// Reason
	// unknown
	// END DEBUGGING
	// SwitchInlineMonitoringReasonUnknown captures enum value "unknown"
	SwitchInlineMonitoringReasonUnknown string = "unknown"
)

// prop value enum
func (o *SwitchInlineMonitoring) validateReasonEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, switchInlineMonitoringTypeReasonPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *SwitchInlineMonitoring) validateReason(formats strfmt.Registry) error {
	if swag.IsZero(o.Reason) { // not required
		return nil
	}

	// value enum
	if err := o.validateReasonEnum("info"+"."+"monitoring"+"."+"reason", "body", *o.Reason); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this switch inline monitoring based on the context it is used
func (o *SwitchInlineMonitoring) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateMonitored(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateReason(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SwitchInlineMonitoring) contextValidateMonitored(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "info"+"."+"monitoring"+"."+"monitored", "body", o.Monitored); err != nil {
		return err
	}

	return nil
}

func (o *SwitchInlineMonitoring) contextValidateReason(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "info"+"."+"monitoring"+"."+"reason", "body", o.Reason); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SwitchInlineMonitoring) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SwitchInlineMonitoring) UnmarshalBinary(b []byte) error {
	var res SwitchInlineMonitoring
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SwitchInlineSnmp switch inline snmp
swagger:model switch_inline_snmp
*/
type SwitchInlineSnmp struct {

	// Community String or SNMPv3 Username.
	User *string `json:"user,omitempty"`

	// SNMP Version.
	// Enum: ["snmpv1","snmpv2c","snmpv3"]
	Version *string `json:"version,omitempty"`
}

// Validate validates this switch inline snmp
func (o *SwitchInlineSnmp) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var switchInlineSnmpTypeVersionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["snmpv1","snmpv2c","snmpv3"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		switchInlineSnmpTypeVersionPropEnum = append(switchInlineSnmpTypeVersionPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// switch_inline_snmp
	// SwitchInlineSnmp
	// version
	// Version
	// snmpv1
	// END DEBUGGING
	// SwitchInlineSnmpVersionSnmpv1 captures enum value "snmpv1"
	SwitchInlineSnmpVersionSnmpv1 string = "snmpv1"

	// BEGIN DEBUGGING
	// switch_inline_snmp
	// SwitchInlineSnmp
	// version
	// Version
	// snmpv2c
	// END DEBUGGING
	// SwitchInlineSnmpVersionSnmpv2c captures enum value "snmpv2c"
	SwitchInlineSnmpVersionSnmpv2c string = "snmpv2c"

	// BEGIN DEBUGGING
	// switch_inline_snmp
	// SwitchInlineSnmp
	// version
	// Version
	// snmpv3
	// END DEBUGGING
	// SwitchInlineSnmpVersionSnmpv3 captures enum value "snmpv3"
	SwitchInlineSnmpVersionSnmpv3 string = "snmpv3"
)

// prop value enum
func (o *SwitchInlineSnmp) validateVersionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, switchInlineSnmpTypeVersionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *SwitchInlineSnmp) validateVersion(formats strfmt.Registry) error {
	if swag.IsZero(o.Version) { // not required
		return nil
	}

	// value enum
	if err := o.validateVersionEnum("info"+"."+"snmp"+"."+"version", "body", *o.Version); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this switch inline snmp based on context it is used
func (o *SwitchInlineSnmp) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SwitchInlineSnmp) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SwitchInlineSnmp) UnmarshalBinary(b []byte) error {
	var res SwitchInlineSnmp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}