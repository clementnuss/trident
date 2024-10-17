// Code generated by go-swagger; DO NOT EDIT.

package networking

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
)

// NewPerformanceFcInterfaceMetricGetParams creates a new PerformanceFcInterfaceMetricGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPerformanceFcInterfaceMetricGetParams() *PerformanceFcInterfaceMetricGetParams {
	return &PerformanceFcInterfaceMetricGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPerformanceFcInterfaceMetricGetParamsWithTimeout creates a new PerformanceFcInterfaceMetricGetParams object
// with the ability to set a timeout on a request.
func NewPerformanceFcInterfaceMetricGetParamsWithTimeout(timeout time.Duration) *PerformanceFcInterfaceMetricGetParams {
	return &PerformanceFcInterfaceMetricGetParams{
		timeout: timeout,
	}
}

// NewPerformanceFcInterfaceMetricGetParamsWithContext creates a new PerformanceFcInterfaceMetricGetParams object
// with the ability to set a context for a request.
func NewPerformanceFcInterfaceMetricGetParamsWithContext(ctx context.Context) *PerformanceFcInterfaceMetricGetParams {
	return &PerformanceFcInterfaceMetricGetParams{
		Context: ctx,
	}
}

// NewPerformanceFcInterfaceMetricGetParamsWithHTTPClient creates a new PerformanceFcInterfaceMetricGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewPerformanceFcInterfaceMetricGetParamsWithHTTPClient(client *http.Client) *PerformanceFcInterfaceMetricGetParams {
	return &PerformanceFcInterfaceMetricGetParams{
		HTTPClient: client,
	}
}

/*
PerformanceFcInterfaceMetricGetParams contains all the parameters to send to the API endpoint

	for the performance fc interface metric get operation.

	Typically these are written to a http.Request.
*/
type PerformanceFcInterfaceMetricGetParams struct {

	/* FcInterfaceUUID.

	   The unique identifier of the Fibre Channel interface.

	*/
	FcInterfaceUUID string

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* Timestamp.

	   The timestamp of the performance data.


	   Format: date-time
	*/
	Timestamp strfmt.DateTime

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the performance fc interface metric get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PerformanceFcInterfaceMetricGetParams) WithDefaults() *PerformanceFcInterfaceMetricGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the performance fc interface metric get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PerformanceFcInterfaceMetricGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the performance fc interface metric get params
func (o *PerformanceFcInterfaceMetricGetParams) WithTimeout(timeout time.Duration) *PerformanceFcInterfaceMetricGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the performance fc interface metric get params
func (o *PerformanceFcInterfaceMetricGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the performance fc interface metric get params
func (o *PerformanceFcInterfaceMetricGetParams) WithContext(ctx context.Context) *PerformanceFcInterfaceMetricGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the performance fc interface metric get params
func (o *PerformanceFcInterfaceMetricGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the performance fc interface metric get params
func (o *PerformanceFcInterfaceMetricGetParams) WithHTTPClient(client *http.Client) *PerformanceFcInterfaceMetricGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the performance fc interface metric get params
func (o *PerformanceFcInterfaceMetricGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFcInterfaceUUID adds the fcInterfaceUUID to the performance fc interface metric get params
func (o *PerformanceFcInterfaceMetricGetParams) WithFcInterfaceUUID(fcInterfaceUUID string) *PerformanceFcInterfaceMetricGetParams {
	o.SetFcInterfaceUUID(fcInterfaceUUID)
	return o
}

// SetFcInterfaceUUID adds the fcInterfaceUuid to the performance fc interface metric get params
func (o *PerformanceFcInterfaceMetricGetParams) SetFcInterfaceUUID(fcInterfaceUUID string) {
	o.FcInterfaceUUID = fcInterfaceUUID
}

// WithFields adds the fields to the performance fc interface metric get params
func (o *PerformanceFcInterfaceMetricGetParams) WithFields(fields []string) *PerformanceFcInterfaceMetricGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the performance fc interface metric get params
func (o *PerformanceFcInterfaceMetricGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithTimestamp adds the timestamp to the performance fc interface metric get params
func (o *PerformanceFcInterfaceMetricGetParams) WithTimestamp(timestamp strfmt.DateTime) *PerformanceFcInterfaceMetricGetParams {
	o.SetTimestamp(timestamp)
	return o
}

// SetTimestamp adds the timestamp to the performance fc interface metric get params
func (o *PerformanceFcInterfaceMetricGetParams) SetTimestamp(timestamp strfmt.DateTime) {
	o.Timestamp = timestamp
}

// WriteToRequest writes these params to a swagger request
func (o *PerformanceFcInterfaceMetricGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param fc_interface.uuid
	if err := r.SetPathParam("fc_interface.uuid", o.FcInterfaceUUID); err != nil {
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

	// path param timestamp
	if err := r.SetPathParam("timestamp", o.Timestamp.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamPerformanceFcInterfaceMetricGet binds the parameter fields
func (o *PerformanceFcInterfaceMetricGetParams) bindParamFields(formats strfmt.Registry) []string {
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