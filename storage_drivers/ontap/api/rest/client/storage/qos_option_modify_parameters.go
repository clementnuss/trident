// Code generated by go-swagger; DO NOT EDIT.

package storage

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

// NewQosOptionModifyParams creates a new QosOptionModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewQosOptionModifyParams() *QosOptionModifyParams {
	return &QosOptionModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewQosOptionModifyParamsWithTimeout creates a new QosOptionModifyParams object
// with the ability to set a timeout on a request.
func NewQosOptionModifyParamsWithTimeout(timeout time.Duration) *QosOptionModifyParams {
	return &QosOptionModifyParams{
		timeout: timeout,
	}
}

// NewQosOptionModifyParamsWithContext creates a new QosOptionModifyParams object
// with the ability to set a context for a request.
func NewQosOptionModifyParamsWithContext(ctx context.Context) *QosOptionModifyParams {
	return &QosOptionModifyParams{
		Context: ctx,
	}
}

// NewQosOptionModifyParamsWithHTTPClient creates a new QosOptionModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewQosOptionModifyParamsWithHTTPClient(client *http.Client) *QosOptionModifyParams {
	return &QosOptionModifyParams{
		HTTPClient: client,
	}
}

/*
QosOptionModifyParams contains all the parameters to send to the API endpoint

	for the qos option modify operation.

	Typically these are written to a http.Request.
*/
type QosOptionModifyParams struct {

	/* Info.

	   QoS option information
	*/
	Info *models.QosOption

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the qos option modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QosOptionModifyParams) WithDefaults() *QosOptionModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the qos option modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QosOptionModifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the qos option modify params
func (o *QosOptionModifyParams) WithTimeout(timeout time.Duration) *QosOptionModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the qos option modify params
func (o *QosOptionModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the qos option modify params
func (o *QosOptionModifyParams) WithContext(ctx context.Context) *QosOptionModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the qos option modify params
func (o *QosOptionModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the qos option modify params
func (o *QosOptionModifyParams) WithHTTPClient(client *http.Client) *QosOptionModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the qos option modify params
func (o *QosOptionModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the qos option modify params
func (o *QosOptionModifyParams) WithInfo(info *models.QosOption) *QosOptionModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the qos option modify params
func (o *QosOptionModifyParams) SetInfo(info *models.QosOption) {
	o.Info = info
}

// WriteToRequest writes these params to a swagger request
func (o *QosOptionModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}