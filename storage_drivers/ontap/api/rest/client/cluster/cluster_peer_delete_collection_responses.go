// Code generated by go-swagger; DO NOT EDIT.

package cluster

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

// ClusterPeerDeleteCollectionReader is a Reader for the ClusterPeerDeleteCollection structure.
type ClusterPeerDeleteCollectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClusterPeerDeleteCollectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewClusterPeerDeleteCollectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewClusterPeerDeleteCollectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewClusterPeerDeleteCollectionOK creates a ClusterPeerDeleteCollectionOK with default headers values
func NewClusterPeerDeleteCollectionOK() *ClusterPeerDeleteCollectionOK {
	return &ClusterPeerDeleteCollectionOK{}
}

/*
ClusterPeerDeleteCollectionOK describes a response with status code 200, with default header values.

OK
*/
type ClusterPeerDeleteCollectionOK struct {
}

// IsSuccess returns true when this cluster peer delete collection o k response has a 2xx status code
func (o *ClusterPeerDeleteCollectionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cluster peer delete collection o k response has a 3xx status code
func (o *ClusterPeerDeleteCollectionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cluster peer delete collection o k response has a 4xx status code
func (o *ClusterPeerDeleteCollectionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cluster peer delete collection o k response has a 5xx status code
func (o *ClusterPeerDeleteCollectionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cluster peer delete collection o k response a status code equal to that given
func (o *ClusterPeerDeleteCollectionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cluster peer delete collection o k response
func (o *ClusterPeerDeleteCollectionOK) Code() int {
	return 200
}

func (o *ClusterPeerDeleteCollectionOK) Error() string {
	return fmt.Sprintf("[DELETE /cluster/peers][%d] clusterPeerDeleteCollectionOK", 200)
}

func (o *ClusterPeerDeleteCollectionOK) String() string {
	return fmt.Sprintf("[DELETE /cluster/peers][%d] clusterPeerDeleteCollectionOK", 200)
}

func (o *ClusterPeerDeleteCollectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewClusterPeerDeleteCollectionDefault creates a ClusterPeerDeleteCollectionDefault with default headers values
func NewClusterPeerDeleteCollectionDefault(code int) *ClusterPeerDeleteCollectionDefault {
	return &ClusterPeerDeleteCollectionDefault{
		_statusCode: code,
	}
}

/*
	ClusterPeerDeleteCollectionDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 4653079 | Unable to delete peer relationship. |
| 4663070 | Unable to delete cluster peer relationship due to an ongoing Vserver migration. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type ClusterPeerDeleteCollectionDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this cluster peer delete collection default response has a 2xx status code
func (o *ClusterPeerDeleteCollectionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cluster peer delete collection default response has a 3xx status code
func (o *ClusterPeerDeleteCollectionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cluster peer delete collection default response has a 4xx status code
func (o *ClusterPeerDeleteCollectionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cluster peer delete collection default response has a 5xx status code
func (o *ClusterPeerDeleteCollectionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cluster peer delete collection default response a status code equal to that given
func (o *ClusterPeerDeleteCollectionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the cluster peer delete collection default response
func (o *ClusterPeerDeleteCollectionDefault) Code() int {
	return o._statusCode
}

func (o *ClusterPeerDeleteCollectionDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /cluster/peers][%d] cluster_peer_delete_collection default %s", o._statusCode, payload)
}

func (o *ClusterPeerDeleteCollectionDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /cluster/peers][%d] cluster_peer_delete_collection default %s", o._statusCode, payload)
}

func (o *ClusterPeerDeleteCollectionDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ClusterPeerDeleteCollectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ClusterPeerDeleteCollectionBody cluster peer delete collection body
swagger:model ClusterPeerDeleteCollectionBody
*/
type ClusterPeerDeleteCollectionBody struct {

	// cluster peer response inline records
	ClusterPeerResponseInlineRecords []*models.ClusterPeer `json:"records,omitempty"`
}

// Validate validates this cluster peer delete collection body
func (o *ClusterPeerDeleteCollectionBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateClusterPeerResponseInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ClusterPeerDeleteCollectionBody) validateClusterPeerResponseInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(o.ClusterPeerResponseInlineRecords) { // not required
		return nil
	}

	for i := 0; i < len(o.ClusterPeerResponseInlineRecords); i++ {
		if swag.IsZero(o.ClusterPeerResponseInlineRecords[i]) { // not required
			continue
		}

		if o.ClusterPeerResponseInlineRecords[i] != nil {
			if err := o.ClusterPeerResponseInlineRecords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this cluster peer delete collection body based on the context it is used
func (o *ClusterPeerDeleteCollectionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateClusterPeerResponseInlineRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ClusterPeerDeleteCollectionBody) contextValidateClusterPeerResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.ClusterPeerResponseInlineRecords); i++ {

		if o.ClusterPeerResponseInlineRecords[i] != nil {
			if err := o.ClusterPeerResponseInlineRecords[i].ContextValidate(ctx, formats); err != nil {
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
func (o *ClusterPeerDeleteCollectionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ClusterPeerDeleteCollectionBody) UnmarshalBinary(b []byte) error {
	var res ClusterPeerDeleteCollectionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}