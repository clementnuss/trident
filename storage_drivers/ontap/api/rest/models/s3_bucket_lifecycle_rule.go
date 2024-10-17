// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// S3BucketLifecycleRule Information about the lifecycle management rule of a bucket.
//
// swagger:model s3_bucket_lifecycle_rule
type S3BucketLifecycleRule struct {

	// links
	Links *S3BucketLifecycleRuleInlineLinks `json:"_links,omitempty"`

	// abort incomplete multipart upload
	AbortIncompleteMultipartUpload *S3BucketLifecycleRuleInlineAbortIncompleteMultipartUpload `json:"abort_incomplete_multipart_upload,omitempty"`

	// Specifies the name of the bucket. Bucket name is a string that can only contain the following combination of ASCII-range alphanumeric characters 0-9, a-z, ".", and "-".
	// Example: bucket1
	// Max Length: 63
	// Min Length: 3
	BucketName *string `json:"bucket_name,omitempty"`

	// Specifies whether or not the associated rule is enabled.
	Enabled *bool `json:"enabled,omitempty"`

	// expiration
	Expiration *S3BucketLifecycleRuleInlineExpiration `json:"expiration,omitempty"`

	// Bucket lifecycle management rule identifier. The length of the name can range from 0 to 256 characters.
	// Max Length: 256
	// Min Length: 0
	Name *string `json:"name,omitempty"`

	// non current version expiration
	NonCurrentVersionExpiration *S3BucketLifecycleRuleInlineNonCurrentVersionExpiration `json:"non_current_version_expiration,omitempty"`

	// object filter
	ObjectFilter *S3BucketLifecycleRuleInlineObjectFilter `json:"object_filter,omitempty"`

	// svm
	Svm *S3BucketLifecycleRuleInlineSvm `json:"svm,omitempty"`

	// Specifies the unique identifier of the bucket.
	// Example: 414b29a1-3b26-11e9-bd58-0050568ea055
	// Read Only: true
	// Format: uuid
	UUID *strfmt.UUID `json:"uuid,omitempty"`
}

// Validate validates this s3 bucket lifecycle rule
func (m *S3BucketLifecycleRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAbortIncompleteMultipartUpload(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBucketName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpiration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNonCurrentVersionExpiration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObjectFilter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSvm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3BucketLifecycleRule) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *S3BucketLifecycleRule) validateAbortIncompleteMultipartUpload(formats strfmt.Registry) error {
	if swag.IsZero(m.AbortIncompleteMultipartUpload) { // not required
		return nil
	}

	if m.AbortIncompleteMultipartUpload != nil {
		if err := m.AbortIncompleteMultipartUpload.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("abort_incomplete_multipart_upload")
			}
			return err
		}
	}

	return nil
}

func (m *S3BucketLifecycleRule) validateBucketName(formats strfmt.Registry) error {
	if swag.IsZero(m.BucketName) { // not required
		return nil
	}

	if err := validate.MinLength("bucket_name", "body", *m.BucketName, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("bucket_name", "body", *m.BucketName, 63); err != nil {
		return err
	}

	return nil
}

func (m *S3BucketLifecycleRule) validateExpiration(formats strfmt.Registry) error {
	if swag.IsZero(m.Expiration) { // not required
		return nil
	}

	if m.Expiration != nil {
		if err := m.Expiration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("expiration")
			}
			return err
		}
	}

	return nil
}

func (m *S3BucketLifecycleRule) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MinLength("name", "body", *m.Name, 0); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 256); err != nil {
		return err
	}

	return nil
}

func (m *S3BucketLifecycleRule) validateNonCurrentVersionExpiration(formats strfmt.Registry) error {
	if swag.IsZero(m.NonCurrentVersionExpiration) { // not required
		return nil
	}

	if m.NonCurrentVersionExpiration != nil {
		if err := m.NonCurrentVersionExpiration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("non_current_version_expiration")
			}
			return err
		}
	}

	return nil
}

func (m *S3BucketLifecycleRule) validateObjectFilter(formats strfmt.Registry) error {
	if swag.IsZero(m.ObjectFilter) { // not required
		return nil
	}

	if m.ObjectFilter != nil {
		if err := m.ObjectFilter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("object_filter")
			}
			return err
		}
	}

	return nil
}

func (m *S3BucketLifecycleRule) validateSvm(formats strfmt.Registry) error {
	if swag.IsZero(m.Svm) { // not required
		return nil
	}

	if m.Svm != nil {
		if err := m.Svm.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm")
			}
			return err
		}
	}

	return nil
}

func (m *S3BucketLifecycleRule) validateUUID(formats strfmt.Registry) error {
	if swag.IsZero(m.UUID) { // not required
		return nil
	}

	if err := validate.FormatOf("uuid", "body", "uuid", m.UUID.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this s3 bucket lifecycle rule based on the context it is used
func (m *S3BucketLifecycleRule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAbortIncompleteMultipartUpload(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExpiration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNonCurrentVersionExpiration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateObjectFilter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSvm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUUID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3BucketLifecycleRule) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *S3BucketLifecycleRule) contextValidateAbortIncompleteMultipartUpload(ctx context.Context, formats strfmt.Registry) error {

	if m.AbortIncompleteMultipartUpload != nil {
		if err := m.AbortIncompleteMultipartUpload.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("abort_incomplete_multipart_upload")
			}
			return err
		}
	}

	return nil
}

func (m *S3BucketLifecycleRule) contextValidateExpiration(ctx context.Context, formats strfmt.Registry) error {

	if m.Expiration != nil {
		if err := m.Expiration.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("expiration")
			}
			return err
		}
	}

	return nil
}

func (m *S3BucketLifecycleRule) contextValidateNonCurrentVersionExpiration(ctx context.Context, formats strfmt.Registry) error {

	if m.NonCurrentVersionExpiration != nil {
		if err := m.NonCurrentVersionExpiration.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("non_current_version_expiration")
			}
			return err
		}
	}

	return nil
}

func (m *S3BucketLifecycleRule) contextValidateObjectFilter(ctx context.Context, formats strfmt.Registry) error {

	if m.ObjectFilter != nil {
		if err := m.ObjectFilter.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("object_filter")
			}
			return err
		}
	}

	return nil
}

func (m *S3BucketLifecycleRule) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

	if m.Svm != nil {
		if err := m.Svm.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm")
			}
			return err
		}
	}

	return nil
}

func (m *S3BucketLifecycleRule) contextValidateUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "uuid", "body", m.UUID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *S3BucketLifecycleRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *S3BucketLifecycleRule) UnmarshalBinary(b []byte) error {
	var res S3BucketLifecycleRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// S3BucketLifecycleRuleInlineAbortIncompleteMultipartUpload Specifies a way to perform abort_incomplete_multipart_upload action on filtered objects within a bucket. It cannot be specified with tags.
//
// swagger:model s3_bucket_lifecycle_rule_inline_abort_incomplete_multipart_upload
type S3BucketLifecycleRuleInlineAbortIncompleteMultipartUpload struct {

	// links
	Links *S3BucketLifecycleRuleInlineAbortIncompleteMultipartUploadInlineLinks `json:"_links,omitempty"`

	// Number of days of initiation after which uploads can be aborted.
	AfterInitiationDays *int64 `json:"after_initiation_days,omitempty"`
}

// Validate validates this s3 bucket lifecycle rule inline abort incomplete multipart upload
func (m *S3BucketLifecycleRuleInlineAbortIncompleteMultipartUpload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3BucketLifecycleRuleInlineAbortIncompleteMultipartUpload) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("abort_incomplete_multipart_upload" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this s3 bucket lifecycle rule inline abort incomplete multipart upload based on the context it is used
func (m *S3BucketLifecycleRuleInlineAbortIncompleteMultipartUpload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3BucketLifecycleRuleInlineAbortIncompleteMultipartUpload) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("abort_incomplete_multipart_upload" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *S3BucketLifecycleRuleInlineAbortIncompleteMultipartUpload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *S3BucketLifecycleRuleInlineAbortIncompleteMultipartUpload) UnmarshalBinary(b []byte) error {
	var res S3BucketLifecycleRuleInlineAbortIncompleteMultipartUpload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// S3BucketLifecycleRuleInlineAbortIncompleteMultipartUploadInlineLinks s3 bucket lifecycle rule inline abort incomplete multipart upload inline links
//
// swagger:model s3_bucket_lifecycle_rule_inline_abort_incomplete_multipart_upload_inline__links
type S3BucketLifecycleRuleInlineAbortIncompleteMultipartUploadInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this s3 bucket lifecycle rule inline abort incomplete multipart upload inline links
func (m *S3BucketLifecycleRuleInlineAbortIncompleteMultipartUploadInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3BucketLifecycleRuleInlineAbortIncompleteMultipartUploadInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("abort_incomplete_multipart_upload" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this s3 bucket lifecycle rule inline abort incomplete multipart upload inline links based on the context it is used
func (m *S3BucketLifecycleRuleInlineAbortIncompleteMultipartUploadInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3BucketLifecycleRuleInlineAbortIncompleteMultipartUploadInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("abort_incomplete_multipart_upload" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *S3BucketLifecycleRuleInlineAbortIncompleteMultipartUploadInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *S3BucketLifecycleRuleInlineAbortIncompleteMultipartUploadInlineLinks) UnmarshalBinary(b []byte) error {
	var res S3BucketLifecycleRuleInlineAbortIncompleteMultipartUploadInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// S3BucketLifecycleRuleInlineExpiration Specifies a way to perform expiration action on filtered objects within a bucket.
//
// swagger:model s3_bucket_lifecycle_rule_inline_expiration
type S3BucketLifecycleRuleInlineExpiration struct {

	// links
	Links *S3BucketLifecycleRuleInlineExpirationInlineLinks `json:"_links,omitempty"`

	// Cleanup object delete markers.
	ExpiredObjectDeleteMarker *bool `json:"expired_object_delete_marker,omitempty"`

	// Number of days since creation after which objects can be deleted. This cannot be used along with object_expiry_date.
	// Example: 100
	ObjectAgeDays *int64 `json:"object_age_days,omitempty"`

	// Specific date from when objects can expire. This cannot be used with object_age_days.
	// Example: 2039-09-23 00:00:00
	// Format: date-time
	ObjectExpiryDate *strfmt.DateTime `json:"object_expiry_date,omitempty"`
}

// Validate validates this s3 bucket lifecycle rule inline expiration
func (m *S3BucketLifecycleRuleInlineExpiration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObjectExpiryDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3BucketLifecycleRuleInlineExpiration) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("expiration" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

func (m *S3BucketLifecycleRuleInlineExpiration) validateObjectExpiryDate(formats strfmt.Registry) error {
	if swag.IsZero(m.ObjectExpiryDate) { // not required
		return nil
	}

	if err := validate.FormatOf("expiration"+"."+"object_expiry_date", "body", "date-time", m.ObjectExpiryDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this s3 bucket lifecycle rule inline expiration based on the context it is used
func (m *S3BucketLifecycleRuleInlineExpiration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3BucketLifecycleRuleInlineExpiration) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("expiration" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *S3BucketLifecycleRuleInlineExpiration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *S3BucketLifecycleRuleInlineExpiration) UnmarshalBinary(b []byte) error {
	var res S3BucketLifecycleRuleInlineExpiration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// S3BucketLifecycleRuleInlineExpirationInlineLinks s3 bucket lifecycle rule inline expiration inline links
//
// swagger:model s3_bucket_lifecycle_rule_inline_expiration_inline__links
type S3BucketLifecycleRuleInlineExpirationInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this s3 bucket lifecycle rule inline expiration inline links
func (m *S3BucketLifecycleRuleInlineExpirationInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3BucketLifecycleRuleInlineExpirationInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("expiration" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this s3 bucket lifecycle rule inline expiration inline links based on the context it is used
func (m *S3BucketLifecycleRuleInlineExpirationInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3BucketLifecycleRuleInlineExpirationInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("expiration" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *S3BucketLifecycleRuleInlineExpirationInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *S3BucketLifecycleRuleInlineExpirationInlineLinks) UnmarshalBinary(b []byte) error {
	var res S3BucketLifecycleRuleInlineExpirationInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// S3BucketLifecycleRuleInlineLinks s3 bucket lifecycle rule inline links
//
// swagger:model s3_bucket_lifecycle_rule_inline__links
type S3BucketLifecycleRuleInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this s3 bucket lifecycle rule inline links
func (m *S3BucketLifecycleRuleInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3BucketLifecycleRuleInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this s3 bucket lifecycle rule inline links based on the context it is used
func (m *S3BucketLifecycleRuleInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3BucketLifecycleRuleInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *S3BucketLifecycleRuleInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *S3BucketLifecycleRuleInlineLinks) UnmarshalBinary(b []byte) error {
	var res S3BucketLifecycleRuleInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// S3BucketLifecycleRuleInlineNonCurrentVersionExpiration Specifies a way to perform non_current_version_expiration action on filtered objects within a bucket.
//
// swagger:model s3_bucket_lifecycle_rule_inline_non_current_version_expiration
type S3BucketLifecycleRuleInlineNonCurrentVersionExpiration struct {

	// links
	Links *S3BucketLifecycleRuleInlineNonCurrentVersionExpirationInlineLinks `json:"_links,omitempty"`

	// Number of latest non-current versions to be retained.
	NewNonCurrentVersions *int64 `json:"new_non_current_versions,omitempty"`

	// Number of days after which non-current versions can be deleted.
	NonCurrentDays *int64 `json:"non_current_days,omitempty"`
}

// Validate validates this s3 bucket lifecycle rule inline non current version expiration
func (m *S3BucketLifecycleRuleInlineNonCurrentVersionExpiration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3BucketLifecycleRuleInlineNonCurrentVersionExpiration) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("non_current_version_expiration" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this s3 bucket lifecycle rule inline non current version expiration based on the context it is used
func (m *S3BucketLifecycleRuleInlineNonCurrentVersionExpiration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3BucketLifecycleRuleInlineNonCurrentVersionExpiration) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("non_current_version_expiration" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *S3BucketLifecycleRuleInlineNonCurrentVersionExpiration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *S3BucketLifecycleRuleInlineNonCurrentVersionExpiration) UnmarshalBinary(b []byte) error {
	var res S3BucketLifecycleRuleInlineNonCurrentVersionExpiration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// S3BucketLifecycleRuleInlineNonCurrentVersionExpirationInlineLinks s3 bucket lifecycle rule inline non current version expiration inline links
//
// swagger:model s3_bucket_lifecycle_rule_inline_non_current_version_expiration_inline__links
type S3BucketLifecycleRuleInlineNonCurrentVersionExpirationInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this s3 bucket lifecycle rule inline non current version expiration inline links
func (m *S3BucketLifecycleRuleInlineNonCurrentVersionExpirationInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3BucketLifecycleRuleInlineNonCurrentVersionExpirationInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("non_current_version_expiration" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this s3 bucket lifecycle rule inline non current version expiration inline links based on the context it is used
func (m *S3BucketLifecycleRuleInlineNonCurrentVersionExpirationInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3BucketLifecycleRuleInlineNonCurrentVersionExpirationInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("non_current_version_expiration" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *S3BucketLifecycleRuleInlineNonCurrentVersionExpirationInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *S3BucketLifecycleRuleInlineNonCurrentVersionExpirationInlineLinks) UnmarshalBinary(b []byte) error {
	var res S3BucketLifecycleRuleInlineNonCurrentVersionExpirationInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// S3BucketLifecycleRuleInlineObjectFilter Specifies a way to filter objects within a bucket.
//
// swagger:model s3_bucket_lifecycle_rule_inline_object_filter
type S3BucketLifecycleRuleInlineObjectFilter struct {

	// links
	Links *S3BucketLifecycleRuleInlineObjectFilterInlineLinks `json:"_links,omitempty"`

	// A prefix that is matched against object-names within a bucket.
	// Example: /logs
	Prefix *string `json:"prefix,omitempty"`

	// Size of the object greater than specified for which the corresponding lifecycle rule is to be applied.
	// Example: 10240
	SizeGreaterThan *int64 `json:"size_greater_than,omitempty"`

	// Size of the object smaller than specified for which the corresponding lifecycle rule is to be applied.
	// Example: 10485760
	SizeLessThan *int64 `json:"size_less_than,omitempty"`

	// An array of key-value paired tags of the form <tag> or <tag=value>.
	//
	// Example: ["project1=projA","project2=projB"]
	Tags []*string `json:"tags,omitempty"`
}

// Validate validates this s3 bucket lifecycle rule inline object filter
func (m *S3BucketLifecycleRuleInlineObjectFilter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3BucketLifecycleRuleInlineObjectFilter) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("object_filter" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this s3 bucket lifecycle rule inline object filter based on the context it is used
func (m *S3BucketLifecycleRuleInlineObjectFilter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3BucketLifecycleRuleInlineObjectFilter) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("object_filter" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *S3BucketLifecycleRuleInlineObjectFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *S3BucketLifecycleRuleInlineObjectFilter) UnmarshalBinary(b []byte) error {
	var res S3BucketLifecycleRuleInlineObjectFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// S3BucketLifecycleRuleInlineObjectFilterInlineLinks s3 bucket lifecycle rule inline object filter inline links
//
// swagger:model s3_bucket_lifecycle_rule_inline_object_filter_inline__links
type S3BucketLifecycleRuleInlineObjectFilterInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this s3 bucket lifecycle rule inline object filter inline links
func (m *S3BucketLifecycleRuleInlineObjectFilterInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3BucketLifecycleRuleInlineObjectFilterInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("object_filter" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this s3 bucket lifecycle rule inline object filter inline links based on the context it is used
func (m *S3BucketLifecycleRuleInlineObjectFilterInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3BucketLifecycleRuleInlineObjectFilterInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("object_filter" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *S3BucketLifecycleRuleInlineObjectFilterInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *S3BucketLifecycleRuleInlineObjectFilterInlineLinks) UnmarshalBinary(b []byte) error {
	var res S3BucketLifecycleRuleInlineObjectFilterInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// S3BucketLifecycleRuleInlineSvm Specifies the name of the SVM where this bucket exists.
//
// swagger:model s3_bucket_lifecycle_rule_inline_svm
type S3BucketLifecycleRuleInlineSvm struct {

	// links
	Links *S3BucketLifecycleRuleInlineSvmInlineLinks `json:"_links,omitempty"`

	// The name of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: svm1
	Name *string `json:"name,omitempty"`

	// The unique identifier of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this s3 bucket lifecycle rule inline svm
func (m *S3BucketLifecycleRuleInlineSvm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3BucketLifecycleRuleInlineSvm) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this s3 bucket lifecycle rule inline svm based on the context it is used
func (m *S3BucketLifecycleRuleInlineSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3BucketLifecycleRuleInlineSvm) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *S3BucketLifecycleRuleInlineSvm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *S3BucketLifecycleRuleInlineSvm) UnmarshalBinary(b []byte) error {
	var res S3BucketLifecycleRuleInlineSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// S3BucketLifecycleRuleInlineSvmInlineLinks s3 bucket lifecycle rule inline svm inline links
//
// swagger:model s3_bucket_lifecycle_rule_inline_svm_inline__links
type S3BucketLifecycleRuleInlineSvmInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this s3 bucket lifecycle rule inline svm inline links
func (m *S3BucketLifecycleRuleInlineSvmInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3BucketLifecycleRuleInlineSvmInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this s3 bucket lifecycle rule inline svm inline links based on the context it is used
func (m *S3BucketLifecycleRuleInlineSvmInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3BucketLifecycleRuleInlineSvmInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *S3BucketLifecycleRuleInlineSvmInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *S3BucketLifecycleRuleInlineSvmInlineLinks) UnmarshalBinary(b []byte) error {
	var res S3BucketLifecycleRuleInlineSvmInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}