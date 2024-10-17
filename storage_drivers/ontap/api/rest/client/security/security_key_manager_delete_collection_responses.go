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

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// SecurityKeyManagerDeleteCollectionReader is a Reader for the SecurityKeyManagerDeleteCollection structure.
type SecurityKeyManagerDeleteCollectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecurityKeyManagerDeleteCollectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSecurityKeyManagerDeleteCollectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSecurityKeyManagerDeleteCollectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSecurityKeyManagerDeleteCollectionOK creates a SecurityKeyManagerDeleteCollectionOK with default headers values
func NewSecurityKeyManagerDeleteCollectionOK() *SecurityKeyManagerDeleteCollectionOK {
	return &SecurityKeyManagerDeleteCollectionOK{}
}

/*
SecurityKeyManagerDeleteCollectionOK describes a response with status code 200, with default header values.

OK
*/
type SecurityKeyManagerDeleteCollectionOK struct {
}

// IsSuccess returns true when this security key manager delete collection o k response has a 2xx status code
func (o *SecurityKeyManagerDeleteCollectionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this security key manager delete collection o k response has a 3xx status code
func (o *SecurityKeyManagerDeleteCollectionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this security key manager delete collection o k response has a 4xx status code
func (o *SecurityKeyManagerDeleteCollectionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this security key manager delete collection o k response has a 5xx status code
func (o *SecurityKeyManagerDeleteCollectionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this security key manager delete collection o k response a status code equal to that given
func (o *SecurityKeyManagerDeleteCollectionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the security key manager delete collection o k response
func (o *SecurityKeyManagerDeleteCollectionOK) Code() int {
	return 200
}

func (o *SecurityKeyManagerDeleteCollectionOK) Error() string {
	return fmt.Sprintf("[DELETE /security/key-managers][%d] securityKeyManagerDeleteCollectionOK", 200)
}

func (o *SecurityKeyManagerDeleteCollectionOK) String() string {
	return fmt.Sprintf("[DELETE /security/key-managers][%d] securityKeyManagerDeleteCollectionOK", 200)
}

func (o *SecurityKeyManagerDeleteCollectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSecurityKeyManagerDeleteCollectionDefault creates a SecurityKeyManagerDeleteCollectionDefault with default headers values
func NewSecurityKeyManagerDeleteCollectionDefault(code int) *SecurityKeyManagerDeleteCollectionDefault {
	return &SecurityKeyManagerDeleteCollectionDefault{
		_statusCode: code,
	}
}

/*
	SecurityKeyManagerDeleteCollectionDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 65536208 | Failed to delete the SVM Key ID. |
| 65536233 | Internal error. Deletion of km_wrapped_kdb key database has failed for the Onboard Key Manager. |
| 65536234 | Internal error. Deletion of cluster_kdb key database has failed for the Onboard Key Manager. |
| 65536239 | Encrypted volumes are found for the SVM. |
| 65536242 | One or more self-encrypting drives are assigned an authentication key. |
| 65536243 | Cannot determine authentication key presence on one or more self-encrypting drives. |
| 65536800 | Failed to lookup onboard keys. |
| 65536813 | Encrypted kernel core files found. |
| 65536817 | Failed to determine if key manager is safe to disable. |
| 65536827 | Failed to determine if the SVM has any encrypted volumes. |
| 65536828 | External key management is not enabled for the SVM. |
| 65536867 | Encrypted volumes are found for the SVM. |
| 196608301 | Failed to determine the type of encryption. |
| 196608305 | NAE aggregates are found in the cluster. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
  - name: KEYMANAGER_MESSAGE_ERR_KM_DISABLE_ENC_CORE_CHECK_TIMEOUT
    message: Failed to disable the key manager because of a timeout when checking for encrypted cores.
*/
type SecurityKeyManagerDeleteCollectionDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this security key manager delete collection default response has a 2xx status code
func (o *SecurityKeyManagerDeleteCollectionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this security key manager delete collection default response has a 3xx status code
func (o *SecurityKeyManagerDeleteCollectionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this security key manager delete collection default response has a 4xx status code
func (o *SecurityKeyManagerDeleteCollectionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this security key manager delete collection default response has a 5xx status code
func (o *SecurityKeyManagerDeleteCollectionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this security key manager delete collection default response a status code equal to that given
func (o *SecurityKeyManagerDeleteCollectionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the security key manager delete collection default response
func (o *SecurityKeyManagerDeleteCollectionDefault) Code() int {
	return o._statusCode
}

func (o *SecurityKeyManagerDeleteCollectionDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /security/key-managers][%d] security_key_manager_delete_collection default %s", o._statusCode, payload)
}

func (o *SecurityKeyManagerDeleteCollectionDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /security/key-managers][%d] security_key_manager_delete_collection default %s", o._statusCode, payload)
}

func (o *SecurityKeyManagerDeleteCollectionDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SecurityKeyManagerDeleteCollectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
SecurityKeyManagerDeleteCollectionBody security key manager delete collection body
swagger:model SecurityKeyManagerDeleteCollectionBody
*/
type SecurityKeyManagerDeleteCollectionBody struct {

	// security key manager response inline records
	SecurityKeyManagerResponseInlineRecords []*models.SecurityKeyManager `json:"records,omitempty"`
}

// Validate validates this security key manager delete collection body
func (o *SecurityKeyManagerDeleteCollectionBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSecurityKeyManagerResponseInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SecurityKeyManagerDeleteCollectionBody) validateSecurityKeyManagerResponseInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(o.SecurityKeyManagerResponseInlineRecords) { // not required
		return nil
	}

	for i := 0; i < len(o.SecurityKeyManagerResponseInlineRecords); i++ {
		if swag.IsZero(o.SecurityKeyManagerResponseInlineRecords[i]) { // not required
			continue
		}

		if o.SecurityKeyManagerResponseInlineRecords[i] != nil {
			if err := o.SecurityKeyManagerResponseInlineRecords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this security key manager delete collection body based on the context it is used
func (o *SecurityKeyManagerDeleteCollectionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSecurityKeyManagerResponseInlineRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SecurityKeyManagerDeleteCollectionBody) contextValidateSecurityKeyManagerResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.SecurityKeyManagerResponseInlineRecords); i++ {

		if o.SecurityKeyManagerResponseInlineRecords[i] != nil {
			if err := o.SecurityKeyManagerResponseInlineRecords[i].ContextValidate(ctx, formats); err != nil {
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
func (o *SecurityKeyManagerDeleteCollectionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SecurityKeyManagerDeleteCollectionBody) UnmarshalBinary(b []byte) error {
	var res SecurityKeyManagerDeleteCollectionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}