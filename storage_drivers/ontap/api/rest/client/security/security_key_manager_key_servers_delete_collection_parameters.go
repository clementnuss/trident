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

// NewSecurityKeyManagerKeyServersDeleteCollectionParams creates a new SecurityKeyManagerKeyServersDeleteCollectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSecurityKeyManagerKeyServersDeleteCollectionParams() *SecurityKeyManagerKeyServersDeleteCollectionParams {
	return &SecurityKeyManagerKeyServersDeleteCollectionParams{
		requestTimeout: cr.DefaultTimeout,
	}
}

// NewSecurityKeyManagerKeyServersDeleteCollectionParamsWithTimeout creates a new SecurityKeyManagerKeyServersDeleteCollectionParams object
// with the ability to set a timeout on a request.
func NewSecurityKeyManagerKeyServersDeleteCollectionParamsWithTimeout(timeout time.Duration) *SecurityKeyManagerKeyServersDeleteCollectionParams {
	return &SecurityKeyManagerKeyServersDeleteCollectionParams{
		requestTimeout: timeout,
	}
}

// NewSecurityKeyManagerKeyServersDeleteCollectionParamsWithContext creates a new SecurityKeyManagerKeyServersDeleteCollectionParams object
// with the ability to set a context for a request.
func NewSecurityKeyManagerKeyServersDeleteCollectionParamsWithContext(ctx context.Context) *SecurityKeyManagerKeyServersDeleteCollectionParams {
	return &SecurityKeyManagerKeyServersDeleteCollectionParams{
		Context: ctx,
	}
}

// NewSecurityKeyManagerKeyServersDeleteCollectionParamsWithHTTPClient creates a new SecurityKeyManagerKeyServersDeleteCollectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewSecurityKeyManagerKeyServersDeleteCollectionParamsWithHTTPClient(client *http.Client) *SecurityKeyManagerKeyServersDeleteCollectionParams {
	return &SecurityKeyManagerKeyServersDeleteCollectionParams{
		HTTPClient: client,
	}
}

/*
SecurityKeyManagerKeyServersDeleteCollectionParams contains all the parameters to send to the API endpoint

	for the security key manager key servers delete collection operation.

	Typically these are written to a http.Request.
*/
type SecurityKeyManagerKeyServersDeleteCollectionParams struct {

	/* ConnectivityClusterAvailability.

	   Filter by connectivity.cluster_availability
	*/
	ConnectivityClusterAvailability *bool

	/* ConnectivityNodeStatesNodeName.

	   Filter by connectivity.node_states.node.name
	*/
	ConnectivityNodeStatesNodeName *string

	/* ConnectivityNodeStatesNodeUUID.

	   Filter by connectivity.node_states.node.uuid
	*/
	ConnectivityNodeStatesNodeUUID *string

	/* ConnectivityNodeStatesState.

	   Filter by connectivity.node_states.state
	*/
	ConnectivityNodeStatesState *string

	/* ContinueOnFailure.

	   Continue even when the operation fails on one of the records.
	*/
	ContinueOnFailure *bool

	/* CreateRemoveTimeout.

	   Filter by create_remove_timeout
	*/
	CreateRemoveTimeout *int64

	/* Force.

	   Set the force flag to "true" to bypass out of quorum checks when removing a primary key server.

	*/
	Force *bool

	/* Info.

	   Info specification
	*/
	Info SecurityKeyManagerKeyServersDeleteCollectionBody

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

	/* SecondaryKeyServers.

	   Filter by secondary_key_servers
	*/
	SecondaryKeyServers *string

	/* SerialRecords.

	   Perform the operation on the records synchronously.
	*/
	SerialRecords *bool

	/* Server.

	   Filter by server
	*/
	Server *string

	/* Timeout.

	   Filter by timeout
	*/
	Timeout *int64

	/* Username.

	   Filter by username
	*/
	Username *string

	/* UUID.

	   External key manager UUID
	*/
	UUID string

	requestTimeout time.Duration
	Context        context.Context
	HTTPClient     *http.Client
}

// WithDefaults hydrates default values in the security key manager key servers delete collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityKeyManagerKeyServersDeleteCollectionParams) WithDefaults() *SecurityKeyManagerKeyServersDeleteCollectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the security key manager key servers delete collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityKeyManagerKeyServersDeleteCollectionParams) SetDefaults() {
	var (
		continueOnFailureDefault = bool(false)

		forceDefault = bool(false)

		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)

		serialRecordsDefault = bool(false)
	)

	val := SecurityKeyManagerKeyServersDeleteCollectionParams{
		ContinueOnFailure: &continueOnFailureDefault,
		Force:             &forceDefault,
		ReturnRecords:     &returnRecordsDefault,
		ReturnTimeout:     &returnTimeoutDefault,
		SerialRecords:     &serialRecordsDefault,
	}

	val.requestTimeout = o.requestTimeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithRequestTimeout adds the timeout to the security key manager key servers delete collection params
func (o *SecurityKeyManagerKeyServersDeleteCollectionParams) WithRequestTimeout(timeout time.Duration) *SecurityKeyManagerKeyServersDeleteCollectionParams {
	o.SetRequestTimeout(timeout)
	return o
}

// SetRequestTimeout adds the timeout to the security key manager key servers delete collection params
func (o *SecurityKeyManagerKeyServersDeleteCollectionParams) SetRequestTimeout(timeout time.Duration) {
	o.requestTimeout = timeout
}

// WithContext adds the context to the security key manager key servers delete collection params
func (o *SecurityKeyManagerKeyServersDeleteCollectionParams) WithContext(ctx context.Context) *SecurityKeyManagerKeyServersDeleteCollectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the security key manager key servers delete collection params
func (o *SecurityKeyManagerKeyServersDeleteCollectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the security key manager key servers delete collection params
func (o *SecurityKeyManagerKeyServersDeleteCollectionParams) WithHTTPClient(client *http.Client) *SecurityKeyManagerKeyServersDeleteCollectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the security key manager key servers delete collection params
func (o *SecurityKeyManagerKeyServersDeleteCollectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConnectivityClusterAvailability adds the connectivityClusterAvailability to the security key manager key servers delete collection params
func (o *SecurityKeyManagerKeyServersDeleteCollectionParams) WithConnectivityClusterAvailability(connectivityClusterAvailability *bool) *SecurityKeyManagerKeyServersDeleteCollectionParams {
	o.SetConnectivityClusterAvailability(connectivityClusterAvailability)
	return o
}

// SetConnectivityClusterAvailability adds the connectivityClusterAvailability to the security key manager key servers delete collection params
func (o *SecurityKeyManagerKeyServersDeleteCollectionParams) SetConnectivityClusterAvailability(connectivityClusterAvailability *bool) {
	o.ConnectivityClusterAvailability = connectivityClusterAvailability
}

// WithConnectivityNodeStatesNodeName adds the connectivityNodeStatesNodeName to the security key manager key servers delete collection params
func (o *SecurityKeyManagerKeyServersDeleteCollectionParams) WithConnectivityNodeStatesNodeName(connectivityNodeStatesNodeName *string) *SecurityKeyManagerKeyServersDeleteCollectionParams {
	o.SetConnectivityNodeStatesNodeName(connectivityNodeStatesNodeName)
	return o
}

// SetConnectivityNodeStatesNodeName adds the connectivityNodeStatesNodeName to the security key manager key servers delete collection params
func (o *SecurityKeyManagerKeyServersDeleteCollectionParams) SetConnectivityNodeStatesNodeName(connectivityNodeStatesNodeName *string) {
	o.ConnectivityNodeStatesNodeName = connectivityNodeStatesNodeName
}

// WithConnectivityNodeStatesNodeUUID adds the connectivityNodeStatesNodeUUID to the security key manager key servers delete collection params
func (o *SecurityKeyManagerKeyServersDeleteCollectionParams) WithConnectivityNodeStatesNodeUUID(connectivityNodeStatesNodeUUID *string) *SecurityKeyManagerKeyServersDeleteCollectionParams {
	o.SetConnectivityNodeStatesNodeUUID(connectivityNodeStatesNodeUUID)
	return o
}

// SetConnectivityNodeStatesNodeUUID adds the connectivityNodeStatesNodeUuid to the security key manager key servers delete collection params
func (o *SecurityKeyManagerKeyServersDeleteCollectionParams) SetConnectivityNodeStatesNodeUUID(connectivityNodeStatesNodeUUID *string) {
	o.ConnectivityNodeStatesNodeUUID = connectivityNodeStatesNodeUUID
}

// WithConnectivityNodeStatesState adds the connectivityNodeStatesState to the security key manager key servers delete collection params
func (o *SecurityKeyManagerKeyServersDeleteCollectionParams) WithConnectivityNodeStatesState(connectivityNodeStatesState *string) *SecurityKeyManagerKeyServersDeleteCollectionParams {
	o.SetConnectivityNodeStatesState(connectivityNodeStatesState)
	return o
}

// SetConnectivityNodeStatesState adds the connectivityNodeStatesState to the security key manager key servers delete collection params
func (o *SecurityKeyManagerKeyServersDeleteCollectionParams) SetConnectivityNodeStatesState(connectivityNodeStatesState *string) {
	o.ConnectivityNodeStatesState = connectivityNodeStatesState
}

// WithContinueOnFailure adds the continueOnFailure to the security key manager key servers delete collection params
func (o *SecurityKeyManagerKeyServersDeleteCollectionParams) WithContinueOnFailure(continueOnFailure *bool) *SecurityKeyManagerKeyServersDeleteCollectionParams {
	o.SetContinueOnFailure(continueOnFailure)
	return o
}

// SetContinueOnFailure adds the continueOnFailure to the security key manager key servers delete collection params
func (o *SecurityKeyManagerKeyServersDeleteCollectionParams) SetContinueOnFailure(continueOnFailure *bool) {
	o.ContinueOnFailure = continueOnFailure
}

// WithCreateRemoveTimeout adds the createRemoveTimeout to the security key manager key servers delete collection params
func (o *SecurityKeyManagerKeyServersDeleteCollectionParams) WithCreateRemoveTimeout(createRemoveTimeout *int64) *SecurityKeyManagerKeyServersDeleteCollectionParams {
	o.SetCreateRemoveTimeout(createRemoveTimeout)
	return o
}

// SetCreateRemoveTimeout adds the createRemoveTimeout to the security key manager key servers delete collection params
func (o *SecurityKeyManagerKeyServersDeleteCollectionParams) SetCreateRemoveTimeout(createRemoveTimeout *int64) {
	o.CreateRemoveTimeout = createRemoveTimeout
}

// WithForce adds the force to the security key manager key servers delete collection params
func (o *SecurityKeyManagerKeyServersDeleteCollectionParams) WithForce(force *bool) *SecurityKeyManagerKeyServersDeleteCollectionParams {
	o.SetForce(force)
	return o
}

// SetForce adds the force to the security key manager key servers delete collection params
func (o *SecurityKeyManagerKeyServersDeleteCollectionParams) SetForce(force *bool) {
	o.Force = force
}

// WithInfo adds the info to the security key manager key servers delete collection params
func (o *SecurityKeyManagerKeyServersDeleteCollectionParams) WithInfo(info SecurityKeyManagerKeyServersDeleteCollectionBody) *SecurityKeyManagerKeyServersDeleteCollectionParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the security key manager key servers delete collection params
func (o *SecurityKeyManagerKeyServersDeleteCollectionParams) SetInfo(info SecurityKeyManagerKeyServersDeleteCollectionBody) {
	o.Info = info
}

// WithReturnRecords adds the returnRecords to the security key manager key servers delete collection params
func (o *SecurityKeyManagerKeyServersDeleteCollectionParams) WithReturnRecords(returnRecords *bool) *SecurityKeyManagerKeyServersDeleteCollectionParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the security key manager key servers delete collection params
func (o *SecurityKeyManagerKeyServersDeleteCollectionParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the security key manager key servers delete collection params
func (o *SecurityKeyManagerKeyServersDeleteCollectionParams) WithReturnTimeout(returnTimeout *int64) *SecurityKeyManagerKeyServersDeleteCollectionParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the security key manager key servers delete collection params
func (o *SecurityKeyManagerKeyServersDeleteCollectionParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSecondaryKeyServers adds the secondaryKeyServers to the security key manager key servers delete collection params
func (o *SecurityKeyManagerKeyServersDeleteCollectionParams) WithSecondaryKeyServers(secondaryKeyServers *string) *SecurityKeyManagerKeyServersDeleteCollectionParams {
	o.SetSecondaryKeyServers(secondaryKeyServers)
	return o
}

// SetSecondaryKeyServers adds the secondaryKeyServers to the security key manager key servers delete collection params
func (o *SecurityKeyManagerKeyServersDeleteCollectionParams) SetSecondaryKeyServers(secondaryKeyServers *string) {
	o.SecondaryKeyServers = secondaryKeyServers
}

// WithSerialRecords adds the serialRecords to the security key manager key servers delete collection params
func (o *SecurityKeyManagerKeyServersDeleteCollectionParams) WithSerialRecords(serialRecords *bool) *SecurityKeyManagerKeyServersDeleteCollectionParams {
	o.SetSerialRecords(serialRecords)
	return o
}

// SetSerialRecords adds the serialRecords to the security key manager key servers delete collection params
func (o *SecurityKeyManagerKeyServersDeleteCollectionParams) SetSerialRecords(serialRecords *bool) {
	o.SerialRecords = serialRecords
}

// WithServer adds the server to the security key manager key servers delete collection params
func (o *SecurityKeyManagerKeyServersDeleteCollectionParams) WithServer(server *string) *SecurityKeyManagerKeyServersDeleteCollectionParams {
	o.SetServer(server)
	return o
}

// SetServer adds the server to the security key manager key servers delete collection params
func (o *SecurityKeyManagerKeyServersDeleteCollectionParams) SetServer(server *string) {
	o.Server = server
}

// WithTimeout adds the timeout to the security key manager key servers delete collection params
func (o *SecurityKeyManagerKeyServersDeleteCollectionParams) WithTimeout(timeout *int64) *SecurityKeyManagerKeyServersDeleteCollectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the security key manager key servers delete collection params
func (o *SecurityKeyManagerKeyServersDeleteCollectionParams) SetTimeout(timeout *int64) {
	o.Timeout = timeout
}

// WithUsername adds the username to the security key manager key servers delete collection params
func (o *SecurityKeyManagerKeyServersDeleteCollectionParams) WithUsername(username *string) *SecurityKeyManagerKeyServersDeleteCollectionParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the security key manager key servers delete collection params
func (o *SecurityKeyManagerKeyServersDeleteCollectionParams) SetUsername(username *string) {
	o.Username = username
}

// WithUUID adds the uuid to the security key manager key servers delete collection params
func (o *SecurityKeyManagerKeyServersDeleteCollectionParams) WithUUID(uuid string) *SecurityKeyManagerKeyServersDeleteCollectionParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the security key manager key servers delete collection params
func (o *SecurityKeyManagerKeyServersDeleteCollectionParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *SecurityKeyManagerKeyServersDeleteCollectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.requestTimeout); err != nil {
		return err
	}
	var res []error

	if o.ConnectivityClusterAvailability != nil {

		// query param connectivity.cluster_availability
		var qrConnectivityClusterAvailability bool

		if o.ConnectivityClusterAvailability != nil {
			qrConnectivityClusterAvailability = *o.ConnectivityClusterAvailability
		}
		qConnectivityClusterAvailability := swag.FormatBool(qrConnectivityClusterAvailability)
		if qConnectivityClusterAvailability != "" {

			if err := r.SetQueryParam("connectivity.cluster_availability", qConnectivityClusterAvailability); err != nil {
				return err
			}
		}
	}

	if o.ConnectivityNodeStatesNodeName != nil {

		// query param connectivity.node_states.node.name
		var qrConnectivityNodeStatesNodeName string

		if o.ConnectivityNodeStatesNodeName != nil {
			qrConnectivityNodeStatesNodeName = *o.ConnectivityNodeStatesNodeName
		}
		qConnectivityNodeStatesNodeName := qrConnectivityNodeStatesNodeName
		if qConnectivityNodeStatesNodeName != "" {

			if err := r.SetQueryParam("connectivity.node_states.node.name", qConnectivityNodeStatesNodeName); err != nil {
				return err
			}
		}
	}

	if o.ConnectivityNodeStatesNodeUUID != nil {

		// query param connectivity.node_states.node.uuid
		var qrConnectivityNodeStatesNodeUUID string

		if o.ConnectivityNodeStatesNodeUUID != nil {
			qrConnectivityNodeStatesNodeUUID = *o.ConnectivityNodeStatesNodeUUID
		}
		qConnectivityNodeStatesNodeUUID := qrConnectivityNodeStatesNodeUUID
		if qConnectivityNodeStatesNodeUUID != "" {

			if err := r.SetQueryParam("connectivity.node_states.node.uuid", qConnectivityNodeStatesNodeUUID); err != nil {
				return err
			}
		}
	}

	if o.ConnectivityNodeStatesState != nil {

		// query param connectivity.node_states.state
		var qrConnectivityNodeStatesState string

		if o.ConnectivityNodeStatesState != nil {
			qrConnectivityNodeStatesState = *o.ConnectivityNodeStatesState
		}
		qConnectivityNodeStatesState := qrConnectivityNodeStatesState
		if qConnectivityNodeStatesState != "" {

			if err := r.SetQueryParam("connectivity.node_states.state", qConnectivityNodeStatesState); err != nil {
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

	if o.CreateRemoveTimeout != nil {

		// query param create_remove_timeout
		var qrCreateRemoveTimeout int64

		if o.CreateRemoveTimeout != nil {
			qrCreateRemoveTimeout = *o.CreateRemoveTimeout
		}
		qCreateRemoveTimeout := swag.FormatInt64(qrCreateRemoveTimeout)
		if qCreateRemoveTimeout != "" {

			if err := r.SetQueryParam("create_remove_timeout", qCreateRemoveTimeout); err != nil {
				return err
			}
		}
	}

	if o.Force != nil {

		// query param force
		var qrForce bool

		if o.Force != nil {
			qrForce = *o.Force
		}
		qForce := swag.FormatBool(qrForce)
		if qForce != "" {

			if err := r.SetQueryParam("force", qForce); err != nil {
				return err
			}
		}
	}
	if err := r.SetBodyParam(o.Info); err != nil {
		return err
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

	if o.SecondaryKeyServers != nil {

		// query param secondary_key_servers
		var qrSecondaryKeyServers string

		if o.SecondaryKeyServers != nil {
			qrSecondaryKeyServers = *o.SecondaryKeyServers
		}
		qSecondaryKeyServers := qrSecondaryKeyServers
		if qSecondaryKeyServers != "" {

			if err := r.SetQueryParam("secondary_key_servers", qSecondaryKeyServers); err != nil {
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

	if o.Server != nil {

		// query param server
		var qrServer string

		if o.Server != nil {
			qrServer = *o.Server
		}
		qServer := qrServer
		if qServer != "" {

			if err := r.SetQueryParam("server", qServer); err != nil {
				return err
			}
		}
	}

	if o.Timeout != nil {

		// query param timeout
		var qrTimeout int64

		if o.Timeout != nil {
			qrTimeout = *o.Timeout
		}
		qTimeout := swag.FormatInt64(qrTimeout)
		if qTimeout != "" {

			if err := r.SetQueryParam("timeout", qTimeout); err != nil {
				return err
			}
		}
	}

	if o.Username != nil {

		// query param username
		var qrUsername string

		if o.Username != nil {
			qrUsername = *o.Username
		}
		qUsername := qrUsername
		if qUsername != "" {

			if err := r.SetQueryParam("username", qUsername); err != nil {
				return err
			}
		}
	}

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}