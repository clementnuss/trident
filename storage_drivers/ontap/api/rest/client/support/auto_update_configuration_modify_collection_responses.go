// Code generated by go-swagger; DO NOT EDIT.

package support

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

// AutoUpdateConfigurationModifyCollectionReader is a Reader for the AutoUpdateConfigurationModifyCollection structure.
type AutoUpdateConfigurationModifyCollectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AutoUpdateConfigurationModifyCollectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAutoUpdateConfigurationModifyCollectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAutoUpdateConfigurationModifyCollectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAutoUpdateConfigurationModifyCollectionOK creates a AutoUpdateConfigurationModifyCollectionOK with default headers values
func NewAutoUpdateConfigurationModifyCollectionOK() *AutoUpdateConfigurationModifyCollectionOK {
	return &AutoUpdateConfigurationModifyCollectionOK{}
}

/*
AutoUpdateConfigurationModifyCollectionOK describes a response with status code 200, with default header values.

OK
*/
type AutoUpdateConfigurationModifyCollectionOK struct {
}

// IsSuccess returns true when this auto update configuration modify collection o k response has a 2xx status code
func (o *AutoUpdateConfigurationModifyCollectionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this auto update configuration modify collection o k response has a 3xx status code
func (o *AutoUpdateConfigurationModifyCollectionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this auto update configuration modify collection o k response has a 4xx status code
func (o *AutoUpdateConfigurationModifyCollectionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this auto update configuration modify collection o k response has a 5xx status code
func (o *AutoUpdateConfigurationModifyCollectionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this auto update configuration modify collection o k response a status code equal to that given
func (o *AutoUpdateConfigurationModifyCollectionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the auto update configuration modify collection o k response
func (o *AutoUpdateConfigurationModifyCollectionOK) Code() int {
	return 200
}

func (o *AutoUpdateConfigurationModifyCollectionOK) Error() string {
	return fmt.Sprintf("[PATCH /support/auto-update/configurations][%d] autoUpdateConfigurationModifyCollectionOK", 200)
}

func (o *AutoUpdateConfigurationModifyCollectionOK) String() string {
	return fmt.Sprintf("[PATCH /support/auto-update/configurations][%d] autoUpdateConfigurationModifyCollectionOK", 200)
}

func (o *AutoUpdateConfigurationModifyCollectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAutoUpdateConfigurationModifyCollectionDefault creates a AutoUpdateConfigurationModifyCollectionDefault with default headers values
func NewAutoUpdateConfigurationModifyCollectionDefault(code int) *AutoUpdateConfigurationModifyCollectionDefault {
	return &AutoUpdateConfigurationModifyCollectionDefault{
		_statusCode: code,
	}
}

/*
	AutoUpdateConfigurationModifyCollectionDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 262179 | Unexpected argument. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type AutoUpdateConfigurationModifyCollectionDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this auto update configuration modify collection default response has a 2xx status code
func (o *AutoUpdateConfigurationModifyCollectionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this auto update configuration modify collection default response has a 3xx status code
func (o *AutoUpdateConfigurationModifyCollectionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this auto update configuration modify collection default response has a 4xx status code
func (o *AutoUpdateConfigurationModifyCollectionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this auto update configuration modify collection default response has a 5xx status code
func (o *AutoUpdateConfigurationModifyCollectionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this auto update configuration modify collection default response a status code equal to that given
func (o *AutoUpdateConfigurationModifyCollectionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the auto update configuration modify collection default response
func (o *AutoUpdateConfigurationModifyCollectionDefault) Code() int {
	return o._statusCode
}

func (o *AutoUpdateConfigurationModifyCollectionDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /support/auto-update/configurations][%d] auto_update_configuration_modify_collection default %s", o._statusCode, payload)
}

func (o *AutoUpdateConfigurationModifyCollectionDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /support/auto-update/configurations][%d] auto_update_configuration_modify_collection default %s", o._statusCode, payload)
}

func (o *AutoUpdateConfigurationModifyCollectionDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AutoUpdateConfigurationModifyCollectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
AutoUpdateConfigurationModifyCollectionBody auto update configuration modify collection body
swagger:model AutoUpdateConfigurationModifyCollectionBody
*/
type AutoUpdateConfigurationModifyCollectionBody struct {

	// links
	Links *models.AutoUpdateConfigurationInlineLinks `json:"_links,omitempty"`

	// The action to be taken by the alert source as specified by the user.
	// Example: confirm
	// Enum: ["confirm","dismiss","automatic"]
	Action *string `json:"action,omitempty"`

	// auto update configuration response inline records
	AutoUpdateConfigurationResponseInlineRecords []*models.AutoUpdateConfiguration `json:"records,omitempty"`

	// Category for the configuration row.
	// Example: disk_fw
	// Read Only: true
	Category *string `json:"category,omitempty"`

	// description
	Description *models.AutoUpdateConfigurationInlineDescription `json:"description,omitempty"`

	// Unique identifier for the configuration row.
	// Example: 572361f3-e769-439d-9c04-2ba48a08ff47
	// Read Only: true
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this auto update configuration modify collection body
func (o *AutoUpdateConfigurationModifyCollectionBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateAutoUpdateConfigurationResponseInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AutoUpdateConfigurationModifyCollectionBody) validateLinks(formats strfmt.Registry) error {
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

var autoUpdateConfigurationModifyCollectionBodyTypeActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["confirm","dismiss","automatic"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		autoUpdateConfigurationModifyCollectionBodyTypeActionPropEnum = append(autoUpdateConfigurationModifyCollectionBodyTypeActionPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// AutoUpdateConfigurationModifyCollectionBody
	// AutoUpdateConfigurationModifyCollectionBody
	// action
	// Action
	// confirm
	// END DEBUGGING
	// AutoUpdateConfigurationModifyCollectionBodyActionConfirm captures enum value "confirm"
	AutoUpdateConfigurationModifyCollectionBodyActionConfirm string = "confirm"

	// BEGIN DEBUGGING
	// AutoUpdateConfigurationModifyCollectionBody
	// AutoUpdateConfigurationModifyCollectionBody
	// action
	// Action
	// dismiss
	// END DEBUGGING
	// AutoUpdateConfigurationModifyCollectionBodyActionDismiss captures enum value "dismiss"
	AutoUpdateConfigurationModifyCollectionBodyActionDismiss string = "dismiss"

	// BEGIN DEBUGGING
	// AutoUpdateConfigurationModifyCollectionBody
	// AutoUpdateConfigurationModifyCollectionBody
	// action
	// Action
	// automatic
	// END DEBUGGING
	// AutoUpdateConfigurationModifyCollectionBodyActionAutomatic captures enum value "automatic"
	AutoUpdateConfigurationModifyCollectionBodyActionAutomatic string = "automatic"
)

// prop value enum
func (o *AutoUpdateConfigurationModifyCollectionBody) validateActionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, autoUpdateConfigurationModifyCollectionBodyTypeActionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *AutoUpdateConfigurationModifyCollectionBody) validateAction(formats strfmt.Registry) error {
	if swag.IsZero(o.Action) { // not required
		return nil
	}

	// value enum
	if err := o.validateActionEnum("info"+"."+"action", "body", *o.Action); err != nil {
		return err
	}

	return nil
}

func (o *AutoUpdateConfigurationModifyCollectionBody) validateAutoUpdateConfigurationResponseInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(o.AutoUpdateConfigurationResponseInlineRecords) { // not required
		return nil
	}

	for i := 0; i < len(o.AutoUpdateConfigurationResponseInlineRecords); i++ {
		if swag.IsZero(o.AutoUpdateConfigurationResponseInlineRecords[i]) { // not required
			continue
		}

		if o.AutoUpdateConfigurationResponseInlineRecords[i] != nil {
			if err := o.AutoUpdateConfigurationResponseInlineRecords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *AutoUpdateConfigurationModifyCollectionBody) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(o.Description) { // not required
		return nil
	}

	if o.Description != nil {
		if err := o.Description.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "description")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this auto update configuration modify collection body based on the context it is used
func (o *AutoUpdateConfigurationModifyCollectionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateAutoUpdateConfigurationResponseInlineRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateCategory(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateDescription(ctx, formats); err != nil {
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

func (o *AutoUpdateConfigurationModifyCollectionBody) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (o *AutoUpdateConfigurationModifyCollectionBody) contextValidateAutoUpdateConfigurationResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.AutoUpdateConfigurationResponseInlineRecords); i++ {

		if o.AutoUpdateConfigurationResponseInlineRecords[i] != nil {
			if err := o.AutoUpdateConfigurationResponseInlineRecords[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *AutoUpdateConfigurationModifyCollectionBody) contextValidateCategory(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "info"+"."+"category", "body", o.Category); err != nil {
		return err
	}

	return nil
}

func (o *AutoUpdateConfigurationModifyCollectionBody) contextValidateDescription(ctx context.Context, formats strfmt.Registry) error {

	if o.Description != nil {
		if err := o.Description.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "description")
			}
			return err
		}
	}

	return nil
}

func (o *AutoUpdateConfigurationModifyCollectionBody) contextValidateUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "info"+"."+"uuid", "body", o.UUID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AutoUpdateConfigurationModifyCollectionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AutoUpdateConfigurationModifyCollectionBody) UnmarshalBinary(b []byte) error {
	var res AutoUpdateConfigurationModifyCollectionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AutoUpdateConfigurationInlineLinks auto update configuration inline links
swagger:model auto_update_configuration_inline__links
*/
type AutoUpdateConfigurationInlineLinks struct {

	// self
	Self *models.Href `json:"self,omitempty"`
}

// Validate validates this auto update configuration inline links
func (o *AutoUpdateConfigurationInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AutoUpdateConfigurationInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this auto update configuration inline links based on the context it is used
func (o *AutoUpdateConfigurationInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AutoUpdateConfigurationInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (o *AutoUpdateConfigurationInlineLinks) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AutoUpdateConfigurationInlineLinks) UnmarshalBinary(b []byte) error {
	var res AutoUpdateConfigurationInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AutoUpdateConfigurationInlineDescription Description of the configuration row.
// ONTAP Message Codes
// | Code       | Description |
// | ---------- | ----------- |
// | 131072401 | Storage Firmware |
// | 131072402 | SP/BMC Firmware |
// | 131072403 | System Files |
//
swagger:model auto_update_configuration_inline_description
*/
type AutoUpdateConfigurationInlineDescription struct {

	// Argument code
	// Read Only: true
	Code *string `json:"code,omitempty"`

	// Message argument
	// Read Only: true
	Message *string `json:"message,omitempty"`
}

// Validate validates this auto update configuration inline description
func (o *AutoUpdateConfigurationInlineDescription) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this auto update configuration inline description based on the context it is used
func (o *AutoUpdateConfigurationInlineDescription) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateCode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateMessage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AutoUpdateConfigurationInlineDescription) contextValidateCode(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "info"+"."+"description"+"."+"code", "body", o.Code); err != nil {
		return err
	}

	return nil
}

func (o *AutoUpdateConfigurationInlineDescription) contextValidateMessage(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "info"+"."+"description"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AutoUpdateConfigurationInlineDescription) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AutoUpdateConfigurationInlineDescription) UnmarshalBinary(b []byte) error {
	var res AutoUpdateConfigurationInlineDescription
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}