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

// NewDuoCollectionGetParams creates a new DuoCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDuoCollectionGetParams() *DuoCollectionGetParams {
	return &DuoCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDuoCollectionGetParamsWithTimeout creates a new DuoCollectionGetParams object
// with the ability to set a timeout on a request.
func NewDuoCollectionGetParamsWithTimeout(timeout time.Duration) *DuoCollectionGetParams {
	return &DuoCollectionGetParams{
		timeout: timeout,
	}
}

// NewDuoCollectionGetParamsWithContext creates a new DuoCollectionGetParams object
// with the ability to set a context for a request.
func NewDuoCollectionGetParamsWithContext(ctx context.Context) *DuoCollectionGetParams {
	return &DuoCollectionGetParams{
		Context: ctx,
	}
}

// NewDuoCollectionGetParamsWithHTTPClient creates a new DuoCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewDuoCollectionGetParamsWithHTTPClient(client *http.Client) *DuoCollectionGetParams {
	return &DuoCollectionGetParams{
		HTTPClient: client,
	}
}

/*
DuoCollectionGetParams contains all the parameters to send to the API endpoint

	for the duo collection get operation.

	Typically these are written to a http.Request.
*/
type DuoCollectionGetParams struct {

	/* APIHost.

	   Filter by api_host
	*/
	APIHost *string

	/* AutoPush.

	   Filter by auto_push
	*/
	AutoPush *bool

	/* Comment.

	   Filter by comment
	*/
	Comment *string

	/* FailMode.

	   Filter by fail_mode
	*/
	FailMode *string

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* Fingerprint.

	   Filter by fingerprint
	*/
	Fingerprint *string

	/* HTTPProxy.

	   Filter by http_proxy
	*/
	HTTPProxy *string

	/* IntegrationKey.

	   Filter by integration_key
	*/
	IntegrationKey *string

	/* IsEnabled.

	   Filter by is_enabled
	*/
	IsEnabled *bool

	/* MaxPrompts.

	   Filter by max_prompts
	*/
	MaxPrompts *int64

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecords *int64

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderBy []string

	/* OwnerName.

	   Filter by owner.name
	*/
	OwnerName *string

	/* OwnerUUID.

	   Filter by owner.uuid
	*/
	OwnerUUID *string

	/* PushInfo.

	   Filter by push_info
	*/
	PushInfo *bool

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

	/* Status.

	   Filter by status
	*/
	Status *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the duo collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DuoCollectionGetParams) WithDefaults() *DuoCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the duo collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DuoCollectionGetParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)
	)

	val := DuoCollectionGetParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the duo collection get params
func (o *DuoCollectionGetParams) WithTimeout(timeout time.Duration) *DuoCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the duo collection get params
func (o *DuoCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the duo collection get params
func (o *DuoCollectionGetParams) WithContext(ctx context.Context) *DuoCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the duo collection get params
func (o *DuoCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the duo collection get params
func (o *DuoCollectionGetParams) WithHTTPClient(client *http.Client) *DuoCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the duo collection get params
func (o *DuoCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIHost adds the aPIHost to the duo collection get params
func (o *DuoCollectionGetParams) WithAPIHost(aPIHost *string) *DuoCollectionGetParams {
	o.SetAPIHost(aPIHost)
	return o
}

// SetAPIHost adds the apiHost to the duo collection get params
func (o *DuoCollectionGetParams) SetAPIHost(aPIHost *string) {
	o.APIHost = aPIHost
}

// WithAutoPush adds the autoPush to the duo collection get params
func (o *DuoCollectionGetParams) WithAutoPush(autoPush *bool) *DuoCollectionGetParams {
	o.SetAutoPush(autoPush)
	return o
}

// SetAutoPush adds the autoPush to the duo collection get params
func (o *DuoCollectionGetParams) SetAutoPush(autoPush *bool) {
	o.AutoPush = autoPush
}

// WithComment adds the comment to the duo collection get params
func (o *DuoCollectionGetParams) WithComment(comment *string) *DuoCollectionGetParams {
	o.SetComment(comment)
	return o
}

// SetComment adds the comment to the duo collection get params
func (o *DuoCollectionGetParams) SetComment(comment *string) {
	o.Comment = comment
}

// WithFailMode adds the failMode to the duo collection get params
func (o *DuoCollectionGetParams) WithFailMode(failMode *string) *DuoCollectionGetParams {
	o.SetFailMode(failMode)
	return o
}

// SetFailMode adds the failMode to the duo collection get params
func (o *DuoCollectionGetParams) SetFailMode(failMode *string) {
	o.FailMode = failMode
}

// WithFields adds the fields to the duo collection get params
func (o *DuoCollectionGetParams) WithFields(fields []string) *DuoCollectionGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the duo collection get params
func (o *DuoCollectionGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithFingerprint adds the fingerprint to the duo collection get params
func (o *DuoCollectionGetParams) WithFingerprint(fingerprint *string) *DuoCollectionGetParams {
	o.SetFingerprint(fingerprint)
	return o
}

// SetFingerprint adds the fingerprint to the duo collection get params
func (o *DuoCollectionGetParams) SetFingerprint(fingerprint *string) {
	o.Fingerprint = fingerprint
}

// WithHTTPProxy adds the hTTPProxy to the duo collection get params
func (o *DuoCollectionGetParams) WithHTTPProxy(hTTPProxy *string) *DuoCollectionGetParams {
	o.SetHTTPProxy(hTTPProxy)
	return o
}

// SetHTTPProxy adds the httpProxy to the duo collection get params
func (o *DuoCollectionGetParams) SetHTTPProxy(hTTPProxy *string) {
	o.HTTPProxy = hTTPProxy
}

// WithIntegrationKey adds the integrationKey to the duo collection get params
func (o *DuoCollectionGetParams) WithIntegrationKey(integrationKey *string) *DuoCollectionGetParams {
	o.SetIntegrationKey(integrationKey)
	return o
}

// SetIntegrationKey adds the integrationKey to the duo collection get params
func (o *DuoCollectionGetParams) SetIntegrationKey(integrationKey *string) {
	o.IntegrationKey = integrationKey
}

// WithIsEnabled adds the isEnabled to the duo collection get params
func (o *DuoCollectionGetParams) WithIsEnabled(isEnabled *bool) *DuoCollectionGetParams {
	o.SetIsEnabled(isEnabled)
	return o
}

// SetIsEnabled adds the isEnabled to the duo collection get params
func (o *DuoCollectionGetParams) SetIsEnabled(isEnabled *bool) {
	o.IsEnabled = isEnabled
}

// WithMaxPrompts adds the maxPrompts to the duo collection get params
func (o *DuoCollectionGetParams) WithMaxPrompts(maxPrompts *int64) *DuoCollectionGetParams {
	o.SetMaxPrompts(maxPrompts)
	return o
}

// SetMaxPrompts adds the maxPrompts to the duo collection get params
func (o *DuoCollectionGetParams) SetMaxPrompts(maxPrompts *int64) {
	o.MaxPrompts = maxPrompts
}

// WithMaxRecords adds the maxRecords to the duo collection get params
func (o *DuoCollectionGetParams) WithMaxRecords(maxRecords *int64) *DuoCollectionGetParams {
	o.SetMaxRecords(maxRecords)
	return o
}

// SetMaxRecords adds the maxRecords to the duo collection get params
func (o *DuoCollectionGetParams) SetMaxRecords(maxRecords *int64) {
	o.MaxRecords = maxRecords
}

// WithOrderBy adds the orderBy to the duo collection get params
func (o *DuoCollectionGetParams) WithOrderBy(orderBy []string) *DuoCollectionGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the duo collection get params
func (o *DuoCollectionGetParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithOwnerName adds the ownerName to the duo collection get params
func (o *DuoCollectionGetParams) WithOwnerName(ownerName *string) *DuoCollectionGetParams {
	o.SetOwnerName(ownerName)
	return o
}

// SetOwnerName adds the ownerName to the duo collection get params
func (o *DuoCollectionGetParams) SetOwnerName(ownerName *string) {
	o.OwnerName = ownerName
}

// WithOwnerUUID adds the ownerUUID to the duo collection get params
func (o *DuoCollectionGetParams) WithOwnerUUID(ownerUUID *string) *DuoCollectionGetParams {
	o.SetOwnerUUID(ownerUUID)
	return o
}

// SetOwnerUUID adds the ownerUuid to the duo collection get params
func (o *DuoCollectionGetParams) SetOwnerUUID(ownerUUID *string) {
	o.OwnerUUID = ownerUUID
}

// WithPushInfo adds the pushInfo to the duo collection get params
func (o *DuoCollectionGetParams) WithPushInfo(pushInfo *bool) *DuoCollectionGetParams {
	o.SetPushInfo(pushInfo)
	return o
}

// SetPushInfo adds the pushInfo to the duo collection get params
func (o *DuoCollectionGetParams) SetPushInfo(pushInfo *bool) {
	o.PushInfo = pushInfo
}

// WithReturnRecords adds the returnRecords to the duo collection get params
func (o *DuoCollectionGetParams) WithReturnRecords(returnRecords *bool) *DuoCollectionGetParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the duo collection get params
func (o *DuoCollectionGetParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the duo collection get params
func (o *DuoCollectionGetParams) WithReturnTimeout(returnTimeout *int64) *DuoCollectionGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the duo collection get params
func (o *DuoCollectionGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithStatus adds the status to the duo collection get params
func (o *DuoCollectionGetParams) WithStatus(status *string) *DuoCollectionGetParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the duo collection get params
func (o *DuoCollectionGetParams) SetStatus(status *string) {
	o.Status = status
}

// WriteToRequest writes these params to a swagger request
func (o *DuoCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.APIHost != nil {

		// query param api_host
		var qrAPIHost string

		if o.APIHost != nil {
			qrAPIHost = *o.APIHost
		}
		qAPIHost := qrAPIHost
		if qAPIHost != "" {

			if err := r.SetQueryParam("api_host", qAPIHost); err != nil {
				return err
			}
		}
	}

	if o.AutoPush != nil {

		// query param auto_push
		var qrAutoPush bool

		if o.AutoPush != nil {
			qrAutoPush = *o.AutoPush
		}
		qAutoPush := swag.FormatBool(qrAutoPush)
		if qAutoPush != "" {

			if err := r.SetQueryParam("auto_push", qAutoPush); err != nil {
				return err
			}
		}
	}

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

	if o.FailMode != nil {

		// query param fail_mode
		var qrFailMode string

		if o.FailMode != nil {
			qrFailMode = *o.FailMode
		}
		qFailMode := qrFailMode
		if qFailMode != "" {

			if err := r.SetQueryParam("fail_mode", qFailMode); err != nil {
				return err
			}
		}
	}

	if o.Fields != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	if o.Fingerprint != nil {

		// query param fingerprint
		var qrFingerprint string

		if o.Fingerprint != nil {
			qrFingerprint = *o.Fingerprint
		}
		qFingerprint := qrFingerprint
		if qFingerprint != "" {

			if err := r.SetQueryParam("fingerprint", qFingerprint); err != nil {
				return err
			}
		}
	}

	if o.HTTPProxy != nil {

		// query param http_proxy
		var qrHTTPProxy string

		if o.HTTPProxy != nil {
			qrHTTPProxy = *o.HTTPProxy
		}
		qHTTPProxy := qrHTTPProxy
		if qHTTPProxy != "" {

			if err := r.SetQueryParam("http_proxy", qHTTPProxy); err != nil {
				return err
			}
		}
	}

	if o.IntegrationKey != nil {

		// query param integration_key
		var qrIntegrationKey string

		if o.IntegrationKey != nil {
			qrIntegrationKey = *o.IntegrationKey
		}
		qIntegrationKey := qrIntegrationKey
		if qIntegrationKey != "" {

			if err := r.SetQueryParam("integration_key", qIntegrationKey); err != nil {
				return err
			}
		}
	}

	if o.IsEnabled != nil {

		// query param is_enabled
		var qrIsEnabled bool

		if o.IsEnabled != nil {
			qrIsEnabled = *o.IsEnabled
		}
		qIsEnabled := swag.FormatBool(qrIsEnabled)
		if qIsEnabled != "" {

			if err := r.SetQueryParam("is_enabled", qIsEnabled); err != nil {
				return err
			}
		}
	}

	if o.MaxPrompts != nil {

		// query param max_prompts
		var qrMaxPrompts int64

		if o.MaxPrompts != nil {
			qrMaxPrompts = *o.MaxPrompts
		}
		qMaxPrompts := swag.FormatInt64(qrMaxPrompts)
		if qMaxPrompts != "" {

			if err := r.SetQueryParam("max_prompts", qMaxPrompts); err != nil {
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

	if o.OwnerName != nil {

		// query param owner.name
		var qrOwnerName string

		if o.OwnerName != nil {
			qrOwnerName = *o.OwnerName
		}
		qOwnerName := qrOwnerName
		if qOwnerName != "" {

			if err := r.SetQueryParam("owner.name", qOwnerName); err != nil {
				return err
			}
		}
	}

	if o.OwnerUUID != nil {

		// query param owner.uuid
		var qrOwnerUUID string

		if o.OwnerUUID != nil {
			qrOwnerUUID = *o.OwnerUUID
		}
		qOwnerUUID := qrOwnerUUID
		if qOwnerUUID != "" {

			if err := r.SetQueryParam("owner.uuid", qOwnerUUID); err != nil {
				return err
			}
		}
	}

	if o.PushInfo != nil {

		// query param push_info
		var qrPushInfo bool

		if o.PushInfo != nil {
			qrPushInfo = *o.PushInfo
		}
		qPushInfo := swag.FormatBool(qrPushInfo)
		if qPushInfo != "" {

			if err := r.SetQueryParam("push_info", qPushInfo); err != nil {
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

	if o.Status != nil {

		// query param status
		var qrStatus string

		if o.Status != nil {
			qrStatus = *o.Status
		}
		qStatus := qrStatus
		if qStatus != "" {

			if err := r.SetQueryParam("status", qStatus); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamDuoCollectionGet binds the parameter fields
func (o *DuoCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamDuoCollectionGet binds the parameter order_by
func (o *DuoCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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