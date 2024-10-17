// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

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

// NewNfsTLSInterfaceGetParams creates a new NfsTLSInterfaceGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNfsTLSInterfaceGetParams() *NfsTLSInterfaceGetParams {
	return &NfsTLSInterfaceGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNfsTLSInterfaceGetParamsWithTimeout creates a new NfsTLSInterfaceGetParams object
// with the ability to set a timeout on a request.
func NewNfsTLSInterfaceGetParamsWithTimeout(timeout time.Duration) *NfsTLSInterfaceGetParams {
	return &NfsTLSInterfaceGetParams{
		timeout: timeout,
	}
}

// NewNfsTLSInterfaceGetParamsWithContext creates a new NfsTLSInterfaceGetParams object
// with the ability to set a context for a request.
func NewNfsTLSInterfaceGetParamsWithContext(ctx context.Context) *NfsTLSInterfaceGetParams {
	return &NfsTLSInterfaceGetParams{
		Context: ctx,
	}
}

// NewNfsTLSInterfaceGetParamsWithHTTPClient creates a new NfsTLSInterfaceGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewNfsTLSInterfaceGetParamsWithHTTPClient(client *http.Client) *NfsTLSInterfaceGetParams {
	return &NfsTLSInterfaceGetParams{
		HTTPClient: client,
	}
}

/*
NfsTLSInterfaceGetParams contains all the parameters to send to the API endpoint

	for the nfs tls interface get operation.

	Typically these are written to a http.Request.
*/
type NfsTLSInterfaceGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* InterfaceUUID.

	   Network interface UUID.
	*/
	InterfaceUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the nfs tls interface get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NfsTLSInterfaceGetParams) WithDefaults() *NfsTLSInterfaceGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the nfs tls interface get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NfsTLSInterfaceGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the nfs tls interface get params
func (o *NfsTLSInterfaceGetParams) WithTimeout(timeout time.Duration) *NfsTLSInterfaceGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the nfs tls interface get params
func (o *NfsTLSInterfaceGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the nfs tls interface get params
func (o *NfsTLSInterfaceGetParams) WithContext(ctx context.Context) *NfsTLSInterfaceGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the nfs tls interface get params
func (o *NfsTLSInterfaceGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the nfs tls interface get params
func (o *NfsTLSInterfaceGetParams) WithHTTPClient(client *http.Client) *NfsTLSInterfaceGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the nfs tls interface get params
func (o *NfsTLSInterfaceGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the nfs tls interface get params
func (o *NfsTLSInterfaceGetParams) WithFields(fields []string) *NfsTLSInterfaceGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the nfs tls interface get params
func (o *NfsTLSInterfaceGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithInterfaceUUID adds the interfaceUUID to the nfs tls interface get params
func (o *NfsTLSInterfaceGetParams) WithInterfaceUUID(interfaceUUID string) *NfsTLSInterfaceGetParams {
	o.SetInterfaceUUID(interfaceUUID)
	return o
}

// SetInterfaceUUID adds the interfaceUuid to the nfs tls interface get params
func (o *NfsTLSInterfaceGetParams) SetInterfaceUUID(interfaceUUID string) {
	o.InterfaceUUID = interfaceUUID
}

// WriteToRequest writes these params to a swagger request
func (o *NfsTLSInterfaceGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Fields != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	// path param interface.uuid
	if err := r.SetPathParam("interface.uuid", o.InterfaceUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamNfsTLSInterfaceGet binds the parameter fields
func (o *NfsTLSInterfaceGetParams) bindParamFields(formats strfmt.Registry) []string {
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