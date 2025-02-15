// Code generated by go-swagger; DO NOT EDIT.

package s_a_n

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

// LunFormDataDeleteReader is a Reader for the LunFormDataDelete structure.
type LunFormDataDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LunFormDataDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLunFormDataDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewLunFormDataDeleteAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLunFormDataDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLunFormDataDeleteOK creates a LunFormDataDeleteOK with default headers values
func NewLunFormDataDeleteOK() *LunFormDataDeleteOK {
	return &LunFormDataDeleteOK{}
}

/*
LunFormDataDeleteOK describes a response with status code 200, with default header values.

OK
*/
type LunFormDataDeleteOK struct {
}

// IsSuccess returns true when this lun form data delete o k response has a 2xx status code
func (o *LunFormDataDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this lun form data delete o k response has a 3xx status code
func (o *LunFormDataDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this lun form data delete o k response has a 4xx status code
func (o *LunFormDataDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this lun form data delete o k response has a 5xx status code
func (o *LunFormDataDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this lun form data delete o k response a status code equal to that given
func (o *LunFormDataDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the lun form data delete o k response
func (o *LunFormDataDeleteOK) Code() int {
	return 200
}

func (o *LunFormDataDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE //storage/luns/{uuid}][%d] lunFormDataDeleteOK", 200)
}

func (o *LunFormDataDeleteOK) String() string {
	return fmt.Sprintf("[DELETE //storage/luns/{uuid}][%d] lunFormDataDeleteOK", 200)
}

func (o *LunFormDataDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLunFormDataDeleteAccepted creates a LunFormDataDeleteAccepted with default headers values
func NewLunFormDataDeleteAccepted() *LunFormDataDeleteAccepted {
	return &LunFormDataDeleteAccepted{}
}

/*
LunFormDataDeleteAccepted describes a response with status code 202, with default header values.

Accepted
*/
type LunFormDataDeleteAccepted struct {
	Payload *models.LunJobLinkResponse
}

// IsSuccess returns true when this lun form data delete accepted response has a 2xx status code
func (o *LunFormDataDeleteAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this lun form data delete accepted response has a 3xx status code
func (o *LunFormDataDeleteAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this lun form data delete accepted response has a 4xx status code
func (o *LunFormDataDeleteAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this lun form data delete accepted response has a 5xx status code
func (o *LunFormDataDeleteAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this lun form data delete accepted response a status code equal to that given
func (o *LunFormDataDeleteAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the lun form data delete accepted response
func (o *LunFormDataDeleteAccepted) Code() int {
	return 202
}

func (o *LunFormDataDeleteAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE //storage/luns/{uuid}][%d] lunFormDataDeleteAccepted %s", 202, payload)
}

func (o *LunFormDataDeleteAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE //storage/luns/{uuid}][%d] lunFormDataDeleteAccepted %s", 202, payload)
}

func (o *LunFormDataDeleteAccepted) GetPayload() *models.LunJobLinkResponse {
	return o.Payload
}

func (o *LunFormDataDeleteAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LunJobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLunFormDataDeleteDefault creates a LunFormDataDeleteDefault with default headers values
func NewLunFormDataDeleteDefault(code int) *LunFormDataDeleteDefault {
	return &LunFormDataDeleteDefault{
		_statusCode: code,
	}
}

/*
	LunFormDataDeleteDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 1254197 | The LUN is mapped and cannot be deleted without specifying the `allow_delete_while_mapped` query parameter. |
| 5374705 | Deleting the LUN is not allowed because it is part of an application. |
| 5374865 | The LUN's aggregate is offline. The aggregate must be online to modify or remove the LUN. |
| 5374866 | The LUN's volume is offline. The volume must be online to modify or remove the LUN. |
| 5374875 | The specified LUN was not found. |
| 5374876 | The specified LUN was not found. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type LunFormDataDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this lun form data delete default response has a 2xx status code
func (o *LunFormDataDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this lun form data delete default response has a 3xx status code
func (o *LunFormDataDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this lun form data delete default response has a 4xx status code
func (o *LunFormDataDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this lun form data delete default response has a 5xx status code
func (o *LunFormDataDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this lun form data delete default response a status code equal to that given
func (o *LunFormDataDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the lun form data delete default response
func (o *LunFormDataDeleteDefault) Code() int {
	return o._statusCode
}

func (o *LunFormDataDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE //storage/luns/{uuid}][%d] lun_form_data_delete default %s", o._statusCode, payload)
}

func (o *LunFormDataDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE //storage/luns/{uuid}][%d] lun_form_data_delete default %s", o._statusCode, payload)
}

func (o *LunFormDataDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *LunFormDataDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
