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

// NewVvolBindingDeleteCollectionParams creates a new VvolBindingDeleteCollectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVvolBindingDeleteCollectionParams() *VvolBindingDeleteCollectionParams {
	return &VvolBindingDeleteCollectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVvolBindingDeleteCollectionParamsWithTimeout creates a new VvolBindingDeleteCollectionParams object
// with the ability to set a timeout on a request.
func NewVvolBindingDeleteCollectionParamsWithTimeout(timeout time.Duration) *VvolBindingDeleteCollectionParams {
	return &VvolBindingDeleteCollectionParams{
		timeout: timeout,
	}
}

// NewVvolBindingDeleteCollectionParamsWithContext creates a new VvolBindingDeleteCollectionParams object
// with the ability to set a context for a request.
func NewVvolBindingDeleteCollectionParamsWithContext(ctx context.Context) *VvolBindingDeleteCollectionParams {
	return &VvolBindingDeleteCollectionParams{
		Context: ctx,
	}
}

// NewVvolBindingDeleteCollectionParamsWithHTTPClient creates a new VvolBindingDeleteCollectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewVvolBindingDeleteCollectionParamsWithHTTPClient(client *http.Client) *VvolBindingDeleteCollectionParams {
	return &VvolBindingDeleteCollectionParams{
		HTTPClient: client,
	}
}

/*
VvolBindingDeleteCollectionParams contains all the parameters to send to the API endpoint

	for the vvol binding delete collection operation.

	Typically these are written to a http.Request.
*/
type VvolBindingDeleteCollectionParams struct {

	/* ContinueOnFailure.

	   Continue even when the operation fails on one of the records.
	*/
	ContinueOnFailure *bool

	/* Count.

	   Filter by count
	*/
	Count *int64

	/* DeleteAllReferences.

	   Forces deletion of the binding regardless of the reference count value.

	*/
	DeleteAllReferences *bool

	/* ID.

	   Filter by id
	*/
	ID *int64

	/* Info.

	   Info specification
	*/
	Info VvolBindingDeleteCollectionBody

	/* IsOptimal.

	   Filter by is_optimal
	*/
	IsOptimal *bool

	/* ProtocolEndpointName.

	   Filter by protocol_endpoint.name
	*/
	ProtocolEndpointName *string

	/* ProtocolEndpointUUID.

	   Filter by protocol_endpoint.uuid
	*/
	ProtocolEndpointUUID *string

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

	/* SecondaryID.

	   Filter by secondary_id
	*/
	SecondaryID *string

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

	/* VvolName.

	   Filter by vvol.name
	*/
	VvolName *string

	/* VvolUUID.

	   Filter by vvol.uuid
	*/
	VvolUUID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the vvol binding delete collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VvolBindingDeleteCollectionParams) WithDefaults() *VvolBindingDeleteCollectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the vvol binding delete collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VvolBindingDeleteCollectionParams) SetDefaults() {
	var (
		continueOnFailureDefault = bool(false)

		deleteAllReferencesDefault = bool(false)

		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)

		serialRecordsDefault = bool(false)
	)

	val := VvolBindingDeleteCollectionParams{
		ContinueOnFailure:   &continueOnFailureDefault,
		DeleteAllReferences: &deleteAllReferencesDefault,
		ReturnRecords:       &returnRecordsDefault,
		ReturnTimeout:       &returnTimeoutDefault,
		SerialRecords:       &serialRecordsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the vvol binding delete collection params
func (o *VvolBindingDeleteCollectionParams) WithTimeout(timeout time.Duration) *VvolBindingDeleteCollectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the vvol binding delete collection params
func (o *VvolBindingDeleteCollectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the vvol binding delete collection params
func (o *VvolBindingDeleteCollectionParams) WithContext(ctx context.Context) *VvolBindingDeleteCollectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the vvol binding delete collection params
func (o *VvolBindingDeleteCollectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the vvol binding delete collection params
func (o *VvolBindingDeleteCollectionParams) WithHTTPClient(client *http.Client) *VvolBindingDeleteCollectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the vvol binding delete collection params
func (o *VvolBindingDeleteCollectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContinueOnFailure adds the continueOnFailure to the vvol binding delete collection params
func (o *VvolBindingDeleteCollectionParams) WithContinueOnFailure(continueOnFailure *bool) *VvolBindingDeleteCollectionParams {
	o.SetContinueOnFailure(continueOnFailure)
	return o
}

// SetContinueOnFailure adds the continueOnFailure to the vvol binding delete collection params
func (o *VvolBindingDeleteCollectionParams) SetContinueOnFailure(continueOnFailure *bool) {
	o.ContinueOnFailure = continueOnFailure
}

// WithCount adds the count to the vvol binding delete collection params
func (o *VvolBindingDeleteCollectionParams) WithCount(count *int64) *VvolBindingDeleteCollectionParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the vvol binding delete collection params
func (o *VvolBindingDeleteCollectionParams) SetCount(count *int64) {
	o.Count = count
}

// WithDeleteAllReferences adds the deleteAllReferences to the vvol binding delete collection params
func (o *VvolBindingDeleteCollectionParams) WithDeleteAllReferences(deleteAllReferences *bool) *VvolBindingDeleteCollectionParams {
	o.SetDeleteAllReferences(deleteAllReferences)
	return o
}

// SetDeleteAllReferences adds the deleteAllReferences to the vvol binding delete collection params
func (o *VvolBindingDeleteCollectionParams) SetDeleteAllReferences(deleteAllReferences *bool) {
	o.DeleteAllReferences = deleteAllReferences
}

// WithID adds the id to the vvol binding delete collection params
func (o *VvolBindingDeleteCollectionParams) WithID(id *int64) *VvolBindingDeleteCollectionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the vvol binding delete collection params
func (o *VvolBindingDeleteCollectionParams) SetID(id *int64) {
	o.ID = id
}

// WithInfo adds the info to the vvol binding delete collection params
func (o *VvolBindingDeleteCollectionParams) WithInfo(info VvolBindingDeleteCollectionBody) *VvolBindingDeleteCollectionParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the vvol binding delete collection params
func (o *VvolBindingDeleteCollectionParams) SetInfo(info VvolBindingDeleteCollectionBody) {
	o.Info = info
}

// WithIsOptimal adds the isOptimal to the vvol binding delete collection params
func (o *VvolBindingDeleteCollectionParams) WithIsOptimal(isOptimal *bool) *VvolBindingDeleteCollectionParams {
	o.SetIsOptimal(isOptimal)
	return o
}

// SetIsOptimal adds the isOptimal to the vvol binding delete collection params
func (o *VvolBindingDeleteCollectionParams) SetIsOptimal(isOptimal *bool) {
	o.IsOptimal = isOptimal
}

// WithProtocolEndpointName adds the protocolEndpointName to the vvol binding delete collection params
func (o *VvolBindingDeleteCollectionParams) WithProtocolEndpointName(protocolEndpointName *string) *VvolBindingDeleteCollectionParams {
	o.SetProtocolEndpointName(protocolEndpointName)
	return o
}

// SetProtocolEndpointName adds the protocolEndpointName to the vvol binding delete collection params
func (o *VvolBindingDeleteCollectionParams) SetProtocolEndpointName(protocolEndpointName *string) {
	o.ProtocolEndpointName = protocolEndpointName
}

// WithProtocolEndpointUUID adds the protocolEndpointUUID to the vvol binding delete collection params
func (o *VvolBindingDeleteCollectionParams) WithProtocolEndpointUUID(protocolEndpointUUID *string) *VvolBindingDeleteCollectionParams {
	o.SetProtocolEndpointUUID(protocolEndpointUUID)
	return o
}

// SetProtocolEndpointUUID adds the protocolEndpointUuid to the vvol binding delete collection params
func (o *VvolBindingDeleteCollectionParams) SetProtocolEndpointUUID(protocolEndpointUUID *string) {
	o.ProtocolEndpointUUID = protocolEndpointUUID
}

// WithReturnRecords adds the returnRecords to the vvol binding delete collection params
func (o *VvolBindingDeleteCollectionParams) WithReturnRecords(returnRecords *bool) *VvolBindingDeleteCollectionParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the vvol binding delete collection params
func (o *VvolBindingDeleteCollectionParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the vvol binding delete collection params
func (o *VvolBindingDeleteCollectionParams) WithReturnTimeout(returnTimeout *int64) *VvolBindingDeleteCollectionParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the vvol binding delete collection params
func (o *VvolBindingDeleteCollectionParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSecondaryID adds the secondaryID to the vvol binding delete collection params
func (o *VvolBindingDeleteCollectionParams) WithSecondaryID(secondaryID *string) *VvolBindingDeleteCollectionParams {
	o.SetSecondaryID(secondaryID)
	return o
}

// SetSecondaryID adds the secondaryId to the vvol binding delete collection params
func (o *VvolBindingDeleteCollectionParams) SetSecondaryID(secondaryID *string) {
	o.SecondaryID = secondaryID
}

// WithSerialRecords adds the serialRecords to the vvol binding delete collection params
func (o *VvolBindingDeleteCollectionParams) WithSerialRecords(serialRecords *bool) *VvolBindingDeleteCollectionParams {
	o.SetSerialRecords(serialRecords)
	return o
}

// SetSerialRecords adds the serialRecords to the vvol binding delete collection params
func (o *VvolBindingDeleteCollectionParams) SetSerialRecords(serialRecords *bool) {
	o.SerialRecords = serialRecords
}

// WithSvmName adds the svmName to the vvol binding delete collection params
func (o *VvolBindingDeleteCollectionParams) WithSvmName(svmName *string) *VvolBindingDeleteCollectionParams {
	o.SetSvmName(svmName)
	return o
}

// SetSvmName adds the svmName to the vvol binding delete collection params
func (o *VvolBindingDeleteCollectionParams) SetSvmName(svmName *string) {
	o.SvmName = svmName
}

// WithSvmUUID adds the svmUUID to the vvol binding delete collection params
func (o *VvolBindingDeleteCollectionParams) WithSvmUUID(svmUUID *string) *VvolBindingDeleteCollectionParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the vvol binding delete collection params
func (o *VvolBindingDeleteCollectionParams) SetSvmUUID(svmUUID *string) {
	o.SvmUUID = svmUUID
}

// WithVvolName adds the vvolName to the vvol binding delete collection params
func (o *VvolBindingDeleteCollectionParams) WithVvolName(vvolName *string) *VvolBindingDeleteCollectionParams {
	o.SetVvolName(vvolName)
	return o
}

// SetVvolName adds the vvolName to the vvol binding delete collection params
func (o *VvolBindingDeleteCollectionParams) SetVvolName(vvolName *string) {
	o.VvolName = vvolName
}

// WithVvolUUID adds the vvolUUID to the vvol binding delete collection params
func (o *VvolBindingDeleteCollectionParams) WithVvolUUID(vvolUUID *string) *VvolBindingDeleteCollectionParams {
	o.SetVvolUUID(vvolUUID)
	return o
}

// SetVvolUUID adds the vvolUuid to the vvol binding delete collection params
func (o *VvolBindingDeleteCollectionParams) SetVvolUUID(vvolUUID *string) {
	o.VvolUUID = vvolUUID
}

// WriteToRequest writes these params to a swagger request
func (o *VvolBindingDeleteCollectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Count != nil {

		// query param count
		var qrCount int64

		if o.Count != nil {
			qrCount = *o.Count
		}
		qCount := swag.FormatInt64(qrCount)
		if qCount != "" {

			if err := r.SetQueryParam("count", qCount); err != nil {
				return err
			}
		}
	}

	if o.DeleteAllReferences != nil {

		// query param delete_all_references
		var qrDeleteAllReferences bool

		if o.DeleteAllReferences != nil {
			qrDeleteAllReferences = *o.DeleteAllReferences
		}
		qDeleteAllReferences := swag.FormatBool(qrDeleteAllReferences)
		if qDeleteAllReferences != "" {

			if err := r.SetQueryParam("delete_all_references", qDeleteAllReferences); err != nil {
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

	if o.IsOptimal != nil {

		// query param is_optimal
		var qrIsOptimal bool

		if o.IsOptimal != nil {
			qrIsOptimal = *o.IsOptimal
		}
		qIsOptimal := swag.FormatBool(qrIsOptimal)
		if qIsOptimal != "" {

			if err := r.SetQueryParam("is_optimal", qIsOptimal); err != nil {
				return err
			}
		}
	}

	if o.ProtocolEndpointName != nil {

		// query param protocol_endpoint.name
		var qrProtocolEndpointName string

		if o.ProtocolEndpointName != nil {
			qrProtocolEndpointName = *o.ProtocolEndpointName
		}
		qProtocolEndpointName := qrProtocolEndpointName
		if qProtocolEndpointName != "" {

			if err := r.SetQueryParam("protocol_endpoint.name", qProtocolEndpointName); err != nil {
				return err
			}
		}
	}

	if o.ProtocolEndpointUUID != nil {

		// query param protocol_endpoint.uuid
		var qrProtocolEndpointUUID string

		if o.ProtocolEndpointUUID != nil {
			qrProtocolEndpointUUID = *o.ProtocolEndpointUUID
		}
		qProtocolEndpointUUID := qrProtocolEndpointUUID
		if qProtocolEndpointUUID != "" {

			if err := r.SetQueryParam("protocol_endpoint.uuid", qProtocolEndpointUUID); err != nil {
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

	if o.SecondaryID != nil {

		// query param secondary_id
		var qrSecondaryID string

		if o.SecondaryID != nil {
			qrSecondaryID = *o.SecondaryID
		}
		qSecondaryID := qrSecondaryID
		if qSecondaryID != "" {

			if err := r.SetQueryParam("secondary_id", qSecondaryID); err != nil {
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

	if o.VvolName != nil {

		// query param vvol.name
		var qrVvolName string

		if o.VvolName != nil {
			qrVvolName = *o.VvolName
		}
		qVvolName := qrVvolName
		if qVvolName != "" {

			if err := r.SetQueryParam("vvol.name", qVvolName); err != nil {
				return err
			}
		}
	}

	if o.VvolUUID != nil {

		// query param vvol.uuid
		var qrVvolUUID string

		if o.VvolUUID != nil {
			qrVvolUUID = *o.VvolUUID
		}
		qVvolUUID := qrVvolUUID
		if qVvolUUID != "" {

			if err := r.SetQueryParam("vvol.uuid", qVvolUUID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}