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

// NewSecurityExternalRoleMappingModifyCollectionParams creates a new SecurityExternalRoleMappingModifyCollectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSecurityExternalRoleMappingModifyCollectionParams() *SecurityExternalRoleMappingModifyCollectionParams {
	return &SecurityExternalRoleMappingModifyCollectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSecurityExternalRoleMappingModifyCollectionParamsWithTimeout creates a new SecurityExternalRoleMappingModifyCollectionParams object
// with the ability to set a timeout on a request.
func NewSecurityExternalRoleMappingModifyCollectionParamsWithTimeout(timeout time.Duration) *SecurityExternalRoleMappingModifyCollectionParams {
	return &SecurityExternalRoleMappingModifyCollectionParams{
		timeout: timeout,
	}
}

// NewSecurityExternalRoleMappingModifyCollectionParamsWithContext creates a new SecurityExternalRoleMappingModifyCollectionParams object
// with the ability to set a context for a request.
func NewSecurityExternalRoleMappingModifyCollectionParamsWithContext(ctx context.Context) *SecurityExternalRoleMappingModifyCollectionParams {
	return &SecurityExternalRoleMappingModifyCollectionParams{
		Context: ctx,
	}
}

// NewSecurityExternalRoleMappingModifyCollectionParamsWithHTTPClient creates a new SecurityExternalRoleMappingModifyCollectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewSecurityExternalRoleMappingModifyCollectionParamsWithHTTPClient(client *http.Client) *SecurityExternalRoleMappingModifyCollectionParams {
	return &SecurityExternalRoleMappingModifyCollectionParams{
		HTTPClient: client,
	}
}

/*
SecurityExternalRoleMappingModifyCollectionParams contains all the parameters to send to the API endpoint

	for the security external role mapping modify collection operation.

	Typically these are written to a http.Request.
*/
type SecurityExternalRoleMappingModifyCollectionParams struct {

	/* Comment.

	   Filter by comment
	*/
	Comment *string

	/* ContinueOnFailure.

	   Continue even when the operation fails on one of the records.
	*/
	ContinueOnFailure *bool

	/* ExternalRole.

	   Filter by external_role
	*/
	ExternalRole *string

	/* Info.

	   Info specification
	*/
	Info SecurityExternalRoleMappingModifyCollectionBody

	/* OntapRoleName.

	   Filter by ontap_role.name
	*/
	OntapRoleName *string

	/* Provider.

	   Filter by provider
	*/
	Provider *string

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

	/* Timestamp.

	   Filter by timestamp
	*/
	Timestamp *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the security external role mapping modify collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityExternalRoleMappingModifyCollectionParams) WithDefaults() *SecurityExternalRoleMappingModifyCollectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the security external role mapping modify collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityExternalRoleMappingModifyCollectionParams) SetDefaults() {
	var (
		continueOnFailureDefault = bool(false)

		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)

		serialRecordsDefault = bool(false)
	)

	val := SecurityExternalRoleMappingModifyCollectionParams{
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

// WithTimeout adds the timeout to the security external role mapping modify collection params
func (o *SecurityExternalRoleMappingModifyCollectionParams) WithTimeout(timeout time.Duration) *SecurityExternalRoleMappingModifyCollectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the security external role mapping modify collection params
func (o *SecurityExternalRoleMappingModifyCollectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the security external role mapping modify collection params
func (o *SecurityExternalRoleMappingModifyCollectionParams) WithContext(ctx context.Context) *SecurityExternalRoleMappingModifyCollectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the security external role mapping modify collection params
func (o *SecurityExternalRoleMappingModifyCollectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the security external role mapping modify collection params
func (o *SecurityExternalRoleMappingModifyCollectionParams) WithHTTPClient(client *http.Client) *SecurityExternalRoleMappingModifyCollectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the security external role mapping modify collection params
func (o *SecurityExternalRoleMappingModifyCollectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithComment adds the comment to the security external role mapping modify collection params
func (o *SecurityExternalRoleMappingModifyCollectionParams) WithComment(comment *string) *SecurityExternalRoleMappingModifyCollectionParams {
	o.SetComment(comment)
	return o
}

// SetComment adds the comment to the security external role mapping modify collection params
func (o *SecurityExternalRoleMappingModifyCollectionParams) SetComment(comment *string) {
	o.Comment = comment
}

// WithContinueOnFailure adds the continueOnFailure to the security external role mapping modify collection params
func (o *SecurityExternalRoleMappingModifyCollectionParams) WithContinueOnFailure(continueOnFailure *bool) *SecurityExternalRoleMappingModifyCollectionParams {
	o.SetContinueOnFailure(continueOnFailure)
	return o
}

// SetContinueOnFailure adds the continueOnFailure to the security external role mapping modify collection params
func (o *SecurityExternalRoleMappingModifyCollectionParams) SetContinueOnFailure(continueOnFailure *bool) {
	o.ContinueOnFailure = continueOnFailure
}

// WithExternalRole adds the externalRole to the security external role mapping modify collection params
func (o *SecurityExternalRoleMappingModifyCollectionParams) WithExternalRole(externalRole *string) *SecurityExternalRoleMappingModifyCollectionParams {
	o.SetExternalRole(externalRole)
	return o
}

// SetExternalRole adds the externalRole to the security external role mapping modify collection params
func (o *SecurityExternalRoleMappingModifyCollectionParams) SetExternalRole(externalRole *string) {
	o.ExternalRole = externalRole
}

// WithInfo adds the info to the security external role mapping modify collection params
func (o *SecurityExternalRoleMappingModifyCollectionParams) WithInfo(info SecurityExternalRoleMappingModifyCollectionBody) *SecurityExternalRoleMappingModifyCollectionParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the security external role mapping modify collection params
func (o *SecurityExternalRoleMappingModifyCollectionParams) SetInfo(info SecurityExternalRoleMappingModifyCollectionBody) {
	o.Info = info
}

// WithOntapRoleName adds the ontapRoleName to the security external role mapping modify collection params
func (o *SecurityExternalRoleMappingModifyCollectionParams) WithOntapRoleName(ontapRoleName *string) *SecurityExternalRoleMappingModifyCollectionParams {
	o.SetOntapRoleName(ontapRoleName)
	return o
}

// SetOntapRoleName adds the ontapRoleName to the security external role mapping modify collection params
func (o *SecurityExternalRoleMappingModifyCollectionParams) SetOntapRoleName(ontapRoleName *string) {
	o.OntapRoleName = ontapRoleName
}

// WithProvider adds the provider to the security external role mapping modify collection params
func (o *SecurityExternalRoleMappingModifyCollectionParams) WithProvider(provider *string) *SecurityExternalRoleMappingModifyCollectionParams {
	o.SetProvider(provider)
	return o
}

// SetProvider adds the provider to the security external role mapping modify collection params
func (o *SecurityExternalRoleMappingModifyCollectionParams) SetProvider(provider *string) {
	o.Provider = provider
}

// WithReturnRecords adds the returnRecords to the security external role mapping modify collection params
func (o *SecurityExternalRoleMappingModifyCollectionParams) WithReturnRecords(returnRecords *bool) *SecurityExternalRoleMappingModifyCollectionParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the security external role mapping modify collection params
func (o *SecurityExternalRoleMappingModifyCollectionParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the security external role mapping modify collection params
func (o *SecurityExternalRoleMappingModifyCollectionParams) WithReturnTimeout(returnTimeout *int64) *SecurityExternalRoleMappingModifyCollectionParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the security external role mapping modify collection params
func (o *SecurityExternalRoleMappingModifyCollectionParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSerialRecords adds the serialRecords to the security external role mapping modify collection params
func (o *SecurityExternalRoleMappingModifyCollectionParams) WithSerialRecords(serialRecords *bool) *SecurityExternalRoleMappingModifyCollectionParams {
	o.SetSerialRecords(serialRecords)
	return o
}

// SetSerialRecords adds the serialRecords to the security external role mapping modify collection params
func (o *SecurityExternalRoleMappingModifyCollectionParams) SetSerialRecords(serialRecords *bool) {
	o.SerialRecords = serialRecords
}

// WithTimestamp adds the timestamp to the security external role mapping modify collection params
func (o *SecurityExternalRoleMappingModifyCollectionParams) WithTimestamp(timestamp *string) *SecurityExternalRoleMappingModifyCollectionParams {
	o.SetTimestamp(timestamp)
	return o
}

// SetTimestamp adds the timestamp to the security external role mapping modify collection params
func (o *SecurityExternalRoleMappingModifyCollectionParams) SetTimestamp(timestamp *string) {
	o.Timestamp = timestamp
}

// WriteToRequest writes these params to a swagger request
func (o *SecurityExternalRoleMappingModifyCollectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Comment != nil {

		// query param comment
		var qrComment string

		if o.Comment != nil {
			qrComment = *o.Comment
		}
		qComment := qrComment
		if qComment != "" {

			if err := r.SetQueryParam("comment", qComment); err != nil {
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

	if o.ExternalRole != nil {

		// query param external_role
		var qrExternalRole string

		if o.ExternalRole != nil {
			qrExternalRole = *o.ExternalRole
		}
		qExternalRole := qrExternalRole
		if qExternalRole != "" {

			if err := r.SetQueryParam("external_role", qExternalRole); err != nil {
				return err
			}
		}
	}
	if err := r.SetBodyParam(o.Info); err != nil {
		return err
	}

	if o.OntapRoleName != nil {

		// query param ontap_role.name
		var qrOntapRoleName string

		if o.OntapRoleName != nil {
			qrOntapRoleName = *o.OntapRoleName
		}
		qOntapRoleName := qrOntapRoleName
		if qOntapRoleName != "" {

			if err := r.SetQueryParam("ontap_role.name", qOntapRoleName); err != nil {
				return err
			}
		}
	}

	if o.Provider != nil {

		// query param provider
		var qrProvider string

		if o.Provider != nil {
			qrProvider = *o.Provider
		}
		qProvider := qrProvider
		if qProvider != "" {

			if err := r.SetQueryParam("provider", qProvider); err != nil {
				return err
			}
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

	if o.Timestamp != nil {

		// query param timestamp
		var qrTimestamp string

		if o.Timestamp != nil {
			qrTimestamp = *o.Timestamp
		}
		qTimestamp := qrTimestamp
		if qTimestamp != "" {

			if err := r.SetQueryParam("timestamp", qTimestamp); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}