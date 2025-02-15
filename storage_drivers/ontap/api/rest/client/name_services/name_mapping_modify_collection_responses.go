// Code generated by go-swagger; DO NOT EDIT.

package name_services

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

// NameMappingModifyCollectionReader is a Reader for the NameMappingModifyCollection structure.
type NameMappingModifyCollectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NameMappingModifyCollectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNameMappingModifyCollectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNameMappingModifyCollectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNameMappingModifyCollectionOK creates a NameMappingModifyCollectionOK with default headers values
func NewNameMappingModifyCollectionOK() *NameMappingModifyCollectionOK {
	return &NameMappingModifyCollectionOK{}
}

/*
NameMappingModifyCollectionOK describes a response with status code 200, with default header values.

OK
*/
type NameMappingModifyCollectionOK struct {
}

// IsSuccess returns true when this name mapping modify collection o k response has a 2xx status code
func (o *NameMappingModifyCollectionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this name mapping modify collection o k response has a 3xx status code
func (o *NameMappingModifyCollectionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this name mapping modify collection o k response has a 4xx status code
func (o *NameMappingModifyCollectionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this name mapping modify collection o k response has a 5xx status code
func (o *NameMappingModifyCollectionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this name mapping modify collection o k response a status code equal to that given
func (o *NameMappingModifyCollectionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the name mapping modify collection o k response
func (o *NameMappingModifyCollectionOK) Code() int {
	return 200
}

func (o *NameMappingModifyCollectionOK) Error() string {
	return fmt.Sprintf("[PATCH /name-services/name-mappings][%d] nameMappingModifyCollectionOK", 200)
}

func (o *NameMappingModifyCollectionOK) String() string {
	return fmt.Sprintf("[PATCH /name-services/name-mappings][%d] nameMappingModifyCollectionOK", 200)
}

func (o *NameMappingModifyCollectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNameMappingModifyCollectionDefault creates a NameMappingModifyCollectionDefault with default headers values
func NewNameMappingModifyCollectionDefault(code int) *NameMappingModifyCollectionDefault {
	return &NameMappingModifyCollectionDefault{
		_statusCode: code,
	}
}

/*
	NameMappingModifyCollectionDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 65798185   | Failed to resolve the specified hostname |
| 65798179   | Cannot swap entries because one or both entries have host name or address configured.|
|            | Delete and re-create the new entry at the specified position.|
*/
type NameMappingModifyCollectionDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this name mapping modify collection default response has a 2xx status code
func (o *NameMappingModifyCollectionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this name mapping modify collection default response has a 3xx status code
func (o *NameMappingModifyCollectionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this name mapping modify collection default response has a 4xx status code
func (o *NameMappingModifyCollectionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this name mapping modify collection default response has a 5xx status code
func (o *NameMappingModifyCollectionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this name mapping modify collection default response a status code equal to that given
func (o *NameMappingModifyCollectionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the name mapping modify collection default response
func (o *NameMappingModifyCollectionDefault) Code() int {
	return o._statusCode
}

func (o *NameMappingModifyCollectionDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /name-services/name-mappings][%d] name_mapping_modify_collection default %s", o._statusCode, payload)
}

func (o *NameMappingModifyCollectionDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /name-services/name-mappings][%d] name_mapping_modify_collection default %s", o._statusCode, payload)
}

func (o *NameMappingModifyCollectionDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NameMappingModifyCollectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
NameMappingModifyCollectionBody name mapping modify collection body
swagger:model NameMappingModifyCollectionBody
*/
type NameMappingModifyCollectionBody struct {

	// links
	Links *models.NameMappingInlineLinks `json:"_links,omitempty"`

	// Client workstation IP Address which is matched when searching for the pattern.
	//   You can specify the value in any of the following formats:
	// * As an IPv4 address with a subnet mask expressed as a number of bits; for instance, 10.1.12.0/24
	// * As an IPv6 address with a subnet mask expressed as a number of bits; for instance, fd20:8b1e:b255:4071::/64
	// * As an IPv4 address with a network mask; for instance, 10.1.16.0/255.255.255.0
	// * As a hostname
	//
	// Example: 10.254.101.111/28
	ClientMatch *string `json:"client_match,omitempty"`

	// Direction in which the name mapping is applied. The possible values are:
	//   * krb_unix  - Kerberos principal name to UNIX user name
	//   * win_unix  - Windows user name to UNIX user name
	//   * unix_win  - UNIX user name to Windows user name mapping
	//   * s3_unix   - S3 user name to UNIX user name mapping
	//   * s3_win    - S3 user name to Windows user name mapping
	//
	// Example: win_unix
	// Enum: ["win_unix","unix_win","krb_unix","s3_unix","s3_win"]
	Direction *string `json:"direction,omitempty"`

	// Position in the list of name mappings.
	// Example: 1
	// Maximum: 2.147483647e+09
	// Minimum: 1
	Index *int64 `json:"index,omitempty"`

	// name mapping response inline records
	NameMappingResponseInlineRecords []*models.NameMapping `json:"records,omitempty"`

	// Pattern used to match the name while searching for a name that can be used as a replacement. The pattern is a UNIX-style regular expression. Regular expressions are case-insensitive when mapping from Windows to UNIX, and they are case-sensitive for mappings from Kerberos to UNIX and UNIX to Windows.
	// Example: ENGCIFS_AD_USER
	// Max Length: 256
	// Min Length: 1
	Pattern *string `json:"pattern,omitempty"`

	// The name that is used as a replacement, if the pattern associated with this entry matches.
	// Example: unix_user1
	// Max Length: 256
	// Min Length: 1
	Replacement *string `json:"replacement,omitempty"`

	// svm
	Svm *models.NameMappingInlineSvm `json:"svm,omitempty"`
}

// Validate validates this name mapping modify collection body
func (o *NameMappingModifyCollectionBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDirection(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateIndex(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNameMappingResponseInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePattern(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateReplacement(formats); err != nil {
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

func (o *NameMappingModifyCollectionBody) validateLinks(formats strfmt.Registry) error {
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

var nameMappingModifyCollectionBodyTypeDirectionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["win_unix","unix_win","krb_unix","s3_unix","s3_win"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		nameMappingModifyCollectionBodyTypeDirectionPropEnum = append(nameMappingModifyCollectionBodyTypeDirectionPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// NameMappingModifyCollectionBody
	// NameMappingModifyCollectionBody
	// direction
	// Direction
	// win_unix
	// END DEBUGGING
	// NameMappingModifyCollectionBodyDirectionWinUnix captures enum value "win_unix"
	NameMappingModifyCollectionBodyDirectionWinUnix string = "win_unix"

	// BEGIN DEBUGGING
	// NameMappingModifyCollectionBody
	// NameMappingModifyCollectionBody
	// direction
	// Direction
	// unix_win
	// END DEBUGGING
	// NameMappingModifyCollectionBodyDirectionUnixWin captures enum value "unix_win"
	NameMappingModifyCollectionBodyDirectionUnixWin string = "unix_win"

	// BEGIN DEBUGGING
	// NameMappingModifyCollectionBody
	// NameMappingModifyCollectionBody
	// direction
	// Direction
	// krb_unix
	// END DEBUGGING
	// NameMappingModifyCollectionBodyDirectionKrbUnix captures enum value "krb_unix"
	NameMappingModifyCollectionBodyDirectionKrbUnix string = "krb_unix"

	// BEGIN DEBUGGING
	// NameMappingModifyCollectionBody
	// NameMappingModifyCollectionBody
	// direction
	// Direction
	// s3_unix
	// END DEBUGGING
	// NameMappingModifyCollectionBodyDirectionS3Unix captures enum value "s3_unix"
	NameMappingModifyCollectionBodyDirectionS3Unix string = "s3_unix"

	// BEGIN DEBUGGING
	// NameMappingModifyCollectionBody
	// NameMappingModifyCollectionBody
	// direction
	// Direction
	// s3_win
	// END DEBUGGING
	// NameMappingModifyCollectionBodyDirectionS3Win captures enum value "s3_win"
	NameMappingModifyCollectionBodyDirectionS3Win string = "s3_win"
)

// prop value enum
func (o *NameMappingModifyCollectionBody) validateDirectionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, nameMappingModifyCollectionBodyTypeDirectionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *NameMappingModifyCollectionBody) validateDirection(formats strfmt.Registry) error {
	if swag.IsZero(o.Direction) { // not required
		return nil
	}

	// value enum
	if err := o.validateDirectionEnum("info"+"."+"direction", "body", *o.Direction); err != nil {
		return err
	}

	return nil
}

func (o *NameMappingModifyCollectionBody) validateIndex(formats strfmt.Registry) error {
	if swag.IsZero(o.Index) { // not required
		return nil
	}

	if err := validate.MinimumInt("info"+"."+"index", "body", *o.Index, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("info"+"."+"index", "body", *o.Index, 2.147483647e+09, false); err != nil {
		return err
	}

	return nil
}

func (o *NameMappingModifyCollectionBody) validateNameMappingResponseInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(o.NameMappingResponseInlineRecords) { // not required
		return nil
	}

	for i := 0; i < len(o.NameMappingResponseInlineRecords); i++ {
		if swag.IsZero(o.NameMappingResponseInlineRecords[i]) { // not required
			continue
		}

		if o.NameMappingResponseInlineRecords[i] != nil {
			if err := o.NameMappingResponseInlineRecords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *NameMappingModifyCollectionBody) validatePattern(formats strfmt.Registry) error {
	if swag.IsZero(o.Pattern) { // not required
		return nil
	}

	if err := validate.MinLength("info"+"."+"pattern", "body", *o.Pattern, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("info"+"."+"pattern", "body", *o.Pattern, 256); err != nil {
		return err
	}

	return nil
}

func (o *NameMappingModifyCollectionBody) validateReplacement(formats strfmt.Registry) error {
	if swag.IsZero(o.Replacement) { // not required
		return nil
	}

	if err := validate.MinLength("info"+"."+"replacement", "body", *o.Replacement, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("info"+"."+"replacement", "body", *o.Replacement, 256); err != nil {
		return err
	}

	return nil
}

func (o *NameMappingModifyCollectionBody) validateSvm(formats strfmt.Registry) error {
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

// ContextValidate validate this name mapping modify collection body based on the context it is used
func (o *NameMappingModifyCollectionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateNameMappingResponseInlineRecords(ctx, formats); err != nil {
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

func (o *NameMappingModifyCollectionBody) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (o *NameMappingModifyCollectionBody) contextValidateNameMappingResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.NameMappingResponseInlineRecords); i++ {

		if o.NameMappingResponseInlineRecords[i] != nil {
			if err := o.NameMappingResponseInlineRecords[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *NameMappingModifyCollectionBody) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

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
func (o *NameMappingModifyCollectionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NameMappingModifyCollectionBody) UnmarshalBinary(b []byte) error {
	var res NameMappingModifyCollectionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
NameMappingInlineLinks name mapping inline links
swagger:model name_mapping_inline__links
*/
type NameMappingInlineLinks struct {

	// self
	Self *models.Href `json:"self,omitempty"`
}

// Validate validates this name mapping inline links
func (o *NameMappingInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NameMappingInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this name mapping inline links based on the context it is used
func (o *NameMappingInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NameMappingInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (o *NameMappingInlineLinks) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NameMappingInlineLinks) UnmarshalBinary(b []byte) error {
	var res NameMappingInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
NameMappingInlineSvm SVM, applies only to SVM-scoped objects.
swagger:model name_mapping_inline_svm
*/
type NameMappingInlineSvm struct {

	// links
	Links *models.NameMappingInlineSvmInlineLinks `json:"_links,omitempty"`

	// The name of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: svm1
	Name *string `json:"name,omitempty"`

	// The unique identifier of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this name mapping inline svm
func (o *NameMappingInlineSvm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NameMappingInlineSvm) validateLinks(formats strfmt.Registry) error {
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

// ContextValidate validate this name mapping inline svm based on the context it is used
func (o *NameMappingInlineSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NameMappingInlineSvm) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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
func (o *NameMappingInlineSvm) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NameMappingInlineSvm) UnmarshalBinary(b []byte) error {
	var res NameMappingInlineSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
NameMappingInlineSvmInlineLinks name mapping inline svm inline links
swagger:model name_mapping_inline_svm_inline__links
*/
type NameMappingInlineSvmInlineLinks struct {

	// self
	Self *models.Href `json:"self,omitempty"`
}

// Validate validates this name mapping inline svm inline links
func (o *NameMappingInlineSvmInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NameMappingInlineSvmInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this name mapping inline svm inline links based on the context it is used
func (o *NameMappingInlineSvmInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NameMappingInlineSvmInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (o *NameMappingInlineSvmInlineLinks) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NameMappingInlineSvmInlineLinks) UnmarshalBinary(b []byte) error {
	var res NameMappingInlineSvmInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
