// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// KerberosInterfaceModifyCollectionReader is a Reader for the KerberosInterfaceModifyCollection structure.
type KerberosInterfaceModifyCollectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KerberosInterfaceModifyCollectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKerberosInterfaceModifyCollectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewKerberosInterfaceModifyCollectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewKerberosInterfaceModifyCollectionOK creates a KerberosInterfaceModifyCollectionOK with default headers values
func NewKerberosInterfaceModifyCollectionOK() *KerberosInterfaceModifyCollectionOK {
	return &KerberosInterfaceModifyCollectionOK{}
}

/*
KerberosInterfaceModifyCollectionOK describes a response with status code 200, with default header values.

OK
*/
type KerberosInterfaceModifyCollectionOK struct {
}

// IsSuccess returns true when this kerberos interface modify collection o k response has a 2xx status code
func (o *KerberosInterfaceModifyCollectionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kerberos interface modify collection o k response has a 3xx status code
func (o *KerberosInterfaceModifyCollectionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kerberos interface modify collection o k response has a 4xx status code
func (o *KerberosInterfaceModifyCollectionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kerberos interface modify collection o k response has a 5xx status code
func (o *KerberosInterfaceModifyCollectionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kerberos interface modify collection o k response a status code equal to that given
func (o *KerberosInterfaceModifyCollectionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the kerberos interface modify collection o k response
func (o *KerberosInterfaceModifyCollectionOK) Code() int {
	return 200
}

func (o *KerberosInterfaceModifyCollectionOK) Error() string {
	return fmt.Sprintf("[PATCH /protocols/nfs/kerberos/interfaces][%d] kerberosInterfaceModifyCollectionOK", 200)
}

func (o *KerberosInterfaceModifyCollectionOK) String() string {
	return fmt.Sprintf("[PATCH /protocols/nfs/kerberos/interfaces][%d] kerberosInterfaceModifyCollectionOK", 200)
}

func (o *KerberosInterfaceModifyCollectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewKerberosInterfaceModifyCollectionDefault creates a KerberosInterfaceModifyCollectionDefault with default headers values
func NewKerberosInterfaceModifyCollectionDefault(code int) *KerberosInterfaceModifyCollectionDefault {
	return &KerberosInterfaceModifyCollectionDefault{
		_statusCode: code,
	}
}

/*
	KerberosInterfaceModifyCollectionDefault describes a response with status code -1, with default header values.

	ONTAP Error Response codes

| Error codes | Description |
| ----------- | ----------- |
| 4           | The entry doesn't exist|
| 3276801     | Failed to bind service principal name on LIF.|
| 3276809     | Failed to disable NFS Kerberos on LIF.|
| 3276832     | Failed to insert Kerberos attributes to database.|
| 3276842     | Internal error. Failed to import Kerberos keytab file into the management databases. Contact technical support for assistance.|
| 3276861     | Kerberos is already enabled/disabled on this LIF.|
| 3276862     | Kerberos service principal name is required.|
| 3276889     | Failed to enable NFS Kerberos on LIF.|
| 3276937     | Failed to lookup the Vserver for the virtual interface.|
| 3276941     | Kerberos is a required field.|
| 3276942     | Service principal name is invalid. It must of the format:"nfs/<LIF-FQDN>@REALM"|
| 3276944     | Internal error. Reason: Failed to initialize the Kerberos context|
| 3276945     | Internal error. Reason: Failed to parse the service principal name|
| 3276951     | Warning: Skipping unsupported encryption type for service principal name|
| 3276952     | "organizational_unit" option cannot be used for "Other" vendor.|
| 3276965     | Account sharing across Vservers is not allowed. Use a different service principal name unique within the first 15 characters.|
| 3277019     | Cannot specify -force when enabling Kerberos.|
| 3277020     | Modifying the NFS Kerberos configuration for a LIF that is not configured for NFS is not supported.|
| 3277043     | Keytab import failed due to missing keys. Keys for encryption types are required for  Vserver but found no matching keys for service principal name. Generate the keytab file with all required keys and try again.|
*/
type KerberosInterfaceModifyCollectionDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this kerberos interface modify collection default response has a 2xx status code
func (o *KerberosInterfaceModifyCollectionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this kerberos interface modify collection default response has a 3xx status code
func (o *KerberosInterfaceModifyCollectionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this kerberos interface modify collection default response has a 4xx status code
func (o *KerberosInterfaceModifyCollectionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this kerberos interface modify collection default response has a 5xx status code
func (o *KerberosInterfaceModifyCollectionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this kerberos interface modify collection default response a status code equal to that given
func (o *KerberosInterfaceModifyCollectionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the kerberos interface modify collection default response
func (o *KerberosInterfaceModifyCollectionDefault) Code() int {
	return o._statusCode
}

func (o *KerberosInterfaceModifyCollectionDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/nfs/kerberos/interfaces][%d] kerberos_interface_modify_collection default %s", o._statusCode, payload)
}

func (o *KerberosInterfaceModifyCollectionDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/nfs/kerberos/interfaces][%d] kerberos_interface_modify_collection default %s", o._statusCode, payload)
}

func (o *KerberosInterfaceModifyCollectionDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *KerberosInterfaceModifyCollectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
KerberosInterfaceModifyCollectionBody kerberos interface modify collection body
swagger:model KerberosInterfaceModifyCollectionBody
*/
type KerberosInterfaceModifyCollectionBody struct {

	// links
	Links *models.KerberosInterfaceInlineLinks `json:"_links,omitempty"`

	// Specifies if Kerberos is enabled.
	Enabled *bool `json:"enabled,omitempty"`

	// encryption types
	// Read Only: true
	EncryptionTypes []*string `json:"encryption_types,omitempty"`

	// Specifies whether the server should ignore any error encountered while deleting the corresponding machine account on the KDC and also disables Kerberos on the LIF.
	Force *bool `json:"force,omitempty"`

	// interface
	Interface *models.KerberosInterfaceInlineInterface `json:"interface,omitempty"`

	// kerberos interface response inline records
	KerberosInterfaceResponseInlineRecords []*models.KerberosInterface `json:"records,omitempty"`

	// Load keytab from URI
	KeytabURI *string `json:"keytab_uri,omitempty"`

	// Specifies the machine account to create in Active Directory.
	MachineAccount *string `json:"machine_account,omitempty"`

	// Organizational unit
	OrganizationalUnit *string `json:"organizational_unit,omitempty"`

	// Account creation password
	Password *string `json:"password,omitempty"`

	// Service principal name. Valid in PATCH.
	Spn *string `json:"spn,omitempty"`

	// svm
	Svm *models.KerberosInterfaceInlineSvm `json:"svm,omitempty"`

	// Account creation user name
	User *string `json:"user,omitempty"`
}

// Validate validates this kerberos interface modify collection body
func (o *KerberosInterfaceModifyCollectionBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateEncryptionTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateInterface(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateKerberosInterfaceResponseInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSvm(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *KerberosInterfaceModifyCollectionBody) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(o.Links) { // not required
		return nil
	}

	if o.Links != nil {
		if err := o.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

var kerberosInterfaceModifyCollectionBodyEncryptionTypesItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["des","des3","aes_128","aes_256"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		kerberosInterfaceModifyCollectionBodyEncryptionTypesItemsEnum = append(kerberosInterfaceModifyCollectionBodyEncryptionTypesItemsEnum, v)
	}
}

func (o *KerberosInterfaceModifyCollectionBody) validateEncryptionTypesItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, kerberosInterfaceModifyCollectionBodyEncryptionTypesItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *KerberosInterfaceModifyCollectionBody) validateEncryptionTypes(formats strfmt.Registry) error {
	if swag.IsZero(o.EncryptionTypes) { // not required
		return nil
	}

	for i := 0; i < len(o.EncryptionTypes); i++ {
		if swag.IsZero(o.EncryptionTypes[i]) { // not required
			continue
		}

		// value enum
		if err := o.validateEncryptionTypesItemsEnum("info"+"."+"encryption_types"+"."+strconv.Itoa(i), "body", *o.EncryptionTypes[i]); err != nil {
			return err
		}

	}

	return nil
}

func (o *KerberosInterfaceModifyCollectionBody) validateInterface(formats strfmt.Registry) error {
	if swag.IsZero(o.Interface) { // not required
		return nil
	}

	if o.Interface != nil {
		if err := o.Interface.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "interface")
			}
			return err
		}
	}

	return nil
}

func (o *KerberosInterfaceModifyCollectionBody) validateKerberosInterfaceResponseInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(o.KerberosInterfaceResponseInlineRecords) { // not required
		return nil
	}

	for i := 0; i < len(o.KerberosInterfaceResponseInlineRecords); i++ {
		if swag.IsZero(o.KerberosInterfaceResponseInlineRecords[i]) { // not required
			continue
		}

		if o.KerberosInterfaceResponseInlineRecords[i] != nil {
			if err := o.KerberosInterfaceResponseInlineRecords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *KerberosInterfaceModifyCollectionBody) validateSvm(formats strfmt.Registry) error {
	if swag.IsZero(o.Svm) { // not required
		return nil
	}

	if o.Svm != nil {
		if err := o.Svm.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "svm")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this kerberos interface modify collection body based on the context it is used
func (o *KerberosInterfaceModifyCollectionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateEncryptionTypes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateInterface(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateKerberosInterfaceResponseInlineRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateSvm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *KerberosInterfaceModifyCollectionBody) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if o.Links != nil {
		if err := o.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

func (o *KerberosInterfaceModifyCollectionBody) contextValidateEncryptionTypes(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "info"+"."+"encryption_types", "body", []*string(o.EncryptionTypes)); err != nil {
		return err
	}

	for i := 0; i < len(o.EncryptionTypes); i++ {

		if err := validate.ReadOnly(ctx, "info"+"."+"encryption_types"+"."+strconv.Itoa(i), "body", o.EncryptionTypes[i]); err != nil {
			return err
		}

	}

	return nil
}

func (o *KerberosInterfaceModifyCollectionBody) contextValidateInterface(ctx context.Context, formats strfmt.Registry) error {

	if o.Interface != nil {
		if err := o.Interface.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "interface")
			}
			return err
		}
	}

	return nil
}

func (o *KerberosInterfaceModifyCollectionBody) contextValidateKerberosInterfaceResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.KerberosInterfaceResponseInlineRecords); i++ {

		if o.KerberosInterfaceResponseInlineRecords[i] != nil {
			if err := o.KerberosInterfaceResponseInlineRecords[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *KerberosInterfaceModifyCollectionBody) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

	if o.Svm != nil {
		if err := o.Svm.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "svm")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *KerberosInterfaceModifyCollectionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *KerberosInterfaceModifyCollectionBody) UnmarshalBinary(b []byte) error {
	var res KerberosInterfaceModifyCollectionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
KerberosInterfaceInlineLinks kerberos interface inline links
swagger:model kerberos_interface_inline__links
*/
type KerberosInterfaceInlineLinks struct {

	// self
	Self *models.Href `json:"self,omitempty"`
}

// Validate validates this kerberos interface inline links
func (o *KerberosInterfaceInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *KerberosInterfaceInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(o.Self) { // not required
		return nil
	}

	if o.Self != nil {
		if err := o.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this kerberos interface inline links based on the context it is used
func (o *KerberosInterfaceInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *KerberosInterfaceInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if o.Self != nil {
		if err := o.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *KerberosInterfaceInlineLinks) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *KerberosInterfaceInlineLinks) UnmarshalBinary(b []byte) error {
	var res KerberosInterfaceInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
KerberosInterfaceInlineInterface Network interface
swagger:model kerberos_interface_inline_interface
*/
type KerberosInterfaceInlineInterface struct {

	// links
	Links *models.KerberosInterfaceInlineInterfaceInlineLinks `json:"_links,omitempty"`

	// ip
	IP *models.KerberosInterfaceInlineInterfaceInlineIP `json:"ip,omitempty"`

	// The name of the interface. If only the name is provided, the SVM scope
	// must be provided by the object this object is embedded in.
	//
	// Example: lif1
	Name *string `json:"name,omitempty"`

	// The UUID that uniquely identifies the interface.
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this kerberos interface inline interface
func (o *KerberosInterfaceInlineInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateIP(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *KerberosInterfaceInlineInterface) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(o.Links) { // not required
		return nil
	}

	if o.Links != nil {
		if err := o.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "interface" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

func (o *KerberosInterfaceInlineInterface) validateIP(formats strfmt.Registry) error {
	if swag.IsZero(o.IP) { // not required
		return nil
	}

	if o.IP != nil {
		if err := o.IP.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "interface" + "." + "ip")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this kerberos interface inline interface based on the context it is used
func (o *KerberosInterfaceInlineInterface) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateIP(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *KerberosInterfaceInlineInterface) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if o.Links != nil {
		if err := o.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "interface" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

func (o *KerberosInterfaceInlineInterface) contextValidateIP(ctx context.Context, formats strfmt.Registry) error {

	if o.IP != nil {
		if err := o.IP.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "interface" + "." + "ip")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *KerberosInterfaceInlineInterface) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *KerberosInterfaceInlineInterface) UnmarshalBinary(b []byte) error {
	var res KerberosInterfaceInlineInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
KerberosInterfaceInlineInterfaceInlineLinks kerberos interface inline interface inline links
swagger:model kerberos_interface_inline_interface_inline__links
*/
type KerberosInterfaceInlineInterfaceInlineLinks struct {

	// self
	Self *models.Href `json:"self,omitempty"`
}

// Validate validates this kerberos interface inline interface inline links
func (o *KerberosInterfaceInlineInterfaceInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *KerberosInterfaceInlineInterfaceInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(o.Self) { // not required
		return nil
	}

	if o.Self != nil {
		if err := o.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "interface" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this kerberos interface inline interface inline links based on the context it is used
func (o *KerberosInterfaceInlineInterfaceInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *KerberosInterfaceInlineInterfaceInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if o.Self != nil {
		if err := o.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "interface" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *KerberosInterfaceInlineInterfaceInlineLinks) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *KerberosInterfaceInlineInterfaceInlineLinks) UnmarshalBinary(b []byte) error {
	var res KerberosInterfaceInlineInterfaceInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
KerberosInterfaceInlineInterfaceInlineIP IP information
swagger:model kerberos_interface_inline_interface_inline_ip
*/
type KerberosInterfaceInlineInterfaceInlineIP struct {

	// address
	Address *models.IPAddressReadonly `json:"address,omitempty"`
}

// Validate validates this kerberos interface inline interface inline ip
func (o *KerberosInterfaceInlineInterfaceInlineIP) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *KerberosInterfaceInlineInterfaceInlineIP) validateAddress(formats strfmt.Registry) error {
	if swag.IsZero(o.Address) { // not required
		return nil
	}

	if o.Address != nil {
		if err := o.Address.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "interface" + "." + "ip" + "." + "address")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this kerberos interface inline interface inline ip based on the context it is used
func (o *KerberosInterfaceInlineInterfaceInlineIP) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *KerberosInterfaceInlineInterfaceInlineIP) contextValidateAddress(ctx context.Context, formats strfmt.Registry) error {

	if o.Address != nil {
		if err := o.Address.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "interface" + "." + "ip" + "." + "address")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *KerberosInterfaceInlineInterfaceInlineIP) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *KerberosInterfaceInlineInterfaceInlineIP) UnmarshalBinary(b []byte) error {
	var res KerberosInterfaceInlineInterfaceInlineIP
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
KerberosInterfaceInlineSvm SVM, applies only to SVM-scoped objects.
swagger:model kerberos_interface_inline_svm
*/
type KerberosInterfaceInlineSvm struct {

	// links
	Links *models.KerberosInterfaceInlineSvmInlineLinks `json:"_links,omitempty"`

	// The name of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: svm1
	Name *string `json:"name,omitempty"`

	// The unique identifier of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this kerberos interface inline svm
func (o *KerberosInterfaceInlineSvm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *KerberosInterfaceInlineSvm) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(o.Links) { // not required
		return nil
	}

	if o.Links != nil {
		if err := o.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "svm" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this kerberos interface inline svm based on the context it is used
func (o *KerberosInterfaceInlineSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *KerberosInterfaceInlineSvm) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if o.Links != nil {
		if err := o.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "svm" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *KerberosInterfaceInlineSvm) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *KerberosInterfaceInlineSvm) UnmarshalBinary(b []byte) error {
	var res KerberosInterfaceInlineSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
KerberosInterfaceInlineSvmInlineLinks kerberos interface inline svm inline links
swagger:model kerberos_interface_inline_svm_inline__links
*/
type KerberosInterfaceInlineSvmInlineLinks struct {

	// self
	Self *models.Href `json:"self,omitempty"`
}

// Validate validates this kerberos interface inline svm inline links
func (o *KerberosInterfaceInlineSvmInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *KerberosInterfaceInlineSvmInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(o.Self) { // not required
		return nil
	}

	if o.Self != nil {
		if err := o.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "svm" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this kerberos interface inline svm inline links based on the context it is used
func (o *KerberosInterfaceInlineSvmInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *KerberosInterfaceInlineSvmInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if o.Self != nil {
		if err := o.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "svm" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *KerberosInterfaceInlineSvmInlineLinks) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *KerberosInterfaceInlineSvmInlineLinks) UnmarshalBinary(b []byte) error {
	var res KerberosInterfaceInlineSvmInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}