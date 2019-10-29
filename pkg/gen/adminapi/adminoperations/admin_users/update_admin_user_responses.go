// Code generated by go-swagger; DO NOT EDIT.

package admin_users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	adminmessages "github.com/transcom/mymove/pkg/gen/adminmessages"
)

// UpdateAdminUserOKCode is the HTTP code returned for type UpdateAdminUserOK
const UpdateAdminUserOKCode int = 200

/*UpdateAdminUserOK Successfully updated Admin User

swagger:response updateAdminUserOK
*/
type UpdateAdminUserOK struct {

	/*
	  In: Body
	*/
	Payload *adminmessages.AdminUser `json:"body,omitempty"`
}

// NewUpdateAdminUserOK creates UpdateAdminUserOK with default headers values
func NewUpdateAdminUserOK() *UpdateAdminUserOK {

	return &UpdateAdminUserOK{}
}

// WithPayload adds the payload to the update admin user o k response
func (o *UpdateAdminUserOK) WithPayload(payload *adminmessages.AdminUser) *UpdateAdminUserOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update admin user o k response
func (o *UpdateAdminUserOK) SetPayload(payload *adminmessages.AdminUser) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateAdminUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateAdminUserBadRequestCode is the HTTP code returned for type UpdateAdminUserBadRequest
const UpdateAdminUserBadRequestCode int = 400

/*UpdateAdminUserBadRequest Invalid Request

swagger:response updateAdminUserBadRequest
*/
type UpdateAdminUserBadRequest struct {
}

// NewUpdateAdminUserBadRequest creates UpdateAdminUserBadRequest with default headers values
func NewUpdateAdminUserBadRequest() *UpdateAdminUserBadRequest {

	return &UpdateAdminUserBadRequest{}
}

// WriteResponse to the client
func (o *UpdateAdminUserBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// UpdateAdminUserUnauthorizedCode is the HTTP code returned for type UpdateAdminUserUnauthorized
const UpdateAdminUserUnauthorizedCode int = 401

/*UpdateAdminUserUnauthorized Must be authenticated to use this end point

swagger:response updateAdminUserUnauthorized
*/
type UpdateAdminUserUnauthorized struct {
}

// NewUpdateAdminUserUnauthorized creates UpdateAdminUserUnauthorized with default headers values
func NewUpdateAdminUserUnauthorized() *UpdateAdminUserUnauthorized {

	return &UpdateAdminUserUnauthorized{}
}

// WriteResponse to the client
func (o *UpdateAdminUserUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// UpdateAdminUserForbiddenCode is the HTTP code returned for type UpdateAdminUserForbidden
const UpdateAdminUserForbiddenCode int = 403

/*UpdateAdminUserForbidden Not authorized to update an admin user

swagger:response updateAdminUserForbidden
*/
type UpdateAdminUserForbidden struct {
}

// NewUpdateAdminUserForbidden creates UpdateAdminUserForbidden with default headers values
func NewUpdateAdminUserForbidden() *UpdateAdminUserForbidden {

	return &UpdateAdminUserForbidden{}
}

// WriteResponse to the client
func (o *UpdateAdminUserForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// UpdateAdminUserInternalServerErrorCode is the HTTP code returned for type UpdateAdminUserInternalServerError
const UpdateAdminUserInternalServerErrorCode int = 500

/*UpdateAdminUserInternalServerError Server error

swagger:response updateAdminUserInternalServerError
*/
type UpdateAdminUserInternalServerError struct {
}

// NewUpdateAdminUserInternalServerError creates UpdateAdminUserInternalServerError with default headers values
func NewUpdateAdminUserInternalServerError() *UpdateAdminUserInternalServerError {

	return &UpdateAdminUserInternalServerError{}
}

// WriteResponse to the client
func (o *UpdateAdminUserInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}