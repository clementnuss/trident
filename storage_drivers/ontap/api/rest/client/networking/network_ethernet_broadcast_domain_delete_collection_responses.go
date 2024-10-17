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

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// NetworkEthernetBroadcastDomainDeleteCollectionReader is a Reader for the NetworkEthernetBroadcastDomainDeleteCollection structure.
type NetworkEthernetBroadcastDomainDeleteCollectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NetworkEthernetBroadcastDomainDeleteCollectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNetworkEthernetBroadcastDomainDeleteCollectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNetworkEthernetBroadcastDomainDeleteCollectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNetworkEthernetBroadcastDomainDeleteCollectionOK creates a NetworkEthernetBroadcastDomainDeleteCollectionOK with default headers values
func NewNetworkEthernetBroadcastDomainDeleteCollectionOK() *NetworkEthernetBroadcastDomainDeleteCollectionOK {
	return &NetworkEthernetBroadcastDomainDeleteCollectionOK{}
}

/*
NetworkEthernetBroadcastDomainDeleteCollectionOK describes a response with status code 200, with default header values.

OK
*/
type NetworkEthernetBroadcastDomainDeleteCollectionOK struct {
}

// IsSuccess returns true when this network ethernet broadcast domain delete collection o k response has a 2xx status code
func (o *NetworkEthernetBroadcastDomainDeleteCollectionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this network ethernet broadcast domain delete collection o k response has a 3xx status code
func (o *NetworkEthernetBroadcastDomainDeleteCollectionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this network ethernet broadcast domain delete collection o k response has a 4xx status code
func (o *NetworkEthernetBroadcastDomainDeleteCollectionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this network ethernet broadcast domain delete collection o k response has a 5xx status code
func (o *NetworkEthernetBroadcastDomainDeleteCollectionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this network ethernet broadcast domain delete collection o k response a status code equal to that given
func (o *NetworkEthernetBroadcastDomainDeleteCollectionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the network ethernet broadcast domain delete collection o k response
func (o *NetworkEthernetBroadcastDomainDeleteCollectionOK) Code() int {
	return 200
}

func (o *NetworkEthernetBroadcastDomainDeleteCollectionOK) Error() string {
	return fmt.Sprintf("[DELETE /network/ethernet/broadcast-domains][%d] networkEthernetBroadcastDomainDeleteCollectionOK", 200)
}

func (o *NetworkEthernetBroadcastDomainDeleteCollectionOK) String() string {
	return fmt.Sprintf("[DELETE /network/ethernet/broadcast-domains][%d] networkEthernetBroadcastDomainDeleteCollectionOK", 200)
}

func (o *NetworkEthernetBroadcastDomainDeleteCollectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNetworkEthernetBroadcastDomainDeleteCollectionDefault creates a NetworkEthernetBroadcastDomainDeleteCollectionDefault with default headers values
func NewNetworkEthernetBroadcastDomainDeleteCollectionDefault(code int) *NetworkEthernetBroadcastDomainDeleteCollectionDefault {
	return &NetworkEthernetBroadcastDomainDeleteCollectionDefault{
		_statusCode: code,
	}
}

/*
	NetworkEthernetBroadcastDomainDeleteCollectionDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 1967103 | A broadcast domain with ports cannot be deleted. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type NetworkEthernetBroadcastDomainDeleteCollectionDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this network ethernet broadcast domain delete collection default response has a 2xx status code
func (o *NetworkEthernetBroadcastDomainDeleteCollectionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this network ethernet broadcast domain delete collection default response has a 3xx status code
func (o *NetworkEthernetBroadcastDomainDeleteCollectionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this network ethernet broadcast domain delete collection default response has a 4xx status code
func (o *NetworkEthernetBroadcastDomainDeleteCollectionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this network ethernet broadcast domain delete collection default response has a 5xx status code
func (o *NetworkEthernetBroadcastDomainDeleteCollectionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this network ethernet broadcast domain delete collection default response a status code equal to that given
func (o *NetworkEthernetBroadcastDomainDeleteCollectionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the network ethernet broadcast domain delete collection default response
func (o *NetworkEthernetBroadcastDomainDeleteCollectionDefault) Code() int {
	return o._statusCode
}

func (o *NetworkEthernetBroadcastDomainDeleteCollectionDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /network/ethernet/broadcast-domains][%d] network_ethernet_broadcast_domain_delete_collection default %s", o._statusCode, payload)
}

func (o *NetworkEthernetBroadcastDomainDeleteCollectionDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /network/ethernet/broadcast-domains][%d] network_ethernet_broadcast_domain_delete_collection default %s", o._statusCode, payload)
}

func (o *NetworkEthernetBroadcastDomainDeleteCollectionDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NetworkEthernetBroadcastDomainDeleteCollectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
NetworkEthernetBroadcastDomainDeleteCollectionBody network ethernet broadcast domain delete collection body
swagger:model NetworkEthernetBroadcastDomainDeleteCollectionBody
*/
type NetworkEthernetBroadcastDomainDeleteCollectionBody struct {

	// broadcast domain response inline records
	BroadcastDomainResponseInlineRecords []*models.BroadcastDomain `json:"records,omitempty"`
}

// Validate validates this network ethernet broadcast domain delete collection body
func (o *NetworkEthernetBroadcastDomainDeleteCollectionBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateBroadcastDomainResponseInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NetworkEthernetBroadcastDomainDeleteCollectionBody) validateBroadcastDomainResponseInlineRecords(formats strfmt.Registry) error {
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

// ContextValidate validate this network ethernet broadcast domain delete collection body based on the context it is used
func (o *NetworkEthernetBroadcastDomainDeleteCollectionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateBroadcastDomainResponseInlineRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NetworkEthernetBroadcastDomainDeleteCollectionBody) contextValidateBroadcastDomainResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (o *NetworkEthernetBroadcastDomainDeleteCollectionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NetworkEthernetBroadcastDomainDeleteCollectionBody) UnmarshalBinary(b []byte) error {
	var res NetworkEthernetBroadcastDomainDeleteCollectionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}