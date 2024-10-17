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

// NewUnixGroupDeleteCollectionParams creates a new UnixGroupDeleteCollectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUnixGroupDeleteCollectionParams() *UnixGroupDeleteCollectionParams {
	return &UnixGroupDeleteCollectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUnixGroupDeleteCollectionParamsWithTimeout creates a new UnixGroupDeleteCollectionParams object
// with the ability to set a timeout on a request.
func NewUnixGroupDeleteCollectionParamsWithTimeout(timeout time.Duration) *UnixGroupDeleteCollectionParams {
	return &UnixGroupDeleteCollectionParams{
		timeout: timeout,
	}
}

// NewUnixGroupDeleteCollectionParamsWithContext creates a new UnixGroupDeleteCollectionParams object
// with the ability to set a context for a request.
func NewUnixGroupDeleteCollectionParamsWithContext(ctx context.Context) *UnixGroupDeleteCollectionParams {
	return &UnixGroupDeleteCollectionParams{
		Context: ctx,
	}
}

// NewUnixGroupDeleteCollectionParamsWithHTTPClient creates a new UnixGroupDeleteCollectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewUnixGroupDeleteCollectionParamsWithHTTPClient(client *http.Client) *UnixGroupDeleteCollectionParams {
	return &UnixGroupDeleteCollectionParams{
		HTTPClient: client,
	}
}

/*
UnixGroupDeleteCollectionParams contains all the parameters to send to the API endpoint

	for the unix group delete collection operation.

	Typically these are written to a http.Request.
*/
type UnixGroupDeleteCollectionParams struct {

	/* ContinueOnFailure.

	   Continue even when the operation fails on one of the records.
	*/
	ContinueOnFailure *bool

	/* ID.

	   Filter by id
	*/
	ID *int64

	/* Info.

	   Info specification
	*/
	Info UnixGroupDeleteCollectionBody

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

	/* SvmName.

	   Filter by svm.name
	*/
	SvmName *string

	/* SvmUUID.

	   Filter by svm.uuid
	*/
	SvmUUID *string

	/* UsersName.

	   Filter by users.name
	*/
	UsersName *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the unix group delete collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UnixGroupDeleteCollectionParams) WithDefaults() *UnixGroupDeleteCollectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the unix group delete collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UnixGroupDeleteCollectionParams) SetDefaults() {
	var (
		continueOnFailureDefault = bool(false)

		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)

		serialRecordsDefault = bool(false)
	)

	val := UnixGroupDeleteCollectionParams{
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

// WithTimeout adds the timeout to the unix group delete collection params
func (o *UnixGroupDeleteCollectionParams) WithTimeout(timeout time.Duration) *UnixGroupDeleteCollectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the unix group delete collection params
func (o *UnixGroupDeleteCollectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the unix group delete collection params
func (o *UnixGroupDeleteCollectionParams) WithContext(ctx context.Context) *UnixGroupDeleteCollectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the unix group delete collection params
func (o *UnixGroupDeleteCollectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the unix group delete collection params
func (o *UnixGroupDeleteCollectionParams) WithHTTPClient(client *http.Client) *UnixGroupDeleteCollectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the unix group delete collection params
func (o *UnixGroupDeleteCollectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContinueOnFailure adds the continueOnFailure to the unix group delete collection params
func (o *UnixGroupDeleteCollectionParams) WithContinueOnFailure(continueOnFailure *bool) *UnixGroupDeleteCollectionParams {
	o.SetContinueOnFailure(continueOnFailure)
	return o
}

// SetContinueOnFailure adds the continueOnFailure to the unix group delete collection params
func (o *UnixGroupDeleteCollectionParams) SetContinueOnFailure(continueOnFailure *bool) {
	o.ContinueOnFailure = continueOnFailure
}

// WithID adds the id to the unix group delete collection params
func (o *UnixGroupDeleteCollectionParams) WithID(id *int64) *UnixGroupDeleteCollectionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the unix group delete collection params
func (o *UnixGroupDeleteCollectionParams) SetID(id *int64) {
	o.ID = id
}

// WithInfo adds the info to the unix group delete collection params
func (o *UnixGroupDeleteCollectionParams) WithInfo(info UnixGroupDeleteCollectionBody) *UnixGroupDeleteCollectionParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the unix group delete collection params
func (o *UnixGroupDeleteCollectionParams) SetInfo(info UnixGroupDeleteCollectionBody) {
	o.Info = info
}

// WithName adds the name to the unix group delete collection params
func (o *UnixGroupDeleteCollectionParams) WithName(name *string) *UnixGroupDeleteCollectionParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the unix group delete collection params
func (o *UnixGroupDeleteCollectionParams) SetName(name *string) {
	o.Name = name
}

// WithReturnRecords adds the returnRecords to the unix group delete collection params
func (o *UnixGroupDeleteCollectionParams) WithReturnRecords(returnRecords *bool) *UnixGroupDeleteCollectionParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the unix group delete collection params
func (o *UnixGroupDeleteCollectionParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the unix group delete collection params
func (o *UnixGroupDeleteCollectionParams) WithReturnTimeout(returnTimeout *int64) *UnixGroupDeleteCollectionParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the unix group delete collection params
func (o *UnixGroupDeleteCollectionParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSerialRecords adds the serialRecords to the unix group delete collection params
func (o *UnixGroupDeleteCollectionParams) WithSerialRecords(serialRecords *bool) *UnixGroupDeleteCollectionParams {
	o.SetSerialRecords(serialRecords)
	return o
}

// SetSerialRecords adds the serialRecords to the unix group delete collection params
func (o *UnixGroupDeleteCollectionParams) SetSerialRecords(serialRecords *bool) {
	o.SerialRecords = serialRecords
}

// WithSvmName adds the svmName to the unix group delete collection params
func (o *UnixGroupDeleteCollectionParams) WithSvmName(svmName *string) *UnixGroupDeleteCollectionParams {
	o.SetSvmName(svmName)
	return o
}

// SetSvmName adds the svmName to the unix group delete collection params
func (o *UnixGroupDeleteCollectionParams) SetSvmName(svmName *string) {
	o.SvmName = svmName
}

// WithSvmUUID adds the svmUUID to the unix group delete collection params
func (o *UnixGroupDeleteCollectionParams) WithSvmUUID(svmUUID *string) *UnixGroupDeleteCollectionParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the unix group delete collection params
func (o *UnixGroupDeleteCollectionParams) SetSvmUUID(svmUUID *string) {
	o.SvmUUID = svmUUID
}

// WithUsersName adds the usersName to the unix group delete collection params
func (o *UnixGroupDeleteCollectionParams) WithUsersName(usersName *string) *UnixGroupDeleteCollectionParams {
	o.SetUsersName(usersName)
	return o
}

// SetUsersName adds the usersName to the unix group delete collection params
func (o *UnixGroupDeleteCollectionParams) SetUsersName(usersName *string) {
	o.UsersName = usersName
}

// WriteToRequest writes these params to a swagger request
func (o *UnixGroupDeleteCollectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.ID != nil {

		// query param id
		var qrID int64

		if o.ID != nil {
			qrID = *o.ID
		}
		qID := swag.FormatInt64(qrID)
		if qID != "" {

			if err := r.SetQueryParam("id", qID); err != nil {
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

	if o.UsersName != nil {

		// query param users.name
		var qrUsersName string

		if o.UsersName != nil {
			qrUsersName = *o.UsersName
		}
		qUsersName := qrUsersName
		if qUsersName != "" {

			if err := r.SetQueryParam("users.name", qUsersName); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}