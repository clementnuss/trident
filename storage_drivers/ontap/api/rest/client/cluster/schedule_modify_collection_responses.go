// Code generated by go-swagger; DO NOT EDIT.

package cluster

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

// ScheduleModifyCollectionReader is a Reader for the ScheduleModifyCollection structure.
type ScheduleModifyCollectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ScheduleModifyCollectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewScheduleModifyCollectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewScheduleModifyCollectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewScheduleModifyCollectionOK creates a ScheduleModifyCollectionOK with default headers values
func NewScheduleModifyCollectionOK() *ScheduleModifyCollectionOK {
	return &ScheduleModifyCollectionOK{}
}

/*
ScheduleModifyCollectionOK describes a response with status code 200, with default header values.

OK
*/
type ScheduleModifyCollectionOK struct {
}

// IsSuccess returns true when this schedule modify collection o k response has a 2xx status code
func (o *ScheduleModifyCollectionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this schedule modify collection o k response has a 3xx status code
func (o *ScheduleModifyCollectionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this schedule modify collection o k response has a 4xx status code
func (o *ScheduleModifyCollectionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this schedule modify collection o k response has a 5xx status code
func (o *ScheduleModifyCollectionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this schedule modify collection o k response a status code equal to that given
func (o *ScheduleModifyCollectionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the schedule modify collection o k response
func (o *ScheduleModifyCollectionOK) Code() int {
	return 200
}

func (o *ScheduleModifyCollectionOK) Error() string {
	return fmt.Sprintf("[PATCH /cluster/schedules][%d] scheduleModifyCollectionOK", 200)
}

func (o *ScheduleModifyCollectionOK) String() string {
	return fmt.Sprintf("[PATCH /cluster/schedules][%d] scheduleModifyCollectionOK", 200)
}

func (o *ScheduleModifyCollectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewScheduleModifyCollectionDefault creates a ScheduleModifyCollectionDefault with default headers values
func NewScheduleModifyCollectionDefault(code int) *ScheduleModifyCollectionDefault {
	return &ScheduleModifyCollectionDefault{
		_statusCode: code,
	}
}

/*
	ScheduleModifyCollectionDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 458788 | The schedule specified is not a valid schedule. |
| 459760 | The schedule specified is not a valid schedule. |
| 459761 | Schedule cannot be modified on this cluster because it is replicated from the remote cluster. |
| 460783 | As this is a MetroCluster configuration and the local cluster is waiting for switchback, changes to non-system schedules are not allowed. |
| 461785 | A cron schedule cannot have an interval field. |
| 461786 | An interval schedule cannot have a cron field. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type ScheduleModifyCollectionDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this schedule modify collection default response has a 2xx status code
func (o *ScheduleModifyCollectionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this schedule modify collection default response has a 3xx status code
func (o *ScheduleModifyCollectionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this schedule modify collection default response has a 4xx status code
func (o *ScheduleModifyCollectionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this schedule modify collection default response has a 5xx status code
func (o *ScheduleModifyCollectionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this schedule modify collection default response a status code equal to that given
func (o *ScheduleModifyCollectionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the schedule modify collection default response
func (o *ScheduleModifyCollectionDefault) Code() int {
	return o._statusCode
}

func (o *ScheduleModifyCollectionDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /cluster/schedules][%d] schedule_modify_collection default %s", o._statusCode, payload)
}

func (o *ScheduleModifyCollectionDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /cluster/schedules][%d] schedule_modify_collection default %s", o._statusCode, payload)
}

func (o *ScheduleModifyCollectionDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ScheduleModifyCollectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ScheduleModifyCollectionBody schedule modify collection body
swagger:model ScheduleModifyCollectionBody
*/
type ScheduleModifyCollectionBody struct {

	// links
	Links *models.ScheduleInlineLinks `json:"_links,omitempty"`

	// cluster
	Cluster *models.ScheduleInlineCluster `json:"cluster,omitempty"`

	// cron
	Cron *models.ScheduleInlineCron `json:"cron,omitempty"`

	// An ISO-8601 duration formatted string.
	// Example: P1DT2H3M4S
	Interval *string `json:"interval,omitempty"`

	// Schedule name. Required in the URL or POST body.
	// Max Length: 256
	// Min Length: 1
	Name *string `json:"name,omitempty"`

	// schedule response inline records
	ScheduleResponseInlineRecords []*models.Schedule `json:"records,omitempty"`

	// If the schedule is owned by a data SVM, then the scope is set to svm. Otherwise it will be set to cluster.
	// Read Only: true
	// Enum: ["cluster","svm"]
	Scope *string `json:"scope,omitempty"`

	// svm
	Svm *models.ScheduleInlineSvm `json:"svm,omitempty"`

	// Schedule type
	// Read Only: true
	// Enum: ["cron","interval"]
	Type *string `json:"type,omitempty"`

	// Job schedule UUID
	// Example: 4ea7a442-86d1-11e0-ae1c-123478563412
	// Read Only: true
	// Format: uuid
	UUID *strfmt.UUID `json:"uuid,omitempty"`
}

// Validate validates this schedule modify collection body
func (o *ScheduleModifyCollectionBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateCron(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateScheduleResponseInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateScope(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSvm(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ScheduleModifyCollectionBody) validateLinks(formats strfmt.Registry) error {
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

func (o *ScheduleModifyCollectionBody) validateCluster(formats strfmt.Registry) error {
	if swag.IsZero(o.Cluster) { // not required
		return nil
	}

	if o.Cluster != nil {
		if err := o.Cluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "cluster")
			}
			return err
		}
	}

	return nil
}

func (o *ScheduleModifyCollectionBody) validateCron(formats strfmt.Registry) error {
	if swag.IsZero(o.Cron) { // not required
		return nil
	}

	if o.Cron != nil {
		if err := o.Cron.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "cron")
			}
			return err
		}
	}

	return nil
}

func (o *ScheduleModifyCollectionBody) validateName(formats strfmt.Registry) error {
	if swag.IsZero(o.Name) { // not required
		return nil
	}

	if err := validate.MinLength("info"+"."+"name", "body", *o.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("info"+"."+"name", "body", *o.Name, 256); err != nil {
		return err
	}

	return nil
}

func (o *ScheduleModifyCollectionBody) validateScheduleResponseInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(o.ScheduleResponseInlineRecords) { // not required
		return nil
	}

	for i := 0; i < len(o.ScheduleResponseInlineRecords); i++ {
		if swag.IsZero(o.ScheduleResponseInlineRecords[i]) { // not required
			continue
		}

		if o.ScheduleResponseInlineRecords[i] != nil {
			if err := o.ScheduleResponseInlineRecords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var scheduleModifyCollectionBodyTypeScopePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["cluster","svm"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		scheduleModifyCollectionBodyTypeScopePropEnum = append(scheduleModifyCollectionBodyTypeScopePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// ScheduleModifyCollectionBody
	// ScheduleModifyCollectionBody
	// scope
	// Scope
	// cluster
	// END DEBUGGING
	// ScheduleModifyCollectionBodyScopeCluster captures enum value "cluster"
	ScheduleModifyCollectionBodyScopeCluster string = "cluster"

	// BEGIN DEBUGGING
	// ScheduleModifyCollectionBody
	// ScheduleModifyCollectionBody
	// scope
	// Scope
	// svm
	// END DEBUGGING
	// ScheduleModifyCollectionBodyScopeSvm captures enum value "svm"
	ScheduleModifyCollectionBodyScopeSvm string = "svm"
)

// prop value enum
func (o *ScheduleModifyCollectionBody) validateScopeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, scheduleModifyCollectionBodyTypeScopePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *ScheduleModifyCollectionBody) validateScope(formats strfmt.Registry) error {
	if swag.IsZero(o.Scope) { // not required
		return nil
	}

	// value enum
	if err := o.validateScopeEnum("info"+"."+"scope", "body", *o.Scope); err != nil {
		return err
	}

	return nil
}

func (o *ScheduleModifyCollectionBody) validateSvm(formats strfmt.Registry) error {
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

var scheduleModifyCollectionBodyTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["cron","interval"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		scheduleModifyCollectionBodyTypeTypePropEnum = append(scheduleModifyCollectionBodyTypeTypePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// ScheduleModifyCollectionBody
	// ScheduleModifyCollectionBody
	// type
	// Type
	// cron
	// END DEBUGGING
	// ScheduleModifyCollectionBodyTypeCron captures enum value "cron"
	ScheduleModifyCollectionBodyTypeCron string = "cron"

	// BEGIN DEBUGGING
	// ScheduleModifyCollectionBody
	// ScheduleModifyCollectionBody
	// type
	// Type
	// interval
	// END DEBUGGING
	// ScheduleModifyCollectionBodyTypeInterval captures enum value "interval"
	ScheduleModifyCollectionBodyTypeInterval string = "interval"
)

// prop value enum
func (o *ScheduleModifyCollectionBody) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, scheduleModifyCollectionBodyTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *ScheduleModifyCollectionBody) validateType(formats strfmt.Registry) error {
	if swag.IsZero(o.Type) { // not required
		return nil
	}

	// value enum
	if err := o.validateTypeEnum("info"+"."+"type", "body", *o.Type); err != nil {
		return err
	}

	return nil
}

func (o *ScheduleModifyCollectionBody) validateUUID(formats strfmt.Registry) error {
	if swag.IsZero(o.UUID) { // not required
		return nil
	}

	if err := validate.FormatOf("info"+"."+"uuid", "body", "uuid", o.UUID.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this schedule modify collection body based on the context it is used
func (o *ScheduleModifyCollectionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateCron(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateScheduleResponseInlineRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateScope(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateSvm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateUUID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ScheduleModifyCollectionBody) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (o *ScheduleModifyCollectionBody) contextValidateCluster(ctx context.Context, formats strfmt.Registry) error {

	if o.Cluster != nil {
		if err := o.Cluster.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "cluster")
			}
			return err
		}
	}

	return nil
}

func (o *ScheduleModifyCollectionBody) contextValidateCron(ctx context.Context, formats strfmt.Registry) error {

	if o.Cron != nil {
		if err := o.Cron.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "cron")
			}
			return err
		}
	}

	return nil
}

func (o *ScheduleModifyCollectionBody) contextValidateScheduleResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.ScheduleResponseInlineRecords); i++ {

		if o.ScheduleResponseInlineRecords[i] != nil {
			if err := o.ScheduleResponseInlineRecords[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *ScheduleModifyCollectionBody) contextValidateScope(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "info"+"."+"scope", "body", o.Scope); err != nil {
		return err
	}

	return nil
}

func (o *ScheduleModifyCollectionBody) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

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

func (o *ScheduleModifyCollectionBody) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "info"+"."+"type", "body", o.Type); err != nil {
		return err
	}

	return nil
}

func (o *ScheduleModifyCollectionBody) contextValidateUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "info"+"."+"uuid", "body", o.UUID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ScheduleModifyCollectionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ScheduleModifyCollectionBody) UnmarshalBinary(b []byte) error {
	var res ScheduleModifyCollectionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ScheduleInlineLinks schedule inline links
swagger:model schedule_inline__links
*/
type ScheduleInlineLinks struct {

	// self
	Self *models.Href `json:"self,omitempty"`
}

// Validate validates this schedule inline links
func (o *ScheduleInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ScheduleInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this schedule inline links based on the context it is used
func (o *ScheduleInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ScheduleInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (o *ScheduleInlineLinks) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ScheduleInlineLinks) UnmarshalBinary(b []byte) error {
	var res ScheduleInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ScheduleInlineCluster The cluster that owns the schedule. Defaults to the local cluster.
swagger:model schedule_inline_cluster
*/
type ScheduleInlineCluster struct {

	// Cluster name
	// Example: cluster1
	Name *string `json:"name,omitempty"`

	// Cluster UUID
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	// Format: uuid
	UUID *strfmt.UUID `json:"uuid,omitempty"`
}

// Validate validates this schedule inline cluster
func (o *ScheduleInlineCluster) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateUUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ScheduleInlineCluster) validateUUID(formats strfmt.Registry) error {
	if swag.IsZero(o.UUID) { // not required
		return nil
	}

	if err := validate.FormatOf("info"+"."+"cluster"+"."+"uuid", "body", "uuid", o.UUID.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this schedule inline cluster based on context it is used
func (o *ScheduleInlineCluster) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ScheduleInlineCluster) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ScheduleInlineCluster) UnmarshalBinary(b []byte) error {
	var res ScheduleInlineCluster
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ScheduleInlineCron Details for schedules of type cron.
swagger:model schedule_inline_cron
*/
type ScheduleInlineCron struct {

	// The days of the month the schedule runs. Leave empty for all.
	Days []*int64 `json:"days,omitempty"`

	// The hours of the day the schedule runs. Leave empty for all.
	Hours []*int64 `json:"hours,omitempty"`

	// The minutes the schedule runs. Required on POST for a cron schedule.
	Minutes []*int64 `json:"minutes,omitempty"`

	// The months of the year the schedule runs. Leave empty for all.
	Months []*int64 `json:"months,omitempty"`

	// The weekdays the schedule runs. Leave empty for all.
	Weekdays []*int64 `json:"weekdays,omitempty"`
}

// Validate validates this schedule inline cron
func (o *ScheduleInlineCron) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDays(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateHours(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMinutes(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMonths(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateWeekdays(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ScheduleInlineCron) validateDays(formats strfmt.Registry) error {
	if swag.IsZero(o.Days) { // not required
		return nil
	}

	for i := 0; i < len(o.Days); i++ {
		if swag.IsZero(o.Days[i]) { // not required
			continue
		}

		if err := validate.MinimumInt("info"+"."+"cron"+"."+"days"+"."+strconv.Itoa(i), "body", *o.Days[i], 1, false); err != nil {
			return err
		}

		if err := validate.MaximumInt("info"+"."+"cron"+"."+"days"+"."+strconv.Itoa(i), "body", *o.Days[i], 31, false); err != nil {
			return err
		}

	}

	return nil
}

func (o *ScheduleInlineCron) validateHours(formats strfmt.Registry) error {
	if swag.IsZero(o.Hours) { // not required
		return nil
	}

	for i := 0; i < len(o.Hours); i++ {
		if swag.IsZero(o.Hours[i]) { // not required
			continue
		}

		if err := validate.MinimumInt("info"+"."+"cron"+"."+"hours"+"."+strconv.Itoa(i), "body", *o.Hours[i], 0, false); err != nil {
			return err
		}

		if err := validate.MaximumInt("info"+"."+"cron"+"."+"hours"+"."+strconv.Itoa(i), "body", *o.Hours[i], 23, false); err != nil {
			return err
		}

	}

	return nil
}

func (o *ScheduleInlineCron) validateMinutes(formats strfmt.Registry) error {
	if swag.IsZero(o.Minutes) { // not required
		return nil
	}

	for i := 0; i < len(o.Minutes); i++ {
		if swag.IsZero(o.Minutes[i]) { // not required
			continue
		}

		if err := validate.MinimumInt("info"+"."+"cron"+"."+"minutes"+"."+strconv.Itoa(i), "body", *o.Minutes[i], 0, false); err != nil {
			return err
		}

		if err := validate.MaximumInt("info"+"."+"cron"+"."+"minutes"+"."+strconv.Itoa(i), "body", *o.Minutes[i], 59, false); err != nil {
			return err
		}

	}

	return nil
}

func (o *ScheduleInlineCron) validateMonths(formats strfmt.Registry) error {
	if swag.IsZero(o.Months) { // not required
		return nil
	}

	for i := 0; i < len(o.Months); i++ {
		if swag.IsZero(o.Months[i]) { // not required
			continue
		}

		if err := validate.MinimumInt("info"+"."+"cron"+"."+"months"+"."+strconv.Itoa(i), "body", *o.Months[i], 1, false); err != nil {
			return err
		}

		if err := validate.MaximumInt("info"+"."+"cron"+"."+"months"+"."+strconv.Itoa(i), "body", *o.Months[i], 12, false); err != nil {
			return err
		}

	}

	return nil
}

func (o *ScheduleInlineCron) validateWeekdays(formats strfmt.Registry) error {
	if swag.IsZero(o.Weekdays) { // not required
		return nil
	}

	for i := 0; i < len(o.Weekdays); i++ {
		if swag.IsZero(o.Weekdays[i]) { // not required
			continue
		}

		if err := validate.MinimumInt("info"+"."+"cron"+"."+"weekdays"+"."+strconv.Itoa(i), "body", *o.Weekdays[i], 0, false); err != nil {
			return err
		}

		if err := validate.MaximumInt("info"+"."+"cron"+"."+"weekdays"+"."+strconv.Itoa(i), "body", *o.Weekdays[i], 6, false); err != nil {
			return err
		}

	}

	return nil
}

// ContextValidate validates this schedule inline cron based on context it is used
func (o *ScheduleInlineCron) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ScheduleInlineCron) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ScheduleInlineCron) UnmarshalBinary(b []byte) error {
	var res ScheduleInlineCron
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ScheduleInlineSvm SVM, applies only to SVM-scoped objects.
swagger:model schedule_inline_svm
*/
type ScheduleInlineSvm struct {

	// links
	Links *models.ScheduleInlineSvmInlineLinks `json:"_links,omitempty"`

	// The name of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: svm1
	Name *string `json:"name,omitempty"`

	// The unique identifier of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this schedule inline svm
func (o *ScheduleInlineSvm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ScheduleInlineSvm) validateLinks(formats strfmt.Registry) error {
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

// ContextValidate validate this schedule inline svm based on the context it is used
func (o *ScheduleInlineSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ScheduleInlineSvm) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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
func (o *ScheduleInlineSvm) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ScheduleInlineSvm) UnmarshalBinary(b []byte) error {
	var res ScheduleInlineSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ScheduleInlineSvmInlineLinks schedule inline svm inline links
swagger:model schedule_inline_svm_inline__links
*/
type ScheduleInlineSvmInlineLinks struct {

	// self
	Self *models.Href `json:"self,omitempty"`
}

// Validate validates this schedule inline svm inline links
func (o *ScheduleInlineSvmInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ScheduleInlineSvmInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this schedule inline svm inline links based on the context it is used
func (o *ScheduleInlineSvmInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ScheduleInlineSvmInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (o *ScheduleInlineSvmInlineLinks) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ScheduleInlineSvmInlineLinks) UnmarshalBinary(b []byte) error {
	var res ScheduleInlineSvmInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}