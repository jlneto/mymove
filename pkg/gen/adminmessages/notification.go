// Code generated by go-swagger; DO NOT EDIT.

package adminmessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Notification notification
// swagger:model Notification
type Notification struct {

	// created at
	// Required: true
	// Format: datetime
	CreatedAt *strfmt.DateTime `json:"created_at"`

	// email
	// Required: true
	Email *string `json:"email"`

	// id
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// notification type
	// Required: true
	// Enum: [MOVE_REVIEWED_EMAIL MOVE_PAYMENT_REMINDER_EMAIL]
	NotificationType *string `json:"notification_type"`

	// service member id
	// Required: true
	// Format: uuid
	ServiceMemberID *strfmt.UUID `json:"service_member_id"`

	// ses message id
	// Required: true
	SesMessageID *string `json:"ses_message_id"`
}

// Validate validates this notification
func (m *Notification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNotificationType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceMemberID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSesMessageID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Notification) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", m.CreatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("created_at", "body", "datetime", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Notification) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	return nil
}

func (m *Notification) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

var notificationTypeNotificationTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["MOVE_REVIEWED_EMAIL","MOVE_PAYMENT_REMINDER_EMAIL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		notificationTypeNotificationTypePropEnum = append(notificationTypeNotificationTypePropEnum, v)
	}
}

const (

	// NotificationNotificationTypeMOVEREVIEWEDEMAIL captures enum value "MOVE_REVIEWED_EMAIL"
	NotificationNotificationTypeMOVEREVIEWEDEMAIL string = "MOVE_REVIEWED_EMAIL"

	// NotificationNotificationTypeMOVEPAYMENTREMINDEREMAIL captures enum value "MOVE_PAYMENT_REMINDER_EMAIL"
	NotificationNotificationTypeMOVEPAYMENTREMINDEREMAIL string = "MOVE_PAYMENT_REMINDER_EMAIL"
)

// prop value enum
func (m *Notification) validateNotificationTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, notificationTypeNotificationTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Notification) validateNotificationType(formats strfmt.Registry) error {

	if err := validate.Required("notification_type", "body", m.NotificationType); err != nil {
		return err
	}

	// value enum
	if err := m.validateNotificationTypeEnum("notification_type", "body", *m.NotificationType); err != nil {
		return err
	}

	return nil
}

func (m *Notification) validateServiceMemberID(formats strfmt.Registry) error {

	if err := validate.Required("service_member_id", "body", m.ServiceMemberID); err != nil {
		return err
	}

	if err := validate.FormatOf("service_member_id", "body", "uuid", m.ServiceMemberID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Notification) validateSesMessageID(formats strfmt.Registry) error {

	if err := validate.Required("ses_message_id", "body", m.SesMessageID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Notification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Notification) UnmarshalBinary(b []byte) error {
	var res Notification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}