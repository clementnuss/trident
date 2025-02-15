// Code generated by go-swagger; DO NOT EDIT.

package s_a_n

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

// IscsiCredentialsModifyCollectionReader is a Reader for the IscsiCredentialsModifyCollection structure.
type IscsiCredentialsModifyCollectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IscsiCredentialsModifyCollectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIscsiCredentialsModifyCollectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIscsiCredentialsModifyCollectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIscsiCredentialsModifyCollectionOK creates a IscsiCredentialsModifyCollectionOK with default headers values
func NewIscsiCredentialsModifyCollectionOK() *IscsiCredentialsModifyCollectionOK {
	return &IscsiCredentialsModifyCollectionOK{}
}

/*
IscsiCredentialsModifyCollectionOK describes a response with status code 200, with default header values.

OK
*/
type IscsiCredentialsModifyCollectionOK struct {
}

// IsSuccess returns true when this iscsi credentials modify collection o k response has a 2xx status code
func (o *IscsiCredentialsModifyCollectionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this iscsi credentials modify collection o k response has a 3xx status code
func (o *IscsiCredentialsModifyCollectionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this iscsi credentials modify collection o k response has a 4xx status code
func (o *IscsiCredentialsModifyCollectionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this iscsi credentials modify collection o k response has a 5xx status code
func (o *IscsiCredentialsModifyCollectionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this iscsi credentials modify collection o k response a status code equal to that given
func (o *IscsiCredentialsModifyCollectionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the iscsi credentials modify collection o k response
func (o *IscsiCredentialsModifyCollectionOK) Code() int {
	return 200
}

func (o *IscsiCredentialsModifyCollectionOK) Error() string {
	return fmt.Sprintf("[PATCH /protocols/san/iscsi/credentials][%d] iscsiCredentialsModifyCollectionOK", 200)
}

func (o *IscsiCredentialsModifyCollectionOK) String() string {
	return fmt.Sprintf("[PATCH /protocols/san/iscsi/credentials][%d] iscsiCredentialsModifyCollectionOK", 200)
}

func (o *IscsiCredentialsModifyCollectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIscsiCredentialsModifyCollectionDefault creates a IscsiCredentialsModifyCollectionDefault with default headers values
func NewIscsiCredentialsModifyCollectionDefault(code int) *IscsiCredentialsModifyCollectionDefault {
	return &IscsiCredentialsModifyCollectionDefault{
		_statusCode: code,
	}
}

/*
	IscsiCredentialsModifyCollectionDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 2621462 | An SVM with the specified UUID does not exist. |
| 2621706 | Both the SVM UUID and SVM name were supplied, but they do not refer to the same SVM. |
| 2621707 | No SVM was specified. Either `svm.name` or `svm.uuid` must be supplied. |
| 5374145 | The iSCSI security password must contain an even number of valid hex digits. |
| 5374147 | The CHAP inbound and outbound passwords must be different. |
| 5374149 | The inbound user and password properties are required for CHAP authentication. |
| 5374150 | Outbound CHAP authentication requires an outbound password. |
| 5374155 | The functionality is not supported for the default security credential. |
| 5374855 | The value for property `initiator_address.ranges.start` is greater than the value for property `initiator_address.ranges.end`. |
| 5374856 | The value for property `initiator_address.ranges.start` does not belong to the same IP address family as the value for property `initiator_address.ranges.end`. |
| 5374895 | The iSCSI security credential does not exist on the specified SVM. |
| 5374900 | Setting the CHAP authentication properties are not supported with authentication types _none_ or _deny_. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type IscsiCredentialsModifyCollectionDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this iscsi credentials modify collection default response has a 2xx status code
func (o *IscsiCredentialsModifyCollectionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this iscsi credentials modify collection default response has a 3xx status code
func (o *IscsiCredentialsModifyCollectionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this iscsi credentials modify collection default response has a 4xx status code
func (o *IscsiCredentialsModifyCollectionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this iscsi credentials modify collection default response has a 5xx status code
func (o *IscsiCredentialsModifyCollectionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this iscsi credentials modify collection default response a status code equal to that given
func (o *IscsiCredentialsModifyCollectionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the iscsi credentials modify collection default response
func (o *IscsiCredentialsModifyCollectionDefault) Code() int {
	return o._statusCode
}

func (o *IscsiCredentialsModifyCollectionDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/san/iscsi/credentials][%d] iscsi_credentials_modify_collection default %s", o._statusCode, payload)
}

func (o *IscsiCredentialsModifyCollectionDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/san/iscsi/credentials][%d] iscsi_credentials_modify_collection default %s", o._statusCode, payload)
}

func (o *IscsiCredentialsModifyCollectionDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *IscsiCredentialsModifyCollectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
IscsiCredentialsModifyCollectionBody iscsi credentials modify collection body
swagger:model IscsiCredentialsModifyCollectionBody
*/
type IscsiCredentialsModifyCollectionBody struct {

	// links
	Links *models.IscsiCredentialsInlineLinks `json:"_links,omitempty"`

	// The iSCSI authentication type. Required in POST; optional in PATCH.
	//
	// Enum: ["chap","none","deny"]
	AuthenticationType *string `json:"authentication_type,omitempty"`

	// chap
	Chap *models.IscsiCredentialsInlineChap `json:"chap,omitempty"`

	// The iSCSI initiator to which the credentials apply. Required in POST.
	//
	// Example: iqn.1998-01.com.corp.iscsi:name1
	Initiator *string `json:"initiator,omitempty"`

	// initiator address
	InitiatorAddress *models.IscsiCredentialsInlineInitiatorAddress `json:"initiator_address,omitempty"`

	// iscsi credentials response inline records
	IscsiCredentialsResponseInlineRecords []*models.IscsiCredentials `json:"records,omitempty"`

	// svm
	Svm *models.IscsiCredentialsInlineSvm `json:"svm,omitempty"`
}

// Validate validates this iscsi credentials modify collection body
func (o *IscsiCredentialsModifyCollectionBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateAuthenticationType(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateChap(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateInitiatorAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateIscsiCredentialsResponseInlineRecords(formats); err != nil {
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

func (o *IscsiCredentialsModifyCollectionBody) validateLinks(formats strfmt.Registry) error {
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

var iscsiCredentialsModifyCollectionBodyTypeAuthenticationTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["chap","none","deny"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		iscsiCredentialsModifyCollectionBodyTypeAuthenticationTypePropEnum = append(iscsiCredentialsModifyCollectionBodyTypeAuthenticationTypePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// IscsiCredentialsModifyCollectionBody
	// IscsiCredentialsModifyCollectionBody
	// authentication_type
	// AuthenticationType
	// chap
	// END DEBUGGING
	// IscsiCredentialsModifyCollectionBodyAuthenticationTypeChap captures enum value "chap"
	IscsiCredentialsModifyCollectionBodyAuthenticationTypeChap string = "chap"

	// BEGIN DEBUGGING
	// IscsiCredentialsModifyCollectionBody
	// IscsiCredentialsModifyCollectionBody
	// authentication_type
	// AuthenticationType
	// none
	// END DEBUGGING
	// IscsiCredentialsModifyCollectionBodyAuthenticationTypeNone captures enum value "none"
	IscsiCredentialsModifyCollectionBodyAuthenticationTypeNone string = "none"

	// BEGIN DEBUGGING
	// IscsiCredentialsModifyCollectionBody
	// IscsiCredentialsModifyCollectionBody
	// authentication_type
	// AuthenticationType
	// deny
	// END DEBUGGING
	// IscsiCredentialsModifyCollectionBodyAuthenticationTypeDeny captures enum value "deny"
	IscsiCredentialsModifyCollectionBodyAuthenticationTypeDeny string = "deny"
)

// prop value enum
func (o *IscsiCredentialsModifyCollectionBody) validateAuthenticationTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, iscsiCredentialsModifyCollectionBodyTypeAuthenticationTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *IscsiCredentialsModifyCollectionBody) validateAuthenticationType(formats strfmt.Registry) error {
	if swag.IsZero(o.AuthenticationType) { // not required
		return nil
	}

	// value enum
	if err := o.validateAuthenticationTypeEnum("info"+"."+"authentication_type", "body", *o.AuthenticationType); err != nil {
		return err
	}

	return nil
}

func (o *IscsiCredentialsModifyCollectionBody) validateChap(formats strfmt.Registry) error {
	if swag.IsZero(o.Chap) { // not required
		return nil
	}

	if o.Chap != nil {
		if err := o.Chap.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "chap")
			}
			return err
		}
	}

	return nil
}

func (o *IscsiCredentialsModifyCollectionBody) validateInitiatorAddress(formats strfmt.Registry) error {
	if swag.IsZero(o.InitiatorAddress) { // not required
		return nil
	}

	if o.InitiatorAddress != nil {
		if err := o.InitiatorAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "initiator_address")
			}
			return err
		}
	}

	return nil
}

func (o *IscsiCredentialsModifyCollectionBody) validateIscsiCredentialsResponseInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(o.IscsiCredentialsResponseInlineRecords) { // not required
		return nil
	}

	for i := 0; i < len(o.IscsiCredentialsResponseInlineRecords); i++ {
		if swag.IsZero(o.IscsiCredentialsResponseInlineRecords[i]) { // not required
			continue
		}

		if o.IscsiCredentialsResponseInlineRecords[i] != nil {
			if err := o.IscsiCredentialsResponseInlineRecords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *IscsiCredentialsModifyCollectionBody) validateSvm(formats strfmt.Registry) error {
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

// ContextValidate validate this iscsi credentials modify collection body based on the context it is used
func (o *IscsiCredentialsModifyCollectionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateChap(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateInitiatorAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateIscsiCredentialsResponseInlineRecords(ctx, formats); err != nil {
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

func (o *IscsiCredentialsModifyCollectionBody) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (o *IscsiCredentialsModifyCollectionBody) contextValidateChap(ctx context.Context, formats strfmt.Registry) error {

	if o.Chap != nil {
		if err := o.Chap.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "chap")
			}
			return err
		}
	}

	return nil
}

func (o *IscsiCredentialsModifyCollectionBody) contextValidateInitiatorAddress(ctx context.Context, formats strfmt.Registry) error {

	if o.InitiatorAddress != nil {
		if err := o.InitiatorAddress.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "initiator_address")
			}
			return err
		}
	}

	return nil
}

func (o *IscsiCredentialsModifyCollectionBody) contextValidateIscsiCredentialsResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.IscsiCredentialsResponseInlineRecords); i++ {

		if o.IscsiCredentialsResponseInlineRecords[i] != nil {
			if err := o.IscsiCredentialsResponseInlineRecords[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *IscsiCredentialsModifyCollectionBody) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (o *IscsiCredentialsModifyCollectionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *IscsiCredentialsModifyCollectionBody) UnmarshalBinary(b []byte) error {
	var res IscsiCredentialsModifyCollectionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
IscsiCredentialsInlineLinks iscsi credentials inline links
swagger:model iscsi_credentials_inline__links
*/
type IscsiCredentialsInlineLinks struct {

	// self
	Self *models.Href `json:"self,omitempty"`
}

// Validate validates this iscsi credentials inline links
func (o *IscsiCredentialsInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *IscsiCredentialsInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this iscsi credentials inline links based on the context it is used
func (o *IscsiCredentialsInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *IscsiCredentialsInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (o *IscsiCredentialsInlineLinks) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *IscsiCredentialsInlineLinks) UnmarshalBinary(b []byte) error {
	var res IscsiCredentialsInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
IscsiCredentialsInlineChap Challenge-Handshake Authentication Protocol (CHAP) credentials.
//
swagger:model iscsi_credentials_inline_chap
*/
type IscsiCredentialsInlineChap struct {

	// inbound
	Inbound *models.IscsiCredentialsInlineChapInlineInbound `json:"inbound,omitempty"`

	// outbound
	Outbound *models.IscsiCredentialsInlineChapInlineOutbound `json:"outbound,omitempty"`
}

// Validate validates this iscsi credentials inline chap
func (o *IscsiCredentialsInlineChap) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateInbound(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateOutbound(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *IscsiCredentialsInlineChap) validateInbound(formats strfmt.Registry) error {
	if swag.IsZero(o.Inbound) { // not required
		return nil
	}

	if o.Inbound != nil {
		if err := o.Inbound.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "chap" + "." + "inbound")
			}
			return err
		}
	}

	return nil
}

func (o *IscsiCredentialsInlineChap) validateOutbound(formats strfmt.Registry) error {
	if swag.IsZero(o.Outbound) { // not required
		return nil
	}

	if o.Outbound != nil {
		if err := o.Outbound.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "chap" + "." + "outbound")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this iscsi credentials inline chap based on the context it is used
func (o *IscsiCredentialsInlineChap) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateInbound(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateOutbound(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *IscsiCredentialsInlineChap) contextValidateInbound(ctx context.Context, formats strfmt.Registry) error {

	if o.Inbound != nil {
		if err := o.Inbound.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "chap" + "." + "inbound")
			}
			return err
		}
	}

	return nil
}

func (o *IscsiCredentialsInlineChap) contextValidateOutbound(ctx context.Context, formats strfmt.Registry) error {

	if o.Outbound != nil {
		if err := o.Outbound.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "chap" + "." + "outbound")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *IscsiCredentialsInlineChap) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *IscsiCredentialsInlineChap) UnmarshalBinary(b []byte) error {
	var res IscsiCredentialsInlineChap
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
IscsiCredentialsInlineChapInlineInbound Inbound CHAP credentials.
//
swagger:model iscsi_credentials_inline_chap_inline_inbound
*/
type IscsiCredentialsInlineChapInlineInbound struct {

	// The inbound CHAP password. Write-only; optional in POST and PATCH.
	//
	// Max Length: 512
	// Min Length: 1
	Password *string `json:"password,omitempty"`

	// The inbound CHAP user name. Optional in POST and PATCH.
	//
	// Max Length: 128
	// Min Length: 1
	User *string `json:"user,omitempty"`
}

// Validate validates this iscsi credentials inline chap inline inbound
func (o *IscsiCredentialsInlineChapInlineInbound) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *IscsiCredentialsInlineChapInlineInbound) validatePassword(formats strfmt.Registry) error {
	if swag.IsZero(o.Password) { // not required
		return nil
	}

	if err := validate.MinLength("info"+"."+"chap"+"."+"inbound"+"."+"password", "body", *o.Password, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("info"+"."+"chap"+"."+"inbound"+"."+"password", "body", *o.Password, 512); err != nil {
		return err
	}

	return nil
}

func (o *IscsiCredentialsInlineChapInlineInbound) validateUser(formats strfmt.Registry) error {
	if swag.IsZero(o.User) { // not required
		return nil
	}

	if err := validate.MinLength("info"+"."+"chap"+"."+"inbound"+"."+"user", "body", *o.User, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("info"+"."+"chap"+"."+"inbound"+"."+"user", "body", *o.User, 128); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this iscsi credentials inline chap inline inbound based on context it is used
func (o *IscsiCredentialsInlineChapInlineInbound) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *IscsiCredentialsInlineChapInlineInbound) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *IscsiCredentialsInlineChapInlineInbound) UnmarshalBinary(b []byte) error {
	var res IscsiCredentialsInlineChapInlineInbound
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
IscsiCredentialsInlineChapInlineOutbound Output CHAP credentials.</br>
// To clear previously set outbound CHAP credentials, set property `chap.outbound.user` to an empty string in PATCH.
//
swagger:model iscsi_credentials_inline_chap_inline_outbound
*/
type IscsiCredentialsInlineChapInlineOutbound struct {

	// The outbound CHAP password. Write-only; optional in POST and PATCH.
	//
	// Max Length: 512
	// Min Length: 1
	Password *string `json:"password,omitempty"`

	// The outbound CHAP user name. Optional in POST and PATCH.</br>
	// To clear previously set outbound CHAP credentials, set this property to an empty string in PATCH.
	//
	// Max Length: 128
	// Min Length: 1
	User *string `json:"user,omitempty"`
}

// Validate validates this iscsi credentials inline chap inline outbound
func (o *IscsiCredentialsInlineChapInlineOutbound) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *IscsiCredentialsInlineChapInlineOutbound) validatePassword(formats strfmt.Registry) error {
	if swag.IsZero(o.Password) { // not required
		return nil
	}

	if err := validate.MinLength("info"+"."+"chap"+"."+"outbound"+"."+"password", "body", *o.Password, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("info"+"."+"chap"+"."+"outbound"+"."+"password", "body", *o.Password, 512); err != nil {
		return err
	}

	return nil
}

func (o *IscsiCredentialsInlineChapInlineOutbound) validateUser(formats strfmt.Registry) error {
	if swag.IsZero(o.User) { // not required
		return nil
	}

	if err := validate.MinLength("info"+"."+"chap"+"."+"outbound"+"."+"user", "body", *o.User, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("info"+"."+"chap"+"."+"outbound"+"."+"user", "body", *o.User, 128); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this iscsi credentials inline chap inline outbound based on context it is used
func (o *IscsiCredentialsInlineChapInlineOutbound) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *IscsiCredentialsInlineChapInlineOutbound) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *IscsiCredentialsInlineChapInlineOutbound) UnmarshalBinary(b []byte) error {
	var res IscsiCredentialsInlineChapInlineOutbound
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
IscsiCredentialsInlineInitiatorAddress Initiator address ranges.
//
swagger:model iscsi_credentials_inline_initiator_address
*/
type IscsiCredentialsInlineInitiatorAddress struct {

	// masks
	Masks []*models.IPInfo `json:"masks,omitempty"`

	// ranges
	Ranges []*models.IPAddressRange `json:"ranges,omitempty"`
}

// Validate validates this iscsi credentials inline initiator address
func (o *IscsiCredentialsInlineInitiatorAddress) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMasks(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRanges(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *IscsiCredentialsInlineInitiatorAddress) validateMasks(formats strfmt.Registry) error {
	if swag.IsZero(o.Masks) { // not required
		return nil
	}

	for i := 0; i < len(o.Masks); i++ {
		if swag.IsZero(o.Masks[i]) { // not required
			continue
		}

		if o.Masks[i] != nil {
			if err := o.Masks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "initiator_address" + "." + "masks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *IscsiCredentialsInlineInitiatorAddress) validateRanges(formats strfmt.Registry) error {
	if swag.IsZero(o.Ranges) { // not required
		return nil
	}

	for i := 0; i < len(o.Ranges); i++ {
		if swag.IsZero(o.Ranges[i]) { // not required
			continue
		}

		if o.Ranges[i] != nil {
			if err := o.Ranges[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "initiator_address" + "." + "ranges" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this iscsi credentials inline initiator address based on the context it is used
func (o *IscsiCredentialsInlineInitiatorAddress) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateMasks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateRanges(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *IscsiCredentialsInlineInitiatorAddress) contextValidateMasks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Masks); i++ {

		if o.Masks[i] != nil {
			if err := o.Masks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "initiator_address" + "." + "masks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *IscsiCredentialsInlineInitiatorAddress) contextValidateRanges(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Ranges); i++ {

		if o.Ranges[i] != nil {
			if err := o.Ranges[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "initiator_address" + "." + "ranges" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *IscsiCredentialsInlineInitiatorAddress) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *IscsiCredentialsInlineInitiatorAddress) UnmarshalBinary(b []byte) error {
	var res IscsiCredentialsInlineInitiatorAddress
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
IscsiCredentialsInlineSvm SVM, applies only to SVM-scoped objects.
swagger:model iscsi_credentials_inline_svm
*/
type IscsiCredentialsInlineSvm struct {

	// links
	Links *models.IscsiCredentialsInlineSvmInlineLinks `json:"_links,omitempty"`

	// The name of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: svm1
	Name *string `json:"name,omitempty"`

	// The unique identifier of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this iscsi credentials inline svm
func (o *IscsiCredentialsInlineSvm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *IscsiCredentialsInlineSvm) validateLinks(formats strfmt.Registry) error {
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

// ContextValidate validate this iscsi credentials inline svm based on the context it is used
func (o *IscsiCredentialsInlineSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *IscsiCredentialsInlineSvm) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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
func (o *IscsiCredentialsInlineSvm) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *IscsiCredentialsInlineSvm) UnmarshalBinary(b []byte) error {
	var res IscsiCredentialsInlineSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
IscsiCredentialsInlineSvmInlineLinks iscsi credentials inline svm inline links
swagger:model iscsi_credentials_inline_svm_inline__links
*/
type IscsiCredentialsInlineSvmInlineLinks struct {

	// self
	Self *models.Href `json:"self,omitempty"`
}

// Validate validates this iscsi credentials inline svm inline links
func (o *IscsiCredentialsInlineSvmInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *IscsiCredentialsInlineSvmInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this iscsi credentials inline svm inline links based on the context it is used
func (o *IscsiCredentialsInlineSvmInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *IscsiCredentialsInlineSvmInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (o *IscsiCredentialsInlineSvmInlineLinks) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *IscsiCredentialsInlineSvmInlineLinks) UnmarshalBinary(b []byte) error {
	var res IscsiCredentialsInlineSvmInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
