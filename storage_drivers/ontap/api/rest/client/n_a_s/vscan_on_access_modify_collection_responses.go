// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

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

// VscanOnAccessModifyCollectionReader is a Reader for the VscanOnAccessModifyCollection structure.
type VscanOnAccessModifyCollectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VscanOnAccessModifyCollectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVscanOnAccessModifyCollectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVscanOnAccessModifyCollectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVscanOnAccessModifyCollectionOK creates a VscanOnAccessModifyCollectionOK with default headers values
func NewVscanOnAccessModifyCollectionOK() *VscanOnAccessModifyCollectionOK {
	return &VscanOnAccessModifyCollectionOK{}
}

/*
VscanOnAccessModifyCollectionOK describes a response with status code 200, with default header values.

OK
*/
type VscanOnAccessModifyCollectionOK struct {
}

// IsSuccess returns true when this vscan on access modify collection o k response has a 2xx status code
func (o *VscanOnAccessModifyCollectionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this vscan on access modify collection o k response has a 3xx status code
func (o *VscanOnAccessModifyCollectionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this vscan on access modify collection o k response has a 4xx status code
func (o *VscanOnAccessModifyCollectionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this vscan on access modify collection o k response has a 5xx status code
func (o *VscanOnAccessModifyCollectionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this vscan on access modify collection o k response a status code equal to that given
func (o *VscanOnAccessModifyCollectionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the vscan on access modify collection o k response
func (o *VscanOnAccessModifyCollectionOK) Code() int {
	return 200
}

func (o *VscanOnAccessModifyCollectionOK) Error() string {
	return fmt.Sprintf("[PATCH /protocols/vscan/{svm.uuid}/on-access-policies][%d] vscanOnAccessModifyCollectionOK", 200)
}

func (o *VscanOnAccessModifyCollectionOK) String() string {
	return fmt.Sprintf("[PATCH /protocols/vscan/{svm.uuid}/on-access-policies][%d] vscanOnAccessModifyCollectionOK", 200)
}

func (o *VscanOnAccessModifyCollectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewVscanOnAccessModifyCollectionDefault creates a VscanOnAccessModifyCollectionDefault with default headers values
func NewVscanOnAccessModifyCollectionDefault(code int) *VscanOnAccessModifyCollectionDefault {
	return &VscanOnAccessModifyCollectionDefault{
		_statusCode: code,
	}
}

/*
	VscanOnAccessModifyCollectionDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 10027033   | Configurations for an On-Access policy associated with a data SVM which was created by SVM owned by the cluster cannot be modified. However, the policy can be enabled or disabled. |
| 10027046   | The specified SVM is not the owner of the specified policy. Check for the correct SVM who owns the policy. |
| 10027101   | The file size must be in the range 1KB to 1TB |
| 10027107   | The include extensions list cannot be empty. Specify at least one extension for inclusion. |
| 10027109   | The specified CIFS path is invalid. It must be in the form \"\\dir1\\dir2\" or \"\\dir1\\dir2\\\". |
| 10027249   | The On-Access policy updated successfully but failed to enable/disable the policy. The reason for an enable policy operation failure might be that another policy is enabled. Disable the already enabled policy and then enable the policy. The reason for a disable policy operation failure might be that Vscan is enabled on the SVM. Disable the Vscan first and then disable the policy. |
| 10027250   | The On-Access policy cannot be enabled/disabled. The reason for an enable policy operation failure might be that another policy is enabled. Disable the already enabled policy and then enable the policy. The reason for a disable policy operation failure might be that Vscan is enabled on the SVM. Disable the Vscan and then disable the policy. |
| 10027253   | The number of paths specified exceeds the configured maximum number of paths. You cannot specify more than the maximum number of configured paths. |
| 10027254   | The number of extensions specified exceeds the configured maximum number of extensions. You cannot specify more than the maximum number of configured extensions. |
*/
type VscanOnAccessModifyCollectionDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this vscan on access modify collection default response has a 2xx status code
func (o *VscanOnAccessModifyCollectionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this vscan on access modify collection default response has a 3xx status code
func (o *VscanOnAccessModifyCollectionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this vscan on access modify collection default response has a 4xx status code
func (o *VscanOnAccessModifyCollectionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this vscan on access modify collection default response has a 5xx status code
func (o *VscanOnAccessModifyCollectionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this vscan on access modify collection default response a status code equal to that given
func (o *VscanOnAccessModifyCollectionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the vscan on access modify collection default response
func (o *VscanOnAccessModifyCollectionDefault) Code() int {
	return o._statusCode
}

func (o *VscanOnAccessModifyCollectionDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/vscan/{svm.uuid}/on-access-policies][%d] vscan_on_access_modify_collection default %s", o._statusCode, payload)
}

func (o *VscanOnAccessModifyCollectionDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/vscan/{svm.uuid}/on-access-policies][%d] vscan_on_access_modify_collection default %s", o._statusCode, payload)
}

func (o *VscanOnAccessModifyCollectionDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *VscanOnAccessModifyCollectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
VscanOnAccessModifyCollectionBody vscan on access modify collection body
swagger:model VscanOnAccessModifyCollectionBody
*/
type VscanOnAccessModifyCollectionBody struct {

	// Status of the On-Access Vscan policy
	Enabled *bool `json:"enabled,omitempty"`

	// Specifies if scanning is mandatory. File access is denied if there are no external virus-scanning servers available for virus scanning.
	Mandatory *bool `json:"mandatory,omitempty"`

	// On-Access policy name
	// Example: on-access-test
	// Max Length: 256
	// Min Length: 1
	Name *string `json:"name,omitempty"`

	// scope
	Scope *models.VscanOnAccessInlineScope `json:"scope,omitempty"`

	// svm
	Svm *models.VscanOnAccessInlineSvm `json:"svm,omitempty"`

	// vscan on access response inline records
	VscanOnAccessResponseInlineRecords []*models.VscanOnAccess `json:"records,omitempty"`
}

// Validate validates this vscan on access modify collection body
func (o *VscanOnAccessModifyCollectionBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateScope(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSvm(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateVscanOnAccessResponseInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *VscanOnAccessModifyCollectionBody) validateName(formats strfmt.Registry) error {
	if swag.IsZero(o.Name) { // not required
		return nil
	}

	if err := validate.MinLength("info"+"."+"name", "body", *o.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("info"+"."+"name", "body", *o.Name, 256); err != nil {
		return err
	}

	return nil
}

func (o *VscanOnAccessModifyCollectionBody) validateScope(formats strfmt.Registry) error {
	if swag.IsZero(o.Scope) { // not required
		return nil
	}

	if o.Scope != nil {
		if err := o.Scope.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "scope")
			}
			return err
		}
	}

	return nil
}

func (o *VscanOnAccessModifyCollectionBody) validateSvm(formats strfmt.Registry) error {
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

func (o *VscanOnAccessModifyCollectionBody) validateVscanOnAccessResponseInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(o.VscanOnAccessResponseInlineRecords) { // not required
		return nil
	}

	for i := 0; i < len(o.VscanOnAccessResponseInlineRecords); i++ {
		if swag.IsZero(o.VscanOnAccessResponseInlineRecords[i]) { // not required
			continue
		}

		if o.VscanOnAccessResponseInlineRecords[i] != nil {
			if err := o.VscanOnAccessResponseInlineRecords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this vscan on access modify collection body based on the context it is used
func (o *VscanOnAccessModifyCollectionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateScope(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateSvm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateVscanOnAccessResponseInlineRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *VscanOnAccessModifyCollectionBody) contextValidateScope(ctx context.Context, formats strfmt.Registry) error {

	if o.Scope != nil {
		if err := o.Scope.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "scope")
			}
			return err
		}
	}

	return nil
}

func (o *VscanOnAccessModifyCollectionBody) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

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

func (o *VscanOnAccessModifyCollectionBody) contextValidateVscanOnAccessResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.VscanOnAccessResponseInlineRecords); i++ {

		if o.VscanOnAccessResponseInlineRecords[i] != nil {
			if err := o.VscanOnAccessResponseInlineRecords[i].ContextValidate(ctx, formats); err != nil {
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
func (o *VscanOnAccessModifyCollectionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VscanOnAccessModifyCollectionBody) UnmarshalBinary(b []byte) error {
	var res VscanOnAccessModifyCollectionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
VscanOnAccessInlineScope vscan on access inline scope
swagger:model vscan_on_access_inline_scope
*/
type VscanOnAccessInlineScope struct {

	// List of file extensions for which scanning is not performed.
	// Example: ["mp*","txt"]
	// Max Items: 300
	// Min Items: 1
	ExcludeExtensions []*string `json:"exclude_extensions,omitempty"`

	// List of file paths for which scanning must not be performed.
	// Example: ["\\dir1\\dir2\\name","\\vol\\a b","\\vol\\a,b\\"]
	// Max Items: 100
	// Min Items: 1
	ExcludePaths []*string `json:"exclude_paths,omitempty"`

	// List of file extensions to be scanned.
	// Example: ["mp*","txt"]
	// Max Items: 300
	// Min Items: 1
	IncludeExtensions []*string `json:"include_extensions,omitempty"`

	// Maximum file size, in bytes, allowed for scanning.
	// Example: 2147483648
	// Maximum: 1.099511627776e+12
	// Minimum: 1024
	MaxFileSize *int64 `json:"max_file_size,omitempty"`

	// Scan only files opened with execute-access.
	OnlyExecuteAccess *bool `json:"only_execute_access,omitempty"`

	// Specifies whether or not read-only volume can be scanned.
	ScanReadonlyVolumes *bool `json:"scan_readonly_volumes,omitempty"`

	// Specifies whether or not files without any extension can be scanned.
	ScanWithoutExtension *bool `json:"scan_without_extension,omitempty"`
}

// Validate validates this vscan on access inline scope
func (o *VscanOnAccessInlineScope) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateExcludeExtensions(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateExcludePaths(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateIncludeExtensions(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMaxFileSize(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *VscanOnAccessInlineScope) validateExcludeExtensions(formats strfmt.Registry) error {
	if swag.IsZero(o.ExcludeExtensions) { // not required
		return nil
	}

	iExcludeExtensionsSize := int64(len(o.ExcludeExtensions))

	if err := validate.MinItems("info"+"."+"scope"+"."+"exclude_extensions", "body", iExcludeExtensionsSize, 1); err != nil {
		return err
	}

	if err := validate.MaxItems("info"+"."+"scope"+"."+"exclude_extensions", "body", iExcludeExtensionsSize, 300); err != nil {
		return err
	}

	return nil
}

func (o *VscanOnAccessInlineScope) validateExcludePaths(formats strfmt.Registry) error {
	if swag.IsZero(o.ExcludePaths) { // not required
		return nil
	}

	iExcludePathsSize := int64(len(o.ExcludePaths))

	if err := validate.MinItems("info"+"."+"scope"+"."+"exclude_paths", "body", iExcludePathsSize, 1); err != nil {
		return err
	}

	if err := validate.MaxItems("info"+"."+"scope"+"."+"exclude_paths", "body", iExcludePathsSize, 100); err != nil {
		return err
	}

	for i := 0; i < len(o.ExcludePaths); i++ {
		if swag.IsZero(o.ExcludePaths[i]) { // not required
			continue
		}

		if err := validate.MinLength("info"+"."+"scope"+"."+"exclude_paths"+"."+strconv.Itoa(i), "body", *o.ExcludePaths[i], 1); err != nil {
			return err
		}

		if err := validate.MaxLength("info"+"."+"scope"+"."+"exclude_paths"+"."+strconv.Itoa(i), "body", *o.ExcludePaths[i], 255); err != nil {
			return err
		}

	}

	return nil
}

func (o *VscanOnAccessInlineScope) validateIncludeExtensions(formats strfmt.Registry) error {
	if swag.IsZero(o.IncludeExtensions) { // not required
		return nil
	}

	iIncludeExtensionsSize := int64(len(o.IncludeExtensions))

	if err := validate.MinItems("info"+"."+"scope"+"."+"include_extensions", "body", iIncludeExtensionsSize, 1); err != nil {
		return err
	}

	if err := validate.MaxItems("info"+"."+"scope"+"."+"include_extensions", "body", iIncludeExtensionsSize, 300); err != nil {
		return err
	}

	return nil
}

func (o *VscanOnAccessInlineScope) validateMaxFileSize(formats strfmt.Registry) error {
	if swag.IsZero(o.MaxFileSize) { // not required
		return nil
	}

	if err := validate.MinimumInt("info"+"."+"scope"+"."+"max_file_size", "body", *o.MaxFileSize, 1024, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("info"+"."+"scope"+"."+"max_file_size", "body", *o.MaxFileSize, 1.099511627776e+12, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this vscan on access inline scope based on context it is used
func (o *VscanOnAccessInlineScope) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *VscanOnAccessInlineScope) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VscanOnAccessInlineScope) UnmarshalBinary(b []byte) error {
	var res VscanOnAccessInlineScope
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
VscanOnAccessInlineSvm SVM, applies only to SVM-scoped objects.
swagger:model vscan_on_access_inline_svm
*/
type VscanOnAccessInlineSvm struct {

	// links
	Links *models.VscanOnAccessInlineSvmInlineLinks `json:"_links,omitempty"`

	// The name of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: svm1
	Name *string `json:"name,omitempty"`

	// The unique identifier of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this vscan on access inline svm
func (o *VscanOnAccessInlineSvm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *VscanOnAccessInlineSvm) validateLinks(formats strfmt.Registry) error {
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

// ContextValidate validate this vscan on access inline svm based on the context it is used
func (o *VscanOnAccessInlineSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *VscanOnAccessInlineSvm) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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
func (o *VscanOnAccessInlineSvm) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VscanOnAccessInlineSvm) UnmarshalBinary(b []byte) error {
	var res VscanOnAccessInlineSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
VscanOnAccessInlineSvmInlineLinks vscan on access inline svm inline links
swagger:model vscan_on_access_inline_svm_inline__links
*/
type VscanOnAccessInlineSvmInlineLinks struct {

	// self
	Self *models.Href `json:"self,omitempty"`
}

// Validate validates this vscan on access inline svm inline links
func (o *VscanOnAccessInlineSvmInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *VscanOnAccessInlineSvmInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this vscan on access inline svm inline links based on the context it is used
func (o *VscanOnAccessInlineSvmInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *VscanOnAccessInlineSvmInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (o *VscanOnAccessInlineSvmInlineLinks) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VscanOnAccessInlineSvmInlineLinks) UnmarshalBinary(b []byte) error {
	var res VscanOnAccessInlineSvmInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}