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

// NewExportPolicyModifyCollectionParams creates a new ExportPolicyModifyCollectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExportPolicyModifyCollectionParams() *ExportPolicyModifyCollectionParams {
	return &ExportPolicyModifyCollectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExportPolicyModifyCollectionParamsWithTimeout creates a new ExportPolicyModifyCollectionParams object
// with the ability to set a timeout on a request.
func NewExportPolicyModifyCollectionParamsWithTimeout(timeout time.Duration) *ExportPolicyModifyCollectionParams {
	return &ExportPolicyModifyCollectionParams{
		timeout: timeout,
	}
}

// NewExportPolicyModifyCollectionParamsWithContext creates a new ExportPolicyModifyCollectionParams object
// with the ability to set a context for a request.
func NewExportPolicyModifyCollectionParamsWithContext(ctx context.Context) *ExportPolicyModifyCollectionParams {
	return &ExportPolicyModifyCollectionParams{
		Context: ctx,
	}
}

// NewExportPolicyModifyCollectionParamsWithHTTPClient creates a new ExportPolicyModifyCollectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewExportPolicyModifyCollectionParamsWithHTTPClient(client *http.Client) *ExportPolicyModifyCollectionParams {
	return &ExportPolicyModifyCollectionParams{
		HTTPClient: client,
	}
}

/*
ExportPolicyModifyCollectionParams contains all the parameters to send to the API endpoint

	for the export policy modify collection operation.

	Typically these are written to a http.Request.
*/
type ExportPolicyModifyCollectionParams struct {

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
	Info ExportPolicyModifyCollectionBody

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

	/* RulesAllowDeviceCreation.

	   Filter by rules.allow_device_creation
	*/
	RulesAllowDeviceCreation *bool

	/* RulesAllowSuid.

	   Filter by rules.allow_suid
	*/
	RulesAllowSuid *bool

	/* RulesAnonymousUser.

	   Filter by rules.anonymous_user
	*/
	RulesAnonymousUser *string

	/* RulesChownMode.

	   Filter by rules.chown_mode
	*/
	RulesChownMode *string

	/* RulesClientsMatch.

	   Filter by rules.clients.match
	*/
	RulesClientsMatch *string

	/* RulesIndex.

	   Filter by rules.index
	*/
	RulesIndex *int64

	/* RulesNtfsUnixSecurity.

	   Filter by rules.ntfs_unix_security
	*/
	RulesNtfsUnixSecurity *string

	/* RulesProtocols.

	   Filter by rules.protocols
	*/
	RulesProtocols *string

	/* RulesRoRule.

	   Filter by rules.ro_rule
	*/
	RulesRoRule *string

	/* RulesRwRule.

	   Filter by rules.rw_rule
	*/
	RulesRwRule *string

	/* RulesSuperuser.

	   Filter by rules.superuser
	*/
	RulesSuperuser *string

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

// WithDefaults hydrates default values in the export policy modify collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExportPolicyModifyCollectionParams) WithDefaults() *ExportPolicyModifyCollectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the export policy modify collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExportPolicyModifyCollectionParams) SetDefaults() {
	var (
		continueOnFailureDefault = bool(false)

		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)

		serialRecordsDefault = bool(false)
	)

	val := ExportPolicyModifyCollectionParams{
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

// WithTimeout adds the timeout to the export policy modify collection params
func (o *ExportPolicyModifyCollectionParams) WithTimeout(timeout time.Duration) *ExportPolicyModifyCollectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the export policy modify collection params
func (o *ExportPolicyModifyCollectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the export policy modify collection params
func (o *ExportPolicyModifyCollectionParams) WithContext(ctx context.Context) *ExportPolicyModifyCollectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the export policy modify collection params
func (o *ExportPolicyModifyCollectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the export policy modify collection params
func (o *ExportPolicyModifyCollectionParams) WithHTTPClient(client *http.Client) *ExportPolicyModifyCollectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the export policy modify collection params
func (o *ExportPolicyModifyCollectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContinueOnFailure adds the continueOnFailure to the export policy modify collection params
func (o *ExportPolicyModifyCollectionParams) WithContinueOnFailure(continueOnFailure *bool) *ExportPolicyModifyCollectionParams {
	o.SetContinueOnFailure(continueOnFailure)
	return o
}

// SetContinueOnFailure adds the continueOnFailure to the export policy modify collection params
func (o *ExportPolicyModifyCollectionParams) SetContinueOnFailure(continueOnFailure *bool) {
	o.ContinueOnFailure = continueOnFailure
}

// WithID adds the id to the export policy modify collection params
func (o *ExportPolicyModifyCollectionParams) WithID(id *int64) *ExportPolicyModifyCollectionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the export policy modify collection params
func (o *ExportPolicyModifyCollectionParams) SetID(id *int64) {
	o.ID = id
}

// WithInfo adds the info to the export policy modify collection params
func (o *ExportPolicyModifyCollectionParams) WithInfo(info ExportPolicyModifyCollectionBody) *ExportPolicyModifyCollectionParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the export policy modify collection params
func (o *ExportPolicyModifyCollectionParams) SetInfo(info ExportPolicyModifyCollectionBody) {
	o.Info = info
}

// WithName adds the name to the export policy modify collection params
func (o *ExportPolicyModifyCollectionParams) WithName(name *string) *ExportPolicyModifyCollectionParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the export policy modify collection params
func (o *ExportPolicyModifyCollectionParams) SetName(name *string) {
	o.Name = name
}

// WithReturnRecords adds the returnRecords to the export policy modify collection params
func (o *ExportPolicyModifyCollectionParams) WithReturnRecords(returnRecords *bool) *ExportPolicyModifyCollectionParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the export policy modify collection params
func (o *ExportPolicyModifyCollectionParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the export policy modify collection params
func (o *ExportPolicyModifyCollectionParams) WithReturnTimeout(returnTimeout *int64) *ExportPolicyModifyCollectionParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the export policy modify collection params
func (o *ExportPolicyModifyCollectionParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithRulesAllowDeviceCreation adds the rulesAllowDeviceCreation to the export policy modify collection params
func (o *ExportPolicyModifyCollectionParams) WithRulesAllowDeviceCreation(rulesAllowDeviceCreation *bool) *ExportPolicyModifyCollectionParams {
	o.SetRulesAllowDeviceCreation(rulesAllowDeviceCreation)
	return o
}

// SetRulesAllowDeviceCreation adds the rulesAllowDeviceCreation to the export policy modify collection params
func (o *ExportPolicyModifyCollectionParams) SetRulesAllowDeviceCreation(rulesAllowDeviceCreation *bool) {
	o.RulesAllowDeviceCreation = rulesAllowDeviceCreation
}

// WithRulesAllowSuid adds the rulesAllowSuid to the export policy modify collection params
func (o *ExportPolicyModifyCollectionParams) WithRulesAllowSuid(rulesAllowSuid *bool) *ExportPolicyModifyCollectionParams {
	o.SetRulesAllowSuid(rulesAllowSuid)
	return o
}

// SetRulesAllowSuid adds the rulesAllowSuid to the export policy modify collection params
func (o *ExportPolicyModifyCollectionParams) SetRulesAllowSuid(rulesAllowSuid *bool) {
	o.RulesAllowSuid = rulesAllowSuid
}

// WithRulesAnonymousUser adds the rulesAnonymousUser to the export policy modify collection params
func (o *ExportPolicyModifyCollectionParams) WithRulesAnonymousUser(rulesAnonymousUser *string) *ExportPolicyModifyCollectionParams {
	o.SetRulesAnonymousUser(rulesAnonymousUser)
	return o
}

// SetRulesAnonymousUser adds the rulesAnonymousUser to the export policy modify collection params
func (o *ExportPolicyModifyCollectionParams) SetRulesAnonymousUser(rulesAnonymousUser *string) {
	o.RulesAnonymousUser = rulesAnonymousUser
}

// WithRulesChownMode adds the rulesChownMode to the export policy modify collection params
func (o *ExportPolicyModifyCollectionParams) WithRulesChownMode(rulesChownMode *string) *ExportPolicyModifyCollectionParams {
	o.SetRulesChownMode(rulesChownMode)
	return o
}

// SetRulesChownMode adds the rulesChownMode to the export policy modify collection params
func (o *ExportPolicyModifyCollectionParams) SetRulesChownMode(rulesChownMode *string) {
	o.RulesChownMode = rulesChownMode
}

// WithRulesClientsMatch adds the rulesClientsMatch to the export policy modify collection params
func (o *ExportPolicyModifyCollectionParams) WithRulesClientsMatch(rulesClientsMatch *string) *ExportPolicyModifyCollectionParams {
	o.SetRulesClientsMatch(rulesClientsMatch)
	return o
}

// SetRulesClientsMatch adds the rulesClientsMatch to the export policy modify collection params
func (o *ExportPolicyModifyCollectionParams) SetRulesClientsMatch(rulesClientsMatch *string) {
	o.RulesClientsMatch = rulesClientsMatch
}

// WithRulesIndex adds the rulesIndex to the export policy modify collection params
func (o *ExportPolicyModifyCollectionParams) WithRulesIndex(rulesIndex *int64) *ExportPolicyModifyCollectionParams {
	o.SetRulesIndex(rulesIndex)
	return o
}

// SetRulesIndex adds the rulesIndex to the export policy modify collection params
func (o *ExportPolicyModifyCollectionParams) SetRulesIndex(rulesIndex *int64) {
	o.RulesIndex = rulesIndex
}

// WithRulesNtfsUnixSecurity adds the rulesNtfsUnixSecurity to the export policy modify collection params
func (o *ExportPolicyModifyCollectionParams) WithRulesNtfsUnixSecurity(rulesNtfsUnixSecurity *string) *ExportPolicyModifyCollectionParams {
	o.SetRulesNtfsUnixSecurity(rulesNtfsUnixSecurity)
	return o
}

// SetRulesNtfsUnixSecurity adds the rulesNtfsUnixSecurity to the export policy modify collection params
func (o *ExportPolicyModifyCollectionParams) SetRulesNtfsUnixSecurity(rulesNtfsUnixSecurity *string) {
	o.RulesNtfsUnixSecurity = rulesNtfsUnixSecurity
}

// WithRulesProtocols adds the rulesProtocols to the export policy modify collection params
func (o *ExportPolicyModifyCollectionParams) WithRulesProtocols(rulesProtocols *string) *ExportPolicyModifyCollectionParams {
	o.SetRulesProtocols(rulesProtocols)
	return o
}

// SetRulesProtocols adds the rulesProtocols to the export policy modify collection params
func (o *ExportPolicyModifyCollectionParams) SetRulesProtocols(rulesProtocols *string) {
	o.RulesProtocols = rulesProtocols
}

// WithRulesRoRule adds the rulesRoRule to the export policy modify collection params
func (o *ExportPolicyModifyCollectionParams) WithRulesRoRule(rulesRoRule *string) *ExportPolicyModifyCollectionParams {
	o.SetRulesRoRule(rulesRoRule)
	return o
}

// SetRulesRoRule adds the rulesRoRule to the export policy modify collection params
func (o *ExportPolicyModifyCollectionParams) SetRulesRoRule(rulesRoRule *string) {
	o.RulesRoRule = rulesRoRule
}

// WithRulesRwRule adds the rulesRwRule to the export policy modify collection params
func (o *ExportPolicyModifyCollectionParams) WithRulesRwRule(rulesRwRule *string) *ExportPolicyModifyCollectionParams {
	o.SetRulesRwRule(rulesRwRule)
	return o
}

// SetRulesRwRule adds the rulesRwRule to the export policy modify collection params
func (o *ExportPolicyModifyCollectionParams) SetRulesRwRule(rulesRwRule *string) {
	o.RulesRwRule = rulesRwRule
}

// WithRulesSuperuser adds the rulesSuperuser to the export policy modify collection params
func (o *ExportPolicyModifyCollectionParams) WithRulesSuperuser(rulesSuperuser *string) *ExportPolicyModifyCollectionParams {
	o.SetRulesSuperuser(rulesSuperuser)
	return o
}

// SetRulesSuperuser adds the rulesSuperuser to the export policy modify collection params
func (o *ExportPolicyModifyCollectionParams) SetRulesSuperuser(rulesSuperuser *string) {
	o.RulesSuperuser = rulesSuperuser
}

// WithSerialRecords adds the serialRecords to the export policy modify collection params
func (o *ExportPolicyModifyCollectionParams) WithSerialRecords(serialRecords *bool) *ExportPolicyModifyCollectionParams {
	o.SetSerialRecords(serialRecords)
	return o
}

// SetSerialRecords adds the serialRecords to the export policy modify collection params
func (o *ExportPolicyModifyCollectionParams) SetSerialRecords(serialRecords *bool) {
	o.SerialRecords = serialRecords
}

// WithSvmName adds the svmName to the export policy modify collection params
func (o *ExportPolicyModifyCollectionParams) WithSvmName(svmName *string) *ExportPolicyModifyCollectionParams {
	o.SetSvmName(svmName)
	return o
}

// SetSvmName adds the svmName to the export policy modify collection params
func (o *ExportPolicyModifyCollectionParams) SetSvmName(svmName *string) {
	o.SvmName = svmName
}

// WithSvmUUID adds the svmUUID to the export policy modify collection params
func (o *ExportPolicyModifyCollectionParams) WithSvmUUID(svmUUID *string) *ExportPolicyModifyCollectionParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the export policy modify collection params
func (o *ExportPolicyModifyCollectionParams) SetSvmUUID(svmUUID *string) {
	o.SvmUUID = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *ExportPolicyModifyCollectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.RulesAllowDeviceCreation != nil {

		// query param rules.allow_device_creation
		var qrRulesAllowDeviceCreation bool

		if o.RulesAllowDeviceCreation != nil {
			qrRulesAllowDeviceCreation = *o.RulesAllowDeviceCreation
		}
		qRulesAllowDeviceCreation := swag.FormatBool(qrRulesAllowDeviceCreation)
		if qRulesAllowDeviceCreation != "" {

			if err := r.SetQueryParam("rules.allow_device_creation", qRulesAllowDeviceCreation); err != nil {
				return err
			}
		}
	}

	if o.RulesAllowSuid != nil {

		// query param rules.allow_suid
		var qrRulesAllowSuid bool

		if o.RulesAllowSuid != nil {
			qrRulesAllowSuid = *o.RulesAllowSuid
		}
		qRulesAllowSuid := swag.FormatBool(qrRulesAllowSuid)
		if qRulesAllowSuid != "" {

			if err := r.SetQueryParam("rules.allow_suid", qRulesAllowSuid); err != nil {
				return err
			}
		}
	}

	if o.RulesAnonymousUser != nil {

		// query param rules.anonymous_user
		var qrRulesAnonymousUser string

		if o.RulesAnonymousUser != nil {
			qrRulesAnonymousUser = *o.RulesAnonymousUser
		}
		qRulesAnonymousUser := qrRulesAnonymousUser
		if qRulesAnonymousUser != "" {

			if err := r.SetQueryParam("rules.anonymous_user", qRulesAnonymousUser); err != nil {
				return err
			}
		}
	}

	if o.RulesChownMode != nil {

		// query param rules.chown_mode
		var qrRulesChownMode string

		if o.RulesChownMode != nil {
			qrRulesChownMode = *o.RulesChownMode
		}
		qRulesChownMode := qrRulesChownMode
		if qRulesChownMode != "" {

			if err := r.SetQueryParam("rules.chown_mode", qRulesChownMode); err != nil {
				return err
			}
		}
	}

	if o.RulesClientsMatch != nil {

		// query param rules.clients.match
		var qrRulesClientsMatch string

		if o.RulesClientsMatch != nil {
			qrRulesClientsMatch = *o.RulesClientsMatch
		}
		qRulesClientsMatch := qrRulesClientsMatch
		if qRulesClientsMatch != "" {

			if err := r.SetQueryParam("rules.clients.match", qRulesClientsMatch); err != nil {
				return err
			}
		}
	}

	if o.RulesIndex != nil {

		// query param rules.index
		var qrRulesIndex int64

		if o.RulesIndex != nil {
			qrRulesIndex = *o.RulesIndex
		}
		qRulesIndex := swag.FormatInt64(qrRulesIndex)
		if qRulesIndex != "" {

			if err := r.SetQueryParam("rules.index", qRulesIndex); err != nil {
				return err
			}
		}
	}

	if o.RulesNtfsUnixSecurity != nil {

		// query param rules.ntfs_unix_security
		var qrRulesNtfsUnixSecurity string

		if o.RulesNtfsUnixSecurity != nil {
			qrRulesNtfsUnixSecurity = *o.RulesNtfsUnixSecurity
		}
		qRulesNtfsUnixSecurity := qrRulesNtfsUnixSecurity
		if qRulesNtfsUnixSecurity != "" {

			if err := r.SetQueryParam("rules.ntfs_unix_security", qRulesNtfsUnixSecurity); err != nil {
				return err
			}
		}
	}

	if o.RulesProtocols != nil {

		// query param rules.protocols
		var qrRulesProtocols string

		if o.RulesProtocols != nil {
			qrRulesProtocols = *o.RulesProtocols
		}
		qRulesProtocols := qrRulesProtocols
		if qRulesProtocols != "" {

			if err := r.SetQueryParam("rules.protocols", qRulesProtocols); err != nil {
				return err
			}
		}
	}

	if o.RulesRoRule != nil {

		// query param rules.ro_rule
		var qrRulesRoRule string

		if o.RulesRoRule != nil {
			qrRulesRoRule = *o.RulesRoRule
		}
		qRulesRoRule := qrRulesRoRule
		if qRulesRoRule != "" {

			if err := r.SetQueryParam("rules.ro_rule", qRulesRoRule); err != nil {
				return err
			}
		}
	}

	if o.RulesRwRule != nil {

		// query param rules.rw_rule
		var qrRulesRwRule string

		if o.RulesRwRule != nil {
			qrRulesRwRule = *o.RulesRwRule
		}
		qRulesRwRule := qrRulesRwRule
		if qRulesRwRule != "" {

			if err := r.SetQueryParam("rules.rw_rule", qRulesRwRule); err != nil {
				return err
			}
		}
	}

	if o.RulesSuperuser != nil {

		// query param rules.superuser
		var qrRulesSuperuser string

		if o.RulesSuperuser != nil {
			qrRulesSuperuser = *o.RulesSuperuser
		}
		qRulesSuperuser := qrRulesSuperuser
		if qRulesSuperuser != "" {

			if err := r.SetQueryParam("rules.superuser", qRulesSuperuser); err != nil {
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