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

// NewExportRuleClientsDeleteCollectionParams creates a new ExportRuleClientsDeleteCollectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExportRuleClientsDeleteCollectionParams() *ExportRuleClientsDeleteCollectionParams {
	return &ExportRuleClientsDeleteCollectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExportRuleClientsDeleteCollectionParamsWithTimeout creates a new ExportRuleClientsDeleteCollectionParams object
// with the ability to set a timeout on a request.
func NewExportRuleClientsDeleteCollectionParamsWithTimeout(timeout time.Duration) *ExportRuleClientsDeleteCollectionParams {
	return &ExportRuleClientsDeleteCollectionParams{
		timeout: timeout,
	}
}

// NewExportRuleClientsDeleteCollectionParamsWithContext creates a new ExportRuleClientsDeleteCollectionParams object
// with the ability to set a context for a request.
func NewExportRuleClientsDeleteCollectionParamsWithContext(ctx context.Context) *ExportRuleClientsDeleteCollectionParams {
	return &ExportRuleClientsDeleteCollectionParams{
		Context: ctx,
	}
}

// NewExportRuleClientsDeleteCollectionParamsWithHTTPClient creates a new ExportRuleClientsDeleteCollectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewExportRuleClientsDeleteCollectionParamsWithHTTPClient(client *http.Client) *ExportRuleClientsDeleteCollectionParams {
	return &ExportRuleClientsDeleteCollectionParams{
		HTTPClient: client,
	}
}

/*
ExportRuleClientsDeleteCollectionParams contains all the parameters to send to the API endpoint

	for the export rule clients delete collection operation.

	Typically these are written to a http.Request.
*/
type ExportRuleClientsDeleteCollectionParams struct {

	/* ContinueOnFailure.

	   Continue even when the operation fails on one of the records.
	*/
	ContinueOnFailure *bool

	/* Index.

	   Export Rule Index
	*/
	Index int64

	/* Info.

	   Info specification
	*/
	Info ExportRuleClientsDeleteCollectionBody

	/* PolicyID.

	   Export Policy ID
	*/
	PolicyID int64

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

// WithDefaults hydrates default values in the export rule clients delete collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExportRuleClientsDeleteCollectionParams) WithDefaults() *ExportRuleClientsDeleteCollectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the export rule clients delete collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExportRuleClientsDeleteCollectionParams) SetDefaults() {
	var (
		continueOnFailureDefault = bool(false)

		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)

		serialRecordsDefault = bool(false)
	)

	val := ExportRuleClientsDeleteCollectionParams{
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

// WithTimeout adds the timeout to the export rule clients delete collection params
func (o *ExportRuleClientsDeleteCollectionParams) WithTimeout(timeout time.Duration) *ExportRuleClientsDeleteCollectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the export rule clients delete collection params
func (o *ExportRuleClientsDeleteCollectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the export rule clients delete collection params
func (o *ExportRuleClientsDeleteCollectionParams) WithContext(ctx context.Context) *ExportRuleClientsDeleteCollectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the export rule clients delete collection params
func (o *ExportRuleClientsDeleteCollectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the export rule clients delete collection params
func (o *ExportRuleClientsDeleteCollectionParams) WithHTTPClient(client *http.Client) *ExportRuleClientsDeleteCollectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the export rule clients delete collection params
func (o *ExportRuleClientsDeleteCollectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContinueOnFailure adds the continueOnFailure to the export rule clients delete collection params
func (o *ExportRuleClientsDeleteCollectionParams) WithContinueOnFailure(continueOnFailure *bool) *ExportRuleClientsDeleteCollectionParams {
	o.SetContinueOnFailure(continueOnFailure)
	return o
}

// SetContinueOnFailure adds the continueOnFailure to the export rule clients delete collection params
func (o *ExportRuleClientsDeleteCollectionParams) SetContinueOnFailure(continueOnFailure *bool) {
	o.ContinueOnFailure = continueOnFailure
}

// WithIndex adds the index to the export rule clients delete collection params
func (o *ExportRuleClientsDeleteCollectionParams) WithIndex(index int64) *ExportRuleClientsDeleteCollectionParams {
	o.SetIndex(index)
	return o
}

// SetIndex adds the index to the export rule clients delete collection params
func (o *ExportRuleClientsDeleteCollectionParams) SetIndex(index int64) {
	o.Index = index
}

// WithInfo adds the info to the export rule clients delete collection params
func (o *ExportRuleClientsDeleteCollectionParams) WithInfo(info ExportRuleClientsDeleteCollectionBody) *ExportRuleClientsDeleteCollectionParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the export rule clients delete collection params
func (o *ExportRuleClientsDeleteCollectionParams) SetInfo(info ExportRuleClientsDeleteCollectionBody) {
	o.Info = info
}

// WithPolicyID adds the policyID to the export rule clients delete collection params
func (o *ExportRuleClientsDeleteCollectionParams) WithPolicyID(policyID int64) *ExportRuleClientsDeleteCollectionParams {
	o.SetPolicyID(policyID)
	return o
}

// SetPolicyID adds the policyId to the export rule clients delete collection params
func (o *ExportRuleClientsDeleteCollectionParams) SetPolicyID(policyID int64) {
	o.PolicyID = policyID
}

// WithReturnRecords adds the returnRecords to the export rule clients delete collection params
func (o *ExportRuleClientsDeleteCollectionParams) WithReturnRecords(returnRecords *bool) *ExportRuleClientsDeleteCollectionParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the export rule clients delete collection params
func (o *ExportRuleClientsDeleteCollectionParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the export rule clients delete collection params
func (o *ExportRuleClientsDeleteCollectionParams) WithReturnTimeout(returnTimeout *int64) *ExportRuleClientsDeleteCollectionParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the export rule clients delete collection params
func (o *ExportRuleClientsDeleteCollectionParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSerialRecords adds the serialRecords to the export rule clients delete collection params
func (o *ExportRuleClientsDeleteCollectionParams) WithSerialRecords(serialRecords *bool) *ExportRuleClientsDeleteCollectionParams {
	o.SetSerialRecords(serialRecords)
	return o
}

// SetSerialRecords adds the serialRecords to the export rule clients delete collection params
func (o *ExportRuleClientsDeleteCollectionParams) SetSerialRecords(serialRecords *bool) {
	o.SerialRecords = serialRecords
}

// WriteToRequest writes these params to a swagger request
func (o *ExportRuleClientsDeleteCollectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param index
	if err := r.SetPathParam("index", swag.FormatInt64(o.Index)); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.Info); err != nil {
		return err
	}

	// path param policy.id
	if err := r.SetPathParam("policy.id", swag.FormatInt64(o.PolicyID)); err != nil {
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