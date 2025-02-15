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

// NewLunMapCollectionGetParams creates a new LunMapCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLunMapCollectionGetParams() *LunMapCollectionGetParams {
	return &LunMapCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLunMapCollectionGetParamsWithTimeout creates a new LunMapCollectionGetParams object
// with the ability to set a timeout on a request.
func NewLunMapCollectionGetParamsWithTimeout(timeout time.Duration) *LunMapCollectionGetParams {
	return &LunMapCollectionGetParams{
		timeout: timeout,
	}
}

// NewLunMapCollectionGetParamsWithContext creates a new LunMapCollectionGetParams object
// with the ability to set a context for a request.
func NewLunMapCollectionGetParamsWithContext(ctx context.Context) *LunMapCollectionGetParams {
	return &LunMapCollectionGetParams{
		Context: ctx,
	}
}

// NewLunMapCollectionGetParamsWithHTTPClient creates a new LunMapCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewLunMapCollectionGetParamsWithHTTPClient(client *http.Client) *LunMapCollectionGetParams {
	return &LunMapCollectionGetParams{
		HTTPClient: client,
	}
}

/*
LunMapCollectionGetParams contains all the parameters to send to the API endpoint

	for the lun map collection get operation.

	Typically these are written to a http.Request.
*/
type LunMapCollectionGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* IgroupInitiators.

	   Filter by igroup.initiators
	*/
	IgroupInitiators *string

	/* IgroupName.

	   Filter by igroup.name
	*/
	IgroupName *string

	/* IgroupOsType.

	   Filter by igroup.os_type
	*/
	IgroupOsType *string

	/* IgroupProtocol.

	   Filter by igroup.protocol
	*/
	IgroupProtocol *string

	/* IgroupReplicated.

	   Filter by igroup.replicated
	*/
	IgroupReplicated *bool

	/* IgroupUUID.

	   Filter by igroup.uuid
	*/
	IgroupUUID *string

	/* LogicalUnitNumber.

	   Filter by logical_unit_number
	*/
	LogicalUnitNumber *int64

	/* LunName.

	   Filter by lun.name
	*/
	LunName *string

	/* LunNodeName.

	   Filter by lun.node.name
	*/
	LunNodeName *string

	/* LunNodeUUID.

	   Filter by lun.node.uuid
	*/
	LunNodeUUID *string

	/* LunSmbcReplicated.

	   Filter by lun.smbc.replicated
	*/
	LunSmbcReplicated *bool

	/* LunUUID.

	   Filter by lun.uuid
	*/
	LunUUID *string

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecords *int64

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderBy []string

	/* ReportingNodesName.

	   Filter by reporting_nodes.name
	*/
	ReportingNodesName *string

	/* ReportingNodesUUID.

	   Filter by reporting_nodes.uuid
	*/
	ReportingNodesUUID *string

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

// WithDefaults hydrates default values in the lun map collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LunMapCollectionGetParams) WithDefaults() *LunMapCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the lun map collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LunMapCollectionGetParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)
	)

	val := LunMapCollectionGetParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the lun map collection get params
func (o *LunMapCollectionGetParams) WithTimeout(timeout time.Duration) *LunMapCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the lun map collection get params
func (o *LunMapCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the lun map collection get params
func (o *LunMapCollectionGetParams) WithContext(ctx context.Context) *LunMapCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the lun map collection get params
func (o *LunMapCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the lun map collection get params
func (o *LunMapCollectionGetParams) WithHTTPClient(client *http.Client) *LunMapCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the lun map collection get params
func (o *LunMapCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the lun map collection get params
func (o *LunMapCollectionGetParams) WithFields(fields []string) *LunMapCollectionGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the lun map collection get params
func (o *LunMapCollectionGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithIgroupInitiators adds the igroupInitiators to the lun map collection get params
func (o *LunMapCollectionGetParams) WithIgroupInitiators(igroupInitiators *string) *LunMapCollectionGetParams {
	o.SetIgroupInitiators(igroupInitiators)
	return o
}

// SetIgroupInitiators adds the igroupInitiators to the lun map collection get params
func (o *LunMapCollectionGetParams) SetIgroupInitiators(igroupInitiators *string) {
	o.IgroupInitiators = igroupInitiators
}

// WithIgroupName adds the igroupName to the lun map collection get params
func (o *LunMapCollectionGetParams) WithIgroupName(igroupName *string) *LunMapCollectionGetParams {
	o.SetIgroupName(igroupName)
	return o
}

// SetIgroupName adds the igroupName to the lun map collection get params
func (o *LunMapCollectionGetParams) SetIgroupName(igroupName *string) {
	o.IgroupName = igroupName
}

// WithIgroupOsType adds the igroupOsType to the lun map collection get params
func (o *LunMapCollectionGetParams) WithIgroupOsType(igroupOsType *string) *LunMapCollectionGetParams {
	o.SetIgroupOsType(igroupOsType)
	return o
}

// SetIgroupOsType adds the igroupOsType to the lun map collection get params
func (o *LunMapCollectionGetParams) SetIgroupOsType(igroupOsType *string) {
	o.IgroupOsType = igroupOsType
}

// WithIgroupProtocol adds the igroupProtocol to the lun map collection get params
func (o *LunMapCollectionGetParams) WithIgroupProtocol(igroupProtocol *string) *LunMapCollectionGetParams {
	o.SetIgroupProtocol(igroupProtocol)
	return o
}

// SetIgroupProtocol adds the igroupProtocol to the lun map collection get params
func (o *LunMapCollectionGetParams) SetIgroupProtocol(igroupProtocol *string) {
	o.IgroupProtocol = igroupProtocol
}

// WithIgroupReplicated adds the igroupReplicated to the lun map collection get params
func (o *LunMapCollectionGetParams) WithIgroupReplicated(igroupReplicated *bool) *LunMapCollectionGetParams {
	o.SetIgroupReplicated(igroupReplicated)
	return o
}

// SetIgroupReplicated adds the igroupReplicated to the lun map collection get params
func (o *LunMapCollectionGetParams) SetIgroupReplicated(igroupReplicated *bool) {
	o.IgroupReplicated = igroupReplicated
}

// WithIgroupUUID adds the igroupUUID to the lun map collection get params
func (o *LunMapCollectionGetParams) WithIgroupUUID(igroupUUID *string) *LunMapCollectionGetParams {
	o.SetIgroupUUID(igroupUUID)
	return o
}

// SetIgroupUUID adds the igroupUuid to the lun map collection get params
func (o *LunMapCollectionGetParams) SetIgroupUUID(igroupUUID *string) {
	o.IgroupUUID = igroupUUID
}

// WithLogicalUnitNumber adds the logicalUnitNumber to the lun map collection get params
func (o *LunMapCollectionGetParams) WithLogicalUnitNumber(logicalUnitNumber *int64) *LunMapCollectionGetParams {
	o.SetLogicalUnitNumber(logicalUnitNumber)
	return o
}

// SetLogicalUnitNumber adds the logicalUnitNumber to the lun map collection get params
func (o *LunMapCollectionGetParams) SetLogicalUnitNumber(logicalUnitNumber *int64) {
	o.LogicalUnitNumber = logicalUnitNumber
}

// WithLunName adds the lunName to the lun map collection get params
func (o *LunMapCollectionGetParams) WithLunName(lunName *string) *LunMapCollectionGetParams {
	o.SetLunName(lunName)
	return o
}

// SetLunName adds the lunName to the lun map collection get params
func (o *LunMapCollectionGetParams) SetLunName(lunName *string) {
	o.LunName = lunName
}

// WithLunNodeName adds the lunNodeName to the lun map collection get params
func (o *LunMapCollectionGetParams) WithLunNodeName(lunNodeName *string) *LunMapCollectionGetParams {
	o.SetLunNodeName(lunNodeName)
	return o
}

// SetLunNodeName adds the lunNodeName to the lun map collection get params
func (o *LunMapCollectionGetParams) SetLunNodeName(lunNodeName *string) {
	o.LunNodeName = lunNodeName
}

// WithLunNodeUUID adds the lunNodeUUID to the lun map collection get params
func (o *LunMapCollectionGetParams) WithLunNodeUUID(lunNodeUUID *string) *LunMapCollectionGetParams {
	o.SetLunNodeUUID(lunNodeUUID)
	return o
}

// SetLunNodeUUID adds the lunNodeUuid to the lun map collection get params
func (o *LunMapCollectionGetParams) SetLunNodeUUID(lunNodeUUID *string) {
	o.LunNodeUUID = lunNodeUUID
}

// WithLunSmbcReplicated adds the lunSmbcReplicated to the lun map collection get params
func (o *LunMapCollectionGetParams) WithLunSmbcReplicated(lunSmbcReplicated *bool) *LunMapCollectionGetParams {
	o.SetLunSmbcReplicated(lunSmbcReplicated)
	return o
}

// SetLunSmbcReplicated adds the lunSmbcReplicated to the lun map collection get params
func (o *LunMapCollectionGetParams) SetLunSmbcReplicated(lunSmbcReplicated *bool) {
	o.LunSmbcReplicated = lunSmbcReplicated
}

// WithLunUUID adds the lunUUID to the lun map collection get params
func (o *LunMapCollectionGetParams) WithLunUUID(lunUUID *string) *LunMapCollectionGetParams {
	o.SetLunUUID(lunUUID)
	return o
}

// SetLunUUID adds the lunUuid to the lun map collection get params
func (o *LunMapCollectionGetParams) SetLunUUID(lunUUID *string) {
	o.LunUUID = lunUUID
}

// WithMaxRecords adds the maxRecords to the lun map collection get params
func (o *LunMapCollectionGetParams) WithMaxRecords(maxRecords *int64) *LunMapCollectionGetParams {
	o.SetMaxRecords(maxRecords)
	return o
}

// SetMaxRecords adds the maxRecords to the lun map collection get params
func (o *LunMapCollectionGetParams) SetMaxRecords(maxRecords *int64) {
	o.MaxRecords = maxRecords
}

// WithOrderBy adds the orderBy to the lun map collection get params
func (o *LunMapCollectionGetParams) WithOrderBy(orderBy []string) *LunMapCollectionGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the lun map collection get params
func (o *LunMapCollectionGetParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithReportingNodesName adds the reportingNodesName to the lun map collection get params
func (o *LunMapCollectionGetParams) WithReportingNodesName(reportingNodesName *string) *LunMapCollectionGetParams {
	o.SetReportingNodesName(reportingNodesName)
	return o
}

// SetReportingNodesName adds the reportingNodesName to the lun map collection get params
func (o *LunMapCollectionGetParams) SetReportingNodesName(reportingNodesName *string) {
	o.ReportingNodesName = reportingNodesName
}

// WithReportingNodesUUID adds the reportingNodesUUID to the lun map collection get params
func (o *LunMapCollectionGetParams) WithReportingNodesUUID(reportingNodesUUID *string) *LunMapCollectionGetParams {
	o.SetReportingNodesUUID(reportingNodesUUID)
	return o
}

// SetReportingNodesUUID adds the reportingNodesUuid to the lun map collection get params
func (o *LunMapCollectionGetParams) SetReportingNodesUUID(reportingNodesUUID *string) {
	o.ReportingNodesUUID = reportingNodesUUID
}

// WithReturnRecords adds the returnRecords to the lun map collection get params
func (o *LunMapCollectionGetParams) WithReturnRecords(returnRecords *bool) *LunMapCollectionGetParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the lun map collection get params
func (o *LunMapCollectionGetParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the lun map collection get params
func (o *LunMapCollectionGetParams) WithReturnTimeout(returnTimeout *int64) *LunMapCollectionGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the lun map collection get params
func (o *LunMapCollectionGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSvmName adds the svmName to the lun map collection get params
func (o *LunMapCollectionGetParams) WithSvmName(svmName *string) *LunMapCollectionGetParams {
	o.SetSvmName(svmName)
	return o
}

// SetSvmName adds the svmName to the lun map collection get params
func (o *LunMapCollectionGetParams) SetSvmName(svmName *string) {
	o.SvmName = svmName
}

// WithSvmUUID adds the svmUUID to the lun map collection get params
func (o *LunMapCollectionGetParams) WithSvmUUID(svmUUID *string) *LunMapCollectionGetParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the lun map collection get params
func (o *LunMapCollectionGetParams) SetSvmUUID(svmUUID *string) {
	o.SvmUUID = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *LunMapCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.IgroupInitiators != nil {

		// query param igroup.initiators
		var qrIgroupInitiators string

		if o.IgroupInitiators != nil {
			qrIgroupInitiators = *o.IgroupInitiators
		}
		qIgroupInitiators := qrIgroupInitiators
		if qIgroupInitiators != "" {

			if err := r.SetQueryParam("igroup.initiators", qIgroupInitiators); err != nil {
				return err
			}
		}
	}

	if o.IgroupName != nil {

		// query param igroup.name
		var qrIgroupName string

		if o.IgroupName != nil {
			qrIgroupName = *o.IgroupName
		}
		qIgroupName := qrIgroupName
		if qIgroupName != "" {

			if err := r.SetQueryParam("igroup.name", qIgroupName); err != nil {
				return err
			}
		}
	}

	if o.IgroupOsType != nil {

		// query param igroup.os_type
		var qrIgroupOsType string

		if o.IgroupOsType != nil {
			qrIgroupOsType = *o.IgroupOsType
		}
		qIgroupOsType := qrIgroupOsType
		if qIgroupOsType != "" {

			if err := r.SetQueryParam("igroup.os_type", qIgroupOsType); err != nil {
				return err
			}
		}
	}

	if o.IgroupProtocol != nil {

		// query param igroup.protocol
		var qrIgroupProtocol string

		if o.IgroupProtocol != nil {
			qrIgroupProtocol = *o.IgroupProtocol
		}
		qIgroupProtocol := qrIgroupProtocol
		if qIgroupProtocol != "" {

			if err := r.SetQueryParam("igroup.protocol", qIgroupProtocol); err != nil {
				return err
			}
		}
	}

	if o.IgroupReplicated != nil {

		// query param igroup.replicated
		var qrIgroupReplicated bool

		if o.IgroupReplicated != nil {
			qrIgroupReplicated = *o.IgroupReplicated
		}
		qIgroupReplicated := swag.FormatBool(qrIgroupReplicated)
		if qIgroupReplicated != "" {

			if err := r.SetQueryParam("igroup.replicated", qIgroupReplicated); err != nil {
				return err
			}
		}
	}

	if o.IgroupUUID != nil {

		// query param igroup.uuid
		var qrIgroupUUID string

		if o.IgroupUUID != nil {
			qrIgroupUUID = *o.IgroupUUID
		}
		qIgroupUUID := qrIgroupUUID
		if qIgroupUUID != "" {

			if err := r.SetQueryParam("igroup.uuid", qIgroupUUID); err != nil {
				return err
			}
		}
	}

	if o.LogicalUnitNumber != nil {

		// query param logical_unit_number
		var qrLogicalUnitNumber int64

		if o.LogicalUnitNumber != nil {
			qrLogicalUnitNumber = *o.LogicalUnitNumber
		}
		qLogicalUnitNumber := swag.FormatInt64(qrLogicalUnitNumber)
		if qLogicalUnitNumber != "" {

			if err := r.SetQueryParam("logical_unit_number", qLogicalUnitNumber); err != nil {
				return err
			}
		}
	}

	if o.LunName != nil {

		// query param lun.name
		var qrLunName string

		if o.LunName != nil {
			qrLunName = *o.LunName
		}
		qLunName := qrLunName
		if qLunName != "" {

			if err := r.SetQueryParam("lun.name", qLunName); err != nil {
				return err
			}
		}
	}

	if o.LunNodeName != nil {

		// query param lun.node.name
		var qrLunNodeName string

		if o.LunNodeName != nil {
			qrLunNodeName = *o.LunNodeName
		}
		qLunNodeName := qrLunNodeName
		if qLunNodeName != "" {

			if err := r.SetQueryParam("lun.node.name", qLunNodeName); err != nil {
				return err
			}
		}
	}

	if o.LunNodeUUID != nil {

		// query param lun.node.uuid
		var qrLunNodeUUID string

		if o.LunNodeUUID != nil {
			qrLunNodeUUID = *o.LunNodeUUID
		}
		qLunNodeUUID := qrLunNodeUUID
		if qLunNodeUUID != "" {

			if err := r.SetQueryParam("lun.node.uuid", qLunNodeUUID); err != nil {
				return err
			}
		}
	}

	if o.LunSmbcReplicated != nil {

		// query param lun.smbc.replicated
		var qrLunSmbcReplicated bool

		if o.LunSmbcReplicated != nil {
			qrLunSmbcReplicated = *o.LunSmbcReplicated
		}
		qLunSmbcReplicated := swag.FormatBool(qrLunSmbcReplicated)
		if qLunSmbcReplicated != "" {

			if err := r.SetQueryParam("lun.smbc.replicated", qLunSmbcReplicated); err != nil {
				return err
			}
		}
	}

	if o.LunUUID != nil {

		// query param lun.uuid
		var qrLunUUID string

		if o.LunUUID != nil {
			qrLunUUID = *o.LunUUID
		}
		qLunUUID := qrLunUUID
		if qLunUUID != "" {

			if err := r.SetQueryParam("lun.uuid", qLunUUID); err != nil {
				return err
			}
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

	if o.OrderBy != nil {

		// binding items for order_by
		joinedOrderBy := o.bindParamOrderBy(reg)

		// query array param order_by
		if err := r.SetQueryParam("order_by", joinedOrderBy...); err != nil {
			return err
		}
	}

	if o.ReportingNodesName != nil {

		// query param reporting_nodes.name
		var qrReportingNodesName string

		if o.ReportingNodesName != nil {
			qrReportingNodesName = *o.ReportingNodesName
		}
		qReportingNodesName := qrReportingNodesName
		if qReportingNodesName != "" {

			if err := r.SetQueryParam("reporting_nodes.name", qReportingNodesName); err != nil {
				return err
			}
		}
	}

	if o.ReportingNodesUUID != nil {

		// query param reporting_nodes.uuid
		var qrReportingNodesUUID string

		if o.ReportingNodesUUID != nil {
			qrReportingNodesUUID = *o.ReportingNodesUUID
		}
		qReportingNodesUUID := qrReportingNodesUUID
		if qReportingNodesUUID != "" {

			if err := r.SetQueryParam("reporting_nodes.uuid", qReportingNodesUUID); err != nil {
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

// bindParamLunMapCollectionGet binds the parameter fields
func (o *LunMapCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamLunMapCollectionGet binds the parameter order_by
func (o *LunMapCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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
