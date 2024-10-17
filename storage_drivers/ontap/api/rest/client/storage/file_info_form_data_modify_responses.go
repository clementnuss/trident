// Code generated by go-swagger; DO NOT EDIT.

package storage

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

// FileInfoFormDataModifyReader is a Reader for the FileInfoFormDataModify structure.
type FileInfoFormDataModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FileInfoFormDataModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFileInfoFormDataModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFileInfoFormDataModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFileInfoFormDataModifyOK creates a FileInfoFormDataModifyOK with default headers values
func NewFileInfoFormDataModifyOK() *FileInfoFormDataModifyOK {
	return &FileInfoFormDataModifyOK{}
}

/*
FileInfoFormDataModifyOK describes a response with status code 200, with default header values.

OK
*/
type FileInfoFormDataModifyOK struct {
}

// IsSuccess returns true when this file info form data modify o k response has a 2xx status code
func (o *FileInfoFormDataModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this file info form data modify o k response has a 3xx status code
func (o *FileInfoFormDataModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this file info form data modify o k response has a 4xx status code
func (o *FileInfoFormDataModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this file info form data modify o k response has a 5xx status code
func (o *FileInfoFormDataModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this file info form data modify o k response a status code equal to that given
func (o *FileInfoFormDataModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the file info form data modify o k response
func (o *FileInfoFormDataModifyOK) Code() int {
	return 200
}

func (o *FileInfoFormDataModifyOK) Error() string {
	return fmt.Sprintf("[PATCH //storage/volumes/{volume.uuid}/files/{path}][%d] fileInfoFormDataModifyOK", 200)
}

func (o *FileInfoFormDataModifyOK) String() string {
	return fmt.Sprintf("[PATCH //storage/volumes/{volume.uuid}/files/{path}][%d] fileInfoFormDataModifyOK", 200)
}

func (o *FileInfoFormDataModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFileInfoFormDataModifyDefault creates a FileInfoFormDataModifyDefault with default headers values
func NewFileInfoFormDataModifyDefault(code int) *FileInfoFormDataModifyDefault {
	return &FileInfoFormDataModifyDefault{
		_statusCode: code,
	}
}

/*
	FileInfoFormDataModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 918235 | A volume with UUID {volume.uuid} was not found. |
| 6488081 | The {field} field is not supported for PATCH operations. |
| 6488082 | Failed to rename {path}. |
| 6488083 | Failed to rename {path} to {path} because a directory named {path} already exists. |
| 6488111 | The fields 'fill_enabled' and 'overwrite_enabled' are not supported in the directory modify operation. |
| 6488117 | Permission denied. |
| 6488119 | Operation not supported. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type FileInfoFormDataModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this file info form data modify default response has a 2xx status code
func (o *FileInfoFormDataModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this file info form data modify default response has a 3xx status code
func (o *FileInfoFormDataModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this file info form data modify default response has a 4xx status code
func (o *FileInfoFormDataModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this file info form data modify default response has a 5xx status code
func (o *FileInfoFormDataModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this file info form data modify default response a status code equal to that given
func (o *FileInfoFormDataModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the file info form data modify default response
func (o *FileInfoFormDataModifyDefault) Code() int {
	return o._statusCode
}

func (o *FileInfoFormDataModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH //storage/volumes/{volume.uuid}/files/{path}][%d] file_info_form_data_modify default %s", o._statusCode, payload)
}

func (o *FileInfoFormDataModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH //storage/volumes/{volume.uuid}/files/{path}][%d] file_info_form_data_modify default %s", o._statusCode, payload)
}

func (o *FileInfoFormDataModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *FileInfoFormDataModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}