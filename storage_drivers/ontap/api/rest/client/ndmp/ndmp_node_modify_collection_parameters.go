// Code generated by go-swagger; DO NOT EDIT.

package ndmp

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

// NewNdmpNodeModifyCollectionParams creates a new NdmpNodeModifyCollectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNdmpNodeModifyCollectionParams() *NdmpNodeModifyCollectionParams {
	return &NdmpNodeModifyCollectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNdmpNodeModifyCollectionParamsWithTimeout creates a new NdmpNodeModifyCollectionParams object
// with the ability to set a timeout on a request.
func NewNdmpNodeModifyCollectionParamsWithTimeout(timeout time.Duration) *NdmpNodeModifyCollectionParams {
	return &NdmpNodeModifyCollectionParams{
		timeout: timeout,
	}
}

// NewNdmpNodeModifyCollectionParamsWithContext creates a new NdmpNodeModifyCollectionParams object
// with the ability to set a context for a request.
func NewNdmpNodeModifyCollectionParamsWithContext(ctx context.Context) *NdmpNodeModifyCollectionParams {
	return &NdmpNodeModifyCollectionParams{
		Context: ctx,
	}
}

// NewNdmpNodeModifyCollectionParamsWithHTTPClient creates a new NdmpNodeModifyCollectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewNdmpNodeModifyCollectionParamsWithHTTPClient(client *http.Client) *NdmpNodeModifyCollectionParams {
	return &NdmpNodeModifyCollectionParams{
		HTTPClient: client,
	}
}

/*
NdmpNodeModifyCollectionParams contains all the parameters to send to the API endpoint

	for the ndmp node modify collection operation.

	Typically these are written to a http.Request.
*/
type NdmpNodeModifyCollectionParams struct {

	/* AuthenticationTypes.

	   Filter by authentication_types
	*/
	AuthenticationTypes *string

	/* ContinueOnFailure.

	   Continue even when the operation fails on one of the records.
	*/
	ContinueOnFailure *bool

	/* Enabled.

	   Filter by enabled
	*/
	Enabled *bool

	/* Info.

	   Info specification
	*/
	Info NdmpNodeModifyCollectionBody

	/* NodeName.

	   Filter by node.name
	*/
	NodeName *string

	/* NodeUUID.

	   Filter by node.uuid
	*/
	NodeUUID *string

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

	/* User.

	   Filter by user
	*/
	User *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ndmp node modify collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NdmpNodeModifyCollectionParams) WithDefaults() *NdmpNodeModifyCollectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ndmp node modify collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NdmpNodeModifyCollectionParams) SetDefaults() {
	var (
		continueOnFailureDefault = bool(false)

		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)

		serialRecordsDefault = bool(false)
	)

	val := NdmpNodeModifyCollectionParams{
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

// WithTimeout adds the timeout to the ndmp node modify collection params
func (o *NdmpNodeModifyCollectionParams) WithTimeout(timeout time.Duration) *NdmpNodeModifyCollectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ndmp node modify collection params
func (o *NdmpNodeModifyCollectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ndmp node modify collection params
func (o *NdmpNodeModifyCollectionParams) WithContext(ctx context.Context) *NdmpNodeModifyCollectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ndmp node modify collection params
func (o *NdmpNodeModifyCollectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ndmp node modify collection params
func (o *NdmpNodeModifyCollectionParams) WithHTTPClient(client *http.Client) *NdmpNodeModifyCollectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ndmp node modify collection params
func (o *NdmpNodeModifyCollectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthenticationTypes adds the authenticationTypes to the ndmp node modify collection params
func (o *NdmpNodeModifyCollectionParams) WithAuthenticationTypes(authenticationTypes *string) *NdmpNodeModifyCollectionParams {
	o.SetAuthenticationTypes(authenticationTypes)
	return o
}

// SetAuthenticationTypes adds the authenticationTypes to the ndmp node modify collection params
func (o *NdmpNodeModifyCollectionParams) SetAuthenticationTypes(authenticationTypes *string) {
	o.AuthenticationTypes = authenticationTypes
}

// WithContinueOnFailure adds the continueOnFailure to the ndmp node modify collection params
func (o *NdmpNodeModifyCollectionParams) WithContinueOnFailure(continueOnFailure *bool) *NdmpNodeModifyCollectionParams {
	o.SetContinueOnFailure(continueOnFailure)
	return o
}

// SetContinueOnFailure adds the continueOnFailure to the ndmp node modify collection params
func (o *NdmpNodeModifyCollectionParams) SetContinueOnFailure(continueOnFailure *bool) {
	o.ContinueOnFailure = continueOnFailure
}

// WithEnabled adds the enabled to the ndmp node modify collection params
func (o *NdmpNodeModifyCollectionParams) WithEnabled(enabled *bool) *NdmpNodeModifyCollectionParams {
	o.SetEnabled(enabled)
	return o
}

// SetEnabled adds the enabled to the ndmp node modify collection params
func (o *NdmpNodeModifyCollectionParams) SetEnabled(enabled *bool) {
	o.Enabled = enabled
}

// WithInfo adds the info to the ndmp node modify collection params
func (o *NdmpNodeModifyCollectionParams) WithInfo(info NdmpNodeModifyCollectionBody) *NdmpNodeModifyCollectionParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the ndmp node modify collection params
func (o *NdmpNodeModifyCollectionParams) SetInfo(info NdmpNodeModifyCollectionBody) {
	o.Info = info
}

// WithNodeName adds the nodeName to the ndmp node modify collection params
func (o *NdmpNodeModifyCollectionParams) WithNodeName(nodeName *string) *NdmpNodeModifyCollectionParams {
	o.SetNodeName(nodeName)
	return o
}

// SetNodeName adds the nodeName to the ndmp node modify collection params
func (o *NdmpNodeModifyCollectionParams) SetNodeName(nodeName *string) {
	o.NodeName = nodeName
}

// WithNodeUUID adds the nodeUUID to the ndmp node modify collection params
func (o *NdmpNodeModifyCollectionParams) WithNodeUUID(nodeUUID *string) *NdmpNodeModifyCollectionParams {
	o.SetNodeUUID(nodeUUID)
	return o
}

// SetNodeUUID adds the nodeUuid to the ndmp node modify collection params
func (o *NdmpNodeModifyCollectionParams) SetNodeUUID(nodeUUID *string) {
	o.NodeUUID = nodeUUID
}

// WithReturnRecords adds the returnRecords to the ndmp node modify collection params
func (o *NdmpNodeModifyCollectionParams) WithReturnRecords(returnRecords *bool) *NdmpNodeModifyCollectionParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the ndmp node modify collection params
func (o *NdmpNodeModifyCollectionParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the ndmp node modify collection params
func (o *NdmpNodeModifyCollectionParams) WithReturnTimeout(returnTimeout *int64) *NdmpNodeModifyCollectionParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the ndmp node modify collection params
func (o *NdmpNodeModifyCollectionParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSerialRecords adds the serialRecords to the ndmp node modify collection params
func (o *NdmpNodeModifyCollectionParams) WithSerialRecords(serialRecords *bool) *NdmpNodeModifyCollectionParams {
	o.SetSerialRecords(serialRecords)
	return o
}

// SetSerialRecords adds the serialRecords to the ndmp node modify collection params
func (o *NdmpNodeModifyCollectionParams) SetSerialRecords(serialRecords *bool) {
	o.SerialRecords = serialRecords
}

// WithUser adds the user to the ndmp node modify collection params
func (o *NdmpNodeModifyCollectionParams) WithUser(user *string) *NdmpNodeModifyCollectionParams {
	o.SetUser(user)
	return o
}

// SetUser adds the user to the ndmp node modify collection params
func (o *NdmpNodeModifyCollectionParams) SetUser(user *string) {
	o.User = user
}

// WriteToRequest writes these params to a swagger request
func (o *NdmpNodeModifyCollectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AuthenticationTypes != nil {

		// query param authentication_types
		var qrAuthenticationTypes string

		if o.AuthenticationTypes != nil {
			qrAuthenticationTypes = *o.AuthenticationTypes
		}
		qAuthenticationTypes := qrAuthenticationTypes
		if qAuthenticationTypes != "" {

			if err := r.SetQueryParam("authentication_types", qAuthenticationTypes); err != nil {
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

	if o.Enabled != nil {

		// query param enabled
		var qrEnabled bool

		if o.Enabled != nil {
			qrEnabled = *o.Enabled
		}
		qEnabled := swag.FormatBool(qrEnabled)
		if qEnabled != "" {

			if err := r.SetQueryParam("enabled", qEnabled); err != nil {
				return err
			}
		}
	}
	if err := r.SetBodyParam(o.Info); err != nil {
		return err
	}

	if o.NodeName != nil {

		// query param node.name
		var qrNodeName string

		if o.NodeName != nil {
			qrNodeName = *o.NodeName
		}
		qNodeName := qrNodeName
		if qNodeName != "" {

			if err := r.SetQueryParam("node.name", qNodeName); err != nil {
				return err
			}
		}
	}

	if o.NodeUUID != nil {

		// query param node.uuid
		var qrNodeUUID string

		if o.NodeUUID != nil {
			qrNodeUUID = *o.NodeUUID
		}
		qNodeUUID := qrNodeUUID
		if qNodeUUID != "" {

			if err := r.SetQueryParam("node.uuid", qNodeUUID); err != nil {
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

	if o.User != nil {

		// query param user
		var qrUser string

		if o.User != nil {
			qrUser = *o.User
		}
		qUser := qrUser
		if qUser != "" {

			if err := r.SetQueryParam("user", qUser); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}