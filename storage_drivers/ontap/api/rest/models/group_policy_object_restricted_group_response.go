// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GroupPolicyObjectRestrictedGroupResponse group policy object restricted group response
//
// swagger:model group_policy_object_restricted_group_response
type GroupPolicyObjectRestrictedGroupResponse struct {

	// links
	Links *GroupPolicyObjectRestrictedGroupResponseInlineLinks `json:"_links,omitempty"`

	// group policy object restricted group response inline records
	GroupPolicyObjectRestrictedGroupResponseInlineRecords []*GroupPolicyObjectRestrictedGroup `json:"records,omitempty"`

	// Number of restricted group info.
	// Example: 1
	NumRecords *int64 `json:"num_records,omitempty"`
}

// Validate validates this group policy object restricted group response
func (m *GroupPolicyObjectRestrictedGroupResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroupPolicyObjectRestrictedGroupResponseInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GroupPolicyObjectRestrictedGroupResponse) validateLinks(formats strfmt.Registry) error {
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

func (m *GroupPolicyObjectRestrictedGroupResponse) validateGroupPolicyObjectRestrictedGroupResponseInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(m.GroupPolicyObjectRestrictedGroupResponseInlineRecords) { // not required
		return nil
	}

	for i := 0; i < len(m.GroupPolicyObjectRestrictedGroupResponseInlineRecords); i++ {
		if swag.IsZero(m.GroupPolicyObjectRestrictedGroupResponseInlineRecords[i]) { // not required
			continue
		}

		if m.GroupPolicyObjectRestrictedGroupResponseInlineRecords[i] != nil {
			if err := m.GroupPolicyObjectRestrictedGroupResponseInlineRecords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this group policy object restricted group response based on the context it is used
func (m *GroupPolicyObjectRestrictedGroupResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGroupPolicyObjectRestrictedGroupResponseInlineRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GroupPolicyObjectRestrictedGroupResponse) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *GroupPolicyObjectRestrictedGroupResponse) contextValidateGroupPolicyObjectRestrictedGroupResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.GroupPolicyObjectRestrictedGroupResponseInlineRecords); i++ {

		if m.GroupPolicyObjectRestrictedGroupResponseInlineRecords[i] != nil {
			if err := m.GroupPolicyObjectRestrictedGroupResponseInlineRecords[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GroupPolicyObjectRestrictedGroupResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GroupPolicyObjectRestrictedGroupResponse) UnmarshalBinary(b []byte) error {
	var res GroupPolicyObjectRestrictedGroupResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GroupPolicyObjectRestrictedGroupResponseInlineLinks group policy object restricted group response inline links
//
// swagger:model group_policy_object_restricted_group_response_inline__links
type GroupPolicyObjectRestrictedGroupResponseInlineLinks struct {

	// next
	Next *Href `json:"next,omitempty"`

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this group policy object restricted group response inline links
func (m *GroupPolicyObjectRestrictedGroupResponseInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GroupPolicyObjectRestrictedGroupResponseInlineLinks) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(m.Next) { // not required
		return nil
	}

	if m.Next != nil {
		if err := m.Next.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "next")
			}
			return err
		}
	}

	return nil
}

func (m *GroupPolicyObjectRestrictedGroupResponseInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this group policy object restricted group response inline links based on the context it is used
func (m *GroupPolicyObjectRestrictedGroupResponseInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNext(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GroupPolicyObjectRestrictedGroupResponseInlineLinks) contextValidateNext(ctx context.Context, formats strfmt.Registry) error {

	if m.Next != nil {
		if err := m.Next.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "next")
			}
			return err
		}
	}

	return nil
}

func (m *GroupPolicyObjectRestrictedGroupResponseInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *GroupPolicyObjectRestrictedGroupResponseInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GroupPolicyObjectRestrictedGroupResponseInlineLinks) UnmarshalBinary(b []byte) error {
	var res GroupPolicyObjectRestrictedGroupResponseInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}