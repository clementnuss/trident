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

// NewGcpKmsModifyCollectionParams creates a new GcpKmsModifyCollectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGcpKmsModifyCollectionParams() *GcpKmsModifyCollectionParams {
	return &GcpKmsModifyCollectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGcpKmsModifyCollectionParamsWithTimeout creates a new GcpKmsModifyCollectionParams object
// with the ability to set a timeout on a request.
func NewGcpKmsModifyCollectionParamsWithTimeout(timeout time.Duration) *GcpKmsModifyCollectionParams {
	return &GcpKmsModifyCollectionParams{
		timeout: timeout,
	}
}

// NewGcpKmsModifyCollectionParamsWithContext creates a new GcpKmsModifyCollectionParams object
// with the ability to set a context for a request.
func NewGcpKmsModifyCollectionParamsWithContext(ctx context.Context) *GcpKmsModifyCollectionParams {
	return &GcpKmsModifyCollectionParams{
		Context: ctx,
	}
}

// NewGcpKmsModifyCollectionParamsWithHTTPClient creates a new GcpKmsModifyCollectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGcpKmsModifyCollectionParamsWithHTTPClient(client *http.Client) *GcpKmsModifyCollectionParams {
	return &GcpKmsModifyCollectionParams{
		HTTPClient: client,
	}
}

/*
GcpKmsModifyCollectionParams contains all the parameters to send to the API endpoint

	for the gcp kms modify collection operation.

	Typically these are written to a http.Request.
*/
type GcpKmsModifyCollectionParams struct {

	/* CallerAccount.

	   Filter by caller_account
	*/
	CallerAccount *string

	/* CloudkmsHost.

	   Filter by cloudkms_host
	*/
	CloudkmsHost *string

	/* ContinueOnFailure.

	   Continue even when the operation fails on one of the records.
	*/
	ContinueOnFailure *bool

	/* EkmipReachabilityCode.

	   Filter by ekmip_reachability.code
	*/
	EkmipReachabilityCode *string

	/* EkmipReachabilityMessage.

	   Filter by ekmip_reachability.message
	*/
	EkmipReachabilityMessage *string

	/* EkmipReachabilityNodeName.

	   Filter by ekmip_reachability.node.name
	*/
	EkmipReachabilityNodeName *string

	/* EkmipReachabilityNodeUUID.

	   Filter by ekmip_reachability.node.uuid
	*/
	EkmipReachabilityNodeUUID *string

	/* EkmipReachabilityReachable.

	   Filter by ekmip_reachability.reachable
	*/
	EkmipReachabilityReachable *bool

	/* GoogleReachabilityCode.

	   Filter by google_reachability.code
	*/
	GoogleReachabilityCode *string

	/* GoogleReachabilityMessage.

	   Filter by google_reachability.message
	*/
	GoogleReachabilityMessage *string

	/* GoogleReachabilityReachable.

	   Filter by google_reachability.reachable
	*/
	GoogleReachabilityReachable *bool

	/* Info.

	   Info specification
	*/
	Info GcpKmsModifyCollectionBody

	/* KeyName.

	   Filter by key_name
	*/
	KeyName *string

	/* KeyRingLocation.

	   Filter by key_ring_location
	*/
	KeyRingLocation *string

	/* KeyRingName.

	   Filter by key_ring_name
	*/
	KeyRingName *string

	/* OauthHost.

	   Filter by oauth_host
	*/
	OauthHost *string

	/* OauthURL.

	   Filter by oauth_url
	*/
	OauthURL *string

	/* Port.

	   Filter by port
	*/
	Port *int64

	/* PrivilegedAccount.

	   Filter by privileged_account
	*/
	PrivilegedAccount *string

	/* ProjectID.

	   Filter by project_id
	*/
	ProjectID *string

	/* ProxyHost.

	   Filter by proxy_host
	*/
	ProxyHost *string

	/* ProxyPort.

	   Filter by proxy_port
	*/
	ProxyPort *int64

	/* ProxyType.

	   Filter by proxy_type
	*/
	ProxyType *string

	/* ProxyUsername.

	   Filter by proxy_username
	*/
	ProxyUsername *string

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

	/* Scope.

	   Filter by scope
	*/
	Scope *string

	/* SerialRecords.

	   Perform the operation on the records synchronously.
	*/
	SerialRecords *bool

	/* StateClusterState.

	   Filter by state.cluster_state
	*/
	StateClusterState *bool

	/* StateCode.

	   Filter by state.code
	*/
	StateCode *string

	/* StateMessage.

	   Filter by state.message
	*/
	StateMessage *string

	/* SvmName.

	   Filter by svm.name
	*/
	SvmName *string

	/* SvmUUID.

	   Filter by svm.uuid
	*/
	SvmUUID *string

	/* UUID.

	   Filter by uuid
	*/
	UUID *string

	/* VerifyHost.

	   Filter by verify_host
	*/
	VerifyHost *bool

	/* VerifyIP.

	   Filter by verify_ip
	*/
	VerifyIP *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the gcp kms modify collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GcpKmsModifyCollectionParams) WithDefaults() *GcpKmsModifyCollectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the gcp kms modify collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GcpKmsModifyCollectionParams) SetDefaults() {
	var (
		continueOnFailureDefault = bool(false)

		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)

		serialRecordsDefault = bool(false)
	)

	val := GcpKmsModifyCollectionParams{
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

// WithTimeout adds the timeout to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) WithTimeout(timeout time.Duration) *GcpKmsModifyCollectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) WithContext(ctx context.Context) *GcpKmsModifyCollectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) WithHTTPClient(client *http.Client) *GcpKmsModifyCollectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCallerAccount adds the callerAccount to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) WithCallerAccount(callerAccount *string) *GcpKmsModifyCollectionParams {
	o.SetCallerAccount(callerAccount)
	return o
}

// SetCallerAccount adds the callerAccount to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) SetCallerAccount(callerAccount *string) {
	o.CallerAccount = callerAccount
}

// WithCloudkmsHost adds the cloudkmsHost to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) WithCloudkmsHost(cloudkmsHost *string) *GcpKmsModifyCollectionParams {
	o.SetCloudkmsHost(cloudkmsHost)
	return o
}

// SetCloudkmsHost adds the cloudkmsHost to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) SetCloudkmsHost(cloudkmsHost *string) {
	o.CloudkmsHost = cloudkmsHost
}

// WithContinueOnFailure adds the continueOnFailure to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) WithContinueOnFailure(continueOnFailure *bool) *GcpKmsModifyCollectionParams {
	o.SetContinueOnFailure(continueOnFailure)
	return o
}

// SetContinueOnFailure adds the continueOnFailure to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) SetContinueOnFailure(continueOnFailure *bool) {
	o.ContinueOnFailure = continueOnFailure
}

// WithEkmipReachabilityCode adds the ekmipReachabilityCode to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) WithEkmipReachabilityCode(ekmipReachabilityCode *string) *GcpKmsModifyCollectionParams {
	o.SetEkmipReachabilityCode(ekmipReachabilityCode)
	return o
}

// SetEkmipReachabilityCode adds the ekmipReachabilityCode to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) SetEkmipReachabilityCode(ekmipReachabilityCode *string) {
	o.EkmipReachabilityCode = ekmipReachabilityCode
}

// WithEkmipReachabilityMessage adds the ekmipReachabilityMessage to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) WithEkmipReachabilityMessage(ekmipReachabilityMessage *string) *GcpKmsModifyCollectionParams {
	o.SetEkmipReachabilityMessage(ekmipReachabilityMessage)
	return o
}

// SetEkmipReachabilityMessage adds the ekmipReachabilityMessage to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) SetEkmipReachabilityMessage(ekmipReachabilityMessage *string) {
	o.EkmipReachabilityMessage = ekmipReachabilityMessage
}

// WithEkmipReachabilityNodeName adds the ekmipReachabilityNodeName to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) WithEkmipReachabilityNodeName(ekmipReachabilityNodeName *string) *GcpKmsModifyCollectionParams {
	o.SetEkmipReachabilityNodeName(ekmipReachabilityNodeName)
	return o
}

// SetEkmipReachabilityNodeName adds the ekmipReachabilityNodeName to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) SetEkmipReachabilityNodeName(ekmipReachabilityNodeName *string) {
	o.EkmipReachabilityNodeName = ekmipReachabilityNodeName
}

// WithEkmipReachabilityNodeUUID adds the ekmipReachabilityNodeUUID to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) WithEkmipReachabilityNodeUUID(ekmipReachabilityNodeUUID *string) *GcpKmsModifyCollectionParams {
	o.SetEkmipReachabilityNodeUUID(ekmipReachabilityNodeUUID)
	return o
}

// SetEkmipReachabilityNodeUUID adds the ekmipReachabilityNodeUuid to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) SetEkmipReachabilityNodeUUID(ekmipReachabilityNodeUUID *string) {
	o.EkmipReachabilityNodeUUID = ekmipReachabilityNodeUUID
}

// WithEkmipReachabilityReachable adds the ekmipReachabilityReachable to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) WithEkmipReachabilityReachable(ekmipReachabilityReachable *bool) *GcpKmsModifyCollectionParams {
	o.SetEkmipReachabilityReachable(ekmipReachabilityReachable)
	return o
}

// SetEkmipReachabilityReachable adds the ekmipReachabilityReachable to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) SetEkmipReachabilityReachable(ekmipReachabilityReachable *bool) {
	o.EkmipReachabilityReachable = ekmipReachabilityReachable
}

// WithGoogleReachabilityCode adds the googleReachabilityCode to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) WithGoogleReachabilityCode(googleReachabilityCode *string) *GcpKmsModifyCollectionParams {
	o.SetGoogleReachabilityCode(googleReachabilityCode)
	return o
}

// SetGoogleReachabilityCode adds the googleReachabilityCode to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) SetGoogleReachabilityCode(googleReachabilityCode *string) {
	o.GoogleReachabilityCode = googleReachabilityCode
}

// WithGoogleReachabilityMessage adds the googleReachabilityMessage to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) WithGoogleReachabilityMessage(googleReachabilityMessage *string) *GcpKmsModifyCollectionParams {
	o.SetGoogleReachabilityMessage(googleReachabilityMessage)
	return o
}

// SetGoogleReachabilityMessage adds the googleReachabilityMessage to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) SetGoogleReachabilityMessage(googleReachabilityMessage *string) {
	o.GoogleReachabilityMessage = googleReachabilityMessage
}

// WithGoogleReachabilityReachable adds the googleReachabilityReachable to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) WithGoogleReachabilityReachable(googleReachabilityReachable *bool) *GcpKmsModifyCollectionParams {
	o.SetGoogleReachabilityReachable(googleReachabilityReachable)
	return o
}

// SetGoogleReachabilityReachable adds the googleReachabilityReachable to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) SetGoogleReachabilityReachable(googleReachabilityReachable *bool) {
	o.GoogleReachabilityReachable = googleReachabilityReachable
}

// WithInfo adds the info to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) WithInfo(info GcpKmsModifyCollectionBody) *GcpKmsModifyCollectionParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) SetInfo(info GcpKmsModifyCollectionBody) {
	o.Info = info
}

// WithKeyName adds the keyName to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) WithKeyName(keyName *string) *GcpKmsModifyCollectionParams {
	o.SetKeyName(keyName)
	return o
}

// SetKeyName adds the keyName to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) SetKeyName(keyName *string) {
	o.KeyName = keyName
}

// WithKeyRingLocation adds the keyRingLocation to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) WithKeyRingLocation(keyRingLocation *string) *GcpKmsModifyCollectionParams {
	o.SetKeyRingLocation(keyRingLocation)
	return o
}

// SetKeyRingLocation adds the keyRingLocation to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) SetKeyRingLocation(keyRingLocation *string) {
	o.KeyRingLocation = keyRingLocation
}

// WithKeyRingName adds the keyRingName to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) WithKeyRingName(keyRingName *string) *GcpKmsModifyCollectionParams {
	o.SetKeyRingName(keyRingName)
	return o
}

// SetKeyRingName adds the keyRingName to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) SetKeyRingName(keyRingName *string) {
	o.KeyRingName = keyRingName
}

// WithOauthHost adds the oauthHost to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) WithOauthHost(oauthHost *string) *GcpKmsModifyCollectionParams {
	o.SetOauthHost(oauthHost)
	return o
}

// SetOauthHost adds the oauthHost to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) SetOauthHost(oauthHost *string) {
	o.OauthHost = oauthHost
}

// WithOauthURL adds the oauthURL to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) WithOauthURL(oauthURL *string) *GcpKmsModifyCollectionParams {
	o.SetOauthURL(oauthURL)
	return o
}

// SetOauthURL adds the oauthUrl to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) SetOauthURL(oauthURL *string) {
	o.OauthURL = oauthURL
}

// WithPort adds the port to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) WithPort(port *int64) *GcpKmsModifyCollectionParams {
	o.SetPort(port)
	return o
}

// SetPort adds the port to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) SetPort(port *int64) {
	o.Port = port
}

// WithPrivilegedAccount adds the privilegedAccount to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) WithPrivilegedAccount(privilegedAccount *string) *GcpKmsModifyCollectionParams {
	o.SetPrivilegedAccount(privilegedAccount)
	return o
}

// SetPrivilegedAccount adds the privilegedAccount to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) SetPrivilegedAccount(privilegedAccount *string) {
	o.PrivilegedAccount = privilegedAccount
}

// WithProjectID adds the projectID to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) WithProjectID(projectID *string) *GcpKmsModifyCollectionParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) SetProjectID(projectID *string) {
	o.ProjectID = projectID
}

// WithProxyHost adds the proxyHost to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) WithProxyHost(proxyHost *string) *GcpKmsModifyCollectionParams {
	o.SetProxyHost(proxyHost)
	return o
}

// SetProxyHost adds the proxyHost to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) SetProxyHost(proxyHost *string) {
	o.ProxyHost = proxyHost
}

// WithProxyPort adds the proxyPort to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) WithProxyPort(proxyPort *int64) *GcpKmsModifyCollectionParams {
	o.SetProxyPort(proxyPort)
	return o
}

// SetProxyPort adds the proxyPort to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) SetProxyPort(proxyPort *int64) {
	o.ProxyPort = proxyPort
}

// WithProxyType adds the proxyType to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) WithProxyType(proxyType *string) *GcpKmsModifyCollectionParams {
	o.SetProxyType(proxyType)
	return o
}

// SetProxyType adds the proxyType to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) SetProxyType(proxyType *string) {
	o.ProxyType = proxyType
}

// WithProxyUsername adds the proxyUsername to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) WithProxyUsername(proxyUsername *string) *GcpKmsModifyCollectionParams {
	o.SetProxyUsername(proxyUsername)
	return o
}

// SetProxyUsername adds the proxyUsername to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) SetProxyUsername(proxyUsername *string) {
	o.ProxyUsername = proxyUsername
}

// WithReturnRecords adds the returnRecords to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) WithReturnRecords(returnRecords *bool) *GcpKmsModifyCollectionParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) WithReturnTimeout(returnTimeout *int64) *GcpKmsModifyCollectionParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithScope adds the scope to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) WithScope(scope *string) *GcpKmsModifyCollectionParams {
	o.SetScope(scope)
	return o
}

// SetScope adds the scope to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) SetScope(scope *string) {
	o.Scope = scope
}

// WithSerialRecords adds the serialRecords to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) WithSerialRecords(serialRecords *bool) *GcpKmsModifyCollectionParams {
	o.SetSerialRecords(serialRecords)
	return o
}

// SetSerialRecords adds the serialRecords to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) SetSerialRecords(serialRecords *bool) {
	o.SerialRecords = serialRecords
}

// WithStateClusterState adds the stateClusterState to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) WithStateClusterState(stateClusterState *bool) *GcpKmsModifyCollectionParams {
	o.SetStateClusterState(stateClusterState)
	return o
}

// SetStateClusterState adds the stateClusterState to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) SetStateClusterState(stateClusterState *bool) {
	o.StateClusterState = stateClusterState
}

// WithStateCode adds the stateCode to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) WithStateCode(stateCode *string) *GcpKmsModifyCollectionParams {
	o.SetStateCode(stateCode)
	return o
}

// SetStateCode adds the stateCode to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) SetStateCode(stateCode *string) {
	o.StateCode = stateCode
}

// WithStateMessage adds the stateMessage to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) WithStateMessage(stateMessage *string) *GcpKmsModifyCollectionParams {
	o.SetStateMessage(stateMessage)
	return o
}

// SetStateMessage adds the stateMessage to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) SetStateMessage(stateMessage *string) {
	o.StateMessage = stateMessage
}

// WithSvmName adds the svmName to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) WithSvmName(svmName *string) *GcpKmsModifyCollectionParams {
	o.SetSvmName(svmName)
	return o
}

// SetSvmName adds the svmName to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) SetSvmName(svmName *string) {
	o.SvmName = svmName
}

// WithSvmUUID adds the svmUUID to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) WithSvmUUID(svmUUID *string) *GcpKmsModifyCollectionParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) SetSvmUUID(svmUUID *string) {
	o.SvmUUID = svmUUID
}

// WithUUID adds the uuid to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) WithUUID(uuid *string) *GcpKmsModifyCollectionParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) SetUUID(uuid *string) {
	o.UUID = uuid
}

// WithVerifyHost adds the verifyHost to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) WithVerifyHost(verifyHost *bool) *GcpKmsModifyCollectionParams {
	o.SetVerifyHost(verifyHost)
	return o
}

// SetVerifyHost adds the verifyHost to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) SetVerifyHost(verifyHost *bool) {
	o.VerifyHost = verifyHost
}

// WithVerifyIP adds the verifyIP to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) WithVerifyIP(verifyIP *bool) *GcpKmsModifyCollectionParams {
	o.SetVerifyIP(verifyIP)
	return o
}

// SetVerifyIP adds the verifyIp to the gcp kms modify collection params
func (o *GcpKmsModifyCollectionParams) SetVerifyIP(verifyIP *bool) {
	o.VerifyIP = verifyIP
}

// WriteToRequest writes these params to a swagger request
func (o *GcpKmsModifyCollectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CallerAccount != nil {

		// query param caller_account
		var qrCallerAccount string

		if o.CallerAccount != nil {
			qrCallerAccount = *o.CallerAccount
		}
		qCallerAccount := qrCallerAccount
		if qCallerAccount != "" {

			if err := r.SetQueryParam("caller_account", qCallerAccount); err != nil {
				return err
			}
		}
	}

	if o.CloudkmsHost != nil {

		// query param cloudkms_host
		var qrCloudkmsHost string

		if o.CloudkmsHost != nil {
			qrCloudkmsHost = *o.CloudkmsHost
		}
		qCloudkmsHost := qrCloudkmsHost
		if qCloudkmsHost != "" {

			if err := r.SetQueryParam("cloudkms_host", qCloudkmsHost); err != nil {
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

	if o.EkmipReachabilityCode != nil {

		// query param ekmip_reachability.code
		var qrEkmipReachabilityCode string

		if o.EkmipReachabilityCode != nil {
			qrEkmipReachabilityCode = *o.EkmipReachabilityCode
		}
		qEkmipReachabilityCode := qrEkmipReachabilityCode
		if qEkmipReachabilityCode != "" {

			if err := r.SetQueryParam("ekmip_reachability.code", qEkmipReachabilityCode); err != nil {
				return err
			}
		}
	}

	if o.EkmipReachabilityMessage != nil {

		// query param ekmip_reachability.message
		var qrEkmipReachabilityMessage string

		if o.EkmipReachabilityMessage != nil {
			qrEkmipReachabilityMessage = *o.EkmipReachabilityMessage
		}
		qEkmipReachabilityMessage := qrEkmipReachabilityMessage
		if qEkmipReachabilityMessage != "" {

			if err := r.SetQueryParam("ekmip_reachability.message", qEkmipReachabilityMessage); err != nil {
				return err
			}
		}
	}

	if o.EkmipReachabilityNodeName != nil {

		// query param ekmip_reachability.node.name
		var qrEkmipReachabilityNodeName string

		if o.EkmipReachabilityNodeName != nil {
			qrEkmipReachabilityNodeName = *o.EkmipReachabilityNodeName
		}
		qEkmipReachabilityNodeName := qrEkmipReachabilityNodeName
		if qEkmipReachabilityNodeName != "" {

			if err := r.SetQueryParam("ekmip_reachability.node.name", qEkmipReachabilityNodeName); err != nil {
				return err
			}
		}
	}

	if o.EkmipReachabilityNodeUUID != nil {

		// query param ekmip_reachability.node.uuid
		var qrEkmipReachabilityNodeUUID string

		if o.EkmipReachabilityNodeUUID != nil {
			qrEkmipReachabilityNodeUUID = *o.EkmipReachabilityNodeUUID
		}
		qEkmipReachabilityNodeUUID := qrEkmipReachabilityNodeUUID
		if qEkmipReachabilityNodeUUID != "" {

			if err := r.SetQueryParam("ekmip_reachability.node.uuid", qEkmipReachabilityNodeUUID); err != nil {
				return err
			}
		}
	}

	if o.EkmipReachabilityReachable != nil {

		// query param ekmip_reachability.reachable
		var qrEkmipReachabilityReachable bool

		if o.EkmipReachabilityReachable != nil {
			qrEkmipReachabilityReachable = *o.EkmipReachabilityReachable
		}
		qEkmipReachabilityReachable := swag.FormatBool(qrEkmipReachabilityReachable)
		if qEkmipReachabilityReachable != "" {

			if err := r.SetQueryParam("ekmip_reachability.reachable", qEkmipReachabilityReachable); err != nil {
				return err
			}
		}
	}

	if o.GoogleReachabilityCode != nil {

		// query param google_reachability.code
		var qrGoogleReachabilityCode string

		if o.GoogleReachabilityCode != nil {
			qrGoogleReachabilityCode = *o.GoogleReachabilityCode
		}
		qGoogleReachabilityCode := qrGoogleReachabilityCode
		if qGoogleReachabilityCode != "" {

			if err := r.SetQueryParam("google_reachability.code", qGoogleReachabilityCode); err != nil {
				return err
			}
		}
	}

	if o.GoogleReachabilityMessage != nil {

		// query param google_reachability.message
		var qrGoogleReachabilityMessage string

		if o.GoogleReachabilityMessage != nil {
			qrGoogleReachabilityMessage = *o.GoogleReachabilityMessage
		}
		qGoogleReachabilityMessage := qrGoogleReachabilityMessage
		if qGoogleReachabilityMessage != "" {

			if err := r.SetQueryParam("google_reachability.message", qGoogleReachabilityMessage); err != nil {
				return err
			}
		}
	}

	if o.GoogleReachabilityReachable != nil {

		// query param google_reachability.reachable
		var qrGoogleReachabilityReachable bool

		if o.GoogleReachabilityReachable != nil {
			qrGoogleReachabilityReachable = *o.GoogleReachabilityReachable
		}
		qGoogleReachabilityReachable := swag.FormatBool(qrGoogleReachabilityReachable)
		if qGoogleReachabilityReachable != "" {

			if err := r.SetQueryParam("google_reachability.reachable", qGoogleReachabilityReachable); err != nil {
				return err
			}
		}
	}
	if err := r.SetBodyParam(o.Info); err != nil {
		return err
	}

	if o.KeyName != nil {

		// query param key_name
		var qrKeyName string

		if o.KeyName != nil {
			qrKeyName = *o.KeyName
		}
		qKeyName := qrKeyName
		if qKeyName != "" {

			if err := r.SetQueryParam("key_name", qKeyName); err != nil {
				return err
			}
		}
	}

	if o.KeyRingLocation != nil {

		// query param key_ring_location
		var qrKeyRingLocation string

		if o.KeyRingLocation != nil {
			qrKeyRingLocation = *o.KeyRingLocation
		}
		qKeyRingLocation := qrKeyRingLocation
		if qKeyRingLocation != "" {

			if err := r.SetQueryParam("key_ring_location", qKeyRingLocation); err != nil {
				return err
			}
		}
	}

	if o.KeyRingName != nil {

		// query param key_ring_name
		var qrKeyRingName string

		if o.KeyRingName != nil {
			qrKeyRingName = *o.KeyRingName
		}
		qKeyRingName := qrKeyRingName
		if qKeyRingName != "" {

			if err := r.SetQueryParam("key_ring_name", qKeyRingName); err != nil {
				return err
			}
		}
	}

	if o.OauthHost != nil {

		// query param oauth_host
		var qrOauthHost string

		if o.OauthHost != nil {
			qrOauthHost = *o.OauthHost
		}
		qOauthHost := qrOauthHost
		if qOauthHost != "" {

			if err := r.SetQueryParam("oauth_host", qOauthHost); err != nil {
				return err
			}
		}
	}

	if o.OauthURL != nil {

		// query param oauth_url
		var qrOauthURL string

		if o.OauthURL != nil {
			qrOauthURL = *o.OauthURL
		}
		qOauthURL := qrOauthURL
		if qOauthURL != "" {

			if err := r.SetQueryParam("oauth_url", qOauthURL); err != nil {
				return err
			}
		}
	}

	if o.Port != nil {

		// query param port
		var qrPort int64

		if o.Port != nil {
			qrPort = *o.Port
		}
		qPort := swag.FormatInt64(qrPort)
		if qPort != "" {

			if err := r.SetQueryParam("port", qPort); err != nil {
				return err
			}
		}
	}

	if o.PrivilegedAccount != nil {

		// query param privileged_account
		var qrPrivilegedAccount string

		if o.PrivilegedAccount != nil {
			qrPrivilegedAccount = *o.PrivilegedAccount
		}
		qPrivilegedAccount := qrPrivilegedAccount
		if qPrivilegedAccount != "" {

			if err := r.SetQueryParam("privileged_account", qPrivilegedAccount); err != nil {
				return err
			}
		}
	}

	if o.ProjectID != nil {

		// query param project_id
		var qrProjectID string

		if o.ProjectID != nil {
			qrProjectID = *o.ProjectID
		}
		qProjectID := qrProjectID
		if qProjectID != "" {

			if err := r.SetQueryParam("project_id", qProjectID); err != nil {
				return err
			}
		}
	}

	if o.ProxyHost != nil {

		// query param proxy_host
		var qrProxyHost string

		if o.ProxyHost != nil {
			qrProxyHost = *o.ProxyHost
		}
		qProxyHost := qrProxyHost
		if qProxyHost != "" {

			if err := r.SetQueryParam("proxy_host", qProxyHost); err != nil {
				return err
			}
		}
	}

	if o.ProxyPort != nil {

		// query param proxy_port
		var qrProxyPort int64

		if o.ProxyPort != nil {
			qrProxyPort = *o.ProxyPort
		}
		qProxyPort := swag.FormatInt64(qrProxyPort)
		if qProxyPort != "" {

			if err := r.SetQueryParam("proxy_port", qProxyPort); err != nil {
				return err
			}
		}
	}

	if o.ProxyType != nil {

		// query param proxy_type
		var qrProxyType string

		if o.ProxyType != nil {
			qrProxyType = *o.ProxyType
		}
		qProxyType := qrProxyType
		if qProxyType != "" {

			if err := r.SetQueryParam("proxy_type", qProxyType); err != nil {
				return err
			}
		}
	}

	if o.ProxyUsername != nil {

		// query param proxy_username
		var qrProxyUsername string

		if o.ProxyUsername != nil {
			qrProxyUsername = *o.ProxyUsername
		}
		qProxyUsername := qrProxyUsername
		if qProxyUsername != "" {

			if err := r.SetQueryParam("proxy_username", qProxyUsername); err != nil {
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

	if o.Scope != nil {

		// query param scope
		var qrScope string

		if o.Scope != nil {
			qrScope = *o.Scope
		}
		qScope := qrScope
		if qScope != "" {

			if err := r.SetQueryParam("scope", qScope); err != nil {
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

	if o.StateClusterState != nil {

		// query param state.cluster_state
		var qrStateClusterState bool

		if o.StateClusterState != nil {
			qrStateClusterState = *o.StateClusterState
		}
		qStateClusterState := swag.FormatBool(qrStateClusterState)
		if qStateClusterState != "" {

			if err := r.SetQueryParam("state.cluster_state", qStateClusterState); err != nil {
				return err
			}
		}
	}

	if o.StateCode != nil {

		// query param state.code
		var qrStateCode string

		if o.StateCode != nil {
			qrStateCode = *o.StateCode
		}
		qStateCode := qrStateCode
		if qStateCode != "" {

			if err := r.SetQueryParam("state.code", qStateCode); err != nil {
				return err
			}
		}
	}

	if o.StateMessage != nil {

		// query param state.message
		var qrStateMessage string

		if o.StateMessage != nil {
			qrStateMessage = *o.StateMessage
		}
		qStateMessage := qrStateMessage
		if qStateMessage != "" {

			if err := r.SetQueryParam("state.message", qStateMessage); err != nil {
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

	if o.VerifyHost != nil {

		// query param verify_host
		var qrVerifyHost bool

		if o.VerifyHost != nil {
			qrVerifyHost = *o.VerifyHost
		}
		qVerifyHost := swag.FormatBool(qrVerifyHost)
		if qVerifyHost != "" {

			if err := r.SetQueryParam("verify_host", qVerifyHost); err != nil {
				return err
			}
		}
	}

	if o.VerifyIP != nil {

		// query param verify_ip
		var qrVerifyIP bool

		if o.VerifyIP != nil {
			qrVerifyIP = *o.VerifyIP
		}
		qVerifyIP := swag.FormatBool(qrVerifyIP)
		if qVerifyIP != "" {

			if err := r.SetQueryParam("verify_ip", qVerifyIP); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}