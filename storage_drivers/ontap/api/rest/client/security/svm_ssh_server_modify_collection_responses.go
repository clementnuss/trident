// Code generated by go-swagger; DO NOT EDIT.

package security

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

// SvmSSHServerModifyCollectionReader is a Reader for the SvmSSHServerModifyCollection structure.
type SvmSSHServerModifyCollectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SvmSSHServerModifyCollectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSvmSSHServerModifyCollectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSvmSSHServerModifyCollectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSvmSSHServerModifyCollectionOK creates a SvmSSHServerModifyCollectionOK with default headers values
func NewSvmSSHServerModifyCollectionOK() *SvmSSHServerModifyCollectionOK {
	return &SvmSSHServerModifyCollectionOK{}
}

/*
SvmSSHServerModifyCollectionOK describes a response with status code 200, with default header values.

OK
*/
type SvmSSHServerModifyCollectionOK struct {
}

// IsSuccess returns true when this svm Ssh server modify collection o k response has a 2xx status code
func (o *SvmSSHServerModifyCollectionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this svm Ssh server modify collection o k response has a 3xx status code
func (o *SvmSSHServerModifyCollectionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this svm Ssh server modify collection o k response has a 4xx status code
func (o *SvmSSHServerModifyCollectionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this svm Ssh server modify collection o k response has a 5xx status code
func (o *SvmSSHServerModifyCollectionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this svm Ssh server modify collection o k response a status code equal to that given
func (o *SvmSSHServerModifyCollectionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the svm Ssh server modify collection o k response
func (o *SvmSSHServerModifyCollectionOK) Code() int {
	return 200
}

func (o *SvmSSHServerModifyCollectionOK) Error() string {
	return fmt.Sprintf("[PATCH /security/ssh/svms][%d] svmSshServerModifyCollectionOK", 200)
}

func (o *SvmSSHServerModifyCollectionOK) String() string {
	return fmt.Sprintf("[PATCH /security/ssh/svms][%d] svmSshServerModifyCollectionOK", 200)
}

func (o *SvmSSHServerModifyCollectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSvmSSHServerModifyCollectionDefault creates a SvmSSHServerModifyCollectionDefault with default headers values
func NewSvmSSHServerModifyCollectionDefault(code int) *SvmSSHServerModifyCollectionDefault {
	return &SvmSSHServerModifyCollectionDefault{
		_statusCode: code,
	}
}

/*
	SvmSSHServerModifyCollectionDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 10682372 | There must be at least one key exchange algorithm associated with the SSH configuration. |
| 10682373 | There must be at least one cipher associated with the SSH configuration. |
| 10682375 | Failed to modify SSH key exchange algorithms. |
| 10682378 | Failed to modify SSH ciphers. |
| 10682399 | Key exchange algorithm not supported in FIPS-enabled mode. |
| 10682400 | Failed to modify SSH MAC algorithms. |
| 10682401 | MAC algorithm not supported in FIPS-enabled mode. |
| 10682403 | There must be at least one MAC algorithm with the SSH configuration. |
| 10682413 | Failed to modify maximum authentication retry attempts. |
| 10682418 | Cipher not supported in FIPS-enabled mode. |
| 10682420 | To modify the SSH configuration of the admin SVM, use the /api/security/ssh REST API. |
| 10682423 | There must be at least one host key algorithm associated with the SSH configuration. |
| 10682424 | Host key algorithm not supported in FIPS enabled mode. |
| 10682425 | Failed to modify Host key algorithms. |
| 10682426 | Failed to modify _ssh-rsa_ enabled status for publickey algorithms configuration. |
| 10682428 | Cipher not supported in FIPS enabled mode. |
| 10682429 | Adding 'diffie_hellman_group16_sha512' or 'diffie_hellman_group18_sha512' to the SSH key exchange algorithms list requires an effective cluster version of ONTAP 9.16.1 or later. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type SvmSSHServerModifyCollectionDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this svm ssh server modify collection default response has a 2xx status code
func (o *SvmSSHServerModifyCollectionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this svm ssh server modify collection default response has a 3xx status code
func (o *SvmSSHServerModifyCollectionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this svm ssh server modify collection default response has a 4xx status code
func (o *SvmSSHServerModifyCollectionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this svm ssh server modify collection default response has a 5xx status code
func (o *SvmSSHServerModifyCollectionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this svm ssh server modify collection default response a status code equal to that given
func (o *SvmSSHServerModifyCollectionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the svm ssh server modify collection default response
func (o *SvmSSHServerModifyCollectionDefault) Code() int {
	return o._statusCode
}

func (o *SvmSSHServerModifyCollectionDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /security/ssh/svms][%d] svm_ssh_server_modify_collection default %s", o._statusCode, payload)
}

func (o *SvmSSHServerModifyCollectionDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /security/ssh/svms][%d] svm_ssh_server_modify_collection default %s", o._statusCode, payload)
}

func (o *SvmSSHServerModifyCollectionDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SvmSSHServerModifyCollectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
SvmSSHServerModifyCollectionBody svm SSH server modify collection body
swagger:model SvmSSHServerModifyCollectionBody
*/
type SvmSSHServerModifyCollectionBody struct {

	// links
	Links *models.SvmSSHServerInlineLinks `json:"_links,omitempty"`

	// Enables or disables the _ssh-rsa_ signature scheme, which uses the SHA-1 hash algorithm, for RSA keys in public key algorithms. If this flag is _false_, older SSH implementations might fail to authenticate using RSA keys. This flag should be enabled only as a temporary measure until legacy SSH client implementations can be upgraded or reconfigured with another key type, for example: ECDSA.
	//
	IsRsaInPublickeyAlgorithmsEnabled *bool `json:"is_rsa_in_publickey_algorithms_enabled,omitempty"`

	// Maximum authentication retries allowed before closing the connection.
	// Maximum: 6
	// Minimum: 2
	MaxAuthenticationRetryCount *int64 `json:"max_authentication_retry_count,omitempty"`

	// svm
	Svm *models.SvmSSHServerInlineSvm `json:"svm,omitempty"`

	// Ciphers for encrypting the data.
	// Example: ["aes256_ctr","aes192_ctr","aes128_ctr"]
	SvmSSHServerInlineCiphers []*models.Cipher `json:"ciphers,omitempty"`

	// Host key algorithms. The host key algorithm 'ssh_ed25519' can be configured only in non-FIPS mode.
	// Example: ["ecdsa_sha2_nistp256","ssh_ed25519","ssh_rsa"]
	SvmSSHServerInlineHostKeyAlgorithms []*models.HostKeyAlgorithm `json:"host_key_algorithms,omitempty"`

	// Key exchange algorithms.
	// Example: ["diffie_hellman_group_exchange_sha256","ecdh_sha2_nistp256"]
	SvmSSHServerInlineKeyExchangeAlgorithms []*models.KeyExchangeAlgorithm `json:"key_exchange_algorithms,omitempty"`

	// MAC algorithms.
	// Example: ["hmac_sha2_512","hmac_sha2_512_etm"]
	SvmSSHServerInlineMacAlgorithms []*models.MacAlgorithm `json:"mac_algorithms,omitempty"`

	// svm ssh server response inline records
	SvmSSHServerResponseInlineRecords []*models.SvmSSHServer `json:"records,omitempty"`
}

// Validate validates this svm SSH server modify collection body
func (o *SvmSSHServerModifyCollectionBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMaxAuthenticationRetryCount(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSvm(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSvmSSHServerInlineCiphers(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSvmSSHServerInlineHostKeyAlgorithms(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSvmSSHServerInlineKeyExchangeAlgorithms(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSvmSSHServerInlineMacAlgorithms(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSvmSSHServerResponseInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SvmSSHServerModifyCollectionBody) validateLinks(formats strfmt.Registry) error {
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

func (o *SvmSSHServerModifyCollectionBody) validateMaxAuthenticationRetryCount(formats strfmt.Registry) error {
	if swag.IsZero(o.MaxAuthenticationRetryCount) { // not required
		return nil
	}

	if err := validate.MinimumInt("info"+"."+"max_authentication_retry_count", "body", *o.MaxAuthenticationRetryCount, 2, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("info"+"."+"max_authentication_retry_count", "body", *o.MaxAuthenticationRetryCount, 6, false); err != nil {
		return err
	}

	return nil
}

func (o *SvmSSHServerModifyCollectionBody) validateSvm(formats strfmt.Registry) error {
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

func (o *SvmSSHServerModifyCollectionBody) validateSvmSSHServerInlineCiphers(formats strfmt.Registry) error {
	if swag.IsZero(o.SvmSSHServerInlineCiphers) { // not required
		return nil
	}

	for i := 0; i < len(o.SvmSSHServerInlineCiphers); i++ {
		if swag.IsZero(o.SvmSSHServerInlineCiphers[i]) { // not required
			continue
		}

		if o.SvmSSHServerInlineCiphers[i] != nil {
			if err := o.SvmSSHServerInlineCiphers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "ciphers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *SvmSSHServerModifyCollectionBody) validateSvmSSHServerInlineHostKeyAlgorithms(formats strfmt.Registry) error {
	if swag.IsZero(o.SvmSSHServerInlineHostKeyAlgorithms) { // not required
		return nil
	}

	for i := 0; i < len(o.SvmSSHServerInlineHostKeyAlgorithms); i++ {
		if swag.IsZero(o.SvmSSHServerInlineHostKeyAlgorithms[i]) { // not required
			continue
		}

		if o.SvmSSHServerInlineHostKeyAlgorithms[i] != nil {
			if err := o.SvmSSHServerInlineHostKeyAlgorithms[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "host_key_algorithms" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *SvmSSHServerModifyCollectionBody) validateSvmSSHServerInlineKeyExchangeAlgorithms(formats strfmt.Registry) error {
	if swag.IsZero(o.SvmSSHServerInlineKeyExchangeAlgorithms) { // not required
		return nil
	}

	for i := 0; i < len(o.SvmSSHServerInlineKeyExchangeAlgorithms); i++ {
		if swag.IsZero(o.SvmSSHServerInlineKeyExchangeAlgorithms[i]) { // not required
			continue
		}

		if o.SvmSSHServerInlineKeyExchangeAlgorithms[i] != nil {
			if err := o.SvmSSHServerInlineKeyExchangeAlgorithms[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "key_exchange_algorithms" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *SvmSSHServerModifyCollectionBody) validateSvmSSHServerInlineMacAlgorithms(formats strfmt.Registry) error {
	if swag.IsZero(o.SvmSSHServerInlineMacAlgorithms) { // not required
		return nil
	}

	for i := 0; i < len(o.SvmSSHServerInlineMacAlgorithms); i++ {
		if swag.IsZero(o.SvmSSHServerInlineMacAlgorithms[i]) { // not required
			continue
		}

		if o.SvmSSHServerInlineMacAlgorithms[i] != nil {
			if err := o.SvmSSHServerInlineMacAlgorithms[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "mac_algorithms" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *SvmSSHServerModifyCollectionBody) validateSvmSSHServerResponseInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(o.SvmSSHServerResponseInlineRecords) { // not required
		return nil
	}

	for i := 0; i < len(o.SvmSSHServerResponseInlineRecords); i++ {
		if swag.IsZero(o.SvmSSHServerResponseInlineRecords[i]) { // not required
			continue
		}

		if o.SvmSSHServerResponseInlineRecords[i] != nil {
			if err := o.SvmSSHServerResponseInlineRecords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this svm SSH server modify collection body based on the context it is used
func (o *SvmSSHServerModifyCollectionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateSvm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateSvmSSHServerInlineCiphers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateSvmSSHServerInlineHostKeyAlgorithms(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateSvmSSHServerInlineKeyExchangeAlgorithms(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateSvmSSHServerInlineMacAlgorithms(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateSvmSSHServerResponseInlineRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SvmSSHServerModifyCollectionBody) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (o *SvmSSHServerModifyCollectionBody) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

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

func (o *SvmSSHServerModifyCollectionBody) contextValidateSvmSSHServerInlineCiphers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.SvmSSHServerInlineCiphers); i++ {

		if o.SvmSSHServerInlineCiphers[i] != nil {
			if err := o.SvmSSHServerInlineCiphers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "ciphers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *SvmSSHServerModifyCollectionBody) contextValidateSvmSSHServerInlineHostKeyAlgorithms(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.SvmSSHServerInlineHostKeyAlgorithms); i++ {

		if o.SvmSSHServerInlineHostKeyAlgorithms[i] != nil {
			if err := o.SvmSSHServerInlineHostKeyAlgorithms[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "host_key_algorithms" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *SvmSSHServerModifyCollectionBody) contextValidateSvmSSHServerInlineKeyExchangeAlgorithms(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.SvmSSHServerInlineKeyExchangeAlgorithms); i++ {

		if o.SvmSSHServerInlineKeyExchangeAlgorithms[i] != nil {
			if err := o.SvmSSHServerInlineKeyExchangeAlgorithms[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "key_exchange_algorithms" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *SvmSSHServerModifyCollectionBody) contextValidateSvmSSHServerInlineMacAlgorithms(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.SvmSSHServerInlineMacAlgorithms); i++ {

		if o.SvmSSHServerInlineMacAlgorithms[i] != nil {
			if err := o.SvmSSHServerInlineMacAlgorithms[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "mac_algorithms" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *SvmSSHServerModifyCollectionBody) contextValidateSvmSSHServerResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.SvmSSHServerResponseInlineRecords); i++ {

		if o.SvmSSHServerResponseInlineRecords[i] != nil {
			if err := o.SvmSSHServerResponseInlineRecords[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *SvmSSHServerModifyCollectionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SvmSSHServerModifyCollectionBody) UnmarshalBinary(b []byte) error {
	var res SvmSSHServerModifyCollectionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SvmSSHServerInlineLinks svm ssh server inline links
swagger:model svm_ssh_server_inline__links
*/
type SvmSSHServerInlineLinks struct {

	// self
	Self *models.Href `json:"self,omitempty"`
}

// Validate validates this svm ssh server inline links
func (o *SvmSSHServerInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SvmSSHServerInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this svm ssh server inline links based on the context it is used
func (o *SvmSSHServerInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SvmSSHServerInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (o *SvmSSHServerInlineLinks) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SvmSSHServerInlineLinks) UnmarshalBinary(b []byte) error {
	var res SvmSSHServerInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SvmSSHServerInlineSvm SVM name and UUID for which the SSH server is configured.
swagger:model svm_ssh_server_inline_svm
*/
type SvmSSHServerInlineSvm struct {

	// links
	Links *models.SvmSSHServerInlineSvmInlineLinks `json:"_links,omitempty"`

	// The name of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: svm1
	Name *string `json:"name,omitempty"`

	// The unique identifier of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this svm ssh server inline svm
func (o *SvmSSHServerInlineSvm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SvmSSHServerInlineSvm) validateLinks(formats strfmt.Registry) error {
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

// ContextValidate validate this svm ssh server inline svm based on the context it is used
func (o *SvmSSHServerInlineSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SvmSSHServerInlineSvm) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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
func (o *SvmSSHServerInlineSvm) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SvmSSHServerInlineSvm) UnmarshalBinary(b []byte) error {
	var res SvmSSHServerInlineSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SvmSSHServerInlineSvmInlineLinks svm ssh server inline svm inline links
swagger:model svm_ssh_server_inline_svm_inline__links
*/
type SvmSSHServerInlineSvmInlineLinks struct {

	// self
	Self *models.Href `json:"self,omitempty"`
}

// Validate validates this svm ssh server inline svm inline links
func (o *SvmSSHServerInlineSvmInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SvmSSHServerInlineSvmInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this svm ssh server inline svm inline links based on the context it is used
func (o *SvmSSHServerInlineSvmInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SvmSSHServerInlineSvmInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (o *SvmSSHServerInlineSvmInlineLinks) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SvmSSHServerInlineSvmInlineLinks) UnmarshalBinary(b []byte) error {
	var res SvmSSHServerInlineSvmInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
