// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// GroupPolicyObjectCentralAccessPolicySettingsArrayInline List of central access policies.
// Example: ["p1","p2"]
//
// swagger:model group_policy_object_central_access_policy_settings_array_inline
type GroupPolicyObjectCentralAccessPolicySettingsArrayInline []*string

// Validate validates this group policy object central access policy settings array inline
func (m GroupPolicyObjectCentralAccessPolicySettingsArrayInline) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this group policy object central access policy settings array inline based on context it is used
func (m GroupPolicyObjectCentralAccessPolicySettingsArrayInline) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}