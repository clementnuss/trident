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

// NewDuoCreateParams creates a new DuoCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDuoCreateParams() *DuoCreateParams {
	return &DuoCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDuoCreateParamsWithTimeout creates a new DuoCreateParams object
// with the ability to set a timeout on a request.
func NewDuoCreateParamsWithTimeout(timeout time.Duration) *DuoCreateParams {
	return &DuoCreateParams{
		timeout: timeout,
	}
}

// NewDuoCreateParamsWithContext creates a new DuoCreateParams object
// with the ability to set a context for a request.
func NewDuoCreateParamsWithContext(ctx context.Context) *DuoCreateParams {
	return &DuoCreateParams{
		Context: ctx,
	}
}

// NewDuoCreateParamsWithHTTPClient creates a new DuoCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDuoCreateParamsWithHTTPClient(client *http.Client) *DuoCreateParams {
	return &DuoCreateParams{
		HTTPClient: client,
	}
}

/*
DuoCreateParams contains all the parameters to send to the API endpoint

	for the duo create operation.

	Typically these are written to a http.Request.
*/
type DuoCreateParams struct {

	/* Info.

	   Duo configurations to create profile.
	*/
	Info *models.Duo

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecords *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the duo create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DuoCreateParams) WithDefaults() *DuoCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the duo create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DuoCreateParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(false)
	)

	val := DuoCreateParams{
		ReturnRecords: &returnRecordsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the duo create params
func (o *DuoCreateParams) WithTimeout(timeout time.Duration) *DuoCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the duo create params
func (o *DuoCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the duo create params
func (o *DuoCreateParams) WithContext(ctx context.Context) *DuoCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the duo create params
func (o *DuoCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the duo create params
func (o *DuoCreateParams) WithHTTPClient(client *http.Client) *DuoCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the duo create params
func (o *DuoCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the duo create params
func (o *DuoCreateParams) WithInfo(info *models.Duo) *DuoCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the duo create params
func (o *DuoCreateParams) SetInfo(info *models.Duo) {
	o.Info = info
}

// WithReturnRecords adds the returnRecords to the duo create params
func (o *DuoCreateParams) WithReturnRecords(returnRecords *bool) *DuoCreateParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the duo create params
func (o *DuoCreateParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WriteToRequest writes these params to a swagger request
func (o *DuoCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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