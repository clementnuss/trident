// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ManagementProtocols Cluster-wide security protocols related information.
//
//
// swagger:model management_protocols
type ManagementProtocols struct {

	// Indicates whether or not security protocol rsh is enabled on the cluster.
	RshEnabled bool `json:"rsh_enabled,omitempty"`

	// Indicates whether or not security protocol telnet is enabled on the cluster.
	TelnetEnabled bool `json:"telnet_enabled,omitempty"`
}

// Validate validates this management protocols
func (m *ManagementProtocols) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this management protocols based on context it is used
func (m *ManagementProtocols) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ManagementProtocols) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ManagementProtocols) UnmarshalBinary(b []byte) error {
	var res ManagementProtocols
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}