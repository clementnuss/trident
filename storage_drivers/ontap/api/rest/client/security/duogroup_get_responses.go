// Code generated by go-swagger; DO NOT EDIT.

package security

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

// DuogroupGetReader is a Reader for the DuogroupGet structure.
type DuogroupGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DuogroupGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDuogroupGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDuogroupGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDuogroupGetOK creates a DuogroupGetOK with default headers values
func NewDuogroupGetOK() *DuogroupGetOK {
	return &DuogroupGetOK{}
}

/*
DuogroupGetOK describes a response with status code 200, with default header values.

OK
*/
type DuogroupGetOK struct {
	Payload *models.Duogroup
}

// IsSuccess returns true when this duogroup get o k response has a 2xx status code
func (o *DuogroupGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this duogroup get o k response has a 3xx status code
func (o *DuogroupGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this duogroup get o k response has a 4xx status code
func (o *DuogroupGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this duogroup get o k response has a 5xx status code
func (o *DuogroupGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this duogroup get o k response a status code equal to that given
func (o *DuogroupGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the duogroup get o k response
func (o *DuogroupGetOK) Code() int {
	return 200
}

func (o *DuogroupGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/authentication/duo/groups/{owner.uuid}/{name}][%d] duogroupGetOK %s", 200, payload)
}

func (o *DuogroupGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/authentication/duo/groups/{owner.uuid}/{name}][%d] duogroupGetOK %s", 200, payload)
}

func (o *DuogroupGetOK) GetPayload() *models.Duogroup {
	return o.Payload
}

func (o *DuogroupGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Duogroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDuogroupGetDefault creates a DuogroupGetDefault with default headers values
func NewDuogroupGetDefault(code int) *DuogroupGetDefault {
	return &DuogroupGetDefault{
		_statusCode: code,
	}
}

/*
DuogroupGetDefault describes a response with status code -1, with default header values.

Error
*/
type DuogroupGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this duogroup get default response has a 2xx status code
func (o *DuogroupGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this duogroup get default response has a 3xx status code
func (o *DuogroupGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this duogroup get default response has a 4xx status code
func (o *DuogroupGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this duogroup get default response has a 5xx status code
func (o *DuogroupGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this duogroup get default response a status code equal to that given
func (o *DuogroupGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the duogroup get default response
func (o *DuogroupGetDefault) Code() int {
	return o._statusCode
}

func (o *DuogroupGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/authentication/duo/groups/{owner.uuid}/{name}][%d] duogroup_get default %s", o._statusCode, payload)
}

func (o *DuogroupGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/authentication/duo/groups/{owner.uuid}/{name}][%d] duogroup_get default %s", o._statusCode, payload)
}

func (o *DuogroupGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DuogroupGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}