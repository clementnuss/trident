// Code generated by go-swagger; DO NOT EDIT.

package storage

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

// NewSnapshotPolicyScheduleModifyCollectionParams creates a new SnapshotPolicyScheduleModifyCollectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSnapshotPolicyScheduleModifyCollectionParams() *SnapshotPolicyScheduleModifyCollectionParams {
	return &SnapshotPolicyScheduleModifyCollectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSnapshotPolicyScheduleModifyCollectionParamsWithTimeout creates a new SnapshotPolicyScheduleModifyCollectionParams object
// with the ability to set a timeout on a request.
func NewSnapshotPolicyScheduleModifyCollectionParamsWithTimeout(timeout time.Duration) *SnapshotPolicyScheduleModifyCollectionParams {
	return &SnapshotPolicyScheduleModifyCollectionParams{
		timeout: timeout,
	}
}

// NewSnapshotPolicyScheduleModifyCollectionParamsWithContext creates a new SnapshotPolicyScheduleModifyCollectionParams object
// with the ability to set a context for a request.
func NewSnapshotPolicyScheduleModifyCollectionParamsWithContext(ctx context.Context) *SnapshotPolicyScheduleModifyCollectionParams {
	return &SnapshotPolicyScheduleModifyCollectionParams{
		Context: ctx,
	}
}

// NewSnapshotPolicyScheduleModifyCollectionParamsWithHTTPClient creates a new SnapshotPolicyScheduleModifyCollectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewSnapshotPolicyScheduleModifyCollectionParamsWithHTTPClient(client *http.Client) *SnapshotPolicyScheduleModifyCollectionParams {
	return &SnapshotPolicyScheduleModifyCollectionParams{
		HTTPClient: client,
	}
}

/*
SnapshotPolicyScheduleModifyCollectionParams contains all the parameters to send to the API endpoint

	for the snapshot policy schedule modify collection operation.

	Typically these are written to a http.Request.
*/
type SnapshotPolicyScheduleModifyCollectionParams struct {

	/* ContinueOnFailure.

	   Continue even when the operation fails on one of the records.
	*/
	ContinueOnFailure *bool

	/* Count.

	   Filter by count
	*/
	Count *int64

	/* Info.

	   Info specification
	*/
	Info SnapshotPolicyScheduleModifyCollectionBody

	/* Prefix.

	   Filter by prefix
	*/
	Prefix *string

	/* RetentionPeriod.

	   Filter by retention_period
	*/
	RetentionPeriod *string

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

	/* ScheduleName.

	   Filter by schedule.name
	*/
	ScheduleName *string

	/* ScheduleUUID.

	   Filter by schedule.uuid
	*/
	ScheduleUUID *string

	/* SerialRecords.

	   Perform the operation on the records synchronously.
	*/
	SerialRecords *bool

	/* SnapmirrorLabel.

	   Filter by snapmirror_label
	*/
	SnapmirrorLabel *string

	/* SnapshotPolicyName.

	   Filter by snapshot_policy.name
	*/
	SnapshotPolicyName *string

	/* SnapshotPolicyUUID.

	   Snapshot policy UUID
	*/
	SnapshotPolicyUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the snapshot policy schedule modify collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnapshotPolicyScheduleModifyCollectionParams) WithDefaults() *SnapshotPolicyScheduleModifyCollectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the snapshot policy schedule modify collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnapshotPolicyScheduleModifyCollectionParams) SetDefaults() {
	var (
		continueOnFailureDefault = bool(false)

		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)

		serialRecordsDefault = bool(false)
	)

	val := SnapshotPolicyScheduleModifyCollectionParams{
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

// WithTimeout adds the timeout to the snapshot policy schedule modify collection params
func (o *SnapshotPolicyScheduleModifyCollectionParams) WithTimeout(timeout time.Duration) *SnapshotPolicyScheduleModifyCollectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the snapshot policy schedule modify collection params
func (o *SnapshotPolicyScheduleModifyCollectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the snapshot policy schedule modify collection params
func (o *SnapshotPolicyScheduleModifyCollectionParams) WithContext(ctx context.Context) *SnapshotPolicyScheduleModifyCollectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the snapshot policy schedule modify collection params
func (o *SnapshotPolicyScheduleModifyCollectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the snapshot policy schedule modify collection params
func (o *SnapshotPolicyScheduleModifyCollectionParams) WithHTTPClient(client *http.Client) *SnapshotPolicyScheduleModifyCollectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the snapshot policy schedule modify collection params
func (o *SnapshotPolicyScheduleModifyCollectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContinueOnFailure adds the continueOnFailure to the snapshot policy schedule modify collection params
func (o *SnapshotPolicyScheduleModifyCollectionParams) WithContinueOnFailure(continueOnFailure *bool) *SnapshotPolicyScheduleModifyCollectionParams {
	o.SetContinueOnFailure(continueOnFailure)
	return o
}

// SetContinueOnFailure adds the continueOnFailure to the snapshot policy schedule modify collection params
func (o *SnapshotPolicyScheduleModifyCollectionParams) SetContinueOnFailure(continueOnFailure *bool) {
	o.ContinueOnFailure = continueOnFailure
}

// WithCount adds the count to the snapshot policy schedule modify collection params
func (o *SnapshotPolicyScheduleModifyCollectionParams) WithCount(count *int64) *SnapshotPolicyScheduleModifyCollectionParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the snapshot policy schedule modify collection params
func (o *SnapshotPolicyScheduleModifyCollectionParams) SetCount(count *int64) {
	o.Count = count
}

// WithInfo adds the info to the snapshot policy schedule modify collection params
func (o *SnapshotPolicyScheduleModifyCollectionParams) WithInfo(info SnapshotPolicyScheduleModifyCollectionBody) *SnapshotPolicyScheduleModifyCollectionParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the snapshot policy schedule modify collection params
func (o *SnapshotPolicyScheduleModifyCollectionParams) SetInfo(info SnapshotPolicyScheduleModifyCollectionBody) {
	o.Info = info
}

// WithPrefix adds the prefix to the snapshot policy schedule modify collection params
func (o *SnapshotPolicyScheduleModifyCollectionParams) WithPrefix(prefix *string) *SnapshotPolicyScheduleModifyCollectionParams {
	o.SetPrefix(prefix)
	return o
}

// SetPrefix adds the prefix to the snapshot policy schedule modify collection params
func (o *SnapshotPolicyScheduleModifyCollectionParams) SetPrefix(prefix *string) {
	o.Prefix = prefix
}

// WithRetentionPeriod adds the retentionPeriod to the snapshot policy schedule modify collection params
func (o *SnapshotPolicyScheduleModifyCollectionParams) WithRetentionPeriod(retentionPeriod *string) *SnapshotPolicyScheduleModifyCollectionParams {
	o.SetRetentionPeriod(retentionPeriod)
	return o
}

// SetRetentionPeriod adds the retentionPeriod to the snapshot policy schedule modify collection params
func (o *SnapshotPolicyScheduleModifyCollectionParams) SetRetentionPeriod(retentionPeriod *string) {
	o.RetentionPeriod = retentionPeriod
}

// WithReturnRecords adds the returnRecords to the snapshot policy schedule modify collection params
func (o *SnapshotPolicyScheduleModifyCollectionParams) WithReturnRecords(returnRecords *bool) *SnapshotPolicyScheduleModifyCollectionParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the snapshot policy schedule modify collection params
func (o *SnapshotPolicyScheduleModifyCollectionParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the snapshot policy schedule modify collection params
func (o *SnapshotPolicyScheduleModifyCollectionParams) WithReturnTimeout(returnTimeout *int64) *SnapshotPolicyScheduleModifyCollectionParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the snapshot policy schedule modify collection params
func (o *SnapshotPolicyScheduleModifyCollectionParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithScheduleName adds the scheduleName to the snapshot policy schedule modify collection params
func (o *SnapshotPolicyScheduleModifyCollectionParams) WithScheduleName(scheduleName *string) *SnapshotPolicyScheduleModifyCollectionParams {
	o.SetScheduleName(scheduleName)
	return o
}

// SetScheduleName adds the scheduleName to the snapshot policy schedule modify collection params
func (o *SnapshotPolicyScheduleModifyCollectionParams) SetScheduleName(scheduleName *string) {
	o.ScheduleName = scheduleName
}

// WithScheduleUUID adds the scheduleUUID to the snapshot policy schedule modify collection params
func (o *SnapshotPolicyScheduleModifyCollectionParams) WithScheduleUUID(scheduleUUID *string) *SnapshotPolicyScheduleModifyCollectionParams {
	o.SetScheduleUUID(scheduleUUID)
	return o
}

// SetScheduleUUID adds the scheduleUuid to the snapshot policy schedule modify collection params
func (o *SnapshotPolicyScheduleModifyCollectionParams) SetScheduleUUID(scheduleUUID *string) {
	o.ScheduleUUID = scheduleUUID
}

// WithSerialRecords adds the serialRecords to the snapshot policy schedule modify collection params
func (o *SnapshotPolicyScheduleModifyCollectionParams) WithSerialRecords(serialRecords *bool) *SnapshotPolicyScheduleModifyCollectionParams {
	o.SetSerialRecords(serialRecords)
	return o
}

// SetSerialRecords adds the serialRecords to the snapshot policy schedule modify collection params
func (o *SnapshotPolicyScheduleModifyCollectionParams) SetSerialRecords(serialRecords *bool) {
	o.SerialRecords = serialRecords
}

// WithSnapmirrorLabel adds the snapmirrorLabel to the snapshot policy schedule modify collection params
func (o *SnapshotPolicyScheduleModifyCollectionParams) WithSnapmirrorLabel(snapmirrorLabel *string) *SnapshotPolicyScheduleModifyCollectionParams {
	o.SetSnapmirrorLabel(snapmirrorLabel)
	return o
}

// SetSnapmirrorLabel adds the snapmirrorLabel to the snapshot policy schedule modify collection params
func (o *SnapshotPolicyScheduleModifyCollectionParams) SetSnapmirrorLabel(snapmirrorLabel *string) {
	o.SnapmirrorLabel = snapmirrorLabel
}

// WithSnapshotPolicyName adds the snapshotPolicyName to the snapshot policy schedule modify collection params
func (o *SnapshotPolicyScheduleModifyCollectionParams) WithSnapshotPolicyName(snapshotPolicyName *string) *SnapshotPolicyScheduleModifyCollectionParams {
	o.SetSnapshotPolicyName(snapshotPolicyName)
	return o
}

// SetSnapshotPolicyName adds the snapshotPolicyName to the snapshot policy schedule modify collection params
func (o *SnapshotPolicyScheduleModifyCollectionParams) SetSnapshotPolicyName(snapshotPolicyName *string) {
	o.SnapshotPolicyName = snapshotPolicyName
}

// WithSnapshotPolicyUUID adds the snapshotPolicyUUID to the snapshot policy schedule modify collection params
func (o *SnapshotPolicyScheduleModifyCollectionParams) WithSnapshotPolicyUUID(snapshotPolicyUUID string) *SnapshotPolicyScheduleModifyCollectionParams {
	o.SetSnapshotPolicyUUID(snapshotPolicyUUID)
	return o
}

// SetSnapshotPolicyUUID adds the snapshotPolicyUuid to the snapshot policy schedule modify collection params
func (o *SnapshotPolicyScheduleModifyCollectionParams) SetSnapshotPolicyUUID(snapshotPolicyUUID string) {
	o.SnapshotPolicyUUID = snapshotPolicyUUID
}

// WriteToRequest writes these params to a swagger request
func (o *SnapshotPolicyScheduleModifyCollectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetBodyParam(o.Info); err != nil {
		return err
	}

	if o.Prefix != nil {

		// query param prefix
		var qrPrefix string

		if o.Prefix != nil {
			qrPrefix = *o.Prefix
		}
		qPrefix := qrPrefix
		if qPrefix != "" {

			if err := r.SetQueryParam("prefix", qPrefix); err != nil {
				return err
			}
		}
	}

	if o.RetentionPeriod != nil {

		// query param retention_period
		var qrRetentionPeriod string

		if o.RetentionPeriod != nil {
			qrRetentionPeriod = *o.RetentionPeriod
		}
		qRetentionPeriod := qrRetentionPeriod
		if qRetentionPeriod != "" {

			if err := r.SetQueryParam("retention_period", qRetentionPeriod); err != nil {
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

	if o.ScheduleName != nil {

		// query param schedule.name
		var qrScheduleName string

		if o.ScheduleName != nil {
			qrScheduleName = *o.ScheduleName
		}
		qScheduleName := qrScheduleName
		if qScheduleName != "" {

			if err := r.SetQueryParam("schedule.name", qScheduleName); err != nil {
				return err
			}
		}
	}

	if o.ScheduleUUID != nil {

		// query param schedule.uuid
		var qrScheduleUUID string

		if o.ScheduleUUID != nil {
			qrScheduleUUID = *o.ScheduleUUID
		}
		qScheduleUUID := qrScheduleUUID
		if qScheduleUUID != "" {

			if err := r.SetQueryParam("schedule.uuid", qScheduleUUID); err != nil {
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

	if o.SnapmirrorLabel != nil {

		// query param snapmirror_label
		var qrSnapmirrorLabel string

		if o.SnapmirrorLabel != nil {
			qrSnapmirrorLabel = *o.SnapmirrorLabel
		}
		qSnapmirrorLabel := qrSnapmirrorLabel
		if qSnapmirrorLabel != "" {

			if err := r.SetQueryParam("snapmirror_label", qSnapmirrorLabel); err != nil {
				return err
			}
		}
	}

	if o.SnapshotPolicyName != nil {

		// query param snapshot_policy.name
		var qrSnapshotPolicyName string

		if o.SnapshotPolicyName != nil {
			qrSnapshotPolicyName = *o.SnapshotPolicyName
		}
		qSnapshotPolicyName := qrSnapshotPolicyName
		if qSnapshotPolicyName != "" {

			if err := r.SetQueryParam("snapshot_policy.name", qSnapshotPolicyName); err != nil {
				return err
			}
		}
	}

	// path param snapshot_policy.uuid
	if err := r.SetPathParam("snapshot_policy.uuid", o.SnapshotPolicyUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
