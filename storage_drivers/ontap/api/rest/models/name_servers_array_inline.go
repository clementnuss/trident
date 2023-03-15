// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NameServersArrayInline The list of IP addresses of the DNS servers. Addresses can be either
// IPv4 or IPv6 addresses.
//
// Example: ["10.224.65.20","2001:db08:a0b:12f0::1"]
//
// swagger:model name_servers_array_inline
type NameServersArrayInline []*string

// Validate validates this name servers array inline
func (m NameServersArrayInline) Validate(formats strfmt.Registry) error {
	var res []error

	iNameServersArrayInlineSize := int64(len(m))

	if err := validate.MaxItems("", "body", iNameServersArrayInlineSize, 3); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this name servers array inline based on context it is used
func (m NameServersArrayInline) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}