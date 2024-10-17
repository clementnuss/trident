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

// NewLocalCifsGroupDeleteCollectionParams creates a new LocalCifsGroupDeleteCollectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLocalCifsGroupDeleteCollectionParams() *LocalCifsGroupDeleteCollectionParams {
	return &LocalCifsGroupDeleteCollectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLocalCifsGroupDeleteCollectionParamsWithTimeout creates a new LocalCifsGroupDeleteCollectionParams object
// with the ability to set a timeout on a request.
func NewLocalCifsGroupDeleteCollectionParamsWithTimeout(timeout time.Duration) *LocalCifsGroupDeleteCollectionParams {
	return &LocalCifsGroupDeleteCollectionParams{
		timeout: timeout,
	}
}

// NewLocalCifsGroupDeleteCollectionParamsWithContext creates a new LocalCifsGroupDeleteCollectionParams object
// with the ability to set a context for a request.
func NewLocalCifsGroupDeleteCollectionParamsWithContext(ctx context.Context) *LocalCifsGroupDeleteCollectionParams {
	return &LocalCifsGroupDeleteCollectionParams{
		Context: ctx,
	}
}

// NewLocalCifsGroupDeleteCollectionParamsWithHTTPClient creates a new LocalCifsGroupDeleteCollectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewLocalCifsGroupDeleteCollectionParamsWithHTTPClient(client *http.Client) *LocalCifsGroupDeleteCollectionParams {
	return &LocalCifsGroupDeleteCollectionParams{
		HTTPClient: client,
	}
}

/*
LocalCifsGroupDeleteCollectionParams contains all the parameters to send to the API endpoint

	for the local cifs group delete collection operation.

	Typically these are written to a http.Request.
*/
type LocalCifsGroupDeleteCollectionParams struct {

	/* ContinueOnFailure.

	   Continue even when the operation fails on one of the records.
	*/
	ContinueOnFailure *bool

	/* Description.

	   Filter by description
	*/
	Description *string

	/* Info.

	   Info specification
	*/
	Info LocalCifsGroupDeleteCollectionBody

	/* MembersName.

	   Filter by members.name
	*/
	MembersName *string

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

	/* Sid.

	   Filter by sid
	*/
	Sid *string

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

// WithDefaults hydrates default values in the local cifs group delete collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LocalCifsGroupDeleteCollectionParams) WithDefaults() *LocalCifsGroupDeleteCollectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the local cifs group delete collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LocalCifsGroupDeleteCollectionParams) SetDefaults() {
	var (
		continueOnFailureDefault = bool(false)

		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)

		serialRecordsDefault = bool(false)
	)

	val := LocalCifsGroupDeleteCollectionParams{
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

// WithTimeout adds the timeout to the local cifs group delete collection params
func (o *LocalCifsGroupDeleteCollectionParams) WithTimeout(timeout time.Duration) *LocalCifsGroupDeleteCollectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the local cifs group delete collection params
func (o *LocalCifsGroupDeleteCollectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the local cifs group delete collection params
func (o *LocalCifsGroupDeleteCollectionParams) WithContext(ctx context.Context) *LocalCifsGroupDeleteCollectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the local cifs group delete collection params
func (o *LocalCifsGroupDeleteCollectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the local cifs group delete collection params
func (o *LocalCifsGroupDeleteCollectionParams) WithHTTPClient(client *http.Client) *LocalCifsGroupDeleteCollectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the local cifs group delete collection params
func (o *LocalCifsGroupDeleteCollectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContinueOnFailure adds the continueOnFailure to the local cifs group delete collection params
func (o *LocalCifsGroupDeleteCollectionParams) WithContinueOnFailure(continueOnFailure *bool) *LocalCifsGroupDeleteCollectionParams {
	o.SetContinueOnFailure(continueOnFailure)
	return o
}

// SetContinueOnFailure adds the continueOnFailure to the local cifs group delete collection params
func (o *LocalCifsGroupDeleteCollectionParams) SetContinueOnFailure(continueOnFailure *bool) {
	o.ContinueOnFailure = continueOnFailure
}

// WithDescription adds the description to the local cifs group delete collection params
func (o *LocalCifsGroupDeleteCollectionParams) WithDescription(description *string) *LocalCifsGroupDeleteCollectionParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the local cifs group delete collection params
func (o *LocalCifsGroupDeleteCollectionParams) SetDescription(description *string) {
	o.Description = description
}

// WithInfo adds the info to the local cifs group delete collection params
func (o *LocalCifsGroupDeleteCollectionParams) WithInfo(info LocalCifsGroupDeleteCollectionBody) *LocalCifsGroupDeleteCollectionParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the local cifs group delete collection params
func (o *LocalCifsGroupDeleteCollectionParams) SetInfo(info LocalCifsGroupDeleteCollectionBody) {
	o.Info = info
}

// WithMembersName adds the membersName to the local cifs group delete collection params
func (o *LocalCifsGroupDeleteCollectionParams) WithMembersName(membersName *string) *LocalCifsGroupDeleteCollectionParams {
	o.SetMembersName(membersName)
	return o
}

// SetMembersName adds the membersName to the local cifs group delete collection params
func (o *LocalCifsGroupDeleteCollectionParams) SetMembersName(membersName *string) {
	o.MembersName = membersName
}

// WithName adds the name to the local cifs group delete collection params
func (o *LocalCifsGroupDeleteCollectionParams) WithName(name *string) *LocalCifsGroupDeleteCollectionParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the local cifs group delete collection params
func (o *LocalCifsGroupDeleteCollectionParams) SetName(name *string) {
	o.Name = name
}

// WithReturnRecords adds the returnRecords to the local cifs group delete collection params
func (o *LocalCifsGroupDeleteCollectionParams) WithReturnRecords(returnRecords *bool) *LocalCifsGroupDeleteCollectionParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the local cifs group delete collection params
func (o *LocalCifsGroupDeleteCollectionParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the local cifs group delete collection params
func (o *LocalCifsGroupDeleteCollectionParams) WithReturnTimeout(returnTimeout *int64) *LocalCifsGroupDeleteCollectionParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the local cifs group delete collection params
func (o *LocalCifsGroupDeleteCollectionParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSerialRecords adds the serialRecords to the local cifs group delete collection params
func (o *LocalCifsGroupDeleteCollectionParams) WithSerialRecords(serialRecords *bool) *LocalCifsGroupDeleteCollectionParams {
	o.SetSerialRecords(serialRecords)
	return o
}

// SetSerialRecords adds the serialRecords to the local cifs group delete collection params
func (o *LocalCifsGroupDeleteCollectionParams) SetSerialRecords(serialRecords *bool) {
	o.SerialRecords = serialRecords
}

// WithSid adds the sid to the local cifs group delete collection params
func (o *LocalCifsGroupDeleteCollectionParams) WithSid(sid *string) *LocalCifsGroupDeleteCollectionParams {
	o.SetSid(sid)
	return o
}

// SetSid adds the sid to the local cifs group delete collection params
func (o *LocalCifsGroupDeleteCollectionParams) SetSid(sid *string) {
	o.Sid = sid
}

// WithSvmName adds the svmName to the local cifs group delete collection params
func (o *LocalCifsGroupDeleteCollectionParams) WithSvmName(svmName *string) *LocalCifsGroupDeleteCollectionParams {
	o.SetSvmName(svmName)
	return o
}

// SetSvmName adds the svmName to the local cifs group delete collection params
func (o *LocalCifsGroupDeleteCollectionParams) SetSvmName(svmName *string) {
	o.SvmName = svmName
}

// WithSvmUUID adds the svmUUID to the local cifs group delete collection params
func (o *LocalCifsGroupDeleteCollectionParams) WithSvmUUID(svmUUID *string) *LocalCifsGroupDeleteCollectionParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the local cifs group delete collection params
func (o *LocalCifsGroupDeleteCollectionParams) SetSvmUUID(svmUUID *string) {
	o.SvmUUID = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *LocalCifsGroupDeleteCollectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Description != nil {

		// query param description
		var qrDescription string

		if o.Description != nil {
			qrDescription = *o.Description
		}
		qDescription := qrDescription
		if qDescription != "" {

			if err := r.SetQueryParam("description", qDescription); err != nil {
				return err
			}
		}
	}
	if err := r.SetBodyParam(o.Info); err != nil {
		return err
	}

	if o.MembersName != nil {

		// query param members.name
		var qrMembersName string

		if o.MembersName != nil {
			qrMembersName = *o.MembersName
		}
		qMembersName := qrMembersName
		if qMembersName != "" {

			if err := r.SetQueryParam("members.name", qMembersName); err != nil {
				return err
			}
		}
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

	if o.Sid != nil {

		// query param sid
		var qrSid string

		if o.Sid != nil {
			qrSid = *o.Sid
		}
		qSid := qrSid
		if qSid != "" {

			if err := r.SetQueryParam("sid", qSid); err != nil {
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