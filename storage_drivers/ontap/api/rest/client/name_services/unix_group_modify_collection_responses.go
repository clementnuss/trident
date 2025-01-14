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

// UnixGroupModifyCollectionReader is a Reader for the UnixGroupModifyCollection structure.
type UnixGroupModifyCollectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UnixGroupModifyCollectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUnixGroupModifyCollectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUnixGroupModifyCollectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUnixGroupModifyCollectionOK creates a UnixGroupModifyCollectionOK with default headers values
func NewUnixGroupModifyCollectionOK() *UnixGroupModifyCollectionOK {
	return &UnixGroupModifyCollectionOK{}
}

/*
UnixGroupModifyCollectionOK describes a response with status code 200, with default header values.

OK
*/
type UnixGroupModifyCollectionOK struct {
}

// IsSuccess returns true when this unix group modify collection o k response has a 2xx status code
func (o *UnixGroupModifyCollectionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this unix group modify collection o k response has a 3xx status code
func (o *UnixGroupModifyCollectionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this unix group modify collection o k response has a 4xx status code
func (o *UnixGroupModifyCollectionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this unix group modify collection o k response has a 5xx status code
func (o *UnixGroupModifyCollectionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this unix group modify collection o k response a status code equal to that given
func (o *UnixGroupModifyCollectionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the unix group modify collection o k response
func (o *UnixGroupModifyCollectionOK) Code() int {
	return 200
}

func (o *UnixGroupModifyCollectionOK) Error() string {
	return fmt.Sprintf("[PATCH /name-services/unix-groups][%d] unixGroupModifyCollectionOK", 200)
}

func (o *UnixGroupModifyCollectionOK) String() string {
	return fmt.Sprintf("[PATCH /name-services/unix-groups][%d] unixGroupModifyCollectionOK", 200)
}

func (o *UnixGroupModifyCollectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUnixGroupModifyCollectionDefault creates a UnixGroupModifyCollectionDefault with default headers values
func NewUnixGroupModifyCollectionDefault(code int) *UnixGroupModifyCollectionDefault {
	return &UnixGroupModifyCollectionDefault{
		_statusCode: code,
	}
}

/*
	UnixGroupModifyCollectionDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 23724141   | Duplicate group ID. Group ID must be unique.|
*/
type UnixGroupModifyCollectionDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this unix group modify collection default response has a 2xx status code
func (o *UnixGroupModifyCollectionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this unix group modify collection default response has a 3xx status code
func (o *UnixGroupModifyCollectionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this unix group modify collection default response has a 4xx status code
func (o *UnixGroupModifyCollectionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this unix group modify collection default response has a 5xx status code
func (o *UnixGroupModifyCollectionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this unix group modify collection default response a status code equal to that given
func (o *UnixGroupModifyCollectionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the unix group modify collection default response
func (o *UnixGroupModifyCollectionDefault) Code() int {
	return o._statusCode
}

func (o *UnixGroupModifyCollectionDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /name-services/unix-groups][%d] unix_group_modify_collection default %s", o._statusCode, payload)
}

func (o *UnixGroupModifyCollectionDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /name-services/unix-groups][%d] unix_group_modify_collection default %s", o._statusCode, payload)
}

func (o *UnixGroupModifyCollectionDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UnixGroupModifyCollectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UnixGroupModifyCollectionBody unix group modify collection body
swagger:model UnixGroupModifyCollectionBody
*/
type UnixGroupModifyCollectionBody struct {

	// links
	Links *models.UnixGroupInlineLinks `json:"_links,omitempty"`

	// Number of UNIX groups and members.
	CurrentCount *int64 `json:"current_count,omitempty"`

	// UNIX group ID of the specified user.
	//
	ID *int64 `json:"id,omitempty"`

	// UNIX group name to be added to the local database.
	//
	// Example: group1
	Name *string `json:"name,omitempty"`

	// Indicates whether or not the validation for the specified UNIX group name is disabled.
	SkipNameValidation *bool `json:"skip_name_validation,omitempty"`

	// svm
	Svm *models.UnixGroupInlineSvm `json:"svm,omitempty"`

	// unix group inline users
	// Read Only: true
	UnixGroupInlineUsers []*models.UnixGroupInlineUsersInlineArrayItem `json:"users,omitempty"`

	// unix group response inline records
	UnixGroupResponseInlineRecords []*models.UnixGroup `json:"records,omitempty"`
}

// Validate validates this unix group modify collection body
func (o *UnixGroupModifyCollectionBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSvm(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUnixGroupInlineUsers(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUnixGroupResponseInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UnixGroupModifyCollectionBody) validateLinks(formats strfmt.Registry) error {
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

func (o *UnixGroupModifyCollectionBody) validateSvm(formats strfmt.Registry) error {
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

func (o *UnixGroupModifyCollectionBody) validateUnixGroupInlineUsers(formats strfmt.Registry) error {
	if swag.IsZero(o.UnixGroupInlineUsers) { // not required
		return nil
	}

	for i := 0; i < len(o.UnixGroupInlineUsers); i++ {
		if swag.IsZero(o.UnixGroupInlineUsers[i]) { // not required
			continue
		}

		if o.UnixGroupInlineUsers[i] != nil {
			if err := o.UnixGroupInlineUsers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "users" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *UnixGroupModifyCollectionBody) validateUnixGroupResponseInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(o.UnixGroupResponseInlineRecords) { // not required
		return nil
	}

	for i := 0; i < len(o.UnixGroupResponseInlineRecords); i++ {
		if swag.IsZero(o.UnixGroupResponseInlineRecords[i]) { // not required
			continue
		}

		if o.UnixGroupResponseInlineRecords[i] != nil {
			if err := o.UnixGroupResponseInlineRecords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this unix group modify collection body based on the context it is used
func (o *UnixGroupModifyCollectionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateSvm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateUnixGroupInlineUsers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateUnixGroupResponseInlineRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UnixGroupModifyCollectionBody) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (o *UnixGroupModifyCollectionBody) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

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

func (o *UnixGroupModifyCollectionBody) contextValidateUnixGroupInlineUsers(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "info"+"."+"users", "body", []*models.UnixGroupInlineUsersInlineArrayItem(o.UnixGroupInlineUsers)); err != nil {
		return err
	}

	for i := 0; i < len(o.UnixGroupInlineUsers); i++ {

		if o.UnixGroupInlineUsers[i] != nil {
			if err := o.UnixGroupInlineUsers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "users" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *UnixGroupModifyCollectionBody) contextValidateUnixGroupResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.UnixGroupResponseInlineRecords); i++ {

		if o.UnixGroupResponseInlineRecords[i] != nil {
			if err := o.UnixGroupResponseInlineRecords[i].ContextValidate(ctx, formats); err != nil {
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
func (o *UnixGroupModifyCollectionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UnixGroupModifyCollectionBody) UnmarshalBinary(b []byte) error {
	var res UnixGroupModifyCollectionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UnixGroupInlineLinks unix group inline links
swagger:model unix_group_inline__links
*/
type UnixGroupInlineLinks struct {

	// self
	Self *models.Href `json:"self,omitempty"`
}

// Validate validates this unix group inline links
func (o *UnixGroupInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UnixGroupInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this unix group inline links based on the context it is used
func (o *UnixGroupInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UnixGroupInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (o *UnixGroupInlineLinks) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UnixGroupInlineLinks) UnmarshalBinary(b []byte) error {
	var res UnixGroupInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UnixGroupInlineSvm SVM, applies only to SVM-scoped objects.
swagger:model unix_group_inline_svm
*/
type UnixGroupInlineSvm struct {

	// links
	Links *models.UnixGroupInlineSvmInlineLinks `json:"_links,omitempty"`

	// The name of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: svm1
	Name *string `json:"name,omitempty"`

	// The unique identifier of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this unix group inline svm
func (o *UnixGroupInlineSvm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UnixGroupInlineSvm) validateLinks(formats strfmt.Registry) error {
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

// ContextValidate validate this unix group inline svm based on the context it is used
func (o *UnixGroupInlineSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UnixGroupInlineSvm) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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
func (o *UnixGroupInlineSvm) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UnixGroupInlineSvm) UnmarshalBinary(b []byte) error {
	var res UnixGroupInlineSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UnixGroupInlineSvmInlineLinks unix group inline svm inline links
swagger:model unix_group_inline_svm_inline__links
*/
type UnixGroupInlineSvmInlineLinks struct {

	// self
	Self *models.Href `json:"self,omitempty"`
}

// Validate validates this unix group inline svm inline links
func (o *UnixGroupInlineSvmInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UnixGroupInlineSvmInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this unix group inline svm inline links based on the context it is used
func (o *UnixGroupInlineSvmInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UnixGroupInlineSvmInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (o *UnixGroupInlineSvmInlineLinks) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UnixGroupInlineSvmInlineLinks) UnmarshalBinary(b []byte) error {
	var res UnixGroupInlineSvmInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UnixGroupInlineUsersInlineArrayItem unix group inline users inline array item
swagger:model unix_group_inline_users_inline_array_item
*/
type UnixGroupInlineUsersInlineArrayItem struct {

	// UNIX user who belongs to the specified UNIX group and the SVM.
	//
	Name *string `json:"name,omitempty"`
}

// Validate validates this unix group inline users inline array item
func (o *UnixGroupInlineUsersInlineArrayItem) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this unix group inline users inline array item based on context it is used
func (o *UnixGroupInlineUsersInlineArrayItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UnixGroupInlineUsersInlineArrayItem) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UnixGroupInlineUsersInlineArrayItem) UnmarshalBinary(b []byte) error {
	var res UnixGroupInlineUsersInlineArrayItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
