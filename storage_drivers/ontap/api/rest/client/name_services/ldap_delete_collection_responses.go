// Code generated by go-swagger; DO NOT EDIT.

package name_services

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

// LdapDeleteCollectionReader is a Reader for the LdapDeleteCollection structure.
type LdapDeleteCollectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LdapDeleteCollectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLdapDeleteCollectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLdapDeleteCollectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLdapDeleteCollectionOK creates a LdapDeleteCollectionOK with default headers values
func NewLdapDeleteCollectionOK() *LdapDeleteCollectionOK {
	return &LdapDeleteCollectionOK{}
}

/*
LdapDeleteCollectionOK describes a response with status code 200, with default header values.

OK
*/
type LdapDeleteCollectionOK struct {
}

// IsSuccess returns true when this ldap delete collection o k response has a 2xx status code
func (o *LdapDeleteCollectionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ldap delete collection o k response has a 3xx status code
func (o *LdapDeleteCollectionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ldap delete collection o k response has a 4xx status code
func (o *LdapDeleteCollectionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ldap delete collection o k response has a 5xx status code
func (o *LdapDeleteCollectionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ldap delete collection o k response a status code equal to that given
func (o *LdapDeleteCollectionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ldap delete collection o k response
func (o *LdapDeleteCollectionOK) Code() int {
	return 200
}

func (o *LdapDeleteCollectionOK) Error() string {
	return fmt.Sprintf("[DELETE /name-services/ldap][%d] ldapDeleteCollectionOK", 200)
}

func (o *LdapDeleteCollectionOK) String() string {
	return fmt.Sprintf("[DELETE /name-services/ldap][%d] ldapDeleteCollectionOK", 200)
}

func (o *LdapDeleteCollectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLdapDeleteCollectionDefault creates a LdapDeleteCollectionDefault with default headers values
func NewLdapDeleteCollectionDefault(code int) *LdapDeleteCollectionDefault {
	return &LdapDeleteCollectionDefault{
		_statusCode: code,
	}
}

/*
LdapDeleteCollectionDefault describes a response with status code -1, with default header values.

Error
*/
type LdapDeleteCollectionDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this ldap delete collection default response has a 2xx status code
func (o *LdapDeleteCollectionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ldap delete collection default response has a 3xx status code
func (o *LdapDeleteCollectionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ldap delete collection default response has a 4xx status code
func (o *LdapDeleteCollectionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ldap delete collection default response has a 5xx status code
func (o *LdapDeleteCollectionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ldap delete collection default response a status code equal to that given
func (o *LdapDeleteCollectionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ldap delete collection default response
func (o *LdapDeleteCollectionDefault) Code() int {
	return o._statusCode
}

func (o *LdapDeleteCollectionDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /name-services/ldap][%d] ldap_delete_collection default %s", o._statusCode, payload)
}

func (o *LdapDeleteCollectionDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /name-services/ldap][%d] ldap_delete_collection default %s", o._statusCode, payload)
}

func (o *LdapDeleteCollectionDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *LdapDeleteCollectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
LdapDeleteCollectionBody ldap delete collection body
swagger:model LdapDeleteCollectionBody
*/
type LdapDeleteCollectionBody struct {

	// ldap service response inline records
	LdapServiceResponseInlineRecords []*models.LdapService `json:"records,omitempty"`
}

// Validate validates this ldap delete collection body
func (o *LdapDeleteCollectionBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLdapServiceResponseInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *LdapDeleteCollectionBody) validateLdapServiceResponseInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(o.LdapServiceResponseInlineRecords) { // not required
		return nil
	}

	for i := 0; i < len(o.LdapServiceResponseInlineRecords); i++ {
		if swag.IsZero(o.LdapServiceResponseInlineRecords[i]) { // not required
			continue
		}

		if o.LdapServiceResponseInlineRecords[i] != nil {
			if err := o.LdapServiceResponseInlineRecords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this ldap delete collection body based on the context it is used
func (o *LdapDeleteCollectionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLdapServiceResponseInlineRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *LdapDeleteCollectionBody) contextValidateLdapServiceResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.LdapServiceResponseInlineRecords); i++ {

		if o.LdapServiceResponseInlineRecords[i] != nil {
			if err := o.LdapServiceResponseInlineRecords[i].ContextValidate(ctx, formats); err != nil {
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
func (o *LdapDeleteCollectionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LdapDeleteCollectionBody) UnmarshalBinary(b []byte) error {
	var res LdapDeleteCollectionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}