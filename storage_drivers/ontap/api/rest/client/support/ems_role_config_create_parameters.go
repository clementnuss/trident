// Code generated by go-swagger; DO NOT EDIT.

package support

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

// NewEmsRoleConfigCreateParams creates a new EmsRoleConfigCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEmsRoleConfigCreateParams() *EmsRoleConfigCreateParams {
	return &EmsRoleConfigCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEmsRoleConfigCreateParamsWithTimeout creates a new EmsRoleConfigCreateParams object
// with the ability to set a timeout on a request.
func NewEmsRoleConfigCreateParamsWithTimeout(timeout time.Duration) *EmsRoleConfigCreateParams {
	return &EmsRoleConfigCreateParams{
		timeout: timeout,
	}
}

// NewEmsRoleConfigCreateParamsWithContext creates a new EmsRoleConfigCreateParams object
// with the ability to set a context for a request.
func NewEmsRoleConfigCreateParamsWithContext(ctx context.Context) *EmsRoleConfigCreateParams {
	return &EmsRoleConfigCreateParams{
		Context: ctx,
	}
}

// NewEmsRoleConfigCreateParamsWithHTTPClient creates a new EmsRoleConfigCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewEmsRoleConfigCreateParamsWithHTTPClient(client *http.Client) *EmsRoleConfigCreateParams {
	return &EmsRoleConfigCreateParams{
		HTTPClient: client,
	}
}

/*
EmsRoleConfigCreateParams contains all the parameters to send to the API endpoint

	for the ems role config create operation.

	Typically these are written to a http.Request.
*/
type EmsRoleConfigCreateParams struct {

	/* Info.

	   Information specification
	*/
	Info *models.EmsRoleConfig

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecords *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ems role config create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EmsRoleConfigCreateParams) WithDefaults() *EmsRoleConfigCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ems role config create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EmsRoleConfigCreateParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(false)
	)

	val := EmsRoleConfigCreateParams{
		ReturnRecords: &returnRecordsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the ems role config create params
func (o *EmsRoleConfigCreateParams) WithTimeout(timeout time.Duration) *EmsRoleConfigCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ems role config create params
func (o *EmsRoleConfigCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ems role config create params
func (o *EmsRoleConfigCreateParams) WithContext(ctx context.Context) *EmsRoleConfigCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ems role config create params
func (o *EmsRoleConfigCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ems role config create params
func (o *EmsRoleConfigCreateParams) WithHTTPClient(client *http.Client) *EmsRoleConfigCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ems role config create params
func (o *EmsRoleConfigCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the ems role config create params
func (o *EmsRoleConfigCreateParams) WithInfo(info *models.EmsRoleConfig) *EmsRoleConfigCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the ems role config create params
func (o *EmsRoleConfigCreateParams) SetInfo(info *models.EmsRoleConfig) {
	o.Info = info
}

// WithReturnRecords adds the returnRecords to the ems role config create params
func (o *EmsRoleConfigCreateParams) WithReturnRecords(returnRecords *bool) *EmsRoleConfigCreateParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the ems role config create params
func (o *EmsRoleConfigCreateParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WriteToRequest writes these params to a swagger request
func (o *EmsRoleConfigCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	if o.ReturnRecords != nil {

		// query param return_records
		var qrReturnRecords bool

		if o.ReturnRecords != nil {
			qrReturnRecords = *o.ReturnRecords
		}
		qReturnRecords := swag.FormatBool(qrReturnRecords)
		if qReturnRecords != "" {

			if err := r.SetQueryParam("return_records", qReturnRecords); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}