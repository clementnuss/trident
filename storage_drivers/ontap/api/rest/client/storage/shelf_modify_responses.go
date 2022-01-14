// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// ShelfModifyReader is a Reader for the ShelfModify structure.
type ShelfModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ShelfModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewShelfModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewShelfModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewShelfModifyOK creates a ShelfModifyOK with default headers values
func NewShelfModifyOK() *ShelfModifyOK {
	return &ShelfModifyOK{}
}

/* ShelfModifyOK describes a response with status code 200, with default header values.

OK
*/
type ShelfModifyOK struct {
}

func (o *ShelfModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /storage/shelves/{uid}][%d] shelfModifyOK ", 200)
}

func (o *ShelfModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewShelfModifyDefault creates a ShelfModifyDefault with default headers values
func NewShelfModifyDefault(code int) *ShelfModifyDefault {
	return &ShelfModifyDefault{
		_statusCode: code,
	}
}

/* ShelfModifyDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 17825872 | Shelf locate request failed because shelf \"<name>\" was not found. |
| 17825873 | Shelf locate request failed because shelf \"<name>\" does not support this command. |
| 17825874 | Shelf locate request failed for shelf \"<name>\" with an unknown error. |
| 17825875 | Shelf locate request failed for shelf \"<name>\" because shelf modules are unreachable. |

*/
type ShelfModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the shelf modify default response
func (o *ShelfModifyDefault) Code() int {
	return o._statusCode
}

func (o *ShelfModifyDefault) Error() string {
	return fmt.Sprintf("[PATCH /storage/shelves/{uid}][%d] shelf_modify default  %+v", o._statusCode, o.Payload)
}
func (o *ShelfModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ShelfModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}