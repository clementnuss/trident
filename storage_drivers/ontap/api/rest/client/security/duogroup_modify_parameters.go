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

// NewDuogroupModifyParams creates a new DuogroupModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDuogroupModifyParams() *DuogroupModifyParams {
	return &DuogroupModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDuogroupModifyParamsWithTimeout creates a new DuogroupModifyParams object
// with the ability to set a timeout on a request.
func NewDuogroupModifyParamsWithTimeout(timeout time.Duration) *DuogroupModifyParams {
	return &DuogroupModifyParams{
		timeout: timeout,
	}
}

// NewDuogroupModifyParamsWithContext creates a new DuogroupModifyParams object
// with the ability to set a context for a request.
func NewDuogroupModifyParamsWithContext(ctx context.Context) *DuogroupModifyParams {
	return &DuogroupModifyParams{
		Context: ctx,
	}
}

// NewDuogroupModifyParamsWithHTTPClient creates a new DuogroupModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewDuogroupModifyParamsWithHTTPClient(client *http.Client) *DuogroupModifyParams {
	return &DuogroupModifyParams{
		HTTPClient: client,
	}
}

/*
DuogroupModifyParams contains all the parameters to send to the API endpoint

	for the duogroup modify operation.

	Typically these are written to a http.Request.
*/
type DuogroupModifyParams struct {

	/* Info.

	   Duo group modification details.
	*/
	Info *models.Duogroup

	/* Name.

	   Group name.
	*/
	Name string

	/* OwnerUUID.

	   Account owner UUID
	*/
	OwnerUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the duogroup modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DuogroupModifyParams) WithDefaults() *DuogroupModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the duogroup modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DuogroupModifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the duogroup modify params
func (o *DuogroupModifyParams) WithTimeout(timeout time.Duration) *DuogroupModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the duogroup modify params
func (o *DuogroupModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the duogroup modify params
func (o *DuogroupModifyParams) WithContext(ctx context.Context) *DuogroupModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the duogroup modify params
func (o *DuogroupModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the duogroup modify params
func (o *DuogroupModifyParams) WithHTTPClient(client *http.Client) *DuogroupModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the duogroup modify params
func (o *DuogroupModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the duogroup modify params
func (o *DuogroupModifyParams) WithInfo(info *models.Duogroup) *DuogroupModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the duogroup modify params
func (o *DuogroupModifyParams) SetInfo(info *models.Duogroup) {
	o.Info = info
}

// WithName adds the name to the duogroup modify params
func (o *DuogroupModifyParams) WithName(name string) *DuogroupModifyParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the duogroup modify params
func (o *DuogroupModifyParams) SetName(name string) {
	o.Name = name
}

// WithOwnerUUID adds the ownerUUID to the duogroup modify params
func (o *DuogroupModifyParams) WithOwnerUUID(ownerUUID string) *DuogroupModifyParams {
	o.SetOwnerUUID(ownerUUID)
	return o
}

// SetOwnerUUID adds the ownerUuid to the duogroup modify params
func (o *DuogroupModifyParams) SetOwnerUUID(ownerUUID string) {
	o.OwnerUUID = ownerUUID
}

// WriteToRequest writes these params to a swagger request
func (o *DuogroupModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}