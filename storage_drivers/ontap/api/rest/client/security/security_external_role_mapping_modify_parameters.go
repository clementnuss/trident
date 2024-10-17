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
	"github.com/go-openapi/swag"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// NewSecurityExternalRoleMappingModifyParams creates a new SecurityExternalRoleMappingModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSecurityExternalRoleMappingModifyParams() *SecurityExternalRoleMappingModifyParams {
	return &SecurityExternalRoleMappingModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSecurityExternalRoleMappingModifyParamsWithTimeout creates a new SecurityExternalRoleMappingModifyParams object
// with the ability to set a timeout on a request.
func NewSecurityExternalRoleMappingModifyParamsWithTimeout(timeout time.Duration) *SecurityExternalRoleMappingModifyParams {
	return &SecurityExternalRoleMappingModifyParams{
		timeout: timeout,
	}
}

// NewSecurityExternalRoleMappingModifyParamsWithContext creates a new SecurityExternalRoleMappingModifyParams object
// with the ability to set a context for a request.
func NewSecurityExternalRoleMappingModifyParamsWithContext(ctx context.Context) *SecurityExternalRoleMappingModifyParams {
	return &SecurityExternalRoleMappingModifyParams{
		Context: ctx,
	}
}

// NewSecurityExternalRoleMappingModifyParamsWithHTTPClient creates a new SecurityExternalRoleMappingModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewSecurityExternalRoleMappingModifyParamsWithHTTPClient(client *http.Client) *SecurityExternalRoleMappingModifyParams {
	return &SecurityExternalRoleMappingModifyParams{
		HTTPClient: client,
	}
}

/*
SecurityExternalRoleMappingModifyParams contains all the parameters to send to the API endpoint

	for the security external role mapping modify operation.

	Typically these are written to a http.Request.
*/
type SecurityExternalRoleMappingModifyParams struct {

	/* ExternalRole.

	   External Identity provider role.
	*/
	ExternalRole string

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* Info.

	   External role mapping modification details.
	*/
	Info *models.SecurityExternalRoleMapping

	/* Provider.

	   Type of the external identity provider.
	*/
	Provider string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the security external role mapping modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityExternalRoleMappingModifyParams) WithDefaults() *SecurityExternalRoleMappingModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the security external role mapping modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityExternalRoleMappingModifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the security external role mapping modify params
func (o *SecurityExternalRoleMappingModifyParams) WithTimeout(timeout time.Duration) *SecurityExternalRoleMappingModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the security external role mapping modify params
func (o *SecurityExternalRoleMappingModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the security external role mapping modify params
func (o *SecurityExternalRoleMappingModifyParams) WithContext(ctx context.Context) *SecurityExternalRoleMappingModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the security external role mapping modify params
func (o *SecurityExternalRoleMappingModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the security external role mapping modify params
func (o *SecurityExternalRoleMappingModifyParams) WithHTTPClient(client *http.Client) *SecurityExternalRoleMappingModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the security external role mapping modify params
func (o *SecurityExternalRoleMappingModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExternalRole adds the externalRole to the security external role mapping modify params
func (o *SecurityExternalRoleMappingModifyParams) WithExternalRole(externalRole string) *SecurityExternalRoleMappingModifyParams {
	o.SetExternalRole(externalRole)
	return o
}

// SetExternalRole adds the externalRole to the security external role mapping modify params
func (o *SecurityExternalRoleMappingModifyParams) SetExternalRole(externalRole string) {
	o.ExternalRole = externalRole
}

// WithFields adds the fields to the security external role mapping modify params
func (o *SecurityExternalRoleMappingModifyParams) WithFields(fields []string) *SecurityExternalRoleMappingModifyParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the security external role mapping modify params
func (o *SecurityExternalRoleMappingModifyParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithInfo adds the info to the security external role mapping modify params
func (o *SecurityExternalRoleMappingModifyParams) WithInfo(info *models.SecurityExternalRoleMapping) *SecurityExternalRoleMappingModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the security external role mapping modify params
func (o *SecurityExternalRoleMappingModifyParams) SetInfo(info *models.SecurityExternalRoleMapping) {
	o.Info = info
}

// WithProvider adds the provider to the security external role mapping modify params
func (o *SecurityExternalRoleMappingModifyParams) WithProvider(provider string) *SecurityExternalRoleMappingModifyParams {
	o.SetProvider(provider)
	return o
}

// SetProvider adds the provider to the security external role mapping modify params
func (o *SecurityExternalRoleMappingModifyParams) SetProvider(provider string) {
	o.Provider = provider
}

// WriteToRequest writes these params to a swagger request
func (o *SecurityExternalRoleMappingModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param external_role
	if err := r.SetPathParam("external_role", o.ExternalRole); err != nil {
		return err
	}

	if o.Fields != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	// path param provider
	if err := r.SetPathParam("provider", o.Provider); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamSecurityExternalRoleMappingModify binds the parameter fields
func (o *SecurityExternalRoleMappingModifyParams) bindParamFields(formats strfmt.Registry) []string {
	fieldsIR := o.Fields

	var fieldsIC []string
	for _, fieldsIIR := range fieldsIR { // explode []string

		fieldsIIV := fieldsIIR // string as string
		fieldsIC = append(fieldsIC, fieldsIIV)
	}

	// items.CollectionFormat: "csv"
	fieldsIS := swag.JoinByFormat(fieldsIC, "csv")

	return fieldsIS
}