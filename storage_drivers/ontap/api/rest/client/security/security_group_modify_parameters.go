// Code generated by go-swagger; DO NOT EDIT.

package security

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// NewSecurityGroupModifyParams creates a new SecurityGroupModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSecurityGroupModifyParams() *SecurityGroupModifyParams {
	return &SecurityGroupModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSecurityGroupModifyParamsWithTimeout creates a new SecurityGroupModifyParams object
// with the ability to set a timeout on a request.
func NewSecurityGroupModifyParamsWithTimeout(timeout time.Duration) *SecurityGroupModifyParams {
	return &SecurityGroupModifyParams{
		timeout: timeout,
	}
}

// NewSecurityGroupModifyParamsWithContext creates a new SecurityGroupModifyParams object
// with the ability to set a context for a request.
func NewSecurityGroupModifyParamsWithContext(ctx context.Context) *SecurityGroupModifyParams {
	return &SecurityGroupModifyParams{
		Context: ctx,
	}
}

// NewSecurityGroupModifyParamsWithHTTPClient creates a new SecurityGroupModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewSecurityGroupModifyParamsWithHTTPClient(client *http.Client) *SecurityGroupModifyParams {
	return &SecurityGroupModifyParams{
		HTTPClient: client,
	}
}

/*
SecurityGroupModifyParams contains all the parameters to send to the API endpoint

	for the security group modify operation.

	Typically these are written to a http.Request.
*/
type SecurityGroupModifyParams struct {

	/* Info.

	   Group configuration modification details.
	*/
	Info *models.SecurityGroup

	/* Name.

	   Group name.
	*/
	Name string

	/* OwnerUUID.

	   Group owner. Used to identify a cluster or an SVM.
	*/
	OwnerUUID string

	/* Type.

	   Group type.
	*/
	Type string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the security group modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityGroupModifyParams) WithDefaults() *SecurityGroupModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the security group modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityGroupModifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the security group modify params
func (o *SecurityGroupModifyParams) WithTimeout(timeout time.Duration) *SecurityGroupModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the security group modify params
func (o *SecurityGroupModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the security group modify params
func (o *SecurityGroupModifyParams) WithContext(ctx context.Context) *SecurityGroupModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the security group modify params
func (o *SecurityGroupModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the security group modify params
func (o *SecurityGroupModifyParams) WithHTTPClient(client *http.Client) *SecurityGroupModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the security group modify params
func (o *SecurityGroupModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the security group modify params
func (o *SecurityGroupModifyParams) WithInfo(info *models.SecurityGroup) *SecurityGroupModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the security group modify params
func (o *SecurityGroupModifyParams) SetInfo(info *models.SecurityGroup) {
	o.Info = info
}

// WithName adds the name to the security group modify params
func (o *SecurityGroupModifyParams) WithName(name string) *SecurityGroupModifyParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the security group modify params
func (o *SecurityGroupModifyParams) SetName(name string) {
	o.Name = name
}

// WithOwnerUUID adds the ownerUUID to the security group modify params
func (o *SecurityGroupModifyParams) WithOwnerUUID(ownerUUID string) *SecurityGroupModifyParams {
	o.SetOwnerUUID(ownerUUID)
	return o
}

// SetOwnerUUID adds the ownerUuid to the security group modify params
func (o *SecurityGroupModifyParams) SetOwnerUUID(ownerUUID string) {
	o.OwnerUUID = ownerUUID
}

// WithType adds the typeVar to the security group modify params
func (o *SecurityGroupModifyParams) WithType(typeVar string) *SecurityGroupModifyParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the security group modify params
func (o *SecurityGroupModifyParams) SetType(typeVar string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *SecurityGroupModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// path param owner.uuid
	if err := r.SetPathParam("owner.uuid", o.OwnerUUID); err != nil {
		return err
	}

	// path param type
	if err := r.SetPathParam("type", o.Type); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}