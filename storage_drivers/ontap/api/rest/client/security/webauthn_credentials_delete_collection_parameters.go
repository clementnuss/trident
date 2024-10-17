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
)

// NewWebauthnCredentialsDeleteCollectionParams creates a new WebauthnCredentialsDeleteCollectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWebauthnCredentialsDeleteCollectionParams() *WebauthnCredentialsDeleteCollectionParams {
	return &WebauthnCredentialsDeleteCollectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWebauthnCredentialsDeleteCollectionParamsWithTimeout creates a new WebauthnCredentialsDeleteCollectionParams object
// with the ability to set a timeout on a request.
func NewWebauthnCredentialsDeleteCollectionParamsWithTimeout(timeout time.Duration) *WebauthnCredentialsDeleteCollectionParams {
	return &WebauthnCredentialsDeleteCollectionParams{
		timeout: timeout,
	}
}

// NewWebauthnCredentialsDeleteCollectionParamsWithContext creates a new WebauthnCredentialsDeleteCollectionParams object
// with the ability to set a context for a request.
func NewWebauthnCredentialsDeleteCollectionParamsWithContext(ctx context.Context) *WebauthnCredentialsDeleteCollectionParams {
	return &WebauthnCredentialsDeleteCollectionParams{
		Context: ctx,
	}
}

// NewWebauthnCredentialsDeleteCollectionParamsWithHTTPClient creates a new WebauthnCredentialsDeleteCollectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewWebauthnCredentialsDeleteCollectionParamsWithHTTPClient(client *http.Client) *WebauthnCredentialsDeleteCollectionParams {
	return &WebauthnCredentialsDeleteCollectionParams{
		HTTPClient: client,
	}
}

/*
WebauthnCredentialsDeleteCollectionParams contains all the parameters to send to the API endpoint

	for the webauthn credentials delete collection operation.

	Typically these are written to a http.Request.
*/
type WebauthnCredentialsDeleteCollectionParams struct {

	/* ContinueOnFailure.

	   Continue even when the operation fails on one of the records.
	*/
	ContinueOnFailure *bool

	/* Info.

	   Info specification
	*/
	Info WebauthnCredentialsDeleteCollectionBody

	/* ReturnRecords.

	   The default is true for GET calls.  When set to false, only the number of records is returned.

	   Default: true
	*/
	ReturnRecords *bool

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached.

	   Default: 15
	*/
	ReturnTimeout *int64

	/* SerialRecords.

	   Perform the operation on the records synchronously.
	*/
	SerialRecords *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the webauthn credentials delete collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WebauthnCredentialsDeleteCollectionParams) WithDefaults() *WebauthnCredentialsDeleteCollectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the webauthn credentials delete collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WebauthnCredentialsDeleteCollectionParams) SetDefaults() {
	var (
		continueOnFailureDefault = bool(false)

		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)

		serialRecordsDefault = bool(false)
	)

	val := WebauthnCredentialsDeleteCollectionParams{
		ContinueOnFailure: &continueOnFailureDefault,
		ReturnRecords:     &returnRecordsDefault,
		ReturnTimeout:     &returnTimeoutDefault,
		SerialRecords:     &serialRecordsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the webauthn credentials delete collection params
func (o *WebauthnCredentialsDeleteCollectionParams) WithTimeout(timeout time.Duration) *WebauthnCredentialsDeleteCollectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the webauthn credentials delete collection params
func (o *WebauthnCredentialsDeleteCollectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the webauthn credentials delete collection params
func (o *WebauthnCredentialsDeleteCollectionParams) WithContext(ctx context.Context) *WebauthnCredentialsDeleteCollectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the webauthn credentials delete collection params
func (o *WebauthnCredentialsDeleteCollectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the webauthn credentials delete collection params
func (o *WebauthnCredentialsDeleteCollectionParams) WithHTTPClient(client *http.Client) *WebauthnCredentialsDeleteCollectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the webauthn credentials delete collection params
func (o *WebauthnCredentialsDeleteCollectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContinueOnFailure adds the continueOnFailure to the webauthn credentials delete collection params
func (o *WebauthnCredentialsDeleteCollectionParams) WithContinueOnFailure(continueOnFailure *bool) *WebauthnCredentialsDeleteCollectionParams {
	o.SetContinueOnFailure(continueOnFailure)
	return o
}

// SetContinueOnFailure adds the continueOnFailure to the webauthn credentials delete collection params
func (o *WebauthnCredentialsDeleteCollectionParams) SetContinueOnFailure(continueOnFailure *bool) {
	o.ContinueOnFailure = continueOnFailure
}

// WithInfo adds the info to the webauthn credentials delete collection params
func (o *WebauthnCredentialsDeleteCollectionParams) WithInfo(info WebauthnCredentialsDeleteCollectionBody) *WebauthnCredentialsDeleteCollectionParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the webauthn credentials delete collection params
func (o *WebauthnCredentialsDeleteCollectionParams) SetInfo(info WebauthnCredentialsDeleteCollectionBody) {
	o.Info = info
}

// WithReturnRecords adds the returnRecords to the webauthn credentials delete collection params
func (o *WebauthnCredentialsDeleteCollectionParams) WithReturnRecords(returnRecords *bool) *WebauthnCredentialsDeleteCollectionParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the webauthn credentials delete collection params
func (o *WebauthnCredentialsDeleteCollectionParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the webauthn credentials delete collection params
func (o *WebauthnCredentialsDeleteCollectionParams) WithReturnTimeout(returnTimeout *int64) *WebauthnCredentialsDeleteCollectionParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the webauthn credentials delete collection params
func (o *WebauthnCredentialsDeleteCollectionParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSerialRecords adds the serialRecords to the webauthn credentials delete collection params
func (o *WebauthnCredentialsDeleteCollectionParams) WithSerialRecords(serialRecords *bool) *WebauthnCredentialsDeleteCollectionParams {
	o.SetSerialRecords(serialRecords)
	return o
}

// SetSerialRecords adds the serialRecords to the webauthn credentials delete collection params
func (o *WebauthnCredentialsDeleteCollectionParams) SetSerialRecords(serialRecords *bool) {
	o.SerialRecords = serialRecords
}

// WriteToRequest writes these params to a swagger request
func (o *WebauthnCredentialsDeleteCollectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ContinueOnFailure != nil {

		// query param continue_on_failure
		var qrContinueOnFailure bool

		if o.ContinueOnFailure != nil {
			qrContinueOnFailure = *o.ContinueOnFailure
		}
		qContinueOnFailure := swag.FormatBool(qrContinueOnFailure)
		if qContinueOnFailure != "" {

			if err := r.SetQueryParam("continue_on_failure", qContinueOnFailure); err != nil {
				return err
			}
		}
	}
	if err := r.SetBodyParam(o.Info); err != nil {
		return err
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

	if o.ReturnTimeout != nil {

		// query param return_timeout
		var qrReturnTimeout int64

		if o.ReturnTimeout != nil {
			qrReturnTimeout = *o.ReturnTimeout
		}
		qReturnTimeout := swag.FormatInt64(qrReturnTimeout)
		if qReturnTimeout != "" {

			if err := r.SetQueryParam("return_timeout", qReturnTimeout); err != nil {
				return err
			}
		}
	}

	if o.SerialRecords != nil {

		// query param serial_records
		var qrSerialRecords bool

		if o.SerialRecords != nil {
			qrSerialRecords = *o.SerialRecords
		}
		qSerialRecords := swag.FormatBool(qrSerialRecords)
		if qSerialRecords != "" {

			if err := r.SetQueryParam("serial_records", qSerialRecords); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}