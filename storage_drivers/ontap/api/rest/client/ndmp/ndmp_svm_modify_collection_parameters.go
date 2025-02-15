// Code generated by go-swagger; DO NOT EDIT.

package ndmp

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

// NewNdmpSvmModifyCollectionParams creates a new NdmpSvmModifyCollectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNdmpSvmModifyCollectionParams() *NdmpSvmModifyCollectionParams {
	return &NdmpSvmModifyCollectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNdmpSvmModifyCollectionParamsWithTimeout creates a new NdmpSvmModifyCollectionParams object
// with the ability to set a timeout on a request.
func NewNdmpSvmModifyCollectionParamsWithTimeout(timeout time.Duration) *NdmpSvmModifyCollectionParams {
	return &NdmpSvmModifyCollectionParams{
		timeout: timeout,
	}
}

// NewNdmpSvmModifyCollectionParamsWithContext creates a new NdmpSvmModifyCollectionParams object
// with the ability to set a context for a request.
func NewNdmpSvmModifyCollectionParamsWithContext(ctx context.Context) *NdmpSvmModifyCollectionParams {
	return &NdmpSvmModifyCollectionParams{
		Context: ctx,
	}
}

// NewNdmpSvmModifyCollectionParamsWithHTTPClient creates a new NdmpSvmModifyCollectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewNdmpSvmModifyCollectionParamsWithHTTPClient(client *http.Client) *NdmpSvmModifyCollectionParams {
	return &NdmpSvmModifyCollectionParams{
		HTTPClient: client,
	}
}

/*
NdmpSvmModifyCollectionParams contains all the parameters to send to the API endpoint

	for the ndmp svm modify collection operation.

	Typically these are written to a http.Request.
*/
type NdmpSvmModifyCollectionParams struct {

	/* AuthenticationTypes.

	   Filter by authentication_types
	*/
	AuthenticationTypes *string

	/* ContinueOnFailure.

	   Continue even when the operation fails on one of the records.
	*/
	ContinueOnFailure *bool

	/* Enabled.

	   Filter by enabled
	*/
	Enabled *bool

	/* Info.

	   Info specification
	*/
	Info NdmpSvmModifyCollectionBody

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

	/* SvmName.

	   Filter by svm.name
	*/
	SvmName *string

	/* SvmUUID.

	   Filter by svm.uuid
	*/
	SvmUUID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ndmp svm modify collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NdmpSvmModifyCollectionParams) WithDefaults() *NdmpSvmModifyCollectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ndmp svm modify collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NdmpSvmModifyCollectionParams) SetDefaults() {
	var (
		continueOnFailureDefault = bool(false)

		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)

		serialRecordsDefault = bool(false)
	)

	val := NdmpSvmModifyCollectionParams{
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

// WithTimeout adds the timeout to the ndmp svm modify collection params
func (o *NdmpSvmModifyCollectionParams) WithTimeout(timeout time.Duration) *NdmpSvmModifyCollectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ndmp svm modify collection params
func (o *NdmpSvmModifyCollectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ndmp svm modify collection params
func (o *NdmpSvmModifyCollectionParams) WithContext(ctx context.Context) *NdmpSvmModifyCollectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ndmp svm modify collection params
func (o *NdmpSvmModifyCollectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ndmp svm modify collection params
func (o *NdmpSvmModifyCollectionParams) WithHTTPClient(client *http.Client) *NdmpSvmModifyCollectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ndmp svm modify collection params
func (o *NdmpSvmModifyCollectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthenticationTypes adds the authenticationTypes to the ndmp svm modify collection params
func (o *NdmpSvmModifyCollectionParams) WithAuthenticationTypes(authenticationTypes *string) *NdmpSvmModifyCollectionParams {
	o.SetAuthenticationTypes(authenticationTypes)
	return o
}

// SetAuthenticationTypes adds the authenticationTypes to the ndmp svm modify collection params
func (o *NdmpSvmModifyCollectionParams) SetAuthenticationTypes(authenticationTypes *string) {
	o.AuthenticationTypes = authenticationTypes
}

// WithContinueOnFailure adds the continueOnFailure to the ndmp svm modify collection params
func (o *NdmpSvmModifyCollectionParams) WithContinueOnFailure(continueOnFailure *bool) *NdmpSvmModifyCollectionParams {
	o.SetContinueOnFailure(continueOnFailure)
	return o
}

// SetContinueOnFailure adds the continueOnFailure to the ndmp svm modify collection params
func (o *NdmpSvmModifyCollectionParams) SetContinueOnFailure(continueOnFailure *bool) {
	o.ContinueOnFailure = continueOnFailure
}

// WithEnabled adds the enabled to the ndmp svm modify collection params
func (o *NdmpSvmModifyCollectionParams) WithEnabled(enabled *bool) *NdmpSvmModifyCollectionParams {
	o.SetEnabled(enabled)
	return o
}

// SetEnabled adds the enabled to the ndmp svm modify collection params
func (o *NdmpSvmModifyCollectionParams) SetEnabled(enabled *bool) {
	o.Enabled = enabled
}

// WithInfo adds the info to the ndmp svm modify collection params
func (o *NdmpSvmModifyCollectionParams) WithInfo(info NdmpSvmModifyCollectionBody) *NdmpSvmModifyCollectionParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the ndmp svm modify collection params
func (o *NdmpSvmModifyCollectionParams) SetInfo(info NdmpSvmModifyCollectionBody) {
	o.Info = info
}

// WithReturnRecords adds the returnRecords to the ndmp svm modify collection params
func (o *NdmpSvmModifyCollectionParams) WithReturnRecords(returnRecords *bool) *NdmpSvmModifyCollectionParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the ndmp svm modify collection params
func (o *NdmpSvmModifyCollectionParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the ndmp svm modify collection params
func (o *NdmpSvmModifyCollectionParams) WithReturnTimeout(returnTimeout *int64) *NdmpSvmModifyCollectionParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the ndmp svm modify collection params
func (o *NdmpSvmModifyCollectionParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSerialRecords adds the serialRecords to the ndmp svm modify collection params
func (o *NdmpSvmModifyCollectionParams) WithSerialRecords(serialRecords *bool) *NdmpSvmModifyCollectionParams {
	o.SetSerialRecords(serialRecords)
	return o
}

// SetSerialRecords adds the serialRecords to the ndmp svm modify collection params
func (o *NdmpSvmModifyCollectionParams) SetSerialRecords(serialRecords *bool) {
	o.SerialRecords = serialRecords
}

// WithSvmName adds the svmName to the ndmp svm modify collection params
func (o *NdmpSvmModifyCollectionParams) WithSvmName(svmName *string) *NdmpSvmModifyCollectionParams {
	o.SetSvmName(svmName)
	return o
}

// SetSvmName adds the svmName to the ndmp svm modify collection params
func (o *NdmpSvmModifyCollectionParams) SetSvmName(svmName *string) {
	o.SvmName = svmName
}

// WithSvmUUID adds the svmUUID to the ndmp svm modify collection params
func (o *NdmpSvmModifyCollectionParams) WithSvmUUID(svmUUID *string) *NdmpSvmModifyCollectionParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the ndmp svm modify collection params
func (o *NdmpSvmModifyCollectionParams) SetSvmUUID(svmUUID *string) {
	o.SvmUUID = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *NdmpSvmModifyCollectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AuthenticationTypes != nil {

		// query param authentication_types
		var qrAuthenticationTypes string

		if o.AuthenticationTypes != nil {
			qrAuthenticationTypes = *o.AuthenticationTypes
		}
		qAuthenticationTypes := qrAuthenticationTypes
		if qAuthenticationTypes != "" {

			if err := r.SetQueryParam("authentication_types", qAuthenticationTypes); err != nil {
				return err
			}
		}
	}

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

	if o.Enabled != nil {

		// query param enabled
		var qrEnabled bool

		if o.Enabled != nil {
			qrEnabled = *o.Enabled
		}
		qEnabled := swag.FormatBool(qrEnabled)
		if qEnabled != "" {

			if err := r.SetQueryParam("enabled", qEnabled); err != nil {
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

	if o.SvmName != nil {

		// query param svm.name
		var qrSvmName string

		if o.SvmName != nil {
			qrSvmName = *o.SvmName
		}
		qSvmName := qrSvmName
		if qSvmName != "" {

			if err := r.SetQueryParam("svm.name", qSvmName); err != nil {
				return err
			}
		}
	}

	if o.SvmUUID != nil {

		// query param svm.uuid
		var qrSvmUUID string

		if o.SvmUUID != nil {
			qrSvmUUID = *o.SvmUUID
		}
		qSvmUUID := qrSvmUUID
		if qSvmUUID != "" {

			if err := r.SetQueryParam("svm.uuid", qSvmUUID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
