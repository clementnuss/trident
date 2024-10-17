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

// NewStorageAvailabilityZoneCollectionGetParams creates a new StorageAvailabilityZoneCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStorageAvailabilityZoneCollectionGetParams() *StorageAvailabilityZoneCollectionGetParams {
	return &StorageAvailabilityZoneCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStorageAvailabilityZoneCollectionGetParamsWithTimeout creates a new StorageAvailabilityZoneCollectionGetParams object
// with the ability to set a timeout on a request.
func NewStorageAvailabilityZoneCollectionGetParamsWithTimeout(timeout time.Duration) *StorageAvailabilityZoneCollectionGetParams {
	return &StorageAvailabilityZoneCollectionGetParams{
		timeout: timeout,
	}
}

// NewStorageAvailabilityZoneCollectionGetParamsWithContext creates a new StorageAvailabilityZoneCollectionGetParams object
// with the ability to set a context for a request.
func NewStorageAvailabilityZoneCollectionGetParamsWithContext(ctx context.Context) *StorageAvailabilityZoneCollectionGetParams {
	return &StorageAvailabilityZoneCollectionGetParams{
		Context: ctx,
	}
}

// NewStorageAvailabilityZoneCollectionGetParamsWithHTTPClient creates a new StorageAvailabilityZoneCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewStorageAvailabilityZoneCollectionGetParamsWithHTTPClient(client *http.Client) *StorageAvailabilityZoneCollectionGetParams {
	return &StorageAvailabilityZoneCollectionGetParams{
		HTTPClient: client,
	}
}

/*
StorageAvailabilityZoneCollectionGetParams contains all the parameters to send to the API endpoint

	for the storage availability zone collection get operation.

	Typically these are written to a http.Request.
*/
type StorageAvailabilityZoneCollectionGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecords *int64

	/* Name.

	   Filter by name
	*/
	Name *string

	/* NodesName.

	   Filter by nodes.name
	*/
	NodesName *string

	/* NodesUUID.

	   Filter by nodes.uuid
	*/
	NodesUUID *string

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderBy []string

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

	/* SpaceAvailable.

	   Filter by space.available
	*/
	SpaceAvailable *int64

	/* SpaceDelayedFrees.

	   Filter by space.delayed_frees
	*/
	SpaceDelayedFrees *int64

	/* SpaceEfficiencyWithoutSnapshotsRatio.

	   Filter by space.efficiency_without_snapshots.ratio
	*/
	SpaceEfficiencyWithoutSnapshotsRatio *float64

	/* SpaceEfficiencyWithoutSnapshotsSavings.

	   Filter by space.efficiency_without_snapshots.savings
	*/
	SpaceEfficiencyWithoutSnapshotsSavings *int64

	/* SpaceFullThresholdPercent.

	   Filter by space.full_threshold_percent
	*/
	SpaceFullThresholdPercent *int64

	/* SpaceInactiveData.

	   Filter by space.inactive_data
	*/
	SpaceInactiveData *int64

	/* SpaceLogicalUserDataWithoutSnapshots.

	   Filter by space.logical_user_data_without_snapshots
	*/
	SpaceLogicalUserDataWithoutSnapshots *int64

	/* SpaceNearlyFullThresholdPercent.

	   Filter by space.nearly_full_threshold_percent
	*/
	SpaceNearlyFullThresholdPercent *int64

	/* SpacePhysicalUsed.

	   Filter by space.physical_used
	*/
	SpacePhysicalUsed *int64

	/* SpacePhysicalUsedPercent.

	   Filter by space.physical_used_percent
	*/
	SpacePhysicalUsedPercent *int64

	/* SpacePhysicalUserDataWithoutSnapshots.

	   Filter by space.physical_user_data_without_snapshots
	*/
	SpacePhysicalUserDataWithoutSnapshots *int64

	/* SpaceSize.

	   Filter by space.size
	*/
	SpaceSize *int64

	/* SpaceTotalMetadataUsed.

	   Filter by space.total_metadata_used
	*/
	SpaceTotalMetadataUsed *int64

	/* SpaceUnusable.

	   Filter by space.unusable
	*/
	SpaceUnusable *int64

	/* UUID.

	   Filter by uuid
	*/
	UUID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the storage availability zone collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StorageAvailabilityZoneCollectionGetParams) WithDefaults() *StorageAvailabilityZoneCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the storage availability zone collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StorageAvailabilityZoneCollectionGetParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)
	)

	val := StorageAvailabilityZoneCollectionGetParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the storage availability zone collection get params
func (o *StorageAvailabilityZoneCollectionGetParams) WithTimeout(timeout time.Duration) *StorageAvailabilityZoneCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage availability zone collection get params
func (o *StorageAvailabilityZoneCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the storage availability zone collection get params
func (o *StorageAvailabilityZoneCollectionGetParams) WithContext(ctx context.Context) *StorageAvailabilityZoneCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage availability zone collection get params
func (o *StorageAvailabilityZoneCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage availability zone collection get params
func (o *StorageAvailabilityZoneCollectionGetParams) WithHTTPClient(client *http.Client) *StorageAvailabilityZoneCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage availability zone collection get params
func (o *StorageAvailabilityZoneCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the storage availability zone collection get params
func (o *StorageAvailabilityZoneCollectionGetParams) WithFields(fields []string) *StorageAvailabilityZoneCollectionGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the storage availability zone collection get params
func (o *StorageAvailabilityZoneCollectionGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithMaxRecords adds the maxRecords to the storage availability zone collection get params
func (o *StorageAvailabilityZoneCollectionGetParams) WithMaxRecords(maxRecords *int64) *StorageAvailabilityZoneCollectionGetParams {
	o.SetMaxRecords(maxRecords)
	return o
}

// SetMaxRecords adds the maxRecords to the storage availability zone collection get params
func (o *StorageAvailabilityZoneCollectionGetParams) SetMaxRecords(maxRecords *int64) {
	o.MaxRecords = maxRecords
}

// WithName adds the name to the storage availability zone collection get params
func (o *StorageAvailabilityZoneCollectionGetParams) WithName(name *string) *StorageAvailabilityZoneCollectionGetParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the storage availability zone collection get params
func (o *StorageAvailabilityZoneCollectionGetParams) SetName(name *string) {
	o.Name = name
}

// WithNodesName adds the nodesName to the storage availability zone collection get params
func (o *StorageAvailabilityZoneCollectionGetParams) WithNodesName(nodesName *string) *StorageAvailabilityZoneCollectionGetParams {
	o.SetNodesName(nodesName)
	return o
}

// SetNodesName adds the nodesName to the storage availability zone collection get params
func (o *StorageAvailabilityZoneCollectionGetParams) SetNodesName(nodesName *string) {
	o.NodesName = nodesName
}

// WithNodesUUID adds the nodesUUID to the storage availability zone collection get params
func (o *StorageAvailabilityZoneCollectionGetParams) WithNodesUUID(nodesUUID *string) *StorageAvailabilityZoneCollectionGetParams {
	o.SetNodesUUID(nodesUUID)
	return o
}

// SetNodesUUID adds the nodesUuid to the storage availability zone collection get params
func (o *StorageAvailabilityZoneCollectionGetParams) SetNodesUUID(nodesUUID *string) {
	o.NodesUUID = nodesUUID
}

// WithOrderBy adds the orderBy to the storage availability zone collection get params
func (o *StorageAvailabilityZoneCollectionGetParams) WithOrderBy(orderBy []string) *StorageAvailabilityZoneCollectionGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the storage availability zone collection get params
func (o *StorageAvailabilityZoneCollectionGetParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithReturnRecords adds the returnRecords to the storage availability zone collection get params
func (o *StorageAvailabilityZoneCollectionGetParams) WithReturnRecords(returnRecords *bool) *StorageAvailabilityZoneCollectionGetParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the storage availability zone collection get params
func (o *StorageAvailabilityZoneCollectionGetParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the storage availability zone collection get params
func (o *StorageAvailabilityZoneCollectionGetParams) WithReturnTimeout(returnTimeout *int64) *StorageAvailabilityZoneCollectionGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the storage availability zone collection get params
func (o *StorageAvailabilityZoneCollectionGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSpaceAvailable adds the spaceAvailable to the storage availability zone collection get params
func (o *StorageAvailabilityZoneCollectionGetParams) WithSpaceAvailable(spaceAvailable *int64) *StorageAvailabilityZoneCollectionGetParams {
	o.SetSpaceAvailable(spaceAvailable)
	return o
}

// SetSpaceAvailable adds the spaceAvailable to the storage availability zone collection get params
func (o *StorageAvailabilityZoneCollectionGetParams) SetSpaceAvailable(spaceAvailable *int64) {
	o.SpaceAvailable = spaceAvailable
}

// WithSpaceDelayedFrees adds the spaceDelayedFrees to the storage availability zone collection get params
func (o *StorageAvailabilityZoneCollectionGetParams) WithSpaceDelayedFrees(spaceDelayedFrees *int64) *StorageAvailabilityZoneCollectionGetParams {
	o.SetSpaceDelayedFrees(spaceDelayedFrees)
	return o
}

// SetSpaceDelayedFrees adds the spaceDelayedFrees to the storage availability zone collection get params
func (o *StorageAvailabilityZoneCollectionGetParams) SetSpaceDelayedFrees(spaceDelayedFrees *int64) {
	o.SpaceDelayedFrees = spaceDelayedFrees
}

// WithSpaceEfficiencyWithoutSnapshotsRatio adds the spaceEfficiencyWithoutSnapshotsRatio to the storage availability zone collection get params
func (o *StorageAvailabilityZoneCollectionGetParams) WithSpaceEfficiencyWithoutSnapshotsRatio(spaceEfficiencyWithoutSnapshotsRatio *float64) *StorageAvailabilityZoneCollectionGetParams {
	o.SetSpaceEfficiencyWithoutSnapshotsRatio(spaceEfficiencyWithoutSnapshotsRatio)
	return o
}

// SetSpaceEfficiencyWithoutSnapshotsRatio adds the spaceEfficiencyWithoutSnapshotsRatio to the storage availability zone collection get params
func (o *StorageAvailabilityZoneCollectionGetParams) SetSpaceEfficiencyWithoutSnapshotsRatio(spaceEfficiencyWithoutSnapshotsRatio *float64) {
	o.SpaceEfficiencyWithoutSnapshotsRatio = spaceEfficiencyWithoutSnapshotsRatio
}

// WithSpaceEfficiencyWithoutSnapshotsSavings adds the spaceEfficiencyWithoutSnapshotsSavings to the storage availability zone collection get params
func (o *StorageAvailabilityZoneCollectionGetParams) WithSpaceEfficiencyWithoutSnapshotsSavings(spaceEfficiencyWithoutSnapshotsSavings *int64) *StorageAvailabilityZoneCollectionGetParams {
	o.SetSpaceEfficiencyWithoutSnapshotsSavings(spaceEfficiencyWithoutSnapshotsSavings)
	return o
}

// SetSpaceEfficiencyWithoutSnapshotsSavings adds the spaceEfficiencyWithoutSnapshotsSavings to the storage availability zone collection get params
func (o *StorageAvailabilityZoneCollectionGetParams) SetSpaceEfficiencyWithoutSnapshotsSavings(spaceEfficiencyWithoutSnapshotsSavings *int64) {
	o.SpaceEfficiencyWithoutSnapshotsSavings = spaceEfficiencyWithoutSnapshotsSavings
}

// WithSpaceFullThresholdPercent adds the spaceFullThresholdPercent to the storage availability zone collection get params
func (o *StorageAvailabilityZoneCollectionGetParams) WithSpaceFullThresholdPercent(spaceFullThresholdPercent *int64) *StorageAvailabilityZoneCollectionGetParams {
	o.SetSpaceFullThresholdPercent(spaceFullThresholdPercent)
	return o
}

// SetSpaceFullThresholdPercent adds the spaceFullThresholdPercent to the storage availability zone collection get params
func (o *StorageAvailabilityZoneCollectionGetParams) SetSpaceFullThresholdPercent(spaceFullThresholdPercent *int64) {
	o.SpaceFullThresholdPercent = spaceFullThresholdPercent
}

// WithSpaceInactiveData adds the spaceInactiveData to the storage availability zone collection get params
func (o *StorageAvailabilityZoneCollectionGetParams) WithSpaceInactiveData(spaceInactiveData *int64) *StorageAvailabilityZoneCollectionGetParams {
	o.SetSpaceInactiveData(spaceInactiveData)
	return o
}

// SetSpaceInactiveData adds the spaceInactiveData to the storage availability zone collection get params
func (o *StorageAvailabilityZoneCollectionGetParams) SetSpaceInactiveData(spaceInactiveData *int64) {
	o.SpaceInactiveData = spaceInactiveData
}

// WithSpaceLogicalUserDataWithoutSnapshots adds the spaceLogicalUserDataWithoutSnapshots to the storage availability zone collection get params
func (o *StorageAvailabilityZoneCollectionGetParams) WithSpaceLogicalUserDataWithoutSnapshots(spaceLogicalUserDataWithoutSnapshots *int64) *StorageAvailabilityZoneCollectionGetParams {
	o.SetSpaceLogicalUserDataWithoutSnapshots(spaceLogicalUserDataWithoutSnapshots)
	return o
}

// SetSpaceLogicalUserDataWithoutSnapshots adds the spaceLogicalUserDataWithoutSnapshots to the storage availability zone collection get params
func (o *StorageAvailabilityZoneCollectionGetParams) SetSpaceLogicalUserDataWithoutSnapshots(spaceLogicalUserDataWithoutSnapshots *int64) {
	o.SpaceLogicalUserDataWithoutSnapshots = spaceLogicalUserDataWithoutSnapshots
}

// WithSpaceNearlyFullThresholdPercent adds the spaceNearlyFullThresholdPercent to the storage availability zone collection get params
func (o *StorageAvailabilityZoneCollectionGetParams) WithSpaceNearlyFullThresholdPercent(spaceNearlyFullThresholdPercent *int64) *StorageAvailabilityZoneCollectionGetParams {
	o.SetSpaceNearlyFullThresholdPercent(spaceNearlyFullThresholdPercent)
	return o
}

// SetSpaceNearlyFullThresholdPercent adds the spaceNearlyFullThresholdPercent to the storage availability zone collection get params
func (o *StorageAvailabilityZoneCollectionGetParams) SetSpaceNearlyFullThresholdPercent(spaceNearlyFullThresholdPercent *int64) {
	o.SpaceNearlyFullThresholdPercent = spaceNearlyFullThresholdPercent
}

// WithSpacePhysicalUsed adds the spacePhysicalUsed to the storage availability zone collection get params
func (o *StorageAvailabilityZoneCollectionGetParams) WithSpacePhysicalUsed(spacePhysicalUsed *int64) *StorageAvailabilityZoneCollectionGetParams {
	o.SetSpacePhysicalUsed(spacePhysicalUsed)
	return o
}

// SetSpacePhysicalUsed adds the spacePhysicalUsed to the storage availability zone collection get params
func (o *StorageAvailabilityZoneCollectionGetParams) SetSpacePhysicalUsed(spacePhysicalUsed *int64) {
	o.SpacePhysicalUsed = spacePhysicalUsed
}

// WithSpacePhysicalUsedPercent adds the spacePhysicalUsedPercent to the storage availability zone collection get params
func (o *StorageAvailabilityZoneCollectionGetParams) WithSpacePhysicalUsedPercent(spacePhysicalUsedPercent *int64) *StorageAvailabilityZoneCollectionGetParams {
	o.SetSpacePhysicalUsedPercent(spacePhysicalUsedPercent)
	return o
}

// SetSpacePhysicalUsedPercent adds the spacePhysicalUsedPercent to the storage availability zone collection get params
func (o *StorageAvailabilityZoneCollectionGetParams) SetSpacePhysicalUsedPercent(spacePhysicalUsedPercent *int64) {
	o.SpacePhysicalUsedPercent = spacePhysicalUsedPercent
}

// WithSpacePhysicalUserDataWithoutSnapshots adds the spacePhysicalUserDataWithoutSnapshots to the storage availability zone collection get params
func (o *StorageAvailabilityZoneCollectionGetParams) WithSpacePhysicalUserDataWithoutSnapshots(spacePhysicalUserDataWithoutSnapshots *int64) *StorageAvailabilityZoneCollectionGetParams {
	o.SetSpacePhysicalUserDataWithoutSnapshots(spacePhysicalUserDataWithoutSnapshots)
	return o
}

// SetSpacePhysicalUserDataWithoutSnapshots adds the spacePhysicalUserDataWithoutSnapshots to the storage availability zone collection get params
func (o *StorageAvailabilityZoneCollectionGetParams) SetSpacePhysicalUserDataWithoutSnapshots(spacePhysicalUserDataWithoutSnapshots *int64) {
	o.SpacePhysicalUserDataWithoutSnapshots = spacePhysicalUserDataWithoutSnapshots
}

// WithSpaceSize adds the spaceSize to the storage availability zone collection get params
func (o *StorageAvailabilityZoneCollectionGetParams) WithSpaceSize(spaceSize *int64) *StorageAvailabilityZoneCollectionGetParams {
	o.SetSpaceSize(spaceSize)
	return o
}

// SetSpaceSize adds the spaceSize to the storage availability zone collection get params
func (o *StorageAvailabilityZoneCollectionGetParams) SetSpaceSize(spaceSize *int64) {
	o.SpaceSize = spaceSize
}

// WithSpaceTotalMetadataUsed adds the spaceTotalMetadataUsed to the storage availability zone collection get params
func (o *StorageAvailabilityZoneCollectionGetParams) WithSpaceTotalMetadataUsed(spaceTotalMetadataUsed *int64) *StorageAvailabilityZoneCollectionGetParams {
	o.SetSpaceTotalMetadataUsed(spaceTotalMetadataUsed)
	return o
}

// SetSpaceTotalMetadataUsed adds the spaceTotalMetadataUsed to the storage availability zone collection get params
func (o *StorageAvailabilityZoneCollectionGetParams) SetSpaceTotalMetadataUsed(spaceTotalMetadataUsed *int64) {
	o.SpaceTotalMetadataUsed = spaceTotalMetadataUsed
}

// WithSpaceUnusable adds the spaceUnusable to the storage availability zone collection get params
func (o *StorageAvailabilityZoneCollectionGetParams) WithSpaceUnusable(spaceUnusable *int64) *StorageAvailabilityZoneCollectionGetParams {
	o.SetSpaceUnusable(spaceUnusable)
	return o
}

// SetSpaceUnusable adds the spaceUnusable to the storage availability zone collection get params
func (o *StorageAvailabilityZoneCollectionGetParams) SetSpaceUnusable(spaceUnusable *int64) {
	o.SpaceUnusable = spaceUnusable
}

// WithUUID adds the uuid to the storage availability zone collection get params
func (o *StorageAvailabilityZoneCollectionGetParams) WithUUID(uuid *string) *StorageAvailabilityZoneCollectionGetParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the storage availability zone collection get params
func (o *StorageAvailabilityZoneCollectionGetParams) SetUUID(uuid *string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *StorageAvailabilityZoneCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Fields != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	if o.MaxRecords != nil {

		// query param max_records
		var qrMaxRecords int64

		if o.MaxRecords != nil {
			qrMaxRecords = *o.MaxRecords
		}
		qMaxRecords := swag.FormatInt64(qrMaxRecords)
		if qMaxRecords != "" {

			if err := r.SetQueryParam("max_records", qMaxRecords); err != nil {
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

	if o.NodesName != nil {

		// query param nodes.name
		var qrNodesName string

		if o.NodesName != nil {
			qrNodesName = *o.NodesName
		}
		qNodesName := qrNodesName
		if qNodesName != "" {

			if err := r.SetQueryParam("nodes.name", qNodesName); err != nil {
				return err
			}
		}
	}

	if o.NodesUUID != nil {

		// query param nodes.uuid
		var qrNodesUUID string

		if o.NodesUUID != nil {
			qrNodesUUID = *o.NodesUUID
		}
		qNodesUUID := qrNodesUUID
		if qNodesUUID != "" {

			if err := r.SetQueryParam("nodes.uuid", qNodesUUID); err != nil {
				return err
			}
		}
	}

	if o.OrderBy != nil {

		// binding items for order_by
		joinedOrderBy := o.bindParamOrderBy(reg)

		// query array param order_by
		if err := r.SetQueryParam("order_by", joinedOrderBy...); err != nil {
			return err
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

	if o.SpaceAvailable != nil {

		// query param space.available
		var qrSpaceAvailable int64

		if o.SpaceAvailable != nil {
			qrSpaceAvailable = *o.SpaceAvailable
		}
		qSpaceAvailable := swag.FormatInt64(qrSpaceAvailable)
		if qSpaceAvailable != "" {

			if err := r.SetQueryParam("space.available", qSpaceAvailable); err != nil {
				return err
			}
		}
	}

	if o.SpaceDelayedFrees != nil {

		// query param space.delayed_frees
		var qrSpaceDelayedFrees int64

		if o.SpaceDelayedFrees != nil {
			qrSpaceDelayedFrees = *o.SpaceDelayedFrees
		}
		qSpaceDelayedFrees := swag.FormatInt64(qrSpaceDelayedFrees)
		if qSpaceDelayedFrees != "" {

			if err := r.SetQueryParam("space.delayed_frees", qSpaceDelayedFrees); err != nil {
				return err
			}
		}
	}

	if o.SpaceEfficiencyWithoutSnapshotsRatio != nil {

		// query param space.efficiency_without_snapshots.ratio
		var qrSpaceEfficiencyWithoutSnapshotsRatio float64

		if o.SpaceEfficiencyWithoutSnapshotsRatio != nil {
			qrSpaceEfficiencyWithoutSnapshotsRatio = *o.SpaceEfficiencyWithoutSnapshotsRatio
		}
		qSpaceEfficiencyWithoutSnapshotsRatio := swag.FormatFloat64(qrSpaceEfficiencyWithoutSnapshotsRatio)
		if qSpaceEfficiencyWithoutSnapshotsRatio != "" {

			if err := r.SetQueryParam("space.efficiency_without_snapshots.ratio", qSpaceEfficiencyWithoutSnapshotsRatio); err != nil {
				return err
			}
		}
	}

	if o.SpaceEfficiencyWithoutSnapshotsSavings != nil {

		// query param space.efficiency_without_snapshots.savings
		var qrSpaceEfficiencyWithoutSnapshotsSavings int64

		if o.SpaceEfficiencyWithoutSnapshotsSavings != nil {
			qrSpaceEfficiencyWithoutSnapshotsSavings = *o.SpaceEfficiencyWithoutSnapshotsSavings
		}
		qSpaceEfficiencyWithoutSnapshotsSavings := swag.FormatInt64(qrSpaceEfficiencyWithoutSnapshotsSavings)
		if qSpaceEfficiencyWithoutSnapshotsSavings != "" {

			if err := r.SetQueryParam("space.efficiency_without_snapshots.savings", qSpaceEfficiencyWithoutSnapshotsSavings); err != nil {
				return err
			}
		}
	}

	if o.SpaceFullThresholdPercent != nil {

		// query param space.full_threshold_percent
		var qrSpaceFullThresholdPercent int64

		if o.SpaceFullThresholdPercent != nil {
			qrSpaceFullThresholdPercent = *o.SpaceFullThresholdPercent
		}
		qSpaceFullThresholdPercent := swag.FormatInt64(qrSpaceFullThresholdPercent)
		if qSpaceFullThresholdPercent != "" {

			if err := r.SetQueryParam("space.full_threshold_percent", qSpaceFullThresholdPercent); err != nil {
				return err
			}
		}
	}

	if o.SpaceInactiveData != nil {

		// query param space.inactive_data
		var qrSpaceInactiveData int64

		if o.SpaceInactiveData != nil {
			qrSpaceInactiveData = *o.SpaceInactiveData
		}
		qSpaceInactiveData := swag.FormatInt64(qrSpaceInactiveData)
		if qSpaceInactiveData != "" {

			if err := r.SetQueryParam("space.inactive_data", qSpaceInactiveData); err != nil {
				return err
			}
		}
	}

	if o.SpaceLogicalUserDataWithoutSnapshots != nil {

		// query param space.logical_user_data_without_snapshots
		var qrSpaceLogicalUserDataWithoutSnapshots int64

		if o.SpaceLogicalUserDataWithoutSnapshots != nil {
			qrSpaceLogicalUserDataWithoutSnapshots = *o.SpaceLogicalUserDataWithoutSnapshots
		}
		qSpaceLogicalUserDataWithoutSnapshots := swag.FormatInt64(qrSpaceLogicalUserDataWithoutSnapshots)
		if qSpaceLogicalUserDataWithoutSnapshots != "" {

			if err := r.SetQueryParam("space.logical_user_data_without_snapshots", qSpaceLogicalUserDataWithoutSnapshots); err != nil {
				return err
			}
		}
	}

	if o.SpaceNearlyFullThresholdPercent != nil {

		// query param space.nearly_full_threshold_percent
		var qrSpaceNearlyFullThresholdPercent int64

		if o.SpaceNearlyFullThresholdPercent != nil {
			qrSpaceNearlyFullThresholdPercent = *o.SpaceNearlyFullThresholdPercent
		}
		qSpaceNearlyFullThresholdPercent := swag.FormatInt64(qrSpaceNearlyFullThresholdPercent)
		if qSpaceNearlyFullThresholdPercent != "" {

			if err := r.SetQueryParam("space.nearly_full_threshold_percent", qSpaceNearlyFullThresholdPercent); err != nil {
				return err
			}
		}
	}

	if o.SpacePhysicalUsed != nil {

		// query param space.physical_used
		var qrSpacePhysicalUsed int64

		if o.SpacePhysicalUsed != nil {
			qrSpacePhysicalUsed = *o.SpacePhysicalUsed
		}
		qSpacePhysicalUsed := swag.FormatInt64(qrSpacePhysicalUsed)
		if qSpacePhysicalUsed != "" {

			if err := r.SetQueryParam("space.physical_used", qSpacePhysicalUsed); err != nil {
				return err
			}
		}
	}

	if o.SpacePhysicalUsedPercent != nil {

		// query param space.physical_used_percent
		var qrSpacePhysicalUsedPercent int64

		if o.SpacePhysicalUsedPercent != nil {
			qrSpacePhysicalUsedPercent = *o.SpacePhysicalUsedPercent
		}
		qSpacePhysicalUsedPercent := swag.FormatInt64(qrSpacePhysicalUsedPercent)
		if qSpacePhysicalUsedPercent != "" {

			if err := r.SetQueryParam("space.physical_used_percent", qSpacePhysicalUsedPercent); err != nil {
				return err
			}
		}
	}

	if o.SpacePhysicalUserDataWithoutSnapshots != nil {

		// query param space.physical_user_data_without_snapshots
		var qrSpacePhysicalUserDataWithoutSnapshots int64

		if o.SpacePhysicalUserDataWithoutSnapshots != nil {
			qrSpacePhysicalUserDataWithoutSnapshots = *o.SpacePhysicalUserDataWithoutSnapshots
		}
		qSpacePhysicalUserDataWithoutSnapshots := swag.FormatInt64(qrSpacePhysicalUserDataWithoutSnapshots)
		if qSpacePhysicalUserDataWithoutSnapshots != "" {

			if err := r.SetQueryParam("space.physical_user_data_without_snapshots", qSpacePhysicalUserDataWithoutSnapshots); err != nil {
				return err
			}
		}
	}

	if o.SpaceSize != nil {

		// query param space.size
		var qrSpaceSize int64

		if o.SpaceSize != nil {
			qrSpaceSize = *o.SpaceSize
		}
		qSpaceSize := swag.FormatInt64(qrSpaceSize)
		if qSpaceSize != "" {

			if err := r.SetQueryParam("space.size", qSpaceSize); err != nil {
				return err
			}
		}
	}

	if o.SpaceTotalMetadataUsed != nil {

		// query param space.total_metadata_used
		var qrSpaceTotalMetadataUsed int64

		if o.SpaceTotalMetadataUsed != nil {
			qrSpaceTotalMetadataUsed = *o.SpaceTotalMetadataUsed
		}
		qSpaceTotalMetadataUsed := swag.FormatInt64(qrSpaceTotalMetadataUsed)
		if qSpaceTotalMetadataUsed != "" {

			if err := r.SetQueryParam("space.total_metadata_used", qSpaceTotalMetadataUsed); err != nil {
				return err
			}
		}
	}

	if o.SpaceUnusable != nil {

		// query param space.unusable
		var qrSpaceUnusable int64

		if o.SpaceUnusable != nil {
			qrSpaceUnusable = *o.SpaceUnusable
		}
		qSpaceUnusable := swag.FormatInt64(qrSpaceUnusable)
		if qSpaceUnusable != "" {

			if err := r.SetQueryParam("space.unusable", qSpaceUnusable); err != nil {
				return err
			}
		}
	}

	if o.UUID != nil {

		// query param uuid
		var qrUUID string

		if o.UUID != nil {
			qrUUID = *o.UUID
		}
		qUUID := qrUUID
		if qUUID != "" {

			if err := r.SetQueryParam("uuid", qUUID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamStorageAvailabilityZoneCollectionGet binds the parameter fields
func (o *StorageAvailabilityZoneCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
	fieldsIR := o.Fields

	var fieldsIC []string
	for _, fieldsIIR := range fieldsIR { // explode []string

		fieldsIIV := fieldsIIR // string as string
		fieldsIC = append(fieldsIC, fieldsIIV)
	}

	// items.CollectionFormat: "csv"
	fieldsIS := swag.JoinByFormat(fieldsIC, "csv")

	return fieldsIS
}

// bindParamStorageAvailabilityZoneCollectionGet binds the parameter order_by
func (o *StorageAvailabilityZoneCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
	orderByIR := o.OrderBy

	var orderByIC []string
	for _, orderByIIR := range orderByIR { // explode []string

		orderByIIV := orderByIIR // string as string
		orderByIC = append(orderByIC, orderByIIV)
	}

	// items.CollectionFormat: "csv"
	orderByIS := swag.JoinByFormat(orderByIC, "csv")

	return orderByIS
}