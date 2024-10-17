// Code generated by go-swagger; DO NOT EDIT.

package name_services

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

// NewUnixGroupUserDeleteCollectionParams creates a new UnixGroupUserDeleteCollectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUnixGroupUserDeleteCollectionParams() *UnixGroupUserDeleteCollectionParams {
	return &UnixGroupUserDeleteCollectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUnixGroupUserDeleteCollectionParamsWithTimeout creates a new UnixGroupUserDeleteCollectionParams object
// with the ability to set a timeout on a request.
func NewUnixGroupUserDeleteCollectionParamsWithTimeout(timeout time.Duration) *UnixGroupUserDeleteCollectionParams {
	return &UnixGroupUserDeleteCollectionParams{
		timeout: timeout,
	}
}

// NewUnixGroupUserDeleteCollectionParamsWithContext creates a new UnixGroupUserDeleteCollectionParams object
// with the ability to set a context for a request.
func NewUnixGroupUserDeleteCollectionParamsWithContext(ctx context.Context) *UnixGroupUserDeleteCollectionParams {
	return &UnixGroupUserDeleteCollectionParams{
		Context: ctx,
	}
}

// NewUnixGroupUserDeleteCollectionParamsWithHTTPClient creates a new UnixGroupUserDeleteCollectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewUnixGroupUserDeleteCollectionParamsWithHTTPClient(client *http.Client) *UnixGroupUserDeleteCollectionParams {
	return &UnixGroupUserDeleteCollectionParams{
		HTTPClient: client,
	}
}

/*
UnixGroupUserDeleteCollectionParams contains all the parameters to send to the API endpoint

	for the unix group user delete collection operation.

	Typically these are written to a http.Request.
*/
type UnixGroupUserDeleteCollectionParams struct {

	/* ContinueOnFailure.

	   Continue even when the operation fails on one of the records.
	*/
	ContinueOnFailure *bool

	/* Info.

	   Info specification
	*/
	Info UnixGroupUserDeleteCollectionBody

	/* Name.

	   Filter by name
	*/
	Name *string

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

	/* SvmUUID.

	   UUID of the SVM to which this object belongs.
	*/
	SvmUUID string

	/* UnixGroupName.

	   UNIX group name.
	*/
	UnixGroupName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the unix group user delete collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UnixGroupUserDeleteCollectionParams) WithDefaults() *UnixGroupUserDeleteCollectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the unix group user delete collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UnixGroupUserDeleteCollectionParams) SetDefaults() {
	var (
		continueOnFailureDefault = bool(false)

		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)

		serialRecordsDefault = bool(false)
	)

	val := UnixGroupUserDeleteCollectionParams{
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

// WithTimeout adds the timeout to the unix group user delete collection params
func (o *UnixGroupUserDeleteCollectionParams) WithTimeout(timeout time.Duration) *UnixGroupUserDeleteCollectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the unix group user delete collection params
func (o *UnixGroupUserDeleteCollectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the unix group user delete collection params
func (o *UnixGroupUserDeleteCollectionParams) WithContext(ctx context.Context) *UnixGroupUserDeleteCollectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the unix group user delete collection params
func (o *UnixGroupUserDeleteCollectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the unix group user delete collection params
func (o *UnixGroupUserDeleteCollectionParams) WithHTTPClient(client *http.Client) *UnixGroupUserDeleteCollectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the unix group user delete collection params
func (o *UnixGroupUserDeleteCollectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContinueOnFailure adds the continueOnFailure to the unix group user delete collection params
func (o *UnixGroupUserDeleteCollectionParams) WithContinueOnFailure(continueOnFailure *bool) *UnixGroupUserDeleteCollectionParams {
	o.SetContinueOnFailure(continueOnFailure)
	return o
}

// SetContinueOnFailure adds the continueOnFailure to the unix group user delete collection params
func (o *UnixGroupUserDeleteCollectionParams) SetContinueOnFailure(continueOnFailure *bool) {
	o.ContinueOnFailure = continueOnFailure
}

// WithInfo adds the info to the unix group user delete collection params
func (o *UnixGroupUserDeleteCollectionParams) WithInfo(info UnixGroupUserDeleteCollectionBody) *UnixGroupUserDeleteCollectionParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the unix group user delete collection params
func (o *UnixGroupUserDeleteCollectionParams) SetInfo(info UnixGroupUserDeleteCollectionBody) {
	o.Info = info
}

// WithName adds the name to the unix group user delete collection params
func (o *UnixGroupUserDeleteCollectionParams) WithName(name *string) *UnixGroupUserDeleteCollectionParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the unix group user delete collection params
func (o *UnixGroupUserDeleteCollectionParams) SetName(name *string) {
	o.Name = name
}

// WithReturnRecords adds the returnRecords to the unix group user delete collection params
func (o *UnixGroupUserDeleteCollectionParams) WithReturnRecords(returnRecords *bool) *UnixGroupUserDeleteCollectionParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the unix group user delete collection params
func (o *UnixGroupUserDeleteCollectionParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the unix group user delete collection params
func (o *UnixGroupUserDeleteCollectionParams) WithReturnTimeout(returnTimeout *int64) *UnixGroupUserDeleteCollectionParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the unix group user delete collection params
func (o *UnixGroupUserDeleteCollectionParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSerialRecords adds the serialRecords to the unix group user delete collection params
func (o *UnixGroupUserDeleteCollectionParams) WithSerialRecords(serialRecords *bool) *UnixGroupUserDeleteCollectionParams {
	o.SetSerialRecords(serialRecords)
	return o
}

// SetSerialRecords adds the serialRecords to the unix group user delete collection params
func (o *UnixGroupUserDeleteCollectionParams) SetSerialRecords(serialRecords *bool) {
	o.SerialRecords = serialRecords
}

// WithSvmUUID adds the svmUUID to the unix group user delete collection params
func (o *UnixGroupUserDeleteCollectionParams) WithSvmUUID(svmUUID string) *UnixGroupUserDeleteCollectionParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the unix group user delete collection params
func (o *UnixGroupUserDeleteCollectionParams) SetSvmUUID(svmUUID string) {
	o.SvmUUID = svmUUID
}

// WithUnixGroupName adds the unixGroupName to the unix group user delete collection params
func (o *UnixGroupUserDeleteCollectionParams) WithUnixGroupName(unixGroupName string) *UnixGroupUserDeleteCollectionParams {
	o.SetUnixGroupName(unixGroupName)
	return o
}

// SetUnixGroupName adds the unixGroupName to the unix group user delete collection params
func (o *UnixGroupUserDeleteCollectionParams) SetUnixGroupName(unixGroupName string) {
	o.UnixGroupName = unixGroupName
}

// WriteToRequest writes these params to a swagger request
func (o *UnixGroupUserDeleteCollectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
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

	// path param svm.uuid
	if err := r.SetPathParam("svm.uuid", o.SvmUUID); err != nil {
		return err
	}

	// path param unix_group.name
	if err := r.SetPathParam("unix_group.name", o.UnixGroupName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}