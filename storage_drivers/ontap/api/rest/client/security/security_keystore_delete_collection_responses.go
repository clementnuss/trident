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

// SecurityKeystoreDeleteCollectionReader is a Reader for the SecurityKeystoreDeleteCollection structure.
type SecurityKeystoreDeleteCollectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecurityKeystoreDeleteCollectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSecurityKeystoreDeleteCollectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSecurityKeystoreDeleteCollectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSecurityKeystoreDeleteCollectionOK creates a SecurityKeystoreDeleteCollectionOK with default headers values
func NewSecurityKeystoreDeleteCollectionOK() *SecurityKeystoreDeleteCollectionOK {
	return &SecurityKeystoreDeleteCollectionOK{}
}

/*
SecurityKeystoreDeleteCollectionOK describes a response with status code 200, with default header values.

OK
*/
type SecurityKeystoreDeleteCollectionOK struct {
}

// IsSuccess returns true when this security keystore delete collection o k response has a 2xx status code
func (o *SecurityKeystoreDeleteCollectionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this security keystore delete collection o k response has a 3xx status code
func (o *SecurityKeystoreDeleteCollectionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this security keystore delete collection o k response has a 4xx status code
func (o *SecurityKeystoreDeleteCollectionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this security keystore delete collection o k response has a 5xx status code
func (o *SecurityKeystoreDeleteCollectionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this security keystore delete collection o k response a status code equal to that given
func (o *SecurityKeystoreDeleteCollectionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the security keystore delete collection o k response
func (o *SecurityKeystoreDeleteCollectionOK) Code() int {
	return 200
}

func (o *SecurityKeystoreDeleteCollectionOK) Error() string {
	return fmt.Sprintf("[DELETE /security/key-stores][%d] securityKeystoreDeleteCollectionOK", 200)
}

func (o *SecurityKeystoreDeleteCollectionOK) String() string {
	return fmt.Sprintf("[DELETE /security/key-stores][%d] securityKeystoreDeleteCollectionOK", 200)
}

func (o *SecurityKeystoreDeleteCollectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSecurityKeystoreDeleteCollectionDefault creates a SecurityKeystoreDeleteCollectionDefault with default headers values
func NewSecurityKeystoreDeleteCollectionDefault(code int) *SecurityKeystoreDeleteCollectionDefault {
	return &SecurityKeystoreDeleteCollectionDefault{
		_statusCode: code,
	}
}

/*
	SecurityKeystoreDeleteCollectionDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 262155 | This operation requires an effective cluster version of 9.14.0 or later. |
| 65538905 | The keystore configuration is currently enabled and cannot be deleted. |
| 65538907 | The method is not yet supported for deleting the given UUID's type of configuration. |
| 65538908 | The specified keystore configuration UUID either does not exist or corresponds to a keystore configuration that is not supported by this operation. |
| 65539521 | An effective cluster version of ONTAP 9.16.1 or later is required to delete an inactive key manager on the admin SVM. |
| 65539522 | This command does not support disabling key manager configurations with the specified keystore type on the admin SVM. |
| 65539593 | Cannot delete an inactive key manager configuration while the keystore is in the switching state. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type SecurityKeystoreDeleteCollectionDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this security keystore delete collection default response has a 2xx status code
func (o *SecurityKeystoreDeleteCollectionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this security keystore delete collection default response has a 3xx status code
func (o *SecurityKeystoreDeleteCollectionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this security keystore delete collection default response has a 4xx status code
func (o *SecurityKeystoreDeleteCollectionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this security keystore delete collection default response has a 5xx status code
func (o *SecurityKeystoreDeleteCollectionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this security keystore delete collection default response a status code equal to that given
func (o *SecurityKeystoreDeleteCollectionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the security keystore delete collection default response
func (o *SecurityKeystoreDeleteCollectionDefault) Code() int {
	return o._statusCode
}

func (o *SecurityKeystoreDeleteCollectionDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /security/key-stores][%d] security_keystore_delete_collection default %s", o._statusCode, payload)
}

func (o *SecurityKeystoreDeleteCollectionDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /security/key-stores][%d] security_keystore_delete_collection default %s", o._statusCode, payload)
}

func (o *SecurityKeystoreDeleteCollectionDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SecurityKeystoreDeleteCollectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
SecurityKeystoreDeleteCollectionBody security keystore delete collection body
swagger:model SecurityKeystoreDeleteCollectionBody
*/
type SecurityKeystoreDeleteCollectionBody struct {

	// security keystore response inline records
	SecurityKeystoreResponseInlineRecords []*models.SecurityKeystore `json:"records,omitempty"`
}

// Validate validates this security keystore delete collection body
func (o *SecurityKeystoreDeleteCollectionBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSecurityKeystoreResponseInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SecurityKeystoreDeleteCollectionBody) validateSecurityKeystoreResponseInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(o.SecurityKeystoreResponseInlineRecords) { // not required
		return nil
	}

	for i := 0; i < len(o.SecurityKeystoreResponseInlineRecords); i++ {
		if swag.IsZero(o.SecurityKeystoreResponseInlineRecords[i]) { // not required
			continue
		}

		if o.SecurityKeystoreResponseInlineRecords[i] != nil {
			if err := o.SecurityKeystoreResponseInlineRecords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this security keystore delete collection body based on the context it is used
func (o *SecurityKeystoreDeleteCollectionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSecurityKeystoreResponseInlineRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SecurityKeystoreDeleteCollectionBody) contextValidateSecurityKeystoreResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.SecurityKeystoreResponseInlineRecords); i++ {

		if o.SecurityKeystoreResponseInlineRecords[i] != nil {
			if err := o.SecurityKeystoreResponseInlineRecords[i].ContextValidate(ctx, formats); err != nil {
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
func (o *SecurityKeystoreDeleteCollectionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SecurityKeystoreDeleteCollectionBody) UnmarshalBinary(b []byte) error {
	var res SecurityKeystoreDeleteCollectionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}