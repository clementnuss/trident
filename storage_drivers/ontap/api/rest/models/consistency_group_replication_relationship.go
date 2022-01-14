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

// ConsistencyGroupReplicationRelationship consistency group replication relationship
//
// swagger:model consistency_group_replication_relationship
type ConsistencyGroupReplicationRelationship struct {

	// links
	Links *SelfLink `json:"_links,omitempty"`

	// Indicates whether or not this consistency group is the source for replication.
	//
	// Read Only: true
	IsSource *bool `json:"is_source,omitempty"`

	// The unique identifier of the SnapMirror relationship.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	// Read Only: true
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this consistency group replication relationship
func (m *ConsistencyGroupReplicationRelationship) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConsistencyGroupReplicationRelationship) validateLinks(formats strfmt.Registry) error {
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

// ContextValidate validate this consistency group replication relationship based on the context it is used
func (m *ConsistencyGroupReplicationRelationship) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIsSource(ctx, formats); err != nil {
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

func (m *ConsistencyGroupReplicationRelationship) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ConsistencyGroupReplicationRelationship) contextValidateIsSource(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "is_source", "body", m.IsSource); err != nil {
		return err
	}

	return nil
}

func (m *ConsistencyGroupReplicationRelationship) contextValidateUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "uuid", "body", string(m.UUID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConsistencyGroupReplicationRelationship) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsistencyGroupReplicationRelationship) UnmarshalBinary(b []byte) error {
	var res ConsistencyGroupReplicationRelationship
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}