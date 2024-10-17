// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// NfsTLSInterfaceGetReader is a Reader for the NfsTLSInterfaceGet structure.
type NfsTLSInterfaceGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NfsTLSInterfaceGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNfsTLSInterfaceGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNfsTLSInterfaceGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNfsTLSInterfaceGetOK creates a NfsTLSInterfaceGetOK with default headers values
func NewNfsTLSInterfaceGetOK() *NfsTLSInterfaceGetOK {
	return &NfsTLSInterfaceGetOK{}
}

/*
NfsTLSInterfaceGetOK describes a response with status code 200, with default header values.

OK
*/
type NfsTLSInterfaceGetOK struct {
	Payload *models.NfsTLSInterface
}

// IsSuccess returns true when this nfs Tls interface get o k response has a 2xx status code
func (o *NfsTLSInterfaceGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this nfs Tls interface get o k response has a 3xx status code
func (o *NfsTLSInterfaceGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this nfs Tls interface get o k response has a 4xx status code
func (o *NfsTLSInterfaceGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this nfs Tls interface get o k response has a 5xx status code
func (o *NfsTLSInterfaceGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this nfs Tls interface get o k response a status code equal to that given
func (o *NfsTLSInterfaceGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the nfs Tls interface get o k response
func (o *NfsTLSInterfaceGetOK) Code() int {
	return 200
}

func (o *NfsTLSInterfaceGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/nfs/tls/interfaces/{interface.uuid}][%d] nfsTlsInterfaceGetOK %s", 200, payload)
}

func (o *NfsTLSInterfaceGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/nfs/tls/interfaces/{interface.uuid}][%d] nfsTlsInterfaceGetOK %s", 200, payload)
}

func (o *NfsTLSInterfaceGetOK) GetPayload() *models.NfsTLSInterface {
	return o.Payload
}

func (o *NfsTLSInterfaceGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NfsTLSInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNfsTLSInterfaceGetDefault creates a NfsTLSInterfaceGetDefault with default headers values
func NewNfsTLSInterfaceGetDefault(code int) *NfsTLSInterfaceGetDefault {
	return &NfsTLSInterfaceGetDefault{
		_statusCode: code,
	}
}

/*
	NfsTLSInterfaceGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response codes

| Error codes | Description |
| ----------- | ----------- |
| 262197      | The value provided is invalid for field \"fields\".|
*/
type NfsTLSInterfaceGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this nfs tls interface get default response has a 2xx status code
func (o *NfsTLSInterfaceGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this nfs tls interface get default response has a 3xx status code
func (o *NfsTLSInterfaceGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this nfs tls interface get default response has a 4xx status code
func (o *NfsTLSInterfaceGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this nfs tls interface get default response has a 5xx status code
func (o *NfsTLSInterfaceGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this nfs tls interface get default response a status code equal to that given
func (o *NfsTLSInterfaceGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the nfs tls interface get default response
func (o *NfsTLSInterfaceGetDefault) Code() int {
	return o._statusCode
}

func (o *NfsTLSInterfaceGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/nfs/tls/interfaces/{interface.uuid}][%d] nfs_tls_interface_get default %s", o._statusCode, payload)
}

func (o *NfsTLSInterfaceGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/nfs/tls/interfaces/{interface.uuid}][%d] nfs_tls_interface_get default %s", o._statusCode, payload)
}

func (o *NfsTLSInterfaceGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NfsTLSInterfaceGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}