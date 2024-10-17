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

// NewCifsSearchPathModifyCollectionParams creates a new CifsSearchPathModifyCollectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCifsSearchPathModifyCollectionParams() *CifsSearchPathModifyCollectionParams {
	return &CifsSearchPathModifyCollectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCifsSearchPathModifyCollectionParamsWithTimeout creates a new CifsSearchPathModifyCollectionParams object
// with the ability to set a timeout on a request.
func NewCifsSearchPathModifyCollectionParamsWithTimeout(timeout time.Duration) *CifsSearchPathModifyCollectionParams {
	return &CifsSearchPathModifyCollectionParams{
		timeout: timeout,
	}
}

// NewCifsSearchPathModifyCollectionParamsWithContext creates a new CifsSearchPathModifyCollectionParams object
// with the ability to set a context for a request.
func NewCifsSearchPathModifyCollectionParamsWithContext(ctx context.Context) *CifsSearchPathModifyCollectionParams {
	return &CifsSearchPathModifyCollectionParams{
		Context: ctx,
	}
}

// NewCifsSearchPathModifyCollectionParamsWithHTTPClient creates a new CifsSearchPathModifyCollectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewCifsSearchPathModifyCollectionParamsWithHTTPClient(client *http.Client) *CifsSearchPathModifyCollectionParams {
	return &CifsSearchPathModifyCollectionParams{
		HTTPClient: client,
	}
}

/*
CifsSearchPathModifyCollectionParams contains all the parameters to send to the API endpoint

	for the cifs search path modify collection operation.

	Typically these are written to a http.Request.
*/
type CifsSearchPathModifyCollectionParams struct {

	/* ContinueOnFailure.

	   Continue even when the operation fails on one of the records.
	*/
	ContinueOnFailure *bool

	/* Index.

	   Filter by index
	*/
	Index *int64

	/* Info.

	   Info specification
	*/
	Info CifsSearchPathModifyCollectionBody

	/* NewIndex.

	   New position for the home directory search path
	*/
	NewIndex *int64

	/* Path.

	   Filter by path
	*/
	Path *string

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

// WithDefaults hydrates default values in the cifs search path modify collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CifsSearchPathModifyCollectionParams) WithDefaults() *CifsSearchPathModifyCollectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cifs search path modify collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CifsSearchPathModifyCollectionParams) SetDefaults() {
	var (
		continueOnFailureDefault = bool(false)

		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)

		serialRecordsDefault = bool(false)
	)

	val := CifsSearchPathModifyCollectionParams{
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

// WithTimeout adds the timeout to the cifs search path modify collection params
func (o *CifsSearchPathModifyCollectionParams) WithTimeout(timeout time.Duration) *CifsSearchPathModifyCollectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cifs search path modify collection params
func (o *CifsSearchPathModifyCollectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cifs search path modify collection params
func (o *CifsSearchPathModifyCollectionParams) WithContext(ctx context.Context) *CifsSearchPathModifyCollectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cifs search path modify collection params
func (o *CifsSearchPathModifyCollectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cifs search path modify collection params
func (o *CifsSearchPathModifyCollectionParams) WithHTTPClient(client *http.Client) *CifsSearchPathModifyCollectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cifs search path modify collection params
func (o *CifsSearchPathModifyCollectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContinueOnFailure adds the continueOnFailure to the cifs search path modify collection params
func (o *CifsSearchPathModifyCollectionParams) WithContinueOnFailure(continueOnFailure *bool) *CifsSearchPathModifyCollectionParams {
	o.SetContinueOnFailure(continueOnFailure)
	return o
}

// SetContinueOnFailure adds the continueOnFailure to the cifs search path modify collection params
func (o *CifsSearchPathModifyCollectionParams) SetContinueOnFailure(continueOnFailure *bool) {
	o.ContinueOnFailure = continueOnFailure
}

// WithIndex adds the index to the cifs search path modify collection params
func (o *CifsSearchPathModifyCollectionParams) WithIndex(index *int64) *CifsSearchPathModifyCollectionParams {
	o.SetIndex(index)
	return o
}

// SetIndex adds the index to the cifs search path modify collection params
func (o *CifsSearchPathModifyCollectionParams) SetIndex(index *int64) {
	o.Index = index
}

// WithInfo adds the info to the cifs search path modify collection params
func (o *CifsSearchPathModifyCollectionParams) WithInfo(info CifsSearchPathModifyCollectionBody) *CifsSearchPathModifyCollectionParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the cifs search path modify collection params
func (o *CifsSearchPathModifyCollectionParams) SetInfo(info CifsSearchPathModifyCollectionBody) {
	o.Info = info
}

// WithNewIndex adds the newIndex to the cifs search path modify collection params
func (o *CifsSearchPathModifyCollectionParams) WithNewIndex(newIndex *int64) *CifsSearchPathModifyCollectionParams {
	o.SetNewIndex(newIndex)
	return o
}

// SetNewIndex adds the newIndex to the cifs search path modify collection params
func (o *CifsSearchPathModifyCollectionParams) SetNewIndex(newIndex *int64) {
	o.NewIndex = newIndex
}

// WithPath adds the path to the cifs search path modify collection params
func (o *CifsSearchPathModifyCollectionParams) WithPath(path *string) *CifsSearchPathModifyCollectionParams {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the cifs search path modify collection params
func (o *CifsSearchPathModifyCollectionParams) SetPath(path *string) {
	o.Path = path
}

// WithReturnRecords adds the returnRecords to the cifs search path modify collection params
func (o *CifsSearchPathModifyCollectionParams) WithReturnRecords(returnRecords *bool) *CifsSearchPathModifyCollectionParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the cifs search path modify collection params
func (o *CifsSearchPathModifyCollectionParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the cifs search path modify collection params
func (o *CifsSearchPathModifyCollectionParams) WithReturnTimeout(returnTimeout *int64) *CifsSearchPathModifyCollectionParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the cifs search path modify collection params
func (o *CifsSearchPathModifyCollectionParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSerialRecords adds the serialRecords to the cifs search path modify collection params
func (o *CifsSearchPathModifyCollectionParams) WithSerialRecords(serialRecords *bool) *CifsSearchPathModifyCollectionParams {
	o.SetSerialRecords(serialRecords)
	return o
}

// SetSerialRecords adds the serialRecords to the cifs search path modify collection params
func (o *CifsSearchPathModifyCollectionParams) SetSerialRecords(serialRecords *bool) {
	o.SerialRecords = serialRecords
}

// WithSvmName adds the svmName to the cifs search path modify collection params
func (o *CifsSearchPathModifyCollectionParams) WithSvmName(svmName *string) *CifsSearchPathModifyCollectionParams {
	o.SetSvmName(svmName)
	return o
}

// SetSvmName adds the svmName to the cifs search path modify collection params
func (o *CifsSearchPathModifyCollectionParams) SetSvmName(svmName *string) {
	o.SvmName = svmName
}

// WithSvmUUID adds the svmUUID to the cifs search path modify collection params
func (o *CifsSearchPathModifyCollectionParams) WithSvmUUID(svmUUID *string) *CifsSearchPathModifyCollectionParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the cifs search path modify collection params
func (o *CifsSearchPathModifyCollectionParams) SetSvmUUID(svmUUID *string) {
	o.SvmUUID = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *CifsSearchPathModifyCollectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Index != nil {

		// query param index
		var qrIndex int64

		if o.Index != nil {
			qrIndex = *o.Index
		}
		qIndex := swag.FormatInt64(qrIndex)
		if qIndex != "" {

			if err := r.SetQueryParam("index", qIndex); err != nil {
				return err
			}
		}
	}
	if err := r.SetBodyParam(o.Info); err != nil {
		return err
	}

	if o.NewIndex != nil {

		// query param new_index
		var qrNewIndex int64

		if o.NewIndex != nil {
			qrNewIndex = *o.NewIndex
		}
		qNewIndex := swag.FormatInt64(qrNewIndex)
		if qNewIndex != "" {

			if err := r.SetQueryParam("new_index", qNewIndex); err != nil {
				return err
			}
		}
	}

	if o.Path != nil {

		// query param path
		var qrPath string

		if o.Path != nil {
			qrPath = *o.Path
		}
		qPath := qrPath
		if qPath != "" {

			if err := r.SetQueryParam("path", qPath); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}