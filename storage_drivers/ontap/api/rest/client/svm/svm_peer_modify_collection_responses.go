// Code generated by go-swagger; DO NOT EDIT.

package svm

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

// SvmPeerModifyCollectionReader is a Reader for the SvmPeerModifyCollection structure.
type SvmPeerModifyCollectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SvmPeerModifyCollectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSvmPeerModifyCollectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewSvmPeerModifyCollectionAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSvmPeerModifyCollectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSvmPeerModifyCollectionOK creates a SvmPeerModifyCollectionOK with default headers values
func NewSvmPeerModifyCollectionOK() *SvmPeerModifyCollectionOK {
	return &SvmPeerModifyCollectionOK{}
}

/*
SvmPeerModifyCollectionOK describes a response with status code 200, with default header values.

OK
*/
type SvmPeerModifyCollectionOK struct {
}

// IsSuccess returns true when this svm peer modify collection o k response has a 2xx status code
func (o *SvmPeerModifyCollectionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this svm peer modify collection o k response has a 3xx status code
func (o *SvmPeerModifyCollectionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this svm peer modify collection o k response has a 4xx status code
func (o *SvmPeerModifyCollectionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this svm peer modify collection o k response has a 5xx status code
func (o *SvmPeerModifyCollectionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this svm peer modify collection o k response a status code equal to that given
func (o *SvmPeerModifyCollectionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the svm peer modify collection o k response
func (o *SvmPeerModifyCollectionOK) Code() int {
	return 200
}

func (o *SvmPeerModifyCollectionOK) Error() string {
	return fmt.Sprintf("[PATCH /svm/peers][%d] svmPeerModifyCollectionOK", 200)
}

func (o *SvmPeerModifyCollectionOK) String() string {
	return fmt.Sprintf("[PATCH /svm/peers][%d] svmPeerModifyCollectionOK", 200)
}

func (o *SvmPeerModifyCollectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSvmPeerModifyCollectionAccepted creates a SvmPeerModifyCollectionAccepted with default headers values
func NewSvmPeerModifyCollectionAccepted() *SvmPeerModifyCollectionAccepted {
	return &SvmPeerModifyCollectionAccepted{}
}

/*
SvmPeerModifyCollectionAccepted describes a response with status code 202, with default header values.

Accepted
*/
type SvmPeerModifyCollectionAccepted struct {
}

// IsSuccess returns true when this svm peer modify collection accepted response has a 2xx status code
func (o *SvmPeerModifyCollectionAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this svm peer modify collection accepted response has a 3xx status code
func (o *SvmPeerModifyCollectionAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this svm peer modify collection accepted response has a 4xx status code
func (o *SvmPeerModifyCollectionAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this svm peer modify collection accepted response has a 5xx status code
func (o *SvmPeerModifyCollectionAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this svm peer modify collection accepted response a status code equal to that given
func (o *SvmPeerModifyCollectionAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the svm peer modify collection accepted response
func (o *SvmPeerModifyCollectionAccepted) Code() int {
	return 202
}

func (o *SvmPeerModifyCollectionAccepted) Error() string {
	return fmt.Sprintf("[PATCH /svm/peers][%d] svmPeerModifyCollectionAccepted", 202)
}

func (o *SvmPeerModifyCollectionAccepted) String() string {
	return fmt.Sprintf("[PATCH /svm/peers][%d] svmPeerModifyCollectionAccepted", 202)
}

func (o *SvmPeerModifyCollectionAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSvmPeerModifyCollectionDefault creates a SvmPeerModifyCollectionDefault with default headers values
func NewSvmPeerModifyCollectionDefault(code int) *SvmPeerModifyCollectionDefault {
	return &SvmPeerModifyCollectionDefault{
		_statusCode: code,
	}
}

/*
	SvmPeerModifyCollectionDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

<br/>
```
| Error codes | Description |
| ----------- | ----------- |
| 13434889    | Internal error. Wait and retry. |
| 26345575    | The specified peer cluster name and peer cluster UUID do not match. |
| 26345576    | Given peer state is invalid. |
| 26345577    | One of the following is required: applications, state, or name. |
| 26345578    | Internal error. Unable to retrieve local or peer SVM name. |
| 26345579    | The specified field is invalid. |
| 26345581    | Peer cluster name could not be retrieved or validated. |
| 9896077     | The peer relationship is in use by FlexCache. View the FlexCache relationships, delete them and retry the operation. |
| 9896088     | System generated a name for the peer SVM because of a naming conflict. Use the name property to uniquely identify the peer SVM alias name. |
```
<br/>
*/
type SvmPeerModifyCollectionDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this svm peer modify collection default response has a 2xx status code
func (o *SvmPeerModifyCollectionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this svm peer modify collection default response has a 3xx status code
func (o *SvmPeerModifyCollectionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this svm peer modify collection default response has a 4xx status code
func (o *SvmPeerModifyCollectionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this svm peer modify collection default response has a 5xx status code
func (o *SvmPeerModifyCollectionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this svm peer modify collection default response a status code equal to that given
func (o *SvmPeerModifyCollectionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the svm peer modify collection default response
func (o *SvmPeerModifyCollectionDefault) Code() int {
	return o._statusCode
}

func (o *SvmPeerModifyCollectionDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /svm/peers][%d] svm_peer_modify_collection default %s", o._statusCode, payload)
}

func (o *SvmPeerModifyCollectionDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /svm/peers][%d] svm_peer_modify_collection default %s", o._statusCode, payload)
}

func (o *SvmPeerModifyCollectionDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SvmPeerModifyCollectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
SvmPeerModifyCollectionBody svm peer modify collection body
swagger:model SvmPeerModifyCollectionBody
*/
type SvmPeerModifyCollectionBody struct {

	// links
	Links *models.SvmPeerInlineLinks `json:"_links,omitempty"`

	// Use this to suspend, resume or delete the SVM peer relationship even if the remote cluster is not accessible due to, for example, network connectivity issues.
	Force *bool `json:"force,omitempty"`

	// A peer SVM alias name to avoid a name conflict on the local cluster.
	Name *string `json:"name,omitempty"`

	// peer
	Peer *models.SvmPeerInlinePeer `json:"peer,omitempty"`

	// SVM peering state. To accept a pending SVM peer request, PATCH the state to "peered". To reject a pending SVM peer request, PATCH the state to "rejected". To suspend a peered SVM peer relationship, PATCH the state to "suspended". To resume a suspended SVM peer relationship, PATCH the state to "peered". The states "initiated", "pending", and "initializing" are system-generated and cannot be used for PATCH.
	// Example: peered
	// Enum: ["peered","rejected","suspended","initiated","pending","initializing"]
	State *string `json:"state,omitempty"`

	// svm
	Svm *models.SvmPeerInlineSvm `json:"svm,omitempty"`

	// A list of applications for an SVM peer relationship.
	// Example: ["snapmirror","lun_copy"]
	SvmPeerInlineApplications []*models.SvmPeerApplications `json:"applications,omitempty"`

	// svm peer response inline records
	SvmPeerResponseInlineRecords []*models.SvmPeer `json:"records,omitempty"`

	// SVM peer relationship UUID
	// Read Only: true
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this svm peer modify collection body
func (o *SvmPeerModifyCollectionBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePeer(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSvm(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSvmPeerInlineApplications(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSvmPeerResponseInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SvmPeerModifyCollectionBody) validateLinks(formats strfmt.Registry) error {
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

func (o *SvmPeerModifyCollectionBody) validatePeer(formats strfmt.Registry) error {
	if swag.IsZero(o.Peer) { // not required
		return nil
	}

	if o.Peer != nil {
		if err := o.Peer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "peer")
			}
			return err
		}
	}

	return nil
}

var svmPeerModifyCollectionBodyTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["peered","rejected","suspended","initiated","pending","initializing"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		svmPeerModifyCollectionBodyTypeStatePropEnum = append(svmPeerModifyCollectionBodyTypeStatePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// SvmPeerModifyCollectionBody
	// SvmPeerModifyCollectionBody
	// state
	// State
	// peered
	// END DEBUGGING
	// SvmPeerModifyCollectionBodyStatePeered captures enum value "peered"
	SvmPeerModifyCollectionBodyStatePeered string = "peered"

	// BEGIN DEBUGGING
	// SvmPeerModifyCollectionBody
	// SvmPeerModifyCollectionBody
	// state
	// State
	// rejected
	// END DEBUGGING
	// SvmPeerModifyCollectionBodyStateRejected captures enum value "rejected"
	SvmPeerModifyCollectionBodyStateRejected string = "rejected"

	// BEGIN DEBUGGING
	// SvmPeerModifyCollectionBody
	// SvmPeerModifyCollectionBody
	// state
	// State
	// suspended
	// END DEBUGGING
	// SvmPeerModifyCollectionBodyStateSuspended captures enum value "suspended"
	SvmPeerModifyCollectionBodyStateSuspended string = "suspended"

	// BEGIN DEBUGGING
	// SvmPeerModifyCollectionBody
	// SvmPeerModifyCollectionBody
	// state
	// State
	// initiated
	// END DEBUGGING
	// SvmPeerModifyCollectionBodyStateInitiated captures enum value "initiated"
	SvmPeerModifyCollectionBodyStateInitiated string = "initiated"

	// BEGIN DEBUGGING
	// SvmPeerModifyCollectionBody
	// SvmPeerModifyCollectionBody
	// state
	// State
	// pending
	// END DEBUGGING
	// SvmPeerModifyCollectionBodyStatePending captures enum value "pending"
	SvmPeerModifyCollectionBodyStatePending string = "pending"

	// BEGIN DEBUGGING
	// SvmPeerModifyCollectionBody
	// SvmPeerModifyCollectionBody
	// state
	// State
	// initializing
	// END DEBUGGING
	// SvmPeerModifyCollectionBodyStateInitializing captures enum value "initializing"
	SvmPeerModifyCollectionBodyStateInitializing string = "initializing"
)

// prop value enum
func (o *SvmPeerModifyCollectionBody) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, svmPeerModifyCollectionBodyTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *SvmPeerModifyCollectionBody) validateState(formats strfmt.Registry) error {
	if swag.IsZero(o.State) { // not required
		return nil
	}

	// value enum
	if err := o.validateStateEnum("info"+"."+"state", "body", *o.State); err != nil {
		return err
	}

	return nil
}

func (o *SvmPeerModifyCollectionBody) validateSvm(formats strfmt.Registry) error {
	if swag.IsZero(o.Svm) { // not required
		return nil
	}

	if o.Svm != nil {
		if err := o.Svm.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "svm")
			}
			return err
		}
	}

	return nil
}

func (o *SvmPeerModifyCollectionBody) validateSvmPeerInlineApplications(formats strfmt.Registry) error {
	if swag.IsZero(o.SvmPeerInlineApplications) { // not required
		return nil
	}

	for i := 0; i < len(o.SvmPeerInlineApplications); i++ {
		if swag.IsZero(o.SvmPeerInlineApplications[i]) { // not required
			continue
		}

		if o.SvmPeerInlineApplications[i] != nil {
			if err := o.SvmPeerInlineApplications[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "applications" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *SvmPeerModifyCollectionBody) validateSvmPeerResponseInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(o.SvmPeerResponseInlineRecords) { // not required
		return nil
	}

	for i := 0; i < len(o.SvmPeerResponseInlineRecords); i++ {
		if swag.IsZero(o.SvmPeerResponseInlineRecords[i]) { // not required
			continue
		}

		if o.SvmPeerResponseInlineRecords[i] != nil {
			if err := o.SvmPeerResponseInlineRecords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this svm peer modify collection body based on the context it is used
func (o *SvmPeerModifyCollectionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidatePeer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateSvm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateSvmPeerInlineApplications(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateSvmPeerResponseInlineRecords(ctx, formats); err != nil {
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

func (o *SvmPeerModifyCollectionBody) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (o *SvmPeerModifyCollectionBody) contextValidatePeer(ctx context.Context, formats strfmt.Registry) error {

	if o.Peer != nil {
		if err := o.Peer.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "peer")
			}
			return err
		}
	}

	return nil
}

func (o *SvmPeerModifyCollectionBody) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

	if o.Svm != nil {
		if err := o.Svm.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "svm")
			}
			return err
		}
	}

	return nil
}

func (o *SvmPeerModifyCollectionBody) contextValidateSvmPeerInlineApplications(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.SvmPeerInlineApplications); i++ {

		if o.SvmPeerInlineApplications[i] != nil {
			if err := o.SvmPeerInlineApplications[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "applications" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *SvmPeerModifyCollectionBody) contextValidateSvmPeerResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.SvmPeerResponseInlineRecords); i++ {

		if o.SvmPeerResponseInlineRecords[i] != nil {
			if err := o.SvmPeerResponseInlineRecords[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *SvmPeerModifyCollectionBody) contextValidateUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "info"+"."+"uuid", "body", o.UUID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SvmPeerModifyCollectionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SvmPeerModifyCollectionBody) UnmarshalBinary(b []byte) error {
	var res SvmPeerModifyCollectionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SvmPeerInlineLinks svm peer inline links
swagger:model svm_peer_inline__links
*/
type SvmPeerInlineLinks struct {

	// self
	Self *models.Href `json:"self,omitempty"`
}

// Validate validates this svm peer inline links
func (o *SvmPeerInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SvmPeerInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this svm peer inline links based on the context it is used
func (o *SvmPeerInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SvmPeerInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (o *SvmPeerInlineLinks) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SvmPeerInlineLinks) UnmarshalBinary(b []byte) error {
	var res SvmPeerInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SvmPeerInlinePeer Details for a peer SVM object.
swagger:model svm_peer_inline_peer
*/
type SvmPeerInlinePeer struct {

	// cluster
	Cluster *models.SvmPeerInlinePeerInlineCluster `json:"cluster,omitempty"`

	// svm
	Svm *models.SvmPeerInlinePeerInlineSvm `json:"svm,omitempty"`
}

// Validate validates this svm peer inline peer
func (o *SvmPeerInlinePeer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSvm(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SvmPeerInlinePeer) validateCluster(formats strfmt.Registry) error {
	if swag.IsZero(o.Cluster) { // not required
		return nil
	}

	if o.Cluster != nil {
		if err := o.Cluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "peer" + "." + "cluster")
			}
			return err
		}
	}

	return nil
}

func (o *SvmPeerInlinePeer) validateSvm(formats strfmt.Registry) error {
	if swag.IsZero(o.Svm) { // not required
		return nil
	}

	if o.Svm != nil {
		if err := o.Svm.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "peer" + "." + "svm")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this svm peer inline peer based on the context it is used
func (o *SvmPeerInlinePeer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateSvm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SvmPeerInlinePeer) contextValidateCluster(ctx context.Context, formats strfmt.Registry) error {

	if o.Cluster != nil {
		if err := o.Cluster.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "peer" + "." + "cluster")
			}
			return err
		}
	}

	return nil
}

func (o *SvmPeerInlinePeer) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

	if o.Svm != nil {
		if err := o.Svm.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "peer" + "." + "svm")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SvmPeerInlinePeer) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SvmPeerInlinePeer) UnmarshalBinary(b []byte) error {
	var res SvmPeerInlinePeer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SvmPeerInlinePeerInlineCluster svm peer inline peer inline cluster
swagger:model svm_peer_inline_peer_inline_cluster
*/
type SvmPeerInlinePeerInlineCluster struct {

	// links
	Links *models.SvmPeerInlinePeerInlineClusterInlineLinks `json:"_links,omitempty"`

	// name
	// Example: cluster2
	Name *string `json:"name,omitempty"`

	// uuid
	// Example: ebe27c49-1adf-4496-8335-ab862aebebf2
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this svm peer inline peer inline cluster
func (o *SvmPeerInlinePeerInlineCluster) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SvmPeerInlinePeerInlineCluster) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(o.Links) { // not required
		return nil
	}

	if o.Links != nil {
		if err := o.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "peer" + "." + "cluster" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this svm peer inline peer inline cluster based on the context it is used
func (o *SvmPeerInlinePeerInlineCluster) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SvmPeerInlinePeerInlineCluster) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if o.Links != nil {
		if err := o.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "peer" + "." + "cluster" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SvmPeerInlinePeerInlineCluster) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SvmPeerInlinePeerInlineCluster) UnmarshalBinary(b []byte) error {
	var res SvmPeerInlinePeerInlineCluster
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SvmPeerInlinePeerInlineClusterInlineLinks svm peer inline peer inline cluster inline links
swagger:model svm_peer_inline_peer_inline_cluster_inline__links
*/
type SvmPeerInlinePeerInlineClusterInlineLinks struct {

	// self
	Self *models.Href `json:"self,omitempty"`
}

// Validate validates this svm peer inline peer inline cluster inline links
func (o *SvmPeerInlinePeerInlineClusterInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SvmPeerInlinePeerInlineClusterInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(o.Self) { // not required
		return nil
	}

	if o.Self != nil {
		if err := o.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "peer" + "." + "cluster" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this svm peer inline peer inline cluster inline links based on the context it is used
func (o *SvmPeerInlinePeerInlineClusterInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SvmPeerInlinePeerInlineClusterInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if o.Self != nil {
		if err := o.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "peer" + "." + "cluster" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SvmPeerInlinePeerInlineClusterInlineLinks) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SvmPeerInlinePeerInlineClusterInlineLinks) UnmarshalBinary(b []byte) error {
	var res SvmPeerInlinePeerInlineClusterInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SvmPeerInlinePeerInlineSvm SVM, applies only to SVM-scoped objects.
swagger:model svm_peer_inline_peer_inline_svm
*/
type SvmPeerInlinePeerInlineSvm struct {

	// links
	Links *models.SvmPeerInlinePeerInlineSvmInlineLinks `json:"_links,omitempty"`

	// The name of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: svm1
	Name *string `json:"name,omitempty"`

	// The unique identifier of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this svm peer inline peer inline svm
func (o *SvmPeerInlinePeerInlineSvm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SvmPeerInlinePeerInlineSvm) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(o.Links) { // not required
		return nil
	}

	if o.Links != nil {
		if err := o.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "peer" + "." + "svm" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this svm peer inline peer inline svm based on the context it is used
func (o *SvmPeerInlinePeerInlineSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SvmPeerInlinePeerInlineSvm) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if o.Links != nil {
		if err := o.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "peer" + "." + "svm" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SvmPeerInlinePeerInlineSvm) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SvmPeerInlinePeerInlineSvm) UnmarshalBinary(b []byte) error {
	var res SvmPeerInlinePeerInlineSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SvmPeerInlinePeerInlineSvmInlineLinks svm peer inline peer inline svm inline links
swagger:model svm_peer_inline_peer_inline_svm_inline__links
*/
type SvmPeerInlinePeerInlineSvmInlineLinks struct {

	// self
	Self *models.Href `json:"self,omitempty"`
}

// Validate validates this svm peer inline peer inline svm inline links
func (o *SvmPeerInlinePeerInlineSvmInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SvmPeerInlinePeerInlineSvmInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(o.Self) { // not required
		return nil
	}

	if o.Self != nil {
		if err := o.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "peer" + "." + "svm" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this svm peer inline peer inline svm inline links based on the context it is used
func (o *SvmPeerInlinePeerInlineSvmInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SvmPeerInlinePeerInlineSvmInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if o.Self != nil {
		if err := o.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "peer" + "." + "svm" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SvmPeerInlinePeerInlineSvmInlineLinks) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SvmPeerInlinePeerInlineSvmInlineLinks) UnmarshalBinary(b []byte) error {
	var res SvmPeerInlinePeerInlineSvmInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SvmPeerInlineSvm Local SVM details
swagger:model svm_peer_inline_svm
*/
type SvmPeerInlineSvm struct {

	// links
	Links *models.SvmPeerInlineSvmInlineLinks `json:"_links,omitempty"`

	// The name of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: svm1
	Name *string `json:"name,omitempty"`

	// The unique identifier of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this svm peer inline svm
func (o *SvmPeerInlineSvm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SvmPeerInlineSvm) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(o.Links) { // not required
		return nil
	}

	if o.Links != nil {
		if err := o.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "svm" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this svm peer inline svm based on the context it is used
func (o *SvmPeerInlineSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SvmPeerInlineSvm) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if o.Links != nil {
		if err := o.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "svm" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SvmPeerInlineSvm) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SvmPeerInlineSvm) UnmarshalBinary(b []byte) error {
	var res SvmPeerInlineSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SvmPeerInlineSvmInlineLinks svm peer inline svm inline links
swagger:model svm_peer_inline_svm_inline__links
*/
type SvmPeerInlineSvmInlineLinks struct {

	// self
	Self *models.Href `json:"self,omitempty"`
}

// Validate validates this svm peer inline svm inline links
func (o *SvmPeerInlineSvmInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SvmPeerInlineSvmInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(o.Self) { // not required
		return nil
	}

	if o.Self != nil {
		if err := o.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "svm" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this svm peer inline svm inline links based on the context it is used
func (o *SvmPeerInlineSvmInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SvmPeerInlineSvmInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if o.Self != nil {
		if err := o.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "svm" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SvmPeerInlineSvmInlineLinks) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SvmPeerInlineSvmInlineLinks) UnmarshalBinary(b []byte) error {
	var res SvmPeerInlineSvmInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}