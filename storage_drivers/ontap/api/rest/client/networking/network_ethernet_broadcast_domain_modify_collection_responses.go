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

// NetworkEthernetBroadcastDomainModifyCollectionReader is a Reader for the NetworkEthernetBroadcastDomainModifyCollection structure.
type NetworkEthernetBroadcastDomainModifyCollectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NetworkEthernetBroadcastDomainModifyCollectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNetworkEthernetBroadcastDomainModifyCollectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNetworkEthernetBroadcastDomainModifyCollectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNetworkEthernetBroadcastDomainModifyCollectionOK creates a NetworkEthernetBroadcastDomainModifyCollectionOK with default headers values
func NewNetworkEthernetBroadcastDomainModifyCollectionOK() *NetworkEthernetBroadcastDomainModifyCollectionOK {
	return &NetworkEthernetBroadcastDomainModifyCollectionOK{}
}

/*
NetworkEthernetBroadcastDomainModifyCollectionOK describes a response with status code 200, with default header values.

OK
*/
type NetworkEthernetBroadcastDomainModifyCollectionOK struct {
}

// IsSuccess returns true when this network ethernet broadcast domain modify collection o k response has a 2xx status code
func (o *NetworkEthernetBroadcastDomainModifyCollectionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this network ethernet broadcast domain modify collection o k response has a 3xx status code
func (o *NetworkEthernetBroadcastDomainModifyCollectionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this network ethernet broadcast domain modify collection o k response has a 4xx status code
func (o *NetworkEthernetBroadcastDomainModifyCollectionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this network ethernet broadcast domain modify collection o k response has a 5xx status code
func (o *NetworkEthernetBroadcastDomainModifyCollectionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this network ethernet broadcast domain modify collection o k response a status code equal to that given
func (o *NetworkEthernetBroadcastDomainModifyCollectionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the network ethernet broadcast domain modify collection o k response
func (o *NetworkEthernetBroadcastDomainModifyCollectionOK) Code() int {
	return 200
}

func (o *NetworkEthernetBroadcastDomainModifyCollectionOK) Error() string {
	return fmt.Sprintf("[PATCH /network/ethernet/broadcast-domains][%d] networkEthernetBroadcastDomainModifyCollectionOK", 200)
}

func (o *NetworkEthernetBroadcastDomainModifyCollectionOK) String() string {
	return fmt.Sprintf("[PATCH /network/ethernet/broadcast-domains][%d] networkEthernetBroadcastDomainModifyCollectionOK", 200)
}

func (o *NetworkEthernetBroadcastDomainModifyCollectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNetworkEthernetBroadcastDomainModifyCollectionDefault creates a NetworkEthernetBroadcastDomainModifyCollectionDefault with default headers values
func NewNetworkEthernetBroadcastDomainModifyCollectionDefault(code int) *NetworkEthernetBroadcastDomainModifyCollectionDefault {
	return &NetworkEthernetBroadcastDomainModifyCollectionDefault{
		_statusCode: code,
	}
}

/*
	NetworkEthernetBroadcastDomainModifyCollectionDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 1376484 | Cannot change the MTU of a VLAN to be greater than the MTU of the port hosting it. |
| 1377267 | The specified IPspace does not exist. |
| 1377269 | Failed to lookup the specified IPspace. |
| 1377560 | Broadcast domain already exists in specified IPspace. |
| 1377575 | Remove associated subnets before deleting this broadcast domain. |
| 1377605 | Moving the system-generated broadcast domain to another IPspace is not supported. |
| 1377609 | Updates are partially complete. Updating broadcast domain attributes on some or all of the ports in the broadcast domain have failed. |
| 1966460 | The specified MTU is either too large or too small. |
| 1967082 | The specified ipspace.name does not match the IPspace name of ipspace.uuid. |
| 1967150 | The specified ipspace.uuid is not valid. |
| 1967151 | The specified ipspace.uuid and ipspace.name do not match. |
| 1967152 | Patching IPspace for a broadcast domain requires an effective cluster version of 9.7 or later. |
| 53280884 | The MTU of the broadcast domain cannot be modified on this platform. |
| 53282013 | Broadcast domain cannot be renamed because the name is reserved by the system. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type NetworkEthernetBroadcastDomainModifyCollectionDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this network ethernet broadcast domain modify collection default response has a 2xx status code
func (o *NetworkEthernetBroadcastDomainModifyCollectionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this network ethernet broadcast domain modify collection default response has a 3xx status code
func (o *NetworkEthernetBroadcastDomainModifyCollectionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this network ethernet broadcast domain modify collection default response has a 4xx status code
func (o *NetworkEthernetBroadcastDomainModifyCollectionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this network ethernet broadcast domain modify collection default response has a 5xx status code
func (o *NetworkEthernetBroadcastDomainModifyCollectionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this network ethernet broadcast domain modify collection default response a status code equal to that given
func (o *NetworkEthernetBroadcastDomainModifyCollectionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the network ethernet broadcast domain modify collection default response
func (o *NetworkEthernetBroadcastDomainModifyCollectionDefault) Code() int {
	return o._statusCode
}

func (o *NetworkEthernetBroadcastDomainModifyCollectionDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /network/ethernet/broadcast-domains][%d] network_ethernet_broadcast_domain_modify_collection default %s", o._statusCode, payload)
}

func (o *NetworkEthernetBroadcastDomainModifyCollectionDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /network/ethernet/broadcast-domains][%d] network_ethernet_broadcast_domain_modify_collection default %s", o._statusCode, payload)
}

func (o *NetworkEthernetBroadcastDomainModifyCollectionDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NetworkEthernetBroadcastDomainModifyCollectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
NetworkEthernetBroadcastDomainModifyCollectionBody network ethernet broadcast domain modify collection body
swagger:model NetworkEthernetBroadcastDomainModifyCollectionBody
*/
type NetworkEthernetBroadcastDomainModifyCollectionBody struct {

	// links
	Links *models.BroadcastDomainInlineLinks `json:"_links,omitempty"`

	// Ports that belong to the broadcast domain
	// Read Only: true
	BroadcastDomainInlinePorts []*models.BroadcastDomainInlinePortsInlineArrayItem `json:"ports,omitempty"`

	// broadcast domain response inline records
	BroadcastDomainResponseInlineRecords []*models.BroadcastDomain `json:"records,omitempty"`

	// ipspace
	Ipspace *models.BroadcastDomainInlineIpspace `json:"ipspace,omitempty"`

	// Maximum transmission unit, largest packet size on this network
	// Example: 1500
	// Minimum: 68
	Mtu *int64 `json:"mtu,omitempty"`

	// Name of the broadcast domain, scoped to its IPspace
	// Example: bd1
	Name *string `json:"name,omitempty"`

	// Broadcast domain UUID
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	// Read Only: true
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this network ethernet broadcast domain modify collection body
func (o *NetworkEthernetBroadcastDomainModifyCollectionBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateBroadcastDomainInlinePorts(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateBroadcastDomainResponseInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateIpspace(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMtu(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NetworkEthernetBroadcastDomainModifyCollectionBody) validateLinks(formats strfmt.Registry) error {
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

func (o *NetworkEthernetBroadcastDomainModifyCollectionBody) validateBroadcastDomainInlinePorts(formats strfmt.Registry) error {
	if swag.IsZero(o.BroadcastDomainInlinePorts) { // not required
		return nil
	}

	for i := 0; i < len(o.BroadcastDomainInlinePorts); i++ {
		if swag.IsZero(o.BroadcastDomainInlinePorts[i]) { // not required
			continue
		}

		if o.BroadcastDomainInlinePorts[i] != nil {
			if err := o.BroadcastDomainInlinePorts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "ports" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *NetworkEthernetBroadcastDomainModifyCollectionBody) validateBroadcastDomainResponseInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(o.BroadcastDomainResponseInlineRecords) { // not required
		return nil
	}

	for i := 0; i < len(o.BroadcastDomainResponseInlineRecords); i++ {
		if swag.IsZero(o.BroadcastDomainResponseInlineRecords[i]) { // not required
			continue
		}

		if o.BroadcastDomainResponseInlineRecords[i] != nil {
			if err := o.BroadcastDomainResponseInlineRecords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *NetworkEthernetBroadcastDomainModifyCollectionBody) validateIpspace(formats strfmt.Registry) error {
	if swag.IsZero(o.Ipspace) { // not required
		return nil
	}

	if o.Ipspace != nil {
		if err := o.Ipspace.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "ipspace")
			}
			return err
		}
	}

	return nil
}

func (o *NetworkEthernetBroadcastDomainModifyCollectionBody) validateMtu(formats strfmt.Registry) error {
	if swag.IsZero(o.Mtu) { // not required
		return nil
	}

	if err := validate.MinimumInt("info"+"."+"mtu", "body", *o.Mtu, 68, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this network ethernet broadcast domain modify collection body based on the context it is used
func (o *NetworkEthernetBroadcastDomainModifyCollectionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateBroadcastDomainInlinePorts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateBroadcastDomainResponseInlineRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateIpspace(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateUUID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NetworkEthernetBroadcastDomainModifyCollectionBody) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (o *NetworkEthernetBroadcastDomainModifyCollectionBody) contextValidateBroadcastDomainInlinePorts(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "info"+"."+"ports", "body", []*models.BroadcastDomainInlinePortsInlineArrayItem(o.BroadcastDomainInlinePorts)); err != nil {
		return err
	}

	for i := 0; i < len(o.BroadcastDomainInlinePorts); i++ {

		if o.BroadcastDomainInlinePorts[i] != nil {
			if err := o.BroadcastDomainInlinePorts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "ports" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *NetworkEthernetBroadcastDomainModifyCollectionBody) contextValidateBroadcastDomainResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.BroadcastDomainResponseInlineRecords); i++ {

		if o.BroadcastDomainResponseInlineRecords[i] != nil {
			if err := o.BroadcastDomainResponseInlineRecords[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *NetworkEthernetBroadcastDomainModifyCollectionBody) contextValidateIpspace(ctx context.Context, formats strfmt.Registry) error {

	if o.Ipspace != nil {
		if err := o.Ipspace.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "ipspace")
			}
			return err
		}
	}

	return nil
}

func (o *NetworkEthernetBroadcastDomainModifyCollectionBody) contextValidateUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "info"+"."+"uuid", "body", o.UUID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *NetworkEthernetBroadcastDomainModifyCollectionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NetworkEthernetBroadcastDomainModifyCollectionBody) UnmarshalBinary(b []byte) error {
	var res NetworkEthernetBroadcastDomainModifyCollectionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
BroadcastDomainInlineLinks broadcast domain inline links
swagger:model broadcast_domain_inline__links
*/
type BroadcastDomainInlineLinks struct {

	// self
	Self *models.Href `json:"self,omitempty"`
}

// Validate validates this broadcast domain inline links
func (o *BroadcastDomainInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *BroadcastDomainInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(o.Self) { // not required
		return nil
	}

	if o.Self != nil {
		if err := o.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this broadcast domain inline links based on the context it is used
func (o *BroadcastDomainInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *BroadcastDomainInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if o.Self != nil {
		if err := o.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *BroadcastDomainInlineLinks) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *BroadcastDomainInlineLinks) UnmarshalBinary(b []byte) error {
	var res BroadcastDomainInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
BroadcastDomainInlineIpspace Applies to both SVM and cluster-scoped objects. Either the UUID or name is supplied on input.
swagger:model broadcast_domain_inline_ipspace
*/
type BroadcastDomainInlineIpspace struct {

	// links
	Links *models.BroadcastDomainInlineIpspaceInlineLinks `json:"_links,omitempty"`

	// IPspace name
	// Example: Default
	Name *string `json:"name,omitempty"`

	// IPspace UUID
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this broadcast domain inline ipspace
func (o *BroadcastDomainInlineIpspace) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *BroadcastDomainInlineIpspace) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(o.Links) { // not required
		return nil
	}

	if o.Links != nil {
		if err := o.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "ipspace" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this broadcast domain inline ipspace based on the context it is used
func (o *BroadcastDomainInlineIpspace) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *BroadcastDomainInlineIpspace) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if o.Links != nil {
		if err := o.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "ipspace" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *BroadcastDomainInlineIpspace) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *BroadcastDomainInlineIpspace) UnmarshalBinary(b []byte) error {
	var res BroadcastDomainInlineIpspace
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
BroadcastDomainInlineIpspaceInlineLinks broadcast domain inline ipspace inline links
swagger:model broadcast_domain_inline_ipspace_inline__links
*/
type BroadcastDomainInlineIpspaceInlineLinks struct {

	// self
	Self *models.Href `json:"self,omitempty"`
}

// Validate validates this broadcast domain inline ipspace inline links
func (o *BroadcastDomainInlineIpspaceInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *BroadcastDomainInlineIpspaceInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(o.Self) { // not required
		return nil
	}

	if o.Self != nil {
		if err := o.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "ipspace" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this broadcast domain inline ipspace inline links based on the context it is used
func (o *BroadcastDomainInlineIpspaceInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *BroadcastDomainInlineIpspaceInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if o.Self != nil {
		if err := o.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "ipspace" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *BroadcastDomainInlineIpspaceInlineLinks) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *BroadcastDomainInlineIpspaceInlineLinks) UnmarshalBinary(b []byte) error {
	var res BroadcastDomainInlineIpspaceInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
BroadcastDomainInlinePortsInlineArrayItem Port UUID along with readable names
swagger:model broadcast_domain_inline_ports_inline_array_item
*/
type BroadcastDomainInlinePortsInlineArrayItem struct {

	// links
	Links *models.BroadcastDomainInlinePortsInlineArrayItemInlineLinks `json:"_links,omitempty"`

	// name
	// Example: e1b
	Name *string `json:"name,omitempty"`

	// node
	Node *models.BroadcastDomainInlinePortsInlineArrayItemInlineNode `json:"node,omitempty"`

	// uuid
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this broadcast domain inline ports inline array item
func (o *BroadcastDomainInlinePortsInlineArrayItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *BroadcastDomainInlinePortsInlineArrayItem) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(o.Links) { // not required
		return nil
	}

	if o.Links != nil {
		if err := o.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (o *BroadcastDomainInlinePortsInlineArrayItem) validateNode(formats strfmt.Registry) error {
	if swag.IsZero(o.Node) { // not required
		return nil
	}

	if o.Node != nil {
		if err := o.Node.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this broadcast domain inline ports inline array item based on the context it is used
func (o *BroadcastDomainInlinePortsInlineArrayItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateNode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *BroadcastDomainInlinePortsInlineArrayItem) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if o.Links != nil {
		if err := o.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (o *BroadcastDomainInlinePortsInlineArrayItem) contextValidateNode(ctx context.Context, formats strfmt.Registry) error {

	if o.Node != nil {
		if err := o.Node.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *BroadcastDomainInlinePortsInlineArrayItem) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *BroadcastDomainInlinePortsInlineArrayItem) UnmarshalBinary(b []byte) error {
	var res BroadcastDomainInlinePortsInlineArrayItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
BroadcastDomainInlinePortsInlineArrayItemInlineLinks broadcast domain inline ports inline array item inline links
swagger:model broadcast_domain_inline_ports_inline_array_item_inline__links
*/
type BroadcastDomainInlinePortsInlineArrayItemInlineLinks struct {

	// self
	Self *models.Href `json:"self,omitempty"`
}

// Validate validates this broadcast domain inline ports inline array item inline links
func (o *BroadcastDomainInlinePortsInlineArrayItemInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *BroadcastDomainInlinePortsInlineArrayItemInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(o.Self) { // not required
		return nil
	}

	if o.Self != nil {
		if err := o.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this broadcast domain inline ports inline array item inline links based on the context it is used
func (o *BroadcastDomainInlinePortsInlineArrayItemInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *BroadcastDomainInlinePortsInlineArrayItemInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if o.Self != nil {
		if err := o.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *BroadcastDomainInlinePortsInlineArrayItemInlineLinks) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *BroadcastDomainInlinePortsInlineArrayItemInlineLinks) UnmarshalBinary(b []byte) error {
	var res BroadcastDomainInlinePortsInlineArrayItemInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
BroadcastDomainInlinePortsInlineArrayItemInlineNode broadcast domain inline ports inline array item inline node
swagger:model broadcast_domain_inline_ports_inline_array_item_inline_node
*/
type BroadcastDomainInlinePortsInlineArrayItemInlineNode struct {

	// Name of node on which the port is located.
	// Example: node1
	Name *string `json:"name,omitempty"`
}

// Validate validates this broadcast domain inline ports inline array item inline node
func (o *BroadcastDomainInlinePortsInlineArrayItemInlineNode) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this broadcast domain inline ports inline array item inline node based on context it is used
func (o *BroadcastDomainInlinePortsInlineArrayItemInlineNode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *BroadcastDomainInlinePortsInlineArrayItemInlineNode) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *BroadcastDomainInlinePortsInlineArrayItemInlineNode) UnmarshalBinary(b []byte) error {
	var res BroadcastDomainInlinePortsInlineArrayItemInlineNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}