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

// NewNameMappingModifyCollectionParams creates a new NameMappingModifyCollectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNameMappingModifyCollectionParams() *NameMappingModifyCollectionParams {
	return &NameMappingModifyCollectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNameMappingModifyCollectionParamsWithTimeout creates a new NameMappingModifyCollectionParams object
// with the ability to set a timeout on a request.
func NewNameMappingModifyCollectionParamsWithTimeout(timeout time.Duration) *NameMappingModifyCollectionParams {
	return &NameMappingModifyCollectionParams{
		timeout: timeout,
	}
}

// NewNameMappingModifyCollectionParamsWithContext creates a new NameMappingModifyCollectionParams object
// with the ability to set a context for a request.
func NewNameMappingModifyCollectionParamsWithContext(ctx context.Context) *NameMappingModifyCollectionParams {
	return &NameMappingModifyCollectionParams{
		Context: ctx,
	}
}

// NewNameMappingModifyCollectionParamsWithHTTPClient creates a new NameMappingModifyCollectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewNameMappingModifyCollectionParamsWithHTTPClient(client *http.Client) *NameMappingModifyCollectionParams {
	return &NameMappingModifyCollectionParams{
		HTTPClient: client,
	}
}

/*
NameMappingModifyCollectionParams contains all the parameters to send to the API endpoint

	for the name mapping modify collection operation.

	Typically these are written to a http.Request.
*/
type NameMappingModifyCollectionParams struct {

	/* ClientMatch.

	   Filter by client_match
	*/
	ClientMatch *string

	/* ContinueOnFailure.

	   Continue even when the operation fails on one of the records.
	*/
	ContinueOnFailure *bool

	/* Direction.

	   Filter by direction
	*/
	Direction *string

	/* Index.

	   Filter by index
	*/
	Index *int64

	/* Info.

	   Info specification
	*/
	Info NameMappingModifyCollectionBody

	/* NewIndex.

	   New position of the Index after a swap is completed.
	*/
	NewIndex *int64

	/* Pattern.

	   Filter by pattern
	*/
	Pattern *string

	/* Replacement.

	   Filter by replacement
	*/
	Replacement *string

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

// WithDefaults hydrates default values in the name mapping modify collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NameMappingModifyCollectionParams) WithDefaults() *NameMappingModifyCollectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the name mapping modify collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NameMappingModifyCollectionParams) SetDefaults() {
	var (
		continueOnFailureDefault = bool(false)

		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)

		serialRecordsDefault = bool(false)
	)

	val := NameMappingModifyCollectionParams{
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

// WithTimeout adds the timeout to the name mapping modify collection params
func (o *NameMappingModifyCollectionParams) WithTimeout(timeout time.Duration) *NameMappingModifyCollectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the name mapping modify collection params
func (o *NameMappingModifyCollectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the name mapping modify collection params
func (o *NameMappingModifyCollectionParams) WithContext(ctx context.Context) *NameMappingModifyCollectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the name mapping modify collection params
func (o *NameMappingModifyCollectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the name mapping modify collection params
func (o *NameMappingModifyCollectionParams) WithHTTPClient(client *http.Client) *NameMappingModifyCollectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the name mapping modify collection params
func (o *NameMappingModifyCollectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClientMatch adds the clientMatch to the name mapping modify collection params
func (o *NameMappingModifyCollectionParams) WithClientMatch(clientMatch *string) *NameMappingModifyCollectionParams {
	o.SetClientMatch(clientMatch)
	return o
}

// SetClientMatch adds the clientMatch to the name mapping modify collection params
func (o *NameMappingModifyCollectionParams) SetClientMatch(clientMatch *string) {
	o.ClientMatch = clientMatch
}

// WithContinueOnFailure adds the continueOnFailure to the name mapping modify collection params
func (o *NameMappingModifyCollectionParams) WithContinueOnFailure(continueOnFailure *bool) *NameMappingModifyCollectionParams {
	o.SetContinueOnFailure(continueOnFailure)
	return o
}

// SetContinueOnFailure adds the continueOnFailure to the name mapping modify collection params
func (o *NameMappingModifyCollectionParams) SetContinueOnFailure(continueOnFailure *bool) {
	o.ContinueOnFailure = continueOnFailure
}

// WithDirection adds the direction to the name mapping modify collection params
func (o *NameMappingModifyCollectionParams) WithDirection(direction *string) *NameMappingModifyCollectionParams {
	o.SetDirection(direction)
	return o
}

// SetDirection adds the direction to the name mapping modify collection params
func (o *NameMappingModifyCollectionParams) SetDirection(direction *string) {
	o.Direction = direction
}

// WithIndex adds the index to the name mapping modify collection params
func (o *NameMappingModifyCollectionParams) WithIndex(index *int64) *NameMappingModifyCollectionParams {
	o.SetIndex(index)
	return o
}

// SetIndex adds the index to the name mapping modify collection params
func (o *NameMappingModifyCollectionParams) SetIndex(index *int64) {
	o.Index = index
}

// WithInfo adds the info to the name mapping modify collection params
func (o *NameMappingModifyCollectionParams) WithInfo(info NameMappingModifyCollectionBody) *NameMappingModifyCollectionParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the name mapping modify collection params
func (o *NameMappingModifyCollectionParams) SetInfo(info NameMappingModifyCollectionBody) {
	o.Info = info
}

// WithNewIndex adds the newIndex to the name mapping modify collection params
func (o *NameMappingModifyCollectionParams) WithNewIndex(newIndex *int64) *NameMappingModifyCollectionParams {
	o.SetNewIndex(newIndex)
	return o
}

// SetNewIndex adds the newIndex to the name mapping modify collection params
func (o *NameMappingModifyCollectionParams) SetNewIndex(newIndex *int64) {
	o.NewIndex = newIndex
}

// WithPattern adds the pattern to the name mapping modify collection params
func (o *NameMappingModifyCollectionParams) WithPattern(pattern *string) *NameMappingModifyCollectionParams {
	o.SetPattern(pattern)
	return o
}

// SetPattern adds the pattern to the name mapping modify collection params
func (o *NameMappingModifyCollectionParams) SetPattern(pattern *string) {
	o.Pattern = pattern
}

// WithReplacement adds the replacement to the name mapping modify collection params
func (o *NameMappingModifyCollectionParams) WithReplacement(replacement *string) *NameMappingModifyCollectionParams {
	o.SetReplacement(replacement)
	return o
}

// SetReplacement adds the replacement to the name mapping modify collection params
func (o *NameMappingModifyCollectionParams) SetReplacement(replacement *string) {
	o.Replacement = replacement
}

// WithReturnRecords adds the returnRecords to the name mapping modify collection params
func (o *NameMappingModifyCollectionParams) WithReturnRecords(returnRecords *bool) *NameMappingModifyCollectionParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the name mapping modify collection params
func (o *NameMappingModifyCollectionParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the name mapping modify collection params
func (o *NameMappingModifyCollectionParams) WithReturnTimeout(returnTimeout *int64) *NameMappingModifyCollectionParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the name mapping modify collection params
func (o *NameMappingModifyCollectionParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSerialRecords adds the serialRecords to the name mapping modify collection params
func (o *NameMappingModifyCollectionParams) WithSerialRecords(serialRecords *bool) *NameMappingModifyCollectionParams {
	o.SetSerialRecords(serialRecords)
	return o
}

// SetSerialRecords adds the serialRecords to the name mapping modify collection params
func (o *NameMappingModifyCollectionParams) SetSerialRecords(serialRecords *bool) {
	o.SerialRecords = serialRecords
}

// WithSvmName adds the svmName to the name mapping modify collection params
func (o *NameMappingModifyCollectionParams) WithSvmName(svmName *string) *NameMappingModifyCollectionParams {
	o.SetSvmName(svmName)
	return o
}

// SetSvmName adds the svmName to the name mapping modify collection params
func (o *NameMappingModifyCollectionParams) SetSvmName(svmName *string) {
	o.SvmName = svmName
}

// WithSvmUUID adds the svmUUID to the name mapping modify collection params
func (o *NameMappingModifyCollectionParams) WithSvmUUID(svmUUID *string) *NameMappingModifyCollectionParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the name mapping modify collection params
func (o *NameMappingModifyCollectionParams) SetSvmUUID(svmUUID *string) {
	o.SvmUUID = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *NameMappingModifyCollectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ClientMatch != nil {

		// query param client_match
		var qrClientMatch string

		if o.ClientMatch != nil {
			qrClientMatch = *o.ClientMatch
		}
		qClientMatch := qrClientMatch
		if qClientMatch != "" {

			if err := r.SetQueryParam("client_match", qClientMatch); err != nil {
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

	if o.Direction != nil {

		// query param direction
		var qrDirection string

		if o.Direction != nil {
			qrDirection = *o.Direction
		}
		qDirection := qrDirection
		if qDirection != "" {

			if err := r.SetQueryParam("direction", qDirection); err != nil {
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

	if o.Pattern != nil {

		// query param pattern
		var qrPattern string

		if o.Pattern != nil {
			qrPattern = *o.Pattern
		}
		qPattern := qrPattern
		if qPattern != "" {

			if err := r.SetQueryParam("pattern", qPattern); err != nil {
				return err
			}
		}
	}

	if o.Replacement != nil {

		// query param replacement
		var qrReplacement string

		if o.Replacement != nil {
			qrReplacement = *o.Replacement
		}
		qReplacement := qrReplacement
		if qReplacement != "" {

			if err := r.SetQueryParam("replacement", qReplacement); err != nil {
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