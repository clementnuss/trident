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

// NewPortsetDeleteCollectionParams creates a new PortsetDeleteCollectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPortsetDeleteCollectionParams() *PortsetDeleteCollectionParams {
	return &PortsetDeleteCollectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPortsetDeleteCollectionParamsWithTimeout creates a new PortsetDeleteCollectionParams object
// with the ability to set a timeout on a request.
func NewPortsetDeleteCollectionParamsWithTimeout(timeout time.Duration) *PortsetDeleteCollectionParams {
	return &PortsetDeleteCollectionParams{
		timeout: timeout,
	}
}

// NewPortsetDeleteCollectionParamsWithContext creates a new PortsetDeleteCollectionParams object
// with the ability to set a context for a request.
func NewPortsetDeleteCollectionParamsWithContext(ctx context.Context) *PortsetDeleteCollectionParams {
	return &PortsetDeleteCollectionParams{
		Context: ctx,
	}
}

// NewPortsetDeleteCollectionParamsWithHTTPClient creates a new PortsetDeleteCollectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewPortsetDeleteCollectionParamsWithHTTPClient(client *http.Client) *PortsetDeleteCollectionParams {
	return &PortsetDeleteCollectionParams{
		HTTPClient: client,
	}
}

/*
PortsetDeleteCollectionParams contains all the parameters to send to the API endpoint

	for the portset delete collection operation.

	Typically these are written to a http.Request.
*/
type PortsetDeleteCollectionParams struct {

	/* AllowDeleteWhileBound.

	     Allows deletion of a portset that is bound to an igroup.<br/>
	Deleting a portset can expand the set of LIFs through which a LUN is available.

	*/
	AllowDeleteWhileBound *bool

	/* ContinueOnFailure.

	   Continue even when the operation fails on one of the records.
	*/
	ContinueOnFailure *bool

	/* IgroupsName.

	   Filter by igroups.name
	*/
	IgroupsName *string

	/* IgroupsUUID.

	   Filter by igroups.uuid
	*/
	IgroupsUUID *string

	/* Info.

	   Info specification
	*/
	Info PortsetDeleteCollectionBody

	/* InterfacesFcName.

	   Filter by interfaces.fc.name
	*/
	InterfacesFcName *string

	/* InterfacesFcUUID.

	   Filter by interfaces.fc.uuid
	*/
	InterfacesFcUUID *string

	/* InterfacesFcWwpn.

	   Filter by interfaces.fc.wwpn
	*/
	InterfacesFcWwpn *string

	/* InterfacesIPIPAddress.

	   Filter by interfaces.ip.ip.address
	*/
	InterfacesIPIPAddress *string

	/* InterfacesIPName.

	   Filter by interfaces.ip.name
	*/
	InterfacesIPName *string

	/* InterfacesIPUUID.

	   Filter by interfaces.ip.uuid
	*/
	InterfacesIPUUID *string

	/* InterfacesUUID.

	   Filter by interfaces.uuid
	*/
	InterfacesUUID *string

	/* Name.

	   Filter by name
	*/
	Name *string

	/* Protocol.

	   Filter by protocol
	*/
	Protocol *string

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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the portset delete collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PortsetDeleteCollectionParams) WithDefaults() *PortsetDeleteCollectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the portset delete collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PortsetDeleteCollectionParams) SetDefaults() {
	var (
		allowDeleteWhileBoundDefault = bool(false)

		continueOnFailureDefault = bool(false)

		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)

		serialRecordsDefault = bool(false)
	)

	val := PortsetDeleteCollectionParams{
		AllowDeleteWhileBound: &allowDeleteWhileBoundDefault,
		ContinueOnFailure:     &continueOnFailureDefault,
		ReturnRecords:         &returnRecordsDefault,
		ReturnTimeout:         &returnTimeoutDefault,
		SerialRecords:         &serialRecordsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the portset delete collection params
func (o *PortsetDeleteCollectionParams) WithTimeout(timeout time.Duration) *PortsetDeleteCollectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the portset delete collection params
func (o *PortsetDeleteCollectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the portset delete collection params
func (o *PortsetDeleteCollectionParams) WithContext(ctx context.Context) *PortsetDeleteCollectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the portset delete collection params
func (o *PortsetDeleteCollectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the portset delete collection params
func (o *PortsetDeleteCollectionParams) WithHTTPClient(client *http.Client) *PortsetDeleteCollectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the portset delete collection params
func (o *PortsetDeleteCollectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAllowDeleteWhileBound adds the allowDeleteWhileBound to the portset delete collection params
func (o *PortsetDeleteCollectionParams) WithAllowDeleteWhileBound(allowDeleteWhileBound *bool) *PortsetDeleteCollectionParams {
	o.SetAllowDeleteWhileBound(allowDeleteWhileBound)
	return o
}

// SetAllowDeleteWhileBound adds the allowDeleteWhileBound to the portset delete collection params
func (o *PortsetDeleteCollectionParams) SetAllowDeleteWhileBound(allowDeleteWhileBound *bool) {
	o.AllowDeleteWhileBound = allowDeleteWhileBound
}

// WithContinueOnFailure adds the continueOnFailure to the portset delete collection params
func (o *PortsetDeleteCollectionParams) WithContinueOnFailure(continueOnFailure *bool) *PortsetDeleteCollectionParams {
	o.SetContinueOnFailure(continueOnFailure)
	return o
}

// SetContinueOnFailure adds the continueOnFailure to the portset delete collection params
func (o *PortsetDeleteCollectionParams) SetContinueOnFailure(continueOnFailure *bool) {
	o.ContinueOnFailure = continueOnFailure
}

// WithIgroupsName adds the igroupsName to the portset delete collection params
func (o *PortsetDeleteCollectionParams) WithIgroupsName(igroupsName *string) *PortsetDeleteCollectionParams {
	o.SetIgroupsName(igroupsName)
	return o
}

// SetIgroupsName adds the igroupsName to the portset delete collection params
func (o *PortsetDeleteCollectionParams) SetIgroupsName(igroupsName *string) {
	o.IgroupsName = igroupsName
}

// WithIgroupsUUID adds the igroupsUUID to the portset delete collection params
func (o *PortsetDeleteCollectionParams) WithIgroupsUUID(igroupsUUID *string) *PortsetDeleteCollectionParams {
	o.SetIgroupsUUID(igroupsUUID)
	return o
}

// SetIgroupsUUID adds the igroupsUuid to the portset delete collection params
func (o *PortsetDeleteCollectionParams) SetIgroupsUUID(igroupsUUID *string) {
	o.IgroupsUUID = igroupsUUID
}

// WithInfo adds the info to the portset delete collection params
func (o *PortsetDeleteCollectionParams) WithInfo(info PortsetDeleteCollectionBody) *PortsetDeleteCollectionParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the portset delete collection params
func (o *PortsetDeleteCollectionParams) SetInfo(info PortsetDeleteCollectionBody) {
	o.Info = info
}

// WithInterfacesFcName adds the interfacesFcName to the portset delete collection params
func (o *PortsetDeleteCollectionParams) WithInterfacesFcName(interfacesFcName *string) *PortsetDeleteCollectionParams {
	o.SetInterfacesFcName(interfacesFcName)
	return o
}

// SetInterfacesFcName adds the interfacesFcName to the portset delete collection params
func (o *PortsetDeleteCollectionParams) SetInterfacesFcName(interfacesFcName *string) {
	o.InterfacesFcName = interfacesFcName
}

// WithInterfacesFcUUID adds the interfacesFcUUID to the portset delete collection params
func (o *PortsetDeleteCollectionParams) WithInterfacesFcUUID(interfacesFcUUID *string) *PortsetDeleteCollectionParams {
	o.SetInterfacesFcUUID(interfacesFcUUID)
	return o
}

// SetInterfacesFcUUID adds the interfacesFcUuid to the portset delete collection params
func (o *PortsetDeleteCollectionParams) SetInterfacesFcUUID(interfacesFcUUID *string) {
	o.InterfacesFcUUID = interfacesFcUUID
}

// WithInterfacesFcWwpn adds the interfacesFcWwpn to the portset delete collection params
func (o *PortsetDeleteCollectionParams) WithInterfacesFcWwpn(interfacesFcWwpn *string) *PortsetDeleteCollectionParams {
	o.SetInterfacesFcWwpn(interfacesFcWwpn)
	return o
}

// SetInterfacesFcWwpn adds the interfacesFcWwpn to the portset delete collection params
func (o *PortsetDeleteCollectionParams) SetInterfacesFcWwpn(interfacesFcWwpn *string) {
	o.InterfacesFcWwpn = interfacesFcWwpn
}

// WithInterfacesIPIPAddress adds the interfacesIPIPAddress to the portset delete collection params
func (o *PortsetDeleteCollectionParams) WithInterfacesIPIPAddress(interfacesIPIPAddress *string) *PortsetDeleteCollectionParams {
	o.SetInterfacesIPIPAddress(interfacesIPIPAddress)
	return o
}

// SetInterfacesIPIPAddress adds the interfacesIpIpAddress to the portset delete collection params
func (o *PortsetDeleteCollectionParams) SetInterfacesIPIPAddress(interfacesIPIPAddress *string) {
	o.InterfacesIPIPAddress = interfacesIPIPAddress
}

// WithInterfacesIPName adds the interfacesIPName to the portset delete collection params
func (o *PortsetDeleteCollectionParams) WithInterfacesIPName(interfacesIPName *string) *PortsetDeleteCollectionParams {
	o.SetInterfacesIPName(interfacesIPName)
	return o
}

// SetInterfacesIPName adds the interfacesIpName to the portset delete collection params
func (o *PortsetDeleteCollectionParams) SetInterfacesIPName(interfacesIPName *string) {
	o.InterfacesIPName = interfacesIPName
}

// WithInterfacesIPUUID adds the interfacesIPUUID to the portset delete collection params
func (o *PortsetDeleteCollectionParams) WithInterfacesIPUUID(interfacesIPUUID *string) *PortsetDeleteCollectionParams {
	o.SetInterfacesIPUUID(interfacesIPUUID)
	return o
}

// SetInterfacesIPUUID adds the interfacesIpUuid to the portset delete collection params
func (o *PortsetDeleteCollectionParams) SetInterfacesIPUUID(interfacesIPUUID *string) {
	o.InterfacesIPUUID = interfacesIPUUID
}

// WithInterfacesUUID adds the interfacesUUID to the portset delete collection params
func (o *PortsetDeleteCollectionParams) WithInterfacesUUID(interfacesUUID *string) *PortsetDeleteCollectionParams {
	o.SetInterfacesUUID(interfacesUUID)
	return o
}

// SetInterfacesUUID adds the interfacesUuid to the portset delete collection params
func (o *PortsetDeleteCollectionParams) SetInterfacesUUID(interfacesUUID *string) {
	o.InterfacesUUID = interfacesUUID
}

// WithName adds the name to the portset delete collection params
func (o *PortsetDeleteCollectionParams) WithName(name *string) *PortsetDeleteCollectionParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the portset delete collection params
func (o *PortsetDeleteCollectionParams) SetName(name *string) {
	o.Name = name
}

// WithProtocol adds the protocol to the portset delete collection params
func (o *PortsetDeleteCollectionParams) WithProtocol(protocol *string) *PortsetDeleteCollectionParams {
	o.SetProtocol(protocol)
	return o
}

// SetProtocol adds the protocol to the portset delete collection params
func (o *PortsetDeleteCollectionParams) SetProtocol(protocol *string) {
	o.Protocol = protocol
}

// WithReturnRecords adds the returnRecords to the portset delete collection params
func (o *PortsetDeleteCollectionParams) WithReturnRecords(returnRecords *bool) *PortsetDeleteCollectionParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the portset delete collection params
func (o *PortsetDeleteCollectionParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the portset delete collection params
func (o *PortsetDeleteCollectionParams) WithReturnTimeout(returnTimeout *int64) *PortsetDeleteCollectionParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the portset delete collection params
func (o *PortsetDeleteCollectionParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSerialRecords adds the serialRecords to the portset delete collection params
func (o *PortsetDeleteCollectionParams) WithSerialRecords(serialRecords *bool) *PortsetDeleteCollectionParams {
	o.SetSerialRecords(serialRecords)
	return o
}

// SetSerialRecords adds the serialRecords to the portset delete collection params
func (o *PortsetDeleteCollectionParams) SetSerialRecords(serialRecords *bool) {
	o.SerialRecords = serialRecords
}

// WithSvmName adds the svmName to the portset delete collection params
func (o *PortsetDeleteCollectionParams) WithSvmName(svmName *string) *PortsetDeleteCollectionParams {
	o.SetSvmName(svmName)
	return o
}

// SetSvmName adds the svmName to the portset delete collection params
func (o *PortsetDeleteCollectionParams) SetSvmName(svmName *string) {
	o.SvmName = svmName
}

// WithSvmUUID adds the svmUUID to the portset delete collection params
func (o *PortsetDeleteCollectionParams) WithSvmUUID(svmUUID *string) *PortsetDeleteCollectionParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the portset delete collection params
func (o *PortsetDeleteCollectionParams) SetSvmUUID(svmUUID *string) {
	o.SvmUUID = svmUUID
}

// WithUUID adds the uuid to the portset delete collection params
func (o *PortsetDeleteCollectionParams) WithUUID(uuid *string) *PortsetDeleteCollectionParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the portset delete collection params
func (o *PortsetDeleteCollectionParams) SetUUID(uuid *string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *PortsetDeleteCollectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AllowDeleteWhileBound != nil {

		// query param allow_delete_while_bound
		var qrAllowDeleteWhileBound bool

		if o.AllowDeleteWhileBound != nil {
			qrAllowDeleteWhileBound = *o.AllowDeleteWhileBound
		}
		qAllowDeleteWhileBound := swag.FormatBool(qrAllowDeleteWhileBound)
		if qAllowDeleteWhileBound != "" {

			if err := r.SetQueryParam("allow_delete_while_bound", qAllowDeleteWhileBound); err != nil {
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

	if o.IgroupsName != nil {

		// query param igroups.name
		var qrIgroupsName string

		if o.IgroupsName != nil {
			qrIgroupsName = *o.IgroupsName
		}
		qIgroupsName := qrIgroupsName
		if qIgroupsName != "" {

			if err := r.SetQueryParam("igroups.name", qIgroupsName); err != nil {
				return err
			}
		}
	}

	if o.IgroupsUUID != nil {

		// query param igroups.uuid
		var qrIgroupsUUID string

		if o.IgroupsUUID != nil {
			qrIgroupsUUID = *o.IgroupsUUID
		}
		qIgroupsUUID := qrIgroupsUUID
		if qIgroupsUUID != "" {

			if err := r.SetQueryParam("igroups.uuid", qIgroupsUUID); err != nil {
				return err
			}
		}
	}
	if err := r.SetBodyParam(o.Info); err != nil {
		return err
	}

	if o.InterfacesFcName != nil {

		// query param interfaces.fc.name
		var qrInterfacesFcName string

		if o.InterfacesFcName != nil {
			qrInterfacesFcName = *o.InterfacesFcName
		}
		qInterfacesFcName := qrInterfacesFcName
		if qInterfacesFcName != "" {

			if err := r.SetQueryParam("interfaces.fc.name", qInterfacesFcName); err != nil {
				return err
			}
		}
	}

	if o.InterfacesFcUUID != nil {

		// query param interfaces.fc.uuid
		var qrInterfacesFcUUID string

		if o.InterfacesFcUUID != nil {
			qrInterfacesFcUUID = *o.InterfacesFcUUID
		}
		qInterfacesFcUUID := qrInterfacesFcUUID
		if qInterfacesFcUUID != "" {

			if err := r.SetQueryParam("interfaces.fc.uuid", qInterfacesFcUUID); err != nil {
				return err
			}
		}
	}

	if o.InterfacesFcWwpn != nil {

		// query param interfaces.fc.wwpn
		var qrInterfacesFcWwpn string

		if o.InterfacesFcWwpn != nil {
			qrInterfacesFcWwpn = *o.InterfacesFcWwpn
		}
		qInterfacesFcWwpn := qrInterfacesFcWwpn
		if qInterfacesFcWwpn != "" {

			if err := r.SetQueryParam("interfaces.fc.wwpn", qInterfacesFcWwpn); err != nil {
				return err
			}
		}
	}

	if o.InterfacesIPIPAddress != nil {

		// query param interfaces.ip.ip.address
		var qrInterfacesIPIPAddress string

		if o.InterfacesIPIPAddress != nil {
			qrInterfacesIPIPAddress = *o.InterfacesIPIPAddress
		}
		qInterfacesIPIPAddress := qrInterfacesIPIPAddress
		if qInterfacesIPIPAddress != "" {

			if err := r.SetQueryParam("interfaces.ip.ip.address", qInterfacesIPIPAddress); err != nil {
				return err
			}
		}
	}

	if o.InterfacesIPName != nil {

		// query param interfaces.ip.name
		var qrInterfacesIPName string

		if o.InterfacesIPName != nil {
			qrInterfacesIPName = *o.InterfacesIPName
		}
		qInterfacesIPName := qrInterfacesIPName
		if qInterfacesIPName != "" {

			if err := r.SetQueryParam("interfaces.ip.name", qInterfacesIPName); err != nil {
				return err
			}
		}
	}

	if o.InterfacesIPUUID != nil {

		// query param interfaces.ip.uuid
		var qrInterfacesIPUUID string

		if o.InterfacesIPUUID != nil {
			qrInterfacesIPUUID = *o.InterfacesIPUUID
		}
		qInterfacesIPUUID := qrInterfacesIPUUID
		if qInterfacesIPUUID != "" {

			if err := r.SetQueryParam("interfaces.ip.uuid", qInterfacesIPUUID); err != nil {
				return err
			}
		}
	}

	if o.InterfacesUUID != nil {

		// query param interfaces.uuid
		var qrInterfacesUUID string

		if o.InterfacesUUID != nil {
			qrInterfacesUUID = *o.InterfacesUUID
		}
		qInterfacesUUID := qrInterfacesUUID
		if qInterfacesUUID != "" {

			if err := r.SetQueryParam("interfaces.uuid", qInterfacesUUID); err != nil {
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

	if o.Protocol != nil {

		// query param protocol
		var qrProtocol string

		if o.Protocol != nil {
			qrProtocol = *o.Protocol
		}
		qProtocol := qrProtocol
		if qProtocol != "" {

			if err := r.SetQueryParam("protocol", qProtocol); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}