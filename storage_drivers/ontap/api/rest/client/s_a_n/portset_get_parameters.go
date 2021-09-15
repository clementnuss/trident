// Code generated by go-swagger; DO NOT EDIT.

package s_a_n

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

// NewPortsetGetParams creates a new PortsetGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPortsetGetParams() *PortsetGetParams {
	return &PortsetGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPortsetGetParamsWithTimeout creates a new PortsetGetParams object
// with the ability to set a timeout on a request.
func NewPortsetGetParamsWithTimeout(timeout time.Duration) *PortsetGetParams {
	return &PortsetGetParams{
		timeout: timeout,
	}
}

// NewPortsetGetParamsWithContext creates a new PortsetGetParams object
// with the ability to set a context for a request.
func NewPortsetGetParamsWithContext(ctx context.Context) *PortsetGetParams {
	return &PortsetGetParams{
		Context: ctx,
	}
}

// NewPortsetGetParamsWithHTTPClient creates a new PortsetGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewPortsetGetParamsWithHTTPClient(client *http.Client) *PortsetGetParams {
	return &PortsetGetParams{
		HTTPClient: client,
	}
}

/* PortsetGetParams contains all the parameters to send to the API endpoint
   for the portset get operation.

   Typically these are written to a http.Request.
*/
type PortsetGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* UUID.

	   The unique identifier of the portset.

	*/
	UUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the portset get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PortsetGetParams) WithDefaults() *PortsetGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the portset get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PortsetGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the portset get params
func (o *PortsetGetParams) WithTimeout(timeout time.Duration) *PortsetGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the portset get params
func (o *PortsetGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the portset get params
func (o *PortsetGetParams) WithContext(ctx context.Context) *PortsetGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the portset get params
func (o *PortsetGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the portset get params
func (o *PortsetGetParams) WithHTTPClient(client *http.Client) *PortsetGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the portset get params
func (o *PortsetGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFieldsQueryParameter adds the fields to the portset get params
func (o *PortsetGetParams) WithFieldsQueryParameter(fields []string) *PortsetGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the portset get params
func (o *PortsetGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithUUIDPathParameter adds the uuid to the portset get params
func (o *PortsetGetParams) WithUUIDPathParameter(uuid string) *PortsetGetParams {
	o.SetUUIDPathParameter(uuid)
	return o
}

// SetUUIDPathParameter adds the uuid to the portset get params
func (o *PortsetGetParams) SetUUIDPathParameter(uuid string) {
	o.UUIDPathParameter = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *PortsetGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FieldsQueryParameter != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUIDPathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamPortsetGet binds the parameter fields
func (o *PortsetGetParams) bindParamFields(formats strfmt.Registry) []string {
	fieldsIR := o.FieldsQueryParameter

	var fieldsIC []string
	for _, fieldsIIR := range fieldsIR { // explode []string

		fieldsIIV := fieldsIIR // string as string
		fieldsIC = append(fieldsIC, fieldsIIV)
	}

	// items.CollectionFormat: "csv"
	fieldsIS := swag.JoinByFormat(fieldsIC, "csv")

	return fieldsIS
}