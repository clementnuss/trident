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

// NewUnixGroupSettingsGetParams creates a new UnixGroupSettingsGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUnixGroupSettingsGetParams() *UnixGroupSettingsGetParams {
	return &UnixGroupSettingsGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUnixGroupSettingsGetParamsWithTimeout creates a new UnixGroupSettingsGetParams object
// with the ability to set a timeout on a request.
func NewUnixGroupSettingsGetParamsWithTimeout(timeout time.Duration) *UnixGroupSettingsGetParams {
	return &UnixGroupSettingsGetParams{
		timeout: timeout,
	}
}

// NewUnixGroupSettingsGetParamsWithContext creates a new UnixGroupSettingsGetParams object
// with the ability to set a context for a request.
func NewUnixGroupSettingsGetParamsWithContext(ctx context.Context) *UnixGroupSettingsGetParams {
	return &UnixGroupSettingsGetParams{
		Context: ctx,
	}
}

// NewUnixGroupSettingsGetParamsWithHTTPClient creates a new UnixGroupSettingsGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewUnixGroupSettingsGetParamsWithHTTPClient(client *http.Client) *UnixGroupSettingsGetParams {
	return &UnixGroupSettingsGetParams{
		HTTPClient: client,
	}
}

/*
UnixGroupSettingsGetParams contains all the parameters to send to the API endpoint

	for the unix group settings get operation.

	Typically these are written to a http.Request.
*/
type UnixGroupSettingsGetParams struct {

	/* Enabled.

	   Filter by enabled
	*/
	EnabledQueryParameter *bool

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecordsQueryParameter *int64

	/* NegativeCacheEnabled.

	   Filter by negative_cache_enabled
	*/
	NegativeCacheEnabledQueryParameter *bool

	/* NegativeTTL.

	   Filter by negative_ttl
	*/
	NegativeTTLQueryParameter *string

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderByQueryParameter []string

	/* PropagationEnabled.

	   Filter by propagation_enabled
	*/
	PropagationEnabledQueryParameter *bool

	/* ReturnRecords.

	   The default is true for GET calls.  When set to false, only the number of records is returned.

	   Default: true
	*/
	ReturnRecordsQueryParameter *bool

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached.

	   Default: 15
	*/
	ReturnTimeoutQueryParameter *int64

	/* SvmName.

	   Filter by svm.name
	*/
	SVMNameQueryParameter *string

	/* SvmUUID.

	   SVM UUID.
	*/
	SVMUUIDPathParameter string

	/* TTL.

	   Filter by ttl
	*/
	TTLQueryParameter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the unix group settings get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UnixGroupSettingsGetParams) WithDefaults() *UnixGroupSettingsGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the unix group settings get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UnixGroupSettingsGetParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(true)

		returnTimeoutQueryParameterDefault = int64(15)
	)

	val := UnixGroupSettingsGetParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the unix group settings get params
func (o *UnixGroupSettingsGetParams) WithTimeout(timeout time.Duration) *UnixGroupSettingsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the unix group settings get params
func (o *UnixGroupSettingsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the unix group settings get params
func (o *UnixGroupSettingsGetParams) WithContext(ctx context.Context) *UnixGroupSettingsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the unix group settings get params
func (o *UnixGroupSettingsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the unix group settings get params
func (o *UnixGroupSettingsGetParams) WithHTTPClient(client *http.Client) *UnixGroupSettingsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the unix group settings get params
func (o *UnixGroupSettingsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnabledQueryParameter adds the enabled to the unix group settings get params
func (o *UnixGroupSettingsGetParams) WithEnabledQueryParameter(enabled *bool) *UnixGroupSettingsGetParams {
	o.SetEnabledQueryParameter(enabled)
	return o
}

// SetEnabledQueryParameter adds the enabled to the unix group settings get params
func (o *UnixGroupSettingsGetParams) SetEnabledQueryParameter(enabled *bool) {
	o.EnabledQueryParameter = enabled
}

// WithFieldsQueryParameter adds the fields to the unix group settings get params
func (o *UnixGroupSettingsGetParams) WithFieldsQueryParameter(fields []string) *UnixGroupSettingsGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the unix group settings get params
func (o *UnixGroupSettingsGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithMaxRecordsQueryParameter adds the maxRecords to the unix group settings get params
func (o *UnixGroupSettingsGetParams) WithMaxRecordsQueryParameter(maxRecords *int64) *UnixGroupSettingsGetParams {
	o.SetMaxRecordsQueryParameter(maxRecords)
	return o
}

// SetMaxRecordsQueryParameter adds the maxRecords to the unix group settings get params
func (o *UnixGroupSettingsGetParams) SetMaxRecordsQueryParameter(maxRecords *int64) {
	o.MaxRecordsQueryParameter = maxRecords
}

// WithNegativeCacheEnabledQueryParameter adds the negativeCacheEnabled to the unix group settings get params
func (o *UnixGroupSettingsGetParams) WithNegativeCacheEnabledQueryParameter(negativeCacheEnabled *bool) *UnixGroupSettingsGetParams {
	o.SetNegativeCacheEnabledQueryParameter(negativeCacheEnabled)
	return o
}

// SetNegativeCacheEnabledQueryParameter adds the negativeCacheEnabled to the unix group settings get params
func (o *UnixGroupSettingsGetParams) SetNegativeCacheEnabledQueryParameter(negativeCacheEnabled *bool) {
	o.NegativeCacheEnabledQueryParameter = negativeCacheEnabled
}

// WithNegativeTTLQueryParameter adds the negativeTTL to the unix group settings get params
func (o *UnixGroupSettingsGetParams) WithNegativeTTLQueryParameter(negativeTTL *string) *UnixGroupSettingsGetParams {
	o.SetNegativeTTLQueryParameter(negativeTTL)
	return o
}

// SetNegativeTTLQueryParameter adds the negativeTtl to the unix group settings get params
func (o *UnixGroupSettingsGetParams) SetNegativeTTLQueryParameter(negativeTTL *string) {
	o.NegativeTTLQueryParameter = negativeTTL
}

// WithOrderByQueryParameter adds the orderBy to the unix group settings get params
func (o *UnixGroupSettingsGetParams) WithOrderByQueryParameter(orderBy []string) *UnixGroupSettingsGetParams {
	o.SetOrderByQueryParameter(orderBy)
	return o
}

// SetOrderByQueryParameter adds the orderBy to the unix group settings get params
func (o *UnixGroupSettingsGetParams) SetOrderByQueryParameter(orderBy []string) {
	o.OrderByQueryParameter = orderBy
}

// WithPropagationEnabledQueryParameter adds the propagationEnabled to the unix group settings get params
func (o *UnixGroupSettingsGetParams) WithPropagationEnabledQueryParameter(propagationEnabled *bool) *UnixGroupSettingsGetParams {
	o.SetPropagationEnabledQueryParameter(propagationEnabled)
	return o
}

// SetPropagationEnabledQueryParameter adds the propagationEnabled to the unix group settings get params
func (o *UnixGroupSettingsGetParams) SetPropagationEnabledQueryParameter(propagationEnabled *bool) {
	o.PropagationEnabledQueryParameter = propagationEnabled
}

// WithReturnRecordsQueryParameter adds the returnRecords to the unix group settings get params
func (o *UnixGroupSettingsGetParams) WithReturnRecordsQueryParameter(returnRecords *bool) *UnixGroupSettingsGetParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the unix group settings get params
func (o *UnixGroupSettingsGetParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the unix group settings get params
func (o *UnixGroupSettingsGetParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *UnixGroupSettingsGetParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the unix group settings get params
func (o *UnixGroupSettingsGetParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WithSVMNameQueryParameter adds the svmName to the unix group settings get params
func (o *UnixGroupSettingsGetParams) WithSVMNameQueryParameter(svmName *string) *UnixGroupSettingsGetParams {
	o.SetSVMNameQueryParameter(svmName)
	return o
}

// SetSVMNameQueryParameter adds the svmName to the unix group settings get params
func (o *UnixGroupSettingsGetParams) SetSVMNameQueryParameter(svmName *string) {
	o.SVMNameQueryParameter = svmName
}

// WithSVMUUIDPathParameter adds the svmUUID to the unix group settings get params
func (o *UnixGroupSettingsGetParams) WithSVMUUIDPathParameter(svmUUID string) *UnixGroupSettingsGetParams {
	o.SetSVMUUIDPathParameter(svmUUID)
	return o
}

// SetSVMUUIDPathParameter adds the svmUuid to the unix group settings get params
func (o *UnixGroupSettingsGetParams) SetSVMUUIDPathParameter(svmUUID string) {
	o.SVMUUIDPathParameter = svmUUID
}

// WithTTLQueryParameter adds the ttl to the unix group settings get params
func (o *UnixGroupSettingsGetParams) WithTTLQueryParameter(ttl *string) *UnixGroupSettingsGetParams {
	o.SetTTLQueryParameter(ttl)
	return o
}

// SetTTLQueryParameter adds the ttl to the unix group settings get params
func (o *UnixGroupSettingsGetParams) SetTTLQueryParameter(ttl *string) {
	o.TTLQueryParameter = ttl
}

// WriteToRequest writes these params to a swagger request
func (o *UnixGroupSettingsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.EnabledQueryParameter != nil {

		// query param enabled
		var qrEnabled bool

		if o.EnabledQueryParameter != nil {
			qrEnabled = *o.EnabledQueryParameter
		}
		qEnabled := swag.FormatBool(qrEnabled)
		if qEnabled != "" {

			if err := r.SetQueryParam("enabled", qEnabled); err != nil {
				return err
			}
		}
	}

	if o.FieldsQueryParameter != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	if o.MaxRecordsQueryParameter != nil {

		// query param max_records
		var qrMaxRecords int64

		if o.MaxRecordsQueryParameter != nil {
			qrMaxRecords = *o.MaxRecordsQueryParameter
		}
		qMaxRecords := swag.FormatInt64(qrMaxRecords)
		if qMaxRecords != "" {

			if err := r.SetQueryParam("max_records", qMaxRecords); err != nil {
				return err
			}
		}
	}

	if o.NegativeCacheEnabledQueryParameter != nil {

		// query param negative_cache_enabled
		var qrNegativeCacheEnabled bool

		if o.NegativeCacheEnabledQueryParameter != nil {
			qrNegativeCacheEnabled = *o.NegativeCacheEnabledQueryParameter
		}
		qNegativeCacheEnabled := swag.FormatBool(qrNegativeCacheEnabled)
		if qNegativeCacheEnabled != "" {

			if err := r.SetQueryParam("negative_cache_enabled", qNegativeCacheEnabled); err != nil {
				return err
			}
		}
	}

	if o.NegativeTTLQueryParameter != nil {

		// query param negative_ttl
		var qrNegativeTTL string

		if o.NegativeTTLQueryParameter != nil {
			qrNegativeTTL = *o.NegativeTTLQueryParameter
		}
		qNegativeTTL := qrNegativeTTL
		if qNegativeTTL != "" {

			if err := r.SetQueryParam("negative_ttl", qNegativeTTL); err != nil {
				return err
			}
		}
	}

	if o.OrderByQueryParameter != nil {

		// binding items for order_by
		joinedOrderBy := o.bindParamOrderBy(reg)

		// query array param order_by
		if err := r.SetQueryParam("order_by", joinedOrderBy...); err != nil {
			return err
		}
	}

	if o.PropagationEnabledQueryParameter != nil {

		// query param propagation_enabled
		var qrPropagationEnabled bool

		if o.PropagationEnabledQueryParameter != nil {
			qrPropagationEnabled = *o.PropagationEnabledQueryParameter
		}
		qPropagationEnabled := swag.FormatBool(qrPropagationEnabled)
		if qPropagationEnabled != "" {

			if err := r.SetQueryParam("propagation_enabled", qPropagationEnabled); err != nil {
				return err
			}
		}
	}

	if o.ReturnRecordsQueryParameter != nil {

		// query param return_records
		var qrReturnRecords bool

		if o.ReturnRecordsQueryParameter != nil {
			qrReturnRecords = *o.ReturnRecordsQueryParameter
		}
		qReturnRecords := swag.FormatBool(qrReturnRecords)
		if qReturnRecords != "" {

			if err := r.SetQueryParam("return_records", qReturnRecords); err != nil {
				return err
			}
		}
	}

	if o.ReturnTimeoutQueryParameter != nil {

		// query param return_timeout
		var qrReturnTimeout int64

		if o.ReturnTimeoutQueryParameter != nil {
			qrReturnTimeout = *o.ReturnTimeoutQueryParameter
		}
		qReturnTimeout := swag.FormatInt64(qrReturnTimeout)
		if qReturnTimeout != "" {

			if err := r.SetQueryParam("return_timeout", qReturnTimeout); err != nil {
				return err
			}
		}
	}

	if o.SVMNameQueryParameter != nil {

		// query param svm.name
		var qrSvmName string

		if o.SVMNameQueryParameter != nil {
			qrSvmName = *o.SVMNameQueryParameter
		}
		qSvmName := qrSvmName
		if qSvmName != "" {

			if err := r.SetQueryParam("svm.name", qSvmName); err != nil {
				return err
			}
		}
	}

	// path param svm.uuid
	if err := r.SetPathParam("svm.uuid", o.SVMUUIDPathParameter); err != nil {
		return err
	}

	if o.TTLQueryParameter != nil {

		// query param ttl
		var qrTTL string

		if o.TTLQueryParameter != nil {
			qrTTL = *o.TTLQueryParameter
		}
		qTTL := qrTTL
		if qTTL != "" {

			if err := r.SetQueryParam("ttl", qTTL); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamUnixGroupSettingsGet binds the parameter fields
func (o *UnixGroupSettingsGetParams) bindParamFields(formats strfmt.Registry) []string {
	fieldsIR := o.FieldsQueryParameter

	var fieldsIC []string
	for _, fieldsIIR := range fieldsIR { // explode []string

		fieldsIIV := fieldsIIR // string as string
		fieldsIC = append(fieldsIC, fieldsIIV)
	}

	// items.CollectionFormat: "csv"
	fieldsIS := swag.JoinByFormat(fieldsIC, "csv")

	return fieldsIS
}

// bindParamUnixGroupSettingsGet binds the parameter order_by
func (o *UnixGroupSettingsGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
	orderByIR := o.OrderByQueryParameter

	var orderByIC []string
	for _, orderByIIR := range orderByIR { // explode []string

		orderByIIV := orderByIIR // string as string
		orderByIC = append(orderByIC, orderByIIV)
	}

	// items.CollectionFormat: "csv"
	orderByIS := swag.JoinByFormat(orderByIC, "csv")

	return orderByIS
}